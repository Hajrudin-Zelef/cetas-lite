---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-11
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["benchmarks", "embedding"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2014, 2208]
sha256: 6e10c6854aaed47905e30dc551c601a18392cb4031a71bc0b169264ee8731913
---

# Outils dev + ingénierie RAG (chunk & corpus)

| | HNSW | IVFFlat |
|---|---|---|
| Rappel | excellent | bon |
| Vitesse requête | très rapide | rapide |
| RAM | plus gourmande | économe |
| Construction | plus lente | rapide |
| **Choix pour toi** | ✅ **défaut** (< 5M vecteurs) | si RAM limitée |

## 67. Requêtes de similarité cosinus (SQL complet)

Opérateurs pgvector (vérifié) : `<->` euclidien, **`<=>` cosinus**,
`<#>` produit intérieur (négatif).

```sql
-- ── Top-5 des chunks les plus proches (distance cosinus) ────
-- :q est ton embedding de question (vector(1536)), passé en paramètre.
SELECT id, source, section,
       content,
       1 - (embedding <=> :q) AS similarite   -- cosinus similarité ∈ [-1, 1]
FROM rag_chunks
WHERE source = 'huawei_cli_ref'               -- filtre relationnel AVANT le tri
ORDER BY embedding <=> :q                     -- l'index HNSW sert ICI
LIMIT 5;

-- ── Avec seuil de similarité ────────────────────────────────
SELECT id, source, section, content,
       1 - (embedding <=> :q) AS similarite
FROM rag_chunks
WHERE 1 - (embedding <=> :q) > 0.75           -- seuil : à calibrer (§101)
ORDER BY embedding <=> :q
LIMIT 10;

-- ── Hybride : vectoriel + plein-texte (tsvector) ────────────
-- (nécessite une colonne search_tsv + index GIN — voir §101)
SELECT id, source, section, content,
       1 - (embedding <=> :q) AS sim_vec,
       ts_rank(search_tsv, plainto_tsquery('french', 'iStack S310')) AS rank_txt
FROM rag_chunks
WHERE search_tsv @@ plainto_tsquery('french', 'iStack S310')
ORDER BY embedding <=> :q
LIMIT 10;
```

**Fonction RPC réutilisable** (appelable depuis Python/JS via l'API
Supabase — pattern vérifié 2026) :

```sql
CREATE OR REPLACE FUNCTION match_chunks(
    query_embedding vector(1536),
    match_threshold float DEFAULT 0.70,
    match_count int DEFAULT 10,
    p_source text DEFAULT NULL
)
RETURNS TABLE (
    id bigint, content text, source text,
    section text, page int, similarite float
)
LANGUAGE sql STABLE
AS $$
    SELECT c.id, c.content, c.source, c.section, c.page,
           1 - (c.embedding <=> query_embedding)
    FROM rag_chunks c
    WHERE (p_source IS NULL OR c.source = p_source)
      AND 1 - (c.embedding <=> query_embedding) > match_threshold
    ORDER BY c.embedding <=> query_embedding
    LIMIT match_count;
$$;
```

## 68. Client Python : insérer et interroger

```bash
pip install supabase  # client officiel Python (vérifié : existe)
```

```python
"""Ingestion + requête pgvector via le client Supabase."""
import os
from supabase import create_client

SUPABASE_URL = os.environ["SUPABASE_URL"]
SUPABASE_KEY = os.environ["SUPABASE_KEY"]   # clé "service_role" côté serveur UNIQUEMENT
sb = create_client(SUPABASE_URL, SUPABASE_KEY)

# ── Insertion d'un chunk ─────────────────────────────────────
def insert_chunk(chunk: dict, embedding: list[float]) -> None:
    sb.table("rag_chunks").insert({
        "chunk_hash": chunk["hash"],
        "content": chunk["text"],
        "embedding": embedding,          # le client sérialise la liste → vector
        "source": chunk["metadata"]["source"],
        "doc_title": chunk["metadata"].get("doc_title"),
        "section": chunk["metadata"].get("section"),
        "page": chunk["metadata"].get("page"),
        "chunk_index": chunk["metadata"]["chunk_index"],
        "corpus_version": "v1",
    }).execute()

# ── Insertion en batch (beaucoup plus rapide) ────────────────
def insert_batch(rows: list[dict]) -> None:
    # rows = liste de dicts comme ci-dessus, par paquets de 100-500
    sb.table("rag_chunks").insert(rows).execute()

# ── Requête via la fonction RPC match_chunks (§67) ───────────
def search(query_embedding: list[float], source: str | None = None, k: int = 10):
    res = sb.rpc("match_chunks", {
        "query_embedding": query_embedding,
        "match_threshold": 0.70,
        "match_count": k,
        "p_source": source,
    }).execute()
    return res.data

# ── Connexion directe psycopg (scripts lourds) ───────────────
# Pour l'ingestion massive, psycopg3 + COPY est plus rapide que le client.
# Port 6543 = pooler (recommandé pour les scripts — vérifié 2026),
# port 5432 = connexion directe.
```

