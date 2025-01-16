package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a single todo item
type Task struct {
	description string
	done        bool
}

// TodoList manages the collection of tasks
type TodoList struct {
	tasks []Task
}

// addTask adds a new task to the list
func (t *TodoList) addTask(description string) {
	task := Task{
		description: description,
		done:        false,
	}
	t.tasks = append(t.tasks, task)
	fmt.Println("Task added successfully!")
}

// viewTasks displays all tasks with their status
func (t *TodoList) viewTasks() {
	if len(t.tasks) == 0 {
		fmt.Println("No tasks in the list!")
		return
	}

	fmt.Println("\nCurrent Tasks:")
	fmt.Println("-------------")
	for i, task := range t.tasks {
		status := " "
		if task.done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.description)
	}
	fmt.Println()
}

// deleteTask removes a task by its index
func (t *TodoList) deleteTask(index int) error {
	if index < 0 || index >= len(t.tasks) {
		return fmt.Errorf("invalid task index")
	}

	// Remove task by creating a new slice without the selected index
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Task deleted successfully!")
	return nil
}

// toggleTask marks a task as done or undone
func (t *TodoList) toggleTask(index int) error {
	if index < 0 || index >= len(t.tasks) {
		return fmt.Errorf("invalid task index")
	}

	t.tasks[index].done = !t.tasks[index].done
	status := "undone"
	if t.tasks[index].done {
		status = "done"
	}
	fmt.Printf("Task marked as %s!\n", status)
	return nil
}

// showMenu displays the available commands
func showMenu() {
	fmt.Println("\nTodo List Menu:")
	fmt.Println("1. Add task")
	fmt.Println("2. View tasks")
	fmt.Println("3. Delete task")
	fmt.Println("4. Toggle task status")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice (1-5): ")
}

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()

		// Read user choice
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			todoList.addTask(description)

		case "2":
			todoList.viewTasks()

		case "3":
			todoList.viewTasks()
			fmt.Print("Enter task number to delete: ")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid input! Please enter a number.")
				continue
			}
			if err := todoList.deleteTask(index - 1); err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			todoList.viewTasks()
			fmt.Print("Enter task number to toggle: ")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid input! Please enter a number.")
				continue
			}
			if err := todoList.toggleTask(index - 1); err != nil {
				fmt.Println("Error:", err)
			}

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}