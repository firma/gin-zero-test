package errno

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"miya/common/validate"
)

var _ Error = (*errno)(nil)

type (
	Error interface {
		GetCode() int
		GetHttpStatusCode() int
		GetErrMsg() string
		GetReason() string

		GetData() interface{}
		GetRawData() []byte

		// Format 补充格式化输出错误信息
		Format(v ...interface{}) Error

		// WithData 设置成功时返回的数据
		WithData(data interface{}) Error

		// WithReason 错误详细描述
		WithReason(reason interface{}) Error

		// WithID 设置当前请求的唯一ID
		WithID(id string) Error

		// 设置http status code
		WithHttpStatusCode(httpStatusCode int) Error

		ToString() string

		ToBytes() []byte

		Error() string

		Render
	}

	errno struct {
		HttpStatusCode int `json:"-"`
		Code           int `json:"code"` // 业务编码
		// DEPRECATED
		ErrMsg    string      `json:"err"`              // 错误描述
		Msg       string      `json:"msg"`              // 错误描述
		Data      interface{} `json:"data"`             // 成功时返回的数据
		Reason    interface{} `json:"reason,omitempty"` // 错误详细描述
		RequestID string      `json:"rid,omitempty"`    // 当前请求的唯一ID，便于问题定位
		NowTime   int64       `json:"nowTime"`          // 时间戳
	}
)

func NewError(code int, errMsg string) Error {
	return errno{
		HttpStatusCode: http.StatusOK,
		Code:           code,
		ErrMsg:         errMsg,
		Msg:            errMsg,
		Data:           make(map[string]interface{}),
		NowTime:        time.Now().Unix(),
	}
}

func (e errno) Error() string {
	return e.Msg
}

func (e errno) GetCode() int {
	return e.Code
}

func (e errno) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e errno) GetErrMsg() string {
	return e.Msg
}

func (e errno) GetReason() string {
	if v, ok := e.Reason.(string); ok {
		return v
	}

	return ""
}

func (e errno) GetData() interface{} {
	return e.Data
}

// 仅支持string/[]byte类型的data
func (e errno) GetRawData() []byte {
	if s, ok := e.Data.(string); ok {
		return []byte(s)
	}

	if s, ok := e.Data.([]byte); ok {
		return s
	}

	return []byte{}
}

// DEPRECATED 请使用 FormatErrMsg
func (e errno) Format(v ...interface{}) Error {
	e.Msg = fmt.Sprintf(e.Msg, v...)

	e.ErrMsg = e.Msg

	return e
}

func (e errno) FormatErrMsg(v ...interface{}) Error {
	e.Msg = fmt.Sprintf(e.Msg, v...)

	e.ErrMsg = e.Msg

	return e
}

func (e errno) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e errno) WithID(rid string) Error {
	e.RequestID = rid

	return e
}

func (e errno) reset(ctx *gin.Context) Error {
	e.NowTime = time.Now().Unix()

	return e
}

// 如果reason是Error，则会直接使用reason返回
func (e errno) WithReason(reason interface{}) Error {
	if v, ok := reason.(Error); ok {
		return v
	}

	if v, ok := reason.(error); ok {
		e.Reason = validate.TransErr(v).Error()

		return e
	}

	e.Reason = reason
	return e
}

// 请使用net.http status code
func (e errno) WithHttpStatusCode(httpStatusCode int) Error {
	e.HttpStatusCode = httpStatusCode

	return e
}

func (e errno) ToString() string {
	return string(e.ToBytes())
}

func (e errno) ToBytes() []byte {
	if e.Data == nil {
		e.Data = make(map[string]interface{})
	}

	raw, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return raw
}