> **Clés** : `anon` (publique, avec RLS) vs `service_role` (tous droits).
> Tes scripts d'ingestion tournent avec `service_role` **côté serveur
> uniquement**, jamais dans du code distribué.

## 69. Stockage et Auth (bref)

- **Storage** : buckets S3-compatibles pour tes PDF/exports bruts
  (`supabase.storage.from_("corpus-brut").upload(...)`). Politiques d'accès
  via RLS, URLs signées temporaires (`createSignedUrl`).
- **Auth** : inscription/connexion, OAuth (20+ providers), MFA. Chaque
  utilisateur a un `auth.uid()` utilisable dans les politiques RLS.
- **RLS (Row Level Security)** — le jour où tu exposes ton RAG :

```sql
ALTER TABLE rag_chunks ENABLE ROW LEVEL SECURITY;
-- Exemple : chacun ne voit que sa collection
CREATE POLICY "isolation par utilisateur"
    ON rag_chunks FOR SELECT
    USING (auth.uid() = user_id);   -- nécessite une colonne user_id
```

Pour ton usage perso local, RLS est optionnel ; pour une app partagée,
c'est obligatoire.

## 70. Pourquoi Supabase plutôt que… (tableau)

| | Supabase + pgvector | pgvector self-hosted | Pinecone/Qdrant cloud |
|---|---|---|---|
| Opérations | zéro | à toi (backups, updates) | zéro |
| SQL relationnel | ✅ vrai Postgres | ✅ | ❌ / limité |
| Filtres métadonnées | ✅ `WHERE` natif | ✅ | ✅ (selon produit) |
| Coût (ton échelle) | $0-25/mois | coût serveur existant | $25-70/mois |
| Échelle | < ~50M vecteurs/noeud (vérifié) | idem | 100M+ |
| Latence p99 | ~ms (OK < 50 req/s — vérifié) | idem | meilleure à très haute charge |
| Portabilité | ✅ Postgres standard | ✅ | ❌ propriétaire |

**Verdict pour toi :** Supabase (ou pgvector self-hosté sur ton serveur
si tu veux zéro cloud) couvre 100 % de ton besoin. Un vector DB dédié ne
se justifie qu'au-delà de ~5-50M de vecteurs.

## 71. Dimensionnement pour TON RAG (calculs)

Hypothèses : `text-embedding-3-small` = 1536 dims × 4 octets (float32).

```text
1 vecteur  = 1536 × 4 o = 6 144 o ≈ 6 Ko
+ texte du chunk (~2 Ko) + métadonnées + overhead index HNSW (~1.5-2x)

Estimation de TON corpus actuel (ordres de grandeur) :
  Guides (~30 guides × ~2500 lignes ≈ 75k lignes ≈ 600k tokens)
    → ~1 200 chunks de 500 tokens
  CLI Huawei nettoyés (12 620 commandes ≈ 2M tokens)
    → ~4 000 chunks
  Articles web (~300 articles ≈ 1.5M tokens)
    → ~3 000 chunks
  ─────────────────────────────────────────
  TOTAL ≈ 8 000-10 000 chunks ≈ 10 000 × 6 Ko ≈ 60 Mo de vecteurs
          + textes + index ≈ 150-250 Mo au total

→ Tient dans le plan Free (500 Mo) MAIS la pause à 7 jours d'inactivité
  rend le Free inutilisable en pratique → Pro $25/mois si usage régulier,
  ou pgvector self-hosté sur ton serveur (coût = 0 € marginal).
```

**Coût d'embedding** (vérifié : $0.02 / 1M tokens) :
`~4M tokens × $0.02/1M ≈ $0.08` pour tout ton corpus. Le re-embed complet
coûte moins de 10 centimes — n'hésite pas à re-embedder quand ton chunking
change (§102).

**Astuce Matryoshka** (vérifié 2026) : `text-embedding-3-small` supporte
`dimensions=512` à l'appel → vecteurs 3x plus petits, ~95 % du rappel
conservé selon les benchmarks. À tester sur TON jeu de questions (§101)
avant d'adopter.

## 72. Checklist Supabase

