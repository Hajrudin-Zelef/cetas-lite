---
id: collect-261001-general-networking/general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026-5
title: "Création du projet avec toutes les options activées"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Mistral", "OpenAI"]
dates: []
keywords: ["agents", "benchmarks", "claude", "memory", "mistral", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026.md
source_anchor: ""
source_lines: [451, 536]
sha256: 5e0af71a6512278e08dcec770005f2194d6a4f11da3d9e89a90e93ed6b37fc65
---

# Création du projet avec toutes les options activées

Le déploiement le plus rapide reste Vercel : connectez votre dépôt GitHub, ajoutez vos variables d’environnement, et la première URL de production est en ligne en deux minutes. Le plan Hobby est gratuit (100 GB de bande passante, 6 000 builds/mois). Le plan Pro à 20 $ par mois et par membre suffit pour la plupart des petites équipes.

| Hébergeur | Tarif mensuel | Région UE | Spécificité | 
|---|---|---|---|
| Vercel Hobby | 0 € | fra1 (Frankfurt) | 100 GB bande passante, 100k req fonctions | 
| Vercel Pro | 20 €/membre | fra1 + cdg1 (Paris) | 1 TB bande, support, équipes | 
| Netlify Starter | 0 € | eu-west-3 | 100 GB bande, 125k req fonctions | 
| Scaleway Serverless Containers | ~5 € | par1 (Paris) | Souverain UE, RGPD/SecNumCloud-friendly | 
| Coolify (auto-hébergé) | 5 € VPS Hetzner | fsn1, hel1 | Open source, illimité, voir notre tutoriel Coolify | 
| Dokploy | VPS de votre choix | libre | Open source, GUI, déploiement Docker Compose | 

Pour un déploiement souverain en France ou dans l’UE, deux options sérieuses émergent en 2026. Premièrement, **Scaleway Serverless Containers** dans la région Paris, compatible avec un build standalone Next.js (`output: "standalone"` dans `next.config.ts`). Deuxièmement, **Coolify** auto-hébergé sur un VPS Hetzner ou OVH, qui offre une expérience semblable à Vercel mais sans dépendance externe. Idéal pour les organisations soumises au RGPD strict ou à la doctrine Cloud au centre.

```
# Build Docker pour auto-hébergé
# next.config.ts
const nextConfig: NextConfig = {
  output: "standalone",
};
# Dockerfile minimal
FROM node:22-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build
FROM node:22-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["node", "server.js"]
```
## 5 pièges classiques (et comment les éviter)

**Piège 1 — Importer un fichier serveur dans un composant client.** Si vous écrivez `import { prisma } from "@/lib/db"` dans un fichier marqué `"use client"`, vous exposez les secrets de connexion et le code Prisma au bundle navigateur. Next.js détecte la plupart de ces fuites avec l’erreur `Module not found: Can't resolve 'fs'`, mais pas toujours. Ajoutez `import "server-only"` en première ligne de vos fichiers serveur pour bloquer toute importation client.

**Piège 2 — Oublier le singleton Prisma.** Documenté à l’étape 4. En dev, le HMR de Turbopack provoque des fuites de connexion. Le pattern `globalThis.prisma` est non négociable.

**Piège 3 — Cookies et sessions dans un Server Component pré-rendu.** Si vous appelez `cookies()` ou `auth()` dans une route pré-rendue statiquement, Next.js bascule automatiquement la page en dynamique, ce qui peut faire exploser votre facture serverless. Solution : isoler la lecture de session dans un sous-composant entouré de `<Suspense>` avec PPR activé.

**Piège 4 — Cache agressif sur les Server Actions.** Une Server Action qui modifie des données doit appeler `revalidatePath()` ou `revalidateTag()` à la fin, sinon la page index continue d’afficher l’ancienne version pendant 60 secondes (votre durée d’ISR). Beaucoup de développeurs oublient cette étape et accusent à tort le framework d’être bogué.

**Piège 5 — Sous-estimer le coût du middleware Edge.** Le middleware s’exécute sur *chaque* requête, y compris les assets statiques si vous n’avez pas configuré le `matcher`. Une logique d’A/B testing trop lourde ajoute 30 à 50 ms par requête. Limitez le middleware au strict nécessaire et utilisez `matcher` pour exclure `/_next/static`, `/_next/image` et les fichiers favicon/sitemap.

## 8 problèmes courants et leurs solutions

| Symptôme | Cause probable | Solution | 
|---|---|---|
| `Module not found: Can't resolve 'fs'` | Import de module Node dans un Client Component | Déplacer vers un Server Component ou utiliser `dynamic(() => ..., { ssr: false })` | 
| `Hydration mismatch` dans la console | HTML serveur différent du HTML client (date, Math.random) | Wrapper la valeur dans `useEffect` ou ajouter`suppressHydrationWarning` | 
| `Error: Dynamic server usage` | `cookies()` ,`headers()` ou`searchParams` dans une page statique | Marquer la page comme dynamique avec `export const dynamic = "force-dynamic"` | 
| `too many connections` Postgres | Pas de singleton Prisma OU pas de `pgbouncer=true` | Ajouter le singleton (étape 4) et activer pgBouncer | 
| Build OOM (heap out of memory) | Build trop volumineux ou fuite mémoire dans `generateStaticParams` | `NODE_OPTIONS="--max-old-space-size=4096" npm run build` | 
| Images Unsplash non chargées | Domaine manquant dans `remotePatterns` | Ajouter le hostname dans `next.config.ts` (étape 9) | 
| Sessions Auth.js qui se déconnectent | Cookie domain mal configuré sur sous-domaine | Définir `AUTH_TRUST_HOST=true` et`cookies.sessionToken.options.domain` | 
| Server Action 404 en prod | Cache CDN agressif sur POST | Ajouter `Cache-Control: no-store` via middleware sur`/_next/...` | 

## Astuces avancées pour aller plus loin

**Parallel Routes et Intercepting Routes.** Vous pouvez créer des modales préservant l’URL (utile pour partager un lien vers une photo agrandie) en combinant `(.)`, `(..)` et un slot parallèle `@modal`. Cette technique remplace les bibliothèques de modales tierces avec une intégration native au routeur.

**Edge Functions internationales.** Si vous servez des utilisateurs en France et au Québec, configurez `preferredRegion = ["fra1", "yul1"]` dans vos route handlers pour que Vercel route automatiquement vers la région la plus proche. La latence p99 chute de 180 ms à 35 ms pour les utilisateurs québécois.

**Streaming OpenAI avec `after()`.** La fonction `after()` (anciennement `unstable_after`) permet d’exécuter du code après la réponse au client. Idéal pour logger un événement analytics, déclencher un webhook Slack ou rafraîchir un cache, sans pénaliser la latence perçue de l’utilisateur.

**Internationalisation avec next-intl.** Pour un site multilingue FR/EN/DE, la bibliothèque `next-intl` (200 KB, compatible RSC) est devenue le standard depuis l’abandon du routage i18n natif de Next.js dans le Pages Router. Elle gère les messages, les pluriels ICU, les formats de date `Intl.DateTimeFormat` et la négociation linguistique via Accept-Language.

**Intégrer une IA générative.** Pour ajouter un assistant conversationnel à votre blog, combinez Next.js avec le Vercel AI SDK et un modèle Mistral ou Claude. Voir notre guide Mistral Medium 3.5 pour le choix du modèle et notre tutoriel CrewAI pour les agents multi-étapes.

## Performance : benchmarks réels Next.js 16

Pour donner des chiffres concrets, j’ai mesuré le projet de ce tutoriel sur un MacBook Pro M3 Pro et un serveur Hetzner CX22 (2 vCPU, 4 GB RAM). Les résultats illustrent l’écart entre Webpack et Turbopack, et l’amélioration apportée par PPR sur une page mixte statique/dynamique.

| Métrique | Next.js 14 + Webpack | Next.js 16 + Turbopack | Gain | 
|---|---|---|---|
| Démarrage `next dev` (M3 Pro) | 4,1 s | 0,9 s | 4,5× | 
| HMR au save (composant 200 LOC) | 820 ms | 110 ms | 7,5× | 
| Build production (15 pages, 30 composants) | 52 s | 22 s | 2,4× | 
| Bundle Client initial gzippé | 94 KB | 78 KB | -17% | 
| TTFB page d’accueil ISR (fra1) | 62 ms | 58 ms | -7% | 
| TTFB page article PPR streamée | n/a | 41 ms | nouveau | 
| Mémoire `next dev` au bout d’1h | 1,8 GB | 980 MB | -46% | 

