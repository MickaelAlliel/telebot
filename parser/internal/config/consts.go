package config

const (
	ContextKeyBot      = "ctx_bot"
	ContextKeyDbClient = "ctx_db_client"

	ErrorMessage_UnauthorizedUser = "Private chats with me are for authorized users only. Please contact bot owner for assitance."
	ErrorMessage_BadRequest       = "Oh no! It seems your message wasn't in the right format.\nRequired:\namount (10 | 10.5)\ncategory\npayment method"
	ErrorMessage_FailedToSave     = "Oh no! I couldn't save your data...\nPlease try again or contact my owner."

	SuccessMessage_Saved = ":cash: :+1:"
)
