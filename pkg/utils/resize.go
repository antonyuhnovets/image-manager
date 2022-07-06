package utils

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeImg(img image.Image, quality uint) image.Image {
	width := (uint(img.Bounds().Dx()) * quality) / 100
	high := (uint(img.Bounds().Dy()) * quality) / 100
	m := resize.Resize(width, high, img, resize.Lanczos2)
	return m
}
