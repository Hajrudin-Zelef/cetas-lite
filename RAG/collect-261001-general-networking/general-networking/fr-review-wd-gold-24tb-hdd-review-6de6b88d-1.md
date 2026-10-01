---
id: collect-261001-general-networking/general-networking/fr-review-wd-gold-24tb-hdd-review-6de6b88d-1
title: "fr-review-wd-gold-24tb-hdd-review-6de6b88d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-wd-gold-24tb-hdd-review-6de6b88d.md
source_anchor: ""
source_lines: [1, 72]
sha256: 7cd71df29f783c07f68d72c11b4d59b3e081db69ca7d460692a82a792981a91c
---

# fr-review-wd-gold-24tb-hdd-review-6de6b88d

En novembre 2023, WD a lancé le disque dur Gold Enterprise de 24 To, offrant 2 To de plus que son précédent modèle de 22 To. Bien que cette mise à jour puisse paraître modeste, elle représente une amélioration significative de la capacité de stockage, notamment en termes de densité de stockage dans les configurations NAS et serveurs de grande envergure. Suite à notre premier test HL15 , où nous avons mis à l'épreuve quinze disques WD Gold de 24 To, nous sommes maintenant prêts à réévaluer ces disques au sein d'une configuration NAS standard à 8 disques.
La série WD Gold est conçue pour répondre aux divers besoins des entreprises, offrant une gamme complète de disques durs de classe entreprise de 1 To à 24 To. Ces disques sont conçus pour résister aux environnements de stockage exigeants, offrant un MTBF allant jusqu'à 2.5 millions d'heures, une technologie de protection contre les vibrations et une consommation d'énergie réduite grâce à la technologie HelioSeal pour les modèles de plus de 12 To.
Caractéristiques du WD Gold 24 To
WD a choisi de commercialiser ces disques haute capacité exclusivement avec une interface SATA, poursuivant ainsi la tradition de leur conception CMR remplie d'hélium à 7200 24 tr/min, un incontournable de leur gamme depuis plusieurs années. Il est intéressant de noter que le nouveau modèle Gold de 0.3 To réduit légèrement la consommation d'énergie, en utilisant 0.2 watt de moins en fonctionnement et 22 watt de moins en mode veille par rapport à son prédécesseur de XNUMX To.
Alors que la série Gold HDD couvre des capacités de 1 To à 24 To, il est à noter que la fonctionnalité OptiNAND, y compris l'innovant Armor Cache compatible OptiNAND, est exclusive aux modèles de 20 To et plus. Cette technologie met en synergie les avantages des modes activés et désactivés du cache en écriture, offrant aux utilisateurs un mélange de performances et de protection des données sans avoir à choisir entre les deux.
De plus, la série WD Gold offre flexibilité et évolutivité aux entreprises, permettant une configuration de stockage personnalisée capable de gérer des charges de travail d'applications intensives et de prendre en charge jusqu'à 550 To de données par an. Cela fait du disque dur WD Gold 24 To une option intéressante pour les entreprises cherchant à faire évoluer leur infrastructure de stockage de données tout en maintenant des niveaux élevés de fiabilité et de performances.
Gestionnaire de stockage et d'instantanés QNAP
Storage & Snapshots Manager est un outil polyvalent au sein de l'interface NAS QNAP qui permet aux utilisateurs de superviser et de gérer les différents aspects du stockage, y compris les configurations RAID, la santé des disques et la création d'instantanés pour la protection des données.
Le volet Stockage/Instantanés affiche un aperçu des volumes de stockage au sein du pool de stockage. Dans cet exemple, il met en évidence le volume système et plusieurs cibles iSCSI. Les barres indicatrices « Pourcentage utilisé » fournissent un aperçu rapide de l'utilisation de l'espace sur différents volumes.
La fenêtre Storage Pool 1 Management nous montre que tous les disques sont configurés dans une matrice RAID 6, ont une capacité totale de 130.91 To et sont en « bon » état. Cela indique que tout est prêt pour nos tests.
L'image ci-dessous montre la section Disques/VJBOD du gestionnaire de stockage et d'instantanés. Ici, les statuts « Bon » et « Prêt » confirment le déploiement de huit disques durs WD Gold haute capacité. Ce volet comprend également des informations instantanées telles que le numéro de modèle, la capacité du disque, la vitesse actuelle, le type de bus, etc., qui changent en fonction du disque sélectionné dans la liste. Ces informations sont cruciales pour maintenir des performances optimales et anticiper tout problème potentiel.
Spécifications du disque dur WD Gold 24 To
| Référence du modèle | WD241 CRISE | 
| Facteur de forme | 3.5 pouce | 
| Interface | SATA 6 Gb / s | 
| 512n / 512e secteurs utilisateur par lecteur4 | 512e | 
| Capacité formatée | 24TB | 
| Technologie OptiNAND | Oui | 
| conforme RoHS | Oui | 
| Performances |  | 
| Taux de transfert de données (max soutenu) | 298MB / s | 
| RPM | 7200 | 
| Cache | 512MB | 
| Gestion de l'énergie |  | 
| Besoins de puissance moyens (W) |  | 
| Efficacité | 6.8W | 
| Idle | 5.5W | 
| Indice d'efficacité énergétique (W/To, inactif) | 0.2 | 
| Fiabilité |  | 
| MTBF (heures, prévisionnel) | 2,500,000 | 
| Taux d'échec annualisé2 (AFR, %) | 0.35 | 
| Garantie limitée | 5 ans | 
| Environnemental |  | 
| Température de fonctionnement | 5 ° C à 60 ° C | 
| Température hors fonctionnement | -40 ° C à 70 ° C | 
| Choc (lecture/écriture)  En fonctionnement (demi-sinusoïdale, 2ms) | 40G / 40G | 
| Hors fonctionnement (demi-onde sinusoïdale, 2 ms) | 200G | 
| Acoustique (moyenne) |  | 
| Mode inactif | 20 dBA | 
| Mode de recherche | 32 dBA | 
| Dimensions physiques |  | 
| Hauteur (max) | 26.1mm | 
| Longueur (maximale) | 147.0mm | 
| Largeur (± 01 po) | 101.6mm | 
| Poids | 1.47 lb (67 kg)  ± 10% | 
Performances
Analyse synthétique de la charge de travail d'entreprise
Notre processus d'évaluation des disques durs d'entreprise conditionne chaque ensemble de disques à un état stable en utilisant la même charge de travail avec laquelle le périphérique sera testé. Cela implique une lourde charge de 16 threads et une file d'attente exceptionnelle de 16 par thread. L'appareil est ensuite testé à intervalles définis de plusieurs profils de profondeur de thread/file d'attente pour montrer ses performances dans des conditions d'utilisation légère et intensive. Étant donné que les disques durs atteignent rapidement leur niveau de performance nominal, seules les sections principales de chaque test sont représentées graphiquement.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse synthétique de la charge de travail d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées et les valeurs largement publiées, telles que la vitesse de lecture et d'écriture maximale de 4K et 8K 70/30, couramment utilisées pour les disques d'entreprise.
4K
- 100 pour cent de lecture ou 100 pour cent d'écriture
- 100 pour cent 4K
8K70/30
- 70 pour cent en lecture, 30 pour cent en écriture
- 100 pour cent 8K
8K (séquentiel)
- 100 pour cent de lecture ou 100 pour cent d'écriture
- 100 pour cent 8K
128K (séquentiel)
- 100 pour cent de lecture ou 100 pour cent d'écriture
- 100 pour cent 128K
4K 100 % lecture/écriture
Notre premier test mesure les performances aléatoires 4K. Dans ce test, le WD Gold a atteint 4,187 1,417 IOPS en lecture et 4,373 1,398 IOPS en écriture en SMB tout en affichant XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture iSCSI.
Pour une latence moyenne, le WD Gold a atteint 61.13 ms en lecture et 180.51 ms en écriture, et 58.52 ms en lecture/182.92 ms en écriture en SMB et iSCSI, respectivement.
À la latence maximale, les disques WD Gold 24 To ont enregistré 759.98 ms en lecture, 4,322.5 1,228.1 ms en écriture en SMB, 31,059 XNUMX ms en lecture et XNUMX XNUMX ms en écriture en iSCSI.
Concernant l'écart type, en SMB, le WD Gold a atteint 37.06 ms en lecture et 373.254 ms en écriture, tandis que l'iSCSI a enregistré 104.53 ms en lecture et 1,178 XNUMX ms en écriture.
8K 100 % lecture/écriture
