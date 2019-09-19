package errorCode

type ErrData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrorInternalError       = ErrData{Code: 50000, Msg: "内部错误"}
	ErrorDBError             = ErrData{Code: 50001, Msg: "数据库错误"}
	ErrorGeWchatUserInfo     = ErrData{Code: 50001, Msg: "获取微信用户信息失败"}
	ErrorParameterError      = ErrData{Code: 40002, Msg: "参数错误"}
	ErrorNotAuthUser         = ErrData{Code: 40001, Msg: "鉴权失败"}
	ErrorNoPermission        = ErrData{Code: 40003, Msg: "无权限"}
	ErrorNotfind             = ErrData{Code: 40004, Msg: "资源不存在"}
	ErrorNotfindWechatConfig = ErrData{Code: 40014, Msg: "微信配置不存在"}
	ErrorWechatAsToken       = ErrData{Code: 40015, Msg: "微信access_token获取错误"}
	ErrorVerifytCodeError    = ErrData{Code: 40024, Msg: "验证码错误"}
)
