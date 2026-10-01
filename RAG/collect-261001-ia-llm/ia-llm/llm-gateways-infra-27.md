---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-27
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Mistral", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: []
keywords: ["gpu", "amd", "datacenter", "embedding", "embeddings", "gpus", "inference", "kv cache", "llama", "mistral", "nvidia", "qwen"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3171, 3302]
sha256: 2fcf24cb9be243e181438fac3ae16dab76a9a77a651a6c4ec1bd5dbeeee7cb42
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

Causes habituelles d'OOM en inférence : oubli de `torch.no_grad()` (le graphe de gradient reste en mémoire), batch trop gros, `max_new_tokens` délirant (le KV cache grandit avec la génération), deux modèles chargés « pour comparer » sur la même carte. En entraînement : ajoute le gradient checkpointing et la mixed precision (`torch.amp`).

### 164.6. Embeddings avec PyTorch : le chaînon vers ton RAG

Ton RAG indexe avec text-embedding-3-small (API). Le pendant local — utile pour chiffrer le « tout local » ou debugger — tient en 15 lignes avec `sentence-transformers` (qui n'est qu'un wrapper PyTorch + transformers) :

```python
# pip install sentence-transformers
from sentence_transformers import SentenceTransformer
import torch

model = SentenceTransformer("BAAI/bge-m3",  # multilingue, bon sur le FR — à benchmarker vs ton embedding actuel
                            model_kwargs={"torch_dtype": torch.float16})
docs = [
    "La tension de floating d'un bloc VRLA 12 V est de 13,5 à 13,6 V à 20 °C.",
    "Le code C6000 Kyocera signale une rupture de chauffe du four.",
    "Un onduleur double conversion isole totalement la charge du réseau.",
]
emb = model.encode(docs, normalize_embeddings=True)  # (3, 1024) — des tensors numpy
print(emb.shape, emb.dtype)

# Similarité cosinus = produit scalaire sur vecteurs normalisés
import numpy as np
q = model.encode(["Quelle tension pour un bloc batterie 12 V ?"], normalize_embeddings=True)
scores = q @ emb.T
print(scores)  # le doc 0 doit gagner — sinon, ton embedding ou ton chunking a un problème
```

Ce script est aussi ton **test de fumée d'embedding** : si la similarité ne classe pas correctement 10 paires évidentes de ton domaine, inutile d'indexer 10 000 documents — change d'embedding ou de chunking d'abord (sections 140–141).

### 164.7. `torch.compile` et l'inférence rapide : ce qu'il faut en retenir

```python
model = AutoModelForCausalLM.from_pretrained(model_id, torch_dtype=torch.bfloat16, device_map="auto")
model.eval()
model = torch.compile(model, mode="reduce-overhead")  # une ligne — compile le graphe

# Premier appel = compilation (lent, 1 à 5 min) ; les suivants = +10 à 30 %
with torch.no_grad():
    out = model.generate(**inputs, max_new_tokens=200)
```

Limites honnêtes : le warmup de compilation est réel (prévois-le au démarrage du service, pas à la première requête utilisateur) ; certains modèles exotiques « cassent » le graphe (graph breaks) et le gain s'évapore — mesure avant/après. vLLM/SGLang font déjà ce genre d'optimisations (et bien plus) en interne : `torch.compile` manuel sert surtout pour tes scripts maison (embedding batch, reranking, prototypes).

---

## 165. NVIDIA NIM : l'inférence en microservice clé en main

**En une phrase :** un conteneur Docker par modèle, maintenu par NVIDIA, qui embarque poids optimisés + moteur d'inférence (TensorRT-LLM, vLLM ou SGLang selon le profil) + API OpenAI-compatible — tu `docker run`, ça sert.

### 165.1. Principe

NIM (NVIDIA Inference Microservice) répond à une question simple : « que faut-il pour servir un LLM comme n'importe quel service de prod ? » La réponse de NVIDIA : un conteneur qui fait tout, tout seul.

```
docker run ...
      ↓
NIM inspecte ton GPU
      ↓
choisit le profil optimal pour ce GPU (moteur + quantification validés par NVIDIA)
      ↓
télécharge les poids depuis NGC (ou utilise le cache local)
      ↓
démarre le moteur avec une config validée
      ↓
expose une API OpenAI-compatible sur le port 8000
```

Dans le conteneur, toujours les mêmes briques : poids optimisés par NVIDIA · moteur d'inférence (vLLM, TensorRT-LLM ou SGLang selon profil) · serveur API `/v1/chat/completions` · sondes `/v1/health/live` et `/v1/health/ready` · métriques Prometheus sur `/metrics`. Ce ne sont pas des options : c'est le contrat de chaque NIM.

