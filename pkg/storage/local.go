package storage

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/antonyuhnovets/image-manager/pkg/utils"
)

type LocalStorage struct {
	path string
}

func GetLocalStorage(path string) Entity {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return &LocalStorage{path}
}

func (lc *LocalStorage) SaveImg(name string, imgBytes []byte, quality uint) error {
	var err error
	fpath := fmt.Sprintf("%s/%v_%s", lc.path, quality, name)

	img, format, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatalln(err)
	}
	resizedImg := utils.ResizeImg(img, quality)

	out, _ := os.Create(fpath)
	defer out.Close()

	switch format {
	case "jpeg":
		err = jpeg.Encode(out, resizedImg, nil)
		fmt.Println("Image saved in jpeg")
	case "png":
		err = png.Encode(out, resizedImg)
		fmt.Println("Image saved in png")
	}
	return err
}

func (lc *LocalStorage) GetImgByName(name string) (string, error) {
	var fpath string
	files, err := ioutil.ReadDir(lc.path)
	if err != nil {
		return name, err
	}
	fmt.Printf("Reading dir, searching for %s", name)
	for _, file := range files {
		if file.Name() == name {
			fpath = fmt.Sprintf("%s/%s", lc.path, file.Name())
			fmt.Printf("Downloading %s", file.Name())
			return fpath, nil
		}
	}

	return "", errors.New("file not found")
}