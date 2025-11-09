package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/word-pattern/description/?envType=study-plan-v2&envId=top-interview-150
func main() {
	var pattern, s string
	fmt.Scanln(&pattern)
	fmt.Scanln(&s)
	result := wordPattern(pattern, s)
	fmt.Println(result)
}

func wordPattern(pattern string, s string) bool {

	if !(len(s) >= 1 && len(s) <= 3000) {
		return false
	}

	if !(len(pattern) >= 1 && len(pattern) <= 3000) {
		return false
	}

	if !IsValidString(s) {
		return false
	}
	if !isSmallLetters(s) {
		return false
	}

	//сам алгоритм

	parts := strings.Fields(s)

	if len(parts) != len(pattern) {
		return false
	}

	correspondence := map[string]rune{}
	PatternExisting := map[rune]bool{}
	for i, r := range pattern {
		value, exist := correspondence[parts[i]]
		if PatternExisting[r] != exist {
			return false
		}
		if exist {
			if r == value {
				if !PatternExisting[r] {
					return false
				}
			} else {
				return false
			}

		} else {

			PatternExisting[r] = true
			correspondence[parts[i]] = r
		}
	}

	return true
}

func IsValidString(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= 'a' && c <= 'z' || c == ' ') {
			return false
		}
	}
	return true
}

func isSmallLetters(s string) bool {
	for _, char := range s {
		if !(char >= 'a' && char <= 'z' || char == ' ') {
			return false
		}
	}
	return true
}

// AF718JT
