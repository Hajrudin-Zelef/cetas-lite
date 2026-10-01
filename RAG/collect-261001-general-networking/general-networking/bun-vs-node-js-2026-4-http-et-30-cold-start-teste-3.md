---
id: collect-261001-general-networking/general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste-3
title: "Installation à froid (cache vide), mesurée sur Apple M3 Pro"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Lambda", "Mistral"]
dates: []
keywords: ["agent", "aws", "claude", "mistral"]
source: docs/RAG/collect-261001-general-networking/bun-vs-node-js-2026-4-http-et-30-cold-start-teste.md
source_anchor: ""
source_lines: [131, 185]
sha256: 7e298e139ce9f6ebb874d37ce4939589f5a3ad223f138e8c76c0085d00fae4b8
---

# Installation à froid (cache vide), mesurée sur Apple M3 Pro

**Fireship** (Jeff Delaney), 4,1 millions d’abonnés YouTube, a publié le 5 décembre 2025 la vidéo « *Anthropic just bought your favorite JS runtime* » dans laquelle il déclare : « *Bun n’est plus un runtime alternatif, c’est devenu l’infrastructure JavaScript d’Anthropic. Si vous écrivez du JavaScript pour la production, vous devriez l’utiliser au moins en local pour gagner du temps. La grande question est de savoir si Anthropic le rendra payant à terme.* ». Sa note finale : **9/10 pour le DX (developer experience), 7/10 pour l’écosystème**.

**ThePrimeagen** (Michael Paulson), ex-Netflix et streamer Twitch, a comparé Bun et Node.js en direct le 14 mars 2026 sur un projet Hono. Verdict : « *Le moment où j’ai compris que Bun était sérieux, c’est en exécutant un script TypeScript de 12 Ko sans avoir à toucher quoi que ce soit. Sur Node.js 22, je dois encore configurer un `--loader` ou utiliser `tsx`. Pour mes projets perso, je suis 100 % Bun. Pour le boulot, je reste sur Node.js LTS parce que mes ops ne veulent pas entendre parler de Zig.* ». Sa recommandation : adopter Bun en local, garder Node.js en production tant que l’équipe SRE n’a pas validé le runtime.

**Theo Browne**, fondateur de t3.gg et créateur de la stack T3 (Next.js + tRPC + Prisma), a publié sur X le 8 février 2026 un thread dans lequel il explique pourquoi t3.gg a basculé tous ses projets internes sur Bun en novembre 2025 : « *On a divisé par 9 le temps de CI sur GitHub Actions juste en remplaçant `npm ci` par `bun install --frozen-lockfile`. C’est le ROI le plus rapide que j’ai vu sur un changement d’outil.* ». Theo recommande Bun pour les startups qui démarrent en 2026, et déconseille la migration des codebases Node.js de plus de cinq ans tant que la suite de tests n’a pas été audité.

Côté français, **Matt Pocock** (TotalTypeScript) a déclaré en mars 2026 sur le podcast Codeur Mobile : « *Bun n’a pas tué Node.js, il l’a forcé à se moderniser. Sans Bun, on n’aurait jamais eu `node:test` ni `--watch` dans Node.js 22. Le gagnant final, c’est l’écosystème.* ». Une formulation diplomatique mais juste : la concurrence Bun vs Node.js a accouché en 2026 d’un Node.js plus rapide et plus complet qu’à n’importe quel moment de son histoire.

## Tableau des prix d’hébergement et coûts d’exploitation

Le coût d’exploitation est l’argument décisif pour beaucoup de directions techniques françaises. Voici la comparaison des principales offres d’hébergement adaptées à Bun et Node.js, avec les tarifs publiés en avril 2026 :

| Plateforme | Plan | Prix mensuel | Support Bun natif | Support Node.js | 
|---|---|---|---|---|
| Vercel Hobby | Hobby | 0 € | Oui (depuis 2025) | Oui (par défaut) | 
| Vercel Pro | Pro | 20 $/utilisateur | Oui | Oui | 
| Cloudflare Workers | Free | 0 € | Compatibilité partielle | Compat. via nodejs_compat | 
| Cloudflare Workers Paid | Standard | 5 $/mois | Compatibilité partielle | Compat. via nodejs_compat | 
| Railway | Hobby | 5 $/mois | Oui (Bun officiel) | Oui | 
| Fly.io | Pay-as-you-go | ~5 $/mois | Oui (image officielle) | Oui | 
| Render | Starter | 7 $/mois | Oui | Oui | 
| AWS Lambda | Pay-per-call | 0,20 $/M | Via custom runtime | Native | 
| Scaleway Serverless | Free Tier | 0 € | Via container | Native | 
| OVHcloud Functions | Pay-as-you-go | ~3 €/mois | Via container Bun | Native | 

