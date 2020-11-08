package pig_latin

import (
	"fmt"
	"testing"
)

var invalidInputsWiki = []string{
	"",
	"ay",
	"say",
	"day",
	"hello",
	"helloay",
}

func TestInvalidInputWiki(t *testing.T) {
	for i := range invalidInputsWiki {
		isPigLatin := isPigLatinWiki(invalidInputsWiki[i])
		if isPigLatin {
			t.Error(fmt.Sprintf("the word [%s] is a pig Latin word\n", invalidInputsWiki[i]))
		}
	}
}

var validInputsWiki = []string{
	"iay",
	"aay",
	"ellohay",
	"allyay",
}

func TestValidInputWiki(t *testing.T) {
	for i := range validInputsWiki {
		isPigLatin := isPigLatinWiki(validInputsWiki[i])
		if !isPigLatin {
			t.Error(fmt.Sprintf("the word [%s] is a pig Latin word\n", validInputsWiki[i]))
		}
	}
}

func TestCaseWiki(t *testing.T) {
	if (isPigLatinWiki("iya") != isPigLatinWiki("Iya")) ||
		(isPigLatinWiki("i") != isPigLatinWiki("I")) {
		t.Error("the case of letters shouldn't affect the test result")
	}
}
