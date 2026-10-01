---
id: collect-261001-general-networking/general-networking/ibm-granite-4-2-30b-parametres-57-swe-bench-2026-1
title: "Récupérer les poids depuis Hugging Face"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Mistral", "Nvidia"]
dates: []
keywords: ["agent", "agents", "apache", "benchmark", "benchmarks", "deepseek", "llama", "mai", "mistral", "moe", "nvidia", "open source"]
source: docs/RAG/collect-261001-general-networking/ibm-granite-4-2-30b-parametres-57-swe-bench-2026.md
source_anchor: ""
source_lines: [1, 42]
sha256: 949fbf7dba3203f4d865ce723a9eb1d0c0c86d79f42383214b587cfda97911ea
---

# Récupérer les poids depuis Hugging Face

IBM a mis en ligne le 25 août 2026 la famille **Granite 4.2**, trois modèles de langage en poids ouverts (3, 8 et 30 milliards de paramètres) pensés pour les agents d’entreprise plutôt que pour le grand public. Le modèle phare, Granite 4.2 30B, revendique un score de 57,00 sur SWE-Bench Verified, l’un des benchmarks les plus exigeants pour mesurer la capacité d’une IA à corriger du code réel. Diffusés sous licence Apache 2.0, ces modèles arrivent alors que l’Europe cherche des alternatives crédibles aux géants américains et chinois pour équiper ses propres agents IA en entreprise. Voici ce que Granite 4.2 change concrètement, comment il se positionne face à Llama 4, Mistral ou Qwen3.8, et ce qu’il reste encore à démontrer.

## Qu’est-ce que Granite 4.2, la nouvelle famille de modèles d’IBM ?

Granite 4.2 est une famille de modèles de langage **denses et de type décodeur**, c’est-à-dire des transformeurs classiques et non des architectures à mélange d’experts (MoE). IBM la présente comme sa première génération de modèles « raisonnants » pensés spécifiquement pour des agents d’entreprise capables de manipuler des outils réels : environnements de développement logiciel, terminaux et moteurs de recherche web. Trois tailles sont proposées : 3, 8 et 30 milliards de paramètres, ce qui permet de couvrir aussi bien des déploiements légers en périphérie (edge) que des charges de travail plus lourdes nécessitant un raisonnement approfondi.

Le trait distinctif de Granite 4.2 est son **mode de raisonnement commutable** (« switchable thinking mode »). Concrètement, l’utilisateur ou le système peut activer ou désactiver un mode de délibération plus poussé selon la complexité de la tâche, un compromis entre vitesse de réponse et qualité de raisonnement que l’on retrouve désormais chez plusieurs fournisseurs de modèles à vocation professionnelle. Pour les versions 8B et 30B, IBM indique avoir appliqué de l’**apprentissage par renforcement dans des environnements réels** plutôt que dans des simulations abstraites, ce qui expliquerait en partie les scores obtenus sur les benchmarks de programmation.

## Les chiffres clés : tailles, architecture et licence Apache 2.0

Le choix de la licence Apache 2.0 n’est pas anodin. Elle autorise l’usage commercial, la modification et la redistribution des poids sans redevance, à l’inverse de licences plus restrictives utilisées par certains concurrents pour leurs modèles les plus puissants. IBM publie non seulement les poids des trois tailles de Granite 4.2, mais aussi leurs variantes quantifiées, le dépôt de code source et la documentation associée. C’est cette combinaison, ouverture totale des poids et licence permissive, qui distingue Granite du positionnement plus fermé d’une partie du marché.

| Modèle | Paramètres | Architecture | Score SWE-Bench Verified | Usage cible | 
|---|---|---|---|---|
| Granite 4.2 3B | 3 milliards | Dense, décodeur | Non communiqué | Agents légers, périphérie | 
| Granite 4.2 8B | 8 milliards | Dense, décodeur, RL en environnement réel | 47,67 | Agents de développement, terminal | 
| Granite 4.2 30B | 30 milliards | Dense, décodeur, post-entraîné depuis Granite 4.1 30B | 57,00 | Raisonnement complexe, agents d’entreprise | 

