package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var count int

var fooCmd = &cobra.Command{
	Use:   "foo",
	Short: "The foo command",
	RunE:  runFoo,
}

func runFoo(cmd *cobra.Command, args []string) (err error) {
	fmt.Printf("This is the foo command; count=%d, threads=%d\n", count, threads)
	return nil
}

func init() {
	fooCmd.Flags().IntVarP(&count, "count", "c", 0, "Count of foo")
	rootCmd.AddCommand(fooCmd)
}
