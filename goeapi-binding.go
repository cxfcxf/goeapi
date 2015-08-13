package main

import (
    "fmt"
    goeapi "github.com/cxfcxf/goeapi/api"
)

func main() {
    var nodes goeapi.Nodes
    nodes.ParseConfig("config.json")

    //send := []string{"ip address 10.200.6.1/24"}
    //fmt.Println(nodes["arista-laxo"].Configure(send))
    fmt.Println(nodes["arista-laxo"].GetHostName())
    //fmt.Println(nodes["arista-laxo"].RunCmds([]string{"show vlan"}))
}