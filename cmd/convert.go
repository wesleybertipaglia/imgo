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
	Short: "Convert a single image or all images of one type in a directory",
	Example: `  
		imgo convert -i input.jpg --to png -o ./output
		imgo convert -i ./images --from png --to webp
		imgo convert -i ./images --ti png --to webp -o ./output -r 80`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		from, _ := cmd.Flags().GetString("from")
		ti, _ := cmd.Flags().GetString("ti")
		to, _ := cmd.Flags().GetString("to")
		typeAlias, _ := cmd.Flags().GetString("type")
		outputDir, _ := cmd.Flags().GetString("output")
		rate, _ := cmd.Flags().GetInt("rate")

		if from == "" {
			from = ti
		}
		if to == "" {
			to = typeAlias
		}

		if input == "" || to == "" {
			fmt.Println("Usage:")
			fmt.Println("  single image: imgo convert -i <file> --to <png|jpg|webp> [-o ./output]")
			fmt.Println("  directory:    imgo convert -i <directory> --from <png|jpg|...> --to <png|jpg|webp>")
			return
		}

		toExt := utils.NormalizeExt(to)
		encoder, err := utils.GetEncoder(strings.TrimPrefix(toExt, "."))
		if err != nil {
			fmt.Println(err)
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
		if info.IsDir() {
			if from == "" {
				fmt.Println("For a directory, --from is required. Example: imgo convert -i ./images --from png --to webp")
				return
			}
			fromExt := utils.NormalizeExt(from)
			if !utils.IsSupportedInputExt(fromExt) {
				fmt.Printf("Unsupported input type: %s\n", from)
				return
			}
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
			inputs = []string{input}
			if outputDir == "" {
				outputDir = "."
			}
		}

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Println("Failed to create output directory:", err)
			return
		}

		toName := strings.TrimPrefix(toExt, ".")
		var success, failed int
		for _, in := range inputs {
			img, err := imaging.Open(in)
			if err != nil {
				fmt.Printf("Skipping %s: failed to open image: %v\n", in, err)
				failed++
				continue
			}

			baseName := strings.TrimSuffix(filepath.Base(in), filepath.Ext(in))
			outputPath := filepath.Join(outputDir, baseName+"_converted."+strings.ToLower(toName))

			f, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("Skipping %s: failed to create output file: %v\n", in, err)
				failed++
				continue
			}

			if err := encoder.Encode(f, img, rate); err != nil {
				fmt.Printf("Skipping %s: failed to encode image: %v\n", in, err)
				f.Close()
				failed++
				continue
			}
			f.Close()

			fmt.Println("Image converted and saved to", outputPath)
			success++
		}

		if len(inputs) > 1 {
			fmt.Printf("Done: %d converted, %d failed.\n", success, failed)
		}
	},
}

func init() {
	convertCmd.Flags().StringP("input", "i", "", "Input file or directory (required)")
	convertCmd.Flags().String("from", "", "Input image type, e.g. png (required for directory, ignored for single file)")
	convertCmd.Flags().String("ti", "", "Alias for --from, e.g. --ti png")
	convertCmd.Flags().String("to", "", "Output image type, e.g. webp (required)")
	convertCmd.Flags().StringP("type", "t", "", "Alias for --to, e.g. -t webp")
	convertCmd.Flags().StringP("output", "o", "", "Output directory (default: current dir for single file, same as input for directory)")
	convertCmd.Flags().IntP("rate", "r", 80, "Compression rate (default: 80)")
}
