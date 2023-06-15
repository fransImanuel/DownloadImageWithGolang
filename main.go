package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// if runtime.GOOS == "windows" {
	// 	fmt.Println("Hello from Windows")
	// }

	// link_gambar := "https://images.bisnis.com/posts/2019/03/26/904507/shutterstock_568477957.jpg"
	// sub_folder := "tmp"

	createDirTmp(sub_folder)
	// loadImageFromModels(sub_folder, link_gambar)
}

func loadImageFromModels(cont string, link_gambar string) {
	// for i, each := range result {
	// 	_ = i
	// 	if (each.Gambar != "") && (each.Tracking_number != "") {
	// 		loadImageFromURL(cont, each.Gambar, each.Tracking_number)
	// 		compressedImage(cont, each.Tracking_number)
	// 	}
	// }
	nama_file_gambar := "ini_gambar"
	loadImageFromURL(cont, link_gambar, nama_file_gambar)
	compressedImage(cont, nama_file_gambar)
}

func compressedImage(cont string, name string) {
	inFile, err := os.Open("./" + cont + "/" + name + ".jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inFile.Close()

	// Decode original image
	img, _, err := image.Decode(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	resizedImg := resize.Resize(60, 30, img, resize.Lanczos3)

	// Create compressed image file
	outFile, err := os.Create("./" + cont + "/" + name + ".jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, resizedImg, nil); err != nil {
		fmt.Println(err)
		return
	}
}

func createDirTmp(dir string) {
	getCurrWd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("getCurrWd: ", getCurrWd)

	path := getCurrWd + fmt.Sprintf("\\%s", dir)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
			fmt.Println(err)
		}
	}
}

func loadImageFromURL(cont string, url string, name string) {
	response, e := http.Get(url)
	if e != nil {
		fmt.Println(e)
	}
	defer response.Body.Close()

	file, err := os.Create("./" + cont + "/" + name + ".jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
	}
}
