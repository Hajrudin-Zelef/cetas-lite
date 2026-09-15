// Package plugins : extensions externes de Cetas.
//
// Un plugin est un dossier pose dans <home>/plugins/<nom>/ contenant un
// manifeste plugin.json. Chaque plugin declare des outils que l'agent (et le
// chat, selon la famille) peut appeler ; l'outil est soit une commande locale
// (n'importe quel langage : python, node, shell...), soit un appel HTTP.
//
// Manifeste minimal :
//
//	{
//	  "name": "mon-plugin",
//	  "version": "1.0.0",
//	  "description": "Ce que fait le plugin",
//	  "tools": [
//	    {
//	      "name": "resumer",
//	      "description": "Resume un texte",
//	      "parameters": {"type": "object",
//	        "properties": {"texte": {"type": "string"}},
//	        "required": ["texte"]},
//	      "exec": {"command": ["python3", "resumer.py"], "timeout_sec": 30}
//	    }
//	  ]
//	}
//
// Protocole "exec" : les arguments JSON de l'appel sont ecrits sur stdin ;
// le programme repond sur stdout soit du JSON {"result": ...} (ou {"error":
// ...}), soit du texte brut utilise tel quel comme resultat. Le processus
// est tue a la fin du timeout. Aucun shell n'est utilise : la commande est
// un tableau argv exact. Le repertoire de travail est confine au dossier du
// plugin (les ".." sont refuses).
package plugins

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	maxStdout   = 256 << 10 // 256 Ko de sortie max par execution
	maxStderr   = 8 << 10
	maxTimeout  = 300 // secondes
	defaultTout = 60
)

var validName = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

// ExecDef decrit une commande locale (argv, sans shell).
type ExecDef struct {
	Command    []string `json:"command"`
	TimeoutSec int      `json:"timeout_sec,omitempty"`
	Dir        string   `json:"dir,omitempty"`
}

// HTTPDef decrit un appel HTTP (JSON in, JSON ou texte out).
type HTTPDef struct {
	URL        string            `json:"url"`
	Method     string            `json:"method,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	TimeoutSec int               `json:"timeout_sec,omitempty"`
}

// ToolDef : un outil expose par un plugin.
type ToolDef struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters,omitempty"`
	Required    []string       `json:"required,omitempty"`
	Exec        *ExecDef       `json:"exec,omitempty"`
	HTTP        *HTTPDef       `json:"http,omitempty"`
}

// Manifest : le contenu de plugin.json.
type Manifest struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Description string    `json:"description,omitempty"`
	Tools       []ToolDef `json:"tools"`
}

// Def : un outil pret a etre expose au modele.
type Def struct {
	Plugin      string
	Name        string // nom complet : plugin_<plugin>_<outil>
	Description string
	Schema      map[string]any
}

// Info : etat public d'un plugin charge.
type Info struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description,omitempty"`
	Tools       []string `json:"tools"`
}

type loadedTool struct {
	def ToolDef
	dir string // dossier du plugin (workdir de base)
}

type loadedPlugin struct {
	manifest Manifest
	tools    []loadedTool
}

// Manager charge les plugins depuis un dossier et execute leurs outils.
type Manager struct {
	dir     string
	client  *http.Client
	mu      sync.RWMutex
	plugins []*loadedPlugin
	byName  map[string]*loadedTool // nom complet -> outil
	errors  []string
}

// NewManager cree un gestionnaire (sans charger). Load effectue le chargement.
func NewManager(dir string, client *http.Client) *Manager {
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	return &Manager{dir: dir, client: client, byName: map[string]*loadedTool{}}
}

// Dir retourne le dossier des plugins.
func (m *Manager) Dir() string { return m.dir }

// Load (re)charge tous les plugins du dossier. Les erreurs par plugin sont
// conservees et exposees via Errors(), sans faire echouer l'ensemble.
func (m *Manager) Load() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.plugins = nil
	m.byName = map[string]*loadedTool{}
	m.errors = nil

	entries, err := os.ReadDir(m.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // aucun plugin : cas normal
		}
		return fmt.Errorf("lecture %s: %w", m.dir, err)
	}
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		pdir := filepath.Join(m.dir, e.Name())
		lp, err := loadOne(pdir)
		if err != nil {
			m.errors = append(m.errors, fmt.Sprintf("%s: %v", e.Name(), err))
			continue
		}
		m.plugins = append(m.plugins, lp)
		for i := range lp.tools {
			t := &lp.tools[i]
			full := fullName(lp.manifest.Name, t.def.Name)
			if _, dup := m.byName[full]; dup {
				m.errors = append(m.errors, fmt.Sprintf("%s: outil duplique %q", lp.manifest.Name, full))
				continue
			}
			m.byName[full] = t
		}
	}
	sort.Slice(m.plugins, func(i, j int) bool { return m.plugins[i].manifest.Name < m.plugins[j].manifest.Name })
	return nil
}

