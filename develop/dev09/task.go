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

func savePage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}

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
	if err := savePage("https://losst.ru/komanda-wget-linux"); err != nil {
		fmt.Println(err)
	}
}
