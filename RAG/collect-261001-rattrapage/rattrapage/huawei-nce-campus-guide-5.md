---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-5
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [331, 411]
sha256: f4fc67ff6e695e5151d665a10507f508214dc88683ee52d08661d3d275da8008
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Cette architecture « manager + controller + analyzer » intégrés est la différence fondamentale avec eSight (manager seul).

## 12. Les trois couches logiques : management, control, analysis

**Management (gestion)** : c'est la partie « NMS » — inventaire des équipements, gestion des configurations et des firmwares, topologie, alarmes, rapports, gestion des licences, portails multi-tenant. Si vous venez d'eSight, c'est la partie la plus familière.

**Control (contrôle)** : la partie SDN — le contrôleur traduit des **intentions de service** (ex : « les utilisateurs du VLAN 20 ont 100 Mbps garantis ») en configurations poussées aux équipements via NETCONF/YANG. Il orchestre aussi le ZTP, les templates, les politiques d'accès et de QoS. C'est ce qui n'existe pas dans eSight.

**Analysis (analyse)** : la collecte et l'exploitation des données — télémétrie temps réel, statistiques d'expérience par utilisateur/application, détection d'anomalies, aide au diagnostic. La version de base est intégrée à NCE-Campus ; la version avancée (ML, prédiction) est apportée par **CampusInsight**.

L'intérêt de l'intégration : quand une alarme remonte (management), le contrôleur connaît la configuration déployée (control) et l'analyseur sait si c'est un vrai problème d'expérience (analysis) — la corrélation est native, pas bricolée entre trois outils.

## 13. Le contrôleur SDN au cœur du système

Le contrôleur SDN est le cerveau : il maintient une **vue centralisée** de l'état du réseau (topologie, configurations, politiques) et pilote les équipements en mode « intent-based ».

Concrètement, l'administrateur ne configure plus chaque switch en CLI : il définit dans NCE des **sites**, des **templates**, des **politiques** (VLAN, SSID, 802.1X, QoS). Le contrôleur :

1. Valide la cohérence (conflits, ressources).
2. Génère les configurations effectives par équipement (en substituant les variables des templates).
3. Les pousse via NETCONF (edit-config) ou les protocoles adaptés.
4. Vérifie la conformité en continu (configuration consistency verification : get-config vs attendu, avec notification de changement).

Point d'honnêteté : le SDN de campus Huawei reste un **SDN « piloté par templates »** plus qu'un SDN temps réel type data center — les décisions de forwarding restent locales aux équipements ; le contrôleur orchestre la configuration et les politiques, il ne calcule pas chaque flux. C'est un choix d'architecture sain pour un campus (résilience en cas de perte du contrôleur).

## 14. Southbound interfaces : comment NCE parle aux équipements

« Southbound » = les protocoles entre le contrôleur et les équipements gérés. NCE-Campus en utilise plusieurs, selon le type d'équipement et la fonction :

| Protocole | Usage principal | Équipements concernés |
|---|---|---|
| **NETCONF / YANG** | Configuration structurée (déploiement, vérification de conformité) | Switches, routeurs AR récents (modèles supportant YANG) |
| **SNMP** | Supervision, télémétrie classique, gestion des équipements legacy | Tous, y compris équipements traditionnels |
| **CAPWAP** | Tunnels de gestion/contrôle des AP vers le WAC | AP AirEngine (via WAC physique ou fonction WAC) |
| **HTTP/2** | Télémétrie et authentification (selon versions) | Équipements récents |
| **HTTPS** | Gestion web sécurisée, portails | Divers |
| **TCP** | Canaux dédiés (ex : synchro avec composants d'authentification distants) | Composants d'authentification |

La documentation V300R022C10 mentionne explicitement NETCONF/SNMP/HTTP/2/HTTPS/TCP comme canaux southbound. Le choix du protocole est automatique selon les capacités annoncées par l'équipement lors de l'onboarding.

## 15. NETCONF/YANG — le protocole de configuration structurée

NETCONF (RFC 6241/6242) est le protocole moderne de configuration réseau : contrairement à SNMP (orienté supervision) ou au CLI (texte libre), il manipule des **données structurées** décrites par des modèles **YANG** (RFC 7950).

Pourquoi c'est important dans NCE :

- **Idempotence et validation** : une opération `edit-config` est validée contre le modèle YANG avant application — moins d'erreurs de syntaxe qu'en CLI.
- **Transactions** : possibilité de valider/commit avec rollback (candidate datastore) — base du rollback automatisé.
- **Vérification de conformité** : NCE peut faire `get-config` et comparer avec la configuration attendue (configuration consistency verification), et recevoir des **notifications de changement** quand quelqu'un modifie un équipement en CLI hors NCE (détection du « shadow IT » local — voir cas pratique 14).
- **Interopérabilité** : YANG est un standard — les mêmes concepts servent pour l'automatisation par scripts/Ansible.

Limite honnête : tous les équipements ne supportent pas YANG de façon égale ; les gammes récentes (CloudEngine, NetEngine AR V5, AirEngine) sont les mieux couvertes. Les équipements anciens restent gérés en SNMP/CLI-like — vérifier la **matrice de compatibilité** de la version NCE cible pour chaque modèle (S310, AR720 : vérifier le support NETCONF/YANG exact selon la version logicielle).

## 16. SNMP — la télémétrie et le legacy

SNMP reste utilisé par NCE-Campus pour :

- La **supervision** des équipements traditionnels ou tiers (au sens large : tout ce qui n'est pas piloté en NETCONF).
- Les **tests SLA** : la documentation Monitoring and O&M (V300R020C10) indique que les tâches SLA s'appuient sur le protocole **NQA** des équipements, configurables depuis NCE pour les équipements gérés en SNMP.
- La **découverte** et la collecte d'indicateurs de performance de base (CPU, mémoire, interfaces).

Pour Zelef, SNMP reste le protocole de supervision de ses onduleurs et équipements d'énergie — mais **pas via NCE** : c'est le domaine d'eSight ou d'un superviseur dédié. Ne pas confondre les périmètres.

Bonnes pratiques SNMP sous NCE : utiliser **SNMPv3** (authPriv) partout, bannir les community strings `public`/`private`, filtrer les accès SNMP par ACL de management.

## 17. CAPWAP — la gestion des AP AirEngine

Les AP Wi-Fi Huawei (AirEngine : AP361, AP761, etc.) fonctionnent historiquement en mode **fit AP** pilotés par un contrôleur sans-fil (**WAC** — Wireless Access Controller, anciennement AC) via le protocole **CAPWAP** (RFC 5415) : le WAC centralise la gestion radio, les SSID, l'authentification, le roaming.

Dans l'architecture NCE-Campus :

- NCE-Campus **orchestre** la couche WLAN (profils radio, SSID, politiques) et la pousse vers les WAC/AP.
- Les AP peuvent être gérés en mode central (via WAC) ; NCE s'interface au-dessus pour l'automatisation et la supervision unifiée.
- Pour les petits déploiements, des AP peuvent fonctionner avec des fonctions WAC intégrées/virtuelles selon les modèles — vérifier par modèle.

Point de vigilance : la documentation V300R022C10 note des **restrictions** sur certaines fonctions avancées selon les générations d'équipements (ex : l'authentification sans-fil directe, l'identification de terminaux/application exigent des équipements V5 comme dispositifs d'accès). Toujours vérifier la matrice « fonction × modèle » avant de promettre une fonction à un métier.

## 18. HTTP/2, HTTPS et autres canaux southbound

Compléments aux protocoles principaux :

