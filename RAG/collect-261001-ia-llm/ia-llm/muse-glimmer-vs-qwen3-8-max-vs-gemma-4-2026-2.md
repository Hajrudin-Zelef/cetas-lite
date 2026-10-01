---
id: collect-261001-ia-llm/ia-llm/muse-glimmer-vs-qwen3-8-max-vs-gemma-4-2026-2
title: "Exemple : servir Muse Glimmer quantifié en local avec llama.cpp"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Google", "Hugging Face", "Meta"]
dates: []
keywords: ["muse", "agents", "apache", "benchmark", "benchmarks", "embeddings", "gpu", "gqa", "moe", "multimodal", "open weights", "qwen"]
source: docs/RAG/collect-261001-ia-llm/muse-glimmer-vs-qwen3-8-max-vs-gemma-4-2026.md
source_anchor: ""
source_lines: [31, 74]
sha256: 4454e37947260264dff44c916391c74dc010cc4aafcfc34825872b6b53bb5650
---

# Exemple : servir Muse Glimmer quantifié en local avec llama.cpp

Le déploiement local est hors de portée pour un particulier ou une petite équipe. Même avec seulement 95 milliards de paramètres actifs par passage, il faut stocker l’intégralité du mélange d’experts en mémoire pour un service à faible latence, ce qui pousse Qwen3.8-2.4T-A95B vers des configurations multi-GPU ou multi-TPU de classe data center. Aucun fournisseur ne communique de configuration “un seul GPU” pour ce modèle, contrairement à Muse Glimmer.

Sur les benchmarks, Qwen3.8-Max dispose du chiffre le plus concret et le plus vérifiable des trois modèles de ce comparatif : un score de **67,7 % sur SWE-Bench Pro**, obtenu avec la configuration d’échafaudage (scaffolding) propre à Qwen, ce qui le place parmi les modèles ouverts les plus compétents pour les tâches de codage et d’agents logiciels. Plusieurs analyses, dont celle de MarkTechPost, présentent également le modèle comme quasi état de l’art sur MMLU et GPQA parmi les modèles ouverts, sans toutefois publier de pourcentage exact dans les extraits accessibles (source).

Côté licence, la variante 27 milliards de paramètres de la même génération Qwen3.8 reste sous Apache 2.0, mais le checkpoint 2,4 billions de paramètres Max utilise une **licence Qwen personnalisée**, plus restrictive, sans que le détail exact des clauses de seuil de revenu soit repris dans les communications publiques consultées. Pour une entreprise, cela signifie qu’il faut lire attentivement le fichier de licence du dépôt avant tout déploiement commercial à grande échelle.

Une alternative existe pour qui ne veut pas gérer un cluster : l’API cloud Qwen3.8-Max facturée **2 dollars par million de tokens en entrée et 6 dollars par million de tokens en sortie**, avec un accès à la fenêtre de contexte complète de 1 million de tokens. C’est le compromis choisi par la majorité des équipes qui veulent la puissance du modèle sans investir dans l’infrastructure de calcul nécessaire à l’auto-hébergement.

## Gemma 4 12B de Google DeepMind : le multimodal sans encodeur

Gemma 4 12B a été présenté par Google Developers le 3 juin 2026 comme “un modèle multimodal unifié, sans encodeur” (source). La rupture technique tient dans sa description officielle : “aucun encodeur multimodal séparé. Les entrées vision et audio circulent directement dans le tronc du LLM” (source). Concrètement, plutôt que de faire passer une image par un module de vision distinct puis de la projeter dans l’espace du langage, Gemma 4 tokenise directement image, audio et vidéo dans une représentation partagée traitée par le même transformeur de 12 milliards de paramètres que le texte.

Le modèle prend en charge quatre modalités en entrée : texte, image, audio et vidéo, ce qui en fait le plus polyvalent des trois modèles comparés ici sur le plan strictement multimodal, Muse Glimmer se limitant au texte et à l’image, et Qwen3.8-2.4T-A95B restant essentiellement textuel dans sa variante open weights. En contrepartie, Gemma 4 reste le plus petit des trois par le nombre de paramètres, avec seulement 12 milliards.

