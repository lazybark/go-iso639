package detect

import (
	"testing"
)

func TestGetDominantScript(t *testing.T) {
	tests := []struct {
		text     string
		expected string
	}{
		{"Hello, World!", "Latin"},
		{"Hello, World! 안녕 세계", "Latin"},
		{"Hello, World! 안녕 세계 1231234123121231232323", "Latin"},
		{"Привет, мир!", "Cyrillic"},
		{"Привет, мир! Hello", "Cyrillic"},
		{"مرحبا بالعالم", "Arabic"},
		{"Hello مرحبا بالعالم", "Arabic"},
		{"שלום עולםمرحبا بال ", "Hebrew"},
		{"שלום עולם", "Hebrew"},
		{"नमस्ते दुनिया", "Devanagari"},
		{"こんにちは世界", "Hiragana"},
		{"안녕하세요 세계", "Hangul"},
		{"ሰላም ልዑል", "Ethiopic"},
		{"ᎣᏏᏲ ᎡᎶᎯ", "Cherokee"},
		{"", ""},
	}

	for _, test := range tests {
		result := GetDominantScript(test.text)
		if result != test.expected {
			t.Errorf("For text %q, expected %q, but got %q", test.text, test.expected, result)
		}
	}
}
