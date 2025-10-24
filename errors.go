package yunlai

import "fmt"

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
	// ReasonMatchNotExist 赛事不存在
	ReasonMatchNotExist = "MATCH_NOT_EXIST"
)

// Error 自定义错误类型
type Error struct {
	// Code HTTP状态码
	Code int `json:"code"`
	// Reason 业务错误码
	Reason string `json:"reason"`
	// Message 错误信息
	Message string `json:"message"`
	// Metadata 元数据
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

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
