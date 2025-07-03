package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/huahuayu/rule-reviewer-mcp/internal/mcp"
)

func main() {
	// Create MCP server
	server := mcp.NewServer()
	
	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Setup stdin/stdout communication
	scanner := bufio.NewScanner(os.Stdin)
	
	// Process MCP messages
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		// Parse MCP message
		var message mcp.Message
		if err := json.Unmarshal([]byte(line), &message); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}
		
		// Handle message
		response, err := server.HandleMessage(ctx, message)
		if err != nil {
			log.Printf("Error handling message: %v", err)
			continue
		}
		
		// Send response
		if response != nil {
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}
			fmt.Println(string(responseBytes))
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from stdin: %v", err)
	}
} 