package gitutils

import (
	"fmt"
	"os"
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

func (v *VersionTag) AddGitTag(message, commitHash string) {
	cmd := exec.Command("git", "tag", "-a", v.String(), commitHash, "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error creating tag: ", string(output))
		os.Exit(128)
	}
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

func ParseTag(tag string) VersionTag {
	res := VersionTag{}
	re := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if re.MatchString(tag) {
		segments := strings.Split(strings.TrimPrefix(tag, "v"), ".")
		var err error
		res.Major, err = strconv.Atoi(segments[0])
		if err != nil {
			panic(err)
		}
		res.Minor, err = strconv.Atoi(segments[1])
		if err != nil {
			panic(err)
		}
		res.Patch, err = strconv.Atoi(segments[2])
		if err != nil {
			panic(err)
		}
	}
	return res
}

// gets the tags from the git repository
func (g *GitTags) getTags() {
	cmd := exec.Command("git", "tag")
	bTags, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	g.rawTags = bTags
}

// returns the last tag
func (g *GitTags) LastTag() VersionTag {
	g.getTags()
	rawTagsStr := string(g.rawTags)
	tags := strings.Split(string(rawTagsStr), "\n")
	tags = tags[:len(tags)-1]
	if len(tags) == 0 {
		return VersionTag{
			Major: 0,
			Minor: 0,
			Patch: 0,
		}
	}
	sortedTags := tagSort(tags)
	lastTag := sortedTags[len(tags)-1]
	t := ParseTag(lastTag)
	return t
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
	tagA := ParseTag(a)
	tagB := ParseTag(b)

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
