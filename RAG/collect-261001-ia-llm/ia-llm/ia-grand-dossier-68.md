---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-68
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "DeepSeek", "Hugging Face", "Lambda", "Mistral", "OpenAI", "United States", "vLLM"]
dates: []
keywords: ["agents", "apache", "awq", "aws", "deepseek", "embedding", "embeddings", "fine-tuning", "gpu", "inference", "memory", "mistral"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5608, 5770]
sha256: 5c1e63c4e82b5d3ff67564069bf797b097599594bfc5db7d4166820a27ab2a7b
---

# IA — Le grand dossier

- **Tech** : moteur **FireAttention** + FireOptimizer (auto-tuning), speculative decoding, batching workload-aware.
- **Prix** : 0,20–3,00 $/1M typique ; tiers Standard/Priority/Fast (ex. DeepSeek V4 Pro : 1,74/3,48 $ en Standard, 2,61/5,22 $ en Priority) ; batch -50 %, prompt caching -50 % (cumulables).
- **Points forts** : function calling, JSON garanti, grammaires contraintes — pensés pour les **agents** ; clients : Cursor, Notion, DoorDash, Quora.
- **Crédit** : 1 $ offert à l'inscription.
- **Quand l'utiliser** : agents en production où la fiabilité et la latence priment sur le prix plancher.

### A.7. DeepInfra — le prix plancher

- **Depuis** : 2022. **Position** : minimiser le coût unitaire de l'inférence open-weight (~77 modèles, infra US propre dont B200).
- **Prix** : 0,02–1,50 $/1M ; batch -50 % (24 h) ; ex. GPT-OSS-120B ~0,09/0,45 $/1M (parmi les moins chers).
- **Conformité** : SOC 2, ISO 27001, GDPR, **HIPAA** — rare à ce niveau de prix.
- **Quand l'utiliser** : bulk (résumés de masse, labeling, evals, backfill RAG). Pas pour : la latence extrême.

### A.8. RunPod — le GPU du peuple (dev)

- **Offres** : Pods (VM GPU), Serverless (scale-to-zero, cold starts <200 ms), Clusters ; facturation à la seconde ; **egress 0 $**.
- **Prix (sept. 2026)** : H100 2,69–3,29 $/h ; 4090 ~0,44–0,74 $/h ; A100-80 ~1,89 $/h (Secure Cloud).
- **Points forts** : variété GPU, API-first, templates communautaires (vLLM en 1 clic).
- **Quand l'utiliser** : dev/expérimentation, inférence serverless, petits clusters. Le premier à tester avant d'acheter du GPU.

### A.9. Lambda Labs — le neocloud sérieux

- **Offres** : instances on-demand, réservées, 1-Click Clusters, Private Cloud ; stack ML optimisée ; support d'ingénieurs ML.
- **Prix** : H100 2,99–3,99 $/h ; A100-80 ~1,79 $/h ; **B200 6,69 $/h — le moins cher publié** (août 2026) ; egress 0 $ ; pas de spot.
- **Quand l'utiliser** : entraînement et inférence « sans surprise » quand RunPod/Vast semblent trop artisanaux et AWS trop cher.

### A.10. Vast.ai — le moins cher (marketplace)

- **Modèle** : marketplace P2P (particuliers + data centers), prix au spot, disponibilité variable.
- **Prix (sept. 2026)** : 4090 dès ~0,14 $/h (spot) ; H100 ~1,47–2,60 $/h ; egress 0 $.
- **Points durs** : l'hôte peut couper (spot) ; qualité réseau inégale ; à réserver au batch tolérant aux pannes, pas à la prod interactive.
- **Quand l'utiliser** : fine-tuning pas cher, expérimentation, inférence batch avec checkpoints fréquents.

### A.11. Hugging Face — l'écosystème

- **Offres** : Hub (1M+ modèles) ; **Inference Providers** (routage vers des hébergeurs partenaires derrière l'API HF) ; Inference Endpoints (dédiés, à l'heure) ; Jobs (conteneurs GPU one-shot).
- **Prix** : Free/PRO (9 $/mois) ; endpoints à l'heure d'instance ; serverless au token selon provider routé.
- **Quand l'utiliser** : tout ce qui touche aux poids ouverts y transite ; endpoints pour déployer un modèle du Hub en 10 minutes.

### A.12. Mistral AI — l'option européenne

- **Double jeu** : poids ouverts (Apache 2.0/MIT : Small, Nemo, Devstral, Ministral…) + **La Plateforme** (API hébergée UE).
- **Prix** : Large 3 : 0,50/1,50 $/1M ; Medium 3.5 : 1,50/7,50 $ ; Small ~0,10/0,30 $ ; Nemo 0,02/0,04 $ (à vérifier).
- **Argument** : souveraineté UE, Le Chat, conformité aux appels d'offres publics européens.
- **Vigilance licences** : Medium 3.5 = MIT modifiée (plafond de revenus) — lire avant usage commercial à grande échelle.

