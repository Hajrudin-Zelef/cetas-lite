---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-15
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Perplexity"]
dates: ["2026-07-28", "2026-09-27"]
keywords: ["mcp", "agent", "agentic", "agents", "benchmark", "claude", "perplexity"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [2351, 2490]
sha256: a2081d3e5e17b7166f9cb3a4bce995a385330fa72deca51f18bda572cc9cba6c
---

# MCP, LSP, code-server, websearch — Guide pratique

1. **Tools** (le modèle décide d'appeler), **Resources** (l'humain/host
   attache au contexte), **Prompts** (l'humain choisit le workflow).
2. Parce que stdout **est** le canal du protocole : chaque ligne doit être
   un message JSON-RPC complet. Un log parasite = JSON invalide = dialogue
   cassé. Les logs vont sur **stderr**.
3. HTTP+SSE répartissait une conversation sur **2 endpoints + une session**
   (sessions collantes obligatoires). Streamable HTTP = **1 endpoint**,
   chaque POST porte sa réponse, **sans état** → compatible load balancer,
   reconnexions simples.
4. **Dépréciés/supprimés** : sampling → appel direct de l'API LLM côté
   serveur ; roots → chemins en paramètres d'outils / URIs de resources ;
   `initialize`+`Mcp-Session-Id` → protocole sans état (`_meta` par requête
   + `server/discover`).
5. Des **instructions malveillantes cachées dans les métadonnées** du tool
   (description, noms de paramètres...), invisibles à l'humain mais lues par
   le modèle → exfiltration ou actions non voulues. La description fait partie
   du **contexte du modèle** : c'est une entrée d'instruction déguisée.
6. `destructiveHint: true` (peut détruire), `readOnlyHint: false`,
   `idempotentHint: false` → le host **doit demander confirmation**, et
   l'opérateur doit isoler ce tool (permissions deny par défaut).
7. `textDocument/references` (appelants) ; `textDocument/definition`
   (définition). (+ `hover` pour le type.)
8. Le serveur LSP n'utilise pas le **même environnement Python** que le code
   (mauvais venv/interpréteur sélectionné) → il ne voit pas les dépendances
   installées ailleurs.
9. Le terminal intégré et le « hot reload » passent par websocket ; sans les
   headers `Upgrade`, l'IDE charge mais le **terminal reste mort** et
   certaines extensions dysfonctionnent.
10. Les snippets « basic » sont souvent insuffisants → l'agent **relance des
    recherches** (×3 appels + tokens de raisonnement). Une API « chère » qui
    rend du contenu propre coûte moins cher **par réponse correcte** qu'une
    API pas chère qui oblige à multiplier les appels.

## 117. Pour aller plus loin (sources vérifiées sept 2026)

- Spec et docs officielles MCP : `modelcontextprotocol.io` (spec `2026-07-28`).
- SDK Python : `github.com/modelcontextprotocol/python-sdk` (README avec
  l'exemple `MCPServer` repris au §29).
- SDK TypeScript : `github.com/modelcontextprotocol/typescript-sdk`
  (README v2 : paquets éclatés, `registerTool`).
- Inspector : `@modelcontextprotocol/inspector` via npx.
- Sécurité : chapitre C10-MCP de l'OWASP AISVS
  (`github.com/owasp/aisvs` — tool poisoning, ATPA, rug pulls) ;
  disclosures Ox Security (avril 2026, faille STDIO) et Socket (février 2026,
  ver SANDWORM_MODE) — à lire avant toute prod.
- Catalogue de serveurs : `github.com/adrianlerer/50-essential-mcp-servers`
  (50 serveurs avec statuts officiel/communautaire).
- Comparateur d'API de recherche : `github.com/teionarr/agentic-search-arena`
  (benchmark aveugle Tavily/Exa/Brave/Perplexity... sur TES requêtes).
- Article comparatif prix 2026 : `sourceforge.net/articles/best-web-search-apis-to-integrate-for-ai-agents/`.
- LSP : spec sur `microsoft.github.io/language-server-protocol` (📌 URL à
  confirmer — chercher « LSP specification ») ; `neovim.io/doc/user/lsp.html`
  pour le client Neovim.
- code-server : dépôt `coder/code-server` sur GitHub, docs `coder.com/docs/code-server`.

## 118. Noms non identifiés / points à vérifier

- **« code-serve »** : confirmé = **code-server** (coquille probable). ✅ résolu.
- **GitLab MCP officiel** : existence d'un serveur GitLab officiel non
  vérifiée → si besoin, chercher « GitLab MCP server » ou utiliser l'API via
  un serveur maison. 📌 NON VÉRIFIÉ (sept 2026).
- **Image Docker `codercom/code-server`** : nom historique toujours vu en
  2026, mais le dépôt GitHub a migré vers `coder/code-server` → vérifier le
  nom exact de l'image sur Docker Hub avant usage. 📌 À VÉRIFIER.
- **Versions exactes** (SDK `mcp`, `@modelcontextprotocol/*`, code-server) :
  notées au 27/09/2026, elles auront bougé → `pip index versions mcp` /
  `npm view` le jour J. 📌 À VÉRIFIER.
- **Prix des API websearch** : ordres de grandeur sept 2026, les grilles
  changent vite → vérifier sur le site du fournisseur avant engagement. 📌 À VÉRIFIER.
- **Forme exacte du lancement stdio via `Client([...])`** (SDK Python v2) :
  non recoupée à 100 % → suivre `py.sdk.modelcontextprotocol.io` (section Clients). 📌 À VÉRIFIER.

## 119. Changelog du guide

- **2026-09-27 (v1)** : création. Recherche web préalable (27/09/2026) :
  spec MCP 2026-07-28, SDK Python `mcp` (v1.27+/v2 bêta), SDK TS
  (1.29+/v2), serveurs existants, sécu (OWASP AISVS, Ox avril 2026,
  SANDWORM fév 2026), LSP, code-server, API websearch + prix.
  Convention : les mentions « 📌 à vérifier » signalent ce qui bouge vite.

## 120. Architecture cible : le RAG de Zelef, branché en MCP

Comment tout ce guide s'assemble pour ton projet :

```
Tes guides .md (corpus) ──┐
Scripts collect_*.py ─────┼──► Serveur MCP « doc-rag » (Partie II)
Prompts diagnostic ───────┘         │ stdio (poste) / Streamable HTTP (équipe)
                                    ▼
                        ┌───────────────────────┐
                        │ Claude Code / agent   │◄── websearch (Brave/Tavily MCP)
                        │  + mcp-language-      │    pour la fraîcheur
                        │    server (pyright)   │
                        └───────────────────────┘
                                    │ code via code-server (navigateur)
                                    ▼
                        Réponse sourcée : corpus d'abord, web ensuite
```

Étapes concrètes, dans l'ordre :
1. Fais tourner `doc-rag` en stdio sur ton poste (§29-32) — 1 heure de travail.
2. Branche-le dans Claude Code (§35) + 5 serveurs utiles (§56).
3. Ajoute `mcp-language-server` quand tu fais refactorer tes scripts.
4. Passe `doc-rag` en Streamable HTTP + nginx (§33, §40) quand ton équipe
   l'utilise — avec auth et user dédié (§66, §68).
5. Ajoute le MCP Brave/Tavily (§51) avec budget (§109) pour la veille.
6. Expose le tout via code-server (§88-97) pour l'astreinte nomade.

💡 Ton avantage compétitif n'est pas le modèle — c'est **ton corpus**
(40+ guides vérifiés) + **tes prompts** (procédures terrain). MCP est juste
le tuyau standard qui les rend utilisables par n'importe quel agent.

## 121. Mémo commandes (une page à garder)

```bash
# --- SDK Python ---
uv add "mcp[cli]"                                   # installer
uv run mcp dev server.py                            # tester (Inspector)
uv run mcp run server.py --transport streamable-http --port 8000
printf '%s\n' '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | uv run python server.py

# --- Claude Code ---
claude mcp add doc-rag -- uv run --directory ~/mcp/doc-server server.py
claude mcp add --transport http doc-rag http://127.0.0.1:8000/mcp
claude mcp list                                     # vérifier

# --- Inspector ---
npx -y @modelcontextprotocol/inspector

# --- Serveurs utiles ---
claude mcp add fs -- npx -y @modelcontextprotocol/server-filesystem ~/workspace
claude mcp add playwright -- npx -y @playwright/mcp
claude mcp add context7 -- npx -y @upstash/context7-mcp

# --- LSP ---
npm i -g pyright && pyright-langserver --version
go install golang.org/x/tools/gopls@latest
rustup component add rust-analyzer
# Neovim : :LspInfo  :LspLog  :LspRestart

