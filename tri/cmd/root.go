/*
Copyright © 2023 Anderson Bosa
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long: `Tri will help you get more done in less time.
It's designed to be as simple as possible to help
you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	/* `:=` assignment operator; declares & assigns in one operation */
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		/* err are just values. No exceptions in Go. Errors should be handled when the occur */
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --data-file")
	}

	rootCmd.PersistentFlags().StringVar(
		&dataFile,
		"datafile",
		home+string(os.PathSeparator)+".tri_todo.json",
		"data file to store todos",
	)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "", "config file (default is $HOME/.tri.yml)",
	)
}

// This function initializes the configuration settings for a Go program using Viper.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tri")

	/* Read config file if is found */
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration from file: ", viper.ConfigFileUsed())
	}
}
