package pig_latin

import (
	"bytes"
	"strings"
)

const (
	validOneLetterWords = "ai"
	vowelsStr           = "aeiuo"
	suffix              = "ay"
)

var (
	suffixLen        = len(suffix)
	testLetterOffset = suffixLen + 1
)

/*
the isPigLatinWiki algorithm what follows wiki rules
https://en.wikipedia.org/wiki/Pig_Latin
*/
func isPigLatinWiki(word string) bool {
	word = strings.ToLower(word)

	// any pig Latin word has to be >= 3 digits (a - > aay, I -> Iay)
	if len(word) < 3 {
		return false
	}
	// a pig Latin word has to contain suffix 'ay'. If it doesn't, it's not a pig Latin word
	if !strings.HasSuffix(word, suffix) {
		return false
	}
	// if it's the smallest pig Latin word, so it has to starts with valid one word vowel (a, I)
	if len(word) == 3 && !strings.ContainsAny(word[:1], validOneLetterWords) {
		return false
	}
	// it's not a pig Latin word if the word starts with consonant, but it's a vowel right before the suffix
	wordLen := len(word)
	if !strings.ContainsAny(word[:1], vowelsStr) &&
		strings.ContainsAny(word[wordLen-testLetterOffset:wordLen-suffixLen], vowelsStr) {
		return false
	}

	return true
}

/*
an implementation of isPigLatin algorithm without the suffix 'ay' ->
we have an original word and the pig Latin candidate
*/
func isPigLatin(origin, candidate string) bool {
	origin = strings.ToLower(origin)
	candidate = strings.ToLower(candidate)

	if len(origin) != len(candidate) {
		return false
	}
	offset := 0
	if !strings.ContainsAny(candidate[:1], vowelsStr) {
		prefLen := getPrefixLen(origin)
		if prefLen >= 0 {
			offset = prefLen
		}
	}

	originRunes := []rune(origin)
	candidateRunes := []rune(candidate)

	for i := offset; i < len(originRunes); i++ {
		if originRunes[i] != candidateRunes[i-offset] {
			return false
		}
	}

	for i := 0; i < offset; i++ {
		if originRunes[i] != candidateRunes[len(candidateRunes)-offset+i] {
			return false
		}
	}

	return true
}

var (
	vowels = []byte{'a', 'e', 'i', 'u', 'o'}
)

func getPrefixLen(word string) int {
	foundVowel := false
	for i, r := range word {
		if !foundVowel && bytes.ContainsRune(vowels, r) {
			foundVowel = true
			continue
		}
		if foundVowel && !bytes.ContainsRune(vowels, r) {
			// found a consonant
			return i
		}
	}

	return -1
}
