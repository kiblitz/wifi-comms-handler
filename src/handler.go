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
  addr := "[" + c.RemoteAddr().String() + "]"
  defer fmt.Println("closing connection with ", addr)
  fmt.Println("creating connection with ", addr)
  for {
    data, err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
      fmt.Println(addr, err)
      return
    }
    msg := strings.TrimSpace(data)
    switch {
      case msg == "quit":
        return 
      default:
        c.Write([]byte("starting task\n"))
        res := comms.HandleComm(addr, MAP, msg)
        c.Write([]byte(res + "finished task\n"))
    }
  }
}
