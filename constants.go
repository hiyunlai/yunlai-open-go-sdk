package yunlai

// LotteryType 彩种类型
type LotteryType = int64

const (
	LotteryTypePls  LotteryType = 163 // 排列三
	LotteryTypePlw  LotteryType = 164 // 排列五
	LotteryTypeSsq  LotteryType = 166 // 双色球
	LotteryTypeFc3d LotteryType = 167 // 福彩3D
	LotteryTypeQxc  LotteryType = 168 // 七星彩
	LotteryTypeBqc  LotteryType = 171 // 6场半全场
	LotteryTypeSfc  LotteryType = 172 // 胜负游戏
	LotteryTypeRxj  LotteryType = 173 // 任选9场
	LotteryTypeZjq  LotteryType = 174 // 4场进球
	LotteryTypeKl8  LotteryType = 185 // 快乐8
	LotteryTypeQlc  LotteryType = 186 // 七乐彩
	LotteryTypeDlt  LotteryType = 188 // 大乐透
	LotteryTypeBjdc LotteryType = 226 // 北京单场
	LotteryTypeJczq LotteryType = 227 // 竞彩足球
	LotteryTypeJclq LotteryType = 228 // 竞彩篮球
	LotteryTypeGyj  LotteryType = 229 // 冠亚军
	LotteryTypeGj   LotteryType = 230 // 冠军
)

// LotteryTypeNameMapping 彩种类型名称映射表
var LotteryTypeNameMapping = map[LotteryType]string{
	LotteryTypePls:  "排列3",
	LotteryTypePlw:  "排列5",
	LotteryTypeSsq:  "双色球",
	LotteryTypeFc3d: "福彩3D",
	LotteryTypeQxc:  "七星彩",
	LotteryTypeBqc:  "6场半全场",
	LotteryTypeSfc:  "胜负游戏",
	LotteryTypeRxj:  "任选9场",
	LotteryTypeZjq:  "4场进球",
	LotteryTypeKl8:  "快乐8",
	LotteryTypeQlc:  "七乐彩",
	LotteryTypeDlt:  "大乐透",
	LotteryTypeBjdc: "北京单场",
	LotteryTypeJczq: "竞彩足球",
	LotteryTypeJclq: "竞彩篮球",
	LotteryTypeGyj:  "冠亚军",
	LotteryTypeGj:   "冠军",
}

// SubType 竞彩足球子玩法类型
type SubType = int64

// 传统足彩子玩法
const (
	SubTypeCtzcDs SubType = 1 // 传统足彩单式
	SubTypeCtzcFs SubType = 3 // 传统足彩复式
)

// 北京单场子玩法
const (
	SubTypeBjdcSpf  SubType = 1 // 胜平负
	SubTypeBjdcZjq  SubType = 2 // 总进球数
	SubTypeBjdcSxds SubType = 3 // 上下单双
	SubTypeBjdcDcbf SubType = 4 // 单场比分
	SubTypeBjdcBqc  SubType = 5 // 半全场
	SubTypeBjdcSfgg SubType = 7 // 胜负过关
)

// 竞彩足球子玩法
const (
	SubTypeJczqRqspf SubType = 1 // 让球胜平负
	SubTypeJczqZjq   SubType = 2 // 总进球
	SubTypeJczqBf    SubType = 3 // 比分
	SubTypeJczqBqc   SubType = 4 // 半全场
	SubTypeJczqHhgg  SubType = 5 // 混合过关
	SubTypeJczqSpf   SubType = 6 // 胜平负
)

// 竞彩篮球子玩法
const (
	SubTypeJclqSf   SubType = 1 // 胜负
	SubTypeJclqRfsf SubType = 2 // 让分胜负
	SubTypeJclqSfc  SubType = 3 // 胜分差
	SubTypeJclqDxf  SubType = 4 // 大小分
	SubTypeJclqHhgg SubType = 5 // 混合过关
)

// LotterySubTypeNameMapping 彩种子玩法名称映射表
var LotterySubTypeNameMapping = map[LotteryType]map[SubType]string{
	LotteryTypeBqc: {
		SubTypeCtzcDs: "单式",
		SubTypeCtzcFs: "复式",
	},
	LotteryTypeSfc: {
		SubTypeCtzcDs: "单式",
		SubTypeCtzcFs: "复式",
	},
	LotteryTypeRxj: {
		SubTypeCtzcDs: "单式",
		SubTypeCtzcFs: "复式",
	},
	LotteryTypeZjq: {
		SubTypeCtzcDs: "单式",
		SubTypeCtzcFs: "复式",
	},
	LotteryTypeBjdc: {
		SubTypeBjdcSpf:  "胜平负",
		SubTypeBjdcZjq:  "总进球数",
		SubTypeBjdcSxds: "上下单双",
		SubTypeBjdcDcbf: "比分",
		SubTypeBjdcBqc:  "半全场",
		SubTypeBjdcSfgg: "胜负过关",
	},
	LotteryTypeJczq: {
		SubTypeJczqRqspf: "让球胜平负",
		SubTypeJczqZjq:   "总进球",
		SubTypeJczqBf:    "比分",
		SubTypeJczqBqc:   "半全场",
		SubTypeJczqHhgg:  "混合过关",
		SubTypeJczqSpf:   "胜平负",
	},
	LotteryTypeJclq: {
		SubTypeJclqSf:   "胜负",
		SubTypeJclqRfsf: "让分胜负",
		SubTypeJclqSfc:  "胜分差",
		SubTypeJclqDxf:  "大小分",
		SubTypeJclqHhgg: "混合过关",
	},
}

// MxnType 过关方式
type MxnType = string

const (
	MxnType0c0 MxnType = "0c0" // 无过关

	MxnType1c1 MxnType = "1c1" // 单关

	MxnType2c1 MxnType = "2c1" // 2串1
	MxnType2c3 MxnType = "2c3" // 2串3

	MxnType3c1 MxnType = "3c1" // 3串1
	MxnType3c3 MxnType = "3c3" // 3串3
	MxnType3c4 MxnType = "3c4" // 3串4
	MxnType3c6 MxnType = "3c6" // 3串6
	MxnType3c7 MxnType = "3c7" // 3串7

	MxnType4c1  MxnType = "4c1"  // 4串1
	MxnType4c4  MxnType = "4c4"  // 4串4
	MxnType4c5  MxnType = "4c5"  // 4串5
	MxnType4c6  MxnType = "4c6"  // 4串6
	MxnType4c10 MxnType = "4c10" // 4串10
	MxnType4c11 MxnType = "4c11" // 4串11
	MxnType4c14 MxnType = "4c14" // 4串14
	MxnType4c15 MxnType = "4c15" // 4串15

	MxnType5c1  MxnType = "5c1"  // 5串1
	MxnType5c5  MxnType = "5c5"  // 5串5
	MxnType5c6  MxnType = "5c6"  // 5串6
	MxnType5c10 MxnType = "5c10" // 5串10
	MxnType5c15 MxnType = "5c15" // 5串15
	MxnType5c16 MxnType = "5c16" // 5串16
	MxnType5c20 MxnType = "5c20" // 5串20
	MxnType5c25 MxnType = "5c25" // 5串25
	MxnType5c26 MxnType = "5c26" // 5串26
	MxnType5c30 MxnType = "5c30" // 5串30
	MxnType5c31 MxnType = "5c31" // 5串31

	MxnType6c1  MxnType = "6c1"  // 6串1
	MxnType6c6  MxnType = "6c6"  // 6串6
	MxnType6c7  MxnType = "6c7"  // 6串7
	MxnType6c15 MxnType = "6c15" // 6串15
	MxnType6c20 MxnType = "6c20" // 6串20
	MxnType6c21 MxnType = "6c21" // 6串21
	MxnType6c22 MxnType = "6c22" // 6串22
	MxnType6c35 MxnType = "6c35" // 6串35
	MxnType6c41 MxnType = "6c41" // 6串41
	MxnType6c42 MxnType = "6c42" // 6串42
	MxnType6c50 MxnType = "6c50" // 6串50
	MxnType6c56 MxnType = "6c56" // 6串56
	MxnType6c57 MxnType = "6c57" // 6串57
	MxnType6c62 MxnType = "6c62" // 6串62
	MxnType6c63 MxnType = "6c63" // 6串63

	MxnType7c1   MxnType = "7c1"   // 7串1
	MxnType7c7   MxnType = "7c7"   // 7串7
	MxnType7c8   MxnType = "7c8"   // 7串8
	MxnType7c21  MxnType = "7c21"  // 7串21
	MxnType7c35  MxnType = "7c35"  // 7串35
	MxnType7c120 MxnType = "7c120" // 7串120
	MxnType7c127 MxnType = "7c127" // 7串127

	MxnType8c1   MxnType = "8c1"   // 8串1
	MxnType8c8   MxnType = "8c8"   // 8串8
	MxnType8c9   MxnType = "8c9"   // 8串9
	MxnType8c28  MxnType = "8c28"  // 8串28
	MxnType8c56  MxnType = "8c56"  // 8串56
	MxnType8c70  MxnType = "8c70"  // 8串70
	MxnType8c247 MxnType = "8c247" // 8串247
	MxnType8c255 MxnType = "8c255" // 8串255

	MxnType9c1  MxnType = "9c1"  // 9串1
	MxnType10c1 MxnType = "10c1" // 10串1
	MxnType11c1 MxnType = "11c1" // 11串1
	MxnType12c1 MxnType = "12c1" // 12串1
	MxnType13c1 MxnType = "13c1" // 13串1
	MxnType14c1 MxnType = "14c1" // 14串1
	MxnType15c1 MxnType = "15c1" // 15串1
)

