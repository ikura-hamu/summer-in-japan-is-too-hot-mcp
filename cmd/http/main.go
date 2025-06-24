package main

import (
	"flag"
	"log"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	var port = flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	mcpServer := summermcp.NewServer()

	log.Printf("Starting HTTP MCP server on port %s", *port)
	httpServer := server.NewStreamableHTTPServer(mcpServer)
	if err := httpServer.Start(":" + *port); err != nil {
		log.Fatal(err)
	}
}
