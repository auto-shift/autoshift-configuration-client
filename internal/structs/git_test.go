package structs

import (
	"reflect"
	"testing"
)

func TestGitVars_ReadGitConfigs(t *testing.T) {
	type fields struct {
		gMap map[string]GVars
	}
	tests := []struct {
		name   string
		fields fields
		want   GitVars
	}{
		// TODO: Add test cases.
		{
			name:   "test1",
			fields: fields{},
			want: GitVars{
				GMap: map[string]GVars{"gitVars": {
					GitDir:  "Not Set",
					GitUrl:  "Not Set",
					GitUser: "Not Set",
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gVars := GitVars{
				GMap: tt.fields.gMap,
			}
			if got := gVars.ReadGitConfigs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitVars.ReadGitConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitVars_WriteGitConfigs(t *testing.T) {
	type fields struct {
		GMap GitMap
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				GMap: map[string]GVars{"gitVars": {
					GitDir:  "Test Dir",
					GitUrl:  "Test URL",
					GitUser: "Test User",
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gVars := GitVars{
				GMap: tt.fields.GMap,
			}
			gVars.WriteGitConfigs()
		})
	}
}

func TestGitVars_UpdateGitVars(t *testing.T) {
	type fields struct {
		GMap GitMap
	}
	type args struct {
		gVars GVars
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				GMap: map[string]GVars{
					"gitVars": {
						GitDir:  "Not Set",
						GitUrl:  "Not Set",
						GitUser: "Not Set",
					},
				},
			},
			args: args{
				GVars{
					GitDir:  "Not Set",
					GitUrl:  "Not Set",
					GitUser: "Not Set",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars := GitVars{
				GMap: tt.fields.GMap,
			}
			vars.UpdateGitVars(tt.args.gVars)
		})
	}
}
