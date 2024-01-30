package version

import (
	"testing"
)

func check(t *testing.T, name string, got string, expected string) bool {
	if got != expected {
		t.Errorf("check %s failed: got: %q, expected %q\n", name, got, expected)
	}
	return true
}

func TestDecode(t *testing.T) {
	inputs := []string{
		`{"GitCommit":"deadbeef","GitTreeState":"dirty","GitVersion":"0.0.1","BuildDate":"now"}`,
	}
	expects := [][]func(v *Version) bool{
		{
			func(v *Version) bool { return check(t, "GitCommit", v.GitCommit, "deadbeef") },
			func(v *Version) bool { return check(t, "GitTreeState", v.GitTreeState, "dirty") },
			func(v *Version) bool { return check(t, "GitVersion", v.GitVersion, "0.0.1") },
			func(v *Version) bool { return check(t, "BuildDate", v.BuildDate, "now") },
		},
	}
	for i, input := range inputs {
		var (
			v, err   = Decode([]byte(input))
			expected = expects[i]
		)
		if err != nil {
			t.Error(err)
		}
		for _, check := range expected {
			check(v)
		}
	}
}
