---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-2
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "Oracle"]
dates: []
keywords: ["agent", "license", "research", "valuation"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [140, 289]
sha256: 727f2490471973ad8b996038dfe423d5bacb1032b6511654000cf7a78dc70201
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**Règle de choix terrain :**
- Moins de ~200 équipements, besoin simple (alarmes + topo + perfs + backup
  de configs) → **Standard** suffit dans 90 % des cas.
- Besoin de WLAN unifié ou d'analyse de trafic (NTA) → **Standard** minimum.
- Multi-sites avec NMS locaux remontant vers un NMS central, ou besoin de
  haute disponibilité en cluster → **Professional**.
- Maquette, labo, tout petit site → **Compact**.

## 5. eSight face à iMaster NCE : la question de la fin de vie, honnêtement

C'est LA question à se poser avant d'investir. Voici les faits vérifiés :

- **eSight est toujours commercialisé** : la brochure « Huawei eSight 23.1 »
  est datée du **27 mars 2024** (document « Huawei confidential » diffusé
  par les partenaires). Le produit reçoit donc encore des versions.
- **iMaster NCE est la plateforme stratégique de Huawei** pour l'avenir :
  système autonome de gestion et de contrôle réseau pour l'entreprise,
  qui intègre gestion traditionnelle + fonctions de contrôle SDN + analyse
  de données basée sur les intentions de service, avec moteur IA
  (recommandation de topologies, provisioning automatique, telemetry).
  Les analystes (Appledore Research, 2023) décrivent NCE comme la gamme
  « destinée à remplacer les EMS historiques » de Huawei.
- **Aucune date officielle de fin de vie (EoL/EoS) d'eSight n'a été trouvée
  dans les sources publiques consultées** → **à vérifier sur la
  documentation officielle** auprès de votre partenaire Huawei.

**Positionnement honnête pour un chef de service :**

| Critère | eSight | iMaster NCE |
|---|---|---|
| Nature | NMS classique (FCAPS) | Plateforme autonome (management + contrôle + analyse) |
| Idéal pour | Superviser un parc existant, multi-constructeur partiel | Nouveaux déploiements campus/SD-WAN, automatisation |
| SDN / provisioning auto | Non (ou marginal) | Oui (cœur du produit) |
| Courbe d'apprentissage | Modérée | Plus forte |
| Investissement | Licence perpétuelle + SnS historiquement | Modèle plus orienté souscription |

**Trajectoire de migration à anticiper :** si vous déployez eSight aujourd'hui
sur un parc existant, prévoyez dans votre feuille de route à 3-5 ans une
évaluation de NCE (notamment NCE-Campus pour le Wi-Fi/filaire). Concrètement :
gardez vos processus d'exploitation (runbooks, nommage, gestion des alarmes)
**indépendants de l'outil** pour que la migration ne soit qu'un changement
de console, pas une réinvention. Les exports réguliers d'inventaire et de
configurations depuis eSight sont votre assurance-vie pour toute migration.

## 6. Forces et limites (retours d'exploitants)

Synthèse de retours d'utilisateurs (plateformes d'avis, 2022-2026) :

**Forces :**
- Stabilité saluée (notes 9/10) : « pas de crashs ni de glitches majeurs ».
- Parfaitement intégré aux équipements Huawei (MIB propriétaires, NTA).
- Installation « one-click » simple sur appliance Huawei fournie.
- Scalabilité correcte dans l'écosystème Huawei.

**Limites :**
- Installation sur VM jugée « très complexe » quand ça coince ;
  prévoir un admin qui connaît le produit.
- Support multi-constructeur perfectible (très bon sur Huawei, moyen
  sur le reste — un avis le note 2-3/10 hors Huawei).
- Cluster/redondance : procédure jugée complexe.
- Documentation parfois ardue à naviguer pour les cas d'erreur.

**Traduction pour Zelef** : eSight est un excellent choix si votre parc est
majoritairement Huawei ; prévoyez de la compétence interne ou un partenaire
pour l'installation et la montée en charge, et ne lui demandez pas d'être
un Zabbix générique.

