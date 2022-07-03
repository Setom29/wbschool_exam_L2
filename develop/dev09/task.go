package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func savePage(url string, filename string) error {
	// vreate response
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	// save response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// if len(flag.Args()) == 0 {
	// 	fmt.Println("url not given")
	// }
	if err := savePage("https://losst.ru/komanda-wget-linux", "index.html"); err != nil {
		fmt.Println(err)
	}
}
