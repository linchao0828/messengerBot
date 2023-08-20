package consts

const (
	LogId             = "log_id"
	CurrentUserId     = "current_user_id"
	SessionKey        = "messengerBotsession"
	ManagerSessionKey = "mgcsession"
)

const (
	ConfValOn  = "on"  // 开启（图形验证码、邀请码）
	ConfValOff = "off" // 关闭（图形验证码、邀请码）
)

const (
	OK           = 10000
	Fail         = 10001
	BadRequest   = 10400
	NoAuthorized = 10401
	NotFound     = 10404
	AlreadyExist = 10409

	MobileCodeVerifyFail = 11000
	UserNotFound         = 11001
	NotVerify            = 11002
	IdCardVerifyRepeat   = 11003
	ContentIllegal       = 11004
	UserUnderAge         = 11005
	InviteCodeVerifyFail = 11006
	InviteCodeLimited    = 11007
	InviteCodeExpired    = 11008
	ChatCountLimit       = 11009
	GuestUserChatLimit   = 11010
	PayPasswordNotMatch  = 11011

	LockCountFail              = 12001
	OverBuyLimit               = 12002
	SoldOut                    = 12003
	BuySelf                    = 12004
	LockTokenFail              = 12005
	OverGiveCountLimit         = 12006
	OverSmbCountLimit          = 12007
	PreOverBuyLimit            = 12008
	LaunchTimeNotYet           = 12009
	MarketPriceLimit           = 12010
	OverCancelOrderCountLimit  = 12011
	OverAvatarCountLimit       = 12012
	MarketPriceZero            = 12013
	OverWaitPayOrderCountLimit = 12014
	TokenLocked                = 12015
	PassInvisible              = 12016
	TokenExchanged             = 12017
	CharacterGiverSelf         = 12018
	CharacterCannotDelete      = 12019

	// admin

	SkuOverPassQuota       = 20001
	PassNotReady           = 20002
	InviteCodeSummaryExist = 20003
)
