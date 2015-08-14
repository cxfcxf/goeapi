package api

import (
	"fmt"
)

func (n *Node) IntConfigure(name string, cmds []string) string {
	return n.Configure(append([]string{fmt.Sprintf("interface %s", name)}, cmds...))
}