Positionnement dans la gamme NVIDIA : **NeMo** = entraîner/fine-tuner · **NIM** = servir · **TensorRT-LLM** = le moteur d'optimisation sous le capot. Tu ne fine-tunes pas avec NIM, tu sers avec.

### 165.2. Déploiement (vérifié sept 2026)

Prérequis : GPU NVIDIA (datacenter, workstation ou RTX récent — il existe des NIM « RTX » pour postes locaux, produit distinct de l'offre cloud), driver récent, Docker + nvidia-container-toolkit, et une **clé API NGC** (gratuite, depuis ton compte NVIDIA) pour tirer les conteneurs du registre `nvcr.io`.

```bash
# 1. Login au registre NVIDIA (clé API NGC — mets la vraie dans un secret, jamais en clair)
echo "$NGC_API_KEY" | docker login nvcr.io --username '$oauthtoken' --password-stdin

# 2. Lancer un NIM (exemple : Llama 3.1 8B instruct — adapte le tag à ton modèle)
docker run -d --name nim-llama \
  --gpus all \
  -p 8000:8000 \
  -e NGC_API_KEY="$NGC_API_KEY" \
  -v "$HOME/.cache/nim:/opt/nim/.cache" \
  nvcr.io/nim/meta/llama-3.1-8b-instruct:latest

# 3. Attendre que le modèle soit prêt (téléchargement + optimisation au 1er lancement)
curl http://localhost:8000/v1/health/ready

# 4. Inférence — même format qu'OpenAI
curl http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "meta/llama-3.1-8b-instruct",
    "messages": [{"role": "user", "content": "Qu'est-ce qu'un bypass de maintenance sur un onduleur ?"}],
    "max_tokens": 300
  }'
```

En Kubernetes (Helm, vérifié sur des déploiements 2026) :

```bash
helm repo add nim https://helm.ngc.nvidia.com/nim
helm install nim-llama nim/meta-llama3-3-70b-instruct \
  --namespace nim --create-namespace \
  --set model.ngcAPIKey="$NGC_API_KEY"
# Le pod expose /v1/chat/completions : ton code existant ne change pas
```

Notes de prod :
- **Premier démarrage = long** : téléchargement des poids + construction du moteur TensorRT-LLM (plusieurs minutes à dizaines de minutes selon modèle/GPU). Monte un volume de cache (`/opt/nim/.cache`) pour ne payer ce coût qu'une fois.
- Catalogue : des dizaines de modèles packagés (Llama, Mistral, Qwen, Gemma, embeddings, rerankers, vision…) sur `build.nvidia.com` — vérifie que *ton* modèle y existe avant d'architecturer autour de NIM.
- NIM existe aussi en version **workstation/RTX** (Windows via WSL2) pour faire tourner des modèles localement sur un PC RTX — pratique pour une démo ou un poste de dev, pas pour servir une équipe.

### 165.3. Cas d'usage — et quand NE PAS l'utiliser

| Cas | NIM pertinent ? | Pourquoi |
|---|---|---|
| Servir 1–N modèles en prod sur GPU NVIDIA, sans tuner de moteur | ✅ Oui | C'est exactement le produit : profil validé, API standard, sondes/métriques incluses |
| Équipe qui veut une API OpenAI-compatible interne rapidement | ✅ Oui | `docker run` + clé NGC, pas de code de serving à écrire |
| Kubernetes avec autoscaling | ✅ Oui | Chart Helm officiel, HPA sur les métriques `/metrics` |
| Fine-tuner un modèle | ❌ Non | C'est NeMo (ou Unsloth, section 162) |
| GPU non-NVIDIA (AMD, Apple Silicon, CPU) | ❌ Non | NIM = NVIDIA uniquement ; vLLM/SGLang/Ollama sinon |
| Modèle exotique non packagé par NVIDIA | ❌ Non | Pas de NIM officiel = retour à vLLM/SGLang manuel |
| Budget serré / spot Vast.ai | ⚠️ Réfléchir | Le conteneur est lourd (plusieurs Go) et le 1er boot coûte cher en temps ; sur du spot préempté souvent, vLLM léger peut être plus agile |

Ordre de grandeur perf (tiers, 2026) : Llama 3.3 70B en NIM sur H100 ≈ 2 400 tok/s vs ~1 200 tok/s en vLLM vanille sur le même hardware — l'écart vient du backend TensorRT-LLM compilé spécifiquement pour ce couple modèle/GPU. C'est le deal NIM : tu échanges du contrôle (choix du moteur, flags fins) contre de la perf validée et du temps d'ingénierie économisé.

### 165.4. NIM dans ta stack RAG

