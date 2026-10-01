---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-15
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Baidu", "DeepSeek", "Google", "Meta", "Mistral", "OpenAI", "United States"]
dates: []
keywords: ["agents", "agi", "apache", "benchmarks", "compute", "copyright", "deepseek", "diffusion", "distillation", "distribution", "embeddings", "fine-tuning"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [951, 1013]
sha256: 50244a5304ba7a697e629f0fb1d8b184955b574540784c7a3ee8452b9f93f692
---

# IA — Le grand dossier

**1. « Plus gros = toujours meilleur. »**
Réalité : vrai **à données et entraînement égaux** (lois de scaling), faux en pratique — Chinchilla (2022) a montré qu'un 70B bien entraîné bat un 280B sous-entraîné. Depuis 2024, les petits modèles sur-entraînés et distillés dominent le rapport qualité/prix.

**2. « L'open source est en retard de 2 ans sur le fermé. »**
Réalité : l'écart s'est réduit à **quelques mois** sur la plupart des tâches (Llama 3, Qwen, DeepSeek-V3/R1). Le fermé garde l'avantage sur la frontière extrême (raisonnement profond, multimodalité intégrée) et sur le produit (UX, distribution).

**3. « On va manquer de données, c'est la fin. »**
Réalité : le **texte humain de qualité** se raréfie (Epoch AI : 2026-2032), mais trois échappatoires existent — données synthétiques/distillation, RL avec récompenses vérifiables (sans données humaines, voie AlphaGo Zero), et données multimodales/capteurs (ordres de grandeur supérieurs). Le mur concerne le **pré-entraînement texte naïf**, pas l'apprentissage en général.

**4. « GPT-4 a N milliards de paramètres » (N au choix).**
Réalité : **OpenAI n'a jamais publié** l'architecture, la taille ni les données de GPT-4. Tout chiffre circulant est une spéculation (les estimations MoE ~1 800B totaux sont des rumeurs — « à vérifier », à ne jamais citer comme fait).

**5. « L'IA va remplacer les sysadmins. »**
Réalité : l'IA déplace le travail, elle ne le supprime pas — comme chaque vague d'automatisation (virtualisation, cloud, IaC). Le poste évolue vers : serving de modèles, FinOps IA, sécurité des agents, énergétique des baies. Les équipes qui adoptent l'IA comme outil remplacent celles qui l'ignorent, pas l'inverse.

**6. « Il faut des GPU à 30 000 $ pour faire de l'IA en entreprise. »**
Réalité : un **RAG d'entreprise** tourne sur une RTX 4090 (24 Go, ~2 000 $) avec un 8B-32B quantifié, voire sur CPU pour les embeddings. Les H100 sont pour la frontière, pas pour l'usage courant.

**7. « Le RAG, c'est juste un vector DB + un prompt. »**
Réalité : c'est l'architecture de base, mais la qualité vient du **chunking**, du **reranking**, de l'**évaluation** et du **nettoyage du corpus** — 80 % du travail est data, 20 % modèle. (C'est d'ailleurs exactement la méthode de Zelef : il scrape et nettoie lui-même.)

**8. « Les modèles ne font que recracher leur entraînement. »**
Réalité : ils **généralisent** au-delà (cf. GPT-3 résolvant des additions absentes du training set, R1-Zero découvrant l'auto-vérification par RL). La mémorisation existe (et pose des problèmes de copyright/fuites), mais ne résume pas le comportement.

**9. « L'alignement (RLHF), c'est juste de la censure. »**
Réalité : le RLHF aligne d'abord sur **l'utilité** (suivre les instructions — InstructGPT 2022) ; la « censure » (refus) n'en est qu'une facette, et elle est paramétrable (system prompts, fine-tuning). Sans alignement, pas de produit : GPT-3 brut ne « répond » pas aux questions, il les « continue ».

**10. « La Chine est en retard à cause des sanctions. »**
Réalité : les sanctions ont **forcé l'efficacité** — DeepSeek-V3/R1 (2024-2025) rivalise avec la frontière US pour une fraction du coût, précisément parce que l'équipe était contrainte aux H800 bridés. Le retard éventuel est sur le **compute brut**, pas sur l'algorithmique.

**11. « L'inférence, c'est gratuit une fois le modèle entraîné. »**
Réalité : c'est **le poste dominant** en 2026 (débat « 80/20 reversal »). Servir des milliards de requêtes — surtout avec des agents qui « réfléchissent » — coûte plus cher que l'entraînement. Le coût se pilote : quantification, routage, caching, budgets.

**12. « Un modèle open-weight = open source. »**
Réalité : **non**. « Open-weight » = poids téléchargeables ; « open source » (définition OSI) exige aussi données, code d'entraînement, licence libre. Llama a une licence **communautaire restrictive** ; Mistral et Qwen sont en **Apache 2.0** (plus libre) ; DeepSeek-R1 en **MIT**. La licence conditionne l'usage commercial — la lire avant de déployer.

**13. « Les benchmarks disent quel modèle est le meilleur. »**
Réalité : les benchmarks mesurent des **tâches précises**, sont sujets à la **contamination** (le test était dans le training) et au **gaming** (optimiser pour le test). Un modèle n° 1 sur MMLU peut être moins bon **pour votre cas d'usage**. Règle : toujours évaluer sur **vos** données, **vos** questions.

**14. « L'AGI est pour 2027, c'est mathématique. »**
Réalité : c'est une **prédiction** (Amodei : « Machines of Loving Grace », 2024 ; d'autres disent 2030+, LeCun dit « pas avec les LLM »). Aucune loi de la physique ne donne une date. Les extrapolations de courbes passées sont des opinions, pas des faits — les présenter comme telles.

**15. « Le scaling est mort, regardez GPT-5. »**
Réalité : le **scaling naïf du pré-entraînement** montre des rendements décroissants (vrai), mais le progrès a **changé d'axe** : test-time compute (o1 : AIME 13 % → 83 %), efficacité (R1 : niveau o1 pour ~294 k$ de RL — « à vérifier »), agents. Confondre « la fin d'un levier » avec « la fin du progrès » est l'erreur symétrique de celle des maximalistes.

---

## 12. Portraits approfondis : 10 figures à connaître absolument

> Complément de la section 2 : des fiches plus longues sur des personnes dont le nom revient sans cesse, avec le fait vérifié qui les rend incontournables.

**12.1. Fei-Fei Li — la femme sans qui rien n'existait.**
Professeure à Stanford, elle lance en **2009** le projet **ImageNet** : ~15 millions d'images étiquetées à la main via Mechanical Turk, 22 000 catégories. À l'époque, ses collègues la prennent pour une excentrique (« trop de données, pas assez d'idées »). C'est **son** dataset qui rend AlexNet possible (2012). Elle dirige ensuite le Stanford HAI (Human-Centered AI, 2019) et fonde **World Labs** (2024, ~1 Md$ levé — « à vérifier » le montant exact) sur les **world models** spatiaux. Leçon : en IA, **les datasets sont des contributions scientifiques de premier rang**, pas de la plomberie.

**12.2. Ian Goodfellow — l'inventeur des GAN.**
Doctorant de Yoshua Bengio à Montréal, il invente les **GAN en 2014** (l'anecdote — vraie, racontée par Goodfellow lui-même — : l'idée lui vient dans un bar en discutant avec des amis). Parcours : Google Brain, Apple (directeur ML, 2019-2022 — départ lié au retour au bureau imposé, épisode public), puis retour à la recherche. Les GAN ont dominé la génération d'images jusqu'à la diffusion (2021-2022) et restent utilisées (data augmentation, deepfakes — d'où les débats).

**12.3. Jeff Dean — l'infrastructure faite homme.**
Chez Google depuis 1999, co-auteur de **MapReduce, BigTable, Spanner**, cofondateur de **Google Brain** (2011) avec Andrew Ng, puis « Senior Fellow ». C'est lui qui a rendu possible l'entraînement distribué à l'échelle Google (DistBelief, puis TensorFlow). En 2023, il prend la tête de **Google DeepMind Research** après la fusion. Figure du « camp systèmes » : pour lui, le progrès IA est d'abord un problème d'**ingénierie à grande échelle** — position qui parle directement aux sysadmins.

**12.4. Andrew Ng — le vulgarisateur en chef.**
Cofondateur de Google Brain (2011), chief scientist de **Baidu** (2014-2017, il y monte l'équipe IA à 1 300 personnes — chiffre public de l'époque), cofondateur de **Coursera** (2012 : le cours ML suivi par des millions de personnes), fondateur de **Landing AI** et de **DeepLearning.AI**. Position : comme LeCun, il minimise le risque existentiel et insiste sur les **risques concrets** (biais, concentration) ; il défend l'open source et l'IA appliquée aux pays en développement.

