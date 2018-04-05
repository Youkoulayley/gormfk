package gormfk

import (
	"testing"

	"github.com/jinzhu/gorm"
)

type Conf struct {
	Name string
}

func TestReduceModelToName(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test with working model",
			args{
				&Conf{},
			},
			"conf",
		},
		{
			"Test with empty model",
			args{
				"",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReduceModelToName(tt.args.model)
			if got != tt.want {
				t.Errorf("ReduceModelToName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMany2ManyFIndex(t *testing.T) {
	type args struct {
		db          *gorm.DB
		parentModel interface{}
		childModel  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// Todo : set a DB
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Many2ManyFIndex(tt.args.db, tt.args.parentModel, tt.args.childModel)
		})
	}
}
