package yunlai

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

// getTestClient 获取测试客户端（所有测试文件共用）
func getTestClient() *Client {
	return NewClientWithApiKey("hIOg6sxoBqNb78TSCyzQeVcuDHlGkYW4")
}

// ==================== Match 测试 ====================

// TestGetMatchBetList 测试获取赛事投注列表
func TestGetMatchBetList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetMatchBetList(ctx, LotteryTypeJczq, SubTypeJczqHhgg)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetMatchBetList Response:\n%s\n\n", resultJSON)
}

// TestGetMatchBetInfo 测试获取赛事投注信息
func TestGetMatchBetInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetMatchBetInfo(ctx, LotteryTypeJczq, SubTypeJczqHhgg, "2510282001")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetMatchBetInfo Response:\n%s\n\n", resultJSON)
}

// TestGetCtzcBetList 测试获取传统足彩投注列表
func TestGetCtzcBetList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetCtzcBetList(ctx, LotteryTypeSfc)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetCtzcBetList Response:\n%s\n\n", resultJSON)
}

// TestGetCtzcBetInfo 测试获取传统足彩投注信息
func TestGetCtzcBetInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetCtzcBetInfo(ctx, LotteryTypeSfc, "25153")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetCtzcBetInfo Response:\n%s\n\n", resultJSON)
}

// TestGetDigitalBetList 测试获取数字彩投注列表
func TestGetDigitalBetList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDigitalBetList(ctx, LotteryTypeDlt)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDigitalBetList Response:\n%s\n\n", resultJSON)
}

// TestGetDigitalBetInfo 测试获取数字彩投注信息
func TestGetDigitalBetInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDigitalBetInfo(ctx, LotteryTypeDlt, "25122")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDigitalBetInfo Response:\n%s\n\n", resultJSON)
}

// TestGetGyjBetInfo 测试获取冠亚军投注信息
func TestGetGyjBetInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetGyjBetInfo(ctx, LotteryTypeGyj, "1118245")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetGyjBetInfo Response:\n%s\n\n", resultJSON)
}

// ==================== Live 测试 ====================

// TestGetLiveMatchLotteryMapping 测试获取彩种赛事对应比分直播赛事
func TestGetLiveMatchLotteryMapping(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchLotteryMapping(ctx, LotteryTypeJczq, "2510271001")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchLotteryMapping Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchList 测试获取比分直播赛事列表
func TestGetLiveMatchList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchList(ctx, SportIdFootball, "", "", LotteryTypeJczq)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchList Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchInfo 测试获取比分直播赛事信息
func TestGetLiveMatchInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchInfo(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchInfo Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchHistory 测试获取比分直播赛事历史对阵
func TestGetLiveMatchHistory(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchHistory(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchHistory Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchLineup 测试获取比分直播赛事阵容
func TestGetLiveMatchLineup(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchLineup(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchLineup Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchNews 测试获取比分直播赛事新闻情报
func TestGetLiveMatchNews(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchNews(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchNews Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchIntelligenceNews 测试获取比分直播赛事专属情报
func TestGetLiveMatchIntelligenceNews(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchIntelligenceNews(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchIntelligenceNews Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchOdds 测试获取比分直播赛事百家盘赔率
func TestGetLiveMatchOdds(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchOdds(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchOdds Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchPhaseAll 测试获取比分直播赛事直播全数据
func TestGetLiveMatchPhaseAll(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchPhaseAll(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchPhaseAll Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchPhaseEvent 测试获取比分直播赛事事件直播
func TestGetLiveMatchPhaseEvent(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchPhaseEvent(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchPhaseEvent Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchPhaseText 测试获取比分直播赛事文字直播
func TestGetLiveMatchPhaseText(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchPhaseText(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchPhaseText Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchScore 测试获取比分直播赛事比分详情
func TestGetLiveMatchScore(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchScore(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchScore Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchStatic 测试获取比分直播赛事数据统计
func TestGetLiveMatchStatic(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchStatic(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchStatic Response:\n%s\n\n", resultJSON)
}

// TestGetLiveMatchPlayerStatic 测试获取比分直播赛事球员数据统计
func TestGetLiveMatchPlayerStatic(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLiveMatchPlayerStatic(ctx, "12725688")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLiveMatchPlayerStatic Response:\n%s\n\n", resultJSON)
}

// ==================== Draw 测试 ====================

// TestGetLotteryDrawHomeList 测试获取竞彩开奖大厅
func TestGetLotteryDrawHomeList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetLotteryDrawHomeList(ctx)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetLotteryDrawHomeList Response:\n%s\n\n", resultJSON)
}

// TestGetDrawHistoryList 测试获取竞彩开奖历史列表
func TestGetDrawHistoryList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDrawHistoryList(ctx, LotteryTypeJczq, "2025-10-27")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDrawHistoryList Response:\n%s\n\n", resultJSON)
}

// TestGetCtzcHomeResultList 测试获取传统足彩列表
func TestGetCtzcHomeResultList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetCtzcHomeResultList(ctx)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetCtzcHomeResultList Response:\n%s\n\n", resultJSON)
}

// TestGetCtzcResultList 测试获取传统足彩历史列表
func TestGetCtzcResultList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetCtzcResultList(ctx, LotteryTypeSfc, "10")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetCtzcResultList Response:\n%s\n\n", resultJSON)
}

// TestGetCtzcResultInfo 测试获取传统足彩赛果详情
func TestGetCtzcResultInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetCtzcResultInfo(ctx, LotteryTypeSfc, "25153")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetCtzcResultInfo Response:\n%s\n\n", resultJSON)
}

// TestGetDigitalHomeResultList 测试获取数字彩列表
func TestGetDigitalHomeResultList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDigitalHomeResultList(ctx)
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDigitalHomeResultList Response:\n%s\n\n", resultJSON)
}

// TestGetDigitalResultList 测试获取数字彩历史列表
func TestGetDigitalResultList(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDigitalResultList(ctx, LotteryTypeSsq, "10")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDigitalResultList Response:\n%s\n\n", resultJSON)
}

// TestGetDigitalResultInfo 测试获取数字彩彩果详情
func TestGetDigitalResultInfo(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	result, err := client.GetDigitalResultInfo(ctx, LotteryTypeDlt, "25122")
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("GetDigitalResultInfo Response:\n%s\n\n", resultJSON)
}
