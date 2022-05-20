package config

const (
	ContextKeyBot      = "ctx_bot"
	ContextKeyDbClient = "ctx_db_client"

	ErrorMessage_UnauthorizedUser = "Private chats with me are for authorized users only. Please contact bot owner for assitance."
	ErrorMessage_BadRequest       = "Oh no! It seems your message wasn't in the right format.\nRequired:\namount (10 | 10.5)\ncategory\npayment method"
	ErrorMessage_FailedToSave     = "Oh no! I couldn't save your data...\nPlease try again or contact my owner."

	// https://apps.timwhitlock.info/emoji/tables/unicode
	// Emojis can be sent using UNICODE values
	SuccessMessage_Saved = "\xF0\x9F\x92\xB0 \xF0\x9F\x91\x8D" // :moneybag: :thumbsup:
	StartMessage         = "Hello \xF0\x9F\x99\x8B! I am here to help you be on top of your spending, whenever you make an expense you can send me the details and I'll save it for you :)\nYou can do that in the following format:\n\namount\ncategory\npayment method\n\nExample:\n\n120.5\nאוכל\nאשראי"
)