// MxnTypeNameMapping 过关方式名称映射表
var MxnTypeNameMapping = map[MxnType]string{
	MxnType0c0: "无过关",

	MxnType1c1: "单关",

	MxnType2c1: "2串1",
	MxnType2c3: "2串3",

	MxnType3c1: "3串1",
	MxnType3c3: "3串3",
	MxnType3c4: "3串4",
	MxnType3c6: "3串6",
	MxnType3c7: "3串7",

	MxnType4c1:  "4串1",
	MxnType4c4:  "4串4",
	MxnType4c5:  "4串5",
	MxnType4c6:  "4串6",
	MxnType4c10: "4串10",
	MxnType4c11: "4串11",
	MxnType4c14: "4串14",
	MxnType4c15: "4串15",

	MxnType5c1:  "5串1",
	MxnType5c5:  "5串5",
	MxnType5c6:  "5串6",
	MxnType5c10: "5串10",
	MxnType5c15: "5串15",
	MxnType5c16: "5串16",
	MxnType5c20: "5串20",
	MxnType5c25: "5串25",
	MxnType5c26: "5串26",
	MxnType5c30: "5串30",
	MxnType5c31: "5串31",

	MxnType6c1:  "6串1",
	MxnType6c6:  "6串6",
	MxnType6c7:  "6串7",
	MxnType6c15: "6串15",
	MxnType6c20: "6串20",
	MxnType6c21: "6串21",
	MxnType6c22: "6串22",
	MxnType6c35: "6串35",
	MxnType6c41: "6串41",
	MxnType6c42: "6串42",
	MxnType6c50: "6串50",
	MxnType6c56: "6串56",
	MxnType6c57: "6串57",
	MxnType6c62: "6串62",
	MxnType6c63: "6串63",

	MxnType7c1:   "7串1",
	MxnType7c7:   "7串7",
	MxnType7c8:   "7串8",
	MxnType7c21:  "7串21",
	MxnType7c35:  "7串35",
	MxnType7c120: "7串120",
	MxnType7c127: "7串127",

	MxnType8c1:   "8串1",
	MxnType8c8:   "8串8",
	MxnType8c9:   "8串9",
	MxnType8c28:  "8串28",
	MxnType8c56:  "8串56",
	MxnType8c70:  "8串70",
	MxnType8c247: "8串247",
	MxnType8c255: "8串255",

	MxnType9c1:  "9串1",
	MxnType10c1: "10串1",
	MxnType11c1: "11串1",
	MxnType12c1: "12串1",
	MxnType13c1: "13串1",
	MxnType14c1: "14串1",
	MxnType15c1: "15串1",
}

// MxnTypeCountMapping 过关方式过关数映射表
var MxnTypeCountMapping = map[MxnType]int{
	MxnType0c0: 0,

	MxnType1c1: 1,

	MxnType2c1: 2,
	MxnType2c3: 2,

	MxnType3c1: 3,
	MxnType3c3: 3,
	MxnType3c4: 3,
	MxnType3c6: 3,
	MxnType3c7: 3,

	MxnType4c1:  4,
	MxnType4c4:  4,
	MxnType4c5:  4,
	MxnType4c6:  4,
	MxnType4c10: 4,
	MxnType4c11: 4,
	MxnType4c14: 4,
	MxnType4c15: 4,

	MxnType5c1:  5,
	MxnType5c5:  5,
	MxnType5c6:  5,
	MxnType5c10: 5,
	MxnType5c15: 5,
	MxnType5c16: 5,
	MxnType5c20: 5,
	MxnType5c25: 5,
	MxnType5c26: 5,
	MxnType5c30: 5,
	MxnType5c31: 5,

	MxnType6c1:  6,
	MxnType6c6:  6,
	MxnType6c7:  6,
	MxnType6c15: 6,
	MxnType6c20: 6,
	MxnType6c21: 6,
	MxnType6c22: 6,
	MxnType6c35: 6,
	MxnType6c41: 6,
	MxnType6c42: 6,
	MxnType6c50: 6,
	MxnType6c56: 6,
	MxnType6c57: 6,
	MxnType6c62: 6,
	MxnType6c63: 6,

	MxnType7c1:   7,
	MxnType7c7:   7,
	MxnType7c8:   7,
	MxnType7c21:  7,
	MxnType7c35:  7,
	MxnType7c120: 7,
	MxnType7c127: 7,

	MxnType8c1:   8,
	MxnType8c8:   8,
	MxnType8c9:   8,
	MxnType8c28:  8,
	MxnType8c56:  8,
	MxnType8c70:  8,
	MxnType8c247: 8,
	MxnType8c255: 8,

	MxnType9c1:  9,
	MxnType10c1: 10,
	MxnType11c1: 11,
	MxnType12c1: 12,
	MxnType13c1: 13,
	MxnType14c1: 14,
	MxnType15c1: 15,
}

// MxnTypeNumsMapping 过关方式串关数映射表
var MxnTypeNumsMapping = map[MxnType][]int{
	MxnType0c0: {0},

	MxnType1c1: {1},

	MxnType2c1: {2},
	MxnType2c3: {1, 2},

	MxnType3c1: {3},
	MxnType3c3: {2},
	MxnType3c4: {2, 3},
	MxnType3c6: {1, 2},
	MxnType3c7: {1, 2, 3},

	MxnType4c1:  {4},
	MxnType4c4:  {3},
	MxnType4c5:  {3, 4},
	MxnType4c6:  {2},
	MxnType4c10: {1, 2},
	MxnType4c11: {2, 3, 4},
	MxnType4c14: {1, 2, 3},
	MxnType4c15: {1, 2, 3, 4},

	MxnType5c1:  {5},
	MxnType5c5:  {4},
	MxnType5c6:  {4, 5},
	MxnType5c10: {2},
	MxnType5c15: {1, 2},
	MxnType5c16: {3, 4, 5},
	MxnType5c20: {2, 3},
	MxnType5c25: {1, 2, 3},
	MxnType5c26: {2, 3, 4, 5},
	MxnType5c30: {1, 2, 3, 4},
	MxnType5c31: {1, 2, 3, 4, 5},

	MxnType6c1:  {6},
	MxnType6c6:  {5},
	MxnType6c7:  {5, 6},
	MxnType6c15: {2},
	MxnType6c20: {3},
	MxnType6c21: {1, 2},
	MxnType6c22: {4, 5, 6},
	MxnType6c35: {2, 3},
	MxnType6c41: {1, 2, 3},
	MxnType6c42: {3, 4, 5, 6},
	MxnType6c50: {2, 3, 4},
	MxnType6c56: {1, 2, 3, 4},
	MxnType6c57: {2, 3, 4, 5, 6},
	MxnType6c62: {1, 2, 3, 4, 5},
	MxnType6c63: {1, 2, 3, 4, 5, 6},

	MxnType7c1:   {7},
	MxnType7c7:   {6},
	MxnType7c8:   {6, 7},
	MxnType7c21:  {5},
	MxnType7c35:  {4},
	MxnType7c120: {2, 3, 4, 5, 6, 7},
	MxnType7c127: {1, 2, 3, 4, 5, 6, 7},

	MxnType8c1:   {8},
	MxnType8c8:   {7},
	MxnType8c9:   {7, 8},
	MxnType8c28:  {6},
	MxnType8c56:  {5},
	MxnType8c70:  {4},
	MxnType8c247: {2, 3, 4, 5, 6, 7, 8},
	MxnType8c255: {1, 2, 3, 4, 5, 6, 7, 8},

	MxnType9c1:  {9},
	MxnType10c1: {10},
	MxnType11c1: {11},
	MxnType12c1: {12},
	MxnType13c1: {13},
	MxnType14c1: {14},
	MxnType15c1: {15},
}

// SupportType 支持过关方式
type SupportType = int64

const (
	SupportTypeOnlyDg SupportType = 0 // 仅单关
	SupportTypeOnlyGg SupportType = 1 // 仅过关
	SupportTypeDgGg   SupportType = 2 // 单关和过关
	SupportTypeNoSale SupportType = 3 // 未开售
)

// BetCode 投注码
type BetCode = string

// BetCodeAll 全中（无效场次）
const BetCodeAll BetCode = "*"

// 投注码组：竞彩足球 - 让球胜平负
const (
	BetCodeJczqRqspfS BetCode = "1-3" // 胜
	BetCodeJczqRqspfP BetCode = "1-1" // 平
	BetCodeJczqRqspfF BetCode = "1-0" // 负
)

// 投注码组：竞彩足球 - 总进球
const (
	BetCodeJczqZjq0 BetCode = "2-0" // 0球
	BetCodeJczqZjq1 BetCode = "2-1" // 1球
	BetCodeJczqZjq2 BetCode = "2-2" // 2球
	BetCodeJczqZjq3 BetCode = "2-3" // 3球
	BetCodeJczqZjq4 BetCode = "2-4" // 4球
	BetCodeJczqZjq5 BetCode = "2-5" // 5球
	BetCodeJczqZjq6 BetCode = "2-6" // 6球
	BetCodeJczqZjq7 BetCode = "2-7" // 7球
)

