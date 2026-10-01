---
id: collect-250926-servers-hardware/servers-hardware/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-4
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Anthropic", "MiniMax", "Nvidia", "OpenRouter", "vLLM"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "blackwell", "claude", "diffusion", "gpu", "moe", "nvidia", "opus 4", "vllm"]
source: docs/RAG/clean4/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [81, 108]
sha256: 3f4cbc843adb0abcc5f686261c6ae0739599bd2bb2a78873143304678660d236
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

Cependant, cette capacité accrue a un coût important. La consommation électrique dépasse 8 kW. Les charges thermiques nécessitent une infrastructure de refroidissement dédiée au centre de données. Le niveau sonore empêche tout déploiement en dehors des salles machines spécialement conçues à cet effet. De plus, les délais de livraison s'étendent souvent de six à dix-huit mois en raison des contraintes d'approvisionnement persistantes sur les plateformes GPU d'entreprise.
Le Grando occupe une position différente. Pour les organisations qui privilégient un déploiement rapide, des environnements d'exploitation faciles à gérer et des charges de travail d'inférence ou de création plutôt qu'un entraînement distribué à grande échelle, les compromis sont souvent avantageux. Les équipes qui ont besoin de leur matériel immédiatement, dans un environnement opérationnel, peuvent trouver l'approche du Grando en matière de densité plus pratique que d'attendre une plateforme qu'elles ne pourront pas déployer efficacement une fois disponible.
Résultats des tests de performance du Comino Grando
Configuration du système
- Châssis: Comino Grando
- Carte mère: Support ASRock GENOAD8X-2T/BCM
- CPU: AMD EPYC 9474F 48C
- Mémoire: 512 Go de mémoire DDR5
- GPU: 8 x NVIDIA RTX PRO 6000
- Stockage: M.2 SSD
Service Claude Code – MiniMax M2.5
Au-delà des benchmarks traditionnels d'inférence LLM brute, nous souhaitions évaluer les performances de ce matériel dans un flux de travail de codage automatisé, notamment en gérant plusieurs sessions Claude Code simultanées à l'aide d'un modèle hébergé localement. Ce cas d'utilisation a un impact direct sur la productivité des équipes de développement : combien d'ingénieurs peuvent utiliser simultanément un assistant de codage IA hébergé sur un seul nœud avant que l'expérience ne se dégrade ?
Pour tester cela, nous avons conçu un environnement de test qui génère un ensemble de données de problèmes de programmation de difficulté moyenne (comme l'implémentation d'un cache LRU, la création d'une application de gestion de tâches en ligne de commande, l'écriture d'un convertisseur Markdown et la construction d'une API REST) et exécute chaque session Claude Code dans un conteneur Docker distinct, en interaction avec le serveur vLLM local. Un proxy transparent, placé entre les sessions et le point de terminaison d'inférence, capture les métriques de chaque requête pour chaque instance de Claude Code. Le modèle utilisé était MiniMax M2.5, déployé via vLLM sur les huit GPU NVIDIA RTX PRO 6000 du système. Bien qu'il ne s'agisse pas du modèle de programmation le mieux classé dans les palmarès publics, M2.5 est un modèle performant que de nombreux utilisateurs, y compris nos collègues développeurs, utilisent en local.
Pour établir un point de référence, nous utilisons le débit de sortie moyen de Claude Opus 4.6 d'Anthropic, mesuré via OpenRouter.ai, l'un des services de routage les plus utilisés pour l'accès aux API en production. Ce débit de référence correspond à environ 37 jetons par seconde et par requête API.
Nous avons mesuré deux indicateurs clés : le nombre moyen de jetons de sortie par seconde et par session Claude Code (ce que chaque développeur constate) et le nombre total de jetons de sortie par seconde pour toutes les sessions (le travail total produit par le serveur).
D'après les résultats, une session Claude Code simultanée unique offre un débit de 67.3 tok/s par utilisateur et un débit agrégé de 64.7 tok/s. Avec deux sessions simultanées, le débit par instance diminue légèrement à 57.4 tok/s, tandis que le débit agrégé atteint 95.1 tok/s, le traitement par lots de vLLM commençant à amortir la surcharge. Quatre sessions simultanées maintiennent un débit de 49.2 tok/s par utilisateur, offrant une expérience toujours très réactive pour les flux de travail de codage interactifs, tandis que le débit agrégé atteint 177.2 tok/s. Huit sessions représentent la configuration optimale pour le débit agrégé, avec un pic à 206.7 tok/s, tandis que le débit par instance se stabilise à 38.7 tok/s, un niveau confortable pour la génération et l'itération de code en temps réel.
Avec 16 sessions simultanées, le système présente le compromis classique du traitement par lots : le débit par instance chute à 31.1 tok/s et le débit agrégé à 105.8 tok/s. Cela suggère qu'à ce niveau de concurrence, le modèle 230B MiniMax M2.5 atteint les limites de ce que huit GPU peuvent supporter sans introduire de latence significative pour chaque utilisateur. La baisse globale observée entre 8 et 16 sessions reflète les besoins en bande passante mémoire d'une architecture MoE de grande taille soumise à une charge de décodage simultanée importante, plutôt qu'une inefficacité de l'ordonnancement.
Pour les organisations qui évaluent une infrastructure d'IA auto-hébergée pour leurs outils de développement, le Grando constitue une solution convaincante. Doté d'un modèle 230B de pointe, il peut gérer sans problème jusqu'à huit sessions Claude Code simultanées avec un débit offrant une véritable interactivité, et des vitesses par utilisateur dépassant 38 requêtes/seconde en pic de production agrégée. Des équipes de quatre à huit ingénieurs peuvent ainsi travailler à un débit quasi optimal sans perte de réactivité perceptible.
L'architecture à refroidissement liquide rend ce niveau de puissance de calcul utilisable dans des environnements où les serveurs GPU traditionnels ne peuvent pas fonctionner. Le système est suffisamment silencieux pour être installé dans un bureau de start-up, une petite salle des machines ou un coin dédié d'un espace de travail ouvert. Les systèmes à refroidissement par air avec une densité de GPU similaire atteignent généralement 90 dB ou plus, un niveau sonore suffisamment élevé pour nécessiter un espace dédié dans un centre de données ou, au minimum, une salle serveur fermée avec un traitement acoustique performant. Le Grando s'adapte à l'équipe qui l'utilise. Grâce à la localisation complète des données, l'absence de coûts d'API par jeton et un contrôle total sur le choix du modèle, il offre une solution auto-hébergée qui évolue avec une équipe de développement en pleine croissance sans nécessiter d'infrastructure de centre de données ni d'augmentation systématique des coûts.
Service en ligne vLLM – Performances d'inférence LLM
vLLM est l'un des moteurs d'inférence et de diffusion à haut débit les plus populaires pour les LLM. Le benchmark de diffusion en ligne de vLLM évalue les performances réelles de ce moteur d'inférence en cas de requêtes simultanées. Il simule les charges de travail de production en envoyant des requêtes à un serveur vLLM en cours d'exécution, avec des paramètres configurables tels que le débit de requêtes, les longueurs des entrées et sorties, et le nombre de clients simultanés. Ce benchmark mesure des indicateurs clés, notamment le débit (jetons par seconde), le temps d'obtention du premier jeton et le temps d'obtention du jeton de sortie (TPOT), permettant ainsi aux utilisateurs de comprendre les performances de vLLM sous différentes conditions de charge.
Nous avons testé les performances d'inférence sur une suite complète de modèles couvrant diverses architectures, échelles de paramètres et stratégies de quantification afin d'évaluer le débit sous différents profils de concurrence.
Résumé des résultats
| Modèle | La précision | Égal (256/256) | Préremplissage important (8k/1k) | Décodage intensif (1k/8k) | 
|---|---|---|---|---|
| Comino Grando avec 8× RTX PRO 6000 Blackwell — Résultats d'inférence vLLM (tok/s, pic à BS=256) |  |  |  |  | 
| GPT-OSS 20B | ep_dp1 | 17,280 | 32,061 | 11,187 | 
| GPT-OSS 120B | ep_dp1 | 11,726 | 21,636 | 7,570 | 
