---
id: collect-261001-general-networking/general-networking/fr-review-synology-e10m20-t1-review-a0ab7b39-1
title: "fr-review-synology-e10m20-t1-review-a0ab7b39"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "ethernet"]
source: docs/RAG/collect-261001-general-networking/fr-review-synology-e10m20-t1-review-a0ab7b39.md
source_anchor: ""
source_lines: [1, 51]
sha256: 6183907e3838d97dd97b8857146247e11909e457d5e235c72ab63e3af5bd77df
---

# fr-review-synology-e10m20-t1-review-a0ab7b39

Synology propose une belle gamme de périphériques NAS qui sont parfaits pour les TPE et les PME. Le seul inconvénient est que les NAS de milieu de gamme ne sont généralement pas conçus pour des performances et une vitesse de réseau maximales. Grâce à une nouvelle carte d'extension, ils n'ont plus besoin de l'être. Avec l'introduction du Synology E10M20-T1 AIC, les utilisateurs peuvent obtenir des E/S et une bande passante beaucoup plus élevées avec une seule carte.
Pour la bande passante, le Synology E10M20-T1 est livré avec un port 10GbE, augmentant considérablement les E/S d'un NAS qui a GbE intégré. En plus des performances réseau améliorées, la carte est livrée avec deux emplacements SSD NVMe pour les disques M.2 (facteurs de forme 2280 et 22110). Cela donne aux utilisateurs des performances d'E/S élevées et ajoute deux emplacements NVMe pour le cache SSD sans renoncer aux baies de lecteur. De cette façon, les utilisateurs peuvent charger leur NAS avec des disques durs haute capacité et toujours disposer d'un cache SSD. Cette carte est également particulièrement utile dans certains des anciens systèmes NAS de Synology, qui n'ont pas d'emplacements SSD intégrés.
La carte Synology E10M20-T1 est disponible dès aujourd'hui au prix de 250 $ . Pour ceux qui le souhaitent, nous proposons également une présentation vidéo de la carte.
Spécifications du Synology E10M20-T1
| Généralités |  | 
| Interface de bus hôte | PCIe 3.0 x8 | 
| Hauteur du support | Profil bas et pleine hauteur | 
| Taille (Hauteur x Largeur x Profondeur) | 71.75 mm x 200.05 mm x 17.70 mm | 
| Température de fonctionnement | 0 ° C à 40 ° C (° F à 32 104 ° F) | 
| Température de stockage | -20 ° C à 60 ° C (° F à -5 140 ° F) | 
| Humidité relative | 5% à 95% RH | 
| Garantie | 5 ans | 
| Stockage |  | 
| Interface de stockage | PCIe NVMe | 
| Facteur de forme pris en charge | 22110/2280 | 
| Type et quantité de connecteur | Clé M, 2 emplacements | 
| Réseau |  | 
| Conformité aux spécifications IEEE | Ethernet IEEE 802.3an 10 Gb/s Ethernet IEEE 802.3bz 2.5 Gb/s / 5 Gb/s Gigabit Ethernet IEEE 802.3ab Ethernet rapide IEEE 802.3u Contrôle de flux IEEE 802.3x | 
| Taux de transfert de données | 10 Gbps | 
| Mode de fonctionnement réseau | Full Duplex | 
| Caractéristiques supportées | Cadre géant de 9 Ko Déchargement de la somme de contrôle TCP/UDP/IP Négociation automatique entre 100 Mb/s, 1 Gb/s, 2.5 Gb/s, 5 Gb/s et 10 Gb/s | 
| Compatibilité |  | 
| Modèles appliqués SSD NVMe | Série SA : SA3600, SA3400 Série 20: RS820RP +, RS820 + Série 19 : DS2419+, DS1819+ Série 18 : RS2818RP+, DS3018xs, DS1618+ | 
Concevoir et construire
Le Synology E10M20-T1 est un AIC HHFL qui s'adaptera à certains modèles de Synology NAS. D'un côté se trouvent des dissipateurs thermiques qui s'étendent sur toute la longueur de la carte.
Le retrait du dissipateur thermique avec quatre vis à l'arrière donne un accès aux deux baies SSD M.2 NVMe. Dans l'ensemble, il est facile de sécuriser les lecteurs sur la carte.
Le verso de la carte est relativement spartiate. Synology comprend également un support de taille normale dans la boîte si cela est nécessaire, et des pastilles de contact thermiques pour les SSD.
Performances
Pour tester le Synology E10M20-T1, nous l'avons installé dans un Synology DS1819+ . Nous avons installé des disques durs WD Red de 14 To dans les baies . Pour le cache, nous avons utilisé des SSD Synology SNV3400-400G . Nous avons testé les disques en configurations iSCSI et CIFS en RAID 6, avec et sans cache.
Analyse synthétique de la charge de travail d'entreprise
Notre processus de référence de disque dur d'entreprise préconditionne chaque disque dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads avec une file d'attente exceptionnelle de 16 par thread. Il est ensuite testé à des intervalles définis dans plusieurs profils de profondeur de thread/file d'attente pour montrer les performances en cas d'utilisation légère et intensive. Étant donné que les disques durs atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4K et 8K 70/30, qui est couramment utilisée pour les disques d'entreprise.
- 4K
  - 100 % de lecture ou 100 % d'écriture
  - 100% 4K
