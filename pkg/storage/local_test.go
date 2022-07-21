// Unit tests for local storage
// Images saving in test-storage dir

package storage

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"testing"
)

func setTestStorage() LocalStorage {
	lc := LocalStorage{"./test-storage"}

	return lc
}

func TestGetImgByName(t *testing.T) {
	imgName := "test-image.jpeg"
	storage := setTestStorage()

	path, err := storage.GetImgByName(imgName)
	if err != nil {
		t.Fatal(err)
	}

	if path != fmt.Sprintf("%s/%s", storage.path, imgName) {
		t.Fatal("Incorrect full path")
	}
}

func TestSaveImg(t *testing.T) {
	storage := setTestStorage()
	imgName := "test-image.jpeg"
	path := fmt.Sprintf("%s/%s", storage.path, imgName)

	imgBytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	img, format, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		t.Fatal(err)
	}

	err = storage.SaveImg("50_test_save", format, img)
	if err != nil {
		t.Fatal(err)
	}

	files, err := ioutil.ReadDir(storage.path)
	if err != nil {
		t.Fatal(err)
	}

	savedFile := ""

	for _, file := range files {
		if file.Name() == "50_test_save" {
			savedFile = file.Name()
		}
	}

	if savedFile == "" {
		t.Fatal("File is not saved")
	}
}
