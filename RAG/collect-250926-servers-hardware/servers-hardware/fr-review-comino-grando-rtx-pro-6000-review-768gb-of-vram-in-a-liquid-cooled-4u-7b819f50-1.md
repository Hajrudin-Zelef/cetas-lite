---
id: collect-250926-servers-hardware/servers-hardware/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-1
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "attention", "blackwell", "gpu", "intel", "nvidia"]
source: docs/RAG/clean4/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [1, 45]
sha256: 1f436daece2842b553fac1a09992c6610e2a32eeeb3b1f621c4b124d8081c329
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

Comino nous a récemment envoyé la dernière version du Comino Grando pour test. Cette configuration inclut huit cartes NVIDIA RTX PRO 6000 Blackwell, chacune dotée de 96 Go de VRAM, pour un total de 768 Go de mémoire GPU. Nous avions testé le Comino en 2024, avec une configuration de six RTX 4090 offrant 144 Go de mémoire GPU, ainsi qu'une version équipée de NVIDIA H100 . Cette nouvelle version représente un bond générationnel significatif, tant en termes de capacité mémoire brute que de gamme de charges de travail prises en charge.
Le Grando est une plateforme 4U conçue spécifiquement pour résoudre le conflit crucial entre la puissance de calcul GPU haute densité et la gestion thermique. Alors que les châssis classiques à refroidissement par air s'effondrent sous la consommation continue de plus de 600 W des cartes graphiques professionnelles modernes, le Grando adopte une approche fondamentalement différente. Entièrement conçu autour d'une architecture à refroidissement liquide, il est capable de dissiper une chaleur massive de 6.5 kW en continu. Il ne s'agit pas d'une adaptation ou d'un ajout de dernière minute : le châssis entier, de l'agencement inversé de la carte mère à son système de raccords rapides à code couleur, a été pensé autour du circuit de refroidissement.
Le résultat est une plateforme capable d'accueillir huit GPU professionnels à pleine consommation (TDP) dans un châssis 4U unique, fonctionnant 24 h/24 et 7 j/7 dans des environnements ambiants de 3 à 38 °C, sans limitation thermique, sans les nuisances sonores d'un refroidissement par air à haut régime et sans compromettre la maintenance. Pour les entreprises déployant à grande échelle des charges de travail d'inférence IA, d'apprentissage automatique ou de simulation haute performance, le Grando offre une solution véritablement rare : un serveur qui ne vous oblige pas à choisir entre densité, dissipation thermique et fiabilité.
Spécifications du Comino Grando
Le tableau ci-dessous présente les spécifications physiques et les configurations matérielles prises en charge pour la plateforme Comino Grando.
| Spécifications / Fonctionnalités | Comino Grando | 
|---|---|
| Serveur Comino Grando et station de travail rackable |  | 
| Capacité frigorifique | 6.5 kW (Maximum 6 500 W à 20 °C de température d'air d'admission) | 
| Cartes mères | Jusqu'à EATX et EBB | 
| GPU (Serveur) | Jusqu'à 8 ; NVIDIA : RTX A6000, RTX 6000 ADA, RTX PRO 6000, A40, L40, L40S, A100, H100, H200 | 
| GPU (Station de travail rackable) | Jusqu'à 6 ; NVIDIA : 3090, 4090, 5080, 5090, RTX A6000, RTX 6000 ADA, RTX PRO 6000, A40, L40, L40S, A100, H100, H200 ; AMD : W7800, W7900 | 
| CPU | Jusqu'à 2 ; Processeur mono-socket : Intel Xeon W-2400/2500 et 3400/3500, Intel Xeon Scalable 4e génération, 5e génération, Xeon 6, AMD Threadripper PRO 5000WX, 7000WX, 9000WX, AMD EPYC 9004/9005 Double socket : Intel Xeon Scalable de 4e et 5e génération, Xeon 6, AMD EPYC 9004/9005 | 
| RAM | Jusqu'à 2TB | 
| Disques M2 | Jusqu'à 8 emplacements NVMe | 
| Stockage | Cages remplaçables à chaud sur le panneau arrière : jusqu’à 4 SSD remplaçables à chaud (4 x 7 mm ou 2 x 15 mm) et jusqu’à 4 autres (4 x 7 mm ou 2 x 15 mm) à la place d’une 4e alimentation ; Cage interne de 3.5″ jusqu’à 4 x 3.5″ ou 4 x 2.5″ 15 mm ou 12 x 2.5″ 7 mm ; Emplacements internes 2.5″ : jusqu’à 4 SSD 2.5″ de 7 mm | 
| Alimentation et tension de fonctionnement | Jusqu'à 4 alimentations CRPS remplaçables à chaud de 1 000 W à 180-264 V Jusqu'à 4 alimentations CRPS remplaçables à chaud de 1 000 W à 90-140 V Modes de redondance : 4+0, 3+1, 2+2 | 
| Le niveau de bruit | 39dB-70dB | 
| Lan | Jusqu'à 2 x 10 Gbit/s sur la carte mère et jusqu'à 400 Gbit/s en PCIe | 
| OS | Ubuntu / Windows 11 (Pro/Famille) / Serveur Windows | 
| Spécifications physiques et de refroidissement |  | 
| Refroidissement liquide | Processeur avec VRM et GPU avec GDDR et VRM | 
| Réservoir | Comino personnalisé 450ml avec pompes intégrées | 
| Ventilateurs | 3x Ultra High Flow 6200 tr/min (niveau sonore élevé) ou 3x Haut Débit 3000 tr/min (faible niveau sonore) | 
| Installation | Montage en rack 19 pouces ou utilisation autonome comme station de travail | 
| Espace rack requis | 4U | 
| Taille | 439 x 681 x 177 mm (sans poignées ni parties saillantes) | 
| Poids | 4 GPU : 49 kg (net), 67 kg (brut) 6 GPU : 52 kg (net), 70 kg (brut) 8 GPU : 55 kg (net), 72 kg (brut) | 
| Plage de températures de fonctionnement et de stockage | Conservation : -5 à 50 °C / 23 à 122 °F Température de fonctionnement : 3 à 38 °C / 38 à 100 °F | 
| Système de surveillance Comino (CMS) |  | 
| Marché | Carte de contrôle avec capteurs et logiciel pour la surveillance en temps réel | 
| Avantages clés | Système de refroidissement et surveillance du processeur/GPU, interface web, journal du système de refroidissement, surveillance centralisée pour les groupes de travail | 
| Capteurs et appareils connectés | Température (air et liquide de refroidissement), humidité relative, tension, débit du liquide de refroidissement, niveau de liquide de refroidissement dans le réservoir, ventilateurs, pompes, carte mère, écran et boutons | 
| Possibilités d'intégration | Mettez en place une surveillance via une API REST et envoyez les données des capteurs à un logiciel de surveillance (par exemple, Zabbix, Grafana) ou à des bases de données (par exemple, InfluxDB). | 
| Exigences techniques du CMS |  | 
| OS | Windows 11 / 10 Ubuntu 22.04/20.4 (Dépendance pour Ubuntu : le système cible doit avoir les utilitaires nvidia-smi et sensors installés) | 
| Navigateurs Web | Mozilla Firefox, Google Chrome, Chromium, Apple Safari, Microsoft Edge (Attention : Internet Explorer 11 n’est pas pris en charge) | 
| Disque dur | 300MB | 
| version du firmware du contrôleur | 1.0.6 ou plus récent | 
| Version PCB du contrôleur | 2.xx.xx | 
Conception, construction et densité des GPU
Agencement et déploiement du châssis
Le serveur Grando est un modèle d'optimisation de l'espace, avec des dimensions de 17.3 x 26.8 x 6.97 cm (4U). Contrairement aux serveurs traditionnels, il place l'arrière de la carte mère à l'avant du châssis, inversant ainsi l'agencement interne classique. Ceci garantit que les composants refroidis par air, tels que les modules de RAM et les VRM, bénéficient d'un flux d'air frais optimal avant que celui-ci n'atteigne le radiateur de refroidissement liquide situé à l'arrière.
Le châssis lui-même est fabriqué selon les mêmes normes rigoureuses, avec une construction en acier massif et une finition peinture époxy noire mate appliquée à l'intérieur comme à l'extérieur. Ce choix délibéré s'étend aux tubes, aux câbles, au radiateur et au masque de soudure du circuit imprimé, témoignant d'une volonté affirmée de proposer une esthétique soignée et professionnelle. De plus, le système offre une grande flexibilité d'installation et peut être utilisé aussi bien en rack 19 pouces qu'en unité de bureau autonome. Selon la configuration, son poids varie entre 148 et 159 kg.
Plaques froides et blocs de refroidissement pour GPU
