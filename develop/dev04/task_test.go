package main

import (
	"strings"
	"testing"
)

func TestFindAnagrans(t *testing.T) {
	result := map[string][]string{
		"пятак":  {"пятка", "тяпка"},
		"слиток": {"листок", "столик"},
		"каир":   {"ирак", "раки"},
	}

	dictionary := []string{
		"Пятак",
		"Пятак",
		"пятка",
		"Тяпка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Потоп",
		"Каир",
		"Ирак",
		"раки",
	}

	myResult := findAnagrams(dictionary)

	for key, val := range myResult {
		var (
			anagrams []string
			ok       bool
		)
		if anagrams, ok = result[key]; !ok {
			t.Errorf("excess key: %s", key)
		}
		joinedResultAnagrams := strings.Join(anagrams, " ")
		joinedMyAnagrams := strings.Join(val, " ")
		if joinedMyAnagrams != joinedResultAnagrams {
			t.Errorf("wrong anagrams:\nShould: %s\nGot: %s\n", joinedResultAnagrams, joinedMyAnagrams)
		}
	}
}
