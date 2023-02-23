package example

//Write a method to replace all the spaces in a string with ‘%20’.
//You may assume that the string has sufficient space at the end to hold the additional
//characters and that you are given the “true” length of the string.

func URLify(str string, trueLength int) string {
	var charReplace, newStr string

	for _, char := range str {
		if string(char) == " " {
			charReplace = "%20"
			trueLength = trueLength + 2
		} else {
			charReplace = string(char)
		}

		newStr = newStr + charReplace

		if len(newStr) == trueLength {
			return newStr
		}
	}

	return ""
}
