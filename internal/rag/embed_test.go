package rag

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// Les dimensions sont figees ici : elles ont ete validees contre le
// catalogue OpenRouter reel le 2026-09-24. Un changement cote provider
// (ex. MRL) doit faire echouer ce test, pas passer inapercu.
func TestEmbedModelRegistry(t *testing.T) {
	want := map[string]int{
		"openai-text-embedding-3-small": 1536,
		"openai-text-embedding-3-large": 3072,
		"mistralai-mistral-embed-2312":  1024,
		"qwen-qwen3-embedding-4b":       2560,
		"google-gemini-embedding-001":   3072,
	}
	if len(EmbedModels) != len(want) {
		t.Fatalf("%d modeles, attendus %d", len(EmbedModels), len(want))
	}
	for _, m := range EmbedModels {
		d, ok := want[m.Slug]
		if !ok {
			t.Fatalf("modele inattendu: %s", m.Slug)
		}
		if m.Dims != d {
			t.Fatalf("%s: %d dims, attendues %d", m.Slug, m.Dims, d)
		}
		if !strings.Contains(m.OpenRouterID, "/") {
			t.Fatalf("%s: identifiant OpenRouter invalide: %q", m.Slug, m.OpenRouterID)
		}
	}
	if _, ok := LookupEmbedModel(DefaultEmbedModel); !ok {
		t.Fatalf("modele par defaut introuvable: %s", DefaultEmbedModel)
	}
	if _, ok := LookupEmbedModel("inexistant"); ok {
		t.Fatal("modele inexistant trouve")
	}
}

// fakeEmbedServer : serveur d'embedding factice. capture la requete,
// repond avec les vecteurs de respond (dans l'ordre des index).
func fakeEmbedServer(t *testing.T, respond [][]float32, status int, captured *embedRequest) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var er embedRequest
		if err := json.NewDecoder(r.Body).Decode(&er); err != nil {
			t.Errorf("requete illisible: %v", err)
		}
		if captured != nil {
			*captured = er
		}
		if got := r.Header.Get("Authorization"); got != "Bearer cle-test" {
			t.Errorf("Authorization = %q", got)
		}
		w.WriteHeader(status)
		if status != http.StatusOK {
			return
		}
		data := make([]struct {
			Embedding []float32 `json:"embedding"`
			Index     int       `json:"index"`
		}, len(respond))
		// Reponse volontairement melangee : l'ordre de sortie doit suivre
		// les index, pas l'ordre d'arrivee.
		for i, v := range respond {
			j := len(respond) - 1 - i
			data[j].Embedding = v
			data[j].Index = i
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"data": data})
	}))
}

func testEmbedder(t *testing.T, srv *httptest.Server, model EmbedModel) *Embedder {
	t.Helper()
	old := openRouterEmbedURL
	openRouterEmbedURL = srv.URL
	t.Cleanup(func() { openRouterEmbedURL = old })
	return NewOpenRouterEmbedder("cle-test", model, srv.Client())
}

