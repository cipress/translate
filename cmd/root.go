package cmd

import (
	"fmt"
	"github.com/hankmartinez/translate"
	"github.com/spf13/cobra"
	"net/url"
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
	var escapedQuery string
	for _, q := range args {
		escaped := url.PathEscape(q + " ")
		escapedQuery += escaped
	}
	res, err := translate.Translate(sl, tl, escapedQuery)
	if err != nil {
		return fmt.Errorf("could not translate [%v]: %v", escapedQuery, err)
	}
	fmt.Printf("%v\n", res)
	return nil
}

//Execute translate command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
