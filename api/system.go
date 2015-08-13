package api

import (
	"fmt"
	"encoding/json"
	"strings"
	"regexp"
)

type Rcfg struct {
	Jsonrpc	string             `json:"jsonrpc"`
	Result 	[]map[string]interface{}
	Id		int
}

func (n *Node) ParseEosConfig() string {
	res := []byte(n.RunCmds([]string{"enable", "show running-config"}, "text"))
	var ret Rcfg
	err := json.Unmarshal(res ,&ret)
	if err != nil {panic(err)}
	if str, ok := ret.Result[1]["output"].(string); ok {
		return str
	}
	return "can not convert interface to string"
}

func (n *Node) GetHostName() string {
	re := regexp.MustCompile("hostname")
	config := strings.Split(n.ParseEosConfig(), "\n")
	for _, line := range config {
		if re.MatchString(line) {
			return strings.Split(line, " ")[1]
		}
	}
	return "No Hostname Defined"
}

func (n *Node) GetIpRouting() bool {
	config := strings.Split(n.ParseEosConfig(), "\n")
	for _, line := range config {
		if line == "ip routing" {
			return true
		}
	}
	return false
}

func (n *Node) SetHostName(hostname string) string {
	ins := fmt.Sprintf("hostname %s", hostname)
	return n.Configure([]string{ins})
}

func (n *Node) SetIpRouting() string {
	return n.Configure([]string{"ip routing"})
}