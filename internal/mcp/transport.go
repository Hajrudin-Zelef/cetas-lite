package mcp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"time"
)

const maxMessageSize = 8 << 20

type transport interface {
	Call(ctx context.Context, method string, params any) (json.RawMessage, error)
	Notify(ctx context.Context, method string, params any) error
	Close() error
}

func newTransport(cfg ServerConfig) (transport, error) {
	kind, err := cfg.transport()
	if err != nil {
		return nil, err
	}
	switch kind {
	case "stdio":
		return newStdioTransport(cfg)
	case "http":
		return newHTTPTransport(cfg), nil
	}
	return nil, fmt.Errorf("transport non supporte: %s", kind)
}

// ── stdio ────────────────────────────────────────────────────────────

// limitedSyncBuffer capture le stderr d'un serveur MCP en le bornant : un
// serveur bavard ne peut pas faire gonfler la memoire sur une longue session.
type limitedSyncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
	max int
}

func (b *limitedSyncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.buf.Len() >= b.max {
		return len(p), nil // tronque : on garde le debut
	}
	if len(p) > b.max-b.buf.Len() {
		p = p[:b.max-b.buf.Len()]
	}
	return b.buf.Write(p)
}

func (b *limitedSyncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

type stdioTransport struct {
	cmd     *exec.Cmd
	stdin   io.WriteCloser
	job     uintptr // Job Object Windows (0 sous Unix)
	mu      sync.Mutex
	pending map[int64]chan rpcMessage
	nextID  int64
	closed  bool
	err     error
}

func newStdioTransport(cfg ServerConfig) (*stdioTransport, error) {
	cmd := exec.Command(cfg.Command, cfg.Args...)
	env := os.Environ()
	keys := make([]string, 0, len(cfg.Env))
	for k := range cfg.Env {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		env = append(env, k+"="+cfg.Env[k])
	}
	cmd.Env = env
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr := &limitedSyncBuffer{max: 64 << 10}
	cmd.Stderr = stderr
	isolatePreStart(cmd)
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("demarrage %s: %w", cfg.Command, err)
	}
	t := &stdioTransport{cmd: cmd, stdin: stdin, pending: map[int64]chan rpcMessage{}, job: isolatePostStart(cmd)}
	go t.readLoop(stdout, stderr)
	return t, nil
}

func (t *stdioTransport) readLoop(stdout io.Reader, stderr *limitedSyncBuffer) {
	sc := bufio.NewScanner(stdout)
	sc.Buffer(make([]byte, 0, 64*1024), maxMessageSize)
	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 {
			continue
		}
		var msg rpcMessage
		if err := json.Unmarshal(line, &msg); err != nil {
			continue
		}
		if msg.Method != "" && msg.ID != nil {
			t.reply(*msg.ID, &rpcError{Code: -32601, Message: "method not found"})
			continue
		}
		if msg.ID != nil {
			t.deliver(*msg.ID, msg)
		}
	}
	t.finish(fmt.Errorf("serveur mcp termine: %s", strings.TrimSpace(stderr.String())))
}

func (t *stdioTransport) deliver(id int64, msg rpcMessage) {
	t.mu.Lock()
	ch, ok := t.pending[id]
	delete(t.pending, id)
	t.mu.Unlock()
	if ok {
		ch <- msg
	}
}

func (t *stdioTransport) finish(err error) {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return
	}
	t.closed = true
	if err != nil {
		t.err = err
	}
	for id, ch := range t.pending {
		delete(t.pending, id)
		close(ch)
	}
	t.mu.Unlock()
}

func (t *stdioTransport) Call(ctx context.Context, method string, params any) (json.RawMessage, error) {
	raw, err := marshalParams(params)
	if err != nil {
		return nil, err
	}
	t.mu.Lock()
	if t.closed {
		err := t.err
		t.mu.Unlock()
		return nil, err
	}
	t.nextID++
	id := t.nextID
	ch := make(chan rpcMessage, 1)
	t.pending[id] = ch
	data, err := json.Marshal(rpcMessage{JSONRPC: rpcVersion, ID: &id, Method: method, Params: raw})
	if err == nil {
		_, err = t.stdin.Write(append(data, '\n'))
	}
	if err != nil {
		delete(t.pending, id)
		t.mu.Unlock()
		return nil, err
	}
	t.mu.Unlock()

	select {
	case <-ctx.Done():
		// Nettoyage : sans ca, chaque appel annule laisse un canal orphelin
		// dans pending (fuite memoire sur les longues sessions).
		t.mu.Lock()
		delete(t.pending, id)
		t.mu.Unlock()
		return nil, ctx.Err()
	case m, ok := <-ch:
		if !ok {
			t.mu.Lock()
			err := t.err
			t.mu.Unlock()
			if err == nil {
				err = fmt.Errorf("serveur mcp indisponible")
			}
			return nil, err
		}
		if m.Error != nil {
			return nil, m.Error
		}
		return m.Result, nil
	}
}

