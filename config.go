package yunlai

import (
	"log"
	"time"
)

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

// Logger 日志接口
type Logger interface {
	Printf(format string, v ...interface{})
}

// Option 配置选项函数
type Option func(*Config)

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

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithRetry 设置重试次数
func WithRetry(retry int) Option {
	return func(c *Config) {
		c.Retry = retry
	}
}

// WithLogger 设置日志记录器
func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// DefaultLogger 默认日志记录器
type DefaultLogger struct{}

// Printf 实现Logger接口
func (l *DefaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
