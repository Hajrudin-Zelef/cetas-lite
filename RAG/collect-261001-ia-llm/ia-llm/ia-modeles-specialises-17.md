---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-17
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "xAI"]
dates: []
keywords: ["agi", "attention", "bedrock", "benchmark", "benchmarks", "claude", "cohere", "deepseek", "embedding", "embeddings", "fine-tuning", "fp8"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1710, 1833]
sha256: c8ae46889b7892bb9a9b9b5dece5dc0875e851e54ce7db57edc34f6c72ec7593
---

# Encyclopédie des modèles IA — Volume 3

- **SWE-bench à 80 %+ en 2026** : le benchmark est en voie de saturation — quand les meilleurs dépassent 85 %, il ne discrimine plus. C'est le cycle de vie normal : un benchmark naît discriminant, meurt saturé, est remplacé (d'où HLE, ARC-AGI-2).
- **LMArena mesure la préférence, pas la justesse.** Un modèle « agréable » (formatage soigné, ton confiant) gagne des votes même s'il hallucine. Pour un RAG technique, préfère les benchmarks de factualité.

## 119. Vendor-reported vs indépendants : l'esprit critique

Trois niveaux de confiance, par ordre croissant :

1. **Chiffre fournisseur, benchmark maison ou obscur.** Valeur indicative uniquement. Exemple : « 9,2× débit supérieur vs autres modèles omni ouverts (comparateur non nommé) » (NVIDIA, Nano Omni) — non reproduit, comparateur anonyme.
2. **Chiffre fournisseur, benchmark standard (MTEB, SWE-bench).** Crédible mais optimisé : le fournisseur a choisi la config, le prompt, parfois le split. Exemple : les 80,6 % SWE-bench de Gemini 3.1 Pro et DeepSeek V4-Pro — même score, mais les harnesses diffèrent peut-être.
3. **Benchmark indépendant.** Artificial Analysis (vitesse/prix/qualité), Zilliz/Milvus (needle-in-a-haystack embeddings), les benchs communautaires (anton-abyzov pour Nano Omni). Moins flatteurs, plus robustes.

Drapeaux rouges (vus dans les recherches) :

- 🚩 Benchmark **contesé sur l'accès au split privé** (Voyage 4 / RTEB) — le score est potentiellement contaminé.
- 🚩 Comparateur **non nommé** (« vs autres modèles ouverts ») — on ne peut pas reproduire.
- 🚩 Score annoncé pour la **famille**, pas le checkpoint (Nemotron RTEB 78,46) — le modèle précis peut être en dessous.
- 🚩 Chiffres repris de **synthèses GitHub** sans fiche primaire (Opus 4.8, GPT-5.6 Sol, Grok 4.5) — deux degrés de séparation de la source.
- 🚩 Prix de **sources secondaires** (OpenRouter, fal.ai, MuAPI) présentés comme grille officielle — le prix réel dépend du chemin.

## 120. La checklist esprit critique (à appliquer à chaque fiche)

Avant de citer un chiffre ou de choisir un modèle sur un benchmark :

- [ ] Le benchmark est-il **standard et public** (MTEB, SWE-bench) ou maison/obscur ?
- [ ] Le score concerne-t-il **ce checkpoint précis** ou la famille ?
- [ ] La source est-elle **primaire** (fiche officielle, papier) ou secondaire (presse, synthèse GitHub) ?
- [ ] Existe-t-il une mesure **indépendante** qui confirme l'ordre de grandeur ?
- [ ] Le benchmark mesure-t-il **mon usage** (retrieval FR, pas classification EN) ?
- [ ] Les prix viennent-ils de la **grille officielle** ou d'un agrégateur ?
- [ ] Les dates sont-elles cohérentes (annonce vs sortie vs GA) ?
- [ ] Les contradictions entre sources sont-elles **signalées** ou gommées ?

Si 3+ cases sont douteuses : le chiffre est une indication, pas un fait. C'est exactement la discipline appliquée dans la Partie A de ce volume.

## 121. Prix API 2026 : les ordres de grandeur

Fourchettes constatées dans les recherches (grilles officielles quand précisé, secondaires sinon — voir les réserves par modèle) :

**Embeddings (par million de tokens) :**

| Modèle | Prix | Source |
|---|---|---|
| Voyage 4 lite | $0,02 | Docs officielles |
| Voyage 4 / 4-large | $0,06 / $0,12 | Docs officielles |
| Cohere Embed v4 | ~$0,12 (Bedrock) | Secondaire |
| Gemini Embedding 2 (texte) | ~$0,20 | Secondaire (OpenRouter) |

**Rerankers :**

| Modèle | Prix | Source |
|---|---|---|
| Cohere Rerank 4.0 Fast | $2,00 / 1000 search units | Microsoft/Azure |
| Cohere Rerank 4.0 Pro | $2,50 / 1000 search units | Microsoft/Azure |
| Voyage rerank-2.5-lite | ~$0,02/M (?) | Secondaire, unité non vérifiée |

