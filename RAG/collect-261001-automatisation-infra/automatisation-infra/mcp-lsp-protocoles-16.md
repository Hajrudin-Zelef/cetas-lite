---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-16
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI", "SpaceX"]
dates: ["2025-11-25", "2026-07-28", "2026-08-22", "2026-09-02", "2026-09-14", "2026-09-19", "2026-09-22", "2026-09-25", "2026-09-27", "2026-09-29"]
keywords: ["mcp", "agent", "agentic", "agents", "claude", "lean", "research"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [2491, 2638]
sha256: e6c72f06677e642d1bdf4e976b4415bb080ae705b6702bedc9eb206480e1e203
---

# --- code-server ---
curl -fsSL https://code-server.dev/install.sh | sh
code-server --bind-addr 127.0.0.1:8080 ~/projets
docker run -d -p 127.0.0.1:8080:8080 -v "$HOME/projets:/home/coder/projets" \
  -e PASSWORD='...' codercom/code-server:latest   # 📌 nom d'image à vérifier

# --- websearch (Tavily, ex. §108) ---
export TAVILY_API_KEY='...'   # jamais en dur, jamais commité
python websearch_tool.py
```

## 122. Checklist de mise en production du serveur doc-rag

- [ ] `pytest` vert, dont le test anti path-traversal (§42)
- [ ] Testé dans l'Inspector : les 3 tools + 2 resources + 2 prompts répondent
- [ ] Descriptions relues : précises, avec exemple, sans promesse excessive
- [ ] Bornes en place : `limit` ≤ 20, `max_chars` ≤ 20000, timeout handlers
- [ ] Tourne sous l'utilisateur dédié `mcp`, volume corpus en `:ro`
- [ ] Streamable HTTP derrière nginx TLS (§40), auth par clé/OAuth (§18)
- [ ] Anti-rebinding activé si bind local + client navigateur (§19)
- [ ] Logs JSON sur stderr → Loki, sans secrets (§23)
- [ ] Version épinglée, pas de `latest`, changelog tenu (§41)
- [ ] Surveillance du diff `tools/list` (détection rug pull, §61)
- [ ] Permissions Claude Code : allowlist `mcp__doc-rag__*` en lecture (§69)
- [ ] Budget websearch configuré si l'agent chaîne doc-rag + recherche (§109)
- [ ] Documenté pour l'équipe : README avec les 3 tools et 2 exemples d'usage

*Fin du guide — bon bricolage, et que tes agents restent sous ton contrôle. 🔧*

## 123. À venir — annonces vérifiées au 27/09/2026

> Méthode : recherche web effectuée le 27/09/2026. Seules les
> annonces officielles ou vérifiables figurent ici ; les rumeurs sont
> étiquetées **RUMEUR non confirmée**. Si rien d'annoncé pour un
> périmètre, c'est écrit noir sur blanc.

### 123.1 MCP — la roadmap officielle (22/08/2026)

Les mainteneurs du cœur MCP **David Soria Parra et Den Delimarsky** ont
publié la roadmap officielle le **22/08/2026** (« The New MCP
Roadmap »). Cinq priorités pour la prochaine révision de la spec et
au-delà — **sans dates ni numéros de version annoncés** :

1. **Primitives de messagerie agentique** (agent-to-agent) — faire
   communiquer les agents entre eux via le protocole, pas seulement
   agent→serveur.
2. **Unification et durcissement du transport HTTP-natif** — finir la
   convergence vers Streamable HTTP (HTTP+SSE est déprécié depuis
   2026-07-28), scalabilité horizontale sans état.
3. **Identité d'agent et sécurité entreprise** — prouver *quelle*
   charge de travail appelle ton serveur MCP (sans stocker de secrets
   longue durée), audit trails structurés, SSO, RBAC, patterns
   gateway/proxy. ⚠️ Ne couvre pas l'auth *derrière* le serveur (vers
   ta BDD/ton SI) : deux problèmes d'auth distincts.
4. **Primitives améliorées** — stabilisation de Tasks (expérimental
   depuis 2025), workflow primitives (state machines, checkpoints),
   streaming compressé.
5. **Meilleure expérience développeur des SDK** — qualité des SDK
   officiels (TypeScript, Python) et communautaires.

Gouvernance : MCP est sous l'**Agentic AI Foundation** (Linux
Foundation), avec des Working Groups (Transports, Auth, Registry…)
et des SEPs (propositions d'évolution). Source : *CData*, « The
Official MCP Roadmap, Explained for Enterprise Teams », 14/09/2026
(décortique l'annonce du 22/08).

### 123.2 MCP — le stable reste 2026-07-28 (rappel d'impact)

Au **25/09/2026**, la spec stable reste **2026-07-28** (TypeScript
SDK v2 l'implémente ; la ligne v1.x maintenance vise 2025-11-25).
Rappel des ruptures à anticiper pour ton serveur doc-rag quand tu
monteras de version :

- **Protocole sans état** : handshake `initialize` et `Mcp-Session-Id`
  supprimés — déployable derrière un simple load balancer / en
  serverless. Nouveau RPC obligatoire **`server/discover`**.
- **MRTR** (multi-round-trip requests) remplace les appels initiés par
  le serveur (sampling/élicitation/roots) : le serveur répond
  `resultType: "input_required"`, le client réessaie avec les réponses.
- **Tasks** devient l'extension officielle `io.modelcontextprotocol/
  tasks` (sémantique de retry/timeout encore en stabilisation —
  retour d'expérience production bienvenu par les mainteneurs).
- **Auth** : Dynamic Client Registration (RFC 7591) **déprécié** au
  profit des **CIMD** (Client ID Metadata Documents) ; validation
  `iss` RFC 9207 ; credentials par émetteur.
- **Dépréciés** (cycle de 12 mois) : Roots, Sampling, Logging,
  transport HTTP+SSE → migrer vers Streamable HTTP.
- Ops : métadonnées de cache obligatoires sur les listes, headers
  `Mcp-Method`/`Mcp-Name`, conventions OpenTelemetry, ordre
  déterministe des tools (ami du prompt caching).
- Sources : *ai-xblog.com*, « MCP 2026: Latest Spec & Security »
  (25/09/2026) ; synthèse *palaia* (GitHub, 22/09/2026).

**Non traité par la roadmap** (constat tiers, à garder en tête) :
efficacité du contexte (overhead des descriptions de tools), batching,
modèle de sécurité fondamental contre l'injection de prompt. Le
recommandé reste : peu de tools, descriptions resserrées, exécution
via CLI quand possible. Source : *wbingli/ai-agent-research*
(rapport roadmap, 25/09/2026).

### 123.3 LSP — rien d'officialisé au 27/09/2026

En septembre 2026, **LSP 3.18** reste à la fois « la dernière version »
et « en développement » sur le site de spec de Microsoft : aucune
nouvelle version officialisée. Côté implémentations, l'activité est du
côté des serveurs (ex. IntelliJ 2026.1 élargit le support LSP :
formatting, code lens, rename) plutôt que du protocole lui-même.
**Rien d'officialisé au 27/09/2026** pour une évolution du protocole.

### 123.4 code-server — suivi mensuel, pas de roadmap feature

- **v4.138.0 (19/09/2026)** = VS Code 1.138.0 ; v4.137.0 (11/09) ;
  v4.136.2 (08/09). Le projet suit le rythme mensuel de VS Code,
  changelog tenu sur GitHub. **Aucune roadmap de fonctionnalités
  publiée** — rien d'officialisé au 27/09/2026.
- En revanche, l'éditeur **Coder** (même société) pousse fort sur
  l'agentique auto-hébergé : **Coder Agent Relay** annoncé le
  **02/09/2026** (GlobeNewswire, avec SpaceXAI comme partenaire de
  lancement) — exécution des agents cloud (Cursor Cloud Agents) dans
  des workspaces Coder sur ton infra ; **Coder Agents** (agent natif,
  agnostique de modèle) en bêta. Pertinent si tu veux un jour exposer
  ton doc-rag + agents à l'équipe sans sortir du réseau interne.

### 123.5 APIs de recherche web — pas d'annonce majeure en septembre

- **Tavily, Exa, Firecrawl, Kagi** : aucune annonce officielle de
  fonctionnalité retrouvée pour septembre 2026. L'activité est côté
  écosystème (ex. *mcp-omnisearch* 0.1.0 ajoute Tavily Research en
  polling asynchrone + modes crawl/map).
- À noter (déjà sorti, pas « à venir ») : **Brave** a lancé début 2026
  une **AI Search API** dédiée — endpoint séparé de la Web Search API,
  contenu pré-traité et structuré pour ingestion LLM/agents (RAG,
  assistants). Si tu dois choisir un provider pour chaîner doc-rag +
  recherche web avec un budget tokens serré, c'est l'option « lean »
  documentée ; Tavily reste l'option « snippets riches ».
- Source : *webpronews.com* (fév. 2026, reprise de l'annonce Brave).

### 123.6 Calendrier — OpenAI DevDay (29/09/2026)

**DevDay OpenAI, 29/09/2026, San Francisco** : l'IA agentique est
l'axe majeur annoncé (Agents API en bêta publique depuis début
septembre, Agents SDK, infra managée, **MCP au programme des
sujets**). Prochain point de contrôle naturel pour cette section :
une annonce MCP-côté-OpenAI (ou une évolution Agents API) y est
plausible — à vérifier début octobre.

### 123.7 Tableau récapitulatif (statut au 27/09/2026)

