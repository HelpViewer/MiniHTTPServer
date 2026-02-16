# MiniHTTPServer

Mini HTTP server for spinning HelpViewer quickly on local environments without changing standard CORS policy on browser apps. Intended primarily for development purposes only, not for production deployment.

## Install

1. Cloning this repository to your local machine
2. OCI (Podman/Docker) image

### Build locally

This call will provide static build of source code (no external lib or binary dependencies).

> CGO_ENABLED=0 \\  
> go build -trimpath -ldflags="-s -w" -o main

The same approach is used for OCI image preparation by CI/CD pipeline.

### OCI image

Base : **scratch**

Linux systems:

> podman pull ghcr.io/helpviewer/minihttpserver  
> mkdir ./www  
> podman run -p 80:8080 -v ./www:/www ghcr.io/helpviewer/minihttpserver  

Webserver is accessible from port **80** via your browser with use of your IP for external network (inside Podman network access via container alias and port 8080).

## Use with local launch

1. Run:
  - buid.bat (on Windows)
  - buid.sh (on Linux)

2. Create directory **www** on the same level as **build** script
3. Run ./main or .\main.exe

### Parameters

Parameters map:  
main --port [port] --dir [directory] --addr [IP address]

| Parameter | Default | meaning |
| --- | --- | --- |
| **--port** | 8080 | HTTP port on which the server listens for requests |
| **--dir** | www/ | The root directory with files to be served |
| **--addr** | 127.0.0.1: | Binding IP address. |

> [!CAUTION]
> **--addr** parameter important notice:  
Value: **127.0.0.1:** (default) keeps the server published only within the computer.  
When you start the server with the value **:**, you publish the server to the connected network so that external computers can access your server.