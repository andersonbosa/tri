/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"tri/todo"

	"github.com/spf13/cobra"
)

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid position\n ", err)
	}

	if i > 0 && i < len(items) {

		/* TODO: move to isolated function */
		if items[i-1].Done == true {
			items[i-1].Done = false
			fmt.Printf("%q %v\n", items[i-1].Text, "marked as TODO")

		} else {
			items[i-1].Done = true
			fmt.Printf("%q %v\n", items[i-1].Text, "marked as DONE")
		}

		sort.Sort(todo.ByPri(items))

		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doens't match any items")
	}

}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark item as Done",
	Aliases: []string{"do"},
	// Long: ``,
	Run: doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
