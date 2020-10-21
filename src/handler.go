package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

func main() {
  args := os.Args
  if len(args) == 1 {
    fmt.Println("no port number provided")
    return
  }

  PORT := ":" + args[1]
  l, err := net.Listen("tcp", PORT)
  if err != nil {
    fmt.Println(err)
    return
  }

  go handleServerComms(l)

  for {
    c, err := l.Accept()
    if err != nil {
      fmt.Println(err)
      return
    }
    go handleConn(c)
  }
}

func handleServerComms(l net.Listener) {
  defer os.Exit(0)
  defer l.Close()
  for {
    text, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
      fmt.Println(err)
    }
    comm := strings.TrimSpace(text) 
    switch {
      case comm == "quit":
        return
    }
  } 
}

func handleConn(c net.Conn) {
  defer fmt.Println("closing connection with ", c.RemoteAddr().String())
  fmt.Println("creating connection with ", c.RemoteAddr().String())
  for {
    data, err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
      fmt.Println(err)
      return
    }
    msg := strings.TrimSpace(data)
    switch {
      case msg == "quit":
        return 
    }
  }
}
