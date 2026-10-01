---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-5
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "embedding", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [644, 751]
sha256: 2f4ec81419a980897368ecf943e16ee37af806b54c1c06fe5e3af81d499bad65
---

# macOS (Homebrew)

```
cd web
npm install -g vercel
vercel --prod
# Configurez les variables d'environnement dans le dashboard Vercel :
# NEXT_PUBLIC_SUPABASE_URL
# NEXT_PUBLIC_SUPABASE_ANON_KEY
# SUPABASE_SERVICE_ROLE_KEY (pour usages server-only)
```
Pensez à ajouter votre domaine Vercel à la liste des **Redirect URLs** autorisées dans *Authentication → URL Configuration* du dashboard Supabase, sous peine d’échec de tous les callbacks OAuth en production.

## Étape 12 : Branching pour les pull requests et workflows CI/CD

Le **Branching** Supabase, désormais inclus sur tous les plans depuis mai 2026, crée automatiquement une base PostgreSQL éphémère pour chaque pull request GitHub. Les migrations s’appliquent à la branche, les tests d’intégration tournent sur des données isolées, et la branche est détruite à la fusion ou à la fermeture de la PR.

Activez le Branching dans le dashboard, puis ajoutez ce workflow GitHub Actions :

```
# .github/workflows/supabase.yml
name: Supabase CI
on:
  pull_request:
    paths: ['supabase/**', 'web/**']
  push:
    branches: [main]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: supabase/setup-cli@v1
        with:
          version: 2.6.8
      - run: supabase db lint
      - run: supabase test db
  deploy:
    needs: test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: supabase/setup-cli@v1
      - run: supabase link --project-ref ${{ secrets.SUPABASE_PROJECT_REF }}
      - run: supabase db push
      - run: supabase functions deploy --project-ref ${{ secrets.SUPABASE_PROJECT_REF }}
        env:
          SUPABASE_ACCESS_TOKEN: ${{ secrets.SUPABASE_ACCESS_TOKEN }}
```
Ce pipeline garantit qu’aucune migration cassée n’atteint la production : `supabase db lint` détecte les SQL non-portables, `supabase test db` exécute pgTAP sur les fonctions stockées et politiques RLS.

## Étape 13 : Observabilité, logs et performance en production

Le dashboard Supabase intègre, depuis la généralisation de l’éditeur SQL inline dans Supabase Studio entre le 10 et le 24 février 2025, des logs détaillés inspectables avec un panneau latéral, des rapports SQL personnalisables via SQL Snippets, et un AI Assistant capable de réécrire les requêtes lentes. La vue **Reports** agrège les métriques par endpoint API, latence p95, taille moyenne des réponses et nombre d’erreurs 4xx/5xx – l’équivalent d’un Datadog basique sans coût supplémentaire.

Pour le monitoring approfondi, exportez les logs vers Logflare, Datadog ou Grafana Loki via le connecteur natif. La métrique critique à surveiller reste le **compute utilization** de l’instance Postgres – au-delà de 80 % soutenu, planifiez le passage à un compute add-on (Small : 25 $/mois, Large : 110 $/mois, 16XL : 3 730 $/mois). Le tier Free utilise des shared compute units : leur limite stricte explique l’erreur `compute usage exceeded` que rencontrent les projets en croissance trop rapide.

| Plan | Prix mensuel | Database | Storage | MAU inclus | Edge Functions | Realtime connexions | 
|---|---|---|---|---|---|---|
| Free | 0 $ | 500 MB | 1 GB | 50 000 | 500 000 invocations | 200 | 
| Pro | 25 $/projet | 8 GB | 100 GB | 100 000 (puis 0,00325 $/MAU) | 2 millions invocations | 500 | 
| Team | 599 $/projet | 100 GB+ (custom) | 500 GB+ | 100 000 (puis 0,00325 $/MAU) | 2 millions invocations + SSO | 10 000+ | 
| Enterprise | Sur devis | Illimitée | Illimité | Illimité | Custom | Custom | 

## Pièges courants à éviter avec Supabase en production

