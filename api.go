package yunlai

import (
	"context"
	"fmt"
)

// ==================== 业务接口 ====================

// GetMatchBetList 获取赛事投注列表（北单、竞足、竞篮）
func (c *Client) GetMatchBetList(ctx context.Context, lotteryType LotteryType, subType SubType) (*MatchBetList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"subType":     fmt.Sprintf("%d", subType),
	})
	return request[MatchBetList](c, ctx, "/v1/match/bet-list", queryParams, nil)
}

// GetMatchBetInfo 获取赛事投注信息（北单、竞足、竞篮）
func (c *Client) GetMatchBetInfo(ctx context.Context, lotteryType LotteryType, subType SubType, matchCode string) (*MatchBetInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"subType":     fmt.Sprintf("%d", subType),
		"matchCode":   matchCode,
	})
	return request[MatchBetInfo](c, ctx, "/v1/match/bet-info", queryParams, nil)
}

// GetCtzcBetList 获取传统足彩投注列表
func (c *Client) GetCtzcBetList(ctx context.Context, lotteryType LotteryType) (*CtzcBetList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
	})
	return request[CtzcBetList](c, ctx, "/v1/match/ctzc/bet-list", queryParams, nil)
}

// GetCtzcBetInfo 获取传统足彩投注信息
func (c *Client) GetCtzcBetInfo(ctx context.Context, lotteryType LotteryType, issue string) (*CtzcBetInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"issue":       issue,
	})
	return request[CtzcBetInfo](c, ctx, "/v1/match/ctzc/bet-info", queryParams, nil)
}

// GetDigitalBetList 获取数字彩投注列表
func (c *Client) GetDigitalBetList(ctx context.Context, lotteryType LotteryType) (*DigitalBetList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
	})
	return request[DigitalBetList](c, ctx, "/v1/match/digital/bet-list", queryParams, nil)
}

// GetDigitalBetInfo 获取数字彩投注信息
func (c *Client) GetDigitalBetInfo(ctx context.Context, lotteryType LotteryType, issue string) (*DigitalBetInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"issue":       issue,
	})
	return request[DigitalBetInfo](c, ctx, "/v1/match/digital/bet-info", queryParams, nil)
}

// GetGyjBetInfo 获取冠亚军投注信息
func (c *Client) GetGyjBetInfo(ctx context.Context, lotteryType LotteryType, issue string) (*GyjBetInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"issue":       issue,
	})
	return request[GyjBetInfo](c, ctx, "/v1/match/gyj/bet-info", queryParams, nil)
}

// GetLiveMatchLotteryMapping 获取彩种赛事对应比分直播赛事
func (c *Client) GetLiveMatchLotteryMapping(ctx context.Context, lotteryType LotteryType, matchCode string) (*LiveMatchLotteryMapping, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"matchCode":   matchCode,
	})
	return request[LiveMatchLotteryMapping](c, ctx, "/v1/live/match/lottery/mapping", queryParams, nil)
}

// GetLiveMatchList 获取比分直播赛事列表
func (c *Client) GetLiveMatchList(ctx context.Context, sportId SportId, startTimeStr, endTimeStr string, lotteryType LotteryType) (*LiveMatchList, error) {
	queryParams := buildQueryParams(map[string]string{
		"sportId":      sportId,
		"startTimeStr": startTimeStr,
		"endTimeStr":   endTimeStr,
		"lotteryType":  fmt.Sprintf("%d", lotteryType),
	})
	return request[LiveMatchList](c, ctx, "/v1/live/match", queryParams, nil)
}

// GetLiveMatchInfo 获取比分直播赛事信息
func (c *Client) GetLiveMatchInfo(ctx context.Context, matchId string) (*LiveMatchInfo, error) {
	path := fmt.Sprintf("/v1/live/match/%s", matchId)
	return request[LiveMatchInfo](c, ctx, path, nil, nil)
}

// GetLiveMatchHistory 获取比分直播赛事历史对阵
func (c *Client) GetLiveMatchHistory(ctx context.Context, matchId string) (*LiveMatchHistory, error) {
	path := fmt.Sprintf("/v1/live/match/%s/history", matchId)
	return request[LiveMatchHistory](c, ctx, path, nil, nil)
}

// GetLiveMatchLineup 获取比分直播赛事阵容
func (c *Client) GetLiveMatchLineup(ctx context.Context, matchId string) (*LiveMatchLineup, error) {
	path := fmt.Sprintf("/v1/live/match/%s/lineup", matchId)
	return request[LiveMatchLineup](c, ctx, path, nil, nil)
}

// GetLiveMatchNews 获取比分直播赛事新闻情报
func (c *Client) GetLiveMatchNews(ctx context.Context, matchId string) (*LiveMatchNews, error) {
	path := fmt.Sprintf("/v1/live/match/%s/news", matchId)
	return request[LiveMatchNews](c, ctx, path, nil, nil)
}

// GetLiveMatchIntelligenceNews 获取比分直播赛事专属情报
func (c *Client) GetLiveMatchIntelligenceNews(ctx context.Context, matchId string) (*LiveMatchIntelligenceNews, error) {
	path := fmt.Sprintf("/v1/live/match/%s/news/intelligence", matchId)
	return request[LiveMatchIntelligenceNews](c, ctx, path, nil, nil)
}

// GetLiveMatchOdds 获取比分直播赛事百家盘赔率
func (c *Client) GetLiveMatchOdds(ctx context.Context, matchId string) (*LiveMatchOdds, error) {
	path := fmt.Sprintf("/v1/live/match/%s/odds", matchId)
	return request[LiveMatchOdds](c, ctx, path, nil, nil)
}

// GetLiveMatchPhaseAll 获取比分直播赛事直播全数据
func (c *Client) GetLiveMatchPhaseAll(ctx context.Context, matchId string) (*LiveMatchPhaseAll, error) {
	path := fmt.Sprintf("/v1/live/match/%s/phase/all", matchId)
	return request[LiveMatchPhaseAll](c, ctx, path, nil, nil)
}

