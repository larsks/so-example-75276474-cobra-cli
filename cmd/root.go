package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var threads int

var rootCmd = &cobra.Command{
	Use:   "clitest",
	Short: "A simple cli test",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	fmt.Printf("This is the root command, threads=%d\n", threads)
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")
}