// 投注码组：竞彩足球 - 比分
const (
	BetCodeJczqBfS  BetCode = "3-0"  // 胜其他
	BetCodeJczqBf10 BetCode = "3-1"  // 1:0
	BetCodeJczqBf20 BetCode = "3-2"  // 2:0
	BetCodeJczqBf21 BetCode = "3-3"  // 2:1
	BetCodeJczqBf30 BetCode = "3-4"  // 3:0
	BetCodeJczqBf31 BetCode = "3-5"  // 3:1
	BetCodeJczqBf32 BetCode = "3-6"  // 3:2
	BetCodeJczqBf40 BetCode = "3-7"  // 4:0
	BetCodeJczqBf41 BetCode = "3-8"  // 4:1
	BetCodeJczqBf42 BetCode = "3-9"  // 4:2
	BetCodeJczqBf50 BetCode = "3-10" // 5:0
	BetCodeJczqBf51 BetCode = "3-11" // 5:1
	BetCodeJczqBf52 BetCode = "3-12" // 5:2
	BetCodeJczqBfP  BetCode = "3-13" // 平其他
	BetCodeJczqBf00 BetCode = "3-14" // 0:0
	BetCodeJczqBf11 BetCode = "3-15" // 1:1
	BetCodeJczqBf22 BetCode = "3-16" // 2:2
	BetCodeJczqBf33 BetCode = "3-17" // 3:3
	BetCodeJczqBfF  BetCode = "3-18" // 负其他
	BetCodeJczqBf01 BetCode = "3-19" // 0:1
	BetCodeJczqBf02 BetCode = "3-20" // 0:2
	BetCodeJczqBf12 BetCode = "3-21" // 1:2
	BetCodeJczqBf03 BetCode = "3-22" // 0:3
	BetCodeJczqBf13 BetCode = "3-23" // 1:3
	BetCodeJczqBf23 BetCode = "3-24" // 2:3
	BetCodeJczqBf04 BetCode = "3-25" // 0:4
	BetCodeJczqBf14 BetCode = "3-26" // 1:4
	BetCodeJczqBf24 BetCode = "3-27" // 2:4
	BetCodeJczqBf05 BetCode = "3-28" // 0:5
	BetCodeJczqBf15 BetCode = "3-29" // 1:5
	BetCodeJczqBf25 BetCode = "3-30" // 2:5
)

// 投注码组：竞彩足球 - 半全场
const (
	BetCodeJczqBqcSS BetCode = "4-0" // 胜胜
	BetCodeJczqBqcSP BetCode = "4-1" // 胜平
	BetCodeJczqBqcSF BetCode = "4-2" // 胜负
	BetCodeJczqBqcPS BetCode = "4-3" // 平胜
	BetCodeJczqBqcPP BetCode = "4-4" // 平平
	BetCodeJczqBqcPF BetCode = "4-5" // 平负
	BetCodeJczqBqcFS BetCode = "4-6" // 负胜
	BetCodeJczqBqcFP BetCode = "4-7" // 负平
	BetCodeJczqBqcFF BetCode = "4-8" // 负负
)

// 投注码组：竞彩足球 - 胜平负
const (
	BetCodeJczqSpfS BetCode = "6-3" // 胜
	BetCodeJczqSpfP BetCode = "6-1" // 平
	BetCodeJczqSpfF BetCode = "6-0" // 负
)

// 投注码组：竞彩篮球 - 胜负
const (
	BetCodeJclqSfF BetCode = "1-0" // 客胜
	BetCodeJclqSfS BetCode = "1-1" // 主胜
)

// 投注码组：竞彩篮球 - 让分胜负
const (
	BetCodeJclqRfsfF BetCode = "2-0" // 客胜
	BetCodeJclqRfsfS BetCode = "2-1" // 主胜
)

// 投注码组：竞彩篮球 - 胜分差
const (
	BetCodeJclqSfcF15   BetCode = "3-0"  // 客胜(1-5)
	BetCodeJclqSfcF610  BetCode = "3-1"  // 客胜(6-10)
	BetCodeJclqSfcF1115 BetCode = "3-2"  // 客胜(11-15)
	BetCodeJclqSfcF1620 BetCode = "3-3"  // 客胜(16-20)
	BetCodeJclqSfcF2125 BetCode = "3-4"  // 客胜(21-25)
	BetCodeJclqSfcF26   BetCode = "3-5"  // 客胜(26+)
	BetCodeJclqSfcS15   BetCode = "3-6"  // 主胜(1-5)
	BetCodeJclqSfcS610  BetCode = "3-7"  // 主胜(6-10)
	BetCodeJclqSfcS1115 BetCode = "3-8"  // 主胜(11-15)
	BetCodeJclqSfcS1620 BetCode = "3-9"  // 主胜(16-20)
	BetCodeJclqSfcS2125 BetCode = "3-10" // 主胜(21-25)
	BetCodeJclqSfcS26   BetCode = "3-11" // 主胜(26+)
)

// 投注码组：竞彩篮球 - 让分胜负
const (
	BetCodeJclqDxfD BetCode = "4-0" // 大分
	BetCodeJclqDxfX BetCode = "4-1" // 小分
)

// 投注码组：北京单场 - 胜平负
const (
	BetCodeBjdcSpfS BetCode = "1-3" // 胜
	BetCodeBjdcSpfP BetCode = "1-1" // 平
	BetCodeBjdcSpfF BetCode = "1-0" // 负
)

// 投注码组：北京单场 - 总进球
const (
	BetCodeBjdcZjq0 BetCode = "2-0" // 0球
	BetCodeBjdcZjq1 BetCode = "2-1" // 1球
	BetCodeBjdcZjq2 BetCode = "2-2" // 2球
	BetCodeBjdcZjq3 BetCode = "2-3" // 3球
	BetCodeBjdcZjq4 BetCode = "2-4" // 4球
	BetCodeBjdcZjq5 BetCode = "2-5" // 5球
	BetCodeBjdcZjq6 BetCode = "2-6" // 6球
	BetCodeBjdcZjq7 BetCode = "2-7" // 7球
)

// 投注码组：北京单场 - 上下单双
const (
	BetCodeBjdcSxdsSD BetCode = "3-0" // 上单
	BetCodeBjdcSxdsSS BetCode = "3-1" // 上双
	BetCodeBjdcSxdsXD BetCode = "3-2" // 下单
	BetCodeBjdcSxdsXS BetCode = "3-3" // 下双
)

// 投注码组：北京单场 - 单场比分
const (
	BetCodeBjdcBfS  BetCode = "4-0"  // 胜其他
	BetCodeBjdcBf10 BetCode = "4-1"  // 1:0
	BetCodeBjdcBf20 BetCode = "4-2"  // 2:0
	BetCodeBjdcBf21 BetCode = "4-3"  // 2:1
	BetCodeBjdcBf30 BetCode = "4-4"  // 3:0
	BetCodeBjdcBf31 BetCode = "4-5"  // 3:1
	BetCodeBjdcBf32 BetCode = "4-6"  // 3:2
	BetCodeBjdcBf40 BetCode = "4-7"  // 4:0
	BetCodeBjdcBf41 BetCode = "4-8"  // 4:1
	BetCodeBjdcBf42 BetCode = "4-9"  // 4:2
	BetCodeBjdcBfP  BetCode = "4-10" // 平其他
	BetCodeBjdcBf00 BetCode = "4-11" // 0:0
	BetCodeBjdcBf11 BetCode = "4-12" // 1:1
	BetCodeBjdcBf22 BetCode = "4-13" // 2:2
	BetCodeBjdcBf33 BetCode = "4-14" // 3:3
	BetCodeBjdcBfF  BetCode = "4-15" // 负其他
	BetCodeBjdcBf01 BetCode = "4-16" // 0:1
	BetCodeBjdcBf02 BetCode = "4-17" // 0:2
	BetCodeBjdcBf12 BetCode = "4-18" // 1:2
	BetCodeBjdcBf03 BetCode = "4-19" // 0:3
	BetCodeBjdcBf13 BetCode = "4-20" // 1:3
	BetCodeBjdcBf23 BetCode = "4-21" // 2:3
	BetCodeBjdcBf04 BetCode = "4-22" // 0:4
	BetCodeBjdcBf14 BetCode = "4-23" // 1:4
	BetCodeBjdcBf24 BetCode = "4-24" // 2:4
)

// 投注码组：北京单场 - 半全场
const (
	BetCodeBjdcBqcSS BetCode = "5-0" // 胜胜
	BetCodeBjdcBqcSP BetCode = "5-1" // 胜平
	BetCodeBjdcBqcSF BetCode = "5-2" // 胜负
	BetCodeBjdcBqcPS BetCode = "5-3" // 平胜
	BetCodeBjdcBqcPP BetCode = "5-4" // 平平
	BetCodeBjdcBqcPF BetCode = "5-5" // 平负
	BetCodeBjdcBqcFS BetCode = "5-6" // 负胜
	BetCodeBjdcBqcFP BetCode = "5-7" // 负平
	BetCodeBjdcBqcFF BetCode = "5-8" // 负负
)

