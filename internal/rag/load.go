package rag

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const manifestFile = "manifest.json"

var (
	docExts   = map[string]bool{".md": true, ".markdown": true, ".txt": true, ".text": true}
	skipNames = map[string]bool{"readme.md": true, "index.md": true}
)

type manifestDoc struct {
	Corpus string          `json:"corpus"`
	Title  string          `json:"title"`
	Chunks []manifestChunk `json:"chunks"`
}

type manifestChunk struct {
	Path     string   `json:"path"`
	Title    string   `json:"title"`
	Domain   string   `json:"domain"`
	Task     string   `json:"task"`
	Actors   []string `json:"actors"`
	Dates    []string `json:"dates"`
	Keywords []string `json:"keywords"`
	Section  string   `json:"section"`
	DeltaOf  string   `json:"delta_of"`
}

// Load construit l'index depuis un dossier racine.
//
// Chaque sous-dossier contenant un manifest.json est lu en mode riche
// (facettes + chemins de chunks) ; les autres sont lus en mode brut (un
// fichier .md/.txt = un chunk). Un dossier absent ou vide donne un index
// vide (Ready=false) sans erreur : le RAG est alors inactif et ne coute rien.
func Load(root string) (*Index, error) {
	start := time.Now()
	root = strings.TrimSpace(root)
	if root == "" {
		return emptyIndex(), nil
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return emptyIndex(), nil
	}
	entries, err := os.ReadDir(abs)
	if err != nil {
		return nil, err
	}

	var chunks []chunk
	var errs []string
	corpora := 0
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
			continue
		}
		if !e.IsDir() {
			if isDoc(name) && !skipNames[strings.ToLower(name)] {
				if c, ok := rawChunk(abs, abs, name); ok {
					chunks = append(chunks, c)
					corpora++
				}
			}
			continue
		}
		dir := filepath.Join(abs, name)
		if _, statErr := os.Stat(filepath.Join(dir, manifestFile)); statErr == nil {
			cs, err := loadManifest(abs, dir, name)
			if err != nil {
				errs = append(errs, name+": "+err.Error())
				continue
			}
			chunks = append(chunks, cs...)
			corpora++
			continue
		}
		if cs := loadRawDir(abs, dir); len(cs) > 0 {
			chunks = append(chunks, cs...)
			corpora++
		}
	}
	return buildIndex(chunks, corpora, errs, start), nil
}

func loadManifest(root, dir, corpusName string) ([]chunk, error) {
	raw, err := os.ReadFile(filepath.Join(dir, manifestFile))
	if err != nil {
		return nil, err
	}
	var doc manifestDoc
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}
	corpus := strings.TrimSpace(doc.Corpus)
	if corpus == "" {
		corpus = corpusName
	}
	out := make([]chunk, 0, len(doc.Chunks))
	for _, mc := range doc.Chunks {
		rel := filepath.ToSlash(strings.TrimSpace(mc.Path))
		if rel == "" {
			continue
		}
		body, ok := readChunkFile(root, dir, rel)
		if !ok {
			continue
		}
		body = stripFrontMatter(body)
		if len(strings.Fields(body)) == 0 {
			continue
		}
		domain := strings.TrimSpace(mc.Domain)
		if domain == "" {
			domain = domainOf(rel)
		}
		out = append(out, chunk{
			path: rel, corpus: corpus, title: mc.Title, domain: domain,
			task: mc.Task, actors: mc.Actors, dates: mc.Dates, keywords: mc.Keywords,
			section: mc.Section, deltaOf: mc.DeltaOf, body: body,
		})
	}
	return out, nil
}

func loadRawDir(root, dir string) []chunk {
	var out []chunk
	_ = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		name := d.Name()
		if d.IsDir() {
			if p != dir && (strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")) {
				return fs.SkipDir
			}
			return nil
		}
		if !isDoc(name) || skipNames[strings.ToLower(name)] {
			return nil
		}
		if c, ok := rawChunk(root, filepath.Dir(p), name); ok {
			out = append(out, c)
		}
		return nil
	})
	sort.Slice(out, func(i, j int) bool { return out[i].path < out[j].path })
	return out
}

func rawChunk(root, dir, name string) (chunk, bool) {
	p := filepath.Join(dir, name)
	if !within(root, p) {
		return chunk{}, false
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return chunk{}, false
	}
	body := string(b)
	body = stripFrontMatter(body)
	if len(strings.Fields(body)) == 0 {
		return chunk{}, false
	}
	rel, err := filepath.Rel(root, p)
	if err != nil {
		return chunk{}, false
	}
	rel = filepath.ToSlash(rel)
	domain := ""
	if d, err := filepath.Rel(root, dir); err == nil && d != "." {
		domain = filepath.ToSlash(d)
	}
	return chunk{
		path: rel, corpus: topSegment(rel), title: titleOf(body, name),
		domain: domain, body: body,
	}, true
}

func readChunkFile(root, dir, rel string) (string, bool) {
	cands := []string{filepath.Join(root, filepath.FromSlash(rel))}
	if i := strings.IndexByte(rel, '/'); i >= 0 {
		cands = append(cands, filepath.Join(dir, filepath.FromSlash(rel[i+1:])))
	}
	cands = append(cands, filepath.Join(dir, filepath.Base(rel)))
	for _, p := range cands {
		if !within(root, p) {
			continue
		}
		if b, err := os.ReadFile(p); err == nil {
			return string(b), true
		}
	}
	return "", false
}

func within(root, p string) bool {
	abs, err := filepath.Abs(p)
	if err != nil {
		return false
	}
	r, err := filepath.Abs(root)
	if err != nil {
		return false
	}
	return abs == r || strings.HasPrefix(abs, r+string(os.PathSeparator))
}

func isDoc(name string) bool { return docExts[strings.ToLower(filepath.Ext(name))] }

func topSegment(rel string) string {
	if i := strings.IndexByte(rel, '/'); i >= 0 {
		return rel[:i]
	}
	return ""
}

// stripFrontMatter retire un en-tete YAML (--- ... ---) en tete de document :
// il duplique des metadonnees deja portees par le manifeste et pollue les
// snippets. Sans bloc fermant, le corps est rendu tel quel.
func stripFrontMatter(body string) string {
	if !strings.HasPrefix(body, "---\n") && !strings.HasPrefix(body, "---\r\n") {
		return body
	}
	lines := strings.Split(body, "\n")
	limit := len(lines)
	if limit > 60 {
		limit = 60
	}
	for i := 1; i < limit; i++ {
		if strings.TrimSpace(strings.TrimRight(lines[i], "\r")) == "---" {
			return strings.TrimLeft(strings.Join(lines[i+1:], "\n"), "\n")
		}
	}
	return body
}

func domainOf(rel string) string {
	parts := strings.Split(rel, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}

func titleOf(body, name string) string {
	for _, line := range strings.Split(body, "\n") {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "#") {
			if h := strings.TrimSpace(strings.TrimLeft(t, "#")); h != "" {
				return truncateRunes(h, 120)
			}
			continue
		}
		return truncateRunes(t, 120)
	}
	base := name
	if i := strings.LastIndexByte(base, '.'); i > 0 {
		base = base[:i]
	}
	return base
}

func truncateRunes(s string, n int) string {
	if n <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n])
}
