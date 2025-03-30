package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var divideCmd = &cobra.Command{
	Use:     "div",
	Aliases: []string{"divide"},
	Short:   "Divides multiple numbers",
	Long: `Carry out a division operation on the numbers provided. 
Divides all the numbers together and prints out the result. Only numerical values are allowed.`,
	Example: `zero div 20.5 30 50 70
zero divide --file data.json
zero div -- 10 5 -20 20`,
	Run: divideRun,
}

func divideRun(cmd *cobra.Command, args []string) {
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

		result = Divide(values)
	} else {
		values, err := getValues(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing the add command '%s'\n", err)
			os.Exit(1)
		}

		result = Divide(values)
	}

	fmt.Println("result: ", result)
}

func init() {
	rootCmd.AddCommand(divideCmd)
	divideCmd.Flags().StringVar(&file, "file", "", "File containing numbers/data")
}
