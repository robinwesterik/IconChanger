package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

type imageData struct {
	path   string
	width  int
	height int
}

// This function finds the width and height of the image
func resFind(input string) (width int, height int) {
	r, _ := vips.NewImageFromFile(input)
	return r.Width(), r.Height()
}

// This function parses the imageData from the list of images
func parse(imagePaths []string) []imageData {
	var images []imageData
	for _, path := range imagePaths {
		w, h := resFind(path)
		images = append(images, imageData{
			path:   path,
			width:  w,
			height: h,
		})
	}
	return images
}

// This function filters a list of files by extension. Only png, jpg, and jpeg files are returned
func filter(files []string) []string {
	var filteredFiles []string
	for _, file := range files {
		if filepath.Ext(file) == ".png" || filepath.Ext(file) == ".jpg" || filepath.Ext(file) == ".jpeg" {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}

// This function scans folders/subfolders and returns a list of files
func scan(path string) []string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return files
}

func exportPNG(input string, output string, width int, height int) {
	o, _ := vips.NewImageFromFile(input)
	o.Thumbnail(width, height, vips.InterestingAll)
	ep := vips.NewDefaultPNGExportParams()
	image1bytes, _, _ := o.Export(ep)
	os.WriteFile(output, image1bytes, 0644)
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	l := parse(filter(scan(os.Args[1])))
	for _, i := range l {
		exportPNG(os.Args[2], i.path, i.width, i.height)
	}
}
