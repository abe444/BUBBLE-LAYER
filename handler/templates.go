package handler

import (
	"fmt"
	"html/template"
	"strings"
)

// used in html templates for page list
func IntRange(n int) []int {
	res := make([]int, n)
	if n == 0 {
		res[0] = 0
		return res
	}
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return res
}

func FormatBody(input string) template.HTML {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "&gt;") {
			lines[i] = fmt.Sprintf("<span class=\"greentext\">%s</span>", line)
		}
	}
	return template.HTML(strings.Join(lines, "<br>\n"))
}
