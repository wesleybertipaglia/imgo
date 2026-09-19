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
	Short: "Compress a single image or all images of one type in a directory",
	Example: `  imgo compress -i input.jpg -r 80
  imgo compress -i ./images --from jpg -r 80
  imgo compress -i ./images --from png -r 80 -o ./output`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		from, _ := cmd.Flags().GetString("from")
		ti, _ := cmd.Flags().GetString("ti")
		rate, _ := cmd.Flags().GetInt("rate")
		format, _ := cmd.Flags().GetString("format")
		outputDir, _ := cmd.Flags().GetString("output")

		if from == "" {
			from = ti
		}

		if input == "" {
			fmt.Println("Usage:")
			fmt.Println("  single image: imgo compress -i <file> [-r 80]")
			fmt.Println("  directory:    imgo compress -i <directory> --from <png|jpg|...> [-r 80] [-o ./output]")
			return
		}

		if rate <= 0 || rate > 100 {
			rate = 80
		}

		info, err := os.Stat(input)
		if err != nil {
			fmt.Printf("Input not found: %s\n", input)
			return
		}

		var inputs []string
		var outFormat string
		if info.IsDir() {
			if from == "" {
				fmt.Println("For a directory, --from is required. Example: imgo compress -i ./images --from jpg -r 80")
				return
			}
			fromExt := utils.NormalizeExt(from)
			if !utils.IsSupportedInputExt(fromExt) {
				fmt.Printf("Unsupported input type: %s\n", from)
				return
			}
			outFormat = format
			if outFormat == "" {
				outFormat = strings.TrimPrefix(fromExt, ".")
			}
			outFormat = strings.ToLower(outFormat)

			inputs, err = utils.CollectImagesByType(input, fromExt)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(inputs) == 0 {
				fmt.Printf("No %s images found in %s\n", fromExt, input)
				return
			}
			if outputDir == "" {
				outputDir = input
			}
		} else {
			if !utils.IsSupportedImage(input) {
				fmt.Printf("Unsupported input file: %s\n", input)
				return
			}
			outFormat = format
			if outFormat == "" {
				outFormat = strings.TrimPrefix(strings.ToLower(filepath.Ext(input)), ".")
			}
			outFormat = strings.ToLower(outFormat)
			inputs = []string{input}
			if outputDir == "" {
				outputDir = filepath.Dir(input)
			}
		}

		encoder, err := utils.GetEncoder(outFormat)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Println("Failed to create output directory:", err)
			return
		}

		var success, failed int
		for _, in := range inputs {
			img, err := imaging.Open(in)
			if err != nil {
				fmt.Printf("Skipping %s: failed to open image: %v\n", in, err)
				failed++
				continue
			}

			baseName := strings.TrimSuffix(filepath.Base(in), filepath.Ext(in))
			outputPath := filepath.Join(outputDir, baseName+"_compressed."+outFormat)

			f, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("Skipping %s: failed to create output file: %v\n", in, err)
				failed++
				continue
			}

			err = encoder.Encode(f, img, rate)
			f.Close()
			if err != nil {
				fmt.Printf("Skipping %s: failed to encode image: %v\n", in, err)
				failed++
			} else {
				fmt.Println("Compressed image saved to", outputPath)
				success++
			}
		}

		if len(inputs) > 1 {
			fmt.Printf("Done: %d compressed, %d failed.\n", success, failed)
		}
	},
}

func init() {
	compressCmd.Flags().StringP("input", "i", "", "Input file or directory (required)")
	compressCmd.Flags().String("from", "", "Input image type, e.g. jpg (required for directory, ignored for single file)")
	compressCmd.Flags().String("ti", "", "Alias for --from, e.g. --ti jpg")
	compressCmd.Flags().IntP("rate", "r", 80, "Compression rate (default: 80)")
	compressCmd.Flags().StringP("format", "f", "", "Output format (default: same as input)")
	compressCmd.Flags().StringP("output", "o", "", "Output directory (default: same as input)")
}
