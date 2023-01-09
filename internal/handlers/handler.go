package handlers

import rkecho "github.com/rookie-ninja/rk-echo/boot"

type Handler interface {
	RegisterAPI(echoEntry *rkecho.EchoEntry)
}
