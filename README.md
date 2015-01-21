[![Build Status](https://drone.io/github.com/twcclan/go-guid/status.png)](https://drone.io/github.com/twcclan/go-guid/latest)

# go-guid
A webservice to compute W:ET GUIDs from etkeys

# download
Get pre-built binary here: https://drone.io/github.com/twcclan/go-guid/files

# performance / resource requirements (as tested on a small virtual machine, single 2 GHz Opteron core on Xen, 256 MB RAM)
The server uses about 1 KB of resident memory during normal operation, about half as much when idle.
Generating a single GUID takes about 10 µs, which in turn means the server mentioned above could generate about 100k GUIDs per second.