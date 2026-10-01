---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-6
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [779, 929]
sha256: 7855ce31fe1a3e042cc935e0e96d293a02b36c7227193d7ab0a2f82227345f12
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

1. **`arp broadcast enable` oublié** sur les sous-interfaces → DHCP et ping KO. Premier réflexe de dépannage de ce TP.
2. **`dhcp enable` oublié** au niveau global → le service ne tourne pas, `dhcp select global` ne suffit pas.
3. **Pool dont le network ne correspond à aucune interface** → aucune adresse attribuée, sans message d'erreur explicite.
4. **Relais : route de retour manquante** sur R1 vers 192.168.30.0/24 → l'Offer n'arrive jamais au client. Toujours vérifier le routage dans les deux sens.
5. **Bail trop long en lab** : pour voir les renouvellements, utiliser un bail court (ex. 10 minutes) sur un pool de test.
6. **Deux serveurs DHCP sur le même broadcast domain** (en production : box opérateur + serveur) → adresses incohérentes. En lab : ne jamais activer DHCP sur deux équipements du même VLAN.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Sous-interfaces 802.1Q + `arp broadcast enable` | 4 |
| Serveur DHCP : 2 pools complets (network, gateway, DNS, excluded, lease) | 6 |
| PC en DHCP : adresses correctes par VLAN | 3 |
| Relais DHCP fonctionnel vers le site C | 4 |
| Vérifications (`display ip pool`, statistiques) + capture DORA commentée | 3 |

### Durée estimée
**1 h 30** (dont 20 min pour le relais).

### Fiche animateur — points à insister
- Le DHCP, c'est 4 paramètres critiques : **adresse, masque, passerelle, DNS**. Un poste qui "a une IP mais pas Internet" = souvent passerelle ou DNS manquants → `ipconfig /all` est le premier réflexe.
- `excluded-ip-address` : en production, toujours réserver le bas de plage pour les équipements fixes (imprimantes, AP, onduleurs avec carte réseau...). Faire le lien avec le métier de l'équipe.
- Le relais DHCP est indispensable dès qu'on a plusieurs sites/VLAN : un seul serveur centralisé, des relais sur chaque routeur d'accès.
- Montrer `display dhcp server statistics` : les compteurs qui n'augmentent pas indiquent où le dialogue se bloque (Discover sans Offer = problème serveur/pool ; Offer sans Request = problème réseau retour).

---

## 7. TP5 — Routage statique et floating static

### Objectif
Maîtriser les routes statiques : route vers un réseau distant, route par défaut, et route flottante (floating static) comme mécanisme de backup.

### Prérequis
Adressage IP, sous-interfaces (TP4).

### Topologie

```
LAN A (192.168.10.0/24)                    LAN B (192.168.20.0/24)
PC1 ──SW1── GE0/0/1 R1 (AR2220) GE0/0/2 ══════════ GE0/0/2 R2 (AR2220) GE0/0/1──SW2── PC2
              .254      10.0.12.1/30   lien principal    10.0.12.2/30      .254
                              GE0/0/3 ══════════ GE0/0/3
                              10.0.13.1/30   lien backup    10.0.13.2/30
```

- **R1** : GE0/0/1 = 192.168.10.254/24 (LAN A), GE0/0/2 = 10.0.12.1/30 (lien principal), GE0/0/3 = 10.0.13.1/30 (lien backup).
- **R2** : GE0/0/1 = 192.168.20.254/24 (LAN B), GE0/0/2 = 10.0.12.2/30, GE0/0/3 = 10.0.13.2/30.
- SW1 : E0/0/1 (PC1) en access VLAN 10 ; SW2 : idem VLAN 20. Trunks vers les routeurs, sous-interfaces comme au TP4 (ou interfaces physiques en access si on simplifie : ici on utilise des interfaces physiques routées directement pour le LAN, plus simple : GE0/0/1 en `undo portswitch`).

### Énoncé

1. Adresser toutes les interfaces selon le plan.
2. Sur R1 : route statique vers 192.168.20.0/24 via 10.0.12.2 (lien principal) ; route flottante via 10.0.13.2 avec **préférence 100** (la statique par défaut = préférence 60).
3. Sur R2 : miroir (vers 192.168.10.0/24 via 10.0.12.1, backup via 10.0.13.1 préférence 100).
4. Tester PC1 ↔ PC2. Vérifier dans la table de routage quelle route est active.
5. **Test de bascule** : couper le lien principal (shutdown GE0/0/2 sur R1) → vérifier que la route flottante prend le relais, ping continu pendant la coupure.
6. Ajouter une route par défaut sur R1 et R2 vers un "ISP" simulé (R3 avec 10.0.99.0/30, ou un Cloud) — comprendre `0.0.0.0/0`.

