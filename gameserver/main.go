package main

import (
    "github.com/tiancaiamao/ouster/config"
    "log"
    "net"
)

func main() {
    log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
    log.Println("Starting the server.")

    Initialize()

    listener, err := net.Listen("tcp", config.GameServerPort)
    checkError(err)

    log.Println("Game Server OK.")

    for {
        conn, err := listener.Accept()
        if err == nil {
            go handleClient(conn)
        }
    }
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func handleClient(conn net.Conn) {
    log.Println("accept a connection...")

    agent := NewAgent(conn)
    go agent.Loop()
}
