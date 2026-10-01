---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-14
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meituan", "MiniMax", "Moonshot", "Xiaomi", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["agents", "apache", "asic", "attention", "benchmarks", "cyber", "deepseek", "exploit", "glm", "gqa", "kimi", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1254, 1350]
sha256: f56c34a2be9f500768d9be2e4cb44ba34665b17870ff7232dd2b8a6311e1f63a
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

| Modèle | Labo | Date | Archi | Total / actifs | Contexte | Prix $/1M (in/out) | Licence |
|---|---|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | Alibaba | 16/02/26 | MoE, GDN | 397B / 17B | ~262K | n.v. | Apache 2.0 |
| Ling-2.5-1T / Ring-2.5-1T | Ant | 16/02/26 | 1T MoE | 1T / n.v. | 1M | n.v. | Ouverts (n.v.) |
| Qwen3-Coder-Next | Alibaba | 02/26 | MoE, GDN | 80B / 3B | 256K | ~0,11 / 0,80 | Apache 2.0 |
| MiMo-V2-Pro | Xiaomi | 18/03/26 | MoE | >1T / 42B | n.v. | n.v. | Fermé |
| MiniMax-M2.7 | MiniMax | 18/03/26 | MoE, GQA full | 230B / 10B | ~200K | 0,30 / 1,20 | Ouverts (n.v.) |
| Qwen3.6-Plus | Alibaba | 02/04/26 | n.v. (fermé) | n.v. | 1M | ~0,29 / 1,74 | Fermé |
| LongCat-Next | Meituan | 02/04/26 | MoE multimodal DiNA | ~70B / 3B | n.v. | n.v. | MIT |
| Qwen3.6-35B-A3B | Alibaba | 16/04/26 | MoE, GDN | 35B / 3B | 256K | ~0,14 / 1,00 | Apache 2.0 |
| Qwen3.6-Max-Preview | Alibaba | 20/04/26 | MoE (presse) | 35B / 3B | 256K | n.v. | Fermé |
| Kimi K2.6 | Moonshot | 20/04/26 | MoE, MLA | 1T / 32B | 256K | 0,60 / 2,50 | Modified MIT |
| MiMo-V2.5 / V2.5-Pro | Xiaomi | 22/04/26 | MoE | 310B / 1,02T | n.v. | n.v. | MIT |
| DeepSeek V4-Pro | DeepSeek | 24/04/26 | MoE, attn. hybride | 1,6T / 49B | 1M | 0,435 / 0,87–0,93 | MIT |
| DeepSeek V4-Flash | DeepSeek | 24/04/26 | MoE, NVFP4 | 284B / 13B | 1M | n.v. | MIT |
| Qwen3.7-Max | Alibaba | 20/05/26 | n.v. (fermé) | n.v. | 1M | ~1,25 / 3,75 | Fermé |
| Qwen3.7-Plus | Alibaba | 20/05/26 | n.v. (fermé) | n.v. | 1M ? | 0,40 / 1,60 | Fermé |
| MiniMax-M3 | MiniMax | 01/06/26 | MoE | ~440B ? | 1M | n.v. | Ouverts (n.v.) |
| Kimi K2.7 Code | Moonshot | 12/06/26 | MoE, MLA | 1T / 32B | 256K | 0,95 / 4,00 | Modified MIT |
| GLM-5.2 | Zhipu | 16/06/26 | MoE, MLA+DSA | ~750B / 40B | 1M | 1,40 / 4,40 | MIT |
| LongCat-2.0 | Meituan | 30/06/26 | MoE, sparse attn | 1,6T / 48B | 1M natif | n.v. | MIT (poids ?) |
| Tencent Hy3 | Tencent | 06/07/26 | MoE + MTP | 295B / 21B | 256K | n.v. (post-gratuit) | Apache 2.0 |
| MiniMax-H3 | MiniMax | 17/07/26 | Dense 33B omni | 33B | multimodal | n.v. | Community (restrictive) |
| Ling-3.0-flash | Ant | 23/07/26 | MoE, KDA+MLA | 124B / 5,1B | 256K | 0,075 / 0,22 | MIT |
| Kimi K3 | Moonshot | 27/07/26 | MoE, KDA+MLA | 2,8T / 104B | 1M | 1,20 / 4,80 | Modified MIT (?) |
| Qwen3.8-Max | Alibaba | 03/08/26 | MoE | 2,4T / 95B | 1M | 2,00 / 6,00 | Ouverts (n.v.) |
| Qwen3.8-27B | Alibaba | 14/08/26 | Dense, GDN | 27,78B | 256K | n.v. | Apache 2.0 |
| GLM-5.3 | Zhipu | 14/08/26 | MoE (base 5.2) | ~744B / 40B | 1M | 1,40 / 4,40 | MIT (?) |
| GLM-5.3-Flash | Zhipu | 27/08/26 | MoE | 320B / 18B | 1M | 0,15 / 0,50 | MIT |
| Qwen3.8-Flash | Alibaba | 02/09/26 | MoE, GDN+sparse | 176B / 6B | 256K | 0,16 / 0,47 | Ouverts (n.v.) |
| humain-m3 | HUMAIN×MiniMax | 03/09/26 | MoE (lignée M3) | 428B | n.v. | n.v. | Preview |
| Ling-3.0-flash-VL | Ant | 09/09/26 | MoE + vision | ~125B / 5,1B | 128–256K | n.v. | MIT |
| Ling-3.0-flash-Fin | Ant | 09/09/26 | MoE (base 3.0) | 124B / 5,1B | n.v. | n.v. | Ouverts (n.v.) |
| DeepSeek V4.1-Flash | DeepSeek | 10/09/26 | MoE, enc.-dec. | 552B / 8–16B | 1M | 0,15 / 0,60 | MIT |
| MiMo-V2.6-Pro-RL | Xiaomi | 21/09/26 | MoE, omnimodal | 1,02T / 42B | 1M | 0,435 / 0,87 | MIT |
| MiMo-V2.6-Flash-RL | Xiaomi | 21/09/26 | MoE, omnimodal | 309B / 15B | 1M | 0,14 / 0,28 | MIT |

