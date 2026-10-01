---
id: collect-261001-general-networking/general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026-2
title: "Création du projet avec toutes les options activées"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-next-js-16-full-stack-en-13-etapes-2026.md
source_anchor: ""
source_lines: [56, 163]
sha256: c4ab1d84208c9379c383166962477851ea6649bae238626bdfe3deb1514565a7
---

# Création du projet avec toutes les options activées

```
src/
├── app/
│   ├── layout.tsx        # Layout racine (obligatoire)
│   ├── page.tsx          # Route /
│   ├── globals.css       # Styles globaux Tailwind
│   ├── (auth)/           # Groupe de routes sans URL
│   │   ├── login/page.tsx     # Route /login
│   │   └── signup/page.tsx    # Route /signup
│   ├── blog/
│   │   ├── page.tsx           # Route /blog
│   │   ├── [slug]/page.tsx    # Route dynamique /blog/mon-article
│   │   └── loading.tsx        # Skeleton de chargement
│   ├── api/
│   │   └── posts/route.ts     # API REST GET/POST /api/posts
│   └── error.tsx              # Boundary d'erreur global
├── components/
├── lib/
│   ├── db.ts             # Client Prisma
│   └── auth.ts           # Configuration Auth.js
└── middleware.ts         # Middleware Edge (protection routes)
```
La règle la plus mal comprise est la suivante : **par défaut, tous les composants sont des Server Components**. Cela signifie qu’ils s’exécutent côté serveur, ne sont jamais envoyés au navigateur sous forme de JavaScript, et peuvent accéder directement à la base de données ou aux secrets. Pour basculer un composant en Client Component (avec hooks `useState`, gestionnaires d’événements onClick, etc.), il faut ajouter la directive `'use client'` sur la première ligne du fichier.

Cette dualité est la source de 80 % des erreurs en début d’apprentissage. Si vous voyez `You're importing a component that needs useState. It only works in a Client Component`, vous savez maintenant quoi faire : ajouter `'use client'` en haut du fichier concerné, ou refactoriser pour garder la logique d’état dans un composant enfant.

## Étape 3 : Configurer Tailwind CSS 4 et la typographie

Le starter `create-next-app` installe Tailwind CSS 4, qui utilise un nouveau modèle de configuration basé sur CSS (plus de `tailwind.config.ts` obligatoire). Pour un blog, ajoutons le plugin `@tailwindcss/typography` qui fournit la classe `prose` indispensable pour styliser le contenu Markdown ou MDX. Installons aussi `clsx` et `tailwind-merge` pour gérer les variantes conditionnelles.

```
npm install @tailwindcss/typography clsx tailwind-merge
# src/app/globals.css
@import "tailwindcss";
@plugin "@tailwindcss/typography";
@theme {
  --color-brand-50: #f0f9ff;
  --color-brand-500: #0ea5e9;
  --color-brand-900: #0c4a6e;
  --font-sans: "Inter", system-ui, sans-serif;
}
body { @apply bg-white text-slate-900 antialiased; }
```
Créez ensuite un utilitaire `cn` dans `src/lib/utils.ts` qui combinera `clsx` et `tailwind-merge` pour fusionner intelligemment des classes Tailwind sans conflits. Cette fonction est utilisée dans la quasi-totalité des projets sérieux, et c’est aussi la convention adoptée par shadcn/ui, la bibliothèque de composants la plus populaire de l’écosystème React en 2026.

```
// src/lib/utils.ts
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
```
Pour les polices, Next.js intègre `next/font/google` qui télécharge la police au build, l’héberge en local et l’inline en CSS critique. Le score Core Web Vitals s’en trouve immédiatement amélioré (élimination du CLS lié au flash of unstyled text). Ajoutez Inter dans `layout.tsx` avec `import { Inter } from "next/font/google"`, puis appliquez `className={inter.className}` sur la balise `html`.

## Étape 4 : Mettre en place Prisma et PostgreSQL

Pour un blog production-ready, nous utilisons PostgreSQL via Prisma 6. Installez Postgres en local avec Docker (`docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:17`) ou utilisez une instance hébergée chez Supabase, Neon ou Scaleway. Pour cet article, je passe par Neon, qui offre un tier gratuit suffisant pour le dev et stocke les données dans la région `eu-central-1` (Frankfurt).

```
# Installation de Prisma
npm install prisma @prisma/client
npm install -D prisma
# Initialisation
npx prisma init --datasource-provider postgresql
# Édition du fichier prisma/schema.prisma
model Post {
  id        String   @id @default(cuid())
  slug      String   @unique
  title     String
  excerpt   String
  content   String   @db.Text
  published Boolean  @default(false)
  authorId  String
  author    User     @relation(fields: [authorId], references: [id])
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt
  @@index([published, createdAt])
}
model User {
  id    String  @id @default(cuid())
  email String  @unique
  name  String?
  posts Post[]
}
```
Renseignez votre `DATABASE_URL` dans le fichier `.env` (au format `postgresql://user:pass@host:5432/db?sslmode=require`). Pour Neon, ajoutez `&pgbouncer=true&connection_limit=1` à l’URL — c’est obligatoire pour éviter d’épuiser le pool de connexions sur un environnement serverless où chaque invocation peut ouvrir une nouvelle connexion. Cette erreur de configuration est la première cause de panne en production chez les nouveaux utilisateurs de Next.js avec Prisma.

```
# Génération du client TypeScript et création des tables
npx prisma generate
npx prisma migrate dev --name init
# src/lib/db.ts (singleton pour éviter les fuites de connexion)
import { PrismaClient } from "@prisma/client";
const globalForPrisma = globalThis as unknown as {
  prisma: PrismaClient | undefined;
};
export const prisma =
  globalForPrisma.prisma ??
  new PrismaClient({ log: ["error", "warn"] });
if (process.env.NODE_ENV !== "production") globalForPrisma.prisma = prisma;
```
Le pattern singleton ci-dessus est **obligatoire**. En mode développement, Next.js recharge à chaud les modules, ce qui crée une nouvelle instance Prisma à chaque modification de fichier. Sans le singleton, vous saturez votre instance Postgres en quelques secondes (erreur `too many connections`). C’est la pitfall numéro 1 documentée sur GitHub Discussions Vercel depuis Next.js 13.

## Étape 5 : Récupérer les données depuis un Server Component

Le principal intérêt des React Server Components est de pouvoir interroger directement la base de données depuis le composant qui rend la page, sans passer par un endpoint REST ou GraphQL intermédiaire. Cela simplifie l’architecture, réduit la latence (un seul aller-retour réseau pour le navigateur) et améliore la sécurité (les secrets de connexion ne quittent jamais le serveur).

