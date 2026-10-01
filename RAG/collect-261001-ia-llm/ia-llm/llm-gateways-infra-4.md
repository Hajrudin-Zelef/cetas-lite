---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-4
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cerebras", "Cohere", "DeepSeek", "Google", "Groq", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Z.ai", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agent", "astra", "benchmarks", "claude", "cohere", "deepseek", "embeddings", "gemini", "glm", "gpt-6", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [366, 474]
sha256: e5f966d65ed254bfdd5e4488b2efd2a10c086a51e61772442a840f05796488b0
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Fournisseur | Modèle repère | In / Out par 1M | Latence typique | Free tier | Idéal pour |
|---|---|---|---|---|---|
| OpenRouter | `deepseek-v4.1-flash` | $0.15 / $0.60 | Moyenne (routage) | 50 req/j | Comparer, fallback |
| OpenRouter | `glm-5.3-flash` (promo) | $0.075 / $0.25 | Moyenne | via `:free` | Pas cher + contexte 1,3M |
| Groq payant | `gpt-oss-20b` | $0.075 / $0.30 | **~1 000 tok/s** | 30 RPM | Temps réel, triage |
| Groq payant | `gpt-oss-120b` | $0.15 / $0.60 | ~500 tok/s | 30 RPM | Agent rapide costaud |
| Cerebras payant | `llama-3.1-8b` | $0.10 / $0.10 | **~2 000 tok/s** | 1M tok/j | Haut volume simple |
| Cerebras payant | `gpt-oss-120b` | $0.35 / $0.75 | ~2 000 tok/s | 1M tok/j | Vitesse max gros modèle |
| OpenAI direct | `gpt-6-luna` | $0.10 / $0.50 | Moyenne | Non | Volume raisonné fermé |
| OpenAI direct | `gpt-6-sol` | $2 / $10 | Moyenne | Non | Code / raisonnement |
| OpenAI direct | `gpt-6-astra` | $10 / $50 | Moyenne | Non | Frontière |
| Anthropic direct | `claude-haiku-4.5` | $1 / $5 | Faible | Non | Rapide + économe fermé |
| Anthropic direct | `claude-sonnet-5` | $2 / $10 | Moyenne | Non | Code agentique |
| Anthropic direct | `claude-opus-5.5` | $4 / $20 | Moyenne-haute | Non | Raisonnement dur |
| DeepSeek direct | `deepseek-v4.1-flash` | $0.15 / $0.60 | Moyenne | 5M tokens offerts (inscription) | Rapport qualité/prix |
| Local (Ollama) | `qwen3:32b` Q4 | $0 (élec) | 10–60 tok/s (RTX 4090) | ∞ | Confidentiel, volume |

## 31. Arbre de décision (angle budget Zelef)

1. **Question simple / triage / résumé** → Groq free (`gpt-oss-20b`) ou Cerebras free (`llama-3.1-8b`). Coût : 0.
2. **Besoin de comparer des modèles** → OpenRouter, script section 11. Coût : quasi 0 avec `:free`.
3. **Raisonnement difficile, code complexe** → DeepSeek V4.1 (direct ou OpenRouter, $0.15/$0.60) d'abord ; si insuffisant, `gpt-6-sol` ($2/$10) ou `claude-sonnet-5` ($2/$10).
4. **Temps réel / voix** → Groq ou Cerebras (free puis payant).
5. **Gros volume répétitif** (> 50M tokens/mois) → calculer le seuil du local (section 134) : souvent Ollama/vLLM sur RTX 4090.
6. **Données sensibles / hors-ligne** → local obligatoire, pas de débat.
7. **Embeddings** → TEI local (partie G), pas l'API OpenAI.

## 32. Le pattern « routeur intelligent » (recommandé pour ton RAG)

```python
def choisir_modele(question: str, budget: str = "eco") -> tuple[str, str]:
    """Retourne (base_url, model). Logique : gratuit d'abord, payant si besoin."""
    if budget == "eco":
        # 1. Triage / simple -> Groq gratuit
        return ("https://api.groq.com/openai/v1", "openai/gpt-oss-20b")
    # 2. Standard -> Cerebras gratuit (1M tok/j)
    # 3. Complexe -> OpenRouter DeepSeek V4.1 ($0.15/$0.60)
    # 4. Très complexe -> gpt-6-sol ou claude-sonnet-5
```

En pratique : un petit classifieur (lui-même sur modèle gratuit) estime la difficulté, puis route. Les systèmes multi-fournisseurs sérieux ajoutent : **cache sémantique** (même question → réponse stockée, coût 0), **déduplication**, et **limites par utilisateur/jour**.

## 33. Coût réel d'une question RAG (exemple chiffré)

Hypothèses : 5 chunks de 800 tokens récupérés + prompt système 500 tokens + question 100 tokens = **4 600 tokens in**, réponse **400 tokens out**.

