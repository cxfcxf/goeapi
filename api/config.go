package api

import (
	"io/ioutil"
	"encoding/json"
)

type Nodes map[string]*Node

type Node struct {
    Ipaddr  string  `json:"ipaddress"`
    User    string  `json:"username"`
    Pass    string  `json:"password"`
    Trans   string  `json:"transport"`
}

type Eoscfg struct {
    Jsonrpc string   `json:"jsonrpc"`
    Result  []map[string]interface{}
    Id      int
}

func (ns *Nodes) ParseEapiConfig(file string) { 
    f, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(f, &ns)
    if err != nil {
        panic(err)
    }
}

func (n *Node) ParseEosRunConfig() string {
    res := []byte(n.RunCmds([]string{"enable", "show running-config"}, "text"))

    var ret []map[string]interface{}
    err := json.Unmarshal(res ,&ret)
    if err != nil {panic(err)}

    if str, ok := ret[1]["output"].(string); ok {
        return str
    }
    return "can not convert interface to string"
}

func (n *Node) ParseEosStartConfig() string {
    res := []byte(n.RunCmds([]string{"enable", "show startup-config"}, "text"))
    
    var ret []map[string]interface{}
    err := json.Unmarshal(res ,&ret)
    if err != nil {panic(err)}

    if str, ok := ret[1]["output"].(string); ok {
        return str
    }
    return "can not convert interface to string"
}