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
