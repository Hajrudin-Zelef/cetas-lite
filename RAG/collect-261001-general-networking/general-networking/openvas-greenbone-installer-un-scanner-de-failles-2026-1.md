---
id: collect-261001-general-networking/general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026-1
title: "Vérifier l'espace disque disponible (minimum 20 Go recommandé)"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["benchmarks", "open source"]
source: docs/RAG/collect-261001-general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026.md
source_anchor: ""
source_lines: [1, 38]
sha256: 0ae1e90dbdc0b3396700e0130209ac9de6b7afb725131f0076e95edcd2455bb7
---

# Vérifier l'espace disque disponible (minimum 20 Go recommandé)

Selon le classement Wiz publié en 2026, OpenVAS s’impose comme l’outil de gestion des vulnérabilités le plus utilisé dans les infrastructures open source en Europe, porté par un feed de sécurité mis à jour quotidiennement et couvrant plus de 160 000 tests de détection. Avec la directive NIS2 qui impose désormais des scans de vulnérabilités réguliers à des dizaines de milliers d’entités européennes, la question n’est plus de savoir s’il faut scanner son réseau, mais avec quel outil et à quelle fréquence. Ce tutoriel détaille l’installation complète de Greenbone Community Edition (OpenVAS) sur Debian et Ubuntu, du premier scan jusqu’à l’automatisation des rapports, avec la version 25.0.6 publiée le 6 août 2026.

## Qu’est-ce qu’OpenVAS et pourquoi Greenbone domine la gestion de vulnérabilités en 2026

OpenVAS (Open Vulnerability Assessment Scanner) est le moteur de scan au cœur de Greenbone Community Edition, une suite open source de gestion des vulnérabilités. Concrètement, l’outil interroge chaque machine de votre réseau avec des milliers de tests appelés NVT (Network Vulnerability Tests), compare les résultats à des bases de données de failles connues (CVE), puis génère un rapport classé par gravité selon le score CVSS. Ce n’est pas un simple scanner de ports comme peut l’être Nmap : OpenVAS va chercher les versions de logiciels exposées, teste les mauvaises configurations, et signale les correctifs manquants.

La branche actuelle, OPENVAS SCAN 25.0, est classée en statut de cycle de vie « Mature » par Greenbone, avec un patch 25.0.6 daté du 6 août 2026 et un manuel officiel pour cette même branche 25.0 recensé sur Greenbone.net en août 2026. Ce numéro de suite ne doit pas être confondu avec la version du composant scanner lui-même : sur le dépôt GitHub officiel, la dernière release taguée du moteur OpenVAS Scanner reste la v23.50.24, publiée en août 2026, un numéro également repris par la fiche Wikipedia du projet dans son infobox mise à jour le 31 août 2026 — les deux schémas de version (produit Greenbone vs composant scanner) coexistent et peuvent prêter à confusion lors d’un audit de conformité. Une branche antérieure, OPENVAS SCAN 24.10, reste maintenue avec un patch 24.10.11 publié le 19 février 2026, pour les environnements qui n’ont pas encore migré. Côté appliance, Greenbone a lancé OPENVAS OS 2.0 le 1er juillet 2026, avec un patch 2.0.1 daté du 30 juillet 2026, une version « New » pensée pour les déploiements clé en main en VM ou sur matériel dédié.

Ce qui distingue Greenbone des scanners commerciaux comme Nessus ou Qualys, c’est le modèle open source : le moteur de scan et l’interface web (GSA, Greenbone Security Assistant) sont librement installables, seul l’accès au feed « Enterprise » le plus complet est payant. Le feed communautaire, mis à jour chaque jour, reste largement suffisant pour la majorité des PME, collectivités et laboratoires de test. C’est cette gratuité couplée à une couverture large qui explique pourquoi l’outil revient si souvent dans les benchmarks 2026, y compris pour des cas d’usage sur des hyperviseurs comme Proxmox VE 8.0+ ou Huawei FusionCompute 8.0.