// Errors retourne les erreurs de chargement non fatales.
func (m *Manager) Errors() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return append([]string(nil), m.errors...)
}

// Plugins liste les plugins charges.
func (m *Manager) Plugins() []Info {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Info, 0, len(m.plugins))
	for _, lp := range m.plugins {
		tools := make([]string, 0, len(lp.tools))
		for _, t := range lp.tools {
			tools = append(tools, fullName(lp.manifest.Name, t.def.Name))
		}
		out = append(out, Info{
			Name:        lp.manifest.Name,
			Version:     lp.manifest.Version,
			Description: lp.manifest.Description,
			Tools:       tools,
		})
	}
	return out
}

// Defs retourne les outils exposes au modele.
func (m *Manager) Defs() []Def {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var out []Def
	for _, lp := range m.plugins {
		for _, t := range lp.tools {
			schema := map[string]any{"type": "object"}
			if t.def.Parameters != nil {
				schema = t.def.Parameters
			}
			desc := t.def.Description
			if desc == "" {
				desc = "Outil du plugin " + lp.manifest.Name
			}
			out = append(out, Def{
				Plugin:      lp.manifest.Name,
				Name:        fullName(lp.manifest.Name, t.def.Name),
				Description: desc + " [plugin " + lp.manifest.Name + "]",
				Schema:      schema,
			})
		}
	}
	return out
}

// Execute lance l'outil de nom complet avec les arguments JSON donnes.
func (m *Manager) Execute(ctx context.Context, full, argsJSON string) (string, error) {
	m.mu.RLock()
	t, ok := m.byName[full]
	m.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("outil plugin inconnu: %s", full)
	}
	if t.def.Exec != nil {
		return runExec(ctx, t, argsJSON)
	}
	if t.def.HTTP != nil {
		return m.runHTTP(ctx, t, argsJSON)
	}
	return "", fmt.Errorf("outil %s sans exec ni http", full)
}

func fullName(plugin, tool string) string {
	return "plugin_" + plugin + "_" + tool
}

func loadOne(dir string) (*loadedPlugin, error) {
	raw, err := os.ReadFile(filepath.Join(dir, "plugin.json"))
	if err != nil {
		return nil, fmt.Errorf("plugin.json: %w", err)
	}
	var mf Manifest
	if err := json.Unmarshal(raw, &mf); err != nil {
		return nil, fmt.Errorf("plugin.json invalide: %w", err)
	}
	if !validName.MatchString(mf.Name) {
		return nil, fmt.Errorf("nom de plugin invalide %q", mf.Name)
	}
	if filepath.Base(dir) != mf.Name {
		return nil, fmt.Errorf("le dossier %q ne correspond pas au nom %q", filepath.Base(dir), mf.Name)
	}
	if len(mf.Tools) == 0 {
		return nil, fmt.Errorf("aucun outil declare")
	}
	lp := &loadedPlugin{manifest: mf}
	seen := map[string]bool{}
	for _, td := range mf.Tools {
		if !validName.MatchString(td.Name) {
			return nil, fmt.Errorf("nom d'outil invalide %q", td.Name)
		}
		if seen[td.Name] {
			return nil, fmt.Errorf("outil duplique %q", td.Name)
		}
		seen[td.Name] = true
		if full := fullName(mf.Name, td.Name); len(full) > 64 {
			return nil, fmt.Errorf("nom d'outil trop long %q (64 max)", full)
		}
		if (td.Exec == nil) == (td.HTTP == nil) {
			return nil, fmt.Errorf("outil %q : exactement un de exec/http requis", td.Name)
		}
		if td.Exec != nil {
			if len(td.Exec.Command) == 0 {
				return nil, fmt.Errorf("outil %q : command vide", td.Name)
			}
			if td.Exec.TimeoutSec < 0 || td.Exec.TimeoutSec > maxTimeout {
				return nil, fmt.Errorf("outil %q : timeout_sec hors borne (0..%d)", td.Name, maxTimeout)
			}
			if _, err := confinedDir(dir, td.Exec.Dir); err != nil {
				return nil, fmt.Errorf("outil %q : %w", td.Name, err)
			}
		}
		if td.HTTP != nil {
			if strings.TrimSpace(td.HTTP.URL) == "" {
				return nil, fmt.Errorf("outil %q : url vide", td.Name)
			}
			if td.HTTP.TimeoutSec < 0 || td.HTTP.TimeoutSec > maxTimeout {
				return nil, fmt.Errorf("outil %q : timeout_sec hors borne (0..%d)", td.Name, maxTimeout)
			}
		}
		lp.tools = append(lp.tools, loadedTool{def: td, dir: dir})
	}
	return lp, nil
}

