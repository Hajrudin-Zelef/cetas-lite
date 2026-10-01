---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-3
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "Microsoft", "Oracle"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [290, 438]
sha256: 2ceeb9c23fc7a2562361bab6d94171900cdab4c92acee5bcb51f0b739ad0ac8e
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

Selon le datasheet (versions historiques — **à vérifier sur la
documentation officielle** pour les versions récentes) :

| OS du serveur eSight | Base supportée |
|---|---|
| Windows Server | MySQL 5.5 ou Microsoft SQL Server 2008 R2 |
| SUSE Linux Enterprise Server | Oracle Database 11g R2 |

**Enjeux terrain :**
- La base est **le premier goulot d'étranglement** : quand eSight rame,
  c'est elle dans 80 % des cas (voir cas n°13 en section 14).
- Prévoyez le **dimensionnement disque** en fonction de la **durée de
  rétention** des données de performance et d'alarmes (voir section 3).
- La base doit être **sauvegardée séparément** du système (outil
  *Database Backup & Restore* fourni — voir section 12).
- Ne jamais bidouiller la base à la main (UPDATE/DELETE en SQL direct) :
  vous casseriez la cohérence applicative et le support Huawei.

## 11. Protocoles de supervision : le langage entre eSight et les équipements

| Protocole | Sens | Usage dans eSight | Port usuel |
|---|---|---|---|
| SNMP v1/v2c/v3 | eSight ↔ équipement | Découverte, polling d'état et de compteurs, traps | UDP 161 (poll), UDP 162 (traps) |
| Syslog | Équipement → eSight | Remontée d'événements/logs, source d'alarmes | UDP 514 |
| NetStream (NTA) | Équipement → eSight | Analyse de trafic (qui parle à qui, top talkers) | UDP configuré (souvent 5000+) |
| ICMP | eSight → équipement | Test de joignabilité (ping) pendant la découverte | — |
| SSH / Telnet | eSight → équipement | Sauvegarde/restauration de configurations, Smart Config Tool | TCP 22 / 23 |
| HTTP/HTTPS | Navigateur → eSight | Console d'administration web | TCP 80/443 (ports exacts **à vérifier sur la documentation officielle**) |
| SMTP | eSight → relais mail | Notifications d'alarmes par e-mail | TCP 25/587 |
| SMPP/API SMS | eSight → passerelle | Notifications d'alarmes par SMS | Selon opérateur |

**Choix structurant : SNMP v2c vs v3.**
- **v2c** : simple (community string), mais **non chiffré et authentification
  faible**. Acceptable uniquement sur un VLAN de management isolé.
- **v3** : authentification + chiffrement (authPriv). **C'est le standard
  à viser** pour tout nouveau déploiement. eSight gère les deux.
- Règle d'or : la community SNMP ou les credentials v3 utilisés par eSight
  doivent être **dédiés à la supervision** (un compte par outil), jamais
  partagés avec d'autres usages.

## 12. Haute disponibilité : le cluster à deux nœuds

Le constructeur annonce : *« eSight supports two-node clusters in hot
standby mode »* (cluster à 2 nœuds en **secours à chaud**).

- Principe : deux serveurs eSight, un actif et un en standby qui prend
  le relais en cas de panne (bascule).
- D'après le datasheet, la fonction **Double Système n'est supportée que
  sous Linux** et relève de l'édition **Professional**.
- Protection DR (Disaster Recovery) également mentionnée pour la
  continuité de service.

**Lecture terrain :**
- La HA ne se bricole pas après coup : si vous visez 200+ équipements
  ou une supervision critique (site industriel, H24), **prévoyez le
  cluster dès l'installation**.
- Les retours d'exploitants qualifient la mise en cluster de « complexe » :
  faites-vous accompagner par le partenaire/intégrateur pour cette étape.
- Même avec un cluster, gardez des **sauvegardes externalisées**
  (section 12) : la HA protège de la panne serveur, pas de la corruption
  de base ni de l'erreur humaine.
