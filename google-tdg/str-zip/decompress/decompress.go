package decompress

import (
	"strconv"
	"strings"
)

func findCloseBracket(str string, openBracket int) int {
	balance := 1
	for i := openBracket + 1; i < len(str); i++ {
		switch str[i] {
		case '[':
			balance++
		case ']':
			if balance == 1 {
				return i
			}
			balance--
		}
	}
	panic("unbalanced string")
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseNumber(s string, start int, end *int) int {
	for i := start; i < len(s); i++ {
		if s[i] == '[' {
			*end = i
			r, _ := strconv.Atoi(s[start:i])
			return r
		} else if !isDigit(s[i]) {
			panic("invalid multiplier")
		}
	}
	panic("string finished")
}

func Decompress_recursive(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str); {
		c := str[i]
		if !isDigit(c) {
			builder.WriteByte(c)
			i++
		} else {
			var openBracket int
			multiplier := parseNumber(str, i, &openBracket)
			closeBracket := findCloseBracket(str, openBracket)
			substr := Decompress_recursive(str[openBracket+1 : closeBracket])
			builder.WriteString(strings.Repeat(substr, multiplier))
			i = closeBracket + 1
		}
	}
	return builder.String()
}

type Zip struct {
	Repeats int
	builder strings.Builder
}

func Decompress(str string) string {
	stack := []Zip{{}}
	top := 0
	for i := 0; i < len(str); {
		c := str[i]
		if c == ']' {
			substr := stack[top].builder.String()
			count := stack[top].Repeats
			stack = stack[0:top]
			top--
			stack[top].builder.WriteString(strings.Repeat(substr, count))
			i++
		} else if !isDigit(c) {
			stack[top].builder.WriteByte(c)
			i++
		} else {
			var openBracket int
			multiplier := parseNumber(str, i, &openBracket)
			stack = append(stack, Zip{Repeats: multiplier})
			top++
			i = openBracket + 1
		}
	}
	return stack[0].builder.String()
}
