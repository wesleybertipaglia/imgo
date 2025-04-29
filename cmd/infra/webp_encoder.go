package infra

import (
	"image"
	"io"

	"github.com/chai2010/webp"
)

type WebpEncoder struct{}

func (e WebpEncoder) Encode(w io.Writer, img image.Image, quality int) error {
	return webp.Encode(w, img, &webp.Options{
		Lossless: false,
		Quality:  float32(quality),
	})
}
