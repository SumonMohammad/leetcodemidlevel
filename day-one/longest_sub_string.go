package main

import "fmt"

func findLongestSubString(s string) int {
	lastSeen := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, ch := range s {
		if lastIndex, found := lastSeen[ch]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		lastSeen[ch] = i

		currentLen := i - start + 1

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
func main() {
	fmt.Println(findLongestSubString("abccababcd"))
}
