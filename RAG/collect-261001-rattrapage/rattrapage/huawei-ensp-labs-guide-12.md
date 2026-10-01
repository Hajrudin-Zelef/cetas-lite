---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-12
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1738, 1898]
sha256: 33b81b88d27c01fb0bc117e429349f42dd91da80692c1a52ac1d180efb60cf42
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

```
                    ┌──────────── zone trust ────────────┐
                    │                                    │
PC1 (.11) ──SW1── GE0/0/1 FW1 (USG5500) GE0/0/2 ── R-ISP ── "Internet" (Server 203.0.113.10)
PC2 (.12)   VLAN10   192.168.10.254/24   198.51.100.1/30 │ 198.51.100.2/30
Serveur (.100, web)  └────── zone untrust ──────────────┘
                     zone DMZ (optionnelle) : GE0/0/3, 192.168.50.254/24, serveur DMZ
```

- **FW1 (USG5500)** : GE0/0/1 en zone **trust** (192.168.10.254/24), GE0/0/2 en zone **untrust** (198.51.100.1/30).
- **R-ISP** : 198.51.100.2/30 + LAN 203.0.113.0/24 avec un Server public.
- Politique par défaut d'un USG : **tout est refusé entre zones** sauf ce qui est explicitement autorisé.

> **Note de cadrage** : eNSP propose l'USG5500 (syntaxe `firewall zone` + `policy interzone`), vos USG6000 utilisent la même logique de zones avec la syntaxe `security-policy` plus récente. Les concepts (zones, interzone, NAT dans la politique) sont identiques : ce TP entraîne au raisonnement "zones", transférable tel quel.

### Énoncé

1. Créer les zones, y affecter les interfaces, adresser.
2. Vérifier que **par défaut rien ne passe** (PC1 ping 203.0.113.10 → échec) : comprendre le deny implicite.
3. Créer une politique **trust → untrust** autorisant tout (ou au moins ICMP/HTTP/DNS).
4. Configurer le **NAT sortant** (Easy IP) sur l'interface untrust.
5. Tester : PC1 ping et web vers 203.0.113.10 → OK.
6. Créer une politique **untrust → trust** autorisant **uniquement** HTTP vers le serveur interne 192.168.10.100 (port mapping via NAT server ou destination NAT).
7. Tester depuis le Server public : `http://198.51.100.1` → page du serveur interne ; mais `ping 198.51.100.1` → échec (ICMP non autorisé en inbound).
8. (Optionnel) Ajouter une zone **DMZ** avec un serveur, politique untrust→dmz (HTTP only) et trust→dmz (tout).

### Correction pas à pas

**Zones et interfaces :**

```
[FW1]firewall zone trust
[FW1-zone-trust]add interface GigabitEthernet 0/0/1
[FW1-zone-trust]quit
[FW1]firewall zone untrust
[FW1-zone-untrust]add interface GigabitEthernet 0/0/2
[FW1-zone-untrust]quit

[FW1]interface GigabitEthernet 0/0/1
[FW1-GigabitEthernet0/0/1]ip address 192.168.10.254 24
[FW1]interface GigabitEthernet 0/0/2
[FW1-GigabitEthernet0/0/2]ip address 198.51.100.1 30
[FW1]ip route-static 0.0.0.0 0.0.0.0 198.51.100.2
```

> Sur l'USG5500 d'eNSP, les interfaces sont déjà en mode L3 (pas de `undo portswitch` nécessaire en général — vérifier avec `display ip interface brief`).

**Test du deny par défaut :**

```
PC1> ping 203.0.113.10    # ÉCHEC : aucune politique interzone n'existe encore.
```

**Politique trust → untrust :**

```
[FW1]policy interzone trust untrust outbound
[FW1-policy-interzone-trust-untrust-outbound]policy 1
[FW1-policy-interzone-trust-untrust-outbound-1]action permit
[FW1-policy-interzone-trust-untrust-outbound-1]policy source 192.168.10.0 mask 24
[FW1-policy-interzone-trust-untrust-outbound-1]quit
[FW1-policy-interzone-trust-untrust-outbound]quit
```

> `outbound` vu depuis la zone source (trust) : c'est le sens trust→untrust. Sans `policy source`, la règle s'applique à tout le trafic de la zone.

**NAT sortant (Easy IP) :**

```
[FW1]acl 2000
[FW1-acl-basic-2000]rule permit source 192.168.10.0 0.0.0.255
[FW1-acl-basic-2000]quit
[FW1]interface GigabitEthernet 0/0/2
[FW1-GigabitEthernet0/0/2]nat outbound 2000
[FW1-GigabitEthernet0/0/2]quit
```

