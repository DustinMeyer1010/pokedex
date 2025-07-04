package utils

import (
	"regexp"
	"strings"
)

func CleanInput(text string) []string {
	text = strings.TrimSpace(text)
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	result := strings.Split(text, " ")

	if result[0] == "" {
		return []string{}
	}

	return result

}
