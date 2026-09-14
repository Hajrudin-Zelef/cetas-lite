//go:build windows

package plugins

import "testing"

// TestQuoteWindowsArg verifie la citation MSVCRT des arguments : chaque
// forme produite doit se re-decoder en l'argument d'origine.
func TestQuoteWindowsArg(t *testing.T) {
	cas := []struct{ in, want string }{
		{"simple", "simple"},
		{"", `""`},
		{"a b", `"a b"`},
		{"a\"b", `"a\"b"`},
		{`a\b`, `a\b`},
		{`a\ b`, `"a\ b"`},
		{`trail\ x`, `"trail\\ x"`},
		{`C:\Program Files\app.exe`, `"C:\Program Files\app.exe"`},
		{`q\"`, `"q\\\""`},
	}
	for _, c := range cas {
		if got := quoteWindowsArg(c.in); got != c.want {
			t.Errorf("quoteWindowsArg(%q) = %q, voulu %q", c.in, got, c.want)
		}
	}
}
