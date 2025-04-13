package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgtool",
	Short: "A CLI tool to convert and compress images",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(compressCmd)
}
