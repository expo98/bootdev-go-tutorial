package main

func isValidPassword(password string) bool {
	// ?

	var safePasswordLength bool
	var upperCasePasswordCharacter bool
	var digitPasswordCharacter bool

	if 5 <= len(password) && len(password) <= 12 && safePasswordLength == false {
		safePasswordLength = true
	} else {
		return false
	}

	for _, character := range password {
		if 'A' <= character && character <= 'Z' && upperCasePasswordCharacter == false {
			upperCasePasswordCharacter = true
		}

		if '0' <= character && character <= '9' && digitPasswordCharacter == false {
			digitPasswordCharacter = true
		}

		if digitPasswordCharacter && upperCasePasswordCharacter {
			return true
		}
	}

	return false
}
