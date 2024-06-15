package main

import (
	"testing"
)

// MockServiceはServiceインターフェースを実装するモック
type MockService struct{}

// DoSomethingはMockServiceのモックメソッド
func (m MockService) DoSomething() string {
	return "Mocked Response"
}

// TestClientUseServiceはClientのUseServiceメソッドのテスト
func TestClientUseService(t *testing.T) {
	// MockServiceのインスタンスを作成
	mockService := MockService{}

	// MockServiceのインスタンスをClientに注入
	client := NewClient(mockService)

	// モックされたServiceのメソッドの結果を検証
	if result := client.service.DoSomething(); result != "Mocked Response" {
		t.Errorf("expected 'Mocked Response', got '%s'", result)
	}
}
