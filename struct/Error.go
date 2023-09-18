package _struct

import "fmt"

type ErrorCode int

// NoPass 无密码或密码错误
const NoPass ErrorCode = 1

// OpenDict 打开字典出现错误
const OpenDict ErrorCode = 2

// OpenRar 打开rar压缩包出现错误
const OpenRar ErrorCode = 3

// ReadRar 读取rar压缩包出现错误
const ReadRar ErrorCode = 4

// SuccessPass 成功的密码
const SuccessPass ErrorCode = 5

// OpenZip 打开Zip包出现错误
const OpenZip ErrorCode = 6

// Error 返回错误
type Error struct {
	Is   bool      // 是否有错误
	Code ErrorCode // 错误码
	Msg  string    // 错误信息(人看的)
	Err  error     // 错误信息(详细报错)
}

func NewErr() Error {
	return Error{Is: false}
}

func (e *Error) Print() {
	fmt.Println("出现错误:", e.Msg, " ", e.Err, "错误id", e.Code)
}