| Fournisseur / modèle | Coût par question | 1 000 questions/jour | 30 jours |
|---|---|---|---|
| Groq free / Cerebras free | $0 | $0 | $0 |
| DeepSeek V4.1 ($0.15/$0.60) | $0.00093 | $0.93 | ~$28 |
| gpt-6-luna ($0.10/$0.50) | $0.00066 | $0.66 | ~$20 |
| gpt-6-sol ($2/$10) | $0.0132 | $13.20 | ~$396 |
| claude-opus-5.5 ($4/$20) | $0.0264 | $26.40 | ~$792 |
| Local Ollama (élec 300W à 0.20 €/kWh) | ~$0.00006 | $0.06 | ~$2 |

Lecture : le gratuit absorbe des milliers de questions/jour ; le saut vers un modèle frontière se justifie uniquement si la qualité l'exige **mesurée** (éval, pas intuition).

## 34. Comparer honnêtement : la méthode

1. Fige **20–50 questions réelles** de ton domaine (pas des benchmarks génériques).
2. Pour chaque modèle : même prompt système, même `temperature=0`, même contexte.
3. Note : exactitude (barème 0/1/2), hallucinations, format respecté, tokens consommés, latence.
4. Calcule le **coût par bonne réponse** = coût total / nombre de réponses correctes. C'est le seul KPI qui compte.
5. Re-teste trimestriellement : les modèles et prix bougent vite (cf. promos GLM/Mercury).

---

# PARTIE E — FREELLMAPI ET L'ÉCOSYSTÈME GRATUIT RÉEL

## 35. « freellmapi » : identification — C'EST UN VRAI PROJET

Bonne nouvelle : `freellmapi` existe. C'est **FreeLLMAPI** (github.com/tashfeenahmed/freellmapi, licence MIT, site freellmapi.co) : un **proxy auto-hébergé compatible OpenAI** qui **empile les tiers gratuits de ~28 fournisseurs** (~1,7 à 7 Md tokens/mois cumulés selon la version du catalogue) derrière **un seul endpoint `/v1`**.
Principe : tu y enregistres tes clés gratuites (Google, Groq, Cerebras, Mistral, OpenRouter, GitHub Models, Cohere, Cloudflare, HuggingFace, Z.ai, NVIDIA, OVH AI Endpoints, Pollinations, LLM7, Kilo, OpenCode Zen, AI Horde…), le routeur choisit le meilleur modèle dispo par requête, **bascule automatiquement** au suivant en cas de rate-limit (429), et suit la conso par clé pour rester sous chaque plafond gratuit. Clés stockées **chiffrées**. Tu peux aussi brancher tes endpoints locaux (Ollama, llama.cpp, vLLM).
Modèle : open source gratuit en self-hosted (image Docker `ghcr.io/freellmapi`) ; offre hébergée « premium » à **$19/an** (catalogue live auto-à-jour). Mention légale du projet : **« personal experimentation only »** — pas pour de la prod commerciale.

## 36. FreeLLMAPI : installation (Docker)

```bash
# 1. Récupérer l'image (vérifié : ghcr.io/freellmapi via le README officiel)
docker pull ghcr.io/hatch/freellmapi 2>/dev/null || docker pull ghcr.io/freellmapi/freellmapi
# En pratique, suivre le README du repo pour le nom exact d'image (il a pu changer) :
# https://github.com/tashfeenahmed/freellmapi

# 2. Lancer (exemple type — adapter aux variables du README)
docker run -d --name freellmapi \
  -p 8787:8787 \
  -v freellmapi-data:/data \
  --restart unless-stopped \
  ghcr.io/freellmapi/freellmapi:latest

# 3. Ouvrir l'UI locale, y coller tes clés gratuites (Groq, Gemini, Cerebras...)
# 4. Appeler comme une API OpenAI :
curl http://localhost:8787/v1/chat/completions \
  -H "Authorization: Bearer flm_FAKEFAKE" \
  -H "Content-Type: application/json" \
  -d '{"model": "auto", "messages": [{"role": "user", "content": "Bonjour"}]}'
```

Le catalogue de modèles se met à jour via un flux signé (pas besoin de `git pull` pour les nouveaux modèles gratuits). Vérifie toujours le README pour les variables d'environnement exactes — le projet évolue vite.

## 37. FreeLLMAPI : intérêts et limites pour toi

**Intérêts** : un seul endpoint pour tout ton gratuit ; failover auto quand Groq te met un 429 à 14h ; suivi de conso par provider ; tu peux y brancher ton Ollama local comme « provider » de plus.
**Limites** : c'est un projet communautaire (pas de SLA) ; les quotas gratuits restent ceux des fournisseurs (il ne les augmente pas, il les **additionne**) ; « personal experimentation only » → pas pour servir ton équipe en prod ; chaque clé gratuite a ses propres **conditions d'utilisation des données** (section 42) — le proxy ne les efface pas.

## 38. L'écosystème gratuit réel en 2026 (vérifié le 27/09/2026)

Tableau des tiers **permanents, sans carte bancaire** (les chiffres bougent — re-vérifier avant de dimensionner) :

