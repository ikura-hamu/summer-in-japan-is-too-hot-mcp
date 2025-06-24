package main

import (
	"context"
	"fmt"
	"log"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// MCPサーバーを作成
	mcpServer := summermcp.NewServer()

	// In-processクライアントを作成
	mcpClient, err := client.NewInProcessClient(mcpServer)
	if err != nil {
		log.Fatal(err)
	}
	defer mcpClient.Close()

	ctx := context.Background()

	// クライアントを初期化
	_, err = mcpClient.Initialize(ctx, mcp.InitializeRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("In-process MCP client initialized successfully")

	// ツールを呼び出してテスト
	result, err := mcpClient.CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "make_japan_cool",
			Arguments: map[string]interface{}{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 結果を表示
	if len(result.Content) > 0 {
		if textContent, ok := mcp.AsTextContent(result.Content[0]); ok {
			fmt.Printf("Tool result: %s\n", textContent.Text)
		}
	}

	log.Println("In-process test completed successfully")
}
