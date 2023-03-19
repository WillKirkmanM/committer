package utilities

import (
	"fmt"
	"os"
)

func ParseCliArgs() {
	res := ValidateArgs()

	if res == true {
		fmt.Println("All Args are Correct!")
	}
}

func ValidateArgs() (result bool) {
	hasMessage := false
	hasScope := false

	for key, value := range ListArgs() {

		if key == "-m" || key == "--message" {
			hasMessage = true
		}
		if key == "-s" || key == "--scope" {
			hasScope = true
		}

		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
	return hasMessage == hasScope
}

func ListArgs() map[string]string {
	resultMap := make(map[string]string)
	ArgsArray := os.Args[1:]

	for index := range ArgsArray {
		firstCharacter := string([]rune(ArgsArray[index])[0:1])

		if firstCharacter == "-" {
			wordAfter := string([]rune(ArgsArray[index]))

			index = index + 1
			for word := range ArgsArray[index:] {
				firstCharacter := ArgsArray[index+word][0:1]
				if firstCharacter == "-" {
					break
				}
				// fmt.Println(ArgsArray[index:])
				resultMap[wordAfter] += ArgsArray[index:][word] + " "
			}
		}
	}
	return resultMap
}
