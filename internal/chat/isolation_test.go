package chat

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestWrapCommandNone(t *testing.T) {
	tokens := []string{"ls", "-la", "src"}
	got := wrapCommand(IsolationNone, "/tmp/ws", tokens)
	if !reflect.DeepEqual(got, tokens) {
		t.Fatalf("got %v", got)
	}
}

func TestWrapCommandBwrap(t *testing.T) {
	tokens := []string{"echo", "hi"}
	got := wrapCommand(IsolationBwrap, "/tmp/ws", tokens)
	if len(got) == 0 || got[0] != "bwrap" {
		t.Fatalf("argv = %v", got)
	}
	sep := -1
	for i, a := range got {
		if a == "--" {
			sep = i
			break
		}
	}
	if sep < 0 {
		t.Fatalf("separateur -- absent: %v", got)
	}
	if !reflect.DeepEqual(got[sep+1:], tokens) {
		t.Fatalf("commande interne = %v", got[sep+1:])
	}
	if !contains(got, "/tmp/ws") {
		t.Fatalf("bind du workspace absent: %v", got)
	}
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func TestResolveIsolationNoneAndInvalid(t *testing.T) {
	if got, err := ResolveIsolation("", "/tmp/ws"); err != nil || got != IsolationNone {
		t.Fatalf("vide -> %q %v", got, err)
	}
	if got, err := ResolveIsolation(IsolationNone, "/tmp/ws"); err != nil || got != IsolationNone {
		t.Fatalf("none -> %q %v", got, err)
	}
	if _, err := ResolveIsolation("containers", "/tmp/ws"); err == nil {
		t.Fatal("mode invalide doit echouer")
	}
}

func TestResolveIsolationExplicitBwrap(t *testing.T) {
	if _, err := exec.LookPath("bwrap"); err != nil {
		if _, err := ResolveIsolation(IsolationBwrap, t.TempDir()); err == nil {
			t.Fatal("bwrap absent mais ResolveIsolation a reussi")
		}
		return
	}
	if got, err := ResolveIsolation(IsolationBwrap, t.TempDir()); err != nil || got != IsolationBwrap {
		t.Logf("bwrap present mais sonde indisponible: %v", err)
	}
}

func TestResolveIsolationAuto(t *testing.T) {
	got, err := ResolveIsolation("auto", t.TempDir())
	if err != nil {
		t.Fatalf("auto: %v", err)
	}
	if got != IsolationNone && got != IsolationBwrap {
		t.Fatalf("auto -> %q", got)
	}
}
