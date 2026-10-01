---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-4
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [526, 747]
sha256: 3237354a3285eb2368ab767eccc6886becb05b73f829810c8b0e2bb488c433da
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

Exemple de trame (noms d'interface à adapter au modèle de carte réel) :

```
[AGENCE-DAKAR-AR720]interface Cellular 0/0/0
[AGENCE-DAKAR-AR720-Cellular0/0/0]apn-profile apn-internet
[AGENCE-DAKAR-AR720-Cellular0/0/0]quit
[AGENCE-DAKAR-AR720]apn-profile apn-internet
[AGENCE-DAKAR-AR720-apn-profile-apn-internet]apn internet.operateur.exemple
[AGENCE-DAKAR-AR720-apn-profile-apn-internet]quit
```

Vérifier le signal : `display cellular 0/0/0` (niveau RSRP/SINR — viser un montage
d'antenne correct, éviter de laisser le routeur au fond d'une armoire métallique fermée).

## 25. MTU/MSS sur le WAN : le détail qui casse les VPN

Le PPPoE réduit la MTU utile (1492 au lieu de 1500, à cause de l'en-tête PPPoE). Sans
ajustement, certains sites web ne chargent pas et les tunnels VPN rament.

```
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]tcp adjust-mss 1400
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Règle pratique : MSS = MTU du lien − 100 environ (1400 pour PPPoE, 1440 pour de l'IP
classique derrière NAT). À ajuster si des applications spécifiques coincent.

---
---

# Partie C — LAN et services locaux

## 26. Adressage du LAN : plan et configuration

Plan type d'une agence (à adapter, mais **documenter et ne plus changer**) :

| Réseau | Usage | Interface |
|---|---|---|
| 192.168.10.0/24 | LAN bureautique | GE0/0/2 (ou Vlanif10) |
| 192.168.20.0/24 | Wi-Fi invités | Vlanif20 |
| 192.168.30.0/24 | Téléphonie / imprimantes | Vlanif30 |
| 192.168.99.0/24 | Management | Vlanif99 |

Configuration de base :

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/2
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]ip address 192.168.10.1 255.255.255.0
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]quit
```

Ou en mode VLAN (recommandé dès qu'il y a plusieurs réseaux) :

```
[AGENCE-DAKAR-AR720]vlan 10
[AGENCE-DAKAR-AR720-vlan10]quit
[AGENCE-DAKAR-AR720]interface Vlanif 10
[AGENCE-DAKAR-AR720-Vlanif10]ip address 192.168.10.1 255.255.255.0
[AGENCE-DAKAR-AR720-Vlanif10]quit
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/2
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]port link-type access
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]port default vlan 10
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]quit
```

## 27. Serveur DHCP : pool de base

```
[AGENCE-DAKAR-AR720]dhcp enable
[AGENCE-DAKAR-AR720]ip pool lan-bureautique
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]gateway-list 192.168.10.1
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]network 192.168.10.0 mask 255.255.255.0
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]dns-list 192.168.10.1 8.8.8.8
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]excluded-ip-address 192.168.10.1 192.168.10.20
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]lease day 1 hour 0 minute 0
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]quit
[AGENCE-DAKAR-AR720]interface Vlanif 10
[AGENCE-DAKAR-AR720-Vlanif10]dhcp select global
[AGENCE-DAKAR-AR720-Vlanif10]quit
```

- `dhcp enable` : activer le service DHCP **globalement** (oubli classique).
- `excluded-ip-address` : réserver la plage des équipements fixes (imprimantes, serveurs).
- `lease day 1` : bail d'un jour (réseau bureautique stable). Pour un Wi-Fi invités :
  bail court (quelques heures).

## 28. DHCP : réservations (baux statiques) pour imprimantes et équipements

Les copieurs/imprimantes réseau et les équipements fixes doivent avoir une IP stable.
Deux méthodes : réservation DHCP (recommandé, centralisé) ou IP fixe sur l'équipement.

```
[AGENCE-DAKAR-AR720]ip pool lan-bureautique
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]static-bind ip-address 192.168.10.21 mac-address aaaa-bbbb-cccc
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]quit
```

Format MAC Huawei : `aaaa-bbbb-cccc`. Vérifier les baux attribués :

```
display ip pool name lan-bureautique used
```

Astuce terrain : tenir un tableau (dans le dossier du site) : IP / MAC / équipement /
emplacement. Quand l'imprimante « disparaît » du réseau, c'est le premier document qu'on
ouvre.

## 29. DHCP : options avancées (option 43, 66, 150…)

Les options DHCP servent à pousser des paramètres spécifiques : serveur TFTP pour des
téléphones IP (option 66/150), contrôleur Wi-Fi (option 43), etc.

```
[AGENCE-DAKAR-AR720]ip pool lan-telephonie
[AGENCE-DAKAR-AR720-ip-pool-lan-telephonie]option 66 ip-address 192.168.30.5
[AGENCE-DAKAR-AR720-ip-pool-lan-telephonie]quit
```

Pour des options brutes (hexadécimal), utiliser `option <code> hex <valeur>` — **syntaxe
à vérifier selon la version VRP**, les options constructeur (ex. option 43 Huawei pour
les AP) ont des formats spécifiques documentés dans le guide WLAN.

## 30. DNS : le routeur comme relais + DNS statiques

Faire du routeur le DNS du LAN simplifie la vie (un seul point à changer si les DNS du
FAI changent) :

```
[AGENCE-DAKAR-AR720]dns resolve
[AGENCE-DAKAR-AR720]dns server 8.8.8.8
[AGENCE-DAKAR-AR720]dns server 1.1.1.1
[AGENCE-DAKAR-AR720]interface Vlanif 10
[AGENCE-DAKAR-AR720-Vlanif10]dhcp select global
```

(avec `dns-list 192.168.10.1 ...` dans le pool, cf. section 27 : les clients interrogent
le routeur, qui relaie.)

Entrées statiques utiles (intranet, imprimantes) :

```
[AGENCE-DAKAR-AR720]ip host imprimante-accueil 192.168.10.21
[AGENCE-DAKAR-AR720]ip host intranet-siege 10.0.0.10
```

Vérification : `display dns dynamic-host` / `ping imprimante-accueil`.

## 31. Relais DHCP (quand le serveur est ailleurs)

Si le serveur DHCP est au siège (ou sur un serveur Windows du site), le routeur fait
relais :

```
[AGENCE-DAKAR-AR720]dhcp enable
[AGENCE-DAKAR-AR720]interface Vlanif 20
[AGENCE-DAKAR-AR720-Vlanif20]dhcp select relay
[AGENCE-DAKAR-AR720-Vlanif20]dhcp relay server-ip 10.0.0.5
[AGENCE-DAKAR-AR720-Vlanif20]quit
```

Cas typique : Wi-Fi invités avec portail captif hébergé ailleurs, ou centralisation des
baux DHCP au siège pour audit.

## 32. VLAN : segmentation minimale recommandée

Même sur une petite agence, segmenter en 3 VLAN change tout pour la sécurité :

| VLAN | Nom | Réseau | Règle |
|---|---|---|---|
| 10 | Bureautique | 192.168.10.0/24 | Accès Internet + VPN siège |
| 20 | Invités | 192.168.20.0/24 | Internet seul, isolé du reste |
| 30 | Équipements | 192.168.30.0/24 | Imprimantes, caméras : pas d'Internet direct |

```
[AGENCE-DAKAR-AR720]vlan batch 10 20 30
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/3
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/3]port link-type trunk
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/3]port trunk allow-pass vlan 10 20 30
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/3]quit
```

Le filtrage inter-VLAN se fait ensuite avec les zones de sécurité (Partie G).

## 33. STP et boucles : le minimum vital

Sur les ports d'accès vers les PC, activer **STP edge** (portfast) pour une montée de
lien rapide, et **BPDU protection** pour couper net tout port qui reçoit des BPDU
(quelqu'un qui branche un petit switch sauvage) :

```
[AGENCE-DAKAR-AR720]stp enable
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/2
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]stp edged-port enable
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]quit
[AGENCE-DAKAR-AR720]stp bpdu-protection
```

Une boucle de switching = tempête de broadcast = réseau mort en quelques secondes.
Sur le terrain, 90 % des « le réseau est lent d'un coup » inexpliqués sont des boucles
(câble branché deux fois, petit switch de bureau rebouclé).

## 34. Agrégation et storm-control (protection broadcast)

Limiter les tempêtes sur les ports d'accès :

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/2
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]storm-control broadcast min-rate 1000 max-rate 2000
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]quit
```

(valeurs en kbps, à ajuster selon la taille du site). Ça ne remplace pas la suppression
de la boucle, mais ça évite que tout le site tombe en attendant l'intervention.

---
---

# Partie D — NAT

## 35. NAT : les trois usages sur l'AR720

