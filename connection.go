package nm

import (
	"errors"

	"launchpad.net/~jamesh/go-dbus/trunk"
)

type Connection struct {
	proxy      *dbus.ObjectProxy
	Data       map[string]map[string]*dbus.Variant
	ObjectPath dbus.ObjectPath
}

func newConnection(conn *dbus.Connection, objPath dbus.ObjectPath) (*Connection, error) {
	c := &Connection{
		ObjectPath: objPath,
	}

	c.proxy = conn.Object(NM_UNIQ_NAME, c.ObjectPath)

	if err := c.Refresh(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Connection) Refresh() error {
	if c.proxy == nil {
		return errors.New("Proxy not set.")
	}

	ret, err := c.proxy.Call(NM_CONN_IFACE, "GetSettings")
	if err != nil {
		return err
	}

	if err = ret.Args(&c.Data); err != nil {
		return err
	}

	return nil
}
