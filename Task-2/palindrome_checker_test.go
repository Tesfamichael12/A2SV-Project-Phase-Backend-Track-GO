package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"racecar", true},
		{"hello", false},
		{"Was it a car or a cat I saw?", true},
		{"No lemon, no melon", true},
	}
	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		if got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
