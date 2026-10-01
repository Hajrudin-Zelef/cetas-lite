---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-6
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Fireworks AI", "Huawei", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-04-20", "2026-06-12", "2026-06-15", "2026-07-27", "2026-09-27"]
keywords: ["agent", "agents", "ascend", "attention", "benchmarks", "claude", "cyber", "deepseek", "fine-tuning", "glm", "gpu", "int4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [501, 595]
sha256: 511b3bc777a0cca17b1d47addc503bc4232d2070686b1987c103ff99419af511
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : Kimi K2.6 (ID API : `kimi-k2.6`). **Sortie : 20 avril 2026**. Statut : disponible, poids ouverts.
- **Architecture** : MoE — **1T (1 billion) total / 32B actifs** par token ; **61 couches** (1 bloc dense +
  60 blocs MoE) ; **384 experts routés** (8 sélectionnés + 1 partagé ; config top-8-of-385) ; SwiGLU ;
  vocabulaire 163 840 tokens ; **livré nativement quantifié INT4** (pas de release haute précision).
- **Attention** : **MLA (Multi-head Latent Attention)** style DeepSeek-V3 + YaRN RoPE scaling.
- **Contexte** : **262 144 tokens** (256K).
- **Fonctionnalités** : multimodal natif (encodeur vision **MoonViT**, ~400M params : texte + image + vidéo) ;
  agents long-horizon (essaims jusqu'à **300 agents**, **4 000 étapes** coordonnées) ; thinking activé par
  défaut, température **fixée à 1,0** sur l'API ; raisonnement retourné dans le champ `reasoning_content`.
- **Poids** : **ouverts, Modified MIT** (HF `moonshotai/Kimi-K2.6`).
- **Inférence** : MLA → KV cache compressé (latent) ; INT4 natif ; guidance : 8× H200 (baseline), 8×
  B200/B300 ; vLLM + FlashInfer-MoE, CUTLASS maison (templates NVFP4), FlashAttention-2, parallélisme EP/TP/PP.
- **Benchmarks** : à sa sortie, « modèle open-weight le plus fort » (Artificial Analysis), à quelques
  points des systèmes fermés sur SWE-Bench Pro et HLE ; LMArena 1355 ; MMLU 88,0 ; HumanEval 87,0 ;
  SWE-bench 80,2 % (agrégateur tiers).
- **Prix API** : $0,60 / 1M input, $2,50 / 1M output (best-of-ai, juil. 2026) — prix officiel Moonshot non vérifié.
- **Déploiement** : vLLM, SGLang, KTransformers, API Moonshot (OpenAI-compatible).

## 40. Kimi K2.7 Code (« Kimi 2.7 »)

- **Nom exact** : **Kimi K2.7 Code** (ID API : `kimi-k2.7-code`) — « Kimi 2.7 » seul n'existe pas, c'est
  l'abréviation usuelle. **Sortie : 12 juin 2026** (mode HighSpeed : 15 juin ; dispo Cloudflare Workers AI
  le jour J). Statut : disponible, poids ouverts.
- **Architecture** : **identique à K2.6** (1T total / 32B actifs, 384 experts 8+1, 61 couches, MLA, SwiGLU,
  vocab ~160K) — construit directement sur le checkpoint K2.6, post-training spécialisé code ; **nativement INT4**.
- **Contexte** : 262 144 tokens ; **sortie max par défaut 32 768 tokens** (non relevable à « illimité » sur l'API).
- **Fonctionnalités** : code-first — +21,8 % sur Kimi Code Bench v2, +11,0 % Program Bench, +31,5 % MLS Bench
  Lite vs K2.6 ; **~30 % de tokens de raisonnement en moins** que K2.6 sur les tâches de code ; thinking
  **forcé** (toujours actif) ; multimodal (texte+image+vidéo via MoonViT, 4K image / FHD vidéo recommandés) ;
  sessions >12 h en continu (revendication Moonshot).
- **Poids** : **ouverts, Modified MIT** (HF `moonshotai/Kimi-K2.7-Code`).
- **Benchmarks** : MCP-Mark **81,1** (devant Claude Opus 4.8 : 76,4 ; derrière GPT-5.5 : 92,9). Réserve :
  Moonshot n'avait soumis **aucun bench indépendant** au 15/06/2026 (critique TechTimes) — chiffres
  constructeur.
- **Prix API** : **$0,95 / 1M input, $4,00 / 1M output, $0,19 / 1M cached input** (prix Moonshot) ; variante
  **HighSpeed** à 2× ($1,90/$8,00) pour ~180–260 tok/s.
- **Déploiement** : API Moonshot, Cloudflare Workers AI (`@cf/moonshotai/kimi-k2.7-code`), Fireworks,
  OpenRouter, vLLM/SGLang/KTransformers en self-hosted.

## 41. Kimi K3 (2,8T paramètres)

- **Nom exact** : Kimi K3 (HF `moonshotai/Kimi-K3`). Annonce ~**16 juillet 2026**, **poids open-source le
  27 juillet 2026**. Statut : disponible, poids ouverts — le plus gros modèle open-weight de l'histoire
  à sa sortie.
- **Architecture** : MoE très sparse — **2,8T (2 800B) total / 104B actifs** par token (16 experts activés
  sur 896). **Divergence non résolue** : une source (table 99sono) indique « ~50B actifs » avec le même
  « 16/896 » — incohérence non tranchée au 27/09/2026.
- **Attention** : hybride **Kimi Delta Attention (KDA, attention linéaire maison) + MLA** — support natif
  du contexte 1M.
- **Contexte** : **1 048 576 tokens (1M)**.
- **Fonctionnalités** : vision native ; « thinking mode » toujours actif ; coding, agents, raisonnement
  complexe — en tête des modèles ouverts sur plusieurs benchs.
- **Poids** : **ouverts, Modified MIT** (licence exacte à confirmer sur la carte HF — la carte
  `moonshotai/Kimi-K3` n'a pas été lue directement). **Téléchargement ~1,4–1,56 To** (96 shards
  safetensors) — le poids est la barrière, pas la licence.
- **Inférence** : KDA rend l'inférence long-contexte « plusieurs fois moins chère » ; **support Day-0** :
  Mooncake (réutilisation KV cache cross-instance, disaggregation PD/EPD), SGLang, vLLM, TokenSpeed.
  Guidance Moonshot : supernœud **64+ accélérateurs** ; minimum pratique cité : 4 nœuds / 32 GPU en
  400 Gbps RoCEv2/InfiniBand.
- **Benchmarks** : SWE-bench Verified **85,6 %**, Terminal-Bench 3.0 **79,8 %**, LMArena Elo **1512**
  (Tier-1), AA Intelligence **65,0**.
- **Prix API** : $1,20 / 1M input, $4,80 / 1M output (matrice comparative tierce) — prix officiel Moonshot
  non vérifié.
- **Fine-tuning** : non vérifié (poids ouverts → possible en théorie ; aucune recette vérifiée au 27/09).

## 42. Kimi : l'attention KDA + MLA

Le **Kimi Delta Attention (KDA)** est l'attention linéaire maison de Moonshot, combinée au MLA
(style DeepSeek) dans K3. Le principe est le même que le Gated DeltaNet de Qwen : remplacer une partie
des couches d'attention quadratique par des couches linéaires à état récurrent fixe, pour un coût mémoire
qui n'explose plus avec la longueur du contexte. K3 est le premier modèle à 1M de contexte natif servi
avec cette architecture hybride à l'échelle de 2,8T paramètres. Fait notable : le KDA de Moonshot a
essaimé — l'hybride « 35 couches KDA + 7 couches MLA » du Ling-3.0-flash d'Ant Group reprend
explicitement la brique Kimi (voir section 74). L'open-weight fait circuler les idées d'architecture
plus vite que les papiers.

## 43. Kimi : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| Kimi K2.6 | 20/04/2026 | MoE, MLA + YaRN | 1T / 32B | 256K | 0,60 / 2,50 (tiers) | Modified MIT |
| Kimi K2.7 Code | 12/06/2026 | MoE, MLA (base K2.6) | 1T / 32B | 256K | 0,95 / 4,00 (×2 HighSpeed) | Modified MIT |
| Kimi K3 | 27/07/2026 (poids) | MoE, KDA + MLA | 2,8T / 104B (divergence : ~50B) | 1M | 1,20 / 4,80 (tiers) | Modified MIT (à confirmer) |

## 44. Zhipu AI / Z.ai : portrait de l'entreprise

Zhipu AI (opérant à l'international sous la marque **Z.ai**), issue de l'université Tsinghua, est devenue
en 2026 la première entreprise IA chinoise cotée à innover au niveau frontier en open-weight. Sa famille
**GLM** a marqué février 2026 avec GLM-5 (744B/40B, entraîné entièrement sur puces **Huawei Ascend**,
zéro NVIDIA — fait géopolitique majeur corroboré par Reuters), puis l'été avec **GLM-5.2** (juin),
**GLM-5.3** (août, positionné « cyber ») et **GLM-5.3-Flash** (août). Zhipu livre aussi le harnais agent
officiel **ZCode** (style Claude Code) avec ses modèles.

## 45. GLM-5.2

