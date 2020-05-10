package version

import (
	"testing"
)

func TestGet(t *testing.T) {
	versionSpec := Get()

	if versionSpec.Label != "version_test_label" {
		t.Errorf("Expected version_test_label, got %s", versionSpec.Label)
	}

	if versionSpec.GitCommit != "version_git_commit" {
		t.Errorf("Expected version_git_commit, got %s", versionSpec.GitCommit)
	}
}
