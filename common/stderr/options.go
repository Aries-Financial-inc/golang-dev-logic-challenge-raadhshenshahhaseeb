package stderr

// Service Layer Errors

const (
	ErrGettingUnderlyingPriceFromToken ErrorCode = iota + optionsServiceErrCode
)

// Controller Layer Errors

const (
	ErrBindingTheRequest ErrorCode = iota + optionsControllerErrCode
	ErrBidCannotBeGreaterThanAsk
	ErrBidAndAskCannotBeNegative
	ErrStrikeCannotBeNegative
	ErrExpiryCannotBeInPast
	ErrPositionNotSupportedOrInvalid
	ErrQuoteNotSupportedOrInvalid
)
