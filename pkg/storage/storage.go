// Set up entity for different types of storage
// Make service scalable

package storage

import (
	"github.com/antonyuhnovets/image-manager/pkg/config"
)

// Declare storage methods
type Entity interface {
	SaveImg(string, []byte, uint) error
	GetImgByName(string) (string, error)
}

// Get methods for setted storage type
func GetStorage(cfg *config.Config) Entity {
	switch cfg.STORAGE {
	default:
		return SetLocalStorage(cfg.STORAGE_PATH)
	}
}
