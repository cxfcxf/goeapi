package api

import (
	"fmt"
	"strings"
)

func (n *Node) UsersGetAll() []string {
	re := regexp.MustCompile("username")
	var users []string
	config := strings.Split(n.ParseEosRunConfig(), "\n")
    for _, line := range config {
        if re.MatchString(line) {
            users = append(users, line)
        }
    }
    return users
}

func (n *Node) UsersGet(username string) string {
	re := regexp.MustCompile(username)
	for _, user := n.UsersGetAll() {
		if re.MatchString(user) {
			return user
		}
	}
	return "can not find user"
}

func (n *Node) UsersCreate(name string, nopassword bool, secret string, encryption string) string {
	return "place holder"
}

func (n *Node) UsersDelete(name string) string {
	return "place holder"
}

func (n *Node) UsersDefault(name string) string {
	return "place holder"
}

func (n *Node) UsersSetPrivilege(name string, level int) string {
	return "place holder"
}

func (n *Node) UsersSetRole(name string, role string) string {
	return "place holder"
}