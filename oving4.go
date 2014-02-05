package main

import (
        "net"
        "fmt"
		"time"
)


func main(){

    // TCP
    // CLIENT
    //go tcp_client()

    
    // SERVER
    //go tcp_server()
    
	// UDP
	

	//go udp_reciver()

	go udp_sender()


	neverQuit := make(chan string)
	<- neverQuit
    
}

func tcp_receiver(c net.Conn) {
    defer c.Close()
    for {
        buf := make([]byte, 1024)
        c.Read(buf)
        fmt.Println(string(buf))
    }
}

func tcp_client() {
    serverAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.161:34933")
    PrintError(err)

    con, err := net.DialTCP("tcp", nil, serverAddr);
    PrintError(err)
    
    
    go tcp_receiver(con)

    con.Write([]byte("hello tcp world\x00"))
    con.Write([]byte("so fancy\x00"))


    con.Write([]byte("Connect to: 129.241.187.155:34933\x00"))
}

func tcp_server() {
    listener, err := net.Listen("tcp", ":34933")
    PrintError(err)

    defer listener.Close()
    for {
        newCon, err := listener.Accept()
        PrintError(err)

        go tcp_receiver(newCon)
        newCon.Write([]byte("nc:hello tcp world\x00"))
        newCon.Write([]byte("nc:so fancy\x00"))

    }
}


func udp_reciver() {
    newBuf := make([]byte, 1024)
    addr, err := net.ResolveUDPAddr("udp", ":20020")
    PrintError(err)

    sock, err := net.ListenUDP("udp", addr)
	PrintError(err)

	for {
        _, _, err = sock.ReadFromUDP(newBuf)
        PrintError(err)

        fmt.Println(string(newBuf))
    }
}

func udp_sender() {
    serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20020")
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)

	for {
		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write([]byte("Fremdeles pa?"))
		PrintError(err2)
	}
}


func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}

// udp_sender(newBuf []byte, sock *net.UDPConn)




































