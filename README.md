# 云来开放平台 Go SDK

云来开放平台官方 Go 语言 SDK，提供便捷的 API 调用接口。

## 特性

- ✅ 简洁的 API 设计
- ✅ 泛型架构，类型安全
- ✅ 模块化设计（Live、Match、Draw）
- ✅ API Key 认证
- ✅ 自动请求重试
- ✅ 上下文支持
- ✅ 完善的错误处理
- ✅ 日志记录支持

## 安装

```bash
go get github.com/hiyunlai/yunlai-open-sdk
```

## 快速开始

### 方式一：使用便捷方法（推荐）

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    yunlai "github.com/hiyunlai/yunlai-open-sdk"
)

func main() {
    // 创建客户端（只需提供 API Key）
    client := yunlai.NewClientWithApiKey("your_api_key")
    
    // 调用 Live 模块接口
    ctx := context.Background()
    matchInfo, err := client.GetLiveMatchInfo(ctx, "12345678")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("赛事: %s vs %s\n", 
        matchInfo.HomeTeamInfo.CnName, 
        matchInfo.AwayTeamInfo.CnName)
}
```

### 方式二：使用配置选项

```go
client := yunlai.NewClient(
    yunlai.WithApiKey("your_api_key"),
    yunlai.WithTimeout(60 * time.Second),
    yunlai.WithRetry(5),
)
```

## 业务模块

SDK 按照开放平台的业务领域划分为三大模块：

### Live 模块 - 比分直播

获取实时比赛数据、赛事信息、球队数据等。

```go
// 获取比分直播赛事信息
matchInfo, err := client.GetLiveMatchInfo(ctx, matchID)
```

**数据模型**：
- `LiveMatchInfo` - 赛事信息
- `LiveInfo` - 比分直播信息
- `TeamInfo` - 球队信息
- `TournamentInfo` - 联赛信息
- `LotteryMapping` - 彩票场次映射
- 等...

### Match 模块 - 竞彩投注（待实现）

获取竞彩足球、竞彩篮球、北京单场等投注和赛果信息。

### Draw 模块 - 开奖信息（待实现）

获取数字彩、传统足彩等开奖历史和详情信息。

## 配置选项

```go
client := yunlai.NewClientWithApiKey("your_api_key",
    yunlai.WithBaseURL("https://api.58st.cn"),    // 自定义 API 基础 URL（可选）
    yunlai.WithTimeout(30 * time.Second),         // 请求超时时间（可选，默认30秒）
    yunlai.WithRetry(3),                          // 重试次数（可选，默认3次）
    yunlai.WithLogger(logger),                    // 自定义日志器（可选）
)
```

## 错误处理

SDK 提供了统一的错误处理机制：

```go
matchInfo, err := client.GetLiveMatchInfo(ctx, matchID)
if err != nil {
    // 判断是否为 API 错误
    if apiErr, ok := err.(*yunlai.Error); ok {
        fmt.Printf("错误码: %d\n", apiErr.Code)
        fmt.Printf("错误原因: %s\n", apiErr.Reason)
        fmt.Printf("错误信息: %s\n", apiErr.Message)
        // 错误元数据
        fmt.Printf("元数据: %v\n", apiErr.Metadata)
    } else {
        // 其他错误（如网络错误）
        log.Fatal(err)
    }
    return
}
```

## 完整示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    yunlai "github.com/hiyunlai/yunlai-open-sdk"
)

func main() {
    // 创建客户端
    client := yunlai.NewClientWithApiKey("your_api_key",
        yunlai.WithTimeout(60*time.Second),
    )
    
    // 创建上下文
    ctx := context.Background()
    
    // 获取比分直播赛事信息
    matchInfo, err := client.GetLiveMatchInfo(ctx, "12345678")
    if err != nil {
        if apiErr, ok := err.(*yunlai.Error); ok {
            log.Printf("API错误: [%d:%s] %s", 
                apiErr.Code, apiErr.Reason, apiErr.Message)
        } else {
            log.Fatal(err)
        }
        return
    }
    
    // 输出赛事信息
    fmt.Printf("赛事ID: %s\n", matchInfo.MatchID)
    fmt.Printf("运动类型: %s\n", matchInfo.SportID)
    
    if matchInfo.LiveInfo != nil {
        fmt.Printf("比赛状态: %s\n", matchInfo.LiveInfo.StatusName)
    }
    
    if matchInfo.TournamentInfo != nil {
        fmt.Printf("联赛: %s\n", matchInfo.TournamentInfo.CnName)
    }
    
    if matchInfo.HomeTeamInfo != nil && matchInfo.AwayTeamInfo != nil {
        fmt.Printf("对阵: %s vs %s\n", 
            matchInfo.HomeTeamInfo.CnName,
            matchInfo.AwayTeamInfo.CnName)
    }
    
    // 比分信息
    if matchInfo.LiveInfo != nil {
        if matchInfo.LiveInfo.HomeTeamScore != nil {
            fmt.Printf("主队比分: %s\n", matchInfo.LiveInfo.HomeTeamScore.Score)
        }
        if matchInfo.LiveInfo.AwayTeamScore != nil {
            fmt.Printf("客队比分: %s\n", matchInfo.LiveInfo.AwayTeamScore.Score)
        }
    }
}
```

## 自定义日志

```go
import "log"

type MyLogger struct{}

func (l *MyLogger) Printf(format string, v ...interface{}) {
    log.Printf("[YUNLAI] "+format, v...)
}

client := yunlai.NewClientWithApiKey("your_api_key",
    yunlai.WithLogger(&MyLogger{}),
)
```

## 开发

### 运行测试

```bash
go test -v
```

### 编译

```bash
go build ./...
```

## 许可证

MIT License
