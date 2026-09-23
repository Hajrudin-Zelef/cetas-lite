package docs

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/dslipak/pdf"
)

const maxText = 512 << 10

type Kind string

const (
	KindText Kind = "text"
	KindPDF  Kind = "pdf"
	KindHTML Kind = "html"
)

var structured = map[string]Kind{
	".pdf":   KindPDF,
	".html":  KindHTML,
	".htm":   KindHTML,
	".xhtml": KindHTML,
}

var textExts = map[string]bool{
	".md": true, ".markdown": true, ".txt": true, ".text": true, ".log": true,
	".csv": true, ".tsv": true, ".json": true, ".yaml": true, ".yml": true,
	".toml": true, ".ini": true, ".cfg": true, ".conf": true, ".env": true,
	".py": true, ".js": true, ".mjs": true, ".cjs": true, ".ts": true, ".tsx": true, ".jsx": true,
	".go": true, ".rs": true, ".java": true, ".kt": true, ".c": true, ".h": true,
	".cpp": true, ".hpp": true, ".cs": true, ".rb": true, ".php": true,
	".sh": true, ".bash": true, ".zsh": true, ".sql": true, ".xml": true,
	".css": true, ".scss": true, ".vue": true, ".svelte": true, ".swift": true,
	".m": true, ".mm": true, ".pl": true, ".lua": true, ".r": true, ".dart": true,
	".scala": true, ".clj": true, ".ex": true, ".exs": true, ".erl": true, ".hs": true,
	".elm": true, ".tf": true, ".gradle": true, ".properties": true,
}

func IsSupported(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	if _, ok := structured[ext]; ok {
		return true
	}
	return textExts[ext]
}

func KindOf(name string) Kind {
	ext := strings.ToLower(filepath.Ext(name))
	if k, ok := structured[ext]; ok {
		return k
	}
	return KindText
}

func Extract(name string, data []byte) (string, Kind, error) {
	kind := KindOf(name)
	switch kind {
	case KindPDF:
		text, err := extractPDF(data)
		if err != nil {
			return "", kind, err
		}
		return capText(text), kind, nil
	case KindHTML:
		md, err := htmltomarkdown.ConvertString(string(data))
		if err != nil {
			return capText(string(data)), kind, nil
		}
		return capText(md), kind, nil
	default:
		return capText(string(data)), kind, nil
	}
}

func extractPDF(data []byte) (string, error) {
	if r, err := pdf.NewReader(bytes.NewReader(data), int64(len(data))); err == nil {
		if rc, err := r.GetPlainText(); err == nil {
			b, _ := io.ReadAll(io.LimitReader(rc, maxText))
			if s := strings.TrimSpace(string(b)); s != "" {
				return s, nil
			}
		}
	}
	if s, err := pdftotext(data); err == nil && strings.TrimSpace(s) != "" {
		return strings.TrimSpace(s), nil
	}
	return "", errors.New("PDF extraction unavailable (install poppler-utils/pdftotext for better support)")
}

func pdftotext(data []byte) (string, error) {
	bin, err := exec.LookPath("pdftotext")
	if err != nil {
		return "", err
	}
	tmp, err := os.CreateTemp("", "cetas-*.pdf")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return "", err
	}
	tmp.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, bin, "-q", "-enc", "UTF-8", tmp.Name(), "-").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func capText(s string) string {
	if len(s) > maxText {
		return s[:maxText] + "\n… [document tronque]"
	}
	return s
}