- Détails de mise en œuvre (VIP, synchro de base, procédure de bascule)
  → **à vérifier sur la documentation officielle** de votre version.

## 13. NMS hiérarchique : superviser des superviseurs

En édition Professional, eSight supporte le **network management hiérarchique** :
des eSight « subordonnés » (par site, par filiale, par région) remontent
leurs alarmes, topologies et performances vers un eSight « supérieur »
au siège.

Cas d'usage typiques :
- Groupe multi-sites : chaque site a son eSight local (autonomie en cas
  de coupure WAN), le siège a la vision consolidée.
- Séparation des responsabilités : l'équipe locale gère son périmètre,
  la direction voit les indicateurs globaux.

**Attention** : chaque niveau ajoute de la latence de remontée d'alarmes
et de la complexité de licences. Ne hiérarchisez que si l'organisation
l'exige vraiment.

## 14. Interfaces externes : eSight n'est pas une île

La brochure 23.1 liste les interfaces externes fournies :

1. **Southbound** — accès aux équipements (SNMP, syslog, NetStream…).
2. **Northbound** — intégration vers un NMS de niveau supérieur
   (interface SNMP northbound en édition Standard/Professional,
   voir section 13).
3. **Administrator login** — console web.
4. **SMS server** — envoi de SMS d'alarmes.
5. **Authentication server** — intégration à un serveur d'authentification
   (annuaire d'entreprise, voir section 11).

C'est ce qui permet d'inscrire eSight dans un écosystème : un eSight qui
remonte ses alarmes critiques vers un manager-of-managers, qui envoie
des SMS via la passerelle de l'entreprise, qui authentifie les exploitants
sur l'AD — c'est un eSight bien intégré (section 13).

## 15. Sécurité de l'architecture : le NMS est une cible

Un NMS concentre des pouvoirs dangereux : il connaît la topologie complète,
possède des accès SNMP (parfois en écriture) et SSH sur tous les équipements,
et stocke les configurations (avec leurs secrets). **Compromettre eSight =
compromettre le réseau.**

Règles non négociables :
- eSight vit sur un **VLAN de management dédié**, filtré par firewall ;
  seuls les équipements supervisés et les postes d'exploitation y accèdent.
- **HTTPS uniquement** pour la console (pas de HTTP en clair).
- Comptes nominatifs (section 11), jamais de compte générique partagé
  pour l'exploitation courante.
- SNMP v3 authPriv partout où c'est possible ; à défaut, v2c avec des
  ACL strictes côté équipements (n'autoriser que l'IP d'eSight).
- Patchs eSight et OS suivis (section 16).

---

# 3. DIMENSIONNEMENT

## 16. Tableau de dimensionnement constructeur

Le datasheet officiel donne les configurations minimales par palier
d'équipements gérés (**données issues du datasheet historique** ;
les OS cités — Windows Server 2008, MySQL 5.5 — datent de cette édition :
**à vérifier sur la documentation officielle** pour les prérequis
actuels, qui ont évolué).

**Édition Compact (jusqu'à 20 nœuds) :**

| Palier | CPU | Mémoire | Disque |
|---|---|---|---|
| 0 – 20 nœuds | Dual-core 2 GHz ou + | 4 Go | 40 Go |

**Édition Standard :**

| Équipements gérés | CPU | Mémoire | Disque |
|---|---|---|---|
| 0 – 200 | 1 × dual-core 2 GHz ou + | 4 Go | 60 Go |
| 200 – 500 | 2 × dual-core 2 GHz ou + | 4 Go | 60 Go |
| 500 – 2 000 | 2 × quad-core 2 GHz ou + | 8 Go | 120 Go |
| 2 000 – 5 000 | 2 × quad-core 2 GHz ou + | 16 Go | 250 Go |

Le constructeur recommande des **serveurs PC** (serveurs x86 standards).

## 17. Lire ce tableau en chef de service : les règles de marge

Le tableau donne des **minimums**. En pratique :

