package main

import (
	"fmt"
	exiftool "github.com/barasher/go-exiftool"
	"github.com/h2non/bimg"
	"os"
)

func main() {
	pdf := "./sample.pdf"
	buffer, err := bimg.Read(pdf)
	if err != nil {
		fmt.Println(err)
	}

	jpeg, err := bimg.NewImage(buffer).Convert(bimg.JPEG)
	if err != nil {
		fmt.Println(err)
	}

	if bimg.NewImage(jpeg).Type() == "jpeg" {
		fmt.Println("JPEG")
	}

	file, err := os.Create("./converted.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(jpeg)
	if err != nil {
		fmt.Println(err)
	}

	// Get Exif
	jpeg_path := "./converted.jpeg"
	getExif(jpeg_path)

	// Remove images
	os.Remove(jpeg_path)
}

func getExif(file string) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Println("Error when initializing Exiftool:", err)
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(file)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}
	}
}
