package main

import (
    "fmt"
    goeapi "./api"
)

func main() {
    var nodes goeapi.Nodes
    nodes.ParseEapiConfig("config.json")

    //send := []string{"ip address 10.200.6.1/24"}
    //fmt.Println(nodes["arista-laxo"].Configure(send))
    //send := []string{"hostname arista-home"}
    fmt.Println(nodes["arista-laxo"].GetHostName())
    //fmt.Println(nodes["arista-laxo"].RunCmds([]string{"show vlan"}))
}
