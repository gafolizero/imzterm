package load

import (
	"image"
	"log"
	"os"
)

func Load(filePath string) *image.NRGBA {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()

	if err != nil {
		log.Println("Cannot read file: ", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot read file: ", err)
	}

	return img.(*image.NRGBA)
}
