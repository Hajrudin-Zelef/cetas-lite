package chat

// Outils GitHub de l'agent : repos, issues et pull requests via l'API REST
// GitHub, authentifiés avec le token du connecteur GitHub (chiffré dans le
// coffre, exposé par Sandbox.GitHubToken).
//
// Lecture (sans approbation, autorisés en mode plan) :
//   GitHubRepos, GitHubIssues, GitHubIssueGet, GitHubPRs
// Écriture (approbation requise, refusées en mode plan) :
//   GitHubRepoCreate, GitHubIssueCreate, GitHubIssueComment,
//   GitHubPRCreate, GitHubPRMerge
//
// Sans compte GitHub connecté, chaque outil renvoie une erreur explicite qui
// guide vers Configuration → Connecteurs → GitHub.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cetas-lite/internal/provider"
)

const (
	githubDefaultTimeout = 15 * time.Second
	githubMaxBody        = 2 * 1024 * 1024
	githubAPIVersion     = "2022-11-28"
	githubBodyPreview    = 1500
)

// githubAPIBase est la base de l'API GitHub ; variable (et non constante)
// pour permettre aux tests d'injecter un serveur local.
var githubAPIBase = "https://api.github.com"

// githubNameRe valide owner/repo : lettres, chiffres, -, _ et .
// (empêche toute injection de chemin dans l'URL d'API).
var githubNameRe = regexp.MustCompile(`^[A-Za-z0-9_.\-]+$`)

// githubToolSchemas déclare les schémas des outils GitHub.
func githubToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	owner := map[string]any{"type": "string", "description": "Repository owner (user or org)"}
	repo := map[string]any{"type": "string", "description": "Repository name"}
	state := map[string]any{"type": "string", "description": "open, closed or all (default open)", "enum": []string{"open", "closed", "all"}}
	limit := map[string]any{"type": "integer", "description": "Max results (default 30, max 100)"}
	number := map[string]any{"type": "integer", "description": "Issue or pull request number"}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubRepos", Description: "Liste tes dépôts GitHub (compte connecté), les plus récents d'abord. Lecture seule.", Parameters: str(map[string]any{
			"limit": limit,
		})}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubIssues", Description: "Liste les issues d'un dépôt GitHub. Lecture seule.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo, "state": state, "limit": limit,
		}, "owner", "repo")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubIssueGet", Description: "Affiche une issue GitHub avec ses commentaires. Lecture seule.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo, "number": number,
		}, "owner", "repo", "number")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubPRs", Description: "Liste les pull requests d'un dépôt GitHub. Lecture seule.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo, "state": state, "limit": limit,
		}, "owner", "repo")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubRepoCreate", Description: "Crée un dépôt GitHub sur ton compte. Approbation requise.", Parameters: str(map[string]any{
			"name":        map[string]any{"type": "string", "description": "Repository name"},
			"description": map[string]any{"type": "string"},
			"private":     map[string]any{"type": "boolean", "description": "Private repository (default true)"},
		}, "name")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubIssueCreate", Description: "Crée une issue GitHub. Approbation requise.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo,
			"title":  map[string]any{"type": "string", "description": "Issue title"},
			"body":   map[string]any{"type": "string", "description": "Issue body (markdown)"},
			"labels": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
		}, "owner", "repo", "title")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubIssueComment", Description: "Commente une issue (ou pull request) GitHub. Approbation requise.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo, "number": number,
			"body": map[string]any{"type": "string", "description": "Comment body (markdown)"},
		}, "owner", "repo", "number", "body")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubPRCreate", Description: "Crée une pull request GitHub. Approbation requise.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo,
			"title": map[string]any{"type": "string", "description": "Pull request title"},
			"head":  map[string]any{"type": "string", "description": "Branch containing the changes"},
			"base":  map[string]any{"type": "string", "description": "Branch to merge into (default main)"},
			"body":  map[string]any{"type": "string", "description": "Pull request body (markdown)"},
		}, "owner", "repo", "title", "head")}},
		{Type: "function", Function: provider.ToolFunction{Name: "GitHubPRMerge", Description: "Fusionne (merge) une pull request GitHub. Approbation requise.", Parameters: str(map[string]any{
			"owner": owner, "repo": repo, "number": number,
			"merge_method": map[string]any{"type": "string", "description": "merge, squash or rebase (default merge)", "enum": []string{"merge", "squash", "rebase"}},
		}, "owner", "repo", "number")}},
	}
}

