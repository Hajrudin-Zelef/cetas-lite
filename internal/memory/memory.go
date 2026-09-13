package memory

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	indexFile     = "MEMORY.md"
	indexHeader   = "# Index memoire\n"
	maxReadLines  = 500
	defaultSearch = 8
	maxSearch     = 30
)

var (
	errOutside        = errors.New("chemin hors memoire")
	ErrAlreadyApplied = errors.New("deja a jour")
	nameRe            = regexp.MustCompile(`^[A-Za-z0-9._-]+\.md$`)
)

type Page struct {
	Name  string
	Title string
}

type Hit struct {
	File    string
	Title   string
	Snippet string
}

type Store struct {
	root string
}

func New(root string) *Store {
	root = strings.TrimSpace(root)
	if root == "" {
		return &Store{}
	}
	if abs, err := filepath.Abs(root); err == nil {
		root = abs
	}
	_ = os.MkdirAll(root, 0o700)
	return &Store{root: root}
}

func (s *Store) enabled() bool { return s != nil && s.root != "" }

func sanitizeUser(user string) (string, error) {
	safe := strings.ReplaceAll(user, "/", "_")
	safe = strings.ReplaceAll(safe, "\\", "_")
	safe = strings.TrimSpace(safe)
	if safe == "" || safe == "." || safe == ".." {
		return "", errors.New("utilisateur invalide")
	}
	return safe, nil
}

func (s *Store) userDir(user string) (string, error) {
	if !s.enabled() {
		return "", errors.New("memoire non configuree")
	}
	safe, err := sanitizeUser(user)
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(filepath.Join(s.root, safe))
	if err != nil {
		return "", err
	}
	rootAbs, err := filepath.Abs(s.root)
	if err != nil {
		return "", err
	}
	if abs != rootAbs && !strings.HasPrefix(abs, rootAbs+string(os.PathSeparator)) {
		return "", errOutside
	}
	return abs, nil
}

func safeName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("nom vide")
	}
	if !strings.HasSuffix(strings.ToLower(name), ".md") {
		name += ".md"
	}
	if !nameRe.MatchString(name) {
		return "", errors.New("nom invalide (alphanum, ._-)")
	}
	return name, nil
}

func (s *Store) pagePath(user, name string) (string, error) {
	dir, err := s.userDir(user)
	if err != nil {
		return "", err
	}
	fn, err := safeName(name)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fn), nil
}

func writeRaw(p string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o600)
}

func titleOf(content string) string {
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(line), "#"))
		if t != "" {
			return t
		}
	}
	return ""
}

func (s *Store) List(user string) []Page {
	dir, err := s.userDir(user)
	if err != nil {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	out := []Page{}
	for _, e := range entries {
		if e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(e.Name()), ".md") {
			continue
		}
		if strings.EqualFold(e.Name(), indexFile) {
			continue
		}
		title := ""
		if b, err := os.ReadFile(filepath.Join(dir, e.Name())); err == nil {
			title = titleOf(string(b))
		}
		out = append(out, Page{Name: e.Name(), Title: title})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func (s *Store) Read(user, name string, offset, limit int) (string, error) {
	p, err := s.pagePath(user, name)
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("page '%s' introuvable", name)
		}
		return "", err
	}
	lines := strings.Split(string(b), "\n")
	if offset <= 0 {
		offset = 1
	}
	if limit <= 0 || limit > maxReadLines {
		limit = maxReadLines
	}
	var out strings.Builder
	for i := offset - 1; i < len(lines) && i < offset-1+limit; i++ {
		fmt.Fprintf(&out, "%d\t%s\n", i+1, lines[i])
	}
	return strings.TrimRight(out.String(), "\n"), nil
}

