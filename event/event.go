package event

const (
	EVENT_BASE = 0x10000000

	// User Registry
	EVENT_SEND_REGISTRY_CODE_VIA_SMS = EVENT_BASE + 101
	EVENT_SEND_RESET_CODE_VIA_SMS    = EVENT_BASE + 102
	EVENT_SEND_RESET_CODE_VIA_EMAIL  = EVENT_BASE + 103
	EVENT_SEND_EMAIL_REGISTRY_CODE   = EVENT_BASE + 104
	EVENT_USER_REGISTERED            = EVENT_BASE + 105
	EVENT_USER_SIGNIN                = EVENT_BASE + 106
	EVENT_USER_SIGNOUT               = EVENT_BASE + 107
	EVENT_USER_UPDATE_PROFILE        = EVENT_BASE + 108
	EVENT_USER_UPDATE_LOCATION       = EVENT_BASE + 109
	EVENT_USER_UPDATE_SETTINGS       = EVENT_BASE + 110
	EVENT_USER_UPDATE_SIGNATURE      = EVENT_BASE + 111
	EVENT_USER_UPDATE_DEVICETOKEN    = EVENT_BASE + 112

	// 用户邮箱或手机验证
	EVENT_USER_VERIFIED_BY_PHONE = EVENT_BASE + 120
	EVENT_USER_VERIFIED_BY_EMAIL = EVENT_BASE + 121
	EVENT_USER_CODE_VERIFIED_OK  = EVENT_BASE + 122
	EVENT_USER_RESET_PASSWORD    = EVENT_BASE + 123
	EVENT_RESET_PASSWORD         = EVENT_BASE + 124

	// Change friendship
	EVENT_FRIEND_FOLLOW         = EVENT_BASE + 201
	EVENT_FRIEND_UNFOLLOW       = EVENT_BASE + 202
	EVENT_FRIEND_BATCH_FOLLOW   = EVENT_BASE + 203
	EVENT_FRIEND_BATCH_UNFOLLOW = EVENT_BASE + 204
	EVENT_FRIEND_CLOSEFRIEND    = EVENT_BASE + 205
	EVENT_FRIEND_UNCLOSEFRIEND  = EVENT_BASE + 206
	EVENT_FRIEND_BLOCK          = EVENT_BASE + 207
	EVENT_FRIEND_UNBLOCK        = EVENT_BASE + 208
	EVENT_FRIEND_UPDATE_NOTE    = EVENT_BASE + 209
	EVENT_FRIEND_UPDATE_COUNT   = EVENT_BASE + 210

	// Feed operations
	EVENT_FEED_CREATE         = EVENT_BASE + 301
	EVENT_FEED_UPDATE         = EVENT_BASE + 302
	EVENT_FEED_REMOVE         = EVENT_BASE + 303
	EVENT_FEED_LIKE           = EVENT_BASE + 304
	EVENT_FEED_UNLIKE         = EVENT_BASE + 305
	EVENT_FEED_COMMENT_POST   = EVENT_BASE + 306
	EVENT_FEED_COMMENT_REMOVE = EVENT_BASE + 307
	EVENT_FEED_VIEW           = EVENT_BASE + 308

	// Tag
	EVENT_TAG_CREATE = EVENT_BASE + 401
	EVENT_TAG_UPDATE = EVENT_BASE + 402
	EVENT_TAG_REMOVE = EVENT_BASE + 403

	// View History 用户间互相查看的记录
	EVENT_VIEW_USER_PROFILE = EVENT_BASE + 601

	// 各种作弊
	EVENT_CHEATING_LIKE_FEED = EVENT_BASE + 700 // 点赞
	EVENT_CHEATING_USER_AUTH = EVENT_BASE + 701 // 用户启动App

	// Notify Push
	EVENT_NOTIFY_FEED_LIKE      = EVENT_BASE + 800 //有人给我点赞了
	EVENT_NOTIFY_FEED_COMMENT   = EVENT_BASE + 801 //有人给我评论了
	EVENT_NOTIFY_FRIEND_FOLLOW  = EVENT_BASE + 802 //有人关注我了
	EVENT_NOTIFY_READ           = EVENT_BASE + 803 //已读通知
	EVENT_NOTIFY_COMMENT_REMOVE = EVENT_BASE + 804 //删除评论

	// Trend Push
	EVENT_TREND_FEED_LIKE     = EVENT_BASE + 804 //好友赞了某个FEED
	EVENT_TREND_FRIEND_FOLLOW = EVENT_BASE + 805 //好友关注了某个人
	EVENT_TREND_READ          = EVENT_BASE + 806 //已读动态

	// Chat Push
	EVENT_CHAT_CREATE = EVENT_BASE + 900 //添加一条聊天记录
	EVENT_CHAT_DELETE = EVENT_BASE + 901 //删除一条聊天记录
)
