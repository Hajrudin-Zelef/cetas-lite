---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-3
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [433, 662]
sha256: ff02c96dd24c55652af064316b301e5a64192ab963fe578fc01c83884d2f877e
---

# VRP — Le système d'exploitation transversal Huawei

```vrp
<AR720>display current-configuration | include ospf
<AR720>display current-configuration | begin interface GigabitEthernet0/0/0
<AR720>display logbuffer | exclude INFO
<AR720>display ip routing-table | include 192.168.10.0
```

Options disponibles après `|` : `include`, `exclude`, `begin` (et parfois
`count`). Combinez : `display current-configuration | include vlan`.

## 26. Gérer la pagination : `screen-length`

```vrp
<AR720>screen-length 0 temporary     # désactive la pagination pour la session
<AR720>screen-length 24 temporary    # 24 lignes par page (défaut)
```

Le mot-clé `temporary` limite l'effet à la session courante. En system view,
`screen-length 0` (sans temporary) le rend persistant — utile pour les
scripts de sauvegarde (voir section 93).

## 27. L'historique des commandes

```vrp
<AR720>display history-command
  display version
  display current-configuration
  system-view
  ...
```

Affiche les 10 dernières commandes. Avec `↑`/`↓`, on les rappelle et on les
édite — précieux pour rejouer une commande en changeant un paramètre.

## 28. Les alias de commandes (command alias)

VRP permet de définir des alias pour les admins habitués à Cisco :

```vrp
[AR720]command-alias enable
[AR720]command-alias mapping show display
[AR720]command-alias mapping sh display
```

Après cela, `show version` fonctionne comme `display version`. À utiliser
avec parcimonie en équipe : un alias personnel qui n'existe pas sur les
autres boîtiers crée de la confusion. Documentez-les si vous les utilisez.

## 29. Le mode `diagnose` — zone dangereuse

```vrp
[AR720]diagnose
[AR720-diagnose]display ...
```

La vue `diagnose` donne accès à des commandes de diagnostic bas niveau
(mémoire, processus). **Ne l'utilisez que sur consigne du support Huawei** :
certaines commandes peuvent redémarrer un processus ou l'équipement.

## 30. Résumé navigation — pense-bête

```text
<Nom>  --system-view-->  [Nom]  --interface X-->  [Nom-X]
  ^                         |                         |
  |--- return / Ctrl+Z -----+---- return / Ctrl+Z -----+
  |                         |                         |
  +-- quit (quitte session)-+------ quit -------------+
```

---


## 31. `display version` — la carte d'identité

```vrp
<AR720>display version
Huawei Versatile Routing Platform Software
VRP (R) software, Version 5.170 (AR720 V300R022C00SPC500)
Copyright (C) 2011-2024 Huawei Technologies Co., Ltd.
HUAWEI AR720 uptime is 45 days, 3 hours, 12 minutes
...
```

À relever systématiquement : **version VRP exacte**, **modèle**, **uptime**.
C'est la première commande de tout dépannage et de tout upgrade.

## 32. `display device` et étiquettes électroniques

```vrp
<AR720>display device
<AR720>display device elabel        # numéro de série, version matérielle
<AR720>display device fan           # ventilateurs
<AR720>display device power         # alimentations
<AR720>display device temperature all  # sondes de température
```

`display device elabel` donne le **numéro de série** — indispensable pour
ouvrir un ticket TAC ou vérifier une garantie.

## 33. Interfaces : `display interface` et `display interface brief`

```vrp
<AR720>display interface brief
PHY: Physical
*down: administratively down
(l): loopback
(s): spoofing
(b): BFD down
(e): ETHOAM down
InUti/OutUti: input utility/output utility
Interface                   PHY      Protocol  InUti OutUti   inErrors  outErrors
GigabitEthernet0/0/0        up       up           0%     0%          0          0
GigabitEthernet0/0/1        down     down         0%     0%          0          0
NULL0                       up       up           0%     0%          0          0
```

Deux états à distinguer : **PHY** (couche physique) et **Protocol**
(couche liaison). `up/up` = OK. `down/down` = câble ou distant. `up/down` =
problème de protocole (négociation, encapsulation). `*down` = administrativement
coupé (`shutdown`).

Détail d'une interface :

