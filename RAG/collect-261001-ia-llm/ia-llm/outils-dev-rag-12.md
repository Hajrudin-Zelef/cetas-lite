---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-12
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei", "OpenAI"]
dates: []
keywords: ["decode", "embedding", "embeddings", "valuation"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2209, 2405]
sha256: c6f08317396edd33843c8bbd06760ae232b7a7c7b0780190270ab66b570e7d91
---

# Outils dev + ingénierie RAG (chunk & corpus)

- [ ] Extension `vector` activée, table `rag_chunks` créée (§64)
- [ ] Index HNSW créé **après** le chargement initial (§65)
- [ ] `ANALYZE` après chaque gros chargement
- [ ] Fonction `match_chunks` déployée et testée (§67)
- [ ] Secrets (`SUPABASE_URL`, clés) en variables d'environnement
- [ ] `service_role` uniquement côté serveur
- [ ] Stratégie Free (proto) vs Pro/self-hosted décidée (§71)

## 73. Pense-bête Supabase

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE INDEX ON rag_chunks USING hnsw (embedding vector_cosine_ops);
SELECT id, 1 - (embedding <=> :q) AS s FROM rag_chunks
ORDER BY embedding <=> :q LIMIT 5;
SET hnsw.ef_search = 100;   -- plus de rappel à la requête
```

```bash
supabase start          # stack locale
supabase db push        # pousse les migrations
```

---

# PARTIE B — INGÉNIERIE RAG

## 74. Pourquoi le chunking est la décision la plus impactante

Le chunking = découper tes documents en morceaux (« chunks ») avant de les
embedder. C'est **le paramètre le plus impactant** de tout le pipeline
(consensus des évaluations 2025-2026) :

```text
Chunks trop petits (64-128 tokens)  → retrieval précis MAIS le LLM reçoit
                                      des miettes sans contexte → hallucinations
Chunks trop gros (2000+ tokens)     → l'embedding moyenne trop de concepts
                                      → le chunk matche tout et rien
Sweet spot (vérifié)                → 256-1024 tokens, 512 = point de départ
                                      le plus courant en production
```

Ton embedding `text-embedding-3-small` accepte jusqu'à 8191 tokens en
entrée, mais sa **zone optimale** est ~256-512 tokens : assez de signal
sémantique, pas de dilution.

## 75. Vocabulaire du chunking

| Terme | Définition | Exemple |
|---|---|---|
| Chunk | Morceau de texte embeddé comme une unité | une section `##` |
| Token | Unité du modèle (~4 caractères en anglais, ~2.5 en code) | `tiktoken` pour compter |
| Fenêtre | Taille du chunk en tokens | 512 |
| Overlap | Recouvrement entre chunks consécutifs | 50 tokens (10 %) |
| Stride | Pas d'avancement = fenêtre − overlap | 462 |
| Orphelin | Chunk sans contexte (titre seul, phrase coupée) | à éviter |

## 76. Compter les tokens pour de vrai (tiktoken)