Vercel et Railway sont aujourd’hui les deux seules plateformes qui proposent Bun comme runtime de premier rang, sans contournement par container ou custom runtime. Cloudflare Workers ne supporte pas encore l’intégralité du runtime Bun mais offre depuis octobre 2025 un compat layer qui couvre 80 % des APIs Node.js. Sur AWS Lambda, il faut passer par un layer Lambda Bun (publié officiellement par AWS en mars 2026) ou un container ARM64 ; la performance gagne 12 à 30× sur le cold start mais la complexité opérationnelle reste plus élevée que pour Node.js. Pour les organisations européennes soumises à des contraintes RGPD strictes, **Scaleway** et **OVHcloud** proposent des images Bun en version 1.3 sur leurs offres serverless souveraines.

## Cinq cas d’usage et recommandations Bun vs Node.js

Plutôt que de désigner un grand vainqueur, le bon angle est de matcher chaque cas d’usage avec le runtime adapté. Voici cinq scénarios couvrant la majorité des projets web français en 2026.

### Cas 1 : Startup SaaS qui démarre une nouvelle codebase

Recommandation : **Bun**. Aucun héritage à porter, équipe jeune, gain immédiat sur le DX, CI/CD divisée par 8, possibilité de packager en SEA pour les CLI internes. Choisir Hono ou Elysia comme framework HTTP, Drizzle ORM pour PostgreSQL, et le runtime `bun test` intégré. Mistral AI et Doctolib backend internal tools en 2026 ont fait ce choix.

### Cas 2 : E-commerce existant en Node.js 18 avec 200 000 LOC

Recommandation : **Node.js 22 LTS**. Migrer 200 000 lignes vers un runtime alternatif coûte 6 à 12 mois de travail pour un ROI incertain ; mieux vaut investir dans la mise à jour Node.js 18 → 22 (qui apporte déjà le node:test stable, le mode –watch et le strip-types) et profiter du nouvel HTTP keep-alive amélioré. Veepee, Vinted backend Node, et Le Bon Coin ont confirmé cette stratégie en Q1 2026.

### Cas 3 : API serverless à fort trafic (10 M+ requêtes/jour)

Recommandation : **Bun** sur Vercel Edge ou Cloudflare Workers via nodejs_compat. L’économie sur les cold starts est immédiate (12-30×), et la facture serverless chute de 40 à 60 %. Contentful a publié un cas d’étude le 22 mars 2026 montrant un passage de 47 000 $ à 18 000 $ par mois en migrant 250 fonctions Lambda vers Bun.

### Cas 4 : Outillage interne CLI et scripts d’automatisation

Recommandation : **Bun**. Le mode `bun build --compile` produit un binaire portable de 96 Mo qui démarre en 3 ms et se distribue sans Node.js sur la machine cible. C’est le choix d’Anthropic pour Claude Code, de t3.gg pour ses CLI internes et de Datadog pour son agent de migration de configurations.

### Cas 5 : Backend Nest.js entreprise avec décorateurs et microservices

Recommandation : **Node.js 22 LTS** pour le moment. Bien que NestJS 11 affiche 99,2 % de compatibilité avec Bun, certains plugins (Swagger, GraphQL Federation, Microservices Kafka) souffrent encore de bugs subtils. Si vous démarrez un projet Nest.js neuf en 2026, faites un POC sur Bun, mais ne migrez pas des microservices critiques tant que la version Bun 1.4 (prévue T3 2026) n’a pas livré la compat `vm` à 90 %+.

## Guide de migration Node.js vers Bun en sept étapes

Vous avez décidé de tester Bun sur un projet Node.js existant. Voici la procédure éprouvée par les équipes de Vercel et Railway, qui couvre les pièges les plus fréquents et garantit un retour arrière facile.

