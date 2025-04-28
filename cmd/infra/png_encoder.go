package infra

import (
	"errors"
	"image"
	"io"
	"os"

	"github.com/disintegration/imaging"
)

type PngEncoder struct{}

func (e PngEncoder) Encode(w io.Writer, img image.Image) error {
	if file, ok := w.(*os.File); ok {
		return imaging.Save(img, file.Name())
	}

	return errors.ErrUnsupported
}
