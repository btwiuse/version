package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	k8s "k8s.io/apimachinery/pkg/version"
)

var (
	MajorString        string
	MinorString        string
	GitVersionString   string
	GitCommitString    string
	GitTreeStateString string
	GitBranchString    string
	BuildDateString    string
)

var Info = &Version{
	Major:        MajorString,
	Minor:        MinorString,
	GitVersion:   GitVersionString,
	GitCommit:    GitCommitString,
	GitTreeState: GitTreeStateString,
	BuildDate:    BuildDateString,
	GoVersion:    runtime.Version(),
	Compiler:     runtime.Compiler,
	Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
}

func GetVersion() *Version {
	return Info
}

// Version is alias to k8s.io/apimachinery/pkg/version.Info
type Version = k8s.Info

func Decode(data []byte) (*Version, error) {
	v := &Version{}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
