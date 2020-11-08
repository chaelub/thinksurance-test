package pig_latin

import (
	"fmt"
	"testing"
)

var invalidInputs = []string{
	"",
	"ya",
	"sya",
	"day",
	"hello",
	"helloya",
}

func TestInvalidInput(t *testing.T) {
	for i := range invalidInputs {
		isPigLatin := isPigLatin(invalidInputs[i])
		if isPigLatin {
			t.Error(fmt.Sprintf("the word [%s] is a pig Latin word\n", invalidInputs[i]))
		}
	}
}

var validInputs = []string{
	"iya",
	"aya",
	"ellohya",
	"allyya",
}

func TestValidInput(t *testing.T) {
	for i := range validInputs {
		isPigLatin := isPigLatin(validInputs[i])
		if !isPigLatin {
			t.Error(fmt.Sprintf("the word [%s] is a pig Latin word\n", validInputs[i]))
		}
	}
}

func TestCase(t *testing.T) {
	if (isPigLatin("iya") != isPigLatin("Iya")) ||
		(isPigLatin("i") != isPigLatin("I")) {
		t.Error("the case of letters shouldn't affect the test result")
	}
}
