package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Analysis(str string) (string, error) {
	if strings.ContainsAny(str, ".- ") {
		for _, ch := range str {
			if ch != '.' && ch != '-' && ch != ' ' && ch != '\n' && ch != '\r' {
				// Значит, это всё же текст, останавливаем выполнение этого условия
				break
			}
		}
		result := morse.ToText(str)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("Не удалось распознать код")
		}
		return result, nil
	} else {
		result := morse.ToMorse(str)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("Не удалось конвертировать текст в Morse")
		}
		return result, nil
	}
}
