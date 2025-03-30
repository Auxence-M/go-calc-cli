package cmd

import (
	"fmt"
	"os"
	"strconv"
)

func Add(args []float64) float64 {
	result := 0.0

	for _, i := range args {
		result = result + i
	}

	return result
}

func Subtract(args []float64) float64 {
	result := args[0]

	for i := 1; i < len(args); i++ {
		result = result - args[i]
	}
	return result
}

func Multiply(args []float64) float64 {
	result := 1.0

	for _, i := range args {
		result = result * i
	}
	return result
}

func Divide(args []float64) float64 {
	result := args[0]

	for i := 1; i < len(args); i++ {
		result = result / args[i]
	}
	return result
}

func getValues(args []string) ([]float64, error) {
	var values []float64
	for _, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while parsing the values '%s'\n", err)
			os.Exit(1)
		}
		values = append(values, val)
	}
	return values, nil
}
