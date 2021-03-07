package main

import (
	"gtodo/src"
	"gtodo/src/api/http"
	"gtodo/src/infra/service"

	"github.com/golobby/container"
)

func main() {
	container.Singleton(func() src.IContextService {
		return service.NewServiceContext()
	})
	webServer := http.NewWebServer()
	webServer.LoadRouterV1().Start()
}
