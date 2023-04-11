package main

import (
	"errors"
	"strings"
)

// Stringservice提供有关字符串的操作。
type StringService interface {
	// 转化为大写
	Uppercase(string) (string, error)
	// 计数
	Count(string) int
}

// stringService是Stringservice的具体实现
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
