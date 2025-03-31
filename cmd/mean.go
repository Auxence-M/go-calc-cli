package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var meanCmd = &cobra.Command{
	Use:   "mean",
	Short: "Mean of a group of values",
	Long:  `Calculates the mean or average of a group of values and prints ot the result`,
	Example: `zero mean 20.5 30 50 70
zero mean --file data.json
zero mean -- 10 5 -20 20`,
	Run: meanRun,
}

func meanRun(cmd *cobra.Command, args []string) {
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

		result = Mean(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Mean(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(meanCmd)
	meanCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
