package main

import (
	"context"
	"fmt"
	"log"

	summermcp "github.com/ikura-hamu/summer-in-japan-is-too-hot-mcp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	// MCPサーバーを作成
	server := summermcp.NewServer()

	// In-memoryトランスポートを作成
	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	// サーバーをgoroutineで実行
	go func() {
		if err := server.Run(ctx, serverTransport); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// クライアントを作成
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "0.0.0"}, nil)

	// クライアントを接続
	session, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Println("In-process MCP client initialized successfully")

	// ツールを呼び出してテスト
	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "make_japan_cool",
		Arguments: map[string]any{},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 結果を表示
	if len(result.Content) > 0 {
		if textContent, ok := result.Content[0].(*mcp.TextContent); ok {
			fmt.Printf("Tool result: %s\n", textContent.Text)
		}
	}

	log.Println("In-process test completed successfully")
}
