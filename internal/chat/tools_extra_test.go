package chat

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/vfs"
)

func extraSandbox(t *testing.T) *Sandbox {
	t.Helper()
	lfs, err := vfs.NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	return NewSandboxFS(lfs)
}

// TestExtraToolSchemas vérifie que les 8 outils sont déclarés avec leurs
// champs requis.
func TestExtraToolSchemas(t *testing.T) {
	schemas := ToolSchemas()
	byName := map[string]bool{}
	required := map[string][]string{}
	for _, s := range schemas {
		byName[s.Function.Name] = true
		if p, ok := s.Function.Parameters.(map[string]any); ok {
			if r, ok := p["required"].([]string); ok {
				required[s.Function.Name] = r
			}
		}
	}
	// Phase 3 : Cat est fusionné dans Read — plus de schéma Cat.
	for _, n := range []string{"Tree", "Echo", "Mkdir", "Mv", "Sed", "Awk", "Curl"} {
		if !byName[n] {
			t.Fatalf("schema manquant: %s", n)
		}
	}
	check := map[string]string{
		"Echo": "text", "Mkdir": "path",
		"Mv": "src", "Sed": "expression", "Awk": "program", "Curl": "url",
	}
	for tool, field := range check {
		found := false
		for _, r := range required[tool] {
			if r == field {
				found = true
			}
		}
		if !found {
			t.Fatalf("%s: champ requis %q absent", tool, field)
		}
	}
	if len(required["Mv"]) != 2 {
		t.Fatalf("Mv: src+dst requis, got %v", required["Mv"])
	}
}

// TestToolTreeCatEcho vérifie les outils de lecture pure.
func TestToolTreeCatEcho(t *testing.T) {
	ctx := context.Background()
	sb := extraSandbox(t)
	for _, f := range []string{"src/main.go", "src/util.go", "README.md"} {
		if r := sb.Execute(ctx, "Write", `{"file_path":"`+f+`","content":"x"}`); !strings.HasPrefix(r.Text, "[ok]") {
			t.Fatalf("write %s: %s", f, r.Text)
		}
	}
	r := sb.Execute(ctx, "Tree", `{}`)
	if !strings.Contains(r.Text, "src/") || !strings.Contains(r.Text, "README.md") {
		t.Fatalf("tree: %s", r.Text)
	}
	if !strings.Contains(r.Text, "└──") && !strings.Contains(r.Text, "├──") {
		t.Fatalf("tree: branches ascii absentes: %s", r.Text)
	}
	r = sb.Execute(ctx, "Tree", `{"path":"src","max_depth":1}`)
	if !strings.Contains(r.Text, "main.go") || strings.Contains(r.Text, "README.md") {
		t.Fatalf("tree src: %s", r.Text)
	}
	// Phase 3 : Cat n'est plus un schéma, mais reste un alias de Read
	// (filet de sécurité) : head -> limit, tail -> dernières lignes.
	sb.Execute(ctx, "Write", `{"file_path":"n.txt","content":"l1\nl2\nl3\nl4\nl5"}`)
	r = sb.Execute(ctx, "Cat", `{"file_path":"n.txt","head":2}`)
	if r.Text != "l1\nl2\n… (5 lignes au total, affichees 1-2)" {
		t.Fatalf("cat head: %q", r.Text)
	}
	r = sb.Execute(ctx, "Cat", `{"file_path":"n.txt","tail":2}`)
	if r.Text != "l4\nl5" {
		t.Fatalf("cat tail: %q", r.Text)
	}
	r = sb.Execute(ctx, "Cat", `{"file_path":"n.txt","head":1,"tail":1}`)
	if !strings.Contains(r.Text, "[error]") {
		t.Fatalf("cat head+tail: aurait du echouer: %s", r.Text)
	}
	// Confinement.
	r = sb.Execute(ctx, "Cat", `{"file_path":"../evil"}`)
	if !strings.Contains(r.Text, "[error]") {
		t.Fatalf("cat confinement: %s", r.Text)
	}
	// Echo.
	r = sb.Execute(ctx, "Echo", `{"text":"bonjour"}`)
	if r.Text != "bonjour" {
		t.Fatalf("echo: %q", r.Text)
	}
}