// toolGitHub exécute un outil GitHub : vérifie le token puis dispatche.
func (s *Sandbox) toolGitHub(ctx context.Context, name string, args map[string]any) ToolResult {
	token := s.githubToken()
	if token == "" {
		return ToolResult{Text: "[erreur] GitHub non connecté : connectez votre compte dans Configuration → Connecteurs → GitHub, puis réessayez."}
	}
	switch name {
	case "GitHubRepos":
		return s.ghRepos(ctx, token, args)
	case "GitHubIssues":
		return s.ghIssues(ctx, token, args)
	case "GitHubIssueGet":
		return s.ghIssueGet(ctx, token, args)
	case "GitHubPRs":
		return s.ghPRs(ctx, token, args)
	case "GitHubRepoCreate":
		return s.ghRepoCreate(ctx, token, args)
	case "GitHubIssueCreate":
		return s.ghIssueCreate(ctx, token, args)
	case "GitHubIssueComment":
		return s.ghIssueComment(ctx, token, args)
	case "GitHubPRCreate":
		return s.ghPRCreate(ctx, token, args)
	case "GitHubPRMerge":
		return s.ghPRMerge(ctx, token, args)
	default:
		return ToolResult{Text: "[erreur] outil inconnu: " + name}
	}
}

// ghOwnerRepo valide et retourne owner/repo, ou un message d'erreur.
func ghOwnerRepo(args map[string]any) (string, string, string) {
	owner, repo := strArg(args, "owner"), strArg(args, "repo")
	if owner == "" || repo == "" {
		return "", "", "[erreur] GitHub : owner et repo sont requis."
	}
	if !githubNameRe.MatchString(owner) || !githubNameRe.MatchString(repo) {
		return "", "", "[erreur] GitHub : owner/repo invalide (caractères autorisés : lettres, chiffres, -, _, .)."
	}
	return owner, repo, ""
}

// ghLimit borne le paramètre limit (défaut 30, max 100).
func ghLimit(args map[string]any) int {
	n := intArg(args, "limit")
	if n <= 0 {
		return 30
	}
	if n > 100 {
		return 100
	}
	return n
}

// ghState valide le paramètre state (défaut open).
func ghState(args map[string]any) (string, string) {
	st := strings.ToLower(strings.TrimSpace(strArg(args, "state")))
	if st == "" {
		return "open", ""
	}
	if st != "open" && st != "closed" && st != "all" {
		return "", "[erreur] GitHub : state doit être open, closed ou all."
	}
	return st, ""
}

// ghNumber valide le paramètre number (> 0).
func ghNumber(args map[string]any) (int, string) {
	n := intArg(args, "number")
	if n <= 0 {
		return 0, "[erreur] GitHub : number doit être un entier positif."
	}
	return n, ""
}

// ghDo exécute un appel à l'API GitHub et retourne le statut et le corps
// (borné à githubMaxBody).
func ghDo(ctx context.Context, token, method, path string, query url.Values, payload any) (int, []byte, error) {
	var rdr io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return 0, nil, err
		}
		rdr = bytes.NewReader(b)
	}
	u := githubAPIBase + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u, rdr)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", githubAPIVersion)
	req.Header.Set("User-Agent", "cetas-lite")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{Timeout: githubDefaultTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(io.LimitReader(resp.Body, githubMaxBody))
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, b, nil
}

// ghMessage extrait le champ "message" d'une réponse d'erreur GitHub.
func ghMessage(body []byte) string {
	var m map[string]any
	if err := json.Unmarshal(body, &m); err != nil {
		return strings.TrimSpace(string(body))
	}
	if s, ok := m["message"].(string); ok && s != "" {
		if errs, ok := m["errors"].([]any); ok && len(errs) > 0 {
			if e0, ok := errs[0].(map[string]any); ok {
				if ms, ok := e0["message"].(string); ok && ms != "" {
					return s + " — " + ms
				}
			}
		}
		return s
	}
	return strings.TrimSpace(string(body))
}

