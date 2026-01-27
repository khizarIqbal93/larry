package uv

import (
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

func UpdateUvVersion(rawTOML, increment string) string {
	parsed := parse(rawTOML)
	tag := getUVVersion(parsed)
	switch increment {
	case "major":
		tag.IncrementMajor()
	case "minor":
		tag.IncrementMinor()
	case "patch":
		tag.IncrementPatch()
	}

	parsed["project"].(map[string]any)["version"] = tag.String()
	updatedToml, err := toml.Marshal(parsed)
	if err != nil {
		panic(err.Error())
	}

	return string(updatedToml)

}
