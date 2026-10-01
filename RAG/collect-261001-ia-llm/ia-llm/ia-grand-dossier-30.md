---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-30
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "Cohere", "DeepSeek", "Google", "Groq", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "StepFun", "United States", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-07-31", "2026-09-27"]
keywords: ["apache", "cohere", "deepseek", "diffusion", "embeddings", "glm", "kimi", "license", "llama", "mistral", "moe", "muse"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2063, 2141]
sha256: 9c13a30f95f2ad869f92e6f88261352d1684f39714adb5211cb3e152738e564f
---

# IA — Le grand dossier

### 2.5. Mistral AI (France — le champion européen)

| Modèle (API La Plateforme) | Entrée /1M | Sortie /1M | Notes |
|---|---|---|---|
| **Mistral Large 3** | 0,50 $ | 1,50 $ | Flagship, 256K contexte |
| **Mistral Medium 3.5** | 1,50 $ | 7,50 $ | Licence MIT modifiée (plafond de revenus) — à vérifier |
| **Mistral Small** (ex 3.1) | ~0,10 $ | ~0,30 $ | Edge / haut volume (à vérifier) |
| **Devstral 2** (123B) | 0,40 $ | 0,90 $ | Code, 256K ; retiré le 31/07/2026 → remplacé par Small 4 (à vérifier) |
| **Mistral Nemo** | 0,02 $ | 0,04 $ | Le moins cher |

