package main

import "gtodo/src/api/http"

func main() {
	webServer := http.NewWebServer()
	webServer.LoadRouterV1().Start()
}
