package cmd

import (
	"fmt"
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
			fmt.Println("Please provide input file and output type.")
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

		baseName := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
		outputPath := filepath.Join(outputDir, baseName+"."+outputType)

		switch strings.ToLower(outputType) {
		case "jpg", "jpeg":
			err = imaging.Save(img, outputPath, imaging.JPEGQuality(95))
		case "png":
			err = imaging.Save(img, outputPath)
		default:
			fmt.Println("Unsupported format:", outputType)
			return
		}

		if err != nil {
			fmt.Println("Failed to save image:", err)
		} else {
			fmt.Println("Image converted and saved to", outputPath)
		}
	},
}

func init() {
	convertCmd.Flags().StringP("input", "i", "", "Input image path")
	convertCmd.Flags().StringP("type", "t", "", "Output format (png, jpg, webp)")
	convertCmd.Flags().StringP("output", "o", "", "Output directory (default: current)")
}
