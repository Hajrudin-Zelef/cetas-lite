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
	"cetas-lite/internal/vfs"
)

const (
	toolMaxOutput = 8000
	// Lignes max renvoyées par Read quand aucun "limit" n'est précisé.
	defaultReadLines = 400
	bashDefault      = 10
	bashMax          = 60
	scriptDefault    = 30
	scriptMax        = 60
	lsMaxEntries     = 500
	globMaxEntries   = 500
	grepMaxMatches   = 100
	grepMaxFile      = 2 * 1024 * 1024
)

type DiffLine struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

type ToolResult struct {
	Text string
	Diff []DiffLine
	// Meta : donnees structurees jointes au delta SSE (ex. sources d'une
	// recherche web : {"sources": [{"title","url"}], "search_provider": ...}).
	Meta map[string]any
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

var execBannedFlags = map[string]bool{"-c": true, "--eval": true, "-e": true, "-exec": true, "-execdir": true}

// execEvalFlagPrefixes : par binaire, les préfixes de flags d'évaluation
// de code interdits. Le contrôle d'égalité exacte laissait passer les
// formes collées (python3 -c<charge>, node --eval=<charge>) : pour ces
// binaires on contrôle par préfixe. Scopé par binaire pour ne pas casser
// les usages légitimes (grep -c, sed -e... restent soumis au seul
// contrôle exact ci-dessus).
var execEvalFlagPrefixes = map[string][]string{
	"python":  {"-c"},
	"python3": {"-c"},
	"node":    {"-e", "--eval"},
}

// bannedEvalFlag détecte un flag d'évaluation de code, y compris collé à
// sa charge (python3 -cprint(1), node --eval=...).
func bannedEvalFlag(binary, tok string) bool {
	for _, pfx := range execEvalFlagPrefixes[binary] {
		if strings.HasPrefix(tok, pfx) {
			return true
		}
	}
	return false
}

// awkBashBinaries : variantes d'awk admises dans Bash.
var awkBashBinaries = map[string]bool{"awk": true, "gawk": true, "mawk": true, "nawk": true}

// bashEmbeddedProgram reconstitue le programme sed/awk embarqué dans une
// commande Bash : les options sont ignorées (y compris -e<script> collé),
// ainsi que l'argument des options qui en prennent un (awk -v var=val,
// -F separateur, -f fichier ; sed -f fichier) — sans quoi cet argument
// serait pris pour le programme et masquerait le vrai programme risqué.
// Le premier opérande non-option restant est le programme (la suite =
// fichiers).
func bashEmbeddedProgram(binary string, args []string) string {
	var takesArg map[string]bool
	switch {
	case awkBashBinaries[binary]:
		takesArg = map[string]bool{"-v": true, "-F": true, "-f": true, "-W": true}
	case binary == "sed":
		takesArg = map[string]bool{"-f": true}
	}
	var parts []string
	skipNext := false
	for _, t := range args {
		if t == "" {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}
		if strings.HasPrefix(t, "-") {
			if rest, ok := strings.CutPrefix(t, "-e"); ok && rest != "" {
				parts = append(parts, rest)
			} else if rest, ok := strings.CutPrefix(t, "--expression="); ok && rest != "" {
				parts = append(parts, rest)
			} else if takesArg[t] {
				skipNext = true
			}
			// Sinon : option simple ou collée (-F:, -n...) ignorée.
			continue
		}
		if len(parts) == 0 {
			parts = append(parts, t)
		}
		break
	}
	return strings.Join(parts, "\n")
}

// bashProgramRisk applique aux programmes sed/awk embarqués via Bash les
// mêmes détections que les outils dédiés (sedRisk, awkRiskRe).
func bashProgramRisk(binary, prog string) string {
	switch {
	case binary == "sed":
		if exec, write := sedRisk(prog); exec || write {
			return "programme sed a effets de bord detecte : utilise l'outil Sed dedie (approbation requise)"
		}
	case awkBashBinaries[binary]:
		if awkRiskRe.MatchString(prog) {
			return "programme awk a effets de bord detecte : utilise l'outil Awk dedie (approbation requise)"
		}
	}
	return ""
}

// gitShowsConfig détecte les commandes git qui affichent la configuration
// (donc l'extraHeader d'authentification) : leur sortie est expurgée du token.
func gitShowsConfig(binary string, tokens []string) bool {
	if binary != "git" || len(tokens) < 2 {
		return false
	}
	// La sous-commande peut être précédée d'options globales
	// (git -c key=val config --list) : on les saute pour la trouver.
	i := 1
	for i < len(tokens) && strings.HasPrefix(tokens[i], "-") {
		if tokens[i] == "-c" || tokens[i] == "-C" {
			i++ // -c/-C prennent une valeur
		}
		i++
	}
	if i >= len(tokens) {
		return false
	}
	rest := tokens[i+1:]
	switch tokens[i] {
	case "config":
		for _, t := range rest {
			if t == "--list" || t == "-l" {
				return true
			}
		}
	case "var":
		for _, t := range rest {
			if t == "-l" {
				return true
			}
		}
	}
	return false
}

// toolAliases : outils historiques fusionnés -> outil canonique (phase 3).
// Le modèle ne voit que le canonique dans les schémas, mais tout appel
// résiduel avec l'ancien nom (natif, DSML, pseudo-texte) est réécrit.
var toolAliases = map[string]string{
	"Cat": "Read",
}

// canonicalToolName réécrit le nom d'un outil fusionné vers son canonique
// (insensible à la casse). Retourne le nom inchangé sinon.
func canonicalToolName(name string) string {
	for old, canon := range toolAliases {
		if strings.EqualFold(old, name) {
			return canon
		}
	}
	return name
}

func ToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	out := []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "Ls", Description: "List the workspace tree. Call me first.", Parameters: str(map[string]any{})}},
		{Type: "function", Function: provider.ToolFunction{Name: "Read", Description: "Read a file (offset/limit pagination).", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string", "description": "Relative path"},
			"offset":    map[string]any{"type": "integer", "description": "First line (1-based)"},
			"limit":     map[string]any{"type": "integer", "description": "Max number of lines"},
		}, "file_path")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Write", Description: "Write or overwrite a file.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string"},
			"content":   map[string]any{"type": "string"},
		}, "file_path", "content")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Edit", Description: "Replace one occurrence of text.", Parameters: str(map[string]any{
			"file_path": map[string]any{"type": "string"},
			"old":       map[string]any{"type": "string", "description": "Exact text to replace"},
			"new":       map[string]any{"type": "string", "description": "Replacement text"},
		}, "file_path", "old", "new")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Grep", Description: "Search for a pattern in workspace files.", Parameters: str(map[string]any{
			"pattern": map[string]any{"type": "string"},
			"path":    map[string]any{"type": "string", "description": "Default: workspace root"},
			"limit":   map[string]any{"type": "integer"},
		}, "pattern")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Glob", Description: "Find files by pattern (e.g. **/*.go).", Parameters: str(map[string]any{
			"pattern": map[string]any{"type": "string"},
		}, "pattern")}},
		{Type: "function", Function: provider.ToolFunction{Name: "Bash", Description: "Run an allowed command (no shell, no pipes). Default 10s timeout, 60s max.", Parameters: str(map[string]any{
			"command": map[string]any{"type": "string"},
			"timeout": map[string]any{"type": "integer", "description": "Seconds (default 10, max 60)"},
		}, "command")}},
		{Type: "function", Function: provider.ToolFunction{Name: "RunScript", Description: "Run a python/node script in the sandbox. Default 30s timeout, 60s max.", Parameters: str(map[string]any{
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
	return append(append(out, extraToolSchemas()...), githubToolSchemas()...)
}

// requiredFields centralise les champs requis par outil. Source unique
// utilisee par Execute et par la validation pre-approbation (agent.go) :
// un appel dont un champ requis manque ou est vide ne doit jamais arriver
// jusqu'a la carte d'approbation utilisateur.
func requiredFields(name string) []string {
	switch name {
	case "Read", "Cat":
		return []string{"file_path"}
	case "Write":
		return []string{"file_path", "content"}
	case "Edit":
		return []string{"file_path", "old", "new"}
	case "Grep", "Glob":
		return []string{"pattern"}
	case "Bash":
		return []string{"command"}
	case "RunScript":
		return []string{"language", "code"}
	case "Echo":
		return []string{"text"}
	case "Mkdir":
		return []string{"path"}
	case "Mv":
		return []string{"src", "dst"}
	case "Sed":
		return []string{"expression"}
	case "Awk":
		return []string{"program"}
	case "Curl":
		return []string{"url"}
	}
	return nil
}

func (s *Sandbox) Execute(ctx context.Context, name, argsJSON string) ToolResult {
	args := map[string]any{}
	if strings.TrimSpace(argsJSON) != "" {
		_ = json.Unmarshal([]byte(argsJSON), &args)
	}
	switch name {
	case "Ls":
		return s.toolLs(ctx)
	case "Read":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolRead(ctx, args)
	case "Write":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolWrite(ctx, args)
	case "Edit":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolEdit(ctx, args)
	case "Grep":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolGrep(ctx, args)
	case "Glob":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolGlob(ctx, args)
	case "Bash":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return ToolResult{Text: s.toolBash(ctx, args)}
	case "RunScript":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return ToolResult{Text: s.toolRunScript(ctx, args)}
	case "TodoWrite":
		return ToolResult{Text: "[ok] liste mise a jour"}
	case "Tree":
		return s.toolTree(ctx, args)
	case "Cat":
		// Phase 3 : Cat est fusionné dans Read et n'est plus annoncé au
		// modèle. Filet de sécurité : un appel résiduel est réécrit en
		// Read (head -> limit ; tail -> dernières lignes).
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolCatAsRead(ctx, args)
	case "Echo":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolEcho(ctx, args)
	case "Mkdir":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolMkdir(ctx, args)
	case "Mv":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolMv(ctx, args)
	case "Sed":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolSed(ctx, args)
	case "Awk":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolAwk(ctx, args)
	case "Curl":
		if f := missingArg(args, requiredFields(name)...); f != "" {
			return ToolResult{Text: missingArgErr(name, f)}
		}
		return s.toolCurl(ctx, args)
	case "GitHubRepos", "GitHubIssues", "GitHubIssueGet", "GitHubPRs",
		"GitHubRepoCreate", "GitHubIssueCreate", "GitHubIssueComment",
		"GitHubPRCreate", "GitHubPRMerge":
		return s.toolGitHub(ctx, name, args)
	default:
		return ToolResult{Text: "[erreur] outil inconnu: " + name}
	}
}

// missingArg retourne le premier champ requis absent ou vide.
func missingArg(args map[string]any, fields ...string) string {
	for _, f := range fields {
		v, ok := args[f]
		if !ok {
			return f
		}
		if s, ok := v.(string); ok && strings.TrimSpace(s) == "" {
			return f
		}
	}
	return ""
}

func missingArgErr(tool, field string) string {
	return "[erreur] " + tool + " : champ requis manquant ou vide : \"" + field +
		"\" — renvoie l'appel avec ce champ renseigne."
}

func (s *Sandbox) toolLs(ctx context.Context) ToolResult {
	var entries []string
	walkErr := s.fs.Walk(ctx, func(e vfs.Entry) error {
		entries = append(entries, e.Path)
		if len(entries) >= lsMaxEntries {
			return errWalkStopSandbox
		}
		return nil
	})
	sort.Strings(entries)
	out := fmt.Sprintf("%d entrees", len(entries))
	if len(entries) > 0 {
		out += "\n" + strings.Join(entries, "\n")
	}
	// B3 : erreur de parcours remontée au modèle, pas ignorée.
	if walkErr != nil && !errors.Is(walkErr, errWalkStopSandbox) {
		out += "\n[avertissement] parcours incomplet"
	}
	return ToolResult{Text: out}
}

var errWalkStopSandbox = errors.New("stop")

func (s *Sandbox) toolRead(ctx context.Context, args map[string]any) ToolResult {
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
	lim := intArg(args, "limit")
	switch {
	case lim > 0 && start+lim < end:
		end = start + lim
	case lim <= 0 && end-start > defaultReadLines:
		// Garde-fou : un Read sans limite sur un gros fichier noierait le
		// contexte du tour (tokens). Le compteur ci-dessous indique au
		// modèle comment paginer avec offset/limit.
		end = start + defaultReadLines
	}
	out := strings.Join(lines[start:end], "\n")
	if end < total {
		out += fmt.Sprintf("\n… (%d lignes au total, affichees %d-%d)", total, start+1, end)
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolWrite(ctx context.Context, args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	content := strArg(args, "content")
	if strings.TrimSpace(rel) == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	clean, err := s.fs.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if clean == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	mode := os.FileMode(0o644)
	existed := false
	if _, err := s.fs.Stat(ctx, clean); err == nil {
		existed = true
	}
	if err := s.fs.WriteFile(ctx, clean, []byte(content), mode); err != nil {
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

func (s *Sandbox) toolEdit(ctx context.Context, args map[string]any) ToolResult {
	rel := strArg(args, "file_path")
	oldText := strArg(args, "old")
	newText := strArg(args, "new")
	if strings.TrimSpace(rel) == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	if oldText == "" {
		return ToolResult{Text: "[erreur] old vide"}
	}
	clean, err := s.fs.Resolve(rel)
	if err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	if clean == "" {
		return ToolResult{Text: "[erreur] chemin vide"}
	}
	b, err := s.fs.ReadFile(ctx, clean)
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
	if err := s.fs.WriteFile(ctx, clean, []byte(updated), 0o644); err != nil {
		return ToolResult{Text: "[erreur] " + err.Error()}
	}
	return ToolResult{Text: fmt.Sprintf("[ok] %s modifie (1 remplacement)", rel), Diff: lineDiff(oldText, newText)}
}

func (s *Sandbox) toolGrep(ctx context.Context, args map[string]any) ToolResult {
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
	base := ""
	if rel := strArg(args, "path"); rel != "" && rel != "." {
		clean, err := s.fs.Resolve(rel)
		if err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		base = clean
	}
	var matches []string
	stopped := false
	walkFile := func(rel string) {
		if stopped {
			return
		}
		fi, err := s.fs.Stat(ctx, rel)
		if err != nil || fi.IsDir || fi.Size > grepMaxFile {
			return
		}
		b, err := s.fs.ReadFile(ctx, rel)
		if err != nil || strings.ContainsRune(string(b), 0) {
			return
		}
		for i, line := range strings.Split(string(b), "\n") {
			if re.MatchString(line) {
				matches = append(matches, fmt.Sprintf("%s:%d: %s", rel, i+1, line))
				if len(matches) >= limit {
					stopped = true
					return
				}
			}
		}
	}
	if base != "" {
		if fi, err := s.fs.Stat(ctx, base); err == nil && !fi.IsDir {
			walkFile(base)
		} else {
			_ = s.fs.Walk(ctx, func(e vfs.Entry) error {
				if base != "" && e.Path != base && !strings.HasPrefix(e.Path, base+"/") {
					return nil
				}
				if !e.IsDir {
					walkFile(strings.TrimSuffix(e.Path, "/"))
				}
				if stopped {
					return errWalkStopSandbox
				}
				return nil
			})
		}
	} else {
		walkErr := s.fs.Walk(ctx, func(e vfs.Entry) error {
			if !e.IsDir {
				walkFile(strings.TrimSuffix(e.Path, "/"))
			}
			if stopped {
				return errWalkStopSandbox
			}
			return nil
		})
		// B3 : erreur de parcours remontée au modèle, pas ignorée.
		if walkErr != nil && !errors.Is(walkErr, errWalkStopSandbox) {
			matches = append(matches, "[avertissement] parcours incomplet")
		}
	}
	out := fmt.Sprintf("%d resultats", len(matches))
	if len(matches) > 0 {
		out += "\n" + strings.Join(matches, "\n")
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolGlob(ctx context.Context, args map[string]any) ToolResult {
	pattern := strings.TrimSpace(strArg(args, "pattern"))
	if pattern == "" {
		return ToolResult{Text: "[erreur] pattern vide"}
	}
	// B2 : un motif comme **/**/** provoque une explosion combinatoire
	// dans matchSegments — on borne le nombre de segments **.
	if n := countGlobStars(pattern); n > maxGlobStars {
		return ToolResult{Text: fmt.Sprintf("[erreur] motif trop complexe : %d segments ** (max %d)", n, maxGlobStars)}
	}
	pattern = strings.TrimPrefix(filepath.ToSlash(pattern), "./")
	var files []string
	walkErr := s.fs.Walk(ctx, func(e vfs.Entry) error {
		if e.IsDir {
			return nil
		}
		rel := strings.TrimSuffix(e.Path, "/")
		if globMatch(pattern, rel) {
			files = append(files, rel)
			if len(files) >= globMaxEntries {
				return errWalkStopSandbox
			}
		}
		return nil
	})
	sort.Strings(files)
	var out string
	if len(files) == 0 {
		out = "aucun fichier pour le motif: " + pattern
	} else {
		out = fmt.Sprintf("%d fichiers pour '%s'\n%s", len(files), pattern, strings.Join(files, "\n"))
	}
	// B3 : une erreur de parcours n'est plus silencieuse — avertissement
	// au modèle (l'arrêt volontaire sur quota n'en est pas une).
	if walkErr != nil && !errors.Is(walkErr, errWalkStopSandbox) {
		out += "\n[avertissement] parcours incomplet"
	}
	return ToolResult{Text: out}
}

func (s *Sandbox) toolBash(ctx context.Context, args map[string]any) string {
	if runtime.GOOS == "windows" && !s.fs.Remote() {
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
	// H6 : seul le nom de base était contrôlé mais le chemin complet était
	// exécuté — outils/ls pouvait être un script malveillant du workspace.
	// On impose le nom simple (résolution PATH), jamais un chemin.
	if strings.Contains(tokens[0], "/") {
		return "[erreur] chemin de binaire interdit (utilise le nom simple, ex. \"ls\"): " + tokens[0]
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
		// H1 : formes collées (python3 -c<charge>, node --eval=<charge>).
		if bannedEvalFlag(binary, tok) {
			return "[erreur] flag interdit: " + tok
		}
		if msg := s.checkArg(tok); msg != "" {
			return "[erreur] " + msg
		}
	}
	// H2/H5/H7 : un programme sed/awk à effets de bord passé via Bash
	// échappait aux détections des outils dédiés. On applique ici les mêmes
	// heuristiques : refusé dans Bash, le modèle passe par l'outil dédié
	// (qui déclenche l'approbation).
	if msg := bashProgramRisk(binary, bashEmbeddedProgram(binary, tokens[1:])); msg != "" {
		return "[erreur] " + msg
	}
	timeout := intArg(args, "timeout")
	if timeout <= 0 {
		timeout = bashDefault
	}
	if timeout > bashMax {
		timeout = bashMax
	}
	// Environnement git : authentifie les opérations GitHub via le token
	// connecté, sans modifier aucun fichier du dépôt.
	var env map[string]string
	if binary == "git" {
		env = gitEnv(s.githubToken())
	}
	argv := tokens
	if !s.fs.Remote() {
		argv = wrapCommand(s.Isolation, s.localRoot(), tokens)
	}
	out, err := s.fs.Exec(ctx, argv[0], argv[1:], env, time.Duration(timeout)*time.Second)
	switch {
	case isTimeoutErr(err):
		return fmt.Sprintf("[timeout apres %ds]", timeout)
	case errors.Is(ctx.Err(), context.Canceled):
		return "[commande interrompue]"
	}
	exit := exitCodeOf(err)
	parts := []string{fmt.Sprintf("exit: %d", exit)}
	if out := truncate(out, toolMaxOutput); out != "" {
		parts = append(parts, "sortie:\n"+out)
	}
	if err != nil && exit == 0 {
		parts = append(parts, fmt.Sprintf("[erreur] %v", err))
	}
	result := strings.Join(parts, "\n\n")
	// H4 : le token GitHub ne remonte jamais au modèle, même via
	// `git config --list` (l'extraHeader l'y affiche en clair).
	if tok := s.githubToken(); tok != "" && gitShowsConfig(binary, tokens) {
		result = strings.ReplaceAll(result, tok, "***")
	}
	return result
}

// isTimeoutErr détecte un dépassement de délai renvoyé par un FS.
func isTimeoutErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "delai depasse")
}

// exitCodeOf extrait le code de sortie d'une commande (locale ou SSH).
func exitCodeOf(err error) int {
	if err == nil {
		return 0
	}
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		return ee.ExitCode()
	}
	// *ssh.ExitError expose ExitStatus() sans hériter d'exec.ExitError.
	type exitStater interface{ ExitStatus() int }
	var es exitStater
	if errors.As(err, &es) {
		return es.ExitStatus()
	}
	return 1
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
	ext := "py"
	if lang == "node" {
		ext = "js"
	}
	// Le script est déposé dans le FS (local ou distant) puis exécuté
	// via le même chemin que Bash : comportement identique partout.
	rel := fmt.Sprintf(".runscript_tmp/script-%d.%s", time.Now().UnixNano(), ext)
	if err := s.fs.WriteFile(ctx, rel, []byte(code), 0o600); err != nil {
		return "[erreur] ecriture script impossible: " + err.Error()
	}
	defer s.fs.Remove(ctx, rel)
	env := map[string]string{
		"PATH": envOr("PATH", "/usr/local/bin:/usr/bin:/bin"),
		"LANG": "C.UTF-8",
	}
	if !s.fs.Remote() {
		env["HOME"] = s.localRoot()
	}
	argv := []string{runner, rel}
	if !s.fs.Remote() {
		argv = wrapCommand(s.Isolation, s.localRoot(), argv)
	}
	out, err := s.fs.Exec(ctx, argv[0], argv[1:], env, time.Duration(timeout)*time.Second)
	switch {
	case isTimeoutErr(err):
		return fmt.Sprintf("[timeout apres %ds]", timeout)
	case errors.Is(ctx.Err(), context.Canceled):
		return "[commande interrompue]"
	}
	parts := []string{fmt.Sprintf("exit: %d", exitCodeOf(err))}
	if out := truncate(out, toolMaxOutput); out != "" {
		parts = append(parts, "sortie:\n"+out)
	}
	if err != nil && exitCodeOf(err) == 0 {
		parts = append(parts, fmt.Sprintf("[erreur] %v", err))
	}
	return strings.Join(parts, "\n\n")
}

func skipEntry(name string) bool {
	return vfs.SkipEntry(name)
}

// maxGlobStars borne les segments ** d'un motif Glob (B2 : explosion
// combinatoire dans matchSegments au-delà).
const maxGlobStars = 8

// countGlobStars compte les segments exactement "**" d'un motif.
func countGlobStars(pattern string) int {
	n := 0
	for _, seg := range strings.Split(pattern, "/") {
		if seg == "**" {
			n++
		}
	}
	return n
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

// toolModelMaxChars borne la taille d'un résultat d'outil injecté dans le
// contexte du modèle. Sans borne, un simple Read sur un gros README fait
// exploser les tokens d'entrée (constaté : 14k tokens pour 2 requêtes
// simples). Le modèle peut toujours paginer (Read offset/limit) ou affiner
// sa requête pour obtenir la suite. L'affichage UI garde sa propre borne
// (toolMaxOutput) : les deux consommateurs sont indépendants.
const toolModelMaxChars = 6000

// Phase 2 : troncature intelligente. Au lieu de couper brutalement en tête,
// on conserve le début ET la fin du résultat (la fin contient souvent la
// conclusion ou le statut), avec un marqueur indiquant ce qui a été omis.
// Le découpage se fait par lignes quand c'est pertinent (Read, Grep, Bash),
// avec repli caractère pour les textes sans retours à la ligne (JSON
// minifié...). Le total reste sous toolModelMaxChars.
func truncateToolForModel(s string) string {
	r := []rune(s)
	if len(r) <= toolModelMaxChars {
		return s
	}
	const headBudget = 4000
	const tailBudget = 1500
	hint := "pagine avec Read offset/limit ou affine ta requete pour obtenir la suite"
	if lines := strings.Split(s, "\n"); len(lines) > 1 {
		var head []string
		used := 0
		for _, ln := range lines {
			if w := len([]rune(ln)) + 1; used+w > headBudget && len(head) > 0 {
				break
			} else {
				head = append(head, ln)
				used += w
			}
		}
		var tail []string
		used = 0
		for i := len(lines) - 1; i >= 0; i-- {
			if w := len([]rune(lines[i])) + 1; used+w > tailBudget && len(tail) > 0 {
				break
			} else {
				tail = append([]string{lines[i]}, tail...)
				used += w
			}
		}
		if omitted := len(lines) - len(head) - len(tail); omitted > 0 {
			marker := fmt.Sprintf("\n…[%d lignes omises — resultat tronque pour limiter le contexte ; %s]\n", omitted, hint)
			return strings.Join(head, "\n") + marker + strings.Join(tail, "\n")
		}
		// Chevauchement tête/queue (peu de lignes très longues) : repli caractère.
	}
	omitted := len(r) - headBudget - tailBudget
	return string(r[:headBudget]) +
		fmt.Sprintf("\n…[%d caracteres omis — resultat tronque pour limiter le contexte ; %s]\n", omitted, hint) +
		string(r[len(r)-tailBudget:])
}

// Phase 2 : erreurs d'outils uniformes et actionnables. Toute erreur remonte
// au modèle sous la forme "[erreur] <outil> : <cause> — <consigne>", pour
// qu'il se corrige en UN tour au lieu de dériver. Les messages déjà
// explicites (nom de l'outil ou consigne de correction présents) sont
// laissés intacts ; seuls les messages bruts (ex. erreur système nue)
// sont enrichis. Appliqué au point de passage unique toolRegistry.execute.
func uniformToolError(tool string, tr ToolResult) ToolResult {
	const pfx = "[erreur]"
	t := tr.Text
	if !strings.HasPrefix(t, pfx) {
		return tr
	}
	rest := strings.TrimSpace(strings.TrimPrefix(t, pfx))
	lower := strings.ToLower(rest)
	if strings.Contains(rest, tool) ||
		strings.Contains(lower, "renvoie") || strings.Contains(lower, "corrige") ||
		strings.Contains(lower, "utilise") || strings.Contains(lower, "ajoute") {
		return tr // déjà uniforme et actionnable
	}
	tr.Text = pfx + " " + tool + " : " + rest + " — corrige les arguments et renvoie l'appel."
	return tr
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
