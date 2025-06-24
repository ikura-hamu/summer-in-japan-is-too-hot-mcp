package main

import (
	"log"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	mcpServer := summermcp.NewServer()

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatal(err)
	}
}
