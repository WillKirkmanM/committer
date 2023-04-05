package utilities

import (
	"fmt"
	"log"
)

func Commit() {
	commitType, message, scope, err := ParseCliArgs()
	if err != nil {
		log.Fatal(err)
	}

	commitMessage := fmt.Sprintf("%s(%s): %s\n", commitType, scope, message)
	fmt.Println(commitMessage)

	AddFile(scope)
	GitCommit(message)
	Confirm("push the files")
	PushFile()
}
