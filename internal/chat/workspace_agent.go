package chat

import (
	"strings"

	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
	"cetas-lite/internal/workspace"
)

// SetWorkspaceManager branche le gestionnaire de projets (upload local /
// SFTP) sur le moteur. L'agent travaille alors sur le FS du projet choisi.
func (e *Engine) SetWorkspaceManager(m *workspace.Manager) {
	e.mu.Lock()
	e.wsProjects = m
	e.mu.Unlock()
}

func (e *Engine) workspaceManager() *workspace.Manager {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.wsProjects
}

// agentSandbox construit le sandbox du tour : projet (local ou SFTP) si
// in.ProjectID est renseigné, sinon l'espace partagé historique.
func (e *Engine) agentSandbox(in TurnInput) (*Sandbox, error) {
	var sb *Sandbox
	if pid := strings.TrimSpace(in.ProjectID); pid != "" {
		if wm := e.workspaceManager(); wm != nil {
			fsys, err := wm.OpenFS(pid)
			if err != nil {
				return nil, err
			}
			sb = NewSandboxFS(fsys)
		}
	}
	if sb == nil {
		root, err := userWorkspacePath(e.workspace, in.User)
		if err != nil {
			return nil, err
		}
		var err2 error
		sb, err2 = NewSandbox(root)
		if err2 != nil {
			return nil, err2
		}
	}
	sb.AllowScript = e.scriptAllowed()
	sb.Isolation = e.isolation()
	sb.GitHubToken = e.githubTokenFunc()
	return sb, nil
}

// githubTokenFunc retourne une fonction lisant le token GitHub connecté
// (déchiffré depuis le store), ou "" si aucun compte n'est connecté.
func (e *Engine) githubTokenFunc() func() string {
	st := e.st
	return func() string {
		if st == nil {
			return ""
		}
		ct, ok := st.GetSecret("github_token")
		if !ok {
			return ""
		}
		v, err := vault.Open(st)
		if err != nil {
			return ""
		}
		b, err := v.Decrypt(ct, []byte("github_token"))
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(b))
	}
}

// githubLogin retourne le login GitHub connecté (pour le prompt), ou "".
func githubLogin(st *store.Store) string {
	if st == nil {
		return ""
	}
	if b, ok := st.GetMeta("github_login"); ok {
		return strings.TrimSpace(string(b))
	}
	return ""
}