func (s *Store) Content(user, name string) string {
	p, err := s.pagePath(user, name)
	if err != nil {
		return ""
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (s *Store) Add(user, name, content string) error {
	fn, err := safeName(name)
	if err != nil {
		return err
	}
	if strings.EqualFold(fn, indexFile) {
		return errors.New("nom reserve a l'index")
	}
	p, err := s.pagePath(user, fn)
	if err != nil {
		return err
	}
	if _, err := os.Stat(p); err == nil {
		return errors.New("la page existe deja — utilise mem_edit pour la modifier")
	}
	body := strings.TrimRight(content, "\n") + "\n"
	if err := writeRaw(p, []byte(body)); err != nil {
		return err
	}
	s.indexUpsert(user, fn)
	return nil
}

func (s *Store) Edit(user, name, oldText, newText string) error {
	if oldText == "" {
		return errors.New("old vide")
	}
	fn, err := safeName(name)
	if err != nil {
		return err
	}
	p, err := s.pagePath(user, fn)
	if err != nil {
		return err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("page '%s' introuvable", name)
		}
		return err
	}
	content := string(b)
	n := strings.Count(content, oldText)
	if n == 0 {
		if newText != "" && strings.Contains(content, newText) {
			return ErrAlreadyApplied
		}
		return errors.New("old introuvable dans la page")
	}
	if n > 1 {
		return fmt.Errorf("old apparait %d fois — ajoute du contexte pour le rendre unique", n)
	}
	updated := strings.Replace(content, oldText, newText, 1)
	if err := writeRaw(p, []byte(updated)); err != nil {
		return err
	}
	s.indexUpsert(user, fn)
	return nil
}

func (s *Store) Delete(user, name string) error {
	fn, err := safeName(name)
	if err != nil {
		return err
	}
	if strings.EqualFold(fn, indexFile) {
		return errors.New("nom reserve a l'index")
	}
	p, err := s.pagePath(user, fn)
	if err != nil {
		return err
	}
	if _, err := os.Stat(p); err != nil {
		return errors.New("page introuvable: " + fn)
	}
	if err := os.Remove(p); err != nil {
		return err
	}
	s.indexRemove(user, fn)
	return nil
}

func (s *Store) Search(user, query string, limit int) []Hit {
	if limit <= 0 || limit > maxSearch {
		limit = defaultSearch
	}
	terms := uniqueTerms(strings.ToLower(query))
	if len(terms) == 0 {
		return nil
	}
	type doc struct {
		p       Page
		name    string
		title   string
		content string
		hay     string
	}
	docs := []doc{}
	df := map[string]int{}
	for _, p := range s.List(user) {
		content := s.Content(user, p.Name)
		d := doc{
			p:       p,
			name:    strings.ToLower(p.Name),
			title:   strings.ToLower(p.Title),
			content: content,
			hay:     strings.ToLower(p.Name + "\n" + p.Title + "\n" + content),
		}
		docs = append(docs, d)
		for _, t := range terms {
			if strings.Contains(d.hay, t) {
				df[t]++
			}
		}
	}
	n := len(docs)
	type scored struct {
		hit     Hit
		matched int
		score   float64
	}
	var ranked []scored
	for _, d := range docs {
		matched := 0
		score := 0.0
		for _, t := range terms {
			cnt := strings.Count(d.hay, t)
			if cnt == 0 {
				continue
			}
			matched++
			idf := math.Log(float64(n+1)/float64(df[t]+1)) + 1
			tf := 1.0 + math.Log(float64(cnt))
			field := 1.0
			if strings.Contains(d.name, t) {
				field += 4
			}
			if strings.Contains(d.title, t) {
				field += 2
			}
			score += idf * tf * field
		}
		if matched == 0 {
			continue
		}
		ranked = append(ranked, scored{
			hit:     Hit{File: d.p.Name, Title: d.p.Title, Snippet: snippetAround(d.content, terms)},
			matched: matched,
			score:   score,
		})
	}
	sort.SliceStable(ranked, func(i, j int) bool {
		if ranked[i].matched != ranked[j].matched {
			return ranked[i].matched > ranked[j].matched
		}
		return ranked[i].score > ranked[j].score
	})
	out := []Hit{}
	for i, r := range ranked {
		if i >= limit {
			break
		}
		out = append(out, r.hit)
	}
	return out
}

func (s *Store) Index(user string) string {
	return s.Content(user, indexFile)
}

func (s *Store) indexUpsert(user, name string) {
	fn, err := safeName(name)
	if err != nil || strings.EqualFold(fn, indexFile) {
		return
	}
	dir, err := s.userDir(user)
	if err != nil {
		return
	}
	idxPath := filepath.Join(dir, indexFile)
	content := indexHeader
	if b, err := os.ReadFile(idxPath); err == nil {
		content = string(b)
	}
	title := fn
	if p, err := s.pagePath(user, fn); err == nil {
		if b, err := os.ReadFile(p); err == nil {
			if t := titleOf(string(b)); t != "" {
				title = t
			}
		}
	}
	line := "- [" + title + "](" + fn + ")"
	ref := "](" + fn + ")"
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	found := false
	for i, l := range lines {
		if strings.Contains(l, ref) {
			lines[i] = line
			found = true
			break
		}
	}
	if !found {
		lines = append(lines, line)
	}
	body := strings.TrimRight(strings.Join(lines, "\n"), "\n") + "\n"
	_ = writeRaw(idxPath, []byte(body))
}

func (s *Store) indexRemove(user, name string) {
	fn, err := safeName(name)
	if err != nil || strings.EqualFold(fn, indexFile) {
		return
	}
	dir, err := s.userDir(user)
	if err != nil {
		return
	}
	idxPath := filepath.Join(dir, indexFile)
	b, err := os.ReadFile(idxPath)
	if err != nil {
		return
	}
	ref := "](" + fn + ")"
	lines := strings.Split(string(b), "\n")
	out := make([]string, 0, len(lines))
	changed := false
	for _, l := range lines {
		if strings.Contains(l, ref) {
			changed = true
			continue
		}
		out = append(out, l)
	}
	if !changed {
		return
	}
	body := strings.TrimRight(strings.Join(out, "\n"), "\n") + "\n"
	_ = writeRaw(idxPath, []byte(body))
}

func uniqueTerms(q string) []string {
	seen := map[string]bool{}
	var out []string
	for _, t := range strings.Fields(q) {
		if !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	return out
}

func snippetAround(content string, terms []string) string {
	flat := strings.Join(strings.Fields(content), " ")
	low := strings.ToLower(flat)
	idx := -1
	for _, t := range terms {
		if i := strings.Index(low, t); i >= 0 && (idx < 0 || i < idx) {
			idx = i
		}
	}
	if idx < 0 {
		if len(flat) > 160 {
			return flat[:160] + "…"
		}
		return flat
	}
	start := idx - 60
	if start < 0 {
		start = 0
	}
	end := idx + 100
	if end > len(flat) {
		end = len(flat)
	}
	s := flat[start:end]
	if start > 0 {
		s = "…" + s
	}
	if end < len(flat) {
		s += "…"
	}
	return s
}
