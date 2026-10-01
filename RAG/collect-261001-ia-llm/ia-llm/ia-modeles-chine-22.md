---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-22
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "LongCat", "MiniMax", "Moonshot", "OpenRouter", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["attention", "deepseek", "fp4", "fp8", "glm", "kimi", "kv cache", "moe", "multimodal", "nvfp4", "omni", "quantization"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [2167, 2289]
sha256: 5e8b1d64c6e6fb81211ba227e764e1c612d7908ac9dfff3dc60a02ff9f0fcc07
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

| Modèle | Bench | Score | Source |
|---|---|---|---|
| Kimi K3 | SWE-bench Verified | **85,6 %** | vendor |
| DeepSeek V4-Pro Max | SWE-bench Verified | 80,6 % | vendor |
| Qwen3.7-Max | SWE-bench Verified | 80,4 % | vendor |
| Qwen3.6-Plus | SWE-bench Verified | 78,8 % | vendor |
| Tencent Hy3 | SWE-bench Verified | 78,0 % | vendor |
| Qwen3.6-27B | SWE-bench | 77,2 % | agrégateur tiers |
| Qwen3.8-27B | SWE-bench Pro | **61,7 %** | agrégateur tiers |
| Qwen3.8-Flash | SWE-bench Pro | 62,5 % | vendor/agrég. |
| GLM-5.2 | SWE-bench Pro | 62,1 % | vendor |
| MiMo-V2.6-Pro-RL | DeepSWE v1.1 | 71,9 % | vendor |
| DeepSeek V4.1-Flash | DeepSWE v1.1 | **74,2 %** | vendor |
| GLM-5.3 | DeepSWE v1.1 | 66,9 % | vendor |
| GLM-5.3 | SWE-bench Verified | 57,8 % vs 82,4 % | **divergence non résolue** |

Lecture : le trio de tête (K3, V4-Pro, Qwen3.7-Max) se tient en 5 points sur le Verified ; sur le Pro
(version dure), ce sont les petits ouverts (Qwen3.8-27B, Qwen3.8-Flash, GLM-5.2) qui mènent la danse des
ouverts. Et la divergence GLM-5.3 rappelle que **tout ce tableau est vendor-reported** : à prendre comme
une carte des revendications, pas comme un classement.

## 194. Face-à-face : qui fait quoi en multimodal

| Modèle | Texte | Image (in) | Image (out) | Vidéo (in) | Vidéo (out) | Audio (in) | Audio (out) |
|---|---|---|---|---|---|---|---|
| MiMo-V2.6-Pro-RL | ✓ | ✓ | n.v. | ✓ | n.v. | ✓ | n.v. |
| LongCat-Next | ✓ | ✓ | ✓ | ✓ | compréhension seule | ✓ | ✓ (24 kHz) |
| GLM-5.3-Flash | ✓ | ✓ | — | ✓ | — | — | — |
| Qwen3.8-Max | ✓ | ✓ | n.v. | ✓ | n.v. | n.v. | — |
| Tencent Hy3 | ✓ | — | — | — | — | — | — |
| MiniMax-H3 | (prompt) | ✓ (cond.) | — | ✓ (cond.) | ✓ 4–15 s | ✓ (cond.) | ✓ stéréo 32 kHz |
| Ming-Flash-Omni-2.0 | ✓ | n.v. | n.v. | n.v. | n.v. | ✓ (parole+musique) | n.v. |

(— = non ; n.v. = non vérifié ; cond. = en conditioning d'entrée.) Constat : **aucun modèle du volume ne
fait tout** — l'« omni » reste un horizon. Le texte+image+vidéo en entrée est devenu standard (Qwen, GLM,
MiMo, Kimi) ; la génération multimodale reste l'exception (LongCat-Next pour l'image+voix, MiniMax-H3
pour la vidéo+audio).

## 195. Face-à-face : les réglages « thinking »

| Modèle | Thinking par défaut | Réglage | Désactivable |
|---|---|---|---|
| DeepSeek V4.1-Flash | on | low / high / max | oui |
| Qwen3.8-Max | xhigh (défaut) | low / medium / xhigh | oui (/think, /no_think) |
| GLM-5.2 | réglable | non-thinking → High → Max | oui |
| GLM-5.3 | max (défaut) | low / high / max | **non** (`disabled` rejeté) |
| Kimi K2.7 Code | forcé | — | non |
| Kimi K3 | toujours actif | — | n.v. |
| Tencent Hy3 | hybride | no_think / think_low / think_high | oui (no_think) |
| Ling-3.0-flash | on (CoT) | — | oui |
| Qwen3.8-27B | on | — | oui |

Tendance : le thinking **réglable et facturable** (plus d'effort = meilleurs scores mais plus de tokens
= plus cher, cf. Qwen3.8-Max-0902). Le GLM-5.3 qui **interdit** de couper le thinking est le signal
faible : les labos ne font plus confiance à l'utilisateur pour doser le raisonnement — ou ne veulent pas
qu'on mesure le modèle sans.

## 196. Les noms qui n'ont rien donné (non trouvés)

