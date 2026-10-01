---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-1
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["compute", "mai", "sol"]
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [1, 83]
sha256: 5f1e6b7d99b536522dcd289d3278bf18b3fec44898c8bbc47dc16bce1035e075
---

# macOS (Homebrew)

**Tutoriel Supabase 2026** – Avec plus de **101 948 étoiles GitHub**, **12 315 forks** et la version **@supabase/supabase-js 2.105.3** publiée sur npm – toujours la branche v2.x que la documentation officielle de la CLI Supabase, actualisée en août 2026, désigne comme SDK par défaut –, Supabase s’est imposé en France comme la principale alternative open-source à Firebase. La plateforme combine une base PostgreSQL gérée, une authentification multi-fournisseurs, du stockage S3-compatible, du temps réel via WebSockets et des Edge Functions Deno – le tout exposé par une API REST auto-générée. Ce guide pratique de 13 étapes vous montre comment construire une application complète Next.js 16 + Supabase, déployable en production, en moins de 90 minutes.

*Publié le 06 avril 2026 – Tutoriel testé sur Ubuntu 24.04 LTS, macOS 15 Sequoia et Windows 11 24H2 avec Node.js 22 LTS.*

## Pourquoi choisir Supabase en 2026 : un BaaS open-source dominant

Supabase n’est plus une promesse : c’est devenu en 2026 le **backend-as-a-service open-source** le plus adopté par les développeurs français et européens. Là où Firebase impose un modèle propriétaire et un verrouillage Google Cloud, Supabase repose sur une stack 100 % open-source – PostgreSQL 16, GoTrue pour l’authentification, Realtime en Elixir, Storage en Go, et PostgREST pour exposer automatiquement la base sous forme d’API REST. Ce socle technique permet à un développeur de cloner toute son infrastructure en local avec la CLI Supabase, de l’auto-héberger sur Hetzner ou Scaleway pour respecter le RGPD, ou de migrer vers le cloud managé sans changer une ligne de code applicatif.

L’écosystème a explosé : la plateforme dépasse les **101 948 étoiles GitHub** en mai 2026, avec **12 315 forks** actifs et un rythme de contributions qui place le projet parmi les 30 premiers dépôts mondiaux. Côté financement, Supabase a levé plus de 120 millions de dollars cumulés depuis sa création, dont une série B de 80 millions menée par Felicis Ventures, valorisant la société à environ 2 milliards de dollars selon les dernières fuites du marché secondaire. Pour les équipes françaises, l’argument décisif reste la maîtrise des données : trois régions européennes – Francfort, Paris et Dublin – permettent de garder le trafic utilisateur sur le sol UE, condition indispensable pour les secteurs santé, finance et secteur public.

Ce tutoriel construit une application réelle : une plateforme de gestion de tickets support avec authentification email + magic link, table PostgreSQL avec Row Level Security, upload de pièces jointes vers Supabase Storage, mises à jour en temps réel via Realtime, et une Edge Function Deno qui envoie un email de notification. Le code complet est compatible avec Next.js 16.2.4 et React 19.2.6, déployable sur Vercel ou auto-hébergé via Docker.

## Prérequis : versions exactes pour suivre ce tutoriel sans erreur

Avant d’écrire la première ligne de code, vérifiez que votre poste de développement dispose des outils suivants. Toutes les versions ont été validées en avril 2026 sur les trois systèmes d’exploitation majeurs. Une version trop ancienne de Node.js ou de Docker provoquera des erreurs cryptiques sur les Edge Functions et la CLI Supabase.

| Outil | Version minimale | Version testée | Commande de vérification | 
|---|---|---|---|
| Node.js | 20.0.0 LTS | 22.14.0 LTS | `node --version` | 
| npm | 10.0.0 | 10.9.2 | `npm --version` | 
| Docker Desktop | 4.30 | 4.38.0 | `docker --version` | 
| Supabase CLI | 1.220.0 | 2.6.8 | `supabase --version` | 
| Git | 2.40 | 2.45.2 | `git --version` | 
| @supabase/supabase-js | 2.100.0 | 2.105.3 | `npm view @supabase/supabase-js version` | 
| @supabase/ssr | 0.10.0 | 0.10.2 | `npm view @supabase/ssr version` | 
| Next.js | 15.0.0 | 16.2.4 | `npx next --version` | 
| TypeScript | 5.4.0 | 6.0.3 | `npx tsc --version` | 
| Deno (Edge Functions) | 2.0 | 2.4.1 | `deno --version` | 

