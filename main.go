package main

import (
	"fmt"
	"github.com/ckone4you/golangtest/cmd"
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My CLI App",
}

func init() {

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	profile, _ := rootCmd.Flags().GetString("awsIntegration-profile")
	region, _ := rootCmd.Flags().GetString("awsIntegration-region")
	rootCmd.AddCommand(greetCmd)
	// Add subcommands for each tool
	rootCmd.AddCommand(cmd.NewKubectlCommand())
	rootCmd.AddCommand(cmd.NewAWSCommand(profile, region))
	//rootCmd.AddCommand(cmd.NewHelmCommand())
	//rootCmd.AddCommand(cmd.NewBatsCommand())

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	logrus.Error(err)
	fmt.Println("An error occurred while executing the command. Please check the logs for more details.")
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Print a greeting message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s!\n", name)
	},
}
