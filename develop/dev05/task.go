package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
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

type flags struct {
	A       int
	B       int
	C       int
	c       int
	i       bool
	v       bool
	F       bool
	n       bool
	pattern string
}

func parseFlags() (*flags, error) {
	A := flag.Int("A", 0, `-A - "after" print +N lines after a match`)
	B := flag.Int("B", 0, `-B - "before" print +N lines to match`)
	C := flag.Int("C", 0, `-C - "context" (A+B) print ±N lines around the match`)
	c := flag.Int("c", 0, `-c - "count" (number of rows)`)
	i := flag.Bool("i", false, `-i - "ignore-case" (ignore case)`)
	v := flag.Bool("v", false, `-v - "invert" (instead of a match, exclude)`)
	F := flag.Bool("F", false, `-F - "fixed", exact match with string, not pattern`)
	n := flag.Bool("n", false, `-n - "line num", print line number`)

	*i = true
	*n = true

	flag.Parse()

	if len(flag.Args()) == 0 {
		return nil, errors.New("the pattern must be specified")
	}
	pattern := flag.Args()[0]

	if *i {
		pattern = strings.ToLower(pattern)
	}
	return &flags{*A, *B, *C, *c, *i, *v, *F, *n, pattern}, nil
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

func patternSearch(strs []string, f *flags) map[int]bool {
	indexMap := make(map[int]bool)
	for ind, str := range strs {
		// limiting number of indexes by -c
		if (f.c == len(indexMap)) && (f.c > 0) {
			break
		}
		// ignore case
		if f.i {
			str = strings.ToLower(str)
		}
		if f.F {
			if strings.Contains(str, f.pattern) {
				indexMap[ind] = true
			}
		} else {
			//use regexp
			match, _ := regexp.MatchString(f.pattern, str)
			if f.v && !match {
				// exclude match
				indexMap[ind] = true
			} else if !f.v && match {
				// include match
				indexMap[ind] = true
			}
		}
	}
	return indexMap
}

func getStrsNearMatch(strs []string, indexMap map[int]bool, f *flags) {
	// count the correct number of rows
	if f.C > 0 {
		if f.A == 0 {
			f.A = f.C
		}
		if f.B == 0 {
			f.B = f.C
		}
	}
	// get indexes before and after the match
	indexesArr := make([]int, 0)
	var ind int
	for key := range indexMap {
		for i := 0; i < f.B; i++ {
			ind = key - i - 1
			if ind >= 0 {
				indexesArr = append(indexesArr, ind)
			}
		}
		for i := 0; i < f.A; i++ {
			ind = key + i + 1
			if ind < len(strs) {
				indexesArr = append(indexesArr, ind)
			}
		}
	}
	for _, el := range indexesArr {
		indexMap[el] = true
	}
}

func printStrings(strs []string, indexMap map[int]bool, f *flags) {
	keys := make([]int, 0)
	for k := range indexMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if f.n {
		for _, key := range keys {
			fmt.Print(key, strs[key], "\n")
		}
	} else {
		for _, key := range keys {
			fmt.Print(strs[key], "\n")
		}
	}
}

func customGrep() error {
	f, err := parseFlags()
	if err != nil {
		return err
	}
	strs, err := readFile("test1.txt")
	if err != nil {
		return err
	}
	// get map of indexes
	indexMap := patternSearch(strs, f)
	// add indexes near the match string
	getStrsNearMatch(strs, indexMap, f)
	// print strings with or without the numbers
	printStrings(strs, indexMap, f)
	return nil
}

func main() {
	err := customGrep()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
