package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosflorencio/pr-description-ai/ai"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

const (
	OpenAPIEnvVar = "OPENAI_API_KEY"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pr-description-ai",
	Short: "Generate PR Description from git changes",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		value := os.Getenv(OpenAPIEnvVar)
		if value == "" {
			fmt.Printf("Error: Required environment variable %s is not set.\n", OpenAPIEnvVar)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		currentBranch, err := ai.CurrentBranchName()
		if err != nil {
			log.Fatal("Could not get current branch name, are you in a git repository?")
		}

		changes, err := ai.CompareGitChanges(targetBranch)

		if err != nil {
			log.Fatal("Couldn't get the git changes, is the target branch correct?")
		}

		if len(changes) == 0 {
			fmt.Println("No changes found.")
			return
		}

		prompt := fmt.Sprintf(`
		Generate a Pull Request description. Current branch name: %s
		The description should be in markdown format using the following template:

		## Motivation
		[insert here the possible motivation for the changes in a small paragraph]

		## Changes
		[insert here the description of the changes in a list]

		Here are the git changes:
		%s
		`, currentBranch, changes)

		response, err := ai.ChatGPT(prompt, gptModel)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(response)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var targetBranch string
var gptModel string

func init() {
	rootCmd.Flags().StringVarP(&targetBranch, "branch", "b", "develop", "Target branch to compare against")
	rootCmd.Flags().StringVarP(&gptModel, "model", "m", openai.GPT4, "OpenAI model to use")
}
