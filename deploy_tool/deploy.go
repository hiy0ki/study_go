package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// 引数でIPのリストとコマンドのリストのファイル名を受け取る
	// TODO flagパッケージへ変更する
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("argment error")
	}

	fmt.Println(args)

	// コマンドのリストをメモリへ
	fp, err := os.Open(args[1])
	if err != nil {
		log.Fatal("command file open error")
	}
	defer fp.Close()

	commands := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// 空行はスキップをする
		if scanner.Text() == "" {
			continue
		}
		// 配列に格納する
		commands = append(commands, scanner.Text())
		fmt.Println(scanner.Text())
	}

	fmt.Println(commands)
	fmt.Println(len(commands))
	// IPのリストからIPを一つ取り出し、接続をする
	// ファイルのコマンドを実行する
	// 結果を標準出力へ

	ip := "192.168.33.10"
	port := "22"
	user := "vagrant"

	buf, err := ioutil.ReadFile("")
	if err != nil {
		panic(err)
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

	fmt.Println("hogehogehoge")
	/*
		for _, command := range commands {
			//session.Run(command)
			o, _ := session.Output(command)
			fmt.Println(o)
		}
	*/
	//session.Run("echo \"hogehoge\" > empty.txt")
}
