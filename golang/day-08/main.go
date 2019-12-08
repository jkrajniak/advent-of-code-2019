package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Wide = 25
	Height = 6
)

var LayerSize = Wide * Height

func Calculate(layers string, layerSize int) int {
	numZeros := math.Inf(1)
	answer := 0
	for i := 0; i < len(layers); i = i + layerSize {
		layer := layers[i:i+layerSize]
		zeros := strings.Count(layer, "0")
		if zeros == 0 {
			continue
		}
		if float64(zeros) < numZeros {
			answer = strings.Count(layer, "1") * strings.Count(layer, "2")
			numZeros = float64(zeros)
		}
	}
	return answer
}

func ComputeImage(layers string, layerSize int) []int {
	imagePixels := make([][]int, layerSize)
	rawImage := make([]int, layerSize)

	// Build layers of pixels.
	for i := 0; i < len(layers); i = i + layerSize {
		layer := layers[i:i+layerSize]
		for i, c := range layer {
			intVal, _ := strconv.Atoi(string(c))
			imagePixels[i] = append(imagePixels[i], intVal)
		}
	}

	for i := 0; i < layerSize; i++ {
		rawImage[i] = DecodePixel(imagePixels[i])
	}

	return rawImage
}

func DecodePixel(pixels []int) int {
	for _, v := range pixels {
		if v == 0 || v == 1 {
			return v
		}
	}
	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		answer := Calculate(line, LayerSize)
		fmt.Println(answer)
		rawImage := ComputeImage(line, LayerSize)
		for i, p := range rawImage {
			if i % Wide == 0 {
				fmt.Println("")
			}
			fmt.Print(p)
		}
		MakeImage(rawImage)
	}
}

func MakeImage(data []int) {
	img := image.NewGray(image.Rect(0, 0, Wide, Height))
	pixel := 0
	for y := 0; y < Height; y++ {
		for x := 0; x < Wide; x++ {
			if data[pixel] == 0 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
			pixel++
		}
	}
	f, _ := os.Create("image.png")
	
	png.Encode(f, img)
}


func convertToIntSlice(s []string) ([]int, error) {
	var ints []int
	for _, s := range s {
		i, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
