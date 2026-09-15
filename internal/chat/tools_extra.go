package chat

// Outils Unix "encadrés" : Tree, Cat, Echo (lecture seule, sans approbation),
// Mkdir, Mv (mutations confinées, approbation requise), Sed, Awk (binaires
// système sans shell ; approbation seulement pour les usages à effets de
// bord), Curl (HTTP(S) en Go pur : approbation, timeout, taille bornée).
//
// Cadre commun à tous :
// - chemins relatifs au workspace, confinement vérifié (Resolve/checkArg) ;
// - sorties bornées (toolMaxOutput) ;
// - exécution via s.fs.Exec : identique en local et sur projet SFTP distant
//   (SSH). Sed/Awk exigent le binaire sur la machine d'exécution.

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/vfs"
)

const (
	treeDefaultDepth = 3
	treeMaxDepth     = 6
	treeMaxEntries   = 500
	curlDefaultTO    = 15
	curlMaxTO        = 30
	curlMaxBody      = 2 * 1024 * 1024
	sedAwkMaxTO      = 60
	echoMaxLen       = 4000
)

// extraToolSchemas déclare les schémas des outils Unix encadrés.
func extraToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "Tree", Description: "Show the workspace directory tree (bounded depth). Read-only.", Parameters: str(map[string]any{
			"path":      map[string]any{"type": "string", "description": "Dossier relatif, defaut: racine"},
			"max_depth": map[string]any{"type": "integer", "description": "Profondeur max (defaut 3, max 6)"},
		})}},
		{Type: "function", Function: provider.ToolFunction{Name: "Cat", Description: "Print a file (head/tail line selection). Read-only.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string"},
			"head":      map[string]any{"type": "integer", "description": "N premieres lignes"},
			"tail":      map[string]any{"type": "integer", "description": "N dernieres lignes (exclusif avec head)"},
		}, "file_path")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Echo", Description: "Print text back (4000 chars max). Read-only.", Parameters: str(map[string]any{
			"text": map[string]any{"type": "string"},
		}, "text")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Mkdir", Description: "Create directories in the workspace (mkdir -p). Needs approval.", Parameters: str(map[string]any{
			"path": map[string]any{"type": "string", "description": "Chemin relatif a creer"},
		}, "path")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Mv", Description: "Move/rename a file or directory inside the workspace. Needs approval.", Parameters: str(map[string]any{
			"src":       map[string]any{"type": "string"},
			"dst":       map[string]any{"type": "string"},
			"overwrite": map[string]any{"type": "boolean", "description": "Ecraser la destination existante (defaut false)"},
		}, "src", "dst")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Sed", Description: "Stream-edit text with GNU sed (no shell). Stream mode is read-only; in_place=true modifies the file and needs approval, like Edit.", Parameters: str(map[string]any{
			"expression": map[string]any{"type": "string", "description": "Expression sed, ex. s/foo/bar/g"},
			"file":       map[string]any{"type": "string", "description": "Fichier du workspace (ou input)"},
			"input":      map[string]any{"type": "string", "description": "Texte d'entree au lieu d'un fichier"},
			"in_place":   map[string]any{"type": "boolean", "description": "Modifier le fichier sur place (defaut false)"},
		}, "expression")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Awk", Description: "Process text with awk (no shell). Pure stream processing is read-only; programs with side effects (system(), file writes, pipes to commands) need approval.", Parameters: str(map[string]any{
			"program":         map[string]any{"type": "string", "description": "Programme awk, ex. {print $2}"},
			"file":            map[string]any{"type": "string", "description": "Fichier du workspace (ou input)"},
			"input":           map[string]any{"type": "string", "description": "Texte d'entree au lieu d'un fichier"},
			"field_separator": map[string]any{"type": "string", "description": "Separateur de champs (-F)"},
		}, "program")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Curl", Description: "HTTP(S) request (framed: needs approval, 30s max, 2MB cap, http/https only).", Parameters: str(map[string]any{
			"url":     map[string]any{"type": "string"},
			"method":  map[string]any{"type": "string", "enum": []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, "description": "Defaut GET"},
			"headers": map[string]any{"type": "object", "description": "En-tetes HTTP"},
			"body":    map[string]any{"type": "string", "description": "Corps de requete"},
			"timeout": map[string]any{"type": "integer", "description": "Secondes (defaut 15, max 30)"},
		}, "url")}},
	}
}

