package mcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("CETAS_MCP_HELPER") == "1" {
		runFakeServer()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

// CETAS_MCP_HANG=1 : le serveur lit mais ne repond jamais (simule un serveur
// bloque — pour tester timeouts et nettoyages).
func runFakeServer() {
	hang := os.Getenv("CETAS_MCP_HANG") == "1"
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
	w := bufio.NewWriter(os.Stdout)
	for sc.Scan() {
		if hang {
			continue // lit, ne repond jamais
		}
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var msg rpcMessage
		if json.Unmarshal(line, &msg) != nil {
			continue
		}
		if msg.ID == nil {
			continue
		}
		resp := rpcMessage{JSONRPC: "2.0", ID: msg.ID}
		switch msg.Method {
		case "initialize":
			resp.Result = json.RawMessage(`{"protocolVersion":"2024-11-05","capabilities":{},"serverInfo":{"name":"fake","version":"0"}}`)
		case "tools/list":
			resp.Result = json.RawMessage(`{"tools":[{"name":"echo","description":"Echo text","inputSchema":{"type":"object","properties":{"text":{"type":"string"}},"required":["text"]}},{"name":"add","description":"Add two ints","inputSchema":{"type":"object","properties":{"a":{"type":"integer"},"b":{"type":"integer"}}}}]}`)
		case "tools/call":
			var p struct {
				Name      string         `json:"name"`
				Arguments map[string]any `json:"arguments"`
			}
			_ = json.Unmarshal(msg.Params, &p)
			switch p.Name {
			case "echo":
				text, _ := p.Arguments["text"].(string)
				resp.Result = json.RawMessage(fmt.Sprintf(`{"content":[{"type":"text","text":%s}],"isError":false}`, quoteJSON(text)))
			case "add":
				a, _ := p.Arguments["a"].(float64)
				b, _ := p.Arguments["b"].(float64)
				resp.Result = json.RawMessage(fmt.Sprintf(`{"content":[{"type":"text","text":"%d"}],"isError":false}`, int(a+b)))
			default:
				resp.Result = json.RawMessage(`{"content":[{"type":"text","text":"outil inconnu"}],"isError":true}`)
			}
		default:
			resp.Error = &rpcError{Code: -32601, Message: "method not found"}
		}
		data, _ := json.Marshal(resp)
		_, _ = w.Write(append(data, '\n'))
		_ = w.Flush()
	}
}

func quoteJSON(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}
