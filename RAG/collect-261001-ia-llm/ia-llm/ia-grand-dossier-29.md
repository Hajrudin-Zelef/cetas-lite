---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-29
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Cohere", "DeepSeek", "Glasswing", "Google", "Groq", "Microsoft", "Mistral", "OpenAI", "Oracle", "United States", "vLLM", "xAI"]
dates: ["2026-06-17", "2026-09-01", "2026-09-03", "2026-09-10", "2026-09-15", "2026-09-24", "2026-11-21", "2026-12-11"]
keywords: ["agent", "agents", "apache", "astra", "aws", "bedrock", "chatgpt", "claude", "cohere", "copilot", "deepseek", "distribution"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1959, 2062]
sha256: 621495ed9bed66319408cdc1bf570695567f9a693bc88d78608847d07a64208f
---

# IA — Le grand dossier

| Offre | Description | Modèle économique |
|---|---|---|
| **Azure OpenAI Service** | GPT (OpenAI) hébergé sur Azure, avec les contrôles d'entreprise Microsoft (VNet, clés managées, résidence UE) | $/1M tokens (tarifs OpenAI + options PTU) |
| **Azure AI Foundry** (ex-AI Studio) | Catalogue multi-modèles : GPT, Claude (via « Microsoft Foundry »), Llama, Mistral, Phi, DeepSeek… + outils d'évaluation, fine-tuning, agents | $/1M tokens ou PTU (Provisioned Throughput Units) |
| **Instances GPU** | ND H100 v5 (~6,98 $/h/GPU), NC, NV ; partenariat OpenAI historique | $/heure ; 8×H100 ~26,8 $/h (sept. 2026, à vérifier) |
| **Microsoft Phi** | Petits modèles ouverts (MIT) : Phi-4-mini et suivants, pensés pour l'edge et le CPU | Poids gratuits ; API ~0,075 $/1M in / 0,30 $/1M out (à vérifier) |

**Positionnement** : le cloud des entreprises Microsoft (AD, M365, Copilot). Azure AI Foundry est l'équivalent de Bedrock côté Microsoft. Les **PTU** (débit provisionné) sont l'option « SLA de latence » : on réserve un débit tokens/s au lieu de subir le trafic partagé — pertinent pour un RAG en production interne.

### 1.3. Google Cloud

**Ce qu'ils vendent en IA :**

| Offre | Description | Modèle économique |
|---|---|---|
| **Vertex AI** | Plateforme unifiée : API Gemini, Model Garden (Llama, Mistral, Claude, DeepSeek selon régions), entraînement, déploiement | $/1M tokens ; ex. Gemini 3.1 Pro ~2,00 $/1M in / 12,00 $/1M out (vérifié sept. 2026) |
| **Google AI Studio** | Console + API directe Gemini, **free tier généreux sans carte bancaire** (quotas journaliers) | Gratuit (limites) puis $/1M tokens |
| **GCE GPU** | A3 (H100, ~11,06 $/h/GPU on-demand — le plus cher des hyperscalers), A2 (A100), G2 (L4) | $/heure |
| **TPU v5e/v5p/Trillium** | Puces maison, excellentes pour l'inférence à grande échelle et l'entraînement JAX/PyTorch | $/heure/chip, souvent -30 % vs GPU équivalent |
| **Gemma** | Modèles ouverts Google (Apache 2.0 pour Gemma 4), pendant « local » de Gemini | Poids gratuits |

**Positionnement** : le meilleur rapport qualité/prix sur les **gros contextes** (fenêtres 1M tokens natives) et le **free tier** le plus utilisable pour prototyper. **Gotcha documenté** : les prompts du free tier peuvent servir à améliorer les produits Google — jamais de données confidentielles sur le tier gratuit ; le tier payant est en zéro-rétention (vérifié sept. 2026).

### 1.4. Oracle Cloud (OCI)

**Ce qu'ils vendent en IA :**

| Offre | Description | Modèle économique |
|---|---|---|
| **OCI Generative AI** | API managée : Cohere Command, Llama, etc. | $/1M tokens |
| **GPU bare metal** | BM.GPU.H100.8 : 8×H100 bare metal avec RDMA 2×200 Gb/s, ~10,00 $/h/GPU (juil. 2026, à vérifier) | $/heure |
| **OCI Data Science** | Notebooks, entraînement, déploiement | $/heure |

**Positionnement** : agressif sur les prix du bare metal GPU pour les charges HPC/IA, et historiquement le partenaire d'infrastructure d'OpenAI (avec Microsoft). Intéressant pour du **serveur dédié d'inférence** (style vLLM multi-GPU) sans la taxe des trois grands.

### 1.5. Les hyperscalers en une phrase chacun

- **AWS** : le supermarché (Bedrock) + la puce maison la moins chère (Trainium/Inferentia) ; gare à l'egress.
- **Azure** : le cloud des DSI Microsoft ; Foundry + PTU pour la prod régulée.
- **Google Cloud** : gros contextes et free tier ; TPU pour l'échelle ; Gemma pour le local.
- **Oracle** : bare metal GPU agressif ; second couteau devenu premier sur l'infra IA lourde.

---

## 2. Les labs : qui crée les modèles

Distinction fondamentale pour ton RAG : le **lab** conçoit et entraîne le modèle ; le **provider** le sert via une API. Un même modèle (ex. Llama 4, DeepSeek V4) est servi par une dizaine de providers à des prix différents — d'où l'intérêt des routeurs (section 4).

### 2.1. OpenAI (États-Unis, fermé sauf GPT-OSS)

