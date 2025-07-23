package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseString(input string) bool {
	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' && ch != '\n' && ch != '\r' && ch != '\t' {
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
	if !strings.Contains(strTrim, " ") && !strings.Contains(strTrim, "\n") && !strings.Contains(strTrim, "\r") && !strings.Contains(strTrim, "\t") {
		// Нет пробелов - возвращаем исходный текст
		return str, nil
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
