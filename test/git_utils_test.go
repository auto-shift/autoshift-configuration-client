package utils

import (
	"testing"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

func TestGitClone(t *testing.T) {
	type args struct {
		gitUser string
		gitPass string
		gitUrl  string
		gitDir  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			//test paramters
			//e.g.
			//name: <name>
			//args: args{arg1:"", arg2: "", etc }
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.GitClone(tt.args.gitUser, tt.args.gitPass, tt.args.gitUrl, tt.args.gitDir)
		})
	}
}
