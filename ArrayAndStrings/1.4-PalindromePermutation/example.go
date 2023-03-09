package example

import "strings"

// Given a string, write a function to check if it is a permutation of a palindrome.
//  A palindrome does not to be limited to just dictionary words.

func IsPermutationOfPalindrome(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	letters := map[string]bool{}

	for _, s := range str {

		if _, ok := letters[string(s)]; ok {
			delete(letters, string(s))
		} else {
			letters[string(s)] = true
		}
	}

	return len(letters) <= 1
}
