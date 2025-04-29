package cmd

import (
	"fmt"
	"imgo/cmd/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress images to JPEG, PNG, or WebP",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		rate, _ := cmd.Flags().GetInt("rate")
		format, _ := cmd.Flags().GetString("format")

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

		if format == "" {
			format = strings.TrimPrefix(filepath.Ext(input), ".")
		}

		encoder, err := utils.GetEncoder(strings.ToLower(format))
		if err != nil {
			fmt.Println(err)
			return
		}

		baseName := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
		outputPath := filepath.Join(filepath.Dir(input), baseName+"_compressed."+format)

		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			return
		}
		defer f.Close()

		err = encoder.Encode(f, img, rate)
		if err != nil {
			fmt.Println("Failed to encode image:", err)
		} else {
			fmt.Println("Compressed image saved to", outputPath)
		}
	},
}

func init() {
	compressCmd.Flags().StringP("input", "i", "", "Input image path")
	compressCmd.Flags().IntP("rate", "r", 80, "Compression rate (default: 80)")
	compressCmd.Flags().StringP("format", "f", "", "Output format (default: input format)")
}