### Correction pas à pas

**Adressage R1 :**

```
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]undo portswitch
[R1-GigabitEthernet0/0/1]ip address 192.168.10.254 24
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]undo portswitch
[R1-GigabitEthernet0/0/2]ip address 10.0.12.1 30
[R1]interface GigabitEthernet 0/0/3
[R1-GigabitEthernet0/0/3]undo portswitch
[R1-GigabitEthernet0/0/3]ip address 10.0.13.1 30
```

> Sur les AR eNSP, les ports GE sont en mode switch par défaut : `undo portswitch` les passe en mode routé (L3). Alternative : utiliser des sous-interfaces comme au TP4.

**Routes statiques R1 :**

```
[R1]ip route-static 192.168.20.0 24 10.0.12.2
# Préférence par défaut 60 : route principale.

[R1]ip route-static 192.168.20.0 24 10.0.13.2 preference 100
# Floating static : inactive tant que la principale est valide.

[R1]ip route-static 0.0.0.0 0.0.0.0 10.0.12.2
# Route par défaut (exemple : vers l'ISP via R2, à adapter).
```

**R2 (miroir) :**

```
[R2]ip route-static 192.168.10.0 24 10.0.12.1
[R2]ip route-static 192.168.10.0 24 10.0.13.1 preference 100
```

**Test de bascule :**

```
# Ping continu depuis PC1 vers 192.168.20.11 (PC2)
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]shutdown
# Observer 1 à 3 paquets perdus puis reprise via 10.0.13.0/30.

[R1]display ip routing-table 192.168.20.0
# La route active pointe maintenant vers 10.0.13.2 avec Proto=Static, Pre=100.

[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]undo shutdown
# Retour automatique sur la route préférence 60 (préemption immédiate en statique).
```

### Vérifications

```
[R1]display ip routing-table
# Colonnes : Destination/Mask, Proto (Direct/Static), Pre (préférence), Cost, NextHop, Interface.

[R1]display ip routing-table 192.168.20.0 verbose
# Détail : préférence, tag, état actif/inactif.

[R1]tracert 192.168.20.11
# Visualiser le chemin emprunté (doit passer par 10.0.12.2 en nominal, 10.0.13.2 en backup).

[R1]display ip route-static
# Toutes les statiques configurées, y compris les inactives.
```

### Pièges classiques

1. **Route configurée d'un seul côté** : le ping aller passe, le retour n'a pas de route → échec. Le routage statique est **bidirectionnel par construction** : toujours configurer les deux sens (ou une défaut).
2. **Next-hop injoignable directement** : une statique dont le next-hop n'est pas sur un réseau connecté est invalide (reste inactive). Vérifier avec `display ip routing-table` : la route n'apparaît pas.
3. **Préférence inversée** : mettre `preference 100` sur la principale et rien sur le backup → le backup devient la route active en permanence. Règle : **plus la préférence est petite, plus la route est préférée** (0 = connecté, 60 = statique, 10 = OSPF interne...).
4. **Oublier `undo portswitch`** : l'interface reste en L2, `ip address` est refusé.
5. **Masque /30 mal calculé** : 10.0.12.1/30 et 10.0.12.2/30 sont bien dans le même sous-réseau (10.0.12.0/30 : .0 réseau, .1-.2 hôtes, .3 broadcast). Erreur fréquente : utiliser .3 (broadcast) comme adresse.
6. **Floating static et détection de panne** : la statique ne tombe que si l'interface du next-hop tombe (lien direct). Si la panne est **au-delà** du next-hop, la statique reste active et le trafic part dans le vide → en production, coupler avec de la détection (BFD, NQA, IP SLA). À mentionner même si hors lab.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Adressage complet et cohérent (interfaces + PC) | 5 |
| Routes statiques principales dans les deux sens | 5 |
| Floating static avec préférence correcte | 4 |
| Test de bascule documenté (ping continu + table de routage) | 4 |
| Route par défaut comprise et configurée | 2 |

### Durée estimée
**1 heure**.

