package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
	} else {
		wordSlice, charCount, characters := processInput(args)
		aL := getVowels(characters, charCount)
		newSlice := mirrorVowels(wordSlice, aL)
		for _, v := range newSlice {
			for _, e := range v {
				z01.PrintRune(e)
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func processInput(args []string) ([]string, int, string) {
	result := []string{}
	characters := ""
	word := ""
	wordCount := 0
	for _, v := range args {
		for i, e := range v {
			if e != ' ' && i < len(v) {
				word += string(e)
				characters += string(e)
				wordCount++
				continue
			} else {
				result = append(result, word)
				word = ""
			}
		}
		result = append(result, word)
		word = ""
	}
	return result, wordCount, characters
}

func getVowels(s string, charCount int) []string {
	a := []string{}
	var loopCount int = charCount
	for i := 0; i < loopCount; i++ {
		if isVowel(string(s[i])) {
			a = append(a, string(s[i]))
		}
	}
	return a
}

func isVowel(c string) bool {
	vowels := "AEIOUaeiou"
	for _, v := range vowels {
		if c == string(v) {
			return true
		}
	}
	return false
}

func mirrorVowels(wS []string, aL []string) []string {
	mirrorCharCount := len(aL) - 1
	newWordSlice := wS
	counter := 0
	if mirrorCharCount > 0 {
		for i, v := range wS {
			currentWord := ""
			for _, e := range v {
				if isVowel(string(e)) {
					currentWord += string(aL[mirrorCharCount-counter])
					counter++
				} else {
					currentWord += string(e)
				}
			}
			newWordSlice[i] = currentWord
		}
	}
	return newWordSlice
}
