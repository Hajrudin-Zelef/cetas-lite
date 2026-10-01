---
id: collect-261001-ia-llm/ia-llm/k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026-1
title: "k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Falcon", "Moonshot", "TII"]
dates: []
keywords: ["agents", "apache", "attention", "benchmarks", "fine-tuning", "kimi", "mai", "moe", "open source", "research", "sol", "valuation"]
source: docs/RAG/collect-261001-ia-llm/k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026.md
source_anchor: ""
source_lines: [1, 41]
sha256: 43f7c4fa8bd741a5a3137eecfd4b9eb6786783407e71e71d01088474f04b5a44
---

# k2-horizon-abou-dabi-publie-6-ia-ouvertes-375-md-2026

Le 3 septembre 2026, un institut de recherche basé à Abou Dabi a publié en une seule journée six modèles d’intelligence artificielle allant de 0,9 à 375 milliards de paramètres, avec les poids, le code, les données d’entraînement ET les recettes de fabrication en accès libre. L’annonce de l’Institute of Foundation Models (IFM), rattaché à la Mohamed bin Zayed University of Artificial Intelligence (MBZUAI), rebat les cartes du débat sur l’IA ouverte à un moment où Bruxelles cherche justement des alternatives transparentes aux modèles fermés américains et chinois.

Baptisée **K2 Horizon**, cette famille de modèles se distingue par une amplitude rarement vue chez un seul éditeur : un modèle de 0,9 milliard de paramètres taillé pour tourner sur une montre connectée jusqu’à un modèle phare de 375 milliards de paramètres pensé pour les charges de travail agentiques en entreprise. Le tout sous licence Apache 2.0, la licence open source la plus permissive du marché. Pour un institut installé à Paris, Abou Dabi et dans la Silicon Valley, l’enjeu dépasse la simple prouesse technique : il touche directement au débat français et européen sur la souveraineté numérique.

## K2 Horizon : six modèles, une seule philosophie d’ouverture

Le lancement du 3 septembre 2026 ne concerne pas un modèle unique mais une flotte complète de six systèmes, chacun calibré pour un usage matériel précis. Le plus petit, K2 Horizon 0,9B, vise les appareils très contraints comme les montres connectées et les lunettes intelligentes, avec une fenêtre de contexte de 128 000 tokens obtenue via une extension YaRN RoPE. Les modèles 3,7B et 7B, tous deux denses, ciblent les smartphones et les déploiements à faible coût, avec une fenêtre de contexte native de 524 288 tokens héritée des dernières phases d’entraînement.

Le modèle 32B, dense lui aussi, est présenté par IFM comme le plus puissant de la gamme dans cette catégorie, pensé pour l’hébergement local et les serveurs on-premise avec 512 000 tokens de contexte. Vient ensuite K2 Horizon 36B-A4B, un modèle épars combinant une architecture Mixture-of-Value-Attention (MoVA) et un mélange d’experts (MoE) avec seulement 4 milliards de paramètres actifs par requête, pensé pour absorber des charges de travail quotidiennes lourdes à moindre coût de calcul. Enfin, le modèle phare K2 Horizon 375B-A23B est un système MoE épars avec 375 milliards de paramètres au total et 23 milliards de paramètres actifs par token, positionné par IFM parmi les meilleurs modèles de sa catégorie de taille pour le raisonnement général, le code et les tâches agentiques.

Ce qui distingue vraiment cette sortie du reste du marché open source, c’est le niveau de transparence. IFM ne se contente pas de publier des poids figés : l’institut diffuse aussi le code d’entraînement, les points de contrôle intermédiaires (checkpoints) et, pour plusieurs modèles, les données et les recettes de construction complètes. Cette approche permet à des chercheurs extérieurs de reproduire les résultats et d’auditer la manière dont chaque modèle a été construit, une pratique encore rare à cette échelle de paramètres.

## Le tableau complet de la famille K2 Horizon

Voici la répartition détaillée des six modèles annoncés le 3 septembre 2026, avec leur architecture, leur fenêtre de contexte et leur usage cible tel que défini par IFM.

