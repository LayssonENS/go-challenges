package example

//Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.

func CheckPermutation(strOne, strTwo string) bool {

	if len(strOne) != len(strTwo) {
		return false
	}

	mapOne := make(map[string]int)

	for _, char := range strOne {
		mapOne[string(char)] += 1
	}

	for _, char := range strTwo {
		mapOne[string(char)] -= 1

		if mapOne[string(char)] < 0 {
			return false
		}

	}

	return true
}
