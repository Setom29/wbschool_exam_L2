package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func isAnagram(a, b string) bool {
	runeA := []rune(a)
	runeB := []rune(b)
	if len(runeA) != len(runeB) {
		return false
	}
	sort.Slice(runeA, func(i, j int) bool {
		return runeA[i] < runeA[j]
	})
	sort.Slice(runeB, func(i, j int) bool {
		return runeB[i] < runeB[j]
	})
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			return false
		}
	}
	return true
}
func getAnagramsMap(arr []string) map[string][]string {
	anagramMap := make(map[string][]string)
	for _, el := range arr {
		el = strings.ToLower(el)
		if len(anagramMap) == 0 {
			anagramMap[el] = make([]string, 0)
		} else {
			var counter int
			for key, _ := range anagramMap {
				isAn := isAnagram(key, el)
				if isAn {
					anagramMap[key] = append(anagramMap[key], el)
					break
				} else if (!isAn) && (counter == len(anagramMap)-1) {
					anagramMap[el] = make([]string, 0)
					break
				}
			}
		}
	}
	for key, val := range anagramMap {
		if len(val) == 0 {
			delete(anagramMap, key)
		} else {
			sort.Slice(anagramMap[key], func(i, j int) bool {
				return anagramMap[key][i] < anagramMap[key][j]
			})
		}

	}
	return anagramMap
}

func main() {
	arr := []string{"пятак", "пятка", "листок", "тяпка", "столик", "слиток"}
	fmt.Println(getAnagramsMap(arr))
}
