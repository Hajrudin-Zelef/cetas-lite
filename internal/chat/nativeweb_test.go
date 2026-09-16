package chat

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/search"
	"cetas-lite/internal/store"
)

type stubWebTools struct{}

func (stubWebTools) Search(ctx context.Context, query string, maxResults int) search.Result {
	return search.Result{}
}
func (stubWebTools) Fetch(ctx context.Context, rawURL string) (string, string, error) {
	return "", "", nil
}

func newNativeWebEngine(t *testing.T) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "native.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	e := NewEngine(provider.NewRegistry(), nil, st, nil, t.TempDir())
	e.SetSearcher(stubWebTools{})
	return e
}

func setWebSearchMode(t *testing.T, e *Engine, user, mode string) {
	t.Helper()
	if err := e.st.PutSetting(user, "ui", []byte(`{"websearch_mode":"`+mode+`"}`)); err != nil {
		t.Fatal(err)
	}
}

func TestWebSearchModeDefaut(t *testing.T) {
	e := newNativeWebEngine(t)
	if got := e.webSearchMode("u"); got != WebSearchAuto {
		t.Fatalf("mode par defaut attendu auto, obtenu %q", got)
	}
}

func TestNativeWebFor(t *testing.T) {
	e := newNativeWebEngine(t)
	cas := []struct {
		user, providerID, model, mode string
		want                          bool
	}{
		{"u", "openrouter", "openai/gpt-5", "auto", true},
		{"u", "openrouter", "openai/gpt-5", "natif", true},
		{"u", "openrouter", "openai/gpt-5", "outils", false},
		{"u", "openrouter", "openai/gpt-5", "off", false},
		{"u", "deepseek", "deepseek-chat", "auto", true},              // natif via enable_search
		{"u", "deepseek", "deepseek-chat", "natif", true},             // idem
		{"u", "deepseek", "deepseek-chat", "outils", false},           // outils forces
		{"u", "deepseek", "deepseek-chat", "off", false},              // coupe
		{"u", "openrouter", "perplexity/sonar:online", "auto", false}, // deja en ligne
	}
	for _, c := range cas {
		setWebSearchMode(t, e, c.user, c.mode)
		if got := e.nativeWebFor(c.user, c.providerID, c.model); got != c.want {
			t.Errorf("nativeWebFor(%s,%s,%s,%s) = %v, voulu %v", c.user, c.providerID, c.model, c.mode, got, c.want)
		}
	}
}

func TestWebToolsFor(t *testing.T) {
	e := newNativeWebEngine(t)
	mkIn := func(web bool) TurnInput { return TurnInput{User: "u", Web: web} }
	setWebSearchMode(t, e, "u", "auto")
	if !e.webToolsFor(mkIn(true)) {
		t.Error("auto+web : outils attendus")
	}
	setWebSearchMode(t, e, "u", "outils")
	if !e.webToolsFor(mkIn(true)) {
		t.Error("outils+web : outils attendus")
	}
	setWebSearchMode(t, e, "u", "natif")
	if e.webToolsFor(mkIn(true)) {
		t.Error("natif : pas d'outils web")
	}
	setWebSearchMode(t, e, "u", "off")
	if e.webToolsFor(mkIn(true)) {
		t.Error("off : pas d'outils web")
	}
	if e.webToolsFor(mkIn(false)) {
		t.Error("web=false : pas d'outils web")
	}
	// Recherche approfondie : les outils sont requis meme en mode "natif"
	// (le natif est desactive pour ce tour), mais pas si le web est coupe.
	setWebSearchMode(t, e, "u", "natif")
	deep := TurnInput{User: "u", Web: true, WebDepth: "deep"}
	if !e.webToolsFor(deep) {
		t.Error("deep+natif : outils attendus (natif desactive pour ce tour)")
	}
	setWebSearchMode(t, e, "u", "off")
	if e.webToolsFor(deep) {
		t.Error("deep+off : pas d'outils web")
	}
}

