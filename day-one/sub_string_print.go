package main

import "fmt"

func printSubString(s string) string {
	lastSeen := make(map[rune]int)
	maxlen := 0
	start := 0
	maxStart := 0

	for i, ch := range s {
		if lastIndex, found := lastSeen[ch]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		lastSeen[ch] = i

		currentLen := i - start + 1

		if currentLen > maxlen {
			maxlen = currentLen
			maxStart = start
		}
	}

	return s[maxStart : maxStart+maxlen]
}
func main() {

	result := printSubString("hellobabyloveyou")
	fmt.Println(result)
}
