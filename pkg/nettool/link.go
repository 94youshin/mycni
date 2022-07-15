package nettool

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

func CreateOrUpdateBridge(name, ip string, mtu int) (*netlink.Bridge, error) {
	//TODO
	return nil, nil
}

func SetupVeth(netns ns.NetNS, br *netlink.Bridge, ifname, ip, gwip string, mtu int) error {
	//TODO
	return nil
}

func makeVethPair(name string, mtu int) (string, *netlink.Veth, error) {
	peerName, err := generateRandomVethName()
	if err != nil {
		return "", nil, err
	}
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name: name,
			MTU:  mtu,
		},
		PeerName: peerName,
	}

	err = netlink.LinkAdd(veth)
	switch {
	case err == nil:
		return peerName, veth, nil
	case os.IsExist(err):
		return "", nil, fmt.Errorf("failed to create veth because veth name %q already exists", name)
	default:
		return "", nil, fmt.Errorf("failed to create veth %q with error: %v", name, err)
	}

}

// generateRandomVethName generate random string start with "veth"
func generateRandomVethName() (string, error) {
	rd := make([]byte, 4)
	if _, err := rand.Read(rd); err != nil {
		return "", fmt.Errorf("failed to generate random veth name: %v", err)
	}
	return fmt.Sprintf("veth%x", rd), nil
}

// GetVethIPInNS return the IP address for the ifName in container namespace
func GetVethIPInNS(netns ns.NetNS, ifName string) (string, error) {
	ip := ""
	err := netns.Do(func(_ ns.NetNS) error {
		l, err := netlink.LinkByName(ifName)
		if err != nil {
			return fmt.Errorf("failed to lookup veth %q in %q: %v", ifName, netns.Path(), err)
		}
		veth, ok := l.(*netlink.Veth)
		if !ok {
			return fmt.Errorf("link %s already exists but is not a veth type", ifName)
		}
		addrs, err := netlink.AddrList(veth, netlink.FAMILY_ALL)
		if err != nil {
			return fmt.Errorf("failed to list address for veth %q:%v", ifName, err)
		}
		switch {
		case len(addrs) > 1:
			return fmt.Errorf("unexpected address for veth %q:%v", ifName, addrs)
		case len(addrs) == 1:
			ip = addrs[0].IPNet.String()
		default:
			return fmt.Errorf("no address set for veth %q", ifName)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return ip, nil
}
