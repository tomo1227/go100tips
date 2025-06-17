package main_test

import (
	"errors"
	main "go100tips/cmd/mock/matryer_moq"
	"testing"
)

func TestNobita_fly(t *testing.T) {
	tests := []struct {
		name           string
		expectedOutput string
		setupMock      func(*FourDimensionalPocketMock)
	}{
		{
			name:           "成功時",
			expectedOutput: "タケコプター🚁を使って飛んでいく!",
			setupMock: func(m *FourDimensionalPocketMock) {
				m.GetTakecopterFunc = func() (string, error) {
					return "タケコプター🚁", nil
				}
			},
		},
		{
			name:           "失敗時",
			expectedOutput: "飛べません🐧",
			setupMock: func(m *FourDimensionalPocketMock) {
				m.GetTakecopterFunc = func() (string, error) {
					return "", errors.New("error")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPocket := &FourDimensionalPocketMock{}
			tt.setupMock(mockPocket)

			nobita := main.NewNobita(mockPocket)
			got := nobita.Fly()

			if got != tt.expectedOutput {
				t.Errorf("got = %s, want = %s", got, tt.expectedOutput)
			}
		})
	}
}