L’empreinte mémoire documentée est de 26,7 Go en BF16, ce qui correspond à peu près à ce qu’on attend d’un modèle de 12 milliards de paramètres incluant les embeddings multimodaux additionnels. Cela signifie qu’un GPU de 32 Go, comme une RTX 5090 ou une carte professionnelle d’entrée de gamme, suffit pour une inférence en pleine précision. Aucune fiche technique consultée ne publie de chiffre exact pour les versions quantifiées en 4 ou 8 bits, mais la taille du modèle laisse penser qu’une quantification poussée le ferait tenir sur des GPU grand public de 12 à 16 Go, une estimation à vérifier au cas par cas selon le format de quantification retenu.

Sur les benchmarks chiffrés, les sources consultées ne publient pas de score MMLU, GPQA Diamond, SWE-Bench ou AIME 2025 précis pour Gemma 4 12B à la date de rédaction de cet article : Google communique sur la capacité multimodale unifiée plutôt que sur des tableaux de scores comparatifs, et la fiche technique complète n’était pas accessible dans les sources consultées. C’est un manque notable face à la transparence relative de Qwen sur SWE-Bench Pro.

La licence suit le modèle Gemma habituel : des poids téléchargeables gratuitement depuis Hugging Face et Kaggle, sous une licence Gemma propriétaire à Google, permissive pour un usage commercial mais distincte d’Apache 2.0, avec des clauses spécifiques (notamment une restriction d’usage pour entraîner un modèle concurrent des services Google) qui nécessitent une lecture attentive avant intégration en production. Une passerelle API existe également via Google AI Studio et Vertex AI, facturée au token, mais aucun tarif précis par million de tokens n’était communiqué publiquement dans les sources consultées.

## Tableau comparatif des spécifications techniques

| Critère | Muse Glimmer 30B (Meta) | Qwen3.8-2.4T-A95B (Alibaba) | Gemma 4 12B (Google DeepMind) | 
|---|---|---|---|
| Paramètres totaux | ~29,6-30 Md (dense) | 2,4 billions (MoE) | 12 Md (dense) | 
| Paramètres actifs par requête | ~30 Md (pas de MoE) | 95 Md | 12 Md | 
| Architecture | Transformeur dense causal, GQA 32:2 | Mélange d’experts (MoE) | Multimodal unifié, sans encodeur séparé | 
| Fenêtre de contexte | 131 072 tokens | 256 000 tokens (poids ouverts) / 1M (API) | Non précisée dans la fiche publique | 
| Modalités en entrée | Texte + image (vidéo partielle) | Texte (checkpoint open weights) | Texte, image, audio, vidéo | 
| Sortie | Texte uniquement | Texte uniquement | Texte (génération multimodale non confirmée) | 
| Licence | Apache 2.0 | Licence Qwen personnalisée | Licence Gemma (Google) | 
| Date de publication des poids | 10 août 2026 | 12-13 août 2026 | 3 juin 2026 | 
| Taille BF16 (pleine précision) | ~55-60 Go | Échelle data center (multi-fichiers) | 26,7 Go | 
| Taille quantifiée (4 bits environ) | 17-20 Go | Non pertinent à cette échelle | Non publiée officiellement | 
| Matériel minimal recommandé | 1 GPU 24 Go (RTX 4090/5090) | Cluster multi-GPU/TPU classe data center | 1 GPU 32 Go en BF16 | 
| Débit mesuré | 233,4 tokens/s sur RTX 5090 (DFlash) | Non communiqué pour un GPU unique | Non communiqué | 
| Accès API tiers | Hébergeurs tiers, tarifs variables | 2 $ / 6 $ par million de tokens (in/out) | Google AI Studio / Vertex AI, tarif non public | 

## Benchmarks : ce que disent (et ne disent pas) les chiffres publics

La comparaison directe des trois modèles sur un même jeu de benchmarks se heurte à un problème réel en cet été 2026 : les trois éditeurs n’ont pas publié les mêmes suites de tests au même niveau de détail. C’est un point important à souligner plutôt que de combler artificiellement les trous avec des estimations présentées comme des faits.

Qwen3.8-Max est le mieux documenté sur le plan agentique et logiciel : son score de 67,7 % sur SWE-Bench Pro, obtenu avec le scaffolding propre à Alibaba, en fait une référence sérieuse pour les tâches de résolution de bugs et de génération de code multi-fichiers. C’est un chiffre directement comparable aux scores publiés par d’autres laboratoires sur le même benchmark, ce qui permet un positionnement assez précis face aux modèles fermés. Le comparatif d’Aireiter souligne toutefois que les tableaux de benchmarks complets restaient “en mouvement” juste après la sortie, et renvoie au rapport technique d’Alibaba pour les chiffres définitifs (source).

