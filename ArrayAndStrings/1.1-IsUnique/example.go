package example

//Is unique : implement an algorithm to determine if a string has all unique character. What if you cannot use additional
//data structures?

func IsUnique(str string) bool {

	unique := make(map[string]int)

	for _, char := range str {
		unique[string(char)] += 1
	}

	for _, char := range str {
		if unique[string(char)] == 1 {
			return true
		}
	}

	return false
}
