package main

import (
	"context"
	"log"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := summermcp.NewServer()

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
