package main

import (
	"fmt"
	"os"

	"github.com/kdawg500/larry/parse"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: larry <message> <major|minor|patch>")
		os.Exit(1)
	}

	message := args[1]
	action := args[2]

	if action != "major" && action != "minor" && action != "patch" {
		fmt.Println("Usage: larry <message> <major|minor|patch>")
		os.Exit(1)
	}

	logs := parse.GitLogs{}
	tags := parse.GitTags{}
	commit := logs.LastCommitHash()
	tag := tags.LastTag()
	switch action {
	case "major":
		tag.IncrementMajor(message, commit)
	case "minor":
		tag.IncrementMinor(message, commit)
	case "patch":
		tag.IncrementPatch(message, commit)
	}

	fmt.Println(tag.String())
}
