---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-19
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "LongCat", "Meituan", "MiniMax", "Moonshot", "OpenRouter", "United States", "Xiaomi", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["agent", "apache", "attention", "attribution", "benchmarks", "compute", "deepseek", "distribution", "glm", "kimi", "multimodal", "omni"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1845, 1956]
sha256: bffb8dde851ca31bf3b949b171ab22a5f3bb3fd293e9125020722a4b4c3f34aa
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

## 165. MiniMax : MaxProof, l'or aux IMO/USAMO

Le framework **MaxProof** (generative-verifier RL + evolutionary search) a fait passer MiniMax-M3
**au-dessus du seuil médaille d'or humaine à l'IMO 2025 et à l'USAMO 2026** (blog MiniMax, 9 juin 2026).
C'est un résultat de post-training, pas de pré-entraînement : la méthode (vérificateur génératif + recherche
évolutionnaire) compte autant que le modèle. À mettre en regard du Ring-2.5-1T d'Ant (IMO 2025 : 35/42,
niveau or) : en 2026, **deux labos chinois ouverts revendiquent le niveau médaille d'or en maths de
compétition**. Le raisonnement mathématique frontier n'est plus un monopole fermé — et c'est un
pré-entraînement + RL ciblé, pas un scaling brut, qui l'a apporté.

## 166. Tencent : la fenêtre gratuite du 6 au 21 juillet 2026

Pour lancer Hy3, Tencent a ouvert une **fenêtre gratuite** : Nous Portal + OpenRouter `:free` du ~6 au
~21 juillet 2026 (fenêtre fermée depuis). C'est une tactique de distribution agressive : inonder les
comparateurs (OpenRouter, LMArena) et les harnais (Hermes Agent de NousResearch a intégré Hy3 pendant la
fenêtre — voir le guide dev.to « now that the free window closed ») pour acheter de la notoriété au prix
du compute. Le prix courant post-fenêtre est **non vérifié** au 27/09/2026 — Tencent n'a pas communiqué
de grille publique dans les sources consultées. Leçon : une fenêtre gratuite crée un pic d'adoption qui
ne dit rien du prix d'équilibre ; toujours vérifier le tarif en vigueur avant d'architecturer dessus.

## 167. Face-à-face : les six attentions chinoises

| Labo | Nom | Principe | Modèles |
|---|---|---|---|
| DeepSeek | Hybride CSA+HCA+SWA | 3 sparsités d'attention combinées | V4, V4.1-Flash |
| Alibaba | Gated DeltaNet (+ Gated Attention) | Attention linéaire récurrente + blocs standard | Qwen3.5→3.8 |
| Moonshot | KDA + MLA | Attention linéaire maison + cache latent | Kimi K3 (Ant reprend le KDA) |
| Zhipu | MLA + DSA + IndexShare | Sparse dynamique + partage d'indices top-k | GLM-5.2/5.3 |
| Ant | KDA + MLA (hybride 35+7) | Reprise du KDA Moonshot | Ling-3.0 |
| Meituan | LongCat Sparse Attention / ScMoE | Sparse attention + overlap calcul/communication | LongCat-2.0 / Flash |

Constat : **zéro labo chinois majeur n'utilise l'attention dense vanilla** sur ses flagships 2026. Tous
ont une brique linéaire ou sparse. Et les briques circulent (KDA Moonshot → Ant). L'attention n'est plus
un choix d'architecture, c'est une **commodité optimisée** — la différenciation s'est déplacée vers le
post-training (RL, environnements, harnais).

## 168. Face-à-face : le coût du million de tokens (input)

| Segment | Modèle | $/1M input | Contexte |
|---|---|---|---|
| Efficace ouvert | Ling-3.0-flash | **0,075** | 256K |
| Flash | DeepSeek V4.1-Flash | 0,15 (0,30 peak) | 1M |
| Flash | GLM-5.3-Flash | 0,15 | 1M |
| Flash | Qwen3.8-Flash | 0,16 | 256K |
| Mid ouvert | MiniMax-M2 | 0,30 | ~200K |
| Mid fermé | Qwen3.7-Plus | 0,40 | 1M ? |
| Flagship ouvert | MiMo-V2.6-Pro-RL | 0,435 | 1M |
| Flagship ouvert | DeepSeek V4-Pro | 0,435 | 1M |
| Flagship ouvert | Kimi K3 | 1,20 | 1M |
| Flagship fermé | Qwen3.7-Max | ~1,25 | 1M |
| Flagship ouvert | GLM-5.2/5.3 | 1,40 | 1M |
| Flagship ouvert | Qwen3.8-Max | 2,00 | 1M |

Lecture : **un facteur 27×** entre le moins cher (Ling-3.0-flash) et le plus cher (Qwen3.8-Max) de ce
volume — pour des modèles qui se côtoient dans les mêmes benchmarks. Le prix ne mesure plus la qualité,
il mesure le positionnement.

## 169. Face-à-face : les licences

