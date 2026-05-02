package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/khizarIqbal93/larry/gitutils"
	"github.com/khizarIqbal93/larry/uv"
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
	uvPtr := flag.Bool("uv", false, "option to update version in pyproject.toml in uv projects")
	flag.Parse()

	args := flag.Args()

	err := clipboard.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logs := gitutils.GitLogs{}
	tags := gitutils.GitTags{}

	commit, err := logs.LastCommitHash()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	tag, err := tags.LastTag()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println(larryAscii)
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Usage: larry -m <message> (optional) increment <major|minor|patch>")
		os.Exit(1)
	}

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
		err = tag.IncrementMajor().AddGitTag(*messagePtr, commit)
	case MINOR:
		err = tag.IncrementMinor().AddGitTag(*messagePtr, commit)
	case PATCH:
		err = tag.IncrementPatch().AddGitTag(*messagePtr, commit)
	}
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	if *uvPtr {
		pyprojectPath := "pyproject.toml"
		if _, err := os.Stat(pyprojectPath); os.IsNotExist(err) {
			color.Red("pyproject.toml not found in current directory")
			os.Exit(1)
		}

		content, err := os.ReadFile(pyprojectPath)
		if err != nil {
			color.Red("Failed to read pyproject.toml: %s", err.Error())
			os.Exit(1)
		}

		updatedContent := uv.UpdateUvVersion(string(content), tag.String())
		err = os.WriteFile(pyprojectPath, []byte(updatedContent), 0644)
		if err != nil {
			color.Red("Failed to write pyproject.toml: %s", err.Error())
			os.Exit(1)
		}

		color.Green("Updated version in pyproject.toml")
	}

	color.Cyan("Tag created: %s \nfor commit: %s", tag.String(), commit)
	clipboard.Write(clipboard.FmtText, []byte(tag.String()))
}
