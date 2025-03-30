package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var subtractCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"subtract"},
	Short:   "Subtract multiple numbers",
	Long: `Carry out a subtraction operation on the numbers provided. 
Subtract all the numbers together and prints out the result. Only numerical values are allowed.`,
	Example: `zero sub 20.5 30 50 70
zero subtract --file data.json
zero sub -- 10 5 -20 20`,
	Run: subRun,
}

func subRun(cmd *cobra.Command, args []string) {
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

		result = Subtract(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Subtract(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(subtractCmd)
	subtractCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
