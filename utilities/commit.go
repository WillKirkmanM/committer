package utilities

import (
	"fmt"
)

func Commit() {
	commitType := GetCommitType()
	scope := GetScope()
	message := GetCommitMessage()

	fmt.Println(commitType, scope, message)
}