**Ne compte jamais en caractères** : 1000 caractères = ~250 tokens de prose
anglaise mais ~500 tokens de code (vérifié : ratio chars/token — prose ~4.0,
prose technique dense ~3.5, code ~2.5, JSON minifié ~2.0). Un chunk calibré
« 1000 caractères » peut dépasser la limite du modèle **en silence**
(troncature invisible à l'embedding).

```bash
pip install tiktoken
```

```python
"""Compter les tokens comme OpenAI les compte."""
import tiktoken

# text-embedding-3-small utilise l'encodage cl100k_base (vérifié)
enc = tiktoken.encoding_for_model("text-embedding-3-small")
# ou : tiktoken.get_encoding("cl100k_base")

def count_tokens(text: str) -> int:
    return len(enc.encode(text))

def truncate_tokens(text: str, max_tokens: int) -> str:
    toks = enc.encode(text)
    return enc.decode(toks[:max_tokens])

# Repères sur TON corpus (ordres de grandeur) :
#   guide .md de 2500 lignes ≈ 90-120k tokens → ~200 chunks de 500 tokens
#   1 commande CLI Huawei nettoyée ≈ 80-150 tokens → ~4-6 commandes par chunk de 500
```

> **Règle :** tout chunker qui décide en *caractères* doit convertir via
> `tiktoken` avant l'embedding, ou tronquer silencieusement.

## 77. Stratégie 1 : chunking fixe (taille/overlap)

Le plus simple : fenêtre de N tokens, pas de N−overlap.

**Comment choisir 500 vs 1000 vs 2000 (chiffres vérifiés 2025-2026) :**

| Taille | Précision retrieval | Contexte pour le LLM | Coût | Usage recommandé |
|---|---|---|---|---|
| 128-256 | ★★★ (factoid) | ★ (miettes) | bas | FAQ, questions factuelles |
| **512** | ★★ | ★★ | moyen | **défaut universel** (docs techniques) |
| 1000-1024 | ★★ | ★★★ | moyen+ | raisonnement multi-hop, procédures longues |
| 2000+ | ★ (dilution) | ★★★ | haut | résumés, rarement pour le retrieval |

Recommandations par cas (sources 2026) :
- **Documentation technique** (ton cas) : **256-512 tokens, overlap 15 %**.
- **QA factoid** : 64-128 tokens, overlap 10 %.
- **Contrats/juridique** : pattern parent-enfant (voir §81).
- **Point de départ universel** : **512 tokens, overlap 50** (10 %).

**Coût de l'overlap** : avec 10 % d'overlap, tu embeddes ~11 % de tokens
en plus. À $0.02/1M tokens, c'est négligeable pour ton volume (§71).

## 78. Stratégie 2 : chunking par sections markdown (## — RECOMMANDÉ pour tes guides)

Tes guides sont structurés en `## N. Titre`. **Ne casse pas cette
structure** : 1 section = 1 chunk (ou N chunks si la section est énorme).
Avantages :
- Chaque chunk a un **titre** → métadonnée `section` gratuite (§84).
- Jamais de phrase coupée au milieu d'une idée.
- Le retrieval ramène des unités **citables** (« voir § Configuration iStack »).

Règle de découpe :
```text
## Section de 300 tokens  → 1 chunk
## Section de 1500 tokens → découper en sous-chunks de ~500 (fixe, voir §79)
                             en gardant le titre de section dans chaque chunk
# Les blocs ```code``` et les tableaux |...| ne sont JAMAIS coupés (voir §82)
```

## 79. Stratégie 3 : chunking récursif (le bon défaut générique)

**RecursiveCharacterTextSplitter** (LangChain, vérifié) : essaie de couper
sur `\n\n` (paragraphes), puis `\n`, puis phrases, puis mots — en
préservant au mieux les frontières naturelles.

```bash
pip install langchain-text-splitters
```

```python
from langchain_text_splitters import RecursiveCharacterTextSplitter

# ⚠ chunk_size est en CARACTÈRES ici, pas en tokens (voir §76) !
splitter = RecursiveCharacterTextSplitter(
    chunk_size=2000,        # ≈ 500 tokens de prose
    chunk_overlap=200,      # ≈ 50 tokens (10 %)
    separators=["\n\n", "\n", ". ", " ", ""],
)
chunks = splitter.split_text(open("guide.md").read())
```

> Pour tes guides : **sections markdown d'abord** (§78), récursif en
> **secours** pour les sections trop longues ou les articles web sans
> structure.

## 80. Stratégie 4 : chunking sémantique (principe et coûts)

Principe : embedder chaque phrase, regrouper les phrases **sémantiquement
proches** (rupture quand la similarité cosinus chute). Les chunks suivent
les changements de sujet, pas des tailles fixes.

```python
# Pseudo-code (langchain expérimental — à vérifier : API exacte en 2026)
from langchain_experimental.text_splitter import SemanticChunker
from langchain_openai import OpenAIEmbeddings

splitter = SemanticChunker(
    OpenAIEmbeddings(model="text-embedding-3-small"),
    breakpoint_threshold_type="percentile",  # rupture à 95e percentile
)
```

**Coûts (vérifié 2026) :** 3 à 5x plus cher (embeddings des phrases +
des chunks), gain **minime** sur textes courts/structurés. **Verdict :**
à tester **uniquement** si tes sections markdown sont très hétérogènes
et que ton jeu d'évaluation (§101) montre un problème. Pas le défaut.

## 81. Stratégie 5 : par proposition + pattern parent-enfant

**Chunking par proposition** : 1 phrase autonome = 1 chunk (idéal pour
le factoid : « L'AP361 consomme 8,8 W en PoE 802.3af »). Retrieval
ultra-précis, mais contexte pauvre → à combiner avec :

**Parent-enfant (small-to-big)** — le pattern « précision + contexte » :
```text
Enfant  (128 tokens)  → indexé et recherché (précision maximale)
Parent  (1024 tokens = la section entière) → renvoyé au LLM (contexte riche)
```

```sql
-- Schéma : le parent est stocké une fois, les enfants pointent vers lui
ALTER TABLE rag_chunks ADD COLUMN parent_id BIGINT REFERENCES rag_chunks(id);
-- Recherche : ORDER BY enfant.embedding <=> :q, puis JOIN pour remonter le parent
```

