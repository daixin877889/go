package errorCode

type ErrData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrorDBError = ErrData{Code: 50001, Msg: "数据库错误"}
)
