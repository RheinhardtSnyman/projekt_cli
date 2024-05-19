package main

import (
	"reflect"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		s       string
		want    picSize
		wantErr bool
	}{
		{"500x500", picSize{500, 500}, false},
		{"5x500", picSize{5, 500}, false},
		{"500-500", picSize{}, true},
		{"50Ax500", picSize{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := parseSize(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSoze(%s) error = %v, wantErr %v", tt.s, err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSize(%s) = %v, wantErr %v", tt.s, got, tt.want)
			}
		})
	}
}

func Test_useFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		filename string
		want     bool
	}{
		{"a/b/file.JPG", true},
		{"a/b/file.jpeg", true},
		{"a/b/file.jeg", false},
		{"a/b/file", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := useFile(tt.filename); got != tt.want {
				t.Errorf("useFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
