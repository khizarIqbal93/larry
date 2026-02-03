package uv

import (
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kdawg500/larry/gitutils"
)

func parse(rawTOML string) map[string]any {
	var pyProjectTOML map[string]any
	// var pyProjectTOML PyProjectTOML
	_, err := toml.Decode(rawTOML, &pyProjectTOML)
	if err != nil {
		panic(err.Error())
	}

	return pyProjectTOML
}

func getUVVersion(parsedTOML map[string]any) gitutils.VersionTag {
	version := parsedTOML["project"].(map[string]any)["version"].(string)
	version = "v" + version
	tag := gitutils.ParseTag(version)
	return tag
}

func UpdateUvVersion(rawTOML, updatedTag string) string {
	newVersion := strings.TrimPrefix(updatedTag, "v")
	re := regexp.MustCompile(`(?m)^(version\s*=\s*")([^"]+)(")`)
	updatedTOML := re.ReplaceAllString(rawTOML, "${1}"+newVersion+"${3}")
	return updatedTOML
}
