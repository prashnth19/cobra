package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, ORT Recovery!")
		},
	}

	rootCmd.Execute()
}
