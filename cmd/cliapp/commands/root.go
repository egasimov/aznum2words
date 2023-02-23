/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"github.com/egasimov/aznum2words"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any cliapp
var rootCmd = &cobra.Command{
	Use:   "aznum2words-cli [arg]",
	Short: "A command line tool used for converting numbers into words in Azerbaijani language",
	Long: `
aznum2words-cli is a CLI tool that is designed to convert positive or negative numbers - integers, floating-point into word representation in the Azerbaijani language.

Use cases:
There are many use cases where the word representation of numbers is useful or even necessary in real life. 

Here are a few examples:
- Legal and financial documents: Legal and financial documents often require the use of the word representation of numbers to avoid confusion and ensure accuracy.
- Check writing: When writing a check, it is common practice to write out the word representation of the amount in addition to the numerical amount to prevent alterations or fraud.
- Customer service: Customer service representatives may need to verbally communicate the word representation of a number to a customer over the phone, particularly when dealing with complex or technical issues.
- Education: In education, the word representation of numbers is often used to teach children how to spell numbers and reinforce their understanding of place value.

`,
	Example: `
	aznum2words-cli "12345"
	aznum2words-cli "-12345"
	aznum2words-cli "123.45"
	aznum2words-cli "-123.45"
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		arg0 := strings.TrimSpace(args[0])

		result, err := aznum2words.SpellNumber(arg0)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, result)
	},
}

// Execute adds all child command to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aznum2words.exe.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
