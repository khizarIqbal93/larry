package uv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/khizarIqbal93/larry/gitutils"
)

func parse(rawTOML string) (map[string]any, error) {
	var pyProjectTOML map[string]any
	if _, err := toml.Decode(rawTOML, &pyProjectTOML); err != nil {
		return nil, fmt.Errorf("failed to parse TOML: %w", err)
	}
	return pyProjectTOML, nil
}

func getUVVersion(parsedTOML map[string]any) (gitutils.VersionTag, error) {
	project, ok := parsedTOML["project"].(map[string]any)
	if !ok {
		return gitutils.VersionTag{}, fmt.Errorf("missing [project] table in pyproject.toml")
	}
	version, ok := project["version"].(string)
	if !ok {
		return gitutils.VersionTag{}, fmt.Errorf("missing version field in [project]")
	}
	return gitutils.ParseTag("v" + version)
}

func UpdateUvVersion(rawTOML, updatedTag string) string {
	newVersion := strings.TrimPrefix(updatedTag, "v")
	re := regexp.MustCompile(`(?m)^(version\s*=\s*")([^"]+)(")`)
	updatedTOML := re.ReplaceAllString(rawTOML, "${1}"+newVersion+"${3}")
	return updatedTOML
}
