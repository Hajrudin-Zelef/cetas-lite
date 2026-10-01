---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-13
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Huawei", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-06-01", "2026-07-06", "2026-07-17", "2026-09-03", "2026-09-27"]
keywords: ["agent", "agents", "apache", "ascend", "attention", "benchmark", "benchmarks", "compute", "deepseek", "fine-tuning", "fp8", "gqa"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1173, 1253]
sha256: cdae6768b18ad134257180b94422dcb4a93d59d62dd0acc9f45fdbbd9c3f83b5
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : humain-m3. Annoncé le **3 septembre 2026** à LEAP par **HUMAIN** (PIF, Arabie saoudite).
- **Architecture** : **428B paramètres MoE**, issu de la lignée MiniMax-M3, pré-entraîné sur **>1T tokens
  de contenu arabe natif**, meilleur score moyen sur 7 benchmarks arabes publics.
- **Statut** : research preview sur HUMAIN Node ; **open-weight release planifiée** — non vérifié si
  effective au 27/09/2026.
- Ce n'est pas un modèle MiniMax pur mais un **dérivé vérifié** : il illustre la stratégie d'export du
  savoir-faire chinois via partenariat souverain (modèle arabe frontier construit sur base chinoise
  ouverte). À surveiller : si la release open-weight se confirme, ce sera le premier frontier arabe
  ouvert de cette génération.

## 91. MiniMax : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| MiniMax-M2 → M2.7 | 10/2025 → 03/2026 | MoE, full GQA, MTP | 230B / 10B | ~192–205K | 0,30 / 1,20 | Ouverts (nom non vérifié) |
| MiniMax-M3 | 01/06/2026 | MoE | ~440B ? (non vérifié) | 1M | non vérifié | Ouverts (non vérifiée) |
| MiniMax-H3 | 17/07/2026 (poids 03/08) | Dense 33B omni vidéo+audio | 33B (dense) | multimodal natif | non vérifié | Community (territoires exclus) |
| humain-m3 (dérivé) | 03/09/2026 | MoE (lignée M3) | 428B / non vérifié | non vérifié | non vérifié | Preview (open-weight planifié) |

## 92. Tencent : portrait de l'entreprise (Hunyuan)

Tencent, via sa division **Hunyuan**, est le dernier géant chinois à avoir dégainé en 2026 — mais avec une
cible claire : **l'agent logiciel**. **Hy3 Preview** (avril 2026) puis **Tencent Hy3 / Hunyuan 3**
(**6 juillet 2026**) sont explicitement optimisés pour les workloads agents (coding repo-scale, CI/CD,
multi-step tool use, workflows financiers). ID API : `hy3`. Tencent joue la carte de l'efficacité
ciblée : 295B total / 21B actifs, Apache 2.0, support day-0 SGLang/vLLM — un positionnement « le meilleur
agent ouvert par dollar de compute » plutôt que la course au billion de paramètres.

## 93. Hy3 Preview (avril 2026)

- **Nom exact** : Hy3 Preview. **Sortie : avril 2026**. Statut : précurseur remplacé par Hy3 production
  le 6 juillet 2026.
- C'est la version « bêta publique » qui a permis à Tencent de roder son positionnement agent avant la
  release : function calling, structured output, context caching, preserved reasoning étaient déjà là.
- Specs détaillées : non vérifiées séparément au 27/09/2026 — la fiche Hy3 ci-dessous fait foi pour la
  lignée.

## 94. Tencent Hy3 (Hunyuan 3)

- **Nom exact** : Tencent Hy3 (Hunyuan 3) — ID API : `hy3`. **Sortie : 6 juillet 2026** (successeur de
  Hy3 Preview, avril 2026). Statut : disponible (poids ouverts + API).
- **Architecture** : **MoE — 295B total / 21B actifs par token** ; 192 experts, 8 actifs par token
  (+ expert partagé selon sources) ; + **3.8B paramètres de couches MTP** (multi-token prediction).
- **Attention** : type non vérifié au 27/09/2026.
- **Contexte** : **256K** (input max 192K, output max 128K).
- **Fonctionnalités** : **hybrid reasoning** — modes `no_think` / `think_low` / `think_high` ; function
  calling, structured output, context caching, preserved reasoning ; conçu pour agents logiciels
  (repo-scale coding, CI/CD, multi-step tool use), knowledge work long-contexte, workflows financiers.
- **Poids / licence** : **ouverts, Apache 2.0** ; checkpoint FP8 officiel ; support **day-0 SGLang**
  (FP8 + speculative decoding) ; vLLM ≥ 0.23 (support natif arch `hy_v3` / `hy_v3_mtp`) ; variantes
  communautaires NVFP4 (`kodelow/Hy3-NVFP4-W4A16`) ; tutoriel SGLang sur Ascend NPU (910B, 2× Atlas 800I A2).
- **Inférence** : ~21.8 tok/s single-stream sur 2× DGX Spark (NVFP4, test communautaire) ; 181 Go en
  NVFP4-W4A16 ; contexte 128K validé en pratique (256K = stretch, selon runbook communautaire).
- **Benchmarks** (Tencent, directionnels, non répliqués indépendamment) : SWE-bench Verified **78.0**,
  SWE-bench Pro **57.9**, GPQA Diamond **90.4**, HLE **53.2**.
- **Prix API** : fenêtre **gratuite** (Nous Portal + OpenRouter `:free` du ~6 au ~21 juillet 2026, fenêtre
  fermée). Prix courant : non vérifié au 27/09/2026.
- **Fine-tuning** : poids ouverts → fine-tuning/quantization/RL post-training possibles en self-hosted ;
  fine-tuning via API : non vérifié.
- **Déploiement** : SGLang, vLLM, Ollama/MLX (quant 4-bit, « Ultra-only » = très gros hardware), Nous
  Portal, OpenRouter.

## 95. Hy3 : l'agent comme cible de design

Là où DeepSeek optimise le ratio benchmark/prix et Kimi l'essaim d'agents, Tencent optimise le **workflow
agent complet** : `preserved reasoning` (le raisonnement survit entre les appels d'outils), context caching
(le cache survit entre les tours), structured output garanti, et trois niveaux de thinking commutables
(`no_think` pour les appels triviaux, `think_high` pour la planification). C'est une philosophie
« infra-agent » : le modèle est conçu comme un composant logiciel fiable plutôt que comme un cerveau
brillant. Le tutoriel officiel SGLang sur **Ascend NPU** (910B, 2× Atlas 800I A2) montre que Tencent
prépare aussi l'inférence domestique — comme Zhipu et Xiaomi, la pile chinoise se découple de NVIDIA.

## 96. Tencent : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| Hy3 Preview | 04/2026 | MoE (non vérifié) | non vérifié | non vérifié | — (preview) | non vérifié |
| Tencent Hy3 | 06/07/2026 | MoE, +3.8B MTP | 295B / 21B | 256K (192K in / 128K out) | gratuit 6→21/07, puis non vérifié | Apache 2.0 |

## 97. Le grand tableau comparatif Chine (février → septembre 2026)

