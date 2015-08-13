package main

import (
	"fmt"
	"strconv"
)

func IsVlan(vid string) bool {
	id, err := strconv.Atoi(vid)
	if err != nil {
		panic(err)
	}
	if id >= 1 && id <= 4095 {
		return true
	}
	return false
}

func (n *Node) VlanGet(vid string) string {
	if IsVlan(vid) {
		command := fmt.Sprint("show vlan %s", vid)
		return n.Configure([]string{command})
	}
	return "vlan id not valid"
}

func (n *Node) VlanGetAll() string {
	return n.Configure([]string{"show vlan"})
}

func (n *Node) VlanCreate(vid string) string {
	if IsVlan(vid) {
    	command := fmt.Sprintf("vlan %s", vid)
    	return n.Configure([]string{command})
	}
	return "vlan id not valid"
}

func (n *Node) VlanDelete(vid string) string {
	if IsVlan(vid) {
    	command := fmt.Sprintf("no vlan %s", vid)
    	return n.Configure([]string{command})
    }
    return "vlan id not valid"
}

func (n *Node) VlanConfigure(vid string, cmds []string) string {
	if IsVlan(vid) {
		inszero := fmt.Sprintf("vlan %s", vid)
		return n.Configure(append([]string{inszero}, cmds...))
	}
	return "vlan id not valid"
}