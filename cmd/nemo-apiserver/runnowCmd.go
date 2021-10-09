package main

import (
	_ "embed"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var runnowCmd = &cobra.Command{
	Use:   "runnow",
	Short: "runnow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("12345")
		fmt.Printf("%s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}
