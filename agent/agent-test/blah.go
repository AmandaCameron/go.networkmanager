package main

import (
	"github.com/AmandaCameron/go.networkmanager"
	"github.com/AmandaCameron/go.networkmanager/agent"
)

type Agent struct {
	// Do Nothing.
}

func (a *Agent) GetSecrets(conn nm.Connection, setting string, hints []string, flags uint32) []map[string]interface{} {
	println("GetSecrets")
	for section, values := range nm.Connection.Data {
		println("==", section, "==")
		for name, value := range values {
			println(" ", name, "=", value)
		}
	}

	return nil
}

func (a *Agent) CancelGetSecrets(conn nm.Connection, setting string) {
	println("CancelGetSecret")
}

func (a *)