---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-21
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "ByteDance", "DeepSeek", "OpenAI"]
dates: ["2026-06-23", "2026-06-24", "2026-09-27"]
keywords: ["agents", "attention", "benchmark", "benchmarks", "claude", "deepseek", "distribution", "embedding", "embeddings", "gpt-6", "multimodal", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2498, 2586]
sha256: e450140725f890c345be9d866da02e23878ca751773bc9d43c4347928b605336
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

> **Rappel de doctrine :** une annonce n'est pas une disponibilité, une disponibilité n'est pas un prix, un prix n'est pas une éligibilité. Les quatre doivent être vérifiés séparément.

*Fin du guide — 159 sections. Section 159 rédigée le 27/09/2026 : seules les annonces avec source et date y figurent ; tout le reste est « rien d'officialisé » ou « RUMEUR non confirmée ». Prix « vérifiés le 27/09/2026 » par recherche web ; le reste « à vérifier » sur les pages officielles avant toute décision budgétaire.*

## 160. Doubao (ByteDance) : le challenger chinois qui compte

**En une phrase :** l'assistant et la famille de modèles de ByteDance (TikTok), exposés via la plateforme cloud Volcano Engine — le rival direct de DeepSeek sur le segment « API pas chère et agents », avec des prix publiés en yuans et un flagship, Seed 2.1 Pro, qui vise le créneau « coding + agents long-chain ».

### 160.1. Positionnement

- Doubao = la marque IA grand public de ByteDance (app mobile gratuite, leader en Chine en 2026). Les modèles s'appellent **Seed** (Doubao Seed 2.1 Pro / Turbo au 27/09/2026).
- Distribution API : **Volcano Engine** (la branche cloud de ByteDance), runtime **Ark**. C'est l'équivalent de « DeepSeek Platform » chez ByteDance : console, endpoints, facturation au token, compatible OpenAI.
- Positionnement prix : volontairement agressif, avec une particularité utile pour ton RAG — le **cache d'entrée facturé très bas** (¥1,20/M sur Seed 2.1 Pro, soit 5x moins cher que l'entrée normale). Si tes prompts RAG réutilisent le même gros préfixe système (instructions + quelques exemples), le cache change la facture.
- Particularité administrative : un compte Volcano Engine « Chine » exige **vérification d'identité réelle** (real-name verification) et recharge en CNY. Il existe une console internationale (Volcano Engine overseas) avec carte bancaire internationale — sinon, les agrégateurs tiers (ex. ofox.ai) exposent Seed 2.1 en USD sans compte Volcano, avec une marge au passage.

### 160.2. Modèles (état au 27/09/2026 — vérifié)

Annoncés le **24/06/2026** à la conférence Volcano Engine FORCE. Deux variantes, contexte **256K** sur les deux, texte + image en entrée, texte en sortie.

| Modèle | ID d'API (Volcano) | Positionnement |
|---|---|---|
| **Doubao Seed 2.1 Pro** | `doubao-seed-2.1-pro` | Flagship « deep thinking » : code complexe, agents long-chain, multimodal. ByteDance revendique des scores de tête sur Terminal Bench 2.1, SWE-Pro, SciCode (code) et OSWorld, MobileWorld, MMMU-Pro (agents) — chiffres vendeur, à vérifier sur tes évals |
| **Doubao Seed 2.1 Turbo** | `doubao-seed-2.1-turbo` | Moitié prix de Pro : trafic entreprise haute fréquence, latence faible |

Génération précédente toujours servie : Seed 2.0 Lite/Mini, Doubao 1.5 Pro 32K, Doubao 1.5 Lite (le moins cher de la gamme). Seed 2.0 Pro a été retiré du catalogue phare le 23/06/2026, remplacé par Seed 2.1 Pro — ne le code plus en dur.

### 160.3. Prix vérifiés le 27/09/2026 (Volcano Engine, CNY/M tokens)

| Modèle | Entrée | Sortie | Cache entrée | Contexte |
|---|---|---|---|---|
| Seed 2.1 Pro | ¥6 | ¥30 | **¥1,20** | 256K |
| Seed 2.1 Turbo | ¥3 | ¥15 | ¥0,60 | 256K |
| Seed 2.0 Pro (≤ 32K) | ¥3,20 | ¥16 | — | 32K (ancienne génération) |
| Doubao 1.5 Lite | ¥0,30 | ¥0,60 | — | le moins cher |

