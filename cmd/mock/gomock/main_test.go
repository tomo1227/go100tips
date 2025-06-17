package main

import (
	"go100tips/cmd/mock/gomock/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNobitaFly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPocket := mock.NewMockFourDimensionalPocket(ctrl)

	tests := []struct {
		name        string
		mockReturn  string
		mockError   error
		expectedFly string
	}{
		{
			name:        "正常なケース",
			mockReturn:  "タケコプター🚁",
			mockError:   nil,
			expectedFly: "タケコプター🚁を使って飛んでいく!",
		},
		{
			name:        "エラーケース",
			mockReturn:  "",
			mockError:   assert.AnError,
			expectedFly: "飛べません🐧",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPocket.EXPECT().GetTakecopter().Return(tt.mockReturn, tt.mockError).Times(1)
			nobita := NewNobita(mockPocket)
			assert.Equal(t, tt.expectedFly, nobita.fly())
		})
	}
}
