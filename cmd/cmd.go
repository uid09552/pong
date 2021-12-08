package main

import "headerexplorer/pkg/server"

func main() {
	server := server.Server{}
	server.Serve()
}
