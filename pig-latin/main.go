package pig_latin

import "strings"

const (
	validOneLetterWords = "ai"
	vowels              = "aeiuo"
	suffix              = "ya"
)

/*
any pig Latin word has to be >= 3 digits (a - > aya, I -> Iya)
a pig Latin word has to contain suffix 'ya'. If it doesn't, it's not a pig Latin word
the word should starts with vowel too if:
	a vowel is right before the suffix 'ya'
*/
func isPigLatin(word string) bool {
	word = strings.ToLower(word)
	if len(word) < 3 {
		return false
	}
	if !strings.HasSuffix(word, suffix) {
		return false
	}
	if len(word) == 3 && !strings.ContainsAny(word, validOneLetterWords) {
		return false
	}

	return true
}
