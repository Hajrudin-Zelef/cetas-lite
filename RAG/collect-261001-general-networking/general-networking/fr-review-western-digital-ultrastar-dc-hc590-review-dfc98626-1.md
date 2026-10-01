---
id: collect-261001-general-networking/general-networking/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626-1
title: "fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Meta", "Nvidia"]
dates: []
keywords: ["amd", "benchmarks", "deepseek", "gpu", "llama", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626.md
source_anchor: ""
source_lines: [1, 43]
sha256: 28da15aaf96a0b361c8e69ebb67838502c87c529b8d57d0c985cd5a89c773839
---

# fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626

Le Western Digital Ultrastar DC HC590 est le dernier disque dur d'entreprise haute capacité de WD. Il offre jusqu'à 26 To de stockage grâce à l'enregistrement magnétique conventionnel (CMR) dans un format standard de 3.5 pouces. Il s'agit du premier disque dur CMR à utiliser une conception à 11 plateaux, offrant une capacité brute supérieure sans nécessiter de modifications de l'infrastructure existante. Cela permet aux environnements d'entreprise et hyperscale d'adapter la densité de stockage avec le même encombrement physique et thermique, tout en préservant la compatibilité avec les conceptions de rack et les boîtiers actuels. En fin de compte, le HC590 est conçu pour les charges de travail gourmandes en stockage et axées sur la lecture, où la fiabilité et la rentabilité priment sur les performances optimales.
Caractéristiques du WD Ultrastar DC HC590
Le HC590 présente des caractéristiques de fiabilité familières à cette catégorie de disques, conçues pour maintenir des performances constantes dans des environnements denses et sujets aux vibrations. Il associe notamment la technologie PMR assistée par énergie (ePMR) à son architecture OptiNAND, qui intègre de la mémoire flash intégrée (iNAND) pour décharger des tâches internes comme la gestion des métadonnées. Ensemble, ces technologies contribuent à accroître la densité surfacique et à réduire la sollicitation des supports magnétiques, améliorant ainsi l'efficacité et le débit lors d'opérations séquentielles à grande échelle ou gourmandes en métadonnées.
Pour garantir des performances stables dans les racks de serveurs denses, le lecteur est également équipé de la technologie RVS (Rotational Vibration Safeguard), qui utilise deux capteurs pour détecter et compenser les vibrations environnementales. Il intègre également la technologie DFH (Dynamic Fly Height), qui ajuste la tête de lecture/écriture à la volée pour une meilleure précision de contact. Le boîtier scellé à l'hélium réduit la résistance interne et la consommation d'énergie par rapport aux solutions à air comprimé.
ArmorCache est également disponible, permettant aux utilisateurs de sélectionner le mode cache en écriture désactivé (WCD) pour améliorer les performances d'écriture aléatoire ou le mode cache en écriture activé (WCE) pour protéger les données du cache en cas de panne de courant. Ceci est particulièrement utile dans les systèmes où la politique de mise en cache est régie par des exigences d'intégrité des données ou de débit.
Côté fiabilité, le disque affiche un MTBF de 2.5 millions d'heures et bénéficie d'une garantie limitée de 5 ans, deux caractéristiques typiques des disques durs d'entreprise. Le HC590 est disponible en versions SAS et SATA, avec des capacités de 24 To et 26 To. Pour ce test, nous avons testé le modèle SATA 26 To.
Spécifications du disque dur WD Ultrastar DC HC590 26 To
| Spécifications | DÉTAILS | 
| Modèle | WD Ultrastar DC HC590 | 
| Capacités | 26TB | 
| Facteur de forme | 3.5 pouce | 
| Interface | SAS 12 Gb/s ou SATA 6 Gb/s | 
| Technologie d'enregistrement | Enregistrement magnétique conventionnel (CMR) | 
| Nombre de plateaux | 11 | 
| scellé à l'hélium | Oui | 
| Technologies Avancées | ePMR, OptiNAND, hauteur de vol dynamique, protection contre les vibrations rotationnelles (RVS) | 
| Options de cache | ArmorCache : WCE (protection des données) / WCD (performances d'écriture) | 
| Temps moyen entre les pannes (MTBF) | 2.5 millions d'heures | 
| Caractéristiques de sécurité | Effacement sécurisé (SE) | 
| Dimensions (L x l x H) | 5.776 "x 4.000" x 1.028 " | 
| Poids | 1.47 lb (environ 667 g) | 
| Garantie | 5-Year Limited Warranty | 
| Référence du modèle | WUH722626AL5204 | 
Performances du WD Ultrastar DC HC590
Avant de plonger dans les benchmarks, voici une liste de disques durs de capacité comparable utilisés pour comparer les performances du disque WD Ultrastar DC HC590 26 To.
Voici le banc d'essai haute performance que nous avons utilisé pour l'analyse comparative du stockage :
- CPU: AMD Ryzen 7 9800X3D
- Carte mère : Asus ROG Crosshair X870E Hero
- RAM : G.SKILL Trident Z5 Royal Series DDR5-6000 (2 x 16 Go)
- GPU: NVIDIA GeForce RTX 4090
- Système d'exploitation : Windows 11 Pro, Ubuntu 24.10 Desktop
Performance synthétique de pointe
Le test FIO est un outil d'analyse comparative flexible et puissant permettant de mesurer les performances des périphériques de stockage, notamment les SSD et les disques durs. Il évalue des indicateurs tels que la bande passante, les IOPS (opérations d'entrée/sortie par seconde) et la latence sous différentes charges de travail, comme les opérations de lecture/écriture séquentielles et aléatoires. Ce test permet d'évaluer les performances maximales des systèmes de stockage, ce qui le rend utile pour comparer différents périphériques ou configurations. Nous avons mesuré les performances maximales en rafale pour ce test, en limitant la charge de travail à 10 Go sur tous les disques durs.
Lors de nos tests FIO, le WD DC HC590 26 To a enregistré des performances séquentielles légèrement inférieures à celles de ses concurrents, mais est resté dans une fourchette proche. En lecture séquentielle à 128 268 secondes, il a atteint 6 Mo/s, soit environ 24 % de moins que le Seagate Exos X5 et environ 280 % de moins que le WD Gold. En écriture séquentielle, il a atteint 2 Mo/s, soit seulement XNUMX % de moins que le WD Gold, et se rapproche de celui du Seagate.
En lecture aléatoire 4K, le HC590 a atteint 198 IOPS, soit environ 7 % de moins que le Seagate et le WD Gold. Ses performances en écriture aléatoire 4K étaient supérieures, avec 663 IOPS, soit seulement 11 % de moins que le Seagate, tout en surpassant le WD Red Pro de plus de 50 % et le WD Gold de près de 2 %.
Dans l’ensemble, le HC590 maintient un débit séquentiel solide et fait preuve de solidité dans les performances d’écriture aléatoire avec seulement des différences modestes par rapport aux autres disques haute capacité testés.
| Test FIO (un débit MB/s/IOPS plus élevé est meilleur) | Lecture séquentielle 128K (1T/64Q) | Écriture séquentielle 128 Ko (1T/64Q) | Lecture 4K aléatoire (16T/32Q) | Écriture 4K aléatoire (16T/32Q) | 
| Seagate x24 24 To | 285 Mo/s (29.42 ms) | 285 Mo/s (29.42 ms) | 210 152.03 IOPS (XNUMX ms) | 749 42.70 IOPS (XNUMX ms) | 
| WD Or 24 To | 283 Mo/s (29.66 ms) | 286 Mo/s (29.36 ms) | 214 148.98 IOPS (XNUMX ms) | 651 49.11 IOPS (XNUMX ms) | 
| WD Red Pro 22 To | 271 Mo/s (31.00 ms) | 276 Mo/s (30.37 ms) | 214 149.17 IOPS (XNUMX ms) | 421 75.92 IOPS (XNUMX ms) | 
| WD Ultrastar DC HC590 26 To | 268 Mo/s (31.28 ms) | 280 Mo/s (30.00 ms) | 198 161.18 IOPS (XNUMX ms) | 663 48.23 IOPS (XNUMX ms) | 
Temps de chargement moyen du LLM
Le test de temps de chargement moyen des LLM a évalué les temps de chargement de trois LLM différents : DeepSeek R1 7B, Meta Llama 3.2 11B et DeepSeek R1 32B. Chaque modèle a été testé 10 fois et le temps de chargement moyen a été calculé. Ce test mesure la capacité du lecteur à charger rapidement des modèles de langage volumineux (LLM) en mémoire. Les temps de chargement des LLM sont essentiels pour les tâches liées à l'IA, notamment pour l'inférence en temps réel et le traitement de grands ensembles de données. Un chargement plus rapide permet au modèle de traiter rapidement les données, améliorant ainsi la réactivité de l'IA et réduisant les temps d'attente.
