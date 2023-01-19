package common

import (
	"net/http"
	"testing"
)

func TestReadBody(t *testing.T) {
	type args struct {
		r     *http.Request
		model any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadBody(tt.args.r, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("ReadBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadIdFromRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadIdFromRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadIdFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadIdFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