// confinedDir resout le workdir d'un outil en le confinant au dossier du plugin.
func confinedDir(pluginDir, sub string) (string, error) {
	if strings.TrimSpace(sub) == "" {
		return pluginDir, nil
	}
	if filepath.IsAbs(sub) {
		return "", fmt.Errorf("dir absolu interdit")
	}
	joined := filepath.Join(pluginDir, sub)
	abs, err := filepath.Abs(joined)
	if err != nil {
		return "", err
	}
	base, err := filepath.Abs(pluginDir)
	if err != nil {
		return "", err
	}
	if abs != base && !strings.HasPrefix(abs, base+string(filepath.Separator)) {
		return "", fmt.Errorf("dir hors du plugin (.. interdit)")
	}
	return abs, nil
}

func timeoutOf(sec int) time.Duration {
	if sec <= 0 {
		sec = defaultTout
	}
	return time.Duration(sec) * time.Second
}

// runExec execute la commande : args JSON sur stdin, resultat sur stdout.
// L'execution brute est deleguee a runExecRaw (specifique a la plateforme) :
// le plugin tourne isole — groupe de processus sous Unix, Job Object sous
// Windows — pour qu'a la fin du timeout (ou a l'annulation) tout l'arbre
// soit tue et qu'aucun orphelin ne survive.
func runExec(ctx context.Context, t *loadedTool, argsJSON string) (string, error) {
	stdout, err := runExecRaw(ctx, t, argsJSON)
	if err != nil {
		return "", err
	}
	out := bytes.TrimSpace(stdout)
	if len(out) == 0 {
		return "", fmt.Errorf("exec: sortie vide")
	}
	// Protocole JSON prefere : {"result": ...} / {"error": ...}.
	var parsed struct {
		Result any    `json:"result"`
		Error  string `json:"error"`
	}
	if json.Unmarshal(out, &parsed) == nil && (parsed.Error != "" || parsed.Result != nil) {
		if parsed.Error != "" {
			return "", fmt.Errorf("plugin: %s", truncateRunes(parsed.Error, 500))
		}
		switch r := parsed.Result.(type) {
		case string:
			return r, nil
		default:
			raw, _ := json.Marshal(r)
			return string(raw), nil
		}
	}
	return truncateRunes(string(out), 64<<10), nil
}

// runHTTP appelle le endpoint HTTP du plugin (JSON in, JSON ou texte out).
func (m *Manager) runHTTP(ctx context.Context, t *loadedTool, argsJSON string) (string, error) {
	h := t.def.HTTP
	method := strings.ToUpper(strings.TrimSpace(h.Method))
	if method == "" {
		method = "POST"
	}
	cctx, cancel := context.WithTimeout(ctx, timeoutOf(h.TimeoutSec))
	defer cancel()
	req, err := http.NewRequestWithContext(cctx, method, h.URL, strings.NewReader(argsJSON))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain")
	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxStdout))
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("http %d: %s", resp.StatusCode, truncateRunes(strings.TrimSpace(string(raw)), 300))
	}
	out := bytes.TrimSpace(raw)
	var parsed struct {
		Result any    `json:"result"`
		Error  string `json:"error"`
	}
	if json.Unmarshal(out, &parsed) == nil && (parsed.Error != "" || parsed.Result != nil) {
		if parsed.Error != "" {
			return "", fmt.Errorf("plugin: %s", truncateRunes(parsed.Error, 500))
		}
		if s, ok := parsed.Result.(string); ok {
			return s, nil
		}
		enc, _ := json.Marshal(parsed.Result)
		return string(enc), nil
	}
	return truncateRunes(string(out), 64<<10), nil
}

type limitedWriter struct {
	w *bytes.Buffer
	n int
}

func (l *limitedWriter) Write(p []byte) (int, error) {
	room := l.n - l.w.Len()
	if room <= 0 {
		return len(p), nil // tronque silencieusement
	}
	if len(p) > room {
		p = p[:room]
	}
	return l.w.Write(p)
}

func truncateRunes(s string, max int) string {
	r := []rune(strings.TrimSpace(s))
	if len(r) <= max {
		return string(r)
	}
	return string(r[:max]) + "\n[...]"
}
