package asciiart

import (
	"os"
	"strings"
)

type AsciiArt struct {
	asciiChars map[rune][8]string
}

func NewAsciiArt() *AsciiArt {
	return &AsciiArt{
		asciiChars: make(map[rune][8]string),
	}
}

func (a *AsciiArt) FillMap() error {
	text, err := os.ReadFile("assets/fonts/standard.txt")
	if err != nil {
		return err
	}

	data := strings.Split(string(text), "\n")

	var currentChar rune = 32
	var currentCharRepresentation [8]string
	var count int
	for i, line := range data {
		if i == 0 || i == len(data)-1 {
			continue
		}
		if len(line) == 0 {
			count = 0
			currentChar++
			currentCharRepresentation = [8]string{}
			continue
		}
		currentCharRepresentation[count] = line
		count++
		if count == 8 {
			a.asciiChars[currentChar] = currentCharRepresentation
		}
	}
	return nil
}

func (a *AsciiArt) AsciiStringBuilder(s string) string {
	var builder strings.Builder
	inputStrInRune := []rune(s)
	for i := 0; i <= 7; i++ {
		for _, char := range inputStrInRune {
			if charRepresentation, found := a.asciiChars[char]; found {
				builder.WriteString(charRepresentation[i])
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (a *AsciiArt) AsciiMake(inputStr string) (string, error) {
	if err := a.FillMap(); err != nil {
		return "", err
	}

	var builder strings.Builder
	arr := strings.Split(inputStr, "\n")
	for _, word := range arr {
		if word == "" && len(arr) > 1 {
			builder.WriteString("\n")
			continue
		}
		builder.WriteString(a.AsciiStringBuilder(word))
	}
	return builder.String(), nil
}
