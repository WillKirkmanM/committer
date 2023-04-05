package utilities

import (
	"fmt"
	"os"
	"strings"
)

func ParseCliArgs() (commitType string, messsage string, scope string, error error) {
	if len(os.Args) <= 1 {
		commitType := GetCommitType()
		scope := GetScope()
		message := GetCommitMessage()

		return commitType, message, scope, nil
	}

	if ValidateArgs() == true {
		args := ListArgs()

		var commitType string
		var message string
		var scope string

		if v, ok := args["-t"]; ok {
			commitType = strings.TrimSpace(v)
		}
		if v, ok := args["--type"]; ok {
			commitType = strings.TrimSpace(v)
		}

		if v, ok := args["-m"]; ok {
			message = v
		}
		if v, ok := args["--message"]; ok {
			message = v
		}

		if v, ok := args["-s"]; ok {
			scope = strings.TrimSpace(v)
		}
		if v, ok := args["--scope"]; ok {
			scope = strings.TrimSpace(v)
		}

		return commitType, message, scope, nil
	}
	return "", "", "", fmt.Errorf("Missing Arguments!") // TODO: Show what arguments need to be provided
}

func ValidateArgs() (result bool) {
	hasMessage := false
	hasScope := false

	for key := range ListArgs() {
		switch key {
		case "-m":
			hasMessage = true
		case "--message":
			hasMessage = true

		case "-s":
			hasScope = true
		case "--scope":
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
