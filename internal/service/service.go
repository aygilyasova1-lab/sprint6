package service
import (
	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"errors"
	"strings"
)
var emptyMessageError = errors.New("нет текста")
func TextDetector(message string) (string, error) {
	message = strings.TrimSpace(message)
	var finishText string
	isMorse := true
	if message == "" {
		return "", emptyMessageError
	}
	for _, r := range message {
    if r != '.' && r != '-' && r != ' ' {
        isMorse = false
        break
    	}
	}
	if isMorse == false {
		finishText = morse.ToMorse(message)
	} else {
		finishText = morse.ToText(message)
	}
	return finishText, nil
}   