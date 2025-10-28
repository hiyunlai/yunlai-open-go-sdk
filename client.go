package yunlai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

// ============================================================================
// 常量定义
// ============================================================================

const (
	// DefaultBaseURL 默认的API基础URL
	DefaultBaseURL = "https://api.58st.cn"
	// DefaultTimeout 默认的请求超时时间
	DefaultTimeout = 30 * time.Second
	// DefaultRetry 默认的重试次数
	DefaultRetry = 3
)

// 业务错误码常量
const (
	// ReasonInvalidRequest 无效的请求
	ReasonInvalidRequest = "INVALID_REQUEST"
	// ReasonNetworkError 网络错误
	ReasonNetworkError = "NETWORK_ERROR"
	// ReasonAPIError API错误
	ReasonAPIError = "API_ERROR"
	// ReasonSignError 签名错误
	ReasonSignError = "SIGN_ERROR"
)

// ============================================================================
// 类型定义
// ============================================================================

// Error 自定义错误类型
type Error struct {
	// Code HTTP状态码
	Code int `json:"code"`
	// Reason 业务错误码
	Reason string `json:"reason"`
	// Message 错误信息
	Message string `json:"message"`
	// Metadata 元数据
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Logger 日志接口
type Logger interface {
	Printf(format string, v ...any)
}

// Config 客户端配置
type Config struct {
	// ApiKey API密钥（用于Bearer Token认证）
	ApiKey string
	// BaseURL API基础URL
	BaseURL string
	// Timeout 请求超时时间
	Timeout time.Duration
	// Retry 重试次数
	Retry int
	// Logger 日志记录器
	Logger Logger
}

// Option 配置选项函数
type Option func(*Config)

// DefaultLogger 默认日志记录器
type DefaultLogger struct{}

// Client 云来开放平台客户端
type Client struct {
	config      *Config
	restyClient *resty.Client
}

// ============================================================================
// 错误相关函数和方法
// ============================================================================

// Error 实现error接口
func (e *Error) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("[%d:%s] %s", e.Code, e.Reason, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError 创建新的错误（使用业务错误码）
func NewError(reason, message string) *Error {
	return &Error{
		Code:    500,
		Reason:  reason,
		Message: message,
	}
}

// NewErrorWithCode 创建带HTTP状态码的错误
func NewErrorWithCode(code int, reason, message string) *Error {
	return &Error{
		Code:    code,
		Reason:  reason,
		Message: message,
	}
}

// IsError 判断是否为指定业务错误码的错误
func IsError(err error, reason string) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.Reason == reason
	}
	return false
}

// ============================================================================
// Logger 实现
// ============================================================================

// Printf 实现Logger接口
func (l *DefaultLogger) Printf(format string, v ...any) {
	log.Printf(format, v...)
}

// ============================================================================
// 配置选项函数
// ============================================================================

// WithApiKey 设置API密钥
func WithApiKey(apiKey string) Option {
	return func(c *Config) {
		c.ApiKey = apiKey
	}
}

// WithBaseURL 设置API基础URL
func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		c.BaseURL = baseURL
	}
}

// WithLogger 设置日志记录器
func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// WithRetry 设置重试次数
func WithRetry(retry int) Option {
	return func(c *Config) {
		c.Retry = retry
	}
}

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// ============================================================================
// Client 构造函数
// ============================================================================

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

// ============================================================================
// Client 辅助函数
// ============================================================================

// request 发送GET请求并解析响应为指定类型
func request[T any](c *Client, ctx context.Context, path string, queryParams map[string]string, headers map[string]string) (*T, error) {
	// 创建resty请求
	restyReq := c.restyClient.R().SetContext(ctx)

	// 设置查询参数
	if queryParams != nil {
		restyReq.SetQueryParams(queryParams)
	}

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

// buildQueryParams 构建查询参数map，自动过滤空值
func buildQueryParams(params map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		if v != "" {
			result[k] = v
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
