---
id: collect-261001-general-networking/general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste-5
title: "Installation à froid (cache vide), mesurée sur Apple M3 Pro"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Lambda"]
dates: []
keywords: ["aws", "benchmark", "benchmarks", "claude"]
source: docs/RAG/collect-261001-general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste.md
source_anchor: ""
source_lines: [274, 294]
sha256: 83ee67cbd447ef6f661d31e2d38d3da7d7ca14719c6a8ff12c131aac99c1dc9b
---

# Installation à froid (cache vide), mesurée sur Apple M3 Pro

Sur le benchmark HTTP simple, oui, environ **1,3× plus rapide que Deno 2.5**. Mais Deno reste meilleur sur les standards web (URL, Fetch, Web Crypto sans polyfill) et offre un permission model plus mature. Le bon choix dépend de votre besoin : Bun pour la vitesse maximale, Deno pour la sécurité et les standards Web stricts.

### Bun supporte-t-il les workers et le multi-thread ?

Oui. Bun 1.3 supporte `worker_threads` à 96 % de compatibilité Node.js, plus une API `Bun.spawn` native. Les SharedArrayBuffer fonctionnent à 100 %. Le module `cluster` est aussi disponible avec `SO_REUSEPORT` natif sous Linux, ce qui simplifie les déploiements multi-cœurs.

### Quel framework HTTP choisir avec Bun en 2026 ?

**Hono** pour les APIs simples et les Workers (le plus rapide), **Elysia** pour la pile « Bun-first » avec validation Zod intégrée, **Fastify 5** si vous migrez depuis Node.js et voulez garder votre code, et **Express 5** pour la rétrocompatibilité maximale. Ces quatre frameworks tournent à 100 % sur Bun en avril 2026.

### Bun fonctionne-t-il sur AWS Lambda en 2026 ?

Oui, depuis le **14 mars 2026**, AWS publie un layer Lambda Bun officiel pour les architectures ARM64 et x86_64. Le runtime gagne 12 à 30× sur le cold start mais nécessite encore un effort de configuration plus important que Node.js 22 (runtime managé natif). Pour les fonctions à fort trafic, le gain est immédiat. Pour les fonctions occasionnelles, restez sur Node.js managé.

### Combien coûte Bun en production ?

Bun est **100 % gratuit et open-source MIT**. Aucun coût de licence. Le coût réel est celui de l’hébergement (Vercel Pro 20 $/mois, Railway 5 $/mois, AWS Lambda à 0,20 $/M d’invocations, Scaleway Serverless free tier). Anthropic n’a pas annoncé de modèle commercial autour de Bun en avril 2026 ; le runtime est traité comme une infrastructure mutualisée pour l’écosystème Claude Code.

### Related Coverage

*Sources externes : Annonce officielle Bun rejoint Anthropic · Communiqué Anthropic du 3 décembre 2025 · Notes de release Node.js 22 LTS · Dépôt GitHub officiel de Bun · TechEmpower Round 23 benchmarks.*
