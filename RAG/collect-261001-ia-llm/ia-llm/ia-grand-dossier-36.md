---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-36
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "Groq", "Hugging Face", "JFrog", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["attention", "awq", "benchmark", "chatgpt", "decode", "disaggregated", "embeddings", "gguf", "gpu", "kimi", "kv cache", "llama"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2594, 2711]
sha256: c5a46f392266c8867d02548e43ef95cf14f7ad35c60f327c35f75ba938a71cfc
---

# ── Télécharger un GGUF (ex. Qwen3 32B Q4) ───────────
pip install huggingface_hub
hf download Qwen/Qwen3-32B-GGUF --include "*Q4_K_M*" --local-dir ./models

# ── Servir en HTTP (API OpenAI-compat, port 8080) ─────
./build/bin/llama-server \
  -m ./models/qwen3-32b-q4_k_m.gguf \
  --port 8080 --host 0.0.0.0 \
  -c 32768 \                 # contexte 32K (défaut 4096)
  --n-gpu-layers 99 \        # tout sur GPU (0 = CPU seul, N = partiel)
  --threads $(nproc) \
  --parallel 4 \             # 4 requêtes simultanées
  --cache-type-k q8_0 \      # quantifie aussi le KV cache -> économise la VRAM
  --api-key "changeme"       # auth basique du serveur

# ── Quantifier soi-même ──────────────────────────────
./build/bin/llama-quantize ./models/modele-f16.gguf ./models/modele-q4_k_m.gguf Q4_K_M

# ── Benchmark rapide : est-ce que ça tient ? ─────────
./build/bin/llama-bench -m ./models/modele.gguf -p 512 -n 128
```

**Leviers perf** : `--n-gpu-layers` (offload partiel si la VRAM est juste : les couches restantes tournent sur CPU, lent mais ça passe), `--cache-type-k q8_0` (KV cache quantifié), `--flash-attn on` (réduit la mémoire d'attention). **Règle** : si ça OOM, baisser `-c` (contexte) avant de baisser la quantification.

#### 5.5.3. vLLM — servir un modèle à une équipe

```bash
# ── Installation (GPU NVIDIA, Python 3.10+) ──────────
pip install vllm            # version stable ; ~v0.10.x au printemps 2026, v0.28 côté vllm-metal (Mac)
# Vérifier : python -c "import vllm; print(vllm.__version__)"

# ── Serveur OpenAI-compat (port 8000) ────────────────
vllm serve Qwen/Qwen3-32B-AWQ \
  --served-model-name qwen3-32b \     # nom logique vu par les clients
  --host 0.0.0.0 --port 8000 \
  --api-key "${VLLM_API_KEY}" \       # auth Bearer
  --gpu-memory-utilization 0.90 \     # part de VRAM allouée
  --max-model-len 32768 \             # contexte max
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \    # parser adapté au modèle (kimi_k2, glm47…)
  --enable-prefix-caching \           # réutilise les préfixes (RAG : énorme gain)
  --quantization awq                  # si poids AWQ (auto-détecté sinon)

# ── Test ─────────────────────────────────────────────
curl http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer ${VLLM_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"model":"qwen3-32b","messages":[{"role":"user","content":"ping"}]}'

# Endpoints utiles : /health (sonde), /metrics (Prometheus), /v1/models
```

**En production** : 1 conteneur Docker par modèle/GPU (`vllm/vllm-openai` sur Docker Hub), derrière LiteLLM (§4) pour les fallbacks vers le cloud. **Multi-LoRA** : plusieurs adaptateurs fine-tunés servis sur les mêmes poids (`--enable-lora --lora-modules`), parfait pour « un modèle de base + un adaptateur par métier ». **Disaggregated prefill/decode** (versions récentes) : séparer le traitement du prompt (prefill) de la génération (decode) sur des GPU différents — le pattern des gros déploiements.

#### 5.5.4. LM Studio — la voie GUI

1. Télécharger sur `lmstudio.ai` (Mac/Windows/Linux, gratuit).
2. Onglet Discover : chercher `qwen3`, télécharger la variante `Q4_K_M`.
3. Onglet Developer : démarrer le serveur local (`http://localhost:1234`, API OpenAI-compat) → brancher n'importe quel client.
4. Idéal pour : faire essayer un LLM local à un collègue non-dev, tester un modèle avant de l'industrialiser.

#### 5.5.5. text-generation-webui & Open WebUI

