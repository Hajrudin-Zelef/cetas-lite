package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"cetas-lite/internal/config"
	"cetas-lite/internal/rag"
	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
)

// Lot 2 — sous-commande `rag-vectors` : gestion des vecteurs d'embedding.
//
//	cetas-lite rag-vectors status            etat des vecteurs par modele
//	cetas-lite rag-vectors build [--model S] construit les vecteurs (cout API)
//	cetas-lite rag-vectors probe "requete"   top cosinus (calibration du seuil)
//
// La cle OpenRouter est lue dans le coffre (comme `serve`), sinon via
// OPENROUTER_API_KEY. Les vecteurs vivent dans <rag>/.vectors/<slug>.bin,
// ignores par le loader de corpus.

func runRagVectors(args []string) error {
	if len(args) == 0 {
		printRagVectorsUsage()
		return nil
	}
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if strings.TrimSpace(cfg.RagDir) == "" {
		return fmt.Errorf("dossier RAG non configure (CETAS_LITE_RAG_DIR)")
	}
	switch args[0] {
	case "build":
		return ragVectorsBuild(cfg, flagValue(args[1:], "model", rag.DefaultEmbedModel))
	case "status":
		return ragVectorsStatus(cfg)
	case "probe":
		rest := args[1:]
		q := probeQuery(rest)
		if strings.TrimSpace(q) == "" {
			return fmt.Errorf("usage: rag-vectors probe \"requete\" [--model S] [--top N]")
		}
		return ragVectorsProbe(cfg, q, flagValue(rest, "model", rag.DefaultEmbedModel), flagInt(rest, "top", 5))
	default:
		printRagVectorsUsage()
		return fmt.Errorf("action inconnue: %s", args[0])
	}
}

// probeQuery : la requete est le premier argument positionnel (les flags
// --model/--top/--query et leurs valeurs sont sautes), ou --query.
func probeQuery(args []string) string {
	if q := flagValue(args, "query", ""); q != "" {
		return q
	}
	var parts []string
	skip := false
	for _, a := range args {
		if skip {
			skip = false
			continue
		}
		if a == "--model" || a == "--top" {
			skip = true
			continue
		}
		if strings.HasPrefix(a, "--") {
			continue
		}
		parts = append(parts, a)
	}
	return strings.Join(parts, " ")
}

func printRagVectorsUsage() {
	fmt.Println("usage: cetas-lite rag-vectors <status|build|probe> [options]")
	fmt.Println("  status                 etat des vecteurs par modele")
	fmt.Println("  build [--model SLUG]   construit les vecteurs (appel API OpenRouter)")
	fmt.Println("  probe \"requete\" [--model SLUG] [--top N]")
	fmt.Println("                         affiche les meilleurs cosinus (calibration)")
	fmt.Println("modeles:")
	for _, m := range rag.EmbedModels {
		fmt.Printf("  %-38s %d dims\n", m.Slug, m.Dims)
	}
}

// embedKeyForCLI : cle d'embedding lue dans le coffre (priorite), sinon
// l'environnement de secours. Le coffre d'abord : l'env n'est qu'une
// facilite de test, jamais la source de verite (le commentaire d'origine
// annoncait cet ordre sans l'appliquer). Backend desktop => rag_desktop_key
// (repli OPENROUTER_API_KEY non, les deux cles ne sont pas interchangeables).
func embedKeyForCLI(cfg *config.Config, backend string) (string, error) {
	secret := "rag_desktop_key"
	envName := "CETAS_LITE_EMBED_KEY"
	if backend == "openrouter" {
		secret = "openrouter"
		envName = "OPENROUTER_API_KEY"
	}
	st, err := store.Open(cfg.DBPath)
	if err == nil {
		defer st.Close()
		if ct, ok := st.GetSecret(secret); ok {
			v, verr := vault.Open(st)
			if verr == nil {
				if pt, derr := v.Decrypt(ct, []byte(secret)); derr == nil && strings.TrimSpace(string(pt)) != "" {
					return string(pt), nil
				}
			}
		}
	}
	if k := strings.TrimSpace(os.Getenv(envName)); k != "" {
		return k, nil
	}
	return "", fmt.Errorf("cle %q absente du coffre et %s non defini", secret, envName)
}

func lookupModelOrFail(slug string) (rag.EmbedModel, error) {
	m, ok := rag.LookupEmbedModel(slug)
	if !ok {
		return rag.EmbedModel{}, fmt.Errorf("modele inconnu: %s", slug)
	}
	return m, nil
}

