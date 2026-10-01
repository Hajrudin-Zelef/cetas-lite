---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-4
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: ["2026-07-28"]
keywords: ["mcp"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [475, 662]
sha256: 13a31377933d47c138db809d93abca3aaf57d28916dabfb3ce2b2abba0712c91
---

# MCP, LSP, code-server, websearch — Guide pratique

- **Pagination** : `tools/list`, `resources/list`, `prompts/list` acceptent
  un `cursor` ; le serveur renvoie `nextCursor`. Obligatoire dès que le
  catalogue dépasse quelques dizaines d'entrées (sinon le contexte du
  modèle explose).
- **Complétion** (`completion/complete`) : autocomplétion des arguments
  (ex : le client demande les valeurs possibles de `nom_doc`). Pratique
  pour les hosts avec UI.
- **Notifications** : `notifications/tools/list_changed` et
  `notifications/resources/list_changed` préviennent le client de
  re-lister. `notifications/cancelled` annule. `notifications/progress`
  suit l'avancement (utile en Streamable HTTP/SSE).

## 25. Gestion d'erreurs : les codes à connaître

| Code JSON-RPC | Sens | Quand le renvoyer |
|---|---|---|
| `-32700` | Parse error | JSON illisible |
| `-32600` | Invalid request | Message mal formé |
| `-32601` | Method not found | Méthode inconnue |
| `-32602` | Invalid params | Arguments invalides |
| `-32603` | Internal error | Exception non gérée côté serveur |

Et côté **métier** : ne renvoie PAS une erreur JSON-RPC quand l'outil échoue
« normalement » (fichier absent, SQL en erreur). Renvoie un résultat avec
`"isError": true` et un message clair dans `content`. Le modèle sait alors
lire l'erreur et s'adapter ; une erreur protocole, elle, casse l'appel.

💡 Soigne les messages d'erreur métier : c'est le **seul feedback** que le
modèle reçoit pour se corriger (« Table `onduleurs` introuvable, tables
disponibles : ... »).

## 26. Checklist : mon serveur est-il conforme (spec 2026-07-28) ?

- [ ] Répond à `server/discover` (identité, versions, capacités)
- [ ] Lit `_meta` (protocolVersion, clientInfo) sur chaque requête
- [ ] stdio : stdout = JSON-RPC uniquement, logs sur stderr
- [ ] Streamable HTTP : un seul endpoint, `202` sur notifications, SSE pour le progress
- [ ] Pas de `initialize`/`Mcp-Session-Id` imposé, pas de `sampling`, pas de `roots`
- [ ] `tools/list` paginé si > 50 outils
- [ ] Descriptions d'outils précises + annotations honnêtes
- [ ] Erreurs métier → `isError: true`, pas d'exception protocole
- [ ] Auth sur tout endpoint distant ; anti-rebinding en local
- [ ] Version du protocole annoncée = une version existante (ex : `2026-07-28`)

# PARTIE II — ÉCRIRE SON SERVEUR MCP

## 27. Choisir son SDK

| SDK | Paquet | Quand le choisir |
|---|---|---|
| **Python** | `mcp` (PyPI) | Ton écosystème (scripts RAG, Ansible) ; le plus simple |
| **TypeScript** | `@modelcontextprotocol/sdk` (+ paquets v2) | Serveur web existant en Node, intégration Express/Fastify/Hono |
| **Autres** (Go, Rust, Java, Kotlin, Swift...) | SDKs officiels | Contrainte perf ou écosystème — pas nécessaire pour commencer |

💡 Pour Zeef : **Python**. Tes scripts de collecte RAG sont en Python, ton
corpus est local, et le SDK Python est le plus documenté. Le SDK TypeScript
ne devient intéressant que si tu dois greffer MCP sur une app Node existante.

État des versions (fin sept 2026, 📌 à re-vérifier le jour J) :
- Python : `mcp` 1.27+ stable ; **v2 en bêta** (`2.0.0b1`) qui renomme
  `FastMCP` → `MCPServer` et ajoute un vrai `Client` unifié.
- TypeScript : 1.29+ stable ; **v2** sortie avec paquets éclatés
  (`@modelcontextprotocol/server`, `/client`, `/node`, `/express`...),
  `registerTool` remplace `tool()`.

L'exemple ci-dessous utilise l'API **v2** (`MCPServer`), qui est l'avenir ;
un encadré donne l'équivalence v1 (`FastMCP`).

