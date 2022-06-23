package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dump-dumper",
	Short: "dump-dumper is a simple cli tool to streamline your dockerize database backup process",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to dump-dumper")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
