package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var multiplyCmd = &cobra.Command{
	Use:     "mul",
	Aliases: []string{"multiply"},
	Short:   "Multiplies multiple numbers",
	Long: `Carry out a multiplication operation on the numbers provided. 
Multiplies all the numbers together and prints out the result. Only numerical values are allowed.`,
	Example: `zero mul 20.5 30 50 70
zero multiply --file data.json
zero mul -- 10 5 -20 20`,
	Run: multiplyRun,
}

func multiplyRun(cmd *cobra.Command, args []string) {
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

		result = Multiply(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Multiply(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
	multiplyCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
