/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run -n [name] -v [command]",
	Short: "Run a saved command",
	Long: `Run saved command, add a commnad, or delete a command For example:

ne run -n "git_undo" -v "git reset --soft HEAD~1"
ne run git_undo
ne run -d git_indo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("name", "n", "name", "memorable command name")
	rootCmd.PersistentFlags().StringP("value", "v", "value", "original command name")
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
