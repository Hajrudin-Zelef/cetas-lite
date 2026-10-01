---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-14
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Mistral", "Moonshot", "OpenAI", "vLLM"]
dates: []
keywords: ["awq", "claude", "cohere", "deepseek", "embedding", "embeddings", "fine-tuning", "fp8", "gguf", "gpt-6", "gptq", "gqa"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1392, 1507]
sha256: d6662cf54590f9800c6c7ae0ae50ab7886ed6a1d214714c769770c92a8d3ec25
---

# Encyclopédie des modèles IA — Volume 3

1. **Français technique.** Le modèle doit produire un français propre avec du vocabulaire métier (pas « le UPS » mais « l'onduleur », pas de calques de l'anglais). Les modèles multilingues forts en 2026 : Qwen (3.5/3.6/3.8), Gemma 4 (140+ langues), Mistral (français natif de l'équipe), Claude (excellent en français, fermé).
2. **Contexte utile ≥ 32K.** Avec 10–20 chunks de 1000–1500 tokens + prompt système, tu es à 15–30K tokens par requête. 32K est le plancher confortable ; 128K si tu veux du rerank-large-then-stuff.
3. **Suivi d'instructions + citations.** Le modèle doit citer ses sources (n° de chunk) et dire « je ne sais pas » quand les chunks ne contiennent pas la réponse — ça s'évalue, ça ne se suppose pas.
4. **Prix/latence.** En API : fourchette 2026 de ~$0,14/M (Qwen3.6-35B tiers) à $4–5/M (flagships). En local : le coût est la carte, pas le token.
5. **Poids ouverts vs API.** Local = confidentialité totale (tes docs ne sortent pas), coût fixe, mais tu gères l'infra. API = zéro infra, mais dépendance + données chez un tiers.

## 100. Recommandation générateur 2026 (pour ton usage)

| Profil | Choix | Pourquoi |
|---|---|---|
| **API, qualité/prix FR** | Qwen3.6-35B-A3B (~$0,14/$1,00 tiers) ou Qwen3.8-27B | Excellent français, 262K contexte, pas cher |
| **API, qualité max** | Claude Opus 5.5 / Sonnet 5, GPT-6 Sol | Meilleur suivi d'instructions FR ; $4–5+/M |
| **API, économique** | DeepSeek V4-Pro ($0,435/M, cache hit $0,0036) | MoE 1,6T, 1M contexte ; prix imbattable avec cache |
| **Local 24 Go** | Qwen3.8-27B Q4_K_M (18 Go) | 262K contexte, français excellent, tient sur 4090 |
| **Local 16 Go** | Qwen3.5-14B / Mistral équivalent Q4 | Bon compromis |
| **Local, français natif** | Mistral Medium 3.5 / Magistral | Équipe française, bon FR ; licence Modified MIT à lire |

Le choix « sweet spot » 2026 pour un RAG FR en local : **un Qwen 27B en Q4_K_M sur 24 Go**. Assez gros pour bien suivre les instructions et citer, assez petit pour tenir sur une carte grand public avec du cache.

## 101. Architecture RAG complète proposée (avec les modèles du moment)

```
┌─────────────┐
│  INGESTION   │  Tes scripts existants + découpage (sections 109-111)
│  (une fois)  │  → chunks + métadonnées (section 112)
└──────┬──────┘
       │  Embedding (section 95) : text-embedding-3-small (actuel)
       │  ou Qwen3-Embedding-8B (self-host) → vecteurs en Qdrant/pgvector
       ▼
┌─────────────┐
│  REQUÊTE     │  Question FR → embedding → top-50 vectoriel
└──────┬──────┘
       │  Rerank (section 98) : jina-reranker-v3.5 ou Qwen3-Reranker-0.6B
       │  → top-5 à top-10, triés par pertinence
       ▼
┌─────────────┐
│ GÉNÉRATION   │  Prompt système FR + chunks cités + question
│              │  → Qwen3.8-27B local (Q4_K_M) ou API selon budget
└──────┬──────┘
       │  Réponse avec citations [chunk n°], « je ne sais pas » si vide
       ▼
     Utilisateur
```

Options à brancher plus tard (pas au jour 1) : HyDE (section 113) si le rappel plafonne, fine-tuning d'embeddings (section 92) si le vocabulaire métier coince, LoRA comportemental (section 89) si le format des réponses doit être standardisé.

## 102. Le pipeline chiffré : coût par question

Hypothèses : question FR moyenne, 15 chunks de 1200 tokens récupérés, rerankés en top-8, prompt final ~20K tokens, réponse ~800 tokens.

**En API (ordre de grandeur, grilles 2026) :**

| Poste | Tokens | Prix unit. | Coût/question |
|---|---|---|---|
| Embedding question | ~50 | $0,02/M | ~$0,000001 (négligeable) |
| Rerank (Cohere Fast) | 15 docs | $2,00/1000 SU | ~$0,002–0,01 |
| Génération (Qwen3.6 tiers) | 20K in / 0,8K out | $0,14/$1,00 | ~$0,0036 |
| Génération (flagship $4/$20) | 20K in / 0,8K out | $4/$20 | ~$0,096 |
| **Total éco** | | | **~$0,005–0,015/question** |
| **Total flagship** | | | **~$0,10/question** |

Lecture : 1000 questions/mois en mode éco = **$5–15/mois**. En flagship = ~$100/mois. Le rerank et l'embedding sont des poussières ; le générateur fait 90 %+ du coût. D'où l'intérêt du prompt caching (section 122) : avec 90 % de remise sur le préfixe en cache, le mode flagship tombe à ~$0,02/question.

**En local :** coût marginal ≈ électricité. Une 4090 consomme ~450W en charge ; à 10 s/question, 1000 questions = ~2,8 kWh ≈ **~0,50 €/mois**. Le coût est la carte (amortissement), pas l'usage.

## 103. Dimensionnement VRAM : la règle des octets

Tout le dimensionnement tient en une formule :

```
VRAM_nécessaire = poids_modèle + KV_cache + marge_moteur
```

Avec :

| Poste | Formule | Exemple 7B FP16, 32K |
|---|---|---|
| Poids | paramètres × octets/paramètre | 7B × 2 = 14 Go |
| KV cache | section 59 (par token × tokens) | 128 Kio × 32K = 4 Go |
| Marge moteur | 10–20 % (fragmentation, activations, overhead CUDA) | ~2–3 Go |
| **Total** | | **~20–21 Go** |

Règles de pouce :

- **Poids** : FP16/BF16 = 2 octets/param ; INT8 = 1 ; INT4 ≈ 0,5–0,6 (overhead des échelles) ; GGUF Q4_K_M ≈ 0,55–0,6.
- **Ne jamais dimensionner à 100 %.** Au-delà de ~90 % de VRAM, les allocateurs CUDA fragmentent et les perfs s'effondrent. Vise 70–80 % d'utilisation max.
- **Le batch multiplie le cache, pas les poids.** 4 requêtes simultanées = 1× poids + 4× cache. C'est le cache qui tue le multi-utilisateur, pas le modèle.

## 104. Tableau : poids des modèles selon format

| Modèle | FP16/BF16 | INT8/FP8 | INT4 (GPTQ/AWQ) | GGUF Q4_K_M | GGUF Q8_0 |
|---|---|---|---|---|---|
| 7B–8B | ~14–16 Go | ~7–8 Go | ~4–5 Go | ~4–5 Go | ~7–8 Go |
| 9B–14B | ~18–28 Go | ~9–14 Go | ~5–8 Go | ~6–10 Go | ~10–15 Go |
| 27B–32B | ~54–64 Go | ~27–32 Go | ~15–19 Go | ~16–20 Go | ~29–34 Go |
| 70B–72B | ~140 Go | ~70 Go | ~38–42 Go | ~40–44 Go | ~75 Go |
| 122B-A10B (MoE) | ~244 Go | ~128 Go (FP8 off.) | ~71 Go (Q4_K_M) | ~71 Go | ~121 Go |
| 1T (MoE, ex Kimi K2) | ~2000 Go | ~1000 Go | ~500 Go (natif) | — | — |

Chiffres réels croisés des recherches : Qwen3.5-122B-A10B FP8 officiel = 128,4 Go ; Q4_K_M ≈ 71 Go ; Q8_0 ≈ 121 Go. Qwen3.8-27B Ollama Q4_K_M = 18 Go. Kimi K2 INT4 natif ≈ 500 Go (d'où 8× H200).

## 105. Tableau : KV cache par contexte (rappel opérationnel)

Modèles GQA-8, head_dim 128, cache FP16. Ajoute aux poids du tableau précédent.

| Modèle | 4K | 8K | 32K | 128K | 1M |
|---|---|---|---|---|---|
| 7B (32 couches) | 0,5 Go | 1 Go | 4 Go | 16 Go | 128 Go |
| 13B (40 couches) | 0,6 Go | 1,25 Go | 5 Go | 20 Go | 160 Go |
| 27B (48 couches, type Qwen3.8) | 0,75 Go | 1,5 Go | 6 Go | 24 Go | 192 Go |
| 70B (80 couches) | 1,25 Go | 2,5 Go | 10 Go | 40 Go | 320 Go |

En FP8 : divise par 2. Lecture pour un RAG : avec 20K tokens de contexte par requête (15 chunks), un 27B consomme ~3,75 Go de cache par requête en FP16. Sur une 4090 (24 Go) avec le modèle en Q4 (18 Go), il reste ~4 Go utiles → **1 requête à la fois** confortablement, 2 en serré. Pour servir 4 utilisateurs simultanés, il faut soit une 48 Go, soit du FP8 partout, soit un modèle plus petit.

## 106. vLLM vs Ollama vs llama.cpp : que choisir

