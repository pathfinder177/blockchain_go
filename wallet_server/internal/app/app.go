package app

import (
	"wallet_server/internal/controller/server"
)

const (
	appPort string = ":3003"
)

func Run() {
	server.StartServer(appPort)
}
