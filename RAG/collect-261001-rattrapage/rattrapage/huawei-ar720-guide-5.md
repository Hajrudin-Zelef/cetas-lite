---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-5
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [748, 966]
sha256: 483896f23061c8a9f542c5d83ccfa94881e114abaf7787110ac1d91713a13e7a
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

1. **NAPT (outbound)** : tout le LAN sort sur Internet avec l'IP du WAN. Le cas de base,
   obligatoire.
2. **NAT server (port mapping)** : exposer un serveur interne (caméra, VPN, applicatif)
   depuis Internet.
3. **NAT statique 1:1** : mapper toute une IP publique vers une IP privée (DMZ).

Sur VRP, le NAT se configure sur l'interface de sortie (outbound) avec une ACL qui
désigne le trafic à translater.

## 36. NAPT sortant : la configuration de base

```
[AGENCE-DAKAR-AR720]acl number 2000
[AGENCE-DAKAR-AR720-acl-basic-2000]rule 5 permit source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2000]rule 10 permit source 192.168.20.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2000]rule 15 permit source 192.168.30.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2000]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat outbound 2000
[AGENCE-DAKAR-AR720-Dialer1]quit
```

- L'ACL 2000 (basique) liste les réseaux autorisés à sortir. **Tout réseau oublié ici =
  pas d'Internet** (cause fréquente du « le VLAN invités n'a pas Internet »).
- `nat outbound 2000` sans adresse = NAPT avec l'IP de l'interface (le mode par défaut
  et le plus courant).
- Répéter `nat outbound 2000` sur **chaque** interface WAN (Dialer 1, GE0/0/1…).

Vérification :

```
display nat outbound
display nat session table
```

## 37. NAPT avec pool d'adresses (plusieurs IP publiques)

Si le FAI fournit un bloc d'IP publiques (ex. /29) :

```
[AGENCE-DAKAR-AR720]nat address-group 1 197.155.10.35 197.155.10.38
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat outbound 2000 address-group 1
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Le routeur répartit les sessions sur le pool. Réserver une IP du bloc pour le NAT server
(ne pas l'inclure dans l'address-group si elle sert à exposer un serveur).

## 38. NAT server : exposer un serveur interne (port mapping)

Exemple : un serveur de vidéosurveillance (192.168.30.10) accessible depuis Internet sur
le port 8080, et un accès VPN L2TP (ports UDP 500/4500/1701 — voir Partie F).

```
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat server protocol tcp global current-interface 8080 inside 192.168.30.10 8080
[AGENCE-DAKAR-AR720-Dialer1]quit
```

- `global current-interface` : utilise l'IP courante de l'interface (pratique en PPPoE/
  DHCP où l'IP change). Avec une IP fixe, on peut écrire l'IP en dur.
- `protocol tcp ... 8080 inside ... 8080` : port externe 8080 → port interne 8080.
  On peut faire du port mapping (externe 8080 → interne 80).

**Règle de sécurité** : chaque `nat server` doit s'accompagner d'une politique de
sécurité WAN→LAN qui n'autorise QUE ce port (Partie G). Un NAT server sans filtre =
porte ouverte.

## 39. NAT statique 1:1 (DMZ)

```
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat static global 197.155.10.36 inside 192.168.30.10
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Toute l'IP publique 197.155.10.36 est mappée vers le serveur interne. À réserver aux
vrais besoins (serveur avec beaucoup de ports). Vérification : `display nat static`.

## 40. NAT et VPN : l'exclusion du trafic tunnelé

Quand un tunnel IPSec transporte du trafic LAN→LAN entre sites, ce trafic **ne doit pas
être NATé** (sinon le tunnel ne le reconnaît plus). On l'exclut avec une ACL :

```
[AGENCE-DAKAR-AR720]acl number 3000
[AGENCE-DAKAR-AR720-acl-adv-3000]rule 5 deny ip source 192.168.10.0 0.0.0.255 destination 192.168.0.0 0.0.255.255
[AGENCE-DAKAR-AR720-acl-adv-3000]rule 10 permit ip source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3000]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat outbound 3000
[AGENCE-DAKAR-AR720-Dialer1]quit
```

La règle `deny` en premier : le trafic vers le réseau du siège (192.168.0.0/16) n'est
pas NATé, tout le reste oui. **Oublier cette exclusion = tunnel IPSec qui monte mais
aucun trafic ne passe** (classique absolu, voir cas n°1 du dépannage).

## 41. ALG : quand le NAT casse les protocoles (FTP, SIP…)

