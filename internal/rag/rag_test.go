package rag

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode/utf8"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
}

func richCorpus(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "corpusA", "01-topic", "alpha.md"),
		"# Alpha\n\nLe noyau Linux et les serveurs GPU.\n")
	writeFile(t, filepath.Join(root, "corpusA", "01-topic", "beta.md"),
		"# Beta\n\nQuantization et inference locale.\n")
	writeFile(t, filepath.Join(root, "corpusA", manifestFile), `{
  "corpus": "corpusA",
  "title": "Corpus A",
  "relationship": "base",
  "chunks": [
    {"path": "corpusA/01-topic/alpha.md", "title": "Alpha", "domain": "01-topic",
     "task": "analysis", "actors": ["Linux"], "dates": ["2026-01-02"],
     "keywords": ["kernel", "gpu"], "section": "Alpha"},
    {"path": "corpusA/01-topic/beta.md", "title": "Beta", "domain": "01-topic",
     "task": "analysis", "keywords": ["quantization"], "delta_of": "corpusA"}
  ]
}`)
	return root
}

func TestLoadMissingDirIsInactive(t *testing.T) {
	ix, err := Load(filepath.Join(t.TempDir(), "absent"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if ix.Stats().Ready {
		t.Fatal("un dossier absent doit donner un index inactif")
	}
	if r := ix.Search(context.Background(), "linux", 5); r.Ready || len(r.Hits) != 0 {
		t.Fatalf("recherche sur index inactif: %+v", r)
	}
}

func TestLoadManifestRich(t *testing.T) {
	ix, err := Load(richCorpus(t))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if !st.Ready || st.Corpora != 1 || st.Chunks != 2 {
		t.Fatalf("stats inattendues: %+v", st)
	}
	res := ix.Search(context.Background(), "quantization", 5)
	if len(res.Hits) != 1 {
		t.Fatalf("attendu 1 hit, obtenu %d", len(res.Hits))
	}
	h := res.Hits[0]
	if h.Path != "corpusA/01-topic/beta.md" || h.DeltaOf != "corpusA" {
		t.Fatalf("hit inattendu: %+v", h)
	}
	if res := ix.Search(context.Background(), "linux", 5); len(res.Hits) != 1 ||
		res.Hits[0].Actors[0] != "Linux" {
		t.Fatalf("facettes non indexees: %+v", res.Hits)
	}
}

func TestLoadRawFallback(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "perso", "notes.md"),
		"# Mes notes\n\nRecette de cafe filtre et maintenance du VPS.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if !st.Ready || st.Corpora != 1 || st.Chunks != 1 {
		t.Fatalf("stats inattendues: %+v", st)
	}
	res := ix.Search(context.Background(), "cafe", 5)
	if len(res.Hits) != 1 {
		t.Fatalf("attendu 1 hit, obtenu %d", len(res.Hits))
	}
	if res.Hits[0].Title != "Mes notes" {
		t.Fatalf("titre brut inattendu: %q", res.Hits[0].Title)
	}
	if res.Hits[0].Domain != "perso" || res.Hits[0].Corpus != "perso" {
		t.Fatalf("metadonnees brutes inattendues: %+v", res.Hits[0])
	}
	if !strings.Contains(res.Hits[0].Snippet, "cafe") {
		t.Fatalf("snippet sans le terme: %q", res.Hits[0].Snippet)
	}
}

func TestReadLookupAndUnknown(t *testing.T) {
	ix, _ := Load(richCorpus(t))
	full, err := ix.Read("corpusA/01-topic/alpha.md", 1, 2)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if !strings.HasPrefix(full, "1\t# Alpha") {
		t.Fatalf("lecture inattendue: %q", full)
	}
	if _, err := ix.Read("beta.md", 1, 5); err != nil {
		t.Fatalf("lecture par nom de base: %v", err)
	}
	if _, err := ix.Read("inconnu.md", 1, 5); err == nil {
		t.Fatal("chemin inconnu: erreur attendue")
	}
}

