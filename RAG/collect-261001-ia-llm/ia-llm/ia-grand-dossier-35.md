---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-35
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Hugging Face", "Microsoft", "Mistral", "Nvidia", "OpenAI", "TensorRT-LLM", "vLLM"]
dates: ["2026-08-21", "2026-09-05"]
keywords: ["awq", "benchmarks", "blackwell", "chatgpt", "deepseek", "distribution", "embedding", "embeddings", "fp8", "gguf", "gptq", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2499, 2593]
sha256: 7e0ab60a2984dcdb48f2d88545e7f735304839373486735d1f4ab6cd8a84bcb9
---

# IA — Le grand dossier

**Règle de sélection pour un RAG FR** : le modèle doit (1) parler français correctement (Qwen3, Mistral, Llama le font), (2) suivre un format JSON strict pour l'extraction, (3) tenir dans ta VRAM (section 6). Pour les embeddings FR, `Qwen3-Embedding` ou `nomic-embed-text` en local remplacent très bien text-embedding-3-small sur de la doc technique (à valider par eval sur ton corpus — voir §7.1).

### 5.3. Les formats de poids : GGUF, AWQ, GPTQ (et les autres)

Un modèle « brut » (FP16/BF16) est trop gros pour un poste de travail. La **quantification** réduit la précision des poids (16 bits → 8/4/2 bits) avec une perte de qualité mesurée : **Q8 ≈ 100 % de la qualité FP16, Q4 ≈ 92–98 %** (perplexité +0,05–0,2 pt, benchmarks -1–2 % — vérifié sur Llama 4 Scout, sept. 2026).

| Format | Principe | Écosystème | Quand l'utiliser |
|---|---|---|---|
| **GGUF** | Format de **llama.cpp** ; quantifications K (Q4_K_M, Q5_K_M, Q8_0…) mixtes par couche | Ollama, llama.cpp, LM Studio, text-generation-webui, vLLM (lecture) | **Défaut en local** : un fichier, CPU+GPU, le plus large support |
| **AWQ** | Quantification 4-bit « activation-aware » (protège les poids importants) | vLLM, TGI, llama.cpp (partiel) | Serveur **vLLM** : le meilleur débit sur GPU NVIDIA |
| **GPTQ** | 4-bit par couches, le plus ancien standard | vLLM, text-generation-webui, Transformers | Modèles pré-quantifiés existants ; supplanté par AWQ sur vLLM |
| **FP8** | 8-bit natif GPU récents (Ada/Hopper) | vLLM, TensorRT-LLM, certains providers (SiliconFlow) | Serveur pro : -50 % mémoire pour ~0 perte |
| **MLX** | Format Apple (mlx-lm) | Mac uniquement | Le plus rapide sur Apple Silicon |
| **safetensors** | Format de stockage Hugging Face (pas une quantification) | Tout l'écosystème | Distribution des poids d'origine |

**Conversions** : `llama.cpp` fournit `llama-quantize` (safetensors → GGUF Q4_K_M en une commande) ; Hugging Face regorge de variantes déjà quantifiées (chercher `Q4_K_M GGUF` sur le Hub). **Piège** : une quantification trop agressive (Q2, IQ1) sur un petit modèle (≤8B) dégrade vite ; sur un 70B, Q4 reste excellent car la redondance des paramètres compense.

**Le KV cache** (mémoire des conversations en cours) n'est **pas** réduit par la quantification des poids — sauf si l'outil le quantifie aussi (vLLM, llama.cpp le proposent). À 128K de contexte, le KV cache peut ajouter **10 Go+** : c'est lui qui fait OOM, pas les poids (tableaux §6.2).

### 5.4. Choisir sa stack : le tableau

| Stack | Nature | Points forts | Points faibles | Idéal pour |
|---|---|---|---|---|
| **Ollama** | Runtime tout-en-un (Go + llama.cpp/MLX), `localhost:11434`, API OpenAI-compat | Install en 1 commande, `ollama pull/run`, multi-plateforme (Win/Mac/Linux, CUDA/Metal/ROCm/CPU), 176k★, MIT | Moins réglable que llama.cpp pur ; 1 hôte | **Débuter et produire vite** ; RAG local ; dev |
| **llama.cpp** | Moteur C++ de référence, `llama-server` | Le plus rapide en solo, contrôle total (offload GPU/CPU par couche, contextes), builds nightly | Ligne de commande / compilation ; courbe d'apprentissage | **Perf maximale**, edge, intégration embarquée |
| **vLLM** | Serveur Python d'inférence haut débit (PagedAttention, batching continu) | **Débit multi-utilisateurs** imbattable, 200+ architectures, AWQ/GPTQ/FP8, LoRA multiples, `/metrics` Prometheus | GPU NVIDIA (ou gros CPU) ; 1 modèle/GPU en général | **Servir** un modèle à une équipe / une appli |
| **LM Studio** | App desktop (GUI) + serveur local OpenAI-compat | Zéro terminal, catalogue de modèles intégré, fonctionne sur Mac/Win/Linux | Moins scriptable ; desktop d'abord | **Non-dev**, démo, poste de travail |
| **text-generation-webui** (oobabooga) | Interface web + API, très extensible | Extensions, personnages, ExLlamaV2, test rapide de quantifs | Projet communautaire, UX datée | Bidouille, comparaisons |
| **Open WebUI** | Frontend web « ChatGPT local » | Belle UI, RAG intégré, multi-utilisateurs, parle à Ollama/vLLM/OpenAI | Ce n'est qu'un frontend (il faut un backend) | Donner un **chat interne** aux collègues |

