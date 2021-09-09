package main

import (
	"fmt"

	caesercipher "github.com/GibranHL0/Gophercises/HackerRank/CaeserCipher"
	camelcase "github.com/GibranHL0/Gophercises/HackerRank/CamelCase"
)

func main() {
	example := "saveChangesInTheEditor"

	fmt.Println(camelcase.CamelCase(example))

	fmt.Println(caesercipher.CaesarCipher("Ciphering.", 26))

}
