package http

import (
	"gtodo/src"
	v1 "gtodo/src/api/http/v1"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	frameWork *echo.Echo
}

func (ws *WebServer) LoadRouterV1() src.IWebServer {
	var routerV1 *v1.RouterV1 = v1.NewRouterV1(ws.frameWork)
	routerV1.LoadAPI()
	return ws
}

func (ws *WebServer) Start() {
	err := ws.frameWork.Start(":8000")
	if err != nil {
		ws.frameWork.Logger.Fatal(err)
	}
}

func NewWebServer() src.IWebServerSetup {
	return &WebServer{
		echo.New(),
	}
}