func TestSearchCanceledContext(t *testing.T) {
	ix, _ := Load(richCorpus(t))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = ix.Search(ctx, "quantization linux", 5)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("recherche bloquee sur un contexte annule")
	}
}

func TestManagerBackgroundLoadAndConcurrency(t *testing.T) {
	m := New(richCorpus(t))
	if !m.Enabled() {
		t.Fatal("manager attendu actif")
	}
	m.Start()
	deadline := time.Now().Add(2 * time.Second)
	for !m.Ready() && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if !m.Ready() {
		t.Fatal("index non pret apres Start")
	}
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if r := m.Search(context.Background(), "quantization", 3); len(r.Hits) == 0 {
				t.Errorf("hit manquant")
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := m.Reload(); err != nil {
			t.Errorf("Reload: %v", err)
		}
	}()
	wg.Wait()
}

func TestManagerEmptyRootDisabled(t *testing.T) {
	m := New("")
	if m.Enabled() {
		t.Fatal("root vide doit etre desactive")
	}
	if r := m.Search(context.Background(), "x", 5); r.Ready {
		t.Fatal("recherche sur manager desactive")
	}
}

func TestLoadRepoCorpus(t *testing.T) {
	root := filepath.Join("..", "..", "RAG")
	if _, err := os.Stat(filepath.Join(root, "briefing-ia-2026", manifestFile)); err != nil {
		t.Skip("corpus RAG absent")
	}
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if st.Corpora < 4 {
		t.Fatalf("attendu au moins 4 corpus, obtenu %d (%v)", st.Corpora, st.Errors)
	}
	if len(st.Errors) != 0 {
		t.Fatalf("erreurs de chargement: %v", st.Errors)
	}
	if st.Chunks < 400 {
		t.Fatalf("trop peu de chunks: %d", st.Chunks)
	}
	res := ix.Search(context.Background(), "quantization", 5)
	if len(res.Hits) == 0 {
		t.Fatal("aucun hit pour 'quantization'")
	}
	if strings.HasPrefix(res.Hits[0].Snippet, "---") {
		t.Fatalf("front-matter non nettoye: %q", res.Hits[0].Snippet)
	}
}

func BenchmarkSearch(b *testing.B) {
	root := filepath.Join("..", "..", "RAG")
	if _, err := os.Stat(filepath.Join(root, "briefing-ia-2026", manifestFile)); err != nil {
		b.Skip("corpus RAG absent")
	}
	ix, err := Load(root)
	if err != nil || !ix.Stats().Ready {
		b.Skip("index indisponible")
	}
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if len(ix.Search(ctx, "inference quantization gpu memory", 8).Hits) == 0 {
			b.Fatal("aucun hit")
		}
	}
}

func TestExcerptKeepsAccents(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "notes", "ia.md"),
		"# État de l'art\n\nLes données d'entraînement sont coûteuses à collecter.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	// "données" est normalisé en "donnees" pour le scoring, mais l'extrait
	// injecté doit garder le texte d'origine (accents, casse).
	res := ix.Search(context.Background(), "données entraînement", 5)
	if len(res.Hits) != 1 {
		t.Fatalf("attendu 1 hit, obtenu %d", len(res.Hits))
	}
	h := res.Hits[0]
	if !strings.Contains(h.Excerpt, "données") {
		t.Fatalf("extrait dégrade (accents perdus): %q", h.Excerpt)
	}
	if !strings.Contains(h.Excerpt, "État") {
		t.Fatalf("extrait dégrade (casse perdue): %q", h.Excerpt)
	}
	if strings.Contains(h.Excerpt, "donnees") && !strings.Contains(h.Excerpt, "données") {
		t.Fatalf("extrait issu du texte replie au lieu de l'original: %q", h.Excerpt)
	}
}

