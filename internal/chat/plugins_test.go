package chat

import (
	"context"
	"errors"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/plugins"
)

type stubPlugins struct {
	defs   []plugins.Def
	out    string
	err    error
	calls  []string
	infos  []plugins.Info
	perrs  []string
	loaded bool
}

func (s *stubPlugins) Defs() []plugins.Def { return s.defs }
func (s *stubPlugins) Execute(ctx context.Context, name, argsJSON string) (string, error) {
	s.calls = append(s.calls, name)
	return s.out, s.err
}
func (s *stubPlugins) Plugins() []plugins.Info { return s.infos }
func (s *stubPlugins) Errors() []string        { return s.perrs }
func (s *stubPlugins) Load() error {
	s.loaded = true
	return nil
}

func demoPlugins(out string, err error) *stubPlugins {
	return &stubPlugins{
		defs: []plugins.Def{{
			Plugin:      "demo",
			Name:        "plugin_demo_ping",
			Description: "Ping [plugin demo]",
			Schema:      map[string]any{"type": "object"},
		}},
		infos: []plugins.Info{{Name: "demo", Version: "1.0.0", Tools: []string{"plugin_demo_ping"}}},
		out:   out,
		err:   err,
	}
}

func TestPluginToolSchemas(t *testing.T) {
	schemas := PluginToolSchemas(demoPlugins("", nil).defs)
	if len(schemas) != 1 {
		t.Fatalf("schemas = %d", len(schemas))
	}
	fn := schemas[0].Function
	if fn.Name != "plugin_demo_ping" {
		t.Fatalf("nom = %q", fn.Name)
	}
	if len(fn.Description) < 8 || fn.Description[:8] != "[Plugin]" {
		t.Fatalf("description sans prefixe [Plugin] : %q", fn.Description)
	}
}

func TestAgentExposesPluginTools(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetPlugins(demoPlugins("", nil))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
	if !hasToolName(sp.requests()[0].Tools, "plugin_demo_ping") {
		t.Fatal("l'outil plugin doit etre expose en mode agent")
	}
}

func TestPluginExecuteOK(t *testing.T) {
	e := newAgentEngine(t, &scriptedProvider{id: "fake"}, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	stub := demoPlugins("pong", nil)
	e.SetPlugins(stub)
	res := e.pluginExecute(context.Background(), "plugin_demo_ping", `{}`)
	if res.Text != "pong" {
		t.Fatalf("resultat = %q", res.Text)
	}
	if len(stub.calls) != 1 || stub.calls[0] != "plugin_demo_ping" {
		t.Fatalf("appels = %v", stub.calls)
	}
}

func TestPluginExecuteErreur(t *testing.T) {
	e := newAgentEngine(t, &scriptedProvider{id: "fake"}, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetPlugins(demoPlugins("", errors.New("boom")))
	res := e.pluginExecute(context.Background(), "plugin_demo_ping", `{}`)
	if res.Text != "[erreur plugin] boom" {
		t.Fatalf("resultat = %q", res.Text)
	}
}

func TestPluginExecuteSansManager(t *testing.T) {
	e := newAgentEngine(t, &scriptedProvider{id: "fake"}, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	res := e.pluginExecute(context.Background(), "plugin_x_y", `{}`)
	if res.Text == "" {
		t.Fatal("erreur attendue sans gestionnaire")
	}
}

func TestPluginApprovalEtDedup(t *testing.T) {
	if !needsApproval("plugin_demo_ping") {
		t.Fatal("les outils plugin exigent une approbation")
	}
	if dedupableTool("plugin_demo_ping") {
		t.Fatal("les outils plugin ne doivent pas etre dedupliques")
	}
}

func TestPluginStatusEtReload(t *testing.T) {
	e := newAgentEngine(t, &scriptedProvider{id: "fake"}, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	stub := demoPlugins("", nil)
	stub.perrs = []string{"autre: manifeste invalide"}
	e.SetPlugins(stub)
	infos, errs := e.PluginStatus()
	if len(infos) != 1 || infos[0].Name != "demo" {
		t.Fatalf("infos = %+v", infos)
	}
	if len(errs) != 1 {
		t.Fatalf("errs = %v", errs)
	}
	if err := e.PluginReload(); err != nil {
		t.Fatal(err)
	}
	if !stub.loaded {
		t.Fatal("le rechargement doit appeler Load()")
	}
}
