---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-12
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [992, 1083]
sha256: 3454be8de7631f12c0c88824ef5c77f1aa8e605f08405bda91e01cabb5c8b553
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Recommandation : **PEAP-MSCHAPv2** adossé à l'AD pour démarrer (compromis simplicité/sécurité), trajectoire vers **EAP-TLS** pour les postes gérés (plus sûr, pas de mot de passe). Tester **chaque méthode** en maquette avec de vrais postes avant le déploiement (les supplicants ont leurs caprices : voir cas pratique 10).

## 68. MAC-auth et portail captif — les alternatives

Tout le monde ne parle pas 802.1X (imprimantes, caméras, vieux équipements, certains IoT) — prévoir les filets :

- **MAB (MAC Authentication Bypass)** : le switch authentifie l'équipement par son **adresse MAC** auprès du RADIUS. Pratique mais faible (MAC spoofable) → réserver aux équipements connus, avec **liste blanche** de MAC et VLAN dédié (ex : VLAN imprimantes, très filtré).
- **Portail captif** : l'utilisateur ouvre un navigateur, s'identifie sur une page (NCE peut porter le portail). Adapté aux **invités** et au BYOD léger.
- **Stratégie combinée type** : 802.1X en premier → si échec et MAC connue → MAB → sinon portail ou quarantaine. L'ordre et les délais se règlent dans la politique NCE.

Pour Zelef : les **copieurs** (son métier !) sont typiquement en MAB sur VLAN dédié — prévoir ce cas dès la conception 802.1X, pas après.

## 69. Autorisation dynamique : VLAN, ACL, QoS par utilisateur

La vraie puissance du 802.1X piloté par NCE : l'autorisation **change selon qui se connecte** :

- **VLAN dynamique** : le RADIUS renvoie `Tunnel-Private-Group-ID` = le VLAN de l'utilisateur (compta → VLAN 21, technique → VLAN 22). Un même port physique sert à tout le monde, en sécurité.
- **ACL dynamique** : règles de filtrage poussées à la session (ex : les invités ne voient qu'Internet).
- **QoS dynamique** : profil de débit/priorité selon le profil (ex : la visio des dirigeants priorisée).

Tout cela se définit dans les **politiques d'autorisation** de NCE (qui → quoi → sous quelles conditions : groupe AD, horaire, type d'équipement, site). Conseil : commencer avec 3-4 profils simples (employé, invité, IoT, quarantaine), complexifier ensuite — une matrice d'autorisation illisible est une matrice non maintenable.

## 70. QoS filaire — modèles et déploiement

