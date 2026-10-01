---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-8
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Groq", "Hugging Face", "OpenAI", "vLLM"]
dates: []
keywords: ["gpu", "apache", "attention", "awq", "compute", "deepseek", "embedding", "embeddings", "gpt-6", "gpus", "inference", "moe"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [918, 1053]
sha256: adb8d7a915027845f2c4cc0f713e7a921a17aae8130918a9032aef9cb9df78e5
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

if USE_LOCAL:
    LLM_BASE_URL = "http://localhost:8000/v1"      # vLLM (Qwen3-32B-AWQ)
    LLM_MODEL    = "Qwen/Qwen3-32B-AWQ"
    EMB_URL      = "http://localhost:8080"          # TEI (bge-m3)
else:
    LLM_BASE_URL = "https://api.groq.com/openai/v1" # gratuit cloud
    LLM_MODEL    = "openai/gpt-oss-120b"
    EMB_URL      = "https://api.openai.com/v1"      # text-embedding-3-small
    EMB_MODEL    = "text-embedding-3-small"
```

Stratégie hybride réaliste : **embeddings toujours locaux** (TEI, coût 0, qualité équivalente — partie G), **génération locale par défaut** (vLLM/Ollama) avec **fallback cloud gratuit** (Groq) si le serveur est éteint, et **escalade payante** (DeepSeek/gpt-6-sol) sur les questions ratées.

## 64. Checklist self-hosting

- [ ] `nvidia-smi` OK + test Docker GPU (section 61).
- [ ] Ollama installé, `qwen3:32b` (ou équivalent) testé sur 10 questions de ton domaine.
- [ ] Modelfile avec ton prompt système (section 54).
- [ ] vLLM en docker-compose si besoin de débit / multi-utilisateurs (section 56).
- [ ] `--enable-prefix-caching` activé (RAG).
- [ ] Reverse proxy + auth si accès LAN (section 62).
- [ ] Éval qualité local vs cloud sur tes 20–50 questions (section 34) avant de figer.

---

# PARTIE G — TEI (TEXT EMBEDDINGS INFERENCE)

## 65. TEI : présentation et état (sept 2026)

**TEI** (github.com/huggingface/text-embeddings-inference, Apache-2.0) est le serveur d'inférence officiel de Hugging Face pour les **embeddings** (et aussi reranking `/rerank` et classification `/predict`). Écrit en **Rust** (moteur Candle), il fait le **batching dynamique par tokens**, expose des métriques **Prometheus** + OpenTelemetry, et propose des images Docker **CPU et GPU** (par compute capability : `cpu-`, `turing-`, `86-` pour Ampere 86, `89-` pour Ada, `hopper-`…).
État sept 2026 : **projet actif et maintenu** (versions 1.8.x / 1.9.x constatées), utilisé en prod par de nombreuses équipes RAG. C'est **l'alternative locale directe à `text-embedding-3-small`** pour ton RAG : mêmes dimensions paramétrables, coût 0, données qui ne sortent pas.

## 66. TEI : installation Docker (CPU et GPU)

```bash
# --- CPU (suffit pour bge-small/base et l'indexation en batch) ---
docker run -d --name tei \
  -p 8080:80 \
  -v tei-data:/data \
  --restart unless-stopped \
  ghcr.io/huggingface/text-embeddings-inference:cpu-1.8.3 \
  --model-id BAAI/bge-m3

# --- GPU (Ampere 86 = RTX 30xx/A40 ; 89 = Ada/RTX 40xx ; hopper = H100) ---
docker run -d --name tei \
  --gpus all -p 8080:80 \
  -v tei-data:/data \
  --restart unless-stopped \
  ghcr.io/huggingface/text-embeddings-inference:86-1.8.3 \
  --model-id BAAI/bge-m3
```

Le premier démarrage télécharge les poids (~2 Go pour bge-m3) dans `/data` (volume persistant → un seul téléchargement). Santé : `curl localhost:8080/health` → 200. Infos modèle : `curl localhost:8080/info | jq`.

## 67. TEI : modèles d'embeddings recommandés (vérifiés)

| Modèle (`--model-id`) | Dim. | Langues | Contexte | Poids | Pourquoi |
|---|---|---|---|---|---|
| `BAAI/bge-m3` | 1024 | **Multilingue (100+, FR excellent)** | 8 192 | ~2,2 Go | **Choix par défaut pour ton RAG FR** |
| `BAAI/bge-small-en-v1.5` | 384 | Anglais | 512 | ~130 Mo | Ultra-léger, CPU instantané |
| `BAAI/bge-base-en-v1.5` | 768 | Anglais | 512 | ~440 Mo | Bon compromis EN |
| `BAAI/bge-large-en-v1.5` | 1024 | Anglais | 512 | ~1,3 Go | Qualité max EN classique |
| `jinaai/jina-embeddings-v3` | 1024 | Multilingue | 8 192 | ~2,3 Go | Alternative multilingue, Matryoshka |
| `nomic-ai/nomic-embed-text-v1.5` | 768 | Anglais (+multilingue partiel) | 8 192 | ~550 Mo | Long contexte, léger |
| `nomic-ai/nomic-embed-text-v2-moe` | 768 | Multilingue | 8 192 | MoE | Version 2026, à évaluer |
| `Qwen/Qwen3-Embedding-0.6B` (famille) | 1024 | Multilingue | 32K | ~2,4 Go | Long contexte, famille Qwen3 |

**Attention licence** : BGE = MIT (usage commercial OK) ; Jina = Apache-2.0/CC-BY-NC selon version — **vérifier la fiche du modèle exact** avant usage pro. Pour ton RAG FR : **`bge-m3` en premier**, `jina-embeddings-v3` en challenger.

## 68. TEI : configuration (batching, dimensions, truncation)

```bash
docker run -d --name tei \
  -p 8080:80 -v tei-data:/data \
  ghcr.io/huggingface/text-embeddings-inference:cpu-1.8.3 \
  --model-id BAAI/bge-m3 \
  --max-client-batch-size 256 \     # requêtes concurrentes max
  --max-batch-tokens 32768 \         # budget tokens par batch (le batching dynamique s'auto-règle)
  --max-batch-requests 128 \
  --auto-truncate true               # tronque au lieu d'erreur si texte > contexte
```

- **Batching dynamique par tokens** : TEI remplit chaque batch au mieux — pas de calcul manuel `batch × seq²`.
- `--auto-truncate true` : indispensable en indexation (un PDF avec une page monstrueuse ne doit pas faire planter le batch).
- **Dimensions** : `/v1/embeddings` accepte `dimensions` (Matryoshka) sur les modèles qui le supportent (jina-v3, nomic) → ex. 512 au lieu de 1024 = 2× moins de stockage FAISS/Chroma.
- **Normalisation** : TEI renvoie des vecteurs normalisés par défaut selon le modèle — à vérifier pour ta similarité cosinus.

## 69. TEI : appels API (natif + OpenAI-compatible)

```bash
# Natif : un vecteur par input
curl -s localhost:8080/embed -H 'Content-Type: application/json' \
  -d '{"inputs": "Le protocole OSPF utilise Dijkstra."}' | jq '.[0] | length'
# -> 1024  (bge-m3)

# Natif : batch
curl -s localhost:8080/embed -H 'Content-Type: application/json' \
  -d '{"inputs": ["texte 1", "texte 2", "texte 3"]}' | jq 'length'
# -> 3

# OpenAI-compatible : même code que ton RAG actuel !
curl -s localhost:8080/v1/embeddings -H 'Content-Type: application/json' \
  -d '{"input": ["texte 1", "texte 2"], "model": "bge-m3"}' | jq '.data | length'
```

```python
# Python : drop-in replacement de text-embedding-3-small
from openai import OpenAI
emb = OpenAI(base_url="http://localhost:8080/v1", api_key="tei")
v = emb.embeddings.create(input=["panne onduleur", "batterie VRLA"], model="bge-m3")
print(len(v.data[0].embedding))  # 1024
```

→ Migration de ton RAG : changer `base_url` + `model` + `dimensions` (1536 → 1024) et **réindexer** (les vecteurs ne sont pas compatibles entre modèles).

## 70. TEI : reranking (le bonus qui change tout)

TEI sert aussi les **cross-encoders** (`BAAI/bge-reranker-*`) sur `/rerank` — l'étage qui manque à la plupart des RAG maison :

```bash
docker run -d --name tei-rerank -p 8082:80 -v tei-data:/data \
  ghcr.io/huggingface/text-embeddings-inference:cpu-1.8.3 \
  --model-id BAAI/bge-reranker-v2-m3 --auto-truncate false

curl -s localhost:8082/rerank -H 'Content-Type: application/json' -d '{
  "query": "durée de vie batterie onduleur",
  "texts": ["Les batteries VRLA durent 3 à 5 ans à 20°C.",
            "Le protocole OSPF calcule le plus court chemin."]
}' | jq
# -> scores triés : le texte batterie en premier
```

Pipeline RAG pro : **récupération large** (top-50 en cosinus, pas cher) → **rerank** (top-5, précis) → génération. Le rerank coûte ~10 ms/doc en local et améliore la pertinence plus que n'importe quel changement de modèle d'embedding.

## 71. TEI : bench vs API OpenAI text-embedding-3-small