func TestDropWebTools(t *testing.T) {
	tools := []provider.Tool{
		{Function: provider.ToolFunction{Name: "Read"}},
		{Function: provider.ToolFunction{Name: "web_search"}},
		{Function: provider.ToolFunction{Name: "web_fetch"}},
		{Function: provider.ToolFunction{Name: "Bash"}},
	}
	out := dropWebTools(tools)
	if len(out) != 2 || out[0].Function.Name != "Read" || out[1].Function.Name != "Bash" {
		t.Fatalf("filtrage inattendu: %+v", out)
	}
	if len(tools) != 4 {
		t.Fatal("la liste d'origine ne doit pas etre mutee")
	}
}

func TestSearchSourcesDelta(t *testing.T) {
	d := searchSourcesDelta([]provider.WebAnnotation{
		{URL: "https://a.example", Title: "A"},
	}, true)
	if d["phase"] != "done" || d["native"] != true {
		t.Fatalf("delta inattendu: %+v", d)
	}
	srcs, ok := d["sources"].([]map[string]any)
	if !ok || len(srcs) != 1 || srcs[0]["url"] != "https://a.example" {
		t.Fatalf("sources inattendues: %+v", d["sources"])
	}
}

// --- Moteur unique + repli natif -> outils -----------------------------------

func newChatEngine(t *testing.T, sps []*scriptedProvider, fams []alias.Family) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	for _, sp := range sps {
		reg.Set(sp)
	}
	return NewEngine(reg, fams, st, nil, t.TempDir())
}

func hasNativePlugin(req provider.Request) bool {
	pl, ok := req.Extra["plugins"].([]any)
	if !ok || len(pl) == 0 {
		return false
	}
	m, ok := pl[0].(map[string]any)
	return ok && m["id"] == "web"
}

func stubWebWithHit() *stubWeb {
	return &stubWeb{result: search.Result{Provider: "tavily", Hits: []search.Hit{
		{Title: "T", URL: "https://x.test", Description: "d"},
	}}}
}

// En mode chat + auto, un membre OpenRouter natif ne doit PAS declencher la
// pre-recherche locale (pas de double recherche : latence et cout x2).
func TestChatNatifSansPreRechercheLocale(t *testing.T) {
	stub := stubWebWithHit()
	or := &scriptedProvider{id: "openrouter", steps: []scriptStep{{content: "ok"}}}
	e := newChatEngine(t, []*scriptedProvider{or}, plainFamily(alias.Member{Provider: "openrouter", Model: "m"}))
	e.SetSearcher(stub)
	runTurn(t, e, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "quoi de neuf ?", Web: true})

	if n := atomic.LoadInt32(&stub.searchCalls); n != 0 {
		t.Fatalf("pre-recherche locale inattendue avec le natif (appels=%d)", n)
	}
	reqs := or.requests()
	if len(reqs) != 1 || !hasNativePlugin(reqs[0]) {
		t.Fatalf("plugin natif attendu dans l'unique requete : %+v", reqs)
	}
}

// En mode chat + auto, un membre DeepSeek (natif via enable_search) ne doit
// PAS declencher la pre-recherche locale (pas de double recherche).
func TestChatDeepSeekNatifSansPreRecherche(t *testing.T) {
	stub := stubWebWithHit()
	ds := &scriptedProvider{id: "deepseek", steps: []scriptStep{{content: "ok"}}}
	e := newChatEngine(t, []*scriptedProvider{ds}, plainFamily(alias.Member{Provider: "deepseek", Model: "deepseek-chat"}))
	e.SetSearcher(stub)
	runTurn(t, e, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "quoi de neuf ?", Web: true})

	if n := atomic.LoadInt32(&stub.searchCalls); n != 0 {
		t.Fatalf("pre-recherche locale inattendue avec le natif DeepSeek (appels=%d)", n)
	}
	reqs := ds.requests()
	if len(reqs) != 1 {
		t.Fatalf("une requete attendue : %+v", reqs)
	}
	if reqs[0].Extra["enable_search"] != true {
		t.Fatalf("enable_search=true attendu dans Extra : %+v", reqs[0].Extra)
	}
}

