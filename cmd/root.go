package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for performing mathematical operations",
	Long: `zero is a cli tool for performing mathematical operations like addition, multiplication, division and subtraction on a list of values
provided by the user either directly after the command, or written in a json file. 
It can also perform statistical operations like mean, standard deviations etc.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing subcommands '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
