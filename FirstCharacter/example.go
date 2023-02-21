package example

//Given a string S consisting of lowercase Latin Letters, the task is to find the first non-repeating character in S.
//Input: “geeksforgeeks”
//Output: 5
//Explanation: As 5 is first character in the string which does not repeat.

func FirstNonRepeatingCharacterOld(str string) int {
	for i := range str {
		foundDuplicate := false
		for j := range str {
			if str[i] == str[j] && i != j {
				foundDuplicate = true
			}
		}

		if foundDuplicate != true {
			return i
		}
	}
	return -1
}

func FirstNonRepeatingCharacter(str string) int {

	charCountNoRepeat := make(map[string]int)

	for _, char := range str {
		charCountNoRepeat[string(char)] += 1
	}

	for i, char := range str {
		if charCountNoRepeat[string(char)] == 1 {
			return i
		}
	}

	return -1
}
