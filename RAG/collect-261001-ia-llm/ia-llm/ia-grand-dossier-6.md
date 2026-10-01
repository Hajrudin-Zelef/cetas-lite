---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-6
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Apple", "Baidu", "DeepSeek", "Google", "Inflection AI", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Stability AI", "StepFun", "United States", "Z.ai", "xAI"]
dates: ["2026-09-27"]
keywords: ["agi", "apache", "blackwell", "chatgpt", "claude", "compute", "copilot", "deepseek", "diffusion", "distribution", "gemini", "glm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [326, 398]
sha256: cce8ac83bb718986496df29612cbb2935eb0fccffe11f415e7742c4f3ce1e3d7
---

# IA — Le grand dossier

**Jensen Huang (né 1963).**
- Cofondateur et CEO de **NVIDIA** (1993). Le pari CUDA (2006) — rendre les GPU programmables pour le calcul général — est rétrospectivement **le pari infrastructurel le plus rentable de l'histoire de la tech** : sans CUDA, pas de deep learning à grande échelle.
- NVIDIA devient en 2024-2025 l'entreprise la plus valorisée du monde par périodes, avec ~80-90 % du marché des accélérateurs IA (« à vérifier » la part exacte, l'ordre de grandeur est consensuel).
- Jalons hardware : A100 (2020), H100 (2022), **Blackwell B200** (2024), Rubin (annoncé).

