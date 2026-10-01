---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-12
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cerebras", "DeepSeek", "Groq", "OpenAI", "OpenRouter", "vLLM"]
dates: ["2026-09-27"]
keywords: ["awq", "deepseek", "gpt-6", "qwen", "reranker", "sol", "vllm"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1459, 1603]
sha256: 1ebe1394ae292648038ec7fd6f478cebe72a31ea781da3e8040809c0c3d7a77a
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

```
┌─────────────┐     ┌──────────────┐     ┌────────────────┐
│ Indexation  │────▶│ TEI (bge-m3) │────▶│ Chroma/pgvector│  (local, 0 €)
│ (batch)     │     │  :8080       │     │                │
└─────────────┘     └──────────────┘     └───────┬────────┘
                                                │ top-50
                                         ┌──────▼────────┐
                                         │ Reranker TEI  │  (local, 0 €)
                                         │  :8082        │
                                         └──────┬────────┘
                                                │ top-5 + question
                    ┌───────────────────────────▼───────────────────────────┐
                    │ Génération : ordre de tentative                        │
                    │ 1. vLLM local (Qwen3-32B-AWQ) — 0 €                   │
                    │ 2. Groq free (gpt-oss-120b) — 0 €                      │
                    │ 3. Cerebras free (1M tok/j) — 0 €                      │
                    │ 4. OpenRouter DeepSeek V4.1 ($0.15/$0.60) — ~0 €       │
                    │ 5. gpt-6-sol / sonnet-5 — payant, cas durs uniquement  │
                    └───────────────────────────────────────────────────────┘
```

## 102. Le routeur de génération (code complet)

```python
import os, time
from openai import OpenAI

PROVIDERS = [
    ("http://localhost:8000/v1", "Qwen/Qwen3-32B-AWQ", "vllm-local"),
    ("https://api.groq.com/openai/v1", "openai/gpt-oss-120b", "groq"),
    ("https://api.cerebras.ai/v1", "gpt-oss-120b", "cerebras"),
    ("https://openrouter.ai/api/v1", "deepseek/deepseek-v4.1-flash", "openrouter"),
]
KEYS = {
    "groq": os.environ.get("GROQ_API_KEY_FAKE", "gsk_FAKE"),
    "cerebras": os.environ.get("CEREBRAS_API_KEY_FAKE", "csk_FAKE"),
    "openrouter": os.environ.get("OPENROUTER_API_KEY_FAKE", "sk-or-v1_FAKE"),
    "vllm-local": "local",
}

def generer(messages, max_tokens=800):
    last_err = None
    for base_url, model, name in PROVIDERS:
        try:
            client = OpenAI(base_url=base_url, api_key=KEYS[name], timeout=60)
            r = client.chat.completions.create(
                model=model, messages=messages,
                max_tokens=max_tokens, temperature=0.2)
            log_usage(name, model, r.usage)   # section 103
            return r.choices[0].message.content, name
        except Exception as e:
            last_err = e
            time.sleep(1)  # backoff simple ; en prod : exponentiel + jitter
    raise RuntimeError(f"Tous les providers ont échoué : {last_err}")
```

## 103. Logger chaque token (la base du contrôle budgétaire)

```python
import sqlite3, time
db = sqlite3.connect("llm_usage.db")
db.execute("""CREATE TABLE IF NOT EXISTS usage(
  ts REAL, provider TEXT, model TEXT,
  prompt_tok INTEGER, completion_tok INTEGER)""")

def log_usage(provider, model, usage):
    db.execute("INSERT INTO usage VALUES (?,?,?,?,?)",
        (time.time(), provider, model,
         usage.prompt_tokens, usage.completion_tokens))
    db.commit()

# Coût estimé sur 30 jours par provider (prix en dur, à maintenir)
PRIX = {  # (in, out) $/1M — vérifiés le 27/09/2026
    "deepseek/deepseek-v4.1-flash": (0.15, 0.60),
    "openai/gpt-oss-120b": (0.15, 0.60),
    "gpt-oss-120b": (0.35, 0.75),
}
```

Dashboard minimal : une requête SQL par semaine → coût par provider, dérive éventuelle, top modèles. **Sans mesure, pas d'optimisation.**

## 104. Cache sémantique : le multiplicateur de gratuit

