package service
import (
	"strings"
	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"errors"
)
var emptyMessageError = errors.New("нет текста")
func TextDetector(message string) (string, error) {
	var finishText string
	if message == "" {
		return "", emptyMessageError
	}
	if strings.ContainsAny(message, "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM") {
		finishText = morse.ToMorse(message)
	} else {
		finishText = morse.ToText(message)
	}
	return finishText, nil
}   