package main

//Letter Combinations of a Phone Number

//Approach 1: Iterative Cartesian Product -----------------------------------

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	strings := []string{}
	for _, digit := range digits {
		strings = append(strings, digitRep(digit))
	}

	result := []string{""}
	for _, s := range strings {
		temp := []string{}
		for _, prefix := range result {
			for _, char := range s {
				temp = append(temp, prefix+string(char))
			}
		}
		result = temp
	}

	return result
}

func digitRep(num rune) string {
	switch num {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}
