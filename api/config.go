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

func (ns *Nodes) ParseConfig(file string) { 
    f, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(f, &ns)
    if err != nil {
        panic(err)
    }
}