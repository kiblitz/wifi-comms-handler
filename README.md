# wifi-comms-handler
## About
Handler for commands sent over LAN via TCP

## Use
- Clone the repo
- Edit `src/comms/comms.json` to include an array of json objects with attributes `in` (wifi command) and `out` (local command)
- Run the program and specify the PORT
```
go run src/handler.go PORT
```
### Server commands
- `quit` - exit the program

## Examples
### Scanner
#### json
```
[
  {
    "in": "scan",
    "out": "/home/gelos/.scripts/scanner.sh %s"
  }
]
```

`scan` command maps to running a bash script `scanner.sh` carrying over 1 string arg `%s`

#### bash script
```
#!/bin/bash

if [ "$#" -eq 1 ]
then
  scanimage --device escl:http://10.0.0.249:80 --format=png > ~/Downloads/$1.png
else
  echo "wrong number of arguments supplied"
fi
```

bash script runs scanner and outputs to file in `~/Documents` with arg name

#### client source
```
import socket
import sys
# Create a TCP/IP socket
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server_address = ('localhost', 2222)
print(sys.stderr, 'connecting to %s port %s' % server_address)
socket.connect(server_address)

socket.sendall(b'scan output-file\n')

while True:
  print(sock.recv(128))
```

client source code run after server start with name `output-file` (line 11)

#### results
```
$ go run handler.go 2222
creating connection with  [127.0.0.1:54468]
[127.0.0.1:54468] /home/gelos/.scripts/scanner.sh [output-file]
```
```
$ ls | grep output-file
output-file.png
```

scanned image saved to `~/Downloads` folder
