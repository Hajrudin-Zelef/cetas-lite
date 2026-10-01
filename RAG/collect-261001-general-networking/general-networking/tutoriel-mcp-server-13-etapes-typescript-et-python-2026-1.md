---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-1
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "Google", "Microsoft", "OpenAI"]
dates: ["2025-11-25", "2026-07-28"]
keywords: ["agents", "aws", "claude", "deepseek", "gemini", "mai", "mcp", "model context protocol", "tool calling"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: be6bc1660c60e89c700a52da8ea6f56320ae3f2eae5be65b0dc3d0a9b4fd2faa
---

# Forcer une version et lancer l'inspector en mode HTTP

Le **Model Context Protocol** (MCP) est devenu, en moins de deux ans, le standard de facto pour connecter les modèles d’intelligence artificielle à des outils, fichiers et bases de données. Depuis la publication initiale par Anthropic le 5 novembre 2024 jusqu’à la toute dernière spécification stable **2026-07-28** – publiée fin juillet 2026 comme cinquième révision majeure du protocole –, l’écosystème a explosé : plus de 2 000 serveurs MCP publics sont désormais recensés sur npm, Microsoft a même fait passer son propre Microsoft Learn MCP Server en disponibilité générale le 7 novembre 2025, et les clients Claude Desktop, Cursor, Windsurf, VS Code et Cline prennent tous en charge le protocole nativement. Ce tutoriel détaille la construction d’un **serveur MCP** complet, en TypeScript puis en Python, avec authentification OAuth 2.1, transports stdio et Streamable HTTP, et déploiement en production. Les commandes, extraits de code et configurations présentés fonctionnent avec le SDK TypeScript officiel `@modelcontextprotocol/sdk` v1.29.0 et le package Python `mcp` (FastMCP), les quatre SDK Tier 1 (TypeScript, Python, Go, C#) ayant tous livré le support de la spec 2026-07-28 dès le 28 juillet 2026 – contenu à jour au 19 août 2026.

À la fin des treize étapes, vous disposerez d’un serveur MCP réutilisable, exposant des *tools*, *resources* et *prompts*, testé avec **MCP Inspector**, connecté à Claude Desktop et Cursor, empaqueté pour npm et PyPI, déployé sur Cloudflare Workers, et prêt à intégrer des agents LangGraph ou CrewAI. Les pièges classiques (désynchronisation du handshake JSON-RPC, fuites de jetons OAuth, schémas Zod incohérents) sont traités en détail dans la section dédiée.

## Qu’est-ce que Model Context Protocol (MCP) en 2026

Le Model Context Protocol est un **protocole ouvert JSON-RPC 2.0** qui standardise la manière dont les applications LLM (hosts) communiquent avec des sources de données et des outils externes via des serveurs dédiés. Son inspiration directe est le Language Server Protocol (LSP) de Microsoft, qui a unifié les éditeurs et les analyseurs statiques depuis 2016. De la même façon, MCP évite aux développeurs d’écrire N × M intégrations propriétaires entre chaque modèle (Claude, GPT-5, Gemini, DeepSeek) et chaque système (GitHub, Slack, Postgres, Google Drive). Un seul serveur MCP est consommable par tous les clients compatibles, sans modification.

La spécification stable active au 19 août 2026 est **2026-07-28**, cinquième révision majeure du protocole. Son *release candidate* a été verrouillé le 21 mai 2026, laissant dix semaines de validation aux SDK Tier 1 avant publication fin juillet 2026 ; elle rend le cœur du protocole davantage stateless et instaure une fenêtre de dépréciation de douze mois pour les anciennes versions à compter de juillet 2026. Les quatre SDK Tier 1 (TypeScript, Python, Go, C#) ont tous livré ce support dès le 28 juillet 2026, et le SDK Rust a ajouté un support bêta de cette même spec le même mois, en tant que première implémentation hors Tier 1. La version précédente, **2025-11-25** – quatrième révision majeure publiée en novembre 2025 –, avait introduit la toute première politique de dépréciation formelle, ainsi que le support d’OpenID Connect Discovery, les métadonnées d’icônes pour tools/resources/prompts, le consentement incrémental de scopes, l’élicitation en mode URL, le *sampling tool calling*, les OAuth Client ID metadata documents et un support expérimental des *Tasks*. Anthropic a transféré la gouvernance de MCP à la Linux Foundation en 2025, ce qui a ouvert le protocole à la contribution de Microsoft, Google, AWS et OpenAI, via des Spec Enhancement Proposals (SEPs) et des Working Groups publics.

Techniquement, un serveur MCP expose trois primitives : des **tools** (fonctions exécutables, comme « envoyer un email »), des **resources** (données en lecture, comme un fichier ou une requête SQL) et des **prompts** (gabarits réutilisables pour les interactions LLM). Chacune est décrite par un schéma JSON validé côté serveur, ce qui permet au modèle d’appeler l’outil en toute sécurité, avec typage fort. Les transports supportés sont `stdio` (processus local), `SSE` (Server-Sent Events, déprécié au profit de Streamable HTTP), `Streamable HTTP` (transport web moderne unifié) et `WebSocket` (bidirectionnel, utilisé pour le sampling).

## Architecture MCP : Hosts, Clients et Serveurs

L’architecture MCP est stateful et suit un modèle trois rôles clairement délimités. Le **host** est l’application qui héberge le LLM (Claude Desktop, Cursor, un script LangGraph). Il instancie un ou plusieurs **clients**, un par serveur distant, qui maintiennent chacun une connexion JSON-RPC persistante. Chaque **serveur** MCP expose ses capacités via des méthodes comme `tools/list`, `tools/call`, `resources/read` ou `prompts/get`. Le host orchestre, le client dialogue, le serveur répond.

Le handshake d’initialisation est critique. À la connexion, le client envoie `initialize` avec la version de protocole supportée et ses capabilities (e.g. `roots`, `sampling`). Le serveur répond avec la version négociée, son nom, sa version et ses capabilities (`tools`, `resources`, `prompts`, `logging`). Le client confirme avec `notifications/initialized`, puis la session démarre. Tout appel avant ce triptyque échoue avec `-32002 (Server not initialized)`.

La table ci-dessous résume les rôles et flux du protocole dans la version 2025-11-25 :

| Composant | Rôle | Méthodes principales | Transport typique | 
|---|---|---|---|
| Host (Claude Desktop, Cursor) | Exécute le LLM et consomme les capacités | – | gère les clients | 
| Client (par serveur) | Maintient la session JSON-RPC | `initialize` ,`tools/call` | stdio, Streamable HTTP | 
| Serveur MCP | Expose tools, resources, prompts | `tools/list` ,`resources/read` | stdio, HTTP, WebSocket | 
| Transport stdio | Communication locale par flux | JSON-RPC sur stdin/stdout | fork de processus | 
| Transport Streamable HTTP | Communication distante moderne | POST /mcp + SSE | HTTPS + OAuth 2.1 | 

## Prérequis du tutoriel : versions, outils et comptes

Avant d’écrire la moindre ligne, vérifiez l’environnement. Un serveur MCP moderne exige des runtimes récents pour supporter `fetch` natif, les modules ES et les décorateurs asynchrones – d’autant plus depuis que les quatre SDK Tier 1 (TypeScript, Python, Go, C#) ont basculé sur le support de la spécification 2026-07-28 le 28 juillet 2026. Les versions minimales et recommandées au 19 août 2026 sont :

| Outil | Version minimale | Version recommandée | Commande de vérification | 
|---|---|---|---|
| Node.js | 18.0 | 22 LTS | `node --version` | 
| npm | 9.0 | 10.x | `npm --version` | 
| Python | 3.10 | 3.12 | `python3 --version` | 
| uv (gestionnaire Python) | 0.4 | 0.5+ | `uv --version` | 
| TypeScript | 5.2 | 5.6+ | `tsc --version` | 
| @modelcontextprotocol/sdk | 1.20 | 1.29.0 | `npm list @modelcontextprotocol/sdk` | 
| mcp (PyPI) | 1.12 | 1.15+ | `uv pip show mcp` | 
| Claude Desktop | 0.8 | 1.2+ | Menu → À propos | 

Créez également un compte **GitHub** pour tester le serveur MCP GitHub officiel en fin de tutoriel, installez `git` 2.40+, et prévoyez un dossier de travail propre : `mkdir -p ~/mcp-lab && cd ~/mcp-lab`. Si vous comptez déployer sur Cloudflare, créez un compte gratuit et installez `wrangler` 4.x via `npm install -g wrangler`. Pour déboguer, vous utiliserez **MCP Inspector**, livré via `npx @modelcontextprotocol/inspector`.

## Étape 1 – Initialiser un projet MCP TypeScript avec le SDK 1.29

