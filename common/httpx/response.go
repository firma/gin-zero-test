package httpx

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"miya/common/errno"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	Response errno.Error

	ContextFunc func(ctx context.Context) Response

	contextKey struct{}
)

var (
	requestKey = contextKey{}
)

func Json(e ContextFunc) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx := ginCtx.Request.Context()

		ctx = context.WithValue(ctx, requestKey, ginCtx.Request)

		f := e(ctx)

		f.RenderJson(ginCtx)
	}
}

func RequestFromServerContext(ctx context.Context) *http.Request {
	val := ctx.Value(requestKey)

	request, ok := val.(*http.Request)
	if !ok {
		return &http.Request{
			Body: io.NopCloser(bytes.NewBuffer([]byte{})),
		}
	}

	return request
}

func RequestHeaderFromServerContext(ctx context.Context) http.Header {
	req := RequestFromServerContext(ctx)

	return req.Header
}

func RawDataFromServerContext(ctx context.Context) ([]byte, error) {
	req := RequestFromServerContext(ctx)

	return ioutil.ReadAll(req.Body)
}

func ShouldBindFromServerContext(ctx context.Context, obj interface{}) error {
	req := RequestFromServerContext(ctx)

	b := binding.Default(req.Method, filterFlags(req.Header.Get("Content-Type")))

	return b.Bind(req, obj)
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}
