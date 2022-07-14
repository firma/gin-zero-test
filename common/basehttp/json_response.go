package httpresponse

import (
	"github.com/go-playground/validator/v10"
	"miya/common/errorcode"
	util "miya/common/utils"
	validator2 "miya/common/validator"
	"reflect"
)

type Response struct {
	Code    int         `json:"code" xml:"code"`
	Data    interface{} `json:"data,omitempty" xml:"data"`
	Message string      `json:"message" xml:"message"`
	Extra   interface{} `json:"extra,omitempty" xml:"extra"`
}

const ValidationErrorName = "ValidationErrors"

func Success(data ...interface{}) Response {
	response := Response{
		Code:    errorcode.SuccessError.ErrorCode,
		Message: errorcode.SuccessError.ErrorMsg,
	}
	if len(data) == 0 {
		response.Data = map[string]interface{}{}
	} else {
		response.setData(data[0])
	}
	return response
}

func Error(err errorcode.StandardError, data ...interface{}) Response {
	response := Response{
		Code:    err.ErrorCode,
		Message: err.ErrorMsg,
	}

	if data != nil && len(data) == 0 {
		response.Data = map[string]interface{}{}
	} else if len(data) > 0 {
		response.setData(data[0])
	} else {
		response.setData(make([]interface{}, 0))
	}

	return response
}

func (response *Response) setData(data interface{}) {
	r := reflect.TypeOf(data)
	if r.Name() == ValidationErrorName {
		response.Data = map[string]interface{}{}

		// 翻译验证参数错误
		errs := data.(validator.ValidationErrors)
		errMessages := make([]string, 0)

		for _, value := range errs.Translate(validator2.GetTranslator()) {
			errMessages = append(errMessages, util.SnakeString(value))
		}
		response.Extra = errMessages
	} else if r.Kind() == reflect.Slice {
		s := reflect.ValueOf(data)
		if s.Len() == 0 {
			response.Data = map[string]interface{}{}
		} else {
			response.Data = data
		}
	} else {
		response.Data = data
	}
}
