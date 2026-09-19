package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/alias"
)

func logHasReasoning(c *Conversation) bool {
	for _, ev := range c.marshal().Log {
		if _, ok := ev.Delta["reasoning_content"]; ok {
			return true
		}
	}
	return false
}

func TestAgentThinkingOnEnablesReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais", Think: true})
	reqs := sp.requests()
	if len(reqs) == 0 || !reqs[0].EnableReasoning {
		t.Fatal("l'agent avec thinking actif doit activer le raisonnement")
	}
	if reqs[0].ReasoningEffort == "" {
		t.Fatal("effort attendu en mode agent")
	}
}

// TestAgentThinkingOffDisablesReasoning : le bouton Thinking eteint doit
// etre respecte de bout en bout — pas de raisonnement cote provider, pas
// de contenu de raisonnement diffuse, directive systeme "reponse directe".
func TestAgentThinkingOffDisablesReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion privee"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais vite", Think: false})
	reqs := sp.requests()
	if len(reqs) == 0 {
		t.Fatal("aucune requete emise")
	}
	if reqs[0].EnableReasoning {
		t.Fatal("l'agent avec thinking eteint ne doit pas activer le raisonnement provider")
	}
	if logHasReasoning(c) {
		t.Fatal("aucun contenu de raisonnement ne doit etre diffuse quand le thinking est eteint")
	}
}

func TestChatThinkingOffNoReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion privee"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "bonjour"})
	if sp.requests()[0].EnableReasoning {
		t.Fatal("le chat par defaut ne doit pas activer le raisonnement")
	}
	if logHasReasoning(c) {
		t.Fatal("aucun raisonnement ne doit etre diffuse par defaut")
	}
}

func TestChatThinkingOnEmitsReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "bonjour", Think: true})
	if !sp.requests()[0].EnableReasoning {
		t.Fatal("think=true doit activer le raisonnement")
	}
	if !logHasReasoning(c) {
		t.Fatal("le raisonnement doit etre diffuse quand active")
	}
}

func TestResolveEffort(t *testing.T) {
	if got := resolveEffort(false, "court", "default"); got != "" {
		t.Fatalf("effort sans thinking = %q", got)
	}
	if got := resolveEffort(true, "court", "default"); got != "low" {
		t.Fatalf("auto court = %q", got)
	}
	long := make([]rune, 500)
	for i := range long {
		long[i] = 'x'
	}
	if got := resolveEffort(true, string(long), "default"); got != "medium" {
		t.Fatalf("auto long = %q", got)
	}
	if got := resolveEffort(true, "court", "high"); got != "high" {
		t.Fatalf("effort explicite = %q", got)
	}
	// Durcissement : un effort explicite ne survit pas à think=false.
	if got := resolveEffort(false, "court", "high"); got != "" {
		t.Fatalf("effort explicite + think=false = %q, attendu vide", got)
	}
}

func TestResolveAgentEffort(t *testing.T) {
	// Agent : le "Défaut" du composer vaut "low" (la boucle itère déjà,
	// un raisonnement long par itération ralentit l'expérience).
	if got := resolveAgentEffort(true, "default"); got != "low" {
		t.Fatalf("agent defaut = %q, attendu low", got)
	}
	for _, tc := range [][2]string{{"low", "low"}, {"medium", "medium"}, {"high", "high"}} {
		if got := resolveAgentEffort(true, tc[0]); got != tc[1] {
			t.Fatalf("agent effort %q = %q", tc[0], got)
		}
	}
	if got := resolveAgentEffort(false, "high"); got != "" {
		t.Fatalf("agent think=false = %q, attendu vide", got)
	}
}

// Tests Phase 4 : détection des tours mécaniques (thinking allégé).

func TestMechanicalTurn(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"lis le fichier dnsmasq.conf", true},
		{"Lis README.md", true},
		{"montre-moi le dossier src", true},
		{"liste les fichiers", true},
		{"cherche TODO dans le code", true},
		{"read package.json", true},
		{"list files", true},
		{"salut", false},
		{"", false},
		{"lis le fichier et corrige le bug", false},    // verbe d'action
		{"montre le code puis modifie-le", false},      // verbe d'action
		{"cherche l'erreur et fix le problème", false}, // verbe d'action
		{"explique-moi ce fichier", false},             // pas un verbe de lecture
		{strings.Repeat("lis ", 50), false},            // trop long
	}
	for _, tt := range tests {
		if got := mechanicalTurn(tt.in); got != tt.want {
			t.Errorf("mechanicalTurn(%q) = %v, attendu %v", tt.in, got, tt.want)
		}
	}
}

func TestAgentThinkDirective(t *testing.T) {
	// Tour mécanique + thinking + effort par défaut -> directive allégée.
	d := agentThinkDirective(true, "lis le fichier x.txt", "")
	if !strings.Contains(d, "No reasoning needed") {
		t.Errorf("tour mécanique : directive allégée attendue, obtenu %q", d)
	}
	// Effort explicite de l'utilisateur -> directive standard respectée.
	d = agentThinkDirective(true, "lis le fichier x.txt", "high")
	if strings.Contains(d, "No reasoning needed") {
		t.Errorf("effort explicite : la directive standard aurait dû être gardée")
	}
	// Message non mécanique -> directive standard.
	d = agentThinkDirective(true, "corrige le bug de connexion", "")
	if strings.Contains(d, "No reasoning needed") {
		t.Errorf("tâche de code : la directive standard aurait dû être gardée")
	}
	// Thinking désactivé -> réponse directe, jamais de raisonnement.
	d = agentThinkDirective(false, "lis le fichier x.txt", "")
	if !strings.Contains(d, "Answer directly") {
		t.Errorf("thinking off : attendu 'Answer directly', obtenu %q", d)
	}
}
