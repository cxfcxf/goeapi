package main

import (
	"os"
    "fmt"
    "flag"
    goeapi "./api"
)

var eapi = flag.String("eapi", "status", "status/enable/disable")


func main() {
	flag.Parse()

    var nodes goeapi.Nodes
    nodes.ParseEapiConfig("config.json")

    cmds := []string{}

    if *eapi == "status" {
    	cmds = []string{"show management api http-commands"}
    } else if *eapi == "enable" {
    	cmds = []string{"enable", "config", "management api http-commands", "protocol http", "no shutdown"}
    } else if *eapi == "disable" {
    	cmds = []string{"enable", "config", "no management api http-commands"}
    } else {
    	flag.PrintDefaults()
    	os.Exit(1)
    }

    // defualt is keyboardinteractive=true for most of network device
    // only false if you do a "auth mode password"
    fmt.Println(nodes["arista-laxo"].Ssh(true ,cmds))
}