**Le trio gagnant en entreprise** : **vLLM** (servir) + **Ollama** (postes/dev) + **Open WebUI** (interface) — le tout derrière **LiteLLM** (§4) qui expose aussi les providers cloud sous les mêmes noms logiques. Un seul point d'entrée, deux mondes (local + cloud).

### 5.5. Stacks en pratique : installation et usage réel

#### 5.5.1. Ollama — le plus court chemin vers un LLM local

```bash
# ── Installation (Linux) ─────────────────────────────
curl -fsSL https://ollama.com/install.sh | sh
# Windows/macOS : installateur sur ollama.com/download
# Version vérifiée : v0.34.0 (05/09/2026). Vérifier : ollama --version

# ── Premier modèle ───────────────────────────────────
ollama pull qwen3:8b            # ~5,2 Go (Q4_K_M) — généraliste FR correct
ollama pull qwen2.5-coder:7b    # ~4,7 Go — code
ollama pull nomic-embed-text    # 274 Mo — embeddings locaux
ollama pull deepseek-r1:8b      # ~5,2 Go — raisonnement (chaîne de pensée)

ollama run qwen3:8b             # chat interactif
# /bye pour quitter ; /set parameter pour régler (température…)

# ── API (compatible OpenAI) : le RAG parle à http://localhost:11434 ──
curl http://localhost:11434/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model":"qwen3:8b","messages":[{"role":"user","content":"Bonjour"}]}'

# Embeddings locaux (remplace text-embedding-3-small en dev)
curl http://localhost:11434/v1/embeddings \
  -H "Content-Type: application/json" \
  -d '{"model":"nomic-embed-text","input":"onduleur triphasé 40 kVA"}'

# ── Modelfile : pinner ton comportement (équivalent Dockerfile) ──
cat > Modelfile <<'EOF'
FROM qwen3:8b
PARAMETER temperature 0.2
PARAMETER num_ctx 8192
SYSTEM """Tu es un assistant technique pour un service systèmes & énergies.
Réponds en français, de façon concise, avec des commandes vérifiables.
Si tu ne sais pas, dis-le au lieu d'inventer."""
EOF
ollama create tech-fr -f Modelfile
ollama run tech-fr

# ── Multi-modèles / mémoire ──────────────────────────
# OLLAMA_MAX_LOADED_MODELS=2  -> garde 2 modèles en VRAM
# OLLAMA_NUM_PARALLEL=4        -> 4 requêtes simultanées
# OLLAMA_HOST=0.0.0.0          -> exposer sur le LAN (à protéger ! pas d'auth native)
# OLLAMA_KEEP_ALIVE=30m        -> durée de résidence en VRAM
systemctl edit ollama  # puis [Service] Environment="OLLAMA_HOST=0.0.0.0"
```

**Points durs** : pas d'authentification native → ne jamais exposer `11434` sans reverse proxy (Caddy/Nginx + clé) ; un seul hôte (pas de cluster) ; la quantification est imposée par le registre (pour du Q8/AWQ custom → llama.cpp).

**Ollama Cloud** (nouveauté 2026) : `ollama run gpt-oss:120b` peut router vers le cloud Ollama si le modèle ne tient pas en local — pratique, mais les données partent chez Ollama : à désactiver explicitement pour un usage souverain (variable d'environnement, à vérifier sur ta version).

#### 5.5.2. llama.cpp — le moteur nu, performance maximale

```bash
# ── Build (Linux, CUDA) ──────────────────────────────
git clone https://github.com/ggml-org/llama.cpp && cd llama.cpp
cmake -B build -DGGML_CUDA=ON -DCMAKE_CUDA_ARCHITECTURES="89"  # 89=Ada (4090), 90=Hopper, 100=Blackwell
cmake --build build --config Release -j
# Binaires précompilés : releases GitHub (tags b10xxx "bleeding edge",
# vX.Y.Z "stable pour downstream" depuis le 21/08/2026)

