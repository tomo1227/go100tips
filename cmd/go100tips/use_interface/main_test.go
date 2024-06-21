package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDoraemon struct {
	want bool
}

func mockNewDoraemon(want bool) *mockDoraemon {
	return &mockDoraemon{want: want}
}

func mockNewNobita() FourDimensionalPocket {
	return mockDoraemon{}
}

func (d mockDoraemon) getItem() (string, error) {
	if d.want {
		return "", assert.AnError
	}
	return "暗記パン🍞", nil
}

func TestNobita_study(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Success Test",
			want:    "暗記パン🍞を使って勉強できた!",
			wantErr: false,
		},
		{
			name:    "Fail Test",
			want:    "勉強できないよ〜😭",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doraemon := mockNewDoraemon(tt.wantErr)
			nobita := NewNobita(doraemon)
			if got := nobita.study(); got != tt.want {
				t.Errorf("Nobita.study() = %v, want %v", got, tt.want)
			}
		})
	}
}