// GetLiveMatchPhaseEvent 获取比分直播赛事事件直播
func (c *Client) GetLiveMatchPhaseEvent(ctx context.Context, matchId string) (*LiveMatchPhaseEvent, error) {
	path := fmt.Sprintf("/v1/live/match/%s/phase/event", matchId)
	return request[LiveMatchPhaseEvent](c, ctx, path, nil, nil)
}

// GetLiveMatchPhaseText 获取比分直播赛事文字直播
func (c *Client) GetLiveMatchPhaseText(ctx context.Context, matchId string) (*LiveMatchPhaseText, error) {
	path := fmt.Sprintf("/v1/live/match/%s/phase/text", matchId)
	return request[LiveMatchPhaseText](c, ctx, path, nil, nil)
}

// GetLiveMatchScore 获取比分直播赛事比分详情
func (c *Client) GetLiveMatchScore(ctx context.Context, matchId string) (*LiveMatchScoreDetail, error) {
	path := fmt.Sprintf("/v1/live/match/%s/score", matchId)
	return request[LiveMatchScoreDetail](c, ctx, path, nil, nil)
}

// GetLiveMatchStatic 获取比分直播赛事数据统计
func (c *Client) GetLiveMatchStatic(ctx context.Context, matchId string) (*LiveMatchStaticDetail, error) {
	path := fmt.Sprintf("/v1/live/match/%s/static", matchId)
	return request[LiveMatchStaticDetail](c, ctx, path, nil, nil)
}

// GetLiveMatchPlayerStatic 获取比分直播赛事球员数据统计
func (c *Client) GetLiveMatchPlayerStatic(ctx context.Context, matchId string) (*LiveMatchPlayerStaticDetail, error) {
	path := fmt.Sprintf("/v1/live/match/%s/static/player", matchId)
	return request[LiveMatchPlayerStaticDetail](c, ctx, path, nil, nil)
}

// GetLotteryDrawHomeList 获取竞彩开奖大厅
func (c *Client) GetLotteryDrawHomeList(ctx context.Context) (*LotteryDrawHomeList, error) {
	return request[LotteryDrawHomeList](c, ctx, "/v1/draw/list", nil, nil)
}

// GetDrawHistoryList 获取竞彩开奖历史列表
func (c *Client) GetDrawHistoryList(ctx context.Context, lotteryType LotteryType, dateStr string) (*DrawHistoryList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"dateStr":     dateStr,
	})
	return request[DrawHistoryList](c, ctx, "/v1/draw/history/list", queryParams, nil)
}

// GetCtzcHomeResultList 获取传统足彩列表
func (c *Client) GetCtzcHomeResultList(ctx context.Context) (*CtzcHomeResultList, error) {
	return request[CtzcHomeResultList](c, ctx, "/v1/draw/ctzc/list", nil, nil)
}

// GetCtzcResultList 获取传统足彩历史列表
func (c *Client) GetCtzcResultList(ctx context.Context, lotteryType LotteryType, count string) (*CtzcResultList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"count":       count,
	})
	return request[CtzcResultList](c, ctx, "/v1/draw/ctzc/history/list", queryParams, nil)
}

// GetCtzcResultInfo 获取传统足彩赛果详情
func (c *Client) GetCtzcResultInfo(ctx context.Context, lotteryType LotteryType, issue string) (*CtzcResultInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"issue":       issue,
	})
	return request[CtzcResultInfo](c, ctx, "/v1/draw/ctzc", queryParams, nil)
}

// GetDigitalHomeResultList 获取数字彩列表
func (c *Client) GetDigitalHomeResultList(ctx context.Context) (*DigitalHomeResultList, error) {
	return request[DigitalHomeResultList](c, ctx, "/v1/draw/digital/list", nil, nil)
}

// GetDigitalResultList 获取数字彩历史列表
func (c *Client) GetDigitalResultList(ctx context.Context, lotteryType LotteryType, count string) (*DigitalResultList, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"count":       count,
	})
	return request[DigitalResultList](c, ctx, "/v1/draw/digital/history/list", queryParams, nil)
}

// GetDigitalResultInfo 获取数字彩彩果详情
func (c *Client) GetDigitalResultInfo(ctx context.Context, lotteryType LotteryType, issue string) (*DigitalResultInfo, error) {
	queryParams := buildQueryParams(map[string]string{
		"lotteryType": fmt.Sprintf("%d", lotteryType),
		"issue":       issue,
	})
	return request[DigitalResultInfo](c, ctx, "/v1/draw/digital", queryParams, nil)
}

// ==================== 数据模型 ====================

// BetOption 投注选项
type BetOption struct {
	// 投注码
	BetCode string `json:"betCode"`
	// 赔率
	Odds float64 `json:"odds"`
	// 赔率变化: 0 不变, 1 上升, -1 下降
	Change string `json:"change"`
}

// BetInfo 投注信息
type BetInfo struct {
	// 子玩法
	SubType SubType `json:"subType"`
	// 支持过关方式: 0-仅单关, 1-仅过关, 2-单关和过关, 3-未开售
	SupportType string `json:"supportType"`
	// 让球/让分数
	Handicap float64 `json:"handicap"`
	// 投注选项列表
	Options []*BetOption `json:"options"`
}

// MatchInfo 赛事信息
type MatchInfo struct {
	// 赛事ID
	MatchId string `json:"matchId"`
	// 运动类型
	SportId SportId `json:"sportId"`
	// 比赛时间
	MatchTimeStr string `json:"matchTimeStr"`
	// 状态
	Status Status `json:"status"`
	// 状态码
	StatusCode StatusCode `json:"statusCode"`
	// 状态名
	StatusName string `json:"statusName"`
	// 主队比分
	HomeTeamScore *MatchScore `json:"homeTeamScore"`
	// 客队比分
	AwayTeamScore *MatchScore `json:"awayTeamScore"`
}

