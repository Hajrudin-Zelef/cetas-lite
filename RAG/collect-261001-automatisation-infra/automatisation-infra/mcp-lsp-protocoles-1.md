---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-1
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: ["2025-11-25", "2026-07-28"]
keywords: ["mcp", "agent", "agentic", "agents", "claude", "gemini"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1, 135]
sha256: 52f951f430863b12e907c3ae73b2f854825808abc8245d32a4b5cda04f2ce873
---

# MCP, LSP, code-server, websearch — Guide pratique

> **Public :** Zelef, chef de service systèmes & énergies, sysadmin, construit son RAG perso.
> **Objectif :** comprendre et utiliser les 4 briques d'outillage des agents IA modernes :
> MCP (brancher des outils aux modèles), LSP (l'intelligence des éditeurs),
> code-server (VS Code dans le navigateur), websearch (donner le web aux agents).
> **État des lieux :** fin septembre 2026. Les versions et prix bougent vite :
> ce qui est marqué « à vérifier » doit être re-contrôlé le jour où tu l'utilises.

---

## 0. Mode d'emploi de ce guide

- Chaque partie est autonome : tu peux lire dans l'ordre ou piquer ce qu'il te faut.
- Le code Python/Node est réel et testé contre les SDK actuels (SDK Python `mcp`
  1.27+ / v2 bêta ; SDK TypeScript 1.29+ / v2). Les commandes d'installation sont
  celles des dépôts officiels.
- Les prix des API sont ceux constatés fin septembre 2026 : à vérifier avant
  de sortir la carte bleue.
- Convention : `⚠️` = piège classique, `💡` = astuce, `🔒` = sécurité,
  `📌` = à vérifier / incertain.

---

# PARTIE I — MCP : LE PROTOCOLE

## 1. Pourquoi MCP existe (genèse)

Avant MCP (novembre 2024, lancé par Anthropic), chaque assistant IA qui voulait
appeler un outil externe bricolait son intégration : un connecteur maison pour
Slack, un autre pour GitHub, un autre pour ta base Postgres. N assistants × M
outils = N×M intégrations à maintenir. MCP inverse l'équation : **un seul
protocole standard**, chaque outil l'implémente une fois, tous les assistants
compatibles peuvent l'utiliser.

L'analogie officielle : MCP est « l'USB-C des applications IA ». Comme l'USB-C
a standardisé la prise, MCP standardise la façon dont un modèle :
- découvre ce qu'un outil sait faire (liste des tools),
- lit des données (resources),
- reçoit des instructions pré-packagées (prompts).

Fin 2025, la gouvernance est passée sous la **Agentic AI Foundation**
(fondation Linux, créée décembre 2025) : le protocole n'appartient plus à un
seul éditeur. Les clients compatibles en septembre 2026 : Claude Code,
Claude Desktop, VS Code, Cursor, Gemini CLI, Codex, Windsurf, Cline, Continue.

## 2. Le principe en une image

```
┌──────────────┐      MCP (JSON-RPC)      ┌───────────────────┐
│  HOST        │◄────────────────────────►│  SERVEUR MCP      │
│  (Claude     │   1. tools/list          │  "filesystem"     │
│   Code,      │   2. tools/call          │  expose :         │
│   Cursor...) │                          │   - read_file     │
│              │      ┌──────────────┐    │   - write_file    │
│  ┌────────┐  │      │ CLIENT MCP   │    │   - list_dir      │
│  │ Modèle │◄─┼─────►│ (1 par       │    └───────────────────┘
│  │ (LLM)  │  │      │  serveur)    │
│  └────────┘  │      └──────────────┘    ┌───────────────────┐
│              │                         │  SERVEUR MCP      │
└──────────────┘                         │  "postgres"       │
                                         │  expose :         │
                                         │   - query (read)  │
                                         └───────────────────┘
```

Le **host** est l'application (ton IDE, ton agent). Le **client MCP** est la
brique logicielle, à l'intérieur du host, qui parle à **un** serveur. Le
**serveur MCP** expose des capacités. Le modèle ne parle jamais directement
au serveur : il demande au host d'appeler un tool, le host route vers le
bon client, le client appelle le serveur, la réponse remonte au modèle.

💡 Retiens : **1 host → N clients → N serveurs**. Chaque serveur est isolé ;
si le serveur Postgres plante, le serveur filesystem continue.

## 3. Spec 2026-07-28 : ce qui a changé (à connaître absolument)

La version courante de la spec fin septembre 2026 est **2026-07-28**
(remplace 2025-11-25). Ce n'est pas une mise à jour cosmétique, c'est une
refonte. Points clés :

| Sujet | Avant (2025-11-25) | Maintenant (2026-07-28) |
|---|---|---|
| État de connexion | `initialize` + `notifications/initialized`, session via `Mcp-Session-Id` | **Protocole sans état** : chaque requête porte sa version et les capacités client dans `_meta` ; découverte via `server/discover` |
| Transport distant | HTTP + SSE (2 endpoints + session) | **Streamable HTTP** (1 endpoint) ; HTTP+SSE **déprécié** |
| Roots | Négociés à l'init | **Déprécié** → passer les chemins en paramètres d'outils / URIs de resources |
| Sampling (le serveur redemande au LLM) | Requêtes serveur→client | **Déprécié** → pattern MRTR (voir §20) ou appel direct de l'API du LLM |
| Logging | `logging/setLevel`, notifications | **Déprécié** → stderr ou OpenTelemetry |
| Elicitation (dialogue interactif) | `elicitation/create` serveur→client | Remplacé par MRTR : `resultType: "input_required"` |
| Souscriptions resources | `resources/subscribe` + SSE | Fusionné en un flux `subscriptions/listen` |
| `ping` | Requête dédiée | **Supprimé** |

⚠️ Conséquence pratique : les tutos datant de 2024/2025 qui montrent
`initialize`, `sampling/createMessage` ou le transport HTTP+SSE à deux
endpoints sont **obsolètes**. Si tu copies un exemple, vérifie sa date.

## 4. Architecture détaillée : host, client, serveur

- **Host** : l'application que l'utilisateur voit. Ex : Claude Code (terminal),
  Claude Desktop, VS Code + extension, ton propre agent Python. C'est lui qui
  gère les permissions (« autoriser l'appel de `write_file` ? »), l'affichage,
  et la boucle agentique.
- **Client** : objet logiciel créé par le host, **un par serveur**. Il gère
  le transport (spawn du processus stdio, ou HTTP), la négociation de version,
  et expose au host une API simple : `list_tools()`, `call_tool()`,
  `read_resource()`, `get_prompt()`.
- **Serveur** : processus indépendant (binaire, script Python, conteneur) qui
  **déclare** ses capacités et **exécute** les appels. Il ne voit jamais le
  modèle ; il ne voit que des requêtes JSON-RPC.

Le serveur est volontairement « bête » : toute l'intelligence (quel tool
appeler, avec quels arguments) est côté modèle/host. Le serveur valide,
exécute, renvoie.

## 5. Le cycle de vie d'une connexion (spec 2026-07-28)

1. **Découverte** : le client appelle `server/discover` → le serveur annonce
   son identité, les versions de protocole supportées, ses capacités
   (tools ? resources ? prompts ?).
2. **Négociation** : chaque requête suivante embarque dans `_meta` :
   `io.modelcontextprotocol/protocolVersion` + `clientInfo` + `clientCapabilities`.
   Pas de handshake persistant : deux requêtes successives peuvent atterrir
   sur deux instances différentes derrière un load balancer.
3. **Utilisation** : `tools/list`, `tools/call`, `resources/read`, `prompts/get`...
4. **Notifications** : messages sans réponse attendue (`notifications/cancelled`
   pour annuler un appel en cours, `notifications/tools/list_changed` quand
   le catalogue change).
5. **Fin** : en stdio, le client ferme le stdin du serveur → le serveur doit
   se terminer proprement. En Streamable HTTP, rien à fermer : pas d'état.

## 6. JSON-RPC 2.0 sur le fil : à quoi ça ressemble vraiment

MCP = JSON-RPC 2.0 + conventions. Exemple d'appel d'outil, en stdio
(une ligne = un message JSON complet, **sans retour à la ligne interne**) :