// 投注码组：北京单场 - 胜负过关
const (
	BetCodeBjdcSfggS BetCode = "7-3" // 胜
	BetCodeBjdcSfggF BetCode = "7-0" // 负
)

// SportId 运动类型
type SportId = string

const (
	SportIdFootball      SportId = "1"
	SportIdBasketball    SportId = "2"
	SportIdBaseball      SportId = "3"
	SportIdHockey        SportId = "4"
	SportIdTennis        SportId = "5"
	SportIdHandball      SportId = "6"
	SportIdVolleyball    SportId = "7"
	SportIdRugby         SportId = "8"
	SportIdCricket       SportId = "9"
	SportIdSnooker       SportId = "10"
	SportIdBeachFootball SportId = "11"
	SportIdBadminton     SportId = "12"
	SportIdTablesTennis  SportId = "13"
	SportIdGolf          SportId = "14"
	SportIdHorseRacing   SportId = "15"
	SportIdWaterPolo     SportId = "16"
	SportId3X3Basketball SportId = "21"
	SportIdCSGO          SportId = "101"
	SportIdLOL           SportId = "102"
	SportIdHonorOfKings  SportId = "103"
	SportIdDOTA2         SportId = "104"
	SportIdHearthStone   SportId = "105"
	SportIdHeroStorm     SportId = "106"
	SportIdOverWatch     SportId = "107"
	SportIdStarCraft     SportId = "108"
	SportIdStarCraft2    SportId = "109"
	SportIdRocketLeague  SportId = "110"
	SportIdWarCraft3     SportId = "111"
)

// SportIdCnMapping 运动类型中文名映射表
var SportIdCnMapping = map[SportId]string{
	SportIdFootball:      "足球",
	SportIdBasketball:    "篮球",
	SportIdBaseball:      "棒球",
	SportIdHockey:        "冰上曲棍球",
	SportIdTennis:        "网球",
	SportIdHandball:      "手球",
	SportIdVolleyball:    "排球",
	SportIdRugby:         "美式橄榄球",
	SportIdCricket:       "板球",
	SportIdSnooker:       "斯诺克",
	SportIdBeachFootball: "沙滩足球",
	SportIdBadminton:     "羽毛球",
	SportIdTablesTennis:  "乒乓球",
	SportIdGolf:          "高尔夫",
	SportIdHorseRacing:   "赛马",
	SportIdWaterPolo:     "水球",
	SportId3X3Basketball: "3x3篮球",
	SportIdCSGO:          "反恐精英：全球攻势",
	SportIdLOL:           "英雄联盟",
	SportIdHonorOfKings:  "王者荣耀",
	SportIdDOTA2:         "刀塔2",
	SportIdHearthStone:   "炉石传说：魔兽英雄传",
	SportIdHeroStorm:     "风暴英雄",
	SportIdOverWatch:     "守望先锋",
	SportIdStarCraft:     "星际争霸",
	SportIdStarCraft2:    "星际争霸2",
	SportIdRocketLeague:  "火箭联盟",
	SportIdWarCraft3:     "魔兽争霸3",
}

// SportIdEnMapping 运动类型英文名映射表
var SportIdEnMapping = map[SportId]string{
	SportIdFootball:      "Football",
	SportIdBasketball:    "Basketball",
	SportIdBaseball:      "Baseball",
	SportIdHockey:        "Hockey",
	SportIdTennis:        "Tennis",
	SportIdHandball:      "Handball",
	SportIdVolleyball:    "Volleyball",
	SportIdRugby:         "Rugby",
	SportIdCricket:       "Cricket",
	SportIdSnooker:       "Snooker",
	SportIdBeachFootball: "Beach football",
	SportIdBadminton:     "Badminton",
	SportIdTablesTennis:  "Tables tennis",
	SportIdGolf:          "Golf",
	SportIdHorseRacing:   "Horse racing",
	SportIdWaterPolo:     "Waterpolo",
	SportId3X3Basketball: "3X3 Basketball",
	SportIdCSGO:          "CS:GO",
	SportIdLOL:           "LOL",
	SportIdHonorOfKings:  "Honor of Kings",
	SportIdDOTA2:         "DOTA2",
	SportIdHearthStone:   "HearthStone: Heroes Of Warcraft",
	SportIdHeroStorm:     "Hero Storm",
	SportIdOverWatch:     "Overwatch",
	SportIdStarCraft:     "Starcraft",
	SportIdStarCraft2:    "Starcraft 2",
	SportIdRocketLeague:  "Rocket League",
	SportIdWarCraft3:     "Warcraft 3",
}

// 状态码
type Status = string

const (
	StatueNotStarted Status = "1" // 未开始
	StatusInProgress Status = "2" // 进行中
	StatusFinished   Status = "3" // 已结束
	StatusCanceled   Status = "4" // 已取消
)

// StatusCode 子状态码
type StatusCode = string

const (
	StatusCode0   StatusCode = "0"
	StatusCode1   StatusCode = "1"
	StatusCode2   StatusCode = "2"
	StatusCode10  StatusCode = "10"
	StatusCode11  StatusCode = "11"
	StatusCode12  StatusCode = "12"
	StatusCode13  StatusCode = "13"
	StatusCode14  StatusCode = "14"
	StatusCode20  StatusCode = "20"
	StatusCode21  StatusCode = "21"
	StatusCode22  StatusCode = "22"
	StatusCode25  StatusCode = "25"
	StatusCode30  StatusCode = "30"
	StatusCode31  StatusCode = "31"
	StatusCode32  StatusCode = "32"
	StatusCode33  StatusCode = "33"
	StatusCode34  StatusCode = "34"
	StatusCode35  StatusCode = "35"
	StatusCode36  StatusCode = "36"
	StatusCode40  StatusCode = "40"
	StatusCode41  StatusCode = "41"
	StatusCode42  StatusCode = "42"
	StatusCode43  StatusCode = "43"
	StatusCode44  StatusCode = "44"
	StatusCode45  StatusCode = "45"
	StatusCode100 StatusCode = "100"
	StatusCode301 StatusCode = "301"
	StatusCode302 StatusCode = "302"
	StatusCode303 StatusCode = "303"
	StatusCode304 StatusCode = "304"
	StatusCode305 StatusCode = "305"
	StatusCode306 StatusCode = "306"
	StatusCode307 StatusCode = "307"
	StatusCode308 StatusCode = "308"
	StatusCode309 StatusCode = "309"
	StatusCode310 StatusCode = "310"
	StatusCode311 StatusCode = "311"
	StatusCode312 StatusCode = "312"
	StatusCode313 StatusCode = "313"
	StatusCode314 StatusCode = "314"
	StatusCode315 StatusCode = "315"
	StatusCode316 StatusCode = "316"
	StatusCode317 StatusCode = "317"
	StatusCode318 StatusCode = "318"
	StatusCode319 StatusCode = "319"
	StatusCode320 StatusCode = "320"
	StatusCode501 StatusCode = "501"
	StatusCode502 StatusCode = "502"
	StatusCode503 StatusCode = "503"
	StatusCode504 StatusCode = "504"
	StatusCode505 StatusCode = "505"
	StatusCode506 StatusCode = "506"
	StatusCode507 StatusCode = "507"
	StatusCode901 StatusCode = "901"
	StatusCode902 StatusCode = "902"
	StatusCode903 StatusCode = "903"
	StatusCode904 StatusCode = "904"
	StatusCode905 StatusCode = "905"
	StatusCode944 StatusCode = "944"
	StatusCode945 StatusCode = "945"
	StatusCode951 StatusCode = "951"
	StatusCode952 StatusCode = "952"
	StatusCode953 StatusCode = "953"
	StatusCode954 StatusCode = "954"
)

// StatusCodeCnNameMapping 子状态中文名称映射表
var StatusCodeCnNameMapping = map[StatusCode]string{
	StatusCode0:   "未开始",
	StatusCode1:   "上半场",
	StatusCode2:   "下半场",
	StatusCode10:  "进行中",
	StatusCode11:  "第一节",
	StatusCode12:  "第二节",
	StatusCode13:  "第三节",
	StatusCode14:  "第四节",
	StatusCode20:  "加时",
	StatusCode21:  "加时上半场",
	StatusCode22:  "加时下半场",
	StatusCode25:  "点球",
	StatusCode30:  "暂停",
	StatusCode31:  "中场",
	StatusCode32:  "等待加时赛",
	StatusCode33:  "加时赛中场",
	StatusCode34:  "加时赛后",
	StatusCode35:  "等待点球决胜",
	StatusCode36:  "点球决胜后",
	StatusCode40:  "取消",
	StatusCode41:  "延期",
	StatusCode42:  "推迟",
	StatusCode43:  "中断",
	StatusCode44:  "弃赛",
	StatusCode45:  "待定",
	StatusCode100: "完场",
	StatusCode301: "第一盘",
	StatusCode302: "第二盘",
	StatusCode303: "第三盘",
	StatusCode304: "第四盘",
	StatusCode305: "第五盘",
	StatusCode306: "第六盘",
	StatusCode307: "第七盘",
	StatusCode308: "第八盘",
	StatusCode309: "第九盘",
	StatusCode310: "第十盘",
	StatusCode311: "第十一盘",
	StatusCode312: "第十二盘",
	StatusCode313: "第十三盘",
	StatusCode314: "第十四盘",
	StatusCode315: "第十五盘",
	StatusCode316: "第十六盘",
	StatusCode317: "第十七盘",
	StatusCode318: "第十八盘",
	StatusCode319: "第十九盘",
	StatusCode320: "第二十盘",
	StatusCode501: "第一盘",
	StatusCode502: "第二盘",
	StatusCode503: "第三盘",
	StatusCode504: "第四盘",
	StatusCode505: "第五盘",
	StatusCode506: "第六盘",
	StatusCode507: "第七盘",
	StatusCode901: "第一局",
	StatusCode902: "第二局",
	StatusCode903: "第三局",
	StatusCode904: "第四局",
	StatusCode905: "第五局",
	StatusCode944: "茶歇",
	StatusCode945: "午餐",
	StatusCode951: "第一赛日结束",
	StatusCode952: "第二赛日结束",
	StatusCode953: "第三赛日结束",
	StatusCode954: "第四赛日结束",
}

