package impls

import "testing"

func TestLocalRepo(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocalRepo(); got != tt.want {
				t.Errorf("LocalRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