**Test :** `PC1> ping 203.0.113.10` → OK maintenant (politique + NAT).

**Exposition du serveur web interne (untrust → trust, HTTP only) :**

```
# 1. Port mapping (NAT server) comme au TP7 :
[FW1]interface GigabitEthernet 0/0/2
[FW1-GigabitEthernet0/0/2]nat server protocol tcp global 198.51.100.1 www inside 192.168.10.100 www
[FW1-GigabitEthernet0/0/2]quit

# 2. Politique untrust -> trust : UNIQUEMENT http vers le serveur
[FW1]policy interzone untrust trust inbound
[FW1-policy-interzone-untrust-trust-inbound]policy 1
[FW1-policy-interzone-untrust-trust-inbound-1]action permit
[FW1-policy-interzone-untrust-trust-inbound-1]policy destination 192.168.10.100 mask 32
[FW1-policy-interzone-untrust-trust-inbound-1]service http
[FW1-policy-interzone-untrust-trust-inbound-1]quit
[FW1-policy-interzone-untrust-trust-inbound]quit
```

**Tests depuis l'extérieur (Server public 203.0.113.10) :**

```
# http://198.51.100.1  →  OK (page du serveur 192.168.10.100)
# ping 198.51.100.1    →  KO (ICMP non autorisé en untrust->trust : deny implicite)
```

**Variante DMZ (optionnelle) :**

```
[FW1]firewall zone dmz
[FW1-zone-dmz]add interface GigabitEthernet 0/0/3
[FW1]interface GigabitEthernet 0/0/3
[FW1-GigabitEthernet0/0/3]ip address 192.168.50.254 24

# untrust -> dmz : HTTP only vers le serveur DMZ
[FW1]policy interzone untrust dmz inbound
[FW1-policy-interzone-untrust-dmz-inbound]policy 1
[FW1-policy-interzone-untrust-dmz-inbound-1]action permit
[FW1-policy-interzone-untrust-dmz-inbound-1]policy destination 192.168.50.100 mask 32
[FW1-policy-interzone-untrust-dmz-inbound-1]service http
```

### Vérifications

```
[FW1]display zone
# Zones, interfaces membres, priorités (trust=85, dmz=50, untrust=5, local=100).

[FW1]display policy interzone trust untrust outbound
# Règles configurées dans ce sens.

[FW1]display firewall session table
# Sessions suivies par le firewall (stateful) : constater que le retour est autorisé
# automatiquement (pas besoin de règle retour : le firewall est stateful).

[FW1]display nat outbound
[FW1]display nat server
```

> **Point clé à faire verbaliser** : le firewall est **stateful** : une seule règle trust→untrust suffit pour l'aller-retour d'une session initiée depuis le trust. C'est la différence avec une simple ACL de routeur (stateless).

### Pièges classiques

1. **Oublier que le deny est implicite** : "j'ai autorisé trust→untrust, pourquoi le retour ne passe pas ?" → En fait le retour passe (stateful). La vraie question inverse : "j'ai ouvert untrust→trust pour le web, pourquoi le ping entrant ne passe pas ?" → Parce que seul `http` est autorisé. Bien distinguer.
2. **`inbound`/`outbound` inversés** : `policy interzone trust untrust outbound` = dans le sens trust→untrust. Se tromper de sens = règle sans effet.
3. **NAT server sans politique** : le mapping existe mais le firewall bloque → toujours les **deux** (NAT + politique). Ordre de traitement : NAT destination d'abord, puis politique interzone sur l'adresse traduite.
4. **Oublier la route de retour** côté ISP vers 198.51.100.1 (ici directe, mais en production : route vers l'IP publique du firewall).
5. **Zone `local`** : le trafic **à destination du firewall lui-même** (ping de son IP, web d'admin) passe par les politiques vers/depuis la zone `local`. En lab, si le ping vers 198.51.100.1 échoue depuis l'extérieur alors qu'on l'attendait : c'est la politique untrust→local qui manque (et c'est voulu dans l'énoncé).
6. **Tester le port mapping depuis le LAN** : même piège qu'au TP7 (hairpin) → tester depuis l'extérieur.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Zones créées, interfaces affectées et adressées | 4 |
| Compréhension du deny par défaut (test initial documenté) | 2 |
| Politique trust→untrust + NAT sortant fonctionnels | 5 |
| Port mapping + politique untrust→trust restrictive (HTTP only) | 6 |
| Vérifications (`display zone`, sessions stateful) + tests inbound/outbound | 3 |

### Durée estimée
**1 h 30** (dont 20 min pour la variante DMZ).

