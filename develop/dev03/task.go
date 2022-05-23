package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getKeys() {
	byCol := flag.Bool("k", false, "specifying a column to sort")
	byNum := flag.Bool("n", false, "sort by numeric value")
	reverseSort := flag.Bool("r", false, "sort in reverse order")
	keepDuplicate := flag.Bool("u", false, "do not output duplicate lines")
	flag.Parse()
}

func main() {
	var strs []string
	// open the file
	file, err := os.Open("some_strings.txt")

	//handle errors while opening
	if err != nil {
		fmt.Printf("Error when opening file: %s\n", err)
	}
	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		strs = append(strs, fileScanner.Text())
	}
	getKeys()
}
