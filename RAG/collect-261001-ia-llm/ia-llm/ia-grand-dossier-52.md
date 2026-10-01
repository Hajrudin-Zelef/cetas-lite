---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-52
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cohere", "Hugging Face", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["awq", "benchmark", "cohere", "embedding", "embeddings", "fp8", "gguf", "gpu", "inference", "llama", "llama.cpp", "lora"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [4064, 4189]
sha256: dfd8663cad6a5850392c88b006a789464a33e02e9b4c0400f65e47bb2c9e1a38
---

# IA — Le grand dossier

| Moteur | Cas d'usage | Points forts | Limites |
|---|---|---|---|
| **vLLM** | Prod GPU (1 à N GPU) | PagedAttention (KV-cache paginé → +de requêtes concurrentes), API OpenAI-compatible, LoRA multi-adaptateurs à chaud | GPU NVIDIA (CUDA) quasi requis |
| **TGI** (Text Generation Inference, Hugging Face) | Prod, intégration HF | Simple, métriques Prometheus natives, tensor-parallel | Moins flexible que vLLM sur les LoRA |
| **TensorRT-LLM** (NVIDIA) | Perf max sur NVIDIA | Kernels optimisés, FP8 | Complexe à régler, lié à NVIDIA |
| **llama.cpp / Ollama** | Local, CPU, edge, Mac | GGUF quantifiés (Q4_K_M…), tourne sur CPU, **Ollama** = `docker run` du LLM | Débit limité vs GPU serveur |
| **SGLang** | Prod recherche, Structured outputs | Programmation « langage » des appels, constrained decoding | Écosystème plus jeune |
| **Hugging Face TEI** | Microservice d'embeddings | Batching d'embeddings, simple | Embeddings uniquement |
| **NIM** (NVIDIA) | Entreprise | Conteneurs optimisés + support | Licence/écosystème NVIDIA |

**Exemple réel — servir un modèle avec vLLM (OpenAI-compatible) :**

```bash
# 1. Lancement (1 GPU 24 Go, modèle 8B quantifié AWQ)
vllm serve Qwen/Qwen3-8B-AWQ \
  --port 8000 --max-model-len 32768 \
  --gpu-memory-utilization 0.9 \
  --enable-prefix-caching        # réutilise le KV-cache des prompts répétés (gros gain RAG)

# 2. Appel (même client que l'API OpenAI)
curl http://localhost:8000/v1/chat/completions -H "Content-Type: application/json" -d '{
  "model": "Qwen/Qwen3-8B-AWQ",
  "messages": [{"role":"user","content":"Résume ce log en 3 lignes : ..."}],
  "temperature": 0.2, "max_tokens": 300
}'
```

**Exemple réel — Ollama (le plus simple pour commencer en local) :**

```bash
ollama pull qwen3:8b
ollama run qwen3:8b "Explique la différence entre VRF et VLAN en 5 lignes"
# API : http://localhost:11434/api/generate  /  /api/chat  (format propre, documenté)
# Modèles d'embedding : ollama pull nomic-embed-text
```

**Dimensionnement (ordres de grandeur, FP16 sauf mention) :**

| Modèle | Poids FP16 | Q4 (≈) | VRAM d'inférence conseillée |
|---|---|---|---|
| 8B | ~16 Go | ~5 Go | 1× 24 Go confortable (contexte long OK) |
| 32B | ~64 Go | ~19 Go | 1× 48 Go (Q4) ou 2× 48 Go (FP16) |
| 70B | ~140 Go | ~40 Go | 2× 48-80 Go (Q4) ; 2-4× 80 Go (FP16) |

**Supervision** : expose les métriques (TGI et vLLM ont des endpoints Prometheus), alerte sur :
temps inter-token (p99), longueur des files, taux d'erreur, saturation KV-cache. Un LLM en prod
se supervise comme une base de données : latence, files, erreurs.

### 4.3. Stack RAG (ton pipeline — exemple chaîné complet)

**À quoi ça sert** : brancher un LLM sur *tes* documents (tes 25 000+ lignes de guides, tes
PDF constructeurs, tes tickets GLPI). C'est l'architecture de ton app personnelle.