// LotteryInfo 彩种信息
type LotteryInfo struct {
	// 彩种类型
	LotteryType LotteryType `json:"lotteryType"`
	// 场次编码
	MatchCode string `json:"matchCode"`
	// 期次
	Issue string `json:"issue"`
	// 轮次
	Round string `json:"round"`
	// 投注截止时间
	BetEndTimeStr string `json:"betEndTimeStr"`
	// 是否反转显示
	Reverse bool `json:"reverse"`
	// 投注信息列表
	BetInfos []BetInfo `json:"betInfos"`
	// 北单开奖赔率
	SpValue float64 `json:"spValue"`
	// 是否停售
	IsStopSell bool `json:"isStopSell"`
}

// MatchBetItem 赛事投注详细信息
type MatchBetItem struct {
	// 赛事信息
	MatchInfo *MatchInfo `json:"matchInfo"`
	// 彩种信息
	LotteryInfo *LotteryInfo `json:"lotteryInfo"`
	// 联赛信息
	TournamentInfo *TournamentInfo `json:"tournamentInfo"`
	// 主队信息
	HomeTeamInfo *TeamInfo `json:"homeTeamInfo"`
	// 客队信息
	AwayTeamInfo *TeamInfo `json:"awayTeamInfo"`
}

// MatchBetList 赛事投注列表响应（北单、竞足、竞篮）
type MatchBetList struct {
	// 赛事投注列表
	List []*MatchBetItem `json:"list"`
}

// MatchBetInfo 赛事投注信息响应（北单、竞足、竞篮）
type MatchBetInfo struct {
	// 赛事投注信息
	Info *MatchBetItem `json:"info"`
}

// CtzcMatchBetItem 传统足彩投注项
type CtzcMatchBetItem struct {
	// 赛事编号
	MatchNum string `json:"matchNum"`
	// 比赛日期
	StartDateStr string `json:"startDateStr"`
	// 联赛名称
	TournamentName string `json:"tournamentName"`
	// 主队全称
	HomeTeamAllName string `json:"homeTeamAllName"`
	// 主队名称
	HomeTeamName string `json:"homeTeamName"`
	// 客队全称
	AwayTeamAllName string `json:"awayTeamAllName"`
	// 客队名称
	AwayTeamName string `json:"awayTeamName"`
	// 主胜赔率
	HOdds float64 `json:"hOdds"`
	// 平局赔率
	DOdds float64 `json:"dOdds"`
	// 客胜赔率
	AOdds float64 `json:"aOdds"`
}

// CtzcResultPrizeLevel 传统足彩奖等
type CtzcResultPrizeLevel struct {
	// 奖等
	PrizeLevel string `json:"prizeLevel"`
	// 中奖注数
	StakeCount string `json:"stakeCount"`
	// 单注奖金
	StakeAmount string `json:"stakeAmount"`
	// 总奖金
	TotalPrizeAmount string `json:"totalPrizeAmount"`
}

// CtzcBetInfo 传统足彩投注信息
type CtzcBetInfo struct {
	// 期次
	Issue string `json:"issue"`
	// 销售开始时间
	SaleBeginTimeStr string `json:"saleBeginTimeStr"`
	// 销售结束时间
	SaleEndTimeStr string `json:"saleEndTimeStr"`
	// 投注赛事列表
	MatchList []*CtzcMatchBetItem `json:"matchList"`
	// 开奖结果
	DrawResult []string `json:"drawResult"`
	// 奖等列表
	PrizeLevelList []*CtzcResultPrizeLevel `json:"prizeLevelList"`
}

// CtzcBetList 传统足彩投注列表响应
type CtzcBetList struct {
	// 传统足彩在售期次列表
	List []*CtzcBetInfo `json:"list"`
}

// DigitalResultPrizeLevel 数字彩奖等
type DigitalResultPrizeLevel struct {
	// 奖等
	PrizeLevel string `json:"prizeLevel"`
	// 中奖注数
	StakeCount string `json:"stakeCount"`
	// 单注奖金
	StakeAmount float64 `json:"stakeAmount"`
	// 总奖金
	TotalPrizeAmount float64 `json:"totalPrizeAmount"`
}

// DigitalBetInfo 数字彩投注信息
type DigitalBetInfo struct {
	// 期次
	Issue string `json:"issue"`
	// 销售结束时间
	SaleEndTimeStr string `json:"saleEndTimeStr"`
	// 开奖结果
	DrawResult []string `json:"drawResult"`
	// 奖等列表
	PrizeLevelList []*DigitalResultPrizeLevel `json:"prizeLevelList"`
	// 是否已推送彩果
	IsResultPushed bool `json:"isResultPushed"`
}

// DigitalBetList 数字彩投注列表响应
type DigitalBetList struct {
	// 数字彩在售期次列表
	List []*DigitalBetInfo `json:"list"`
}

// GyjBetOption 冠亚军投注可选项
type GyjBetOption struct {
	// 投注编号
	Num string `json:"num"`
	// 投注赔率
	Odds float64 `json:"odds"`
	// 是否可投注
	IsEnabled bool `json:"isEnabled"`
	// 冠军队伍名称
	HomeTeamName string `json:"homeTeamName"`
	// 冠军队伍图标
	HomeTeamPic string `json:"homeTeamPic"`
	// 亚军队伍名称（仅冠亚军）
	AwayTeamName string `json:"awayTeamName"`
	// 亚军队伍图标（仅冠亚军）
	AwayTeamPic string `json:"awayTeamPic"`
}

// GyjBetInfo 冠亚军投注信息
type GyjBetInfo struct {
	// 期次
	Issue string `json:"issue"`
	// 赛事中文名称
	CNName string `json:"cnName"`
	// 投注可选项列表
	Options []*GyjBetOption `json:"options"`
	// 彩果投注编号
	ResultNum string `json:"resultNum"`
	// 彩果赔率
	ResultOdds float64 `json:"resultOdds"`
	// 销售结束时间
	SaleEndTimeStr string `json:"saleEndTimeStr"`
}

