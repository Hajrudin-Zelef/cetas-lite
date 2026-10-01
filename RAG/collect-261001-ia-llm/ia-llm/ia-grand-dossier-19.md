---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-19
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Apple", "Cohere", "DeepSeek", "Google", "Hugging Face", "Inflection AI", "Microsoft", "MiniMax", "Mistral", "OpenAI", "Perplexity", "Stability AI", "StepFun", "United States", "Z.ai", "xAI"]
dates: []
keywords: ["agents", "apache", "attention", "benchmarks", "cohere", "compute", "deepseek", "diffusion", "dpo", "embeddings", "glm", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1235, 1280]
sha256: ef0bb5a45219ec3630a0c55fa93da6bf59aa4b700f52702116f19beec05e914c
---

# IA — Le grand dossier

**Cohere** (Toronto, 2020 — Aidan Gomez, ex-« Attention Is All You Need »). Positionnement **entreprise** : modèles Command, **RAG** (Rerank), embeddings multilingues. Le labo « boring but useful » : moins de hype, beaucoup d'adoption B2B. Pertinent pour un RAG : leur **reranker** est une référence.

**Hugging Face** (2006 → pivot IA 2016). Pas un labo de modèles frontières, mais **l'infrastructure de l'open source** : hub de modèles/datasets, Transformers (la bibliothèque), Spaces, Inference Endpoints. Sans HF, l'écosystème Llama/Qwen/Mistral n'existerait pas à cette vitesse. Valorisation ~4,5 Md$ (2023 — « à vérifier » l'évolution).

**Stability AI** (2022, Emad Mostaque). A financé et popularisé **Stable Diffusion** (août 2022) — la démocratisation de l'image. Turbulences 2023-2024 (départ du fondateur, difficultés financières — « à vérifier » le détail), mais l'impact historique reste.

**Character.AI** (2021, Noam Shazeer — ex-« Attention Is All You Need » — et Daniel De Freitas). Chatbots de personnages, succès grand public fulgurant 2023. **2024 : Shazeer retourne chez Google** dans un accord à ~2,7 Md$ (montant de presse — « à vérifier ») : cas d'école du « acqui-hire déguisé » par les hyperscalers.

**Perplexity** (2022, Aravind Srinivas). Le **moteur de recherche conversationnel** : LLM + retrieval web + citations. Modèle économique : l'abonnement et la pub ; modèle technique : c'est un **RAG sur le web entier** — la preuve produit que l'architecture RAG passe à l'échelle planétaire.

**Inflection AI** (2022, Mustafa Suleyman + Reid Hoffman). Chatbot **Pi** (l'empathique). **2024 : Microsoft absorbe l'équipe** (Suleyman devient CEO de Microsoft AI) — autre cas de consolidation.

**Adept** (2022, ex-« Attention Is All You Need » : Vaswani et al.). Agents qui manipulent les logiciels. **2024 : Amazon absorbe une partie de l'équipe** (« à vérifier » le périmètre) — troisième cas de consolidation la même année.

**Le motif 2024 est clair :** les labos indépendants sans hyperscaler partenaire se font **absorber pour leurs talents**. Seuls survivent indépendants : ceux qui ont un **financement propre** (DeepSeek via High-Flyer), une **cause nationale** (Mistral via la France/l'UE), ou une **échelle déjà atteinte** (Anthropic, xAI).

**Zhipu AI (GLM)** (Pékin, 2019 — issue de Tsinghua). Famille **GLM-4/5**, très forte en chinois/anglais, stratégie open-weight agressive. Avec Qwen (Alibaba), MiniMax et StepFun, elle forme le **deuxième écosystème** : la Chine ne dépend plus des modèles US pour l'usage courant.

**Apple** — le géant silencieux. Pas de labo frontière public, mais : puces **Neural Engine** dans 2+ milliards d'appareils, **Apple Intelligence** (2024 — approche « petit modèle local + cloud privé »), et des publications (OpenELM, DCLIP — « à vérifier » les dernières). Stratégie : **l'IA on-device** comme différenciation vie privée — l'exact inverse du cloud géant.

---

## 17. Lire un rapport technique de modèle comme un pro (guide)

> Les labos publient des « technical reports » (GPT-4, Llama 3, DeepSeek-V3...). Voici comment les lire en 30 minutes pour en extraire l'essentiel.

**Checklist de lecture :**

1. **Architecture (section 2-3).** Dense ou MoE ? Combien d'actifs par token ? Attention classique ou variante (MLA, GQA, sliding window) ? → Détermine le **coût d'inférence**.
2. **Données (data section).** Combien de tokens ? Quelles sources ? Ratio tokens/paramètre ? → Positionne le modèle sur l'axe Chinchilla (sous/sur-entraîné).
3. **Compute.** FLOP totaux ? Type et nombre de GPU ? Durée ? → Permet de vérifier les affirmations de coût (formule C ≈ 6ND).
4. **Post-training.** SFT ? RLHF/DPO/RLVR ? Quelle part de RL ? → Un modèle « base » ≠ un modèle « instruct » : le post-training fait 50 % de l'utilité perçue.
5. **Évaluations.** Quels benchmarks ? Comparé à quoi ? **Chercher les benchmarks où il perd** — un rapport honnête en montre. Se méfier des « evals maison » non reproductibles.
6. **Limites (souvent à la fin).** C'est la section la plus honnête du rapport : hallucinations, biais, langues faibles, long contexte dégradé.
7. **Licence et accès.** Poids ouverts ? Quelle licence (Apache 2.0, MIT, communautaire) ? → Détermine si vous pouvez le déployer (voir idée reçue n° 12).

**Signaux d'alerte :** aucun détail d'architecture ni de données (GPT-4, 2023 — rupture de la norme académique) ; benchmarks uniquement internes ; comparaisons contre des versions affaiblies des concurrents ; absence de section limites.

**Exercice :** prenez le rapport **DeepSeek-V3** (déc. 2024, arXiv:2412.19437) et le rapport **GPT-4** (mars 2023) : comparez ce que chacun révèle. Vous comprendrez en une heure pourquoi la transparence est devenue un **argument géopolitique** (voir 3.2, DeepSeek).

---

## 18. Chronologie des modèles : tous les jalons, 2018 → 2026

> Un modèle par ligne (ou presque) : la frise que tout praticien devrait avoir en tête. « Poids » = ouverts ou fermés.

