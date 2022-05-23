package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func appendLetters(r []rune, l rune, digit []rune) []rune {
	num, err := strconv.Atoi(string(digit))
	if err != nil {
		num = 1
	}
	for i := 0; i < num; i++ {
		r = append(r, l)
	}

	return r
}

func getFullString(s string) (string, error) {
	r := []rune(s)
	// case: empty string
	if len(r) == 0 {
		return "", nil
	}

	arr := []rune{}
	digit := []rune{}
	letter := r[0]
	var err error

	// case: digit on the zero position
	if unicode.IsDigit(r[0]) {
		err = errors.New("Invalid string")
		return "", err
	}

	for i := 0; i < len(r); i++ {
		switch {
		case r[i] == rune('/'):
			arr = append(arr, rune('/'))
		case unicode.IsDigit(r[i]):
			digit = append(digit, r[i])
		default:
			letter = r[i]
		}
		if i == len(r)-1 {
			if letter == rune(-1) {
				arr = appendLetters(arr, r[i], digit)
			} else {
				arr = appendLetters(arr, letter, digit)
			}

		} else if unicode.IsLetter(r[i+1]) {
			arr = appendLetters(arr, letter, digit)
			digit = []rune{}
		}
	}
	return string(arr), nil
}

func main() {
	s := "a4bc2d5e"
	ans, err := getFullString(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(ans)
	}
}