// boolArg lit un argument booléen (bool natif ou chaîne "true"/"1").
func boolArg(args map[string]any, key string) bool {
	switch v := args[key].(type) {
	case bool:
		return v
	case string:
		s := strings.ToLower(strings.TrimSpace(v))
		return s == "true" || s == "1" || s == "yes"
	}
	return false
}

// ---------------- Tree ----------------

func (s *Sandbox) toolTree(ctx context.Context, args map[string]any) ToolResult {
	rel := strings.TrimSpace(strArg(args, "path"))
	depth := intArg(args, "max_depth")
	if depth <= 0 {
		depth = treeDefaultDepth
	}
	if depth > treeMaxDepth {
		depth = treeMaxDepth
	}
	root, err := vfs.Tree(ctx, s.fs, rel, vfs.TreeOptions{MaxDepth: depth, MaxEntries: treeMaxEntries})
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	name := root.Name
	if name == "" {
		name = "."
	}
	var sb strings.Builder
	sb.WriteString(name + "/\n")
	renderTreeChildren(&sb, root.Children, "")
	out := sb.String()
	if root.Truncated {
		out += "… (tronque : profondeur ou nombre d'entrees max atteint)\n"
	}
	return ToolResult{Text: out}
}

func renderTreeChildren(sb *strings.Builder, kids []*vfs.Node, prefix string) {
	for i, k := range kids {
		last := i == len(kids)-1
		branch, cont := "├── ", "│   "
		if last {
			branch, cont = "└── ", "    "
		}
		disp := k.Name
		if k.IsDir {
			disp += "/"
		}
		fmt.Fprintf(sb, "%s%s%s\n", prefix, branch, disp)
		if k.IsDir && len(k.Children) > 0 {
			renderTreeChildren(sb, k.Children, prefix+cont)
		}
	}
}

// ---------------- Cat ----------------

