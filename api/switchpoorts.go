package api

import (
	"fmt"
)

func (n *Node) SwPortGet(name string) {

}

func (n *Node) SwPortCreate(name string) string {
	var cmds []string{fmt.Sprintf("interface %s", name), "no ip address", "switchport"}
	return n.Configure(cmds)
}