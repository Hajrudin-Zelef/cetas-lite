---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-23
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Baidu", "Cohere", "DeepSeek", "Google", "Huawei", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "StepFun", "TSMC", "United States", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "apache", "ascend", "compute", "deepseek", "embeddings", "export controls", "gemini", "glm", "hyperscaler", "kimi"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1528, 1584]
sha256: e80fbb57aee5cac7b0c7448204765c6aac80bac6416984714148b16f4db67d69
---

# IA — Le grand dossier

**Frise multimodale :**
- **2021 — CLIP** (OpenAI) : l'alignement texte-image qui rend tout le reste possible.
- **2022 — Flamingo** (DeepMind) : few-shot visuel — montrer 2-3 exemples image+texte suffit.
- **2023 — GPT-4V** : la vision arrive dans le modèle frontière (lecture de documents, schémas, photos d'équipements).
- **2024 — GPT-4o** (mai) : texte+image+audio **dans un seul réseau**, latence ~320 ms — la conversation vocale naturelle. **Gemini 1.5** : 1M tokens de contexte → des heures de vidéo ingérables d'un coup.
- **2024 — Sora** (fév.) : vidéo 1 minute ; **Veo 3** (2025) : vidéo + audio natif.
- **2025-2026 — les agents « voient »** : computer use (captures d'écran), robotique (RT-2, puis les modèles « physical AI » 2026 : Robostral de Mistral, etc. — « à vérifier » les noms au jour près).

**Pour le RAG de Zelef, c'est stratégique :** ses documents techniques contiennent **schémas unifilaires, photos d'équipements, tableaux, captures** — pas que du texte. Un pipeline RAG 2026 doit donc :
1. **Extraire** le texte (pdftotext/OCR),
2. **Décrire** les images (un modèle vision génère une légende indexable),
3. **Indexer** les deux dans le même espace vectoriel (embeddings multimodaux : CLIP, SigLIP),
4. **Restituer** l'image source dans la réponse (pas seulement le texte).

Le RAG « texte seul » est déjà une architecture legacy pour la documentation technique.

---

## 20. Géopolitique de l'IA : les trois blocs (situation sept. 2026)

> L'IA est devenue une ressource stratégique. Cette annexe résume les positions — faits vérifiés, analyses attribuées.

### 20.1. États-Unis : le maximalisme privé sous perfusion publique

- **Acteurs :** OpenAI, Anthropic, Google DeepMind, Meta, xAI, NVIDIA — tous privés, tous dépendants du **soutien public indirect** (contrats fédéraux, énergie, export controls comme arme).
- **Doctrine :** « l'innovation d'abord » — régulation fédérale minimale (les tentatives type SB 1047 en Californie ont échoué ou été édulcorées — « à vérifier » l'état exact), **contrôles d'export** comme instrument géopolitique (H100/H800 → Chine, 2022-2023, renforcés).
- ** paris :** Stargate (~500 Md$ annoncés, janv. 2025), clusters privés gigantesques (Colossus/xAI).
- **Faiblesse :** dépendance à **TSMC** (Taïwan) pour la gravure — le point de vulnérabilité n° 1 reconnu par tous les stratèges US.

### 20.2. Chine : l'efficacité contrainte comme doctrine

- **Acteurs :** DeepSeek, Alibaba (Qwen), Moonshot (Kimi), Zhipu (GLM), MiniMax, StepFun, Baidu (Ernie), Huawei (Ascend).
- **Doctrine :** contourner les sanctions par **l'algorithmique** (DeepSeek est la preuve de concept), puis **l'ouverture** (poids MIT/Apache) pour capter l'écosystème mondial des développeurs.
- **Atouts :** ingénieurs nombreux et excellents, données industrielles massives, volonté étatique (plans IA successifs depuis 2017).
- **Faiblesses :** accès limité au compute de pointe (H800 bridés, Ascend en rattrapage — « à vérifier » les perfs relatives), censure des modèles domestiques (alignement politique obligatoire — fait documenté).
- **Leçon pour le dossier :** les contrôles d'export ont **accéléré** l'adversaire sur l'axe efficacité — l'effet inverse de l'objectif affiché (analyse partagée par de nombreux commentateurs, dont Thompson 2025).

### 20.3. Europe : la régulation comme produit (et ses limites)

- **Acteurs :** Mistral AI (France), Aleph Alpha (Allemagne — « à vérifier » la trajectoire 2025-2026), pools de recherche (Ellis, CNRS...).
- **Doctrine :** **AI Act** (entré en application progressive 2024-2026) : la première régulation complète au monde, par niveaux de risque. Souveraineté : investissements (Mistral Compute, datacenters — annonces 2025-2026).
- **Atouts :** recherche fondamentale excellente, Mistral comme champion crédible, cadre juridique clair pour les entreprises.
- **Faiblesses :** **pas d'hyperscaler européen** (le cloud IA reste US), compute public insuffisant, fuite des talents vers les salaires US.
- **Positionnement de Mistral :** l'« européen » qui joue le jeu de l'ouvert (Apache 2.0) pour exister face aux géants — stratégie validée par les levées (3 Md€, sept. 2026).

### 20.4. Ce que ça change pour une entreprise (hors politique)

1. **Dépendance :** 100 % des modèles frontières viennent des US ou de Chine — **l'abstraction multi-fournisseur** (6.3) n'est pas un luxe, c'est de la gestion du risque géopolitique.
2. **Conformité :** l'AI Act impose des obligations (transparence, documentation, évaluation) croissantes avec le risque du cas d'usage — un RAG interne documentaire est en **risque faible**, un agent qui décide en production est en **risque élevé**.
3. **Données :** les restrictions de scraping (4.2) + les procès en cours rendent les corpus « aspirés » juridiquement fragiles — **vos données propres** (méthode Zelef) sont l'actif le plus sûr.

---

## 21. Petit dictionnaire des acronymes (40 entrées)

> Pour décoder n'importe quel article, papier ou spec sans ouvrir 10 onglets.

