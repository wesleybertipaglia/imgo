package infra

import (
	"errors"
	"image"
	"io"
	"os"

	"github.com/disintegration/imaging"
)

type JpegEncoder struct{}

func (e JpegEncoder) Encode(w io.Writer, img image.Image, quality int) error {
	if file, ok := w.(*os.File); ok {
		return imaging.Save(img, file.Name(), imaging.JPEGQuality(quality))
	}
	return errors.ErrUnsupported
}
