---
id: collect-261001-general-networking/general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240-1
title: "fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240"
domain: general-networking
role: reference
task: reference
actors: ["Meta", "Samsung"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240.md
source_anchor: ""
source_lines: [1, 48]
sha256: 6b220d267b0888acad286e0cd528d73a79e816a7789a3886a4a33e4cfe3215fb
---

# fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240

Sponsorisé par Western Digital
Le SSD Western Digital Ultrastar® DC SN861 est conçu pour répondre aux besoins de haute performance des centres de données hyperscale et des environnements d'entreprise. Le SN861 prend en charge une interface PCIe® Gen5 et est disponible sous différents facteurs de forme, notamment U.2 et E1.S, lui permettant de s'adapter à plusieurs scénarios de déploiement. Ce n'est cependant pas aussi simple que de fabriquer le SN861 dans différents facteurs de forme ; Western Digital a judicieusement conçu l'ensemble des fonctionnalités du SN861 pour s'aligner sur ses marchés cibles.
L'interface Gen5 confère au SN861 des performances nettement supérieures à celles du SN655 de génération précédente . Les avantages de ce nouveau disque sont bien plus profonds, notamment grâce à des fonctionnalités telles que le placement flexible des données (FDP) au format E1.S. Le FDP réduit l'amplification d'écriture et optimise le placement des données. Le SN861 intègre des fonctions de sécurité avancées, comme la protection des données de bout en bout, le chiffrement AES-XTS et TCG OPAL 2.01. Le contrôleur contribue également à réduire la consommation d'énergie du SSD, avec une moyenne inférieure à 5 watts en veille. De plus, le disque prend en charge plusieurs normes, telles que NVMe® 2.0 et OCP Cloud Spec 2.0.
Bien que les fonctionnalités de sécurité et d'efficacité soient essentielles, chaque actualisation générationnelle inclut une augmentation significative des performances, et le SN861 n'est pas différent. Le disque offre des vitesses de lecture séquentielle allant jusqu'à 13,700 3.3 Mo/s et des IOPS en lecture aléatoire jusqu'à 861 millions, essentielles pour les applications telles que l'IA/ML et l'analyse du Big Data. Les deux versions du SN20 consomment en moyenne 5 watts en fonctionnement et moins de 1 watts au repos. La puissance est réglable, il est donc facile d'ajuster le profil de puissance du disque en fonction de la charge de travail attendue. Les hyperscalers, par exemple, exécutent souvent leurs disques EXNUMX.S à des niveaux de consommation beaucoup plus faibles.
Il est intéressant de noter que même si les deux facteurs de forme du SN861 sont techniquement très similaires dans leur conception, Western Digital a optimisé chaque disque pour des charges de travail spécifiques. Dans la version E1.S, par exemple, cela signifie des fonctionnalités telles que FDP et des performances optimisées pour les charges de travail cloud. Le disque U.2, en revanche, trouvera sa place dans les charges de travail d'entreprise hautes performances et sans aucun doute dans les charges de travail émergentes comme l'IA qui peuvent bénéficier de l'augmentation massive des performances du disque.
EDSFF et FDP
FDP offre des avantages significatifs aux hyperscalers comme Meta en optimisant les performances et la fiabilité de leurs SSD dans des charges de travail telles que CacheLib. FDP réduit le facteur d'amplification d'écriture (WAF), ce qui entraîne des vitesses d'écriture améliorées et une durée de vie prolongée des SSD, ce qui est crucial pour gérer des tâches de traitement de données massives.
La technologie améliore l'organisation des données en regroupant intelligemment les données similaires, en minimisant le surprovisionnement et en réduisant le besoin d'un garbage collection intensif. FDP prend également en charge plusieurs espaces de noms, garantissant des performances cohérentes sur différentes charges de travail. Cette optimisation améliore les performances et l'endurance des applications et réduit considérablement le coût total de possession (TCO) des infrastructures de stockage à grande échelle.
La prise en charge de FDP dans la version E1.S de l'Ultrastar SN861 affirme que le disque est prêt à répondre aux besoins des hyperscalers, mais FDP n'est qu'une partie de l'équation. La version E1.S du disque doit répondre aux exigences de performances à grande échelle, en particulier la qualité de service autour des performances de lecture.
U.2 pour les entreprises
Aussi passionnant que soit le disque E1.S pour les cas d’utilisation à grande échelle, le U.2 SN861 est le disque que la plupart des entreprises adopteront. Nous soumettons le lecteur à une série de tests pour mesurer les performances globales dans notre suite de tests standard.
Fiche technique du SSD Western Digital Ultrastar DC SN861
|  | 1.60TB | 1.92TB | 3.20TB | 3.84TB | 6.40TB | 7.68TB | 
|---|---|---|---|---|---|---|
| Endurance | 3 DWPD | 1 DWPD | 3 DWPD | 1 DWPD | 3 DWPD | 1 DWPD | 
| Sécurité |  |  |  |  |  |  | 
| Facteur de forme |  |  |  |  |  |  | 
| Interface |  |  |  |  |  |  | 
| Spécification NVMe |  |  |  |  |  |  | 
| Performance (projetée) | 1.60TB | 1.92TB | 3.20TB | 3.84TB | 6.40TB | 7.68TB | 
| Débit de lecture (max Mo/s, séquence 128 Ko) | 13,700 | 13,700 | 13,700 | 13,700 | 13,700 | 13,700 | 
| Débit d'écriture (max Go/s, séquence 256 Ko) | 3,600 | 3,600 | 7,200 | 7,200 | 7,500 | 7,500 | 
| Lire IOPS (max, Rnd 4KiB) | 2,100K | 2,100K | 3,300K | 3,300K | 3,300K | 3,300K | 
| Ecrire IOPS (max, Rnd 4KiB) | 350K | 165K | 665K | 330K | 800K | 430K | 
| Latence de lecture (µS) | 65 | 65 | 65 | 65 | 65 | 65 | 
| Latence d'écriture (µS) | 8 | 8 | 8 | 8 | 8 | 8 | 
| Fiabilité |  |  |  |  |  |  | 
| MTTF (heures, projeté) |  |  |  |  |  |  | 
| Taux d'erreurs sur les bits non corrigibles (UBER) |  |  |  |  |  |  | 
| Taux de défaillance annualisé (AFR, projeté) |  |  |  |  |  |  | 
| Garantie limitée (années) |  |  |  |  |  |  | 
| Gestion de l'alimentation (projetée) |  |  |  |  |  |  | 
| Exigence (DC, +/- 10%) |  |  |  |  |  |  | 
| Modes de fonctionnement (moyenne, max) |  |  |  |  |  |  | 
| Inactif (moyen) |  |  |  |  |  |  | 
| Grandeur physique |  |  |  |  |  |  | 
| hauteur z (mm) |  |  |  |  |  |  | 
| Dimensions (largeur x longueur, mm) |  |  |  |  |  |  | 
| Environnemental |  |  |  |  |  |  | 
| Température de fonctionnement (ambiante) |  |  |  |  |  |  | 
| Température hors fonctionnement |  |  |  |  |  |  | 
Pour mesurer les performances des SSD NVMe® Gen5 d'entreprise utilisés dans cette comparaison, nous avons utilisé la suite de tests fio pour les charges de travail extrêmes et Vdbench pour les charges de travail mixtes. Le script fio utilisé est un outil automatisé configuré pour préconditionner et tester légèrement les disques de manière uniforme ; il est disponible ici sur GitHub . Nous l'avons utilisé pour effectuer des tests de lecture et d'écriture séquentielles de 256 Ko afin de mesurer la bande passante maximale, ainsi que des tests de lecture et d'écriture aléatoires de 4 Ko afin de mesurer le débit maximal.
| Débit de pointe et bande passante | Western Digital SN861 7.68 To | KIOXIA CM7-R 7.68 To | Samsung PM1743 7.68 To | Samsung PM9A3 7.68 To | 
| Lecture séquentielle 256K (1T/64Q) | 13,283MB / s | 12,092MB / s | 14,495MB / s | 6,751MB / s | 
| 256 Ko d'écriture séquentielle (1T/64Q) | 7,696MB / s | 5,796MB / s | 6,052MB / s | 4,055MB / s | 
| Lecture aléatoire 4K (8T/32Q) | 2,108,065 IOPS | 1,963,066 IOPS | 1,900,838 IOPS | 1,068,508 IOPS | 
| Écriture aléatoire 4K (8T/32Q) | 473,658 IOPS | 301,061 IOPS | 319,758 IOPS | 206,660 IOPS | 
Lorsque nous examinons les principaux chiffres de performances du Western Digital SN861, il fait bon usage de son interface Gen5. En lecture séquentielle, il mesurait 13.3 Go/s, ce qui arrivait en deuxième position par rapport au Samsung PM1743, qui mesurait 14.5 Go/s. En écriture séquentielle, le SN861 est arrivé en premier, balayant les deux autres modèles Gen5 comparables, avec une vitesse de 7.7 Go/s, avec 6.1 Go/s du Samsung PM1743 comme prochain plus proche.
