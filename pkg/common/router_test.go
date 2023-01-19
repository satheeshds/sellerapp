package common

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		routes Routes
	}
	tests := []struct {
		name string
		args args
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.routes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
