---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-14
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "Cohere", "DeepSeek", "Google", "Groq", "Hugging Face", "Meta", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2026-09-27"]
keywords: ["gpu", "agents", "cohere", "deepseek", "embedding", "embeddings", "fine-tuning", "gemini", "glm", "grok", "inference", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1701, 1836]
sha256: 615bac800ada3bb69eef9cdd5c49401823d762034fef7b27c177294257afe6e8
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

1. **Confondre Groq et Grok** (xAI) — deux entreprises sans rapport.
2. **Promos prises pour des prix** : GLM 5.3 Flash ($0.075/$0.25), Mercury 2.5 ($0.04/$0.15) sont des promos de lancement.
3. **Fenêtres peak DeepSeek** : 2× le prix en semaine 01h–04h / 06h–10h UTC sur la route officielle.
4. **50 req/jour OpenRouter free** : quota d'essai, pas une infra — empiler les gratuits.
5. **429 non géré** : sans retry + backoff, ton appli meurt dès que le free tier sature.
6. **Données dans le gratuit** : Google/Mistral/DeepSeek peuvent entraîner sur tes prompts (opt-out à activer ou à éviter).
7. **Mélanger des embeddings** de deux modèles dans le même index — résultats faux silencieusement.
8. **Oublier `--auto-truncate`** en indexation TEI : un doc trop long fait échouer tout le batch.
9. **Q2/Q3 sans éval** : la dégradation qualité est réelle, à mesurer sur tes questions.
10. **KV cache oublié** dans le dimensionnement : +30–50 % au-delà des poids en RAG à gros contexte.
11. **Ollama exposé sans auth** sur le réseau : pas d'authentification native — reverse proxy obligatoire.
12. **Clé API dans Git** : une fois poussée, elle est compromise — rotation immédiate.
13. **Instance Vast.ai non détruite** : la facture tourne même « stopped ».
14. **Egress Vast.ai** : $0.01–0.02/Go — 140 Go de checkpoint = $2.80 à rapatrier.
15. **Licence du modèle** : CC-BY-NC ou Llama Community mal lue = risque juridique en entreprise.
16. **`:latest` en prod** : épingler versions d'images et `model-id` (reproductibilité).
17. **DeepSeek Harness en preview** : changements cassants possibles — labo, pas prod.
18. **Comparer sans éval** : le « meilleur modèle » n'existe pas hors de TON jeu de questions.

---

# CAS PRATIQUES

## 117. Cas 1 — « Je veux un RAG perso à 0 € »

1. TEI + `bge-m3` en Docker sur ton Proxmox (section 66) → embeddings.
2. Indexer ton corpus (scripts existants), top-50 → reranker `bge-reranker-v2-m3` → top-5.
3. Génération : Ollama `qwen3:32b` si tu as une 4090, sinon routeur section 102 (Groq → Cerebras → OpenRouter free).
4. Cache SQLite des Q/R (section 104).
5. Coût mensuel : **0 €** (+ élec existante). Qualité : à éval
...[truncated 8112 chars]
---

# PENSE-BÊTE (UNE PAGE)

## 123. Pense-bête : endpoints et commandes

```
OpenRouter   https://openrouter.ai/api/v1              clé sk-or-v1-...   GET /api/v1/models = vérité prix
Groq         https://api.groq.com/openai/v1            clé gsk_...        30 RPM free, jamais facturé (429)
Cerebras     https://api.cerebras.ai/v1                clé csk-...        1M tok/j free, 8K ctx cap free
HF Providers https://router.huggingface.co/v1          token hf_...       $0.10/mois free
Ollama       http://localhost:11434/v1                 (pas de clé)       ollama run qwen3:32b
vLLM         http://localhost:8000/v1                  (pas de clé)       --enable-prefix-caching
TEI embed    http://localhost:8080/v1/embeddings       (pas de clé)       bge-m3, 1024 dim
TEI rerank   http://localhost:8082/rerank              (pas de clé)       bge-reranker-v2-m3
DeepSeek H.  http://127.0.0.1:3080                     npx @deepseek-ai/dsh web
```

```
# Santé express de toute la stack locale
for p in 11434 8000 8080 8082; do
  curl -s -o /dev/null -w "port $p : %{http_code}\n" http://localhost:$p/health 2>/dev/null \
    || curl -s -o /dev/null -w "port $p : %{http_code}\n" http://localhost:$p/v1/models
done
nvidia-smi --query-gpu=name,memory.used,memory.total,utilization.gpu --format=csv
```

## 124. Pense-bête : VRAM

```
Poids ≈ params × octets  (FP16: 2 | Q8: 1 | Q4: 0,5)
+ KV cache ≈ 320 Ko/token (70B) → 32K ctx ≈ 10 Go
7B Q4 ≈ 5 Go  |  14B Q4 ≈ 9 Go  |  32B Q4 ≈ 20 Go  |  70B Q4 ≈ 40 Go
4090 24 Go → 32B Q4 + 16K ctx = sweet spot
```

