package nettool

import (
	"net"

	"github.com/vishvananda/netlink"
)

// AddRoute adds a universally-scoped route to a device.
func AddRoute(ipn *net.IPNet, gw net.IP, dev netlink.Link) error {
	return netlink.RouteAdd(&netlink.Route{
		LinkIndex: dev.Attrs().Index,
		Scope:     netlink.SCOPE_UNIVERSE,
		Dst:       ipn,
		Gw:        gw,
	})
}

func AddHostRoute(ipn *net.IPNet, gw net.IP, dev netlink.Link) error {
	//TODO:
	return nil
}

func AddDefaultRoute(gw net.IP, dev netlink.Link) error {
	//TODO:
	return nil
}
