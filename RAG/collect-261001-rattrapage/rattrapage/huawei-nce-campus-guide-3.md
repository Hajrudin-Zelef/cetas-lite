---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-3
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [159, 233]
sha256: c3afbb54f370ca3c27bf621d28919a100684d31f0fdc97ec5f4d87362928716d
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

148. Cas pratique 25 — Coupure électrique et reprise (lien onduleurs)
149. Limite 1 — Courbe d'apprentissage
150. Limite 2 — Dépendance au contrôleur (single point of management)
151. Limite 3 — Coût des licences
152. Limite 4 — Maturité sur petit site (eKit suffit)
153. Limite 5 — Écosystème et documentation
154. Limite 6 — Dépendance au support Huawei/partenaire
155. Synthèse honnête — faut-il y aller ?
156. Quiz — 10 questions
157. Quiz — réponses détaillées
158. Glossaire
159. Pour aller plus loin — ressources et prochaines étapes

---

# PARTIE 1 — POSITIONNEMENT

## 1. Ce qu'est iMaster NCE-Campus — définition officielle

iMaster NCE-Campus est le système de gestion et de contrôle de réseau campus de Huawei, qualifié par le constructeur de « système de gestion et de contrôle de réseau autonome (autonomous driving network) de nouvelle génération pour les réseaux campus ». Concrètement, c'est une plateforme logicielle qui **intègre trois fonctions** historiquement séparées :

- **Management (gestion)** : inventaire, configuration, supervision, alarmes, rapports — le rôle classique d'un NMS comme eSight.
- **Control (contrôle)** : pilotage SDN des équipements via des protocoles structurés (NETCONF/YANG notamment), déploiement automatisé de services, orchestration des politiques.
- **Analysis (analyse)** : collecte massive de données (télémétrie), big data et algorithmes de machine learning pour la visibilité d'expérience, la prédiction de pannes et l'analyse de cause racine.

La promesse marketing — à prendre avec le recul d'un exploitant — est la « full-lifecycle automation » : de la planification à la construction, de l'exploitation à l'optimisation et à la sécurité, tout le cycle de vie du réseau campus serait piloté depuis un point unique, avec une réduction annoncée des coûts OPEX/O&M (Huawei évoque une division par deux des coûts d'O&M dans ses argumentaires — chiffre marketing, à considérer comme un objectif, pas une garantie).

Techniquement, iMaster NCE-Campus se déploie soit **on-premises** (sur serveurs/VM du client), soit en **cloud** (modèle MSP / cloud public Huawei selon les offres commerciales), ce qui ouvre des modèles économiques variés : achat classique ou souscription.

Point important pour Zelef : NCE-Campus gère les familles d'équipements qu'il connaît déjà — **switches CloudEngine (dont S310), routeurs NetEngine AR (dont AR720), AP AirEngine (dont AP361 et AP761), firewalls USG (dont USG6000)** — via une gestion unifiée filaire + sans-fil + WAN de branche.

## 2. La famille iMaster NCE : NCE-Campus, NCE-Fabric, NCE-WAN, NCE-IP

« iMaster NCE » n'est pas un produit unique mais une **famille de contrôleurs** partageant une base technologique commune (SDN, big data, API ouvertes). Il faut les distinguer clairement :

