---
id: collect-261001-general-networking/general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste-2
title: "Installation à froid (cache vide), mesurée sur Apple M3 Pro"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Lambda"]
dates: []
keywords: ["aws", "benchmark", "benchmarks", "claude", "sandbox"]
source: docs/RAG/collect-261001-general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste.md
source_anchor: ""
source_lines: [58, 130]
sha256: d6c6f515b49824b995bd1184884b4eca6ea71f0ff2ced5f6f677e1840e33c152
---

# Installation à froid (cache vide), mesurée sur Apple M3 Pro

Cette accélération provient principalement de trois optimisations Bun : la pré-compilation TypeScript en cache, le linking statique des dépendances Zig dans le binaire, et l’absence de phase d’initialisation V8 (qui prend à elle seule 80 ms sur un Node.js classique). Pour une API qui sert dix millions de requêtes serverless par jour, passer de 120 ms à 8 ms représente une économie d’**environ 18 000 €/mois** sur la facture AWS Lambda Provisioned Concurrency selon les calculs de Vercel publiés en mars 2026.

Côté **Single Executable Application**, c’est le terrain où Bun a creusé l’écart. Anthropic a basé Claude Code sur la fonction `bun build --compile` qui génère un binaire de 96 Mo embarquant runtime, code et assets. Le binaire démarre en 3 ms et se distribue sans Node.js installé sur la machine cible. Node.js 22 propose aussi des SEA via `--build-snapshot`, mais le binaire pèse 165 Mo et démarre en 95 ms.

## Installation de paquets : bun install écrase npm sur tous les indicateurs

Le gestionnaire de paquets intégré `bun install` est la fonctionnalité la plus citée dans les retours d’expérience français. Sur un projet Next.js standard avec 1 247 dépendances et un cache vide :

```
# Installation à froid (cache vide), mesurée sur Apple M3 Pro
bun install   # 4,1 secondes
pnpm install  # 18,7 secondes
yarn install  # 24,3 secondes
npm install   # 87,2 secondes
# Installation tiède (lockfile présent)
bun install   # 0,3 seconde
pnpm install  # 1,9 seconde
yarn install  # 4,5 secondes
npm install   # 12,1 secondes
```
Le ratio **20 à 40× en faveur de Bun** contre npm vient de quatre choix techniques : un cache global sous forme de hardlinks (à la pnpm), un parser de manifeste écrit en Zig, des résolutions DNS parallélisées via Linux io_uring, et un format binaire de lockfile (`bun.lockb`) qui se lit en 50 ms au lieu de 800 ms pour `package-lock.json`. Pour une équipe française qui exécute des CI/CD sur GitHub Actions ou GitLab CI, passer de npm à Bun divise par 8 à 12 le temps moyen d’un job d’installation.

Bun 1.3 a en outre introduit le mode `--frozen-lockfile` par défaut sur la commande `bun ci`, et la prise en charge native du protocole **JSR** (le registre rival de npm créé par Deno Land). Côté Node.js, le projet officiel Corepack stabilise pnpm et yarn comme alternatives, mais aucune n’atteint les performances brutes de Bun. **Theo Browne**, fondateur de t3.gg, déclarait sur son podcast du 8 février 2026 : « *J’ai abandonné pnpm pour Bun il y a six mois. Je n’y reviendrai jamais. Le gain de productivité au quotidien est immense, et le lockfile binaire règle 80 % des merge conflicts* ».

## TypeScript natif : transpilation Bun vs strip-types Node.js 22

L’un des arguments les plus puissants en faveur de Bun en 2025 était la prise en charge TypeScript sans configuration. Node.js a partiellement rattrapé ce retard avec la version 22 LTS, mais la différence reste visible. Voici les trois modes disponibles en avril 2026 :

- **Bun 1.3.13** : exécution directe de fichiers`.ts` et`.tsx` sans flag, sans`tsconfig.json` obligatoire, transpilation en cache disque, support TSX/JSX natif
- **Node.js 22 LTS** : flag`--experimental-strip-types` qui retire les annotations de type sans transpiler, donc pas de support des enums, namespaces, ni des décorateurs legacy
- **Node.js 23** (current) : flag`--experimental-transform-types` qui ajoute la transpilation complète mais reste expérimental

Concrètement, un fichier `server.ts` qui utilise des enums TypeScript ou des décorateurs Nest.js plante immédiatement avec Node.js 22 strip-types. Bun 1.3 l’exécute sans broncher. Pour les codebases Nest.js, TypeORM ou Prisma qui s’appuient sur les décorateurs, la différence est **structurelle**. À l’inverse, ni Bun ni Node.js ne font de type-checking au runtime ; il faut toujours invoquer `tsc --noEmit` ou `bun tsc` pour valider les types en CI. Pour aller plus loin sur l’écosystème JavaScript serveur full-stack, consultez notre tutoriel Bun.js 1.3 avec Elysia qui couvre la prise en main bout en bout.

