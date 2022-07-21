// Abstract interfaces and struct for compress images

package compressor

import (
	"bytes"
	"fmt"
	"image"
)

// Compressor serve as main abstraction for methods realization
type Compressor struct {
	qualities []uint                     // set of values for resizing img
	Handler   func(string, []byte) error // func describe full cycle of img resizing
	Entity                               // set of compressing methods
}

// Compressing methods
type Entity interface {
	Resize(image.Image, uint) image.Image
	Convert([]byte) (image.Image, string, error)
}

// Convert bytes to image and return it's format
func (c *Compressor) Convert(imgByte []byte) (image.Image, string, error) {
	img, format, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

// Method could be realized for different libraries according to requirements
// In this case using ResizerNFNT
func (c *Compressor) Resize(img image.Image, quality uint) image.Image {
	r := &ResizerNFNT{}
	r.GetBounds(img)

	return r.ResizeImg(img, quality)
}

// Fill compressor instance and return it with handler function
// Take as argument function, witch allow save img in storage
func GetCompressor(saver func(string, string, image.Image) error) *Compressor {
	compressor := &Compressor{
		qualities: []uint{25, 50, 75, 100},
	}
	f := func(id string, imgBytes []byte) error {
		img, format, err := compressor.Convert(imgBytes)
		if err != nil {
			return err
		}
		for _, q := range compressor.qualities {
			resized := compressor.Resize(img, uint(q))
			saver(fmt.Sprintf("%v_%s", q, id), format, resized)
		}
		return nil
	}
	compressor.Handler = f

	return compressor
}
