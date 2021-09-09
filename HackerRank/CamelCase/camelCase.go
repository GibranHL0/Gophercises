package camelcase

import "unicode"

func CamelCase(s string) int32 {
	var counter int32 = 1

	if len(s) < 1 {
		return 0
	}

	for _, letter := range s {
		if unicode.IsUpper(letter){
			counter ++
		}
	}

	return counter
}