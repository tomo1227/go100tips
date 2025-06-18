package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDoraemon struct {
	errFunc string
}

func mockNewDoraemon(errFunc string) FourDimensionalPocket {
	return &mockDoraemon{errFunc: errFunc}
}

func (d *mockDoraemon) GetTakecopter() (string, error) {
	if d.errFunc == "GetTakecopter" {
		return "", assert.AnError
	}
	return "タケコプター🚁", nil
}

func (d *mockDoraemon) GetDokodemodoa() (string, error) {
	if d.errFunc == "GetDokodemodoa" {
		return "", assert.AnError
	}
	return "どこでもドア🚪", nil
}

func TestNobita_fly(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		errFunc string
	}{
		{
			name:    "Success Test",
			want:    "タケコプター🚁を使って飛んでいく!",
			errFunc: "",
		},
		{
			name:    "Fail Test",
			want:    "飛べません🐧",
			errFunc: "GetTakecopter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doraemon := mockNewDoraemon(tt.errFunc)
			nobita := NewNobita(doraemon)
			if got := nobita.fly(); got != tt.want {
				t.Errorf("Nobita.fly() = %v, want %v", got, tt.want)
			}
		})
	}
}
