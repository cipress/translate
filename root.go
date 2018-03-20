package translate

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	sl, tl string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&sl, "sl", "en", "source language")
	rootCmd.PersistentFlags().StringVar(&tl, "tl", "it", "target language")
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "a simple cli for translation from google.",
	RunE:  transl}

func transl(_ *cobra.Command, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("please provide at least 1 arg to translate\n")
	}
	var query string
	for _, q := range args {
		escaped := q + " "
		query += escaped
	}
	res, err := Translate(sl, tl, query)
	if err != nil {
		return fmt.Errorf("could not translate [%v]: %v", query, err)
	}
	fmt.Fprintln(os.Stdout, res)
	return nil
}

//Execute translate command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "could not execute command: %v", err)
		os.Exit(1)
	}
}