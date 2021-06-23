package error_code

const (
	Success        = 0
	PermissionDeny = -1
	InternalError  = 50000

	InputError      = 1001
	TbkRequestError = 4000
)

var errMsg = map[int]string{
	Success:        "OK",
	PermissionDeny: "没有权限。",
	InternalError:  "服务器错误。",

	InputError:      "请求参数错误。",
	TbkRequestError: "请求淘宝客出错。",
}

func GetErrorMsg(code int, extra_msg string) string {
	if v, found := errMsg[code]; found {
		return v + extra_msg
	}
	return "未知错误。" + extra_msg
}
