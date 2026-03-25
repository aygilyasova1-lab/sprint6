package service
import (
	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"errors"
)
var emptyMessageError = errors.New("нет текста")
func TextDetector(message string) (string, error) {
	var finishText string
	var isMorse bool
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