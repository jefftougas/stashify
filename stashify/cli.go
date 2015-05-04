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
var pullRequestCreateTitle string
var pullRequestCreateDescription string

var rootCmd = &cobra.Command{
	Use:   "stashify",
	Short: "Atlassian Stash CLI in Go",
	Long:  "Stashify helps automate git workflows from the command line",
	Run:   rootRun,
}

var project = &cobra.Command{
	Use:   "project [command]",
	Short: "Manage projects which house repositories",
	Long:  "Project related management tasks, listing of projects, creation, etc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var pullRequest = &cobra.Command{
	Use:   "pr [command]",
	Short: "Work with pull requests",
	Long:  "Create pull requests, merge, decline etc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var pullRequestCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a new pull request",
	Long:  "Create a new pull request on the current branch by default",
	Run: func(cmd *cobra.Command, args []string) {
		pr := stash.StashPullRequest{Project: Project}
		pr.Create(pullRequestCreateTitle, pullRequestCreateDescription)
	},
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
	rootCmd.AddCommand(pullRequest)
	rootCmd.AddCommand(project)
	project.AddCommand(projectCreate)
	pullRequest.AddCommand(pullRequestCreate)

	rootCmd.PersistentFlags().StringVarP(&Project.Username, "username", "u", "", "Username for Stash")
	rootCmd.PersistentFlags().StringVarP(&Project.Password, "password", "p", "", "Password for Stash")

	project.PersistentFlags().StringVarP(&projectName, "name", "n", "", "New project name, defaults to .stashify.yml project name")
	project.PersistentFlags().StringVarP(&projectKey, "key", "k", "", "New project key, defaults to .stashify.yml project key")

	/* Pull request related sub commands */
	pullRequestCreate.PersistentFlags().StringVarP(&pullRequestCreateTitle, "title", "t", "", "Title for pull request, defaults to commit message title")
	pullRequestCreate.PersistentFlags().StringVarP(&pullRequestCreateDescription, "description", "d", "", "Pull request description, defaults to commit message")
}

func Execute() {
	addCommands()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
