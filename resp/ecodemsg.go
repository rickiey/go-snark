package resp

const (
	// ContextErrCode record log
	ContextErrCode = "context/err/code"
	// OK ok
	OK = 0
	// OkMsg ..
	OkMsg = "SUCCESS"

	// Failed 内部错误
	Failed = -1
	// FailedMsg ..
	FailedMsg = "ERROR"

	// Lose ..
	Lose = 1
	// LoseMsg ..
	LoseMsg = "操作失败"

	// ParamErr ..
	ParamErr = 1000
	// ParamErrMsg ..
	ParamErrMsg = "参数读取错误"
)
