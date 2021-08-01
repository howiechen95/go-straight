package ch0

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestSimpleFor(t *testing.T) {
	for i := 0; i < 10*10000; i++ {
		fmt.Println(i)
	}
}

func TestMultiFor(t *testing.T) {
	var tickets chan byte
	var tokenWg sync.WaitGroup
	max_tickets := 10 * 10000
	tx_size := 10 * 10000

	tickets = make(chan byte, max_tickets)
	for i := 0; i < max_tickets; i++ {
		tickets <- 1
	}

	tokenWg.Add(tx_size)
	i := 0
	for i < tx_size {
		i = i + 1
		<-tickets
		go func() {
			defer func() {
				tokenWg.Done()
				tickets <- 1
			}()
			fmt.Println(i)
		}()
	}

	tokenWg.Wait()
}

func TestFilePath(t *testing.T) {
	fmt.Println(filepath.Join("G", "F", "G"))
	fmt.Println(filepath.Join("G/F", "G"))
	fmt.Println(filepath.Join("gfg", "GFG"))
	fmt.Println(filepath.Join("Geeks", "for", "Geeks"))
}

func TestImg(t *testing.T) {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	//draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)
}

func TestImg001(t *testing.T) {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("20.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	jpg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	jpeg.Encode(file, jpg, nil)
}
