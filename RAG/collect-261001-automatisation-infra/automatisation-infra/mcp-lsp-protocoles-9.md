---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-9
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS", "Anthropic", "Microsoft"]
dates: ["2025-11-25"]
keywords: ["mcp", "agent", "agentic", "aws", "benchmark", "claude", "disclosure", "incident", "sandbox"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1384, 1536]
sha256: f4efb91ea3aba549709ce837d93f995009484c51e5bc3edb43d9fa6e1c2fb8f8
---

# MCP, LSP, code-server, websearch — Guide pratique

Disclosure du 15 avril 2026 : faille systémique dans l'exécution STDIO des
SDK officiels (Python, TypeScript, Java, Rust — ~150 M de téléchargements
cumulés, ~200 000 instances estimées vulnérables). En substance : la façon
dont les SDK lancent les processus serveurs permettait l'exécution de
commande dans certains cas. Réponse d'Anthropic : comportement « by design »,
STDIO = défaut sûr, **la sanitization est la responsabilité du développeur**
— pas de patch protocole.

Leçon pour toi : **ne compte pas sur le SDK pour te protéger**. Valide les
entrées, scope les chemins, sandboxe les processus — c'est ton code, ta
responsabilité.

## 65. Le ver SANDWORM_MODE (février 2026)

Un ver npm (**19 paquets typosquattés**) incluait un module « McpInject » qui
plantait un **faux serveur MCP** (`~/.dev-utils/`) avec des tools aux noms
rassurants (`index_project`, `lint_check`) dont les descriptions exfiltraient
clés SSH, credentials AWS, tokens npm et `.env` — ciblant Claude Code,
Claude Desktop, Cursor, VS Code, Windsurf.

Défenses : ne jamais `npm install` à l'aveugle (vérifie le nom exact du
paquet !), audite `~/.claude.json` / `.mcp.json` / `claude_desktop_config.json`
régulièrement (tout serveur inconnu = incident), permissions d'écriture
restrictives sur ces fichiers.

## 66. Moindre privilège : mise en pratique

- **Filesystem** : args = dossiers explicites, jamais `/`. Le serveur
  officiel le fait ; ton serveur maison doit faire pareil (`is_relative_to`).
- **BDD** : rôle SQL `mcp_readonly` (`GRANT SELECT`), pas de `postgres`/root.
- **Réseau** : si ton serveur n'a pas besoin d'Internet, bloque l'egress
  (firewall, namespace réseau, `openWorldHint: false` honnête).
- **Tokens** : scopes minimaux (repo:read plutôt que repo:write si lecture
  seule), rotation, jamais dans git.
- **Processus** : utilisateur dédié `mcp`, pas root, pas ton user avec ses
  clés SSH.
- **Outils destructeurs** : `destructiveHint: true` + confirmation host
  obligatoire. Sépare `query_read` et `query_write` en deux tools.

## 67. Audit d'un serveur avant adoption (checklist 15 points)

1. [ ] Code source disponible et lu (ou éditeur de confiance) ?
2. [ ] `tools/list` : chaque description lue, pas d'impératif suspect ?
3. [ ] Noms de paramètres/valeurs par défaut passés au crible (§59) ?
4. [ ] Annotations honnêtes (`destructiveHint` là où il faut) ?
5. [ ] Périmètre filesystem/BDD/réseau minimal et explicite ?
6. [ ] Pas de `latest` : version épinglée + hash vérifié ?
7. [ ] Secrets en variables d'environnement, jamais en dur ?
8. [ ] stdio : stdout propre, pas d'exec de commandes construites par
   concaténation (injection) ?
9. [ ] HTTP : HTTPS + auth, anti-rebinding si localhost ?
10. [ ] Logs : pas de secrets en clair, pas d'args sensibles loggés ?
11. [ ] Dépendances : `pip audit` / `npm audit` propres ?
12. [ ] L'auteur est-il identifiable (repo, historique) ou anonyme ?
13. [ ] Quelle donnée transite ? (RGPD, secret industriel ?)
14. [ ] Plan de mise à jour : qui surveille les CVE du serveur ?
15. [ ] Testé dans l'Inspector avant branchement à l'agent ?

Score : < 12/15 → bac à sable obligatoire (§68) ou refus.

## 68. Isolation : sandbox, conteneurs, réseau

