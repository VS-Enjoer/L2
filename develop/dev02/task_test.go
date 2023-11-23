package dev02

import "testing"

func TestMainFunction(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for _, tc := range testCases {
		result, err := main(tc.input)
		if err != nil {
			t.Errorf("Неожиданная ошибка: %v", err)
		}
		if result != tc.expected {
			t.Errorf("Ожидалось %s, но получили %s", tc.expected, result)
		}
	}
}
