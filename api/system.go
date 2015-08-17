package api

import (
    "fmt"
    "strings"
    "regexp"
)

func (n *Node) GetHostName() string {
    re := regexp.MustCompile("hostname")
    config := strings.Split(n.ParseEosRunConfig(), "\n")
    for _, line := range config {
        if re.MatchString(line) {
            return strings.Split(line, " ")[1]
        }
    }
    return "No Hostname Defined"
}

func (n *Node) GetIpRouting() bool {
    config := strings.Split(n.ParseEosRunConfig(), "\n")
    for _, line := range config {
        if line == "ip routing" {
            return true
        }
    }
    return false
}

func (n *Node) SetHostName(hostname string) string {
    return n.Configure([]string{fmt.Sprintf("hostname %s", hostname)})
}

func (n *Node) SetIpRouting() string {
    return n.Configure([]string{"ip routing"})
}

func (n *Node) WriteMemory() string {
    return n.Configure([]string{"write memory"})
}
