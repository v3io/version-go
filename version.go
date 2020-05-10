package version

import (
	"fmt"
)

type Info struct {
	Label     string `json:"label"`
	GitCommit string `json:"git_commit"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// these global variables are initialized by the build process if the build target
// is a standalone binary
var (
	label     = "latest"
	gitCommit = ""
	os        = ""
	arch      = ""
)

// Get returns the version information
func Get() *Info {

	// return the info initialized by the linker during build
	return &Info{
		Label:     label,
		GitCommit: gitCommit,
		OS:        os,
		Arch:      arch,
	}
}

// Set will update the stored version info, used primarily for tests
func Set(info *Info) {
	label = info.Label
	gitCommit = info.GitCommit
	os = info.OS
	arch = info.Arch
}

// String will print out the info
func (s *Info) String() string {
	return fmt.Sprintf("Label: %s, Git commit: %s, OS: %s, Arch: %s",
		s.Label,
		s.GitCommit,
		s.OS,
		s.Arch)
}
