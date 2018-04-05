package gormfk

import "testing"

func TestReduceModelToName(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceModelToName(tt.args.model); got != tt.want {
				t.Errorf("ReduceModelToName() = %v, want %v", got, tt.want)
			}
		})
	}
}
