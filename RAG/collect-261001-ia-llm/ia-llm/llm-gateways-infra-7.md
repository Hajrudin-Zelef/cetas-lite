---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-7
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["gpu", "amd", "awq", "deepseek", "embedding", "embeddings", "fp8", "gguf", "gptq", "gpus", "kv cache", "llama"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [760, 917]
sha256: f9c3e80d7bcf479b28912b0fca6ddd14c6742953434a9cf21154eb4a73d6bc7e
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

Endpoints : `http://localhost:8000/v1/chat/completions`, `/v1/completions`, `/v1/embeddings` (modèles d'embedding supportés), `/metrics` (Prometheus), `/health`.

Paramètres clés : `--max-model-len` (contexte), `--gpu-memory-utilization` (0–0.95, marge pour le système), `--tensor-parallel-size` (= nombre de GPU), `--quantization` (awq, gptq, fp8), `--enable-prefix-caching` (**indispensable en RAG** : le prompt système + chunks récurrents ne sont ré-encodés qu'une fois).

## 56. vLLM : docker-compose de production

```yaml
services:
  vllm:
    image: vllm/vllm-openai:latest
    command: >
      --model Qwen/Qwen3-32B-AWQ
      --quantization awq
      --port 8000
      --gpu-memory-utilization 0.9
      --max-model-len 32768
      --enable-prefix-caching
    ports:
      - "8000:8000"
    volumes:
      - hf-cache:/root/.cache/huggingface
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: 1
              capabilities: [gpu]
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s
      timeout: 10s
      retries: 3

volumes:
  hf-cache:
```

Prérequis : **nvidia-container-toolkit** installé sur l'hôte (`nvidia-smi` doit fonctionner dans le conteneur). Premier démarrage = téléchargement des poids (plusieurs Go) → prévoir le temps et le volume persistant.

## 57. llama.cpp : quand et comment

- **Quand** : pas de GPU NVIDIA (CPU, AMD, Mac, embarqué), ou besoin du contrôle le plus fin (offload partiel couches CPU/GPU).
- **Serveur** : `llama-server -m modele.gguf -ngl 99 -c 16384 --port 8080` (`-ngl` = couches sur GPU, `99` = tout si ça tient).
- **API** : compatible OpenAI sur `/v1/chat/completions`.
- **GGUF** : LE format local standard (quantizations `Q4_K_M`, `Q5_K_M`, `Q8_0`…) — c'est ce qu'Ollama utilise en interne.
- Perf CPU réaliste : 7B Q4 à ~10–20 tok/s sur un bon CPU moderne — suffisant pour du batch, pas pour du temps réel.

## 58. Quantization : GGUF vs AWQ vs GPTQ (l'essentiel)

| Format | Bits | Où ça tourne | Qualité (perte vs FP16) | Usage |
|---|---|---|---|---|
| FP16 / BF16 | 16 | GPU (gros) | Référence (0) | Entraînement, qualité max |
| FP8 | 8 | GPU récents (Hopper/Ada) | Quasi nulle | vLLM `--quantization fp8` |
| **Q8_0 (GGUF)** | 8 | CPU/GPU | Négligeable (~0,1 %) | « Presque FP16 » en 2× moins lourd |
| **Q4_K_M (GGUF)** | ~4,5 | CPU/GPU | Faible (perplexité +1–3 %) | **Standard local** : le meilleur compromis |
| **AWQ** | 4 | GPU uniquement | Faible | vLLM, inférence GPU 4-bit |
| **GPTQ** | 4 (ou 3/8) | GPU uniquement | Faible à moyenne (selon calibration) | Alternative à AWQ |
| Q2/Q3 (GGUF) | 2–3 | CPU/GPU | **Forte** (dégradation visible) | Dernier recours, petits modèles |

Règles pratiques :
1. **Q4_K_M par défaut** : divise la VRAM par ~3,5 pour une perte quasi invisible sur Q&A et code.
2. Sous Q4 (Q3/Q2), les gros modèles (70B) restent souvent meilleurs que les petits en FP16 — à tester sur TON éval (section 34).
3. AWQ > GPTQ en facilité sous vLLM en 2026 ; les deux exigent un GPU.
4. La quantization **n'affecte presque pas la vitesse** sur GPU moderne (le goulot = la bande passante mémoire, réduite par la quantization → souvent plus rapide).

## 59. Dimensionnement VRAM : la formule (à connaître par cœur)

**Poids du modèle** ≈ `nb_paramètres × octets_par_paramètre` :
- FP16/BF16 : **2 octets** → 7B = 14 Go, 70B = 140 Go
- Q8 : **1 octet** → 7B = 7 Go, 70B = 70 Go
- Q4 : **~0,5 octet** → 7B ≈ 3,9 Go, 70B ≈ 40 Go

**KV cache** (mémoire du contexte, EN PLUS des poids) ≈ `2 × couches × têtes_KV × dim_tête × ctx × octets` :
- Ex. Llama 3.3 70B (80 couches, 8 têtes KV, dim 128) en FP16 : **~320 Ko par token** → 32K ctx ≈ **10 Go** de KV cache.
- Le KV cache grandit avec le **contexte ET le batch** : en RAG avec 32K de contexte, prévoir +30–50 % au-delà des poids.

**Tableau de cadrage** (poids seuls, ajouter KV cache + ~2 Go système/CUDA) :

| Modèle | FP16 | Q8 | Q4_K_M | GPU minimal réaliste (Q4 + ctx 16K) |
|---|---|---|---|---|
| 1.5B | 3 Go | 1,5 Go | ~1,1 Go | CPU / iGPU |
| 7B | 14 Go | 7 Go | ~4,7 Go | RTX 3060 12 Go |
| 8B | 16 Go | 8 Go | ~4,9 Go | RTX 4060 Ti 16 Go |
| 13–14B | 28 Go | 14 Go | ~9 Go | RTX 4070 Ti Super 16 Go (limite) |
| 27B | 54 Go | 27 Go | ~17 Go | RTX 4090 24 Go |
| 32B | 64 Go | 32 Go | ~20 Go | RTX 4090 24 Go |
| 70B | 140 Go | 70 Go | ~40 Go | 2× RTX 4090 ou 1× RTX 6000 Ada 48 Go |
| 671B (DeepSeek R1 full, MoE) | ~1 340 Go (FP8 ~670 Go) | — | ~370 Go (Q4) | Cluster 8× H100 / 8× MI300 — pas du local |

Lecture : sur une **RTX 4090 24 Go**, le sweet spot = **32B en Q4** (20 Go poids + KV cache 16K ≈ 3–4 Go). Le 70B Q4 tient à 40 Go + cache → il faut 48 Go (2× 4090 ou A6000/6000 Ada).

## 60. Quel GPU pour quoi (repères 2026, prix indicatifs neufs — à vérifier)

| GPU | VRAM | Prix ordre de grandeur (neuf, 2026) | Idéal LLM local |
|---|---|---|---|
| RTX 3060 12 Go | 12 | ~300 € | 7–8B Q4, embeddings |
| RTX 4070 Super | 12 | ~650 € | 8B Q4 + TEI |
| RTX 4070 Ti Super | 16 | ~850 € | 14B Q4 |
| RTX 4090 | 24 | ~1 800–2 000 € | **32B Q4** — le choix sysadmin |
| RTX 5090 | 32 | ~2 200–2 500 € (dispo variable) | 32B Q8 / 70B Q4 limite |
| RTX 6000 Ada | 48 | ~8 000 €+ | 70B Q4 confortable |
| 2× RTX 4090 | 48 (2×24, tensor-parallel) | ~3 800 € | 70B Q4 via vLLM |
| Mac Studio (Unified 64–128 Go) | partagée | ~2 500–5 000 € | 70B Q4 en unified memory (lent mais ça passe) |

Note pro : en entreprise, une 4090 « gaming » dans un serveur n'a pas de garantie adaptée ; pour de la prod, viser gammes pro (Ada/Ampere pro) ou du cloud (partie I).

## 61. Docker + GPU : le prérequis unique

```bash
# Ubuntu/Debian — NVIDIA Container Toolkit (1 fois par hôte)
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | \
  sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
  sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
  sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
sudo apt update && sudo apt install -y nvidia-container-toolkit
sudo nvidia-ctk runtime configure --runtime=docker
sudo systemctl restart docker

# Test (doit afficher la sortie de nvidia-smi du HOST)
docker run --rm --gpus all nvidia/cuda:12.6.0-base-ubuntu22.04 nvidia-smi
```

Si ce test échoue, rien d'autre (vLLM, Ollama GPU, TEI GPU) ne marchera : c'est le premier diagnostic.

## 62. Exposer ton LLM local en API « prod perso »

Architecture recommandée (tout sur ton serveur, derrière un reverse proxy) :

```
[Clients] --https--> [Caddy/Nginx :443 + basic auth] --> [vLLM :8000]
                                                      --> [Ollama :11434]
                                                      --> [TEI   :8080]
```

```bash
# Exemple Caddy (Caddyfile) — TLS auto + auth basique
llm.maison.lan {
    basic_auth {
        zelef $2a$14$HASH_FAKE_BCRYPT
    }
    reverse_proxy localhost:8000
}
```

Règles : **jamais** d'Ollama/vLLM exposé brut sur Internet (pas d'auth native) ; toujours TLS + mot de passe ou tailscale/WireGuard ; logge les accès.

## 63. Lien avec ton RAG : servir un LLM local pour ton app

Ton RAG a deux appels : **embeddings** (indexation + requête) et **génération** (réponse). Les deux peuvent être locaux :

```python
# config.py — un seul interrupteur pour basculer cloud <-> local
import os
USE_LOCAL = os.getenv("USE_LOCAL", "1") == "1"

