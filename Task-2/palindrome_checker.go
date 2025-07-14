package main

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	cleaned_string := make([]int32, 0, len(s))
	for _, chr := range s {
		if unicode.IsLetter(chr) || unicode.IsDigit(chr) {
			cleaned_string = append(cleaned_string, chr)
		}
	}


	n := len(cleaned_string)
	for i := 0; i < n/2; i++ {
		if cleaned_string[i] != cleaned_string[n-1-i] {
			return false
		}
	}
	return true
}
