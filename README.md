# Rule Reviewer MCP Server

An open source MCP (Model Context Protocol) server that provides intelligent rule review capabilities for development projects. It integrates with Cursor and other AI coding assistants to analyze and review both user-defined rules and project-specific rules.

## Features

- **🔍 Comprehensive Rule Analysis**: Analyzes rule quality, conflicts, and gaps
- **🔄 Duplicate Detection**: Identifies redundant rules, including those that duplicate Cursor's built-in capabilities
- **📊 Conflict Resolution**: Detects and provides recommendations for conflicting rules
- **📝 Multi-format Support**: Handles rules in Markdown, JSON, YAML, and plain text formats
- **🚀 High Performance**: Efficient processing of large rule sets with concurrent analysis
- **🔌 MCP Protocol Compliance**: Full MCP v1.0 protocol implementation for seamless integration

## Installation

### Quick Install (Mac)

```bash
curl -L https://github.com/huahuayu/rule-reviewer-mcp/releases/latest/download/rule-reviewer-mcp -o /tmp/rule-reviewer-mcp && chmod +x /tmp/rule-reviewer-mcp && sudo mv /tmp/rule-reviewer-mcp /usr/local/bin/rule-reviewer-mcp
```

### Prerequisites

- Go 1.24.3 or higher
- Git

### Build from Source

```bash
git clone https://github.com/huahuayu/rule-reviewer-mcp.git
cd rule-reviewer-mcp
go mod download
go build -o rule-reviewer-mcp
```

### Cross-platform Builds

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o rule-reviewer-mcp-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o rule-reviewer-mcp-macos

# Windows
GOOS=windows GOARCH=amd64 go build -o rule-reviewer-mcp.exe
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
      "command": "/path/to/rule-reviewer-mcp",
      "args": []
    }
  }
}
```

### Using the Rule Review Tool

Once configured, you can use the `review_rules` tool in Cursor:

```javascript
// Example tool call
{
  "name": "review_rules",
  "arguments": {
    "user_rules": "# My Rules\n- Use camelCase for variables\n- Always handle errors",
    "project_rules": "# Project Rules\n- Follow Go conventions\n- Use gofmt",
    "project_context": "Go web API project using fiber framework"
  }
}
```

## Rule Formats

### Markdown Format

```markdown
# Code Style Rules

## Variable Naming
- Use camelCase for variables and functions
- Use PascalCase for types and structs

## Error Handling
- Always handle errors explicitly
- Use custom error types for business logic
```

### JSON Format

```json
[
  {
    "title": "Variable Naming",
    "content": "Use camelCase for variables and functions",
    "category": "code_style",
    "priority": "high"
  },
  {
    "title": "Error Handling",
    "content": "Always handle errors explicitly",
    "category": "error_handling",
    "priority": "high"
  }
]
```

### YAML Format

```yaml
variable_naming:
  category: code_style
  priority: high
  content: Use camelCase for variables and functions

error_handling:
  category: error_handling
  priority: high
  content: Always handle errors explicitly
```

### Plain Text Format

```text
Use camelCase for variables and functions
Always handle errors explicitly
Keep functions under 50 lines
Use descriptive function names
```

## Analysis Output

The rule review analysis includes:

### Executive Summary
A brief overview of the rule set health and key findings.

### Critical Issues
High-priority issues that require immediate attention.

### Conflicts Detected
Detailed list of conflicting rules with resolution recommendations.

### Duplicate Rules Found
- Rules that duplicate Cursor's built-in capabilities
- Inter-rule duplicates with consolidation recommendations

### Missing Rules (Gap Analysis)
Suggested new rules organized by category:
- Code style and formatting
- Architecture and design patterns
- Testing requirements
- Documentation standards
- Security practices
- Performance considerations
- Error handling
- Code review process
- Deployment and CI/CD

### Improvement Opportunities
Medium and low priority enhancements.

### Recommended Actions
Prioritized list of next steps.

## API Reference

### MCP Tool: `review_rules`

Analyzes and reviews user rules and project rules for conflicts, duplicates, and gaps.

**Parameters:**
- `user_rules` (string, required): User-defined rules in any supported format
- `project_rules` (string, required): Project-specific rules in any supported format
- `project_context` (string, optional): Additional context about the project

**Returns:**
- Comprehensive analysis report in structured markdown format

## Development

### Project Structure

```
rule-reviewer-mcp/
├── main.go                 # Entry point and MCP server initialization
├── internal/
│   ├── mcp/               # MCP protocol implementation
│   ├── analyzer/          # Rule analysis logic
│   └── prompt/            # Prompt template management
├── pkg/
│   └── rules/             # Rule parsing and validation
├── examples/              # Example rule files and usage
├── docs/                  # Documentation and guides
├── LICENSE               # Open source license
├── README.md             # Project documentation
└── go.mod                # Go module definition
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...

# Run benchmarks
go test -bench=. ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code
go vet ./...

# Static analysis
golangci-lint run
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Write clear, descriptive commit messages
- Include tests for new features
- Update documentation as needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: Report bugs and request features on [GitHub Issues](https://github.com/huahuayu/rule-reviewer-mcp/issues)
- **Discussions**: Join discussions on [GitHub Discussions](https://github.com/huahuayu/rule-reviewer-mcp/discussions)
- **Documentation**: Check the [docs](docs/) directory for additional guides

## Roadmap

- [ ] **Rule Templates**: Pre-built rule sets for common project types
- [ ] **Auto-fixing**: Automated rule conflict resolution
- [ ] **Integration**: Direct integration with popular linters and formatters
- [ ] **Metrics**: Rule compliance tracking and reporting
- [ ] **Web Interface**: Optional web-based UI for rule management
- [ ] **Plugin System**: Extensible architecture for custom rule processors

## Acknowledgments

- [Model Context Protocol](https://modelcontextprotocol.io/) for the MCP specification
- [Cursor](https://cursor.sh/) for AI-powered code editing
- The Go community for excellent tools and libraries

---

**Built with ❤️ in Go** | **Open Source** | **MCP Protocol Compliant** 