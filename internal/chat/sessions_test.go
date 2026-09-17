package chat

// Tests du modèle de sessions : isolation, persistance, suppression
// définitive (anti-résurrection), migration depuis l'ancien format.

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

func newSessionEngine(t *testing.T) (*Engine, *store.Store) {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "bonjour"}}
	reg := provider.NewRegistry()
	reg.Set(fp)
	fams := codeFamily(alias.Member{Provider: "fake", Model: "ok"})
	return NewEngine(reg, fams, st, nil, t.TempDir()), st
}

func turnInput(text string) TurnInput {
	return TurnInput{User: "sam", Family: "code", Mode: "standard", Text: text}
}

func TestSessionCreateAndList(t *testing.T) {
	e, _ := newSessionEngine(t)
	id := e.CreateSession("sam")
	if id == "" {
		t.Fatal("id de session vide")
	}
	if cur := e.CurrentSessionID("sam"); cur != id {
		t.Fatalf("session courante = %q, want %q", cur, id)
	}
	// La session vide apparaît aussitôt dans la liste.
	list := e.ListSessions("sam")
	if len(list) != 1 || list[0].ID != id {
		t.Fatalf("liste = %+v", list)
	}
	if list[0].Title != "(sans titre)" {
		t.Fatalf("titre = %q", list[0].Title)
	}

	runTurn(t, e, "sam", turnInput("salut"))
	list = e.ListSessions("sam")
	if len(list) != 1 {
		t.Fatalf("liste = %+v", list)
	}
	if list[0].Title != "salut" {
		t.Fatalf("titre = %q, want %q", list[0].Title, "salut")
	}
	if list[0].Messages != 2 {
		t.Fatalf("messages = %d, want 2", list[0].Messages)
	}
	if list[0].CreatedAt == 0 || list[0].UpdatedAt == 0 {
		t.Fatalf("dates manquantes: %+v", list[0])
	}
}

func TestSessionIsolation(t *testing.T) {
	e, _ := newSessionEngine(t)

	// Isolation entre utilisateurs.
	idA := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("secret de sam"))
	if got := e.ListSessions("bob"); len(got) != 0 {
		t.Fatalf("bob voit les sessions de sam: %+v", got)
	}
	if e.OpenSession("bob", idA) {
		t.Fatal("bob ne doit pas pouvoir ouvrir la session de sam")
	}

	// Isolation entre sessions du même utilisateur.
	idB := e.CreateSession("sam")
	if len(e.Conversation("sam").MessagesSnapshot()) != 0 {
		t.Fatal("la nouvelle session doit être vierge")
	}
	if !e.OpenSession("sam", idA) {
		t.Fatal("ouverture session A attendue")
	}
	if len(e.Conversation("sam").MessagesSnapshot()) == 0 {
		t.Fatal("les messages de la session A doivent être restaurés")
	}
	if !e.OpenSession("sam", idB) {
		t.Fatal("ouverture session B attendue")
	}
	if len(e.Conversation("sam").MessagesSnapshot()) != 0 {
		t.Fatal("la session B ne doit pas voir les messages de A")
	}
	if e.OpenSession("sam", "inexistante") {
		t.Fatal("ouverture d'une session inconnue doit échouer")
	}
}

func TestSessionListOrder(t *testing.T) {
	e, _ := newSessionEngine(t)
	e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("premiere"))
	idB := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("seconde"))
	list := e.ListSessions("sam")
	if len(list) != 2 {
		t.Fatalf("liste = %+v", list)
	}
	if list[0].ID != idB || list[0].Title != "seconde" {
		t.Fatalf("ordre inattendu: %+v", list)
	}
}

func TestSessionOpenEmitsReset(t *testing.T) {
	e, _ := newSessionEngine(t)
	idA := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("un"))
	idB := e.CreateSession("sam")
	_ = idB

	c := e.Conversation("sam")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ready := make(chan struct{}, 1)
	gotReset := make(chan struct{}, 1)
	go c.Subscribe(ctx, 0, func(ev map[string]any) bool {
		if _, ok := ev["caught_up"]; ok {
			select {
			case ready <- struct{}{}:
			default:
			}
		}
		if r, ok := ev["reset"].(bool); ok && r {
			select {
			case gotReset <- struct{}{}:
			default:
			}
			return false
		}
		return true
	})
	select {
	case <-ready:
	case <-time.After(2 * time.Second):
		t.Fatal("abonnement non prêt")
	}

	if !e.OpenSession("sam", idA) {
		t.Fatal("ouverture attendue")
	}
	select {
	case <-gotReset:
	case <-time.After(2 * time.Second):
		t.Fatal("événement reset attendu après ouverture de session")
	}
	// Le client reconnecte à 0 et rejoue le journal de la session A.
	if len(e.Conversation("sam").MessagesSnapshot()) == 0 {
		t.Fatal("messages de la session A attendus après ouverture")
	}
}