// cliEmbedder : construit l'embedder selon le backend courant (meme logique
// que le runtime) : desktop si CETAS_LITE_EMBED_URL est pose, sinon
// openrouter. --embed-url prime sur l'env pour le backend desktop.
func cliEmbedder(cfg *config.Config, model rag.EmbedModel) (*rag.Embedder, error) {
	backend := cfg.EmbedBackend
	if backend == "" {
		backend = "desktop"
	}
	if backend == "openrouter" {
		key, err := embedKeyForCLI(cfg, "openrouter")
		if err != nil {
			return nil, err
		}
		return rag.NewOpenRouterEmbedder(key, model, nil), nil
	}
	baseURL := cfg.EmbedURL
	if v := embedURLFlag(os.Args); v != "" {
		baseURL = v
	}
	if baseURL == "" {
		return nil, fmt.Errorf("URL desktop absente (--embed-url ou CETAS_LITE_EMBED_URL)")
	}
	key, err := embedKeyForCLI(cfg, "desktop")
	if err != nil {
		return nil, err
	}
	return rag.NewDesktopEmbedder(baseURL, key, model, nil), nil
}

func ragVectorsBuild(cfg *config.Config, slug string) error {
	model, err := lookupModelOrFail(slug)
	if err != nil {
		return err
	}
	emb, err := cliEmbedder(cfg, model)
	if err != nil {
		return err
	}
	backend := cfg.EmbedBackend
	if backend == "" {
		backend = "desktop"
	}
	fmt.Printf("indexation des vecteurs (%s, %d dims, backend=%s)...\n", model.Slug, model.Dims, backend)
	start := time.Now()
	last := 0
	err = rag.BuildVectors(cfg.RagDir, emb, func(done, total int) {
		pct := done * 100 / total
		if pct >= last+10 || done == total {
			fmt.Printf("  %d/%d chunks (%d%%)\n", done, total, pct)
			last = pct
		}
	})
	if err != nil {
		return err
	}
	fmt.Printf("termine en %s : %s\n", time.Since(start).Round(time.Second), rag.VectorFilePath(cfg.RagDir, model.Slug))
	return nil
}

// embedURLFlag : valeur de --embed-url dans les arguments du process.
func embedURLFlag(args []string) string {
	return flagValue(args, "embed-url", "")
}

func ragVectorsStatus(cfg *config.Config) error {
	ix, err := rag.Load(cfg.RagDir)
	if err != nil {
		return fmt.Errorf("chargement corpus: %w", err)
	}
	want := ix.CorpusHash()
	fmt.Printf("corpus: %d chunks\n", ix.Stats().Chunks)
	for _, m := range rag.EmbedModels {
		path := rag.VectorFilePath(cfg.RagDir, m.Slug)
		vs, err := rag.LoadVectorFile(path)
		if err != nil {
			fmt.Printf("  %-38s absent\n", m.Slug)
			continue
		}
		state := "a jour"
		if vs.Hash != want {
			state = "OBSOLETE (corpus modifie)"
		}
		fmt.Printf("  %-38s %d vecteurs, %d dims : %s\n", m.Slug, vs.Count, vs.Dims, state)
	}
	return nil
}

func ragVectorsProbe(cfg *config.Config, query, slug string, top int) error {
	model, err := lookupModelOrFail(slug)
	if err != nil {
		return err
	}
	path := rag.VectorFilePath(cfg.RagDir, model.Slug)
	vs, err := rag.LoadVectorFile(path)
	if err != nil {
		return fmt.Errorf("vecteurs %s indisponibles: %w (lancer `rag-vectors build`)", slug, err)
	}
	emb, err := cliEmbedder(cfg, model)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	qv, err := emb.Embed(ctx, []string{query}, true)
	if err != nil || len(qv) == 0 {
		return fmt.Errorf("embedding impossible: %w", err)
	}
	ix, err := rag.Load(cfg.RagDir)
	if err != nil {
		return fmt.Errorf("chargement corpus: %w", err)
	}
	if vs.Hash != ix.CorpusHash() {
		fmt.Println("attention : vecteurs obsoletes (corpus modifie), scores indicatifs")
	}
	fmt.Printf("requete : %q\n", query)
	for i, h := range vs.TopK(qv[0], top) {
		fmt.Printf("  #%d cos=%.3f  %s\n", i+1, h.Cosine, ix.ChunkRef(h.Idx))
	}
	return nil
}

// flagValue / flagInt : mini-parseur --cle valeur.
func flagValue(args []string, name, def string) string {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "--"+name {
			return args[i+1]
		}
	}
	return def
}

func flagInt(args []string, name string, def int) int {
	var v int
	if _, err := fmt.Sscanf(flagValue(args, name, ""), "%d", &v); err != nil {
		return def
	}
	if v <= 0 {
		return def
	}
	return v
}
