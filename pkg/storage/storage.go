package storage

import (
	"github.com/antonyuhnovets/image-manager/pkg/config"
)

type Entity interface {
	SaveImg(string, []byte, uint) error
	GetImgByName(string) (string, error)
}

func SetStorage(cfg *config.Config) Entity {
	switch cfg.STORAGE {
	default:
		return GetLocalStorage(cfg.STORAGE_PATH)
	}
}
