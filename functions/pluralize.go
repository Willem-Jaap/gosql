package functions

func pluralize(word string) string {
	if len(word) == 0 {
		return word
	}
	lastChar := word[len(word)-1]
	switch lastChar {
	case 's', 'x', 'z':
		return word + "es"
	case 'h':
		if len(word) > 1 && (word[len(word)-2] == 's' || word[len(word)-2] == 'c') {
			return word + "es"
		}
		return word + "s"
	case 'y':
		if len(word) > 1 && (word[len(word)-2] == 'a' || word[len(word)-2] == 'e' || word[len(word)-2] == 'i' || word[len(word)-2] == 'o' || word[len(word)-2] == 'u') {
			return word + "s"
		}
		return word[:len(word)-1] + "ies"
	default:
		return word + "s"
	}
}

func isPlural(word string) bool {
	if len(word) > 1 && word[len(word)-1] == 's' {
		// Most plural words end in "s"
		return true
	}
	if len(word) > 2 && word[len(word)-2:] == "es" {
		// Some plural words end in "es"
		return true
	}
	if len(word) > 2 && word[len(word)-2:] == "en" {
		// Some irregular plural words end in "en"
		return true
	}
	if len(word) > 3 && word[len(word)-3:] == "ies" {
		// Some singular words end in "y" and change to "ies" when pluralized
		return true
	}
	return false
}
