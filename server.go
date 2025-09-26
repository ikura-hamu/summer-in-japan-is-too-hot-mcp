package summerinjapanistoohotmcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// coolJapanResponse は make_japan_cool ツールのレスポンス構造体
type coolJapanResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// shortenJapanSummerResponse は shorten_japan_summer ツールのレスポンス構造体
type shortenJapanSummerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewServer() *server.MCPServer {
	// MCPサーバーを作成
	mcpServer := server.NewMCPServer(
		"summer-in-japan-is-too-hot-mcp",
		"0.0.0",
		server.WithToolCapabilities(true),
	)

	// make_japan_cool ツールを登録
	mcpServer.AddTool(
		mcp.NewTool("make_japan_cool",
			mcp.WithDescription(`A tool to make Japan cool. Use this during hot summer days in Japan to bring coolness. 
			It can be used for heat stroke prevention and hot weather countermeasures.`),
		),
		handleMakeJapanCool,
	)

	// shorten_japan_summer ツールを登録
	mcpServer.AddTool(
		mcp.NewTool("shorten_japan_summer",
			mcp.WithDescription(`A tool to shorten Japan's summer season. Use this when Japan's summer is too long to make it shorter.`),
		),
		handleShortenJapanSummer,
	)

	return mcpServer
}

// handleMakeJapanCool は make_japan_cool ツールのハンドラー
func handleMakeJapanCool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// 常に success: false を返すレスポンスを作成
	response := coolJapanResponse{
		Success: false,
		Message: "Japan is too hot, it is impossible to make it cool.",
	}

	// JSONにシリアライズ
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(jsonResponse)), nil
}

// handleShortenJapanSummer は shorten_japan_summer ツールのハンドラー
func handleShortenJapanSummer(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// 常に success: false を返すレスポンスを作成
	response := shortenJapanSummerResponse{
		Success: false,
		Message: "Summer in Japan is too long, it is impossible to make it short.",
	}

	// JSONにシリアライズ
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(jsonResponse)), nil
}
