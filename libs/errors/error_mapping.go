package errors

const (
	InvalidBodyRequest       = "invalid_body_request"
	InvalidQueryParam        = "invalid_query_param"
	InternalDBError          = "internal_db_error"
	InvalidWSBody            = "invalid_ws_body"
	FailedUpgradeWebsocket   = "failed_upgrade_websocket"
	HeaderNotPresent         = "header_not_present"
	ExpiredToken             = "expired_token"
	Unauthorized             = "unauthorized"
	InvalidRole              = "invalid_role"
	InvalidParameter         = "invalid_parameter"
	InsufficientBalance      = "insufficient_balance"
	UserNotFound             = "user_not_found"
	InvalidToken             = "invalid_token"
	ErrRecordNotFound        = "data not found"
	FailedPublishOrderStatus = "failed_publish_order_status"
	FailedSetOrderStatus     = "failed_set_order_status"
	InvalidAuthentication    = "invalid_authentication"
	CloseConnection          = "close_connection"
)
