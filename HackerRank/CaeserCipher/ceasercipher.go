package caesercipher

import (
	"unicode"
)

func CaesarCipher(s string, k int32) string {
	// 65-90 (A-Z)
	// 97-122 (a-z)

	result := []rune{}

	realSwap := k % 26

	for _, letter := range s {
		if unicode.IsUpper(letter) {
			letter = letter + realSwap

			if letter > 90 {
				letter = letter - 26
			}
		}

		if unicode.IsLower(letter) {
			letter = letter + realSwap

			if letter > 122 {
				letter = letter - 26
			}
		}

		result = append(result, letter)
	}

	return string(result)
}