// ScoreType 比分类型
type ScoreType = string

const (
	ScoreCurrent         ScoreType = "Current"
	ScoreNormalTime      ScoreType = "Normaltime"
	ScoreOvertime        ScoreType = "Overtime"
	ScorePenalties       ScoreType = "Penalties"
	ScoreOt              ScoreType = "ot"
	ScorePt              ScoreType = "pt"
	ScorePeriod1         ScoreType = "Period1"
	ScorePeriod2         ScoreType = "Period2"
	ScorePeriod3         ScoreType = "Period3"
	ScorePeriod4         ScoreType = "Period4"
	ScorePeriod5         ScoreType = "Period5"
	ScorePeriod6         ScoreType = "Period6"
	ScorePeriod7         ScoreType = "Period7"
	ScorePeriod8         ScoreType = "Period8"
	ScorePeriod9         ScoreType = "Period9"
	ScorePeriod10        ScoreType = "Period10"
	ScorePeriod1TieBreak ScoreType = "Period1TieBreak"
	ScorePeriod2TieBreak ScoreType = "Period2TieBreak"
	ScorePeriod3TieBreak ScoreType = "Period3TieBreak"
	ScorePeriod4TieBreak ScoreType = "Period4TieBreak"
	ScorePeriod5TieBreak ScoreType = "Period5TieBreak"
)

// SportScoreTypeCnNameMapping 运动比分类型中文名映射表
var SportScoreTypeCnNameMapping = map[SportId]map[ScoreType]string{
	SportIdFootball: {
		ScoreCurrent:    "当前总比分，若完场则为总比分",
		ScoreNormalTime: "常规时间比分(足球通常90分钟)",
		ScoreOvertime:   "含加时比分，即常规时间+加时的总比分",
		ScorePenalties:  "含点球比分，即常规时间+加时+点球的总比分",
		ScorePeriod1:    "上半场比分",
		ScoreOt:         "仅加时段比分",
		ScorePt:         "仅点球大战比分",
	},
	SportIdBasketball: {
		ScoreCurrent:  "当前总比分，若完场则为总比分",
		ScorePeriod1:  "第一节比分",
		ScorePeriod2:  "第二节比分",
		ScorePeriod3:  "第三节比分",
		ScorePeriod4:  "第四节比分",
		ScoreOvertime: "含加时比分，即常规时间+加时的总比分",
	},
	SportIdTennis: {
		ScoreCurrent:         "盘比分",
		ScorePeriod1:         "第一盘局比分",
		ScorePeriod2:         "第二盘局比分",
		ScorePeriod3:         "第三盘局比分",
		ScorePeriod4:         "第四盘局比分",
		ScorePeriod5:         "第五盘局比分",
		ScorePeriod1TieBreak: "第一局抢七/抢十比分",
		ScorePeriod2TieBreak: "第二局抢七/抢十比分",
		ScorePeriod3TieBreak: "第三局抢七/抢十比分",
		ScorePeriod4TieBreak: "第四局抢七/抢十比分",
		ScorePeriod5TieBreak: "第五局抢七/抢十比分",
	},
	SportIdBaseball: {
		ScoreCurrent:    "盘比分（含加时）",
		ScoreNormalTime: "盘比分（不含加时）",
		ScorePeriod1:    "第一盘比分",
		ScorePeriod2:    "第二盘比分",
		ScorePeriod3:    "第三盘比分",
		ScorePeriod4:    "第四盘比分",
		ScorePeriod5:    "第五盘比分",
		ScorePeriod6:    "第六盘比分",
		ScorePeriod7:    "第七盘比分",
		ScorePeriod8:    "第八盘比分",
		ScorePeriod9:    "第九盘比分",
		ScorePeriod10:   "加时盘的总比分",
	},
}

// RankType 榜单类型
type RankType = string

const (
	RankTypePositionTotal       = "1"
	RankTypePositionHome        = "2"
	RankTypePositionAway        = "3"
	RankTypeSortPositionTotal   = "4"
	RankTypeSortPositionHome    = "5"
	RankTypeSortPositionAway    = "6"
	RankTypeChangeTotal         = "7"
	RankTypeChangeHome          = "8"
	RankTypeChangeAway          = "9"
	RankTypeMatchesTotal        = "10"
	RankTypeMatchesHome         = "11"
	RankTypeMatchesAway         = "12"
	RankTypePointsTotal         = "13"
	RankTypePointsHome          = "14"
	RankTypePointsAway          = "15"
	RankTypeWinTotal            = "16"
	RankTypeWinHome             = "17"
	RankTypeWinAway             = "18"
	RankTypeDrawTotal           = "19"
	RankTypeDrawHome            = "20"
	RankTypeDrawAway            = "21"
	RankTypeLossTotal           = "22"
	RankTypeLossHome            = "23"
	RankTypeLossAway            = "24"
	RankTypeGoalsTotal          = "25"
	RankTypeGoalsTotalHome      = "26"
	RankTypeGoalsTotalAway      = "27"
	RankTypeGoalsForTotal       = "28"
	RankTypeGoalsForHome        = "29"
	RankTypeGoalsForAway        = "30"
	RankTypeGoalsAgainstTotal   = "31"
	RankTypeGoalsAgainstHome    = "32"
	RankTypeGoalsAgainstAway    = "33"
	RankTypeGoalDiffTotal       = "34"
	RankTypeGoalDiffHome        = "35"
	RankTypeGoalDiffAway        = "36"
	RankTypeStreak              = "37"
	RankTypeWinLast10           = "38"
	RankTypeLossLast10          = "39"
	RankTypeWinPctTotal         = "40"
	RankTypeWinPctHome          = "41"
	RankTypeWinPctAway          = "42"
	RankTypeGameBehind          = "43"
	RankTypeGoalsForPgTotal     = "44"
	RankTypeGoalsForPgHome      = "45"
	RankTypeGoalsForPgAway      = "46"
	RankTypeGoalsAgainstPgTotal = "47"
	RankTypeGoalsAgainstPgHome  = "48"
	RankTypeGoalsAgainstPgAway  = "49"
	RankTypeNetRunRate          = "50"
	RankTypeScoresFor           = "51"
	RankTypeScoresAgainst       = "52"
	RankTypeNoResult            = "53"
	RankTypeOversFor            = "54"
	RankTypeOversAgainst        = "55"
	RankTypePercentage          = "56"
)

// RankTypeCnNameMapping 榜单类型中文名映射表
var RankTypeCnNameMapping = map[RankType]string{
	RankTypePositionTotal:       "排名",
	RankTypePositionHome:        "主场排名",
	RankTypePositionAway:        "客场排名",
	RankTypeSortPositionTotal:   "排序",
	RankTypeSortPositionHome:    "主场排序",
	RankTypeSortPositionAway:    "客场排序",
	RankTypeChangeTotal:         "排名变更",
	RankTypeChangeHome:          "主场排名变更",
	RankTypeChangeAway:          "客场排名变更",
	RankTypeMatchesTotal:        "比赛场数",
	RankTypeMatchesHome:         "主场比赛场数",
	RankTypeMatchesAway:         "客场比赛场数",
	RankTypePointsTotal:         "积分",
	RankTypePointsHome:          "主场积分",
	RankTypePointsAway:          "客场积分",
	RankTypeWinTotal:            "胜场",
	RankTypeWinHome:             "主场胜场",
	RankTypeWinAway:             "客场胜场",
	RankTypeDrawTotal:           "平局",
	RankTypeDrawHome:            "主场平局",
	RankTypeDrawAway:            "客场平局",
	RankTypeLossTotal:           "负场",
	RankTypeLossHome:            "主场负场",
	RankTypeLossAway:            "客场负场",
	RankTypeGoalsTotal:          "进失球",
	RankTypeGoalsTotalHome:      "主场进失球",
	RankTypeGoalsTotalAway:      "客场进失球",
	RankTypeGoalsForTotal:       "进球",
	RankTypeGoalsForHome:        "主场进球",
	RankTypeGoalsForAway:        "客场进球",
	RankTypeGoalsAgainstTotal:   "失球",
	RankTypeGoalsAgainstHome:    "主场失球",
	RankTypeGoalsAgainstAway:    "客场失球",
	RankTypeGoalDiffTotal:       "净胜球",
	RankTypeGoalDiffHome:        "主场净胜球",
	RankTypeGoalDiffAway:        "客场净胜球",
	RankTypeStreak:              "连续战绩，+连胜，-连败",
	RankTypeWinLast10:           "近10场胜",
	RankTypeLossLast10:          "近10场负",
	RankTypeWinPctTotal:         "胜率",
	RankTypeWinPctHome:          "主场胜率",
	RankTypeWinPctAway:          "客场胜率",
	RankTypeGameBehind:          "胜差",
	RankTypeGoalsForPgTotal:     "场均进球(得分)",
	RankTypeGoalsForPgHome:      "主场场均进球(得分)",
	RankTypeGoalsForPgAway:      "客场场均进球(得分)",
	RankTypeGoalsAgainstPgTotal: "场均失球(失分)",
	RankTypeGoalsAgainstPgHome:  "主场场均失球(失分)",
	RankTypeGoalsAgainstPgAway:  "客场场均失球(失分)",
	RankTypeNetRunRate:          "净得分率",
	RankTypeScoresFor:           "击球得分",
	RankTypeScoresAgainst:       "投球得分",
	RankTypeNoResult:            "无结果",
	RankTypeOversFor:            "击球轮数",
	RankTypeOversAgainst:        "投球轮数",
	RankTypePercentage:          "胜率",
}

