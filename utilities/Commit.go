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

	fmt.Println("Formatted Message:")
	fmt.Printf("%s(%s): %s\n", commitType, scope, message)
}
