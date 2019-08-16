package errorCode

type ErrData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrorInternalError = ErrData{Code: 50000, Msg: "内部错误"}
	ErrorDBError = ErrData{Code: 50001, Msg: "数据库错误"}
	ErrorParameterError = ErrData{Code: 40002, Msg: "参数错误"}
	ErrorNotAuthUser = ErrData{Code: 40001, Msg: "鉴权失败"}
	ErrorNoPermission = ErrData{Code: 40003, Msg: "无权限"}
	ErrorNotfindrror = ErrData{Code: 40004, Msg: "资源不存在"}
)