// RankTypeEnNameMapping 榜单类型英文名映射表
var RankTypeEnNameMapping = map[RankType]string{
	RankTypePositionTotal:       "positionTotal",
	RankTypePositionHome:        "positionHome",
	RankTypePositionAway:        "positionAway",
	RankTypeSortPositionTotal:   "sortPositionTotal",
	RankTypeSortPositionHome:    "sortPositionHome",
	RankTypeSortPositionAway:    "sortPositionAway",
	RankTypeChangeTotal:         "changeTotal",
	RankTypeChangeHome:          "changeHome",
	RankTypeChangeAway:          "changeAway",
	RankTypeMatchesTotal:        "matchesTotal",
	RankTypeMatchesHome:         "matchesHome",
	RankTypeMatchesAway:         "matchesAway",
	RankTypePointsTotal:         "pointsTotal",
	RankTypePointsHome:          "pointsHome",
	RankTypePointsAway:          "pointsAway",
	RankTypeWinTotal:            "winTotal",
	RankTypeWinHome:             "winHome",
	RankTypeWinAway:             "winAway",
	RankTypeDrawTotal:           "drawTotal",
	RankTypeDrawHome:            "drawHome",
	RankTypeDrawAway:            "drawAway",
	RankTypeLossTotal:           "lossTotal",
	RankTypeLossHome:            "lossHome",
	RankTypeLossAway:            "lossAway",
	RankTypeGoalsTotal:          "goalsTotal",
	RankTypeGoalsTotalHome:      "goalsTotalHome",
	RankTypeGoalsTotalAway:      "goalsTotalAway",
	RankTypeGoalsForTotal:       "goalsForTotal",
	RankTypeGoalsForHome:        "goalsForHome",
	RankTypeGoalsForAway:        "goalsForAway",
	RankTypeGoalsAgainstTotal:   "goalsAgainstTotal",
	RankTypeGoalsAgainstHome:    "goalsAgainstHome",
	RankTypeGoalsAgainstAway:    "goalsAgainstAway",
	RankTypeGoalDiffTotal:       "goalDiffTotal",
	RankTypeGoalDiffHome:        "goalDiffHome",
	RankTypeGoalDiffAway:        "goalDiffAway",
	RankTypeStreak:              "streak",
	RankTypeWinLast10:           "winLast10",
	RankTypeLossLast10:          "lossLast10",
	RankTypeWinPctTotal:         "winPctTotal",
	RankTypeWinPctHome:          "winPctHome",
	RankTypeWinPctAway:          "winPctAway",
	RankTypeGameBehind:          "gameBehind",
	RankTypeGoalsForPgTotal:     "goalsForPgTotal",
	RankTypeGoalsForPgHome:      "goalsForPgHome",
	RankTypeGoalsForPgAway:      "goalsForPgAway",
	RankTypeGoalsAgainstPgTotal: "goalsAgainstPgTotal",
	RankTypeGoalsAgainstPgHome:  "goalsAgainstPgHome",
	RankTypeGoalsAgainstPgAway:  "goalsAgainstPgAway",
	RankTypeNetRunRate:          "netRunRate",
	RankTypeScoresFor:           "scoresFor",
	RankTypeScoresAgainst:       "scoresAgainst",
	RankTypeNoResult:            "noResult",
	RankTypeOversFor:            "oversFor",
	RankTypeOversAgainst:        "oversAgainst",
	RankTypePercentage:          "percentage",
}

// OddsType 赔率类型
type OddsType = int64

const (
	OddsType1   OddsType = 1
	OddsType2   OddsType = 2
	OddsType3   OddsType = 3
	OddsType4   OddsType = 4
	OddsType5   OddsType = 5
	OddsType6   OddsType = 6
	OddsType7   OddsType = 7
	OddsType8   OddsType = 8
	OddsType9   OddsType = 9
	OddsType10  OddsType = 10
	OddsType11  OddsType = 11
	OddsType12  OddsType = 12
	OddsType13  OddsType = 13
	OddsType14  OddsType = 14
	OddsType15  OddsType = 15
	OddsType16  OddsType = 16
	OddsType24  OddsType = 24
	OddsType25  OddsType = 25
	OddsType26  OddsType = 26
	OddsType27  OddsType = 27
	OddsType28  OddsType = 28
	OddsType29  OddsType = 29
	OddsType121 OddsType = 121
	OddsType122 OddsType = 122
	OddsType123 OddsType = 123
	OddsType124 OddsType = 124
	OddsType125 OddsType = 125
	OddsType126 OddsType = 126
	OddsType127 OddsType = 127
	OddsType128 OddsType = 128
	OddsType129 OddsType = 129
	OddsType130 OddsType = 130
	OddsType131 OddsType = 131
	OddsType132 OddsType = 132
	OddsType133 OddsType = 133
	OddsType134 OddsType = 134
	OddsType135 OddsType = 135
	OddsType136 OddsType = 136
	OddsType137 OddsType = 137
	OddsType138 OddsType = 138
	OddsType139 OddsType = 139
	OddsType140 OddsType = 140
	OddsType141 OddsType = 141
	OddsType142 OddsType = 142
	OddsType143 OddsType = 143
	OddsType144 OddsType = 144
	OddsType145 OddsType = 145
	OddsType146 OddsType = 146
	OddsType147 OddsType = 147
	OddsType148 OddsType = 148
	OddsType149 OddsType = 149
	OddsType150 OddsType = 150
	OddsType151 OddsType = 151
	OddsType159 OddsType = 159
	OddsType160 OddsType = 160
	OddsType161 OddsType = 161
)

// OddsTypeCnNameMapping 赔率类型玩法名
var OddsTypeCnNameMapping = map[OddsType]string{
	OddsType1:   "让球胜负(足球常规时间)、让分胜负(篮球含加时)、让局(网球)、让分(排球)、让垒(棒球)",
	OddsType2:   "胜平负(足球常规时间)、胜负(篮球含加时、网球、棒球、排球)",
	OddsType3:   "大小球(足球常规时间)、大小分(篮球含加时)、大小局(网球)、大小分(排球)、大小垒(棒球)",
	OddsType4:   "总进球数",
	OddsType5:   "半场/全场",
	OddsType6:   "比分",
	OddsType7:   "总得分奇偶",
	OddsType8:   "净胜球",
	OddsType9:   "角球大小球",
	OddsType10:  "首球时间",
	OddsType11:  "是否加时",
	OddsType12:  "上半场胜平负",
	OddsType13:  "上半场让球胜负",
	OddsType14:  "上半场大小球",
	OddsType15:  "双胜彩",
	OddsType16:  "让球胜平负",
	OddsType24:  "上下盘单双",
	OddsType25:  "胜分差(篮球)",
	OddsType26:  "上半场角球大小球",
	OddsType27:  "下半场胜平负",
	OddsType28:  "下半场让球胜负",
	OddsType29:  "下半场大小球",
	OddsType121: "让球胜负(足球)、让分胜负(篮球)、让局(网球)、让分(排球)、让垒(棒球)",
	OddsType122: "大小球(足球)、大小分(篮球)、大小局(网球)、大小分(排球)、大小垒(棒球)",
	OddsType123: "比分",
	OddsType124: "总进球数",
	OddsType125: "下一个进球",
	OddsType126: "总得分奇偶",
	OddsType127: "先到x得分",
	OddsType128: "胜平负(足球)、胜负(篮球、网球、棒球、排球)",
	OddsType129: "是否有红牌",
	OddsType130: "总角球数大小",
	OddsType131: "哪队先开球",
	OddsType132: "上半场让球胜负",
	OddsType133: "上半场胜平负",
	OddsType134: "上半场大小球",
	OddsType135: "上半场角球大小球",
	OddsType136: "争球谁能赢",
	OddsType137: "第一节让分胜负",
	OddsType138: "第二节让分胜负",
	OddsType139: "第三节让分胜负",
	OddsType140: "第四节让分胜负",
	OddsType141: "第一节大小球",
	OddsType142: "第二节大小球",
	OddsType143: "第三节大小球",
	OddsType144: "第四节大小球",
	OddsType145: "主队进球数",
	OddsType146: "客队进球数",
	OddsType147: "是否有进球",
	OddsType148: "主队得分大小球",
	OddsType149: "客队得分大小球",
	OddsType150: "半场/全场",
	OddsType151: "让球胜平负",
	OddsType159: "下半场让球胜负",
	OddsType160: "下半场胜平负",
	OddsType161: "下半场大小球",
}

