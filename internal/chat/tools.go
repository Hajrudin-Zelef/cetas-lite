package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"cetas-lite/internal/provider"
)

const (
	toolMaxOutput  = 8000
	bashDefault    = 10
	bashMax        = 60
	scriptDefault  = 30
	scriptMax      = 60
	lsMaxEntries   = 500
	globMaxEntries = 500
	grepMaxMatches = 100
	grepMaxFile    = 2 * 1024 * 1024
)

type DiffLine struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

type ToolResult struct {
	Text string
	Diff []DiffLine
}

var execAllowlist = map[string]bool{
	"ls": true, "cat": true, "grep": true, "git": true, "node": true,
	"python3": true, "python": true, "npm": true, "npx": true, "head": true,
	"tail": true, "wc": true, "find": true, "sed": true, "awk": true,
	"echo": true, "printf": true, "mkdir": true, "touch": true, "rm": true,
	"cp": true, "mv": true, "pwd": true, "date": true, "whoami": true,
	"basename": true, "dirname": true, "sleep": true, "true": true, "false": true,
}

var execBannedTokens = map[string]bool{
	"sudo": true, "su": true, "bash": true, "sh": true, "zsh": true,
	"curl": true, "wget": true,
}

var execBannedFlags = map[string]bool{"-c": true, "--eval": true, "-e": true}

func ToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "Ls", Description: "List the workspace file tree. Use this first.", Parameters: str(map[string]any{})}},
		{Type: "function", Function: provider.ToolFunction{Name: "Read", Description: "Read a file, with offset/limit pagination.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string", "description": "Chemin relatif"},
			"offset":    map[string]any{"type": "integer", "description": "Premiere ligne (1-based)"},
			"limit":     map[string]any{"type": "integer", "description": "Nombre max de lignes"},
		}, "file_path")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Write", Description: "Write or overwrite a file.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string"},
			"content":   map[string]any{"type": "string"},
		}, "file_path", "content")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Edit", Description: "Replace one occurrence of text.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string"},
			"old":       map[string]any{"type": "string", "description": "Texte exact a remplacer"},
			"new":       map[string]any{"type": "string", "description": "Texte de remplacement"},
		}, "file_path", "old", "new")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Grep", Description: "Search for a pattern in workspace files.", Parameters: str(map[string]any{
			"pattern": map[string]any{"type": "string"},
			"path":    map[string]any{"type": "string", "description": "Defaut: racine du workspace"},
			"limit":   map[string]any{"type": "integer"},
		}, "pattern")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Glob", Description: "Find files by pattern (e.g. **/*.go).", Parameters: str(map[string]any{
			"pattern": map[string]any{"type": "string"},
		}, "pattern")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Bash", Description: "Run an allowlisted command (no shell, no pipes). Timeout 10s default, 60s max.", Parameters: str(map[string]any{
			"command": map[string]any{"type": "string"},
			"timeout": map[string]any{"type": "integer", "description": "Secondes (defaut 10, max 60)"},
		}, "command")}},
		{Type: "function", Function: provider.ToolFunction{Name: "RunScript", Description: "Run a python/node script in the sandbox. Timeout 30s default, 60s max.", Parameters: str(map[string]any{
			"language": map[string]any{"type": "string", "enum": []string{"python", "node"}},
			"code":     map[string]any{"type": "string"},
			"timeout":  map[string]any{"type": "integer"},
		}, "language", "code")}},
		{Type: "function", Function: provider.ToolFunction{Name: "TodoWrite", Description: "Update the task list (multi-step plan).", Parameters: str(map[string]any{
			"todos": map[string]any{"type": "array", "items": map[string]any{"type": "object", "properties": map[string]any{
				"content": map[string]any{"type": "string"},
				"status":  map[string]any{"type": "string", "enum": []string{"pending", "in_progress", "completed"}},
			}}},
		}, "todos")}},
	}
}

func (s *Sandbox) Execute(ctx context.Context, name, argsJSON string) ToolResult {
	args := map[string]any{}
	if strings.TrimSpace(argsJSON) != "" {
		_ = json.Unmarshal([]byte(argsJSON), &args)
	}
	switch name {
	case "Ls":
		return s.toolLs()
	case "Read":
		return s.toolRead(args)
	case "Write":
		return s.toolWrite(args)
	case "Edit":
		return s.toolEdit(args)
	case "Grep":
		return s.toolGrep(args)
	case "Glob":
		return s.toolGlob(args)
	case "Bash":
		return ToolResult{Text: s.toolBash(ctx, args)}
	case "RunScript":
		return ToolResult{Text: s.toolRunScript(ctx, args)}
	case "TodoWrite":
		return ToolResult{Text: "[ok] liste mise a jour"}
	default:
		return ToolResult{Text: "[erreur] outil inconnu: " + name}
	}
}

func (s *Sandbox) toolLs() ToolResult {
	var entries []string
	_ = filepath.WalkDir(s.root, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if p == s.root {
			return nil
		}
		if d.IsDir() && skipEntry(d.Name()) {
			return filepath.SkipDir
		}
		if skipEntry(d.Name()) {
			return nil
		}
		rel, rerr := filepath.Rel(s.root, p)
		if rerr != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			rel += "/"
		}
		entries = append(entries, rel)
		if len(entries) >= lsMaxEntries {
			return errors.New("stop")
		}
		return nil
	})
	sort.Strings(entries)
	out := fmt.Sprintf("%d entrees", len(entries))
	if len(entries) > 0 {
		out += "\n" + strings.Join(entries, "\n")
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolRead(args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	p, err := s.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return ToolResult{Text: "[erreur] fichier introuvable: " + rel + ". Utilise Glob pour trouver le bon chemin."}
		}
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	lines := strings.Split(string(b), "\n")
	total := len(lines)
	start := 0
	if off := intArg(args, "offset"); off > 0 {
		start = off - 1
	}
	if start > total {
		start = total
	}
	end := total
	if lim := intArg(args, "limit"); lim > 0 && start+lim < end {
		end = start + lim
	}
	out := strings.Join(lines[start:end], "\n")
	if end < total {
		out += fmt.Sprintf("\n… (%d lignes au total, affichees %d-%d)", total, start+1, end)
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolWrite(args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	content := strArg(args, "content")
	if strings.TrimSpace(rel) == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	p, err := s.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if dir := filepath.Dir(p); dir != "" && dir != s.root {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
	}
	mode := os.FileMode(0o644)
	existed := false
	if fi, err := os.Stat(p); err == nil {
		mode = fi.Mode()
		existed = true
	}
	if err := os.WriteFile(p, []byte(content), mode); err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	verb := "cree"
	if existed {
		verb = "reecrit"
	}
	return ToolResult{
		Text: fmt.Sprintf("[ok] %s %s (%d octets)", rel, verb, len(content)),
		Diff: addedDiff(content),
	}
}

func (s *Sandbox) toolEdit(args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	oldText := strArg(args, "old")
	newText := strArg(args, "new")
	if strings.TrimSpace(rel) == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	if oldText == "" {
		return ToolResult{Text: "[erreur] old vide"}
	}
	p, err := s.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	content := string(b)
	n := strings.Count(content, oldText)
	if n == 0 {
		if newText != "" && strings.Contains(content, newText) {
			return ToolResult{Text: "[ok] deja a jour — le fichier contient deja cette modification"}
		}
		return ToolResult{Text: "[erreur] old introuvable dans le fichier"}
	}
	if n > 1 {
		return ToolResult{Text: fmt.Sprintf("[erreur] old apparait %d fois — ajoute du contexte pour le rendre unique", n)}
	}
	updated := strings.Replace(content, oldText, newText, 1)
	mode := os.FileMode(0o644)
	if fi, err := os.Stat(p); err == nil {
		mode = fi.Mode()
	}
	if err := os.WriteFile(p, []byte(updated), mode); err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	return ToolResult{Text: fmt.Sprintf("[ok] %s modifie (1 remplacement)", rel), Diff: lineDiff(oldText, newText)}
}

func (s *Sandbox) toolGrep(args map[string]any) ToolResult {
	pattern := strArg(args, "pattern")
	if pattern == "" {
		return ToolResult{Text: "[erreur] pattern vide"}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ToolResult{Text: "[erreur] regex invalide: " + err.Error()}
	}
	limit := intArg(args, "limit")
	if limit <= 0 {
		limit = grepMaxMatches
	}
	base := s.root
	if rel := strArg(args, "path"); rel != "" && rel != "." {
		p, err := s.Resolve(rel)
		if err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		base = p
	}
	var matches []string
	walkFile := func(p string) {
		if info, err := os.Stat(p); err != nil || info.Size() > grepMaxFile {
			return
		}
		b, err := os.ReadFile(p)
		if err != nil || strings.ContainsRune(string(b), 0) {
			return
		}
		rel, _ := filepath.Rel(s.root, p)
		rel = filepath.ToSlash(rel)
		for i, line := range strings.Split(string(b), "\n") {
			if re.MatchString(line) {
				matches = append(matches, fmt.Sprintf("%s:%d: %s", rel, i+1, line))
				if len(matches) >= limit {
					return
				}
			}
		}
	}
	if info, err := os.Stat(base); err == nil && !info.IsDir() {
		walkFile(base)
	} else {
		_ = filepath.WalkDir(base, func(p string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() && p != base && skipEntry(d.Name()) {
				return filepath.SkipDir
			}
			if !d.IsDir() && !skipEntry(d.Name()) {
				walkFile(p)
				if len(matches) >= limit {
					return errors.New("stop")
				}
			}
			return nil
		})
	}
	out := fmt.Sprintf("%d resultats", len(matches))
	if len(matches) > 0 {
		out += "\n" + strings.Join(matches, "\n")
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolGlob(args map[string]any) ToolResult {
	pattern := strings.TrimSpace(strArg(args, "pattern"))
	if pattern == "" {
		return ToolResult{Text: "[erreur] pattern vide"}
	}
	pattern = strings.TrimPrefix(filepath.ToSlash(pattern), "./")
	var files []string
	_ = filepath.WalkDir(s.root, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if p == s.root {
			return nil
		}
		if d.IsDir() && skipEntry(d.Name()) {
			return filepath.SkipDir
		}
		if d.IsDir() || skipEntry(d.Name()) {
			return nil
		}
		rel, rerr := filepath.Rel(s.root, p)
		if rerr != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		if globMatch(pattern, rel) {
			files = append(files, rel)
			if len(files) >= globMaxEntries {
				return errors.New("stop")
			}
		}
		return nil
	})
	sort.Strings(files)
	if len(files) == 0 {
		return ToolResult{Text: "aucun fichier pour le motif: " + pattern}
	}
	return ToolResult{Text: fmt.Sprintf("%d fichiers pour '%s'\n%s", len(files), pattern, strings.Join(files, "\n"))}
}

func (s *Sandbox) toolBash(ctx context.Context, args map[string]any) string {
	if runtime.GOOS == "windows" {
		return "[erreur] Bash non supporte sur Windows en v1"
	}
	command := strings.TrimSpace(strArg(args, "command"))
	if command == "" {
		return "[erreur] commande vide"
	}
	tokens, err := splitCommand(command)
	if err != nil {
		return "[erreur] " + err.Error()
	}
	if len(tokens) == 0 {
		return "[erreur] commande vide"
	}
	binary := filepath.Base(tokens[0])
	if !execAllowlist[binary] {
		return "[erreur] commande non autorisee: " + tokens[0]
	}
	for _, tok := range tokens {
		if execBannedTokens[tok] {
			return "[erreur] commande interdite: " + tok
		}
	}
	for _, tok := range tokens[1:] {
		if execBannedFlags[tok] {
			return "[erreur] flag interdit: " + tok
		}
		if msg := s.checkArg(tok); msg != "" {
			return "[erreur] " + msg
		}
	}
	timeout := intArg(args, "timeout")
	if timeout <= 0 {
		timeout = bashDefault
	}
	if timeout > bashMax {
		timeout = bashMax
	}
	cctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(cctx, tokens[0], tokens[1:]...)
	if err := os.MkdirAll(s.root, 0o700); err == nil {
		cmd.Dir = s.root
	}
	cmd.WaitDelay = 2 * time.Second
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	switch {
	case errors.Is(cctx.Err(), context.DeadlineExceeded):
		return fmt.Sprintf("[timeout apres %ds]", timeout)
	case errors.Is(ctx.Err(), context.Canceled):
		return "[commande interrompue]"
	}
	exit := 0
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			exit = ee.ExitCode()
		} else {
			return fmt.Sprintf("[erreur] %v", err)
		}
	}
	parts := []string{fmt.Sprintf("exit: %d", exit)}
	if out := truncate(stdout.String(), toolMaxOutput); out != "" {
		parts = append(parts, "stdout:\n"+out)
	}
	if out := truncate(stderr.String(), toolMaxOutput); out != "" {
		parts = append(parts, "stderr:\n"+out)
	}
	return strings.Join(parts, "\n\n")
}

func filterTools(tools []provider.Tool, drop string) []provider.Tool {
	out := make([]provider.Tool, 0, len(tools))
	for _, t := range tools {
		if t.Function.Name == drop {
			continue
		}
		out = append(out, t)
	}
	return out
}

func (s *Sandbox) checkArg(tok string) string {
	if tok == "" {
		return ""
	}
	norm := strings.ReplaceAll(tok, "\\", "/")
	if strings.Contains(norm, "~") {
		return "argument hors sandbox: " + tok
	}
	if strings.Contains(norm, "=/") || (len(norm) > 2 && norm[0] == '-' && norm[1] != '/' && norm[2] == '/') {
		return "argument hors sandbox: " + tok
	}
	if strings.HasPrefix(norm, "-") {
		if strings.Contains(norm, "..") {
			return "argument hors sandbox: " + tok
		}
		return ""
	}
	if strings.HasPrefix(norm, "/") {
		return "argument absolu interdit: " + tok
	}
	if !strings.Contains(norm, "/") && !strings.Contains(norm, "..") {
		return ""
	}
	clean := path.Clean(norm)
	if clean == ".." || strings.HasPrefix(clean, "../") {
		return "argument hors sandbox: " + tok
	}
	if _, err := s.Resolve(norm); err != nil {
		return "argument hors sandbox: " + tok
	}
	return ""
}

func (s *Sandbox) toolRunScript(ctx context.Context, args map[string]any) string {
	if !s.AllowScript {
		return "[erreur] RunScript desactive (definir CETAS_LITE_ALLOW_SCRIPT=1 pour l'activer)"
	}
	lang := strings.ToLower(strings.TrimSpace(strArg(args, "language")))
	runners := map[string]string{"python": "python3", "node": "node"}
	runner, ok := runners[lang]
	if !ok {
		return "[erreur] Unsupported language: " + lang
	}
	code := strArg(args, "code")
	timeout := intArg(args, "timeout")
	if timeout <= 0 {
		timeout = scriptDefault
	}
	if timeout > scriptMax {
		timeout = scriptMax
	}
	tmpDir := filepath.Join(s.root, ".runscript_tmp")
	if err := os.MkdirAll(tmpDir, 0o700); err != nil {
		return "[erreur] preparation sandbox impossible: " + err.Error()
	}
	ext := "py"
	if lang == "node" {
		ext = "js"
	}
	f, err := os.CreateTemp(tmpDir, "script-*."+ext)
	if err != nil {
		return "[erreur] ecriture script impossible: " + err.Error()
	}
	fpath := f.Name()
	if _, err := f.WriteString(code); err != nil {
		f.Close()
		os.Remove(fpath)
		return "[erreur] ecriture script impossible: " + err.Error()
	}
	f.Close()
	defer os.Remove(fpath)

	cctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(cctx, runner, fpath)
	cmd.Dir = s.root
	cmd.WaitDelay = 2 * time.Second
	cmd.Env = []string{
		"PATH=" + envOr("PATH", "/usr/local/bin:/usr/bin:/bin"),
		"HOME=" + s.root,
		"LANG=C.UTF-8",
	}
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	switch {
	case errors.Is(cctx.Err(), context.DeadlineExceeded):
		return fmt.Sprintf("[timeout apres %ds]", timeout)
	case errors.Is(ctx.Err(), context.Canceled):
		return "[commande interrompue]"
	}
	exit := 0
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			exit = ee.ExitCode()
		} else {
			return fmt.Sprintf("[erreur] %v", err)
		}
	}
	parts := []string{fmt.Sprintf("exit: %d", exit)}
	if out := truncate(stdout.String(), toolMaxOutput); out != "" {
		parts = append(parts, "stdout:\n"+out)
	}
	if out := truncate(stderr.String(), toolMaxOutput); out != "" {
		parts = append(parts, "stderr:\n"+out)
	}
	return strings.Join(parts, "\n\n")
}

func skipEntry(name string) bool {
	return name == ".git" || name == "node_modules" || name == ".runscript_tmp" || strings.HasPrefix(name, ".")
}

func globMatch(pattern, name string) bool {
	pp := strings.Split(pattern, "/")
	nn := strings.Split(name, "/")
	return matchSegments(pp, nn)
}

func matchSegments(pp, nn []string) bool {
	for len(pp) > 0 {
		if pp[0] == "**" {
			if len(pp) == 1 {
				return true
			}
			for i := 0; i <= len(nn); i++ {
				if matchSegments(pp[1:], nn[i:]) {
					return true
				}
			}
			return false
		}
		if len(nn) == 0 {
			return false
		}
		ok, err := path.Match(pp[0], nn[0])
		if err != nil || !ok {
			return false
		}
		pp, nn = pp[1:], nn[1:]
	}
	return len(nn) == 0
}

func addedDiff(content string) []DiffLine {
	var out []DiffLine
	for _, line := range strings.Split(content, "\n") {
		out = append(out, DiffLine{Kind: "+", Text: line})
	}
	return out
}

func lineDiff(oldText, newText string) []DiffLine {
	var out []DiffLine
	for _, line := range strings.Split(oldText, "\n") {
		out = append(out, DiffLine{Kind: "-", Text: line})
	}
	for _, line := range strings.Split(newText, "\n") {
		out = append(out, DiffLine{Kind: "+", Text: line})
	}
	return out
}

func splitCommand(s string) ([]string, error) {
	var tokens []string
	var cur strings.Builder
	inSingle, inDouble, esc := false, false, false
	has := false
	for _, r := range s {
		switch {
		case esc:
			cur.WriteRune(r)
			esc = false
			has = true
		case r == '\\' && !inSingle:
			esc = true
			has = true
		case r == '\'' && !inDouble:
			inSingle = !inSingle
			has = true
		case r == '"' && !inSingle:
			inDouble = !inDouble
			has = true
		case (r == ' ' || r == '\t') && !inSingle && !inDouble:
			if has {
				tokens = append(tokens, cur.String())
				cur.Reset()
				has = false
			}
		default:
			cur.WriteRune(r)
			has = true
		}
	}
	if inSingle || inDouble {
		return nil, errors.New("guillemet non ferme")
	}
	if esc {
		return nil, errors.New("echappement non termine")
	}
	if has {
		tokens = append(tokens, cur.String())
	}
	return tokens, nil
}

func truncate(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "\n…[tronque]"
}

func strArg(args map[string]any, key string) string {
	s, _ := args[key].(string)
	return s
}

func intArg(args map[string]any, key string) int {
	switch v := args[key].(type) {
	case float64:
		return int(v)
	case int:
		return v
	case json.Number:
		n, _ := v.Int64()
		return int(n)
	case string:
		n, _ := strconv.Atoi(v)
		return n
	}
	return 0
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
