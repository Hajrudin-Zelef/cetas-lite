---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-10
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["embedding", "embeddings", "open source"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1836, 2013]
sha256: c4d1a9f9fc1c1258b9b02d43b488acda98f75050796e4f1bc7c03e09abebbbc1
---

# Outils dev + ingénierie RAG (chunk & corpus)

## 59. Pense-bête Puppeteer

```bash
npm install puppeteer
node scrape.js
# Pin la version dans package.json : le Chromium bundlé change sinon.
```

---

## 60. « superbase » : vérification → c'est **Supabase**

Vérifié par recherche web (sept 2026) : **Supabase** (supabase.com) existe
bel et bien — pas « Superbase ». Fondée en 2020 par Paul Copplestone et
Ant Wilson, c'est l'« alternative open source à Firebase » : **un vrai
PostgreSQL managé** + auth + stockage + edge functions + realtime + **
pgvector** (recherche vectorielle). Levée de $500M en juin 2026
(valorisation $10,5 Md), ~10 millions de développeurs revendiqués.

**Pourquoi c'est pertinent pour TON RAG :** tes embeddings
`text-embedding-3-small` (1536 dimensions) vivent dans une colonne
`vector(1536)` **à côté** de tes métadonnées (source, section, page) dans
le même Postgres. Une seule base, du SQL standard, requêtes hybrides
(vectoriel + filtres relationnels) en une requête.

## 61. L'écosystème Supabase en 30 secondes

