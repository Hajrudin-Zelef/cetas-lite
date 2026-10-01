---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-23
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["capex", "cost", "distribution", "license", "sandbox", "training"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [2035, 2113]
sha256: c64babc1c95e1de638fa0cd7ec846e591a22a87f3276bb9ef94610224fe3d152
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **802.1X** : standard de contrôle d'accès réseau par authentification (EAP) avant ouverture d'un port / d'une association Wi-Fi.
- **AAA** : Authentication, Authorization, Accounting — les trois fonctions de la gestion des accès.
- **ADN (Autonomous Driving Network)** : concept Huawei de réseau à gestion autonome (niveaux 1 à 5, du pilotage assisté à l'autonomie complète — vision).
- **API NBI (Northbound Interface)** : interface RESTful exposée par NCE vers les applications (portail, ITSM, scripts) — port documenté 18002, authentification par token `X-ACCESS-TOKEN`.
- **Authentificateur** : en 802.1X, l'équipement (switch/WAC) qui relaie l'authentification et applique la décision.
- **BYOD (Bring Your Own Device)** : usage d'équipements personnels sur le réseau d'entreprise.
- **CampusInsight** : composant d'analyse big data/ML de NCE-Campus (expérience utilisateur, prédiction, optimisation) — souscription séparée.
- **CAPWAP** : protocole de contrôle et de provisionnement des AP sans-fil (tunnels WAC ↔ AP).
- **Composant d'authentification** : service NCE déployable sur les branches (jusqu'à 20) pour l'authentification locale en cas de perte du lien central.
- **DCA (Dynamic Channel Assignment)** : ajustement automatique des canaux radio des AP (RRM).
- **Device-day** : unité de licence NCE = un équipement géré pendant un jour ; licences poolées, grâce de 30 jours en cas de dépassement.
- **EAP** : Extensible Authentication Protocol — cadre des méthodes d'authentification 802.1X (TLS, PEAP, TTLS...).
- **ESDP** : plateforme Huawei de distribution des logiciels/licences (commande et téléchargement).
- **ESN** : Electronic Serial Number — identifiant unique d'un équipement, clé de l'enregistrement ZTP.
- **Intent-based networking** : pilotage par intention (« qui accède à quoi ») traduit automatiquement en configurations.
- **MAB (MAC Authentication Bypass)** : authentification par adresse MAC pour les équipements sans 802.1X (faible sécurité — liste blanche).
- **MSP** : Managed Service Provider — prestataire gérant plusieurs tenants depuis un NCE mutualisé (licences réallouables).
- **NAC (Network Access Control)** : contrôle d'accès au réseau (802.1X, MAB, portail).
- **NETCONF** : protocole de configuration réseau structurée (RFC 6241) ; **YANG** : langage de modélisation des données (RFC 7950).
- **NQA (Network Quality Analysis)** : sondes actives des équipements Huawei pour mesurer la qualité (base du SLA management).
- **OPEX / CAPEX** : dépenses d'exploitation / d'investissement.
- **PPSK (Private Pre-Shared Key)** : clé pré-partagée individuelle (par utilisateur/groupe) avec VLAN/ACL dédiés — alternative au PSK partagé.
- **Qos (Quality of Service)** : mécanismes de priorisation du trafic (DiffServ, DSCP, policing).
- **RBAC (Role-Based Access Control)** : gestion des droits par rôles.
- **RRM (Radio Resource Management)** : gestion automatique des ressources radio (DCA, TPC).
- **SBI (Southbound Interface)** : protocoles entre le contrôleur et les équipements (NETCONF, SNMP, CAPWAP...).
- **SLA** : Service Level Agreement — engagement de qualité de service, mesuré ici via NQA.
- **SNMPv3** : version sécurisée de SNMP (authentification + chiffrement — `authPriv`).
- **SSID** : Service Set Identifier — nom du réseau Wi-Fi diffusé par les AP.
- **TCO (Total Cost of Ownership)** : coût total de possession sur une durée (licences + matériel + services + exploitation).
- **Télémétrie** : collecte de données d'état/performance des équipements (polling SNMP ou streaming HTTP/2).
- **Template** : modèle de configuration versionné avec variables, appliqué par NCE aux équipements.
- **TPC (Transmit Power Control)** : ajustement automatique de la puissance d'émission des AP.
- **VLAN** : Virtual LAN — segmentation logique du réseau local.
- **WAC** : Wireless Access Controller — contrôleur des AP (tunnels CAPWAP).
- **ZTP (Zero-Touch Provisioning)** : provisionnement automatique d'un équipement neuf (DHCP/e-mail/scan), sans intervention experte locale.

---

# PARTIE 18 — POUR ALLER PLUS LOIN

## 159. Pour aller plus loin — ressources et prochaines étapes

**Ressources documentaires (à consulter sur la version exacte cible)** :
- Product Overview iMaster NCE-Campus (ex : V300R022C10) — architecture et fonctions par version.
- Monitoring and O&M Guide (ex : V300R020C10) — upgrade, licences devices, signature database, SLA.
- License Usage Guide (ex : V300R022C00/V300R024C00, MSP Training Manual) — modèle device-day, mode MSP.
- Datasheet iMaster NCE-CampusInsight (ex : V100R025C00) — souscriptions d'analyse.
- Guide d'installation et guide de déploiement HA de la version cible — **la référence** pour le dimensionnement.
- Matrice de compatibilité NCE × équipements — **à exiger du partenaire** avant tout achat.
- Outil EOM/EOFS/EOS Query (e.huawei.com) — vérifier le statut de support d'eSight et des équipements.
- Écosystème développeur Huawei — API NBI, sandbox, plus de 500 API documentées.

**Prochaines étapes proposées pour Zelef** :
1. **Cadrage** : périmètre (quels sites ?), inventaire précis, TCO 5 ans eSight vs NCE avec le partenaire.
2. **Maquette** : 1 switch S310 + 2 AP361 + NCE en lab — valider ZTP, template, 802.1X, supervision.
3. **Formation** : au moins le référent (et un suppléant) formés avant le pilote.
4. **Pilote** : 1 site représentatif non critique, 3 mois, REX écrit, go/no-go.
5. **Industrialisation** : vagues par site, documentation d'exploitation, intégration Zabbix/GLPI via API.
6. **Run** : revues trimestrielles (licences, rôles, rapports), exercice de restauration annuel, veille versions.

**En lien avec tes autres guides** : ce document complète ton guide eSight (supervision transverse), tes guides AP361/AP761/S310/AR720/USG6000 (les équipements pilotés), ton guide onduleurs (l'énergie qui fait tourner tout ça — NCE ne la supervise pas, Zabbix/eSight si), et ton guide eKit (le petit site). L'architecture cible saine : **NCE** pour le campus Huawei automatisé, **eSight/Zabbix** pour le transverse et l'énergie, **eKit** pour les petits sites simples — chacun à sa place, intégrés par API là où ça compte.

---

# PARTIE 19 — APPROFONDISSEMENTS OPÉRATIONNELS

## 160. VXLAN et virtualisation du campus — quand NCE fait du réseau overlay

Au-delà du VLAN classique, NCE-Campus peut déployer des **réseaux virtuels VXLAN** : la documentation Huawei décrit le provisionnement automatisé de VN (Virtual Networks) — « un réseau multifonctionnel, services VXLAN provisionnés en minutes ».

Principe : plutôt que d'étendre des VLAN sur tout le campus (avec les limites du spanning-tree et des domaines de broadcast), on crée des **overlays VXLAN** avec des **VNI** (identifiants de segment), routés/bridgés par les équipements (VTEP). Cas d'usage : segmentation forte entre populations (ex : réseau « production », réseau « invités », réseau « IoT ») avec des politiques de sécurité entre VN.

Recommandation pragmatique : ne pas mettre du VXLAN partout « parce que c'est moderne ». Le VLAN + VRF reste plus simple et suffisant pour la plupart des campus. Réserver VXLAN aux cas où la segmentation doit traverser plusieurs sites/équipements avec des politiques centralisées — et le **tester en maquette** (le dépannage d'un overlay est plus exigeant qu'un VLAN).

## 161. SD-WAN de branche avec l'AR720 — le rôle WAN de NCE

NCE-Campus ne gère pas que le LAN : il orchestre aussi le **WAN des branches** (les AR720 de Zelef sont typiquement des routeurs de branche).