---

## Annexe B — Recettes d'exploitation : la stack complète en Docker Compose

### B.1. `docker-compose.yml` — gateway + observabilité + local + vectoriel

```yaml
# Stack IA hybride — à adapter (ports, volumes, ressources).
# Secrets UNIQUEMENT via fichier .env (jamais dans git).
name: ia-hybride

services:
  postgres:
    image: postgres:16
    environment:
      POSTGRES_DB: litellm
      POSTGRES_USER: litellm
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    volumes: [pgdata:/var/lib/postgresql/data]
    restart: unless-stopped

  litellm:
    image: ghcr.io/berriai/litellm-database:1.97.0
    command: ["--config", "/app/config.yaml", "--port", "4000", "--num_workers", "4"]
    ports: ["4000:4000"]
    volumes: ["./litellm_config.yaml:/app/config.yaml:ro"]
    environment:
      DATABASE_URL: "postgresql://litellm:${POSTGRES_PASSWORD}@postgres:5432/litellm"
      LITELLM_MASTER_KEY: ${LITELLM_MASTER_KEY}
      OPENAI_API_KEY: ${OPENAI_API_KEY}
      ANTHROPIC_API_KEY: ${ANTHROPIC_API_KEY}
      GROQ_API_KEY: ${GROQ_API_KEY}
      CEREBRAS_API_KEY: ${CEREBRAS_API_KEY}
      OPENROUTER_API_KEY: ${OPENROUTER_API_KEY}
    depends_on: [postgres]
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:4000/health/readiness"]
      interval: 30s

  ollama:
    image: ollama/ollama:0.34.0
    ports: ["11434:11434"]
    volumes: [ollama:/root/.ollama]
    deploy: {resources: {reservations: {devices: [{driver: nvidia, count: 1, capabilities: [gpu]}]}}}
    restart: unless-stopped
    # Au premier démarrage : docker exec ollama ollama pull qwen3:8b

  vllm:
    image: vllm/vllm-openai:latest
    ports: ["8000:8000"]
    volumes: [hf-cache:/root/.cache/huggingface]
    command: >
      --model Qwen/Qwen3-32B-AWQ --served-model-name qwen3-32b
      --host 0.0.0.0 --port 8000 --api-key ${VLLM_API_KEY}
      --gpu-memory-utilization 0.90 --max-model-len 32768
      --enable-prefix-caching --enable-auto-tool-choice --tool-call-parser qwen3_coder
    deploy: {resources: {reservations: {devices: [{driver: nvidia, count: 1, capabilities: [gpu]}]}}}
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s

  qdrant:
    image: qdrant/qdrant:latest
    ports: ["6333:6333"]
    volumes: [qdrant:/qdrant/storage]
    restart: unless-stopped

  open-webui:
    image: ghcr.io/open-webui/open-webui:main
    ports: ["3000:8080"]
    environment:
      OLLAMA_BASE_URL: http://ollama:11434
      OPENAI_API_BASE_URL: http://litellm:4000/v1   # tout le cloud via la gateway !
      OPENAI_API_KEY: ${LITELLM_TEAM_KEY}            # clé virtuelle d'équipe
    volumes: [webui:/app/backend/data]
    depends_on: [ollama, litellm]
    restart: unless-stopped

volumes: {pgdata: {}, ollama: {}, hf-cache: {}, qdrant: {}, webui: {}}
```

### B.2. Client Python unifié (une seule base URL pour tout)

```python
# client_ia.py — tout le SI parle à la gateway, jamais aux providers en direct
import os
from openai import OpenAI

gw = OpenAI(base_url="http://litellm-interne.lan:4000/v1",
            api_key=os.environ["LITELLM_TEAM_KEY"])  # clé virtuelle d'équipe (budget plafonné)

def ask(modele_logique: str, prompt: str, system: str = "", **kw) -> str:
    msgs = ([{"role": "system", "content": system}] if system else []) \
           + [{"role": "user", "content": prompt}]
    r = gw.chat.completions.create(model=modele_logique, messages=msgs,
                                   timeout=120, **kw)
    return r.choices[0].message.content

def ask_avec_escalade(prompt: str, system: str = "") -> tuple[str, str]:
    """Triage local/pas cher d'abord, escalade vers 'raisonnement' si confiance basse."""
    import json
    brut = ask("triage", prompt,
               system + "\nTermine ta réponse par une ligne JSON : {\"confiance\": 0.0-1.0}")
    try:
        conf = float(json.loads(brut.strip().splitlines()[-1])["confiance"])
    except Exception:
        conf = 0.0
    if conf < 0.7:
        return ask("raisonnement", prompt, system), "raisonnement"
    return brut, "triage"

def embed(textes: list[str]) -> list[list[float]]:
    r = gw.embeddings.create(model="embeddings", input=textes)
    return [d.embedding for d in r.data]

