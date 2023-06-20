/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func doneRun(cmd *cobra.Command, args []string) {
	fmt.Println("done called")
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
