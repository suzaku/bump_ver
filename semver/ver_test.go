package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("should return error for invalid versions", func(t *testing.T) {
		badVersions := []string{
			"1.-3", "whatever", "asdf.123.1a",
		}
		for _, v := range badVersions {
			_, err := Parse(v)
			assert.NotNil(t, err)
		}
	})
	t.Run("should parse versions correctly", func(t *testing.T) {
		cases := []struct {
			version  string
			expected Ver
		}{
			{
				"1.2.0",
				Ver{1, 2, 0},
			},
			{
				"3.0.0",
				Ver{3, 0, 0},
			},
			{
				"0.3.12",
				Ver{0, 3, 12},
			},
		}
		for _, c := range cases {
			ver, err := Parse(c.version)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, ver)
		}
	})
}

func TestString(t *testing.T) {
	ver := Ver{
		major: 1,
		minor: 2,
		patch: 34,
	}
	verS := fmt.Sprintf("%s", ver)
	assert.Equal(t, "1.2.34", verS)
}
