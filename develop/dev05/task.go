package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	A int
	B int
	C int
	c int
	i int
	v bool
	F bool
	n bool
}

func parseFlags() *Flags {
	A := flag.Int("A", 0, `"after" print +N lines after a match`)
	B := flag.Int("B", 0, `"before" print +N lines to match`)
	C := flag.Int("C", 0, `"context" (A+B) print ±N lines around the match`)
	c := flag.Int("c", 0, `-c - "count" (number of rows)`)
	i := flag.Int("i", 0, `-i - "ignore-case" (ignore case)`)
	v := flag.Bool("v", false, `-v - "invert" (instead of a match, exclude)`)
	F := flag.Bool("F", false, `-F - "fixed", exact match with string, not pattern`)
	n := flag.Bool("n", false, `-n - "line num", print line number`)
	flag.Parse()
	fmt.Println(os.Args)
	return &Flags{*A, *B, *C, *c, *i, *v, *F, *n}
}

func readFile(filename string) ([]string, error) {
	var strs []string
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		strs = append(strs, fileScanner.Text())
	}
	return strs, nil
}

func customGrep() error {
	flags := parseFlags()
	strs, err := readFile("test1.txt")
	if err != nil {
		return err
	}
	fmt.Println(flags)
	fmt.Println(strs[0])

	return nil
}

func main() {
	err := customGrep()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
