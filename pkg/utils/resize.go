package utils

import (
	"image"

	"github.com/nfnt/resize"
)

// Change size of image according to guality
func ResizeImg(img image.Image, quality uint) image.Image {
	width := (uint(img.Bounds().Dx()) * quality) / 100
	height := (uint(img.Bounds().Dy()) * quality) / 100
	m := resize.Resize(width, height, img, resize.Lanczos2)
	return m
}
