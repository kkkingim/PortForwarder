# PortForwarder
TCP Port Forward Tool

Usage: 
  
  ./ktun <localIP>:<localPort>:<remoteIP>:<remotePort>

Example:
  
  *Forward router page to current machine*
  
  ./ktun 0.0.0.0:80:192.168.1.1:80


# Build

## Linux & Macos
```sh
go build -ldflags="-w -s" .
```

## Windows
1. With CMD GUI
```sh
go build -ldflags="-w -s" .
```

2. As Service
```sh
go build -ldflags="-w -s -H=windowsgui" .
```
