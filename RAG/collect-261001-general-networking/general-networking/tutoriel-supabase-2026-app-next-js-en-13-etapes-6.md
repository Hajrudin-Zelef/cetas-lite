---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-6
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Cohere", "Google", "Mistral", "OpenAI", "Stripe", "United States"]
dates: []
keywords: ["apache", "aws", "cohere", "compute", "embeddings", "mistral"]
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [752, 811]
sha256: bd72fca7b9717ed04509b75b1443ecd170d95e457e704cca3791805ffab5eda3
---

# macOS (Homebrew)

L’extension pgvector, passée de la version 0.28.1 à 0.53.0 lors de la mise à jour des images Supabase de février 2026, est activable en une commande SQL : `create extension vector;`. Elle permet de stocker des embeddings (OpenAI, Mistral, Cohere) dans une colonne `vector(1536)` et d’effectuer des recherches par cosine distance avec un index HNSW. Combinée à RLS, c’est une alternative crédible à Pinecone ou Weaviate, avec un coût quasi nul en dessous de 1 million de vecteurs.

## Comparaison Supabase vs Firebase : la décision en 2026

| Critère | Supabase | Firebase | 
|---|---|---|
| Base de données | PostgreSQL 16 (relationnelle, SQL standard) | Firestore (NoSQL document) | 
| Modèle de licence | Open-source (Apache 2.0) | Propriétaire Google Cloud | 
| Self-hosting | Oui, Docker complet | Non possible | 
| Régions UE | Francfort, Paris, Dublin | europe-west1, europe-west3 | 
| Tarification entrée | Free puis 25 $/mois Pro | Spark gratuit puis Blaze pay-as-you-go | 
| Vendor lock-in | Faible (PostgreSQL standard) | Élevé (API propriétaire Firestore) | 
| SDKs officiels | JS, Swift, Kotlin, Flutter, Python | JS, Swift, Kotlin, Flutter, C++, Unity | 
| Conformité RGPD | Native, DPA standard | DPA Google + transferts US | 

Pour une analyse complète des deux plateformes, consultez notre comparatif détaillé Supabase vs Firebase 2026 ci-dessous dans la section Related Coverage. La tendance lourde en 2026 : les startups françaises levant en série A migrent massivement de Firestore vers Supabase pour la souveraineté des données et la sortie de l’écosystème Google Cloud.

## FAQ : Supabase pour les développeurs francophones

### Supabase est-il vraiment gratuit pour un projet de production ?

Oui, le tier Free est exploitable en production pour de petits projets : 50 000 utilisateurs actifs mensuels, 500 MB de base PostgreSQL, 1 GB de stockage, 2 GB de bande passante. Limitation principale : les projets sont mis en pause après 7 jours d’inactivité (aucun trafic), et il faut les réveiller manuellement depuis le dashboard. Pour de la production sérieuse, comptez sur le tier Pro à 25 $ par mois et par projet.

### Peut-on auto-héberger Supabase pour respecter la souveraineté des données ?

Oui, intégralement. Le dépôt github.com/supabase/supabase fournit un docker-compose complet avec tous les services. La documentation officielle couvre Hetzner, Scaleway, AWS, GCP et Azure. Les serveurs français comme OVHcloud Public Cloud ou Scaleway Elements sont parfaitement supportés. Comptez 8 à 16 GB de RAM minimum pour une stack production.

### Comment migrer une base existante PostgreSQL vers Supabase ?

Supabase fournit un outil `supabase db dump` et accepte les imports `pg_dump` standard. Pour une base de plus de 50 GB, utilisez la fonctionnalité **Migration via WAL** du dashboard avec un downtime minimal. Pensez à recréer les politiques RLS – elles ne sont pas exportées par défaut dans les dumps PostgreSQL.

### Quelles sont les alternatives à Supabase en 2026 ?