```
┌───────── INGESTION (offline) ─────────┐      ┌───────── REQUÊTE (online) ─────────┐
│ Docs bruts                            │      │ Question utilisateur                │
│   │ (pdf, md, html)                   │      │   │                                 │
│   ▼                                   │      │   ▼                                 │
│ Nettoyage (dédup,                  │      │   │                                 │
│  en-têtes/pieds de page)              │      │   ▼                                 │
│   ▼                                   │      │ Embedding de la requête             │
│ Chunking (ex: 800 tokens,             │      │   │ (même modèle qu'à l'indexation !) │
│  overlap 120, split sémantique)       │      │   ▼                                 │
│   ▼                                   │      │ Recherche vectorielle (top-k=20)     │
│ Embeddings (ex: text-embedding-3-small │      │   │  (+ filtres métadonnées / ACL)    │
│  1536 dim, ou nomic-embed-text local) │      │   ▼                                 │
│   ▼                                   │      │ Rerank (cross-encoder, top-20 → 5)   │
│ Index vectoriel (pgvector / Qdrant /   │      │   │                                 │
│  Chroma / Milvus) + métadonnées       │      │   ▼                                 │
│  (source, date, ACL, langue)          │      │ Prompt : question + 5 chunks        │
└───────────────────────────────────────┘      │  (avec citations [source p.X])      │
                                               │   ▼                                 │
                                               │ Génération (temp. basse 0.1-0.3)    │
                                               │   ▼                                 │
                                               │ Réponse + sources citées           │
                                               └───────────────────────────────────┘
```

**Les 5 décisions qui font ou cassent un RAG** (retour d'expérience de l'industrie, 2024-2026) :

1. **Chunking** : ni trop gros (bruit) ni trop petits (contexte perdu). Règle de départ : 500-1000
   tokens, overlap 10-20 %, découpage *sémantique* (par section/titre) plutôt que mécanique.
   Pour tes guides : 1 chunk = 1 section numérotée, c'est idéal.
2. **Embeddings** : le modèle d'embedding compte plus que le modèle de génération pour la qualité
   de recherche. Ton choix (text-embedding-3-small, 1536 dim) est solide et économique. En local :
   `nomic-embed-text` ou `bge-m3` (multilingue — pertinent pour ton corpus FR/EN).
3. **Hybride > pur vectoriel** : BM25 (mots-clés) + vectoriel (sémantique) fusionnés (reciprocal
   rank fusion) battent le vectoriel seul sur les termes techniques exacts (« U163 », « C6000 » —
   typiquement tes codes Kyocera !). Indispensable pour un corpus technique.
4. **Rerank** : un cross-encoder (ex. : `bge-reranker`, Cohere Rerank, ou mxbai — « à vérifier »
   sur les versions 2026) relit les top-20 et ne garde que les 5 meilleurs. Gain de qualité massif
   pour un coût faible.
5. **Évaluation** : mesure avec un jeu de questions/réponses de référence (ton propre benchmark :
   50 questions dont tu connais la réponse, score de rappel@k et de fidélité). Sans éval, tu
   pilotes à l'aveugle.

**Exemple chaîné complet et réel (Python, pgvector + sentence-transformers + BM25 hybride) :**

```python
# === 1. INGESTION ===
import re, psycopg2
from sentence_transformers import SentenceTransformer

def chunk_markdown(path, max_tokens=700, overlap=120):
    """Découpe un .md par sections (# ## ###), avec chevauchement."""
    text = open(path, encoding="utf-8").read()
    parts = re.split(r"(?m)^(#{1,3} .+)$", text)  # garde les titres comme délimiteurs
    chunks, buf = [], ""
    for p in parts:
        if len((buf + p).split()) < max_tokens:
            buf += p
        else:
            chunks.append(buf); buf = p[-overlap*4:] + p  # overlap approx en chars
    if buf: chunks.append(buf)
    return [{"text": c, "source": path} for c in chunks if len(c.split()) > 30]

model = SentenceTransformer("nomic-ai/nomic-embed-text-v1.5", trust_remote_code=True)
conn = psycopg2.connect("dbname=rag user=zelef")  # + extension pgvector + tsvector FR
cur = conn.cursor()
cur.execute("""CREATE TABLE IF NOT EXISTS docs(
  id serial PRIMARY KEY, source text, chunk text,
  embedding vector(768),                      -- dim de nomic-embed-text
  tsv tsvector GENERATED ALWAYS AS (to_tsvector('french', chunk)) STORED)""")
cur.execute("CREATE INDEX IF NOT EXISTS ON docs USING hnsw (embedding vector_cosine_ops)")
cur.execute("CREATE INDEX IF NOT EXISTS ON docs USING gin (tsv)")

