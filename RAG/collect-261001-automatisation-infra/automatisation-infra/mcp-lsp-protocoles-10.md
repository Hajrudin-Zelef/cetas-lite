---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-10
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1537, 1720]
sha256: 910ff4c775dc874301361129a43e3960001ec27188398dd5f161ded57a302078
---

# MCP, LSP, code-server, websearch — Guide pratique

1. Client → `initialize` : `processId`, `rootUri` (dossier projet),
   `capabilities` (ce que l'éditeur sait afficher), `initializationOptions`.
2. Serveur → ses `capabilities` (quelles méthodes il supporte :
   `definitionProvider`, `referencesProvider`...).
3. Client → `initialized` (notification).
4. ... session de travail ...
5. Client → `shutdown` puis notification `exit` — le serveur doit se terminer.

⚠️ Le `rootUri` détermine le « workspace » : un serveur lancé sur le mauvais
dossier = diagnostics absurdes. C'est la cause n°1 des « le LSP dit n'importe
quoi » (voir §86).

## 75. Les capacités : ce que tu utilises sans le savoir

| Méthode | Ce qu'elle fait | Raccourci VS Code typique |
|---|---|---|
| `textDocument/hover` | Type + doc au survol | Survol |
| `textDocument/definition` | Aller à la définition | F12 / Ctrl+clic |
| `textDocument/references` | Toutes les références | Maj+F12 |
| `textDocument/completion` | Complétion contextuelle | Ctrl+Espace |
| `textDocument/diagnostic` | Erreurs/avertissements | Soulignés |
| `textDocument/rename` | Renommage sûr | F2 |
| `textDocument/formatting` | Formatage | Maj+Alt+F |
| `textDocument/codeAction` | Quick fixes, imports auto | Ctrl+. |
| `workspace/symbol` | Recherche de symbole globale | Ctrl+T |
| `textDocument/documentSymbol` | Plan du fichier | Maj+Ctrl+O |

💡 Retiens surtout `definition`, `references`, `hover`, `diagnostic` :
ce sont **exactement** les 4 opérations que les agents codeurs appellent
en boucle (voir §87).

## 76. Tableau des serveurs de langage (vérifié sept 2026)

| Langage | Serveur | Installation |
|---|---|---|
| Python | **pyright** (`pyright-langserver --stdio`) | `npm i -g pyright` ou `uvx pyright` |
| Python (alt.) | **basedpyright** (fork commu.) | `uvx basedpyright` |
| TypeScript/JS | **typescript-language-server** | `npm i -g typescript-language-server typescript` |
| Go | **gopls** | `go install golang.org/x/tools/gopls@latest` |
| Rust | **rust-analyzer** | `rustup component add rust-analyzer` |
| C/C++ | **clangd** | `apt install clangd` / avec LLVM |
| Java | **jdtls** (Eclipse JDT LS) | binaire/lanceur dédié |
| Kotlin | **kotlin-lsp** (JetBrains, officiel récent) | release GitHub, nécessite Java |
| Lua | **lua-language-server** | release GitHub |
| Bash | **bash-language-server** | `npm i -g bash-language-server` |
| YAML | **yaml-language-server** | `npm i -g yaml-language-server` |
| TOML | **taplo** (`taplo lsp stdio`) | `cargo install taplo-cli` |
| JSON | **vscode-json-language-server** | `npm i -g vscode-json-languageserver` |
| HTML/CSS | **vscode-html/css-language-server** | idem (vscode-langservers-extracted) |
| Markdown | **marksman** | `brew install marksman` / binaire |
| Dockerfile | **docker-langserver** | `npm i -g dockerfile-language-server-nodejs` |
| Ruby | **ruby-lsp** | `gem install ruby-lsp` |
| PHP | **intelephense** | `npm i -g intelephense` (gratuit partiel) |
| Zig | **zls** | binaire release |
| Dart/Flutter | **dart** (`dart language-server`) | avec le SDK Dart |
| Scala | **metals** | via coursier / VS Code |
| Swift | **sourcekit-lsp** | avec la toolchain Swift |
| C# | **Roslyn** / csharp-ls / OmniSharp | selon époque, via extension |
| Tailwind | **tailwindcss-language-server** | `npm i -g` / via extension |

📌 `basedpyright`, `kotlin-lsp` : existence confirmée, détails d'install
à vérifier le jour J (l'écosystème bouge vite).

## 77. Installation express (les 5 que tu croiseras)

```bash
# Python (le plus tolérant au code cassé — idéal pour agents)
npm i -g pyright
# ou sans npm global : uvx pyright-langserver --stdio

# Go
go install golang.org/x/tools/gopls@latest

# Rust
rustup component add rust-analyzer

# TypeScript
npm i -g typescript-language-server typescript

# C/C++
sudo apt install clangd
```

Vérifie que chaque binaire répond (le LSP parle sur stdin, mais la plupart
ont un `--version`) :
```bash
pyright-langserver --version; gopls version; rust-analyzer --version
```

## 78. VS Code : comment ça marche sous le capot

VS Code est **le** client LSP de référence (le protocole est né pour lui) :
- Les extensions de langage (Python, Go, rust-analyzer...) **embarquent**
  le serveur et le lancent en stdio automatiquement.
- Tu peux forcer le serveur externe : settings `"python.languageServer": "Pylance"`,
  `"rust-analyzer.server.path": "/usr/bin/rust-analyzer"`, etc.
- Le client envoie `didChange` à chaque frappe (souvent en mode incrémental)
  et affiche les diagnostics en quasi temps réel.
- Multi-racines : un serveur par dossier de workspace.

💡 Dans VS Code, 90 % des problèmes LSP se règlent par : bonne extension
installée + bon workspace ouvert + `Developer: Reload Window`.

## 79. Neovim : config LSP moderne (0.10+)

Neovim embarque un client LSP natif (plus besoin de coc.nvim) :

```lua
-- init.lua (Neovim 0.10+)
vim.lsp.enable('pyright')   -- cherche pyright-langserver sur le PATH
vim.lsp.enable('gopls')
vim.lsp.enable('rust_analyzer')

-- Raccourcis essentiels
vim.keymap.set('n', 'gd', vim.lsp.buf.definition)      -- aller à la définition
vim.keymap.set('n', 'gr', vim.lsp.buf.references)      -- références
vim.keymap.set('n', 'K',  vim.lsp.buf.hover)          -- doc au survol
vim.keymap.set('n', '<leader>rn', vim.lsp.buf.rename) -- renommer
vim.keymap.set('n', '<leader>ca', vim.lsp.buf.code_action)
```

Les configs par serveur vivent dans `lsp/<nom>.lua` (ex : `lsp/pyright.lua`
avec `cmd = {'pyright-langserver', '--stdio'}` et `root_markers`).
📌 La forme exacte (`vim.lsp.enable` vs `lspconfig`) dépend de ta version
de Neovim — vérifie `:h lsp`.

## 80. Cursor, Windsurf et les IDE agentiques

Ils réutilisent le client LSP de VS Code (ce sont des forks) et ajoutent
une couche agent : le modèle lui-même émet des requêtes LSP via des tools
(`definition`, `references`...). D'où l'importance du §87 : un agent sans
LSP « devine » le code au texte ; avec LSP, il **navigue** comme un humain
armé de F12.

## 81. Écrire un mini serveur de langage : le concept

Pas besoin d'en écrire un (il en existe pour tout), mais comprendre le
squelette aide au debug :

```
boucle:
  lire header "Content-Length: N"
  lire N octets JSON-RPC
  si method == "initialize" → répondre capabilities
  si method == "textDocument/didOpen" → indexer le fichier (en mémoire)
  si method == "textDocument/definition" → chercher le symbole, répondre URI+range
```

Les SDK existent (`pygls` en Python pour prototyper vite). Cas réel où ça
sert : langage maison / DSL interne (ex : ton format de fiches d'intervention)
→ un petit serveur LSP = complétion + diagnostics dans VS Code en un
après-midi.

## 82. Debug VS Code : les logs du serveur

1. `Ctrl+Maj+U` → panneau **Sortie** → liste déroulante → choisis le serveur
   (ex : « Python », « rust-analyzer »).
2. Pour le niveau trace, dans `settings.json` :
   ```json
   { "rust-analyzer.trace.server": "verbose",
     "python.languageServerLogLevel": "debug" }
   ```
   (chaque extension a sa clé ; cherche `<id>.trace.server` / `logLevel`.)
3. Cherche `initialize`, les `diagnostic` reçus, les erreurs de spawn.

⚠️ Si la liste ne montre aucun serveur : l'extension n'a pas démarré le
binaire → vérifie le PATH et les logs de l'extension elle-même
(« Log (Extension Host) »).

## 83. Debug Neovim : :LspLog

```vim
:LspInfo          " serveurs attachés au buffer courant, root_dir, cmd
:lua vim.lsp.set_log_level("debug")
:LspLog           " ouvre le log (~/.local/state/nvim/lsp.log)
:LspRestart       " relance les serveurs du buffer
```

Le log montre le JSON-RPC brut : idéal pour voir si le serveur répond ou
s'il est bloqué en indexation. Pour un debug fin, `tail -f` le fichier
pendant que tu édites.

## 84. Les pièges classiques (et leurs remèdes)

