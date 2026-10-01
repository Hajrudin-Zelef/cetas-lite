---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-15
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [2227, 2373]
sha256: c6da84ce3536494fe00c5022511d5f6292ebfcf7ec93b50410d6ca39c50f18ac
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**R1.** Fault management (alarmes), Performance management (indicateurs,
tableaux de bord, rapports), Topology management (cartographie),
Configuration management (sauvegarde/comparaison/restauration des configs).
**R2.** L'édition **Standard** (WLAN management et NTA/NetStream y sont
inclus ; la Compact ne les a pas).
**R3.** v2c utilise une community string en clair (authentification faible,
pas de chiffrement) ; v3 apporte authentification et chiffrement
(authPriv). **Privilégier v3 partout où c'est possible.**
**R4.** (1) L'agent SNMP est-il activé sur l'équipement ? (2) La
community / l'utilisateur v3 correspond-il au profil eSight (tester au
snmpwalk) ? (3) Une ACL côté équipement bloque-t-elle l'IP d'eSight ?
**R5.** Acquitter = prendre en charge (« je m'en occupe », avec commentaire
et ticket). Ça n'éteint PAS le défaut et ne le résout pas.
**R6.** Agrégation : regrouper les alarmes répétitives identiques pour
réduire le bruit. Masquage : taire temporairement des alarmes connues
(maintenance). Corrélation : rattacher les alarmes conséquences à la
cause racine.
**R7.** Parce qu'un masquage permanent rend la panne invisible : sans date
de fin, on oublie le masquage et on ne voit plus jamais les alarmes de
cet équipement.
**R8.** La base de données (outil Database Backup & Restore), le système
eSight, les fichiers/actes de licence, et le dossier d'exploitation —
le tout externalisé, avec un test de restauration annuel.
**R9.** Ne pas tout acquitter en panique : trier par heure, identifier la
première alarme (cause racine probable), s'appuyer sur la corrélation,
traiter la cause — les conséquences se solderont en cascade.
**R10.** Réponse honnête : eSight est **toujours commercialisé** (version
23.1 en mars 2024) et reste pertinent pour superviser un parc existant ;
iMaster NCE est la **plateforme stratégique d'avenir** de Huawei
(SDN, automatisation, IA). Aucune date officielle de fin de vie d'eSight
n'est publiée (**à vérifier sur la documentation officielle** / auprès
du partenaire) : on investit dans eSight en gardant des processus
indépendants de l'outil et une feuille de route 3-5 ans vers NCE.

---

# ANNEXE A — RECETTE DE MISE EN SERVICE

## 116. Recette : périmètre et méthode

La **recette** (acceptation) est la phase qui valide que l'installation
d'eSight est conforme au besoin avant de la déclarer « en production ».
Sans recette écrite et signée, les oublis se découvrent en incident.

**Méthode :**
- Recette **par scénarios** (pas par cases à cocher abstraites) : chaque
  scénario décrit une action, le résultat attendu et le résultat constaté.
- Recette **contradictoire** : l'intégrateur déroule, l'équipe interne
  valide et signe.
- **Critère GO/NO-GO** : tout scénario « bloquant » en échec = pas de
  mise en production.

## 117. Recette : infrastructure eSight

| # | Scénario | Résultat attendu |
|---|---|---|
| R-INF-01 | Redémarrer le serveur eSight | Tous les services remontent seuls en < 15 min, console accessible |
| R-INF-02 | Environmental Health Check | 100 % des contrôles au vert |
| R-INF-03 | Couper le réseau 2 min puis rétablir | eSight resynchronise seul, pas d'alarmes fantômes persistantes après synchro manuelle |
| R-INF-04 | Tester l'URL depuis un poste NOC, un poste VPN, une tablette | HTTPS OK partout, certificat reconnu (pas d'avertissement) |
| R-INF-05 | Simuler un disque à 85 % (ou vérifier les seuils) | Alerte espace disque remontée |
| R-INF-06 | Vérifier NTP | Écart < 1 s avec la référence |

## 118. Recette : découverte et supervision

