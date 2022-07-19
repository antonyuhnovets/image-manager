// Realization methods for storage if it local

package storage

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type LocalStorage struct {
	path string
}

// Create dir for local storage if doesn`t exist
func SetLocalStorage(path string) Entity {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	return &LocalStorage{path}
}

// Save image in dir according to format and quality
// Open setted local path, start writer and encode
func (lc *LocalStorage) SaveImg(name string, format string, img image.Image) error {
	fpath := fmt.Sprintf("%s/%s", lc.path, name)

	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer out.Close()

	switch format {
	case "jpeg":
		err = jpeg.Encode(out, img, nil)
		log.Println("Image saved in jpeg")
	case "png":
		err = png.Encode(out, img)
		log.Println("Image saved in png")
	}

	return err
}

// Find image by name and return full path to it
// Walk through dir and chek if file with passed name exist
func (lc *LocalStorage) GetImgByName(name string) (string, error) {
	var fpath string

	files, err := ioutil.ReadDir(lc.path)
	if err != nil {
		return name, err
	}

	log.Printf("Reading dir, searching for %s\n", name)

	for _, file := range files {
		if file.Name() == name {
			fpath = fmt.Sprintf("%s/%s", lc.path, file.Name())

			log.Printf("Downloading %s\n", file.Name())

			return fpath, nil
		}
	}

	return "", errors.New("file not found")
}
