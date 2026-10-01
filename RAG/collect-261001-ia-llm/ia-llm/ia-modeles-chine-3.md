---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-3
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agentic", "agents", "apache", "attention", "benchmarks", "claude", "fine-tuning", "fp8", "gguf", "glm", "gpu", "kimi"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [204, 301]
sha256: e2652ce53f0d6b1145d29256b11d203868d79bcbc1776c616a8d522021a6d705
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

## 16. Qwen3.5-122B-A10B et Qwen3.5-35B-A3B

- **Noms exacts** : Qwen3.5-122B-A10B et Qwen3.5-35B-A3B. **Sortie : 24 février 2026**. Poids ouverts,
  Apache 2.0. 122B : 10B actifs/token, 36 couches ; 35B : 3B actifs/token, 30 couches (chiffres
  communautaires). Même architecture hybride Gated DeltaNet + MoE que le 397B.
- **Le fait marquant** : la communauté note que le 122B-A10B est « très peu différent » du 397B sur les
  benchs malgré plus de 50 % de paramètres en moins — la courbe de scaling s'aplatit et le milieu de
  gamme devient le sweet spot économique.
- **Benchmarks** (SWE-bench Verified) : 35B-A3B **70,0 %** ; 27B dense **75,0 %**.
- **Fine-tuning** : supporté par la communauté — configs Axolotl officielles (LoRA/QLoRA/FFT, FSDP2) pour
  9B, 27B, 35B-A3B, 122B-A10B ; le LoRA doit cibler `linear_attn.in_proj_qkv/z/out_proj` + experts MoE
  (`mlp.experts.gate_up_proj/down_proj`). Fine-tuning via API Alibaba : non vérifié.
- **Inférence** : état GDN ≈ 63 MiB (35B-A3B), ≈ 149 MiB (122B) — mesures communautaires.

## 17. Qwen3.5-27B (dense) et petites tailles

- **Qwen3.5-27B** : sortie 24 février 2026. **Dense 27B**, hybride GDN/attention, Apache 2.0. SWE-bench
  Verified **75,0 %** — meilleur que le 35B-A3B MoE sur ce bench, ce qui a relancé le débat dense vs MoE
  à taille moyenne. GGUF Q4_K_M ≈ 15,7 Go (ordre de grandeur communautaire pour la lignée 27B).
- **Petites tailles** (0.8B / 2B / 4B / 9B) : sortie **2 mars 2026**, denses, Apache 2.0. Destinées à
  l'embarqué et à l'inférence locale légère. Specs fines : non vérifiées au 27/09/2026.
- La gamme 3.5 complète couvre ainsi 0.8B → 397B : la plus large échelle open-weight du marché début 2026.

## 18. Qwen3.5 : l'attention Gated DeltaNet, explication

Le différenciateur technique de toute la génération Qwen 3.5/3.6/3.8 (hors Max fermés) est le **Gated
DeltaNet** : une attention linéaire (complexité O(n) au lieu de O(n²)) mélangée par blocs avec de
l'attention standard (pattern répété, ex. 3 couches GDN + 1 couche d'attention). Au lieu d'un KV cache qui
grossit avec la séquence, les couches linéaires maintiennent un **état récurrent de taille fixe**
(quelques dizaines à ~200 MiB selon la taille du modèle). Conséquence pratique : le long contexte coûte
beaucoup moins cher en mémoire, et c'est ce qui permet à des modèles ouverts de viser 256K–1M de contexte
sans ferme de GPU. C'est la même famille d'idées que le KDA de Kimi et le DSA de GLM : 2026 est l'année
où l'attention linéaire/hybride est devenue le standard chinois.

## 19. Qwen3.6-35B-A3B (« qwen3.6-flash »)

- **Nom exact** : Qwen3.6-35B-A3B (nom API Model Studio : `qwen3.6-flash`). **Sortie : 16 avril 2026**.
  Statut : disponible, poids ouverts.
- **Architecture** : MoE sparse, **35B total / 3B actifs** par token ; 256 experts (8 routés + 1 partagé) ;
  40 couches en 10 groupes de (3× Gated DeltaNet + 1× Gated Attention). Attention hybride GDN + Gated Attention.
- **Contexte** : **262 144 tokens natifs (256K)**, extensible à >1M via YaRN.
- **Fonctionnalités** : vision (texte + image + vidéo), tool use / agents, reasoning (**Thinking
  Preservation** — conservation du raisonnement entre les tours ; modes thinking/non-thinking).
