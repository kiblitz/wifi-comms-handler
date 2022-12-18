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
<p float="left">
  <img src="https://github.com/thisistrivial/thisistrivial/blob/master/.assets/wifi-comms-handler/json.png" width="50%">
</p>

`scan` command maps to running a bash script `scanner.sh` carrying over 1 string arg `%s`

#### bash script
<p float="left">
  <img src="https://github.com/thisistrivial/thisistrivial/blob/master/.assets/wifi-comms-handler/bash.png" width="65%">
</p>

bash script runs scanner and outputs to file in `~/Documents` with arg name

#### client source
<p float="left">
  <img src="https://github.com/thisistrivial/thisistrivial/blob/master/.assets/wifi-comms-handler/client-src.png" width="50%">
</p>

client source code run after server start with name `output-file` (line 11)

#### results
<p float="left">
  <img src="https://github.com/thisistrivial/thisistrivial/blob/master/res/wifi-comms-handler/handler-out.png" width="70%">
  <img src="https://github.com/thisistrivial/thisistrivial/blob/master/res/wifi-comms-handler/res.png" width="55%">
</p>

scanned image saved to `~/Downloads` folder
