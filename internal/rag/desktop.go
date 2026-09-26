package rag

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

// Plateforme RAG desktop (cetasrag) : /embed (BAAI/bge-m3, 1024 dims) et
// /rerank (BAAI/bge-reranker-v2-m3). Ces endpoints sont optionnels : sans
// URL ou sans cle, la jambe semantique reste BM25 seul (fail-open).

// Reranker : client du endpoint /rerank. Construit une fois au demarrage.
type Reranker struct {
	baseURL string
	apiKey  string
	topN    int
	timeout time.Duration
	client  *http.Client
}

// NewReranker : client /rerank. topN <= 0 => defaut 20 ; timeoutMs <= 0 =>
// defaut 4000. client nil => http.DefaultClient.
func NewReranker(baseURL, apiKey string, topN, timeoutMs int, client *http.Client) *Reranker {
	if topN <= 0 {
		topN = 20
	}
	if timeoutMs <= 0 {
		timeoutMs = 4000
	}
	if client == nil {
		client = http.DefaultClient
	}
	return &Reranker{
		baseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		apiKey:  strings.TrimSpace(apiKey),
		topN:    topN,
		timeout: time.Duration(timeoutMs) * time.Millisecond,
		client:  client,
	}
}

// TopN : nombre de candidats envoyes au reranker.
func (r *Reranker) TopN() int { return r.topN }

type rerankRequest struct {
	Query string   `json:"query"`
	Texts []string `json:"texts"`
}

type rerankHit struct {
	Index int     `json:"index"`
	Score float64 `json:"score"`
}

// Rerank : reordonne texts pour query. Retourne les index (dans texts) tries
// par score decroissant. Erreur (timeout, reseau, 5xx, 401) => l'appelant
// garde l'ordre RRF (fail-open). Une reponse vide est une erreur.
func (r *Reranker) Rerank(ctx context.Context, query string, texts []string) ([]rerankHit, error) {
	if r == nil || r.apiKey == "" {
		return nil, fmt.Errorf("reranker indisponible (cle manquante)")
	}
	if len(texts) == 0 {
		return nil, nil
	}
	body, err := json.Marshal(rerankRequest{Query: query, Texts: texts})
	if err != nil {
		return nil, err
	}
	cctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(cctx, http.MethodPost, r.baseURL+"/rerank", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apiKey)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("rerank: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("rerank: cle invalide (401)")
	}
	if resp.StatusCode != http.StatusOK {
		snippet := make([]byte, 300)
		n, _ := io.ReadFull(resp.Body, snippet)
		return nil, fmt.Errorf("rerank: statut %d: %s", resp.StatusCode,
			strings.TrimSpace(string(snippet[:n])))
	}
	var out []rerankHit
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("rerank: reponse illisible: %w", err)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("rerank: reponse vide")
	}
	for _, h := range out {
		if h.Index < 0 || h.Index >= len(texts) {
			return nil, fmt.Errorf("rerank: index %d hors bornes", h.Index)
		}
	}
	return out, nil
}