```python
import hashlib, json
def cache_key(messages):
    return hashlib.sha256(json.dumps(messages).encode()).hexdigest()

# Avant d'appeler generer() : chercher en SQLite/Redis.
# Après : stocker (question_hash -> réponse, provider, date).
# TTL : 7 jours pour la doc technique (change peu), 1 h pour l'actu.
```

En pratique sur un RAG d'équipe : **20–40 % des questions sont des doublons** (FAQ, procédures). Le cache les sert à coût 0 et latence ~5 ms.

## 105. Prompt engineering économique (5 techniques)

1. **Raccourcir le prompt système** : chaque token système est payé à chaque appel — 500 tokens × 10K appels = 5M tokens « invisibles ».
2. **Limiter `max_tokens`** : une réponse de 200 tokens coûte 4× moins qu'une de 800.
3. **Baisser la température** (0–0.2) : réponses plus courtes et déterministes → moins de tokens, cache plus efficace.
4. **Choisir le contexte utile** : top-5 rerankés > top-20 bruts (moins de tokens in, meilleure qualité).
5. **Structurer la sortie** (JSON) : évite les pavés verbeux, parse direct.

## 106. Batch et files d'attente (quand le temps réel n'est pas requis)

- Indexation, résumés de logs, étiquetage : **pas besoin de temps réel** → batcher la nuit sur le moins cher (local, ou batch API −50 % chez Groq).
- File simple : table SQLite `jobs(status, payload)`, worker qui dépile avec le provider le moins cher disponible.
- Exemple : 100K tickets à classifier → Groq free `gpt-oss-20b` à 30 RPM ≈ 56 h… ou **vLLM local** en 2 h à coût élec. Le local gagne dès que le volume est là.

## 107. Monitoring : ce qu'il faut surveiller

| Signal | Où | Alerte si |
|---|---|---|
| Coût / jour par provider | `llm_usage.db` (section 103) | > seuil (ex. 5 €/j) |
| Taux de 429 | logs routeur | > 5 % → changer de provider/ordre |
| Latence p95 génération | logs | > 10 s → fallback trop lent |
| Santé TEI/vLLM/Ollama | `/health`, systemd | down → bascule cloud auto |
| Quotas free restants | dashboard fournisseurs | < 20 % → prévenir |
| Dérive qualité | éval hebdo (20 questions) | score < baseline → rollback modèle |

## 108. Sécurité des clés API (règles dures)

1. **Jamais** de clé en dur dans le code / Git — variables d'environnement ou gestionnaire de secrets (Vault, 1Password CLI, `pass`).
2. Fichier `.env` : dans `.gitignore` dès le premier commit (vérifier avec `git log --all -p | grep sk-` de temps en temps).
3. Une clé **par usage** (dev, prod, tests) avec **limites de dépense** (OpenRouter le permet nativement).
4. Rotation : si une clé fuite (log, screenshot, repo public) → révoquer immédiatement, pas « demain ».
5. En local (Ollama/vLLM/TEI) : pas de clé du tout — un argument de plus pour le local.
6. Sur instance louée (Vast/RunPod) : clés **jetables** à faible plafond, révoquées après le job.

## 109. Sauvegarde et reproductibilité de l'infra

- Versionner : `docker-compose.yml`, `Modelfile`, `.env.example` (sans secrets), scripts de provisionning.
- Épingler : versions d'images (`:cpu-1.8.3`, pas `:latest` en prod), `model-id` + révision HF, versions de drivers.
- Sauvegarder : la base de vecteurs (ou le script de réindexation — moins lourd), `llm_usage.db`, les Modelfiles.
- Documenter : un `INFRA.md` avec le schéma (section 101), les endpoints, les quotas et leurs dates de vérification.

## 110. Coûts électriques du local (ne pas l'oublier)

- RTX 4090 en charge ≈ 300–450 W ; PC complet ≈ 400–600 W.
- À 0.20 €/kWh : 500 W × 24 h = 12 kWh/j ≈ **2.40 €/jour** en 24/7, soit ~72 €/mois.
- En usage intermittent (8 h/j, 5 j/sem) : ~16 €/mois.
- **Le local 24/7 n'est « gratuit » que si la machine tourne déjà** (serveur Proxmox existant — ton cas). Sinon, comparer avec le gratuit cloud (0 €, 0 W).

---

