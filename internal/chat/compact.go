package chat

import (
	"context"
	"fmt"
	"strings"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

const (
	compactThresholdFrac = 0.80
	compactTailFrac      = 0.20
	compactDefaultWindow = 32000
	compactSummaryMin    = 40
)

func msgTokens(m provider.Message) int {
	n := 4
	if s, ok := m.Content.(string); ok {
		n += len(s) / 4
	}
	for _, tc := range m.ToolCalls {
		n += (len(tc.Function.Name) + len(tc.Function.Arguments)) / 4
	}
	return n
}

func estimateTokens(msgs []provider.Message) int {
	total := 0
	for _, m := range msgs {
		total += msgTokens(m)
	}
	return total
}

func compactBounds(msgs []provider.Message, tailBudget int) (head, tailStart int) {
	for head < len(msgs) && msgs[head].Role == "system" {
		head++
	}
	tailStart = len(msgs)
	acc := 0
	for i := len(msgs) - 1; i >= head; i-- {
		acc += msgTokens(msgs[i])
		tailStart = i
		if acc >= tailBudget {
			break
		}
	}
	for tailStart > head && msgs[tailStart].Role != "user" && msgs[tailStart].Role != "assistant" {
		tailStart--
	}
	return head, tailStart
}

func renderTranscript(msgs []provider.Message) string {
	var b strings.Builder
	for _, m := range msgs {
		switch m.Role {
		case "user":
			if s, ok := m.Content.(string); ok {
				fmt.Fprintf(&b, "User: %s\n", s)
			}
		case "assistant":
			if s, ok := m.Content.(string); ok && s != "" {
				fmt.Fprintf(&b, "Assistant: %s\n", s)
			}
			for _, tc := range m.ToolCalls {
				fmt.Fprintf(&b, "Assistant -> tool %s(%s)\n", tc.Function.Name, tc.Function.Arguments)
			}
		case "tool":
			if s, ok := m.Content.(string); ok {
				t := s
				if r := []rune(t); len(r) > 200 {
					t = string(r[:200]) + "..."
				}
				fmt.Fprintf(&b, "Tool result: %s\n", t)
			}
		}
	}
	s := b.String()
	const maxChars = 100000
	if len(s) > maxChars {
		s = "[...debut tronque...]\n" + s[len(s)-maxChars:]
	}
	return s
}

func compactSummaryPrompt() string {
	return `You are a context compactor. Summarize the transcript of older conversation turns densely and faithfully, keeping ONLY:
- The user's current request, goals and constraints
- Concrete findings already gathered (facts, URLs, file paths, values)
- Decisions made and established facts
- State of progress: what is done, what is missing, next step
Rules: no preamble/conclusion, no verbatim quotes, use short bullet points. Be concise while keeping every fact. Write in the SAME language as the conversation.`
}

func summarize(ctx context.Context, p provider.Provider, model, transcript string, budget int) (string, error) {
	resp, err := p.Stream(ctx, provider.Request{
		Model: model,
		Messages: []provider.Message{
			{Role: "system", Content: compactSummaryPrompt()},
			{Role: "user", Content: transcript},
		},
		Temperature: 0.2,
	}, func(provider.Event) bool { return true })
	if err != nil {
		return "", err
	}
	c := strings.TrimSpace(resp.Content)
	if r := []rune(c); len(r) > budget*4 {
		c = strings.TrimSpace(string(r[:budget*4])) + " [...]"
	}
	return c, nil
}

func (e *Engine) maybeCompact(ctx context.Context, c *Conversation, epoch int, members []alias.ResolvedMember) {
	c.mu.Lock()
	msgs := append([]provider.Message(nil), c.Messages...)
	c.mu.Unlock()

	used := estimateTokens(msgs)
	if used < int(float64(compactDefaultWindow)*compactThresholdFrac) {
		return
	}

	tailBudget := int(float64(used) * compactTailFrac)
	head, tailStart := compactBounds(msgs, tailBudget)
	if tailStart <= head {
		return
	}

	torso := msgs[head:tailStart]
	transcript := renderTranscript(torso)
	budget := compactDefaultWindow / 25
	if budget < 400 {
		budget = 400
	}
	if budget > 1600 {
		budget = 1600
	}

	summary := ""
	for _, m := range members {
		p, ok := e.reg.Get(m.Provider)
		if !ok {
			continue
		}
		s, err := summarize(ctx, p, m.Model, transcript, budget)
		if err != nil || len([]rune(strings.TrimSpace(s))) < compactSummaryMin {
			continue
		}
		summary = s
		break
	}
	if summary == "" {
		return
	}

	var pending []provider.Message
	for i := len(torso) - 1; i >= 0; i-- {
		if torso[i].Role == "user" {
			pending = []provider.Message{torso[i]}
			break
		}
	}

	newMsgs := make([]provider.Message, 0, head+2+len(pending)+len(msgs)-tailStart)
	newMsgs = append(newMsgs, msgs[:head]...)
	newMsgs = append(newMsgs,
		provider.Message{Role: "user", Content: "[CONTEXT COMPACTED] Earlier turns were summarized. Summary:\n\n" + summary},
		provider.Message{Role: "assistant", Content: "Understood. Resuming from where I left off."},
	)
	newMsgs = append(newMsgs, pending...)
	newMsgs = append(newMsgs, msgs[tailStart:]...)

	if estimateTokens(newMsgs) > used*4/5 {
		return
	}

	c.mu.Lock()
	if c.epoch == epoch {
		c.Messages = newMsgs
	}
	c.mu.Unlock()
}
