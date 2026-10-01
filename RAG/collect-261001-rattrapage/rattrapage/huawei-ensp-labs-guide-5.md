---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-5
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [606, 778]
sha256: baff0191b8f69a0203898ebd4245ec1ef6b5851967707475ab6d4d9dc12ed721
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

| Critère | Points |
|---|---|
| Mise en évidence du danger de la boucle (observation + explication) | 4 |
| RSTP activé sur les 3 switches, root identifié avec `display stp` | 5 |
| SW1 root primary / SW2 secondary configurés et vérifiés | 5 |
| Edge ports configurés sur les ports PC uniquement | 3 |
| Test de convergence documenté (avant/après coupure) | 3 |

### Durée estimée
**1 h 15** (dont 15 min de capture BPDU).

### Fiche animateur — points à insister
- La boucle L2 est l'incident le plus violent d'un réseau de campus : tempête de broadcast = réseau entier à genoux en quelques secondes. D'où l'importance de RSTP **toujours actif** (ne jamais le désactiver "parce que ça ralentit").
- Le root bridge doit être **choisi**, pas subi : sur le terrain, le root = le cœur (vos S310 d'agrégation, pas un switch d'étage).
- Edge port = uniquement vers équipements terminaux (PC, imprimantes, AP). Jamais vers un autre switch.
- Montrer la différence de convergence : en RSTP la bascule est quasi instantanée ; faire le lien avec la haute disponibilité attendue par les utilisateurs.
- Question piège : "Que se passe-t-il si on branche deux câbles entre les mêmes deux switches sans STP ?" → Boucle immédiate. "Et avec RSTP ?" → Un des deux ports passe en discarding (alternate), le lien sert de backup.

---

## 6. TP4 — Serveur DHCP sur routeur AR

### Objectif
Fournir automatiquement adresses IP, masque, passerelle et DNS via DHCP depuis un routeur AR ; comprendre pools, exclusions, durées de bail ; mettre en place un relais DHCP vers un site distant.

### Prérequis
TP1 (VLAN), notions d'adressage IP.

### Topologie

```
Site A (VLAN 10)                          Site B (VLAN 20)
PC1 ──E0/0/1 SW1 (S3700) GE0/0/1 ── GE0/0/1 R1 (AR2220) GE0/0/2 ── GE0/0/1 SW2 (S3700) E0/0/1── PC3
PC2 ──E0/0/2         (trunk)                (VLAN 10: 192.168.10.254)   (trunk)              E0/0/2── PC4
                                               (VLAN 20: 192.168.20.254, via sous-interfaces)
```

- **R1 (AR2220)** : routeur central, serveur DHCP.
- **SW1 (S3700)** : ports E0/0/1-2 en access VLAN 10, trunk GE0/0/1 vers R1.
- **SW2 (S3700)** : ports E0/0/1-2 en access VLAN 20, trunk GE0/0/1 vers R1.
- R1 utilise des **sous-interfaces** (router-on-a-stick) : GE0/0/1.10 (VLAN 10, 192.168.10.254/24) et GE0/0/2.20 (VLAN 20, 192.168.20.254/24).
- Les PC sont en **DHCP** (double-clic > DHCP activé).

### Énoncé

1. Configurer les VLAN et trunks sur SW1/SW2 (rappel TP2).
2. Configurer les sous-interfaces 802.1Q sur R1 avec les IP passerelles.
3. Activer le **serveur DHCP** sur R1, créer deux pools :
   - `POOL-USERS` : réseau 192.168.10.0/24, passerelle 192.168.10.254, DNS 192.168.10.254 (ou 8.8.8.8 en exemple), **exclure** 192.168.10.1 à 192.168.10.10 (plage réservée aux équipements).
   - `POOL-GUESTS` : réseau 192.168.20.0/24, passerelle 192.168.20.254, bail de 1 heure.
4. Basculer les 4 PC en DHCP, vérifier qu'ils reçoivent une adresse du bon pool.
5. **Relais DHCP** : ajouter un second routeur R2 entre R1 et un "site C" (VLAN 30, 192.168.30.0/24) ; le serveur DHCP reste sur R1 ; configurer `dhcp relay` sur R2.
6. Vérifier les baux (`display ip pool`) et tester la connectivité inter-VLAN via R1.

### Correction pas à pas

**SW1 :**

```
[SW1]vlan batch 10 20
[SW1]interface Ethernet 0/0/1
[SW1-Ethernet0/0/1]port link-type access
[SW1-Ethernet0/0/1]port default vlan 10
[SW1]interface Ethernet 0/0/2
[SW1-Ethernet0/0/2]port link-type access
[SW1-Ethernet0/0/2]port default vlan 10
[SW1]interface GigabitEthernet 0/0/1
[SW1-GigabitEthernet0/0/1]port link-type trunk
[SW1-GigabitEthernet0/0/1]port trunk allow-pass vlan 10
```

**SW2 :** même logique avec VLAN 20 (E0/0/1-2 en access VLAN 20, trunk allow-pass vlan 20).

**R1 — sous-interfaces (router-on-a-stick) :**