// TestToolMkdirMv vérifie les mutations confinées.
func TestToolMkdirMv(t *testing.T) {
	ctx := context.Background()
	sb := extraSandbox(t)
	if r := sb.Execute(ctx, "Mkdir", `{"path":"a/b/c"}`); !strings.Contains(r.Text, "[ok]") {
		t.Fatalf("mkdir: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Mkdir", `{"path":"../evil"}`); !strings.Contains(r.Text, "[error]") {
		t.Fatalf("mkdir confinement: %s", r.Text)
	}
	sb.Execute(ctx, "Write", `{"file_path":"a/f.txt","content":"data"}`)
	if r := sb.Execute(ctx, "Mv", `{"src":"a/f.txt","dst":"a/b/g.txt"}`); !strings.Contains(r.Text, "[ok]") {
		t.Fatalf("mv: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Read", `{"file_path":"a/b/g.txt"}`); !strings.Contains(r.Text, "data") {
		t.Fatalf("mv: contenu perdu: %s", r.Text)
	}
	// Pas d'écrasement sans overwrite.
	sb.Execute(ctx, "Write", `{"file_path":"h.txt","content":"keep"}`)
	if r := sb.Execute(ctx, "Mv", `{"src":"a/b/g.txt","dst":"h.txt"}`); !strings.Contains(r.Text, "[error]") {
		t.Fatalf("mv overwrite implicite: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Mv", `{"src":"a/b/g.txt","dst":"h.txt","overwrite":true}`); !strings.Contains(r.Text, "[ok]") {
		t.Fatalf("mv overwrite=true: %s", r.Text)
	}
	// Hors racine refusé des deux côtés.
	if r := sb.Execute(ctx, "Mv", `{"src":"h.txt","dst":"../evil"}`); !strings.Contains(r.Text, "[error]") {
		t.Fatalf("mv dst confinement: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Mv", `{"src":"../x","dst":"h2.txt"}`); !strings.Contains(r.Text, "[error]") {
		t.Fatalf("mv src confinement: %s", r.Text)
	}
	// Approbations requises.
	for _, n := range []string{"Mkdir", "Mv", "Curl"} {
		if !needsApproval(n) {
			t.Fatalf("%s devrait exiger une approbation", n)
		}
	}
	for _, n := range []string{"Tree", "Cat", "Echo"} {
		if needsApproval(n) {
			t.Fatalf("%s ne devrait pas exiger d'approbation", n)
		}
	}
}

// TestSedRisk vérifie la détection heuristique des effets de bord sed.
func TestSedRisk(t *testing.T) {
	cases := []struct {
		expr      string
		wantExec  bool
		wantWrite bool
	}{
		{"s/foo/bar/g", false, false},
		{"10,20p", false, false},
		{"/^#/d", false, false},
		{"s/x/y/e", true, false},          // flag e : exécute
		{"s/x/echo hi/e", true, false},    // flag e explicite
		{"e whoami", true, false},         // commande e
		{"s/a/b/w /tmp/out", false, true}, // flag w : écrit
		{"w /tmp/out", false, true},       // commande w
		{"w/tmp/out", false, true},        // H7 : w collé au fichier
		{"W/tmp/out", false, true},        // H7 : W collé au fichier
		{"wout.txt", false, true},         // H7 : w + nom relatif collé
		{"r /etc/passwd", false, true},    // H5 : r + chemin
		{"r/etc/passwd", false, true},     // H5 : r collé au chemin
		{"R /etc/passwd", false, true},    // H5 : R + chemin
		{"/root/d", false, false},         // 'r' d'adresse : pas de faux positif
		{"/^warning/d", false, false},     // 'w' d'adresse : pas de faux positif
		{"/pat/w/tmp/x", false, true},     // adresse + commande w réelle
		{"s/where/there/g", false, false}, // 'e' dans un mot : pas de faux positif
		{"s/foo\\/bar/baz/g", false, false},
	}
	for _, c := range cases {
		exec, write := sedRisk(c.expr)
		if exec != c.wantExec || write != c.wantWrite {
			t.Errorf("sedRisk(%q) = exec=%v write=%v, want exec=%v write=%v",
				c.expr, exec, write, c.wantExec, c.wantWrite)
		}
	}
}

// TestAwkRisk vérifie la détection heuristique des effets de bord awk.
func TestAwkRisk(t *testing.T) {
	safe := []string{
		`{print $2}`,
		`NR>1 {sum+=$3} END {print sum}`,
		`$0 ~ /error/ {print NR": "$0}`,
		`BEGIN{FS=","} {print $1}`,
		`length($0) > 80`,
	}
	for _, p := range safe {
		if awkNeedsApproval(map[string]any{"program": p}) {
			t.Errorf("awk faux positif (devrait etre libre): %q", p)
		}
	}
	risky := []string{
		`{system("rm -rf /")}`,
		`{print | "sort"}`,
		`{"ls" | getline x}`,
		`{print $0 > "/tmp/out"}`,
		`{printf "%s" , $0 >> "f"}`,
		`{getline < "secret"}`,
	}
	for _, p := range risky {
		if !awkNeedsApproval(map[string]any{"program": p}) {
			t.Errorf("awk non detecte (devrait exiger approbation): %q", p)
		}
	}
}

// TestNeedsApprovalFor vérifie l'affinage par arguments.
func TestNeedsApprovalFor(t *testing.T) {
	args := func(m map[string]any) map[string]any { return m }
	cases := []struct {
		name string
		args map[string]any
		want bool
	}{
		{"Sed", args(map[string]any{"expression": "s/a/b/g", "file": "f"}), false},
		{"Sed", args(map[string]any{"expression": "s/a/b/g", "file": "f", "in_place": true}), true},
		{"Sed", args(map[string]any{"expression": "s/x/y/e", "file": "f"}), true},
		{"Awk", args(map[string]any{"program": "{print $1}", "file": "f"}), false},
		{"Awk", args(map[string]any{"program": "{system(\"id\")}", "file": "f"}), true},
		{"Mkdir", args(map[string]any{"path": "a"}), true},
		{"Mv", args(map[string]any{"src": "a", "dst": "b"}), true},
		{"Curl", args(map[string]any{"url": "https://x"}), true},
		{"Tree", args(map[string]any{}), false},
		{"Cat", args(map[string]any{"file_path": "f"}), false},
		{"Read", args(map[string]any{"file_path": "f"}), false},
	}
	for _, c := range cases {
		if got := needsApprovalFor(c.name, c.args); got != c.want {
			t.Errorf("needsApprovalFor(%s, %v) = %v, want %v", c.name, c.args, got, c.want)
		}
	}
}

// TestToolSedAwkStream vérifie sed/awk en mode flux (binaires système).
func TestToolSedAwkStream(t *testing.T) {
	if _, err := exec.LookPath("sed"); err != nil {
		t.Skip("sed absent")
	}
	ctx := context.Background()
	sb := extraSandbox(t)
	sb.Execute(ctx, "Write", `{"file_path":"d.txt","content":"foo 1\nbar 2\nfoo 3"}`)
	r := sb.Execute(ctx, "Sed", `{"expression":"s/foo/FOO/g","file":"d.txt"}`)
	if r.Text != "FOO 1\nbar 2\nFOO 3" {
		t.Fatalf("sed stream: %q", r.Text)
	}
	// in_place avec diff.
	r = sb.Execute(ctx, "Sed", `{"expression":"s/FOO/foo/","file":"d.txt","in_place":true}`)
	if !strings.Contains(r.Text, "[ok]") {
		t.Fatalf("sed -i: %s", r.Text)
	}
	if len(r.Diff) == 0 {
		t.Fatal("sed -i: diff attendu")
	}
	r = sb.Execute(ctx, "Read", `{"file_path":"d.txt"}`)
	if !strings.Contains(r.Text, "foo 1") {
		t.Fatalf("sed -i: contenu: %s", r.Text)
	}
	// input texte au lieu de fichier.
	r = sb.Execute(ctx, "Sed", `{"expression":"s/a/b/","input":"aaa"}`)
	if strings.TrimSpace(r.Text) != "baa" {
		t.Fatalf("sed input: %q", r.Text)
	}
	if _, err := exec.LookPath("awk"); err != nil {
		t.Skip("awk absent")
	}
	r = sb.Execute(ctx, "Awk", `{"program":"{print $2}","file":"d.txt"}`)
	if r.Text != "1\n2\n3" {
		t.Fatalf("awk: %q", r.Text)
	}
	r = sb.Execute(ctx, "Awk", `{"program":"$1==\"bar\"","file":"d.txt"}`)
	if strings.TrimSpace(r.Text) != "bar 2" {
		t.Fatalf("awk filtre: %q", r.Text)
	}
}

// TestToolCurl vérifie la version encadrée de curl.
func TestToolCurl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong"))
	}))
	defer srv.Close()
	ctx := context.Background()
	sb := extraSandbox(t)

	r := sb.Execute(ctx, "Curl", `{"url":"`+srv.URL+`"}`)
	if !strings.Contains(r.Text, "non autorisee") {
		t.Fatalf("curl vers loopback aurait du etre bloque (garde SSRF): %s", r.Text)
	}
	// (Le chemin nominal GET/POST n'est plus testable en local : la garde
	// SSRF bloque 127.0.0.1. Voir TestCurlSSRFGuard.)
	// Cadre : schémas non-http refusés.
	for _, bad := range []string{"file:///etc/passwd", "ftp://x", "not a url", ""} {
		r = sb.Execute(ctx, "Curl", `{"url":"`+bad+`"}`)
		if !strings.Contains(r.Text, "[error]") {
			t.Fatalf("curl %q aurait du etre refuse: %s", bad, r.Text)
		}
	}
	r = sb.Execute(ctx, "Curl", `{"url":"`+srv.URL+`","method":"TRACE"}`)
	if !strings.Contains(r.Text, "[error]") {
		t.Fatalf("curl TRACE aurait du etre refuse: %s", r.Text)
	}
}