// EventType 事件类型
type EventType = string

const (
	EventType8   EventType = "8"
	EventType9   EventType = "9"
	EventType18  EventType = "18"
	EventType21  EventType = "21"
	EventType22  EventType = "22"
	EventType23  EventType = "23"
	EventType24  EventType = "24"
	EventType26  EventType = "26"
	EventType27  EventType = "27"
	EventType28  EventType = "28"
	EventType29  EventType = "29"
	EventType30  EventType = "30"
	EventType31  EventType = "31"
	EventType32  EventType = "32"
	EventType33  EventType = "33"
	EventType34  EventType = "34"
	EventType35  EventType = "35"
	EventType40  EventType = "40"
	EventType41  EventType = "41"
	EventType156 EventType = "156"
	EventType157 EventType = "157"
	EventType158 EventType = "158"
	EventType159 EventType = "159"
	EventType160 EventType = "160"
	EventType161 EventType = "161"
	EventType162 EventType = "162"
	EventType1   EventType = "1"
	EventType3   EventType = "3"
	EventType50  EventType = "50"
	EventType83  EventType = "83"
	EventType85  EventType = "85"
	EventType86  EventType = "86"
	EventType94  EventType = "94"
	EventType117 EventType = "117"
	EventType124 EventType = "124"
	EventType127 EventType = "127"
	EventType128 EventType = "128"
	EventType129 EventType = "129"
	EventType134 EventType = "134"
	EventType139 EventType = "139"
	EventType140 EventType = "140"
	EventType141 EventType = "141"
	EventType142 EventType = "142"
)

// SportEventTypeCnNameMapping 时间类型中文名映射表
var SportEventTypeCnNameMapping = map[SportId]map[EventType]string{
	SportIdFootball: {
		EventType1:   "比赛开始",
		EventType3:   "比赛结束",
		EventType8:   "点球",
		EventType9:   "进球",
		EventType18:  "黄牌",
		EventType21:  "黄红牌",
		EventType22:  "红牌",
		EventType23:  "替补",
		EventType24:  "伤停补时显示",
		EventType26:  "任意球",
		EventType27:  "球门球",
		EventType28:  "界外球",
		EventType29:  "越位",
		EventType30:  "角球",
		EventType31:  "射正不进",
		EventType32:  "射偏",
		EventType33:  "门将安全持球",
		EventType34:  "受伤",
		EventType35:  "获得点球",
		EventType40:  "点球射失",
		EventType41:  "点球决胜开始",
		EventType128: "进攻",
		EventType127: "控球",
		EventType129: "危险进攻",
		EventType156: "进球确认(VAR视频裁判事件，0取消进球 1确认进球)",
		EventType157: "罚牌升级确认(VAR视频裁判事件，0不升级 1升级)",
		EventType158: "未进球确认(VAR视频裁判事件，0进球 1未进球)",
		EventType159: "点球确认(VAR视频裁判事件，0点球取消 1确认点球)",
		EventType160: "红牌确认(VAR视频裁判事件，0红牌取消 1确认红牌)",
		EventType161: "不判罚点球确认(VAR视频裁判事件，0判罚点球 1不判罚点球)",
		EventType162: "犯规球员确认(VAR视频裁判事件，0没有犯规 1判罚犯规)",
	},
	SportIdBasketball: {
		EventType1:   "比赛开始",
		EventType3:   "比赛结束",
		EventType50:  "暂停",
		EventType83:  "罚球",
		EventType85:  "罚球得分",
		EventType86:  "罚球未中",
		EventType94:  "犯规",
		EventType117: "未中",
		EventType124: "得分",
		EventType127: "控球",
		EventType128: "进攻",
		EventType134: "赛节结束",
		EventType139: "加时赛",
		EventType140: "第二节开始",
		EventType141: "第三节开始",
		EventType142: "第四节开始",
	},
}

// Weather 天气
type Weather = string

const (
	WeatherGloomy                  Weather = "Gloomy"
	WeatherClearSkies              Weather = "Clear skies"
	WeatherCloudy                  Weather = "Cloudy"
	WeatherSomeCloud               Weather = "Some cloud"
	WeatherSnow                    Weather = "Snow"
	WeatherCloudyRain              Weather = "Cloudy, rain"
	WeatherGloomyRainThunderstorms Weather = "Gloomy, rain and thunderstorms"
	WeatherFog                     Weather = "Fog"
	WeatherMist                    Weather = "mist"
	WeatherGloomyRain              Weather = "Gloomy, rain"
	WeatherSomeCloudRain           Weather = "Some cloud, rain"
)

// WeatherCnNameMapping 天气翻译映射表
var WeatherCnNameMapping = map[Weather]string{
	WeatherGloomy:                  "阴天",
	WeatherClearSkies:              "晴天",
	WeatherCloudy:                  "多云",
	WeatherSomeCloud:               "多云",
	WeatherSnow:                    "雪",
	WeatherCloudyRain:              "阵雨",
	WeatherGloomyRainThunderstorms: "阴转雷暴",
	WeatherFog:                     "有雾",
	WeatherMist:                    "薄雾",
	WeatherGloomyRain:              "阴转雨",
	WeatherSomeCloudRain:           "多云转雨",
}

// WeatherKey 天气K值
type WeatherKey = int64

const (
	WeatherKey9  WeatherKey = 9
	WeatherKey21 WeatherKey = 21
	WeatherKey23 WeatherKey = 23
	WeatherKey24 WeatherKey = 24
	WeatherKey25 WeatherKey = 25
	WeatherKey26 WeatherKey = 26
	WeatherKey27 WeatherKey = 27
	WeatherKey28 WeatherKey = 28
)

// WeatherKeyDescMapping 天气K值描述映射表
var WeatherKeyDescMapping = map[WeatherKey]string{
	WeatherKey9:  "温度",
	WeatherKey21: "天气概述，同weather字段",
	WeatherKey23: "风速值",
	WeatherKey24: "风向、风速单位",
	WeatherKey25: "气压值",
	WeatherKey26: "气压单位",
	WeatherKey27: "湿度值",
	WeatherKey28: "湿度单位",
}

// WeatherWind 风向
type WeatherWind = string

const (
	WeatherWindN   WeatherWind = "N"
	WeatherWindNNE WeatherWind = "NNE"
	WeatherWindNE  WeatherWind = "NE"
	WeatherWindENE WeatherWind = "ENE"
	WeatherWindE   WeatherWind = "E"
	WeatherWindESE WeatherWind = "ESE"
	WeatherWindSE  WeatherWind = "SE"
	WeatherWindSSE WeatherWind = "SSE"
	WeatherWindS   WeatherWind = "S"
	WeatherWindSSW WeatherWind = "SSW"
	WeatherWindSW  WeatherWind = "SW"
	WeatherWindWSW WeatherWind = "WSW"
	WeatherWindW   WeatherWind = "W"
	WeatherWindWNW WeatherWind = "WNW"
	WeatherWindNW  WeatherWind = "NW"
	WeatherWindNNW WeatherWind = "NNW"
)

// WeatherWindCnNameMapping 风向中文名映射表
var WeatherWindCnNameMapping = map[WeatherWind]string{
	WeatherWindN:   "北风",
	WeatherWindNNE: "东北偏北风",
	WeatherWindNE:  "东北风",
	WeatherWindENE: "东北偏东风",
	WeatherWindE:   "东风",
	WeatherWindESE: "东南偏东风",
	WeatherWindSE:  "东南风",
	WeatherWindSSE: "东南偏南风",
	WeatherWindS:   "南风",
	WeatherWindSSW: "西南偏南风",
	WeatherWindSW:  "西南风",
	WeatherWindWSW: "西南偏西风",
	WeatherWindW:   "西风",
	WeatherWindWNW: "西北偏西风",
	WeatherWindNW:  "西北风",
	WeatherWindNNW: "西北偏北风",
}

// LotteryId 彩种ID
type LotteryId = int64

const (
	LotteryIdCtzcSfc   LotteryId = 1
	LotteryIdCtzcBqc   LotteryId = 2
	LotteryIdBjdcSfgg  LotteryId = 3
	LotteryIdBjdcBqc   LotteryId = 4
	LotteryIdBjdcZjq   LotteryId = 5
	LotteryIdBjdcBf    LotteryId = 6
	LotteryIdBjdcRqspf LotteryId = 7
	LotteryIdCtzcJqc   LotteryId = 8
	LotteryIdJclq      LotteryId = 9
	LotteryIdJczq      LotteryId = 10
	LotteryIdJczqSpf   LotteryId = 11
	LotteryIdJczqRqspf LotteryId = 12
	LotteryIdJczqBf    LotteryId = 13
	LotteryIdJczqZjq   LotteryId = 14
	LotteryIdJczqBqc   LotteryId = 15
	LotteryIdJclqSf    LotteryId = 16
	LotteryIdJclqRfsf  LotteryId = 17
	LotteryIdJclqDxf   LotteryId = 18
	LotteryIdBjdcSxds  LotteryId = 19
	LotteryIdJclqSfc   LotteryId = 21
	LotteryIdJclqWait  LotteryId = 22
	LotteryIdJczqWait  LotteryId = 23
)

