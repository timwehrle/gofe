package cmd

import (
	"fmt"
	"gofee/pkg/gofee"
	"log"

	"github.com/spf13/cobra"
)

// The default length of the password
const (
	defaultLength int = 16
)

// Options for the generate command
var options struct {
	length  int
	lowers  bool
	uppers  bool
	digits  bool
	symbols bool
}

func init() {
	genCmd.Flags().BoolVarP(&options.lowers, "exclude-lowers", "w", false, "Exclude lowercase letters")
	genCmd.Flags().BoolVarP(&options.uppers, "exclude-uppers", "u", false, "Exclude uppercase letters")
	genCmd.Flags().BoolVarP(&options.digits, "exclude-digits", "d", false, "Exclude digits")
	genCmd.Flags().BoolVarP(&options.symbols, "exclude-symbols", "s", false, "Exclude symbols")
	genCmd.Flags().IntVarP(&options.length, "length", "l", defaultLength, "Length of the password")
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"generate"},
	Short:   "Generate a password",
	Run: func(cmd *cobra.Command, args []string) {

		config := gofee.PasswordConfig{
			IncludeLowers:  !options.lowers,
			IncludeUppers:  !options.uppers,
			IncludeDigits:  !options.digits,
			IncludeSymbols: !options.symbols,
		}

		pw, err := gofee.Generate(options.length, config)
		if err != nil {
			log.Fatalf("Error generating password: %v", err)
		}

		entropy, err := gofee.CalculateEntropy(len(gofee.Charset), options.length)
		if err != nil {
			log.Fatalf("Error calculating entropy: %v", err)
		}

		fmt.Printf("Entropy: %.2f bits\n", entropy)
		fmt.Println("Password:", pw)
	},
}