- 8K70/30
  - 70 % de lecture, 30 % d'écriture
  - 100% 8K
- 128K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 128K
Dans la première de nos charges de travail d'entreprise, nous avons mesuré un long échantillon de performances 4K aléatoires avec une activité d'écriture à 100 % et de lecture à 100 % pour obtenir nos principaux résultats. Pour CIFS, nous avons vu 170 lectures IPS et 1,461 4,075 écritures IOPS sans cache et 10,950 2,897 lectures IOPS et 1,502 20,021 écritures IOPS avec le cache actif. Pour iSCSI, nous avons vu 22,439 XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture sans cache et en tirant parti du cache dans l'AIC, nous avons vu XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture.
Avec une latence moyenne de 4K, CIFS nous a donné 1,497 176 ms en lecture et 63 ms en écriture sans le cache, puis en l'activant, il est tombé à 23 ms en lecture et 88 ms en écriture. iSCSI a vu 170 ms en lecture et 12.8 ms en écriture, puis en activant le cache, il est tombé à 11.4 ms en lecture et XNUMX ms en écriture.
Ensuite, la latence maximale de 4K. Ici, CIFS a atteint 4,476 3,360 ms en lecture et 339 45 ms en écriture sans le cache ; en tirant parti de l'AIC, les chiffres sont tombés à 1,051 ms en lecture et 6,131 ms en écriture. iSCSI avait 11,951 171 ms en lecture et XNUMX XNUMX ms en écriture sans le cache, et avec lui la latence de lecture est passée à XNUMX XNUMX ms mais la latence d'écriture est tombée à XNUMX ms.
Notre dernier test 4K est l'écart type. Ici, sans le cache, CIFS nous a donné 228 ms en lecture et 288 ms en écriture, avec une latence d'activation du cache réduite à 7.3 ms en lecture et 2 ms en écriture. Pour iSCSI, nous avons de nouveau constaté un pic au lieu d'une baisse des lectures, passant de 69 ms sans cache à 196 ms en cache. Les écritures ont montré une amélioration, passant de 282 ms à 16 ms.
Notre prochain benchmark mesure 100 % de débit séquentiel 8K avec une charge 16T16Q dans des opérations de lecture à 100 % et d'écriture à 100 %. Ici, la configuration CIFS sans le cache avait 13,989 10,770 IOPS en lecture et 13,055 11,443 IOPS en écriture ; après l'activation du cache, les chiffres sont passés à 56,579 30,288 IOPS en lecture et 57,774 33,265 IOPS en écriture. Avec iSCSI, nous avons vu XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS sans le cache activé, avec celui-ci activé, nous avons vu les performances augmenter légèrement pour atteindre XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture.
