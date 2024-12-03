package cmd

import (
	"aoc/cmd/day1"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "aoc",
}

func init() {
	Cmd.AddCommand(day1.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
