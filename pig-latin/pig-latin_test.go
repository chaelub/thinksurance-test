package pig_latin

import (
	"fmt"
	"testing"
)

var validInputs = [][2]string{
	{"tomorrow", "morrowto"},
	{"hello", "llohe"},
	{"I", "I"},
	{"are", "are"},
}

func TestValidInput(t *testing.T) {
	for i := range validInputs {
		isPigLatin := isPigLatin(validInputs[i][0], validInputs[i][1])
		if !isPigLatin {
			t.Error(
				fmt.Sprintf("the word [%s] is a pig Latin for [%s]\n",
					invalidInputs[i][1],
					invalidInputs[i][0],
				),
			)
		}
	}
}

var invalidInputs = [][2]string{
	{"tomorrow", "tomorrow"},
	{"different", "lenght"},
	{"hello", "elloh"},
	{"world", "dlrow"},
}

func TestInvalidInput(t *testing.T) {
	for i := range invalidInputs {
		isPigLatin := isPigLatin(invalidInputs[i][0], invalidInputs[i][1])
		if isPigLatin {
			t.Error(
				fmt.Sprintf("the word [%s] is not a pig Latin for [%s]\n",
					invalidInputs[i][1],
					invalidInputs[i][0],
				),
			)
		}
	}
}
