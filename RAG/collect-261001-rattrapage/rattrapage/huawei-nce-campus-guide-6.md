---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-6
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [412, 491]
sha256: bc75e8359c01746b4cbcb9b313ccc9f0b526046fbae26d7dac2ca7331d7ea4e3
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **HTTP/2** : utilisé pour la télémétrie à haute fréquence (streaming telemetry) et certains échanges d'authentification sur les équipements récents. Plus efficace que le polling SNMP pour la donnée temps réel.
- **HTTPS** : portails captifs, redirection vers les interfaces web des équipements (la documentation note que la redirection SWEB vers le système web des switches/AR n'est pas supportée dans certaines versions — configuration via CLI/SSH dans ce cas).
- **TCP dédié** : les composants d'authentification distants synchronisent les informations d'authentification et d'identification des terminaux avec NCE-Campus via des canaux TCP.

En pratique d'exploitation : ces flux doivent être **autorisés explicitement** dans les pare-feu entre NCE et les sites (voir section 37 — prérequis réseau). Un oubli de flux = un onboarding qui échoue mystérieusement (voir cas pratiques).

---

# PARTIE 3 — BASE DE DONNÉES, ANALYSEUR, HA ET DIMENSIONNEMENT

## 19. Base de données et stockage — que garde NCE ?

NCE-Campus s'appuie sur des bases de données internes (le détail moteur exact dépend de la version et du mode de déploiement — ne pas spéculer : **à vérifier** dans le guide d'installation de la version cible). Fonctionnellement, il conserve :

- **Inventaire** : équipements, modules, versions logicielles, numéros de série/ESN.
- **Configurations** : configurations attendues (templates instanciés), historiques de déploiement, écarts détectés.
- **Données de performance** : indicateurs CPU/mémoire/interfaces, télémétrie — avec des durées de rétention paramétrables (la rétention longue = stockage conséquent : à dimensionner).
- **Alarmes et événements** : base d'alarmes courantes + historique.
- **Données d'authentification** : utilisateurs, sessions, logs d'accès (sensibles — voir sécurité).
- **Rapports** : rapports générés et planifiés.

Conséquences pratiques : prévoir du **stockage redondant** (RAID), une **politique de rétention** (ex : données fines 30 jours, agrégées 1 an — à adapter), et inclure les bases NCE dans la **sauvegarde** (section 107). La volumétrie croît avec le nombre d'équipements et la fréquence de télémétrie — c'est un poste du dimensionnement (sections 23-25).

## 20. iMaster NCE-CampusInsight — l'analyseur (composant distinct)

CampusInsight est la brique « analysis » avancée : plateforme d'analyse big data + ML qui **numérise l'expérience** — par utilisateur, par application, à chaque instant. Ce qu'elle apporte au-delà du NCE de base :

- **Visibilité d'expérience** : qualité perçue par utilisateur/application (pas seulement l'état des équipements).
- **Détection proactive** : identification de problèmes potentiels (utilisateurs, applications, réseau) avant qu'ils ne deviennent des tickets.
- **Root cause analysis** assistée et base de connaissances de pannes (localisation des pannes WLAN/LAN/WAN « en quelques minutes » selon Huawei — objectif à valider en conditions réelles).
- **Optimisation Wi-Fi** : ajustement intelligent des canaux, largeurs de bande et puissances à partir de l'historique de trafic télémétré (test Tolly cité par Huawei : +50 % de performance globale du réseau sans-fil après optimisation — résultat de laboratoire, à considérer comme un potentiel).

Points commerciaux à connaître (datasheet CampusInsight V100R025C00) :

- Souscription par **package logiciel** : package de base d'analyse réseau intelligent (**obligatoire**, tarifé selon le type et la quantité de NE — Network Elements), puis packages optionnels (analyse applicative, optimisation/auto-guérison, analyse de consommation énergétique, copilote O&M).
- C'est donc un **coût additionnel** au NCE-Campus lui-même — à chiffrer séparément.

## 21. Le composant d'authentification — déploiement distribué

