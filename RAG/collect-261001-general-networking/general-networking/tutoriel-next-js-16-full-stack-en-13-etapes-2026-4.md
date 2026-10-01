---
id: collect-261001-general-networking/general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026-4
title: "Création du projet avec toutes les options activées"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026.md
source_anchor: ""
source_lines: [303, 450]
sha256: 5eca06c250b5fe9da820d3c34f48b6cebdfea59068b7d686ed3b37acf4464567
---

# Création du projet avec toutes les options activées

```
// src/app/api/posts/route.ts
import { NextResponse } from "next/server";
import { prisma } from "@/lib/db";
import { auth } from "@/lib/auth";
export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const limit = Number(searchParams.get("limit") ?? 10);
  const posts = await prisma.post.findMany({
    where: { published: true },
    orderBy: { createdAt: "desc" },
    take: Math.min(limit, 50),
  });
  return NextResponse.json({ posts }, {
    headers: { "Cache-Control": "public, s-maxage=60, stale-while-revalidate=300" },
  });
}
export async function POST(request: Request) {
  const session = await auth();
  if (!session?.user) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }
  const body = await request.json();
  const post = await prisma.post.create({ data: { ...body, authorId: session.user.id } });
  return NextResponse.json({ post }, { status: 201 });
}
```
Pour exposer un flux RSS ou Atom (très utile pour un blog), créez `src/app/feed.xml/route.ts` qui retourne du XML avec le bon en-tête `Content-Type: application/xml`. Vous pouvez utiliser la bibliothèque `feed` de Jamie Tan pour générer le contenu correctement formé. Pensez à ajouter une balise `link` dans `layout.tsx` pour que les agrégateurs RSS la détectent automatiquement.

## Étape 9 : Optimisation des images avec next/image

Le composant `Image` de Next.js fait gagner facilement 30 à 60 points sur le Lighthouse Performance Score d’un blog. Il convertit automatiquement les images en AVIF/WebP, génère des tailles responsives, ajoute un placeholder flou pendant le chargement et applique le lazy-loading natif du navigateur. En 2026, c’est un quick-win obligatoire.

```
// next.config.ts
import type { NextConfig } from "next";
const nextConfig: NextConfig = {
  images: {
    formats: ["image/avif", "image/webp"],
    remotePatterns: [
      { protocol: "https", hostname: "images.unsplash.com" },
      { protocol: "https", hostname: "cdn.tech-insider.org" },
    ],
    minimumCacheTTL: 31536000, // 1 an
  },
  experimental: {
    ppr: "incremental", // Partial Prerendering
  },
};
export default nextConfig;
// Utilisation dans un composant
import Image from "next/image";
<Image
  src="/blog/nextjs-16.webp"
  alt="Capture d'écran Next.js 16"
  width={1200}
  height={630}
  priority
  placeholder="blur"
  blurDataURL="data:image/jpeg;base64,..."
/>
```
Attention au coût Vercel : le service d’optimisation est facturé à partir de 5 000 images sources sur le plan Pro. Pour un blog avec beaucoup de visuels, envisagez Cloudflare Images (5 $ par mois pour 100 000 images) ou Bunny.net Optimizer (9,5 $ par mois). Vous pouvez aussi désactiver l’optimiseur intégré avec `unoptimized: true` et servir des images déjà compressées via votre CDN.

## Étape 10 : Partial Prerendering et streaming UI

**Partial Prerendering** (PPR), introduit en preview avec Next.js 14 et désormais activable en mode incremental dans la 16, est la fonctionnalité qui définit la nouvelle ère du rendu web. Elle combine le meilleur du statique (page servie en moins de 100 ms depuis l’edge) et du dynamique (sections personnalisées par utilisateur). Le navigateur reçoit instantanément le shell pré-rendu, et les portions dynamiques sont streamées en HTTP/2 dès qu’elles sont prêtes.

