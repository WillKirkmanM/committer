package utilities

import (
	"fmt"
	"strings"
)

func GetScope() string {
	fmt.Println("What is the Scope of this Commit? E.g: index.js")

	var scope string

	fmt.Scanln(&scope)

	scope = strings.TrimSpace(scope)

	fmt.Print("\033[H\033[2J")

	return scope
}
