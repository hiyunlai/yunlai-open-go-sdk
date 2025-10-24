package yunlai

import (
	"context"
	"fmt"
)

// ==================== 业务接口 ====================

// GetLiveMatchInfo 获取比分直播赛事信息
func (c *Client) GetLiveMatchInfo(ctx context.Context, matchID string) (*LiveMatchInfo, error) {
	path := fmt.Sprintf("/v1/live/match/%s", matchID)
	return request[LiveMatchInfo](c, ctx, path, nil)
}

// ==================== 数据模型 ====================

// LiveMatchInfo 比分直播赛事信息
type LiveMatchInfo struct {
	// 赛事ID
	MatchID string `json:"matchId"`
	// 运动类型
	SportID string `json:"sportId"`
	// 比分直播信息
	LiveInfo *LiveInfo `json:"liveInfo,omitempty"`
	// 联赛信息
	TournamentInfo *TournamentInfo `json:"tournamentInfo,omitempty"`
	// 主队信息
	HomeTeamInfo *TeamInfo `json:"homeTeamInfo,omitempty"`
	// 客队信息
	AwayTeamInfo *TeamInfo `json:"awayTeamInfo,omitempty"`
	// 裁判信息
	RefereeInfo *RefereeInfo `json:"refereeInfo,omitempty"`
	// 场馆信息
	StadiumInfo *StadiumInfo `json:"stadiumInfo,omitempty"`
	// 天气信息
	WeatherInfo *WeatherInfo `json:"weatherInfo,omitempty"`
	// 彩票场次映射
	LotteryMapping *LotteryMapping `json:"lotteryMapping,omitempty"`
}

// LiveInfo 比分直播信息
type LiveInfo struct {
	// 状态 1-未开始 2-进行中 3-已结束 4-已取消
	Status string `json:"status"`
	// 状态码
	StatusCode string `json:"statusCode"`
	// 状态名
	StatusName string `json:"statusName"`
	// 比赛时间
	MatchTimeStr string `json:"matchTimeStr"`
	// 开始时间
	StartTimeStr string `json:"startTimeStr"`
	// 结束时间
	EndTimeStr string `json:"endTimeStr"`
	// 当前阶段开始时间
	CurrentPeriodStartTimeStr string `json:"currentPeriodStartTimeStr"`
	// 比赛时钟是否计时中
	TimeRunning bool `json:"timeRunning"`
	// 足球比赛已进行时间（秒）
	TimePlayed string `json:"timePlayed"`
	// 篮球每节比赛剩余时间（秒）
	TimeRemaining string `json:"timeRemaining"`
	// 主队阵容
	HomeFormation string `json:"homeFormation"`
	// 主队比分
	HomeTeamScore *MatchScore `json:"homeTeamScore,omitempty"`
	// 客队阵容
	AwayFormation string `json:"awayFormation"`
	// 客队比分
	AwayTeamScore *MatchScore `json:"awayTeamScore,omitempty"`
	// 是否为中立场
	IsNeutral bool `json:"isNeutral"`
}

// MatchScore 比分信息
type MatchScore struct {
	// 全场比分（含加时、点球）
	Score string `json:"score"`
	// 正常时间比分（90分钟）
	NormalTimeScore string `json:"normalTimeScore"`
	// 半场比分
	HalfTimeScore string `json:"halfTimeScore"`
}

// TournamentInfo 联赛信息
type TournamentInfo struct {
	// 联赛ID
	ID string `json:"id"`
	// 中文名
	CnName string `json:"cnName"`
	// 中文别名
	CnAlias string `json:"cnAlias"`
	// 颜色
	Color string `json:"color"`
	// 级别
	Level string `json:"level"`
	// LOGO 链接
	LogoURL string `json:"logoUrl"`
}

// TeamInfo 球队信息
type TeamInfo struct {
	// 球队ID
	ID string `json:"id"`
	// 中文名
	CnName string `json:"cnName"`
	// 中文别名
	CnAlias string `json:"cnAlias"`
	// LOGO 地址
	LogoURL string `json:"logoUrl"`
	// 排名
	Rank string `json:"rank"`
	// 联赛排名
	TournamentRank string `json:"tournamentRank"`
	// FIFA 俱乐部排名
	FifaClubRank string `json:"fifaClubRank"`
	// FIFA 国家队排名
	FifaCountryRank string `json:"fifaCountryRank"`
	// FIBA 国家队排名
	FibaCountryRank string `json:"fibaCountryRank"`
}

// RefereeInfo 裁判信息
type RefereeInfo struct {
	// 裁判ID
	ID string `json:"id"`
	// 中文名称
	CnName string `json:"cnName"`
	// 英文名称
	EnName string `json:"enName"`
}

// StadiumInfo 场馆信息
type StadiumInfo struct {
	// 场馆ID
	ID string `json:"id"`
	// 中文名称
	CnName string `json:"cnName"`
	// 英文名称
	EnName string `json:"enName"`
}

// WeatherInfo 天气信息
type WeatherInfo struct {
	// 天气
	Weather string `json:"weather"`
	// 温度
	Temperature string `json:"temperature"`
	// 湿度
	Humidity string `json:"humidity"`
	// 气压
	Pressure string `json:"pressure"`
	// 风向
	WindDirection string `json:"windDirection"`
	// 风速
	WindSpeed string `json:"windSpeed"`
}

// LotteryMapping 彩票场次映射
type LotteryMapping struct {
	// 是否是竞彩足球赛事
	IsJczq bool `json:"isJczq"`
	// 竞彩足球期次
	JczqIssue string `json:"jczqIssue"`
	// 竞彩足球轮次
	JczqRound string `json:"jczqRound"`
	// 竞彩足球场次编码
	JczqMatchCode string `json:"jczqMatchCode"`
	// 是否是竞彩篮球赛事
	IsJclq bool `json:"isJclq"`
	// 竞彩篮球期次
	JclqIssue string `json:"jclqIssue"`
	// 竞彩篮球轮次
	JclqRound string `json:"jclqRound"`
	// 竞彩篮球场次编码
	JclqMatchCode string `json:"jclqMatchCode"`
	// 是否是北京单场赛事
	IsBjdc bool `json:"isBjdc"`
	// 北京单场期次
	BjdcIssue string `json:"bjdcIssue"`
	// 北京单场轮次
	BjdcRound string `json:"bjdcRound"`
	// 北京单场场次编码
	BjdcMatchCode string `json:"bjdcMatchCode"`
	// 是否是胜负彩赛事
	IsSfc bool `json:"isSfc"`
	// 胜负彩期次
	SfcIssue string `json:"sfcIssue"`
	// 胜负彩轮次
	SfcRound string `json:"sfcRound"`
	// 胜负彩场次编码
	SfcMatchCode string `json:"sfcMatchCode"`
	// 是否是半全场赛事
	IsBqc bool `json:"isBqc"`
	// 半全场期次
	BqcIssue string `json:"bqcIssue"`
	// 半全场轮次
	BqcRound string `json:"bqcRound"`
	// 半全场场次编码
	BqcMatchCode string `json:"bqcMatchCode"`
	// 是否是总进球赛事
	IsZjq bool `json:"isZjq"`
	// 总进球期次
	ZjqIssue string `json:"zjqIssue"`
	// 总进球轮次
	ZjqRound string `json:"zjqRound"`
	// 总进球场次编码
	ZjqMatchCode string `json:"zjqMatchCode"`
}
