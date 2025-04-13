package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress JPEG images",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		rate, _ := cmd.Flags().GetInt("rate")

		if input == "" {
			fmt.Println("Please provide an input image.")
			return
		}

		if rate <= 0 || rate > 100 {
			rate = 80
		}

		img, err := imaging.Open(input)
		if err != nil {
			fmt.Println("Failed to open image:", err)
			return
		}

		outputPath := strings.TrimSuffix(input, filepath.Ext(input)) + "_compressed.jpg"

		err = imaging.Save(img, outputPath, imaging.JPEGQuality(rate))
		if err != nil {
			fmt.Println("Failed to save compressed image:", err)
		} else {
			fmt.Println("Compressed image saved to", outputPath)
		}
	},
}

func init() {
	compressCmd.Flags().StringP("input", "i", "", "Input image path")
	compressCmd.Flags().IntP("rate", "r", 80, "Compression rate (default: 80)")
}
