package yunlai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient(
		WithApiKey("test_api_key"),
	)

	if client == nil {
		t.Fatal("client should not be nil")
	}

	if client.config.ApiKey != "test_api_key" {
		t.Errorf("expected ApiKey to be test_api_key, got %s", client.config.ApiKey)
	}
}

func TestNewClientWithApiKey(t *testing.T) {
	client := NewClientWithApiKey("test_api_key")

	if client == nil {
		t.Fatal("client should not be nil")
	}

	if client.config.ApiKey != "test_api_key" {
		t.Errorf("expected ApiKey to be test_api_key, got %s", client.config.ApiKey)
	}

	// 测试带额外选项
	client2 := NewClientWithApiKey("test_api_key",
		WithBaseURL("https://custom.api.com"),
		WithTimeout(60*time.Second),
	)

	if client2.config.BaseURL != "https://custom.api.com" {
		t.Errorf("expected BaseURL to be https://custom.api.com, got %s", client2.config.BaseURL)
	}

	if client2.config.Timeout != 60*time.Second {
		t.Errorf("expected Timeout to be 60s, got %v", client2.config.Timeout)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	client := NewClient(
		WithApiKey("test_api_key"),
		WithBaseURL("https://custom.api.com"),
		WithTimeout(60*time.Second),
		WithRetry(5),
	)

	if client.config.BaseURL != "https://custom.api.com" {
		t.Errorf("expected BaseURL to be https://custom.api.com, got %s", client.config.BaseURL)
	}

	if client.config.Timeout != 60*time.Second {
		t.Errorf("expected Timeout to be 60s, got %v", client.config.Timeout)
	}

	if client.config.Retry != 5 {
		t.Errorf("expected Retry to be 5, got %d", client.config.Retry)
	}
}

func TestClientErrorHandling(t *testing.T) {
	// 创建返回错误的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":400,"reason":"INVALID_PARAM","message":"invalid parameter","metadata":{}}`))
	}))
	defer server.Close()

	// 创建客户端
	client := NewClient(
		WithApiKey("test_api_key"),
		WithBaseURL(server.URL),
	)

	// 发送请求（使用泛型函数测试）
	type TestResponse struct {
		Data string `json:"data"`
	}
	ctx := context.Background()
	_, err := request[TestResponse](client, ctx, "/api/test", nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if apiErr, ok := err.(*Error); ok {
		if apiErr.Code != 400 {
			t.Errorf("expected HTTP status code 400, got %d", apiErr.Code)
		}
		if apiErr.Reason != "INVALID_PARAM" {
			t.Errorf("expected error reason INVALID_PARAM, got %s", apiErr.Reason)
		}
		if apiErr.Message != "invalid parameter" {
			t.Errorf("expected error message 'invalid parameter', got %s", apiErr.Message)
		}
	} else {
		t.Error("expected *Error type")
	}
}
