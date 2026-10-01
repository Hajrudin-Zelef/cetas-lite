---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-5
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [598, 766]
sha256: 9c26c65c189b426a66e54da115cd9a7568a8870da1d46465b35cb542e93ca3e2
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

- [ ] Mot de passe admin initial changé + comptes nominatifs créés
      (section 11) ; le compte admin générique n'est plus utilisé
      au quotidien.
- [ ] HTTPS forcé, certificat remplacé par un certificat d'entreprise
      (le certificat auto-signé d'installation fait hurler les navigateurs
      et habitue les exploitants à cliquer « accepter le risque » —
      une mauvaise habitude de sécurité).
- [ ] Comptes de service SNMP v3 créés (jamais de v2c avec community
      « public »/« private » — voir section 5).
- [ ] Politique de verrouillage : tentatives de login, complexité des
      mots de passe, délai d'expiration de session.
- [ ] Journalisation d'audit activée et vérifiée (qui s'est connecté,
      qui a acquitté quoi — voir section 11).
- [ ] Sauvegarde initiale réalisée et testée (section 12).

## 25. L'outil Environmental Health Check : votre premier réflexe

eSight fournit un **outil de contrôle de santé de l'environnement**
(accessible depuis l'icône eSight Console sur le bureau du serveur →
Tools → Environmental Health Check). D'après la documentation, il vérifie
les paramètres de base de l'OS, de la base de données et du serveur.

**En faire un rituel :**
- Après chaque installation / mise à jour : contrôle systématique.
- En cas de comportement anormal d'eSight (lenteurs, services qui
  tombent) : le lancer **avant** de toucher à quoi que ce soit —
  il pointe souvent la cause (disque plein, mémoire, base injoignable).
- Intégrez-le à la check-list mensuelle (section 16).

---

# 5. DÉCOUVERTE DES ÉQUIPEMENTS

## 26. Principe de la découverte

La découverte = faire connaître à eSight les équipements à superviser.
Deux modes complémentaires :

- **Découverte automatique** : eSight balaie des plages IP (ou un segment),
  interroge chaque adresse en SNMP/ICMP, identifie le type d'équipement
  et l'ajoute à l'inventaire.
- **Ajout manuel** : saisie d'un équipement (IP, profil SNMP) — utile pour
  les équipements isolés, les DMZ, ou quand l'auto-découverte est interdite.

La découverte alimente ensuite **l'inventaire (Resource)**, la **topologie**
et la **collecte d'alarmes/performance**. Tant qu'un équipement n'est pas
découvert *et* correctement paramétré en SNMP, il est invisible : c'est
l'étape fondatrice, à ne pas bâcler.

## 27. Préparer les équipements Huawei en SNMP v2c

Côté équipement (exemple de logique sur un switch Huawei — **syntaxe à
adapter à votre modèle/version**) :

- Activer l'agent SNMP.
- Définir une **community de lecture** dédiée à eSight (exemple fictif :
  `SupervisionLecture2026`).
- Si eSight doit modifier la config via SNMP (rare ; préférer SSH),
  définir aussi une community d'écriture — sinon, lecture seule suffit.
- **Restreindre par ACL** : n'autoriser que l'adresse IP du serveur eSight
  à interroger l'agent SNMP.
- Déclarer eSight comme **destination des traps** (trap-target vers
  l'IP d'eSight, port UDP 162).
- Vérifier que le firewall local de l'équipement (s'il existe) autorise
  SNMP depuis eSight.

💡 Côté eSight, on crée ensuite un **profil SNMP v2c** contenant cette
community, réutilisé pour toute la découverte.

## 28. Préparer les équipements en SNMP v3 (recommandé)

SNMP v3 apporte authentification et chiffrement. Côté équipement :

- Créer un **utilisateur SNMP v3** dédié à eSight (exemple fictif :
  `esight-ro`).
- Niveau de sécurité **`authPriv`** : protocole d'authentification
  (SHA de préférence à MD5) + protocole de chiffrement (AES de
  préférence à DES).
- Mots de passe robustes distincts pour auth et priv, consignés au coffre.
- Vue MIB : restreindre si besoin aux branches utiles (surtout en lecture).
- Déclarer eSight comme destinataire des **inform/traps v3**.

