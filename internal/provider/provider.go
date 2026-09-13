package provider

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Message struct {
	Role       string     `json:"role"`
	Content    any        `json:"content,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

type ToolCall struct {
	Index    int    `json:"index,omitempty"`
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Function Func   `json:"function"`
}

type Func struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Tool struct {
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
}

type ToolFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Request struct {
	Model       string
	Messages    []Message
	Tools       []Tool
	Temperature float64
}

type Response struct {
	Content      string
	Reasoning    string
	ToolCalls    []ToolCall
	Usage        Usage
	FinishReason string
}

type Event struct {
	Content   string
	Reasoning string
	Usage     *Usage
	Err       error
}

type Provider interface {
	ID() string
	Stream(ctx context.Context, req Request, emit func(Event) bool) (Response, error)
}

type Endpoint struct {
	Provider string
	BaseURL  string
	Path     string
	APIKey   string
	Headers  map[string]string
}

type OpenAICompat struct {
	endpoint Endpoint
	client   *http.Client
}

func NewOpenAICompat(e Endpoint, client *http.Client) *OpenAICompat {
	if client == nil {
		client = &http.Client{Timeout: 0}
	}
	if e.Path == "" {
		e.Path = "/chat/completions"
	}
	return &OpenAICompat{endpoint: e, client: client}
}

func (p *OpenAICompat) ID() string { return p.endpoint.Provider }

type streamChunk struct {
	Choices []struct {
		Delta struct {
			Content          string `json:"content"`
			ReasoningContent string `json:"reasoning_content"`
			Reasoning        string `json:"reasoning"`
			ToolCalls        []struct {
				Index    int    `json:"index"`
				ID       string `json:"id"`
				Type     string `json:"type"`
				Function Func   `json:"function"`
			} `json:"tool_calls"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage *Usage `json:"usage"`
}

var ErrStreamCut = errors.New("flux de reponse coupe")

type HTTPError struct {
	Provider string
	Status   int
	Body     string
}

func (e *HTTPError) Error() string {
	msg := e.Body
	if msg == "" {
		msg = http.StatusText(e.Status)
	}
	return fmt.Sprintf("%s a renvoye %d: %s", e.Provider, e.Status, msg)
}

func (p *OpenAICompat) Stream(ctx context.Context, req Request, emit func(Event) bool) (Response, error) {
	var resp Response
	if req.Temperature == 0 {
		req.Temperature = 0.7
	}
	payload := map[string]any{
		"model":       req.Model,
		"messages":    req.Messages,
		"stream":      true,
		"temperature": req.Temperature,
		"stream_options": map[string]any{
			"include_usage": true,
		},
	}
	if len(req.Tools) > 0 {
		payload["tools"] = req.Tools
		payload["parallel_tool_calls"] = false
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return resp, fmt.Errorf("encodage requete: %w", err)
	}

	url := strings.TrimRight(p.endpoint.BaseURL, "/") + p.endpoint.Path
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return resp, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")
	if p.endpoint.APIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.endpoint.APIKey)
	}
	for k, v := range p.endpoint.Headers {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := p.client.Do(httpReq)
	if err != nil {
		if ctx.Err() != nil {
			return resp, ctx.Err()
		}
		return resp, fmt.Errorf("connexion a %s impossible: %w", p.endpoint.Provider, err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(io.LimitReader(httpResp.Body, 2000))
		msg := strings.TrimSpace(string(raw))
		if msg == "" {
			msg = httpResp.Status
		}
		return resp, &HTTPError{Provider: p.endpoint.Provider, Status: httpResp.StatusCode, Body: msg}
	}

	toolCalls := map[int]*ToolCall{}
	sc := bufio.NewScanner(httpResp.Body)
	sc.Buffer(make([]byte, 0, 64*1024), 8<<20)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if data == "" || data == "[DONE]" {
			continue
		}
		var chunk streamChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			continue
		}
		if chunk.Usage != nil {
			resp.Usage = *chunk.Usage
			u := *chunk.Usage
			if !emit(Event{Usage: &u}) {
				return resp, nil
			}
		}
		if len(chunk.Choices) == 0 {
			continue
		}
		ch := chunk.Choices[0]
		if ch.FinishReason != "" {
			resp.FinishReason = ch.FinishReason
		}
		if len(ch.Delta.ToolCalls) > 0 {
			for i, tc := range ch.Delta.ToolCalls {
				idx := tc.Index
				if tc.Index == 0 && i > 0 {
					idx = i
				}
				cur, ok := toolCalls[idx]
				if !ok {
					cur = &ToolCall{Index: idx, Type: "function"}
					toolCalls[idx] = cur
				}
				if tc.ID != "" {
					cur.ID = tc.ID
				}
				if tc.Type != "" {
					cur.Type = tc.Type
				}
				if tc.Function.Name != "" {
					cur.Function.Name = tc.Function.Name
				}
				cur.Function.Arguments += tc.Function.Arguments
			}
		}
		reasoning := ch.Delta.ReasoningContent
		if reasoning == "" {
			reasoning = ch.Delta.Reasoning
		}
		if reasoning != "" {
			resp.Reasoning += reasoning
			if !emit(Event{Reasoning: reasoning}) {
				return resp, nil
			}
		}
		if ch.Delta.Content != "" {
			resp.Content += ch.Delta.Content
			if !emit(Event{Content: ch.Delta.Content}) {
				return resp, nil
			}
		}
	}

	if err := sc.Err(); err != nil {
		if ctx.Err() != nil {
			return resp, ctx.Err()
		}
		return resp, fmt.Errorf("%w: %v", ErrStreamCut, err)
	}
	if ctx.Err() != nil {
		return resp, ctx.Err()
	}

	idxs := make([]int, 0, len(toolCalls))
	for k := range toolCalls {
		idxs = append(idxs, k)
	}
	sortInts(idxs)
	for _, k := range idxs {
		tc := *toolCalls[k]
		if tc.ID == "" {
			tc.ID = fmt.Sprintf("call_%d", k)
		}
		if !json.Valid([]byte(tc.Function.Arguments)) {
			tc.Function.Arguments = "{}"
		}
		resp.ToolCalls = append(resp.ToolCalls, tc)
	}
	return resp, nil
}

func sortInts(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 30 * time.Second,
		},
	}
}
