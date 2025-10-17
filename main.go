package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/kdawg500/larry/parse"
)

//go:embed larry.txt
var larryAscii string

func main() {
	messagePtr := flag.String("m", "", "message for the tag")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println(larryAscii)
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Usage: larry -m <message> (optional) increment <major|minor|patch>")
		os.Exit(1)
	}

	action := args[1]
	if action != "major" && action != "minor" && action != "patch" {
		fmt.Println("Usage: larry -m <message> (optional) increment <major|minor|patch> ")
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
