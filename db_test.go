package cheesedb

import (
	"reflect"
	"testing"
)

func TestNewDb(t *testing.T) {
	tests := []struct {
		name    string
		want    Db
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDb()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDb() got = %v, want %v", got, tt.want)
			}
		})
	}
}