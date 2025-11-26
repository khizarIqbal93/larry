package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kdawg500/larry/parse"
	"golang.design/x/clipboard"
)

//go:embed larry.txt
var larryAscii string

const (
	MAJOR = "major"
	MINOR = "minor"
	PATCH = "patch"
)

func main() {
	messagePtr := flag.String("m", "", "message for the tag")
	flag.Parse()
	args := flag.Args()

	err := clipboard.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(args) < 2 {
		fmt.Println(larryAscii)
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Usage: larry -m <message> (optional) increment <major|minor|patch>")
		os.Exit(1)
	}

	action := args[1]
	if action != MAJOR && action != MINOR && action != PATCH {
		fmt.Println("Usage: larry -m <message> (optional) increment <major|minor|patch> ")
		os.Exit(1)
	}

	logs := parse.GitLogs{}
	tags := parse.GitTags{}
	commit := logs.LastCommitHash()
	tag := tags.LastTag()
	switch action {
	case MAJOR:
		tag.IncrementMajor(*messagePtr, commit)
	case MINOR:
		tag.IncrementMinor(*messagePtr, commit)
	case PATCH:
		tag.IncrementPatch(*messagePtr, commit)
	}

	color.Cyan("Tag created: %s \nfor commit: %s", tag.String(), commit)
	clipboard.Write(clipboard.FmtText, []byte(commit))
}
