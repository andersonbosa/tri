/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"tri/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add wil create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	/*
		- Go does not have constructors
		- if construction is required (initialization prior to use), use a factory
		- Convention is to use New____()
		- or New() WHEN ONLY one type is exported in the package
	*/

	/*
		Composite Literals:
			- []todo.Item{} - an expressions that creats a new value each time it is evaluated
	*/
	items := []todo.Item{}

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, value := range args {
		items = append(items, todo.Item{Text: value})
	}

	if todo.SaveItems(dataFile, items) != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	/*
		- special function
		- called after package variable declarations
		- each package may have multiple init
		- init() order un-guaranteed
	*/
}