Dans les architectures multi-sites, l'authentification des utilisateurs (802.1X, portail) ne doit pas dépendre du lien WAN vers le contrôleur central. NCE-Campus prévoit pour cela un **composant d'authentification** :

- Intégré à NCE-Campus **en tant que service** (pas un produit séparé à acheter dans la logique de base).
- Déployable sur les **branches distantes** : jusqu'à **20 composants** selon la documentation V300R022C10.
- **Synchronisation automatique** avec NCE-Campus central via des canaux TCP : les informations d'authentification des utilisateurs et d'identification des terminaux sont répliquées.
- En cas de coupure du lien vers le central, la branche continue d'authentifier **localement** — continuité de service.

Pour Zelef : sur des sites distants avec des liens WAN peu fiables, prévoir ce composant dans l'architecture. C'est un argument fort face à une authentification 100 % centralisée.

## 22. Haute disponibilité — modes de déploiement

Un contrôleur qui pilote tout le campus devient critique : sa disponibilité doit être traitée comme celle d'un cœur de réseau.

Principes (à valider selon la version — les modes exacts varient) :

- **Déploiement redondant** : NCE-Campus supporte des architectures à haute disponibilité (cluster / actif-passif selon les versions et les guides de déploiement). Exiger du partenaire le **guide de déploiement HA** de la version cible et le tester.
- **Comportement des équipements en cas de perte du contrôleur** : point fondamental — les équipements Huawei **continuent de commuter** avec leur dernière configuration (le plan de données est local). Ce qui s'arrête : les nouveaux déploiements, les changements de politiques, l'authentification centralisée (sauf composant local), la supervision. Le réseau ne « tombe » pas, il se « fige ». C'est rassurant mais à tester (voir cas pratique 11).
- **Bases de données** : la HA doit couvrir les données (réplication) autant que les services.
- **Sauvegarde/restauration** : même en HA, garder une stratégie de backup externe (section 107) — la HA protège de la panne, pas de l'erreur humaine ni du ransomware.

Checklist HA : deux nœuds sur **sites/racks différents**, alimentations et liens redondants, supervision du contrôleur lui-même (qui supervise le superviseur ? — prévoir Zabbix ou équivalent), exercice de bascule **au moins une fois par an**.

## 23. Dimensionnement : petit site (ordre de grandeur)

Ordres de grandeur indicatifs — **à faire valider par le partenaire** avec l'outil de dimensionnement Huawei de la version cible :

- **Périmètre** : 1 site, < 50 équipements (ex : 5-10 switches S310, 10-30 AP, 1-2 AR720, 1 USG6000).
- **Déploiement** : 1 serveur/VM unique (sans HA) ou petite grappe selon l'exigence.
- **Ressources** : VM avec vCPU/RAM/disque « confortables » (les contrôleurs SDN + bases sont gourmands en RAM ; ne pas sous-dimensionner — **valeurs exactes à vérifier** dans le guide d'installation).
- **Remarque honnête** : à cette taille, **le ROI de NCE est discutable** — eKit ou eSight + gestion locale sont souvent plus économiques (voir section 10 et 152). NCE se justifie ici seulement si le site est le pilote d'un déploiement multi-sites.

## 24. Dimensionnement : site moyen (ordre de grandeur)

- **Périmètre** : 1 campus ou 3-10 sites, 50 à 500 équipements administrables.
- **Déploiement** : NCE-Campus en **HA** (2 nœuds), bases répliquées, éventuellement composant d'authentification sur les branches critiques.
- **Ressources** : serveurs dédiés ou VMs sur cluster virtualisé avec réservations de ressources (CPU/RAM garantis — la télémétrie et l'analyse n'aiment pas la contention).
- **Réseau** : liens de management dédiés (VLAN/OOB), NTP commun, DNS.
- **C'est le sweet spot de NCE-Campus** : assez grand pour que l'automatisation paie, assez petit pour rester administrable par une petite équipe formée.

## 25. Dimensionnement : grand campus / multi-sites (ordre de grandeur)

