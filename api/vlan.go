package api

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

func (n *Node) AddTrunkGroups(vid string, name string) string {
    cmd := fmt.Sprintf("trunk group %s", name)
    return n.VlanConfigure(vid, []string{cmd})
}

func (n *Node) RemoveTrunkGroups(vid string, name string) string {
    cmd := fmt.Sprintf("no trunk group %s", name)
    return n.VlanConfigure(vid, []string{cmd})
}

