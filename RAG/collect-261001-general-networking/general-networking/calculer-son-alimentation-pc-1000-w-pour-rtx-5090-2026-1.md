---
id: collect-261001-general-networking/general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026-1
title: "Exemple : RTX 5090 + Ryzen 9 9950X3D"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "attention", "gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026.md
source_anchor: ""
source_lines: [1, 54]
sha256: d2a701261b0c9a44d0c92b682f808971f2716f9c32f0694d4a950df7be13f1ee
---

# Exemple : RTX 5090 + Ryzen 9 9950X3D

Un PC gamer qui redémarre seul en pleine partie, une carte graphique RTX 5090 qui plafonne ses performances, ou pire, un connecteur qui fond dans le boîtier : la majorité de ces incidents remontent à une alimentation mal dimensionnée. Avec l’arrivée des GPU RTX 50 et de la norme ATX 3.1, calculer précisément le wattage nécessaire n’est plus un détail réservé aux passionnés d’overclocking, c’est devenu une étape obligatoire pour tout le monde. Ce tutoriel vous montre comment calculer et choisir son alimentation PC (PSU) en 2026, du calcul de watts jusqu’à l’installation physique, avec des scripts pour automatiser le calcul et vérifier votre consommation réelle.

## Pourquoi le calcul de l’alimentation PC est devenu crucial en 2026

Il y a encore quelques années, choisir une alimentation PC se résumait à prendre un modèle 650 W ou 750 W “au cas où” et à ne plus y penser. Cette approche approximative ne fonctionne plus. Trois évolutions récentes ont rendu le calcul précis indispensable. D’abord, la consommation des cartes graphiques haut de gamme a fortement augmenté : une RTX 5090 consomme à elle seule plus qu’un PC gamer complet d’il y a cinq ans. Ensuite, le connecteur d’alimentation GPU a changé de nature avec le 12V-2×6, qui laisse beaucoup moins de marge d’erreur qu’un ancien connecteur PCIe 8 broches classique. Enfin, la flambée des prix des composants, RAM en tête avec une hausse de l’ordre de 171 % rapportée en 2026, pousse de nombreux acheteurs à vouloir optimiser chaque euro dépensé, y compris sur un composant aussi discret que le bloc d’alimentation.

Le résultat, c’est qu’une alimentation mal dimensionnée n’est plus seulement une question de performance : c’est devenu un enjeu de sécurité matérielle direct. Un sous-dimensionnement provoque des redémarrages en pleine partie ou des écrans bleus aléatoires, tandis qu’un mauvais choix de câblage ou d’adaptateur sur le connecteur GPU peut, dans les cas les plus graves, provoquer une surchauffe localisée du connecteur. À l’inverse, prendre systématiquement le modèle le plus puissant du marché “pour être tranquille” gaspille de l’argent sans bénéfice réel, puisqu’une alimentation surdimensionnée fonctionne loin de son point de rendement optimal la plupart du temps. La méthode présentée dans ce guide vise un équilibre précis entre sécurité et budget, construit à partir des données constructeur officielles plutôt que d’estimations approximatives glanées sur des forums.

## Prérequis : matériel, outils et informations nécessaires

Avant de vous lancer dans le calcul de votre alimentation PC, réunissez les informations et le matériel suivants. Ce tutoriel ne nécessite aucune compétence avancée en électronique, mais quelques prérequis techniques facilitent grandement la démarche.

- **Liste complète de vos composants** : référence exacte du CPU, du GPU, du nombre de SSD/HDD, de ventilateurs et de RGB installés ou prévus.
- **Fiche technique du GPU** : le TGP (Total Graphics Power) ou TDP annoncé par le fabricant, disponible sur la page produit NVIDIA ou AMD.
- **Python 3.10 ou supérieur** installé si vous voulez utiliser le script de calcul fourni plus bas (facultatif mais recommandé).
- **Pilotes GPU à jour** : GeForce Game Ready Driver ou AMD Adrenalin récents pour que`nvidia-smi` ou les outils AMD renvoient des mesures fiables.
- **HWiNFO64 (version 8.x)** pour lire la consommation réelle carte mère/CPU/GPU en temps réel.
- **Tournevis cruciforme PH2** et un bracelet antistatique pour l’installation physique de l’alimentation.
- **Multimètre** (optionnel) si vous voulez vérifier les tensions 12V/5V/3.3V après installation.
- **Budget indicatif** : comptez entre 90 € et 250 € selon la puissance et la certification visées.