*(n.v. = non vérifié au 27/09/2026. Les prix « ~ » viennent de sources tierces.)*

## 98. Timeline février → septembre 2026

- **16 fév.** : Qwen3.5-397B-A17B (Alibaba, ouvert) + Ling-2.5-1T / Ring-2.5-1T (Ant, ouverts). Le mois
  commence par un tir groupé open-weight.
- **Mars** : Qwen3-Coder-Next · MiMo-V2-Pro/Omni/TTS (Xiaomi, fermés — parenthèse) · MiniMax-M2.7
  (self-evolution harness).
- **Avril** : Qwen3.6-Plus (fermé) · Qwen3.6-35B-A3B (ouvert) · Qwen3.6-Max-Preview (fermé — rupture
  « frontier-closed ») · Kimi K2.6 (ouvert, 1T) · MiMo-V2.5 (retour à l'ouvert) · DeepSeek V4-Pro/Flash
  (ouverts, MIT).
- **Mai** : Qwen3.7-Max/Plus (fermés, agents natifs 35 h, puce Zhenwu M890).
- **Juin** : MiniMax-M3 (ouvert, 1M) · Kimi K2.7 Code (ouvert) · GLM-5.2 (ouvert, MIT, AA #51) ·
  LongCat-2.0 (1,6T, ASIC chinois).
- **Juillet** : Tencent Hy3 (ouvert, Apache 2.0) · MiniMax-H3 (vidéo/audio, licence territoriale) ·
  Ling-3.0-flash (ouvert, MIT, $0,075/M) · Kimi K3 (ouvert, 2,8T — record).
- **Août** : Qwen3.8-Max (ouvert, 2,4T — revirement) · Qwen3.8-27B (dense ouvert qui bat Opus 4.6 Max sur
  le code) · GLM-5.3 (ouvert, « cyber ») · GLM-5.3-Flash / « Ox Alpha » (ouvert, 100K puces chinoises).
- **Septembre** : Qwen3.8-Flash (preview Qwen4, ouvert) · humain-m3 (dérivé arabe) · Ling-3.0-flash-VL/Fin ·
  DeepSeek V4.1-Flash (remplace V4-Flash et déprécie V4-Pro) · MiMo-V2.6 (RL retransmis en direct,
  controverse Anthropic).
- **En attente au 27/09** : DeepSeek V4.1 Pro (annoncé, rumeur 28–30 sept.), release open-weight
  d'humain-m3, Qwen4 (teasé via 3.8-Flash).

## 99. Que retenir (1/4) : la course au MoE géant open-weight

En huit mois, le plafond open-weight est passé de ~400B (Qwen3.5, février) à **2,8T paramètres** (Kimi K3,
juillet), avec 2,4T (Qwen3.8-Max), 1,6T (DeepSeek V4-Pro, LongCat-2.0) et 1T (Kimi K2.6, MiMo-V2.6-Pro,
Ling-2.5) entre les deux. Mais le chiffre qui compte n'est pas le total : c'est le **ratio total/actifs**,
passé de ~10× à **20–27×** (K3 : 2,8T/104B, Ling-3.0 : 124B/5,1B, Qwen3.8-Flash : 176B/6B). La sparsité est
devenue l'arme économique : servir 100B actifs coûte comme un dense 100B, mais le modèle « sait » comme
un 2,8T. Conséquence : l'écart open/fermé sur les benchmarks généralistes s'est refermé (K3 à AA 65,
MiMo-V2.6 à 46, GLM-5.3 à 63,8 — des niveaux frontier), et la question n'est plus « l'ouvert peut-il
rivaliser ? » mais « que reste-t-il au fermé ? ». Réponse provisoire : l'agentique très longue (35 h de
Qwen3.7-Max) et l'avance de quelques mois.

## 100. Que retenir (2/4) : le contexte 1M généralisé

Le million de tokens est passé du statut d'exploit (2024–2025) à celui de **standard de gamme** : V4,
V4.1-Flash, Qwen3.6-Plus, Qwen3.7-Max, Qwen3.8-Max, GLM-5.2/5.3/Flash, MiMo-V2.6, Kimi K3, Ling-2.5-1T,
MiniMax-M3, LongCat-2.0 — douze modèles de ce volume l'affichent. Trois moteurs : l'attention
linéaire/hybride (GDN, KDA, DSA+IndexShare) qui borne la mémoire, le KV cache compressé (MLA et dérivés),
et les optimisations système (Mooncake, HiCache, DeepEP). Le 1M « natif » vs « via YaRN » reste une
distinction à surveiller (le 256K stretché n'a pas la même qualité que le 1M entraîné), mais la direction
est sans ambiguïté : en 2026, un modèle chinois sans 1M (ou sans chemin crédible vers 1M) est un modèle
de seconde zone.

## 101. Que retenir (3/4) : les prix en chute libre

La guerre des prix 2026 se lit en trois chiffres. **Input standard** : de ~$1,40 (GLM-5.2, juin) à
**$0,075** (Ling-3.0-flash, juillet) — division par ~19 en six semaines sur le segment efficace.
**Flash** : $0,15–0,16/M input chez DeepSeek, Qwen et Zhipu (alignement au cent près, septembre).
**Cache hit** : $0,003625/M chez DeepSeek V4-Pro, soit 120× moins que le cache miss — la tarification
pousse à architecturer autour du cache. Pendant ce temps, les flagships fermés (Qwen3.7-Max à $1,25/M,
Qwen3.8-Max à $2,00/M, Kimi K3 à $1,20/M) restent 10–25× plus chers que les efficaces ouverts — mais ce
sont eux qui portent les scores frontier. Le marché s'est polarisé : le token banal ne vaut plus rien,
le token « frontier » se paie encore.

## 102. Que retenir (4/4) : l'attention linéaire/hybride, vraie révolution technique