**LLM générateurs (entrée/sortie par million) :**

| Modèle | Prix | Source |
|---|---|---|
| DeepSeek V4-Pro | $0,435 / $0,87–0,93 | Recherche DeepSeek |
| Qwen3.6-35B-A3B (tiers) | ~$0,14 / $1,00 | Secondaire (AirCloud, FP8) |
| Gemini 3 Flash | $0,30 / $1,50 | Guides tiers |
| Gemini 3.1 Pro | ~$2 / $12 | Secondaire |
| Kimi K3 | ~$3 / $15 | Secondaire |
| GPT-5.6 Sol (promo) | $4 / $20 | Secondaire |
| Claude Opus 4.8 | ~$5 / $25 | Secondaire |

Lecture : **un facteur ~30** entre l'entrée de gamme (DeepSeek, Qwen tiers) et les flagships. Pour un RAG où le générateur fait 90 % du coût (section 102), le choix du générateur EST le choix du budget.

## 122. Le prompt caching : la remise invisible

Le **prompt caching** : le fournisseur met en cache le préfixe de tes prompts (prompt système + chunks fréquents) et ne te facture le cache hit qu'une fraction du prix. En 2026, c'est standard chez tous les gros fournisseurs :

| Fournisseur | Remise cache hit (ordre de grandeur) | Seuil |
|---|---|---|
| OpenAI (GPT-5.6/6) | **−90 %** sur les lectures en cache | Cache-write pricing introduit (payer l'écriture) |
| DeepSeek V4-Pro | **−99 %** ($0,435 → $0,003625/M) | Le plus agressif du marché |
| Anthropic (Claude) | −90 %, seuils 512–2048 tokens selon modèle | Documenté par seuils |
| Google (Gemini) | −75 à −90 % selon modèle | — |

Pour un RAG, c'est massif : ton prompt système + tes instructions sont **identiques à chaque requête** → 100 % de hit sur le préfixe. Si les chunks se répètent (questions similaires sur les mêmes guides), le hit rate monte encore.

Stratégie : **mets le stable en premier** (prompt système, instructions, chunks les plus fréquents), le variable en dernier (la question). Le cache fonctionne par préfixe : tout ce qui change invalide la suite. Avec un bon ordonnancement, un RAG flagship à $4/M tombe à ~$0,50/M effectif.

## 123. Le Batch API : −50 % pour les pressés-pas-vite

Le **Batch API** : tu soumets tes requêtes en lot, le fournisseur les traite en différé (typiquement < 24 h), tu paies **−50 %**. Proposé par OpenAI, Google (Gemini Embedding 2 : batch −50 % cité), Anthropic.

Pour un RAG, les cas d'usage :

- **Évaluation** : faire tourner tes 50–500 questions de test en batch = moitié prix. L'évaluation n'a pas besoin de temps réel.
- **HyDE / génération de questions synthétiques** : la génération des paires pour le fine-tuning d'embeddings (section 92) se fait en batch.
- **Re-vectorisation** : si tu changes d'embedding, re-vectoriser 2M tokens en batch divise la facture par 2 (elle est déjà petite).
- ❌ **Pas pour le temps réel** : les questions utilisateurs restent en synchrone plein tarif.

Combiné au prompt caching, l'ordre de grandeur d'un pipeline optimisé : **÷5 à ÷10** par rapport au prix facial synchrone sans cache. Ne jamais chiffrer un projet sur les prix faciaux.

## 124. Calculer le coût d'un RAG : la méthode

Modèle de coût mensuel :

```
Coût_mensuel = N_questions × (coût_embedding_q + coût_rerank + coût_génération)
             + coût_indexation_initiale (amorti)
             + coût_évaluation (batch, amorti)
```

Exemple chiffré : 2000 questions/mois, pipeline éco (section 102) :

| Poste | Calcul | Coût/mois |
|---|---|---|
| Embeddings questions | 2000 × $0,000001 | ~$0,002 |
| Rerank (Cohere Fast) | 2000 × ~$0,005 | ~$10 |
| Génération (Qwen tiers, 20K in) | 2000 × $0,0036 | ~$7,20 |
| Indexation initiale (2M tokens, une fois) | $0,04–0,40 | négligeable amorti |
| Évaluation trimestrielle (500 Q en batch) | ~$1/trimestre | ~$0,33/mois |
| **Total** | | **~$18/mois** |

Leviers dans l'ordre d'impact : 1) passer le générateur en local (→ ~0 €), 2) prompt caching (→ ÷5 sur la génération API), 3) reranker local 0,6B (→ ~0 €), 4) batch pour l'évaluation. Un RAG personnel bien optimisé coûte **moins de 20 €/mois en API** ou **~0 € en local** (amortissement carte).

## 125. Glossaire (50 termes)

**Attention** — mécanisme qui pondère les tokens précédents selon leur pertinence pour le token courant ; coût quadratique en la longueur de séquence.

