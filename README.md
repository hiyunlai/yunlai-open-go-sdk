# 云来开放平台 Go SDK

云来开放平台官方 Go 语言 SDK。

## 安装

```bash
go get github.com/hiyunlai/yunlai-open-sdk
```

## 快速开始

```go
package main

import (
    "context"
    "log"
    
    yunlai "github.com/hiyunlai/yunlai-open-sdk"
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

### Match - 竞彩投注

获取竞彩足球、竞彩篮球、北京单场等投注信息。

```go
// 获取竞彩足球投注列表
betList, _ := client.GetMatchBetList(ctx, "227", "spf")

// 获取传统足彩投注信息
betInfo, _ := client.GetCtzcBetInfo(ctx, "229", "24001")
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

### Draw - 开奖信息

获取数字彩、传统足彩等开奖历史和详情。

```go
// 获取竞彩开奖大厅
drawList, _ := client.GetLotteryDrawHomeList(ctx)

// 获取数字彩彩果详情
resultInfo, _ := client.GetDigitalResultInfo(ctx, "239", "24001")
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

## 配置选项

```go
client := yunlai.NewClientWithApiKey("your_api_key",
    yunlai.WithBaseURL("https://api.58st.cn"),    // 自定义 API 基础 URL
    yunlai.WithTimeout(30 * time.Second),         // 请求超时时间
    yunlai.WithRetry(3),                          // 重试次数
    yunlai.WithLogger(logger),                    // 自定义日志器
)
```

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
