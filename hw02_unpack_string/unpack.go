package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var newSting strings.Builder
	row := strings.Split(s, "")
	skipNext := false

	for idx, char := range row {

		if skipNext {
			skipNext = false
			continue
		}

		var nextChar string

		if idx < len(row)-1 {
			nextChar = row[idx+1]
		} else {
			nextChar = ""
		}

		if char == `\` {

			if _, err := strconv.Atoi(nextChar); err == nil || nextChar == `\` {
				newSting.WriteString(nextChar)
				skipNext = true
				continue
			} else {
				return "", ErrInvalidString
			}
		}

		if digit, err := strconv.Atoi(char); err == nil {

			if _, err := strconv.Atoi(nextChar); err == nil || idx == 0 {
				fmt.Println(" New string", newSting.String())
				return "", ErrInvalidString
			}

			if digit == 0 {
				str := newSting.String()
				newSting.Reset()
				newSting.WriteString(str[:len(str)-1])
			}

			prevChar := row[idx-1]
			repeated := strings.Repeat(prevChar, digit-1)
			newSting.WriteString(repeated)
			continue
		}
		newSting.WriteString(char)
	}
	return newSting.String(), nil
}