func (t *stdioTransport) Notify(ctx context.Context, method string, params any) error {
	raw, err := marshalParams(params)
	if err != nil {
		return err
	}
	data, err := json.Marshal(rpcMessage{JSONRPC: rpcVersion, Method: method, Params: raw})
	if err != nil {
		return err
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return t.err
	}
	_, err = t.stdin.Write(append(data, '\n'))
	return err
}

func (t *stdioTransport) reply(id int64, rerr *rpcError) {
	data, err := json.Marshal(rpcMessage{JSONRPC: rpcVersion, ID: &id, Error: rerr})
	if err != nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return
	}
	_, _ = t.stdin.Write(append(data, '\n'))
}

func (t *stdioTransport) Close() error {
	t.finish(nil)
	killTree(t.cmd, t.job)
	return t.cmd.Wait()
}

// ── http (streamable) ────────────────────────────────────────────────

type httpTransport struct {
	url       string
	headers   map[string]string
	client    *http.Client
	mu        sync.Mutex
	nextID    int64
	sessionID string
}

func newHTTPTransport(cfg ServerConfig) *httpTransport {
	return &httpTransport{
		url:     cfg.URL,
		headers: cfg.Headers,
		client:  &http.Client{Timeout: 30 * time.Second},
	}
}

func (t *httpTransport) Call(ctx context.Context, method string, params any) (json.RawMessage, error) {
	raw, err := marshalParams(params)
	if err != nil {
		return nil, err
	}
	t.mu.Lock()
	t.nextID++
	id := t.nextID
	sid := t.sessionID
	t.mu.Unlock()

	data, err := json.Marshal(rpcMessage{JSONRPC: rpcVersion, ID: &id, Method: method, Params: raw})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	for k, v := range t.headers {
		req.Header.Set(k, v)
	}
	if sid != "" {
		req.Header.Set("Mcp-Session-Id", sid)
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if v := resp.Header.Get("Mcp-Session-Id"); v != "" {
		t.mu.Lock()
		t.sessionID = v
		t.mu.Unlock()
	}
	if resp.StatusCode/100 != 2 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
		return readSSEResult(resp.Body, id)
	}
	var m rpcMessage
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxMessageSize)).Decode(&m); err != nil {
		return nil, err
	}
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Result, nil
}

func (t *httpTransport) Notify(ctx context.Context, method string, params any) error {
	raw, err := marshalParams(params)
	if err != nil {
		return err
	}
	data, err := json.Marshal(rpcMessage{JSONRPC: rpcVersion, Method: method, Params: raw})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	for k, v := range t.headers {
		req.Header.Set(k, v)
	}
	t.mu.Lock()
	sid := t.sessionID
	t.mu.Unlock()
	if sid != "" {
		req.Header.Set("Mcp-Session-Id", sid)
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
	return nil
}

func (t *httpTransport) Close() error {
	t.mu.Lock()
	sid := t.sessionID
	t.mu.Unlock()
	if sid == "" {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, t.url, nil)
	if err != nil {
		return nil
	}
	for k, v := range t.headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Mcp-Session-Id", sid)
	resp, err := t.client.Do(req)
	if err == nil {
		resp.Body.Close()
	}
	return nil
}

func readSSEResult(r io.Reader, id int64) (json.RawMessage, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), maxMessageSize)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if payload == "" {
			continue
		}
		var m rpcMessage
		if json.Unmarshal([]byte(payload), &m) != nil {
			continue
		}
		if m.ID != nil && *m.ID != id {
			continue
		}
		if m.Error != nil {
			return nil, m.Error
		}
		if m.Result != nil {
			return m.Result, nil
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("reponse http sans resultat")
}
