---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-21
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Baidu", "Cohere", "DeepSeek", "Google", "Inflection AI", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Sakana", "United States", "xAI"]
dates: ["2025-01-27"]
keywords: ["agent", "agi", "apache", "attention", "cohere", "compute", "deepseek", "distillation", "embeddings", "gpu", "grok", "kimi"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1362, 1448]
sha256: 5e9aeba31260c239a0e9810bde0cfc3e8384275178626da0dc73b3a639e8a3db
---

# IA — Le grand dossier

| Personne | Était (fait marquant) | Est (sept. 2026, vérifié) |
|---|---|---|
| Geoffrey Hinton | AlexNet 2012, Google 2013-2023 | Alerteur public sur les risques ; Nobel physique 2024 |
| Yann LeCun | CNN 1989, FAIR 2013 | Chief AI Scientist Meta ; sceptique n° 1 du scaling LLM |
| Yoshua Bengio | Embeddings, Mila | Mila ; rapport international sur la sûreté de l'IA |
| Ilya Sutskever | AlexNet, seq2seq, OpenAI chief scientist | Fondateur de **Safe Superintelligence Inc.** (juin 2024) |
| Sam Altman | YC, OpenAI 2015 | CEO d'OpenAI (réintégré nov. 2023) |
| Dario Amodei | VP Recherche OpenAI | CEO d'Anthropic (fondée 2021) |
| Demis Hassabis | DeepMind 2010, AlphaGo, AlphaFold | CEO de **Google DeepMind** ; Nobel chimie 2024 |
| Elon Musk | Cofondateur OpenAI 2015, départ 2018 | Fondateur de **xAI** (2023), Grok, Colossus |
| Andrej Karpathy | OpenAI 2015, Tesla 2017-2022 | **Anthropic** (pretraining, depuis mai 2026) ; ex-Eureka Labs |
| Fei-Fei Li | ImageNet 2009 | Stanford HAI ; fondatrice de **World Labs** (2024) |
| Ian Goodfellow | GAN 2014 | Recherche (ex-Google, ex-Apple) |
| Arthur Mensch | DeepMind (Chinchilla) | CEO de **Mistral AI** (fondée avril 2023) |
| Liang Wenfeng | High-Flyer (quant) | Fondateur de **DeepSeek** (2023) |
| Mustafa Suleyman | DeepMind 2010, Inflection 2022 | CEO de **Microsoft AI** (depuis mars 2024) |
| Jensen Huang | NVIDIA 1993, CUDA 2006 | CEO de NVIDIA (entreprise la plus valorisée par périodes 2024-2025) |
| Ashish Vaswani | 1er auteur « Attention Is All You Need » | Cofondateur d'**Adept** (2021), ex-Google |
| Aidan Gomez | Co-auteur Transformer | Cofondateur/CEO de **Cohere** (2020) |
| Llion Jones | Co-auteur Transformer | Cofondateur de **Sakana AI** (Tokyo, 2023) |
| Noam Shazeer | Co-auteur Transformer, Google 2000-2023 | Retour chez **Google** (2024, accord ~2,7 Md$ — « à vérifier ») |
| John Schulman | PPO 2017, OpenAI | Anthropic (depuis août 2024, alignement) |
| Jeff Dean | Google Brain 2011 | Dirige la recherche de **Google DeepMind** |
| Andrew Ng | Google Brain, Baidu, Coursera | DeepLearning.AI, Landing AI ; vulgarisation |
| Yang Zhilin | Google Brain, Tsinghua | Fondateur de **Moonshot AI** (Kimi, 2023) |

---

## 19. FAQ : 20 questions que tout le monde pose (réponses courtes et sourcées)

**19.1. C'est quoi, la différence entre IA, ML et deep learning ?**
L'IA = le domaine (machines qui font des tâches intelligentes). Le ML = l'IA par apprentissage sur données (vs règles codées). Le deep learning = le ML par réseaux de neurones profonds. Depuis 2012, « IA » désigne en pratique le deep learning.

**19.2. Pourquoi 2012 et pas avant ?**
Les idées (rétropropagation 1986, CNN 1989) existaient, mais il manquait les **données** (ImageNet 2009) et le **compute** (GPU CUDA). AlexNet est la rencontre des trois.

**19.3. Un LLM « comprend »-il vraiment ?**
Débat ouvert (voir 2.5). Position prudente : il **modélise statistiquement** le langage à un niveau qui produit des comportements indiscernables de la compréhension sur beaucoup de tâches, avec des échecs caractéristiques (hallucinations, raisonnement fragile sans CoT). Pour un usage pro : **vérifier, ne pas croire**.

**19.4. C'est quoi, un token, concrètement ?**
Un morceau de mot (~4 caractères / ~0,75 mot anglais). « Onduleur » = ~2-3 tokens. C'est l'unité de découpage du modèle, de mesure du contexte et de facturation des API.