```bash
# text-generation-webui (oobabooga) — bidouille avancée
git clone https://github.com/oobabooga/text-generation-webui && cd text-generation-webui
./start_linux.sh   # installe tout (conda), UI sur http://127.0.0.1:7860, API :5000

# Open WebUI — le "ChatGPT interne" (recommandé pour une équipe)
docker run -d -p 3000:8080 \
  -e OLLAMA_BASE_URL=http://ollama-interne.lan:11434 \
  -v open-webui:/app/backend/data \
  --name open-webui ghcr.io/open-webui/open-webui:main
# Puis : comptes utilisateurs, RAG intégré (upload de PDF -> embeddings),
# "Pipelines" (fonctions), connexion à vLLM/OpenAI via l'UI d'admin.
```

**Architecture interne type** (à mettre dans ton dossier d'exploitation) :

```
[Utilisateurs] -> [Open WebUI :3000] -> [LiteLLM :4000] -> noms logiques
                                                            ├─ local      -> vLLM :8000 (Qwen3-32B-AWQ, H100/4090)
                                                            ├─ local-cpu  -> Ollama :11434 (embeddings, petits modèles)
                                                            ├─ triage     -> Groq/Cerebras (interactif cloud)
                                                            └─ frontier   -> Anthropic/OpenAI (qualité, avec budget)
```

### 5.6. Mises à jour et hygiène d'exploitation

- **Épingler les versions** : `ollama --version`, tag Docker vLLM, commit llama.cpp — un changement de version peut changer le comportement (tokenizer, template de chat). Tester en staging.
- **Templates de chat** : chaque famille a le sien (ChatML, Llama, Qwen…). Ollama et vLLM l'appliquent automatiquement ; en appel brut à llama.cpp, le vérifier.
- **Sécurité** : GGUF = code exécutable potentiel (deseritisation) → ne télécharger que depuis des sources réputées (comptes officiels Hugging Face) ; exposer les ports via reverse proxy + TLS + auth ; journaliser les accès.
- **Sauvegarde** : les poids sont gros (5–80 Go) → un registre interne (simple partage SMB/NFS versionné, ou Nexus/Artifactory) évite de re-télécharger à chaque machine.

---

## 6. Hardware chiffré : dimensionner une machine locale

**La loi fondamentale** : en génération (decode), la vitesse ≈ **bande passante mémoire ÷ taille du modèle**. Un modèle de 40 Go sur une mémoire à 1 000 Go/s plafonne à ~25 tok/s. La VRAM décide **ce qui tient** ; la bande passante décide **à quelle vitesse ça répond**. Le calcul (FLOPS) ne compte presque plus en solo.

### 6.1. Mémoire requise par taille de modèle et quantification

Formule : `mémoire_poids ≈ paramètres × octets/paramètre` (FP16/BF16 = 2 o/B ; Q8 = ~1 o/B ; Q4_K_M = ~0,5–0,6 o/B), **+ 10–20 %** pour le KV cache (contexte 4–8K) et l'overhead du runtime. Valeurs terrain recoupées (sept. 2026, fichiers Ollama/GGUF réels) :

| Modèle | Q4_K_M (fichier) | Q4 : mémoire totale conseillée | Q8_0 (fichier) | Q8 : mémoire totale | F16 (fichier) | F16 : mémoire totale |
|---|---|---|---|---|---|---|
| **3B** | ~2,0 Go | ~3 Go | ~3,5 Go | ~5 Go | ~6 Go | ~8 Go |
| **7B** | ~4,1–4,7 Go | **~6 Go** | ~6,7–8 Go | ~10 Go | ~13,5–14 Go | ~16 Go |
| **8B** | ~4,9–5,2 Go | **~6–8 Go** | ~8–9 Go | ~11 Go | ~16 Go | ~19 Go |
| **13B** | ~8–9 Go | **~10–12 Go** | ~14–15 Go | ~17 Go | ~26–28 Go | ~30 Go |
| **14B** | ~9 Go | **~11 Go** | ~15 Go | ~18 Go | ~28 Go | ~32 Go |
| **27–32B** | ~18–20 Go | **~22–24 Go** | ~32–35 Go | ~38 Go | ~60–64 Go | ~70 Go |
| **70B** | ~40–43 Go | **~48 Go** | ~72–77 Go | ~85 Go | ~140 Go | ~155 Go |
| **120B (GPT-OSS)** | ~65–70 Go (à vérifier) | **~80 Go** | — | — | — | — |
| **405B** | ~243 Go (fichier Ollama) | **~270 Go** | — | — | ~810 Go | — |

Lecture : un **70B Q4** demande ~48 Go de mémoire rapide → ni une 4090 (24 Go) ni une 5090 (32 Go) seules ; il faut 2×24 Go, une carte 48 Go+, ou un Mac à mémoire unifiée ≥64 Go. Un **32B Q4** (~22–24 Go) est le **plafond confortable d'une RTX 4090/5090**.