Côté compte, créez gratuitement un projet sur supabase.com/dashboard. Le tier Free offre 500 MB de base PostgreSQL, 1 GB de stockage, 50 000 utilisateurs actifs mensuels (MAU) et 2 projets simultanés – largement suffisant pour ce tutoriel. Si votre application dépasse ces seuils, le passage au tier Pro à 25 $ par projet et par mois débloque 8 GB de base, 100 GB de stockage et 500 000 invocations Edge Functions incluses, avec 10 $ de crédits compute offerts chaque mois.

## Étape 1 : Installer la CLI Supabase et initialiser le projet local

La CLI Supabase est le couteau suisse du développeur : elle clone un environnement complet – Postgres 16, GoTrue, Storage, Realtime, Studio, Inbucket pour les emails de test – dans des conteneurs Docker locaux. Ce mode local accélère drastiquement le cycle de développement : pas de latence réseau, pas de risque de polluer la base de production, et la possibilité de versionner toutes les migrations dans Git.

```
# macOS (Homebrew)
brew install supabase/tap/supabase
# Linux (npm global)
npm install -g supabase
# Windows (Scoop)
scoop bucket add supabase https://github.com/supabase/scoop-bucket.git
scoop install supabase
# Vérifier l'installation
supabase --version
# Sortie attendue : 2.6.8
# Créer le dossier projet
mkdir tickets-support && cd tickets-support
supabase init
# La CLI génère :
# supabase/config.toml      → configuration locale
# supabase/migrations/      → migrations SQL versionnées
# supabase/seed.sql         → données de test
# supabase/.gitignore       → exclusions Git
```
La commande `supabase init` crée un dossier `supabase/` à la racine du projet. Le fichier `config.toml` contient les ports utilisés par chaque service local – par défaut 54321 pour l’API, 54322 pour Postgres, 54323 pour Studio. Si l’un de ces ports est déjà occupé sur votre poste, modifiez-les dans le fichier de configuration avant de démarrer la stack.

## Étape 2 : Démarrer la stack Supabase locale en Docker

Avec Docker Desktop lancé en arrière-plan, exécutez `supabase start`. Le premier lancement télécharge environ 2 GB d’images Docker – comptez entre 3 et 8 minutes selon votre connexion. Les lancements suivants prennent moins de 30 secondes grâce au cache Docker.

```
supabase start
# Sortie attendue (extrait)
Started supabase local development setup.
         API URL: http://127.0.0.1:54321
     GraphQL URL: http://127.0.0.1:54321/graphql/v1
  S3 Storage URL: http://127.0.0.1:54321/storage/v1/s3
          DB URL: postgresql://postgres:[email protected]:54322/postgres
      Studio URL: http://127.0.0.1:54323
    Inbucket URL: http://127.0.0.1:54324
      JWT secret: super-secret-jwt-token-with-at-least-32-characters-long
        anon key: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
service_role key: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
Notez précieusement les clés `anon key` et `service_role key`. La première sera exposée côté navigateur – elle respecte les politiques RLS. La seconde donne un accès root à la base : elle ne doit JAMAIS apparaître dans le code client. Ouvrez http://127.0.0.1:54323 pour accéder à Supabase Studio en local : l’interface graphique pour gérer tables, politiques RLS, utilisateurs, fonctions et logs.

## Étape 3 : Concevoir le schéma PostgreSQL avec Row Level Security

Supabase repose sur PostgreSQL 16, qui supporte plus de 150 extensions parmi lesquelles pgvector pour l’IA, pg_cron pour les tâches planifiées, postgis pour la géolocalisation, et pg_trgm pour la recherche fuzzy. Pour notre application de tickets, nous avons besoin de deux tables – `tickets` et `messages` – liées par une clé étrangère, avec Row Level Security activée pour garantir qu’un utilisateur ne voit que ses propres données.

Créez votre première migration via la CLI :

