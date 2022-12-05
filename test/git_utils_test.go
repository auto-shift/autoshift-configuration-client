package utils

import (
	"testing"

	"github.com/auto-shift/autoshift-configuration-client/cmd/internal/utils"
)

func TestGitClone(t *testing.T) {
	type args struct {
		gitUser string
		gitPass string
		gitRepo string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.GitClone(tt.args.gitUser, tt.args.gitPass, tt.args.gitRepo)
		})
	}
}
