package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kdawg500/larry/parse"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: larry increment <major|minor|patch> -m <message> (optional)")
		os.Exit(1)
	}

	messagePtr := flag.String("m", "", "message for the tag")
	flag.Parse()

	action := args[2]

	if action != "major" && action != "minor" && action != "patch" {
		fmt.Println("Usage: larry increment <major|minor|patch> -m <message> (optional)")
		os.Exit(1)
	}

	logs := parse.GitLogs{}
	tags := parse.GitTags{}
	commit := logs.LastCommitHash()
	tag := tags.LastTag()
	switch action {
	case "major":
		tag.IncrementMajor(*messagePtr, commit)
	case "minor":
		tag.IncrementMinor(*messagePtr, commit)
	case "patch":
		tag.IncrementPatch(*messagePtr, commit)
	}

	fmt.Printf("Tag created: %s\n for commit: %s\n", tag.String(), commit)
}