func (s *Sandbox) toolCat(ctx context.Context, args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	clean, err := s.fs.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if clean == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	b, err := s.fs.ReadFile(ctx, clean)
	if err != nil {
		if errors.Is(err, vfs.ErrNotFound) {
			return ToolResult{Text: "[erreur] fichier introuvable: " + rel}
		}
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if isBinary(b) {
		return ToolResult{Text: fmt.Sprintf("[binaire] %s : %d octets, contenu non affiche", rel, len(b))}
	}
	lines := strings.Split(string(b), "\n")
	head, tail := intArg(args, "head"), intArg(args, "tail")
	if head > 0 && tail > 0 {
		return ToolResult{Text: "[erreur] head et tail sont exclusifs"}
	}
	total := len(lines)
	sel := lines
	switch {
	case head > 0 && head < total:
		sel = lines[:head]
	case tail > 0 && tail < total:
		sel = lines[total-tail:]
	}
	out := strings.Join(sel, "\n")
	if len(sel) < total {
		out += fmt.Sprintf("\n… (%d lignes au total)", total)
	}
	return ToolResult{Text: truncate(out, toolMaxOutput)}
}

func isBinary(b []byte) bool {
	n := len(b)
	if n > 8192 {
		n = 8192
	}
	for i := 0; i < n; i++ {
		if b[i] == 0 {
			return true
		}
	}
	return false
}

// ---------------- Echo ----------------

func (s *Sandbox) toolEcho(_ context.Context, args map[string]any) ToolResult {
	text := strArg(args, "text")
	if len(text) > echoMaxLen {
		text = text[:echoMaxLen] + "\n… (tronque)"
	}
	return ToolResult{Text: text}
}

// ---------------- Mkdir / Mv ----------------

func (s *Sandbox) toolMkdir(ctx context.Context, args map[string]any) ToolResult {
	rel := strings.TrimSpace(strArg(args, "path"))
	if rel == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	clean, err := s.fs.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if clean == "" {
		return ToolResult{Text: "[erreur] la racine existe deja"}
	}
	if err := s.fs.MkdirAll(ctx, clean); err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	return ToolResult{Text: "[ok] dossier cree: " + clean}
}

func (s *Sandbox) toolMv(ctx context.Context, args map[string]any) ToolResult {
	src, dst := strings.TrimSpace(strArg(args, "src")), strings.TrimSpace(strArg(args, "dst"))
	cleanSrc, err := s.fs.Resolve(src)
	if err != nil {
		return ToolResult{Text: "[erreur] source : " + err.Error()}
	}
	cleanDst, err := s.fs.Resolve(dst)
	if err != nil {
		return ToolResult{Text: "[erreur] destination : " + err.Error()}
	}
	if cleanSrc == "" || cleanDst == "" {
		return ToolResult{Text: "[erreur] la racine ne peut pas etre deplacee"}
	}
	if cleanSrc == cleanDst {
		return ToolResult{Text: "[erreur] source et destination identiques"}
	}
	if _, err := s.fs.Stat(ctx, cleanSrc); err != nil {
		return ToolResult{Text: "[erreur] source introuvable: " + src}
	}
	if _, err := s.fs.Stat(ctx, cleanDst); err == nil && !boolArg(args, "overwrite") {
		return ToolResult{Text: "[erreur] la destination existe deja : " + dst + " (overwrite=false)"}
	}
	if err := s.fs.Rename(ctx, cleanSrc, cleanDst); err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	return ToolResult{Text: "[ok] deplace : " + cleanSrc + " -> " + cleanDst}
}

// ---------------- Sed / Awk ----------------

// sedInputFile prépare la cible de sed/awk : fichier du workspace (confiné)
// ou texte `input` déposé dans un fichier temporaire (nettoyé après usage).
func (s *Sandbox) sedInputFile(ctx context.Context, file, input, tool string) (string, func(), ToolResult, bool) {
	cleanup := func() {}
	if file != "" && input != "" {
		return "", cleanup, ToolResult{Text: "[erreur] un seul de file / input"}, false
	}
	if file != "" {
		if msg := s.checkArg(file); msg != "" {
			return "", cleanup, ToolResult{Text: "[erreur] " + msg}, false
		}
		clean, err := s.fs.Resolve(file)
		if err != nil {
			return "", cleanup, ToolResult{Text: "[erreur] " + err.Error()}, false
		}
		if clean == "" {
			return "", cleanup, ToolResult{Text: "[erreur] chemin vide"}, false
		}
		if st, err := s.fs.Stat(ctx, clean); err != nil || st.IsDir {
			return "", cleanup, ToolResult{Text: "[erreur] fichier introuvable: " + file}, false
		}
		return clean, cleanup, ToolResult{}, true
	}
	if input == "" {
		return "", cleanup, ToolResult{Text: "[erreur] file ou input requis"}, false
	}
	rel := fmt.Sprintf(".runscript_tmp/%s-input-%d.txt", tool, time.Now().UnixNano())
	if err := s.fs.WriteFile(ctx, rel, []byte(input), 0o600); err != nil {
		return "", cleanup, ToolResult{Text: "[erreur] entree temporaire impossible: " + err.Error()}, false
	}
	cleanup = func() { s.fs.Remove(context.Background(), rel) }
	return rel, cleanup, ToolResult{}, true
}

// runBinary exécute un binaire système sans shell, comme toolBash.
func (s *Sandbox) runBinary(ctx context.Context, argv []string, timeout time.Duration) (string, error) {
	if runtime.GOOS == "windows" && !s.fs.Remote() {
		return "", errors.New("non supporte sur Windows en local (utilise un projet SFTP)")
	}
	if !s.fs.Remote() {
		argv = wrapCommand(s.Isolation, s.localRoot(), argv)
	}
	return s.fs.Exec(ctx, argv[0], argv[1:], nil, timeout)
}

func toolTimeout(args map[string]any, def int) time.Duration {
	t := intArg(args, "timeout")
	if t <= 0 {
		t = def
	}
	if t > sedAwkMaxTO {
		t = sedAwkMaxTO
	}
	return time.Duration(t) * time.Second
}

func (s *Sandbox) toolSed(ctx context.Context, args map[string]any) ToolResult {
	expr := strArg(args, "expression")
	if strings.TrimSpace(expr) == "" {
		return ToolResult{Text: "[erreur] expression vide"}
	}
	file := strings.TrimSpace(strArg(args, "file"))
	inPlace := boolArg(args, "in_place")
	target, cleanup, er, ok := s.sedInputFile(ctx, file, strArg(args, "input"), "sed")
	defer cleanup()
	if !ok {
		return er
	}
	var before []byte
	if inPlace {
		if file == "" {
			return ToolResult{Text: "[erreur] in_place exige file (pas input)"}
		}
		var err error
		before, err = s.fs.ReadFile(ctx, target)
		if err != nil {
			return ToolResult{Text: "[erreur] lecture avant modification : " + err.Error()}
		}
	}
	argv := []string{"sed", "-i", "-e", expr, target}
	if !inPlace {
		argv = []string{"sed", "-e", expr, target}
	}
	out, err := s.runBinary(ctx, argv, toolTimeout(args, bashDefault))
	switch {
	case isTimeoutErr(err):
		return ToolResult{Text: "[timeout] sed interrompu"}
	case errors.Is(ctx.Err(), context.Canceled):
		return ToolResult{Text: "[commande interrompue]"}
	}
	if err != nil {
		return ToolResult{Text: "[erreur] sed (exit " + itoa(exitCodeOf(err)) + ") :\n" + truncate(out, toolMaxOutput)}
	}
	if !inPlace {
		return ToolResult{Text: truncate(strings.TrimSuffix(out, "\n"), toolMaxOutput)}
	}
	after, err := s.fs.ReadFile(ctx, target)
	if err != nil {
		return ToolResult{Text: "[erreur] relecture apres sed -i : " + err.Error()}
	}
	diff := lineDiff(string(before), string(after))
	if len(diff) == 0 {
		return ToolResult{Text: "[ok] sed -i : aucun changement dans " + file}
	}
	return ToolResult{Text: "[ok] sed -i applique : " + file, Diff: truncateDiff(diff, 300)}
}

func (s *Sandbox) toolAwk(ctx context.Context, args map[string]any) ToolResult {
	prog := strArg(args, "program")
	if strings.TrimSpace(prog) == "" {
		return ToolResult{Text: "[erreur] programme vide"}
	}
	target, cleanup, er, ok := s.sedInputFile(ctx, strings.TrimSpace(strArg(args, "file")), strArg(args, "input"), "awk")
	defer cleanup()
	if !ok {
		return er
	}
	argv := []string{"awk"}
	if fs := strings.TrimSpace(strArg(args, "field_separator")); fs != "" {
		argv = append(argv, "-F", fs)
	}
	argv = append(argv, prog, target)
	out, err := s.runBinary(ctx, argv, toolTimeout(args, bashDefault))
	switch {
	case isTimeoutErr(err):
		return ToolResult{Text: "[timeout] awk interrompu"}
	case errors.Is(ctx.Err(), context.Canceled):
		return ToolResult{Text: "[commande interrompue]"}
	}
	if err != nil {
		return ToolResult{Text: "[erreur] awk (exit " + itoa(exitCodeOf(err)) + ") :\n" + truncate(out, toolMaxOutput)}
	}
	return ToolResult{Text: truncate(strings.TrimSuffix(out, "\n"), toolMaxOutput)}
}

func itoa(n int) string { return fmt.Sprintf("%d", n) }

// ---------------- Détection d'effets de bord (heuristiques) ----------------

func isWordByte(c byte) bool {
	return c == '_' || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

// sedRisk analyse heuristiquement une expression sed et signale les
// constructions à effets de bord : exécution de commande (commande `e`,
// flag `e` de s///) et écriture de fichier (commandes w/W, flag w de s///).
// Le reste (s///g, p, d, adresses…) est du flux pur, sans approbation.
func sedRisk(expr string) (exec, write bool) {
	i, n := 0, len(expr)
	for i < n {
		c := expr[i]
		if c == ' ' || c == '\t' || c == '\n' || c == ';' || c == '{' || c == '}' {
			i++
			continue
		}
		// Commande s<delim>re<delim>replacement<delim>flags
		if c == 's' && i+1 < n && !isWordByte(expr[i+1]) && expr[i+1] != '\\' {
			prevOK := i == 0 || !isWordByte(expr[i-1])
			if prevOK {
				d := expr[i+1]
				j := i + 2
				segs := 0
				for j < n && segs < 2 {
					if expr[j] == '\\' {
						j += 2
						continue
					}
					if expr[j] == d {
						segs++
					}
					j++
				}
				if segs == 2 {
					for j < n {
						f := expr[j]
						if f == 'e' {
							exec = true
						}
						if f == 'w' || f == 'W' {
							write = true
						}
						if f == ' ' || f == '\t' || f == '\n' || f == ';' || f == '}' {
							break
						}
						j++
					}
					i = j
					continue
				}
			}
		}
		// Commandes d'un caractère : e (exec), w/W (write).
		if (c == 'e' || c == 'w' || c == 'W') && (i == 0 || !isWordByte(expr[i-1])) {
			// 'w' suivi d'un nom de fichier ; 'e' seul. Évite les faux
			// positifs du type "...where..." : exige un séparateur après.
			next := byte(0)
			if i+1 < n {
				next = expr[i+1]
			}
			if next == 0 || next == ' ' || next == '\t' || next == '\n' || next == ';' || next == '}' {
				if c == 'e' {
					exec = true
				} else {
					write = true
				}
			}
		}
		i++
	}
	return exec, write
}

// awkRiskRe repère les programmes awk à effets de bord : appels system(),
// tubes vers/depuis des commandes, écritures de fichiers (>, >>),
// lectures de fichiers (getline <).
var awkRiskRe = regexp.MustCompile(`\bsystem\s*\(|\|\s*["']|\|\s*getline\b|["']\s*\|\s*getline|\bprint[f]?\b[^|>\n]*>>?|\bgetline\b[^|\n]*<`)

// sedNeedsApproval : in_place exige une approbation (comme Edit) ; en mode
// flux, seules les expressions à effets de bord (exec/écriture) la demandent.
func sedNeedsApproval(args map[string]any) bool {
	if boolArg(args, "in_place") {
		return true
	}
	exec, write := sedRisk(strArg(args, "expression"))
	return exec || write
}

// awkNeedsApproval : approbation si le programme a des effets de bord.
func awkNeedsApproval(args map[string]any) bool {
	return awkRiskRe.MatchString(strArg(args, "program"))
}

// ---------------- Curl ----------------

// toolCurl exécute une requête HTTP(S) encadrée : http/https uniquement
// (redirections filtrées), timeout borné, corps de réponse plafonné à 2 Mo,
// binaires signalés sans être affichés.
func (s *Sandbox) toolCurl(ctx context.Context, args map[string]any) ToolResult {
	rawURL := strings.TrimSpace(strArg(args, "url"))
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return ToolResult{Text: "[erreur] URL http(s) valide requise"}
	}
	method := strings.ToUpper(strings.TrimSpace(strArg(args, "method")))
	if method == "" {
		method = http.MethodGet
	}
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead:
	default:
		return ToolResult{Text: "[erreur] methode non supportee: " + method}
	}
	timeout := intArg(args, "timeout")
	if timeout <= 0 {
		timeout = curlDefaultTO
	}
	if timeout > curlMaxTO {
		timeout = curlMaxTO
	}
	var body io.Reader
	if b := strArg(args, "body"); b != "" {
		if len(b) > 1024*1024 {
			return ToolResult{Text: "[erreur] corps de requete > 1 Mo refuse"}
		}
		body = strings.NewReader(b)
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("trop de redirections (max 5)")
			}
			if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
				return errors.New("redirection hors http(s) interdite")
			}
			return nil
		},
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	req.Header.Set("User-Agent", "cetas-lite-agent/1.0")
	if hdrs, ok := args["headers"].(map[string]any); ok {
		for k, v := range hdrs {
			if sv, ok := v.(string); ok && strings.TrimSpace(k) != "" {
				req.Header.Set(k, sv)
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return ToolResult{Text: "[erreur] requete : " + err.Error()}
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, curlMaxBody+1))
	if err != nil {
		return ToolResult{Text: "[erreur] lecture reponse : " + err.Error()}
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "HTTP %s\n", resp.Status)
	if ct := resp.Header.Get("Content-Type"); ct != "" {
		fmt.Fprintf(&sb, "Content-Type: %s\n", ct)
	}
	sb.WriteString("\n")
	capped := len(data) > curlMaxBody
	if capped {
		data = data[:curlMaxBody]
	}
	if isBinary(data) {
		fmt.Fprintf(&sb, "[contenu binaire : %d octets, non affiche]", len(data))
	} else {
		sb.WriteString(truncate(string(data), toolMaxOutput))
	}
	if capped {
		sb.WriteString("\n… (reponse tronquee a 2 Mo)")
	}
	return ToolResult{Text: sb.String()}
}
