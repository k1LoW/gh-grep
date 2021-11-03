package internal

import (
	"unicode/utf8"

	"github.com/fatih/color"
)

func PrintLine(line string, matches [][]int, c *color.Color) string {
	if !utf8.Valid([]byte(line)) {
		return line
	}
	if len(matches) == 0 {
		return line
	}
	pos := 0
	colored := ""
	for _, m := range matches {
		if pos < m[0] {
			colored += line[pos:m[0]]
		}
		colored += c.Sprint(line[m[0]:m[1]])
		pos = m[1]
	}
	colored += line[pos:]

	return colored
}

func PrintOnlyMatches(line string, matches [][]int, c *color.Color) string {
	if !utf8.Valid([]byte(line)) {
		return ""
	}
	if len(matches) == 0 {
		return ""
	}
	for _, m := range matches {
		return c.Sprint(line[m[0]:m[1]])
	}
	return ""
}
