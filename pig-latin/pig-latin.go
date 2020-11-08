package pig_latin

import (
	"strings"
)

const (
	validOneLetterWords = "ai"
	vowels              = "aeiuo"
	suffix              = "ya"
)

var (
	suffixLen        = len(suffix)
	testLetterOffset = suffixLen + 1
)

/*


 */
func isPigLatin(word string) bool {
	word = strings.ToLower(word)

	// any pig Latin word has to be >= 3 digits (a - > aya, I -> Iya)
	if len(word) < 3 {
		return false
	}
	// a pig Latin word has to contain suffix 'ya'. If it doesn't, it's not a pig Latin word
	if !strings.HasSuffix(word, suffix) {
		return false
	}
	// if it's the smallest pig Latin word, so it has to starts with valid one word vowel (a, I)
	if len(word) == 3 && !strings.ContainsAny(word[:1], validOneLetterWords) {
		return false
	}
	// it's not a pig Latin word if the word starts with consonant, but it's a vowel right before the suffix
	wordLen := len(word)
	if !strings.ContainsAny(word[:1], vowels) &&
		strings.ContainsAny(word[wordLen-testLetterOffset:wordLen-suffixLen], vowels) {
		return false
	}

	return true
}

/*
maybe I should solve this task as "is a string a permutation of another string",
because of a strange example in the task description:
	"You have 2 given words as pig Latin candidates (e.g. center → tercen, tomorrow → morrowto).
     Your task is to write a method returning true/false if the given word is a pig Latin or not."
*/
func isPermutation(init, mutated string) bool {
	return false
}
