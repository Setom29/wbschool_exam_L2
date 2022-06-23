package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type Cut struct {
	f    []string
	d    string
	s    bool
	text string
	res  string
}

func newCut() (*Cut, error) {
	var text string
	fields := flag.String("f", "1,2", `-f - "fields" - choose the fields (columns)`)
	d := flag.String("d", "\t", `-d - "delimiter" - use another delimiter`)
	s := flag.Bool("s", false, `-s - "separated" - the strings with delimiter only`)
	flag.Parse()
	if len(flag.Args()) == 0 {
		return nil, errors.New("Wrong args amount.")
	}
	text = flag.Args()[0]
	fmt.Println(*fields, *d, *s)
	return &Cut{f: strings.Split(*fields, ","), d: *d, s: *s, text: text}, nil
}

func (c *Cut) customCut() error {
	res := make([]string, 0, len(c.f))
	textArr := strings.Split(c.text, c.d)
	if len(textArr) <= 1 {
		if c.s {
			res = nil
		} else {
			res = append(res, c.text)
		}

	} else {
		var ind int
		var err error
		for _, v := range c.f {
			ind, err = strconv.Atoi(v)
			if err != nil {
				return errors.New("Wrong type of fields")
			}
			if ind >= len(textArr) || ind < 1 {
				return errors.New("Wrong column number")
			} else {
				res = append(res, textArr[ind])
			}
		}
	}
	c.res = strings.Join(res, " ")
	return nil
}

func main() {
	c, err := newCut()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.customCut()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(c.res)
	}

}
