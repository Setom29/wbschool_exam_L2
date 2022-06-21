package main

import (
	"errors"
	"fmt"
	"os"
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

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sr := []rune(s)
	var s2 string
	var n int
	var backslash bool

	for i, item := range sr {
		// the [0] element is digit
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(sr[i-1]) && sr[i-2] != '\\' {
			return "", ErrInvalidString
		}
		// adding backslash
		if item == '\\' && !backslash {
			backslash = true
			continue
		}
		// letter after backslash
		if backslash && unicode.IsLetter(item) {
			return "", ErrInvalidString
		}
		if backslash {
			s2 += string(item)
			backslash = false
			continue
		}
		if unicode.IsDigit(item) {
			n = int(item - '0')
			if n == 0 {
				s2 = s2[:len(s2)-1]
				continue
			}
			for j := 0; j < n-1; j++ {
				s2 += string(sr[i-1])
			}
			continue
		}
		s2 += string(item)
	}

	return s2, nil
}

func main() {
	s := `qwe\4\5`
	ans, err := Unpack(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(ans)
	}
}

/*
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
}*/