Honnêteté méthodologique : ces noms ont été cherchés (2–3 requêtes chacun) et **n'existent pas** au
27/09/2026. **« Qwen3.8-35B-A3B »** : aucune trace — la série 3.8 ouverte connue = Max + 27B + Flash.
**« Kimi 2.7 »** (standalone) : n'existe pas, c'est Kimi K2.7 Code. **« Qwen3.5/3.6/3.7-Coder »** :
n'existent pas — le codeur 2026 est Qwen3-Coder-Next (février) puis 480B-A35B. **« DeepSeek V4 Pro »**
sans suffixe : c'est la variante V4-Pro (GA 0813), pas un autre modèle. **« DeepSeek V5 », « MiMo-V3 »,
« Ling-4.0 », « LongCat-3.0 »** : hors périmètre, rien n'est apparu dans les résultats — à noter pour la
veille, pas pour ce volume. Si vous croisez ces noms dans la presse, c'est soit une erreur, soit une
rumeur, soit l'avenir.

## 197. Lexique des IDs API (correspondances vérifiées)

| ID API | Modèle réel | Fournisseur |
|---|---|---|
| `deepseek-flash` | DeepSeek V4.1-Flash | DeepSeek |
| `deepseek/deepseek-v4.1-flash` | DeepSeek V4.1-Flash | OpenRouter |
| `deepseek-v4-pro` | DeepSeek V4-Pro (dépréciation en cours) | DeepSeek |
| `kimi-k2.6` | Kimi K2.6 | Moonshot |
| `kimi-k2.7-code` | Kimi K2.7 Code | Moonshot |
| `glm-5.2` | GLM-5.2 | Z.ai |
| `qwen3.6-flash` | Qwen3.6-35B-A3B (ouvert) | Alibaba Model Studio |
| `qwen3.6-plus` | Qwen3.6-Plus (fermé) | Alibaba Model Studio |
| `qwen3.7-max` / `qwen3.7-plus` | Qwen3.7-Max / Plus (fermés) | Alibaba Model Studio |
| `qwen3.8-max` / `qwen3.8-flash` | Qwen3.8-Max / Flash (ouverts) | Alibaba / QwenCloud |
| `hy3` | Tencent Hy3 | Tencent |
| `@cf/moonshotai/kimi-k2.7-code` | Kimi K2.7 Code | Cloudflare Workers AI |
| `stealth/ox-alpha` | GLM-5.3-Flash (était stealth, 20–26/08) | OpenRouter |
| `qwen-plus` / `qwen-max` | IDs génériques historiques, mapping exact non vérifié | docs tierces |

À épingler dans la config : un ID générique (`qwen-plus`) peut changer de modèle sans préavis ; un ID
versionné (`qwen3.8-flash`) est plus stable mais reste un pointeur (cf. snapshots).

## 198. Check-list de mise en production (modèles de ce volume)

Avant de mettre en production un modèle chinois de ce volume, vérifier dans l'ordre : **(1) Licence** —
lire la carte Hugging Face du checkpoint exact (5 modèles majeurs ont une licence non vérifiée au
27/09/2026 ; la clause Modified MIT de Kimi et la restriction territoriale du H3 sont des cas
pièges). **(2) Snapshot** — épingler modèle + snapshot + date (Qwen3.8-Max-0902 ≠ snapshot du 3 août).
**(3) Contexte réel** — distinguer natif / YaRN / stretch (Hy3 : 128K validé en pratique, 256K = stretch).
**(4) Tarification** — peak/off-peak (DeepSeek), cache hit/miss (120× d'écart), plan au token vs Coding
Plan vs abonnement. **(5) Fallback** — un ID peut changer de modèle en 96 h (V4-Pro, août–sept. 2026) ;
prévoir un repli testé. **(6) Réglages** — partir des sampling recommandés du labo (MiniMax, Kimi),
pas des défauts du harnais. **(7) Évaluations** — rejouer sa batterie à chaque changement d'ID ou de
snapshot ; les chiffres vendor-reported ne dispensent pas de mesurer soi-même. **(8) Juridiction** —
données et droit applicable (QwenWork/« state-law risk », H3 territorial).

---

*Fin du volume 1 (provisoire — recompté à la fin).* Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*

## 199. DeepSeek : l'optimiseur Muon (détail d'entraînement)

Parmi les rares détails d'entraînement de V4 qui ont filtré : l'usage de l'**optimiseur Muon** (complété
par de l'AdamW partiel) et du **FP4 QAT** (quantization-aware training) pour les poids MoE. Muon
(MomentUm Orthogonalized by Newtonian updates) orthogonalise les mises à jour du momentum — en pratique,
il converge plus vite qu'AdamW sur les gros transformers, au prix d'une implémentation plus délicate
(distribuée chez DeepSeek). Le FP4 QAT, lui, consiste à entraîner en simulant déjà la quantification 4
bits : les poids MoE naissent « prêts pour le NVFP4 », au lieu d'être quantifiés après coup avec perte.
C'est cohérent avec toute la stratégie V4 : **penser l'inférence dès l'entraînement** (poids NVFP4
natifs, attention hybride pour le KV cache). Réserve : ces détails viennent de notes tierces, pas d'un
rapport officiel — mais ils sont techniquement cohérents avec les formats de poids publiés (BF16/FP8/FP4
natifs sur le Hub).

## 200. Qwen : les ratios d'hybridation (3:1, 75/25, 48+16)

