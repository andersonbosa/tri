/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"tri/todo"

	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

func parsePrettyLine(item todo.Item) string {
	// return strconv.Itoa(item.Priority) + "\t" + item.Text + "\t"
	delimiter := "\t"

	return item.PrettyPosition() +
		delimiter +
		item.PrettyDone() +
		delimiter +
		item.PrettyPriotity() +
		delimiter +
		item.Text +
		delimiter
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, item := range items {
		var prettyLine string

		if allOpt || item.Done == doneOpt {
			prettyLine = parsePrettyLine(item)
		}

		fmt.Fprintln(w, prettyLine)
	}

	w.Flush()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `Listing the todos`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todos")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
