package main

import (
    "fmt"
)

func main() {
    var nodes Nodes
    nodes.ParseConfig("config.json")

    //send := []string{"ip address 10.200.6.1/24"}
    //fmt.Println(nodes["arista-laxo"].Configure(send))
    fmt.Println(nodes["arista-laxo"].VlanConfigure("6", send))
    //fmt.Println(nodes["arista-laxo"].RunCmds([]string{"show vlan"}))
}