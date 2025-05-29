package lengthoflastword

import "testing"


func TestLastWord(t *testing.T) {
	tests := []struct{
		name, s string
		expected int
	}{
		{"Test 1","Hello world",5},
		{"Test 2","   fly me   to   the moon  ",4},
		{"Test 3","luffy is still joyboy",6},
	}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			output := lengthOfLastWord(tc.s)
			if output != tc.expected {
				t.Errorf("string = '%s' \n expected: %d; output: %d", tc.s,tc.expected,output)
			}
		})
	}

}