## Compatibilité APIs Node.js : 98 % atteinte par Bun en 2026

Le grand pari de Bun depuis 2023 est de servir de **drop-in replacement** à Node.js. La doctrine est simple : on doit pouvoir remplacer la commande `node` par `bun` sans toucher au code. Au 6 avril 2026, le tableau de compatibilité publié par l’équipe Bun affiche les pourcentages suivants par module :

| Module Node.js | Compatibilité Bun 1.3.13 | Notes | 
|---|---|---|
| `fs` /`fs/promises` | 100 % | Y compris `watch` et streams | 
| `http` /`https` | 99 % | `http2` stable depuis 1.2 | 
| `net` /`tls` | 98 % | SNI multi-cert OK depuis 1.3 | 
| `crypto` | 97 % | Quelques subtilités sur `scrypt` | 
| `cluster` | 95 % | Mode `SO_REUSEPORT` natif | 
| `worker_threads` | 96 % | SharedArrayBuffer 100 % | 
| `child_process` | 99 % | `spawn` et`exec` 100 % | 
| `perf_hooks` | 95 % | `monitorEventLoopDelay` OK | 
| `dgram` (UDP) | 92 % | Améliorations prévues 1.4 | 
| `vm` | 78 % | Le plus gros écart restant | 

Le module qui pose le plus de problèmes est `vm`, utilisé par les sandbox de code et certains parsers Markdown comme `marked`. Pour le reste, sur les frameworks majeurs, l’équipe Bun fait tourner les suites de tests d’**Express, Fastify, Hono, NestJS, Koa, Hapi et Polka** à chaque release. Express 5 et Fastify 5 passent à 100 %. NestJS 11 passe à 99,2 %. Le seul cadre majeur qui présente encore des incompatibilités est **SvelteKit avec adapter-node**, mais l’adapter `@sveltejs/adapter-bun` publié en novembre 2025 résout le problème. Pour une stack Next.js, suivez notre tutoriel Next.js avec Turbopack.

## Performance comparée sur charges réelles : API REST, WebSocket, batch

Les benchmarks Hello World ne reflètent pas le quotidien d’une équipe SaaS. Pour rendre la comparaison **Bun vs Node.js** opérationnelle, voici trois charges de travail typiques mesurées sur la même machine (AWS c7g.4xlarge, 16 vCPU ARM, 32 Go RAM, Ubuntu 24.04 LTS) :

### API REST avec PostgreSQL et Redis

Une API qui valide un JWT, lit un objet depuis Redis (cache), tombe sur un miss, va chercher les données dans PostgreSQL via le client `postgres.js` et renvoie un JSON de 4 Ko. **Bun 1.3.13 avec Elysia** tient **34 200 req/s** avec une latence p99 de 4,3 ms. **Node.js 22 LTS avec Fastify 5** tient **19 800 req/s** avec une latence p99 de 7,8 ms. L’écart, autour de **1,72×**, est inférieur au benchmark Hello World mais tangible : à charge équivalente, vous économisez près de 40 % d’instances EC2.

### WebSocket à 100 000 clients connectés

Sur 100 000 connexions WebSocket simultanées, broadcast d’un message JSON de 200 octets toutes les secondes : **Bun 1.3.13** consomme **2,1 Go de RAM** et utilise **38 % CPU**. **Node.js 22 LTS** avec `ws` consomme **3,8 Go de RAM** et utilise **71 % CPU**. Le ratio est de **1,8× à 1,9×** en faveur de Bun pour la mémoire et le CPU. Sur un cas d’usage chat ou notifications, ce gain se traduit directement par une instance moins onéreuse.

### Batch de calcul : parsing CSV de 5 Go

Parsing d’un fichier CSV de 5 Go (12 millions de lignes), agrégation des sommes par colonne, écriture du résultat sur disque. **Bun 1.3.13** termine en **47 secondes**, **Node.js 22 LTS** en **96 secondes**, ratio **2,04×**. L’écart vient principalement du parser CSV en Zig et du `Bun.file()` qui utilise `mmap` sous Linux pour éviter les copies mémoire. Pour les pipelines ETL Node.js historiques, il s’agit du gain qui motive le plus de migrations en 2026.

## Avis d’experts : Theo, Fireship, ThePrimeagen au crible

Le verdict des trois personnalités les plus influentes de l’écosystème JavaScript en 2026 dessine une cartographie claire de l’usage de Bun en production.

