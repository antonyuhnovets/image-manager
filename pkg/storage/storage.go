// Set up abstract entity for different types of storage
// Make service scalable and separated

package storage

import (
	"image"

	"github.com/antonyuhnovets/image-manager/pkg/config"
)

// Declare storage methods
type Entity interface {
	SaveImg(string, string, image.Image) error
	GetImgByName(string) (string, error)
}

// Get methods for setted storage type
func GetStorage(cfg *config.Config) Entity {
	switch cfg.Storage {
	default:
		return SetLocalStorage(cfg.Path)
	}
}