// LotteryIdNameMapping 彩种名称映射表
var LotteryIdNameMapping = map[LotteryId]string{
	LotteryIdCtzcSfc:   "14场胜负彩",
	LotteryIdCtzcBqc:   "六场半全场",
	LotteryIdBjdcSfgg:  "北京单场胜负过关",
	LotteryIdBjdcBqc:   "单场半全场胜平负",
	LotteryIdBjdcZjq:   "单场总进球数",
	LotteryIdBjdcBf:    "单场比分",
	LotteryIdBjdcRqspf: "单场让球胜平负",
	LotteryIdCtzcJqc:   "四场进球彩",
	LotteryIdJclq:      "竞彩篮球",
	LotteryIdJczq:      "竞彩足球",
	LotteryIdJczqSpf:   "竞彩胜平负",
	LotteryIdJczqRqspf: "竞彩让球胜平负",
	LotteryIdJczqBf:    "竞彩单场比分",
	LotteryIdJczqZjq:   "竞彩总进球",
	LotteryIdJczqBqc:   "竞彩半全场",
	LotteryIdJclqSf:    "竞彩胜负",
	LotteryIdJclqRfsf:  "竞彩让分胜负",
	LotteryIdJclqDxf:   "竞彩大小分",
	LotteryIdBjdcSxds:  "上下盘单双数",
	LotteryIdJclqSfc:   "竞彩胜分差",
	LotteryIdJclqWait:  "竞彩篮球-暂定赛程",
	LotteryIdJczqWait:  "竞彩足球-暂定赛程",
}

// JczqLotteryId 竞彩足球彩种ID
type JczqLotteryId = int64

const (
	JczqLotteryIdSpf   JczqLotteryId = 2
	JczqLotteryIdZjq   JczqLotteryId = 4
	JczqLotteryIdBqc   JczqLotteryId = 5
	JczqLotteryIdBf    JczqLotteryId = 6
	JczqLotteryIdRqspf JczqLotteryId = 16
)

// JczqLotteryIdNameMapping 竞彩足球彩种名称映射表
var JczqLotteryIdNameMapping = map[JczqLotteryId]string{
	JczqLotteryIdSpf:   "竞彩胜平负",
	JczqLotteryIdZjq:   "竞彩总进球",
	JczqLotteryIdBqc:   "竞彩半全场",
	JczqLotteryIdBf:    "竞彩单场比分",
	JczqLotteryIdRqspf: "竞彩让球胜平负",
}

// JclqLotteryId 竞彩篮球彩种ID
type JclqLotteryId = int64

const (
	JclqLotteryIdRfsf JclqLotteryId = 1
	JclqLotteryIdSf   JclqLotteryId = 2
	JclqLotteryIdDxf  JclqLotteryId = 3
	JclqLotteryIdSfc  JclqLotteryId = 25
)

// JclqLotteryIdNameMapping 竞彩篮球彩种名称映射表
var JclqLotteryIdNameMapping = map[JclqLotteryId]string{
	JclqLotteryIdRfsf: "竞彩让分胜负",
	JclqLotteryIdSf:   "竞彩胜负",
	JclqLotteryIdDxf:  "竞彩大小分",
	JclqLotteryIdSfc:  "胜分差",
}

// JingcaiType 竞彩玩法选项
type JingcaiType = string

const (
	JingcaiTypeS JingcaiType = "1"
	JingcaiTypeP JingcaiType = "x"
	JingcaiTypeF JingcaiType = "2"

	JingcaiType0 JingcaiType = "0"
	JingcaiType1 JingcaiType = "1"
	JingcaiType2 JingcaiType = "2"
	JingcaiType3 JingcaiType = "3"
	JingcaiType4 JingcaiType = "4"
	JingcaiType5 JingcaiType = "5"
	JingcaiType6 JingcaiType = "6"
	JingcaiType7 JingcaiType = "7+"

	JingcaiTypeX2 JingcaiType = "x/2"
	JingcaiType11 JingcaiType = "1/1"
	JingcaiTypeXX JingcaiType = "x/x"
	JingcaiType12 JingcaiType = "1/2"
	JingcaiType1X JingcaiType = "1/x"
	JingcaiType21 JingcaiType = "2/1"
	JingcaiType22 JingcaiType = "2/2"
	JingcaiType2X JingcaiType = "2/x"
	JingcaiTypeX1 JingcaiType = "x/1"

	JingcaiTypeOver  JingcaiType = "over"
	JingcaiTypeUnder JingcaiType = "under"

	JingcaiType1a JingcaiType = "1a"
	JingcaiType1d JingcaiType = "1d"
	JingcaiType1h JingcaiType = "1h"

	JingcaiTypeH1 JingcaiType = "h1"
	JingcaiTypeH2 JingcaiType = "h2"
	JingcaiTypeH3 JingcaiType = "h3"
	JingcaiTypeH4 JingcaiType = "h4"
	JingcaiTypeH5 JingcaiType = "h5"
	JingcaiTypeH6 JingcaiType = "h6"
	JingcaiTypeA1 JingcaiType = "a1"
	JingcaiTypeA2 JingcaiType = "a2"
	JingcaiTypeA3 JingcaiType = "a3"
	JingcaiTypeA4 JingcaiType = "a4"
	JingcaiTypeA5 JingcaiType = "a5"
	JingcaiTypeA6 JingcaiType = "a6"

	JingcaiTypeOdd1  JingcaiType = "odd1"
	JingcaiTypeEven1 JingcaiType = "even1"
	JingcaiTypeOdd2  JingcaiType = "odd2"
	JingcaiTypeEven2 JingcaiType = "even2"
)

// JingcaiTypeSpfNameMapping 竞彩胜平负\竞彩让球胜平负\竞彩胜负\竞彩让分胜负名称映射表
var JingcaiTypeSpfNameMapping = map[JingcaiType]string{
	JingcaiTypeS: "主胜",
	JingcaiTypeP: "平局",
	JingcaiTypeF: "客胜",
}

// JingcaiTypeZjqNameMapping 竞彩总进球名称映射表
var JingcaiTypeZjqNameMapping = map[JingcaiType]string{
	JingcaiType0: "无",
	JingcaiType1: "1球",
	JingcaiType2: "2球",
	JingcaiType3: "3球",
	JingcaiType4: "4球",
	JingcaiType5: "5球",
	JingcaiType6: "6球",
	JingcaiType7: "7球及以上",
}

// JingcaiTypeBqcNameMapping 竞彩半全场称映射表
var JingcaiTypeBqcNameMapping = map[JingcaiType]string{
	JingcaiTypeX2: "平局/客胜",
	JingcaiType11: "主胜/主胜",
	JingcaiTypeXX: "平局/平局",
	JingcaiType12: "主胜/客胜",
	JingcaiType1X: "主胜/平局",
	JingcaiType21: "客胜/主胜",
	JingcaiType22: "客胜/客胜",
	JingcaiType2X: "客胜/平局",
	JingcaiTypeX1: "平局/主胜",
}

// JingcaiTypeDxfNameMapping 竞彩大小分名称映射表
var JingcaiTypeDxfNameMapping = map[JingcaiType]string{
	JingcaiTypeOver:  "大于",
	JingcaiTypeUnder: "小于",
}

// JingcaiTypeBfNameMapping 竞彩比分名称映射表 type 中就是比分，以-分割主客队得分。其它选项参照下表
var JingcaiTypeBfNameMapping = map[JingcaiType]string{
	JingcaiType1a: "客队胜的其它比分",
	JingcaiType1d: "平局的其它比分",
	JingcaiType1h: "主队胜的其它比分",
}

// JingcaiTypeSfcNameMapping 胜分差名称映射表
var JingcaiTypeSfcNameMapping = map[JingcaiType]string{
	JingcaiTypeH1: "主队赢1-5分",
	JingcaiTypeH2: "主队赢6-10分",
	JingcaiTypeH3: "主队赢11-15分",
	JingcaiTypeH4: "主队赢16-20分",
	JingcaiTypeH5: "主队赢21-25分",
	JingcaiTypeH6: "主队赢26+分",
	JingcaiTypeA1: "客队赢1-5分",
	JingcaiTypeA2: "客队赢6-10分",
	JingcaiTypeA3: "客队赢11-15分",
	JingcaiTypeA4: "客队赢16-20分",
	JingcaiTypeA5: "客队赢21-25分",
	JingcaiTypeA6: "客队赢26+分",
}

// JingcaiTypeSxdsNameMapping 上下单双名称映射表
var JingcaiTypeSxdsNameMapping = map[JingcaiType]string{
	JingcaiTypeOdd1:  "上单",
	JingcaiTypeEven1: "上双",
	JingcaiTypeOdd2:  "下单",
	JingcaiTypeEven2: "下双",
}
