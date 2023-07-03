package main

import "fmt"

func main() {
	string := "Hello, World!"
	vowels := "aeiou"
	count := 0
	for _, char := range string {
		if isVowel(char, vowels) {
			count++
		}
	}
	fmt.Println("Количество гласных букв:", count)
}

func isVowel(char rune, vowels string) bool {
	for _, vowel := range vowels {
		if char == vowel {
			return true
		}
	}
	return false
}
