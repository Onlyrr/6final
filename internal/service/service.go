package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseString(input string) bool {
	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}
	return true
}

func Analysis(str string) (string, error) {
	strTrim := strings.TrimSpace(str)
	if strTrim == "" {
		return "", errors.New("пустая строка")
	}

	if isMorseString(strTrim) {
		result := morse.ToText(strTrim)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("не удалось распознать код")
		}
		return result, nil
	} else {
		result := morse.ToMorse(strTrim)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("не удалось конвертировать текст в Morse")
		}
		return result, nil
	}
}
