package task

import (
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пЯтка", "тяПка", "Листок", "слитОк", "блеск", "стоЛик", "кОт", "ток", "ого"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
				"кот":    {"кот", "ток"},
			},
		},
	}

	for _, test := range tests {
		result := FindAnagram(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входных данных %v ожидалось %v, получено %v", test.input, test.expected, result)
		}
	}
}
