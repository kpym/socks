// socks server
// support socks4, socks4a and socks5
// this file is almost a copy of https://github.com/fangdingjun/socks-go/blob/master/examples/server/server.go
package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "time"

    socks "github.com/fangdingjun/socks-go"
)

const version string = "dev"

func main() {
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "socks (version: %s) : socks4, socks4a and socks5 server.\n  Usage: socks [-h|-p <port>]\n\n", version)
        flag.PrintDefaults()
    }
    port := flag.String("p", "1080", "localhost port to serve")
    flag.Parse()

    conn, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Start server on port :%s.", *port)
    for {
        c, err := conn.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        log.Printf("connected from %s", c.RemoteAddr())

        d := net.Dialer{Timeout: 10 * time.Second}
        s := socks.Conn{Conn: c, Dial: d.Dial}
        go s.Serve()
    }
}
