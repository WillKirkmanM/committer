package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func GetCommitMessage() string {
	fmt.Println("What is your Commit Message?")

	var commitMessage string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		commitMessage = scanner.Text()
	}

	fmt.Print("\033[H\033[2J")

	return commitMessage
}
