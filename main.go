package main

import (
	"fmt"
	"strings"
)

func CountWords(text string) map[string]int {
	num := 0
	text = strings.ToLower(text)
	for i, ch := range text {
		if ch == ' ' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == '!' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == ';' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == '?' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == '.' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == ':' {
			num++
			text = text[:i] + text[i+1:]
		} else if ch == ',' {
			num++
			text = text[:i] + text[i+1:]
		}

	}
	var name = map[string]int{
		text: num,
	}
	return name
}
func main() {
	fmt.Println(CountWords("Go Go go!"))
	fmt.Println("Hello WOrld")
}
