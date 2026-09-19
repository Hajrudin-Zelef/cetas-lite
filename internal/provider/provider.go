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
	// ReasoningContent : raisonnement du modele renvoye tel quel dans la
	// suite d'une conversation en mode thinking. DeepSeek l'exige sur
	// TOUS les messages assistant des qu'un seul en porte un (pas
	// seulement ceux avec tool_calls) : sinon le tour suivant echoue
	// en HTTP 400. omitempty : absent quand aucun reasoning n'est en
	// jeu ; withReasoningContentForced force le champ (meme vide) sur
	// le wire des que du reasoning circule dans la conversation.
	ReasoningContent string `json:"reasoning_content,omitempty"`
}

type ToolCall struct {
	Index    int    `json:"index,omitempty"`
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Function Func   `json:"function"`
	// invalid marque un appel irrecevable (nom vide ou arguments JSON
	// irreparables). Non serialise : usage interne uniquement.
	invalid bool
}

// InvalidCall indique si cet appel d'outil doit etre rejete sans execution.
func (t ToolCall) InvalidCall() bool { return t.invalid }

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
	Model           string
	Messages        []Message
	Tools           []Tool
	Temperature     float64
	EnableReasoning bool
	ReasoningEffort string
	// MaxTokens limite le nombre de tokens generes par reponse (0 = defaut du provider).
	MaxTokens int
	// Extra : parametres supplementaires fusionnes dans le payload
	// (ex. plugins de recherche web natifs d'OpenRouter).
	Extra map[string]any
}

// WebAnnotation : citation web renvoyee par une recherche native du provider
// (ex. plugin web d'OpenRouter).
type WebAnnotation struct {
	URL   string
	Title string
}

type Response struct {
	Content      string
	Reasoning    string
	ToolCalls    []ToolCall
	Usage        Usage
	FinishReason string
	// Annotations : citations web d'une recherche native du provider.
	Annotations []WebAnnotation
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

// apiModel retire le suffixe interne de desambiguisation avant l'appel HTTP.
// Le catalogue CETAS nomme distinctement le meme modele selon la passerelle
// OpenCode (-zen pour opencode, -go pour opencode-go) pour que l'UI et les
// capacites (provider/model) ne les confondent pas ; les API OpenCode
// attendent le nom brut (ex. glm-5-zen -> glm-5, hy3-go -> hy3).
func apiModel(providerID, model string) string {
	switch providerID {
	case "opencode":
		return strings.TrimSuffix(model, "-zen")
	case "opencode-go":
		return strings.TrimSuffix(model, "-go")
	}
	return model
}

type streamAnnotation struct {
	Type        string `json:"type"`
	URLCitation struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"url_citation"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type streamChunk struct {
	Choices []struct {
		Delta struct {
			Content          string             `json:"content"`
			ReasoningContent string             `json:"reasoning_content"`
			Reasoning        string             `json:"reasoning"`
			Annotations      []streamAnnotation `json:"annotations"`
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

// webAnnotationsOf extrait les citations web (recherche native, ex. plugin
// web d'OpenRouter) d'un delta de stream, en dedupliquant par URL.
func webAnnotationsOf(raw []streamAnnotation) []WebAnnotation {
	seen := map[string]bool{}
	var out []WebAnnotation
	for _, a := range raw {
		u := a.URLCitation.URL
		if u == "" {
			u = a.URL
		}
		u = strings.TrimSpace(u)
		if u == "" || seen[u] {
			continue
		}
		seen[u] = true
		title := a.URLCitation.Title
		if title == "" {
			title = a.Title
		}
		out = append(out, WebAnnotation{URL: u, Title: strings.TrimSpace(title)})
	}
	return out
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
		"model":       apiModel(p.endpoint.Provider, req.Model),
		"messages":    withReasoningContentForced(req.Messages),
		"stream":      true,
		"temperature": req.Temperature,
		"stream_options": map[string]any{
			"include_usage": true,
		},
	}
	if len(req.Tools) > 0 {
		payload["tools"] = req.Tools
		// Appels d'outils parallèles autorisés (défaut OpenAI) : le modèle
		// regroupe les appels indépendants en un seul tour au lieu d'un
		// aller-retour par outil — gain majeur sur les boucles agentiques.
		// L'assemblage streaming (index triés) et l'exécuteur agent gèrent
		// déjà les appels multiples.
		payload["parallel_tool_calls"] = true
	}
	if req.MaxTokens > 0 {
		payload["max_tokens"] = req.MaxTokens
	}
	for k, v := range req.Extra {
		payload[k] = v
	}
	applyReasoning(payload, p.endpoint.Provider, req.EnableReasoning, req.ReasoningEffort)
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
		if len(ch.Delta.Annotations) > 0 {
			resp.Annotations = append(resp.Annotations, webAnnotationsOf(ch.Delta.Annotations)...)
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
		tc.Function.Name = strings.TrimSpace(tc.Function.Name)
		if tc.Function.Name == "" {
			// Appel sans nom : invalide, l'agent demandera au modele de le renvoyer.
			tc.invalid = true
		} else if fixed, ok := repairJSON(tc.Function.Arguments); ok {
			tc.Function.Arguments = fixed
		} else {
			// Arguments irreparables : invalide plutot qu'executer avec des args vides.
			tc.invalid = true
		}
		resp.ToolCalls = append(resp.ToolCalls, tc)
	}
	return resp, nil
}

// repairJSON tente de reparer des arguments JSON tronques (flux coupe en
// pleine valeur) en fermant les accolades/crochets ouverts. Retourne false
// si le contenu est irreparable : l'appel doit alors etre rejete, jamais
// execute avec des arguments vides ou partiels silencieux.
func repairJSON(raw string) (string, bool) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return "{}", true
	}
	if json.Valid([]byte(s)) {
		return s, true
	}
	// Essaie des points de coupe de plus en plus courts.
	for i := len(s); i > 0; i-- {
		if fixed, ok := closeBrackets(strings.TrimSpace(s[:i])); ok {
			return fixed, true
		}
	}
	return "", false
}

// closeBrackets ferme les accolades/crochets restes ouverts. Si le texte se
// termine au milieu d'une chaine JSON, celle-ci est abandonnee proprement.
func closeBrackets(s string) (string, bool) {
	var stack []byte
	inStr := false
	esc := false
	lastSafe := -1 // dernier index ou couper proprement (hors chaine)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if inStr {
			switch {
			case esc:
				esc = false
			case c == '\\':
				esc = true
			case c == '"':
				inStr = false
				lastSafe = i
			}
			continue
		}
		switch c {
		case '"':
			inStr = true
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return "", false
			}
			stack = stack[:len(stack)-1]
			lastSafe = i
		case ',':
			lastSafe = i
		}
	}
	if inStr {
		if lastSafe < 0 {
			return "", false
		}
		return closeBrackets(s[:lastSafe+1])
	}
	var b strings.Builder
	b.WriteString(s)
	for i := len(stack) - 1; i >= 0; i-- {
		b.WriteByte(stack[i])
	}
	out := b.String()
	if !json.Valid([]byte(out)) {
		return "", false
	}
	return out, true
}

