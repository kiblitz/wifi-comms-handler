package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"

  "./comms"
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

  MAP := comms.GetComms()

  go HandleServerComms(l)

  for {
    c, err := l.Accept()
    if err != nil {
      fmt.Println(err)
      return
    }
    go HandleConn(MAP, c)
  }
}

func HandleServerComms(l net.Listener) {
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

func HandleConn(MAP []comms.Comm, c net.Conn) {
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
      default:
        res := comms.HandleComm(MAP, msg)
        c.Write([]byte(res))
    }
  }
}
