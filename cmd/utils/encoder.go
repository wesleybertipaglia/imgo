package utils

import (
	"image"
	"io"
)

type ImageEncoder interface {
	Encode(w io.Writer, img image.Image, quality int) error
}