## 7. Modèle de licences

D'après les grilles partenaires publiques (ex. S-Part Numbers) :

- Licence **par équipement géré** (1 device), par instance applicative,
  par caméra (vidéosurveillance), par AP/contrôleur (WLAN).
- Chaque licence est couplée à **1 an de souscription et support (SnS)**
  minimum obligatoire (ex. : *« eSight Server Management License,
  1 Year Subscription and Support, 1 Device »*).
- Les composants optionnels (stockage, applicatif, vidéosurveillance,
  micro-ondes, PON) se licencient séparément.

**Points de vigilance terrain :**
- Comptez vos équipements **avant** d'acheter : faites un inventaire
  (script SNMP walk ou export existant) et ajoutez 20 % de marge pour
  la croissance.
- Vérifiez ce qui compte comme « 1 device » : un châssis avec cartes ?
  Un stack ? Un contrôleur + ses AP ? → **à vérifier sur la documentation
  officielle** de votre version, les règles de comptage changent selon
  les composants.
- Renouvelez le SnS : sans support, pas de patchs de sécurité — et un NMS
  non patché qui a les accès SNMP en écriture sur tout le réseau, c'est
  un risque majeur.
- Gardez une **copie des fichiers/actes de licence** dans votre PRA
  (voir section 12).

---

# 2. ARCHITECTURE TECHNIQUE

## 8. Vue d'ensemble de l'architecture

eSight repose sur une architecture **centralisée** avec éventuellement
une **hiérarchie** de NMS (édition Professional) :

```
                    ┌─────────────────────────────────┐
                    │        eSight Server            │
                    │  ┌──────────┐  ┌──────────────┐ │
                    │  │ Services │  │  Base de     │ │
                    │  │ eSight   │◄─┤  données     │ │
                    │  │ (Java)   │  │ (MySQL/SQL   │ │
                    │  └────┬─────┘  │  Server/Oracle)│
                    │       │        └──────────────┘ │
                    └───────┼─────────────────────────┘
                            │ SNMP / syslog / NetStream
                            │ SSH/Telnet / ICMP / HTTP(S)
                ┌───────────┼───────────────┐
                ▼           ▼               ▼
           Routeurs    Commutateurs      Firewalls
           Huawei      (tiers: HP,       WLAN (AC/AP)
           (+ tiers)   Cisco...)         Serveurs, stockage...
```

Le serveur eSight **interroge** (polling SNMP/ICMP) et **écoute**
(traps SNMP, syslog) les équipements. Il n'y a pas d'agent à installer
sur les équipements réseau : tout passe par les protocoles standard.
C'est une architecture **sans agent** côté réseau — un point fort pour
le déploiement.

## 9. Architecture Browser/Server : aucun client à installer

Point structurant : eSight utilise une architecture **B/S (Browser/Server)**.

- L'administration se fait **à 100 % via navigateur web** (interface
  Web 2.0 d'après le constructeur). Aucun client lourd à déployer
  sur les postes des exploitants.
- Conséquence directe : pour donner l'accès à un technicien d'astreinte,
  il suffit d'une URL + un compte (voir section 11). En astreinte de nuit,
  on consulte les alarmes depuis n'importe quel PC — voire une tablette.
- En contrepartie : le serveur eSight doit être **joignable en HTTPS
  depuis tous les postes d'exploitation**, et sa disponibilité conditionne
  tout l'outillage (d'où l'importance du PRA, section 12).
- Le constructeur annonce qu'eSight peut même tourner sur un **PC portable**
  (édition compacte) : utile pour une maquette ou un audit ponctuel sur site.

## 10. La base de données : le cœur (et le point faible) du système

eSight s'appuie sur une base de données relationnelle qui stocke :
l'inventaire des équipements, la topologie, les alarmes (courantes et
historiques), les données de performance, les configurations sauvegardées,
les comptes utilisateurs et les journaux d'audit.