| Membre de la famille | Périmètre | Rôle |
|---|---|---|
| **iMaster NCE-Campus** | Réseaux campus et agences/branches | Gestion unifiée LAN filaire, WLAN, WAN de branche (SD-WAN), authentification des utilisateurs |
| **iMaster NCE-Fabric** | Data centers | Automatisation du fabric DC (VXLAN/EVPN), gestion du cycle de vie des commutateurs de data center |
| **iMaster NCE-WAN** | Réseaux WAN / transport | Contrôleur SD-WAN et orchestration WAN (selon la documentation SD-WAN Huawei, NCE-Campus peut aussi gérer le SD-WAN et « sera le contrôleur SD-WAN mainstream à l'avenir » — formulation Huawei) |
| **iMaster NCE-IP** | Réseaux IP d'opérateurs / entreprises | Gestion des réseaux IP à grande échelle (selon déclinaisons commerciales) |

À cela s'ajoute **iMaster NCE-CampusInsight**, l'analyseur intelligent du campus : plateforme d'analyse big data/ML qui numérise l'expérience utilisateur (par utilisateur, par application, à chaque instant). C'est un **composant indépendant** (licence et souscription séparées), qui se connecte à NCE-Campus.

Pour un chef de service systèmes & énergies comme Zelef, seul **NCE-Campus** (éventuellement + CampusInsight) est pertinent : NCE-Fabric concerne les data centers, NCE-WAN les opérateurs.

## 3. Pourquoi Huawei pousse NCE côté campus — la logique produit

Huawei pousse iMaster NCE-Campus pour des raisons à la fois techniques et commerciales qu'il faut comprendre pour négocier sereinement :

1. **Fin du paradigme CLI + scripts** : Huawei constate — à juste titre — que les réseaux campus deviennent trop complexes (BYOD, IoT, Wi-Fi 6/7, cloud, vidéo) pour être gérés en CLI équipement par équipement. NCE est la réponse « intent-based » : on déclare l'intention (qui accède à quoi, avec quelle qualité), le contrôleur traduit en configurations.
2. **Convergence LAN/WLAN/WAN** : historiquement, le filaire (switches), le Wi-Fi (contrôleur WAC/AC) et les branches (routeurs) se géraient avec des outils différents. NCE-Campus unifie les trois — c'est son argument différenciant principal face à eSight, qui supervise mais n'orchestre pas.
3. **Modèle économique récurrent** : les licences NCE (device-day, souscriptions) génèrent des revenus récurrents, contrairement à la licence perpétuelle classique. C'est une logique d'abonnement assumée par tout le secteur (Cisco DNA, Aruba Central, Juniper Mist font pareil).
4. **IA comme différenciateur marketing** : CampusInsight, l'optimisation radio par ML, la prédiction de pannes — Huawei investit massivement dans l'« autonomous driving network » pour se différencier de Cisco/Aruba.
5. **Cloud et MSP** : NCE permet aux partenaires/intégrateurs de proposer du « network as a service » managé multi-tenant — nouveau canal de vente pour Huawei.

Conséquence pour l'acheteur : NCE n'est pas qu'un « eSight amélioré », c'est un **changement de modèle** (pilotage centralisé + abonnement). Il faut l'évaluer comme tel, pas comme une simple mise à jour.

## 4. ADN : Autonomous Driving Network — ce que ça veut dire concrètement

« Autonomous Driving Network » (ADN) est le concept marketing de Huawei pour les réseaux qui se gèrent (presque) seuls. Derrière le slogan, il y a des fonctions réelles, classées par niveau d'autonomie :

- **Niveau 1-2 (assistance)** : tableaux de bord unifiés, topologie automatique, alarmes corrélées. C'est ce que fait déjà un bon NMS.
- **Niveau 3 (automatisation conditionnelle)** : ZTP (zero-touch provisioning), templates de configuration, mise à jour firmware en masse, déploiement de politiques 802.1X en un clic. C'est le cœur utile de NCE-Campus pour une équipe d'exploitation.
- **Niveau 4 (prédiction)** : CampusInsight analyse la télémétrie pour prédire les pannes (ex : port en dégradation, interférences Wi-Fi croissantes), recommande des optimisations radio. Utile, mais à valider sur le terrain : les faux positifs existent.
- **Niveau 5 (autonomie complète)** : le réseau se répare seul. **N'existe pas en production** à ce jour — à considérer comme une vision, pas une fonctionnalité achetable.

En pratique pour Zelef : attendre de NCE les niveaux 1 à 3 (réels et matures), considérer le niveau 4 comme un bonus à évaluer en pilote, ignorer le niveau 5.

## 5. Relation avec eSight — le positionnement officiel de Huawei

eSight est la plateforme historique de gestion unifiée Huawei (Unified Network Management Platform) : supervision multi-vendeurs (via SNMP), gestion des configurations, alarmes, rapports, topologie. NCE-Campus n'est **pas** présenté par Huawei comme un simple successeur versionné d'eSight, mais comme une **nouvelle génération** fondée sur le SDN.

Le positionnement officiel, tel qu'il ressort de la documentation :

