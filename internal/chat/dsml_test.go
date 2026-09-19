package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/provider"
)

func TestParseDSMLCalls_Basic(t *testing.T) {
	text := `Voici le fichier :
<||DSML||calls>
<||DSML||invoke name="Read">
<||DSML||parameter name="file_path">docs/a.md</||DSML||parameter>
</||DSML||invoke>
</||DSML||calls>
Fin.`
	calls := parseDSMLCalls(text)
	if len(calls) != 1 {
		t.Fatalf("attendu 1 appel, obtenu %d", len(calls))
	}
	if calls[0].Name != "Read" {
		t.Errorf("nom = %q, attendu Read", calls[0].Name)
	}
	if calls[0].Args["file_path"] != "docs/a.md" {
		t.Errorf("file_path = %q", calls[0].Args["file_path"])
	}
}

func TestParseDSMLCalls_Multiple(t *testing.T) {
	text := `<||DSML||calls>
<||DSML||invoke name="Read"><||DSML||parameter name="file_path">a.md</||DSML||parameter></||DSML||invoke>
<||DSML||invoke name="Bash"><||DSML||parameter name="command">go test ./...</||DSML||parameter></||DSML||invoke>
</||DSML||calls>`
	calls := parseDSMLCalls(text)
	if len(calls) != 2 {
		t.Fatalf("attendu 2 appels, obtenu %d", len(calls))
	}
	if calls[1].Name != "Bash" || calls[1].Args["command"] != "go test ./..." {
		t.Errorf("2e appel inattendu : %+v", calls[1])
	}
}

func TestParseDSMLCalls_Unicode(t *testing.T) {
	text := "＜｜｜DSML｜｜calls＞＜｜｜DSML｜｜invoke name=\"Ls\"＞＜｜｜DSML｜｜parameter name=\"path\"＞.＜／｜｜DSML｜｜parameter＞＜／｜｜DSML｜｜invoke＞＜／｜｜DSML｜｜calls＞"
	calls := parseDSMLCalls(text)
	if len(calls) != 1 || calls[0].Name != "Ls" || calls[0].Args["path"] != "." {
		t.Fatalf("variantes pleine-chasse non reconnues : %+v", calls)
	}
}

func TestParseDSMLCalls_Unclosed(t *testing.T) {
	// Bloc traînant en fin de flux : l'appel doit quand même être extrait.
	text := `Texte avant.
<||DSML||calls>
<||DSML||invoke name="Grep">
<||DSML||parameter name="pattern">foo</||DSML||parameter>`
	calls := parseDSMLCalls(text)
	if len(calls) != 1 || calls[0].Name != "Grep" || calls[0].Args["pattern"] != "foo" {
		t.Fatalf("bloc non fermé non extrait : %+v", calls)
	}
}

func TestParseDSMLCalls_NoDSML(t *testing.T) {
	if calls := parseDSMLCalls("Juste du texte normal."); calls != nil {
		t.Fatalf("attendu nil, obtenu %+v", calls)
	}
}

func TestParseDSMLCalls_InvokeSansNom(t *testing.T) {
	text := `<||DSML||calls><||DSML||invoke><||DSML||parameter name="x">1</||DSML||parameter></||DSML||invoke></||DSML||calls>`
	if calls := parseDSMLCalls(text); len(calls) != 0 {
		t.Fatalf("invoke sans nom aurait dû être ignoré : %+v", calls)
	}
}

func TestParseDSMLCalls_GuillemetsSimples(t *testing.T) {
	text := `<||DSML||invoke name='Read'><||DSML||parameter name='file_path'>b.md</||DSML||parameter></||DSML||invoke>`
	calls := parseDSMLCalls(text)
	if len(calls) != 1 || calls[0].Args["file_path"] != "b.md" {
		t.Fatalf("guillemets simples non gérés : %+v", calls)
	}
}

func TestParseDSMLCalls_ParametreMultiligne(t *testing.T) {
	text := "<||DSML||invoke name=\"Write\">\n<||DSML||parameter name=\"content\">ligne1\nligne2\nligne3</||DSML||parameter>\n</||DSML||invoke>"
	calls := parseDSMLCalls(text)
	if len(calls) != 1 || calls[0].Args["content"] != "ligne1\nligne2\nligne3" {
		t.Fatalf("paramètre multiligne non géré : %+v", calls)
	}
}

func TestStripDSMLStreaming(t *testing.T) {
	text := "Bonjour <||DSML||calls><||DSML||invoke name=\"Read\"></||DSML||invoke></||DSML||calls> au revoir"
	got := stripDSMLStreaming(text)
	if strings.Contains(got, "DSML") {
		t.Errorf("DSML restant dans le flux : %q", got)
	}
	if !strings.Contains(got, "Bonjour") || !strings.Contains(got, "au revoir") {
		t.Errorf("texte normal perdu : %q", got)
	}
	// Bloc non fermé : tout ce qui suit est masqué en streaming.
	tail := "Début <||DSML||calls><||DSML||invoke name=\"Read\">"
	if got := stripDSMLStreaming(tail); got != "Début " {
		t.Errorf("masquage streaming = %q, attendu %q", got, "Début ")
	}
}

