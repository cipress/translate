package cmd

import (
	"log"
	"testing"
)

func TestNoArgsFail(t *testing.T) {
	cmd := rootCmd
	if err := cmd.Execute(); err == nil {
		log.Fatalf("an error shoul occur.")
	}
}