**19.5. Pourquoi les modèles « hallucinent » ?**
Parce qu'ils sont entraînés à produire du texte **plausible**, pas du texte **vrai**. La vérité n'est pas dans la fonction objectif (prédire le mot suivant). Contre-feux : RAG (ancrage documentaire), vérificateurs, modèles de raisonnement.

**19.6. Open source ou API pour mon entreprise ?**
Voir 6.3 : API si < ~50-100M tokens/mois ou besoin de la frontière ; self-hosting si volume, confidentialité ou souveraineté. Dans tous les cas : **interface abstraite** (API compatible OpenAI) pour pouvoir changer.

**19.7. Quelle taille de modèle pour un RAG ?**
8B-32B quantifié suffit dans 95 % des cas **si le retrieval est bon**. La qualité du RAG vient du chunking, des embeddings et du reranking — pas des paramètres du générateur. Mesurer avec un gold set (14.3).

**19.8. C'est quoi, Chinchilla, en une phrase ?**
La preuve (DeepMind, 2022) qu'à budget égal, mieux vaut un **petit modèle très entraîné** (20 tokens/paramètre) qu'un géant sous-entraîné — ce qui a tué la course à la taille brute.

**19.9. Pourquoi DeepSeek a fait chuter NVIDIA en bourse ?**
Parce que R1 (janv. 2025) a montré un niveau **o1 pour une fraction du coût**, ce qui a ébranlé la thèse « la valeur de l'IA = le volume de GPU chers ». (-17 %, ~593 Md$ le 27/01/2025 — record US à l'époque.)

**19.10. Le « raisonnement » des modèles, c'est du vrai raisonnement ?**
C'est un **processus de génération de traces intermédiaires** optimisé par RL pour maximiser les bonnes réponses. Ça échoue encore (overthinking, erreurs logiques). Mais les scores (AIME 13 % → 96,7 % en un an) montrent que c'est plus qu'un tour de passe-passe.

**19.11. On va manquer de données ?**
De **texte humain de qualité** : oui, à horizon 2026-2032 au rythme actuel (Epoch AI). Mais : données synthétiques filtrées, RL sans données humaines, multimodal/capteurs. Le mur concerne le pré-entraînement texte naïf.

**19.12. L'AGI, c'est pour quand ?**
Personne ne sait. Fourchettes publiques : Amodei évoque 2026-2027 pour des capacités « prix Nobel » (prédiction) ; d'autres disent 2030+ ; LeCun dit « pas avec les LLM ». Traiter toute date comme une **opinion**, pas un fait.

**19.13. Faut-il avoir peur de l'IA ?**
Deux risques distincts : le **risque existentiel** (débat Hinton/Bengio vs LeCun — non tranché) et les **risques concrets immédiats** (désinformation, deepfakes, concentration, biais, chômage de certaines tâches — avérés). Pour un pro : gérer le concret d'abord.

**19.14. C'est quoi, un agent IA ?**
Un LLM + des outils (web, code, API) + une boucle qui planifie, agit et observe. 2024 : les premiers crédibles (computer use, deep research). 2025-2026 : le principal vecteur de valeur — et de coûts (voir 14.2).

**19.15. Pourquoi tout le monde parle d'inférence en 2026 ?**
Parce que le centre de gravité économique a basculé : ~80 % du compute irait à l'inférence fin 2026 (débat « 80/20 reversal »). Le training ne rapporte plus assez par dollar ; le serving, si.

**19.16. GPU ou TPU ou autre ?**
GPU NVIDIA = le standard flexible (CUDA). TPU = excellent mais captif Google. En 2026, la bataille se joue sur le **silicium d'inférence** (efficacité/watt), pas sur le training brut. Pour un déploiement PME : une carte gamer/ prosumer suffit.

**19.17. C'est quoi, la distillation ?**
Faire apprendre à un petit modèle en imitant un grand (y compris ses raisonnements). C'est comme ça qu'un 14B local fait du « o1-class ». Le levier n° 1 du déploiement économique.

**19.18. Les modèles sont-ils biaisés ?**
Oui, structurellement : ils reflètent leurs données d'entraînement (web = biais humains) et leurs annotateurs RLHF. L'alignement réduit les biais les plus visibles, ne les élimine pas. En contexte pro : évaluer sur **vos** populations et **vos** cas.

**19.19. Que vaut la licence d'un modèle « ouvert » ?**
Tout dépend : **MIT/Apache 2.0** (DeepSeek-R1, Mistral, Qwen) = usage commercial libre ; **licence communautaire** (Llama) = gratuite mais avec restrictions (dont clause sur les très gros déploiements). Toujours lire la licence avant de déployer (idée reçue n° 12).