| Modèle | Paramètres (total / actifs) | Architecture | Contexte | Usage cible | 
|---|---|---|---|---|
| K2 Horizon 0,9B | 0,9 Md / 0,9 Md | Dense | 128 000 tokens | Montres, lunettes, matériel très contraint | 
| K2 Horizon 3,7B | 3,7 Md / 3,7 Md | Dense | 524 288 tokens | Smartphones, fine-tuning mono-nœud | 
| K2 Horizon 7B | 7 Md / 7 Md | Dense | 524 288 tokens | Smartphones, déploiement bas coût, code | 
| K2 Horizon 32B | 32 Md / 32 Md | Dense | 512 000 tokens | Hébergement local, serveurs on-premise | 
| K2 Horizon 36B-A4B | 36 Md / 4 Md | MoE + MoVA (épars) | 512 000 tokens | Charges de travail quotidiennes intensives | 
| K2 Horizon 375B-A23B | 375 Md / 23 Md | MoE (épars) | 524 000 tokens | Raisonnement, code, agents en entreprise | 

## Des scores encore partiels mais des signaux encourageants

Sur le terrain des benchmarks, la communication d’IFM reste pour l’instant incomplète : aucun tableau officiel complet de scores MMLU ou HumanEval n’a été publié pour l’ensemble des six modèles au moment du lancement. Les chiffres disponibles proviennent surtout d’évaluations indépendantes menées dans les jours suivant l’annonce. Le modèle phare 375B-A23B aurait obtenu un score de 47 sur l’Artificial Analysis Intelligence Index, un net progrès par rapport au précédent modèle de raisonnement d’IFM, K2 Think V2, qui plafonnait à 17 sur le même index. Cette évaluation indépendante souligne que le modèle se montre particulièrement performant sur les tâches agentiques et l’exécution de tâches réelles, tout en étant moins dominant sur le raisonnement à forte densité de connaissances que sa taille pourrait le laisser penser.

Côté modèles légers, le plus petit de la famille, K2 Horizon 0,9B, afficherait un score de 48,5 sur AIME 2026 et de 79,9 sur HumanEval+, des résultats notables pour un modèle suffisamment compact pour tourner en quantification sur une montre connectée. Ces chiffres, bien qu’issus de sources indépendantes plutôt que d’un rapport technique officiel exhaustif, positionnent K2 Horizon comme une famille crédible plutôt que comme un simple coup de communication.

## Qui se cache derrière l’Institute of Foundation Models

L’Institute of Foundation Models a été fondé en mai 2025 par la Mohamed bin Zayed University of Artificial Intelligence (MBZUAI), l’université émiratie spécialisée en IA. L’institut opère depuis trois sites : Abou Dabi, la Silicon Valley et Paris. Cette implantation parisienne mérite d’être soulignée, car elle donne à IFM une présence opérationnelle directe sur le sol français, à un moment où le débat sur la souveraineté de l’IA agite les administrations et les entreprises françaises.

Il est important de ne pas confondre IFM avec le Technology Innovation Institute (TII), l’autre grand laboratoire d’IA basé à Abou Dabi, connu pour sa série de modèles Falcon lancée dès 2023 avec Falcon 40B. TII est rattaché à l’Advanced Technology Research Council (ATRC) d’Abou Dabi, tandis qu’IFM dépend de MBZUAI. Les deux organisations sont donc distinctes, même si elles partagent la même ville et une ambition similaire de démocratiser l’IA ouverte depuis le Golfe. TII avait par exemple lancé Falcon-H1 Arabic en janvier 2026, un modèle hybride Mamba-Transformer dédié à la langue arabe.

## Attention à ne pas confondre K2 Horizon et Kimi K2

Le nom “K2” prête à confusion avec la série Kimi K2 de l’éditeur chinois Moonshot AI, mais il s’agit de deux projets totalement distincts, sans aucun lien capitalistique ni technique. Moonshot AI a construit sa propre trajectoire de sorties rapides tout au long de 2026 : Kimi K2 a ouvert la voie, suivi de Kimi K2.5 en janvier 2026 avec un encodeur visuel natif, puis de Kimi K2.6 en avril 2026 (1 000 milliards de paramètres), de Kimi K2.7 Code en juin 2026 pour les tâches de programmation, et enfin de Kimi K3 le 16 juillet 2026, un modèle MoE de 2,8 billions de paramètres avec une fenêtre de contexte d’un million de tokens.