func TestEmbedderRoundtrip(t *testing.T) {
	var captured embedRequest
	vecs := [][]float32{{0.1, 0.2, 0.3}, {0.4, 0.5, 0.6}}
	// Le modele small fait 1536 dims : on triche via un modele de test.
	testModel := EmbedModel{Slug: "test", OpenRouterID: "test/model", Dims: 3}
	srv := fakeEmbedServer(t, vecs, http.StatusOK, &captured)
	defer srv.Close()
	emb := testEmbedder(t, srv, testModel)

	got, err := emb.Embed(context.Background(), []string{"a", "b"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0][0] != 0.1 || got[1][2] != 0.6 {
		t.Fatalf("vecteurs melanges: %v", got)
	}
	if captured.Model != "test/model" {
		t.Fatalf("model = %q", captured.Model)
	}
	if len(captured.Input) != 2 || captured.Input[0] != "a" {
		t.Fatalf("input = %v", captured.Input)
	}
}

func TestEmbedderQueryPrefix(t *testing.T) {
	var captured embedRequest
	qwen, _ := LookupEmbedModel("qwen-qwen3-embedding-4b")
	testModel := EmbedModel{Slug: "t", OpenRouterID: "t/m", Dims: 2, QueryPrefix: qwen.QueryPrefix}
	srv := fakeEmbedServer(t, [][]float32{{0.1, 0.2}}, http.StatusOK, &captured)
	defer srv.Close()
	emb := testEmbedder(t, srv, testModel)

	if _, err := emb.Embed(context.Background(), []string{"question"}, true); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(captured.Input[0], qwen.QueryPrefix) {
		t.Fatalf("prefixe requete absent: %q", captured.Input[0])
	}
	// Les documents ne recoivent jamais le prefixe.
	if _, err := emb.Embed(context.Background(), []string{"question"}, false); err != nil {
		t.Fatal(err)
	}
	if strings.HasPrefix(captured.Input[0], qwen.QueryPrefix) {
		t.Fatalf("prefixe applique a un document: %q", captured.Input[0])
	}
}

func TestEmbedderErrors(t *testing.T) {
	testModel := EmbedModel{Slug: "t", OpenRouterID: "t/m", Dims: 3}
	// Statut non-200.
	srv := fakeEmbedServer(t, nil, http.StatusUnauthorized, nil)
	defer srv.Close()
	emb := testEmbedder(t, srv, testModel)
	if _, err := emb.Embed(context.Background(), []string{"a"}, false); err == nil {
		t.Fatal("erreur attendue (401)")
	}
	// Dimensions incoherentes.
	var captured embedRequest
	srv2 := fakeEmbedServer(t, [][]float32{{0.1}}, http.StatusOK, &captured)
	defer srv2.Close()
	emb2 := testEmbedder(t, srv2, testModel)
	if _, err := emb2.Embed(context.Background(), []string{"a"}, false); err == nil {
		t.Fatal("erreur attendue (dims)")
	}
	// Sans cle : fail-closed ici (l'appelant applique le fail-open).
	emb3 := NewOpenRouterEmbedder("", testModel, nil)
	if _, err := emb3.Embed(context.Background(), []string{"a"}, false); err == nil {
		t.Fatal("erreur attendue (cle manquante)")
	}
}

// TestBuildVectorsBatches : BuildVectors decoupe en lots de 100 et
// preserve l'ordre des chunks (alignement 1:1 avec l'index).
func TestBuildVectorsBatches(t *testing.T) {
	dir := t.TempDir()
	// 250 fichiers => 250 chunks => 3 lots (96+96+58, embedBatchMaxSize=96).
	for i := 0; i < 250; i++ {
		body := fmt.Sprintf("# Titre\n\ncontenu numero %d pomme", i)
		if err := os.WriteFile(fmt.Sprintf(dir+"/doc%03d.md", i), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	ix, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(ix.chunks) != 250 {
		t.Fatalf("%d chunks", len(ix.chunks))
	}
	var batches int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var er embedRequest
		_ = json.NewDecoder(r.Body).Decode(&er)
		batches++
		data := make([]struct {
			Embedding []float32 `json:"embedding"`
			Index     int       `json:"index"`
		}, len(er.Input))
		for i := range er.Input {
			// Vecteur distinctif par position dans le lot et numero de lot.
			data[i].Embedding = []float32{float32(i), float32(batches)}
			data[i].Index = i
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"data": data})
	}))
	defer srv.Close()
	old := openRouterEmbedURL
	openRouterEmbedURL = srv.URL
	defer func() { openRouterEmbedURL = old }()

	model := EmbedModel{Slug: "t", OpenRouterID: "t/m", Dims: 2}
	oldModels := EmbedModels
	EmbedModels = append(EmbedModels, model)
	defer func() { EmbedModels = oldModels }()
	emb := NewOpenRouterEmbedder("cle-test", model, srv.Client())
	if err := BuildVectors(dir, emb, func(_, _ int) {}); err != nil {
		t.Fatal(err)
	}
	if batches != 3 {
		t.Fatalf("%d lots, attendus 3", batches)
	}
	vs, err := LoadVectorFile(VectorFilePath(dir, model.Slug))
	if err != nil {
		t.Fatal(err)
	}
	if vs.Count != 250 || vs.Dims != 2 {
		t.Fatalf("store = %d chunks x %d dims", vs.Count, vs.Dims)
	}
	if vs.Hash != ix.CorpusHash() {
		t.Fatal("hash non aligne sur l'index")
	}
	// Ordre : le chunk i doit porter le vecteur (i%96, lot).
	// Note : LoadVectorFile normalise ; on compare les directions en
	// evitant la division par zero (x=0,lot=1 => vecteur (0,1) OK).
	for i := 0; i < 250; i++ {
		v := vs.vecs[i*2 : (i+1)*2]
		var want [2]float64
		want[0] = float64(i % embedBatchMaxSize)
		want[1] = float64(i/embedBatchMaxSize + 1)
		norm := sqrt(want[0]*want[0] + want[1]*want[1])
		dot := float64(v[0])*want[0]/norm + float64(v[1])*want[1]/norm
		if dot < 0.99999 {
			t.Fatalf("chunk %d: vecteur %v desordonne", i, v)
		}
	}
}
