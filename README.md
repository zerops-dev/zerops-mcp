# Zerops MCP server

## How to run

1. Install `mcphost`

```bash
go install github.com/mark3labs/mcphost@latest
```

2. Create `mcp.json`

```json
{
  "mcpServers": {
    "container": {
      "command": "go",
      "args": ["run", "./cmd/main.go", "run", "--zerops-mcp.log.output", "discard"]
    }
  }
}
```

3. export APIKEY and run

```bash
export ANTHROPIC_API_KEY="API-KEY"
mcphost  --config mcp.json 
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.