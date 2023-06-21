/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* Flags pointers */
var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add wil create a new todo item to the list`,
	Run:   addRun,
}

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
func addRun(cmd *cobra.Command, args []string) {
	items := []todo.Item{}

	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, value := range args {
		newItem := todo.Item{Text: value}
		newItem.SetPriority(priority)

		items = append(items, newItem)
	}

	if todo.SaveItems(viper.GetString("datafile"), items) != nil {
		_ = fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(
		&priority, "priority", "p", 2, "Priority:1,2,3",
	)

	/*
		- special function
		- called after package variable declarations
		- each package may have multiple init
		- init() order un-guaranteed
	*/
}
