package gitutils

import (
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// struct for the git tags
type GitTags struct {
	rawTags []byte
}

type VersionTag struct {
	Major int
	Minor int
	Patch int
}

func (v *VersionTag) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *VersionTag) AddGitTag(message, commitHash string) error {
	cmd := exec.Command("git", "tag", "-a", v.String(), commitHash, "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create tag: %s", strings.TrimSpace(string(output)))
	}
	return nil
}

func (v *VersionTag) IncrementMajor() *VersionTag {
	v.Major = v.Major + 1
	v.Minor = 0
	v.Patch = 0
	return v
}

func (v *VersionTag) IncrementMinor() *VersionTag {
	v.Minor++
	v.Patch = 0
	return v
}

func (v *VersionTag) IncrementPatch() *VersionTag {
	v.Patch++
	return v
}

func ParseTag(tag string) (VersionTag, error) {
	re := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if !re.MatchString(tag) {
		return VersionTag{}, fmt.Errorf("invalid tag format: %q", tag)
	}
	segments := strings.Split(strings.TrimPrefix(tag, "v"), ".")
	major, err := strconv.Atoi(segments[0])
	if err != nil {
		return VersionTag{}, err
	}
	minor, err := strconv.Atoi(segments[1])
	if err != nil {
		return VersionTag{}, err
	}
	patch, err := strconv.Atoi(segments[2])
	if err != nil {
		return VersionTag{}, err
	}
	return VersionTag{Major: major, Minor: minor, Patch: patch}, nil
}

// gets the tags from the git repository
func (g *GitTags) getTags() error {
	cmd := exec.Command("git", "tag")
	bTags, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git tag failed: %w", err)
	}
	g.rawTags = bTags
	return nil
}

// returns the last tag
func (g *GitTags) LastTag() (VersionTag, error) {
	if err := g.getTags(); err != nil {
		return VersionTag{}, err
	}
	raw := strings.TrimSpace(string(g.rawTags))
	if raw == "" {
		return VersionTag{}, nil
	}
	tags := strings.Split(raw, "\n")
	var valid []string
	for _, t := range tags {
		if _, err := ParseTag(t); err == nil {
			valid = append(valid, t)
		}
	}
	if len(valid) == 0 {
		return VersionTag{}, nil
	}
	sorted := tagSort(valid)
	return ParseTag(sorted[len(sorted)-1])
}

// sorts tags in ascending order
//
// v0.1.3, v0.1.4, v0.1.5 ...
func tagSort(tags []string) []string {
	slices.SortStableFunc(tags, func(a, b string) int {
		return compare(a, b)
	})
	return tags
}

// Compares tag a and b. This adheres to semantic versioning vX.X.X, where v0.0.11 > v0.0.10
//
// Returns 1 if a > b otherwise returns -1
func compare(a, b string) int {
	tagA, _ := ParseTag(a)
	tagB, _ := ParseTag(b)

	if tagA.Major > tagB.Major {
		return 1
	}

	if tagA.Major == tagB.Major && tagA.Minor > tagB.Minor {
		return 1
	}

	if tagA.Minor == tagB.Minor && tagA.Patch > tagB.Patch {
		return 1
	}

	return -1
}
