package controllers

import (
	"github.com/antonyuhnovets/image-manager/pkg/storage"
)

// Save image in 4 sizes
func CompressAndSave(id string, s storage.Entity, file []byte) error {
	var err error
	qualities := []uint{100, 75, 50, 25}
	for _, q := range qualities {
		err = s.SaveImg(id, file, uint(q))
		if err != nil {
			return err
		}
	}

	return nil
}