Certains protocoles embarquent des adresses IP dans la charge utile (FTP actif, SIP).
Le NAT de base ne les réécrit pas → ça casse. VRP propose des ALG (Application Level
Gateway) :

```
[AGENCE-DAKAR-AR720]nat alg ftp enable
[AGENCE-DAKAR-AR720]nat alg sip enable
```

N'activer que ce qui sert. Pour la ToIP, préférer un SBC ou un VPN vers le PABX plutôt
que du SIP en NAT direct si possible — le SIP à travers NAT reste une source de tickets.

## 42. Vérifier et purger les sessions NAT

```
display nat session table verbose
display nat statistics
```

Si une translation est « coincée » (vieille session qui bloque un port) :

```
reset nat session table
```

(En vue utilisateur. Attention : ça coupe les sessions actives — à faire en fenêtre de
maintenance ou quand le problème est avéré.)

---
---

# Partie E — Routage

## 43. Routage statique : les bases VRP

```
[AGENCE-DAKAR-AR720]ip route-static 10.0.0.0 255.255.0.0 192.168.10.254
[AGENCE-DAKAR-AR720]ip route-static 10.0.0.0 255.255.0.0 192.168.10.253 preference 100
```

- Le premier paramètre après le réseau = **next-hop** (IP du routeur suivant).
- `preference` : distance administrative (60 par défaut en statique). Plus c'est bas,
  plus c'est prioritaire.
- On peut aussi spécifier l'interface de sortie : `ip route-static 0.0.0.0 0.0.0.0
  Dialer 1` (obligatoire pour les interfaces point-à-point/PPPoE où il n'y a pas de
  next-hop IP fixe).

Vérification : `display ip routing-table`, `display ip routing-table statistics`.

## 44. Route flottante : le secours manuel simple

Deux routes vers la même destination, préférences différentes :

```
[AGENCE-DAKAR-AR720]ip route-static 10.0.0.0 255.255.0.0 172.16.0.1
[AGENCE-DAKAR-AR720]ip route-static 10.0.0.0 255.255.0.0 172.16.1.1 preference 120
```

Tant que le next-hop principal répond (et que l'interface est up), la route de secours
reste inactive. Simple, mais sans NQA ça ne détecte que la perte du lien local —
combiner avec le track/NQA de la Partie B pour un vrai basculement.

## 45. OSPF : pourquoi et quand l'utiliser

Dès qu'on a 3+ sites interconnectés (ou un LAN avec plusieurs routeurs), le statique
devient ingérable. OSPF (dynamique, état de liens) :

- converge seul en cas de panne,
- s'authentifie (MD5),
- supporte le multi-zone (backbone area 0).

Sur l'AR720, OSPF est supporté en IPv4 et OSPFv3 en IPv6 (selon datasheet).

## 46. OSPF : configuration minimale entre deux sites

Côté agence (Router-ID = IP de loopback, à créer) :

```
[AGENCE-DAKAR-AR720]interface LoopBack 0
[AGENCE-DAKAR-AR720-LoopBack0]ip address 10.255.0.11 255.255.255.255
[AGENCE-DAKAR-AR720-LoopBack0]quit
[AGENCE-DAKAR-AR720]ospf 1 router-id 10.255.0.11
[AGENCE-DAKAR-AR720-ospf-1]area 0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 10.255.0.11 0.0.0.0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]quit
[AGENCE-DAKAR-AR720-ospf-1]quit
```

Points clés :

- `router-id` explicite et **unique** par routeur (sinon conflits d'élection).
- Les `network` utilisent des **wildcard masks** (inverses du masque).
- Ne déclarer en OSPF que les interfaces qui doivent parler OSPF. L'interface WAN vers
  Internet : `silent-interface` (voir section 48).

Vérification :

```
display ospf peer
display ospf routing
display ospf interface
```

## 47. OSPF sur tunnel GRE/IPSec : attention au multicast

OSPF utilise du multicast (224.0.0.5/6). Sur un tunnel GRE ça passe (GRE transporte le
multicast). Sur de l'IPSec pur en mode tunnel, il faut vérifier que la crypto ACL couvre
le multicast ou passer par du GRE over IPSec (recommandé, voir section 58).

Pour forcer l'adjacence sur un lien point-à-point (tunnel) :

```
[AGENCE-DAKAR-AR720]interface Tunnel 0/0/1
[AGENCE-DAKAR-AR720-Tunnel0/0/1]ospf network-type p2p
[AGENCE-DAKAR-AR720-Tunnel0/0/1]quit
```

