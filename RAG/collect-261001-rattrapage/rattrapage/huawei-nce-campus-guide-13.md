---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-13
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1084, 1189]
sha256: 4153b93cf9f34520c889733fa182cbbfb7fd03c9e69869771995a8141c1acdca
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Bonnes pratiques : laisser le RRM actif mais **borné** (plages de puissance, canaux autorisés), planifier les changements la **nuit** (un changement de canal en pleine journée = micro-coupures), surveiller les AP « exclus » du RRM (souvent un symptôme : AP mal placé ou en panne partielle).

## 76. Roaming — 802.11k/v/r et pratiques Huawei

Le **roaming** : un client qui se déplace passe d'un AP à l'autre sans couper sa session (critique pour la voix/vidéo sur Wi-Fi).

- **802.11k** : l'AP informe le client des AP voisins (liste de voisins) → le client choisit mieux et plus vite.
- **802.11v** : gestion des transitions (l'AP peut « suggérer » au client de changer d'AP pour équilibrer la charge).
- **802.11r** (Fast BSS Transition) : ré-authentification rapide (évite le cycle 802.1X complet à chaque changement d'AP) — **indispensable** pour la voix sur Wi-Fi.
- Pratiques : activer k/v/r sur le SSID voix, dimensionner le **recouvrement de couverture** (-65 dBm en bordure de cellule pour la voix), tester en marchant avec un appel réel (le seul test qui compte).

Sous NCE : ces options s'activent dans le profil SSID ; la supervision permet de voir les **échecs de roaming** par client (diagnostic fin). Voir cas pratique 9 (roaming dégradé).

## 77. Sécurité WLAN : WPA2/WPA3, 802.1X Wi-Fi, PPSK

- **WPA2-Enterprise (802.1X)** : la base saine pour les employés (EAP-TLS ou PEAP, voir section 67).
- **WPA3-SAE** : remplace le PSK (meilleure résistance aux attaques hors ligne) ; **WPA3-Enterprise** (192-bit) pour les environnements exigeants. Vérifier la compatibilité des clients (parc ancien → mode transition WPA2/WPA3).
- **PPSK (Private Pre-Shared Key)** : une **clé par utilisateur/groupe** (au lieu d'un PSK partagé) avec VLAN/ACL par clé — idéal pour l'IoT, les prestataires, les invités « connus ». NCE gère le cycle de vie des PPSK (création, expiration, révocation).
- **Isolation** : invités isolés entre eux (client isolation) et d'Internet uniquement ; IoT sur VLAN dédié sans accès au LAN.

Règle : **jamais de PSK unique partagé** pour un usage employé — c'est le PPSK ou le 802.1X.

## 78. NCE vs eKit pour le WLAN — quand NCE apporte un plus

eKit gère très bien un Wi-Fi simple (SSID, PSK, quelques AP). NCE apporte un plus quand :

- **802.1X Wi-Fi** avec annuaire d'entreprise (eKit ne le fait pas au même niveau).
- **Multi-sites** avec SSID cohérents et administration centrale.
- **PPSK** à grande échelle (hôtel, résidence, campus étudiant).
- **Analyse** : expérience par utilisateur, optimisation radio ML, rapports.
- **Invités** avec portail personnalisé, vouchers, traçabilité (obligations légales de journalisation — **à vérifier** selon le pays).
- **Roaming voix** finement réglé et supervisé.

En dessous de ces besoins : eKit est plus simple et moins cher. Le Wi-Fi d'une petite agence (3 AP361, 1 SSID PSK) ne justifie pas NCE.

## 79. Invités et BYOD — portail captif via NCE

Le **portail captif** : l'invité s'associe au SSID `INVITES`, ouvre son navigateur, arrive sur une page d'authentification (voucher, SMS, e-mail, sponsor interne, ou simple acceptation des CGU).

Mise en œuvre avec NCE :

- SSID dédié, VLAN invités, ACL restrictives (Internet uniquement, pas de LAN), **débit plafonné** (QoS), **durée de session limitée**.
- Page de portail personnalisée (logo, CGU) — le portail peut être porté par NCE ou externalisé (intégrateurs type Purple : NCE relaie le RADIUS vers la plateforme tierce).
- **Journalisation** : conserver les logs de connexion (exigences légales variables — **à vérifier** localement).
- **BYOD employés** : onboarding avec PPSK ou 802.1X (selon la maturité), VLAN dédié, posture check si exigé.

Pièges : un portail qui ne s'affiche pas (problème DNS/HTTPS — prévoir le contournement), des vouchers qui fuient (durées courtes, un par invité).

## 80. IoT sur le WLAN — segmentation

L'IoT (capteurs, caméras, badges, équipements techniques) est le point faible de la sécurité : des objets peu patchés, parfois avec des stacks réseau fragiles.

Architecture type sous NCE :

- **SSID/VLAN IoT dédié** (ou PPSK par famille d'objets), **jamais** sur le VLAN utilisateurs.
- **Filtrage strict** : l'IoT ne parle qu'à ses serveurs (ACL), pas à Internet en direct (sauf besoin), pas au LAN.
- **Authentification** : PPSK par lot ou 802.1X si l'objet le supporte ; MAB en dernier recours avec liste blanche.
- **Supervision** : les objets qui se taisent (capteur muet) = alarme — définir des seuils d'inactivité.
- **Cycle de vie** : inventaire des objets (qui ? où ? qui est responsable ?), plan de remplacement.

Pour un chef de service énergies : les objets techniques (sondes, automates) passent idéalement en **filaire** sur VLAN IoT — le Wi-Fi reste pour le mobile.

---

# PARTIE 9 — AUTOMATISATION

## 81. Automation — vue d'ensemble

L'automatisation NCE, c'est l'exploitation **à l'échelle** sans multiplier les administrateurs :

1. **ZTP** : onboarding sans toucher (sections 51-59, 82).
2. **Templates + variables** : configuration industrialisée (62-63, 83).
3. **Déploiement par vagues** : changements planifiés, progressifs, réversibles (84).
4. **Firmware en masse** : mises à jour groupées avec politique (85-86).
5. **Rollback** : retour arrière automatisé (87).
6. **API** : intégration aux outils de l'entreprise (88-91).

Philosophie : **tout ce qui est fait deux fois doit être automatisé** ; tout ce qui est automatisé doit être **testé, versionné et réversible**. L'automatisation sans ces trois garde-fous, c'est juste une façon plus rapide de casser plus grand (voir cas pratique 6).

## 82. ZTP détaillé — séquence complète d'un onboarding automatique

Reprenons la séquence ZTP en détail opérationnel (exemple : 20 AP361 pour une nouvelle agence) :

**J-7 — Préparation (au bureau)**
- Sites créés dans NCE, templates validés en maquette, firmware cible déposé sur le serveur de fichiers.
- ESN relevés à la réception et **pré-déclarés** dans le site (liste blanche).
- DHCP du site configuré avec l'option contrôleur (ou procédure e-mail/scan prête).
- Fiche d'intervention rédigée pour le technicien local (branchement, ordre, contacts d'astreinte).

**Jour J — Exécution (sur site)**
1. Le technicien branche les AP (PoE vérifié).
2. Chaque AP : DHCP → découverte NCE → enregistrement (ESN) → affectation au site.
3. NCE pousse le template (SSID, profils radio) puis le firmware si écart.
4. Les AP redémarrent si nécessaire et passent « normal ».
5. Le technicien teste avec son téléphone (SSID visibles, association OK) et remonte la fiche.

**J+1 — Vérification (à distance)**
- Conformité des 20 AP (aucun écart), télémétrie OK, plan de couverture ajusté si besoin.
- Clôture du ticket de déploiement, archivage de la fiche.

Temps typique : quelques heures pour 20 AP avec un technicien non-spécialiste — contre 2-3 jours en configuration manuelle.

## 83. Templates et variables — guide pratique

Méthode pour industrialiser les templates :

