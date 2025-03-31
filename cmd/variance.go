package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var varCmd = &cobra.Command{
	Use:     "var",
	Aliases: []string{"variance"},
	Short:   "Variance of a group of values",
	Long:    `Calculates the variance of a group of values and prints ot the result`,
	Example: `zero var 20.5 30 50 70
zero variance --file data.json
zero var -- 10 5 -20 20`,
	Run: varRun,
}

func varRun(cmd *cobra.Command, args []string) {
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

		result = Variance(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Variance(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(varCmd)
	varCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
