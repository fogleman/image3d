package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/fogleman/image3d"
)

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	return im, err
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run examples/main.go IMAGES")
		return
	}
	var images []image.Image
	for _, filename := range args {
		fmt.Println(filename)
		im, err := loadImage(filename)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, im)
	}
	im := image3d.NewImage3D(images)
	fmt.Println(im.W, im.H, im.D)
	for z := 0; z < im.D-1; z++ {
		z0 := float64(z)
		z1 := z0 + 0.5
		fmt.Println(z0, im.At(float64(im.W/2), float64(im.H/2), z0))
		fmt.Println(z1, im.At(float64(im.W/2), float64(im.H/2), z1))
	}
}
