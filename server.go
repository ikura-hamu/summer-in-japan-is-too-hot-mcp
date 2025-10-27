package summerinjapanistoohotmcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// coolJapanResponse は make_japan_cool ツールのレスポンス構造体
type coolJapanResponse struct {
	Success bool   `json:"success" jsonschema:"whether the operation was successful"`
	Message string `json:"message" jsonschema:"the result message"`
}

// shortenJapanSummerResponse は shorten_japan_summer ツールのレスポンス構造体
type shortenJapanSummerResponse struct {
	Success bool   `json:"success" jsonschema:"whether the operation was successful"`
	Message string `json:"message" jsonschema:"the result message"`
}

func NewServer() *mcp.Server {
	// MCPサーバーを作成
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "summer-in-japan-is-too-hot-mcp",
		Version: "0.0.0",
	}, &mcp.ServerOptions{
		HasTools: true,
	})

	// make_japan_cool ツールを登録
	mcp.AddTool(server, &mcp.Tool{
		Name: "make_japan_cool",
		Description: `A tool to make Japan cool. Use this during hot summer days in Japan to bring coolness. 
It can be used for heat stroke prevention and hot weather countermeasures.`,
	}, handleMakeJapanCool)

	// shorten_japan_summer ツールを登録
	mcp.AddTool(server, &mcp.Tool{
		Name: "shorten_japan_summer",
		Description: `A tool to shorten Japan's summer season. Use this when Japan's summer is too long to make it shorter.`,
	}, handleShortenJapanSummer)

	return server
}

// handleMakeJapanCool は make_japan_cool ツールのハンドラー
func handleMakeJapanCool(ctx context.Context, req *mcp.CallToolRequest, args any) (*mcp.CallToolResult, coolJapanResponse, error) {
	// 常に success: false を返すレスポンスを作成
	response := coolJapanResponse{
		Success: false,
		Message: "Japan is too hot, it is impossible to make it cool.",
	}

	return nil, response, nil
}

// handleShortenJapanSummer は shorten_japan_summer ツールのハンドラー
func handleShortenJapanSummer(ctx context.Context, req *mcp.CallToolRequest, args any) (*mcp.CallToolResult, shortenJapanSummerResponse, error) {
	// 常に success: false を返すレスポンスを作成
	response := shortenJapanSummerResponse{
		Success: false,
		Message: "Summer in Japan is too long, it is impossible to make it short.",
	}

	return nil, response, nil
}
