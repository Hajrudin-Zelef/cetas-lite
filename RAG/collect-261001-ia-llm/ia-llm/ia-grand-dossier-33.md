---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-33
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Groq", "OpenAI"]
dates: ["2026-08-21", "2026-09-27"]
keywords: ["claude", "cost", "deepseek", "embedding", "embeddings", "gemini", "gpt-5.6", "latency", "llama", "luna", "mcp", "terra"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2265, 2414]
sha256: a163edb4175c4658335bf6069778f1ae0f717498ae2dc9455db4df47e5ceb4ca
---

# IA — Le grand dossier

**Économie** : pas de marge sur le token (tu paies le prix public du provider), mais **5,5 % de frais sur les achats de crédits** (min 0,80 $) et +5 % si tu apportes ta propre clé (BYO-key). **Limites** : les routes `:free` sont lentes/en file ; un provider peut logger tes prompts (vérifier chaque fiche modèle) ; un modèle peut disparaître du catalogue (cf. Groq, §3.3).

### 4.3. LiteLLM — la gateway auto-hébergée (le standard)

**LiteLLM** (BerriAI, MIT) existe en deux formes : **bibliothèque Python** (`import litellm` : 100+ providers derrière `litellm.completion(model="groq/llama-...", ...)`) et **serveur proxy** (le cas pro : un conteneur qui expose `:4000` en API compatible OpenAI). Version vérifiée : **v1.97.0** (21/08/2026) ; images Docker cosignées ; endpoint MCP Gateway (`/v1/mcp`) stable depuis la v1.85.

**Démarrage minimal** (Docker, 2 minutes) :

```bash
# 1. Fichier de config (voir 4.4 pour la version complète)
cat > litellm_config.yaml <<'EOF'
model_list:
  - model_name: "rapide"
    litellm_params:
      model: "groq/openai/gpt-oss-120b"
      api_key: "os.environ/GROQ_API_KEY"
EOF

# 2. Lancement
docker run -d --name litellm \
  -p 4000:4000 \
  -v $(pwd)/litellm_config.yaml:/app/config.yaml \
  -e GROQ_API_KEY \
  ghcr.io/berriai/litellm:1.97.0 \
  --config /app/config.yaml --port 4000

# 3. Test
curl -s http://localhost:4000/v1/chat/completions \
  -H "Authorization: Bearer ${LITELLM_MASTER_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"model":"rapide","messages":[{"role":"user","content":"ping"}]}'
```

Note : `api_key: "os.environ/GROQ_API_KEY"` — LiteLLM lit la variable d'environnement au démarrage. **Jamais de clé en clair dans le YAML** (le fichier finit souvent dans git).

### 4.4. Config YAML complète et réutilisable : routage coût/qualité/latence + fallbacks

C'est le cœur opérationnel de cette section. Trois **noms logiques** (`triage`, `raisonnement`, `local`) que ton code appelle ; la gateway décide où ça part vraiment.

