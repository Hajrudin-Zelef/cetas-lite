---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-2
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-04-24", "2026-09-10", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmarks", "claude", "decode", "deepseek", "fp8", "gemini", "gguf"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [104, 203]
sha256: 162d554c7c20d4f33c43c34776be5e7c345b877e364c2631a09fe2f3699dfaae
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : DeepSeek V4.1-Flash (ID API : `deepseek-flash` ; OpenRouter : `deepseek/deepseek-v4.1-flash`).
- **Date de sortie** : **10 septembre 2026**. Statut : disponible (modèle Flash actuel sur l'API DeepSeek).
- **Architecture** : MoE, **552B paramètres backbone** (763B paramètres stockés selon le repo Hugging Face),
  **8B actifs en prefill / 16B en decode**. Nouvelle famille « Causal Encoder-Decoder » (non vérifié en
  détail). Attention hybride KDA + MLA selon une source tierce (non vérifié au 27/09/2026). Un rapport
  technique mentionne aussi une « Engram memory » de 196B (non vérifié).
- **Contexte** : **1M tokens**, sortie max 384K. **Entrée image native** (max 1 024 tokens/image).
- **Fonctionnalités** : reasoning contrôlable (thinking activé par défaut, efforts low/high/max,
  désactivable), tool calling, mode agent. Web search natif : non vérifié.
- **Poids** : **ouverts, licence MIT** (Hugging Face `deepseek-ai/DeepSeek-V4.1-Flash`).
- **Benchmarks clés** (vendor-reported, effort de raisonnement max) : DeepSWE v1.1 **74,2** (vs Claude
  Opus 5 74,0, GPT-5.6 Sol 73,0), victoires sur AutomationBench et CyberGym ; derrière sur Terminal-Bench
  4.0 ; GPQA Diamond 90,9 (V4-Pro : 92,4). Artificial Analysis Intelligence Index : **40**.
- **Prix API** : **$0,15 / 1M input, $0,60 / 1M output** hors peak ; **double en peak** ($0,30 / $1,20).
  Peak : 01:00–04:00 et 06:00–10:00 UTC, jours ouvrés hors jours fériés chinois. Sur OpenRouter off-peak :
  $0,15 / $0,60, cached $0,003/M.
- **KV cache / inférence** : non vérifié au 27/09/2026 — pas de données publiques trouvées. Déploiement :
  API DeepSeek, OpenRouter (DeepInfra, Novita, Venice). Support vLLM/Ollama : non vérifié.

## 9. DeepSeek V4.1 Pro : annoncé, pas sorti

- **Nom exact** : DeepSeek V4.1 Pro (alias attendu `deepseek-v4.1-pro` — non vérifié).
- **Statut au 27/09/2026** : **annoncé, non disponible**. Aucune spec publiée, aucune date de sortie
  annoncée. Une rumeur (orcarouter.ai, 25 septembre 2026) évoque une fenêtre de sortie 28–30 septembre,
  mais rien dans le changelog officiel DeepSeek entre le 10 et le 25 septembre ne l'étaye.
- Fiche honnête : il n'y a rien de plus à dire. C'est le modèle fantôme de cette fin septembre 2026 —
  tout le marché l'attend comme la réponse de DeepSeek au Kimi K3 et au Qwen3.8-Max.

## 10. DeepSeek : l'attention hybride de la famille V4

Le changement technique le plus important de DeepSeek en 2026 est architectural : la famille V4 abandonne
le MLA (Multi-head Latent Attention) pur qui avait fait la réputation de V3 pour une **attention hybride**
combinant CSA (Compressed Sparse Attention, stride 4), HCA (Heavily Compressed Attention, stride 128) et
SWA (sliding window ~128). L'objectif affiché : diviser le KV cache par ~10 par rapport à V3.2 à contexte
égal, ce qui rend le million de tokens économiquement servable. Le revers : trois layouts de cache
simultanés à gérer côté inference engine (d'où les forks vLLM/SGLang maison et l'usage de DeepEP sur
NVLink). Réserve : ces détails viennent de notes communautaires GitHub, pas d'un rapport technique
officiel DeepSeek — le rapport V4 n'a pas été trouvé au 27/09/2026.

## 11. DeepSeek : les prix et le modèle économique

DeepSeek a inventé en 2026 la tarification « peak/off-peak » calée sur les heures ouvrées chinoises
(01:00–04:00 et 06:00–10:00 UTC en peak, prix doublés). Le V4-Pro facture $0,435/M input en cache miss
mais seulement $0,003625/M en cache hit — un ratio de 120× qui pousse à architecturer les applications
autour du cache de contexte. Le V4.1-Flash descend à $0,15/$0,60 hors peak : c'est le niveau de prix qui
a forcé Qwen3.8-Flash ($0,16/$0,47) et GLM-5.3-Flash ($0,15/$0,50) à s'aligner au dixième de cent près.
La guerre des Flash de l'été 2026 se joue littéralement à $0,01 près sur le million de tokens d'entrée.

## 12. DeepSeek : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| V4-Pro (0813) | 24/04/2026 (GA 13/08) | MoE, attention hybride | 1,6T / 49B | 1M | 0,435 / 0,87–0,93 | MIT (ouverts) |
| V4-Flash | 24/04/2026 (retiré 10/09) | MoE, NVFP4 | 284B / 13B | 1M | non vérifié | MIT (ouverts) |
| V4.1-Flash | 10/09/2026 | MoE, causal encoder-decoder | 552B / 8–16B | 1M | 0,15 / 0,60 (×2 peak) | MIT (ouverts) |
| V4.1 Pro | annoncé 10/09, pas sorti | non vérifié | non vérifié | non vérifié | non vérifié | présumé MIT |

## 13. Alibaba Qwen : portrait de l'entreprise

Alibaba (via Alibaba Cloud et l'équipe Qwen) est l'acteur chinois le plus prolifique de 2026 : quatre
générations mineures (3.5 → 3.8) en sept mois, une gamme complète du 0.8B au 2.4T paramètres, et une
stratégie à double détente qui structure tout le marché — **open-weight agressif sur les modèles
efficaces (Apache 2.0), poids fermés sur les flagships** (3.6-Max, 3.7-Max/Plus), puis retour surprise à
l'open-weight sur le flagship 3.8-Max. Alibaba vend aussi du cloud (Model Studio, QwenCloud, Bailian) et
des puces maison (Zhenwu M890, serveurs Panjiu AL128, présentés avec Qwen3.7-Max) : c'est le seul acteur
chinois à jouer simultanément le modèle, le cloud et le silicium.

## 14. La stratégie open-weight / fermé d'Alibaba en 2026

La doctrine 2026 d'Alibaba tient en trois temps. **Février–avril** : tout est ouvert (Qwen3.5 intégralement
Apache 2.0, Qwen3.6-35B-A3B et 27B ouverts). **Avril–mai** : virage « frontier-closed » — Qwen3.6-Max-Preview
est le premier flagship Qwen à poids fermés, suivi des 3.6-Plus, 3.7-Max et 3.7-Plus, tous API uniquement.
**Août** : revirement — Qwen3.8-Max sort en **poids ouverts** (annonce semaine du 10 août 2026), présenté
comme « open-weight multimodal system », et Qwen3.8-Flash est ouvert dès le jour J comme « preview de
l'architecture Qwen4 ». Lecture : Alibaba a testé le fermé pour monétiser l'agentique haut de gamme, puis
a rouvert face à la pression concurrentielle (Kimi K3 à 2.8T ouverts en juillet, GLM-5.3 ouverts en août).
Le pendule open/fermé est l'indicateur avancé de la confiance d'Alibaba dans son avance technique.

## 15. Qwen3.5-397B-A17B (flagship février 2026)

- **Nom exact** : Qwen3.5-397B-A17B. **Sortie : 16 février 2026** (Reuters). Statut : disponible, poids ouverts.
- **Architecture** : MoE hybride — **premier modèle open-weight à combiner Gated DeltaNet (attention
  linéaire) + MoE**. 397B total, **512 experts, routage top-10 + experts partagés, 17B actifs/token**.
  45 couches (chiffre communautaire). Vocabulaire ~250K, 201 langues.
- **Contexte** : natif exact non vérifié au 27/09/2026 — la communauté l'utilise à 262 144 tokens sans problème.
- **Fonctionnalités** : vision (encodeur multimodal, variantes vision-language sur HF), tool use / agents
  (« visual agentic capabilities », vidéos jusqu'à 2h d'après la presse), reasoning (modes
  thinking/non-thinking hérités de Qwen3).
- **Poids** : **ouverts, Apache 2.0** (HF `Qwen/Qwen3.5-397B-A17B`). Formats BF16/FP8 officiels, GGUF
  (Unsloth), NVFP4 (Sehyo, RedHatAI). 122B-A10B en Q4_K_M ≈ 71 Go ; Q8_0 ≈ 121 Go (ordres de grandeur
  communautaires).
- **Inférence** : Gated DeltaNet → état récurrent compact au lieu d'un KV cache plein sur les couches
  linéaires (état ≈ 186 MiB pour 397B — mesure communautaire). Têtes MTP (multi-token prediction)
  présentes, supportées par Ollama (décodage spéculatif).
- **Benchmarks** : Alibaba revendique supériorité sur GPT-5.2, Claude Opus 4.5, Gemini 3 Pro sur plusieurs
  benchs (revendication constructeur, non vérifiée indépendamment).
- **Prix API** : non vérifié — « 60 % moins cher que le prédécesseur » selon Alibaba.
- **Déploiement** : vLLM, SGLang, llama.cpp, Ollama, LM Studio, KTransformers, Megatron-Bridge (NVIDIA
  NeMo), ModelScope.

