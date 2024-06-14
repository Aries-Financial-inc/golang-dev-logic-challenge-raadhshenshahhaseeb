package stderr

const (
	optionsControllerErrCode = (1+iota)*1000 + 1
)

const (
	optionsServiceErrCode = (1+iota)*10000 + 1
	marketPriceServiceErrCode
)

type ErrorCode int

func SuccessCode() int {
	return optionsControllerErrCode - 1
}