| Licence | Modèles | Contrainte |
|---|---|---|
| MIT | DeepSeek (tous), Zhipu GLM, Xiaomi MiMo, Meituan LongCat, Ant Ling-3.0 | Aucune (commercial OK) |
| Apache 2.0 | Qwen (3.5, 3.6, 3.8-27B, Coder), Tencent Hy3 | Aucune (commercial OK, clause brevets) |
| Modified MIT | Kimi K2.6, K2.7 Code, K3 (à confirmer carte HF) | Attribution UI au-delà de 100M MAU / $20M mensuels |
| Community territoriale | MiniMax-H3 | Local interdit US/UE/UK/KR sans autorisation |
| Non vérifiée | MiniMax-M2/M3, Qwen3.8-Max/Flash, GLM-5.3 | **Vérifier la carte HF avant usage commercial** |
| Fermé | Qwen3.6-Max, 3.6-Plus, 3.7-Max/Plus, MiMo-V2-Pro/Omni/TTS | API uniquement |

Règle d'or : **la licence se lit sur la carte Hugging Face du checkpoint exact**, pas dans un article.
Cinq modèles majeurs du volume ont une licence non vérifiée au 27/09/2026 — c'est le premier chantier
avant toute mise en production.

## 170. Face-à-face : les tiers « Flash » de septembre 2026

| Modèle | Date | Total / actifs | Contexte | $/1M in | $/1M out | Licence |
|---|---|---|---|---|---|---|
| DeepSeek V4.1-Flash | 10/09/26 | 552B / 8–16B | 1M | 0,15 (×2 peak) | 0,60 (×2 peak) | MIT |
| GLM-5.3-Flash | 27/08/26 | 320B / 18B | 1M multimodal | 0,15 | 0,50 | MIT |
| Qwen3.8-Flash | 02/09/26 | 176B / 6B | 256K | 0,16 | 0,47 | Ouverts (n.v.) |
| MiMo-V2.6-Flash-RL | 21/09/26 | 309B / 15B | 1M omnimodal | 0,14 | 0,28 | MIT |

Quatre labos, quatre Flash, **un prix d'entrée convergé à $0,14–0,16/M** : c'est un oligopole tarifaire
de fait. Les différenciateurs restants : le multimodal natif (GLM, MiMo), la tarification peak (DeepSeek),
la preview d'architecture future (Qwen). Le V4-Flash original (retiré le 10/09) a été cannibalisé par son
successeur le jour même — la durée de vie d'un tier Flash se compte en mois. Pour l'acheteur : le Flash
est devenu une commodité, on choisit sur les features (vision ? peak ? tool use ?), plus sur le prix.

## 171. Les modèles retirés : V4-Flash et la fragilité des IDs API

Un seul modèle du volume a été **retiré** en 2026 : **DeepSeek V4-Flash** (24 avril → 10 septembre 2026,
4,5 mois de vie). Mais trois autres sont en **fin de vie programmée** : V4-Pro (dépréciation ordonnée vers
V4.1-Flash, statut exact flou), Hy3 Preview (remplacé par Hy3), Qwen3.8-Max snapshot du 3 août (remplacé
par le 0902). La leçon dépasse la Chine : en 2026, **un ID d'API est un pointeur mutable**, pas un
contrat. Bonnes pratiques : épingler modèle + snapshot + date dans la config, monitorer les changelogs
fournisseurs (api-docs.deepseek.com/news/, blogs labos), prévoir un modèle de repli testé (ex. bascule
V4-Pro → V4.1-Flash), et rejouer une batterie d'évaluations à chaque changement d'ID — le comportement
peut changer sans préavis.

## 172. Pour un RAG personnel : que choisir dans ce volume

En restant strictement sur les faits vérifiés : pour un **RAG francophone exigeant** (long contexte,
prix bas, licence propre), les candidats sérieux sont **Ling-3.0-flash** ($0,075/M input, MIT, 256K→1M,
5,1B actifs — le moins cher du volume), **Qwen3.8-27B** (dense ouvert Apache 2.0, 18 Go en Q4_K_M,
tourne sur 1× 24 Go, support Ollama day-0 — le meilleur choix local), **DeepSeek V4.1-Flash** ($0,15/M,
MIT, 1M, entrée image native) et **MiMo-V2.6-Distill-Qwen-9B** (16 Go VRAM, lignée trillion distillée —
le plus accessible en local). Pour l'**agentique** : Kimi K2.6/K3 (essaims, Modified MIT — clause
d'attribution à connaître), GLM-5.2 (ZCode, MIT), Tencent Hy3 (Apache 2.0, preserved reasoning). À éviter
en production sans vérification : tout ce qui est marqué licence non vérifiée (MiniMax-M2/M3, GLM-5.3,
Qwen3.8-Max/Flash) et les modèles retirés ou en dépréciation (V4-Flash, V4-Pro).

## 173. Le local : ce qui tourne vraiment (ordres de grandeur)

