package stashify

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/pmyjavec/stashify/stashify/scm/stash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Project stash.StashProject

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

var projectNew = &cobra.Command{
	Use:   "new [PROJECT NAME]",
	Short: "Create a new Stash project",
	Long:  "Create a project to house repositories inside Atlassian Stash",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

var Config = viper.New()

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

	fmt.Println(Config.Get("url"))
	fmt.Println(Project.Name)

}

func rootRun(cmd *cobra.Command, args []string) {
	return
}

func addCommands() {
	rootCmd.AddCommand(project)

	/* Project Related Sub-Commands */
	project.AddCommand(projectNew)
}

func Execute() {
	addCommands()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
