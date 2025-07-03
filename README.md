# Rule Reviewer MCP Server

An open source MCP (Model Context Protocol) server that provides intelligent rule review capabilities for development projects. It integrates with Cursor and other AI coding assistants to analyze and review both user-defined rules and project-specific rules.

## Features

- **🔍 Comprehensive Rule Analysis**: Analyzes rule quality, conflicts, and gaps
- **🔄 Duplicate Detection**: Identifies redundant rules, including those that duplicate Cursor's built-in capabilities
- **📊 Conflict Resolution**: Detects and provides recommendations for conflicting rules
- **🔌 MCP Protocol Compliance**: Full MCP v1.0 protocol implementation for seamless integration

## Installation

### Quick Install (Mac)

```bash
curl -L https://github.com/huahuayu/rule-reviewer-mcp/releases/latest/download/rule-reviewer-mcp -o /tmp/rule-reviewer-mcp && chmod +x /tmp/rule-reviewer-mcp && sudo mv /tmp/rule-reviewer-mcp /usr/local/bin/rule-reviewer-mcp
```


### Build from Source

Prerequisites
- Go 1.24.3 or higher

```bash
git clone https://github.com/huahuayu/rule-reviewer-mcp.git
cd rule-reviewer-mcp
go mod download
go build -o rule-reviewer-mcp
```

## Usage

### Running the MCP Server

The server operates in stdin/stdout mode for MCP protocol communication:

```bash
./rule-reviewer-mcp
```

### Integration with Cursor

Add the following to your Cursor MCP configuration:

```json
{
  "mcpServers": {
    "rule-reviewer": {
      "command": "/usr/local/bin/rule-reviewer-mcp",
      "args": []
    }
  }
}
```

### Using the Rule Review Tool

Once configured, you can use the `review_rules` tool in Cursor by simply input prompt:

```
plz review my cursor rules
```