Côté eSight : créer le **profil SNMP v3** miroir (même utilisateur,
mêmes protocoles, mêmes mots de passe). **Le moindre écart (SHA vs MD5,
mauvais password) = échec silencieux de la découverte** : c'est la
première chose à vérifier quand « SNMP ne répond pas » (cas n°2,
section 14).

## 29. Créer les profils de découverte dans eSight

Un **profil** regroupe les paramètres d'accès (version SNMP, community ou
utilisateur v3, timeout, retries, paramètres SSH/Telnet pour la gestion
de configuration). Marche à suivre type :

1. Menu Resource/Discovery (**libellé à vérifier sur la documentation
   officielle**) → gestion des profils.
2. Créer un profil par **couple (version SNMP × population d'équipements)** :
   ex. `SNMPv3-Huawei-authPriv`, `SNMPv2c-legacy`, `SNMPv2c-imprimantes`.
3. Régler **timeout** (2-3 s en LAN, plus en WAN) et **retries** (2-3) :
   trop courts = équipements déclarés injoignables à tort ; trop longs =
   découverte interminable.
4. Associer les credentials SSH/Telnet si la sauvegarde de configuration
   est prévue (section 9).

**Bonne pratique** : nommer les profils de façon explicite
(`<proto>-<usage>-<périmètre>`) et documenter dans le dossier
d'exploitation quel profil s'applique à quel parc.

## 30. Lancer une découverte par plages IP

1. Définir les **plages IP** à balayer (ex. : `10.20.0.0/16` pour le LAN,
   en excluant les plages DHCP des postes utilisateurs pour aller vite).
2. Associer le(s) profil(s) SNMP à utiliser, dans l'ordre d'essai.
3. Choisir la **profondeur** : découverte simple (les équipements de la
   plage) vs avec propagation vers les voisins (via LLDP/topologie).
4. Lancer en **fenêtre de maintenance** la première fois (un balayage
   SNMP sur un /16, ça se remarque sur les sondes IDS et ça charge
   les équipements).
5. Suivre l'avancement, puis **contrôler le résultat** : nombre
   d'équipements découverts vs inventaire attendu (section 31).

💡 Découpez les plages par **site/fonction** (une tâche de découverte par
bâtiment ou par VLAN de management) : en cas d'échec, on sait où chercher,
et on peut planifier des redécouvertes ciblées.

## 31. Vérifier que la découverte est complète et saine

Check-list post-découverte :

- [ ] **Quantitatif** : nb d'équipements dans eSight ≈ nb attendu
      (écart > 5 % = enquête).
- [ ] **Qualitatif** : chaque équipement a son **sysName, son type/modèle
      et sa version** correctement identifiés (pas de « unknown device »
      en masse).
- [ ] **Test SNMP** : depuis eSight, un walk de contrôle sur quelques
      équipements (réponse < 1 s en LAN).
- [ ] **Traps** : générer un trap de test (ex. : `linkDown` simulé ou
      reboot d'un équipement de labo) et vérifier sa réception dans
      eSight en moins d'une minute.
- [ ] **Syslog** : vérifier la réception des logs si configurée.
- [ ] **Doublons** : traquer les équipements découverts deux fois
      (multi-homing, plusieurs IP) et fusionner/nettoyer.

## 32. Découverte des équipements tiers (non-Huawei)

eSight gère des équipements de grands constructeurs tiers (HP, Cisco
cités par le constructeur) et des ressources IT (serveurs, imprimantes)
via les **MIB standard** (MIB-II, ENTITY-MIB, IF-MIB…).

- Attendez-vous à une supervision **« générique »** : état up/down,
  interfaces, CPU/mémoire si les MIB standard les exposent — mais pas
  les subtilités propriétaires (d'où la note moyenne hors Huawei
  dans les retours d'exploitants).
- Pour un équipement tiers critique, testez **avant** de généraliser :
  découverte, traps, et 2-3 indicateurs de performance.
- Si un équipement n'est pas reconnu, eSight permet de **personnaliser
  la gestion des équipements tiers** (adaptateurs MIB) — chantier
  avancé, à confier à quelqu'un qui connaît les MIB.

---

# 6. TOPOLOGIE

## 33. Les vues de topologie : la carte du réseau

La topologie est la **vue d'accueil** d'eSight : une carte graphique des
équipements et de leurs liens, avec l'état (couleur selon les alarmes).
C'est l'écran que le NOC garde affiché en permanence.

