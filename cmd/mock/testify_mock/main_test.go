package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDoraemon struct {
	mock.Mock
}

func mockNewDoraemon() *mockDoraemon {
	return &mockDoraemon{}
}

func (m *mockDoraemon) GetTakecopter() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockDoraemon) GetDokodemodoa() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestNobita_fly(t *testing.T) {
	tests := []struct {
		name      string
		mockValue string
		mockErr   error
		expected  string
	}{
		{
			name:      "Success Test",
			mockValue: "タケコプター🚁",
			mockErr:   nil,
			expected:  "タケコプター🚁を使って飛んでいく!",
		},
		{
			name:      "Fail Test",
			mockValue: "",
			mockErr:   assert.AnError,
			expected:  "飛べません🐧",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoraemon := mockNewDoraemon()
			mockDoraemon.On("GetTakecopter").Return(tt.mockValue, tt.mockErr)

			nobita := NewNobita(mockDoraemon)
			got := nobita.fly()

			assert.Equal(t, tt.expected, got)
			mockDoraemon.AssertExpectations(t)
		})
	}
}