## 28. Installation Python : la voie propre (uv)

```bash
# 1. Installer uv (gestionnaire Python moderne, rapide)
curl -LsSf https://astral.sh/uv/install.sh | sh

# 2. Créer le projet
mkdir -p ~/mcp/doc-server && cd ~/mcp/doc-server
uv init --bare .          # ou: uv venv && source .venv/bin/activate

# 3. Installer le SDK (+ CLI : inspector, run, dev)
uv add "mcp[cli]"

# Vérification
uv run python -c "import mcp; print(mcp.__version__)"
```

Alternative pip classique : `pip install "mcp[cli]"`. L'extra `[cli]`
fournit la commande `mcp` (`mcp dev`, `mcp run`, `mcp inspect`...).

📌 Si tu es encore en SDK v1 : `from mcp.server.fastmcp import FastMCP`
au lieu de `from mcp.server import MCPServer`. Le reste de l'exemple est
identique (`@mcp.tool()`, `@mcp.resource()`, `@mcp.prompt()` existent
dans les deux).

## 29. Serveur Python COMPLET : le serveur doc de Zelef (tools)

Scénario : exposer ton corpus de guides (`~/workspace/user/files/*.md`)
via 3 tools : `list_docs`, `search_docs`, `read_doc`. Fichier
`~/mcp/doc-server/server.py` :

```python
"""Serveur MCP 'doc-rag' : expose le corpus documentaire local de Zelef.

Transport par défaut : stdio (lancé par le client).
Test :  uv run mcp dev server.py   (Inspector)
        uv run mcp run server.py --transport streamable-http
"""
from __future__ import annotations

import re
from pathlib import Path

from mcp.server import MCPServer   # SDK v2 ; v1 : from mcp.server.fastmcp import FastMCP as MCPServer

# --- Config ---------------------------------------------------------------
DOC_DIR = Path.home() / "workspace" / "user" / "files"
mcp = MCPServer("doc-rag")          # nom du serveur, visible par les clients


def _iter_docs():
    """Yield (nom, Path) pour chaque .md du corpus, trié."""
    if not DOC_DIR.is_dir():
        return
    for p in sorted(DOC_DIR.glob("*.md")):
        yield p.name, p


def _snippets(text: str, query: str, width: int = 200) -> list[str]:
    """Extraits autour de chaque occurrence (insensible à la casse)."""
    out = []
    for m in re.finditer(re.escape(query), text, re.IGNORECASE):
        start = max(0, m.start() - width // 2)
        end = min(len(text), m.end() + width // 2)
        out.append("..." + text[start:end].replace("\n", " ") + "...")
        if len(out) >= 3:
            break
    return out


# --- Tools ----------------------------------------------------------------
@mcp.tool()
def list_docs() -> list[str]:
    """Liste les noms des documents disponibles dans le corpus local."""
    return [name for name, _ in _iter_docs()]


@mcp.tool()
def search_docs(query: str, limit: int = 5) -> list[dict]:
    """Recherche plein-texte (insensible à la casse) dans le corpus.

    Retourne les documents les plus pertinents avec des extraits.
    `limit` borne le nombre de documents renvoyés (1-20).
    """
    limit = max(1, min(20, limit))
    scored = []
    q = query.lower()
    for name, path in _iter_docs():
        try:
            text = path.read_text(encoding="utf-8", errors="replace")
        except OSError:
            continue
        count = text.lower().count(q)
        if count:
            scored.append((count, name, text))
    scored.sort(reverse=True)
    return [
        {"document": name, "occurrences": n, "extraits": _snippets(t, query)}
        for n, name, t in scored[:limit]
    ]


@mcp.tool()
def read_doc(name: str, max_chars: int = 20000) -> str:
    """Lit un document du corpus par son nom (voir list_docs).

    `max_chars` borne la taille renvoyée pour protéger le contexte du modèle.
    """
    path = DOC_DIR / name
    # 🔒 anti path-traversal : on reste strictement dans DOC_DIR
    if not path.resolve().is_relative_to(DOC_DIR.resolve()) or path.suffix != ".md":
        raise ValueError(f"Document refusé : {name!r}")
    if not path.is_file():
        raise FileNotFoundError(f"Document introuvable : {name!r}")
    text = path.read_text(encoding="utf-8", errors="replace")
    if len(text) > max_chars:
        text = text[:max_chars] + f"\n\n[... tronqué à {max_chars} caractères]"
    return text
```

