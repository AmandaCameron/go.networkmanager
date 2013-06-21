package agent

import (
	"launchpad.net/~jamesh/go-dbus/trunk"

	"github.com/AmandaCameron/go.networkmanager"
)

type SecretAgent interface {
	GetSecrets(nm.Connection, string, []string, uint32) map[string]map[string]*dbus.Variant
	CancelGetSecrets(nm.Connection, string)
	SaveSecrets(nm.Connection)
	DeleteSecrets(nm.Connection)
}

func Register(sys *dbus.Connection, sa SecretAgent, name string) error {
	c := make(chan *dbus.Message)

	sys.RegisterObjectPath("/org/freedesktop/NetworkManager/SecretAgent", c)

	obj := sys.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/AgentManager")

	_, err := obj.Call("org.freedesktop.NetworkManager.AgentManager", "Register", name)

	if err != nil {
		return err
	}

	go func() {
		for {
			msg := <-c

			if msg.Type != dbus.TypeMethodCall {
				continue
			}

			if msg.Interface != "org.freedesktop.NetworkManager.SecretAgent" {
				continue
			}

			if msg.Member == "GetSecrets" {
				var conn nm.Connection
				var setting string
				var hints []string
				var flags uint32

				if err := msg.Args(&conn.Data, &conn.ObjectPath, &setting, &hints, &flags); err != nil {
					continue
				}

				res := sa.GetSecrets(conn, setting, hints, flags)

				if res != nil {
					println("Returning!")

					ret := dbus.NewMethodReturnMessage(msg)
					if err := ret.AppendArgs(res); err != nil {
						println("Error Appending:", err.Error())

						ret = dbus.NewErrorMessage(msg,
							"org.freedesktop.Dbus.Error", "Somebody set us up the bomb!")
					}

					sys.Send(ret)
				} else {
					println("Erroring!")

					ret := dbus.NewErrorMessage(msg,
						"org.freedesktop.NetworkManager.SecretAgent.AgentCanceled",
						"Agent Canceled")

					sys.Send(ret)
				}

			} else if msg.Member == "CancelGetSecrets" {
				var conn nm.Connection
				var setting string

				if err := msg.Args(&conn.ObjectPath, &setting); err != nil {
					continue
				}

				sa.CancelGetSecrets(conn, setting)
			} else if msg.Member == "SaveSecrets" {
				var conn nm.Connection

				if err := msg.Args(&conn.Data, &conn.ObjectPath); err != nil {
					continue
				}

				sa.SaveSecrets(conn)
			} else if msg.Member == "DeleteSecrets" {
				var conn nm.Connection

				if err := msg.Args(&conn.Data, &conn.ObjectPath); err != nil {
					continue
				}

				sa.DeleteSecrets(conn)
			} else {
				println("Unknown method:", msg.Member)
			}
		}
	}()

	return nil
}

func Unregister(sys *dbus.Connection) error {
	obj := sys.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/SecretAgent")

	_, err := obj.Call("org.freedesktop.NetworkManager.AgentManager", "Unregister")

	return err
}
