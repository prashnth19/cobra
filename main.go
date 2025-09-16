package main

import (
	"fmt"
	"log"
	"os"

	// External imports
	"github.com/spf13/cobra"         // single-line import
	"github.com/google/uuid"         // another external dep
	"golang.org/x/text/cases"        // golang.org package
	"golang.org/x/text/language"     // paired dependency

	// Aliased import
	uuidv4 "github.com/google/uuid"

	// Blank import (side-effect only)
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Use cobra
	var rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, ORT Recovery!")
		},
	}

	// Use uuid
	id := uuid.New()
	fmt.Println("Generated UUID:", id)

	// Use aliased uuid
	id2 := uuidv4.New()
	log.Println("Generated UUID (aliased):", id2)

	// Use text casing
	caser := cases.Title(language.English)
	fmt.Println(caser.String("ort recovery test"))

	// Just to keep os in use
	fmt.Println("Args:", os.Args)

	rootCmd.Execute()
}
