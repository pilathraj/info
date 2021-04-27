package main

import (
	"errors"
	"strings"
)

var ErrEmpty = errors.New("Empty string")

type IstringService interface {
	Uppercase(s string) (string, error)
	Count(s string) int
}

type stringService struct{}

func (svc stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}
func (svc stringService) Count(s string) int {
	return len(s)
}
