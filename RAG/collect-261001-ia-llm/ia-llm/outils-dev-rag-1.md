---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-1
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei", "Microsoft", "OpenAI"]
dates: ["2026-09-23", "2026-09-27"]
keywords: ["agent", "agents", "chatgpt", "copilot", "embedding", "embeddings", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1, 167]
sha256: d43a28edc6dd067fd3376ca9e6c4f09936219ca85a77c29738e0c7c049ec2496
---

# Outils dev + ingénierie RAG (chunk & corpus)

> **Public :** Zelef — chef de service systèmes & énergies, sysadmin.
> **Objectif :** maîtriser l'outillage de dev moderne et l'ingénierie RAG
> pour ton pipeline personnel (scripts `collect_*`, Playwright, embeddings
> `text-embedding-3-small` d'OpenAI, corpus de guides + CLI Huawei + articles).
> **État des vérifications :** recherches web effectuées le 27/09/2026.
> Ce qui est marqué **« à vérifier »** n'a pas pu être confirmé — le reste
> est sourcé des docs officielles ou de sources à jour.

## Sommaire

- **PARTIE A — OUTILS DEV**
  - VS Code : §1 à §10
  - GitHub : §11 à §23
  - tmux : §24 à §33
  - ngrok & tunnels : §34 à §44
  - Playwright : §45 à §55
  - Puppeteer : §56 à §59
  - Supabase (pgvector) : §60 à §73
- **PARTIE B — INGÉNIERIE RAG**
  - CHUNK : §74 à §92
  - CORPUS : §93 à §109
  - Prompt engineering : §110 à §124
  - À venir — annonces vérifiées au 27/09/2026 : §125 à §130
  - Fil rouge + glossaire + quiz + pièges (bloc final) : §131 et suivantes

---

# PARTIE A — OUTILS DEV

## 1. Pourquoi VS Code pour le dev IA/agentique (2026)

VS Code est devenu en 2026 l'interface standard pour travailler **avec** des
agents codeurs : fenêtre Agents dédiée, mode agent dans Copilot, sessions
distantes qui survivent aux déconnexions, MCP (Model Context Protocol) natif
pour brancher des outils externes. Pour toi, le trio gagnant :

| Besoin | Fonction VS Code | Intérêt concret |
|---|---|---|
| Coder avec un agent | Mode Agent de Copilot / Codex | L'agent édite, teste, corrige en boucle |
| Développer sur ton serveur | Remote-SSH | Zéro différence avec du local, zéro rsync |
| Brancher tes outils | MCP (`mcp.servers`) | L'agent pilote Supabase, ton API, etc. |
| Standardiser ton env | Profils + devcontainer | Même config sur PC et serveur |

Version à viser : **1.139+** (stable du 23/09/2026 — support des agents
distants en Stable, vérifié).

## 2. Installation et settings.json de base (vérifié)

Installe depuis [code.visualstudio.com](https://code.visualstudio.com/)
(paquet `.deb`/repo Microsoft sur Debian/Ubuntu). Puis `settings.json`
(`Ctrl+Shift+P` → *Preferences: Open User Settings (JSON)*) :

```jsonc
{
  // ── Éditeur ──────────────────────────────────────────────
  "editor.fontSize": 14,
  "editor.tabSize": 4,
  "editor.formatOnSave": true,
  "editor.rulers": [100, 120],
  "editor.wordWrap": "on",
  "files.autoSave": "afterDelay",
  "files.autoSaveDelay": 1000,
  "files.trimTrailingWhitespace": true,
  "files.insertFinalNewline": true,

  // ── Terminal intégré ─────────────────────────────────────
  "terminal.integrated.defaultProfile.linux": "bash",
  "terminal.integrated.fontSize": 13,
  "terminal.integrated.scrollbackLineCount": 20000,

  // ── Python ───────────────────────────────────────────────
  "python.defaultInterpreterPath": "/usr/bin/python3", // adapte (venv !)
  "[python]": {
    "editor.defaultFormatter": "charliermarsh.ruff" // à vérifier : id exact
  },

  // ── Git ──────────────────────────────────────────────────
  "git.autofetch": true,
  "git.confirmSync": false,

  // ── Copilot / mode agent (vérifié sept 2026) ─────────────
  "github.copilot.enable": {
    "*": true,
    "markdown": true,
    "plaintext": false
  },
  "github.copilot.chat.agent.autoRunTerminalCommands": false, // true = l'agent lance seul
  "github.copilot.chat.agent.autoFix": true,
  "chat.agent.maxRequests": 100,          // nb max d'appels outils par tour
  "github.copilot.chat.codeGeneration.useInstructionFiles": true,
  "github.copilot.advanced.codeReferenceFilter": true,

  // ── MCP : serveurs d'outils externes (vérifié sept 2026) ─
  "mcp.servers": {
    // exemple : ajouter plus tard ton serveur Postgres/Supabase
  },

  // ── Fichiers d'instructions custom (vérifié sept 2026) ───
  "chat.instructionsFilesLocations": {
    ".github/instructions": true
  }
}
```

> **Règle d'or :** ne mets jamais de clé API dans `settings.json`
> (il est parfois synchronisé via Settings Sync). Variables d'environnement
> ou fichier `.env` local, toujours.

## 3. Extensions indispensables en 2026 (vérifiées)

Installe en ligne de commande (rapide, scriptable) :

```bash
# ── Python & data ────────────────────────────────────────────
code --install-extension ms-python.python            # Python (Pylance, debug) — vérifié
code --install-extension ms-toolsai.jupyter          # notebooks — vérifié (écosystème)

# ── Remote ───────────────────────────────────────────────────
code --install-extension ms-vscode-remote.remote-ssh # Remote-SSH (Microsoft) — vérifié
# ou le pack complet :
# code --install-extension ms-vscode-remote.vscode-remote-extensionpack

# ── Git ──────────────────────────────────────────────────────
code --install-extension eamodio.gitlens              # GitLens — vérifié

# ── Agents IA ────────────────────────────────────────────────
code --install-extension github.copilot               # GitHub Copilot + chat + agent
code --install-extension openai.chatgpt               # Codex, l'agent d'OpenAI — vérifié sept 2026

# ── Vérifier un id avant d'installer ─────────────────────────
code --list-extensions | grep -i python
```

Extensions « à vérifier » (je ne les ai pas pu confirmer à 100 % pour 2026,
vérifie l'éditeur sur le Marketplace avant d'installer) :
- **Ruff** (`charliermarsh.ruff` — à vérifier) : lint Python ultra-rapide.
- **Error Lens** : erreurs inline.
- **Even Better TOML** : pour `pyproject.toml`.

> **Méfiance typosquatting :** vérifie toujours l'**éditeur** (Microsoft,
> OpenAI, GitHub) et l'**identifiant exact** avant d'installer une extension
> au nom proche d'une extension connue.

## 4. Profils : un profil « RAG » dédié

Les **profils** (`Ctrl+Shift+P` → *Profiles: Create Profile*) permettent de
basculer entre configurations complètes (extensions + settings + snippets).

```text
Profil "RAG" :
  extensions : Python, Jupyter, Remote-SSH, Copilot, GitLens
  settings   : formatOnSave, ruff, interpréteur = ~/rag/.venv/bin/python
Profil "Sysadmin" :
  extensions : Remote-SSH, Ansible (à vérifier), YAML
  settings   : terminal par défaut, thèmes sobres
```

Exporte ton profil : *Profiles: Export Profile* → fichier `.code-profile`
à versionner dans ton repo (voir §13). Quand tu changeras de machine,
import = 30 secondes.

## 5. Snippets : tes raccourcis de code RAG

Fichier → Préférences → Snippets utilisateur → `python.json` :

