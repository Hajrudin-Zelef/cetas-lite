---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-11
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [883, 991]
sha256: a8f91eaaaf7197b985143dfbde6658904afd00bd8d2b5b3a46137ed3cfa14488
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

## 59. Ajouter un AR720 — procédure détaillée

1. **Pré-déclaration** : AR720, ESN, site (souvent **routeur de branche** = point d'entrée du site).
2. **Amorçage** : l'AR720 est fréquemment le **serveur DHCP** du site — paradoxe apparent : il faut une configuration minimale (WAN + DHCP + route vers NCE) avant le ZTP des autres équipements. Deux approches : (a) template AR poussé via ZTP avec option DHCP intégrée, (b) pré-configuration minimale manuelle puis bascule sous NCE.
3. **Enregistrement** : comme les autres (ESN, site).
4. **Template de branche** : WAN (PPPoE/DHCP/fibre selon l'accès), VPN vers le siège (si SD-WAN de branche), DHCP avec option NCE pour les équipements aval, DNS, NTP.
5. **Vérification** : connectivité WAN, tunnel VPN monté, DHCP distribué (tester avec un PC), équipements aval qui s'enregistrent via lui.
6. **Secours** : lien de backup (4G/5G si prévu) testé par bascule réelle.

## 60. Vérifier qu'un équipement est bien géré — checklist

- [ ] L'équipement apparaît dans le **bon site** avec le bon nom
- [ ] État **normal** (ni alarme, ni hors ligne, ni non enregistré)
- [ ] Version logicielle = version cible du site
- [ ] Template appliqué : **conformité vérifiée** (get-config vs attendu, aucun écart inexpliqué)
- [ ] Supervision : télémétrie reçue (CPU, interfaces), alarmes testées (débrancher/rebrancher un port et vérifier l'alarme)
- [ ] Management : ping/SSH depuis NCE (ou via NCE), SNMPv3 fonctionnel
- [ ] Fonctionnel : port d'accès/SSID/VPN testé avec un client réel
- [ ] Documentation : étiquette physique, plan, inventaire à jour

## 61. États des équipements dans NCE (normal, alarme, hors ligne, non enregistré)

La documentation de l'écosystème NCE (API NBI) illustre les états suivis par le contrôleur :

- **Normal (0)** : géré, joignable, sans alarme bloquante.
- **Alarme (1)** : géré mais avec alarme(s) en cours — l'équipement fonctionne mais un problème est signalé (température, port down, etc.).
- **Hors ligne (3)** : ne répond plus au contrôleur — à diagnostiquer (réseau ? alimentation ? panne ?).
- **Non enregistré (4)** : découvert ou déclaré mais pas encore pris en gestion (onboarding en cours ou en échec).

En exploitation : un tableau de bord « par état et par site » est le premier écran du matin. Tout équipement « hors ligne » > 15 min = ticket. Tout « non enregistré » qui stagne = onboarding à reprendre (cas pratiques 1-5).

---

# PARTIE 7 — GESTION DU RÉSEAU FILAIRE

## 62. Templates de configuration — concepts

Le **template** est le cœur de l'automatisation NCE : un modèle de configuration avec des **variables**, appliqué à un site, un groupe ou un équipement. Exemple : le template « Switch d'accès standard » contient les VLAN, le SNMP, le NTP, le 802.1X — avec des variables comme `{{VLAN_GESTION}}`, `{{NOM_SITE}}`.

Cycle de vie d'un template :

1. **Conception** : écrire le template à partir d'une configuration de référence validée (un switch qui marche, pas une théorie).
2. **Variables** : identifier ce qui change par site/équipement (VLAN, noms, adresses).
3. **Test** : appliquer en maquette, vérifier la conformité, tester les services.
4. **Versionnement** : chaque modification = nouvelle version (jamais de modification « à la main » en production).
5. **Déploiement** : par vagues, avec fenêtre de maintenance et rollback prêt (sections 84/87).

Règle d'or : **NCE est la source de vérité**. Toute modification faite en CLI hors NCE sera détectée comme écart (configuration consistency) — et c'est voulu.

## 63. Variables dans les templates — paramétrage par site/équipement

Les variables permettent un template unique pour des sites différents :

| Variable (exemple) | Portée | Exemple de valeur |
|---|---|---|
| `{{NOM_SITE}}` | Site | SIEGE-BAT-A |
| `{{VLAN_GESTION}}` | Site | 10 |
| `{{VLAN_USERS}}` | Site | 20 |
| `{{VLAN_VOIP}}` | Site | 30 |
| `{{VLAN_IOT}}` | Site | 40 |
| `{{NTP_SERVEUR}}` | Global | 192.168.100.10 |
| `{{NOM_SWITCH}}` | Équipement | SW-ACC-01 |
| `{{PORT_UPLINK}}` | Équipement | XGigabitEthernet0/0/1 |

Bonnes pratiques : nommer les variables en **majuscules explicites**, documenter chaque variable (description, exemple, obligatoire/facultatif), valider les valeurs (plages VLAN 1-4094, formats IP) avant déploiement, et **ne jamais** mettre de mot de passe en clair dans un template (utiliser le coffre/secret store de NCE si disponible, sinon procédure sécurisée).

## 64. VLAN — conception et déploiement via NCE

Conception type d'un site (exemple — à adapter) :

| VLAN | Nom | Usage |
|---|---|---|
| 10 | GESTION | Management des équipements |
| 20 | USERS | Postes utilisateurs (802.1X) |
| 30 | VOIP | Téléphonie (voice VLAN) |
| 40 | IOT | Objets connectés (segmenté, filtré) |
| 50 | INVITES | Wi-Fi invités (portail captif) |
| 60 | SERVEURS | Serveurs locaux du site |
| 99 | NATIVE/TRANSPORT | Interconnexions (à sécuriser) |

Déploiement via NCE : les VLAN sont définis dans le template du site (ou la politique), poussés en NETCONF sur les switches, vérifiés par conformité. Le **VLAN de management** doit exister sur tous les équipements et être routé vers NCE. Erreurs classiques : oublier le VLAN sur un trunk d'uplink (site isolé), VLAN voice mal configuré (téléphones sans réseau), VLAN 1 laissé en natif partout (mauvaise pratique de sécurité).

## 65. Politiques d'accès (802.1X) — principes

Le **contrôle d'accès réseau (NAC)** via 802.1X : avant de donner l'accès réseau à un équipement branché sur un port, le switch **authentifie** l'utilisateur/la machine. C'est le « badge d'entrée » du réseau filaire.

Pourquoi c'est structurant avec NCE : déployer 802.1X à la main sur 200 ports = des semaines et des erreurs ; via NCE : une politique définie une fois, poussée partout, avec des **autorisations dynamiques** (VLAN/ACL/QoS selon qui se connecte). C'est typiquement le projet qui **justifie** NCE.

Prérequis : serveur **RADIUS** (le composant d'authentification NCE ou un RADIUS externe), annuaire (AD/LDAP) ou base locale, supplicants configurés sur les postes (natif Windows/macOS), plan de **dérogations** (imprimantes, équipements sans 802.1X → MAB, voir section 68).

## 66. 802.1X : rôles (supplicant, authentificateur, serveur)

Trois rôles, à connaître par cœur pour dépanner :

1. **Supplicant** : le poste/l'équipement qui veut accéder au réseau (le PC avec son client 802.1X).
2. **Authentificateur** : le switch (S310) — il relaie les échanges EAP entre le supplicant et le serveur, et **applique** le résultat (autorise/bloque le port, assigne le VLAN).
3. **Serveur d'authentification** : le RADIUS (NCE ou externe) — il **décide** (vérifie les identifiants dans l'annuaire) et renvoie les attributs d'autorisation (VLAN, ACL, etc.).

Séquence : le port s'ouvre en « non autorisé » → EAPOL entre supplicant et switch → le switch relaie en RADIUS vers le serveur → décision → le switch applique (VLAN dynamique, ACL) et ouvre le port. En cas d'échec : VLAN de quarantaine ou blocage selon la politique.

## 67. 802.1X : méthodes EAP courantes

| Méthode | Authentification | Sécurité | Usage |
|---|---|---|---|
| **EAP-TLS** | Certificats (client + serveur) | Très forte | Postes gérés, le top sécurité (nécessite une PKI) |
| **PEAP-MSCHAPv2** | Identifiant/mot de passe (tunnel TLS) | Bonne | Le plus courant en entreprise (AD) |
| **EAP-TTLS** | Identifiant/mot de passe ou certificat | Bonne | Alternative, souple |
| **EAP-MD5** | Mot de passe (faible) | Faible | À éviter |