| # | Scénario | Résultat attendu |
|---|---|---|
| R-DISC-01 | Découverte d'un switch pilote en SNMP v3 | Inventaire correct : modèle, version, interfaces |
| R-DISC-02 | Provoquer un down/up de port sur le pilote | Trap reçu en < 60 s, alarme visible avec la bonne sévérité |
| R-DISC-03 | Couper le polling (ACL test) 10 min | Alarme « SNMP timeout » puis retour à la normale au rétablissement |
| R-DISC-04 | Vérifier un équipement tiers (ex. imprimante) | A minima up/down + état générique |
| R-DISC-05 | Contrôler l'absence de doublons | Chaque équipement physique = 1 objet dans l'inventaire |
| R-DISC-06 | Sauvegarde de config du pilote | Fichier récupéré, contenu conforme à la config réelle |

## 119. Recette : alarmes et notifications

| # | Scénario | Résultat attendu |
|---|---|---|
| R-ALM-01 | Générer une alarme Critical de test | SMS reçu par l'astreinte en < 5 min avec équipement + libellé |
| R-ALM-02 | Générer une alarme Major de test | E-mail reçu par le groupe, pas de SMS |
| R-ALM-03 | Acquitter avec commentaire + ticket fictif | Alarme acquittée, commentaire visible, traçabilité en audit |
| R-ALM-04 | Simuler un flap (down/up × 5) | Agrégation : 1 alarme avec compteur (pas 10 alarmes) |
| R-ALM-05 | Activer un masquage de test avec date de fin | Pas de remontée pendant le masquage, retour auto après échéance |
| R-ALM-06 | Vérifier le filtrage northbound (si applicable) | Seules les sévérités configurées remontent au NMS supérieur |

## 120. Recette : sauvegardes, comptes et documentation

| # | Scénario | Résultat attendu |
|---|---|---|
| R-BKP-01 | Sauvegarde complète eSight | Fichier généré sur le stockage externe, taille cohérente |
| R-BKP-02 | Restauration sur VM isolée | Console fonctionnelle, inventaire et paramétrage retrouvés |
| R-ACC-01 | Connexion avec chaque rôle | Droits conformes à la matrice (un read-only ne peut pas acquitter) |
| R-ACC-02 | Couper l'AD (si intégré) | Compte local de secours fonctionnel |
| R-DOC-01 | Dossier d'exploitation livré | Schéma, IP, comptes (coffre), procédures, licences — complet et à jour |

**Signature** : recette signée par le chef de service et l'intégrateur,
archivée dans le dossier d'exploitation. C'est le « top départ » officiel.

---

# ANNEXE B — FICHES RÉFLEXES PAR TYPE D'ÉQUIPEMENT

> Que superviser en priorité sur chaque famille d'équipements, et avec
> quels seuils de départ. À affiner par modèle.

## 121. Fiche réflexe : switch d'accès

**À superviser :** état up/down, CPU, mémoire, température, alimentations,
utilisation des ports montants (uplinks), budget **PoE** utilisé/disponible,
erreurs sur les ports.

**Seuils de départ :** uplink > 70 % = warning (tendance) ; CPU > 80 %
persistant = major ; 1 alim sur 2 HS = major ; budget PoE > 90 % = warning
(plus de marge pour ajouter des AP/téléphones).

**Alarmes à ne jamais masquer :** uplink down, équipement injoignable.

## 122. Fiche réflexe : switch cœur / agrégation

**À superviser :** tout ce du switch d'accès, plus : état des membres
de stack/cluster, température par carte, protocoles de routage (voisins
OSPF/BGP), files QoS (drops), spanning-tree (changements de topologie).

**Seuils de départ :** CPU > 70 % = major (un cœur ne doit jamais être
tendu) ; perte d'un membre de stack = major ; changement de root STP =
warning à investiguer.

**Particularité :** chaque alarme sur le cœur est potentiellement un
début de tempête (section 46) — corrélation obligatoire vers les
équipements d'accès qui en dépendent.

## 123. Fiche réflexe : routeur / firewall

**Routeur — à superviser :** état des interfaces WAN, voisins de routage,
CPU (le routage + NetStream coûtent cher), mémoire, tunnels (VPN/IPSec :
état up/down en édition Standard).

**Firewall — à superviser :** CPU, mémoire, sessions (table de sessions
pleine = nouveaux flux rejetés), débit, état du cluster HA (bascule =
major), journaux d'attaques (à corréler, pas à alerter un par un).

**Seuils de départ :** sessions > 80 % de la capacité = warning ;
bascule HA = major + SMS ; tunnel VPN down = major (critical si c'est
le seul lien d'un site).

## 124. Fiche réflexe : AP et contrôleur Wi-Fi (AC)