// LiveMatchInfo 比分直播赛事信息
type LiveMatchInfo struct {
	// 赛事ID
	MatchId string `json:"matchId"`
	// 运动类型
	SportId SportId `json:"sportId"`
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
	Status Status `json:"status"`
	// 状态码
	StatusCode StatusCode `json:"statusCode"`
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
	Id string `json:"id"`
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
	Id string `json:"id"`
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
	Id string `json:"id"`
	// 中文名称
	CnName string `json:"cnName"`
	// 英文名称
	EnName string `json:"enName"`
}

// StadiumInfo 场馆信息
type StadiumInfo struct {
	// 场馆ID
	Id string `json:"id"`
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

// LiveMatchList 赛事列表响应
type LiveMatchList struct {
	// 赛事列表
	List []*LiveMatchInfo `json:"list"`
}

// LiveMatchLotteryMapping 彩种赛事映射响应
type LiveMatchLotteryMapping struct {
	// 赛事ID
	MatchId string `json:"matchId"`
}

// LiveMatchHistory 历史对阵响应
type LiveMatchHistory struct {
	// 两队交锋
	BothHistoryList []*HistoryMatchItem `json:"bothHistoryList"`
	// 主队近期比赛
	HomeHistoryList []*HistoryMatchItem `json:"homeHistoryList"`
	// 客队近期比赛
	AwayHistoryList []*HistoryMatchItem `json:"awayHistoryList"`
}

// HistoryMatchItem 历史对阵信息
type HistoryMatchItem struct {
	// 比赛ID
	MatchId string `json:"matchId"`
	// 比赛时间
	MatchTimeStr string `json:"matchTimeStr"`
	// 联赛ID
	TournamentId string `json:"tournamentId"`
	// 联赛名称
	TournamentName string `json:"tournamentName"`
	// 主队ID
	HomeTeamId string `json:"homeTeamId"`
	// 主队名称
	HomeTeamName string `json:"homeTeamName"`
	// 主队全场比分
	HomeTeamScore string `json:"homeTeamScore"`
	// 主队正常时间比分
	HomeTeamNormalTimeScore string `json:"homeTeamNormalTimeScore"`
	// 主队半场比分
	HomeTeamHalfTimeScore string `json:"homeTeamHalfTimeScore"`
	// 客队ID
	AwayTeamId string `json:"awayTeamId"`
	// 客队名称
	AwayTeamName string `json:"awayTeamName"`
	// 客队全场比分
	AwayTeamScore string `json:"awayTeamScore"`
	// 客队正常时间比分
	AwayTeamNormalTimeScore string `json:"awayTeamNormalTimeScore"`
	// 客队半场比分
	AwayTeamHalfTimeScore string `json:"awayTeamHalfTimeScore"`
}

// LiveMatchLineup 阵容响应
type LiveMatchLineup struct {
	// 主队阵容
	HomeLineup *MatchLineup `json:"homeLineup,omitempty"`
	// 客队阵容
	AwayLineup *MatchLineup `json:"awayLineup,omitempty"`
}

// MatchLineup 赛事阵容
type MatchLineup struct {
	// 主教练
	Coach *MatchLineupItem `json:"coach,omitempty"`
	// 首发列表
	LineupList []*MatchLineupItem `json:"lineupList"`
	// 替补列表
	SubstituteList []*MatchLineupItem `json:"substituteList"`
}

// MatchLineupItem 阵容项
type MatchLineupItem struct {
	// 球员ID
	PlayerId string `json:"playerId"`
	// 球员姓名
	PlayerName string `json:"playerName"`
	// 球员头像地址
	PlayerAvatarURL string `json:"playerAvatarUrl"`
	// 球衣编号
	ShirtNumber string `json:"shirtNumber"`
	// 阵容位置
	Position string `json:"position"`
	// 比赛位置
	MatchPosition string `json:"matchPosition"`
	// 比赛位置名称
	MatchPositionName string `json:"matchPositionName"`
}

// LiveMatchNews 新闻情报响应
type LiveMatchNews struct {
	// 主队情报
	HomeNews *MatchNews `json:"homeNews,omitempty"`
	// 中立情报列表
	NeutralNewsList []string `json:"neutralNewsList"`
	// 客队情报
	AwayNews *MatchNews `json:"awayNews,omitempty"`
}

// MatchNews 赛事情报
type MatchNews struct {
	// 有利情报
	Benefit []string `json:"benefit"`
	// 不利情报
	NoBenefit []string `json:"noBenefit"`
}

// LiveMatchIntelligenceNews 专属情报响应
type LiveMatchIntelligenceNews struct {
	// 专属情报列表
	List []*MatchIntelligenceNewsItem `json:"list"`
}

// MatchIntelligenceNewsItem 赛事专属情报内容
type MatchIntelligenceNewsItem struct {
	// 标题
	Title string `json:"title"`
	// 情报内容
	Content string `json:"content"`
	// 所属球队 0-公共 1-主队 2-客队
	Team string `json:"team"`
	// 标签名称
	TagName string `json:"tagName"`
	// 标签颜色
	TagColor string `json:"tagColor"`
	// 图片链接
	ImageURL string `json:"imageUrl"`
	// 图片宽度
	ImageWidth string `json:"imageWidth"`
	// 图片高度
	ImageHeight string `json:"imageHeight"`
}

// LiveMatchOdds 赔率响应
type LiveMatchOdds struct {
	// 百家盘让分胜负（亚赔）
	RfsfList []*MatchOddsRfsfItem `json:"rfsfList"`
	// 百家盘胜平负（欧赔）
	SpfList []*MatchOddsSpfItem `json:"spfList"`
	// 百家盘大小分（大小球）
	DxfList []*MatchOddsDxfItem `json:"dxfList"`
	// 百家盘凯利指数
	KlList []*MatchOddsKlItem `json:"klList"`
}

// MatchOddsRfsfItem 百家盘让分胜负（亚赔）
type MatchOddsRfsfItem struct {
	// 公司ID
	BookId string `json:"bookId"`
	// 公司名称
	BookName string `json:"bookName"`
	// 初始让球数
	Handicap0 float64 `json:"handicap0"`
	// 初始胜赔率
	Win0 float64 `json:"win0"`
	// 初始负赔率
	Lose0 float64 `json:"lose0"`
	// 即时让球数
	Handicap float64 `json:"handicap"`
	// 即时胜赔率
	Win float64 `json:"win"`
	// 即时负赔率
	Lose float64 `json:"lose"`
}

// MatchOddsSpfItem 百家盘胜平负（欧赔）
type MatchOddsSpfItem struct {
	// 公司ID
	BookId string `json:"bookId"`
	// 公司名称
	BookName string `json:"bookName"`
	// 初始胜赔率
	Win0 float64 `json:"win0"`
	// 初始平赔率
	Draw0 float64 `json:"draw0"`
	// 初始负赔率
	Lose0 float64 `json:"lose0"`
	// 即时胜赔率
	Win float64 `json:"win"`
	// 即时平赔率
	Draw float64 `json:"draw"`
	// 即时负赔率
	Lose float64 `json:"lose"`
}

// MatchOddsDxfItem 百家盘大小分（大小球）
type MatchOddsDxfItem struct {
	// 公司ID
	BookId string `json:"bookId"`
	// 公司名称
	BookName string `json:"bookName"`
	// 初始进球数
	Handicap0 float64 `json:"handicap0"`
	// 初始大球赔率
	Over0 float64 `json:"over0"`
	// 初始小球赔率
	Under0 float64 `json:"under0"`
	// 即时进球数
	Handicap float64 `json:"handicap"`
	// 即时大球赔率
	Over float64 `json:"over"`
	// 即时小球赔率
	Under float64 `json:"under"`
}

// MatchOddsKlItem 百家盘凯利指数
type MatchOddsKlItem struct {
	// 公司ID
	BookId string `json:"bookId"`
	// 公司名称
	BookName string `json:"bookName"`
	// 初始胜赔率指数
	Win0 float64 `json:"win0"`
	// 初始平赔率指数
	Draw0 float64 `json:"draw0"`
	// 初始负赔率指数
	Lose0 float64 `json:"lose0"`
	// 初始返利指数
	Revenue0 float64 `json:"revenue0"`
	// 即时胜赔率指数
	Win float64 `json:"win"`
	// 即时平赔率指数
	Draw float64 `json:"draw"`
	// 即时负赔率指数
	Lose float64 `json:"lose"`
	// 即时返利指数
	Revenue float64 `json:"revenue"`
}

// LiveMatchPhaseAll 直播全数据响应
type LiveMatchPhaseAll struct {
	// 直播列表
	List []*MatchPhaseAllItem `json:"list"`
}

// MatchPhaseAllItem 赛事事件直播
type MatchPhaseAllItem struct {
	// 足球比赛的已进行时间（秒）
	TimePlayed string `json:"timePlayed"`
	// 篮球比赛当前节的剩余时间（秒）
	TimeRemain string `json:"timeRemain"`
	// 补时阶段的第几秒
	TimeInjury string `json:"timeInjury"`
	// 所属球队 0-公共 1-主队 2-客队
	Team string `json:"team"`
	// 比赛状态码
	StatusCode StatusCode `json:"statusCode"`
	// 比赛状态名称
	StatusName string `json:"statusName"`
	// 事件类型ID
	TypeId string `json:"typeId"`
	// 事件名称
	TypeName string `json:"typeName"`
	// 文字直播内容
	Content string `json:"content"`
	// 事件发生时比分
	Score string `json:"score"`
}

// LiveMatchPhaseEvent 事件直播响应
type LiveMatchPhaseEvent struct {
	// 直播列表
	List []*MatchPhaseEventItem `json:"list"`
}

// MatchPhaseEventItem 赛事事件直播
type MatchPhaseEventItem struct {
	// 事件发生时间（比赛进行的秒数）
	Time string `json:"time"`
	// 篮球比赛当前节的剩余时间（秒）
	TimeRemain string `json:"timeRemain"`
	// 补时阶段的第几分钟（分钟）
	TimeInjury string `json:"timeInjury"`
	// 所属球队 0-公共 1-主队 2-客队
	Team string `json:"team"`
	// 事件类型ID
	TypeId string `json:"typeId"`
	// 事件名称
	TypeName string `json:"typeName"`
	// 事件发生后的比分
	Scores string `json:"scores"`
	// 关联球员ID 表示换入或进球球员ID
	PlayerId string `json:"playerId"`
	// 球员名字
	PlayerName string `json:"playerName"`
	// 球员头像
	PlayerPicURL string `json:"playerPicUrl"`
	// 关联球员ID2,表示换出或进球助攻球员ID
	PlayerId2 string `json:"playerId2"`
	// 球员名字2
	PlayerName2 string `json:"playerName2"`
	// 球员头像2
	PlayerPicURL2 string `json:"playerPicUrl2"`
}

// LiveMatchPhaseText 文字直播响应
type LiveMatchPhaseText struct {
	// 直播列表
	List []*MatchPhaseTextItem `json:"list"`
}

// MatchPhaseTextItem 赛事文字直播
type MatchPhaseTextItem struct {
	// 时间 足球：进行分钟数 篮球：当前节剩余时间(分:秒)
	Time string `json:"time"`
	// 所属球队 0-公共 1-主队 2-客队
	Team string `json:"team"`
	// 文字直播内容
	Content string `json:"content"`
}

// LiveMatchScoreDetail 比分详情响应
type LiveMatchScoreDetail struct {
	// 主队比分
	Home *LiveMatchScore `json:"home,omitempty"`
	// 客队比分
	Away *LiveMatchScore `json:"away,omitempty"`
}

// LiveMatchScore 比分直播赛事比分
type LiveMatchScore struct {
	// 当前总比分
	Current string `json:"current"`
	// 常规时间比分（仅足球90分钟）
	NormalTime string `json:"normalTime"`
	// 含加时比分
	Overtime string `json:"overtime"`
	// 篮球第一节比分，足球上半场比分
	Period1 string `json:"period1"`
	// 篮球第二节比分，足球下半场比分
	Period2 string `json:"period2"`
	// 篮球第三节比分
	Period3 string `json:"period3"`
	// 篮球第四节比分
	Period4 string `json:"period4"`
}

// LiveMatchStaticDetail 数据统计响应
type LiveMatchStaticDetail struct {
	// 主队数据统计
	HomeStatic *LiveMatchStatic `json:"homeStatic,omitempty"`
	// 客队数据统计
	AwayStatic *LiveMatchStatic `json:"awayStatic,omitempty"`
}

// LiveMatchStatic 比分直播赛事数据统计
type LiveMatchStatic struct {
	// 控球率
	BallPossession string `json:"ballPossession"`
	// 犯规
	Fouls string `json:"fouls"`
	// 进攻
	Attack string `json:"attack"`
	// 换人
	Substitutions string `json:"substitutions"`
	// 长暂停
	Timeouts string `json:"timeouts"`
	// 角球
	CornerKicks string `json:"cornerKicks"`
	// 传球
	Pass string `json:"pass"`
	// 任意球
	FreeKicks string `json:"freeKicks"`
	// 球门球
	GoalKeeperKicks string `json:"goalKeeperKicks"`
	// 扑救
	GoalKeeperSaves string `json:"goalKeeperSaves"`
	// 射门
	GoalKicks string `json:"goalKicks"`
	// 越位
	Offsides string `json:"offsides"`
	// 射偏
	ShotsOffGoal string `json:"shotsOffGoal"`
	// 射正
	ShotsOnGoal string `json:"shotsOnGoal"`
	// 界外球
	OutOfBounds string `json:"outOfBounds"`
	// 黄牌
	Yellow string `json:"yellow"`
	// 红牌
	Red string `json:"red"`
	// 黄红牌
	YellowRed string `json:"yellowRed"`
	// 危险进攻
	DangerAttack string `json:"dangerAttack"`
	// 伤停
	InjuriesPeriod string `json:"injuriesPeriod"`
	// 安全球率
	BallSafePercentage string `json:"ballSafePercentage"`
	// 点球次数
	Penalties string `json:"penalties"`
	// 点球进球
	PenaltiesMade string `json:"penaltiesMade"`
	// 射门被封
	ShotsBlocked string `json:"shotsBlocked"`
	// 安全球
	BallSafe string `json:"ballSafe"`
	// 危险任意球
	DangerFreeKicks string `json:"dangerFreeKicks"`
	// 直接任意球
	DirectFreeKicks string `json:"directFreeKicks"`
	// 间接任意球
	IndirectFreeKicks string `json:"indirectFreeKicks"`
	// 直接任意球射失
	DirectFreeKicksOff string `json:"directFreeKicksOff"`
	// 得分机会
	BigChances string `json:"bigChances"`
	// 错失得分机会
	BigChancesMissed string `json:"bigChancesMissed"`
	// 快速反击
	CounterAttacks string `json:"counterAttacks"`
	// 禁区内射门
	ShotsInsideBox string `json:"shotsInsideBox"`
	// 禁区外射门
	ShotsOutsideBox string `json:"shotsOutsideBox"`
	// 准确传球
	AccuratePasses string `json:"accuratePasses"`
	// 长传次数
	LongBalls string `json:"longBalls"`
	// 长传成功次数
	LongBallsSuccess string `json:"longBallsSuccess"`
	// 传中次数
	Crosses string `json:"crosses"`
	// 传中成功次数
	CrossesSuccess string `json:"crossesSuccess"`
	// 盘带次数
	Dribbles string `json:"dribbles"`
	// 盘带成功次数
	DribblesSuccess string `json:"dribblesSuccess"`
	// 丢失球权
	PossessionLost string `json:"possessionLost"`
	// 1对1成功
	DuelsWon string `json:"duelsWon"`
	// 头球
	AerialsWon string `json:"aerialsWon"`
	// 铲球次数
	Tackles string `json:"tackles"`
	// 铲球成功次数
	TacklesSuccess string `json:"tacklesSuccess"`
	// 断球
	Interceptions string `json:"interceptions"`
	// 解围
	Clearances string `json:"clearances"`
	// 投篮
	FieldGoalAttempted string `json:"fieldGoalAttempted"`
	// 投篮命中
	FieldGoalMade string `json:"fieldGoalMade"`
	// 三分
	ThreePointAttempted string `json:"threePointAttempted"`
	// 三分命中
	ThreePointMade string `json:"threePointMade"`
	// 罚球
	FreeThrowAttempted string `json:"freeThrowAttempted"`
	// 罚球命中
	FreeThrowMade string `json:"freeThrowMade"`
	// 罚球未中
	FreeThrowOff string `json:"freeThrowOff"`
	// 得分
	Point string `json:"point"`
	// 篮板
	Rebound string `json:"rebound"`
	// 进攻篮板
	OffensiveRebound string `json:"offensiveRebound"`
	// 防守篮板
	DefensiveRebound string `json:"defensiveRebound"`
	// 助攻
	Assist string `json:"assist"`
	// 抢断
	Steal string `json:"steal"`
	// 盖帽
	Block string `json:"block"`
	// 失误
	TurnOver string `json:"turnOver"`
	// 剩余暂停机会
	TimeoutsChance string `json:"timeoutsChance"`
	// 二分
	TwoPointAttempted string `json:"twoPointAttempted"`
	// 二分命中
	TwoPointMade string `json:"twoPointMade"`
	// 连续最高分数
	HighestScore string `json:"highestScore"`
	// 最多领先分数
	LeadingScore string `json:"leadingScore"`
	// 领先时间
	LeadingTime string `json:"leadingTime"`
}

// LiveMatchPlayerStaticDetail 球员统计响应
type LiveMatchPlayerStaticDetail struct {
	// 主队球员统计列表
	HomeList []*LiveMatchPlayerStatic `json:"homeList"`
	// 客队球员统计列表
	AwayList []*LiveMatchPlayerStatic `json:"awayList"`
}

// LiveMatchPlayerStatic 球员技术统计
type LiveMatchPlayerStatic struct {
	// 球员ID
	PlayerId string `json:"playerId"`
	// 篮球相关统计
	BMinutesPlayed         string `json:"bMinutesPlayed"`
	BPoints                string `json:"bPoints"`
	BFreeThrows            string `json:"bFreeThrows"`
	BFreeThrowsMade        string `json:"bFreeThrowsMade"`
	BTwoPointers           string `json:"bTwoPointers"`
	BTwoPointersMade       string `json:"bTwoPointersMade"`
	BThreePointers         string `json:"bThreePointers"`
	BThreePointersMade     string `json:"bThreePointersMade"`
	BFieldGoals            string `json:"bFieldGoals"`
	BFieldGoalsMade        string `json:"bFieldGoalsMade"`
	BRebounds              string `json:"bRebounds"`
	BDefensiveRebounds     string `json:"bDefensiveRebounds"`
	BOffensiveRebounds     string `json:"bOffensiveRebounds"`
	BAssists               string `json:"bAssists"`
	BTurnovers             string `json:"bTurnovers"`
	BSteals                string `json:"bSteals"`
	BBlocks                string `json:"bBlocks"`
	BPersonalFouls         string `json:"bPersonalFouls"`
	BPlusMinus             string `json:"bPlusMinus"`
	BDunkShot              string `json:"bDunkShot"`
	BPersonalFoulsReceived string `json:"bPersonalFoulsReceived"`
	// 足球相关统计
	FMinutesPlayed            string `json:"fMinutesPlayed"`
	FGoals                    string `json:"fGoals"`
	FTotalShots               string `json:"fTotalShots"`
	FOnTargetScoringAttempt   string `json:"fOnTargetScoringAttempt"`
	FShotOffTarget            string `json:"fShotOffTarget"`
	FBlockedScoringAttempt    string `json:"fBlockedScoringAttempt"`
	FFreeKickGoals            string `json:"fFreeKickGoals"`
	FSuccessfulDribbles       string `json:"fSuccessfulDribbles"`
	FDribblesTotal            string `json:"fDribblesTotal"`
	FTotalAllShots            string `json:"fTotalAllShots"`
	FInterceptionWon          string `json:"fInterceptionWon"`
	FClearance                string `json:"fClearance"`
	FTotalTackle              string `json:"fTotalTackle"`
	FBlockedShots             string `json:"fBlockedShots"`
	FOwnGoals                 string `json:"fOwnGoals"`
	FGoalAssist               string `json:"fGoalAssist"`
	FTouches                  string `json:"fTouches"`
	FKeyPasses                string `json:"fKeyPasses"`
	FTotalLongBalls           string `json:"fTotalLongBalls"`
	FAccurateLongBalls        string `json:"fAccurateLongBalls"`
	FTotalCrosses             string `json:"fTotalCrosses"`
	FAccurateCrosses          string `json:"fAccurateCrosses"`
	FTotalPasses              string `json:"fTotalPasses"`
	FAccuratePasses           string `json:"fAccuratePasses"`
	FChippedPasses            string `json:"fChippedPasses"`
	FAccurateChippedPasses    string `json:"fAccurateChippedPasses"`
	FSaves                    string `json:"fSaves"`
	FSavedShotsFromInsideBox  string `json:"fSavedShotsFromInsideBox"`
	FGoodHighClaim            string `json:"fGoodHighClaim"`
	FPunches                  string `json:"fPunches"`
	FTotalKeeperSweeper       string `json:"fTotalKeeperSweeper"`
	FAccurateKeeperSweeper    string `json:"fAccurateKeeperSweeper"`
	FConcededFromInsideBox    string `json:"fConcededFromInsideBox"`
	FConcededFromOutsideBox   string `json:"fConcededFromOutsideBox"`
	FSavesParried             string `json:"fSavesParried"`
	FSavesCaught              string `json:"fSavesCaught"`
	FPenaltiesSaved           string `json:"fPenaltiesSaved"`
	FSavedShotsFromOutsideBox string `json:"fSavedShotsFromOutsideBox"`
	FOffsides                 string `json:"fOffsides"`
	FTotalDuel                string `json:"fTotalDuel"`
	FWonDuel                  string `json:"fWonDuel"`
	FTotalAerial              string `json:"fTotalAerial"`
	FWonAerial                string `json:"fWonAerial"`
	FFoul                     string `json:"fFoul"`
	FWasFouled                string `json:"fWasFouled"`
	FPossessionLostControl    string `json:"fPossessionLostControl"`
	FDispossessed             string `json:"fDispossessed"`
	FTotalDuelsWon            string `json:"fTotalDuelsWon"`
	FTotalDuels               string `json:"fTotalDuels"`
}

// LotteryDrawHomeList 竞彩开奖大厅响应
type LotteryDrawHomeList struct {
	// 竞彩开奖大厅列表信息
	List []*LotteryDrawHomeInfo `json:"list"`
}

// DrawHistoryList 竞彩开奖历史列表响应
type DrawHistoryList struct {
	// 竞彩足球历史开奖结果列表
	List []*DrawInfo `json:"list"`
}

// DigitalResultInfo 数字彩彩果详情响应
type DigitalResultInfo struct {
	// 彩种类型
	LotteryType LotteryType `json:"lotteryType"`
	// 期次
	Issue string `json:"issue"`
	// 开奖号码
	Numbers []string `json:"numbers"`
	// 开奖号码2
	Numbers2 []string `json:"numbers2"`
	// 开奖日期 yyyy-MM-dd
	DrawDateStr string `json:"drawDateStr"`
	// 销量
	Sales float64 `json:"sales"`
	// 奖池
	PoolMoney float64 `json:"poolMoney"`
	// 奖项级别列表,详情接口才返回
	PrizeLevelList []*PrizeLevelInfo `json:"prizeLevelList"`
}

// DigitalHomeResultList 数字彩列表响应
type DigitalHomeResultList struct {
	// 数字彩各彩种最近开奖结果列表
	List []*DigitalResultInfo `json:"list"`
}

// DigitalResultList 数字彩历史列表响应
type DigitalResultList struct {
	// 数字彩开奖结果列表
	List []*DigitalResultInfo `json:"list"`
}

// CtzcResultInfo 传统足彩赛果详情响应
type CtzcResultInfo struct {
	// 彩种类型
	// 171 - 6场半全场
	// 172 - 胜负游戏
	// 173 - 任选9场
	// 174 - 4场进球
	LotteryType LotteryType `json:"lotteryType"`
	// 期次
	Issue string `json:"issue"`
	// 开奖结果
	Numbers []string `json:"numbers"`
	// 开奖日期 yyyy-MM-dd
	DrawDateStr string `json:"drawDateStr"`
	// 开售时间 yyyy-MM-dd HH:mm:ss
	SaleBeginTimeStr string `json:"saleBeginTimeStr"`
	// 停售时间 yyyy-MM-dd HH:mm:ss
	SaleEndTimeStr string `json:"saleEndTimeStr"`
	// 销量
	Sales float64 `json:"sales"`
	// 奖池
	PoolMoney float64 `json:"poolMoney"`
	// 奖项级别列表
	PrizeLevelList []*PrizeLevelInfo `json:"prizeLevelList"`
}

// CtzcHomeResultList 传统足彩列表响应
type CtzcHomeResultList struct {
	// 传统足彩各彩种最近开奖结果列表
	List []*CtzcResultInfo `json:"list"`
}

// CtzcResultList 传统足彩历史列表响应
type CtzcResultList struct {
	// 传统足彩开奖结果列表
	List []*CtzcResultInfo `json:"list"`
}

// LotteryDrawHomeInfo 竞彩开奖大厅信息
type LotteryDrawHomeInfo struct {
	// 彩种类型
	LotteryType LotteryType `json:"lotteryType"`
	// 可选场数
	OptionCount string `json:"optionCount"`
	// 最新开奖信息
	LastDrawInfo DrawInfo `json:"lastDrawInfo"`
}

// DrawInfo 竞彩开奖具体信息
type DrawInfo struct {
	// 赛事ID
	MatchID string `json:"matchId"`
	// 场次编码
	MatchCode string `json:"matchCode"`
	// 周几：周一-周日
	WeekDay string `json:"weekDay"`
	// 轮次
	Round string `json:"round"`
	// 是否单关
	IsSingle bool `json:"isSingle"`
	// 比赛时间
	MatchTimeStr string `json:"matchTimeStr"`
	// 联赛信息
	TournamentInfo TournamentInfo `json:"tournamentInfo"`
	// 主队信息
	HomeTeamInfo TeamInfo `json:"homeTeamInfo"`
	// 客队信息
	AwayTeamInfo TeamInfo `json:"awayTeamInfo"`
	// 开奖是否有效：true 已开奖，false 无效场次
	IsValid bool `json:"isValid"`
	// 主队比分（无效场次时为空）
	HomeTeamScore *MatchScore `json:"homeTeamScore,omitempty"`
	// 客队比分（无效场次时为空）
	AwayTeamScore *MatchScore `json:"awayTeamScore,omitempty"`
	// 玩法开奖信息列表
	GameDrawList []*GameDrawInfo `json:"gameDrawList"`
}

// GameDrawInfo 玩法开奖信息
type GameDrawInfo struct {
	// 玩法类型
	// 竞彩足球子玩法
	// 1 - 让球胜平负
	// 2 - 总进球
	// 3 - 比分
	// 4 - 半全场
	// 6 - 胜平负
	// 竞彩篮球子玩法
	// 1 - 胜负
	// 2 - 让分胜负
	// 3 - 胜分差
	// 4 - 大小分
	SubType SubType `json:"subType"`
	// 让球/让分数
	Handicap float64 `json:"handicap"`
	// 投注码
	BetCode string `json:"betCode"`
	// 赔率
	Odds float64 `json:"odds"`
}

// PrizeLevelInfo 奖项级别信息
type PrizeLevelInfo struct {
	// 奖项级别名称
	PrizeLevelName string `json:"prizeLevelName"`
	// 中奖注数
	StakeCount string `json:"stakeCount"`
	// 单注奖金
	StakeAmount float64 `json:"stakeAmount"`
}

// ==================== Callback - 彩果回调 ====================
// 注意：以下接口不是开放平台提供给合作方调用的接口，
// 而是合作方需要实现的回调接口，由开放平台主动调用通知彩果信息。

// CallbackRequest 彩果回调请求参数
type CallbackRequest struct {
	// 彩种类型
	// 226 - 北京单场
	// 227 - 竞彩足球
	// 228 - 竞彩篮球
	// 171 - 6场半全场
	// 172 - 胜负游戏
	// 173 - 任选9场
	// 174 - 4场进球
	// 163 - 排列三
	// 164 - 排列五
	// 166 - 双色球
	// 167 - 福彩3D
	// 168 - 七星彩
	// 185 - 快乐8
	// 186 - 七乐彩
	// 188 - 大乐透
	// 229 - 冠亚军
	// 230 - 冠军
	LotteryType LotteryType `json:"lotteryType"`

	// 唯一键
	// 不同彩种的格式如下：
	// - 北京单场格式：玩法:场次编码
	// - 竞彩足球篮球格式：场次编码
	// - 传统足彩格式：期次
	// - 数字彩格式：期次
	// - 冠亚军格式：期次
	UnionKey string `json:"unionKey"`
}

// CallbackReply 彩果回调响应数据
type CallbackReply struct{}
