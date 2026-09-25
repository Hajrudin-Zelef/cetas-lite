package rag

import (
	"testing"
)

func TestRelatedTopicQueries_EntitiesFromQuery(t *testing.T) {
	got := RelatedTopicQueries("je veux parler de GLM-5.3", Result{}, 3)
	// Note : minTermLen=2 mange les chiffres isoles (« 5 », « 3 ») —
	// comportement existant du nettoyage, identique cote reutilisation.
	if len(got) != 1 || got[0] != "glm" {
		t.Fatalf("entite attendue [glm], obtenu %v", got)
	}
}

func TestRelatedTopicQueries_ExcludesCurrentQuery(t *testing.T) {
	// L'entite nettoyee egale la requete courante nettoyee : deja
	// recuperee au tour en cours, rien a anticiper.
	got := RelatedTopicQueries("GLM-5.3", Result{}, 3)
	if len(got) != 0 {
		t.Fatalf("requete courante exclue attendue, obtenu %v", got)
	}
}

func TestRelatedTopicQueries_KeywordsFromHits(t *testing.T) {
	res := Result{Hits: []Hit{
		{Title: "Test", Keywords: []string{"NVIDIA RTX 5090", "refroidissement liquide"}},
	}}
	got := RelatedTopicQueries("parle-moi des GPU", res, 5)
	want := map[string]bool{"gpu": true, "nvidia rtx 5090": true, "refroidissement liquide": true}
	if len(got) != len(want) {
		t.Fatalf("attendu %v, obtenu %v", want, got)
	}
	for _, q := range got {
		if !want[q] {
			t.Fatalf("candidat inattendu %q dans %v", q, got)
		}
	}
}

func TestRelatedTopicQueries_TitleEntitiesFallback(t *testing.T) {
	// Sans mots-cles (chunks en mode brut), les entites des titres
	// servent de repli.
	res := Result{Hits: []Hit{{Title: "Kimi K3 : benchmark inference"}}}
	got := RelatedTopicQueries("modeles ouverts", res, 5)
	found := false
	for _, q := range got {
		// « Kimi » en debut de titre n'est pas une entite (regle
		// anti-debut-de-phrase) ; « K3 » (chiffre) l'est.
		if q == "k3" {
			found = true
		}
	}
	if !found {
		t.Fatalf("entite du titre attendue (k3), obtenu %v", got)
	}
}

func TestRelatedTopicQueries_Dedupe(t *testing.T) {
	res := Result{Hits: []Hit{{Title: "x", Keywords: []string{"Kimi K3"}}}}
	got := RelatedTopicQueries("GLM-5.3 vs Kimi K3", res, 5)
	seen := map[string]int{}
	for _, q := range got {
		seen[q]++
	}
	for q, n := range seen {
		if n > 1 {
			t.Fatalf("doublon %q dans %v", q, got)
		}
	}
}

func TestRelatedTopicQueries_Cap(t *testing.T) {
	res := Result{Hits: []Hit{
		{Title: "a", Keywords: []string{"kw1", "kw2", "kw3", "kw4", "kw5"}},
	}}
	got := RelatedTopicQueries("sujet", res, 2)
	if len(got) > 2 {
		t.Fatalf("borne max=2 depassee : %v", got)
	}
}

func TestRelatedTopicQueries_InvalidMax(t *testing.T) {
	if got := RelatedTopicQueries("GLM-5.3", Result{}, 0); got != nil {
		t.Fatalf("max<1 => nil attendu, obtenu %v", got)
	}
}

func TestRelatedTopicQueries_Empty(t *testing.T) {
	// Ni entite ni hit : rien a anticiper (fail-open).
	if got := RelatedTopicQueries("bonjour", Result{}, 3); len(got) != 0 {
		t.Fatalf("vide attendu, obtenu %v", got)
	}
}

func TestCleanQueryKey_ConsistentWithSearch(t *testing.T) {
	// La cle doit etre stable et identique des deux cotes
	// (anticipation et reutilisation).
	a := CleanQueryKey("Parle-moi de GLM-5.3 !")
	b := CleanQueryKey("parle-moi de glm-5.3")
	if a == "" || a != b {
		t.Fatalf("cles instables : %q vs %q", a, b)
	}
	if a != cleanQueryForSearch("Parle-moi de GLM-5.3 !") {
		t.Fatalf("cle differente du nettoyage de recherche : %q", a)
	}
}
