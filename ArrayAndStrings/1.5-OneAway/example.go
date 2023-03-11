package example

//One away: There are three types of edits that can be performed on strings: insert a character,
//remove a character, or replace a character.
//Give two strings, write a function to check if they are one edit (or zero edits) away

func OneAway(strOne, strTwo string) bool {

	letters := map[string]int{}
	count := 0
	minStr := ""

	if len(strOne) <= len(strTwo) {
		minStr = strOne
	} else {
		minStr = strTwo
	}

	for _, min := range minStr {
		letters[string(min)] += 1
	}

	for _, two := range strTwo {
		if _, ok := letters[string(two)]; !ok {
			count++
		}
	}

	if count <= 1 {
		return true
	}

	return false
}
