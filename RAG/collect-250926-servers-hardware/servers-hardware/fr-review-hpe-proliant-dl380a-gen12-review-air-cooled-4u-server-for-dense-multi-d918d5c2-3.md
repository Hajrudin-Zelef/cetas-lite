---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-d918d5c2-3
title: "fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["apache", "benchmark", "benchmarks", "diffusion", "gpu", "intel", "llama", "nvfp4", "nvidia", "open source", "packaging", "valuation"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2.md
source_anchor: ""
source_lines: [53, 83]
sha256: d8e1851dc7866dc3a9206031b6f2dae227f2b9d0148dd838fdddac6b808acb2c
---

# fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2

L'interface signale clairement les points critiques, tels que les certificats auto-signés ou une gestion des clés non configurée, tout en garantissant un fonctionnement sécurisé le cas échéant. Les journaux de sécurité, les politiques de gestion des utilisateurs et les contrôles d'accès sont facilement accessibles, offrant ainsi aux administrateurs une vue d'ensemble complète de la sécurité du système directement depuis l'environnement iLO.
Sous l'onglet Applications HPE, iLO 7 donne accès à des outils intégrés qui optimisent le déploiement et la gestion du cycle de vie des serveurs. À partir de cette interface, les administrateurs peuvent lancer Intelligent Provisioning, un utilitaire intégré conçu pour simplifier l'installation du système d'exploitation, les mises à jour du firmware et la configuration du système sans nécessiter de support externe.
L'onglet Paramètres iLO regroupe toutes les options de configuration et d'administration de l'interface iLO. Les administrateurs peuvent y contrôler l'accès des utilisateurs, la configuration des ports réseau, les méthodes d'authentification et les journaux d'activité. Le menu propose également des options de dépannage, d'application des politiques de sécurité, de gestion des licences et de synchronisation de l'heure.
Les actions rapides, telles que la sauvegarde ou la restauration de la configuration iLO et la réinitialisation, sont facilement accessibles à droite, ce qui simplifie les tâches de maintenance. L'interface, à l'image du reste de l'interface moderne par cartes d'iLO 7, offre une gestion claire et organisée de la sécurité, de la connectivité et des paramètres opérationnels depuis un emplacement centralisé.
Test de performance
Pour évaluer les capacités réelles du DL380a Gen12, nous avons mené une série complète de tests de performance couvrant à la fois l'inférence IA et les charges de travail de calcul générales. Ces tests incluent des benchmarks de service en ligne vLLM pour les grands modèles de langage (LLM) et des benchmarks de la suite de tests Phoronix afin de mesurer le débit du processeur, la bande passante mémoire, l'efficacité du service web et les performances cryptographiques.
Configuration du système
- CPU: 2 processeurs Intel Xeon 6527P
- Mémoire: Kit intelligent HPE 64 Go 2Rx4 PC5-3400B-R (16 disques)
- GPU: 4 x NVIDIA RTX PRO 6000 (96 Go)
- Stockage: 2 x 15.63 To PM1733a U.3
Service en ligne vLLM – Performances d'inférence LLM
vLLM est le moteur d'inférence et de diffusion à haut débit le plus populaire pour les LLM. Le benchmark de diffusion en ligne vLLM est un outil d'évaluation des performances qui mesure les capacités de diffusion réelles de ce moteur d'inférence sous des requêtes simultanées. Il simule les charges de travail de production en envoyant des requêtes à un serveur vLLM en cours d'exécution avec des paramètres configurables, tels que le débit de requêtes, la longueur des entrées/sorties et le nombre de clients simultanés. Le benchmark mesure des indicateurs clés, notamment le débit (tok/s), le temps d'obtention du premier jeton et le temps d'exécution par jeton de sortie, permettant ainsi aux utilisateurs de comprendre les performances de vLLM sous différentes conditions de charge.
Nous avons testé les performances d'inférence sur trois modèles représentatifs couvrant différentes échelles et approches de quantification, en évaluant comment les quatre GPU NVIDIA RTX PRO 6000 du HPE ProLiant DL380a Gen12 gèrent les charges de travail d'inférence de production.
Performances du modèle dense
Les modèles denses représentent l'architecture LLM conventionnelle, où tous les paramètres et activations sont utilisés lors de l'inférence. Nous avons évalué deux configurations de modèles denses : Llama-2-70b-chat-hf et Llama-3.2-90B-Vision-Instruct.
Performances de Llama-2-70B-Chat
En configuration mono-utilisateur (BS=1) avec TP=4, le modèle atteint 32.89 tok/s par utilisateur et un TPOT de 30.18 ms. Avec BS=8, les performances atteignent 15.68 tok/s par utilisateur, pour un débit total de 433.62 tok/s et un TPOT de 35.98 ms. En passant à BS=32, le débit total atteint 741.62 tok/s tout en maintenant 8.00 tok/s par utilisateur et un TPOT de 43.44 ms.
Performance d'instruction de vision du lama 3.2-90B
Avec une station de base (BS) de 1 et un débit total (TP) de 4, le modèle atteint un débit de 20.59 tok/s par utilisateur et un TPOT de 38.27 ms. Avec une BS de 16, les performances passent à 7.20 tok/s par utilisateur, pour un débit total de 806.14 tok/s et un TPOT de 54.98 ms. Le débit total maximal de 1 372,21 tok/s est atteint avec une BS de 128, soit 2.59 tok/s par utilisateur et un TPOT de 122.75 ms.
Performances des types de données à micro-échelle
La micro-échelle représente une approche de quantification avancée qui applique des facteurs d'échelle précis à de petits blocs de poids, plutôt qu'une quantification uniforme à de grands groupes de paramètres. Le format NVFP4 de NVIDIA implémente cette technique grâce à une représentation en virgule flottante par blocs, où chaque bloc de micro-échelle de 8 à 32 valeurs partage un exposant commun servant de facteur d'échelle. Cette approche granulaire préserve la précision numérique tout en assurant une représentation sur 4 bits, maintenant ainsi la plage dynamique essentielle aux architectures de transformateurs. Ce format s'intègre à l'architecture Tensor Core de NVIDIA sur la RTX PRO 6000, permettant un calcul efficace en précision mixte avec décompression à la volée lors des opérations matricielles.
Performances du GPT-OSS-120B
Nous avons évalué le modèle GPT-OSS-120B d'OpenAI avec la quantification NVFP4. En mode mono-utilisateur (TP=2), le modèle atteint 176.09 tok/s par utilisateur avec un TPOT de 5.46 ms, soit la latence la plus faible de notre suite de tests. Avec BS=4 et TP=4, les performances atteignent 105.79 tok/s par utilisateur, pour un débit total de 1155.94 tok/s et un TPOT de 7.79 ms. Avec BS=32 et TP=4, le débit passe à 47.54 tok/s par utilisateur et 3956.44 tok/s au total, avec un TPOT de 13.86 ms. Le débit total maximal de 4015.77 tok/s est atteint avec BS=64, soit 25.38 tok/s par utilisateur et un TPOT de 14.78 ms.
Points de repère Phoronix
Phoronix Test Suite est une plateforme d'analyse comparative automatisée et open source prenant en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue. Nous nous concentrerons sur les tests suivants : Stream, 7-Zip, compilation du noyau Linux, Apache et OpenSSL.
Bande passante de la mémoire de flux
Dans le benchmark Stream, qui mesure le débit mémoire brut, le HPE DL380a Gen12 a atteint un impressionnant score de 542 Go/s, démontrant ainsi la capacité de la plateforme à maintenir des débits de transfert de données élevés sous charge continue. Ce niveau de bande passante rend le système particulièrement performant pour des charges de travail telles que la modélisation de données, la simulation et l'inférence IA, où de grands ensembles de données doivent être transférés rapidement entre la mémoire et les ressources de calcul.
Compression à 7 zips
Le test de compression 7-Zip a mesuré 305 000 MIP, soulignant l'excellente efficacité multithread du système pour les opérations de compression et de décompression exigeantes en calcul. Ces résultats font du DL380a Gen12 un choix idéal pour les environnements impliquant des opérations fréquentes de packaging de données, d'archivage ou de sauvegarde, qui nécessitent des performances CPU constantes et reproductibles.
Compilation du noyau