```yaml
# litellm_config.yaml — passerelle multi-providers
# Secrets via variables d'environnement uniquement.
general_settings:
  master_key: "os.environ/LITELLM_MASTER_KEY"   # clé admin de la gateway
  database_url: "os.environ/DATABASE_URL"       # Postgres -> spend logs, clés virtuelles
  store_model_in_db: true
  # Observabilité : exporter vers Prometheus + logger JSON
  # success_callback: ["prometheus"]  # nécessite le package additionnel

model_list:
  # ── GROUPE "triage" : rapide et pas cher, 3 providers en round-robin ──
  - model_name: "triage"
    litellm_params:
      model: "groq/openai/gpt-oss-120b"          # ~0,15/0,60 $/1M, ~500 tok/s
      api_key: "os.environ/GROQ_API_KEY"
      rpm: 30                                    # limite convenue (free tier)
  - model_name: "triage"
    litellm_params:
      model: "cerebras/gpt-oss-120b"            # ~0,35/0,75 $/1M, ~2000 tok/s
      api_key: "os.environ/CEREBRAS_API_KEY"
  - model_name: "triage"
    litellm_params:
      model: "openrouter/deepseek/deepseek-v4.1-flash"  # ~0,15/0,60 $/1M off-peak
      api_key: "os.environ/OPENROUTER_API_KEY"

  # ── GROUPE "raisonnement" : qualité max, 2 providers ──
  - model_name: "raisonnement"
    litellm_params:
      model: "anthropic/claude-opus-4-8"        # 5,00/25,00 $/1M
      api_key: "os.environ/ANTHROPIC_API_KEY"
      max_tokens: 4096
  - model_name: "raisonnement"
    litellm_params:
      model: "openai/gpt-5.6-terra"             # 2,00/12,00 $/1M
      api_key: "os.environ/OPENAI_API_KEY"

  # ── GROUPE "local" : Ollama sur le LAN, coût nul, données sur site ──
  - model_name: "local"
    litellm_params:
      model: "ollama/qwen3:32b"                 # nom du modèle pullé dans Ollama
      api_base: "http://ollama-interne.lan:11434"

  # ── Embeddings : le RAG de Zelef ──
  - model_name: "embeddings"
    litellm_params:
      model: "openai/text-embedding-3-small"    # 0,02 $/1M, 1536 dims
      api_key: "os.environ/OPENAI_API_KEY"

# ── ROUTAGE : comment choisir dans un groupe ──
router_settings:
  routing_strategy: "usage-based-routing-v2"   # route selon latence+charge (autres: simple-shuffle, least-busy, latency-based-routing)
  num_retries: 3
  retry_after: 5                                # secondes entre retries
  # Cooldown : un deployment en échec répété est écarté N secondes
  cooldown_time: 60
  allowed_fails: 3
  # Politique de retry par type d'erreur
  retry_policy:
    TimeoutError: 2
    RateLimitError: 3
    ContentPolicyViolationError: 0              # ne pas réessayer un refus de modération
    ContextWindowExceededError: 0               # inutile : basculer plutôt (voir fallbacks)

# ── FALLBACKS : que faire quand un groupe échoue ──
fallbacks:
  # Fallbacks généraux (toute erreur)
  - "triage": ["raisonnement"]                  # si les 3 rapides tombent -> qualité
  - "raisonnement": ["triage"]                  # si les frontier tombent -> rapide (dégradé mais vivant)
  # Fallbacks par type d'erreur
  # - {"triage": ["raisonnement"]} avec litellm>=1.5x : voir litellm_params.fallbacks
# Nombre max de bascules par requête
# (défaut raisonnable : 3)

# ── BUDGETS : garde-fous anti-facture-surprise ──
# (clés virtuelles créées via l'UI :4000/ui ou l'API /key/generate)
# Exemple de clé d'équipe : budget 50 $/mois, modèles autorisés restreints.
```

**Stratégies de routage disponibles** (`router_settings.routing_strategy`) :

| Stratégie | Logique | Quand l'utiliser |
|---|---|---|
| `simple-shuffle` | Round-robin aléatoire | Répartir bêtement la charge entre clés |
| `usage-based-routing-v2` | Score = latence + charge + taux d'erreur récents | **Défaut recommandé** : adaptatif |
| `latency-based-routing` | Toujours le plus rapide mesuré | Chat temps réel |
| `cost-based-routing` | Toujours le moins cher du groupe | Batch, backfill RAG |
| `least-busy` | Le deployment le moins sollicité | Fort trafic homogène |

**Fallbacks fins** (au-delà du global) : LiteLLM supporte des fallbacks par **type d'erreur** — `contentPolicyFallbacks` (si refusé pour politique → autre famille de modèle), `contextWindowFallbacks` (si contexte dépassé → modèle à plus grand contexte, ex. Gemini 1M). En YAML natif du proxy, cela se déclare dans la section `fallbacks` étendue ou via `general_settings`; la forme exacte évolue vite — **vérifier la doc de ta version** (`docs.litellm.ai`, v1.97 au 27/09/2026).

### 4.5. Load balancing entre clés d'un même provider

Tu as 3 clés OpenAI (une par équipe) et tu veux répartir sans qu'une seule sature :

```yaml
model_list:
  - model_name: "gpt-partage"
    litellm_params: {model: "openai/gpt-5.6-luna", api_key: "os.environ/OPENAI_KEY_EQUIPE_A", rpm: 500}
  - model_name: "gpt-partage"
    litellm_params: {model: "openai/gpt-5.6-luna", api_key: "os.environ/OPENAI_KEY_EQUIPE_B", rpm: 500}
  - model_name: "gpt-partage"
    litellm_params: {model: "openai/gpt-5.6-luna", api_key: "os.environ/OPENAI_KEY_EQUIPE_C", rpm: 500}
router_settings:
  routing_strategy: "usage-based-routing-v2"
```

Chaque deployment porte son `rpm`/`tpm` : LiteLLM met en cooldown celui qui sature et répartit sur les autres. **Effet secondaire précieux** : ça absorbe aussi les baisses de quota d'un provider sans changer une ligne de code applicatif.

### 4.6. Observabilité : savoir ce que coûte chaque requête

