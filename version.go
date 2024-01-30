package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"

	k8s "k8s.io/apimachinery/pkg/version"
)

var (
	MajorString      string
	MinorString      string
	GitVersionString  string
	GitCommitString  string
	GitTreeStateString   string
	GitBranchString  string
	BuildDateString  string
)

var BuildInfo, HasBuildInfo = debug.ReadBuildInfo()

var Info = &Version{
	Major:      MajorString,
	Minor:      MinorString,
	GitVersion:  GitVersionString,
	GitCommit:  GitCommitString,
	GitTreeState:   GitTreeStateString,
	BuildDate:  BuildDateString,
	GoVersion:  runtime.Version(),
	Compiler:   runtime.Compiler,
	Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
}

func init() {
	if HasBuildInfo {
		var ModVersion = BuildInfo.Main.Version
		if ModVersion == "" {
			return
		}
		if GitVersionString == "" {
			GitVersionString = ModVersion
		}
	}
}

func GetVersion() *Version {
	return Info
}

// Version is modeled after k8s.io/apimachinery/pkg/version
type Version = k8s.Info

func Decode(data []byte) (*Version, error) {
	v := &Version{}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
