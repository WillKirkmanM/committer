package utilities

import (
	"fmt"
	"os"
)

func ParseCliArgs() (messsage string, scope string, error error) {
	if len(os.Args) <= 1 {
		// commitType := GetCommitType()
		scope := GetScope()
		message := GetCommitMessage()

		return message, scope, nil
	}

	if ValidateArgs() == true {
		fmt.Println("All Args are Correct!")
		args := ListArgs()

		var message string
		var scope string

		if v, ok := args["-m"]; ok {
			message = v
		}
		if v, ok := args["--message"]; ok {
			message = v
		}

		if v, ok := args["-s"]; ok {
			scope = v
		}
		if v, ok := args["--scope"]; ok {
			scope = v
		}

		return message, scope, nil
	}
	return "", "", fmt.Errorf("Missing Arguments!")
}

func ValidateArgs() (result bool) {
	hasMessage := false
	hasScope := false

	for key := range ListArgs() {

		if key == "-m" || key == "--message" {
			hasMessage = true
		}
		if key == "-s" || key == "--scope" {
			hasScope = true
		}
	}
	return hasMessage == true && hasScope == true
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
