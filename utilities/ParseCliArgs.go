package utilities

import (
	"fmt"
	"os"
)

func ParseCliArgs() {
	for key, value := range ListArgs() {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
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
