package utils

import (
	"errors"
	"image"
	"imgo/cmd/infra"
	"io"
	"strings"
)

type ImageEncoder interface {
	Encode(w io.Writer, img image.Image, quality int) error
}

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