Double jeu assumé : modèles **ouverts** (Apache 2.0/MIT pour Small, Nemo, Devstral…) + API propriétaire européenne (**La Plateforme**, hébergement UE — argument souveraineté pour les appels d'offres publics). Produit grand public : **Le Chat**.

### 2.6. DeepSeek (Chine — l'agresseur prix)

| Modèle (API officielle) | Entrée /1M | Sortie /1M | Notes |
|---|---|---|---|
| **DeepSeek V4-Pro** | 0,435 $ | 0,87 $ | MIT ; 1,6T params MoE (49B actifs), 1M contexte ; cache-hit : ~0,0036 $/1M |
| **DeepSeek V4-Flash** | 0,14 $ | 0,28 $ | MIT ; version légère, low-latence |
| (historique) V3.2 | 0,27–0,28 $ | 0,42 $ | Cache-hit -90 % |

**Plages horaires** : tarifs « off-peak » vs « peak » (jours ouvrés 01h–04h et 06h–10h UTC) — le prix dépend de l'heure, fait rare à intégrer dans un routeur. Poids **MIT** : la licence la plus permissive du marché pour un modèle de ce niveau. C'est le fournisseur qui a cassé les prix mondiaux en 2025–2026 et forcé OpenAI/Anthropic à baisser les leurs.

### 2.7. Alibaba — Qwen (Chine)

| Modèle (API Model Studio) | Entrée /1M | Sortie /1M | Notes |
|---|---|---|---|
| **Qwen3.7 Max** | 1,25 $ | 3,75 $ | 1M contexte |
| **Qwen3-8B** | 0,05 $ | 0,20 $ | Entrée de gamme |
| **Qwen3-Coder-480B** (ouvert) | — | — | Apache 2.0 ; ~69,6 % SWE-bench Verified |
| **Qwen3-Coder-Next** (80B/3B actifs, ouvert) | — | — | Apache 2.0 ; ~70,6 % SWE-bench à 3B actifs seulement |

Stratégie : **tout ouvrir en Apache 2.0** (Qwen3, Qwen3.5…) pour créer un standard de fait, puis monétiser le cloud (Alibaba Cloud Model Studio) et l'écosystème. C'est la famille open-weight la plus complète en 2026 (chat, code, vision, audio, MoE).

### 2.8. Moonshot AI (Chine — Kimi)

- **Kimi K3** (juil. 2026) : MoE ~2,8T params, 1M contexte, licence Kimi K3 (propre) ; API ~3,00 $/1M in (cache 0,30 $) → ~15,00 $ out — à vérifier.
- **Kimi K2.5/K2.6** : MIT modifiée (plafond 100M MAU) — la clause qui a fait débat ; Kimi K2.6 : ~71,6 % multi-attempt en code agentique.
- **Kimi K2.7-Code** : 1T/32B, ~0,95 $/1M in → 4,00 $ out — à vérifier.
- API : `platform.moonshot.ai` (+ `moonshot.cn` pour la Chine) ; hébergé aussi sur Groq, Together, OpenRouter.

### 2.9. Zhipu AI / Z.ai (Chine — GLM)

- **GLM-5.2** : MIT, 1M contexte, ~1,40 $/1M in → 4,40 $ out ; premier open-weight à battre GPT-5.5 sur SWE-Bench Pro (fait rapporté, à vérifier).
- **GLM-5.3** (753B) : licence GLM-5.3 propre (**abandon de MIT** — fait marquant : la permissivité recule sur le haut de gamme), ~0,66–1,98 $/1M in selon variantes — à vérifier.
- **GLM-5.3-Flash** (321B) : MIT, ~0,10 $/1M in — à vérifier.
- API : BigModel (`bigmodel.cn` / `z.ai`).

### 2.10. Les autres labs à connaître

| Lab | Pays | Modèle phare (sept. 2026) | Licence / accès |
|---|---|---|---|
| **Meta** | US | Llama 4 Scout/Maverick | **Llama Community License** (restrictive : clauses d'usage, pas OSI) ; **Muse Spark** = ligne fermée |
| **Microsoft** | US | Phi-4 / Phi-5 (à vérifier) | MIT — les petits modèles les plus propres juridiquement |
| **Cohere** | CA | Command A / R7B | Propriétaire ; API 2,50 $/10,00 $ (Command A) ; embeddings réputés |
| **AI21 Labs** | IL | Jamba 1.7 Large | Propriétaire ; 2,00 $/8,00 $, 256K, architecture hybride Mamba |
| **MiniMax** | CN | MiniMax M3 (428B/23B) | Licence territoriale restrictive (M2.5) ; API ~0,30 $/1,20 $ — à vérifier |
| **ByteDance (Seed)** | CN | Doubao-Seed / Seedance | API Volcengine ; tarifs agressifs Asie |
| **Tencent** | CN | Hunyuan 3 (HY3, 295B/21B) | **Apache 2.0** (fait vérifié) |
| **01.AI** | CN | Yi | Ouvert (Yi-1.5 Apache 2.0 historiquement) |
| **StepFun** | CN | Step-3.5-Flash | Apache 2.0 — à vérifier |
| **Xiaomi** | CN | MiMo-V2.5/2.6-Pro (1,02T/42B) | MIT (Pro) ; ~0,30–0,75 $/1M in — à vérifier |
| **Meituan** | CN | LongCat-2.0 (1,6T/48B) | MIT ; ~0,75 $/2,95 $ — à vérifier |
| **NVIDIA** | US | Nemotron 3 (Nano/Super/Ultra) | Licence Nemotron OML ; distillations ouvertes de modèles fermés |
| **IBM** | US | Granite | Apache 2.0 ; orienté entreprise (watsonx) |
| **Liquid AI** | US | LFM2 / 2.5 | LFM Open v1.0 (plafond 10M $ revenus) |
| **Poolside** | FR/US | Laguna S 2.1 (118B/8B) | Code ; Malibu annoncé non trouvé au 27/09/2026 |
| **Inception Labs** | US/IL | Mercury-2.5 (dLLM, diffusion) | API ~0,20 $/0,75 $ (promo 0,04 $/0,15 $ sept. 2026) — les modèles de diffusion changent le paradigme latence |
| **Black Forest Labs** | DE | FLUX.1 / FLUX.2 | Image ; licence non-commerciale sur certains poids — à vérifier |

> **Lecture sysadmin** : le marché des labs s'est polarisé. En haut, 4 fermés (OpenAI, Anthropic, Google, xAI) ; en bas, une **dizaine de familles open-weight chinoises + Qwen en Apache 2.0/MIT** qui tirent les prix vers zéro. Entre les deux, Mistral joue l'européen et Meta le « presque ouvert ». Pour un RAG d'entreprise, le choix n'est plus « quel lab » mais « quel couple (modèle, hébergeur) au meilleur prix pour la qualité requise » — exactement ce qu'un routeur automatise.

---

## 3. Top 50 des fournisseurs IA vérifiés

Méthode : chaque entrée ci-dessous **existe réellement** (site/API vérifié ou documenté par des sources recoupées au 27/09/2026). Colonnes : **spécialité** (ce qu'ils font le mieux), **modèle économique** (comment ils facturent), **positionnement** (quand les choisir). Les prix sont des ordres de grandeur vérifiés en septembre 2026 — toujours confirmer sur la page pricing du fournisseur avant d'engager un budget.

### 3.1. Le tableau — 50 fournisseurs