// En mode chat + auto, un membre non natif (ex: fake) utilise la
// pre-recherche locale, sans parametre natif.
func TestChatNonNatifAvecPreRechercheLocale(t *testing.T) {
	stub := stubWebWithHit()
	fk := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newChatEngine(t, []*scriptedProvider{fk}, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetSearcher(stub)
	runTurn(t, e, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "quoi de neuf ?", Web: true})

	if n := atomic.LoadInt32(&stub.searchCalls); n != 1 {
		t.Fatalf("pre-recherche locale attendue (appels=%d)", n)
	}
	reqs := fk.requests()
	if len(reqs) != 1 || hasNativePlugin(reqs[0]) {
		t.Fatalf("aucun plugin natif attendu : %+v", reqs)
	}
	for _, m := range reqs[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "https://x.test") {
				return
			}
		}
	}
	t.Fatal("le contexte de pre-recherche doit etre injecte")
}

// Repli chat : le natif echoue avant toute emission -> une seule
// pre-recherche locale, puis le membre suivant tourne SANS natif.
// Le pool est melange a chaque requete : les deux membres partagent le
// provider "openrouter" (natif supporte), seul l'ordre du tirage varie.
type failFirstProvider struct {
	id   string
	mu   sync.Mutex
	reqs []provider.Request
	once int32
}

func (p *failFirstProvider) ID() string { return p.id }

func (p *failFirstProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	p.mu.Lock()
	p.reqs = append(p.reqs, req)
	p.mu.Unlock()
	if atomic.CompareAndSwapInt32(&p.once, 0, 1) {
		return provider.Response{}, errors.New("natif hs")
	}
	emit(provider.Event{Content: "ok"})
	return provider.Response{Content: "ok"}, nil
}

func (p *failFirstProvider) requests() []provider.Request {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]provider.Request(nil), p.reqs...)
}

func TestChatRepliNatifVersOutils(t *testing.T) {
	stub := stubWebWithHit()
	fp := &failFirstProvider{id: "openrouter"}
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(fp)
	e := NewEngine(reg, plainFamily(
		alias.Member{Provider: "openrouter", Model: "m1"},
		alias.Member{Provider: "openrouter", Model: "m2"},
	), st, nil, t.TempDir())
	e.SetSearcher(stub)
	c := runTurn(t, e, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "quoi de neuf ?", Web: true})

	if n := atomic.LoadInt32(&stub.searchCalls); n != 1 {
		t.Fatalf("repli : 1 pre-recherche locale attendue (appels=%d)", n)
	}
	reqs := fp.requests()
	if len(reqs) != 2 {
		t.Fatalf("2 tentatives attendues, got %d", len(reqs))
	}
	if !hasNativePlugin(reqs[0]) {
		t.Fatalf("le 1er membre tente doit tenter le natif : %+v", reqs[0])
	}
	if hasNativePlugin(reqs[1]) {
		t.Fatalf("le membre de repli ne doit pas reutiliser le natif : %+v", reqs[1])
	}
	for _, m := range reqs[1].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "provider-native") {
				t.Fatal("le membre de repli ne doit plus voir la directive native")
			}
		}
	}
	if txt := logText(c); !strings.Contains(txt, "ok") {
		t.Fatalf("reponse finale attendue, obtenu : %q", txt)
	}
}

// Repli chat : en mode "natif", pas de repli vers les outils.
func TestChatRepliNatifDesactiveEnModeNatif(t *testing.T) {
	stub := stubWebWithHit()
	or := &scriptedProvider{id: "openrouter", steps: []scriptStep{{err: errors.New("natif hs")}}}
	fb := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newChatEngine(t, []*scriptedProvider{or, fb}, plainFamily(
		alias.Member{Provider: "openrouter", Model: "m"},
		alias.Member{Provider: "fake", Model: "m"},
	))
	e.SetSearcher(stub)
	setWebSearchMode(t, e, "sam", WebSearchNative)
	runTurn(t, e, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "quoi de neuf ?", Web: true})

	if n := atomic.LoadInt32(&stub.searchCalls); n != 0 {
		t.Fatalf("mode natif : aucun repli outils attendu (appels=%d)", n)
	}
}

