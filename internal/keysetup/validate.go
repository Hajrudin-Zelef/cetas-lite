package keysetup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// validationTimeout — borne chaque appel de test (setup.py : 30 s).
const validationTimeout = 30 * time.Second

// TestKey valide une clé contre l'API réelle du provider.
// Retourne (true, "") si la clé fonctionne, sinon (false, raison lisible).
// Une erreur réseau (DNS, timeout) est distinguée d'une clé rejetée (401/403).
// Si FallbackModel est défini, il est essayé quand le modèle principal échoue
// (modèle désactivé/renommé côté provider, sans que la clé soit en cause).
func TestKey(p Provider, key string) (bool, string) {
	if p.TestURL == "" {
		return false, "provider désactivé (URL de test non configurée)"
	}
	if strings.TrimSpace(key) == "" {
		return false, "clé vide"
	}

	ctx, cancel := context.WithTimeout(context.Background(), validationTimeout)
	defer cancel()

	models := []string{p.TestModel}
	if p.FallbackModel != "" {
		models = append(models, p.FallbackModel)
	}

	var reason string
	for _, model := range models {
		ok, r := testKeyModel(ctx, p, key, model)
		if ok {
			return true, ""
		}
		reason = r
		// Quota épuisé, provider injoignable : la clé n'est pas en cause,
		// réessayer avec un autre modèle ne changerait rien.
		if strings.Contains(r, "HTTP 429") || strings.HasPrefix(r, "réseau") {
			return false, r
		}
	}
	return false, reason
}

// testKeyModel effectue UNE validation avec un modèle donné.
func testKeyModel(ctx context.Context, p Provider, key, model string) (bool, string) {
	var req *http.Request
	var err error
	switch p.Mode {
	case ModeResponses:
		req, err = responsesRequest(ctx, p, key, model)
	case ModeAnthropic:
		req, err = anthropicRequest(ctx, p, key, model)
	default:
		req, err = chatRequest(ctx, p, key, model)
	}
	if err != nil {
		return false, "requête impossible : " + err.Error()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, "réseau : " + shortNetErr(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

	switch resp.StatusCode {
	case 200, 201:
		return true, ""
	case 401, 403:
		return false, fmt.Sprintf("clé rejetée (HTTP %d) — %s", resp.StatusCode, apiErrMsg(body))
	case 400:
		return false, fmt.Sprintf("requête refusée (HTTP 400) — %s", apiErrMsg(body))
	case 404:
		return false, fmt.Sprintf("endpoint introuvable (HTTP 404) — %s", apiErrMsg(body))
	case 429:
		return false, "quota dépassé (HTTP 429) — la clé semble valide mais le quota est épuisé"
	default:
		if resp.StatusCode >= 500 {
			return false, fmt.Sprintf("erreur serveur du provider (HTTP %d) — réessayez", resp.StatusCode)
		}
		return false, fmt.Sprintf("HTTP %d — %s", resp.StatusCode, apiErrMsg(body))
	}
}

// httpClient — client partagé, timeouts bornés (jamais de client sans timeout).
var httpClient = &http.Client{Timeout: validationTimeout + 5*time.Second}

func chatRequest(ctx context.Context, p Provider, key, model string) (*http.Request, error) {
	payload := map[string]any{
		"model":      model,
		"messages":   []map[string]string{{"role": "user", "content": "Hi"}},
		"max_tokens": 1,
		"stream":     false,
	}
	return jsonBearerRequest(ctx, "POST", p.TestURL, key, payload, p.Headers)
}

func responsesRequest(ctx context.Context, p Provider, key, model string) (*http.Request, error) {
	payload := map[string]any{"model": model, "input": "Hi"}
	return jsonBearerRequest(ctx, "POST", p.TestURL, key, payload, p.Headers)
}

// anthropicRequest — l'API Anthropic native n'est PAS compatible OpenAI :
// auth via x-api-key + version d'API, corps {model, max_tokens, messages}.
func anthropicRequest(ctx context.Context, p Provider, key, model string) (*http.Request, error) {
	payload := map[string]any{
		"model":      model,
		"max_tokens": 1,
		"messages":   []map[string]string{{"role": "user", "content": "Hi"}},
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", p.TestURL, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", key)
	req.Header.Set("anthropic-version", "2023-06-01")
	return req, nil
}

func jsonBearerRequest(ctx context.Context, method, url, key string, payload any, extra map[string]string) (*http.Request, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	for k, v := range extra {
		req.Header.Set(k, v)
	}
	return req, nil
}

// apiErrMsg — extrait un message d'erreur lisible d'un corps de réponse.
func apiErrMsg(body []byte) string {
	var v struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
		Message string `json:"message"`
	}
	if json.Unmarshal(body, &v) == nil {
		if v.Error.Message != "" {
			return truncate(v.Error.Message, 160)
		}
		if v.Message != "" {
			return truncate(v.Message, 160)
		}
	}
	s := strings.TrimSpace(string(body))
	return truncate(s, 160)
}

// shortNetErr — distingue timeout / DNS / connexion refusée.
func shortNetErr(err error) string {
	s := err.Error()
	switch {
	case strings.Contains(s, "deadline exceeded") || strings.Contains(s, "timeout"):
		return "timeout (30 s) — provider injoignable"
	case strings.Contains(s, "no such host"):
		return "DNS — hôte introuvable"
	case strings.Contains(s, "connection refused"):
		return "connexion refusée"
	case strings.Contains(s, "certificate"):
		return "erreur TLS — certificat invalide"
	default:
		return truncate(s, 120)
	}
}

func truncate(s string, n int) string {
	s = strings.Join(strings.Fields(s), " ")
	if len(s) > n {
		return s[:n] + "…"
	}
	if s == "" {
		return "(réponse vide)"
	}
	return s
}
