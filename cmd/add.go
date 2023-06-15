/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
	for _, value := range args {
		fmt.Println(value)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	/*
	* special function
	* called after package variable declarations
	* each package may have multiple init
	* init() order un-guaranteed
	 */
}
