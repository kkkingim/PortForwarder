# PortForwarder
TCP Port Forward Tool

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