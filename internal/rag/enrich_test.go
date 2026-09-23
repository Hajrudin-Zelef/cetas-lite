package rag

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
)

// TestEnrichQueryWithHistory : l'iteration 6 reprend le sujet de
// l'historique quand la requete est elliptique (pronom ou tres courte),
// sans toucher aux requetes deja porteuses d'entites. Fail-open partout.
func TestEnrichQueryWithHistory(t *testing.T) {
	histUser := []string{"Qu'est-ce que Kimi K3 et que vaut-il ?"}
	histFull := []string{
		"Qu'est-ce que Kimi K3 et que vaut-il ?",
		"Kimi K3 est le modèle phare open-weight de Moonshot AI sorti en 2026.",
	}
	cases := []struct {
		name    string
		query   string
		history []string
		// want contient une entite attendue dans le resultat ;
		// vide => resultat strictement inchange.
		want string
	}{
		{
			"pronom elliptique reprend le sujet",
			"oui mais on peut la faire tourner sur combien de cpu ?",
			histUser, "Kimi K3",
		},
		{
			"requete courte reprend le sujet",
			"et le prix ?",
			histUser, "Kimi K3",
		},
		{
			"reponse assistant aussi exploitable",
			"on peut la faire tourner sur combien de cpu ?",
			histFull, "K3",
		},
		{"requete deja porteuse d'entite", "parle-moi de Kimi K3", histFull, ""},
		{"identifiant seul", "GLM-5.3", histFull, ""},
		{"camelCase = entite", "parle-moi de vLLM", histFull, ""},
		{"entite seule", "Kimi", histFull, ""},
		{
			"historique vide",
			"on peut la faire tourner sur combien de cpu ?",
			nil, "",
		},
		{
			"historique sans entite",
			"et le prix ?",
			[]string{"bonjour", "salut, ça va ?"}, "",
		},
		{
			"non elliptique sans entite",
			"quel est ton film préféré ?",
			histFull, "",
		},
	}
	for _, c := range cases {
		got, keys := EnrichQueryWithHistory(c.query, c.history)
		if c.want == "" {
			if got != c.query {
				t.Errorf("%s: EnrichQueryWithHistory(%q) = %q, attendu inchangé", c.name, c.query, got)
			}
			if len(keys) != 0 {
				t.Errorf("%s: entites inattendues %v (requete inchangee)", c.name, keys)
			}
			continue
		}
		if !strings.Contains(got, c.want) {
			t.Errorf("%s: EnrichQueryWithHistory(%q) = %q, attendu contenant %q", c.name, c.query, got, c.want)
		}
		if !strings.HasPrefix(got, c.query) {
			t.Errorf("%s: la requete d'origine doit etre preservee en tete, obtenu %q", c.name, got)
		}
		if len(keys) == 0 {
			t.Errorf("%s: aucune cle d'entite pour une requete enrichie", c.name)
		}
		for _, k := range keys {
			if k != normalize(k) || strings.ContainsAny(k, " .,-") {
				t.Errorf("%s: cle non normalisee %q (doit matcher tokenize)", c.name, k)
			}
		}
	}
}

// TestEnrichEntitiesKeys : les cles de boost sont les tokens post-
// normalisation des groupes ajoutes (meme forme que tokenize), dedupliques.
func TestEnrichEntitiesKeys(t *testing.T) {
	got, keys := EnrichQueryWithHistory(
		"on peut la faire tourner sur combien de cpu ?",
		[]string{"Qu'est-ce que Kimi K3 et que vaut-il ?"},
	)
	if !strings.HasSuffix(got, "Kimi K3") {
		t.Fatalf("requete: %q", got)
	}
	want := []string{"kimi", "k3"}
	if len(keys) != len(want) {
		t.Fatalf("cles = %v, attendu %v", keys, want)
	}
	for i, k := range keys {
		if k != want[i] {
			t.Fatalf("cles = %v, attendu %v", keys, want)
		}
	}
}

// TestEnrichQueryNoDuplicate : un sujet deja present dans la requete
// n'est pas re-ajoute.
func TestEnrichQueryNoDuplicate(t *testing.T) {
	got, _ := EnrichQueryWithHistory(
		"la version K3 consomme combien ?",
		[]string{"Qu'est-ce que Kimi K3 ?"},
	)
	if strings.Count(strings.ToLower(got), "k3") > 1 {
		t.Fatalf("sujet duplique: %q", got)
	}
}

