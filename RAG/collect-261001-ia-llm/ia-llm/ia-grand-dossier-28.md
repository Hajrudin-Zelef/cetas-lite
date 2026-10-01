---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-28
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "Z.ai", "vLLM", "xAI"]
dates: []
keywords: ["agents", "awq", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "cohere", "compute", "datacenter", "deepseek", "diffusion"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1890, 1958]
sha256: a59bdfc152162b85de9522edec2b85e7bd96ce7f434460d50dcf4e89833c3da2
---

# IA — Le grand dossier

- Révolutions : **1.1** (2012, pourquoi) → **1.2** (AlexNet) → **1.3** (jeux/RL) → **1.4** (GAN) → **1.5** (Transformer) → **1.6** (2018, GPT-1/BERT) → **1.7** (GPT-3) → **1.8** (diffusion) → **1.9** (RLHF/ChatGPT) → **1.10** (2023, course LLM) → **1.11** (2024, multimodal/agents/o1) → **1.12** (2025, R1) → **1.13** (2026) → **1.14** (chaînons manquants) → **1.15** (patterns historiques) → **1.16** (multimodalité)
- Personnes : **2.1** (parrains) → **2.2** (équipe Transformer) → **2.3** (fondateurs) → **2.4** (chercheurs) → **2.5** (les deux camps) → **2.6** (où sont-ils)
- Labos : **3.1** (tableau) → **3.2** (fiches détaillées) → **3.3** (dynamiques) → **3.4** (acteurs secondaires)
- Scaling et débat : **4.1** (lois de Kaplan/Chinchilla) → **4.2** (camp du mur) → **4.3** (camp de la continuation) → **4.4** (efficacité : MoE, distillation) → **4.5** (test-time compute) → **4.6** (murs physiques) → **4.7** (synthèse) → **4.8** (les maths) → **4.9** (scénarios 2027-2030) → **4.10** (10 questions ouvertes)
- Timeline : **5** (2012 → 2026, année par année)
- Sysadmin : **6.1** (l'inférence > le training) → **6.2** (hardware et VRAM) → **6.3** (API vs self-hosting) → **6.4** (stack de serving) → **6.5** (compétences) → **6.6** (risques) → **6.8** (sécurité opérationnelle) → **6.9** (feuille de route RAG) → **6.10** (datacenter IA et énergie)
- Méthode : **7** (sources et vérification)
- Compléments : **8** (25 papiers) → **9** (chiffres du compute) → **10** (hardware et infrastructure) → **11** (15 idées reçues) → **12** (10 portraits) → **13** (voix du débat) → **14** (5 cas sysadmin chiffrés) → **15** (20 benchmarks) → **16** (datasets) → **17** (lire un rapport technique) → **18** (frise des modèles) → **19** (FAQ, 20 questions) → **20** (géopolitique) → **21** (40 acronymes) → **22** (checklist annonces) → **23** (12 fiches modèles 2026) → **24** (10 erreurs RAG)
- **25**. Glossaire (60 termes) — **26**. Quiz (10 questions + corrigés, + 5 bonus)

## Comment utiliser ce dossier (mode d'emploi)

- **Première lecture :** sections 1 (révolutions), 4 (débat) et 6 (sysadmin) — l'ossature complète en ~400 lignes.
- **Avant un choix technique :** sections 9 (chiffres), 14 (cas chiffrés), 23 (fiches modèles), 24 (erreurs RAG).
- **Avant une réunion :** la fiche de révision express (40 points) + la checklist annonces (22).
- **Veille :** les « à vérifier » sont des paris datés — les revérifier chaque trimestre, en commençant par les sections 4.9 et 4.10.
- **Contribution :** ce dossier est un document vivant — chaque nouveau rapport technique important mérite sa fiche en section 23.

---


---

# PARTIE 2 — Providers, passerelles/routage, IA locale

> Grand dossier IA de Zelef — Partie 2/3.
> **Sujet** : qui vend l'IA (providers, hyperscalers, labs), comment router entre eux (passerelles, gateways, LiteLLM, OpenRouter), et comment s'en passer (IA locale : modèles ouverts, stacks, hardware chiffré).
> **Périmètre temporel** : faits, prix et versions vérifiés par recherche web au **27 septembre 2026**. Tout prix est en USD sauf mention contraire ; les tarifs API bougent vite — chaque tableau de prix porte sa date de vérification. Les faits incertains sont marqués **« à vérifier »**.
> **Règle d'hygiène** : aucune clé d'API réelle dans ce document. Les exemples de config utilisent des variables d'environnement (`${...}`) — jamais de secret en clair.

---

## Sommaire

1. [Les grands providers : hyperscalers](#1-les-grands-providers--hyperscalers) — AWS, Azure, Google Cloud, Oracle : qui vend quoi
2. [Les labs : qui crée les modèles](#2-les-labs--qui-crée-les-modèles) — OpenAI, Anthropic, Google DeepMind, xAI, Mistral, DeepSeek, Alibaba, Moonshot, Zhipu et les autres
3. [Top 50 des fournisseurs IA vérifiés](#3-top-50-des-fournisseurs-ia-vérifiés) — APIs LLM, clouds GPU, inférence serverless, gateways
4. [Les passerelles et le routage intelligent](#4-les-passerelles-et-le-routage-intelligent) — OpenRouter, LiteLLM, stratégies coût/qualité/latence, fallbacks, observabilité, configs réutilisables
5. [L'IA locale en profondeur](#5-lia-locale-en-profondeur) — pourquoi, modèles ouverts et licences, formats (GGUF, AWQ, GPTQ), stacks (Ollama, llama.cpp, vLLM, LM Studio)
6. [Hardware chiffré : dimensionner une machine locale](#6-hardware-chiffré--dimensionner-une-machine-locale) — tableaux VRAM, Mac vs NVIDIA vs CPU, débits indicatifs
7. [Cas d'usage professionnels concrets](#7-cas-dusage-professionnels-concrets) — RAG local, classification, code, supervision
8. [Comparatif local vs cloud : le tableau de décision](#8-comparatif-local-vs-cloud--le-tableau-de-décision)
9. [Glossaire (40 termes)](#9-glossaire-40-termes)
10. [Quiz — 10 questions corrigées](#10-quiz--10-questions-corrigées)

---

## 1. Les grands providers : hyperscalers

Un **hyperscaler** est un opérateur de cloud à l'échelle planétaire (dizaines de régions, centaines de milliers de serveurs) qui vend désormais l'IA sous trois formes : **(a)** l'accès API à des modèles (les leurs + ceux des labs partenaires), **(b)** la location de GPU à l'heure (instances), **(c)** des puces maison pour l'inférence (moins chères que NVIDIA). Pour un sysadmin, c'est le même réflexe que pour le reste du cloud : on choisit selon la **facturation** (à la requête, au token, à l'heure GPU), la **souveraineté** (région, certifications) et le **verrouillage** (SDK propriétaire vs API compatible OpenAI).

### 1.1. AWS (Amazon Web Services)

**Ce qu'ils vendent en IA :**

| Offre | Description | Modèle économique |
|---|---|---|
| **Amazon Bedrock** | API managée multi-modèles : Claude (Anthropic), Llama (Meta), Mistral, Amazon Nova, DeepSeek (selon régions), Cohere… Un seul endpoint, facturation unifiée | $/1M tokens, tarifs variant par modèle et région |
| **Amazon SageMaker AI** | Plateforme ML complète : entraînement, déploiement d'endpoints, hosting de modèles open-weight (JumpStart) | $/heure d'instance + invocations |
| **EC2 GPU** | Location brute de GPU : P5 (H100), P6 (B200/B300), G6e (L40S), G5 (A10G) ; spot disponible | $/heure/GPU ; ex. H100 ~5,16–6,88 $/h/GPU on-demand (juil.–sept. 2026, à vérifier selon région) |
| **Trainium / Inferentia** | Puces maison : entraînement (Trainium2) et inférence (Inferentia2) ~20–40 % moins chères que NVIDIA à perf équivalente sur charges stables | $/heure d'instance (familles Trn1/Inf2) |
| **Amazon Nova** | Famille de modèles maison (Micro, Lite, Pro, Premier) servie via Bedrock | Ex. Nova Micro : 0,04 $/1M in / 0,14 $/1M out (vérifié sept. 2026) |

**Positionnement** : le choix par défaut des DSI déjà sur AWS. Bedrock est le « supermarché » : on y trouve Claude sans passer par Anthropic, avec la facturation AWS, les IAM, le VPC et les certifications (SOC 2, HIPAA, ISO 27001). **Point de vigilance** : les frais d'**egress** (0,09 $/Go sortant) — télécharger un checkpoint 70B (~140 Go en FP16) coûte ~12,60 $ à chaque fois.

### 1.2. Microsoft Azure

**Ce qu'ils vendent en IA :**