func TestSessionDelete(t *testing.T) {
	e, _ := newSessionEngine(t)
	idA := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("a garder"))
	idB := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("a jeter"))

	cur, ok := e.DeleteSession("sam", idB)
	if !ok {
		t.Fatal("suppression attendue")
	}
	if cur != e.CurrentSessionID("sam") {
		t.Fatal("currentId incohérent")
	}
	list := e.ListSessions("sam")
	if len(list) != 1 || list[0].ID != idA {
		t.Fatalf("liste après suppression = %+v", list)
	}
	if _, ok := e.DeleteSession("sam", idB); ok {
		t.Fatal("seconde suppression doit échouer")
	}
	if e.OpenSession("sam", idB) {
		t.Fatal("ouverture d'une session supprimée doit échouer")
	}
	// La session restante est intacte.
	if !e.OpenSession("sam", idA) || len(e.Conversation("sam").MessagesSnapshot()) == 0 {
		t.Fatal("la session conservée doit rester ouvrable avec ses messages")
	}
}

func TestSessionDeleteCurrent(t *testing.T) {
	e, _ := newSessionEngine(t)
	id := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("salut"))

	cur, ok := e.DeleteSession("sam", id)
	if !ok {
		t.Fatal("suppression attendue")
	}
	if cur == "" || cur == id {
		t.Fatalf("nouvelle session courante attendue, got %q", cur)
	}
	if got := e.CurrentSessionID("sam"); got != cur {
		t.Fatalf("courante = %q, want %q", got, cur)
	}
	if len(e.Conversation("sam").MessagesSnapshot()) != 0 {
		t.Fatal("la remplaçante doit être vierge")
	}
	for _, s := range e.ListSessions("sam") {
		if s.ID == id {
			t.Fatal("la session supprimée ne doit plus être listée")
		}
	}
}

// TestSessionDeleteTombstone : une persistance tardive d'un tour qui était
// en vol au moment de la suppression ne doit PAS ressusciter la session.
func TestSessionDeleteTombstone(t *testing.T) {
	e, _ := newSessionEngine(t)
	id := e.CreateSession("sam")
	runTurn(t, e, "sam", turnInput("salut"))

	// Fantôme : objet conversation portant encore l'ancien ID, avec du
	// contenu (simule le tour en vol qui persiste après la suppression).
	ghost := NewConversation(id, e, nil)
	ghost.load(snapshot{
		ID:       id,
		Messages: []provider.Message{{Role: "user", Content: "fantome"}},
	})

	if _, ok := e.DeleteSession("sam", id); !ok {
		t.Fatal("suppression attendue")
	}
	e.saveSession("sam", ghost)

	if got := e.ListSessions("sam"); len(got) != 0 {
		t.Fatalf("résurrection ! liste = %+v", got)
	}
	if e.OpenSession("sam", id) {
		t.Fatal("la session supprimée ne doit pas être ouvrable")
	}
}

// TestSessionMigration : l'ancien format (bucket archives + clé "active")
// est importé une fois en sessions, sans doublon au second appel.
func TestSessionMigration(t *testing.T) {
	e, st := newSessionEngine(t)

	mkSnap := func(id, text string, ts int64) []byte {
		s := snapshot{
			ID:       id,
			Messages: []provider.Message{{Role: "user", Content: text}},
			Log:      []LogEvent{{Seq: 1, TS: ts, Delta: map[string]any{"user": text}}},
			Seq:      1,
		}
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		return data
	}
	// Archive ancienne (clé horodatée).
	if err := st.ArchiveConversation("sam", "cid-archive", mkSnap("cid-archive", "question archive", 1700000000000)); err != nil {
		t.Fatal(err)
	}
	// Ancienne conversation active.
	if err := st.PutConversation("sam", "active", mkSnap("cid-live", "question live", 1700000001000)); err != nil {
		t.Fatal(err)
	}

	list := e.ListSessions("sam")
	if len(list) != 2 {
		t.Fatalf("liste migrée = %+v", list)
	}
	byTitle := map[string]SessionInfo{}
	for _, s := range list {
		byTitle[s.Title] = s
	}
	if _, ok := byTitle["question archive"]; !ok {
		t.Fatalf("archive non migrée: %+v", list)
	}
	live, ok := byTitle["question live"]
	if !ok {
		t.Fatalf("conversation active non migrée: %+v", list)
	}
	if cur := e.CurrentSessionID("sam"); cur != live.ID {
		t.Fatalf("courante = %q, want %q", cur, live.ID)
	}

	// Idempotence : second appel ne duplique rien.
	if got := e.ListSessions("sam"); len(got) != 2 {
		t.Fatalf("migration non idempotente: %+v", got)
	}
	// L'ancienne clé "active" a disparu.
	if _, ok := st.GetConversation("sam", "active"); ok {
		t.Fatal("la clé active aurait dû être supprimée après migration")
	}
}

func TestSessionMigrationEmptyStore(t *testing.T) {
	e, _ := newSessionEngine(t)
	if got := e.ListSessions("zoe"); len(got) != 0 {
		t.Fatalf("liste vide attendue, got %+v", got)
	}
}
