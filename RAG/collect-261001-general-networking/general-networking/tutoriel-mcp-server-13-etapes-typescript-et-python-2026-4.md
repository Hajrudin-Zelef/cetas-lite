---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-4
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "aws", "bedrock", "claude", "copilot", "distribution", "mcp"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [336, 451]
sha256: 95c802580067d04af90538f60d14c293ad1886d0d815ee6708e07a87bfa1d966
---

# Forcer une version et lancer l'inspector en mode HTTP

L’écosystème d’éditeurs supportant MCP a doublé entre 2025 et 2026. Cursor a ajouté le support natif en v0.43, Windsurf dans sa version Cascade 1.0, et Microsoft a intégré MCP dans GitHub Copilot Agent Mode sur VS Code 1.99+. La configuration est similaire à Claude Desktop mais chaque éditeur utilise son propre fichier.

| Éditeur | Fichier de configuration | Mode agent | 
|---|---|---|
| Claude Desktop | `claude_desktop_config.json` | Chat + slash commands | 
| Cursor | `~/.cursor/mcp.json` ou`.cursor/mcp.json` (projet) | Composer Agent | 
| Windsurf | `~/.codeium/windsurf/mcp_config.json` | Cascade | 
| VS Code (Copilot Agent) | `.vscode/mcp.json` (projet) | Agent Mode | 
| Cline | `~/.config/cline/mcpServers.json` | Autonomous Coder | 

Pour Cursor, créez `.cursor/mcp.json` à la racine du projet avec la même structure que Claude Desktop. L’onglet *MCP* des préférences Cursor permet d’activer/désactiver chaque serveur et de voir les outils disponibles. Pour VS Code, le fichier `.vscode/mcp.json` supporte en plus des *inputs* sécurisés pour les secrets (API keys stockés dans VS Code Secret Storage, jamais en clair) :

```
{
  "inputs": [
    {
      "type": "promptString",
      "id": "github-token",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ],
  "servers": {
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": { "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github-token}" }
    }
  }
}
```
## Étape 10 – Packager et déployer en production

Pour distribuer votre serveur MCP TypeScript sur npm, préparez le `package.json` avec un `bin`, un `files` restreint et une condition de publication. Ajoutez un shebang `#!/usr/bin/env node` en tête du `dist/index.js` (le SDK propose l’helper `chmod` via un script postbuild).

```
{
  "name": "@votreorg/ti-demo-mcp",
  "version": "1.0.0",
  "type": "module",
  "bin": { "ti-demo-mcp": "dist/index.js" },
  "files": ["dist"],
  "scripts": {
    "build": "tsc && chmod +x dist/index.js",
    "prepublishOnly": "npm run build"
  },
  "dependencies": {
    "@modelcontextprotocol/sdk": "1.29.0",
    "zod": "3.24.0"
  }
}
```
Publiez avec `npm publish --access public`. Les utilisateurs l’installeront en un appel dans Claude Desktop : `"command": "npx", "args": ["-y", "@votreorg/ti-demo-mcp"]`. Pour Python, publiez sur PyPI via `uv build && uv publish`. Pour une distribution Docker, une image basée sur `node:22-alpine` suffit en moins de 100 Mo.

Pour un déploiement **Cloudflare Workers**, utilisez `wrangler` 4.x. Cloudflare a publié en 2025 un adaptateur officiel `@modelcontextprotocol/sdk-cloudflare` qui remplace Express par Workers runtime et gère l’OAuth 2.1 via Workers KV. Un déploiement sur AWS Bedrock MCP Gateway, Google Cloud Vertex AI Agent Builder ou Azure AI Studio est également possible depuis début 2026.

```
# Déploiement Cloudflare Workers
npm install -g wrangler@4
wrangler init mcp-worker
# Ajouter votre serveur MCP puis
wrangler deploy
```
## Étape 11 – Tester, monitorer et journaliser

Les tests unitaires d’un serveur MCP ciblent d’abord les handlers (logique métier des tools) et ensuite le handshake complet. Vitest est le framework recommandé pour TypeScript, `pytest` pour Python. Le SDK propose `InMemoryTransport` pour tester sans I/O réelle.

```
import { describe, it, expect } from "vitest";
import { Client } from "@modelcontextprotocol/sdk/client/index.js";
import { InMemoryTransport } from "@modelcontextprotocol/sdk/inMemory.js";
describe("ti-demo-server", () => {
  it("calcule la TVA à 20 %", async () => {
    const [ct, st] = InMemoryTransport.createLinkedPair();
    await server.connect(st);
    const client = new Client({ name: "test", version: "1.0.0" });
    await client.connect(ct);
    const res = await client.callTool({
      name: "calculer_tva",
      arguments: { montant_ht: 100, taux: "20" },
    });
    expect(res.content[0].text).toContain("120,00");
  });
});
```
En production, exposez des métriques Prometheus (`prom-client`) sur `/metrics` : nombre d’appels par tool, latence P50/P95/P99, taux d’erreur JSON-RPC, sessions actives. Un dashboard Grafana de base montre ces quatre séries et identifie immédiatement un outil lent ou une fuite de sessions. Le logging structuré (pino pour Node, structlog pour Python) vers un agrégateur comme Loki ou Datadog complète le dispositif.

## Étape 12 – Intégrer à LangGraph et CrewAI pour agents multi-LLM

Un serveur MCP peut être consommé par n’importe quel framework d’agents. LangGraph (LangChain) expose depuis 2025 un adaptateur `langchain-mcp-adapters` qui convertit les tools MCP en LangChain Tools ; CrewAI propose `crewai-tools-mcp`, et AutoGen ajoute un `MCPToolAdapter` natif. Ces adaptateurs gèrent le handshake, la découverte des tools et la conversion des schémas.

```
from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client
from langchain_mcp_adapters.tools import load_mcp_tools
from langgraph.prebuilt import create_react_agent
from langchain_anthropic import ChatAnthropic
async def run_agent():
    params = StdioServerParameters(command="node", args=["dist/index.js"])
    async with stdio_client(params) as (r, w):
        async with ClientSession(r, w) as session:
            await session.initialize()
            tools = await load_mcp_tools(session)
            agent = create_react_agent(
                ChatAnthropic(model="claude-opus-4-7"),
                tools,
            )
            out = await agent.ainvoke({"messages": "Calcule la TVA de 1250 € à 20 %."})
            print(out["messages"][-1].content)
```
L’agent ReAct interroge la liste des tools exposés par votre serveur, sélectionne `calculer_tva`, passe les arguments extraits de la phrase utilisateur, reçoit la réponse et produit une réponse finale en langage naturel. Tout cela sans écrire une seule ligne de code d’intégration spécifique au modèle.

## Étape 13 – Projet complet : serveur MCP Filesystem d’entreprise

Récapitulons avec un projet de bout en bout. Le serveur `ti-filesystem-mcp` expose deux tools (`lire_fichier`, `lister_dossier`), une resource templatée (`file://{path}`) et un prompt (`resume_code`). Il filtre les chemins via une allowlist pour éviter les escapes, et supporte stdio local ainsi que HTTP distant.

