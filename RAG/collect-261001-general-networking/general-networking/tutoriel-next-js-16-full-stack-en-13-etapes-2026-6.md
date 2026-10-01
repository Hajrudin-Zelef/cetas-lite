---
id: collect-261001-general-networking/general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026-6
title: "Création du projet avec toutes les options activées"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["attention", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026.md
source_anchor: ""
source_lines: [537, 615]
sha256: 6e866dd485e7f2a723965a984b172ab0548bcdeb1d8ee1114511d0eb6139bec2
---

# Création du projet avec toutes les options activées

Le résultat le plus parlant est l’écart sur le HMR (Hot Module Replacement). Quand vous modifiez un composant, Webpack mettait près d’une seconde à actualiser la page ; Turbopack le fait en 110 ms en moyenne. Cumulé sur une journée de développement avec 200 sauvegardes, c’est plus de deux minutes économisées par développeur, par jour. Pour une équipe de cinq, cela représente environ une journée par mois.

## Sécurité : ce que change la mise à jour de mai 2026

La veille déjà, le 6 mai 2026, un avis GitHub signalait un problème de déni de service distinct, corrigé par les versions **Next.js 15.5.16** et **16.2.5** — un premier signal avant que Vercel ne publie, le 7 mai 2026, une mise à jour de sécurité coordonnée corrigeant **13 avis** couvrant le contournement de middleware, le déni de service, le SSRF (Server-Side Request Forgery), l’empoisonnement de cache et le cross-site request forgery. Les versions corrigées de ce second train de correctifs sont **Next.js 15.5.18** et **16.2.6**. Côté React, les patches sont 19.0.6, 19.1.7 et 19.2.6 selon la ligne mineure utilisée.

Si vous tournez en production sur une version antérieure, mettez à jour immédiatement : Vercel a explicitement enjoint tous les projets encore sur les branches **Next.js 13.x et 14.x** à migrer vers 15.5.18 ou 16.2.6, avec `npm install [email protected]` ou `npm install [email protected]` selon votre majeure, puis redéployez. Un guide de mise à jour publié le 26 juin 2026 par Rabinarayan Patra précise en outre que tout projet resté sur une version comprise entre **16.0.0 et 16.2.5** doit impérativement passer à la 16.2.6 pour bénéficier des correctifs. Netlify a confirmé que les projets hébergés chez eux nécessitent aussi le redéploiement pour bénéficier des correctifs. Vercel a poussé les patches automatiquement sur les déploiements existants.

Pour réduire votre surface d’attaque, suivez trois règles. D’abord, configurez les en-têtes HTTP de sécurité dans `next.config.ts` via la fonction `headers()` : Content-Security-Policy, Strict-Transport-Security, X-Content-Type-Options, Referrer-Policy. Ensuite, validez systématiquement les entrées utilisateur avec Zod ou Valibot (jamais `any` ni cast en TypeScript). Enfin, activez les Security Headers recommandées par Next.js et auditez votre site avec securityheaders.com.

## SEO et Core Web Vitals

Next.js 16 fournit nativement des outils SEO de qualité. Le composant `Metadata` de l’App Router génère automatiquement les balises `<title>`, `<meta description>`, Open Graph, Twitter Card et le JSON-LD. Combiné avec le sitemap dynamique (`src/app/sitemap.ts`) et le robots.txt programmable (`src/app/robots.ts`), vous obtenez une base solide pour Google et Bing.

```
// src/app/sitemap.ts
import type { MetadataRoute } from "next";
import { prisma } from "@/lib/db";
export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const posts = await prisma.post.findMany({
    where: { published: true },
    select: { slug: true, updatedAt: true },
  });
  return [
    { url: "https://exemple.fr", lastModified: new Date(), priority: 1.0 },
    ...posts.map(p => ({
      url: `https://exemple.fr/blog/${p.slug}`,
      lastModified: p.updatedAt,
      priority: 0.7,
      changeFrequency: "weekly" as const,
    })),
  ];
}
// src/app/blog/[slug]/page.tsx — metadata dynamique
import type { Metadata } from "next";
export async function generateMetadata({ params }: { params: Promise<{ slug: string }> }): Promise<Metadata> {
  const { slug } = await params;
  const post = await prisma.post.findUnique({ where: { slug } });
  if (!post) return {};
  return {
    title: post.title,
    description: post.excerpt,
    openGraph: { title: post.title, description: post.excerpt, type: "article" },
  };
}
```
Les Core Web Vitals 2026 sont LCP (Largest Contentful Paint), INP (Interaction to Next Paint, qui a remplacé FID en mars 2024) et CLS (Cumulative Layout Shift). Sur un blog Next.js correctement configuré, vous devez viser LCP sous 1,5 s, INP sous 100 ms et CLS sous 0,05. Mesurez avec PageSpeed Insights et corrigez prioritairement le LCP en optimisant l’image principale et en activant `priority` sur celle-ci.

## Comparaison avec les alternatives 2026

Le paysage frameworks JavaScript en 2026 est plus diversifié qu’il ne l’a jamais été. Voici comment Next.js se positionne face aux principales alternatives, en tenant compte des dernières releases majeures de chaque concurrent.

| Framework | Modèle de rendu | Forces | Faiblesses | Communauté | 
|---|---|---|---|---|
| Next.js 16 | RSC + ISR + PPR | Écosystème, Vercel, App Router mature | Couplage Vercel, complexité RSC/Client | Très large | 
| Remix 3 (React Router) | SSR + nested routes | Web standards, mutations sans Server Actions | Communauté plus restreinte | Large | 
| SvelteKit 2 | SSR + Islands | Bundle minuscule, DX excellente | Écosystème UI plus jeune | Moyenne | 
| Nuxt 4 | SSR Vue + ISR | Auto-imports, Nitro polyglotte | Vue, moins demandé en France | Moyenne | 
| Astro 5 | Islands + MDX natif | Performances brutes, multi-frameworks | Moins adapté aux apps lourdes | Moyenne | 
| TanStack Start | RSC + TypeScript-first | Type-safety, signaux fins | Encore en beta, peu de cas d’usage prod | Petite | 

Pour un site de contenu (blog, magazine, documentation), **Astro** reste souvent plus performant brutalement. Pour une SaaS interactive ou une application e-commerce, **Next.js** garde l’avantage écosystème et la cohérence avec React. Pour une startup early-stage qui veut tester rapidement plusieurs idées, **SvelteKit** ou **TanStack Start** peuvent accélérer la vélocité au prix d’une communauté plus réduite. Voir aussi notre tutoriel Supabase + Next.js et tutoriel TanStack Start pour pousser le comparatif plus loin.

## Foire aux questions

### Quelle est la dernière version stable de Next.js en mai 2026 ?

Next.js **16.2.6**, publiée le 7 mai 2026, était la version majeure stable recommandée au printemps, mais plusieurs correctifs se sont enchaînés depuis : un correctif de sécurité coordonné publié par Vercel le 20 juillet 2026 a fait passer les branches supportées à **16.2.11** et **15.5.21**, puis la branche 16 a continué d’avancer jusqu’à la **16.3.5**, désignée dernière version stable sur le site officiel de Next.js au **11 septembre 2026** — c’est elle qu’il faut installer en priorité aujourd’hui. La ligne 15, elle, a continué d’évoluer en parallèle en mode maintenance : la dernière révision publiée est la **15.5.25**, sortie le **31 août 2026**. Attention cependant, Next.js 15 est désormais en **Maintenance LTS** avec une fin de vie fixée au **21 octobre 2026**, deux ans jour pour jour après sa sortie initiale du 21 octobre 2024 (source : HeroDevs) — les équipes encore sur cette ligne ont donc intérêt à planifier leur migration vers la 16 avant cette échéance.

### Faut-il encore apprendre le Pages Router ?

Non, sauf si vous maintenez un projet legacy. L’App Router est la voie recommandée depuis Next.js 13.4 et la quasi-totalité des nouveautés (Server Actions, PPR, after, parallel routes) ne sont pas portées sur le Pages Router. Pour un nouveau projet en 2026, démarrez directement avec l’App Router et oubliez le Pages.

### Quelle base de données choisir avec Next.js ?

PostgreSQL hébergé chez Neon, Supabase ou Scaleway pour 90 % des cas. Si vous avez des besoins très simples (clé-valeur, sessions), regardez Upstash Redis ou Vercel KV. Pour les usages analytiques temps réel, ClickHouse via notre tutoriel ClickHouse est imbattable.

### Vercel ou auto-hébergé : que choisir en France ?

