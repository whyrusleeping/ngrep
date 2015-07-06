package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	timeout := flag.Int("t", 10, "set a timeout in seconds for the read")
	listen := flag.Bool("l", false, "start a listener instead of a dial")
	flag.Parse()

	if len(flag.Args()) < 3 {
		fmt.Printf("USAGE: %s <regex> <host> <port>\n", os.Args[0])
	}
	term := flag.Args()[0]
	host := flag.Args()[1]
	port := flag.Args()[2]

	target := fmt.Sprintf("%s:%s", host, port)

	var con net.Conn
	if *listen {
		list, err := net.Listen("tcp", target)
		checkError(err)

		ncon, err := list.Accept()
		checkError(err)

		con = ncon
	} else {
		ncon, err := net.Dial("tcp", target)
		checkError(err)

		con = ncon
	}

	err := con.SetReadDeadline(time.Now().Add(time.Second * time.Duration(*timeout)))
	checkError(err)

	r := bufio.NewReaderSize(con, 1)

	match, err := regexp.MatchReader(term, r)
	checkError(err)

	if !match {
		os.Exit(1)
	}
}
