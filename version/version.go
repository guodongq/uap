package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	gitVersion   = "v0.0.0"
	gitCommit    = "unknown"
	gitTreeState = "unknown"
	buildDate    = "unknown"
	gitMajor     = "unknown"
	gitMinor     = "unknown"
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitMajor     string `json:"gitMajor"`
	GitMinor     string `json:"gitMinor"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	jsonString, _ := json.Marshal(info)
	return string(jsonString)
}

func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitMajor:     gitMajor,
		GitMinor:     gitMinor,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func SetGitVersion(v string) {
	gitVersion = v
}
