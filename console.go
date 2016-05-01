package console

import (
	"bufio"
	"os"
	"strings"
)

func ReadLine() (string, error) {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimRight(text, "\n\r")
	return text, nil
}
