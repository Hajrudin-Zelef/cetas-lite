package chat

import (
	"errors"
	"path/filepath"
	"testing"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

func testEngine() *Engine {
	return NewEngine(provider.NewRegistry(), nil, nil, nil, "")
}

func TestSpawnAgentValidation(t *testing.T) {
	e := testEngine()
	if _, err := e.SpawnAgent("", TurnInput{Family: "code", Mode: "standard", Text: "x"}); err == nil {
		t.Fatal("utilisateur vide devrait echouer")
	}
	if _, err := e.SpawnAgent("u", TurnInput{Family: "code", Mode: "standard", Text: "  "}); !errors.Is(err, ErrBadMessage) {
		t.Fatalf("message vide: attendu ErrBadMessage, recu %v", err)
	}
	if _, err := e.SpawnAgent("u", TurnInput{Family: "nope", Mode: "x", Text: "x"}); err == nil {
		t.Fatal("famille inconnue devrait echouer")
	}
	// Famille non-agent refusee.
	if _, err := e.SpawnAgent("u", TurnInput{Family: "samagent-nano", Mode: "free", Text: "x"}); err == nil {
		t.Fatal("famille non-agent devrait echouer")
	}
}

func TestSpawnListDeleteAgent(t *testing.T) {
	e := testEngine()
	run, err := e.SpawnAgent("u", TurnInput{Family: "code", Mode: "standard", Text: "bonjour"})
	if err != nil {
		t.Fatalf("spawn: %v", err)
	}
	if run.ID == "" || run.conv == nil {
		t.Fatal("run incomplet")
	}
	// Laisse le tour echoue (pas de provider) se terminer.
	deadline := time.Now().Add(10 * time.Second)
	for run.conv.IsGenerating() && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}
	list := e.ListAgents("u")
	if len(list) != 1 || list[0].ID != run.ID {
		t.Fatalf("liste inattendue: %+v", list)
	}
	if list[0].Status != "done" {
		t.Fatalf("statut attendu done, recu %s", list[0].Status)
	}
	if e.GetAgent("other", run.ID) != nil {
		t.Fatal("un autre utilisateur ne doit pas voir le run")
	}
	if got := e.GetAgent("u", run.ID); got != run {
		t.Fatal("GetAgent devrait retourner le run")
	}
	if !e.DeleteAgent("u", run.ID) {
		t.Fatal("delete devrait reussir")
	}
	if len(e.ListAgents("u")) != 0 {
		t.Fatal("la liste devrait etre vide apres suppression")
	}
	if e.GetAgent("u", run.ID) != nil {
		t.Fatal("le run supprime ne devrait plus exister")
	}
}

func TestMessageAgentBusy(t *testing.T) {
	e := testEngine()
	run, err := e.SpawnAgent("u", TurnInput{Family: "code", Mode: "standard", Text: "t1"})
	if err != nil {
		t.Fatal(err)
	}
	// Second message immediat : le premier tour tourne encore (ou vient de finir).
	err = e.MessageAgent("u", run.ID, TurnInput{Text: "t2"})
	if err != nil && !errors.Is(err, ErrBusy) {
		t.Fatalf("erreur inattendue: %v", err)
	}
	e.StopAgent("u", run.ID)
	e.DeleteAgent("u", run.ID)
}

func TestWorktreeRepoTracking(t *testing.T) {
	e := testEngine()
	e.noteWorktreeRepo("conv1", "/repo/x")
	if got := e.takeWorktreeRepo("conv1"); got != "/repo/x" {
		t.Fatalf("recu %q", got)
	}
	if got := e.takeWorktreeRepo("conv1"); got != "" {
		t.Fatalf("devrait etre consomme, recu %q", got)
	}
}

func TestAgentForConv(t *testing.T) {
	e := testEngine()
	run, err := e.SpawnAgent("u", TurnInput{Family: "code", Mode: "standard", Text: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if got := e.agentForConv(run.convID); got != run {
		t.Fatal("agentForConv devrait retrouver le run")
	}
	run.setWorktreePath("/wt/abc")
	if run.WorktreePath() != "/wt/abc" {
		t.Fatal("worktree path non enregistre")
	}
	e.DeleteAgent("u", run.ID)
}

func TestDeleteUnknownAgentStrict(t *testing.T) {
	e := testEngine()
	if e.DeleteAgent("u", "inexistant") {
		t.Fatal("supprimer un agent inconnu devrait retourner false")
	}
	if e.DeleteAgent("u", "agent:autre") {
		t.Fatal("supprimer un id hors format devrait retourner false")
	}
}

// TestAgentRestoreAfterRestart prouve le cycle complet : creation d'un
// agent, "redemarrage" (nouveau moteur sur le meme store, memoire vide),
// puis restauration et reprise par message. Regression du bug
// "agent introuvable" constate apres redemarrage du service.
func TestAgentRestoreAfterRestart(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	st, err := store.Open(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	e1 := NewEngine(provider.NewRegistry(), nil, st, nil, "")
	run, err := e1.SpawnAgent("u", TurnInput{Family: "code", Mode: "standard", Text: "bonjour"})
	if err != nil {
		t.Fatalf("spawn: %v", err)
	}
	// La persistance a lieu des le debut du tour (StartTurn), pas
	// seulement a sa fin.
	if _, ok := st.GetConversation("u", "agent:"+run.ID); !ok {
		t.Fatal("l'agent devrait etre persiste des le demarrage du tour")
	}
	// Laisse le tour (voue a l'echec sans provider) se terminer.
	deadline := time.Now().Add(10 * time.Second)
	for run.conv.IsGenerating() && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}
	if err := st.Close(); err != nil {
		t.Fatal(err)
	}
	// "Redemarrage" : nouveau moteur, meme store, memoire vide.
	st2, err := store.Open(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer st2.Close()
	e2 := NewEngine(provider.NewRegistry(), nil, st2, nil, "")
	got := e2.GetAgent("u", run.ID)
	if got == nil {
		t.Fatal("l'agent devrait etre restaure depuis le store apres redemarrage")
	}
	if !got.restored {
		t.Error("le run restaure devrait etre marque restored")
	}
	if err := e2.MessageAgent("u", run.ID, TurnInput{Text: "suite"}); err != nil {
		t.Fatalf("MessageAgent apres restauration: %v", err)
	}
	found := false
	for _, s := range e2.ListAgents("u") {
		if s.ID == run.ID {
			found = true
		}
	}
	if !found {
		t.Error("l'agent restaure devrait apparaitre dans ListAgents")
	}
}
