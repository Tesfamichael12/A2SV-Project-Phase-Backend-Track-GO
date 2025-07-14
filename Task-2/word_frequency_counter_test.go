package main

import (
	"reflect"
	"testing"
)

func TestGetWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "",
			expected: map[string]int{},
		},
		{
			input:    "hello",
			expected: map[string]int{"hello": 1},
		},
		{
			input:    "hello world hello",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			input:    "Go go GO",
			expected: map[string]int{"go": 3},
		},
		{
			input:    "a a a b b c",
			expected: map[string]int{"a": 3, "b": 2, "c": 1},
		},
	}

	for _, tt := range tests {
		actual := wordCounter(tt.input)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("wordCounter(%q) = %v; want %v", tt.input, actual, tt.expected)
		}
	}
}