IBM n’a pas encore publié la longueur exacte de la fenêtre de contexte ni la grille tarifaire pour un hébergement via watsonx.ai au moment de la rédaction de cet article. Ces informations figurent habituellement dans le catalogue des modèles de fondation de watsonx quelques semaines après une annonce, comme cela avait été le cas pour Granite 4.1.

## Des scores SWE-Bench qui interpellent : 57 % pour le modèle 30B

SWE-Bench Verified est devenu, au fil de 2025 et 2026, l’une des références les plus citées pour évaluer la capacité réelle d’un modèle à résoudre des tickets de bugs issus de véritables dépôts open source, et non de simples exercices de programmation isolés. Un score de **57,00 pour Granite 4.2 30B** place le modèle dans une fourchette respectable pour un modèle ouvert de cette taille, sachant que les meilleurs modèles propriétaires du marché dépassent aujourd’hui les 70 points sur ce même benchmark. Le modèle 8B, avec 47,67, illustre la logique de compromis d’IBM : offrir un modèle nettement moins gourmand en ressources tout en conservant une capacité de raisonnement agentique correcte.

Il manque toutefois un élément essentiel pour juger pleinement Granite 4.2 : IBM n’a pas publié de comparaison directe, chiffrée, face à Llama 4, Mistral Large 3, Qwen3.8-Max ou DeepSeek V4 sur ce même protocole d’évaluation. Sans ces chiffres officiels, toute affirmation de supériorité ou d’infériorité relèverait de la spéculation. Ce que l’on peut dire avec certitude, c’est qu’IBM entre dans la course des modèles ouverts « orientés code et agents » avec des résultats vérifiables sur au moins un benchmark reconnu, ce qui n’est pas le cas de tous les acteurs du secteur.

## Le mode de raisonnement commutable et l’apprentissage par renforcement

La bascule entre un mode de réponse rapide et un mode de raisonnement approfondi n’est pas une invention d’IBM : elle s’inscrit dans une tendance de fond où les fournisseurs cherchent à réduire les coûts d’inférence sur les tâches simples tout en réservant le calcul intensif aux problèmes complexes. Ce qui distingue Granite 4.2, c’est la manière dont IBM a entraîné les versions 8B et 30B : plutôt que de s’appuyer uniquement sur des jeux de données statiques, les modèles ont été affinés par renforcement **dans des environnements interactifs réels**, notamment des outils de développement logiciel, des interpréteurs de commande (terminal) et des moteurs de recherche web.

### Pourquoi ce choix d’entraînement compte pour les entreprises

Un agent d’entreprise qui doit ouvrir un ticket, exécuter une commande shell ou interroger une base de connaissances interne échoue souvent non pas parce qu’il « ne sait pas », mais parce qu’il n’a jamais été confronté aux frictions réelles de ces environnements : messages d’erreur ambigus, permissions refusées, résultats de recherche contradictoires. En s’entraînant directement dans ces conditions, IBM cherche à réduire l’écart entre les performances mesurées en laboratoire et le comportement observé en production, un problème récurrent pointé par les équipes qui déploient des agents IA à grande échelle depuis 2025.

## De Granite 3.0 à Granite 4.2 : la trajectoire d’IBM en open source

IBM n’improvise pas ce virage vers l’ouverture. L’entreprise avait déjà publié ses premiers modèles de code Granite sous licence Apache 2.0 dès mai 2024, en les rendant disponibles sur Hugging Face et GitHub. La génération Granite 3.0 avait ensuite introduit des variantes à mélange d’experts (MoE) aux côtés de modèles denses classiques, distribuées via Hugging Face, Ollama et les microservices NVIDIA. Granite 4.2 marque une inflexion : IBM abandonne, pour cette génération précise, l’architecture MoE au profit d’un choix pleinement dense, et recentre l’ensemble de la gamme sur le raisonnement agentique plutôt que sur la génération de texte généraliste.

Le modèle 30B de Granite 4.2 n’est d’ailleurs pas parti de zéro : il a été **post-entraîné à partir de Granite 4.1 30B**, la génération précédente déjà intégrée au catalogue de modèles de fondation de watsonx.ai. Cette continuité technique explique en partie la rapidité avec laquelle IBM a pu livrer une nouvelle génération dotée de capacités de raisonnement renforcées, sans repartir d’un pré-entraînement complet.

## watsonx.ai : IBM revendique une nette croissance de ses revenus IA

