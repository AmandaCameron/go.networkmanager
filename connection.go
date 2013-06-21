package nm

import (
	"launchpad.net/~jamesh/go-dbus/trunk"
)

type Connection struct {
	Data       map[string]map[string]*dbus.Variant
	ObjectPath dbus.ObjectPath
}