// TestEnrichQueryRanksKimiFirst : regression du cas capture le 2026-09-23
// sur un fixture reproduisant la structure df du vrai corpus (iteration 6b).
// Plusieurs docs mentionnent « kimi » (idf faible) ; un doc CPU fort porte
// les termes conversationnels rares. Sans boost, le doc CPU domine
// (comme RTX Spark 12.99 > Kimi 8.86 sur le corpus reel) — le test
// echouerait s'il s'appuyait sur Search seul. Avec SearchBoosted, le doc
// Kimi passe en #1.
func TestEnrichQueryRanksKimiFirst(t *testing.T) {
	root := t.TempDir()
	// Docs mentionnant kimi : df eleve -> idf faible.
	writeFile(t, filepath.Join(root, "c1", "kimi-a.md"),
		"# Kimi K3\n\nKimi K3 est un modele open-weight de Moonshot AI.\n")
	writeFile(t, filepath.Join(root, "c1", "kimi-b.md"),
		"# Poids Kimi\n\nTelechargement des poids Kimi K3 et licence Moonshot.\n")
	writeFile(t, filepath.Join(root, "c1", "kimi-c.md"),
		"# Tutoriel Kimi\n\nInstallation de Kimi K3 depuis Hugging Face.\n")
	writeFile(t, filepath.Join(root, "c1", "moonshot.md"),
		"# Moonshot AI\n\nLa societe Moonshot AI derriere Kimi.\n")
	// Doc CPU fort : termes conversationnels rares, df=1 -> idf eleve.
	writeFile(t, filepath.Join(root, "c1", "cpu.md"),
		"# Tourner un modele\n\nUn modele 120B tourne localement sur CPU grand public : combien de CPU faut-il pour le faire tourner en FP4 ?\n")
	// Doc cible : sujet + cpu.
	writeFile(t, filepath.Join(root, "c1", "kimi-cpu.md"),
		"# Kimi K3 sur CPU\n\nDeploiement de Kimi K3 : un CPU unique ne suffit pas, il faut un serveur multi-socket.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	query, ents := EnrichQueryWithHistory(
		"oui mais on peut la faire tourner sur combien de cpu ?",
		[]string{"Qu'est-ce que Kimi K3 et que vaut-il ?"},
	)
	if !strings.Contains(query, "Kimi K3") || len(ents) == 0 {
		t.Fatalf("requete non enrichie: %q (entites %v)", query, ents)
	}
	ctx := context.Background()
	// 1) Sans boost : le doc CPU conversationnel domine (fixture non
	// vacuole — c'est exactement l'echec mesure sur le corpus reel).
	plain := ix.Search(ctx, query, 5)
	if len(plain.Hits) == 0 {
		t.Fatal("aucun hit sans boost")
	}
	if top := plain.Hits[0].Path; strings.HasSuffix(top, "kimi-cpu.md") {
		t.Fatalf("fixture vacuole: Search sans boost classe deja kimi-cpu.md en #1 (%s)", top)
	}
	// 2) Avec boost : le doc Kimi passe en #1.
	boost := map[string]float64{}
	for _, e := range ents {
		boost[e] = 3.0
	}
	boosted := ix.SearchBoosted(ctx, query, boost, 5)
	if len(boosted.Hits) == 0 {
		t.Fatal("aucun hit avec boost")
	}
	if top := boosted.Hits[0].Path; !strings.HasSuffix(top, "kimi-cpu.md") {
		t.Fatalf("boost: top = %q, attendu kimi-cpu.md (requete %q, entites %v)", top, query, ents)
	}
	// 3) Non-regression : boost nil => strictement identique a Search.
	ni := ix.SearchBoosted(ctx, query, nil, 5)
	if len(ni.Hits) != len(plain.Hits) {
		t.Fatalf("boost nil: %d hits vs %d", len(ni.Hits), len(plain.Hits))
	}
	for i := range ni.Hits {
		if ni.Hits[i].Path != plain.Hits[i].Path || ni.Hits[i].Score != plain.Hits[i].Score {
			t.Fatalf("boost nil modifie le classement: [%d] %s %.4f vs %s %.4f",
				i, ni.Hits[i].Path, ni.Hits[i].Score, plain.Hits[i].Path, plain.Hits[i].Score)
		}
	}
	// 4) Non-regression : requete directe a entite => non enrichie (pas
	// de boost en jeu), top-1 stable entre Search et SearchBoosted(nil).
	direct, dents := EnrichQueryWithHistory("kimi k3", []string{"Qu'est-ce que Kimi K3 ?"})
	if len(dents) != 0 || direct != "kimi k3" {
		t.Fatalf("requête directe modifiee: %q (entites %v)", direct, dents)
	}
	d1 := ix.Search(ctx, direct, 3)
	if len(d1.Hits) == 0 {
		t.Fatal("aucun hit pour kimi k3")
	}
	if !strings.Contains(strings.ToLower(d1.Hits[0].Path), "kimi") {
		t.Fatalf("top-1 direct = %q, attendu un doc kimi", d1.Hits[0].Path)
	}
}
