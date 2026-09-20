package chat

import "testing"

func TestMissingArg(t *testing.T) {
	args := map[string]any{"file_path": "a.go", "content": "x"}
	if f := missingArg(args, "file_path", "content"); f != "" {
		t.Fatalf("attendu vide, got %q", f)
	}
	if f := missingArg(args, "file_path", "old"); f != "old" {
		t.Fatalf("attendu old, got %q", f)
	}
	if f := missingArg(map[string]any{"file_path": "  "}, "file_path"); f != "file_path" {
		t.Fatalf("chaine vide non detectee, got %q", f)
	}
	if f := missingArg(map[string]any{}, "pattern"); f != "pattern" {
		t.Fatalf("attendu pattern, got %q", f)
	}
}

func TestVerifyCommandHeuristic(t *testing.T) {
	yes := []string{
		"go build ./...", "go test ./... -race", "go vet ./...",
		"npm test", "cargo test", "pytest -q", "make ci",
	}
	for _, c := range yes {
		if !verifyCommandHeuristic(c) {
			t.Errorf("attendu verification pour %q", c)
		}
	}
	no := []string{
		"ls -la", "go version", "echo hello", "git status",
		"python3 script.py",
	}
	for _, c := range no {
		if verifyCommandHeuristic(c) {
			t.Errorf("faux positif pour %q", c)
		}
	}
}

func TestTrackModification(t *testing.T) {
	modified := map[string]bool{}
	trackModification("Write", map[string]any{"file_path": "a.go"}, modified)
	trackModification("Edit", map[string]any{"file_path": "b.go"}, modified)
	trackModification("Read", map[string]any{"file_path": "c.go"}, modified)
	trackModification("Bash", map[string]any{"command": "ls"}, modified)
	if len(modified) != 2 || !modified["a.go"] || !modified["b.go"] {
		t.Fatalf("suivi incorrect: %v", modified)
	}
	if got := modifiedList(modified); len(got) != 2 {
		t.Fatalf("liste incorrecte: %v", got)
	}
}

func TestRequiredFields(t *testing.T) {
	cases := map[string][]string{
		"Read": {"file_path"}, "Cat": {"file_path"},
		"Write": {"file_path", "content"},
		"Edit":  {"file_path", "old", "new"},
		"Grep":  {"pattern"}, "Glob": {"pattern"},
		"Bash":      {"command"},
		"RunScript": {"language", "code"},
		"Echo":      {"text"},
		"Mkdir":     {"path"},
		"Mv":        {"src", "dst"},
		"Sed":       {"expression"},
		"Awk":       {"program"},
		"Curl":      {"url"},
		// Outils sans champ requis au niveau Execute.
		"Ls": {}, "Tree": {}, "TodoWrite": {},
		"GitHubIssues": {}, "GitHubPRCreate": {}, "Inconnu": {},
	}
	for name, want := range cases {
		got := requiredFields(name)
		if len(got) != len(want) {
			t.Fatalf("%s: attendu %v, got %v", name, want, got)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("%s: attendu %v, got %v", name, want, got)
			}
		}
	}
	// Le cas observe : Edit sans "new" est detecte avant approbation.
	args := map[string]any{"file_path": "a.txt", "old": "foo"}
	if f := missingArg(args, requiredFields("Edit")...); f != "new" {
		t.Fatalf("Edit sans new: attendu \"new\", got %q", f)
	}
	args["new"] = "bar"
	if f := missingArg(args, requiredFields("Edit")...); f != "" {
		t.Fatalf("Edit complet: attendu vide, got %q", f)
	}
}
