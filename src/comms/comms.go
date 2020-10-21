package comms

import (
  "encoding/json"
  "io/ioutil"
  "os/exec"
  "fmt"
  "os"
  "strings"
)

type Comm struct {
  In string
  Out string
}

func GetComms() []Comm {
  jsonFile, err := os.Open("comms/comms.json")
  if err != nil {
    fmt.Println(err)
    return nil
  }
  defer jsonFile.Close()

  var comms []Comm
  byteValue, err := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &comms)
  return comms
}

func HandleComm(comms []Comm, comm string) string {
  split := strings.Split(comm, " ")
  out := GetOut(comms, split[0])
  if out == "" {
    return fmt.Sprintf(split[0], "command not found")
  }
  execStr := GetExec(out, split[1:])
  execLst := strings.Split(execStr, " ")
  cmd := exec.Command(execLst[0], execLst[1:]...)
  stdout, err := cmd.Output()
  if err != nil {
    fmt.Println(err)
  }
  return string(stdout)
}

func GetOut(comms []Comm, in string) string {
  for _, c := range comms {
    if c.In == in {
      return c.Out
    }
  }
  fmt.Println(in, "command not found")
  return ""
}

func GetExec(out string, params []string) string {
  if len(params) == 0 {
    return out
  }
  return GetExec(fmt.Sprintf(out, params[0]), params[1:])
}
