package example

import (
	"strconv"
)

// Implemented a method to perform basic string compression using the counts of repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3. If the compressed string would not become
// smaller than the original string, your method should return the original string.
// You can assume the string has only uppercase and lowercase

func StringCompression(s string) string {

	var strCompression, firstInsert string
	var counterFirst int

	for i, str := range s {

		if firstInsert != string(str) && firstInsert != "" {
			strCompression = strCompression + firstInsert + strconv.Itoa(counterFirst)
			firstInsert = string(str)
			counterFirst = 0
		}

		if firstInsert == string(str) || firstInsert == "" {
			counterFirst++
			firstInsert = string(str)

			if len(s)-1 == i {
				strCompression = strCompression + firstInsert + strconv.Itoa(counterFirst)
			}
			continue
		}

	}

	if len(strCompression) >= len(s) {
		return s
	}

	return strCompression
}
