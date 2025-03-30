package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var file string

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Adds multiple numbers",
	Long: `Carry out an addition operation on the numbers provided. 
Adds all the numbers together and prints out the result. Only numerical values are allowed.`,
	Example: `zero add 20.5 30 50 70
zero add --file data.json
zero add -- 10 5 -20 20`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	result := 0.0

	if file != "" && len(args) > 0 {
		fmt.Println("You cannot provide a file and values at the same time.")
		os.Exit(1)
	}

	if file != "" {
		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while  while reading the file '%s'\n", err)
			os.Exit(1)
		}

		var values []float64
		err = json.Unmarshal(b, &values)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while  while reading the file '%s'\n", err)
			os.Exit(1)
		}

		result = Add(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Add(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
