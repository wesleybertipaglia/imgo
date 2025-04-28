package utils

import (
	"errors"
	"imgo/cmd/infra"
	"strings"
)

func GetEncoder(format string) (ImageEncoder, error) {
	switch strings.ToLower(format) {
	case "webp":
		return infra.WebpEncoder{}, nil
	case "jpg", "jpeg":
		return infra.JpegEncoder{}, nil
	case "png":
		return infra.PngEncoder{}, nil
	default:
		return nil, errors.New("unsupported format: " + format)
	}
}
