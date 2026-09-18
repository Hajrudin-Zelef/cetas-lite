package chat

import (
	"strings"
	"testing"
)

// L'annonce "workspace" en tête de fil doit toujours dire OÙ l'agent
// travaille : un run sans projet affiche "Espace partagé", jamais le nom
// d'un projet affiché ailleurs dans l'UI.
func TestAgentWorkspaceLabel_Shared(t *testing.T) {
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	label, mode := agentWorkspaceLabel(TurnInput{}, sb, "")
	if label != "Espace partagé" || mode != "espace partagé" {
		t.Fatalf("reçu %q / %q", label, mode)
	}
}

func TestAgentWorkspaceLabel_Project(t *testing.T) {
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	in := TurnInput{ProjectID: "proj-123"}
	label, mode := agentWorkspaceLabel(in, sb, "CETAS")
	if label != "CETAS" || mode != "projet local" {
		t.Fatalf("reçu %q / %q", label, mode)
	}
}

func TestAgentWorkspaceLabel_ProjectSansNom(t *testing.T) {
	root := t.TempDir()
	sb, err := NewSandbox(root)
	if err != nil {
		t.Fatal(err)
	}
	in := TurnInput{ProjectID: "proj-123"}
	label, mode := agentWorkspaceLabel(in, sb, "")
	if label == "" || label == "Espace partagé" {
		t.Fatalf("le label devrait replier sur la racine, reçu %q", label)
	}
	if mode != "projet local" {
		t.Fatalf("mode reçu %q", mode)
	}
}

// Règle dure : un message trivial (salut, ok, bien, merci) ne doit JAMAIS
// déclencher d'appels d'outils — c'est ce qui faisait "partir en vrille"
// l'agent pour un simple bonjour.
func TestAgentSystemPrompt_TrivialRule(t *testing.T) {
	p := agentSystemPrompt()
	if !strings.Contains(p, "RULE 1") {
		t.Fatal("le prompt agent devrait contenir la règle RULE 1 (messages triviaux)")
	}
	if !strings.Contains(p, "DO NOT call any tool") {
		t.Fatal("la règle triviale devrait interdire explicitement les appels d'outils")
	}
	for _, w := range []string{"salut", "bien", "merci"} {
		if !strings.Contains(p, w) {
			t.Fatalf("la règle triviale devrait citer %q en exemple", w)
		}
	}
}

// La clause "pas d'outils pour les messages triviaux" doit exister à TOUS
// les niveaux d'effort (avant : seulement en effort faible, donc inactive
// pour le réglage par défaut sur les messages courts).
func TestThinkDirective_TrivialNoToolsAllEfforts(t *testing.T) {
	for _, effort := range []string{"", "low", "medium", "high"} {
		d := thinkDirective(true, true, effort)
		if !strings.Contains(d, "NO tool calls") {
			t.Fatalf("effort %q : la directive agent devrait interdire les outils pour les messages triviaux", effort)
		}
	}
}
