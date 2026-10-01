---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-6
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-07-28"]
keywords: ["mcp", "claude", "memory"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [879, 1076]
sha256: 28707a29bc1539f3d349d5075008bc3e363399ef0ff1dcfb88ef35f968dc13ac
---

# MCP, LSP, code-server, websearch — Guide pratique

- **Claude Desktop** : `~/.config/Claude/claude_desktop_config.json`
  (Linux) avec une clé `mcpServers` au même format (command/args/env).
  Redémarrer l'app après modification.
- **VS Code** : `.vscode/mcp.json` avec clé `servers` :
  ```json
  {"servers": {"doc-rag": {"command": "uv",
    "args": ["run", "--directory", "/home/zelef/mcp/doc-server", "server.py"]}}}
  ```
- **Cursor / Windsurf / Cline** : panneau MCP des settings, même format
  `mcpServers`. La plupart acceptent l'import d'un `.mcp.json`.

⚠️ Chaque host a ses nuances (nom de clé `servers` vs `mcpServers`,
gestion des headers). En cas de doute : la doc du host fait foi, pas ce guide.

## 37. Serveur Node/TypeScript minimal (quand c'est pertinent)

Si un jour tu dois greffer MCP sur une app Node existante
(SDK `@modelcontextprotocol/sdk` 1.29+ ; v2 : paquets éclatés) :

```typescript
// server.ts — SDK v1 stable (1.29.x)
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { z } from "zod";

const server = new McpServer({ name: "doc-rag-ts", version: "1.0.0" });

server.registerTool(
  "search_docs",
  {
    title: "Rechercher dans la documentation",
    description: "Recherche plein-texte dans le corpus documentaire local.",
    inputSchema: {                       // ← un "shape" Zod, pas z.object()
      query: z.string().describe("Mots-clés de recherche"),
      limit: z.number().int().max(20).optional(),
    },
  },
  async ({ query, limit }) => ({
    content: [{ type: "text", text: `Résultats pour ${query} (top ${limit ?? 5})...` }],
  })
);

const transport = new StdioServerTransport();
await server.connect(transport);
```

```bash
npm init -y && npm pkg set type=module
npm i @modelcontextprotocol/sdk@^1.29.0 zod
npx tsc --init --module nodenext --target es2022 --strict
npx tsc && node dist/server.js   # ou : npx tsx server.ts
```

📌 SDK v2 (sorti en 2026) : imports éclatés —
`import { McpServer } from "@modelcontextprotocol/server"`,
`import { StdioServerTransport } from "@modelcontextprotocol/server/stdio"`,
`zod/v4`, et `registerTool` (remplace l'ancien `tool()`).
Vérifie la version installée avant de copier l'un ou l'autre.

## 38. Servir le serveur TS en Streamable HTTP (Express)

```typescript
// http.ts — SDK v2 (paquets @modelcontextprotocol/*)
import express from "express";
import { McpServer } from "@modelcontextprotocol/server";
import { StreamableHTTPServerTransport } from "@modelcontextprotocol/server/streamableHttp";

const app = express();
app.use(express.json());

app.post("/mcp", async (req, res) => {
  const server = new McpServer({ name: "doc-rag-ts", version: "1.0.0" });
  // ... registerTool(...) comme au §37 ...
  const transport = new StreamableHTTPServerTransport({
    sessionIdGenerator: undefined,              // sans état (spec 2026-07-28)
    enableDnsRebindingProtection: true,         // 🔒 anti-rebinding
    allowedHosts: ["127.0.0.1", "mcp.interne.example.com"],
    allowedOrigins: ["https://ide.interne.example.com"],
  });
  await server.connect(transport);
  await transport.handleRequest(req, res, req.body);
});

app.listen(8000, "127.0.0.1");
```

```bash
npm i @modelcontextprotocol/server @modelcontextprotocol/node express
```

💡 Pattern « un serveur par requête » : en mode sans état, on peut créer le
`McpServer` à chaque POST — simple et compatible load balancer. Pour des
handlers lourds à initialiser, mutualise l'objet serveur (il est réutilisable
entre transports).

## 39. Dockeriser un serveur MCP (Python)

```dockerfile
# Dockerfile
FROM python:3.12-slim
RUN pip install --no-cache-dir uv
WORKDIR /app
COPY server.py .
# corpus monté en volume à l'exécution, pas dans l'image
RUN uv pip install --system "mcp[cli]"
EXPOSE 8000
CMD ["python", "server.py"]   # server.py finit par mcp.run(transport="streamable-http", ...)
```

```bash
docker build -t doc-rag-mcp:1.0 .
docker run -d --name doc-rag \
  -p 127.0.0.1:8000:8000 \
  -v ~/workspace/user/files:/data:ro \
  -e DOC_DIR=/data \
  --read-only --tmpfs /tmp \
  doc-rag-mcp:1.0
```

🔒 Durcissement : volume en **read-only** (`:ro`), filesystem racine en
lecture seule, pas de `latest` en prod (tag versionné), `DOC_DIR` en variable
d'environnement (adapte `server.py` : `Path(os.environ.get("DOC_DIR", ...))`).

## 40. Exposer derrière nginx (TLS + auth)

```nginx
# /etc/nginx/sites-available/mcp-doc-rag
server {
    listen 443 ssl;
    server_name mcp.interne.example.com;

    ssl_certificate     /etc/letsencrypt/live/mcp.interne.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/mcp.interne.example.com/privkey.pem;

    # 🔒 SSE / streaming : pas de buffering, timeouts longs
    location /mcp {
        proxy_pass http://127.0.0.1:8000;
        proxy_http_version 1.1;
        proxy_set_header X-Accel-Buffering no;
        proxy_read_timeout 300s;
        proxy_request_buffering off;

        # Auth simple par clé (ou déléguer à OAuth/authelia en amont)
        auth_request /auth;
    }
    location = /auth { proxy_pass http://127.0.0.1:9000/verify; }
}
```

💡 Le SDK conseille aussi d'envoyer des keep-alive SSE (`:`) côté serveur
pour traverser les timeouts des proxies. Teste toujours le streaming de
bout en bout, pas seulement en localhost.

## 41. Versioning et compatibilité protocole

- Annonce une `protocolVersion` réelle dans `server/discover`
  (ex : `2026-07-28`). N'invente pas de numéro.
- Les clients négocient : si ton serveur ne parle qu'une vieille version,
  un client récent peut refuser ou dégrader. **Suis les releases du SDK**,
  c'est lui qui absorbe les changements de spec.
- Versionne ton **catalogue d'outils** comme une API : renommer un tool ou
  changer un schéma = breaking change → version majeure du serveur,
  changelog, et si possible alias de l'ancien nom pendant une transition.
- 🔒 Le « rug pull » (§61) : un serveur qui change ses descriptions d'outils
  **sans** changer de version est un signal d'alerte en audit.

## 42. Tester : stratégie minimale mais sérieuse

```python
# test_server.py — pytest, client in-memory (aucun port, aucun processus)
import pytest
from mcp import Client
from server import mcp, DOC_DIR

@pytest.mark.asyncio
async def test_search_trouve_bypass():
    async with Client(mcp) as client:          # in-memory : direct, rapide
        res = await client.call_tool("search_docs", {"query": "bypass"})
        docs = [r["document"] for r in res.structured_content]
        assert any("onduleur" in d or "ups" in d for d in docs)

@pytest.mark.asyncio
async def test_traversal_bloquee():
    async with Client(mcp) as client:
        res = await client.call_tool("read_doc", {"name": "../../etc/passwd"})
        assert res.isError  # doit échouer proprement, pas crasher
```

```bash
uv add --dev pytest pytest-asyncio
uv run pytest -q
```

💡 Trois niveaux : unitaires in-memory (rapides, CI) → Inspector (visuel,
contrat) → test d'intégration avec le vrai client (Claude Code) une fois.

## 43. Erreurs fréquentes (Python) et remèdes