// messageWireForceReasoning : copie de Message sans omitempty sur
// reasoning_content. Utilise uniquement quand du reasoning circule
// dans la conversation (voir withReasoningContentForced) : DeepSeek
// exige alors le champ sur TOUS les messages assistant, meme vide,
// sous peine de HTTP 400. Garder les champs synchronises avec Message.
type messageWireForceReasoning struct {
	Role             string     `json:"role"`
	Content          any        `json:"content,omitempty"`
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID       string     `json:"tool_call_id,omitempty"`
	ReasoningContent string     `json:"reasoning_content"`
}

// withReasoningContentForced prepare les messages pour l'envoi.
//
// Declenchement : au moins un message assistant porte un
// reasoning_content non vide. Dans ce cas, le champ est force (meme
// vide : "") sur chaque message assistant — c'est le contrat
// thinking de DeepSeek (HTTP 400 « The reasoning_content in the
// thinking mode must be passed back to the API » sinon), partage par
// Kimi et GLM, et traverse tel quel par OpenRouter. Seuls les
// messages assistant sont convertis ; les autres roles gardent leur
// serialisation d'origine a l'identique.
//
// Sans reasoning en jeu, les messages sont renvoyes tels quels : le
// payload est strictement identique a avant (aucun risque de
// regression pour les autres providers).
func withReasoningContentForced(msgs []Message) any {
	forced := false
	for _, m := range msgs {
		if m.Role == "assistant" && m.ReasoningContent != "" {
			forced = true
			break
		}
	}
	if !forced {
		return msgs
	}
	out := make([]any, len(msgs))
	for i, m := range msgs {
		if m.Role != "assistant" {
			out[i] = m
			continue
		}
		out[i] = messageWireForceReasoning{
			Role:             m.Role,
			Content:          m.Content,
			ToolCalls:        m.ToolCalls,
			ToolCallID:       m.ToolCallID,
			ReasoningContent: m.ReasoningContent,
		}
	}
	return out
}

func applyReasoning(payload map[string]any, providerID string, enable bool, effort string) {
	if !enable {
		switch providerID {
		case "openrouter":
			payload["reasoning"] = map[string]any{"enabled": false}
		case "deepseek", "llamacpp", "ollama", "lmstudio":
			payload["chat_template_kwargs"] = map[string]any{"enable_thinking": false}
		}
		return
	}
	switch effort {
	case "low", "medium", "high":
		if providerID == "openrouter" {
			payload["reasoning"] = map[string]any{"effort": effort}
		} else {
			payload["reasoning_effort"] = effort
		}
	}
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