Comptez environ 60 à 90 minutes pour l’ensemble du processus : calcul du wattage, choix du modèle, commande, puis installation et test. Si vous changez uniquement l’alimentation sur un PC déjà monté, l’installation seule prend 20 à 30 minutes.

## Comprendre le TGP et la consommation réelle des GPU RTX 50 et RX 9000

Le point de départ de tout calcul d’alimentation PC gamer, c’est la consommation de la carte graphique. C’est de loin le composant le plus gourmand d’une configuration moderne, et les GPU de génération RTX 50 ont fait grimper les besoins de façon significative par rapport aux générations précédentes. Le TGP (Total Graphics Power) est le chiffre annoncé par le fabricant pour la consommation maximale soutenue de la carte, ventilateurs compris. C’est ce chiffre, et non un TDP théorique, qu’il faut utiliser pour votre calcul.

Le tableau ci-dessous résume le TGP officiel des GPU RTX 50 et de l’équivalent AMD RX 9070 XT, ainsi que la puissance d’alimentation recommandée pour chaque carte, telle que documentée par les fabricants et les guides d’assemblage 2026.

| Carte graphique | TGP / TDP officiel | Connecteur | PSU minimum | PSU recommandé | 
|---|---|---|---|---|
| RTX 5090 | 575 W | 12V-2×6 natif | 1000 W | 1200 W | 
| RTX 5080 | 360 W | 12V-2×6 natif | 750-800 W | 850 W | 
| RTX 5070 Ti | 300 W | 12V-2×6 ou 8-pin | 650 W | 750 W | 
| RTX 5070 | 250 W | 8-pin ou 12V-2×6 | 550 W | 650 W | 
| RX 9070 XT | 304 W | 8-pin x2 | 650 W | 750 W | 

Sur une RTX 5090, le TGP de 575 W représente environ 96 % de la capacité nominale du connecteur 12V-2×6, qui est évalué à 600 W. La marge est donc très faible, ce qui explique pourquoi un mauvais contact ou un câble mal inséré peut provoquer une surchauffe localisée. Sur les cartes moins puissantes comme la RTX 5070 ou la RX 9070 XT, la marge de sécurité électrique reste nettement plus confortable.

## Le connecteur 12V-2×6 et la norme ATX 3.1 : ce qui change

Depuis l’arrivée des RTX 50, la norme ATX 3.1 (associée à PCIe 5.1) s’est imposée comme le standard pour toute alimentation PC gamer haut de gamme. Elle intègre nativement un connecteur 16 broches appelé 12V-2×6, successeur direct du 12VHPWR utilisé sur les RTX 40. Ce connecteur, formalisé par un ECN du PCI-SIG, conserve la même géométrie mais améliore les tolérances de contact pour limiter les défauts d’insertion.

Le sujet est sensible car des cas de connecteurs fondus ont continué à être rapportés en 2025 et 2026 sur des RTX 5090, avec des points chauds pouvant atteindre environ 150 °C en cas de mauvais contact. Dans la grande majorité des cas documentés, le problème provient de trois causes précises : l’utilisation d’anciens câbles 12VHPWR avec un adaptateur vers 12V-2×6, l’usage d’adaptateurs coudés tiers non certifiés, ou une insertion incomplète du connecteur qui n’a pas été enfoncé jusqu’au clic. En réponse, plusieurs fabricants de cartes mères et de PSU (MSI, ASRock, entre autres) ont ajouté des embouts de couleur pour vérifier visuellement l’insertion complète, ainsi que des capteurs de température ou de courant sur certains câbles haut de gamme.

- Utilisez exclusivement un câble natif 12V-2×6 fourni avec votre alimentation ATX 3.1, jamais un adaptateur multiple.
- Évitez les coudes trop serrés à moins de 3,5 cm du connecteur, cela stresse mécaniquement les broches.
- Vérifiez que le connecteur est enfoncé jusqu’au clic audible des deux côtés.
- Si votre PSU n’a pas de câble 12V-2×6 natif et que vous devez utiliser un adaptateur, choisissez uniquement celui fourni par NVIDIA ou le fabricant de votre carte graphique.

## Lire une fiche technique d’alimentation : rails, connecteurs et courbe d’efficacité

Avant d’entrer dans le calcul étape par étape, il est utile de savoir décoder une fiche technique de PSU, car les fabricants n’affichent pas toujours l’information la plus importante en premier. Trois éléments méritent une attention particulière.

