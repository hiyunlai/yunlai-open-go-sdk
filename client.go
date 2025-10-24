package yunlai

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// DefaultBaseURL 默认的API基础URL
	DefaultBaseURL = "https://api.58st.cn"
	// DefaultTimeout 默认的请求超时时间
	DefaultTimeout = 30 * time.Second
	// DefaultRetry 默认的重试次数
	DefaultRetry = 3
)

// Client 云来开放平台客户端
type Client struct {
	config      *Config
	restyClient *resty.Client
}

// NewClient 创建新的客户端实例
func NewClient(opts ...Option) *Client {
	config := &Config{
		BaseURL: DefaultBaseURL,
		Timeout: DefaultTimeout,
		Retry:   DefaultRetry,
	}

	// 应用配置选项
	for _, opt := range opts {
		opt(config)
	}

	// 验证必需的配置
	if config.ApiKey == "" {
		panic("api key is required")
	}

	// 创建resty客户端
	restyClient := resty.New()
	restyClient.SetBaseURL(config.BaseURL)
	restyClient.SetTimeout(config.Timeout)
	restyClient.SetRetryCount(config.Retry)
	
	// 设置Bearer Token
	restyClient.SetAuthToken(config.ApiKey)
	
	// 设置默认请求头
	restyClient.SetHeader("Content-Type", "application/json")
	restyClient.SetHeader("User-Agent", "yunlai-go-sdk/1.0.0")

	// 如果有日志记录器，设置调试模式
	if config.Logger != nil {
		restyClient.SetDebug(true)
		restyClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
			config.Logger.Printf("Request: %s %s", req.Method, req.URL)
			return nil
		})
		restyClient.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
			config.Logger.Printf("Response: %d %s", resp.StatusCode(), string(resp.Body()))
			return nil
		})
	}

	client := &Client{
		config:      config,
		restyClient: restyClient,
	}

	return client
}

// NewClientWithApiKey 使用API密钥创建客户端实例（便捷方法）
func NewClientWithApiKey(apiKey string, opts ...Option) *Client {
	// 将 ApiKey 选项添加到选项列表前面
	allOpts := make([]Option, 0, len(opts)+1)
	allOpts = append(allOpts, WithApiKey(apiKey))
	allOpts = append(allOpts, opts...)
	
	return NewClient(allOpts...)
}

// request 发送GET请求并解析响应为指定类型
func request[T any](c *Client, ctx context.Context, path string, headers map[string]string) (*T, error) {
	// 创建resty请求
	restyReq := c.restyClient.R().SetContext(ctx)

	// 设置自定义请求头
	if headers != nil {
		restyReq.SetHeaders(headers)
	}

	// 执行GET请求
	resp, err := restyReq.Get(path)
	if err != nil {
		return nil, NewError(ReasonNetworkError, fmt.Sprintf("network error: %v", err))
	}

	// 检查HTTP状态码
	if resp.StatusCode() >= 400 {
		// 尝试解析API错误响应
		var apiErr Error
		if err := json.Unmarshal(resp.Body(), &apiErr); err == nil && apiErr.Reason != "" {
			// 成功解析API错误响应
			return nil, &apiErr
		}

		// 无法解析错误响应，返回通用错误
		return nil, NewErrorWithCode(
			resp.StatusCode(),
			ReasonAPIError,
			fmt.Sprintf("API error: HTTP %d", resp.StatusCode()),
		)
	}

	// 解析响应为指定类型
	var result T
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, NewError(
			ReasonAPIError,
			fmt.Sprintf("failed to unmarshal response: %v", err),
		)
	}

	return &result, nil
}
