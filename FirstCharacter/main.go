package main

import "fmt"

func main() {
	str := "asdasdasdsadsadcsad"
	teste := FirstNonRepeatingCharacter(str)
	fmt.Print(teste)
}

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
	characterCounts := map[string]string{}

	for _, char := range str {
		characterCounts[string(char)] += "a"
	}

	for i, char := range str {
		if characterCounts[string(char)] == "hg" {
			return i
		}
	}
	return -1
}
