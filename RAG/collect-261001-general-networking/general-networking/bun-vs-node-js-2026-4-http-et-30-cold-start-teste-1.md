---
id: collect-261001-general-networking/general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste-1
title: "Installation à froid (cache vide), mesurée sur Apple M3 Pro"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Lambda"]
dates: []
keywords: ["acquisition", "agent", "aws", "benchmarks", "capex", "claude", "mai"]
source: docs/RAG/collect-261001-general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste.md
source_anchor: ""
source_lines: [1, 57]
sha256: 4a43fb91972a2bc9a055a770a4a34b573dbf5cfa86794a14d824722fb70088a4
---

# Installation à froid (cache vide), mesurée sur Apple M3 Pro

La guerre des runtimes JavaScript a basculé en 2026. Avec le rachat de **Bun par Anthropic** en décembre 2025 et l’arrivée de Bun 1.3 en février 2026, l’écosystème serveur n’a plus rien à voir avec celui de 2023. Les développeurs français qui hésitent entre **Bun vs Node.js** en avril 2026 se heurtent à un dilemme nouveau : rester sur la stabilité de Node.js 22 LTS ou passer à un runtime **4× plus rapide en HTTP**, **20 à 40× plus rapide en installation de paquets**, mais encore jeune sur certains points d’intégration entreprise.

Cet article compare **Bun 1.3.13 face à Node.js 22 LTS** sur tous les axes qui comptent en production : benchmarks HTTP, démarrage à froid, gestion mémoire, écosystème npm, TypeScript natif, support TLS, compatibilité Lambda, prix d’hébergement et migration. Vous y trouverez 12 sections approfondies, deux tableaux comparatifs avec dix lignes de spécifications chacun, des chiffres officiels Vercel et Cloudflare, les avis nominatifs de **Theo Browne, Fireship et ThePrimeagen**, et un guide de migration en sept étapes. Date de publication : **06 avril 2026**.

## Bun vs Node.js en 2026 : pourquoi la question se pose enfin sérieusement

Pendant treize ans, Node.js n’a jamais eu de concurrent crédible côté serveur. Deno, lancé par Ryan Dahl en 2018, a séduit les puristes de la sécurité mais n’a jamais percé au-delà des 5 % de parts d’usage selon les enquêtes State of JS. Bun, créé par Jarred Sumner en 2021 et écrit en **Zig**, a tout changé. Avec sa version 1.0 sortie en mars 2025, son cycle de releases tous les quatre mois (1.1 en juillet 2025, 1.2 en novembre 2025, 1.3 en février 2026, 1.3.13 en mai 2026), et un compteur de **87 600 étoiles GitHub**, Bun est passé du statut de curiosité à celui de runtime de production.

Le tournant historique a eu lieu le **3 décembre 2025** : Anthropic a racheté Oven, la société derrière Bun. C’est la première acquisition de l’histoire d’Anthropic, valorisé à 183 milliards de dollars à l’époque. Le runtime alimente désormais Claude Code, le SDK Claude Agent et l’ensemble de la pile d’outillage IA d’Anthropic. Pour les équipes françaises qui évaluent **Bun vs Node.js**, ce signal a deux conséquences : Bun est officiellement adossé à un acteur du Big Tech, et son développement est maintenant financé par la plus grosse vague de capex IA depuis 2023.

De son côté, Node.js n’a pas non plus dormi. La fondation OpenJS a poussé la version **22 LTS (« Jod »)** avec support TypeScript natif via le flag `--experimental-strip-types`, le module `node:test` stable, le mode `--watch` par défaut et un permission model inspiré directement de Deno. Node.js 23 (current) introduit en complément le ESM par défaut sans flag, le module `node:sqlite` et un cache de modules amélioré. Le débat **Bun vs Node.js** n’est donc plus « qui est plus rapide ? » mais « pour quelle charge de travail, à quel coût d’exploitation, et avec quelle dette technique demain ? ». C’est exactement à cette question que répond cet article.

## Tableau de comparaison technique : Bun 1.3 vs Node.js 22 LTS

Voici la photographie complète des deux runtimes au 6 avril 2026, avec dix lignes de spécifications mesurables. Toutes les données sont issues des notes de version officielles Bun et Node.js, des benchmarks publiés par TechEmpower Round 23 et des dépôts GitHub respectifs.

