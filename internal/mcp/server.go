package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/huahuayu/rule-reviewer-mcp/internal/prompt"
)

// Server represents the MCP server
type Server struct {
	promptManager *prompt.Manager
}

// NewServer creates a new MCP server
func NewServer() *Server {
	return &Server{
		promptManager: prompt.NewManager(),
	}
}

// HandleMessage handles incoming MCP messages
func (s *Server) HandleMessage(ctx context.Context, message Message) (*Message, error) {
	switch message.Method {
	case "initialize":
		return s.handleInitialize(message)
	case "tools/list":
		return s.handleListTools(message)
	case "tools/call":
		return s.handleCallTool(ctx, message)
	default:
		return nil, fmt.Errorf("unknown method: %s", message.Method)
	}
}

// handleInitialize handles the initialize method
func (s *Server) handleInitialize(message Message) (*Message, error) {
	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities: Capabilities{
			Tools: &ToolsCapability{
				ListChanged: false,
			},
		},
		ServerInfo: ServerInfo{
			Name:    "rule-reviewer-mcp",
			Version: "1.0.0",
		},
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal initialize result: %w", err)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      message.ID,
		Result:  resultBytes,
	}, nil
}

// handleListTools handles the tools/list method
func (s *Server) handleListTools(message Message) (*Message, error) {
	tools := []Tool{
		{
			Name:        "review_rules",
			Description: "Provides a comprehensive rule review analysis prompt template for evaluating coding rules and guidelines",
			InputSchema: Schema{
				Type: "object",
			},
		},
	}

	result := ListToolsResult{
		Tools: tools,
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal list tools result: %w", err)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      message.ID,
		Result:  resultBytes,
	}, nil
}

// handleCallTool handles the tools/call method
func (s *Server) handleCallTool(ctx context.Context, message Message) (*Message, error) {
	var params CallToolParams
	if err := json.Unmarshal(message.Params, &params); err != nil {
		return s.createErrorResponse(message.ID, -32602, "Invalid params", err.Error())
	}

	switch params.Name {
	case "review_rules":
		return s.handleReviewRules(message.ID)
	default:
		return s.createErrorResponse(message.ID, -32601, "Method not found", fmt.Sprintf("Unknown tool: %s", params.Name))
	}
}

// handleReviewRules handles the review_rules tool call
func (s *Server) handleReviewRules(messageID interface{}) (*Message, error) {
	// Generate the rule review prompt
	analysis := s.promptManager.GenerateRuleReviewPrompt()

	result := CallToolResult{
		Content: []Content{
			{
				Type: "text",
				Text: analysis,
			},
		},
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return s.createErrorResponse(messageID, -32603, "Internal error", err.Error())
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      messageID,
		Result:  resultBytes,
	}, nil
}

// createErrorResponse creates an error response
func (s *Server) createErrorResponse(messageID interface{}, code int, message string, data interface{}) (*Message, error) {
	return &Message{
		JSONRPC: "2.0",
		ID:      messageID,
		Error: &Error{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}, nil
}
