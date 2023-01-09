package unitest

import (
	"github.com/labstack/echo/v4"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

type FakeContext struct {
	Method     string
	Body       string
	Param      map[string]string
	QueryParam map[string]string
	Headers    map[string]string
}

func CreateEchoContext(f FakeContext) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := makeUrl(f)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	initParam2Ctx(f, c)
	return c, rec
}

func makeUrl(f FakeContext) *http.Request {
	q := make(url.Values)
	for k, v := range f.QueryParam {
		q.Set(k, v)
	}
	req := httptest.NewRequest(f.Method, "/?"+q.Encode(), strings.NewReader(f.Body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	initHeaders2Ctx(f, req)
	return req
}

func initParam2Ctx(f FakeContext, ctx echo.Context) {
	for k, v := range f.Param {
		ctx.SetParamNames(k)
		ctx.SetParamValues(v)
	}
}

func initHeaders2Ctx(f FakeContext, req *http.Request) {
	for k, v := range f.Headers {
		req.Header.Set(k, v)
	}
}

func CreateFakeLogger() *rkentry.LoggerEntry {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &rkentry.LoggerEntry{
		Logger: logger,
	}
}
