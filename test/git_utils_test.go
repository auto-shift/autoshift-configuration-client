package git_utils

import (
	"testing"

	"github.com/auto-shift/autoshift-configuration-client/cmd/internal/utils"
)

func TestCloneDir(t *testing.T) {
	type args struct {
		repo   string
		branch string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.CloneDir(tt.args.repo, tt.args.branch)
		})
	}
}

func TestGitPull(t *testing.T) {
	type args struct {
		repo   string
		branch string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.GitPull(tt.args.repo, tt.args.branch)
		})
	}
}

func TestGitPush(t *testing.T) {
	type args struct {
		repo   string
		branch string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.GitPush(tt.args.repo, tt.args.branch)
		})
	}
}
