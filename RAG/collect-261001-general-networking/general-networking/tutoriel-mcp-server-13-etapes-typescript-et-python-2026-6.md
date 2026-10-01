---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-6
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: ["2025-06-18", "2025-11-25"]
keywords: ["agent", "aws", "bedrock", "chatgpt", "deepseek", "gemini", "mcp", "memory", "mistral"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [549, 573]
sha256: 46585fff658c1794ef6d0706974720a093b2894109509e45181f837c033c0c1f
---

# Forcer une version et lancer l'inspector en mode HTTP

Les incontournables maintenus par la communauté `modelcontextprotocol` incluent `filesystem`, `github`, `slack`, `postgres`, `sqlite`, `memory`, `fetch`, `brave-search`, `google-drive` et `puppeteer`, tous disponibles via `npx -y @modelcontextprotocol/server-<nom>`. Du côté des éditeurs, le **Microsoft Learn MCP Server** illustre la maturité atteinte par les serveurs officiels : lancé le 12 juin 2025 avec trois outils (search, fetch, code sample search), il a reçu son outil de récupération de documentation le 6 août 2025 – convertissant chaque page en Markdown pour les modèles –, puis son outil de recherche d’échantillons de code le 24 septembre 2025, avant de passer en disponibilité générale le 7 novembre 2025 avec un point de terminaison compatible OpenAI dédié à la recherche approfondie.

### MCP fonctionne-t-il avec GPT-5, Gemini et DeepSeek ?

Oui. OpenAI a ajouté le support MCP côté client dans ChatGPT début 2026, Google Cloud Vertex AI Agent Builder l’intègre depuis mars 2026, et DeepSeek ainsi que Mistral l’exposent via leurs propres clients. Côté serveur, peu importe le modèle consommateur : le protocole est neutre.

### Comment sécuriser un serveur MCP distant ?

Trois couches indispensables : TLS 1.3, OAuth 2.1 avec PKCE et resource indicator (RFC 8707), et rate limiting strict par session (par exemple 60 appels/minute). Ajoutez une allowlist de scopes, validez systématiquement le claim `aud` du token et journalisez chaque appel outil pour l’audit.

### Peut-on monétiser un serveur MCP ?

Oui, via OAuth scopes payants et metered billing. Plusieurs éditeurs (Linear, Notion, Asana) facturent désormais l’accès MCP à leurs API comme un add-on. Les marketplaces Cloudflare, AWS Bedrock et Azure AI Studio proposent depuis Q1 2026 un catalogue commercial de serveurs MCP certifiés.

### Quelle est la différence entre stdio et Streamable HTTP ?

Le transport stdio est local (sous-processus), sans chiffrement, idéal pour un serveur exécuté sur la machine de l’utilisateur (filesystem, accès à un outil CLI). Streamable HTTP fonctionne à distance, avec TLS et OAuth, pour multi-utilisateur ou cloud. L’ancien transport SSE est déprécié depuis la spec 2025-06-18.

### Comment publier un serveur MCP sur le registre officiel ?

Le registre communautaire est hébergé sur GitHub sous `modelcontextprotocol/servers`. Ouvrez une pull request avec votre serveur dans `src/`, en respectant la convention de nommage `server-<nom>`, une suite de tests et un README. Pour un serveur commercial, privilégiez la publication directe sur npm/PyPI et un référencement dans un marketplace cloud.

### Related Coverage

*Sources officielles : Spécification MCP 2025-11-25, TypeScript SDK, Python SDK, Registre officiel de serveurs, Annonce Anthropic, MCP Inspector. Article mis à jour le 23 avril 2026.*
