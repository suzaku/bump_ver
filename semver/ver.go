package semver

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	verPat = regexp.MustCompile(`(?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)`)
)

type Ver struct {
	major, minor, patch int
}

func Parse(verStr string) (Ver, error) {
	matches := verPat.FindStringSubmatch(verStr)
	if len(matches) != 4 {
		return Ver{}, fmt.Errorf("unknown pattern of semver: %q", verStr)
	}
	matchedVals := make(map[string]string, 3)
	for i, name := range verPat.SubexpNames() {
		if i == 0 {
			continue
		}
		matchedVals[name] = matches[i]
	}
	major, _ := strconv.Atoi(matchedVals["major"])
	minor, _ := strconv.Atoi(matchedVals["minor"])
	patch, _ := strconv.Atoi(matchedVals["patch"])
	return Ver{
		major: major,
		minor: minor,
		patch: patch,
	}, nil
}

func (v Ver) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func (v *Ver) BumpMajor() {
	v.major += 1
	v.minor = 0
	v.patch = 0
}

func (v *Ver) BumpMinor() {
	v.minor += 1
	v.patch = 0
}

func (v *Ver) BumpPatch() {
	v.patch += 1
}
