package data

const (
	GAME_PROFIT               = "游戏损益"
	CUSTOMER_LOSS             = "充提客损"
	RECHARGE_AMOUNT           = "充值总额"
	RECHARGE_PEOPLE           = "充值人數"
	FIRST_RECHARGE_AMOUNT     = "首充金额"
	FIRST_RECHARGE_PEOPLE     = "首充人數"
	WITHDRAWAL_TOTAL          = "提现总额"
	WITHDRAWAL_PEOPLE         = "提现人數"
	EXCHANGE_DIAMOND_AMOUNT   = "兑换钻石金额"
	EXCHANGE_DIAMOND_QUANTITY = "兑换钻石数量"
	EXCHANGE_DIAMOND_PEOPLE   = "兑换钻石人数"
	RECHARGE_ONLINE           = "在线充值"
	RECHARGE_BACKGROUND       = "后台充值"
	RECHARGE_CUSTOMER         = "客单充值"
	DEPOSIT_COMPANY           = "公司入款"
	BETS_PEOPLE               = "投注人数"
	BETS_VALID                = "有效投注"
	LOTTERY_TICKETS           = "彩票分红票数"
	LOGIN_PEOPLE              = "登录人数"
	LIVE_PEOPLE               = "直播中"
	ONLINE_PEOPLE             = "在线人数"
	REGISTRANT_PEOPLE         = "注册人数"
	BOUND_PEOPLE              = "绑定人数"
)

const (
	TITLE_NAME  = "標題"
	COLUMN_NAME = "欄位"
)

var title = map[string]string{}
var name = map[string]string{}

func init() {
	title = make(map[string]string)
	name = make(map[string]string)

	//title
	title[GAME_PROFIT] = "提现"
	title[CUSTOMER_LOSS] = "提现"
	title[RECHARGE_AMOUNT] = "充值"
	title[RECHARGE_PEOPLE] = "充值"
	title[FIRST_RECHARGE_AMOUNT] = "首充"
	title[FIRST_RECHARGE_PEOPLE] = "首充"
	title[WITHDRAWAL_TOTAL] = "提现"
	title[WITHDRAWAL_PEOPLE] = "提现"
	title[EXCHANGE_DIAMOND_AMOUNT] = "钻石兑换"
	title[EXCHANGE_DIAMOND_QUANTITY] = "钻石兑换"
	title[EXCHANGE_DIAMOND_PEOPLE] = "钻石兑换"
	title[RECHARGE_ONLINE] = "充值类型"
	title[RECHARGE_BACKGROUND] = "充值类型"
	title[RECHARGE_CUSTOMER] = "充值类型"
	title[DEPOSIT_COMPANY] = "充值类型"
	title[BETS_PEOPLE] = "游戏"
	title[BETS_VALID] = "游戏"
	title[LOTTERY_TICKETS] = "印票"
	title[LOGIN_PEOPLE] = "活跃"
	title[LIVE_PEOPLE] = "在线"
	title[ONLINE_PEOPLE] = "在线"
	title[REGISTRANT_PEOPLE] = "会员"
	title[BOUND_PEOPLE] = "会员"

	// name
	name[GAME_PROFIT] = "游戏损益"
	name[CUSTOMER_LOSS] = "客损"
	name[RECHARGE_AMOUNT] = "总额"
	name[RECHARGE_PEOPLE] = "人数"
	name[FIRST_RECHARGE_AMOUNT] = "金额"
	name[FIRST_RECHARGE_PEOPLE] = "人数"
	name[WITHDRAWAL_TOTAL] = "总额"
	name[WITHDRAWAL_PEOPLE] = "人数"
	name[EXCHANGE_DIAMOND_AMOUNT] = "金额"
	name[EXCHANGE_DIAMOND_QUANTITY] = "数量钻"
	name[EXCHANGE_DIAMOND_PEOPLE] = "人数"
	name[RECHARGE_ONLINE] = "在线"
	name[RECHARGE_BACKGROUND] = "后台(赠)"
	name[RECHARGE_CUSTOMER] = "客单(赠)"
	name[DEPOSIT_COMPANY] = "银行(赠)"
	name[BETS_PEOPLE] = "投注人数"
	name[BETS_VALID] = "有效投注"
	name[LOTTERY_TICKETS] = "主播彩票分红票"
	name[LOGIN_PEOPLE] = "登录人数"
	name[LIVE_PEOPLE] = "直播中"
	name[ONLINE_PEOPLE] = "在线人数"
	name[REGISTRANT_PEOPLE] = "注册量"
	name[BOUND_PEOPLE] = "绑定量"
}

func GetSearchInfo(searchName, dataName string) string {

	switch searchName {
	case TITLE_NAME:
		return title[dataName]
	case COLUMN_NAME:
		return name[dataName]
	default:
		return ""
	}
}

func GetSearchTitleKey(searchName, dataName string) []string {
	var data []string
	for _, v := range GetSearchList() {

		if GetSearchInfo(searchName, v) == dataName {
			data = append(data, v)
		}
	}
	return data
}

func GetSearchList() []string {
	return []string{
		GAME_PROFIT, CUSTOMER_LOSS, RECHARGE_AMOUNT,
		RECHARGE_PEOPLE, FIRST_RECHARGE_AMOUNT,
		FIRST_RECHARGE_PEOPLE, WITHDRAWAL_TOTAL,
		WITHDRAWAL_PEOPLE, EXCHANGE_DIAMOND_AMOUNT,
		EXCHANGE_DIAMOND_QUANTITY, EXCHANGE_DIAMOND_PEOPLE,
		RECHARGE_ONLINE, RECHARGE_BACKGROUND, RECHARGE_CUSTOMER,
		DEPOSIT_COMPANY, BETS_PEOPLE, BETS_VALID, LOTTERY_TICKETS,
		LOGIN_PEOPLE, LIVE_PEOPLE, ONLINE_PEOPLE, REGISTRANT_PEOPLE,
		BOUND_PEOPLE,
	}

}
