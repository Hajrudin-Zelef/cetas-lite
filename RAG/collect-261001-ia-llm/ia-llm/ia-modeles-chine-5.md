---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-5
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "Moonshot", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-05-20", "2026-08-03", "2026-08-14", "2026-09-02"]
keywords: ["agent", "agents", "apache", "attention", "attribution", "benchmarks", "claude", "deepseek", "fine-tuning", "gguf", "hyperscaler", "kimi"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [398, 500]
sha256: 528c4b95e7e1a6f60defe38573c9dac250b3c5ac0f5231dd704d01e4582788ee
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : Qwen3.8-Flash (nom preview : Qwen3.8-Flash-Next) — présenté comme **preview explicite de
  l'architecture Qwen4**. Teaser ModelScope 25 août 2026, annonce **26 août** (Alizila), production
  `qwen3.8-flash` le **2 septembre 2026** (renommage Next → production en une semaine).
- **Statut** : disponible — poids ouverts + API. Licence exacte : non vérifiée (une source dit « Qwen
  License », une autre « non officiellement précisée » — à vérifier sur la carte HF/ModelScope).
- **Architecture** : MoE — **125B backbone + 51B table N-gram = 176B total / 6B actifs** par token ;
  Gated DeltaNet + **Qwen Sparse Attention** (preview archi Qwen4).
- **Contexte** : 262 144 natifs (extensible à ~1M via YaRN) ; déploiement Model Studio listé à 983 616
  tokens de contexte, 131 072 tokens max en sortie ; reasoning « xhigh » par défaut.
- **Fonctionnalités** : multimodal (texte, vision haute résolution, vidéo multi-frame), coding, document
  work, agents.
- **Benchmarks** : SWE-bench Pro **62,5 %**, CoWorkBench **73,9**, JobBench **55,7**, DeepSWE 58,7,
  Toolathlon 73,5, HLE (sans tools) 35,9, GPQA Diamond 91,7, LiveCodeBench v6 **91,9**, MMLU-Pro 73,23
  (chiffres constructeur/agrégateurs — vérification indépendante limitée).
- **Prix API** : **$0,16 / 1M input, $0,47 / 1M output** (QwenCloud/Model Studio) — positionné frontalement
  contre DeepSeek-V4-Flash ($0,15/$0,60).
- **Déploiement** : API QwenCloud/Model Studio (Standard ou Token Plan) ; support **day-0 SGLang** ; poids
  ouverts → vLLM/llama.cpp probables (non vérifié au 27/09).

## 31. Les tiers API « Plus / Flash / Max » démystifiés

En 2026, « Plus / Flash / Max » ne sont plus des modèles mais des **tiers API versionnés** :
`qwen3.6-flash` (= Qwen3.6-35B-A3B ouvert), `qwen3.6-plus` (fermé), `qwen3.7-max` (fermé), `qwen3.7-plus`
(fermé), `qwen3.8-max` (ouvert), `qwen3.8-flash` (ouvert). Les IDs génériques historiques `qwen-plus` /
`qwen-max` restent référencés comme défauts dans des docs tierces récentes (sept. 2026), mais le mapping
exact vers les modèles versionnés côté Model Studio n'a pas été vérifié sur la doc officielle. La doc
OpenClaw (sept. 2026) précise : `qwen3.7-plus` et `qwen3.6-plus` acceptés sur Coding Plan et Standard ;
`qwen3.8-max`, `qwen3.8-flash`, `qwen3.7-max`, `qwen3.6-flash` exigent Standard (pay-as-you-go) ou Token Plan.
Moralité : toujours vérifier l'ID exact facturé, le nom commercial ne dit plus quel modèle tourne derrière.

## 32. Qwen3-Coder-Next (février 2026)

- **Nom exact** : Qwen3-Coder-Next. **Sortie : février 2026**. Statut : disponible, poids ouverts.
- **Note de nommage** : aucun « Qwen3.5-Coder » / « Qwen3.6-Coder » / « Qwen3.7-Coder » n'a été trouvé —
  le codeur 2026 d'Alibaba s'appelle bien « Coder-Next ».
- **Architecture** : MoE **80B total / 3B actifs** ; attention hybride **Gated DeltaNet (75 % linéaire /
  25 % full)** ; contexte **256K** ; licence **Apache 2.0**.
- **Benchmarks** : SWE-bench Verified **70,6 % (SWE-Agent) → 71,1 % (MiniSWE) → 71,3 % (OpenHands)** ;
  SWE-Bench Pro 44,3 %. Entraîné par RL pour tâches agentiques (schémas d'outils style Claude Code).
  Rapport technique : arXiv 2603.00729.
- **Prix API** : ~$0,11 / $0,80 par 1M (tiers). **Déploiement** : Ollama, LM Studio, llama.cpp, vLLM, mlx-lm.

## 33. Qwen3-Coder-480B-A35B

- **Nom exact** : Qwen3-Coder-480B-A35B. Statut : disponible, poids ouverts, **Apache 2.0**.
- **Architecture** : MoE **480B total / 35B actifs**. Contexte 256K (1M par extrapolation).
- **Benchmarks** : SWE-bench Verified **69,6 %**.
- **Prix API** : ~$0,22 in / $0,90 out par 1M (tiers).
- Fait notable : avec 35B actifs par token, c'est le codeur ouvert le plus « dense en actifs » de la gamme
  Qwen 2026 — puissance brute plutôt qu'efficacité, l'inverse du positionnement du Coder-Next (3B actifs).

## 34. Qwen : le déploiement local, mode d'emploi

La gamme ouverte Qwen est la mieux supportée côté self-hosted de tout l'écosystème chinois 2026 : vLLM,
SGLang, llama.cpp, Ollama (support day-0 sur les 27B), LM Studio, KTransformers, Megatron-Bridge, MLX sur
Apple Silicon, ModelScope. Points d'attention : les modèles GDN nécessitent des kernels compatibles
attention linéaire (vérifier la version de vLLM/SGLang) ; les têtes MTP activent le décodage spéculatif
sous Ollama ; les GGUF communautaires (Unsloth, ggml-org, bartowski) sont la voie rapide pour le local.
Pour le fine-tuning, les configs Axolotl officielles (LoRA/QLoRA/FFT, FSDP2) ciblent les projections de
l'attention linéaire + les experts MoE — ne pas appliquer une recette LoRA standard sans vérifier les
noms de modules.

## 35. Qwen : tableau comparatif (3.7 – 3.8 + Coder)

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| Qwen3.7-Max | 20/05/2026 | non vérifié (fermé) | non vérifié | 1M | ~1,25 / 3,75 (tiers) | Fermé |
| Qwen3.7-Plus | 20/05/2026 | non vérifié (fermé) | non vérifié | 1M (non vérifié) | 0,40 / 1,60 | Fermé |
| Qwen3.8-Max | 03/08/2026 (snap. 02/09) | MoE | 2,4T / 95B | 1M | 2,00 / 6,00 | Ouverts (licence non vérifiée) |
| Qwen3.8-27B | 14/08/2026 | Dense, 48 GDN + 16 full | 27,78B | 256K (>1M YaRN) | non vérifié | Apache 2.0 |
| Qwen3.8-Flash | 02/09/2026 | MoE, GDN + Qwen Sparse Attn | 176B / 6B | 256K (~1M YaRN) | 0,16 / 0,47 | Ouverts (licence non vérifiée) |
| Qwen3-Coder-Next | 02/2026 | MoE, GDN 75/25 | 80B / 3B | 256K | ~0,11 / 0,80 (tiers) | Apache 2.0 |
| Qwen3-Coder-480B-A35B | 2026 | MoE | 480B / 35B | 256K (1M extrap.) | ~0,22 / 0,90 (tiers) | Apache 2.0 |

## 36. Qwen : ce que la gamme 2026 raconte

Trois leçons. Un, l'échelle n'est plus un argument : le 122B-A10B fait jeu égal avec le 397B-A17B, le
27B dense bat des MoE plus gros — l'efficacité d'entraînement a rattrapé la taille brute. Deux, Alibaba
a industrialisé le « snapshot » (3.8-Max-0902) : le modèle n'est plus une release mais un flux, avec des
mises à jour qui changent le comportement et le coût. Trois, la compatibilité protocolaire (OpenAI +
Anthropic) et les tiers API versionnés montrent qu'Alibaba vend désormais une plateforme, pas des modèles.

## 37. Moonshot AI : portrait de l'entreprise

Moonshot AI (Pékin), fondée par Yang Zhilin, est le spécialiste chinois de l'agentique à grande échelle.
Sa marque modèle est **Kimi** (plateforme : platform.kimi.ai). La lignée 2026 va de K2.5 (27 janvier) à
K2.6 (20 avril), K2.7 Code (12 juin) puis **K3 (juillet)** — ce dernier étant, à sa sortie, le plus gros
modèle open-weight de l'histoire (2,8T paramètres). Moonshot a deux signatures : les essaims d'agents
(jusqu'à 300 agents coordonnés sur 4 000 étapes) et une licence maison, la « Modified MIT », qui mérite
sa propre section.

## 38. La licence « Modified MIT » de Moonshot : ce qu'elle change

Les poids Kimi sont ouverts sous **licence Modified MIT** : usage commercial, modification, fine-tuning
et revente permis sans royalties — comme la MIT classique — **plus une clause d'attribution commerciale** :
si le produit dépasse **100M d'utilisateurs mensuels ou $20M de revenus mensuels**, l'interface doit
afficher « Kimi K2.6 » (ou le nom du modèle). C'est permissif pour 99,9 % des usages (recherche, PME,
self-hosted, intégration produit), mais ce n'est pas une MIT pure : un hyperscaler qui voudrait servir
Kimi à des centaines de millions d'utilisateurs devrait créditer Moonshot en façade. Pour Zelef et son
RAG personnel : aucun impact, la clause ne se déclenche qu'à l'échelle industrielle.

## 39. Kimi K2.6

