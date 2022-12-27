package errors

// Error 错误结构体
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error error interface
func (e *Error) Error() string {
	return e.Msg
}

// 自定义错误码

const (
	ErrorCodeDefault            = 100000
	ErrorCodeCmdNotFind         = 100001
	ErrorCodeConfigEnvNotFind   = 100002
	ErrorCodeApplicationNotInit = 100003
	ErrorCodeRedisNotExist      = 100004
	ErrorCodeDBNotExist         = 100005
	ErrorCodeParamTypeError     = 100006
)

// 自定义错误

var (
	ErrorDefault            = &Error{Code: ErrorCodeDefault, Msg: "service error"}
	ErrorCmdNotFind         = &Error{Code: ErrorCodeCmdNotFind, Msg: "cmd not find"}
	ErrorConfigEnvNotFind   = &Error{Code: ErrorCodeConfigEnvNotFind, Msg: "config env not find"}
	ErrorApplicationNotInit = &Error{Code: ErrorCodeApplicationNotInit, Msg: "application not init"}
	ErrorRedisNotExist      = &Error{Code: ErrorCodeRedisNotExist, Msg: "redis not exist"}
	ErrorDBNotExist         = &Error{Code: ErrorCodeDBNotExist, Msg: "db not exist"}
	ErrorParamTypeError     = &Error{Code: ErrorCodeParamTypeError, Msg: "param type error"}
)
