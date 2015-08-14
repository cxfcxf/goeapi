package api

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func (n *Node) RunCmds(cmds []string, format string) string {
    data, err := json.Marshal(map[string]interface{}{
            "jsonrpc": "2.0",
            "method": "runCmds",
            "params": map[string]interface{} {
                "version": 1,
                "cmds": cmds,
                "format": format,
                },
            "id": 1,
        })
    if err != nil {
        panic(err)
    }
    eapiurl := fmt.Sprintf("%s://%s:%s@%s/command-api", n.Trans, n.User, n.Pass, n.Ipaddr)

    resp, err := http.Post(eapiurl, "application/json", strings.NewReader(string(data)))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}

func (n *Node) Enable(cmds []string) string {
    return n.RunCmds(append([]string{"enable"}, cmds...), "json")
}

func (n *Node) Configure(cmds []string) string {
    return n.RunCmds(append([]string{"enable", "configure"}, cmds...), "json")
}

func (n *Node) Running_config() string {
    return n.Enable([]string{"show running-config"})
}

func (n *Node) Startup_config() string {
    return n.Enable([]string{"show startup-config"})
}