func TestFoldWithMapRoundtrip(t *testing.T) {
	for _, s := range []string{
		"L'état des données",
		"État",
		"plain ascii",
		"",
		"Ünïcödé — test",
	} {
		folded, fmap := foldWithMap(s)
		if len(fmap) != len(folded)+1 {
			t.Fatalf("fmap de taille %d pour %d octets repliés (%q)", len(fmap), len(folded), s)
		}
		if fmap[len(folded)] != len(s) {
			t.Fatalf("sentinelle finale %d != %d (%q)", fmap[len(folded)], len(s), s)
		}
		for i, off := range fmap[:len(folded)] {
			if off < 0 || off >= len(s) {
				t.Fatalf("offset %d hors bornes pour %q", off, s)
			}
			if !utf8.RuneStart(s[off]) {
				t.Fatalf("offset %d au milieu d'une rune pour %q", off, s)
			}
			_ = i
		}
	}
}

func TestFirstMatchRangeFallback(t *testing.T) {
	body := "Introduction.\n\nL'état des lieux est préoccupant.\n\nConclusion."
	// "état" ne matche pas en ASCII (é), le repli doit re-projeter sur l'original.
	start, end, ok := firstMatchRange(body, []string{"etat"})
	if !ok {
		t.Fatal("correspondance attendue via le repli accent-insensible")
	}
	if !strings.Contains(body[start:end], "état") {
		t.Fatalf("plage [%d:%d] = %q, attendu le mot d'origine accentue", start, end, body[start:end])
	}
	// Chemin direct (ASCII) inchange.
	start, end, ok = firstMatchRange(body, []string{"conclusion"})
	if !ok || body[start:end] != "Conclusion" && body[start:end] != "conclusion" {
		t.Fatalf("chemin direct inattendu: [%d:%d] = %q", start, end, body[start:end])
	}
	// Aucun terme : ok=false.
	if _, _, ok = firstMatchRange(body, []string{"xyzabc"}); ok {
		t.Fatal("ok attendu faux pour un terme absent")
	}
}

// TestCleanQueryForSearch : l'iteration 3 retire les formulations
// conversationnelles et les mots-outils avant BM25, sans toucher aux
// identifiants significatifs. Si le nettoyage vide la requete, l'origine
// est conservee (fail-open).
func TestCleanQueryForSearch(t *testing.T) {
	cases := []struct{ in, want string }{
		{"De nouveau sur kimi k3?", "kimi k3"},
		{"parle-moi de Kimi K3", "kimi k3"},
		{"s'il te plaît, qu'est-ce que vLLM ?", "vllm"},
		{"tell me about kimi k3, please", "kimi k3"},
		// Correctif 2026-09-24 : « nouveau »/« new » ne sont plus
		// conservés comme « significatifs ». Mesuré sur le corpus réel
		// (2487 chunks) : terme rare donc IDF élevé, il éjectait
		// l'entité porteuse (« Il ya du nouveau chez glm » classait un
		// article Kimi contenant « nouveau » #1 devant tout extrait
		// GLM). Le besoin d'information est porté par l'entité.
		{"le nouveau modèle de Google", "google"},
		{"what's new with kimi k3", "kimi k3"},
		// Résidu conversationnel mesuré sur les requêtes GLM fautives
		// (sessions B/C) : « tu », « peut », « dire », « sorti »,
		// « dernier », « pense », « men »/« na » (sms) écrasaient
		// « glm » (fréquent, IDF faible) — ex. « GLM-5.3 est sorti tu
		// peut men dire plus? » classait un article « PC gaming »
		// #1 sans aucun extrait GLM dans le top-3.
		{"Il ya du nouveau chez glm", "glm"},
		{"GLM-5.3 est sorti tu peut men dire plus?", "glm"},
		{"Le modèle glm-5.3 est le dernier je pense", "glm"},
		{"Donc tu na aucune info sur GLM-5.3 ?", "glm"},
		// Identifiants : jamais touches (« 5 »/« 3 » isoles sont de toute
		// facon sous minTermLen, comme dans l'indexation des documents).
		{"GLM-5.3", "glm"},
		{"kimi k3", "kimi k3"},
		// Fail-open : que du remplissage -> requete d'origine.
		{"De nouveau", "De nouveau"},
		{"s'il te plaît", "s'il te plaît"},
	}
	for _, c := range cases {
		if got := cleanQueryForSearch(c.in); got != c.want {
			t.Errorf("cleanQueryForSearch(%q) = %q, attendu %q", c.in, got, c.want)
		}
	}
}