**Liang Wenfeng.**
- Fondateur du fonds quantitatif **High-Flyer** puis de **DeepSeek** (2023, Hangzhou). Particularité : DeepSeek est financé par les profits du trading quantitatif, pas par du VC — d'où une liberté de recherche rare.
- Stratégie : **efficacité algorithmique sous contrainte** (GPU H800 bridés par les contrôles d'export US vers la Chine) + **poids ouverts**. Le choc R1 de janvier 2025 (voir 1.12) l'a rendu célèbre.

**Arthur Mensch (né 1992).**
- Ex-DeepMind (contributeur à **Chinchilla**, Flamingo, Retro), normalien, docteur.
- **Avril 2023** : cofonde **Mistral AI** à Paris avec Guillaume Lample et Timothée Lacroix (ex-Meta, équipe LLaMA).
- Levée seed record européenne : **105 M€** (juin 2023). Positionnement : **open-weight + efficacité + souveraineté européenne**. Série D de **3 Md€** en septembre 2026 (valorisation > 21 Md€ — chiffres de presse, septembre 2026).

**Mustafa Suleyman.**
- Cofondateur de DeepMind (2010), puis d'**Inflection AI** (2022, avec Reid Hoffman) et du chatbot Pi.
- **Mars 2024** : rejoint **Microsoft** comme CEO de **Microsoft AI** (Copilot, etc.). Illustration de la consolidation : les talents circulent vers les hyperscalers.

### 2.4. Les chercheurs-clés à connaître (au-delà des fondateurs)

| Personne | Contribution vérifiée | Période |
|---|---|---|
| **Fei-Fei Li** | Crée **ImageNet** (2009, avec ses étudiants) — sans elle, pas d'AlexNet ; vulgarisatrice majeure | 2009→ |
| **Ian Goodfellow** | Invente les **GAN** (2014, doctorant de Bengio) ; ex-Google/Apple | 2014 |
| **Alex Krizhevsky** | Premier auteur d'**AlexNet** (2012) | 2012 |
| **Andrej Karpathy** | Membre fondateur d'OpenAI (2015), directeur IA de Tesla (2017-2022, Autopilot vision), retour OpenAI 2023-2024, vulgarisateur (« Software 2.0 », Zero-to-Hero), rejoint **Anthropic** (pretraining) en mai 2026 | 2015→ |
| **Jared Kaplan** | Premier auteur des **lois de scaling** OpenAI (2020) ; parti fonder Anthropic | 2020 |
| **Jordan Hoffmann** | Premier auteur du papier **Chinchilla** (DeepMind, 2022) qui corrige Kaplan | 2022 |
| **John Schulman** | Cofondateur OpenAI, inventeur de **PPO** (2017) — l'algorithme du RLHF ; rejoint Anthropic en août 2024 (alignement) | 2017→ |
| **Jeff Dean & Sanjay Ghemawat** | Pionniers des systèmes distribués Google (MapReduce), cofondateurs de **Google Brain** (2011) | 2011→ |
| **Andrew Ng** | Google Brain, cofondateur de Coursera (2012), ex-Baidu ; vulgarisateur n° 1 du ML | 2011→ |
| **Sepp Hochreiter & Jürgen Schmidhuber** | Inventeurs du **LSTM** (1997) — a dominé le NLP avant le Transformer | 1997 |
| **Jonathan Ho** | Premier auteur de **DDPM** (2020) — la diffusion moderne | 2020 |
| **Robin Rombach** | Pilote **Stable Diffusion** (LMU Munich → Stability AI, 2022) | 2022 |
| **Noam Brown** | Libratus (poker, 2017, CMU), puis co-pilote du projet **o1/Strawberry** chez OpenAI | 2017→ |
| **Jason Wei** | Co-auteur du papier **Chain-of-Thought** (2022) et figure du raisonnement chez OpenAI/Google | 2022→ |
| **Chris Olah** | Interprétabilité (OpenAI → Anthropic), circuits des réseaux de neurones | 2017→ |
| **Yann LeCun** | Déjà cité en 2.1 (rappel : il appartient aux deux catégories) | — |

### 2.5. Carte des désaccords : qui pense quoi du futur

| Position | Représentants | Thèse |
|---|---|---|
| **Maximalistes du scaling** | Sam Altman, Dario Amodei (version « sûre »), ex-Gwern | Le paradigme actuel (Transformer + RL + scale) mène à l'AGI ; il faut surtout plus de compute |
| **Prudents / alarmistes** | Geoffrey Hinton, Yoshua Bengio, Ilya Sutskever (via SSI) | Le scaling marche mais devient dangereux ; régulation et recherche sûreté prioritaires |
| **Sceptiques du LLM** | Yann LeCun, Andrew Ng (nuancé) | Les LLM seuls ne feront pas l'AGI ; il faut world models / nouvelles architectures ; le risque existentiel est surjoué |
| **Efficacité d'abord** | Liang Wenfeng, Arthur Mensch | La valeur est dans l'efficacité algo et l'ouverture, pas dans la taille brute |
| **Mur des données** | Elon Musk, Epoch AI (Villalobos et al.) | Le texte humain de qualité s'épuise (2026-2032) ; le scaling pré-training va caler |

Ces camps ne sont pas des caricatures : beaucoup d'acteurs appartiennent à plusieurs cases selon le sujet. Mais cette carte est indispensable pour lire l'actualité IA sans se perdre.

---

## 3. Les grands laboratoires : positionnement de chacun

### 3.1. Tableau comparatif synthétique

| Labo | Fondation / contrôle | Modèles phares (au 27/09/2026) | Ouverture | Positionnement stratégique |
|---|---|---|---|---|
| **OpenAI** | 2015 (non-profit) → 2019 (capped-profit) ; Microsoft partenaire majeur (~13 Md$ au total) | GPT-4o, o1/o3, GPT-5 (août 2025) | Fermé (API) | Le maximaliste : AGI par le scaling, intégration verticale produit (ChatGPT, API, Sora) |
| **Anthropic** | 2021 (ex-OpenAI) ; soutiens Google + Amazon (plusieurs Md$) | Claude (familles 3, 4, Opus 4.7 — « à vérifier » les n° exacts au jour près) | Fermé (API) | « Scale safe » : IA constitutionnelle, RSP, positionnement entreprise/confiance |
| **Google DeepMind** | DeepMind 2010 + Google Brain → fusion avril 2023 | Gemini (2.0, 3.x), Veo, AlphaFold | Fermé (API) + recherche publiée | L'intégré vertical : TPU maison, YouTube/Search/Android comme données et distribution |
| **Meta (FAIR / GenAI)** | FAIR fondé 2013 par LeCun | Llama (2, 3, 4), Llama 3 405B | **Poids ouverts** (licence communautaire) | L'ouverture comme arme : commoditiser les modèles pour vendre pubs + hardware (lunettes, etc.) |
| **xAI** | 2023 (Elon Musk) | Grok (intégré à X) | Partiellement ouvert | Compute massif (Colossus), distribution via X, ton « anti-woke » marketing |
| **DeepSeek** | 2023 (Liang Wenfeng, fonds High-Flyer) | V3, R1, V4 (2026) | **Poids ouverts (MIT)** | Efficacité algorithmique sous contrainte d'export ; casse les prix |
| **Alibaba (Qwen)** | Tongyi Lab | Qwen (2.5, 3), Qwen-Coder | Poids ouverts (Apache 2.0) | Le champion open chinois : écosystème dev, modèles code/multilingues |
| **Moonshot AI** | 2023 (Yang Zhilin) | Kimi (k1.5, K2 — « à vérifier » les versions) | Partiellement ouvert | Contexte long extrême (pionnier du 1M+ tokens côté chinois) |
| **Mistral AI** | Avril 2023 (Paris) | Mistral Large 3, Mixtral, Magistral (raisonnement), Le Chat | **Apache 2.0** (open-weight) | Souveraineté européenne, efficacité MoE, B2B/gouvernements |
| **NVIDIA** | 1993 (Jensen Huang) | (pas de LLM frontaux ; NeMo, NIM) | Écosystème CUDA | Le « vendeur de pelles » : ~80-90 % des accélérateurs IA (« à vérifier » la part exacte) |
| **Autres chinois** | Zhipu (GLM), MiniMax, StepFun, Baidu (Ernie) | GLM-4/5, MiniMax M2, Ernie | Mixte | Écosystème parallèle sous contraintes d'export US |

### 3.2. Fiches détaillées

