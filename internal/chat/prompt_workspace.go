package chat

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/vfs"
)

// Cache du snapshot workspace : le walk FS (FlatList) est refait au plus
// une fois par minute et par espace. Le snapshot est une orientation pour
// le modèle (il doit appeler Ls pour explorer), pas une vérité temps
// réel : un décalage < 60 s est sans conséquence, et les fichiers que le
// modèle vient lui-même de créer sont dans l'historique du tour.
var snapshotCache = struct {
	sync.Mutex
	entries map[string]snapshotEntry
}{entries: map[string]snapshotEntry{}}

type snapshotEntry struct {
	msg provider.Message
	at  time.Time
}

// snapshotCacheTTL : durée de validité d'un snapshot en cache.
const snapshotCacheTTL = 60 * time.Second

// workspaceSnapshotMessage construit le message système décrivant le
// workspace du tour : nom, mode, liste des fichiers et règles strictes de
// lecture. C'est ce qui permet au modèle de ne jamais se perdre : il voit
// la structure réelle au lieu de la deviner.
func workspaceSnapshotMessage(ctx context.Context, sb *Sandbox, projectName string) provider.Message {
	mode := "local"
	if sb.Remote() {
		mode = "remote (SFTP)"
	}
	name := strings.TrimSpace(projectName)
	if name == "" {
		name = sb.Root()
	}
	key := sb.Root() + "\x00" + mode + "\x00" + name

	snapshotCache.Lock()
	if e, ok := snapshotCache.entries[key]; ok && time.Since(e.at) < snapshotCacheTTL {
		snapshotCache.Unlock()
		return e.msg
	}
	snapshotCache.Unlock()

	msg := buildWorkspaceSnapshot(ctx, sb, name, mode)

	snapshotCache.Lock()
	snapshotCache.entries[key] = snapshotEntry{msg: msg, at: time.Now()}
	snapshotCache.Unlock()
	return msg
}

// buildWorkspaceSnapshot fait le walk FS et assemble le message. Le cache
// ci-dessus évite de le refaire à chaque tour.
func buildWorkspaceSnapshot(ctx context.Context, sb *Sandbox, name, mode string) provider.Message {
	var b strings.Builder
	fmt.Fprintf(&b, "WORKSPACE: \"%s\" [%s]\n", name, mode)
	fmt.Fprintf(&b, "Root: %s\n", sb.Root())

	cctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// Plafond volontairement bas : ce snapshot est injecté à CHAQUE tour,
	// chaque fichier listé coûte des tokens d'entrée. 120 entrées suffisent
	// à orienter le modèle ; au-delà, il explore avec Ls/Glob.
	files, truncated, err := vfs.FlatList(cctx, sb.FS(), 120)
	if err != nil {
		b.WriteString("File listing unavailable (" + err.Error() + ") — call Ls first, never guess paths.\n")
	} else if len(files) == 0 {
		b.WriteString("Workspace is empty.\n")
	} else {
		sort.Strings(files)
		b.WriteString(fmt.Sprintf("Files (%d):\n", len(files)))
		for _, f := range files {
			b.WriteString(f + "\n")
		}
		if truncated {
			b.WriteString("… (truncated: call Ls or Glob to explore further)\n")
		}
	}

	b.WriteString("RULES (mandatory):\n" +
		"- All paths are relative to the workspace root, slash-separated, never absolute.\n" +
		"- NEVER guess or invent a path. If unsure, use Glob (patterns) or Ls; a wrong path wastes the whole task.\n" +
		"- Read big files with Read offset/limit pagination; never dump a whole large file at once.\n" +
		"- After creating/deleting files, re-run Ls before referencing new paths.\n" +
		"- Prefer Edit (one exact replacement) over Write for existing files.\n" +
		"- Bash: single commands only, no pipes/redirection; stay inside the workspace.\n")
	if sb.Remote() {
		b.WriteString("- Remote workspace: each tool call has network latency — batch reads, avoid repeated huge Greps.\n")
	}
	return provider.Message{Role: "system", Content: b.String()}
}

// githubPromptMessage décrit l'accès GitHub connecté pour le tour.
func githubPromptMessage(login string) provider.Message {
	return provider.Message{Role: "system", Content: "GITHUB: connected as \"" + login + "\" (token injected automatically, never ask for it).\n" +
		"RULES:\n" +
		"- You MAY use git: status, diff, log, add, commit, push, pull on github.com remotes.\n" +
		"- Before push: run git status and git diff --stat, summarize what will be pushed.\n" +
		"- NEVER force-push (--force) without explicit user approval.\n" +
		"- Commit messages: short imperative summary, no noise.\n" +
		"- If not in a git repo, say so instead of inventing history."}
}
