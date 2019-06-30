package main

import (
	"errors"
	"strings"
)

var ErrEmpty = errors.New("error empty string")

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (s stringService) UpperCase(str string) (string, error) {
	if str == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(str), nil
}

func (s stringService) Count(str string) int {
	return len(str)
}