```
[R1]interface GigabitEthernet 0/0/1.10
[R1-GigabitEthernet0/0/1.10]dot1q termination vid 10
[R1-GigabitEthernet0/0/1.10]ip address 192.168.10.254 24
[R1-GigabitEthernet0/0/1.10]arp broadcast enable
[R1-GigabitEthernet0/0/1.10]quit

[R1]interface GigabitEthernet 0/0/2.20
[R1-GigabitEthernet0/0/2.20]dot1q termination vid 20
[R1-GigabitEthernet0/0/2.20]ip address 192.168.20.254 24
[R1-GigabitEthernet0/0/2.20]arp broadcast enable
[R1-GigabitEthernet0/0/2.20]quit
```

> `arp broadcast enable` est **indispensable** sur les sous-interfaces VRP : sans lui, le routeur ne répond pas aux ARP des clients du VLAN et le ping échoue. C'est LE piège classique de ce TP.

**R1 — serveur DHCP :**

```
[R1]dhcp enable

[R1]ip pool POOL-USERS
[R1-ip-pool-POOL-USERS]network 192.168.10.0 mask 24
[R1-ip-pool-POOL-USERS]gateway-list 192.168.10.254
[R1-ip-pool-POOL-USERS]dns-list 192.168.10.254
[R1-ip-pool-POOL-USERS]excluded-ip-address 192.168.10.1 192.168.10.10
[R1-ip-pool-POOL-USERS]lease day 1 hour 0
[R1-ip-pool-POOL-USERS]quit

[R1]ip pool POOL-GUESTS
[R1-ip-pool-POOL-GUESTS]network 192.168.20.0 mask 24
[R1-ip-pool-POOL-GUESTS]gateway-list 192.168.20.254
[R1-ip-pool-POOL-GUESTS]dns-list 8.8.8.8
[R1-ip-pool-POOL-GUESTS]lease hour 1
[R1-ip-pool-POOL-GUESTS]quit

# Activer DHCP sur les interfaces (mode global) :
[R1]interface GigabitEthernet 0/0/1.10
[R1-GigabitEthernet0/0/1.10]dhcp select global
[R1]interface GigabitEthernet 0/0/2.20
[R1-GigabitEthernet0/0/2.20]dhcp select global
[R1]save
```

**Comment le bon pool est-il choisi ?** Le serveur sélectionne le pool dont le `network` correspond au sous-réseau de l'interface d'arrivée de la requête. D'où l'importance que chaque sous-interface ait son pool.

**PC en DHCP** : double-clic sur chaque PC > cocher **DHCP**. Puis en console du PC : `ipconfig` doit montrer une adresse 192.168.10.x (PC1/PC2) ou 192.168.20.x (PC3/PC4).

**Partie relais — topologie étendue :**

```
R1 GE0/0/3 (192.168.100.1/30) ── (192.168.100.2/30) R2 (AR2220) GE0/0/1.30 (VLAN 30, 192.168.30.254/24, dot1q vid 30)
                                                                                    │
                                                                              SW3 (S3700), PC5/PC6 en DHCP
```

```
# Sur R2 : route vers les pools + relais
[R2]dhcp enable
[R2]interface GigabitEthernet 0/0/1.30
[R2-GigabitEthernet0/0/1.30]dot1q termination vid 30
[R2-GigabitEthernet0/0/1.30]ip address 192.168.30.254 24
[R2-GigabitEthernet0/0/1.30]arp broadcast enable
[R2-GigabitEthernet0/0/1.30]dhcp relay server-ip 192.168.100.1
[R2-GigabitEthernet0/0/1.30]quit

# Routage statique pour joindre le serveur (ou OSPF du TP6 en version avancée)
[R2]ip route-static 0.0.0.0 0.0.0.0 192.168.100.1
[R1]ip route-static 192.168.30.0 24 192.168.100.2

# Sur R1 : pool pour le site C
[R1]ip pool POOL-SITEC
[R1-ip-pool-POOL-SITEC]network 192.168.30.0 mask 24
[R1-ip-pool-POOL-SITEC]gateway-list 192.168.30.254
[R1-ip-pool-POOL-SITEC]quit
```

> Le relais fonctionne parce que la requête DHCP (broadcast) est interceptée par R2, convertie en unicast vers 192.168.100.1 avec l'option `giaddr` = 192.168.30.254 : le serveur choisit alors le pool POOL-SITEC.

### Vérifications

```
[R1]display ip pool
# Pools, plages, baux attribués, adresses utilisées/libres.

[R1]display ip pool name POOL-USERS used
# Détail des baux actifs : IP, MAC, durée restante.

[R1]display dhcp server statistics
# Compteurs Discover/Offer/Request/ACK : idéal pour voir le dialogue.

# Sur PC : ipconfig /all → adresse, masque, passerelle, DNS, bail.
```

**Capture Wireshark** sur le lien SW1↔R1 : filtrer `bootp` → observer **Discover → Offer → Request → ACK** (DORA). Faire remarquer que le client envoie en broadcast (0.0.0.0 → 255.255.255.255) et que le serveur répond en unicast.

### Pièges classiques

