package nm

import (
	"launchpad.net/~jamesh/go-dbus/trunk"
)

func (dev *Device) Mac() (string, error) {
	tmp, err := dev.devTypeGet("HwAddress")
	if err != nil {
		return "", err
	}

	return tmp.(string), nil
}

func (dev *Device) PermMac() (string, error) {
	tmp, err := dev.devTypeGet("PermHwAddress")
	if err != nil {
		return "", err
	}

	return tmp.(string), nil
}

func (dev *Device) AvailConnections() ([]*Connection, error) {
	tmp, err := dev.Get(NM_DEV_IFACE, "AvailableConnections")
	if err != nil {
		return nil, err
	}

	var conns []*Connection

	for _, connObjPath := range tmp.([]interface{}) {
		conns = append(conns, &Connection{
			ObjectPath: connObjPath.(dbus.ObjectPath),
			Data:       nil,
		})
	}

	for _, conn := range conns {
		if err := conn.Refresh(); err != nil {
			return nil, err
		}
	}

	return conns, nil
}