Modèles API au **15/09/2026** (USD/1M tokens, tier standard) :

| Modèle | Entrée /1M | Entrée cachée /1M | Sortie /1M | Notes |
|---|---|---|---|---|
| **GPT-6 Astra** | 10,00 $ | 1,00 $ | 50,00 $ | Flagship du 03/09/2026 ; computer use, code, cybersécurité ; ~1,05M contexte |
| **GPT-5.6 Sol** | 4,00 $ | 0,40 $ | 20,00 $ | Promo jusqu'au 21/11/2026 (tarif normal 5,00/30,00 $) |
| **GPT-5.6 Terra** | 2,00 $ | 0,20 $ | 12,00 $ | Équilibré |
| **GPT-5.6 Luna** | 0,20 $ | 0,02 $ | 1,20 $ | Haut volume, faible latence |
| **GPT-4o** | 2,50 $ | 1,25 $ | 10,00 $ | Encore servi |
| **GPT-4o mini** | 0,15 $ | 0,075 $ | 0,60 $ | Le moins cher du catalogue |
| **o3** | 2,00 $ | 0,50 $ | 8,00 $ | Fin d'API le 11/12/2026 (à vérifier) |
| **text-embedding-3-small** | 0,02 $ | — | — | 1536 dims — **le modèle d'embedding du RAG de Zelef** |
| **text-embedding-3-large** | 0,13 $ | — | — | 3072 dims |

Autres produits : **ChatGPT** (abonnements), **Codex** (agent de code, modèle unifié GPT-5.5/5.6), **Sora** (vidéo ; API fermée le 24/09/2026 — fait vérifié), **API Batch** (-50 % pour jobs différés), **prompt caching** (lecture cache -90 %, écriture +25 % ; durée min 30 min sur 5.6).
Ouverture partielle : **GPT-OSS-20B / 120B** (Apache 2.0, août 2025) — les seuls poids ouverts d'OpenAI, hébergés par presque tous les providers (c'est l'étalon de comparaison des prix inter-providers).

### 2.2. Anthropic (États-Unis, fermé)

Modèles API au **10/09/2026** (USD/1M tokens) :

| Modèle | Entrée /1M | Sortie /1M | Lecture cache /1M | Notes |
|---|---|---|---|---|
| **Claude Fable 5** | 10,00 $ | 50,00 $ | 1,00 $ | Frontière publique (variante de Mythos 5) |
| **Claude Opus 5.5** | 4,00 $ | 20,00 $ | 0,20 $ | Baisse de prix le 24/09/2026 (-20 % in/out, -60 % lecture cache) |
| **Claude Opus 4.8** | 5,00 $ | 25,00 $ | 0,50 $ | 1M contexte |
| **Claude Sonnet 5** | 3,00 $ | 15,00 $ | 0,30 $ | Tarif standard depuis le 01/09/2026 (fini le prix d'appel 2,00 $) |
| **Claude Haiku 4.5** | 1,00 $ | 5,00 $ | 0,10 $ | Haut volume |

Particularités : ratio **sortie = 5× entrée** constant sur toute la gamme ; **prompt caching** très agressif (écriture 1,25× l'entrée en cache 5 min, 2× en cache 1 h ; chaque hit renouvelle la durée) — c'est l'arme n°1 pour un RAG (gros préfixe système réutilisé). **Claude Mythos 5** (mêmes poids que Fable 5 sans les classifieurs de sécurité) : accès restreint (programme Glasswing, US gov) — à vérifier. Produits : **Claude Code** (agent, sièges Pro 20 $/mois, Max 100/200 $/mois), **API**, **Bedrock/Vertex/Foundry** (distribution via hyperscalers).

### 2.3. Google DeepMind (Royaume-Uni/États-Unis)

| Modèle (API) | Entrée /1M | Sortie /1M | Notes (vérifié sept. 2026) |
|---|---|---|---|
| **Gemini 3.1 Pro** | 2,00 $ | 12,00 $ | 1M contexte (4,00/18,00 $ au-delà de 200K) |
| **Gemini 3.1 Flash** | 0,10 $ | 3,00 $ | Meilleur prix/perf haut volume |
| **Gemini 2.5 Flash** | 0,30 $ | 2,50 $ | Dépréciation annoncée 17/06/2026 — migrer vers 3.x |
| **Gemini 2.5 Flash-Lite** | 0,10 $ | 0,40 $ | Déprécié juin 2026 (à vérifier) |
| **Gemma 4** (ouvert) | — | — | Apache 2.0 ; pendant open-weight de Gemini |

Canaux : **AI Studio** (direct, free tier), **Vertex AI** (entreprise), **Gemini app** (abonnements 19,99–29,99 $/mois). Multimodal natif (texte, image, audio, vidéo — Veo pour la vidéo).

### 2.4. xAI (États-Unis)

| Modèle (API) | Entrée /1M | Sortie /1M | Notes |
|---|---|---|---|
| **Grok 4.6** (à vérifier) | ~2,00 $ | ~6,00 $ | Flagship ; ne pas confondre avec **Groq** (l'inférence LPU) |
| **Grok 4.3** | 1,25 $ | 2,50 $ | 1M contexte |
| **Grok 4.1 Fast** | 0,20 $ | 0,50 $ | 2M contexte, haut volume |
| **Grok Build 0.1** | 1,00 $ | 2,00 $ | Orienté code, 256K |

Modèles partiellement ouverts : Grok-1, Grok 2.5 (poids publiés, licence à vérifier). API : `docs.x.ai`. Crédits d'inscription (~25 $) rapportés — à vérifier.

