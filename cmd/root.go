package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/timwehrle/gofee/pkg/gofee"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	// The default length of the password
	defaultLength int = 20

	MAJOR = 0
	MINOR = 1
	PATCH = 0
)

// Options for the generate command
var options struct {
	length       int
	lowers       bool
	uppers       bool
	digits       bool
	symbols      bool
	passwordType string
	quiet        bool
}

func init() {
	rootCmd.Flags().BoolVarP(&options.lowers, "exclude-lowers", "w", false, "exclude lowercase letters")
	rootCmd.Flags().BoolVarP(&options.uppers, "exclude-uppers", "u", false, "exclude uppercase letters")
	rootCmd.Flags().BoolVarP(&options.digits, "exclude-digits", "d", false, "exclude digits")
	rootCmd.Flags().BoolVarP(&options.symbols, "exclude-symbols", "s", false, "exclude symbols")
	rootCmd.Flags().IntVarP(&options.length, "length", "l", defaultLength, "length of the password")
	rootCmd.Flags().StringVarP(&options.passwordType, "type", "t", "", "type of password to generate (pin, memorable)")
	rootCmd.Flags().BoolVarP(&options.quiet, "quiet", "q", false, "print only the password to stdout (no extra info)")

	// Colorize the usage output
	rootCmd.SetOutput(color.Output)
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Examples:`, `{{StyleHeading "Examples:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTemplate)
}

var example = `
gofee --length 20 --exclude-lowers
gofee --length 12 -u -d 
gofee --type pin --length 4
`

var long = `
Gofee is a simple password generator that uses crypto/rand.
It generates a password of a given length using a configurable character set.
`

var rootCmd = &cobra.Command{
	Use:     "gofee",
	Version: fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH),
	Example: example,
	Short:   "Gofee is a simple password generator, reliable and secure.",
	Long:    long,
	Run: func(cmd *cobra.Command, args []string) {
		config := gofee.PasswordConfig{
			IncludeLowers:  !options.lowers,
			IncludeUppers:  !options.uppers,
			IncludeDigits:  !options.digits,
			IncludeSymbols: !options.symbols,
			Type:           options.passwordType,
			RequireClasses: true,
			MinPINLength:   6,
			MinLength:      8,
			MaxLength:      4096,
		}

		pw, charsetUsed, err := gofee.Generate(options.length, config)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		size := len(charsetUsed)
		entropy, err := gofee.CalculateEntropy(size, options.length)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error calculating entropy:", err)
			os.Exit(1)
		}

		if !options.quiet {
			fmt.Fprint(os.Stderr, "Entropy: ")
			color.New(color.FgGreen).Fprintf(os.Stderr, "%.2f bits\n", entropy)
		}

		fmt.Fprintln(os.Stdout, pw)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
