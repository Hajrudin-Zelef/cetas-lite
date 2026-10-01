---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-14
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-07-28"]
keywords: ["mcp", "agent", "agents", "claude", "model context protocol"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [2229, 2350]
sha256: 5929609de8c42e233b2abf98331bf0b31ae47d0f266a9bfd3c391b913382c7e1
---

# MCP, LSP, code-server, websearch — Guide pratique

1. `print()` sur stdout d'un serveur MCP stdio → protocole corrompu.
2. Tuto MCP 2024 copié tel quel → `initialize`/HTTP+SSE obsolètes.
3. `npx -y @modelcontextprotocol/server-postgres` installé sans vérifier →
   paquet archivé/non maintenu.
4. Serveur MCP lancé en **root** → `read_file` peut lire `/etc/shadow`.
5. Description d'outil floue → le modèle ne l'appelle jamais (ou mal).
6. Pas de `max_chars`/`limit` → un doc de 50k tokens sature le contexte.
7. Secret en dur dans `.mcp.json` commité → rotation immédiate requise.
8. `latest` en prod (Docker ou npm) → rug pull silencieux possible.
9. LSP : dossier ouvert au lieu du projet → diagnostics absurdes.
10. Python : serveur LSP sur le mauvais venv → imports « introuvables ».
11. C/C++ sans `compile_commands.json` → clangd devine (mal).
12. code-server exposé sans TLS → shell complet offert au premier venu.
13. code-server : websocket non proxifié → terminal intégré mort.
14. Websearch : 3 appels « basic » au lieu d'1 « advanced » → plus cher.
15. Pas de cache sur `web_search` → l'agent paie 10× la même question.
16. Rate limit ignoré → 429 en pleine démo client.
17. Source web unique citée comme vérité → croiser, toujours.
18. 30 serveurs MCP installés « au cas où » → contexte gonflé + surface
    d'attaque. 5 à 8 suffisent.

## 113. Cas pratiques (scénarios Zelef)

**Cas 1 — « L'agent connaît mes guides. »**
Branche `doc-rag` (Partie II) en stdio dans Claude Code. Demande :
« D'après mes guides, quelle section pour un bypass permanent sur Easy UPS
3S ? » → l'agent appelle `search_docs` puis `read_doc`. Succès = réponse
avec citation du fichier et de la section.

**Cas 2 — « Diagnostic assisté sur site. »**
Depuis ton téléphone (code-server via WireGuard), tu décris le symptôme ;
l'agent lance le prompt `diagnostic_ups` : il fouille ton corpus, propose
3 causes + tests de confirmation. Tu valides, il génère la fiche
d'intervention.

**Cas 3 — « Veille CVE automatisée. »**
Cron hebdo : agent avec Tavily/Brave MCP → cherche « CVE onduleur Eaton
Schneider 2026 », croise avec ton parc (SQLite via server-sqlite), produit
un rapport Markdown avec criticité et action. Coût : quelques centimes.

**Cas 4 — « Code sans hallucination. »**
Sur ton script `collect_pw.py`, l'agent utilise `mcp-language-server`
(pyright) : `definition`/`references` pour cartographier avant de refactorer.
Chaque patch est re-vérifié par `diagnostics` → boucle jusqu'au vert.

**Cas 5 — « Environnement de dev reproductible. »**
Image Docker : code-server + Python + uv + tes serveurs MCP + SearxNG en
fallback. Un `docker compose up` = ton poste de dev complet, sauvegardé,
versionné, accessible en navigateur depuis l'astreinte.

## 114. Glossaire

- **Agent** : programme qui boucle « raisonner → agir (tools) → observer ».
- **Annotations (MCP)** : `readOnlyHint`, `destructiveHint`, `idempotentHint`,
  `openWorldHint` — métadonnées de sécurité d'un tool.
- **ATPA** : empoisonnement via les **sorties** d'outils (CyberArk 2026).
- **Client MCP** : brique du host qui parle à un serveur (1:1).
- **code-server** : VS Code dans le navigateur (projet Coder).
- **Completion (LSP)** : `textDocument/completion` — suggestions contextuelles.
- **Content block** : unité typée de réponse MCP (`text`, `image`, `resource`...).
- **Deny-by-default** : tout interdire sauf allowlist explicite.
- **Diagnostics (LSP)** : erreurs/avertissements publiés par le serveur.
- **DNS rebinding** : attaque qui fait appeler ton localhost depuis le web.
- **Elicitation → MRTR** : dialogue serveur→humain, refondu en 2026-07-28.
- **Exfiltration** : sortie non autorisée de données via un tool.
- **FastMCP** : ancien nom de la classe serveur du SDK Python v1 (→ `MCPServer`).
- **Full-Schema Poisoning** : injection dans tout le schéma d'outil (2026).
- **Host** : application hôte (Claude Code, Cursor...) qui embarque les clients.
- **HTTP+SSE** : ancien transport distant MCP, déprécié.
- **Inspector** : UI officielle de debug des serveurs MCP.
- **isError** : flag d'erreur métier dans un résultat `tools/call`.
- **JSON-RPC 2.0** : protocole d'appel distant sous MCP et LSP.
- **Least privilege** : moindre privilège — le principe cardinal.
- **LSP** : Language Server Protocol — intelligence sémantique des éditeurs.
- **MCP** : Model Context Protocol — standard outils/données pour agents.
- **mcp-remote** : proxy stdio→OAuth/legacy pour serveurs distants.
- **MCPServer** : classe serveur du SDK Python v2.
- **MRTR** : Multi Round-Trip Requests — pattern de dialogue 2026-07-28.
- **Open VSX** : marketplace d'extensions utilisé par code-server.
- **Path traversal** : `../../` pour sortir du dossier autorisé.
- **Prompt (MCP)** : template de workflow paramétrable exposé par un serveur.
- **Rate limit** : quota d'appels/temps d'une API.
- **Resource (MCP)** : donnée adressable par URI, attachée au contexte.
- **Roots** : déclaration de dossiers d'accès — déprécié en 2026-07-28.
- **Rug pull** : serveur qui change son catalogue après audit.
- **Sampling** : le serveur redemandait au LLM — déprécié en 2026-07-28.
- **SearxNG** : métamoteur auto-hébergeable.
- **Server (MCP)** : processus qui expose tools/resources/prompts.
- **SERP** : page de résultats d'un moteur de recherche.
- **SSE** : Server-Sent Events — streaming HTTP unidirectionnel.
- **stdio** : transport local MCP/LSP via stdin/stdout du processus.
- **Streamable HTTP** : transport distant MCP standard (1 endpoint).
- **StructuredContent** : payload validé optionnel d'un résultat d'outil.
- **Tool (MCP)** : action appelable par le modèle (nom + schéma + handler).
- **Tool poisoning** : instructions cachées dans les métadonnées d'outils.
- **WireGuard/Tailscale** : VPN modernes — la bonne façon d'exposer code-server.

## 115. Quiz — 10 questions

1. Quelles sont les **trois primitives** MCP et qui décide de les utiliser
   (modèle, humain, host) ?
2. Pourquoi est-il **interdit** d'écrire des logs sur `stdout` en transport
   stdio ?
3. HTTP+SSE vs Streamable HTTP : quelle différence d'architecture justifie la
   dépréciation du premier ?
4. Que deviennent `sampling`, `roots` et le handshake `initialize` dans la
   spec 2026-07-28 ?
5. C'est quoi le **tool poisoning**, et pourquoi la description d'un tool est
   une surface d'attaque ?
6. Quelles **annotations** mettrais-tu sur un tool `delete_backup`, et
   pourquoi ?
7. En LSP, quelle méthode appelle un agent pour « trouver tous les appelants
   d'une fonction » ? Et pour « où est définie cette variable » ?
8. Pourquoi un serveur LSP Python donne des imports « introuvables » alors
   que le code tourne ?
9. code-server : pourquoi faut-il proxifier le **websocket** derrière nginx,
   et que se passe-t-il sinon ?
10. Ton agent fait 3 recherches « basic » à 8 $/1k au lieu d'1 « advanced » :
    quel est le vrai problème économique, au-delà du prix facial ?

## 116. Quiz — réponses