- **Poids** : **ouverts, Apache 2.0** (HF `Qwen/Qwen3.6-35B-A3B`). GGUF disponible (Ollama, llama.cpp).
  Tient sur 1× H100/H200 en BF16/FP8.
- **Benchmarks** : SWE-bench Verified **73,4 %**, SWE-bench Pro 49,5 %, Terminal-Bench 2.0 51,5 %,
  LiveCodeBench v6 80,4 %, QwenWebBench (frontend) 1397.
- **Prix API** : ~$0,14 / 1M input, $1,00 / 1M output (FP8, fournisseur tiers AirCloud) — prix officiel
  Alibaba non vérifié.
- **Déploiement** : vLLM, SGLang, Ollama, llama.cpp, ModelScope, HF.

## 20. Qwen3.6-27B (dense, « open flagship »)

- **Nom exact** : Qwen3.6-27B. **Sortie : 22 avril 2026**. Statut : disponible, poids ouverts.
- **Architecture** : **dense 27B**, hybride GDN/full attention, multimodal (texte + image + vidéo).
- **Contexte** : 262K natifs d'après un comparatif tiers — source primaire Alibaba non vérifiée.
- **Poids** : **ouverts, Apache 2.0**. GGUF Q4_K_M ≈ 15,7 Go.
- **Benchmarks** : SWE-bench 77,2 % (agrégateur tiers kilo.ai — à prendre avec précaution).
- **Positionnement** : le « dense open flagship » d'Alibaba — la preuve qu'à 27B, un dense bien entraîné
  tient tête aux MoE sparse de la même classe sur le code.

## 21. Qwen3.6-Max-Preview (fermé)

- **Nom exact** : Qwen3.6-Max-Preview. **Sortie : 20 avril 2026**. Statut : disponible, **fermé, API uniquement**.
- **Importance historique** : **premier flagship Qwen à poids fermés** — le moment où Alibaba a rompu avec
  le tout-open-weight.
- **Architecture** : MoE 35B total / **3B actifs** par inférence (classe 35B-A3B) — chiffres rapportés par
  la presse, non publiés par Alibaba.
- **Contexte** : 256 000 tokens. **Fonctionnalités** : texte uniquement au lancement (pas d'image),
  agents/coding, `preserve_thinking` (traces de raisonnement conservées entre tours).
- **Poids** : fermés — pas de poids HF. **Prix API** : non vérifié. Benchmarks : meilleurs scores sur
  6 benchs coding/agents dont SWE-bench Pro et Terminal-Bench 2.0 (revendication Alibaba).
- **Déploiement** : Alibaba Cloud Bailian, Qwen Studio ; compatible API OpenAI **et** Anthropic.

## 22. Qwen3.6-Plus (fermé)

- **Nom exact** : Qwen3.6-Plus. **Sortie : 2 avril 2026** (Caixin Global). Statut : disponible, **fermé,
  API uniquement**.
- **Architecture** : non vérifié (paramètres non publiés). **Contexte : 1M tokens par défaut**.
- **Fonctionnalités** : multimodal (texte, image, vidéo, documents), agentic coding (planifie/teste/itère,
  niveau repo, frontend), « capability loop » (perception → raisonnement → action). Compatible Claude Code,
  OpenClaw, Cline ; **protocole API Anthropic supporté** (on peut pointer un setup Claude Code vers le
  modèle). Intégré à Wukong (plateforme entreprise Alibaba) et à l'app Qwen.
- **Benchmarks** : SWE-bench Verified **78,8** (vs Claude Opus 4.5 : 80,9), Terminal-Bench 2.0 **61,6**
  (devant Opus 4.5 : 59,3), GPQA **90,4** (meilleur des modèles comparés).
- **Prix API** : Bailian — 2 yuans (~$0,29) / 1M input, 12 yuans / 1M output (Caixin).

## 23. Qwen : la compatibilité « protocole Anthropic »

Détail stratégique qui traverse les fiches 3.6-Plus, 3.7-Max et 3.8 : les modèles fermés Qwen supportent le
**protocole API Anthropic** en plus du format OpenAI. Concrètement, un utilisateur peut pointer son
installation Claude Code vers un endpoint Qwen sans changer d'outillage. C'est une prise de guerre
commerciale directe contre Anthropic sur son propre terrain (les agents de code), rendue possible parce que
les harnais type Claude Code/OpenClaw parlent un protocole que n'importe quel fournisseur peut implémenter.
En 2026, l'interopérabilité des protocoles compte autant que les benchmarks.

## 24. Qwen : premier tableau intermédiaire (3.5 – 3.6)

