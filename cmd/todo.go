package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type TodoItem struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todoList []TodoItem

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage todos",
	Long:  `Add, list, or delete todos.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello! you entered in Todo Application 📝.")
		cmd.Help()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Todo",
	Run: func(cmd *cobra.Command, args []string) {
		createTodo()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing all the todos",
	Run: func(cmd *cobra.Command, args []string) {
		listTodos()
	},
}

// function for creating todos
func createTodo() {
	// Prompt the user to enter the task for the todo item
	fmt.Print("Enter your task : ")

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Scan the input until newline character
	if scanner.Scan() {
		task := scanner.Text()

		newTodo := TodoItem{
			ID:   len(todoList) + 1,
			Task: task,
		}
		todoList = append(todoList, newTodo)

		fmt.Printf("Todo created with ID: %d, Task: %s\n", newTodo.ID, newTodo.Task)
	} else {
		fmt.Println("Error reading input:", scanner.Err())
	}
}

// fucntion for listing out all the todos
func listTodos() {
	fmt.Println("Todo List 📄:")
	if len(todoList) == 0 {
		fmt.Println("Currently no todos.Please make one.")
		return
	}
	for _, todo := range todoList {
		fmt.Printf("ID : %d | Task : %s\n", todo.ID, todo.Task)
	}
}

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.AddCommand(createCmd, listCmd)
}
