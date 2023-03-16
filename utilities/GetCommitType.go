package utilities

import (
	"fmt"
)

func GetCommitType() int {
	fmt.Println("Which Commit is it?")

	fmt.Println("1: Feature")
	fmt.Println("2: Documentation")
	fmt.Println("3: Bug Fix")
	fmt.Println("4: Refactor")

	var commitType int

	fmt.Scan(&commitType)

	fmt.Print("\033[H\033[2J")

	return commitType
}
