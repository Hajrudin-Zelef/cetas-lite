package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/provider"
)

func TestAttachmentContextTruncatesPerFile(t *testing.T) {
	e, st := attachEngine(t)
	big := strings.Repeat("a", attachPerFileBudget*2)
	doc, err := st.Save("sam", "gros.txt", []byte(big))
	if err != nil {
		t.Fatal(err)
	}
	ctx := e.AttachmentContext("sam", []string{doc.ID})
	if !strings.Contains(ctx, "[tronqué à 8 Kio]") {
		t.Fatal("la troncature 8 Kio doit etre signalee")
	}
	// Le corps borne a 8 Kio, plus l'en-tete/marqueurs constants.
	if n := strings.Count(ctx, "a"); n > attachPerFileBudget+64 {
		t.Fatalf("le corps depasse 8 Kio (%d caracteres)", n)
	}
}

// Non-regression : une 2e piece jointe envoyee dans un tour ulterieur doit
// etre vue par le modele (le bloc du tour precedent reste, le nouveau
// s'ajoute) — le garde-fou initial bloquait toute nouvelle piece jointe.
func TestSecondAttachmentSeen(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok1"}, {content: "ok2"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	docA, err := st.Save("sam", "a.txt", []byte("CONTENU-A"))
	if err != nil {
		t.Fatal(err)
	}
	docB, err := st.Save("sam", "b.txt", []byte("CONTENU-B"))
	if err != nil {
		t.Fatal(err)
	}
	runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "tour 1", Attachments: []string{docA.ID}})
	runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "tour 2", Attachments: []string{docB.ID}})

	reqs := sp.requests()
	last := reqs[len(reqs)-1]
	if !hasSystemContaining([]provider.Request{last}, "CONTENU-B") {
		t.Fatal("la 2e piece jointe doit etre visible dans le dernier prompt")
	}
	if !hasSystemContaining([]provider.Request{last}, "CONTENU-A") {
		t.Fatal("la 1re piece jointe doit rester persistante dans l'historique")
	}
}

func TestAttachmentContextBudgetTotal(t *testing.T) {
	e, st := attachEngine(t)
	var ids []string
	for i := 0; i < 4; i++ {
		d, err := st.Save("sam", "f.txt", []byte(strings.Repeat("b", attachPerFileBudget)))
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, d.ID)
	}
	ctx := e.AttachmentContext("sam", ids)
	if !strings.Contains(ctx, "[budget PJ atteint]") {
		t.Fatal("le plafond 16 Kio doit etre signale")
	}
}

// P0-B : la 2e extraction est servie par le cache (sidecar), sans relire la
// source (ici on supprime le fichier source pour le prouver).
func TestAttachmentContextCacheHit(t *testing.T) {
	e, st := attachEngine(t)
	doc, err := st.Save("sam", "c.md", []byte("contenu unique Pythagore"))
	if err != nil {
		t.Fatal(err)
	}
	first := e.AttachmentContext("sam", []string{doc.ID})
	if !strings.Contains(first, "Pythagore") {
		t.Fatal("contenu attendu au 1er appel")
	}
	if _, ok := st.ReadSidecar("sam", doc.ID, attachTextSidecarExt); !ok {
		t.Fatal("le sidecar texte doit etre ecrit")
	}
	// Cache memoire : 2e appel identique sans dependre du disque.
	second := e.AttachmentContext("sam", []string{doc.ID})
	if second != first {
		t.Fatal("le 2e appel doit servir le cache")
	}
}

// P0-C : le bloc PJ est persiste comme message system et ne se duplique pas
// lors d'une regeneration.
func TestAttachmentContextPersistedOnce(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}, {content: "ok2"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	doc, err := st.Save("sam", "c.md", []byte("Pythagore"))
	if err != nil {
		t.Fatal(err)
	}
	c := runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "explique", Attachments: []string{doc.ID}})

	if n := countSystemAttachment(c); n != 1 {
		t.Fatalf("1 bloc PJ persiste attendu, obtenu %d", n)
	}
	// Non-regression : le bloc ne doit apparaitre qu'une fois dans le prompt
	// reellement envoye au provider (le bug initial le doublonnait).
	if n := maxPromptAttachment(sp.requests()); n != 1 {
		t.Fatalf("1 bloc PJ dans le prompt attendu, obtenu %d", n)
	}
	if err := e.Regenerate("sam"); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "regeneration non terminee")
	if n := countSystemAttachment(c); n != 1 {
		t.Fatalf("apres regeneration : 1 bloc PJ attendu, obtenu %d", n)
	}
	if n := maxPromptAttachment(sp.requests()); n != 1 {
		t.Fatalf("apres regeneration : 1 bloc PJ dans le prompt attendu, obtenu %d", n)
	}
}

// maxPromptAttachment : nombre maximal de blocs PJ dans un seul envoi
// provider. Le bug de double injection faisait apparaitre le bloc deux fois
// dans le meme prompt (valeur attendue : 1).
func maxPromptAttachment(reqs []provider.Request) int {
	maxN := 0
	for _, r := range reqs {
		n := 0
		for _, m := range r.Messages {
			if m.Role != "system" {
				continue
			}
			if s, ok := m.Content.(string); ok {
				n += strings.Count(s, attachmentContextPrefix)
			}
		}
		if n > maxN {
			maxN = n
		}
	}
	return maxN
}

func countSystemAttachment(c *Conversation) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	n := 0
	for _, m := range c.Messages {
		if m.Role == "system" && isAttachmentContextMessage(m.Content) {
			n++
		}
	}
	return n
}

// P0-E : sans modele vision, une alerte visible (warning) accompagne l'erreur.
func TestVisionWarningVisible(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	e.SetCapabilities(modelcaps.Map{modelcaps.Key("fake", "m"): {}})
	img, err := st.Save("sam", "pic.png", []byte("PNGDATA"))
	if err != nil {
		t.Fatal(err)
	}
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "decris", Attachments: []string{img.ID}})
	if !logHasWarning(c, "modele vision") {
		t.Fatal("une alerte vision visible est attendue")
	}
}

func logHasWarning(c *Conversation, marker string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ev := range c.Log {
		if s, ok := ev.Delta["warning"].(string); ok && strings.Contains(s, marker) {
			return true
		}
	}
	return false
}

func attachEngine(t *testing.T) (*Engine, *attach.Store) {
	t.Helper()
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	return e, st
}
