package main

import "testing"

func TestValidateUrl(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"https://golang.org", true},
		{"http://golang.org", true},
		{"https://golang.org/abc", true},
		{"https://golang.org7abc?id=123", true},
		{"http//golang.org", false},
		{"golang.org", false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := validateUrl(tt.s); got != tt.want {
				t.Errorf("validateUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
