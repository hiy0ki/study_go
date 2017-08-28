package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	user     = flag.String("u", "", "user")
	password = flag.String("p", "", "password")
	port     = flag.Int("P", 22, "port")
)

func run() int {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return 2
	}

	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	hostport := fmt.Sprintf("%s:%d", flag.Arg(0), *port)
	conn, err := ssh.Dial("tcp", hostport, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect %v: %v", hostport, err)
		return 1
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open new session: %v", err)
		return 1
	}
	defer session.Close()

	go func() {
		time.Sleep(5 * time.Second)
		conn.Close()
	}()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	err = session.Run(strings.Join(flag.Args()[1:], " "))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		if ee, ok := err.(*ssh.ExitError); ok {
			return ee.ExitStatus()
		}
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