func TestStripDSMLFinal(t *testing.T) {
	text := "Avant <||DSML||calls><||DSML||invoke name=\"Read\"><||DSML||parameter name=\"file_path\">x</||DSML||parameter></||DSML||invoke></||DSML||calls> après <||DSML||calls><||DSML||invoke name=\"Bash\">"
	got := stripDSMLFinal(text)
	if strings.Contains(got, "DSML") {
		t.Errorf("DSML restant en final : %q", got)
	}
	if !strings.Contains(got, "Avant") || !strings.Contains(got, "après") {
		t.Errorf("texte normal perdu : %q", got)
	}
}

func TestDsmlToToolCalls(t *testing.T) {
	calls := []dsmlCall{{Name: "Read", Args: map[string]string{"file_path": "a.md"}}}
	tcs := dsmlToToolCalls(calls)
	if len(tcs) != 1 {
		t.Fatalf("attendu 1 toolcall")
	}
	tc := tcs[0]
	if tc.ID != "dsml-1" || tc.Function.Name != "Read" {
		t.Errorf("toolcall inattendu : %+v", tc)
	}
	if !strings.Contains(tc.Function.Arguments, `"file_path":"a.md"`) {
		t.Errorf("arguments JSON inattendus : %s", tc.Function.Arguments)
	}
	if tc.InvalidCall() {
		t.Error("le toolcall ne devrait pas être marqué invalide")
	}
}

func TestCanonicalDSMLName(t *testing.T) {
	tools := []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "Read"}},
		{Type: "function", Function: provider.ToolFunction{Name: "Bash"}},
	}
	if got := canonicalDSMLName("read", tools); got != "Read" {
		t.Errorf("canonical(read) = %q", got)
	}
	if got := canonicalDSMLName("BASH", tools); got != "Bash" {
		t.Errorf("canonical(BASH) = %q", got)
	}
	if got := canonicalDSMLName("Inconnu", tools); got != "Inconnu" {
		t.Errorf("nom inconnu modifié : %q", got)
	}
}

func TestDsmlTagPrefixish(t *testing.T) {
	oui := []string{"<", "<||", "<||DS", "<||DSML", "<||DSML||", "<||DSML||c",
		"<||DSML||inv", "<||DSML||invoke name=\"fi", "</||DSML||ca", "< || DSML || p"}
	for _, s := range oui {
		if !dsmlTagPrefixish(s) {
			t.Errorf("prefixish(%q) = false, attendu true", s)
		}
	}
	non := []string{"", "hello", "<||xyz", "<||DSML||xyz", "a < b", "<div>", "<||DSML||calls>"}
	for _, s := range non {
		if dsmlTagPrefixish(s) {
			t.Errorf("prefixish(%q) = true, attendu false", s)
		}
	}
}

// TestStreamFilter_BlocCoupe simule un provider qui découpe une balise sur
// plusieurs chunks : rien du DSML ne doit fuiter, le texte autour passe.
func TestStreamFilter_BlocCoupe(t *testing.T) {
	f := &dsmlStreamFilter{}
	chunks := []string{
		"Voici le résultat : <||",
		"DSML||calls><||DSML||invoke ",
		"name=\"Read\"><||DSML||parameter name=\"file_path\">",
		"docs/a.md</||DSML||parameter></||DSML||invoke></||DSML||",
		"calls> et voilà la suite.",
	}
	var sb strings.Builder
	for _, c := range chunks {
		sb.WriteString(f.push(c))
	}
	sb.WriteString(f.flush())
	got := sb.String()
	if strings.Contains(got, "DSML") {
		t.Errorf("fuite DSML dans le flux : %q", got)
	}
	if got != "Voici le résultat :  et voilà la suite." {
		t.Errorf("texte filtré inattendu : %q", got)
	}
}

func TestStreamFilter_TexteNormalAvecChevrons(t *testing.T) {
	f := &dsmlStreamFilter{}
	got := f.push("a < b et c > d") + f.flush()
	if got != "a < b et c > d" {
		t.Errorf("texte normal altéré : %q", got)
	}
}

func TestStreamFilter_FauxDebutDeBalise(t *testing.T) {
	// "<||D" ressemble à un début de balise mais n'en est pas une.
	f := &dsmlStreamFilter{}
	got := f.push("prix <||D euros") + f.flush()
	if got != "prix <||D euros" {
		t.Errorf("faux positif : %q", got)
	}
}

func TestStreamFilter_BlocNonFermeJeteAuFlush(t *testing.T) {
	f := &dsmlStreamFilter{}
	visible := f.push("Début <||DSML||calls><||DSML||invoke name=\"Read\">")
	visible += f.flush()
	if visible != "Début " {
		t.Errorf("visible = %q, attendu %q", visible, "Début ")
	}
}

func TestStreamFilter_SansDSML(t *testing.T) {
	f := &dsmlStreamFilter{}
	in := "Réponse parfaitement normale, sans aucun balisage."
	if got := f.push(in) + f.flush(); got != in {
		t.Errorf("texte altéré : %q", got)
	}
}
