package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage todos",
	Long:  `Add, list, or delete todos.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello! you entered in Todo Application 📝.")
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