Les principales alternatives sont Firebase (Google), Appwrite (open-source self-hosted), PocketBase (lite Go), Convex (TypeScript-first), Pockethost et Nhost. Supabase reste leader sur le créneau open-source + PostgreSQL grâce à son écosystème npm très mature et la qualité de sa documentation, traduite partiellement en français depuis 2025.

### Supabase est-il compatible RGPD pour des données médicales ou financières ?

Supabase Cloud est certifié SOC 2 Type II et HIPAA-ready sur les tiers Team et Enterprise. Pour la France, le tier Enterprise propose un Data Processing Addendum (DPA) signable avec hébergement exclusif sur la région Paris (eu-west-3). Pour les données ultra-sensibles, l’auto-hébergement sur OVHcloud SecNumCloud reste l’option recommandée.

### Quelle est la différence entre @supabase/supabase-js et @supabase/ssr ?

Le SDK `@supabase/supabase-js` 2.105.3 est le client universel utilisable dans n’importe quel environnement JavaScript (browser, Node, Deno, Cloudflare Workers). Le package `@supabase/ssr` 0.10.2 est une couche d’abstraction spécifiquement pour Next.js, SvelteKit, Remix et Astro qui gère la synchronisation des cookies de session entre serveur et client – indispensable pour le rendu côté serveur authentifié.

### Comment gérer les coûts Supabase à grande échelle ?

Trois leviers : (1) activez **Spend Caps** dans le dashboard pour bloquer toute facturation au-delà d’un seuil ; (2) utilisez les Edge Functions plutôt que des appels DB directs depuis le client pour réduire le compute Postgres ; (3) externalisez les fichiers volumineux vers un CDN tiers (Cloudflare R2, Bunny CDN) tout en gardant les métadonnées dans Supabase Storage. Une équipe française typique reste sous 100 $/mois jusqu’à 100 000 utilisateurs actifs.

### Supabase remplace-t-il vraiment un backend Node.js ou Django ?

Pour 80 % des cas d’usage CRUD avec authentification, oui. La combinaison PostgREST + RLS + Edge Functions couvre l’essentiel. Pour des logiques métier complexes (workflows multi-étapes, intégrations B2B, calculs lourds), gardez un backend dédié en Node.js, Python ou Go qui consomme Supabase via la service_role key. Cette approche hybride est désormais standard chez les scaleups françaises.

## Conclusion : Supabase, le standard du backend français en 2026

Au terme de ce tutoriel, votre application de gestion de tickets dispose d’une authentification complète Magic Link + OAuth GitHub, d’une base PostgreSQL 16 sécurisée par Row Level Security, de stockage de pièces jointes via S3 compatible, de mises à jour temps réel WebSocket, d’une Edge Function Deno qui envoie des notifications email, et d’un pipeline GitHub Actions qui déploie tout automatiquement avec branching par PR. Vous avez écrit moins de 600 lignes de code TypeScript pour une application qui aurait demandé entre 3 000 et 5 000 lignes avec un backend custom Node.js + Express + Passport + multer + Socket.io.

Le pari de Supabase – combiner la robustesse du SQL relationnel et la productivité d’un BaaS moderne – est gagné en 2026. Avec 101 948 étoiles GitHub, une communauté hyper-active – illustrée par le tutoriel vidéo Supabase de 5 heures publié par freeCodeCamp en décembre 2025 et par les plus de 100 tutoriels pas-à-pas Supabase publiés par RapidDev depuis mars 2026 –, des prix prédictibles et un modèle 100 % open-source, c’est aujourd’hui le choix par défaut pour démarrer une application web ou mobile en France et en Europe. Pour aller plus loin, explorez les fonctionnalités Vault (gestion de secrets dans Postgres), Auth Hooks (logique custom sur signup), et l’intégration native Stripe pour les abonnements.

### Related Coverage

**Sources externes** : Documentation officielle Supabase · Tarifs Supabase · Dépôt GitHub Supabase · Documentation Next.js · Documentation PostgreSQL · Deno runtime
