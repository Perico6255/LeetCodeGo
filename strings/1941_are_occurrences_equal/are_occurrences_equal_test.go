package areoccurrencesequal

import "testing"

func TestAreOccurrences(t *testing.T) {
	tests := []struct{
		name, text string
		expected bool
	}{
		{"TestAreOccurrences 1; 'aabbcc' => true","aabbcc",true},
		{"TestAreOccurrences 2; 'aaabbcc' => false","aaabbcc",false},
		{"TestAreOccurrences 3; 'aabbccc' => false","aabbccc",false},
	}

	for _, test  := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := areOccurrencesEqual(test.text)
			if got != test.expected {
				t.Errorf("AreOccurrencesEqual(%s) =  %v; expected: %v",test.text, got, test.expected)
			}
		})
	}
	
}