// TestSearchQueryCleaningRanksKimiFirst : regression du cas Kimi K3 —
// « nouveau », rare et booste en titre, ne doit plus ejecter le vrai sujet.
func TestSearchQueryCleaningRanksKimiFirst(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "c1", "distractor.md"),
		"# Le nouveau modele de Google\n\nDiffusionGemma ecrit son texte d'un bloc et 4 fois plus vite.\n")
	writeFile(t, filepath.Join(root, "c1", "kimi.md"),
		"# Kimi K3\n\nModele ouvert de Moonshot AI : poids, specifications et tutoriel.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	res := ix.Search(context.Background(), "De nouveau sur kimi k3?", 5)
	if len(res.Hits) == 0 {
		t.Fatal("aucun hit")
	}
	if top := res.Hits[0].Path; !strings.HasSuffix(top, "kimi.md") {
		t.Fatalf("top hit = %q, attendu kimi.md", top)
	}
}

// TestSearchQueryCleaningRanksGlmFirst : régression du cas GLM-5.3
// (correctif 2026-09-24) — le résidu conversationnel (« nouveau », « tu »,
// « peut », « dire », « sorti »…), rare donc à IDF élevé, ne doit plus
// éjecter l'entité fréquente (« glm », IDF faible). Sans le correctif,
// « Il ya du nouveau chez glm » classait le distractor #1.
func TestSearchQueryCleaningRanksGlmFirst(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "c1", "distractor.md"),
		"# Un nouveau rival\n\nCe nouveau modèle fait parler avec ses nouveautés.\n")
	writeFile(t, filepath.Join(root, "c1", "glm.md"),
		"# GLM 5\n\nGLM est un modèle ouvert de Z.ai. GLM GLM : poids, tutoriel.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	for _, q := range []string{
		"Il ya du nouveau chez glm",
		"GLM-5.3 est sorti tu peut men dire plus?",
		"Le modèle glm-5.3 est le dernier je pense",
	} {
		res := ix.Search(context.Background(), q, 5)
		if len(res.Hits) == 0 {
			t.Fatalf("aucun hit pour %q", q)
		}
		if top := res.Hits[0].Path; !strings.HasSuffix(top, "glm.md") {
			t.Fatalf("top hit = %q pour %q, attendu glm.md", top, q)
		}
	}
}

// TestSearchCorpusFilters : la recherche restreinte à un corpus ne renvoie
// que des chunks de ce corpus ; corpus inconnu => aucun hit (fail-open) ;
// corpus vide => équivalent à Search.
func TestSearchCorpusFilters(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "cA", "doc.md"),
		"# CUDA toolkit\n\nLe toolkit CUDA de NVIDIA pour le calcul GPU.\n")
	writeFile(t, filepath.Join(root, "cB", "doc.md"),
		"# CUDA toolkit\n\nLe toolkit CUDA de NVIDIA pour le calcul GPU.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	res := ix.SearchCorpus(context.Background(), "cuda toolkit", "cA", 5)
	if len(res.Hits) == 0 {
		t.Fatal("aucun hit dans le corpus cA")
	}
	for _, h := range res.Hits {
		if h.Corpus != "cA" {
			t.Fatalf("hit hors corpus : %q (path %q)", h.Corpus, h.Path)
		}
	}
	if res := ix.SearchCorpus(context.Background(), "cuda toolkit", "nope", 5); len(res.Hits) != 0 {
		t.Fatalf("corpus inconnu : %d hits, attendu 0", len(res.Hits))
	}
	full := ix.Search(context.Background(), "cuda toolkit", 5)
	scoped := ix.SearchCorpus(context.Background(), "cuda toolkit", "", 5)
	if len(full.Hits) != len(scoped.Hits) {
		t.Fatalf("corpus vide : %d hits vs %d en Search", len(scoped.Hits), len(full.Hits))
	}
}
