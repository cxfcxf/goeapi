package api

import (
	"bytes"
	"strings"
	"golang.org/x/crypto/ssh"
)

func PasswordKeyboardInteractive(password string) ssh.KeyboardInteractiveChallenge {
	return func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		answers := make([]string, len(questions))
		// for all question, send password as the answer
		
		for i, _ := range answers {
			answers[i] = string(password)
		}
		return answers, nil
	}
}

func (n *Node) Ssh(interactive bool,cmds []string) string {

	config := new(ssh.ClientConfig)

	if interactive {
		config = &ssh.ClientConfig{
			User: n.User,
			Auth: []ssh.AuthMethod{
				PasswordKeyboardInteractive(n.Pass),
			},
		}
	} else {
	    config = &ssh.ClientConfig{
	    	User: n.User,
	    	Auth: []ssh.AuthMethod{
	    		ssh.Password(n.Pass),
	    	},
	    }
	}

    client, err := ssh.Dial("tcp", n.Ipaddr+":22", config)
    if err != nil {
    	panic(err)
    }

    var so bytes.Buffer
    var se bytes.Buffer

	session, err := client.NewSession()
    if err != nil {
    	panic(err)
    }

    session.Stdout = &so
    session.Stderr = &se

    if err := session.Run(strings.Join(cmds, "\n")); err != nil {
    	panic(err)
    }

    session.Close()

    if len(se.String()) > 0 {
    	return se.String()
    }

    return so.String()

}