Niveaux croissants :
1. **Utilisateur dédié** (`useradd -r mcp`) — 2 minutes, déjà énorme.
2. **Conteneur read-only** (§39) — isole FS et réseau (`--network none`
   si pas besoin d'Internet).
3. **VM / microVM** — pour les serveurs non audités mais nécessaires.
4. **Egress filtering** — proxy sortant avec allowlist de domaines.
5. **Secrets via vault** — le serveur récupère son token à durée de vie
   courte, jamais stocké.

💡 Pour ton `doc-rag` : niveau 1 + volume `:ro` suffit — il ne fait que
lire des `.md` locaux.

## 69. Permissions côté Claude Code : le dernier rempart

```json
{"permissions": {
  "allow": ["mcp__doc-rag__search_docs", "mcp__doc-rag__read_doc",
            "mcp__github__get_issue"],
  "deny":  ["mcp__*__*delete*", "mcp__docker__*exec*"],
  "defaultMode": "acceptEdits" }}
```

- La granularité va jusqu'au tool (`mcp__<serveur>__<outil>`).
- `deny` prime sur `allow` : mets les motifs dangereux en deny.
- En équipe : `managed-mcp.json` (verrouillé admin) + `allowedMcpServers` /
  `deniedMcpServers` par nom, commande ou URL (§35).

## 70. Gouvernance : qui surveille l'écosystème

- **Agentic AI Foundation** (Linux Foundation, déc. 2025) — gouvernance
  ouverte du protocole.
- **OWASP AISVS** — chapitre C10 dédié à la sécurité MCP (référence d'audit).
- **MCPTox** — benchmark de 1 312 cas de test malveillants : à passer sur
  tes serveurs maison.
- Spec 2025-11-25 : groupes de travail + « SDK tiering » — mais **toujours
  pas de vérification d'intégrité obligatoire** des serveurs (avril 2026).
  Traduction : la confiance reste à construire côté opérateur. C'est toi.

## 71. Sécurité MCP : les règles d'or de Zelef

1. Un serveur MCP = du code avec tes droits. Traite-le comme tel.
2. Lis les descriptions d'outils comme tu lis un contrat.
3. Moindre privilège partout : fichiers, BDD, réseau, tokens.
4. Épingle les versions, surveille les diffs de catalogue.
5. Jamais de secret vers un serveur non audité.
6. `doc-rag` (ton serveur) : user dédié, volume read-only, pas d'egress.
7. En cas de doute : bac à sable d'abord, production ensuite.

# PARTIE V — LSP : L'INTELLIGENCE DES ÉDITEURS

## 72. LSP : le problème qu'il résout

Sans LSP : chaque éditeur (VS Code, Neovim, Emacs...) réimplémente
l'analyse sémantique de chaque langage — M éditeurs × N langages. Avec LSP
(Microsoft, 2016, ouvert) : **chaque langage fournit un serveur**, chaque
éditeur parle le protocole — M + N. C'est exactement la même idée que MCP,
dix ans plus tôt, appliquée au code.

Concrètement, le LSP donne à ton éditeur : complétion intelligente,
« aller à la définition », « trouver les références », survol documenté,
diagnostics en temps réel (erreurs soulignées), renommage sûr, formatage,
actions de refactoring. Sans lui, un éditeur n'est qu'un coloriage
syntaxique amélioré.

## 73. Architecture client/serveur

```
┌──────────────┐   JSON-RPC 2.0    ┌────────────────────┐
│ ÉDITEUR      │   sur stdio       │ SERVEUR DE LANGAGE│
│ (client LSP) │◄─────────────────►│ (ex : gopls)       │
│              │                   │                    │
│ - envoie     │  didOpen/didChange│ - parse le projet  │
│   les events │  completion,      │ - connaît les types│
│ - affiche    │  definition,      │ - répond           │
│   les répons.│  diagnostics...   │                    │
└──────────────┘                   └────────────────────┘
```

- Le **client** (intégré à l'éditeur) envoie les événements
  (`textDocument/didOpen`, `didChange`...) et les requêtes.
- Le **serveur** (un binaire par langage, ex : `rust-analyzer`) maintient
  le modèle sémantique du projet et répond.
- Transport historique : **stdio** avec framing `Content-Length` (headers
  HTTP-like, pas du JSON par ligne comme MCP — nuance à connaître si tu
  debug au niveau paquet).

💡 Comme MCP, le serveur est « bête » au sens noble : il ne décide pas,
il **répond à des questions précises** sur le code.

## 74. Le handshake : initialize / initialized / shutdown

