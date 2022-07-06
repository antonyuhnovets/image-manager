package controllers

import (
	"log"

	"github.com/antonyuhnovets/image-manager/pkg/storage"
	"github.com/antonyuhnovets/image-manager/pkg/utils"
)

func CompressAndSave(s storage.Entity, file []byte) {
	var err error
	qualities := []uint{100, 75, 50, 25}
	id := utils.IdGen()
	for _, q := range qualities {
		err = s.SaveImg(id, file, uint(q))
		if err != nil {
			log.Fatalf("Error compressing and saving: %s", err)
		}
	}
}
