package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

const (
	CommonHost = "localhost"
	CommonPort = 1074
)

func handleListen(host string, port int) (int, error) {
	/*
	 AF_INET = Address Family for IPv4
	 ***AF_INET6 Support IPv6 && AF_INET support
	 ***IPv4, AF_UNSPEC Support Both
	 SOCK_STREAM = virtual circuit service
	*/
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		os.NewSyscallError("Error on listening this port", err)
	}

	// Bind the socket to a port
	// SockaddrInet4 --> A Structure, An Ip and a Port
	sa := &syscall.SockaddrInet4{Port: CommonPort}
	//LookUp to Ip
	addrs, err := net.LookupHost(CommonHost)
	if err != nil {
		os.NewSyscallError("Error on convert", err)
	}
	for _, addr := range addrs {
		ip := net.ParseIP(addr).To4()
		copy(sa.Addr[:], ip)
		if err = syscall.Bind(fd, sa); err != nil {
			os.NewSyscallError("bind", err)
		}
	}
	//Listen
	if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		os.NewSyscallError("Error On Listen", err)
	} else {
		fmt.Println("Listen On", CommonHost, ":", CommonPort)
	}
	if err != nil {
		return -1, os.NewSyscallError("Error On Listening This Pors", err)
	}
	return fd, nil
}

func handleConnection(clientSock int, clientAddr syscall.Sockaddr) {
	//Send Message
	message := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"Content-Length: 20\r\n" +
		"\r\n" +
		"<h1>hello world</h1>"
	time.Sleep(150 * time.Millisecond)

	err := syscall.Sendmsg(clientSock, []byte(message), []byte{}, clientAddr, 0)

	if err != nil {
		os.NewSyscallError("Error On Send...", err)
	}

	//Close Connection After SendMsg
	syscall.Close(clientSock)
}

func main() {
	fd, err := handleListen(CommonHost, CommonPort)
	if err != nil {
		os.NewSyscallError("Error On Run Function", err)
	}
	for {
		//Accept
		clientSock, clientAddr, err := syscall.Accept(fd)
		log.Printf("Incoming connection")

		if err != nil {
			os.NewSyscallError("Error On Accept...", err)
		}
		go handleConnection(clientSock, clientAddr)
	}
}
