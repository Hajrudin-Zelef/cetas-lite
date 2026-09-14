package docs

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func makePDF(content string) []byte {
	var b bytes.Buffer
	var offs []int
	obj := func(s string) {
		offs = append(offs, b.Len())
		b.WriteString(s)
	}
	b.WriteString("%PDF-1.4\n")
	obj("1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n")
	obj("2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n")
	obj("3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Contents 4 0 R /Resources << /Font << /F1 5 0 R >> >> >>\nendobj\n")
	stream := "BT /F1 24 Tf 72 700 Td (" + content + ") Tj ET"
	obj(fmt.Sprintf("4 0 obj\n<< /Length %d >>\nstream\n%s\nendstream\nendobj\n", len(stream), stream))
	obj("5 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>\nendobj\n")
	xref := b.Len()
	fmt.Fprintf(&b, "xref\n0 %d\n", len(offs)+1)
	b.WriteString("0000000000 65535 f \n")
	for _, o := range offs {
		fmt.Fprintf(&b, "%010d 00000 n \n", o)
	}
	fmt.Fprintf(&b, "trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", len(offs)+1, xref)
	return b.Bytes()
}

func TestExtractText(t *testing.T) {
	out, kind, err := Extract("notes.md", []byte("# Titre\n\ncorps"))
	if err != nil || kind != KindText {
		t.Fatalf("kind=%v err=%v", kind, err)
	}
	if !strings.Contains(out, "corps") {
		t.Fatalf("out=%q", out)
	}
}

func TestExtractHTML(t *testing.T) {
	out, kind, err := Extract("page.html", []byte("<h1>Titre</h1><p>Texte</p>"))
	if err != nil || kind != KindHTML {
		t.Fatalf("kind=%v err=%v", kind, err)
	}
	if !strings.Contains(out, "Titre") || !strings.Contains(out, "Texte") {
		t.Fatalf("markdown inattendu: %q", out)
	}
}

func TestExtractPDF(t *testing.T) {
	out, kind, err := Extract("doc.pdf", makePDF("Hello PDF"))
	if err != nil {
		t.Fatalf("extraction PDF: %v", err)
	}
	if kind != KindPDF || !strings.Contains(out, "Hello PDF") {
		t.Fatalf("kind=%v out=%q", kind, out)
	}
}

func TestExtractPDFInvalid(t *testing.T) {
	if _, _, err := Extract("bad.pdf", []byte("pas un pdf")); err == nil {
		t.Fatal("PDF invalide doit echouer")
	}
}

func TestIsSupported(t *testing.T) {
	for _, ok := range []string{"a.md", "a.py", "a.pdf", "a.html", "a.json"} {
		if !IsSupported(ok) {
			t.Errorf("%s devrait etre supporte", ok)
		}
	}
	for _, ko := range []string{"a.docx", "a.xlsx", "a.bin"} {
		if IsSupported(ko) {
			t.Errorf("%s ne devrait pas etre supporte", ko)
		}
	}
}
