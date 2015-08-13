package main

import (
	"fmt"
	"regexp"
)

func (n *Node) ParseEosConfig() string {
	return n.RunCmds([]string{"enable", "show running-config"}, "text")
}

func (n *Node) GetHostName() string {
	return n.Configure([]string{"show hostname"})
}

func (n *Node) GetIpRouting() bool {
	// need regexp to find if "no ip routing" exists
	if n.ParseEosConfig().Include("no ip routing") {
		return false
	}
	return true
}

func (n *Node) SetHostName(hostname string) string {
	ins := fmt.Sprintf("hostname %s", hostname)
	return n.Configure([]string{ins})
}

func (n *Node) SetIpRouting() string {
	return n.Configure([]string{"ip routing"})
}