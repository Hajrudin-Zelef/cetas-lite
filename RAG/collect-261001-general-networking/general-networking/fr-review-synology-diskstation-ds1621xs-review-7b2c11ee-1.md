---
id: collect-261001-general-networking/general-networking/fr-review-synology-diskstation-ds1621xs-review-7b2c11ee-1
title: "fr-review-synology-diskstation-ds1621xs-review-7b2c11ee"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel", "valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-synology-diskstation-ds1621xs-review-7b2c11ee.md
source_anchor: ""
source_lines: [1, 55]
sha256: eed657a27a09a0809c2c65410c440728a00da2f089147192cda1a975dce85388
---

# fr-review-synology-diskstation-ds1621xs-review-7b2c11ee

Le Synology DiskStation DS1621xs+ a été lancé mi-septembre dernier comme une mise à jour de la gamme de NAS professionnels haut de gamme de Synology, au format de bureau. Le DS1621xs+ offre six baies pour disques 3.5 pouces, extensibles à 16 baies au total grâce à deux modules JBOD DX517 à 5 baies.
Sous son capot se cache un processeur Intel Xeon D-1527 quadricœur cadencé à 2.2 GHz, 8 Go de mémoire DDR4 ECC SODIMM (extensible jusqu'à 32 Go) et deux baies SSD M.2 2280 NVMe pour la mise en cache. Côté connectivité, le DS1621xs+ offre plusieurs ports intégrés, dont deux ports Gigabit Ethernet et un port RJ-45 10 GbE . Pour les entreprises ayant des besoins encore plus importants, ce NAS propose également un emplacement d'extension PCIe pouvant accueillir des SSD supplémentaires, des modules réseau, ou les deux avec la carte combinée Synology.
Bien que le DS1621xs+ utilise DiskStation Manager (DSM) comme environnement d'exploitation (en plus du système de fichiers Btrfs), sa puissance permet une meilleure évolutivité de certaines fonctionnalités avancées pour les environnements mutualisés. Synology considère ce NAS comme une solution idéale pour les entreprises souhaitant tirer parti de sa suite de services cloud, de collaboration et de virtualisation.
Le Synology DS1621xs+ est livré avec une garantie de 5 ans et a un prix de détail suggéré de 1599.99 $, bien qu'il coûte environ 1550 $ sans accessoires au moment de cette évaluation.
Spécifications du Synology DS1621xs +
| Processeur | Intel Xeon D-1527 4 cœurs 2.2 GHz, Turbo Boost jusqu'à 2.7 GHz | 
| Moteur de chiffrement matériel | Oui (AES-NI) | 
| Mémoire | 8 Go DDR4 ECC SODIMM (extensible jusqu'à 32 Go) | 
| Type de lecteur compatible | 6 disques durs/SSD SATA 3.5″ ou 2.5″ (disques non inclus) 2 x SSD M.2 2280 NVMe (disques non inclus) | 
| Disque remplaçable à chaud | Oui | 
| Port externe | 3 ports USB 3.0 2 ports eSATA | 
| Taille (HxLxP) | 166 x 282 x 243 mm | 
| Poids | 5.3 kg | 
| LAN | 2 ports RJ-1 45 GbE 1 port RJ-10 45 GbE | 
| Réveil sur LAN/WAN | Oui | 
| Emplacement PCIe 3.0 | 1 emplacement x8 à 8 voies Prise en charge des cartes d'interface réseau hautes performances | 
| Mise sous / hors tension programmée | Oui | 
| Ventilateur du système | 2 (92x92x25mm) | 
| Tension d'alimentation d'entrée CA | 100 V à 240 V CA | 
| Fréquence de puissance | 50/60 Hz, monophasé | 
| Température de fonctionnement | 0 ° C à 40 ° C (° F à 32 104 ° F) | 
| La température de stockage | -20 ° C à 60 ° C (° F à -5 140 ° F) | 
| Humidité relative | 5% à 95% RH | 
| Altitude de fonctionnement maximale | 5,000 m (16,400 ft) | 
Synology DS1621xs+ Conception et construction
Le Synology DiskStation DS1621xs+ ressemble au reste des périphériques NAS au format tour de la société. Le châssis est noir mat et composé de métal et de plastique. À l'avant se trouvent les baies de lecteur de 3.5 pouces qui incluent la possibilité de verrouiller et un voyant d'état du lecteur en haut. Le haut de l'appareil de gauche à droite comporte un indicateur d'état, un indicateur d'alerte, un bouton d'alimentation et les voyants LAN. Il y a aussi un port USB 3.0 dans le coin inférieur.
En retournant le NAS vers l'arrière, nous voyons les ventilateurs qui occupent une grande partie de l'immobilier. À gauche, se trouve le port d'alimentation et la fente de sécurité Kensington. À droite, se trouve le slot d'extension PCIe. En bas, de gauche à droite, se trouvent deux ports USB 3.0, un port RJ-10 45GbE, deux ports RJ-1 45GbE, un bouton de réinitialisation et deux ports d'extension.
Le cache M.2 se trouve à un endroit un peu différent, à l'intérieur de l'appareil, sur le côté gauche de la baie de lecteur la plus éloignée.
Performances du Synology DS1621xs+
Pour nos tests, nous avons configuré le Synology DiskStation en RAID 6, en iSCSI et en SMB. Chaque test a été réalisé avec le cache SSD activé et désactivé. Nous avons utilisé un disque dur WD Red de 14 To comme disque dur principal et deux SSD Synology SNV3400 de 400 Go comme cache SSD . Il ne s'agit pas d'une comparaison des performances avec ou sans cache SSD : ce dernier offre de meilleures performances. Ces résultats reflètent les performances attendues, que l'utilisateur choisisse ou non de l'utiliser.
Analyse synthétique de la charge de travail d'entreprise
Notre processus de référence de stockage partagé et de disque dur d'entreprise préconditionne chaque disque dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads avec une file d'attente exceptionnelle de 16 par thread, puis testé à intervalles définis dans plusieurs profils de profondeur de thread/file d'attente pour montrer les performances en cas d'utilisation légère et intensive. Étant donné que les disques durs atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4k et 8k 70/30, qui est couramment utilisée pour les disques d'entreprise.
- 4K
- 
  - 100 % de lecture ou 100 % d'écriture
  - 100% 4K
- 8K 70/30
  - 70 % de lecture, 30 % d'écriture
  - 100% 8K
- 8K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 8K
- 128K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 128K
Dans la première de nos charges de travail d'entreprise, nous avons mesuré un long échantillon de performances 4K aléatoires avec une activité d'écriture à 100 % et de lecture à 100 %. En ce qui concerne les IOPS, le Synology DiskStation DS1621xs+ a pu atteindre 4,540 1,880 IOPS en lecture et 179 1,940 IOPS en écriture en iSCSI et 98,504 IOPS en lecture et 67,453 3,425 IOPS en écriture en SMB. En activant le cache, nous avons vu les chiffres passer à 22,537 XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture dans iSCSI et SMB nous ont donné XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture.
Pour une latence moyenne de 4K, le DS1621xs+ nous a donné des scores iSCSI de 56.38 ms en lecture et 136.15 ms en écriture et des vitesses SMB de 1,426.81 131.9 ms en lecture et 2.6 ms en écriture. Avec le cache, les chiffres sont tombés à 3.8 ms en lecture et 74.7 ms en écriture pour iSCSI et 11.4 ms en lecture et XNUMX ms en écriture pour SMB.
Avec une latence maximale de 4K, le Synology a vu 1,161.9 3,901.2 ms en lecture et 2,729 3,259 ms en écriture pour iSCSI et SMB a atteint 848.7 342.8 ms en lecture et 943 38 ms en écriture. La fonction de cache a ramené les chiffres à XNUMX ms en lecture et XNUMX ms en écriture pour iSCSI et XNUMX ms en lecture et XNUMX ms en écriture pour SMB.
L'écart standard 4K a montré que le NAS était capable d'atteindre 61.8 ms en lecture et 210.3 ms en écriture en iSCSI et 110.9 ms en lecture et 223.5 ms en écriture en SMB. Une fois la mise en cache engagée, nous avons vu 5.5 ms de lecture et 7.2 ms d'écriture pour iSCSI et 65.9 ms de lecture et 1.3 ms d'écriture pour SMB.