```vrp
<AR720>display interface GigabitEthernet 0/0/0
GigabitEthernet0/0/0 current state : UP
Line protocol current state : UP
Description:LIEN_VERS_S310
...
    Input:  123456 packets, 78901234 bytes
    Output: 654321 packets, 98765432 bytes
    Input bandwidth utilization  : 0.01%
    ...
```

Surveillez : `input errors`, `CRC`, `collisions`, `ignored` — un compteur
qui grimpe = câble, SFP ou duplex à vérifier.

## 34. Adressage IP : `display ip interface brief`

```vrp
<AR720>display ip interface brief
*down: administratively down
(s): spoofing  (l): loopback
Interface                         IP Address/Mask      Physical   Protocol
GigabitEthernet0/0/0              192.168.1.1/24       up         up
Vlanif10                          192.168.10.1/24      up         up
LoopBack0                         1.1.1.1/32           up         up(s)
NULL0                             unassigned           up         up(s)
```

L'équivalent direct du `show ip interface brief` Cisco. Vérifiez aussi :

```vrp
<AR720>display ip interface GigabitEthernet 0/0/0   # détail d'une interface
```

## 35. Table de routage : `display ip routing-table`

```vrp
<AR720>display ip routing-table
Route Flags: R - relay, D - download to fib
------------------------------------------------------------------------------
Routing Tables: Public
         Destinations : 12       Routes : 12
Destination/Mask    Proto   Pre  Cost      Flags NextHop         Interface
      0.0.0.0/0     Static  60   0           D   192.168.1.254   GigabitEthernet0/0/0
   192.168.1.0/24   Direct  0    0           D   192.168.1.1     GigabitEthernet0/0/0
   192.168.1.1/32   Direct  0    0           D   127.0.0.1       InLoopBack0
```

Colonnes : `Proto` (Static, Direct, OSPF, BGP, ISIS, RIP), `Pre`
(préférence administrative), `Cost` (métrique), `NextHop`, `Interface`.
Filtres utiles :

```vrp
<AR720>display ip routing-table 192.168.20.0 24        # route spécifique
<AR720>display ip routing-table protocol ospf         # que l'OSPF
<AR720>display ip routing-table statistics            # compteurs par protocole
<AR720>display fib                                    # FIB (forwarding)
```

## 36. Voisins OSPF : `display ospf peer` et `display ospf interface`

```vrp
<AR720>display ospf peer
          OSPF Process 1 with Router ID 1.1.1.1
                  Neighbors
 Area 0.0.0.0 interface 192.168.1.1(GigabitEthernet0/0/0)'s neighbors
 Router ID: 2.2.2.2     Address: 192.168.1.2
   State: Full         Mode:Nbr is  Slave  Priority: 1
   DR: 192.168.1.2  BDR: 192.168.1.1  MTU: 0
   ...
<AR720>display ospf peer brief
<AR720>display ospf interface GigabitEthernet 0/0/0
<AR720>display ospf lsdb
<AR720>display ospf routing
<AR720>display ospf error                       # erreurs OSPF (utile en dépannage)
```

État attendu entre voisins : **Full** (ou **2-Way** en DROther). Tout autre
état durable (Init, ExStart, Exchange) = problème (MTU, area-id, hello
timers, authentification).

## 37. BGP : `display bgp peer` et table BGP

```vrp
<AR720>display bgp peer
 BGP local router ID : 1.1.1.1
 Local AS number : 65001
 Total number of peers : 1                 Peers in established state : 1
  Peer            V          AS  MsgRcvd  MsgSent  OutQ  Up/Down       State PrefRcv
  10.0.0.2        4       65002     1234     1230     0  02:15:30 Established      12
<AR720>display bgp routing-table
<AR720>display bgp routing-table 192.168.20.0
```

## 38. ARP : `display arp`

```vrp
<AR720>display arp
IP ADDRESS      MAC ADDRESS     EXPIRE(M) TYPE        INTERFACE   VPN-INSTANCE
                                          VLAN/CEVLAN
------------------------------------------------------------------------------
192.168.1.2     00e0-4c12-3456            I -         GE0/0/0
192.168.1.254   00e0-4c78-9abc  18        D-0         GE0/0/0
------------------------------------------------------------------------------
Total:2         Dynamic:1       Static:0    Interface:1
```

