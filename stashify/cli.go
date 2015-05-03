package stashify

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/pmyjavec/stashify/stashify/scm/stash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Config = viper.New()
var Project stash.StashProject

/* Flag / options values */
var projectName string
var projectKey string

var rootCmd = &cobra.Command{
	Use:   "stashify",
	Short: "Atlassian Stash CLI in Go",
	Long:  "Stashify automates git workflows from the command line",
	Run:   rootRun,
}

var project = &cobra.Command{
	Use:   "project [command]",
	Short: "Manage projects which house repositories",
	Long:  "Project related management tasks, listing of projects, creation, etc",
}

var projectCreate = &cobra.Command{
	Use:   "create [PROJECT NAME]",
	Short: "Create a new Stash project",
	Long:  "Create a project to house repositories inside Atlassian Stash",
	Run: func(cmd *cobra.Command, args []string) {
		Project.Create(projectName, projectKey)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.Debug("Loading configuration files from ~/.stashify or .stashify.yml")
	Config.SetConfigName(".stashify")
	Config.AddConfigPath("$HOME/")
	Config.AddConfigPath("./")
	Config.ReadInConfig()

	if err := Config.Marshal(&Project); err != nil {
		log.Error(err)
	}

	log.Debug("Loaded Stashify project from", Config.ConfigFileUsed())
	log.Debug(fmt.Sprintf("Stashify project looks like: %+v", Project))
}

func rootRun(cmd *cobra.Command, args []string) {
	return
}

func addCommands() {
	rootCmd.AddCommand(project)
	/* Global Options */
	rootCmd.PersistentFlags().StringVarP(&Project.Username, "username", "u", "", "Username for Stash")
	rootCmd.PersistentFlags().StringVarP(&Project.Password, "password", "p", "", "Password for Stash")

	/* Project Options */
	project.PersistentFlags().StringVarP(&projectName, "name", "n", "", "New project name, defaults to .stashify.yml project name")
	project.PersistentFlags().StringVarP(&projectKey, "key", "k", "", "New project key, defaults to .stashify.yml project key")

	/* Project Related Sub-Commands */
	project.AddCommand(projectCreate)
}

func Execute() {
	addCommands()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