Au-delà du tutoriel pas-à-pas, voici les erreurs les plus fréquentes que rencontrent les équipes françaises lors des premiers déploiements Supabase. Chacune a été observée plusieurs fois sur des projets clients en 2025-2026.

- **Exposer la service_role key côté client** : cette clé contourne RLS. Elle ne doit JAMAIS apparaître dans`NEXT_PUBLIC_*` , dans le code envoyé au navigateur, ni dans un dépôt public. Utilisez-la uniquement dans des Server Actions, Edge Functions ou jobs cron côté serveur.
- **Oublier d’activer RLS sur une table** : par défaut, sans RLS, la clé anonyme donne accès complet à la table – fuite de données garantie. Activez systématiquement`alter table … enable row level security` et créez au moins une policy.
- **Confondre `auth.users` et table publique** : ne stockez jamais de profil utilisateur directement dans`auth.users` . Créez une table`public.profiles` avec une clé étrangère vers`auth.users(id)` et un trigger sur`auth.users` pour la peupler automatiquement.
- **Migrations non versionnées** : modifier le schéma directement dans Studio ne génère pas de migration. Utilisez`supabase db diff` pour capturer les changements ou éditez systématiquement les fichiers SQL.
- **Connexions Postgres saturées** : le tier Free limite à 60 connexions directes. Pour Next.js Serverless, utilisez le pooler PgBouncer sur le port 6543 (mode Transaction), pas le port 5432 direct.
- **Realtime sans publication activée** : oublier`alter publication supabase_realtime add table …` donne un canal silencieux qui ne reçoit jamais d’événement.
- **Bucket Storage public par erreur** : un bucket marqué`public: true` expose tous les fichiers via URL devinable. Préférez`public: false` + URL signées.
- **Magic Link rate limit** : Supabase bloque à 4 emails/heure par IP en local. En production, configurez impérativement un SMTP externe.

## Dépannage : 8 erreurs Supabase et comment les résoudre

| Erreur | Cause probable | Solution | 
|---|---|---|
| `Cannot read properties of undefined (reading 'cookies')` | Mauvais client Supabase (browser dans Server Component) | Importer `createClient` depuis`@/lib/supabase/server` | 
| `JWT expired` | Middleware non configuré ou non exécuté sur la route | Vérifier `matcher` dans`middleware.ts` | 
| `new row violates row-level security policy` | Policy INSERT manquante ou condition incorrecte | Ajouter `create policy … for insert with check (auth.uid() = user_id)` | 
| `relation "tickets" does not exist` | Migration non appliquée sur l’environnement courant | `supabase db reset` en local ou`supabase db push` en prod | 
| `Auth session missing!` | Cookies non transmis (CORS ou domaine cookie incorrect) | Vérifier les Redirect URLs dans le dashboard Supabase | 
| `Connection refused on port 54322` | Stack Supabase non démarrée | Lancer `supabase start` et attendre la fin du pull | 
| `Storage bucket not found` | Bucket non créé sur l’environnement cible | Appliquer la migration storage ou créer le bucket via Studio | 
| `Failed to fetch realtime channel` | WebSocket bloqué par firewall ou table non publiée | Vérifier publication + autoriser `wss://` sur le proxy | 

## Astuces avancées pour passer à l’échelle

### Optimiser les requêtes avec PostgREST embedding

L’API REST auto-générée par PostgREST, épinglé en version 14.5 dans les images de production Supabase depuis la mise à jour d’auto-hébergement de février 2026 (déployée mondialement depuis décembre 2025), supporte le *resource embedding* : récupérez tickets et messages associés en une seule requête HTTP au lieu de N+1.

```
const { data } = await supabase
  .from('tickets')
  .select('id, title, status, messages(id, body, created_at)')
  .eq('id', ticketId)
  .single()
```
### Cache HTTP avec révalidation

Sur Next.js 16, ajoutez `export const revalidate = 60` en tête d’un Server Component pour mettre en cache les réponses Supabase pendant 60 secondes au CDN – diminution massive des coûts compute.

### pgvector pour la recherche sémantique IA

