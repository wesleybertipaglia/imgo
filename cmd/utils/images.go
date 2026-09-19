package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var supportedInputExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tif":  true,
	".tiff": true,
	".bmp":  true,
}

func IsSupportedImage(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return supportedInputExtensions[ext]
}

// NormalizeExt accepts "png", ".png", "PNG" and returns ".png".
func NormalizeExt(ext string) string {
	ext = strings.TrimSpace(strings.ToLower(ext))
	if ext == "" {
		return ""
	}
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return ext
}

func IsSupportedInputExt(ext string) bool {
	return supportedInputExtensions[NormalizeExt(ext)]
}

// CollectImagesByType lists only files with the given extension directly
// inside dir (non-recursive). dir must be a directory.
func CollectImagesByType(dir string, ext string) ([]string, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("input is not a valid directory: %s", dir)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("input must be a directory, got file: %s", dir)
	}

	want := NormalizeExt(ext)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		full := filepath.Join(dir, entry.Name())
		if strings.ToLower(filepath.Ext(full)) == want {
			files = append(files, full)
		}
	}
	return files, nil
}