## 125. Pense-bête : gratuit 2026 (quotas à re-vérifier)

```
Groq 30 RPM | Cerebras 1M tok/j | Gemini Flash ~1500 req/j (données réutilisables !)
OpenRouter :free 50 req/j (1000 si $10 chargés) | Mistral Experiment ~1 Md tok/mois @ 2 RPM
Codestral 30 RPM/2000 j | NVIDIA NIM ~40 RPM | Cohere 1000 req/mois (non commercial)
Cloudflare 10K neurons/j | HF $0.10/mois | DeepSeek 5M tokens offerts (one-shot)
```

## 126. Pense-bête : règles d'or

```
1. Gratuit d'abord, payant sur échec mesuré.      4. Jamais de secret réel sur instance louée.
2. 429 → retry + backoff, toujours.               5. Embeddings : jamais deux modèles dans un index.
3. Données sensibles → local ou ZDR, jamais free. 6. Tout prix : date de vérification notée.
```

---

# POUR ALLER PLUS LOIN

## 127. Feuille de route 4 semaines

- **Semaine 1** : comptes gratuits (Groq, Cerebras, Gemini, OpenRouter, Mistral) + script de comparaison (section 11) sur 20 questions de ton domaine.
- **Semaine 2** : TEI + bge-m3 en Docker, réindexation d'un échantillon, éval précision@5 vs text-embedding-3-small.
- **Semaine 3** : Ollama `qwen3:32b` (ou vLLM si GPU), routeur multi-provider (section 102), cache SQLite.
- **Semaine 4** : reranker TEI dans le pipeline, logging des coûts (section 103), décision : quoi en local, quoi en gratuit, quoi en payant.

## 128. Ressources officielles (vérifiées le 27/09/2026)

- OpenRouter : docs `openrouter.ai/docs`, modèles `openrouter.ai/models`, API `openrouter.ai/api/v1/models`
- Groq : `console.groq.com/docs` (modèles), `groq.com/pricing`
- Cerebras : `inference.cerebras.ai` (docs), `cerebras.ai/pricing` (à vérifier — montants relevés via sources tierces)
- FreeLLMAPI : `github.com/tashfeenahmed/freellmapi`, `freellmapi.co`
- DeepSeek Harness : `github.com/deepseek-ai/deepseek-harness`, doc `deepseek-harness.github.io/deepseek-harness`
- Ollama : `ollama.com`, catalogue `ollama.com/library`
- vLLM : `docs.vllm.ai`
- TEI : `github.com/huggingface/text-embeddings-inference`, spec API `huggingface.github.io/text-embeddings-inference`
- Hugging Face : `huggingface.co/docs/inference-providers`, `huggingface.co/pricing`
- Vast.ai : `vast.ai`, CLI `pip install vastai`
- RunPod : `runpod.io/docs`

## 129. Sujets à creuser ensuite (proposés)

1. **Fine-tuning LoRA** de Qwen3-32B sur tes docs d'onduleurs/réseau (Unsloth/LLaMA-Factory) — le niveau au-dessus du RAG.
2. **Éval RAG automatisée** (RAGAS) branchée sur ton pipeline TEI + reranker.
3. **Agents** : tester DeepSeek Harness sur une tâche réelle (génération de doc technique).
4. **Observabilité** : Prometheus + Grafana sur `/metrics` (vLLM, TEI) et coûts (`llm_usage.db`).
5. **Sécurité** : durcir l'exposition LAN (WireGuard + Caddy) pour l'équipe.

## 130. Historique des vérifications (ce guide)

| Date | Action |
|---|---|
| 27/09/2026 | Recherche web : OpenRouter, Groq, Cerebras, freellmapi, DeepSeek Harness, TEI, HF Inference Providers, Vast.ai, RunPod, free tiers, prix OpenAI/Anthropic |
| 27/09/2026 | Rédaction v1 — 130 sections |
| À faire (trimestriel) | Re-vérifier prix et quotas : ils bougent vite (promos, dépréciations de modèles) |

*Suite du guide — voir Partie M, sections 131–158.*

---

# PARTIE M — APPROFONDISSEMENTS

## 131. LiteLLM : la passerelle open source à héberger soi-même

En plus d'OpenRouter (SaaS) et FreeLLMAPI (gratuit), il existe **LiteLLM** (github.com/BerriAI/litellm, MIT) : un **proxy unifié open source** qui expose 100+ providers derrière une API OpenAI-compatible, avec **fallbacks, retry, cache (Redis), budgets par clé, et logs**.
C'est le choix « entreprise » du self-hosted : tu gardes le contrôle, tu branches tes clés (Groq, OpenAI, ton vLLM local…).

