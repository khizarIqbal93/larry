package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kdawg500/larry/gitutils"
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
	// uvPtr := flag.Bool("uv", false, "option to update version in pyproject.toml in uv projects")
	flag.Parse()
	args := flag.Args()

	err := clipboard.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logs := gitutils.GitLogs{}
	tags := gitutils.GitTags{}
	commit := logs.LastCommitHash()
	tag := tags.LastTag()

	if args[0] == "version" {
		fmt.Println(tag.String())
		return
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

	switch action {
	case MAJOR:
		tag.IncrementMajor().AddGitTag(*messagePtr, commit)
	case MINOR:
		tag.IncrementMinor().AddGitTag(*messagePtr, commit)
	case PATCH:
		tag.IncrementPatch().AddGitTag(*messagePtr, commit)
	}

	color.Cyan("Tag created: %s \nfor commit: %s", tag.String(), commit)
	clipboard.Write(clipboard.FmtText, []byte(tag.String()))
}