// Repli agent : echec natif avant emission -> on retente avec les outils web
// (web_search reinjecte, plugin natif retire), une seule fois.
func TestAgentRepliNatifVersOutils(t *testing.T) {
	or := &scriptedProvider{id: "openrouter", steps: []scriptStep{
		{err: errors.New("natif hs")},
		{content: "ok"},
	}}
	e := newAgentEngine(t, or, codeFamily(alias.Member{Provider: "openrouter", Model: "m"}))
	e.SetSearcher(&stubWeb{})
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "cherche x", Web: true})

	reqs := or.requests()
	if len(reqs) != 2 {
		t.Fatalf("2 requetes attendues (natif puis repli), obtenu %d", len(reqs))
	}
	if !hasNativePlugin(reqs[0]) {
		t.Fatal("la 1re requete doit utiliser le plugin natif")
	}
	if hasToolName(reqs[0].Tools, "web_search") {
		t.Fatal("la 1re requete ne doit pas exposer web_search (doublon avec le natif)")
	}
	if hasNativePlugin(reqs[1]) {
		t.Fatal("la requete de repli ne doit plus utiliser le natif")
	}
	if !hasToolName(reqs[1].Tools, "web_search") {
		t.Fatal("la requete de repli doit reexposer web_search")
	}
}

// Le repli natif -> outils doit remplacer la directive : le membre suivant
// ne doit plus lire "provider-native search".
func TestReplaceWebDirective(t *testing.T) {
	oldDir := searchDirective(true, true)
	newDir := searchDirective(true, false)
	pm := []provider.Message{
		{Role: "system", Content: "Contexte.\n\n" + oldDir},
		{Role: "user", Content: "salut"},
	}
	out := replaceWebDirective(pm, oldDir, newDir)
	s, _ := out[0].Content.(string)
	if strings.Contains(s, "provider-native") {
		t.Fatal("l'ancienne directive native doit etre remplacee")
	}
	if !strings.Contains(s, "web_search and web_fetch") {
		t.Fatal("la nouvelle directive outils doit etre presente")
	}
	if u, _ := out[1].Content.(string); u != "salut" {
		t.Fatal("les messages non-systeme doivent etre intacts")
	}
	// Aucun changement si l'ancienne directive est absente.
	again := replaceWebDirective(out, oldDir, newDir)
	if s2, _ := again[0].Content.(string); s2 != s {
		t.Fatal("remplacement idempotent attendu")
	}
}

// TestAgentDeepWebDisablesNative : en recherche approfondie, le plugin
// natif est desactive (pas de Extra "plugins") meme avec un provider natif
// (openrouter), les outils web_search/web_fetch restent disponibles et la
// directive DEEP WEB RESEARCH est injectee.
func TestAgentDeepWebDisablesNative(t *testing.T) {
	sp := &scriptedProvider{id: "openrouter", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "openrouter", Model: "openai/gpt-5"}))
	e.SetSearcher(stubWebTools{})
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "cherche des infos recentes", Web: true, WebDepth: "deep", Think: false})
	reqs := sp.requests()
	if len(reqs) == 0 {
		t.Fatal("aucune requete emise")
	}
	if reqs[0].Extra != nil {
		t.Fatalf("deep : aucun plugin natif attendu, obtenu Extra=%v", reqs[0].Extra)
	}
	foundSearch, foundFetch := false, false
	for _, tl := range reqs[0].Tools {
		if tl.Function.Name == "web_search" {
			foundSearch = true
		}
		if tl.Function.Name == "web_fetch" {
			foundFetch = true
		}
	}
	if !foundSearch || !foundFetch {
		t.Fatal("deep : web_search et web_fetch attendus dans les outils")
	}
	deepFound := false
	for _, m := range reqs[0].Messages {
		if s, ok := m.Content.(string); ok && m.Role == "system" && strings.Contains(s, "DEEP WEB RESEARCH") {
			deepFound = true
		}
	}
	if !deepFound {
		t.Fatal("deep : directive DEEP WEB RESEARCH attendue dans les messages systeme")
	}
}