// TestCurlSSRFGuard vérifie la garde anti-SSRF de Curl : seules les IP
// publiques sont joignables (même logique que web_fetch).
func TestCurlSSRFGuard(t *testing.T) {
	private := []string{
		"127.0.0.1", "::1", "10.1.2.3", "172.16.5.4", "192.168.1.1",
		"169.254.169.254", "100.64.0.1", "0.0.0.0", "::",
		"224.0.0.1", "fe80::1", "fc00::1",
	}
	for _, s := range private {
		if curlIsPublicIP(net.ParseIP(s)) {
			t.Errorf("curlIsPublicIP(%q) = true, want false", s)
		}
	}
	public := []string{"8.8.8.8", "1.1.1.1", "93.184.216.34", "2001:4860:4860::8888"}
	for _, s := range public {
		if !curlIsPublicIP(net.ParseIP(s)) {
			t.Errorf("curlIsPublicIP(%q) = false, want true", s)
		}
	}
	if curlIsPublicIP(nil) {
		t.Error("curlIsPublicIP(nil) = true, want false")
	}
	// Le composeur refuse de joindre le loopback, par IP comme par nom.
	ctx := context.Background()
	if _, err := curlDialContext(ctx, "tcp", "127.0.0.1:80"); !isSSRFErr(err) {
		t.Errorf("curlDialContext vers 127.0.0.1: got %v, want erreur SSRF", err)
	}
	if _, err := curlDialContext(ctx, "tcp", net.JoinHostPort("localhost", "80")); !isSSRFErr(err) {
		t.Errorf("curlDialContext vers localhost: got %v, want erreur SSRF", err)
	}
}

func isSSRFErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "non autorisee")
}

// TestMarkExternalOutput vérifie le marquage des sorties à source externe
// (M5) : GitHub* et Curl marqués, outils workspace inchangés.
func TestMarkExternalOutput(t *testing.T) {
	for _, n := range []string{"Curl", "GitHubIssues", "GitHubIssueGet", "GitHubPRs", "GitHubRepos", "GitHubIssueCreate"} {
		got := markExternalOutput(n, "contenu")
		if !strings.HasPrefix(got, externalOutputMarker) {
			t.Errorf("markExternalOutput(%q) sans marque: %q", n, got)
		}
	}
	for _, n := range []string{"Read", "Bash", "Grep", "Sed", "Awk", "WebFetch"} {
		if got := markExternalOutput(n, "contenu"); got != "contenu" {
			t.Errorf("markExternalOutput(%q) = %q, want inchange", n, got)
		}
	}
}