Points d'attention :
- Doubao facture **par tranche de longueur d'entrée** (le tarif Pro couvre 0–256K ; vérifie la page pricing si tu dépasses — les longs contextes peuvent avoir un étage tarifaire).
- Tier gratuit : **500K tokens par modèle** après vérification d'identité, valables 30 jours — suffisant pour un benchmark sérieux sans sortir la carte.
- Via agrégateur tiers ofox.ai (USD, sans compte Volcano, vérifié le 27/09/2026) : Pro $0,884/$4,42/M (cache $0,177), Turbo $0,442/$2,212/M (cache $0,085). Pratique pour tester, mais **ce sont des prix de revendeur** : la référence reste la page pricing Volcano.

### 160.4. Appel API — compatible OpenAI

Volcano Ark expose un endpoint OpenAI-compatible. Schéma minimal (clé fictive) :

```python
from openai import OpenAI

client = OpenAI(
    base_url="https://ark.cn-beijing.volces.com/api/v3",  # région à adapter
    api_key="sk-EXEMPLE-FICTIF",
)
resp = client.chat.completions.create(
    model="doubao-seed-2.1-pro",   # ou "doubao-seed-2.1-turbo"
    messages=[{"role": "user", "content": "Résume ce rapport de maintenance en 5 points."}],
)
print(resp.choices[0].message.content)
```

Règles pratiques :
- Les IDs de modèle côté Ark sont des **endpoints à créer dans la console** (chaque déploiement = un endpoint ID). En production, épingle l'endpoint ID, pas le nom marketing.
- Latence Chine → ton infra : si ton RAG tourne hors de Chine, mesure le TTFB réel avant de router du trafic synchrone dessus — un prix bas avec 800 ms de latence réseau ne convient pas à un chat interactif.

### 160.5. Quand l'utiliser dans ta stack

1. **Volume + budget** : si ton RAG génère beaucoup de réponses longues en français technique, Turbo à ¥3/¥15 vaut un benchmark contre DeepSeek V4.1 Flash et GPT-6 Sol sur tes 30 questions d'éval (section 127).
2. **Agents** : Seed 2.1 Pro vise explicitement les agents multi-étapes — à tester là où tu utiliserais Claude ou GPT-6 pour de l'orchestration.
3. **Cache agressif** : ton prompt système RAG est stable → active le cache d'entrée, c'est le poste où Doubao est le plus compétitif.
4. **À ne pas faire** : traiter les benchmarks vendeur (vs Claude Opus 4.6) comme une vérité ; ignorer la contrainte real-name verification si tu dois facturer à une entité hors Chine ; oublier que les prix sont en CNY (taux ~6,79 CNY/USD au 27/09/2026 — recalcule à chaque devis).

### 160.6. Embeddings et modèles annexes : ne pas oublier la brique vectorielle

Ton RAG vit ou meurt par l'embedding (tu es sur text-embedding-3-small). Doubao propose ses propres modèles d'embedding sur Volcano Ark (`doubao-embedding` — plusieurs générations, dont des versions « large » et des modèles vision). Deux usages :
- **Benchmark d'embedding** : si tu testes Doubao comme LLM de génération, teste aussi son embedding sur ton corpus FR technique — un embedding entraîné avec du chinois/anglais peut se comporter différemment sur ton français industriel. Mesure avec tes métriques (section 141 : recall@k, MRR), pas au feeling.
- **Reranker** : Volcano propose aussi des rerankers — l'étage « rerank » après la récupération vectorielle est souvent le gain qualité/coût le plus rentable d'un RAG (quelques centimes pour +10 points de précision).

Règle : ne mélange pas les embeddings entre index et requêtes — tout le corpus doit être réindexé si tu changes de modèle d'embedding. C'est l'opération la plus chère d'un changement de stack RAG : compte tes tokens d'indexation AVANT (section 167.5).

### 160.7. Doubao vs DeepSeek vs GPT-6 Sol : comment arbitrer (sept 2026)

| Critère | Doubao Seed 2.1 Pro | DeepSeek V4 Pro | GPT-6 Sol |
|---|---|---|---|
| Prix entrée /M | ¥6 (~$0,88) | ¥0,66 (~$0,10) hors pic | « moitié prix série 5.6 », non publié |
| Prix sortie /M | ¥30 (~$4,42) | ¥1,98 hors pic | non publié |
| Cache entrée | ¥1,20 (5x moins cher) | — | à vérifier |
| Contexte | 256K | 256K (V4 Pro) | à vérifier |
| Positionnement | Agents long-chain, code | Raisonnement brut, prix cassé | Écosystème OpenAI, prix agressif |
| Accès | Volcano (real-name) / agrégateurs | API directe simple | API OpenAI directe |
| Français technique | À benchmarker | À benchmarker | À benchmarker |

