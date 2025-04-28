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

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert images to different formats",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		outputType, _ := cmd.Flags().GetString("type")
		outputDir, _ := cmd.Flags().GetString("output")

		if input == "" || outputType == "" {
			fmt.Println("Please provide both --input and --type options.")
			return
		}

		if outputDir == "" {
			outputDir = "."
		}

		img, err := imaging.Open(input)
		if err != nil {
			fmt.Println("Failed to open image:", err)
			return
		}

		encoder, err := utils.GetEncoder(outputType)
		if err != nil {
			fmt.Println(err)
			return
		}

		baseName := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
		outputPath := filepath.Join(outputDir, baseName+"_converted."+strings.ToLower(outputType))

		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			return
		}
		defer f.Close()

		if err := encoder.Encode(f, img); err != nil {
			fmt.Println("Failed to encode image:", err)
			return
		}

		fmt.Println("Image converted and saved to", outputPath)
	},
}

func init() {
	convertCmd.Flags().StringP("input", "i", "", "Input image path")
	convertCmd.Flags().StringP("type", "t", "", "Output format (png, jpg, webp)")
	convertCmd.Flags().StringP("output", "o", "", "Output directory (default: current)")
}
