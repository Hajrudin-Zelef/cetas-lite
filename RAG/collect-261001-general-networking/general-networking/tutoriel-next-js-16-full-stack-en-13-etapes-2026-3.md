---
id: collect-261001-general-networking/general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026-3
title: "Création du projet avec toutes les options activées"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026.md
source_anchor: ""
source_lines: [164, 302]
sha256: 5396e2800edf0483b7fcdad8b2c9767d91206c3d4483d6d86f4379d382b4ecae
---

# Création du projet avec toutes les options activées

```
// src/app/blog/page.tsx
import Link from "next/link";
import { prisma } from "@/lib/db";
export const revalidate = 60; // ISR : régénération toutes les 60 secondes
export default async function BlogIndex() {
  const posts = await prisma.post.findMany({
    where: { published: true },
    orderBy: { createdAt: "desc" },
    select: {
      id: true,
      slug: true,
      title: true,
      excerpt: true,
      createdAt: true,
      author: { select: { name: true } },
    },
  });
  return (
    <main className="mx-auto max-w-3xl px-4 py-12">
      <h1 className="text-4xl font-bold tracking-tight">Le blog</h1>
      <ul className="mt-8 space-y-6">
        {posts.map((post) => (
          <li key={post.id} className="border-b border-slate-200 pb-6">
            <Link href={`/blog/${post.slug}`} className="block hover:opacity-80">
              <h2 className="text-xl font-semibold">{post.title}</h2>
              <p className="mt-2 text-slate-600">{post.excerpt}</p>
              <p className="mt-2 text-sm text-slate-400">
                Par {post.author.name} — {post.createdAt.toLocaleDateString("fr-FR")}
              </p>
            </Link>
          </li>
        ))}
      </ul>
    </main>
  );
}
```
Notez la directive `export const revalidate = 60`. Elle active **Incremental Static Regeneration** (ISR) : Next.js va générer la page une fois, la mettre en cache CDN, puis la régénérer en arrière-plan toutes les 60 secondes lors d’une nouvelle requête. C’est le meilleur compromis entre fraîcheur des données et performance, et c’est la stratégie utilisée par Notion, Vercel et The Washington Post sur leurs pages produit.

Pour les pages d’articles individuels, créez `src/app/blog/[slug]/page.tsx`. Le paramètre `slug` est accessible via les props (typage `params: Promise<{ slug: string }>` depuis Next.js 15, async désormais obligatoire). Utilisez `generateStaticParams` pour pré-générer toutes les pages au build, ce qui réduit le coût d’exécution serverless et améliore le TTFB.

## Étape 6 : Authentification avec Auth.js v5 (NextAuth)

L’authentification est un cas d’usage où Next.js brille particulièrement en 2026. Auth.js v5 (anciennement NextAuth.js) supporte nativement l’App Router, les Server Actions et le rendu en RSC. Nous l’utilisons avec l’adaptateur Prisma pour stocker les utilisateurs et sessions en base, plus le provider Google OAuth pour la simplicité.

```
npm install next-auth@beta @auth/prisma-adapter
# src/lib/auth.ts
import NextAuth from "next-auth";
import Google from "next-auth/providers/google";
import { PrismaAdapter } from "@auth/prisma-adapter";
import { prisma } from "@/lib/db";
export const { handlers, auth, signIn, signOut } = NextAuth({
  adapter: PrismaAdapter(prisma),
  providers: [
    Google({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    }),
  ],
  session: { strategy: "database" },
  pages: { signIn: "/login" },
});
# src/app/api/auth/[...nextauth]/route.ts
export { GET, POST } from "@/lib/auth";
# src/middleware.ts
export { auth as middleware } from "@/lib/auth";
export const config = { matcher: ["/admin/:path*"] };
```
Récupérez vos identifiants OAuth depuis la console Google Cloud (`console.cloud.google.com`) en créant un projet, en activant OAuth consent screen, et en ajoutant `http://localhost:3000/api/auth/callback/google` comme URL de redirection autorisée. Pour la production, ajoutez aussi votre domaine de prod. Stockez les secrets dans `.env.local`, jamais dans le dépôt Git.

Le middleware `src/middleware.ts` protège toutes les routes commençant par `/admin`. Si un utilisateur non authentifié essaie d’y accéder, il est redirigé vers `/login`. Ce middleware s’exécute en Edge Runtime, ce qui le rend extrêmement rapide (moins de 5 ms pour la vérification de session). Pour les routes publiques, on peut récupérer la session dans un Server Component via `const session = await auth()`.

## Étape 7 : Server Actions pour créer un article

Les **Server Actions** sont la fonctionnalité phare de Next.js depuis la version 14, et désormais incontournables avec la 16. Elles permettent de définir des fonctions qui s’exécutent côté serveur mais peuvent être appelées directement depuis un formulaire HTML ou un Client Component, sans avoir à créer manuellement un endpoint REST. Le résultat : moins de code, meilleure DX et moins de surface d’attaque.

```
// src/app/admin/posts/new/actions.ts
"use server";
import { z } from "zod";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { prisma } from "@/lib/db";
import { auth } from "@/lib/auth";
const PostSchema = z.object({
  title: z.string().min(5).max(120),
  slug: z.string().regex(/^[a-z0-9-]+$/),
  excerpt: z.string().min(20).max(300),
  content: z.string().min(50),
});
export async function createPost(_: unknown, formData: FormData) {
  const session = await auth();
  if (!session?.user?.id) {
    return { error: "Non authentifié" };
  }
  const parsed = PostSchema.safeParse(Object.fromEntries(formData));
  if (!parsed.success) {
    return { error: parsed.error.flatten().fieldErrors };
  }
  await prisma.post.create({
    data: {
      ...parsed.data,
      authorId: session.user.id,
      published: true,
    },
  });
  revalidatePath("/blog");
  redirect("/blog");
}
```
Côté UI, créez un Client Component qui utilise `useActionState` (renommé depuis `useFormState` dans React 19). Ce hook gère automatiquement l’état du formulaire, les transitions et les erreurs renvoyées par la Server Action. Le résultat est un formulaire qui fonctionne **même sans JavaScript activé** (progressive enhancement), ce qui est rarement atteint dans une SPA traditionnelle.

```
// src/app/admin/posts/new/form.tsx
"use client";
import { useActionState } from "react";
import { createPost } from "./actions";
export function NewPostForm() {
  const [state, formAction, pending] = useActionState(createPost, null);
  return (
    <form action={formAction} className="space-y-4">
      <input name="title" placeholder="Titre" required className="block w-full rounded border p-2" />
      <input name="slug" placeholder="slug-url" required className="block w-full rounded border p-2" />
      <textarea name="excerpt" placeholder="Résumé" required className="block w-full rounded border p-2" />
      <textarea name="content" placeholder="Contenu Markdown" required rows={12} className="block w-full rounded border p-2" />
      {state?.error && <p className="text-red-600 text-sm">{JSON.stringify(state.error)}</p>}
      <button disabled={pending} className="rounded bg-brand-500 px-4 py-2 text-white">
        {pending ? "Publication..." : "Publier"}
      </button>
    </form>
  );
}
```
## Étape 8 : Routes API et streaming avec route handlers

Bien que les Server Actions couvrent 80 % des besoins, certains cas réclament une API REST classique : webhooks, intégrations mobiles, partenaires externes. Next.js permet de créer ces endpoints avec les **route handlers**, fichiers `route.ts` dans le dossier `app`. Ces handlers supportent toutes les méthodes HTTP, le streaming et le retour de réponses typées.

