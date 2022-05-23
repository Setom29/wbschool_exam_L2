package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	s, err := getFullString("")
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, "", s) {
		t.Error(`Expected "", got `, s)
	}
}

func TestCustomString(t *testing.T) {
	s, err := getFullString("a4bc2d5e")
	okStr := "aaaabccddddde"
	if err != nil {
		t.Error(err)
	}
	if !assert.Equal(t, okStr, s) {
		t.Errorf(`Expected "%s", got "%s"`, okStr, s)
	}
}

func TestInvalidString(t *testing.T) {
	s, err := getFullString("45")
	if err == nil {
		t.Errorf(`Expected "Invalid string", got "%s"`, s)
	}
}
