---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-1
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1, 139]
sha256: 5040f6743241dc1cbc81f07dcde8ab2f69137b81f6edfae7d487ded67a8ee068
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

> **NMS (Network Management System) Huawei pour l'entreprise**
> Rédigé le 27 septembre 2026 — orienté terrain, pour chef de service systèmes & énergies.
>
> ⚠️ **Avertissement de méthode** : ce guide s'appuie sur la documentation constructeur
> publique (datasheet eSight, brochure eSight 23.1 de mars 2024, retours d'exploitants).
> Les comportements précis (numéros de ports, chemins de menus, comptes par défaut)
> **dépendent de la version installée**. Chaque fois qu'une information est
> version-dépendante ou non vérifiable publiquement, la mention
> **« à vérifier sur la documentation officielle »** est indiquée explicitement.
> Ne jamais appliquer une procédure d'administration sans l'avoir confrontée
> à la documentation de VOTRE version.

---

## Table des matières

- [1. Présentation d'eSight](#1-présentation-desight)
- [2. Architecture technique](#2-architecture-technique)
- [3. Dimensionnement](#3-dimensionnement)
- [4. Installation](#4-installation)
- [5. Découverte des équipements](#5-découverte-des-équipements)
- [6. Topologie](#6-topologie)
- [7. Gestion des alarmes](#7-gestion-des-alarmes)
- [8. Supervision des performances](#8-supervision-des-performances)
- [9. Gestion des configurations](#9-gestion-des-configurations)
- [10. Gestion WLAN](#10-gestion-wlan)
- [11. Comptes, rôles et audit](#11-comptes-rôles-et-audit)
- [12. Sauvegarde d'eSight et PRA](#12-sauvegarde-desight-et-pra)
- [13. Intégration northbound et écosystème](#13-intégration-northbound-et-écosystème)
- [14. Dépannage : 17 cas terrain](#14-dépannage--17-cas-terrain)
- [15. Bonnes pratiques d'exploitation](#15-bonnes-pratiques-dexploitation)
- [16. Maintenance](#16-maintenance)
- [17. Pour aller plus loin](#17-pour-aller-plus-loin)
- [18. Pense-bête de poche](#18-pense-bête-de-poche)
- [19. Les 20 pièges classiques](#19-les-20-pièges-classiques)
- [20. Glossaire](#20-glossaire)
- [21. Quiz](#21-quiz)
- [Annexe A — Recette de mise en service](#annexe-a--recette-de-mise-en-service)
- [Annexe B — Fiches réflexes par type d'équipement](#annexe-b--fiches-réflexes-par-type-déquipement)
- [Annexe C — Automatisation autour d'eSight](#annexe-c--automatisation-autour-desight)
- [Annexe D — Cas pratiques commentés](#annexe-d--cas-pratiques-commentés)

---

# 1. PRÉSENTATION D'ESIGHT

## 1. Ce qu'est eSight, en une phrase

**Huawei eSight est la plateforme unifiée d'exploitation et de maintenance (O&M)
du système d'information d'entreprise de Huawei** : un NMS qui supervise de façon
centralisée le réseau (routeurs, commutateurs, firewalls, WLAN), les serveurs,
le stockage, la vidéosurveillance, les applications et les liaisons micro-ondes —
qu'ils soient Huawei ou, dans une certaine mesure, d'autres constructeurs.

Le positionnement officiel du constructeur : *« eSight ICT O&M system is a Huawei
new-generation solution for enterprise users to centrally manage the basic network,
unified communications (UC), telepresence meetings, video surveillance devices,
and data centers »* (brochure eSight 23.1, mars 2024).

Concrètement, sur le terrain, eSight fait quatre métiers :

1. **Fault management** — collecte et traite les alarmes (SNMP traps, syslog).
2. **Performance management** — collecte les indicateurs (CPU, mémoire, interfaces)
   et génère tableaux de bord et rapports.
3. **Topology management** — cartographie automatique du réseau.
4. **Configuration management** — sauvegarde, comparaison et restauration
   des configurations des équipements.

## 2. À qui s'adresse eSight

- **PME/ETI et grands comptes** disposant d'un parc majoritairement Huawei
  (ou mixte avec une base Huawei significative).
- **Équipes d'exploitation** (NOC, helpdesk niveau 2/3) qui ont besoin d'un
  point d'entrée unique pour voir l'état du SI.
- **Chefs de service systèmes & réseaux** qui veulent des rapports
  d'activité, des KPI et de la traçabilité pour piloter l'équipe.
- **Intégrateurs** qui déploient du Huawei chez leurs clients et ont besoin
  d'un outil de supervision « constructeur » qui parle parfaitement aux
  équipements Huawei (MIB propriétaires comprises).

Ce n'est PAS un outil pour : superviser uniquement du multi-constructeur
générique (des outils comme Zabbix ou PRTG sont plus agnostiques),
faire du SDN/automatisation avancée (c'est le terrain d'iMaster NCE),
ou gérer du datacenter cloud-native.

## 3. Périmètre fonctionnel : les modules (composants)

eSight est **modulaire** : on achète la plateforme de base puis on ajoute
des composants selon le besoin. D'après la brochure officielle eSight 23.1,
les composants sont :

| Composant | Ce qu'il gère |
|---|---|
| eSight basic management | Plateforme O&M : ressources, topologies, alarmes, performances, rapports, portails |
| eSight network device management | Découverte et maintenance des équipements réseau, interfaces, liens, configurations |
| eSight WLAN management | Ressources du réseau Wi-Fi de campus (contrôleurs AC et points d'accès AP), diagnostic des pannes sans fil, topologie unifiée filaire/sans fil |
| eSight server device management | Ressources serveurs : CPU, mémoire, disques, ports réseau, vues simulées, systèmes d'exploitation |
| eSight storage device management | Ressources, performances et capacité des baies OceanStor (y compris fonctions de protection anti-ransomware) |
| eSight video surveillance management | Ressources de vidéosurveillance (IVS, caméras), topologies de service, performances |
| eSight microwave device management | O&M et montées de version des équipements micro-ondes (RTN) |
| eSight Application Management | Supervision centralisée des OS, bases de données, serveurs web et applicatifs, URLs |
| eSight PON device management | Réseaux GPON : OLT et ONU — état, alarmes, performances, configurations, localisation rapide des pannes |
| eSight network SLA management | Diagnostic périodique automatique des chemins réseau |

💡 **Lecture terrain** : dans la plupart des déploiements PME/ETI, on ne retient
que 2 à 3 modules : *basic management* + *network device management*,
éventuellement + *WLAN management* si le Wi-Fi est critique. Chaque module
ajouté = licence + charge sur le serveur + surface d'administration.
Ne pas acheter « au cas où ».

## 4. Les éditions : Express, Compact, Standard, Professional

Le datasheet constructeur distingue quatre éditions (les trois principales
documentées en détail) :

| Fonction | Compact | Standard | Professional |
|---|---|---|---|
| Alarmes, performance, topologie, fichiers de config | ✅ | ✅ | ✅ |
| Équipements (NE), liens, logs, ressources physiques | ✅ | ✅ | ✅ |
| Étiquettes électroniques, topologie IP | ✅ | ✅ | ✅ |
| Smart Configuration Tool | ✅ | ✅ | ✅ |
| Gestion de la sécurité, accès terminaux | ✅ | ✅ | ✅ |
| System Monitor Tool, Database Backup & Restore Tool, Fault Information Collection Tool | ✅ | ✅ | ✅ |
| Gestion WLAN | ❌ | ✅ | ✅ |
| NTA (analyse de trafic réseau / NetStream) | ❌ | ✅ | ✅ |
| Policy Center, SLA, QoS | ❌ | ✅ | ✅ |
| MPLS VPN, tunnels MPLS, IPSec VPN | ❌ | ✅ | ✅ |
| Gestion des rapports | ❌ | ✅ | ✅ |
| Interface northbound SNMP | ❌ | ✅ | ✅ |
| NMS hiérarchique (supervision de superviseurs) | ❌ | ❌ | ✅ |
| DC nCenter | ❌ | ❌ | ✅ |
| Double système / cluster (Linux uniquement) | ❌ | ❌ | ✅ |

*(Source : Huawei eSight Full Product Datasheet — à vérifier sur la
documentation officielle pour la version exacte que vous déployez,
les périmètres ont pu évoluer.)*

