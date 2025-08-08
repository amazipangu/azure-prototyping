package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toyokumono",
	Short: "This CLI tool helps users avoid repeatedly executing the same command by organizing setup processes, thus speeding up the local environment configuration for cloud-native development.",
	Long: `This CLI tool is designed to improve the efficiency of setting up cloud-native development environments on local machines.
			It organizes repetitive setup processes, eliminating the need for users to repeatedly execute the same commands.
			By automating and streamlining the configuration steps, the tool accelerates the local environment setup, 
			allowing developers to focus on coding and development rather than manual configurations. 
			It supports a faster, more efficient workflow for building and deploying cloud-native applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff here.
		fmt.Println("Toyokumo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
