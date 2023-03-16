package utilities

import (
	"fmt"
)

func GetScope() string {
	fmt.Println("What is the Scope of this Commit? E.g: index.js")

	var scope string

	fmt.Scanln(&scope)

	fmt.Print("\033[H\033[2J")

	return scope
}
