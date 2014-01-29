package main

import (
        "net"
        "fmt"
    //"io"
    //"os"
        //"bufio"
)


func main(){

    
    // TCP
    // CLIENT
    go tcp_client()

    
    // SERVER
    go tcp_server()
    

    // UDP

    udp_reciver()
    


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
    if err != nil {
         fmt.Println(err)
    }

    con, err := net.DialTCP("tcp", nil, serverAddr);
    if err != nil {
         fmt.Println(err)
    }
    
    
    go tcp_receiver(con)

    con.Write([]byte("hello tcp world\x00"))
    con.Write([]byte("so fancy\x00"))


    con.Write([]byte("Connect to: 129.241.187.155:34933\x00"))
}



func tcp_server() {
    listener, err := net.Listen("tcp", ":34933")
    if err != nil {
        fmt.Println(err)
    }
    defer listener.Close()
    for {
        newCon, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
        }
        go tcp_receiver(newCon)
        newCon.Write([]byte("nc:hello tcp world\x00"))
        newCon.Write([]byte("nc:so fancy\x00"))

    }
}



func udp_reciver() {
    newBuf := make([]byte, 1024)
    addr, err := net.ResolveUDPAddr("udp", ":20020")
    if err != nil{
	fmt.Println(err)
    }
    sock, err := net.ListenUDP("udp", addr)
    if err != nil{
        fmt.Println(err)
    }
    
    // syntaksfeil her:
    upd_sender(newBuf []byte, sock *UDPConn)
}

func udp_sender(newBuf []byte, sock *UDPConn) {
    serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20020")
    if err != nil{
        fmt.Println(err)
    }
    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    if err != nil{
        fmt.Println(err)
    }


    _, err2 := con_udp.Write([]byte("hello udp world"))
    if err2 != nil{
        fmt.Println(err2)
    }
    fmt.Println("Du har kommet helt hit")

    
    for {
        _, _, err = sock.ReadFromUDP(newBuf)
        if err != nil{
            fmt.Println(err)
        }
        fmt.Println(string(newBuf))
    }


}










