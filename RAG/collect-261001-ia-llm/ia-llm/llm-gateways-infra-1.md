---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-1
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Google", "Groq", "Hugging Face", "Meta", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Z.ai", "vLLM", "xAI"]
dates: ["2026-08-21", "2026-09-10", "2026-09-27"]
keywords: ["gpu", "agents", "attention", "claude", "cohere", "deepseek", "embedding", "embeddings", "gemini", "glm", "gpt-6", "inference"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1, 102]
sha256: 76c4bc10a05ad2d68f9736d21d1d23cac45c8cbddb1fb605586832b96ac40220
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

> Guide terrain pour Zelef — chef de service systèmes & énergies, sysadmin, RAG personnel (embeddings `text-embedding-3-small` d'OpenAI).
> Objectif : accéder aux LLM au meilleur coût (idéalement 0 €), comprendre les passerelles, savoir héberger soi-même, et dimensionner l'infra.
> Données vérifiées par recherche web le **27/09/2026**. Tout prix non re-vérifié ce jour est marqué « à vérifier ». Aucune clé réelle : exemples fictifs (`sk-...FAKE`).
> Conventions : prix en $/1M tokens sauf mention ; TPS = tokens/seconde ; ctx = fenêtre de contexte.

---

## 1. Comment lire ce guide

1. **Parties A–C** : les trois passerelles rapides (OpenRouter, Groq, Cerebras) — fiches complètes : principe, modèles vérifiés sept 2026, prix, exemples d'appels, limites.
2. **Partie D** : le comparatif qui tranche — quand utiliser quoi, avec ton angle budget.
3. **Partie E** : `freellmapi` (identifié, c'est un vrai projet) + tout l'écosystème **gratuit réel** en 2026, limites chiffrées.
4. **Partie F** : `DeepSeek Harness` (identifié, vrai outil sorti août 2026) + **self-hosting DeepSeek/LLM local** en profondeur : Ollama, vLLM, llama.cpp, quantization, dimensionnement VRAM.
5. **Partie G** : TEI (Text Embeddings Inference) — servir tes embeddings en local, lien direct avec ton RAG.
6. **Partie H** : Hugging Face (Hub, Inference Providers, transformers, datasets, Spaces) — trouver et tester un modèle pour ton RAG.
7. **Partie I** : Vast.ai et RunPod — louer du GPU au lieu d'acheter, pas à pas, comparatif, pièges.
8. **Parties J–L** : stratégie budget, noms non identifiés (honnêteté), glossaire, quiz, pièges, cas pratiques.

## 2. Le paysage en une page (sept 2026)

| Couche | Acteurs | Coût typique | Quand |
|---|---|---|---|
| Passerelle multi-modèles | OpenRouter | 0 à $10/$50 /1M selon modèle | Comparer N modèles, fallback auto |
| Inférence ultra-rapide | Groq (LPU), Cerebras (WSE) | Gratuit (quotas) puis $0.05–$0.85 /1M in | Temps réel, agents, voix |
| Agrégateur de gratuit | freellmapi (self-hosted) | 0 (tiers gratuits empilés) | Expérimentation perso |
| API directes | OpenAI, Anthropic, DeepSeek, Mistral, Google | $0.10 à $50 /1M | Qualité max, SLA, support |
| Local | Ollama, vLLM, llama.cpp | 0 (ton élec + GPU) | Confidentialité, volume, hors-ligne |
| Embeddings locaux | TEI + bge/jina | 0 | Ton RAG sans facture OpenAI |
| GPU à la demande | Vast.ai, RunPod | $0.13–$3.89 /h | Fine-tune, batch, tests lourds |

## 3. Règle d'or du budget (ta contrainte)

1. Commence toujours par le **gratuit** : Groq free tier + Gemini Flash + OpenRouter `:free` + Cerebras free tier couvrent 95 % des usages d'étude et de dev.
2. Ne paie que ce que le gratuit ne fait pas : gros contextes, modèles fermés (GPT-6, Claude), SLA.
3. Mesure avant d'optimiser : logge tokens in/out par appel dès le jour 1 (section 133).
4. Le local (Ollama/TEI) bat tout le monde au-delà d'un certain volume — le seuil se calcule (section 134).

---

# PARTIE A — OPENROUTER

## 4. OpenRouter : présentation

OpenRouter (openrouter.ai) est une **passerelle unifiée** : une seule clé API, un seul endpoint compatible OpenAI (`https://openrouter.ai/api/v1`), et derrière **plus de 300 modèles** de dizaines de fournisseurs (OpenAI, Anthropic, Google, DeepSeek, Meta, Mistral, Qwen, Z.ai, NVIDIA, xAI…).
Tu changes de modèle en changeant une chaîne de caractères — aucun changement de code. C'est l'outil idéal pour **comparer des modèles** sur ton cas d'usage réel et pour mettre en place des **fallbacks** automatiques.
Modèle économique : prépayé (crédits), facturation au token, petite marge au-dessus du prix fournisseur. Pas d'abonnement obligatoire.

## 5. OpenRouter : principe technique

- Endpoint : `POST https://openrouter.ai/api/v1/chat/completions` — même format que l'API OpenAI.
- Auth : `Authorization: Bearer <ta-clé>`.
- Le nom du modèle suit le format `fournisseur/modele` : `anthropic/claude-sonnet-4`, `deepseek/deepseek-v4.1-flash`, `openai/gpt-oss-20b`, `google/gemma-4-31b-it`.
- Variante `:free` : certains modèles ont une route gratuite (`z-ai/glm-5.2:free`), soumise à quotas.
- Route spéciale `openrouter/free` : routeur qui choisit un modèle gratuit au hasard par requête.
- Routage fournisseur : par défaut OpenRouter choisit le fournisseur (le plus rapide / le moins cher) ; tu peux forcer ou ordonner les fournisseurs via le paramètre `provider`.
- Liste des modèles en direct : `GET https://openrouter.ai/api/v1/models` (pas de clé requise) — c'est LA source de vérité pour prix et disponibilité.

## 6. OpenRouter : modèles disponibles vérifiés (sept 2026)

Relevé de sources tierces ayant interrogé l'API live fin août / sept 2026 (prix susceptibles d'évoluer — vérifier sur `openrouter.ai/models/<id>`) :

| Modèle (id) | Contexte | Prix in/out par 1M (vérifié fin août 2026) | Note |
|---|---|---|---|
| `deepseek/deepseek-v4.1-flash` | 1M | $0.15 / $0.60 (hors peak) | Sorti 10/09/2026 ; peak en semaine 01h–04h et 06h–10h UTC : $0.30/$1.20 |
| `deepseek/deepseek-v4-flash-vision-exp` | 1M | $0.22 / $0.66 | Vision expérimentale |
| `qwen/qwen3.8-flash` | 1M | $0.15 / $0.47 | |
| `z-ai/glm-5.3-flash` | 1,3M | **$0.075 / $0.25** (promo lancement) | Prix cassé temporaire |
| `z-ai/glm-5.2` | 1M | $0.95 / $3.00 | |
| `z-ai/glm-4.6` | 203K | $0.50 / $2.00 | |
| `inception/mercury-2.5` | 260K | $0.04 / $0.15 (promo 80 % au lancement) | Tarif normal $0.20/$0.75 — promo ≠ permanent |
| `google/gemma-4-26b-a4b-it` | 262K | $0.06 / $0.33 (route dekallm/bf16) | |
| `mistralai/mistral-small-4` | ctx à vérifier | $0.15 / $0.60 | Bon support JSON/schema |
| `nvidia/nemotron-3.5-lightning` | 1M | $0.08 / $0.20 (route BF16) | |
| `anthropic/claude-sonnet-4` | 200K | $3 / $15 | Référence code |
| `openai/gpt-4o` | 128K | $2.50 / $10 | Ancien mais stable |
| `meta-llama/llama-3.1-405b` | 128K | $3 / $3 | Poids ouverts, gros |

Points d'attention : les prix **varient selon le fournisseur routé** et selon l'heure (DeepSeek a des fenêtres peak). Les promos de lancement (GLM, Mercury) expirent.

## 7. OpenRouter : modèles gratuits (`:free`) vérifiés

Relevé live du 21/08/2026 (`pricing.prompt == "0"`) — la liste bouge, re-vérifier via l'API :

| id | Contexte | Note |
|---|---|---|
| `nvidia/nemotron-3-ultra-550b-a55b:free` | 1M | 550B MoE, le plus gros modèle gratuit du catalogue |
| `z-ai/glm-5.2:free` | 256K | |
| `nvidia/nemotron-3.5-lightning:free` | 1M | |
| `openai/gpt-oss-20b:free` | 131K | Poids ouverts OpenAI |
| `google/gemma-4-31b-it:free` | 262K | |
| `cohere/north-mini-code:free` | 256K | Code |
| `poolside/laguna-s-2.1:free` | 262K | Code |
| `dots-studio/dots-3-note-preview:free` | 512K | |
| `openrouter/free` | 200K | Routeur aléatoire entre modèles gratuits |
| `stealth/ox-alpha` | 1M | Reasoning obligatoire |

**Quotas free** (vérifié sept 2026, sources concordantes) : **50 requêtes/jour** sur les modèles gratuits, **20 req/min** ; passe à **1 000 req/jour** si le compte a déjà acheté au moins $10 de crédits (achat unique, effet permanent). Limite au niveau du compte, tous modèles gratuits confondus.

## 8. OpenRouter : tarification et facturation

