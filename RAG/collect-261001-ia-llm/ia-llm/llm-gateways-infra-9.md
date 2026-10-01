---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-9
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cohere", "Fireworks AI", "Groq", "Hugging Face", "Mistral", "Nebius", "OpenAI", "vLLM"]
dates: []
keywords: ["gpu", "apache", "attribution", "benchmarks", "cohere", "embedding", "embeddings", "gguf", "inference", "llama", "llama.cpp", "mistral"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1054, 1199]
sha256: 60149cbf033515dcedbe78a3b5155e2f604d87c38c051f82a986ec3e682325d4
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Critère | TEI local (`bge-m3`) | OpenAI `text-embedding-3-small` |
|---|---|---|
| Prix (vérifié) | **$0** (élec) | **$0.02 / 1M tokens** (prix public stable) |
| Dimensions | 1024 (bge-m3) | 1536 (réductible Matryoshka) |
| Contexte | 8 192 tokens | 8 191 tokens |
| Multilingue FR | **Excellent** (bge-m3 entraîné multilingue) | Très bon |
| Latence (batch indexation) | ~5–20 ms/doc en CPU, ~1–2 ms en GPU | ~100–300 ms + réseau |
| Débit indexation | Milliers de docs/min (batch) | Limité par RPM/TPM du compte |
| Données | **Ne sortent jamais** | Envoyées à OpenAI (pas d'entraînement par défaut, 30 j rétention) |
| Qualité (MTEB) | bge-m3 : top multilingue historique | Référence solide, légèrement derrière les SOTA 2026 |
| Opérabilité | Un conteneur à superviser | Zéro |

**Calcul** : 1M tokens d'indexation = **$0.02** chez OpenAI. Ça paraît rien — mais à 500M tokens de corpus + réindexations, c'est $10 par passe, et chaque requête coûte aussi. En local : 0, pour toujours. **Pour ton RAG : TEI + bge-m 
...[truncated 12084 chars]
## 75. TEI : service systemd (prod perso)

```ini
# /etc/systemd/system/tei-embed.service
[Unit]
Description=TEI embeddings (bge-m3)
Requires=docker.service
After=docker.service network-online.target
Wants=network-online.target

[Service]
ExecStartPre=-/usr/bin/docker rm -f tei-embed
ExecStart=/usr/bin/docker run --rm --name tei-embed \
  -p 8080:80 \
  -v tei-data:/data \
  ghcr.io/huggingface/text-embeddings-inference:cpu-1.8.3 \
  --model-id BAAI/bge-m3 --port 80 \
  --max-client-batch-size 256 --auto-truncate true
ExecStop=/usr/bin/docker stop tei-embed
Restart=always
RestartSec=5
TimeoutStartSec=600

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now tei-embed
# Premier boot : le téléchargement des poids peut prendre plusieurs minutes
journalctl -u tei-embed -f
```

## 76. Checklist TEI

- [ ] Conteneur TEI + `bge-m3` OK (`/health` → 200, `/info` correct).
- [ ] Test `/embed` et `/v1/embeddings` depuis ton code RAG.
- [ ] Décision dimensions : 1024 natives ou réduites (Matryoshka si modèle supporté).
- [ ] Conteneur reranker `bge-reranker-v2-m3` + test `/rerank`.
- [ ] Pipeline : top-50 cosinus → rerank top-5 → génération (section 74).
- [ ] Réindexation complète du corpus avec le nouveau modèle (jamais de mélange d'embeddings).
- [ ] Éval : comparer précision@5 avant/après sur 50 questions FR (section 34 adaptée).
- [ ] systemd + monitoring (latence p95, file de batch via `/metrics`).

---

# PARTIE H — HUGGING FACE (LE HUB ET SON ÉCOSYSTÈME)

## 77. Hugging Face : ce que c'est (en 2026)

Hugging Face (huggingface.co) = **le GitHub des modèles IA** : ~1M+ de modèles, datasets et démos au catalogue en 2026. Quatre piliers pour toi :
1. **Le Hub** : trouver/télécharger des modèles (LLM, embeddings, rerankers, vision, audio).
2. **Inference Providers** : API serverless unifiée (état sept 2026 : actif, voir section 79).
3. **transformers** : LA librairie Python pour charger et utiliser les modèles.
4. **Datasets + Spaces** : données d'éval et démos web gratuites pour tester avant d'installer.

## 78. Le Hub : trouver un modèle pour ton RAG (méthode)

1. Aller sur `huggingface.co/models`, filtrer : tâche (`feature-extraction` pour embeddings, `text-generation` pour LLM), langue (`french`), licence (`mit`, `apache-2.0`).
2. Lire la **fiche modèle** : dimensions, contexte, licence, benchmarks, « Use this model » (snippet prêt à copier).
3. Vérifier les signaux de confiance : téléchargements/mois, likes, organisation vérifiée (BAAI, nomic-ai, Qwen, jinaai…), date de mise à jour.
4. Tester **sans rien installer** : widget « Inference » sur la fiche, ou un **Space** de démo (section 83).
5. Pour le local : vérifier la présence de fichiers **GGUF** (onglet Files → quantizations) ou la compatibilité **TEI/vLLM/Ollama**.

Exemple de recherche ciblée : `sentence-similarity french` → bge-m3, `jina-embeddings-v3` ; `text-generation qwen3 32b gguf` → variantes quantifiées prêtes pour Ollama/llama.cpp.

## 79. Inference Providers : état vérifié (sept 2026)

- **Actif** : endpoint unifié **`https://router.huggingface.co/v1`** (compatible OpenAI), routage vers **18+ partenaires** (Groq, Fireworks, Mistral, Cohere, Nebius, SambaNova…) — **200+ modèles**, **sans surcoût** (HF répercute le prix fournisseur).
- **Crédits mensuels** (vérifié 2026) : **$0.10/mois** (compte Free), **$2/mois** (PRO) — c'est un filet d'essai, pas une infra.
- **BYOK** : tu peux brancher ta propre clé d'un partenaire et payer le partenaire directement (les crédits HF ne s'appliquent plus, mais tu gardes la facturation directe).
- Le backend historique `hf-inference` (ex-« Inference API serverless ») existe toujours, recentré sur **CPU et petits modèles** (embeddings, classification, BERT/GPT-2) ; les gros LLM routent vers les partenaires.
- Limite serverless : modèles **< 10 Go** (quelques exceptions populaires).
- Verdict pour toi : utile pour **tester un modèle du Hub via API** avant de le self-héberger, pas comme backend de prod (crédits trop faibles).

```python
from openai import OpenAI
client = OpenAI(base_url="https://router.huggingface.co/v1", api_key="hf_FAKEFAKEFAKE")
r = client.chat.completions.create(
    model="Qwen/Qwen3-32B:fireworks-ai",   # syntaxe modele:provider (vérifier la doc)
    messages=[{"role": "user", "content": "Bonjour"}])
```

## 80. transformers : les bases d'usage (indispensable)

```bash
pip install transformers torch --index-url https://download.pytorch.org/whl/cu126
pip install sentence-transformers   # surcouche pratique pour les embeddings
```

```python
# --- Embeddings en 5 lignes (sans TEI, pour prototyper) ---
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("BAAI/bge-m3")
vecs = model.encode(["panne onduleur", "batterie VRLA"], normalize_embeddings=True)
print(vecs.shape)  # (2, 1024)

# --- LLM en local (petit modèle, test) ---
from transformers import pipeline
gen = pipeline("text-generation", model="Qwen/Qwen2.5-3B-Instruct", device_map="auto")
print(gen("Explique le VLAN en une phrase.", max_new_tokens=60)[0]["generated_text"])

# --- Tokenizer seul (compter les tokens avant d'appeler une API payante) ---
from transformers import AutoTokenizer
tok = AutoTokenizer.from_pretrained("Qwen/Qwen3-32B")
print(len(tok("Ton texte ici")["input_ids"]), "tokens")
```

`sentence-transformers` suffit pour prototyper tes embeddings ; **passe à TEI** dès que tu veux du débit (section 65). Le tokenizer local = ton **compteur de tokens gratuit** pour estimer les coûts API avant envoi.

## 81. Datasets : évaluer ton RAG avec des données réelles

```python
from datasets import load_dataset
# Exemple : charger un jeu de questions FR pour tester ton RAG
ds = load_dataset("antoinelouis/colbertv2.0", split="train")  # exemple — adapter
```

Usages : construire ton **jeu d'éval** (50 questions + réponses attendues + docs sources), mesurer précision@k, versionner les jeux de test comme du code. Le Hub héberge aussi des corpus FR utiles pour enrichir un RAG technique (Wikipedia FR, etc.).

## 82. Licences : le piège n°1 du Hub (à lire avant usage pro)

| Licence | Usage commercial | Ce que ça implique |
|---|---|---|
| MIT / Apache-2.0 | **Oui** | Libre, attribution requise |
| BGE (MIT) | Oui | OK pour ton RAG, même pro |
| **Llama Community** | Oui sous conditions | Restrictions (pas d'entraînement de concurrents, etc.) — lire le texte |
| **Qwen** (selon version) | Oui sous conditions | Vérifier la version exacte |
| CC-BY-NC / Non-commercial | **Non** | Interdit en entreprise |
| Gated (accès sur demande) | Selon | Accepter les conditions sur la fiche avant téléchargement |

