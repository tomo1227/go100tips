package main

import (
	"testing"
)

func TestNobita_study(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Success Test",
			want: "暗記パン🍞を使って勉強できた!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nobita := NewNobita()
			result := nobita.study()
			if result != tt.want {
				t.Errorf("expected %s, got %s", tt.want, result)
			}
		})
	}
}
