package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type TodoItem struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todoList []TodoItem

// File path for storing todos
const todoFilePath = "todos.json"

// Function to save todos to a JSON file
func saveTodosToFile(todos []TodoItem) error {
	// Marshal todos slice to JSON
	data, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		return err
	}

	// Write JSON data to file
	err = ioutil.WriteFile(todoFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to load todos from a JSON file
func loadTodosFromFile() ([]TodoItem, error) {
	var todos []TodoItem

	// Read the JSON file
	file, err := ioutil.ReadFile(todoFilePath)
	if err != nil {
		// Return empty slice if file doesn't exist
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err
	}

	// Unmarshal JSON data into todos slice
	err = json.Unmarshal(file, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

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

		// Save todos to file
		err := saveTodosToFile(todoList)
		if err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}

		fmt.Printf("Todo created with ID: %d, Task: %s\n", newTodo.ID, newTodo.Task)
	} else {
		fmt.Println("Error reading input:", scanner.Err())
	}
}

// fucntion for listing out all the todos
func listTodos() {

	// Load todos from file
	todos, err := loadTodosFromFile()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	fmt.Println("Todo List 📄:")
	if len(todos) == 0 {
		fmt.Println("Currently no todos.Please make one.")
		return
	}

	for _, todo := range todos {
		fmt.Printf("ID : %d | Task : %s\n", todo.ID, todo.Task)
	}
}

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.AddCommand(createCmd, listCmd)
}