| Composant | Rôle | Utile pour ton RAG ? |
|---|---|---|
| **Database** (Postgres + pgvector) | Stockage chunks + embeddings | ✅ cœur du système |
| **Auth** | Utilisateurs, OAuth, MFA | plus tard (si tu exposes l'app) |
| **Storage** | Fichiers (S3-compatible) | pour les PDF/exports bruts |
| **Edge Functions** (Deno) | Serverless | pour une API de recherche |
| **Realtime** | Websocket sur changements DB | non prioritaire |
| **CLI** | Dev local via Docker | ✅ (`supabase start`) |

Le code est open source : tu peux **self-hoster** tout ça sur ton serveur
si tu veux quitter le cloud un jour (même Postgres, mêmes requêtes).

## 62. Tarifs 2026 (vérifié sept 2026)

| Plan | Prix | Base de données | À noter |
|---|---|---|---|
| **Free** | $0 | 500 Mo, CPU partagé, 2 projets max | ⚠ **pause après 7 jours d'inactivité** |
| **Pro** | **$25/mois** | 8 Go inclus ($0.125/Go au-delà), 100K MAU | backups quotidiens, support email |
| **Team** | $599/mois | idem + conformité | SOC 2, HIPAA (add-on), SSO |

- **pgvector est inclus dans tous les plans, sans surcoût** (vérifié).
- Le Free suffit pour **prototyper** ; la pause après 7 jours d'inactivité
  le disqualifie pour un usage régulier → Pro à $25/mois.
- **Dimensionnement pour toi (§72)** : ton corpus actuel tient très
  largement dans le Free en volume, mais la pause auto est le vrai sujet.

## 63. Prise en main : CLI + projet local

```bash
# 1. Installer la CLI (Docker requis pour le mode local)
brew install supabase/tap/supabase   # macOS
# Linux : voir docs supabase.com/docs/guides/local-development — à vérifier
# la méthode d'install exacte pour ta distro.

# 2. Initialiser dans ton repo
cd ~/rag-perso
supabase init

# 3. Démarrer la stack locale (Postgres + Auth + Storage + Studio)
supabase start
# → Studio (UI web) : http://127.0.0.1:54323  (à vérifier : port exact)
# → API : http://127.0.0.1:54321              (à vérifier : port exact)

# 4. Lier un projet cloud existant
supabase link --project-ref <ref-de-ton-projet>

# 5. Pousser tes migrations SQL (dossier supabase/migrations/)
supabase db push
```

> **Développement local d'abord** : `supabase start` te donne un Postgres
> avec pgvector **gratuit et offline**. Tu ne touches au cloud que quand
> ça marche.

## 64. Activer pgvector et créer le schéma (SQL complet)

Dans l'éditeur SQL du Studio, ou `psql`, ou `supabase/migrations/` :

```sql
-- ═══════════════════════════════════════════════════════════
-- 001_schema.sql — extension + table des chunks
-- ═══════════════════════════════════════════════════════════

-- 1. Activer l'extension (disponible par défaut sur Supabase — vérifié)
CREATE EXTENSION IF NOT EXISTS vector;

-- 2. Table des chunks : texte + embedding + métadonnées CRITIQUES
CREATE TABLE IF NOT EXISTS rag_chunks (
    id          BIGSERIAL PRIMARY KEY,
    -- Identité du chunk
    chunk_hash  CHAR(64) NOT NULL UNIQUE,   -- sha256 du texte (dédup exacte)
    -- Contenu
    content     TEXT NOT NULL,
    -- Embedding : 1536 = text-embedding-3-small (vérifié)
    embedding   vector(1536),
    -- Métadonnées de filtrage (voir §84 : CRITIQUE pour filtrer)
    source      TEXT NOT NULL,              -- ex. 'huawei_cli_ref', 'guide_debian'
    doc_title   TEXT,                       -- ex. 'NetEngine AR1000V Command Reference'
    section     TEXT,                       -- ex. '## Configuration iStack'
    page        INT,                        -- page du PDF d'origine (si connue)
    lang        CHAR(2) DEFAULT 'fr',
    chunk_index INT NOT NULL,               -- position dans le doc
    -- Traçabilité
    created_at  TIMESTAMPTZ DEFAULT now(),
    corpus_version TEXT NOT NULL DEFAULT 'v1'
);

-- 3. Index B-tree classiques pour les filtres (aussi importants que l'index vectoriel !)
CREATE INDEX IF NOT EXISTS idx_chunks_source  ON rag_chunks (source);
CREATE INDEX IF NOT EXISTS idx_chunks_section ON rag_chunks (section);
CREATE INDEX IF NOT EXISTS idx_chunks_hash    ON rag_chunks (chunk_hash);
```

> **Types pgvector (vérifié, pgvector ≥ 0.7) :**
> - `vector(n)` : float32, le standard.
> - `halfvec(n)` : float16, **2x moins de stockage**, perte de précision mineure.
> - `sparsevec(n)` : vecteurs creux (BM25/TF-IDF).

## 65. Index HNSW : la recherche rapide (SQL complet)

**HNSW** (Hierarchical Navigable Small World) = l'index recommandé pour la
plupart des usages : requêtes en ~O(log n), excellent rappel.

```sql
-- ═══════════════════════════════════════════════════════════
-- 002_indexes.sql — index vectoriels
-- ═══════════════════════════════════════════════════════════

-- HNSW + distance cosinus (le bon choix pour text-embedding-3-small,
-- dont les vecteurs sont quasi-normalisés — voir §67)
CREATE INDEX IF NOT EXISTS idx_chunks_embedding_hnsw
    ON rag_chunks USING hnsw (embedding vector_cosine_ops)
    WITH (m = 16, ef_construction = 64);

-- Après un GROS chargement, toujours :
ANALYZE rag_chunks;
```

Paramètres :
- `m = 16` : connexions par nœud (défaut raisonnable ; 24-32 pour un
  rappel maximal au prix de plus de RAM).
- `ef_construction = 64` : qualité de construction (plus = meilleur rappel,
  construction plus lente).
- **Réglage à la requête** (sans reconstruire l'index) :
  `SET hnsw.ef_search = 100;` (défaut 40 — monte à 100-200 pour un rappel
  maximal sur tes évaluations).

> **Crée l'index APRÈS le chargement initial** : construire HNSW sur une
> table vide puis insérer 100k lignes est beaucoup plus lent que l'inverse.

## 66. Index IVFFlat : l'alternative économe (SQL complet)

**IVFFlat** = partitionne l'espace en `lists` cellules (inverted file).
Moins gourmand en RAM, construction plus rapide, rappel légèrement inférieur.

```sql
-- IVFFlat + distance cosinus
CREATE INDEX IF NOT EXISTS idx_chunks_embedding_ivfflat
    ON rag_chunks USING ivfflat (embedding vector_cosine_ops)
    WITH (lists = 100);

ANALYZE rag_chunks;
```

- `lists` ≈ **√(nombre de lignes)** : 10k lignes → 100 ; 100k → 316 ;
  1M → 1000 (règle vérifiée dans plusieurs guides 2026).
- Réglage à la requête : `SET ivfflat.probes = 10;` (nombre de cellules
  visitées — plus = meilleur rappel, plus lent).
- **N'indexe en IVFFlat qu'avec > quelques milliers de lignes** : en
  dessous, un scan séquentiel est aussi rapide.

