package utilities

import (
	"fmt"
)

func GetCommitType() string {
	fmt.Println("Which Commit is it?")

	fmt.Println("1: Feature")
	fmt.Println("2: Documentation")
	fmt.Println("3: Bug Fix")
	fmt.Println("4: Refactor")

	var commitType int

	fmt.Scan(&commitType)

	fmt.Print("\033[H\033[2J")

	switch commitType {
	case 1:
		return "Feat"
	case 2:
		return "Docs"
	case 3:
		return "Fix"
	case 4:
		return "Refactor"
	default:
		GetCommitType()
	}
	return "" // If this triggers bruh
}
