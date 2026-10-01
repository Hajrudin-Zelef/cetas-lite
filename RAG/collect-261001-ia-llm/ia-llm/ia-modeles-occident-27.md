---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-27
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "SGLang", "SpaceX", "TensorRT-LLM", "vLLM", "xAI"]
dates: ["2026-02-09", "2026-05-28", "2026-07-18", "2026-08-31", "2026-09-04", "2026-09-14", "2026-09-27", "2027-02-04"]
keywords: ["agent", "agents", "apache", "astra", "attention", "blackwell", "chatgpt", "claude", "compute", "consumer", "copilot", "fp8"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [2062, 2148]
sha256: 1980a3d895f9dd538880b1141138f46b8431d5f0b0d0d4c9792d4750ffd052a0
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Licence | Modèles | Usage commercial | Redistribution | Points d'attention |
|---|---|---|---|---|
| **Apache 2.0** | Gemma 4 (toutes), Mistral Small 4, Granite 4.1/4.2, Laguna XS.2 | ✅ | ✅ | la plus sûre juridiquement ; brevets couverts |
| **Modified MIT** | Mistral Medium 3.5 | ✅ avec carve-out | ✅ avec carve-out | **carve-out : restrictions au-delà d'un seuil de revenus** — lire la clause avant usage en grande entreprise |
| **OpenMDW-1.1** (Linux Foundation) | Laguna XS 2.1, S 2.1, Nemotron 3.5 Lightning | ✅ (permissive, rapporté) | ✅ | licence récente (2026), permissive ; vérifier le texte par repo |
| **Llama 4 Community License** (custom) | Llama 4 Scout, Maverick | ✅ avec conditions | ✅ avec conditions | custom Meta : pas l'Apache 2.0 ; seuils d'usage et restrictions spécifiques |
| **NVIDIA Open Model License** | Nemotron 3 Nano/Super/Ultra, Cascade-2 | selon version | selon version | « version Nemotron » — lire le LICENSE du repo |
| **NVIDIA Open Model Agreement** | Nemotron 3 Nano Omni | selon version | selon version | **diffère** de l'Open Model License — ne pas confondre |
| **MIT** | Hermes Agent (harness) | ✅ | ✅ | le harness est MIT, les modèles Hermes ont leur propre licence permissive |
| **Licence non vérifiée / non publiée** | Robostral Navigate (poids non publiés), MiniMax-M2 (nom exact non trouvé — volume Chine), Voxtral TTS (CC BY-NC 4.0 selon une source secondaire non confirmée) | — | — | ne pas déployer sans licence lue |

Règle d'or : **la licence se lit par repo, pas par famille** — la famille Nemotron en compte trois différentes. Et « open-weight » ≠ « open source » : les poids sont téléchargeables, mais les données d'entraînement ne sont (presque) jamais publiées.

## 161. Dossier : l'inférence locale — quel hardware pour quel modèle

Données vérifiées (sources : vendors, fiches HF, tests communautaires) :

| Modèle | Format le plus léger vérifié | Hardware minimal cité | Contexte servi |
|---|---|---|---|
| Nemotron 3.5 Lightning 30B-A3B | NVFP4 ~22 Go | 1× H100 80GB, 1× A100 80GB, DGX Spark (GB10), RTX 5090 32GB (selon config) | jusqu'à 1M (256K documenté single-H100) |
| Gemma 4 26B-A4B | quantifié (24 Go consumer) | GPU consumer 24 Go | 256K |
| Gemma 4 31B | BF16 | 1× H100 80GB | 256K |
| Gemma 4 12B | quantifié | laptop ~16 Go (texte) | 256K |
| Laguna S 2.1 | NVFP4 ~71 Go | 1× NVIDIA DGX Spark | 1M |
| Laguna XS.2 | — | 1 seul GPU / Mac 36 Go | 131K |
| Mistral Small 4 | NVFP4 officiel | 4×H100 / 2×H200 / 1×DGX B200 min (secondaire) | 256K |
| Mistral Medium 3.5 | GGUF (bartowski) | dès 4 GPUs (secondaire) ; vLLM TP8 recommandé | 256K |
| Granite 4.2 8B | GGUF officiel, Ollama `granite4.2:8b` | single-GPU 24–32 Go (sweet spot) | 512K (cible training) |
| Granite 4.2 3B | — | laptop | — |
| Llama 4 Scout | Int4 | single H100 | 10M |
| Llama 4 Maverick | FP8 | 1 hôte DGX H100 (8×H100) | 1M |
| Hermes 4.3-36B | GGUF | consumer hardware (revendiqué) | 128K–512K ? |

Serveurs cités : vLLM, SGLang, llama.cpp, Ollama, TensorRT-LLM, MLX, LiteRT-LM (on-device), NIM. Note : le NVFP4 (Blackwell) est devenu en 2026 le format de référence de l'inférence efficiente (Gemma 4, Nemotron, Small 4, Laguna S 2.1 le proposent en officiel).

## 162. Dossier : les agents de code — comparatif des harnesses

| Harness | Éditeur | Modèle(s) par défaut | Licence | Particularité vérifiée |
|---|---|---|---|---|
| Claude Code | Anthropic | Opus 4.6→5.5 (défaut Opus), Sonnet 5 (Pro) | propriétaire | agent teams, Dynamic Workflows (4.8), `/model fable` opt-in |
| Codex (CLI + cloud) | OpenAI | GPT-5.3-Codex (LTS), 5.4 (jusqu'au 31/08/2026), 5.6 Terra/Luna, GPT-6 Sol | propriétaire | apply-patch, hosted shell ; LTS = argument entreprise |
| Muse Code | Meta | Muse Spark 1.2+ | propriétaire (bêta) | co-entraîné avec le harness ; terminal macOS/Linux |
| `pool` + Shimmer | Poolside | Laguna (XS.2 → S 2.1) | open-source (ACP) | sandbox cloud Shimmer ; protocole ACP |
| Grok Build | xAI | Grok 4.5 (défaut depuis 18/07/2026), 4.6, 4.7 | source-available (contributions externes refusées) | terminal-native, headless `grok -p`, support ACP |
| Hermes Agent | NousResearch | model-agnostic (OpenRouter, Anthropic, OpenAI, Nous…) | **MIT** | GEPA self-evolution, mémoire 3 couches, 20 plateformes messaging |
| Vibe (Code) | Mistral | Mistral Medium 3.5 (agents distants) | propriétaire | Le Chat rebrandé : Work / Code / Chat unifiés |
| GitHub Copilot | GitHub/Microsoft | GPT-5.3-Codex (GA 09/02/2026), GPT-6 Astra (GA 04/09/2026), Grok 4.6 | propriétaire | LTS Codex garanti jusqu'au 04/02/2027 (Business/Enterprise) |

## 163. Dossier : les chiffres d'entreprise 2026 — à manier avec prudence

| Fait | Source citée | Fiabilité |
|---|---|---|
| xAI absorbée par SpaceX ; SpaceX IPO juin 2026 (SPCX, ~$1 800 Mds, $75 Mds levés) | ainvest.com, sept. 2026 | secondaire — chiffres à recouper |
| Musk acquiert Cursor pour $60 Mds (juin 2026) | mui-api | secondaire — à recouper |
| Grok Bot : ~418K utilisateurs hebdo (14/09/2026) | Bloomberg via PYMNTS | presse — plausible |
| Anthropic Series H clôturée (28/05/2026, synchro Opus 4.8) | pilot-shell / the-ledger | secondaire |
| Mistral acquiert Koyeb (fév. 2026) et Emmi AI (~€300M, mai 2026) | redhat-et, presse | secondaire |
| Mistral Compute : 18 000 GPU Blackwell, 44 MW près de Paris | secondaire | non vérifié indépendamment |
| Poolside : levées $626M / $2Mds / $3Mds selon sources | toknow.ai, pathfounders | **contradictoires — à traiter avec prudence** |
| Meta : bug bounty Muse jusqu'à $300K ($130K prompt-injection) | explainx.ai | secondaire |
| Anthropic : $100M crédits + $4M dons (Glasswing) | sources Glasswing | secondaire mais concordant |

Aucun de ces chiffres n'est audité ici : ils éclairent les stratégies (verticalisation xAI/SpaceX, infra Mistral, pivot Meta), pas des valorisations à citer telles quelles.

## 164. Dossier : ce qui a été cherché et NON TROUVÉ — l'inventaire honnête

| Nom recherché | Requêtes | Verdict au 27/09/2026 |
|---|---|---|
| « Claude Sonnet 4.8 » | 2 angles | **NON TROUVÉ** — confusion probable avec Opus 4.8 |
| « Claude Haiku 5 » | 2 angles | **NON TROUVÉ** — le catalogue Vertex rejette l'ID |
| « Gemini 3.5 Pro » | 3 angles | **NON TROUVÉ** — « still nowhere » (plusieurs sources) |
| « Gemini 4 » | 3 angles | **NON TROUVÉ** — aucune annonce |
| « Veo 4 » | 2 angles | **NON TROUVÉ** — remplacé par Gemini Omni |
| « GPT-6 Terra » | 1 angle ciblé | **NON SORTI** — existence future non vérifiée |
| « GPT Pro » (modèle standalone) | 3 angles | **NON TROUVÉ** — = abonnement + variantes -Pro |
| « GPT-5.3 » flagship (sans suffixe) | 2 angles | **NON TROUVÉ** — la 5.3 est Codex + Instant |
| « Poolside Malibu » | 2 angles | **NON TROUVÉ** — « no Malibu family found » |
| « Granite 4.x Code » 2026 | 2 angles | **NON TROUVÉ** — couvert par instruct + CodeAlchemy |
| Muse Glimmer 30B (open-weight) | 1 source unique | **À RECOUPER** — non vérifié au-delà de prompt-atlas |
| « Watermelon » (Meta) | cité par Alexandr Wang | **NON VÉRIFIÉ** — nom de code interne uniquement |
| GPT-6 « aeon » | rumeur TestingCatalog 24–26/09 | **NON VÉRIFIÉ** |
| « ChatGPT Pro Max » $500/mois | rumeur TestingCatalog 24/09 | **NON VÉRIFIÉ** |
| Autre déclinaison Llama 4 en 2026 | 2 angles | **NON TROUVÉ** — rien d'officiel |
| IDs Gemini 3.8 TTS | guide de migration uniquement | **NON VÉRIFIÉ** au-delà |

## 165. Dossier : les conflits de sources non tranchés — récapitulatif du volume

