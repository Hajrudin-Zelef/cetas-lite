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

// EmbedModel : un modele d'embedding adressable.
// Les identifiants et dimensions OpenRouter ont ete valides contre le
// catalogue reel le 2026-09-24 ; bge-m3 est servi par la plateforme RAG
// desktop (cetasrag), 1024 dims.
type EmbedModel struct {
	// Slug : identifiant local, utilise pour le nom du fichier de vecteurs.
	Slug string
	// OpenRouterID : identifiant envoye a l'API OpenRouter (backend
	// "openrouter"). Inutilise par le backend "desktop".
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
	// Plateforme RAG desktop (cetasrag, BAAI/bge-m3) : backend par defaut.
	{Slug: "BAAI/bge-m3", Dims: 1024},
	// Backend de rollback OpenRouter (conserves : retour arriere manuel).
	{Slug: "openai-text-embedding-3-small", OpenRouterID: "openai/text-embedding-3-small", Dims: 1536},
	{Slug: "openai-text-embedding-3-large", OpenRouterID: "openai/text-embedding-3-large", Dims: 3072},
	{Slug: "mistralai-mistral-embed-2312", OpenRouterID: "mistralai/mistral-embed-2312", Dims: 1024},
	{Slug: "qwen-qwen3-embedding-4b", OpenRouterID: "qwen/qwen3-embedding-4b", Dims: 2560,
		QueryPrefix: "Represent this sentence for searching relevant passages: "},
	{Slug: "google-gemini-embedding-001", OpenRouterID: "google/gemini-embedding-001", Dims: 3072},
}

// DefaultEmbedModel : backend desktop (bge-m3, 1024 dims).
const DefaultEmbedModel = "BAAI/bge-m3"

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

// Embedder : client d'embedding. Deux backends : "desktop" (plateforme RAG
// cetasrag, POST /embed, format [[...]]) et "openrouter" (endpoint
// OpenAI-compatible, format {data:[{embedding}]}), garde pour le rollback.
// Le champ desktop distingue les deux au moment de la requete.
type Embedder struct {
	apiKey  string
	model   EmbedModel
	client  *http.Client
	desktop bool
	baseURL string
}

// NewOpenRouterEmbedder : constructeur du backend OpenRouter (rollback).
func NewOpenRouterEmbedder(apiKey string, model EmbedModel, client *http.Client) *Embedder {
	if client == nil {
		client = http.DefaultClient
	}
	return &Embedder{apiKey: strings.TrimSpace(apiKey), model: model, client: client}
}

// NewDesktopEmbedder : constructeur du backend desktop (cetasrag). baseURL
// est l'URL du serveur (ex. https://cetasrag.neva-ci.pro) : /embed et
// /rerank sont ajoutes. Cle vide => embedder indisponible (fail-open).
func NewDesktopEmbedder(baseURL, apiKey string, model EmbedModel, client *http.Client) *Embedder {
	if client == nil {
		client = http.DefaultClient
	}
	return &Embedder{
		apiKey:  strings.TrimSpace(apiKey),
		model:   model,
		client:  client,
		desktop: true,
		baseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/"),
	}
}

// Model : le modele d'embedding utilise.
func (e *Embedder) Model() EmbedModel { return e.model }

// Desktop : vrai si l'embedder parle au serveur desktop (cetasrag).
func (e *Embedder) Desktop() bool { return e != nil && e.desktop }

// embedRequest / embedResponse : format OpenAI-compatible (backend openrouter).
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

// desktopEmbedRequest : format de la plateforme cetasrag (POST /embed).
// "inputs" accepte une chaine seule ou une liste.
type desktopEmbedRequest struct {
	Model  string      `json:"model,omitempty"`
	Inputs interface{} `json:"inputs"`
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
	cctx, cancel := context.WithTimeout(ctx, embedTimeout)
	defer cancel()
	if e.desktop {
		return e.embedDesktop(cctx, in)
	}
	return e.embedOpenRouter(cctx, in)
}

// embedDesktop : POST /embed sur la plateforme cetasrag. Reponse [[...]] :
// un vecteur par entree, dans l'ordre. La dimension est verifiee.
func (e *Embedder) embedDesktop(ctx context.Context, in []string) ([][]float32, error) {
	body, err := json.Marshal(desktopEmbedRequest{Inputs: in})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, e.baseURL+"/embed", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.apiKey)
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embeddings desktop: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("embeddings desktop: cle invalide (401)")
	}
	if resp.StatusCode != http.StatusOK {
		snippet := make([]byte, 400)
		n, _ := io.ReadFull(resp.Body, snippet)
		return nil, fmt.Errorf("embeddings desktop: statut %d: %s", resp.StatusCode,
			strings.TrimSpace(string(snippet[:n])))
	}
	var out [][]float32
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("embeddings desktop: reponse illisible: %w", err)
	}
	if len(out) != len(in) {
		return nil, fmt.Errorf("embeddings desktop: %d vecteurs pour %d textes", len(out), len(in))
	}
	for i, v := range out {
		if len(v) != e.model.Dims {
			return nil, fmt.Errorf("embeddings desktop: %d dims recues, %d attendues (%s)",
				len(v), e.model.Dims, e.model.Slug)
		}
		out[i] = v
	}
	return out, nil
}

// embedOpenRouter : endpoint OpenAI-compatible d'OpenRouter (rollback).
func (e *Embedder) embedOpenRouter(ctx context.Context, in []string) ([][]float32, error) {
	body, err := json.Marshal(embedRequest{Model: e.model.OpenRouterID, Input: in})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, openRouterEmbedURL, bytes.NewReader(body))
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
