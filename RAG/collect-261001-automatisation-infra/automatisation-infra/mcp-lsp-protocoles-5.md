---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-5
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["mcp", "agent", "claude", "memory"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [663, 878]
sha256: b83d17521a708f0650a5d4befa3f0cf8a5b98276bc62562a238dd1709351f2f8
---

# MCP, LSP, code-server, websearch — Guide pratique

Points à noter :
- Les **docstrings deviennent les descriptions** vues par le modèle : soigne-les.
- Les **annotations de type** (`str`, `int`, `list[dict]`) génèrent le
  `inputSchema` automatiquement. Zéro JSON Schema écrit à la main.
- `read_doc` lève `ValueError`/`FileNotFoundError` : le SDK les convertit
  en résultat `isError: true` avec le message — le modèle peut se corriger.
- Le garde anti-traversal (`is_relative_to`) est **non négociable** dès qu'un
  chemin vient du modèle (voir §66).

## 30. Le même serveur : resources (doc://)

Ajoute sous les tools, dans le même `server.py` :

```python
@mcp.resource("doc://corpus/{name}")
def doc_resource(name: str) -> str:
    """Expose chaque document du corpus comme une resource adressable."""
    # Réutilise la logique (et la sécurité) de read_doc
    return read_doc(name, max_chars=100_000)


@mcp.resource("doc://corpus/index")
def doc_index() -> str:
    """Index du corpus : nom + première ligne de chaque document."""
    lines = []
    for name, path in _iter_docs():
        try:
            first = path.read_text(encoding="utf-8", errors="replace").splitlines()
            title = first[0].lstrip("# ").strip() if first else "(vide)"
        except OSError:
            title = "(illisible)"
        lines.append(f"- {name} — {title}")
    return "# Index du corpus\n\n" + "\n".join(lines)
```

Le client peut alors attacher `doc://corpus/onduleurs_ups_guide.md` comme
contexte, sans passer par un appel d'outil. 💡 Les resources sont parfaites
pour les « gros contenus stables » : le modèle ne paie le coût en tokens
que si l'humain les joint.

## 31. Le même serveur : prompts (procédures terrain)

```python
@mcp.prompt()
def diagnostic_ups(symptome: str) -> str:
    """Procédure de diagnostic onduleur : guide l'agent étape par étape."""
    return (
        f"Tu es un technicien énergie senior. Symptôme observé : {symptome}.\n"
        "Suis STRICTEMENT cette procédure :\n"
        "1. Utilise search_docs pour retrouver les sections du guide onduleurs "
        "   liées au symptôme (mots-clés : alarme, bypass, batterie, défaut).\n"
        "2. Liste les 3 causes probables, ordonnées par probabilité.\n"
        "3. Pour chaque cause, donne UN test de confirmation réalisable sur site "
        "   (mesure, lecture d'afficheur, contrôle visuel).\n"
        "4. Termine par la conduite à tenir immédiate (sécurisation) puis le "
        "   plan d'action.\n"
        "N'invente aucune valeur de mesure : si une valeur manque, dis-le."
    )


@mcp.prompt()
def fiche_reflexe(equipement: str) -> str:
    """Génère une fiche réflexe terrain imprimable pour un équipement."""
    return (
        f"À partir du corpus (search_docs + read_doc), produis une fiche réflexe "
        f"pour « {equipement} » : 1 page max, codes défaut essentiels, "
        "procédures d'urgence, numéros utiles à compléter. "
        "Format Markdown, tableaux compacts."
    )
```

💡 C'est comme ça que tes 40+ guides deviennent **actionnables** : n'importe
quel agent connecté à `doc-rag` hérite de tes procédures, sans que tu
réécrives quoi que ce soit.

## 32. Tester en stdio avec l'Inspector

Le **MCP Inspector** est l'outil officiel de debug (interface web locale) :

```bash
cd ~/mcp/doc-server

# Terminal 1 : lance l'inspector (ouvre http://localhost:5173)
npx -y @modelcontextprotocol/inspector

# Dans l'UI :
#  - Transport : STDIO
#  - Command : uv
#  - Arguments : run --with mcp --with-file server.py ... (ou plus simple :)
```

Plus simple, la CLI du SDK :

```bash
# L'inspector pré-configuré sur ton serveur :
uv run mcp dev server.py
```

Dans l'UI tu peux : `tools/list` → appeler `search_docs` avec
`{"query": "bypass", "limit": 3}` → voir le JSON brut aller/retour.
🔍 C'est ici que tu repères les descriptions d'outils floues et les
schémas mal typés **avant** de brancher un vrai agent.

Test sans UI (pipe direct, voir §14) :

```bash
printf '%s\n' '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' \
  | uv run --no-project --with "mcp[cli]" python server.py | head -c 600
```

## 33. Servir en Streamable HTTP

```bash
# Le SDK expose directement le transport HTTP :
uv run mcp run server.py --transport streamable-http --port 8000
# Endpoint : http://127.0.0.1:8000/mcp
```

Ou dans le code (équivalent, pour un déploiement custom) :

```python
if __name__ == "__main__":
    # json_response=True : réponses JSON simples (pas de SSE) — utile derrière
    # un proxy qui bufferise ; False (défaut) : SSE pour le progress.
    mcp.run(transport="streamable-http", json_response=True)
```

Test avec curl (appel `tools/list`) :

```bash
curl -s -X POST http://127.0.0.1:8000/mcp \
  -H 'Content-Type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | head -c 800
```

⚠️ `mcp run` en HTTP sans authentification = OK en labo sur `127.0.0.1`,
**jamais** exposé tel quel sur le réseau (voir §18 et §40).

## 34. Écrire un client Python : le `Client` unifié

Le même paquet `mcp` fournit le client (SDK v2) :

```python
"""client_test.py — dialogue avec le serveur doc-rag, stdio ou HTTP."""
import asyncio
from mcp import Client

async def main() -> None:
    # HTTP distant :
    target = "http://127.0.0.1:8000/mcp"
    # ... ou stdio local : target = ["uv", "run", "server.py"]  (📌 forme exacte
    # du lancement stdio : à vérifier dans la doc du SDK v2 le jour J)

    async with Client(target) as client:
        tools = await client.list_tools()
        print("Tools :", [t.name for t in tools])

        res = await client.call_tool("search_docs", {"query": "bypass", "limit": 2})
        print(res.structured_content or res.content)

        # Lire une resource :
        data = await client.read_resource("doc://corpus/index")
        print(data[:300])

asyncio.run(main())
```

💡 Le client in-memory (`Client(mcp)` avec l'objet serveur direct, sans
transport) est idéal pour les **tests unitaires** : aucun processus, aucun
port (voir §42).

## 35. Brancher le serveur dans Claude Code

`claude mcp add` (CLI) ou fichier `.mcp.json` à la racine du projet :

```bash
# stdio : Claude Code lance le serveur à chaque session
claude mcp add doc-rag -- uv run --directory ~/mcp/doc-server server.py

# HTTP distant (Streamable HTTP) :
claude mcp add --transport http doc-rag http://127.0.0.1:8000/mcp
```

Équivalent en `.mcp.json` (versionnable en équipe) :

```json
{
  "mcpServers": {
    "doc-rag": {
      "type": "stdio",
      "command": "uv",
      "args": ["run", "--directory", "/home/zelef/mcp/doc-server", "server.py"],
      "env": {}
    },
    "doc-rag-distant": {
      "type": "http",
      "url": "https://mcp.interne.example.com/mcp",
      "headers": { "Authorization": "Bearer ${MCP_API_TOKEN}" }
    }
  }
}
```

Points clés :
- `${VAR}` : expansion de variable d'environnement — **jamais** de secret
  en dur dans un fichier commité.
- Scopes : `.mcp.json` (projet, partageable) vs `~/.claude.json`
  (perso, tous projets) vs managed-mcp.json (admin, verrouillé).
- Permissions fines dans `.claude/settings.json` :
  `"allow": ["mcp__doc-rag__search_docs"]` ou `"deny": ["mcp__*"]`.
  Convention de nommage : `mcp__<serveur>__<outil>`.
- `enableAllProjectMcpServers: true` = approbation auto des serveurs du
  projet (pratique, mais 🔒 à n'activer que sur dépôts de confiance).

## 36. Brancher dans Claude Desktop / VS Code / Cursor