La QoS garantit que le trafic important passe même quand le réseau est chargé (ex : la visio ne doit pas être hachée parce que quelqu'un télécharge).

Modèle type (DiffServ) :

| Classe | DSCP | Usage | Traitement |
|---|---|---|---|
| Voix | EF (46) | Téléphonie IP | Priorité stricte, faible latence |
| Visio | AF41 (34) | Visioconférence | Priorité élevée, bande garantie |
| Données critiques | AF31 (26) | Applis métier | Bande garantie |
| Données standard | AF21/BE (0) | Bureautique, web | Best effort |
| Invités/IoT | BE (0) + policing | Non critique | Débit plafonné |

Déploiement via NCE : classification (à l'entrée, selon 802.1X/ACL), marquage, files de priorité et **policing/shaping** sur les switches — le tout dans le template du site. Vérification : générer de la charge en maquette et mesurer (pas de QoS « à l'aveugle »). Erreur classique : marquer sans faire confiance au marquage en amont, ou policer trop bas et casser une appli métier (voir cas pratique 6).

---

# PARTIE 8 — GESTION WLAN CENTRALISÉE

## 71. Gestion WLAN centralisée — principes

Avec NCE-Campus, le Wi-Fi n'est plus configuré AP par AP (ou sur un WAC isolé) : il est **orchestré** comme le filaire — profils radio, SSID et politiques définis centralement, poussés aux WAC/AP, supervisés en unifié.

Bénéfices : cohérence (même SSID partout), déploiement rapide (un nouveau site = un profil), supervision croisée (un problème Wi-Fi visible avec son contexte filaire), et automatisation (ZTP des AP, mises à jour groupées).

Architecture : NCE-Campus → (orchestration) → **WAC** (contrôleur WLAN) → (CAPWAP) → **AP** (AP361/AP761). NCE ne remplace pas le WAC, il le **pilote** et l'intègre à la gestion unifiée.

## 72. Le modèle WAC + AP : comment NCE pilote le Wi-Fi

- Le **WAC** centralise : tunnels CAPWAP vers les AP, gestion des associations, politiques radio de base, authentification 802.11.
- **NCE-Campus** apporte au-dessus : les **profils** (radio, SSID, sécurité), le **déploiement** (ZTP des AP via le WAC), la **supervision unifiée** (le Wi-Fi dans la même topologie que le filaire), les **politiques d'accès** (802.1X Wi-Fi, PPSK, portail) et l'**analyse** (CampusInsight).
- L'authentification Wi-Fi 802.1X transite par le RADIUS (composant NCE ou externe) comme en filaire — **même annuaire, mêmes groupes**, cohérence totale.

Point de vigilance : les fonctions avancées (identification de terminaux, statistiques applicatives, certaines authentifications) exigent des équipements **V5** comme dispositifs d'accès selon la documentation — vérifier la matrice par modèle avant de s'engager.

## 73. Profils radio — conception (canaux, puissance, bandes)

Un **profil radio** définit comment les AP émettent. Conception type :

- **Bandes** : 2,4 GHz (portée, compatibilité — mais encombrée) et 5 GHz (débit, capacité) ; 6 GHz si Wi-Fi 6E/7 et modèles compatibles (vérifier la réglementation locale — **à vérifier**).
- **Largeur de canal** : 20 MHz en 2,4 GHz (toujours) ; 40/80 MHz en 5 GHz selon la densité (plus large = plus de débit mais moins de canaux non chevauchants).
- **Puissance** : **moins fort que le maximum** en intérieur dense (une puissance trop forte = interférences entre AP et clients qui s'accrochent au mauvais AP). Point de départ : 14-17 dBm en 5 GHz en bureau dense, à ajuster par mesure.
- **Canaux 2,4 GHz** : 1, 6, 11 uniquement (les seuls non chevauchants).
- **Canaux 5 GHz** : plan en quinconce, en évitant les canaux DFS si les coupures radar sont problématiques (ou en les assumant).

Le profil se définit une fois dans NCE et s'applique par site/type de zone (bureaux, entrepôt, extérieur). La **planification prédictive** (étude de couverture) reste recommandée pour les zones critiques — NCE optimise, il ne remplace pas une étude.

## 74. SSID — conception et bonnes pratiques

Règles d'un plan SSID propre :

- **Peu de SSID** : chaque SSID consomme de l'airtime (beacons). Au-delà de 3-4 SSID par AP, la performance se dégrade — regrouper par **profils d'accès** plutôt que par service.
- Exemple : `ENTREPRISE` (802.1X, employés), `INVITES` (portail captif, isolé, débit plafonné), `IOT` (PPSK ou 802.1X, VLAN dédié). Éventuellement un SSID `DIRIGEANTS`/spécifique si justifié.
- **Ne pas diffuser le SSID** sensible en clair partout (ou le restreindre par zone/AP).
- **Bande steering** : orienter les clients capables vers le 5 GHz.
- **Sécurité** : WPA3-SAE ou WPA2-Enterprise (jamais de WPA2-PSK partagé pour les employés — le PSK partagé est un secret qui fuit).
- **VLAN par SSID** (ou dynamique par utilisateur) : séparer les flux dès l'AP.

Dans NCE : les SSID se définissent dans les profils WLAN du site, avec la sécurité, le VLAN et les politiques associées — déploiement atomique sur tous les AP du site.

## 75. Planification radio : RRM, DCA, TPC

Le **RRM** (Radio Resource Management) ajuste automatiquement la radio :

- **DCA** (Dynamic Channel Assignment) : change les canaux des AP pour minimiser les interférences, en fonction des mesures.
- **TPC** (Transmit Power Control) : ajuste les puissances d'émission (baisse quand les AP se voient trop, monte pour couvrir un trou).

Avec NCE/CampusInsight, ces algorithmes s'enrichissent de l'historique (ML) : Huawei cite +50 % de performance Wi-Fi après optimisation dans un test Tolly — prendre comme un **potentiel**, pas une promesse contractuelle.