```
// src/app/blog/[slug]/page.tsx
import { Suspense } from "react";
import { prisma } from "@/lib/db";
export const experimental_ppr = true;
async function Comments({ postId }: { postId: string }) {
  // Simulation d'un appel lent (par exemple Disqus, un microservice...)
  await new Promise(r => setTimeout(r, 800));
  const comments = await prisma.comment.findMany({ where: { postId } });
  return <ul>{comments.map(c => <li key={c.id}>{c.body}</li>)}</ul>;
}
export default async function Post({ params }: { params: Promise<{ slug: string }> }) {
  const { slug } = await params;
  const post = await prisma.post.findUnique({ where: { slug } });
  if (!post) return null;
  return (
    <article className="prose mx-auto p-8">
      <h1>{post.title}</h1>
      <div dangerouslySetInnerHTML={{ __html: post.content }} />
      <h2>Commentaires</h2>
      <Suspense fallback={<p>Chargement des commentaires...</p>}>
        <Comments postId={post.id} />
      </Suspense>
    </article>
  );
}
```
Le composant `Suspense` est la clé. Tout ce qui est en dehors de la frontière est statique et pré-rendu au build. Tout ce qui est à l’intérieur attend la résolution de la promesse pour streamer le HTML supplémentaire. Le navigateur affiche immédiatement le titre et le contenu de l’article, et le bloc commentaires apparaît dès que les données sont prêtes. C’est la perception de vitesse maximale.

## Étape 11 : Tester avec Vitest et Playwright

Aucune application sérieuse ne part en production sans tests. Pour Next.js 16, la stack recommandée est **Vitest** pour les tests unitaires et d’intégration (28× plus rapide que Jest en mode watch d’après notre comparatif Vitest vs Jest), et **Playwright** pour les tests end-to-end (multi-navigateurs, multi-OS, voir notre tutoriel Playwright complet).

```
# Installation
npm install -D vitest @vitejs/plugin-react @testing-library/react jsdom
npm install -D @playwright/test
npx playwright install --with-deps chromium
# vitest.config.ts
import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import path from "node:path";
export default defineConfig({
  plugins: [react()],
  test: { environment: "jsdom", globals: true },
  resolve: { alias: { "@": path.resolve(__dirname, "./src") } },
});
# src/lib/utils.test.ts
import { describe, expect, it } from "vitest";
import { cn } from "./utils";
describe("cn", () => {
  it("fusionne les classes Tailwind sans conflit", () => {
    expect(cn("p-2", "p-4")).toBe("p-4");
    expect(cn("text-red-500", false && "text-blue-500")).toBe("text-red-500");
  });
});
```
Pour les tests E2E, créez un fichier `tests/blog.spec.ts` avec Playwright. La grande nouveauté en 2026 est le **tracing instantané** : en cas d’échec, vous obtenez une vidéo, un screenshot et le DOM complet à chaque étape, ce qui réduit drastiquement le temps de debug. Configurez la base URL via `baseURL: "http://localhost:3000"` dans `playwright.config.ts` et lancez la suite avec `npx playwright test`.

## Étape 12 : Build de production et analyse de bundle

Avant de déployer, exécutez `npm run build` pour générer le bundle de production. Next.js 16 a retiré les colonnes `size` et `First Load JS` de la sortie de build (elles étaient trompeuses en architecture RSC). Pour mesurer le bundle envoyé au navigateur, utilisez le plugin `@next/bundle-analyzer` qui affiche un treemap interactif des dépendances client.

```
# Installation et activation
npm install -D @next/bundle-analyzer
# next.config.ts
import bundleAnalyzer from "@next/bundle-analyzer";
const withBundleAnalyzer = bundleAnalyzer({
  enabled: process.env.ANALYZE === "true",
});
export default withBundleAnalyzer(nextConfig);
# Lancement
ANALYZE=true npm run build
# Ouvre automatiquement http://localhost:8888 avec le treemap
```
Cherchez les gros offenders : `moment.js` (250 KB, remplaçable par `date-fns` ou `Temporal`), `lodash` entier (replacer par `lodash-es` avec import nommé pour le tree-shaking), des SDK marketing chargés trop tôt. Sur un blog Next.js bien optimisé, le bundle initial Client doit rester sous **90 KB gzippés**. Au-delà, le Time-to-Interactive sur 3G se dégrade fortement.

## Étape 13 : Déploiement sur Vercel ou en auto-hébergé

