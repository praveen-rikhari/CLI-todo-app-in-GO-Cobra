package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo_app",
	Short: "A basic CLI todo application",
	Long:  "A basic ClI todo list application performing creation, listing & deletion of todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the root command of your todo_app CLI.")
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