Pour resituer OpenVAS dans une stratégie de sécurité complète, il complète des outils déjà couverts sur ce site comme Nmap pour la découverte réseau, Trivy pour les images de conteneurs ou Wazuh pour la corrélation d’événements. Là où Nmap répond à « quels ports sont ouverts » et Trivy à « quelles failles dans mes images Docker », OpenVAS répond à « quelles vulnérabilités exploitables tournent sur mes serveurs et postes de travail, avec quel niveau de gravité ».

## Comprendre les familles de tests NVT et les profils de scan

Avant de lancer le moindre scan, il faut comprendre comment OpenVAS organise ses tests. Les NVT sont regroupés par familles : détection de services, applications web, bases de données, systèmes d’exploitation, équipements réseau, protocoles industriels (SCADA/ICS), et bien d’autres. Chaque famille peut être activée ou désactivée indépendamment dans une configuration de scan personnalisée, ce qui permet d’adapter précisément la portée à votre environnement plutôt que de lancer systématiquement l’intégralité des 160 000 tests du feed.

Greenbone fournit plusieurs profils de scan préconfigurés, accessibles dans Configuration → Scan Configs. Le profil `Full and fast` couvre l’essentiel des familles pertinentes tout en évitant les tests les plus lents ou les plus intrusifs. Le profil `Full and very deep` active davantage de tests, y compris certains considérés comme potentiellement perturbateurs pour les services fragiles, et rallonge la durée du scan de façon significative. Le profil `Discovery`, plus léger, se limite à l’identification des services et versions sans chercher activement à exploiter une faille, ce qui en fait un bon point de départ pour cartographier un réseau inconnu avant d’aller plus loin.

Pour un environnement de production sensible, la meilleure pratique consiste à créer une configuration de scan sur mesure : partez du profil `Full and fast`, dupliquez-le, puis désactivez manuellement les familles de tests non pertinentes pour votre parc (par exemple les tests SCADA si vous n’avez aucun automate industriel, ou les tests spécifiques à des CMS que vous n’utilisez pas). Ce réglage réduit à la fois la durée du scan et le risque de faux positifs liés à des tests hors sujet.

## Prérequis avant d’installer OpenVAS / Greenbone Community Edition

Avant de démarrer, vérifiez que votre environnement correspond aux exigences minimales. Greenbone Community Edition est gourmand en ressources car il combine une base PostgreSQL, un moteur de scan multithread et un feed de plusieurs gigaoctets à synchroniser au premier lancement. Pour les puristes qui préfèrent une installation depuis les sources plutôt que Docker, un script d’installation Debian 12 maintenu sur GitHub (publié en février 2024 et mis à jour en octobre 2025) exige un système Debian 12 propre et entièrement à jour, et expose ensuite l’interface web à l’adresse https://

| Composant | Minimum | Recommandé (PME/collectivité) | 
|---|---|---|
| Système d’exploitation | Debian 12 / Ubuntu 22.04 LTS | Debian 13 / Ubuntu 24.04 LTS | 
| RAM | 4 Go | 16 Go ou plus | 
| CPU | 2 vCPU | 4 vCPU ou plus | 
| Stockage | 20 Go | 100 Go SSD (feed + rapports) | 
| Docker Engine | 24.x | 27.x ou supérieur | 
| Docker Compose | v2.20 | v2.30 ou supérieur | 
| Bande passante Internet | 10 Mbps | 50 Mbps (synchro feed quotidienne) | 

Vous aurez aussi besoin d’un accès root ou sudo sur la machine cible, d’une adresse IP fixe ou d’un nom de domaine interne pour accéder à l’interface web, et idéalement d’un second réseau isolé (VLAN de scan) si vous comptez auditer des systèmes de production. Ne lancez jamais un premier scan complet directement sur un réseau de production sans avoir prévenu les équipes : un scan agressif peut faire planter des équipements réseau ou des IoT fragiles.

Enfin, gardez à l’esprit un point de conformité : au sein de l’Union européenne, scanner un système sans autorisation explicite du propriétaire peut relever de l’article 323-1 du Code pénal français (accès frauduleux à un système de traitement automatisé de données). Ce tutoriel s’adresse à des scans réalisés sur votre propre infrastructure ou dans le cadre d’un mandat écrit.

