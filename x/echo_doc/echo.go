package echo_doc

import (
	"github.com/hwnprsd/go-openapi/x/types"
	"github.com/labstack/echo/v4"
)

type EchoDoc struct {
	app *echo.Echo
	oas *types.OpenAPI
}

func PostRequest[Req, Res any](e *EchoDoc, path string, handler types.PostRequestHandler[Req, Res]) {
}
