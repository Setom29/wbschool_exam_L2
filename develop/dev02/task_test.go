package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// empty string
func TestEmptyString(t *testing.T) {
	s, err := Unpack("")
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, "", s) {
		t.Error(`Expected "", got `, s)
	}
}

func TestCustomString(t *testing.T) {
	s, err := Unpack("a4bc2d5e")
	okStr := "aaaabccddddde"
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, okStr, s) {
		t.Errorf(`Expected "%s", got "%s"`, okStr, s)
	}
}

func TestInvalidString(t *testing.T) {
	s, err := Unpack("45")
	if err == nil {
		t.Errorf(`Expected "Invalid string", got "%s"`, s)
	}
}

func TestBackspaceString1(t *testing.T) {
	s, err := Unpack(`qwe\4\5`)
	okStr := "qwe45"
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, okStr, s) {
		t.Errorf(`Expected "%s", got "%s"`, okStr, s)
	}
}

func TestBackspaceString2(t *testing.T) {
	s, err := Unpack(`qwe\45`)
	okStr := `qwe44444`
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, okStr, s) {
		t.Errorf(`Expected "%s", got "%s"`, okStr, s)
	}
}

func TestBackspaceString3(t *testing.T) {
	s, err := Unpack(`qwe\\5`)
	okStr := `qwe\\\\\`
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, okStr, s) {
		t.Errorf(`Expected "%s", got "%s"`, okStr, s)
	}
}
