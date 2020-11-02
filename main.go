package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

func main() {
	var files []string
	filename := "."

	if len(os.Args) < 2 {
		filename = "."
	} else {
		filename = os.Args[1]
	}

	file, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	switch mode := file.Mode(); {
	case mode.IsDir():

		err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
	case mode.IsRegular():
		files = append(files, filename)
	}

	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ext := filepath.Ext(f)
		var config image.Config

		switch ext {
		case ".webp":
			config, err = webp.DecodeConfig(bytes.NewReader(data))
		case ".gif":
			config, err = gif.DecodeConfig(bytes.NewReader(data))
		case ".jpg", ".jpeg":
			config, err = jpeg.DecodeConfig(bytes.NewReader(data))
		case ".png":
			config, err = png.DecodeConfig(bytes.NewReader(data))
		case ".bmp":
			config, err = bmp.DecodeConfig(bytes.NewReader(data))
		case ".tiff":
			config, err = tiff.DecodeConfig(bytes.NewReader(data))
		default:
			fmt.Printf("%s: Unknown file format!\n", filename)
			continue
		}
		if err != nil {
			fmt.Printf("Error reading file: %s \t ", filename)
			fmt.Println(err)
		}

		fmt.Printf("Type:%s \t Width=%d \t Height=%d \t Name:%s \n", ext, config.Width, config.Height, f)

	}

}
