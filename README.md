# 云来开放平台 Go SDK

云来开放平台官方 Go 语言 SDK。

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## 特性

✅ **零外部依赖** - 只使用 Go 标准库  
✅ **类型安全** - 完整的类型定义和泛型支持  
✅ **易于使用** - 简洁的 API 设计  
✅ **错误处理** - 详细的错误信息和自定义错误类型  
✅ **日志支持** - 可配置的日志记录功能  

## 安装

```bash
go get github.com/hiyunlai/yunlai-open-go-sdk
```

## 快速开始

```go
package main

import (
    "context"
    "log"
    
    yunlai "github.com/hiyunlai/yunlai-open-go-sdk"
)

func main() {
    // 创建客户端
    client := yunlai.NewClientWithApiKey("your_api_key")
    
    // 调用接口
    ctx := context.Background()
    matchInfo, err := client.GetLiveMatchInfo(ctx, "12345678")
    if err != nil {
        log.Fatal(err)
    }
    
    // 使用返回数据
    log.Printf("赛事: %s vs %s", 
        matchInfo.HomeTeamInfo.CnName,
        matchInfo.AwayTeamInfo.CnName)
}
```

## 业务模块

SDK 按业务领域划分为三大模块：

### Match - 竞彩投注

获取竞彩足球、竞彩篮球、北京单场等投注信息。

```go
// 获取竞彩足球投注列表
betList, _ := client.GetMatchBetList(ctx, "227", "5")

// 获取传统足彩投注信息
betInfo, _ := client.GetCtzcBetInfo(ctx, "172", "25001")
```

<details>
<summary>查看全部接口</summary>

- `GetMatchBetList` - 赛事投注列表（北单、竞足、竞篮）
- `GetMatchBetInfo` - 赛事投注信息（北单、竞足、竞篮）
- `GetCtzcBetList` - 传统足彩投注列表
- `GetCtzcBetInfo` - 传统足彩投注信息
- `GetDigitalBetList` - 数字彩投注列表
- `GetDigitalBetInfo` - 数字彩投注信息
- `GetGyjBetInfo` - 冠亚军投注信息

</details>

### Live - 比分直播

获取实时比赛数据、赛事信息、球队数据等。

```go
// 获取赛事信息
matchInfo, _ := client.GetLiveMatchInfo(ctx, matchID)

// 获取赛事列表
matchList, _ := client.GetLiveMatchList(ctx, sportID, startTime, endTime, lotteryType)

// 获取赛事阵容
lineup, _ := client.GetLiveMatchLineup(ctx, matchID)
```

<details>
<summary>查看全部接口</summary>

- `GetLiveMatchInfo` - 赛事信息
- `GetLiveMatchList` - 赛事列表
- `GetLiveMatchLotteryMapping` - 彩种赛事映射
- `GetLiveMatchHistory` - 历史对阵
- `GetLiveMatchLineup` - 赛事阵容
- `GetLiveMatchNews` - 新闻情报
- `GetLiveMatchIntelligenceNews` - 专属情报
- `GetLiveMatchOdds` - 百家盘赔率
- `GetLiveMatchPhaseAll` - 直播全数据
- `GetLiveMatchPhaseEvent` - 事件直播
- `GetLiveMatchPhaseText` - 文字直播
- `GetLiveMatchScore` - 比分详情
- `GetLiveMatchStatic` - 数据统计
- `GetLiveMatchPlayerStatic` - 球员统计

</details>

### Draw - 开奖信息

获取数字彩、传统足彩等开奖历史和详情。

```go
// 获取竞彩开奖大厅
drawList, _ := client.GetLotteryDrawHomeList(ctx)

// 获取数字彩彩果详情
resultInfo, _ := client.GetDigitalResultInfo(ctx, "188", "25001")
```

<details>
<summary>查看全部接口</summary>

- `GetLotteryDrawHomeList` - 竞彩开奖大厅
- `GetDrawHistoryList` - 竞彩开奖历史列表
- `GetDigitalResultInfo` - 数字彩彩果详情
- `GetDigitalHomeResultList` - 数字彩列表
- `GetDigitalResultList` - 数字彩历史列表
- `GetCtzcResultInfo` - 传统足彩赛果详情
- `GetCtzcHomeResultList` - 传统足彩列表
- `GetCtzcResultList` - 传统足彩历史列表

</details>

## Callback - 彩果回调

**重要说明：** Callback 接口不是开放平台提供给合作方调用的接口，而是**合作方需要实现的回调接口**，由开放平台主动调用来通知彩果信息。

### 接口定义

**接口路径：** `POST /v1/callback`

**请求参数：** `CallbackRequest`

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| lotteryType | string | 是 | 彩种类型，见下方彩种类型说明 |
| unionKey | string | 是 | 唯一键，不同彩种格式不同，见下方格式说明 |

**响应参数：** `CallbackReply` (空对象)

### 彩种类型

| 值 | 说明 |
|----|------|
| 226 | 北京单场 |
| 227 | 竞彩足球 |
| 228 | 竞彩篮球 |
| 171 | 6场半全场 |
| 172 | 胜负游戏 |
| 173 | 任选9场 |
| 174 | 4场进球 |
| 163 | 排列三 |
| 164 | 排列五 |
| 166 | 双色球 |
| 167 | 福彩3D |
| 168 | 七星彩 |
| 185 | 快乐8 |
| 186 | 七乐彩 |
| 188 | 大乐透 |
| 229 | 冠亚军 |
| 230 | 冠军 |

### unionKey 格式说明

不同彩种的 `unionKey` 格式不同：

- **北京单场格式：** `玩法:场次编码`
  - 示例：`1:2510598`（胜平负玩法的2510598场次）
- **竞彩足球/篮球格式：** `场次编码`
  - 示例：`2510271001`
- **传统足彩格式：** `期次`
  - 示例：`25001`
- **数字彩格式：** `期次`
  - 示例：`25001`
- **冠亚军格式：** `期次`
  - 示例：`1118245`

### 实现示例

合作方需要实现该回调接口的服务端处理逻辑：

```go
// 示例：使用标准库实现 Callback 接口
func handleCallback(w http.ResponseWriter, r *http.Request) {
    var req yunlai.CallbackRequest
    
    // 解析请求参数
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // 处理彩果通知
    log.Printf("收到彩果通知: 彩种=%s, 唯一键=%s", 
        req.LotteryType, req.UnionKey)
    
    // 根据 lotteryType 和 unionKey 查询彩果详情
    // 然后进行业务处理...
    
    // 返回成功响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(yunlai.CallbackReply{})
}
```

## 配置选项

```go
client := yunlai.NewClientWithApiKey("your_api_key",
    yunlai.WithBaseURL("https://api.58st.cn"),    // 自定义 API 基础 URL
    yunlai.WithTimeout(30 * time.Second),         // 请求超时时间
    yunlai.WithLogger(logger),                    // 自定义日志器
)
```

### 可用配置选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| `WithApiKey(string)` | API 密钥（必需） | - |
| `WithBaseURL(string)` | API 基础 URL | `https://api.58st.cn` |
| `WithTimeout(time.Duration)` | 请求超时时间 | `30s` |
| `WithLogger(Logger)` | 自定义日志记录器 | `nil` |

## 错误处理

```go
matchInfo, err := client.GetLiveMatchInfo(ctx, matchID)
if err != nil {
    if apiErr, ok := err.(*yunlai.Error); ok {
        log.Printf("API错误: [%d:%s] %s", 
            apiErr.Code, apiErr.Reason, apiErr.Message)
    } else {
        log.Fatal(err)
    }
}
```

## 许可证

MIT License
