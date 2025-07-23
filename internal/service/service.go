package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	count := 0
	for _, v := range s {
		if v == '.' || v == '-' || v == ' ' {
			count++
		}
	}

	return count == len([]rune(s))
}

func Analysis(s string) (string, error) {
	text := strings.TrimSpace(s)
	if len(text) == 0 {
		return "", fmt.Errorf("передана пустая строка")
	}

	if isMorse(text) {
		return morse.ToText(text), nil
	}

	return morse.ToMorse(text), nil
}
