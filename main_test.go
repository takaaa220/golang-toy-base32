package main

import (
	"encoding/base32"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input []byte
	}{
		{
			input: []byte("hello world"),
		},
		{
			input: []byte("hello 世界"),
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			if got, expected := Main(tt.input), base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(tt.input); got != expected {
				t.Errorf("Main() = %v, want %v", got, expected)
			}
		})
	}
}
