---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-17
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: ["2026-07-28", "2026-08-22", "2026-09-02", "2026-09-19", "2026-09-25", "2026-09-29"]
keywords: ["mcp", "agent", "agents", "chatgpt", "claude", "copilot", "research"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [2639, 2680]
sha256: 590188bb4f849f6e00440802ffff9ae52c1c8895358785ca83e1afbaccb76526
---

# MCP, LSP, code-server, websearch — Guide pratique

| Périmètre | Annoncé, daté | En cours / bêta | Statut |
|---|---|---|---|
| Spec MCP | Roadmap officielle **22/08/2026** (5 priorités, sans dates) | Stable **2026-07-28** ; Tasks en extension | Prochaine révision : date inconnue |
| LSP | — | 3.18 en développement | **rien d'officialisé** |
| code-server | v4.138.0 (**19/09/2026**) | Suivi mensuel VS Code | **rien d'officialisé** (roadmap) |
| Coder (éditeur) | Agent Relay (**02/09/2026**) | Coder Agents (bêta) | — |
| Websearch APIs | — | Brave AI Search API (déjà lancée, début 2026) | **rien d'annoncé en sept. 2026** |
| Écosystème | DevDay OpenAI (**29/09/2026**) | Agents API (bêta publique) | À surveiller début 10/2026 |

> 🔁 Rituel : relire cette section **après le DevDay (début octobre
> 2026)** puis **début janvier 2027**, et déplacer les éléments sortis
> vers les sections concernées du guide. En particulier : toute
> nouvelle révision stable de la spec MCP impose de re-tester ton
> serveur doc-rag (handshake, `server/discover`, auth CIMD).

### 123.8 Compléments — signaux faibles de septembre 2026 (écosystème)

Quelques signaux qui ne sont pas des annonces officielles mais
méritent une ligne dans ton radar, car ils touchent directement ton
serveur doc-rag :

- **MCP Registry** : l'index central approche les **2 000 serveurs**
  référencés (sept. 2026). La découverte standardisée (`.well-known`,
  Server Cards v2.1) avance dans les Working Groups : publier ton
  doc-rag au Registry avec une carte propre deviendra le moyen normal
  d'être trouvé par des agents tiers — prévois la fiche (nom,
  description, version, auth) dès maintenant.
- **SDKs** : ~97 M de téléchargements mensuels cumulés pour les SDK
  Python + TypeScript ; SDKs communautaires (Java, C#, Go, Rust,
  Swift, Kotlin) actifs. Le chantier « qualité SDK » de la roadmap se
  voit déjà : si tu as bricolé des contournements (retry, timeouts),
  re-vérifie à chaque montée de version du SDK plutôt que de les
  garder ad vitam.
- **Inspector** : le travail continue (support des nouvelles révisions
  de spec) — reste ton banc d'essai avant chaque mise en production.
- **Concurrence des standards** : MCP a dépassé OpenAPI en étoiles
  GitHub (constat du podcast Latent Space avec les créateurs), mais
  OpenAPI garde un outillage plus mature. Pour ton usage interne,
  aucun risque : MCP reste le standard des hosts d'agents (Claude,
  ChatGPT, Copilot, Cursor, VS Code, Zed).
- Sources : synthèse *awesome-mcp-servers* (données écosystème,
  53 jours, re-vérifiée) ; *wbingli/ai-agent-research* (25/09/2026).
