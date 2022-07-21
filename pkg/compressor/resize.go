// Implement resize tool using nfnt library

package compressor

import (
	"image"

	"github.com/nfnt/resize"
)

type ResizerNFNT struct {
	width, height uint
}

// Get image width and height
func (r *ResizerNFNT) GetBounds(img image.Image) {
	r.width = uint(img.Bounds().Dx())
	r.height = uint(img.Bounds().Dy())
}

// Change size of image according to passed guality
func (r *ResizerNFNT) ResizeImg(img image.Image, quality uint) image.Image {
	width := r.width * quality / 100
	height := r.height * quality / 100
	m := resize.Resize(width, height, img, resize.Lanczos2)

	return m
}