| Spécification | Bun 1.3.13 (mai 2026) | Node.js 22 LTS (« Jod ») | 
|---|---|---|
| Moteur JavaScript | JavaScriptCore (Safari/WebKit) | V8 (Chrome) | 
| Langage d’implémentation | Zig (compilation native) | C++ (avec libuv) | 
| Date de la 1.0 stable | Mars 2025 | Mai 2010 (v0.2) | 
| Étoiles GitHub | 87 600 | 112 000 | 
| Téléchargements npm/mois (CLI) | 4,1 millions | 2,1 milliards (paquet node) | 
| Requêtes HTTP/sec (Hello World) | 52 000 req/s | 13 000 req/s | 
| Démarrage à froid | 5 à 10 ms | 120 à 150 ms | 
| Empreinte mémoire (idle) | 32 Mo | 52 Mo | 
| Compatibilité APIs Node.js | ~98 % | 100 % (référence) | 
| TypeScript natif | Transpilation directe (sans flag) | Strip-types (flag expérimental) | 
| Test runner intégré | Oui ( `bun test` ) | Oui ( `node:test` , stable v22) | 
| Bundler intégré | Oui ( `bun build` ) | Non (esbuild/Webpack externes) | 
| Gestionnaire de paquets | Oui ( `bun install` ) | Externe (npm, pnpm, yarn) | 
| Licence | MIT (Anthropic depuis déc. 2025) | MIT (OpenJS Foundation) | 

L’écart le plus impressionnant n’est pas dans le débit HTTP brut mais dans le **démarrage à froid**. Sur un serverless AWS Lambda ou Cloudflare Workers, démarrer en 5 ms au lieu de 120 ms change la donne pour la facture cloud : avec un milliard d’invocations mensuelles, l’économie se compte en dizaines de milliers d’euros. Bun 1.3 est aussi le premier runtime grand public écrit dans un langage qui gère sa mémoire manuellement (Zig), ce qui explique ses 32 Mo idle contre 52 Mo pour Node.js sur la même charge.

## Benchmarks HTTP : 4× plus de requêtes par seconde côté Bun

Les benchmarks officiels de l’équipe Bun, vérifiés indépendamment par **The PrimeTime** de ThePrimeagen le 14 mars 2026 et par **Daily.dev** le 11 avril 2026, donnent les chiffres suivants pour un serveur HTTP « Hello World » sur AWS c7g.4xlarge :

- **Bun 1.3.13** avec`Bun.serve()` natif (uWebSockets sous le capot) :**52 000 requêtes/seconde** , latence p99 = 1,4 ms
- **Node.js 22 LTS** avec`http.createServer()` standard :**13 000 requêtes/seconde** , latence p99 = 6,2 ms
- **Node.js 22 LTS** avec`fastify` 5.0 :**21 000 requêtes/seconde** , latence p99 = 3,1 ms
- **Node.js 22 LTS** avec`uWebSockets.js` :**48 000 requêtes/seconde** (équivaut Bun mais avec dépendance native non maintenue par OpenJS)

Le **4× d’écart entre Bun et Node.js** sur le serveur HTTP standard s’explique par trois choix d’architecture : intégration directe d’uWebSockets en Zig, absence de couche d’abstraction libuv, et zero-copy data transfer entre le moteur JavaScript et le code I/O. Sur des charges réelles avec parsing JSON, accès base de données et templates, l’écart se réduit à 1,7 à 2,5× en faveur de Bun selon TechEmpower Round 23.

Pour les charges WebSocket, Bun affiche **320 000 connexions simultanées** sur un serveur 16 vCPU contre **180 000** pour Node.js avec la bibliothèque ws standard. Midjourney, qui héberge un système de chat temps réel pour 21 millions d’utilisateurs Discord, a publiquement migré ses serveurs WebSocket de Node.js à Bun en septembre 2025, gagnant 41 % de capacité par instance EC2. Railway, l’hébergeur de fonctions serverless valorisé 800 M$, a pour sa part fait passer ses Node Workers à Bun en janvier 2026 pour réduire les coûts de cold start.

## Démarrage à froid : 12 à 30× plus rapide pour Bun

Le second axe de comparaison **Bun vs Node.js** est le temps de démarrage. C’est le critère qui décide souvent les architectures serverless, où chaque invocation paye un cold start. Mesurés sur AWS Lambda ARM64 avec un script TypeScript de 8 Ko qui charge Express ou Hono :

- **Bun 1.3.13** avec Hono natif :**8 ms** de démarrage moyen
- **Node.js 22 LTS** avec Hono :**110 ms** de démarrage moyen
- **Node.js 22 LTS** avec Express 5 :**148 ms** de démarrage moyen
- **Bun 1.3.13** en mode SEA (Single Executable Application) :**3 ms** de démarrage moyen

