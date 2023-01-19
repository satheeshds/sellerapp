package common

import (
	"net/http"
	"testing"
)

func TestWriteErrorResponse(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		statusCode int
		err        error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteErrorResponse(tt.args.w, tt.args.statusCode, tt.args.err)
		})
	}
}

func TestWriteSuccessResponse(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		statusCode int
		model      any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteSuccessResponse(tt.args.w, tt.args.statusCode, tt.args.model)
		})
	}
}

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
