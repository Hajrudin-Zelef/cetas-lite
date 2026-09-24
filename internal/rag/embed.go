package rag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Lot 2 — recherche hybride : BM25 (mots exacts) + embeddings (semantique),
// fusion RRF. L'embedding passe par l'API OpenRouter (endpoint OpenAI-
// compatible /embeddings) avec la cle deja configuree pour l'inference :
// aucun modele local sur le VPS (RAM et sidecar incompatibles).

// EmbedModel : un modele d'embedding adressable via OpenRouter.
// Les identifiants et dimensions ont ete valides contre le catalogue
// OpenRouter reel le 2026-09-24.
type EmbedModel struct {
	// Slug : identifiant local, utilise pour le nom du fichier de vecteurs.
	Slug string
	// OpenRouterID : identifiant envoye a l'API.
	OpenRouterID string
	// Dims : largeur native des vecteurs.
	Dims int
	// QueryPrefix : prefixe d'instruction applique aux requetes seules
	// (jamais aux documents), quand le modele en beneficie (Qwen3).
	QueryPrefix string
}

// EmbedModels : modeles supportes. Changer de modele invalide les vecteurs
// existants (espaces semantiques incompatibles) : un fichier de vecteurs
// est namespacé par Slug, et la bascule reutilise le fichier deja construit
// (voir vectors.go).
var EmbedModels = []EmbedModel{
	{Slug: "openai-text-embedding-3-small", OpenRouterID: "openai/text-embedding-3-small", Dims: 1536},
	{Slug: "openai-text-embedding-3-large", OpenRouterID: "openai/text-embedding-3-large", Dims: 3072},
	{Slug: "mistralai-mistral-embed-2312", OpenRouterID: "mistralai/mistral-embed-2312", Dims: 1024},
	{Slug: "qwen-qwen3-embedding-4b", OpenRouterID: "qwen/qwen3-embedding-4b", Dims: 2560,
		QueryPrefix: "Represent this sentence for searching relevant passages: "},
	{Slug: "google-gemini-embedding-001", OpenRouterID: "google/gemini-embedding-001", Dims: 3072},
}

// DefaultEmbedModel : compromis cout/qualite/taille (15 Mo pour 2487 chunks
// en float32). $0.02/M tokens : l'indexation complete coute quelques centimes.
const DefaultEmbedModel = "openai-text-embedding-3-small"

// EmbedModelSlugFromEnv : modele selectionne via CETAS_LITE_EMBED_MODEL,
// defaut sinon. L'UI de selection (lot 4) ecrira ici.
func EmbedModelSlugFromEnv() string {
	if s := strings.TrimSpace(os.Getenv("CETAS_LITE_EMBED_MODEL")); s != "" {
		return s
	}
	return DefaultEmbedModel
}

// LookupEmbedModel : retrouve un modele par son Slug.
func LookupEmbedModel(slug string) (EmbedModel, bool) {
	for _, m := range EmbedModels {
		if m.Slug == slug {
			return m, true
		}
	}
	return EmbedModel{}, false
}

// openRouterEmbedURL : endpoint embeddings OpenAI-compatible d'OpenRouter.
// Variable (et non constante) pour permettre aux tests de pointer vers un
// serveur factice.
var openRouterEmbedURL = "https://openrouter.ai/api/v1/embeddings"

// Embedder : client d'embedding via OpenRouter. Reutilise la cle API deja
// configuree pour l'inference (aucune cle supplementaire).
type Embedder struct {
	apiKey string
	model  EmbedModel
	client *http.Client
}

// NewOpenRouterEmbedder : construit le client. client nil => http.DefaultClient.
func NewOpenRouterEmbedder(apiKey string, model EmbedModel, client *http.Client) *Embedder {
	if client == nil {
		client = http.DefaultClient
	}
	return &Embedder{apiKey: strings.TrimSpace(apiKey), model: model, client: client}
}

// Model : le modele d'embedding utilise.
func (e *Embedder) Model() EmbedModel { return e.model }

// embedRequest / embedResponse : format OpenAI-compatible.
type embedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type embedResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
}

// embedTimeout : budget d'un appel d'embedding (indexation par lots).
const embedTimeout = 60 * time.Second

// Embed : vectorise texts en un appel. L'ordre de sortie suit l'ordre
// d'entree. fail-closed ici (erreur remontee) : c'est l'appelant
// (hybrid.go, CLI) qui applique le fail-open.
func (e *Embedder) Embed(ctx context.Context, texts []string, forQuery bool) ([][]float32, error) {
	if e == nil || e.apiKey == "" {
		return nil, fmt.Errorf("embedder indisponible (cle manquante)")
	}
	if len(texts) == 0 {
		return nil, nil
	}
	in := texts
	if forQuery && e.model.QueryPrefix != "" {
		in = make([]string, len(texts))
		for i, t := range texts {
			in[i] = e.model.QueryPrefix + t
		}
	}
	body, err := json.Marshal(embedRequest{Model: e.model.OpenRouterID, Input: in})
	if err != nil {
		return nil, err
	}
	cctx, cancel := context.WithTimeout(ctx, embedTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(cctx, http.MethodPost, openRouterEmbedURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.apiKey)
	req.Header.Set("HTTP-Referer", "https://github.com/Hajrudin-Zelef/cetas-lite")
	req.Header.Set("X-Title", "cetas-lite RAG")
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embeddings: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		snippet := make([]byte, 400)
		n, _ := io.ReadFull(resp.Body, snippet)
		return nil, fmt.Errorf("embeddings: statut %d: %s", resp.StatusCode,
			strings.TrimSpace(string(snippet[:n])))
	}
	var er embedResponse
	if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
		return nil, fmt.Errorf("embeddings: reponse illisible: %w", err)
	}
	if len(er.Data) != len(in) {
		return nil, fmt.Errorf("embeddings: %d vecteurs pour %d textes", len(er.Data), len(in))
	}
	out := make([][]float32, len(in))
	for _, d := range er.Data {
		if d.Index < 0 || d.Index >= len(in) {
			return nil, fmt.Errorf("embeddings: index %d hors bornes", d.Index)
		}
		if len(d.Embedding) != e.model.Dims {
			return nil, fmt.Errorf("embeddings: %d dims recues, %d attendues (%s)",
				len(d.Embedding), e.model.Dims, e.model.Slug)
		}
		out[d.Index] = d.Embedding
	}
	return out, nil
}
