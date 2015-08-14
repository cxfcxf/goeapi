package api

import (
    "fmt"
)

func (n *Node) CreateStandardAcl(name string) string {
    return n.Configure([]string{fmt.Sprintf("ip access-list standard %s", name)})
}

func (n *Node) DeleteStandardAcl(name string) string {
    return n.Configure([]string{fmt.Sprintf("no ip access-list standard %s", name)})
}

func (n *Node) DefaultStandardAcl(name string) string {
    return n.Configure([]string{fmt.Sprintf("default ip access-list standard %s", name)})
}

func (n *Node) UpdateEntryStdAcl(name, seqno, action, addr, prefixlen string, log bool) string {
    var cmds []string
    cmds = append(cmds, fmt.Sprintf("ip access-list standard %s", name))
    cmds = append(cmds, fmt.Sprintf("no %s", seqno))
    //entry
    if log {
        cmds = append(cmds, fmt.Sprintf("%s %s %s/%s log", seqno, action, addr, prefixlen))
    } else {
        cmds = append(cmds, fmt.Sprintf("%s %s %s/%s", seqno, action, addr, prefixlen))
    }
    cmds = append(cmds, "exit")
    return n.Configure(cmds)
}

func (n *Node) AddEntryStdAcl(name, action, addr, prefixlen string, log bool) string {
    var cmds []string
    cmds = append(cmds, fmt.Sprintf("ip access-list standard %s", name))
    //entry
    if log {
        cmds = append(cmds, fmt.Sprintf("%s %s/%s log", action, addr, prefixlen))
    } else {
        cmds = append(cmds, fmt.Sprintf("%s %s/%s", action, addr, prefixlen))
    }
    cmds = append(cmds, "exit")
    return n.Configure(cmds)
}

func (n *Node) RemoveEntryStdAcl(name, seqno string) string {
    var cmds []string
    cmds = append(cmds, fmt.Sprintf("ip access-list standard %s", name))
    cmds = append(cmds, fmt.Sprintf("no %s", seqno))
    cmds = append(cmds, "exit")
    return n.Configure(cmds)
}

