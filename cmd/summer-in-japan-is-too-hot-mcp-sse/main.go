package main

import (
	"flag"
	"log"
	"net/http"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	var port = flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	server := summermcp.NewServer()

	log.Printf("Starting SSE MCP server on port %s", *port)
	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		return server
	}, nil)
	if err := http.ListenAndServe(":"+*port, handler); err != nil {
		log.Fatal(err)
	}
}
