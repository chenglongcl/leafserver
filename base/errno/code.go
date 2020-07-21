package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	//2-00-账户相关
	LoginFail = &Errno{Code: 20001, Message: "登录失败"}
)
