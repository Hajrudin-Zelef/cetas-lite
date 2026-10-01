---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-10
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Ant", "LongCat", "Meituan", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "vLLM"]
dates: ["2026-02-11", "2026-02-16", "2026-07-23", "2026-08-05", "2026-09-09"]
keywords: ["agent", "agents", "asic", "attention", "benchmarks", "blackwell", "fine-tuning", "fp4", "fp8", "gptq", "gqa", "int4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [891, 994]
sha256: e2eed084b2057e88e5f4c1fc9b582820f3d7c2474f9f0be20866b1b9ba0df457
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : Ling-3.0-flash (HF `inclusionAI/Ling-3.0-flash`). **Sortie : 23 juillet 2026**. Statut :
  disponible, open weights.
- **Architecture** : MoE, **124B total / 5,1B actifs par token**. Attention hybride : **35 couches KDA
  (Kimi-Delta linear attention) + 7 couches MLA** (gated MLA), **512 experts routés + 1 partagé, top-8**,
  tête MTP (multi-token prediction), 42 couches. Famille `BailingMoeV3`.
- **Contexte** : **256K natif (262 144), extensible à 1M**.
- **Fonctionnalités** : thinking model (CoT activé par défaut, désactivable), tool calling structuré,
  agents de production. **HiCache** (hierarchical cache) : -60 à -80 % TTFT sur longues conversations
  (vendor claim).
- **Poids** : **ouverts, MIT**. Formats : BF16 (~250 Go), FP8 (128,4 Go), INT4 W4A16 (77 Go), MXFP4/FP4
  (70,4 Go, Blackwell only) ; draft DSpark pour speculative decoding.
- **Inférence** : attention linéaire hybride (KDA) → KV cache réduit ; support vLLM (PR #51045, mergée
  2026-08-05) et SGLang (parsers `ling3` reasoning/tool-call). Fine-tuning local possible (poids ouverts)
  — pas d'offre FT API documentée (non vérifié).
- **Benchmarks** : AA Intelligence Index **38** (source tierce best-of-ai).
- **Prix API** : **$0,075 / 1M input, $0,22 / 1M output** (first-party) — le moins cher de tout le volume
  côté input. OpenRouter : variante `:free` à $0.
- **Déploiement** : vLLM, SGLang, Ollama (non vérifié pour Ollama), DGX Spark (GB10, INT4 72 Go),
  H100/H200/B200/GB300 (tp 4–8).

## 71. Ling-3.0-flash-VL (septembre 2026, la vision)

- **Nom exact** : Ling-3.0-flash-VL. **Sortie : 9 septembre 2026** (open-source ; API déjà lancée avant).
  Statut : disponible, open weights.
- **Architecture** : backbone Ling-3.0-flash + encodeur vision 27 couches (0,4B, patches 16px, merging 2×2,
  frames vidéo appariées) via MLP projector. **~125B total / ~5,1B actifs**. Famille BailingMoeV3.
- **Contexte** : 128K natif (doc SGLang Ant) **vs** 256–262K (cryptobriefing) — **incohérence non tranchée** ;
  une source mentionne extension max 1M (non vérifié).
- **Fonctionnalités** : vision native (images, vidéo, documents, UI logicielles), tool use avec **boucle de
  feedback visuel** (auto-correction : l'agent voit le résultat de son action et se corrige), thinking
  activé par défaut. Pas d'audio.
- **Poids** : **ouverts, MIT** — BF16, FP8, INT4 (GPTQ), FP4 (MXFP4, Blackwell).
- **Benchmarks** : AA Intelligence Index **25** (cryptobriefing).
- **Prix API** : non vérifié. **Déploiement** : API via Ling Studio ; SGLang validé (4×GB300 / 4×H200,
  tp 4 ; tp 8 sur H100).

## 72. Ling-3.0-tiny (le nain)

- **Nom exact** : Ling-3.0-tiny. Date : non vérifiée (postérieur au 23 juillet 2026 ; guide de déploiement
  daté septembre 2026). Statut : disponible, open weights (licence non vérifiée).
- **Architecture** : Sparse MoE, **7,9B total / 1,3B actifs**, attention hybride **KDA linéaire + Gated MLA**,
  **128 experts routés**, contexte **128K natif**.
- **Déploiement** : SGLang sur DGX Spark (BF16). **Prix / benchmarks** : non vérifiés.
- Avec 1,3B actifs par token, c'est le modèle « sérieux » le plus frugal du volume — la preuve que la
  recette BailingMoeV3 descend à l'échelle du poste de travail.

## 73. Ling-3.0-flash-Fin (septembre 2026, la finance)

- **Nom exact** : Ling-3.0-flash-Fin. **Sortie : 9 septembre 2026** (annonce Business Wire ; une source
  tierce dit 3 septembre — incohérence mineure non tranchée). Annoncé à l'Inclusion·Conference on the Bund
  (Shanghai). Statut : open-sourcé (licence non vérifiée).
- **Architecture** : MoE, **124B total / 5,1B actifs** (base Ling-3.0-flash, fine-tuné finance).
- **Fonctionnalités** : recherche investissement, conformité, modélisation de valorisation, reporting
  réglementaire ; co-développé avec des institutions financières (CICC investment banking).
- **Benchmarks** (Ant) : FinFIRST, FinSearchComp Verified, FinCRAFT, FinanceAgent v1.1/v2, APEX-Agents,
  SpreadsheetBench v1/v2, τ³-Banking.
- C'est le seul modèle vertical « métier » (finance) open-weight du volume — tous les autres sont des
  généralistes. Signe que la différenciation 2026 se joue aussi sur les déclinaisons sectorielles.

## 74. Ling : l'attention hybride KDA + MLA (empruntée à Kimi)

Le Ling-3.0-flash empile **35 couches d'attention linéaire KDA (Kimi Delta Attention) + 7 couches MLA**.
Le KDA est la brique de Moonshot (voir section 42) : Ant Group l'a adoptée telle quelle dans sa famille
BailingMoeV3. C'est un cas d'école de circulation des architectures en régime open-weight — là où les
labos fermés gardent leurs recettes, les labos ouverts chinois se réutilisent mutuellement les briques
en quelques semaines. Résultat mesurable : 5,1B actifs par token pour un 124B, soit un ratio total/actifs
de 24×, et un KV cache réduit qui rend le 256K–1M servable à $0,075/M input.

## 75. Ling : HiCache, le cache hiérarchique

**HiCache** (hierarchical cache) est l'optimisation d'inférence mise en avant par Ant pour le 3.0-flash :
-60 à -80 % de TTFT (time to first token) sur les longues conversations (vendor claim, non vérifié
indépendamment). Le principe : organiser le cache en hiérarchie pour éviter de recomposer tout le contexte
à chaque tour — critique pour les agents qui accumulent des centaines de milliers de tokens d'historique
d'outils. Avec le KDA (mémoire bornée) + HiCache (réutilisation), Ant attaque le coût du long contexte
sur les deux fronts : stockage et recalcul.

## 76. Ling : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| Ling-mini-2.0 | 10/2025 | MoE, GQA | 16B / 1,4B | non vérifié | non vérifié | Ouverts (non vérifiée) |
| Ling-flash-2.0 | 10/2025 | MoE, GQA | 100B / 6,1B | non vérifié | non vérifié | Ouverts (MIT cité, tiers) |
| Ling-1T | 10/2025 | MoE, GQA | 1T / 50B | non vérifié | non vérifié | Ouverts (non vérifiée) |
| Ling-2.5-1T | 16/02/2026 | 1T, non-thinking | 1T / non vérifié | 1M | non vérifié | Ouverts (non vérifiée) |
| Ring-2.5-1T | 16/02/2026 | 1T, thinking, lin. hybride | 1T / non vérifié | non vérifié | non vérifié | Ouverts (non vérifiée) |
| Ming-Flash-Omni-2.0 | 11/02/2026 | Omni (parole/audio/musique) | non vérifié | non vérifié | non vérifié | non vérifié |
| Ling-2.6-flash | 04/2026 | MoE | 104B / 7,4B | non vérifié | non vérifié | non vérifié (poids ?) |
| Ling-3.0-flash | 23/07/2026 | MoE, KDA + MLA, 512 experts | 124B / 5,1B | 256K (>1M) | 0,075 / 0,22 | MIT |
| Ling-3.0-flash-VL | 09/09/2026 | MoE + encodeur vision | ~125B / 5,1B | 128K–256K (div.) | non vérifié | MIT |
| Ling-3.0-tiny | 09/2026 ? | MoE, KDA + Gated MLA | 7,9B / 1,3B | 128K | non vérifié | Ouverts (non vérifiée) |
| Ling-3.0-flash-Fin | 09/09/2026 | MoE (base 3.0-flash) | 124B / 5,1B | non vérifié | non vérifié | Ouverts (non vérifiée) |

## 77. Meituan LongCat : portrait d'équipe

**LongCat** (龙猫, « chinchilla »... non : « long chat ») est l'équipe IA de **Meituan**, le géant chinois
de la livraison de repas — le profil le plus inattendu du volume. GitHub : `meituan-longcat`. La lignée :
LongCat-Flash-Chat (août 2025), Flash-Thinking (sept. 2025), **LongCat-Next** (avril 2026, multimodal
natif), **LongCat-2.0** (juin 2026, 1,6T paramètres, 1M de contexte natif, entraîné sur ASIC chinois).
Meituan joue la carte « outsider qui frappe fort » : le 2.0 revendique des scores coding au niveau de
GPT-5.5 avec un entraînement 100 % domestique.

## 78. LongCat-Flash-Chat (août 2025)

