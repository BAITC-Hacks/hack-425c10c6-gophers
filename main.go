package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <путь_к_картинке>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Ошибка при декодировании изображения: %v\n", err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	totalPixels := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	redPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			if r8 > 150 && g8 < 100 && b8 < 100 {
				redPixels++
			}
		}
	}

	if float64(redPixels)/float64(totalPixels) > 0.20 {
		fmt.Println("DEFECT")
	} else {
		fmt.Println("OK")
	}
}