// ghAPIError convertit un statut HTTP d'erreur en message actionnable.
func ghAPIError(status int, body []byte) string {
	msg := ghMessage(body)
	if msg != "" {
		msg = " : " + msg
	}
	switch status {
	case 401:
		return "[erreur] GitHub : token invalide ou révoqué (401). Reconnectez le compte dans Configuration → Connecteurs → GitHub" + msg + "."
	case 403:
		return "[erreur] GitHub : accès refusé (403)" + msg + ". Vérifiez les scopes du token et le quota d'API."
	case 404:
		return "[erreur] GitHub : introuvable (404)" + msg + "."
	case 422:
		return "[erreur] GitHub : requête invalide (422)" + msg + "."
	default:
		return fmt.Sprintf("[erreur] GitHub : HTTP %d%s.", status, msg)
	}
}

// ghJSON sérialise une projection en JSON compact, borné à toolMaxOutput.
func ghJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "[erreur] GitHub : sérialisation impossible."
	}
	return truncate(string(b), toolMaxOutput)
}

// ghDecode décode un corps JSON d'API ; en cas d'échec, erreur explicite.
func ghDecode(body []byte, v any) string {
	if err := json.Unmarshal(body, v); err != nil {
		return "[erreur] GitHub : réponse inattendue de l'API."
	}
	return ""
}

// --- projections compactes (économie de tokens) ---

func ghStr(v any) string {
	s, _ := v.(string)
	return s
}

func ghLogin(v any) string {
	if m, ok := v.(map[string]any); ok {
		return ghStr(m["login"])
	}
	return ""
}

func ghLabelNames(v any) []string {
	arr, _ := v.([]any)
	out := make([]string, 0, len(arr))
	for _, e := range arr {
		if m, ok := e.(map[string]any); ok {
			if n := ghStr(m["name"]); n != "" {
				out = append(out, n)
			}
		}
	}
	return out
}

func ghPreview(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > githubBodyPreview {
		return s[:githubBodyPreview] + "…"
	}
	return s
}

func ghProjectRepo(m map[string]any) map[string]any {
	return map[string]any{
		"name":           ghStr(m["name"]),
		"full_name":      ghStr(m["full_name"]),
		"private":        m["private"],
		"description":    ghStr(m["description"]),
		"html_url":       ghStr(m["html_url"]),
		"default_branch": ghStr(m["default_branch"]),
		"updated_at":     ghStr(m["updated_at"]),
	}
}

func ghProjectIssue(m map[string]any) map[string]any {
	return map[string]any{
		"number":     m["number"],
		"title":      ghStr(m["title"]),
		"state":      ghStr(m["state"]),
		"html_url":   ghStr(m["html_url"]),
		"user":       ghLogin(m["user"]),
		"labels":     ghLabelNames(m["labels"]),
		"comments":   m["comments"],
		"created_at": ghStr(m["created_at"]),
		"body":       ghPreview(ghStr(m["body"])),
	}
}

func ghProjectComment(m map[string]any) map[string]any {
	return map[string]any{
		"user":       ghLogin(m["user"]),
		"created_at": ghStr(m["created_at"]),
		"body":       ghPreview(ghStr(m["body"])),
	}
}

func ghProjectPR(m map[string]any) map[string]any {
	head, _ := m["head"].(map[string]any)
	base, _ := m["base"].(map[string]any)
	return map[string]any{
		"number":     m["number"],
		"title":      ghStr(m["title"]),
		"state":      ghStr(m["state"]),
		"draft":      m["draft"],
		"html_url":   ghStr(m["html_url"]),
		"user":       ghLogin(m["user"]),
		"head":       ghStr(head["ref"]),
		"base":       ghStr(base["ref"]),
		"created_at": ghStr(m["created_at"]),
		"body":       ghPreview(ghStr(m["body"])),
	}
}

// --- outils de lecture ---

