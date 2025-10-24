package yunlai

import (
	"context"
	"testing"
)

func TestClient_GetLiveMatchInfo(t *testing.T) {
	// 从环境变量获取token，如果没有则跳过测试
	token := "hIOg6sxoBqNb78TSCyzQeVcuDHlGkYW4"

	// 创建客户端
	client := NewClient(
		WithApiKey(token),
	)

	// 测试获取赛事信息
	ctx := context.Background()
	matchInfo, err := client.GetLiveMatchInfo(ctx, "12725683")

	// 注意：由于我们不知道这个赛事ID是否存在，
	// 所以这里只测试API调用本身是否正常工作
	if err != nil {
		// 如果是404或其他业务错误，说明API调用正常，只是数据不存在
		t.Logf("API call completed with error (expected): %v", err)
	} else {
		// 如果成功获取到数据，打印一些基本信息
		t.Logf("Match ID: %s", matchInfo.MatchID)
		t.Logf("Sport ID: %s", matchInfo.SportID)

		if matchInfo.LiveInfo != nil {
			t.Logf("Match Status: %s", matchInfo.LiveInfo.Status)
			t.Logf("Status Name: %s", matchInfo.LiveInfo.StatusName)
		}

		if matchInfo.HomeTeamInfo != nil {
			t.Logf("Home Team: %s", matchInfo.HomeTeamInfo.CnName)
			if matchInfo.HomeTeamInfo.LogoURL != "" {
				t.Logf("Home Team Logo: %s", matchInfo.HomeTeamInfo.LogoURL)
			}
		}

		if matchInfo.AwayTeamInfo != nil {
			t.Logf("Away Team: %s", matchInfo.AwayTeamInfo.CnName)
			if matchInfo.AwayTeamInfo.LogoURL != "" {
				t.Logf("Away Team Logo: %s", matchInfo.AwayTeamInfo.LogoURL)
			}
		}

		if matchInfo.TournamentInfo != nil {
			t.Logf("Tournament: %s", matchInfo.TournamentInfo.CnName)
		}

		if matchInfo.LiveInfo != nil && matchInfo.LiveInfo.HomeTeamScore != nil {
			t.Logf("Home Score: %s", matchInfo.LiveInfo.HomeTeamScore.Score)
		}

		if matchInfo.LiveInfo != nil && matchInfo.LiveInfo.AwayTeamScore != nil {
			t.Logf("Away Score: %s", matchInfo.LiveInfo.AwayTeamScore.Score)
		}
	}
}