func (s *Sandbox) ghRepos(ctx context.Context, token string, args map[string]any) ToolResult {
	q := url.Values{"per_page": {strconv.Itoa(ghLimit(args))}, "sort": {"updated"}}
	st, body, err := ghDo(ctx, token, "GET", "/user/repos", q, nil)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st < 200 || st >= 300 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var repos []map[string]any
	if msg := ghDecode(body, &repos); msg != "" {
		return ToolResult{Text: msg}
	}
	proj := make([]map[string]any, 0, len(repos))
	for _, r := range repos {
		proj = append(proj, ghProjectRepo(r))
	}
	return ToolResult{Text: ghJSON(proj)}
}

func (s *Sandbox) ghIssues(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	state, errMsg := ghState(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	q := url.Values{"per_page": {strconv.Itoa(ghLimit(args))}, "state": {state}, "sort": {"updated"}}
	st, body, err := ghDo(ctx, token, "GET", "/repos/"+owner+"/"+repo+"/issues", q, nil)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st < 200 || st >= 300 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var issues []map[string]any
	if msg := ghDecode(body, &issues); msg != "" {
		return ToolResult{Text: msg}
	}
	proj := make([]map[string]any, 0, len(issues))
	for _, is := range issues {
		// L'endpoint issues remonte aussi les PRs : on les signale.
		p := ghProjectIssue(is)
		if _, isPR := is["pull_request"]; isPR {
			p["is_pull_request"] = true
		}
		proj = append(proj, p)
	}
	return ToolResult{Text: ghJSON(proj)}
}

func (s *Sandbox) ghIssueGet(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	n, errMsg := ghNumber(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	base := "/repos/" + owner + "/" + repo + "/issues/" + strconv.Itoa(n)
	st, body, err := ghDo(ctx, token, "GET", base, nil, nil)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st < 200 || st >= 300 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var issue map[string]any
	if msg := ghDecode(body, &issue); msg != "" {
		return ToolResult{Text: msg}
	}
	proj := ghProjectIssue(issue)
	st, body, err = ghDo(ctx, token, "GET", base+"/comments", url.Values{"per_page": {"30"}}, nil)
	if err == nil && st >= 200 && st < 300 {
		var comments []map[string]any
		if ghDecode(body, &comments) == "" {
			pc := make([]map[string]any, 0, len(comments))
			for _, c := range comments {
				pc = append(pc, ghProjectComment(c))
			}
			proj["fetched_comments"] = pc
		}
	}
	return ToolResult{Text: ghJSON(proj)}
}

func (s *Sandbox) ghPRs(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	state, errMsg := ghState(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	q := url.Values{"per_page": {strconv.Itoa(ghLimit(args))}, "state": {state}, "sort": {"updated"}}
	st, body, err := ghDo(ctx, token, "GET", "/repos/"+owner+"/"+repo+"/pulls", q, nil)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st < 200 || st >= 300 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var prs []map[string]any
	if msg := ghDecode(body, &prs); msg != "" {
		return ToolResult{Text: msg}
	}
	proj := make([]map[string]any, 0, len(prs))
	for _, p := range prs {
		proj = append(proj, ghProjectPR(p))
	}
	return ToolResult{Text: ghJSON(proj)}
}

// --- outils d'écriture (approbation requise) ---

func (s *Sandbox) ghRepoCreate(ctx context.Context, token string, args map[string]any) ToolResult {
	name := strings.TrimSpace(strArg(args, "name"))
	if name == "" {
		return ToolResult{Text: "[erreur] GitHub : name est requis."}
	}
	if !githubNameRe.MatchString(name) {
		return ToolResult{Text: "[erreur] GitHub : nom de dépôt invalide (caractères autorisés : lettres, chiffres, -, _, .)."}
	}
	private := true
	if v, ok := args["private"].(bool); ok {
		private = v
	}
	payload := map[string]any{
		"name":        name,
		"description": strArg(args, "description"),
		"private":     private,
	}
	st, body, err := ghDo(ctx, token, "POST", "/user/repos", nil, payload)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st != 201 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var r map[string]any
	if msg := ghDecode(body, &r); msg != "" {
		return ToolResult{Text: msg}
	}
	return ToolResult{Text: "[ok] dépôt créé : " + ghStr(r["full_name"]) + " — " + ghStr(r["html_url"])}
}

func (s *Sandbox) ghIssueCreate(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	title := strings.TrimSpace(strArg(args, "title"))
	if title == "" {
		return ToolResult{Text: "[erreur] GitHub : title est requis."}
	}
	payload := map[string]any{"title": title, "body": strArg(args, "body")}
	if l, ok := args["labels"].([]any); ok && len(l) > 0 {
		labels := make([]string, 0, len(l))
		for _, e := range l {
			if str, ok := e.(string); ok && str != "" {
				labels = append(labels, str)
			}
		}
		payload["labels"] = labels
	}
	st, body, err := ghDo(ctx, token, "POST", "/repos/"+owner+"/"+repo+"/issues", nil, payload)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st != 201 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var is map[string]any
	if msg := ghDecode(body, &is); msg != "" {
		return ToolResult{Text: msg}
	}
	return ToolResult{Text: fmt.Sprintf("[ok] issue #%v créée : %s", is["number"], ghStr(is["html_url"]))}
}

func (s *Sandbox) ghIssueComment(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	n, errMsg := ghNumber(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	cbody := strings.TrimSpace(strArg(args, "body"))
	if cbody == "" {
		return ToolResult{Text: "[erreur] GitHub : body est requis."}
	}
	st, body, err := ghDo(ctx, token, "POST",
		"/repos/"+owner+"/"+repo+"/issues/"+strconv.Itoa(n)+"/comments",
		nil, map[string]any{"body": cbody})
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st != 201 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var c map[string]any
	if msg := ghDecode(body, &c); msg != "" {
		return ToolResult{Text: msg}
	}
	return ToolResult{Text: "[ok] commentaire publié : " + ghStr(c["html_url"])}
}

func (s *Sandbox) ghPRCreate(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	title := strings.TrimSpace(strArg(args, "title"))
	head := strings.TrimSpace(strArg(args, "head"))
	if title == "" || head == "" {
		return ToolResult{Text: "[erreur] GitHub : title et head sont requis."}
	}
	base := strings.TrimSpace(strArg(args, "base"))
	if base == "" {
		base = "main"
	}
	if !githubNameRe.MatchString(head) || !githubNameRe.MatchString(base) {
		return ToolResult{Text: "[erreur] GitHub : nom de branche invalide (caractères autorisés : lettres, chiffres, -, _, .)."}
	}
	payload := map[string]any{"title": title, "head": head, "base": base, "body": strArg(args, "body")}
	st, body, err := ghDo(ctx, token, "POST", "/repos/"+owner+"/"+repo+"/pulls", nil, payload)
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st != 201 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var p map[string]any
	if msg := ghDecode(body, &p); msg != "" {
		return ToolResult{Text: msg}
	}
	return ToolResult{Text: fmt.Sprintf("[ok] pull request #%v créée : %s", p["number"], ghStr(p["html_url"]))}
}

func (s *Sandbox) ghPRMerge(ctx context.Context, token string, args map[string]any) ToolResult {
	owner, repo, errMsg := ghOwnerRepo(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	n, errMsg := ghNumber(args)
	if errMsg != "" {
		return ToolResult{Text: errMsg}
	}
	method := strings.ToLower(strings.TrimSpace(strArg(args, "merge_method")))
	if method == "" {
		method = "merge"
	}
	if method != "merge" && method != "squash" && method != "rebase" {
		return ToolResult{Text: "[erreur] GitHub : merge_method doit être merge, squash ou rebase."}
	}
	st, body, err := ghDo(ctx, token, "PUT",
		"/repos/"+owner+"/"+repo+"/pulls/"+strconv.Itoa(n)+"/merge",
		nil, map[string]any{"merge_method": method})
	if err != nil {
		return ToolResult{Text: "[erreur] GitHub : " + err.Error()}
	}
	if st < 200 || st >= 300 {
		return ToolResult{Text: ghAPIError(st, body)}
	}
	var r map[string]any
	if msg := ghDecode(body, &r); msg != "" {
		return ToolResult{Text: msg}
	}
	merged, _ := r["merged"].(bool)
	if !merged {
		return ToolResult{Text: "[erreur] GitHub : PR non fusionnée — " + ghMessage(body) + "."}
	}
	return ToolResult{Text: fmt.Sprintf("[ok] pull request #%d fusionnée (%s).", n, method)}
}
