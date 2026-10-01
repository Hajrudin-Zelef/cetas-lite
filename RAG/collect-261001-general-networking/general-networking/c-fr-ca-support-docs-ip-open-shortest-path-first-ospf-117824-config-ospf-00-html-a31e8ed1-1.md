---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-117824-config-ospf-00-html-a31e8ed1-1
title: "c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-117824-config-ospf-00-html-a31e8ed1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-117824-config-ospf-00-html-a31e8ed1.md
source_anchor: ""
source_lines: [1, 147]
sha256: 158ac10f1ba4a5d4b7beebdd511902ce9e4958f32c7eafee4c758844bfb20553
---

# c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-117824-config-ospf-00-html-a31e8ed1

Ce document décrit comment le protocole OSPF (Open Shortest Path First) entre Nexus et la fonctionnalité Cisco IOS® est mis en oeuvre dans Cisco IOS et Nexus Operating System (NXOS).
Cisco vous recommande de connaître le protocole OSPF.
Les informations contenues dans ce document sont basées sur les versions de matériel et de logiciel suivantes :
Les périphériques Cisco IOS prennent en charge RFC 1583. Cependant, NXOS prend en charge la RFC 2328 et il existe des conceptions dans lesquelles cette différence peut créer des boucles de routage dans le réseau lorsqu'il y a des routes OSPF externes dans le réseau.
La différence entre les RFC 1583 et RFC 2328, en ce qui concerne le choix de la meilleure route parmi plusieurs routes externes, est abordée dans cette section.
Afin de comparer les chemins externes de type 1, observez la somme de la distance à l'adresse de transfert et la métrique de type 1 annoncée (X+Y).  Afin de comparer les chemins externes de type 2, examinez les métriques de type 2 annoncées, puis, si nécessaire, la distance vers les adresses de transfert.
Si le nouveau chemin est plus court, il remplace les chemins actuels dans l'entrée de la table de routage.  Si le nouveau chemin est le même coût, il est ajouté à la liste des chemins de l'entrée de la table de routage.
Les chemins intra-zone qui utilisent des zones autres que les zones de réseau fédérateur sont toujours préférés. Les autres chemins, les chemins de backbone intra-zone et les chemins inter-zone, sont de préférence égale.
R3 et R4 redistribuent le même réseau 172.16.1.0/24 avec la même métrique que la route externe de type E2 OSPF. R6 préfère la route annoncée par R3, car la métrique de transfert vers l'ASBR R3 est inférieure à R4 et le tronçon suivant pour 172.16.1.0/24 est R1. (Selon la RFC 1583, la sélection du chemin est basée uniquement sur le coût.)
R6#sh ip ospf border-routers
OSPF Router with ID (192.168.6.6) (Process ID 1)
Base Topology (MTID 0)
Internal Router Routing Table
Codes: i - Intra-area route, I - Inter-area route
i 192.168.4.4 [51] via 192.168.56.5, GigabitEthernet0/0, ASBR, Area 2, SPF 17
>>>> Cost is 51 to reach R4 ASBR.
i 192.168.1.1 [1] via 192.168.16.1, GigabitEthernet0/1, ABR, Area 2, SPF 17
I 192.168.3.3 [42] via 192.168.16.1, GigabitEthernet0/1, ASBR, Area 2, SPF 17
>>>> Cost is 42 to reach R3 ASBR
R6#sh ip route 172.16.1.0
Routing entry for 172.16.1.0/24
Known via "ospf 1", distance 110, metric 20, type extern 2, forward metric 42
Last update from 192.168.16.1 on GigabitEthernet0/1, 00:02:13 ago
Routing Descriptor Blocks:
* 192.168.16.1, from 192.168.3.3, 00:02:13 ago, via GigabitEthernet0/1
Route metric is 20, traffic share count is 1
R1 préfère la route annoncée par R4 malgré le coût plus élevé car il s’agit d’une route intra-zone vers ASBR. La route ne traverse pas la zone de backbone et le tronçon suivant est R6 (conformément à la RFC 2328).
R1-NXOS# sh ip ospf border-routers
OSPF Process ID 1 VRF default, Internal Routing Table
Codes: i - Intra-area route, I - Inter-area route
intra 192.168.2.2 [40], ABR, Area 0.0.0.0, SPF 18
via 192.168.12.2, Eth4/43
inter 192.168.3.3 [41], ASBR, Area 0.0.0.0, SPF 18 >>>> Cost is 41
via 192.168.12.2, Eth4/43
intra 192.168.4.4 [91], ASBR, Area 0.0.0.2, SPF 18 >>>> Cost is 91
via 192.168.16.6, Eth4/44
switch-R1-NXOS# sh ip route 172.16.1.0
IP Route Table for VRF "default"
'*' denotes best ucast next-hop
'**' denotes best mcast next-hop
'[x/y]' denotes [preference/metric]
'%' in via output denotes VRF
172.16.1.0/24, ubest/mbest: 1/0
*via 192.168.16.6, Eth4/44, [110/20], 00:10:41, ospf-1, type-2
Cela entraîne une boucle dans le réseau, car R6 envoie les paquets à R1 et R1 les renvoie à R6.
R5#traceroute 172.16.1.1 numeric
Type escape sequence to abort.
Tracing the route to 172.16.1.1
VRF info: (vrf in name/id, vrf out name/id)
1 192.168.56.6 4 msec 0 msec 0 msec
2 192.168.16.1 4 msec 0 msec 4 msec
3 192.168.16.6 0 msec 4 msec 0 msec
4 192.168.16.1 4 msec 0 msec 4 msec
5 192.168.16.6 0 msec 4 msec 0 msec
Comme vous le voyez, le paquet fait boucles entre R1 et R6. Pour résoudre ce problème, vous devez modifier la compatibilité RFC sur NXOS.
R1-NXOS(config)# router ospf 1
R1-NXOS(config-router)# rfc1583compatibility
switch-R1-NXOS# sh ip route 172.16.1.0
IP Route Table for VRF "default"
'*' denotes best ucast next-hop
'**' denotes best mcast next-hop
'[x/y]' denotes [preference/metric]
'%' in via output denotes VRF
172.16.1.0/24, ubest/mbest: 1/0
*via 192.168.12.2, Eth4/43, [110/20], 00:00:40, ospf-1, type-2
Maintenant, R1 le pointe correctement vers R2 et la boucle est supprimée du réseau.
R5#traceroute 172.16.1.1 numeric
Type escape sequence to abort.
Tracing the route to 172.16.1.1
VRF info: (vrf in name/id, vrf out name/id)
1 192.168.56.6 0 msec 4 msec 0 msec
2 192.168.16.1 0 msec 0 msec 0 msec
3 192.168.12.2 4 msec 0 msec 0 msec
4 192.168.23.3 4 msec 0 msec 4 msec
5 192.168.23.3 4 msec 0 msec 4 msec
R1 reçoit une route NSSA-External (Type 7) de R6 et une route External (Type 5) de R2 pour le même préfixe 172.16.1.0/24. R1 préfère le type 7, bien que normalement dans OSPF de type 5 soit préféré au type 7.
R1-NXOS# sh ip ospf database nssa-external 172.16.1.0 detail
OSPF Router with ID (192.168.1.1) (Process ID 1 VRF default)
Type-7 AS External Link States (Area 0.0.0.2)
LS age: 914
Options: 0x28 (No TOS-capability, Type 7/5 translation, DC)
LS Type: Type-7 AS-External
Link State ID: 172.16.1.0 (Network address)
Advertising Router: 192.168.4.4 >>>>> Type 7 originated by R4
and installed in the RIB.
LS Seq Number: 0x80000001
Checksum: 0x3696
Length: 36
Network Mask: /24
Metric Type: 2 (Larger than any link state path)
TOS: 0
Metric: 20
Forward Address: 192.168.45.4
External Route Tag: 0>
R1-NXOS# sh ip ospf database external 172.16.1.0 detail
OSPF Router with ID (192.168.1.1) (Process ID 1 VRF default)
Type-5 AS External Link States
LS age: 853
Options: 0x2 (No TOS-capability, No DC)
LS Type: Type-5 AS-External
Link State ID: 172.16.1.0 (Network address)
Advertising Router: 192.168.1.1 >>>>> Since Type 7 is installed
in the RIB, it was converted to type 5
LS Seq Number: 0x80000001
Checksum: 0xb545
Length: 36
Network Mask: /24
Metric Type: 2 (Larger than any link state path)
TOS: 0<
Metric: 20
Forward Address: 192.168.45.4
External Route Tag: 0<
LS age: 596
Options: 0x20 (No TOS-capability, DC)
LS Type: Type-5 AS-External
Link State ID: 172.16.1.0 (Network address)
Advertising Router: 192.168.3.3 >>>>>> Type 5 is also received from R3
LS Seq Number: 0x80000002
Checksum: 0x2250
Length: 36
Network Mask: /24
Metric Type: 2 (Larger than any link state path)>
TOS: 0
Metric: 20<>
Forward Address: 0.0.0.0
External Route Tag: 0
R1-NXOS# sh ip route 172.16.1.0
IP Route Table for VRF "default"
'*' denotes best ucast next-hop
'**' denotes best mcast next-hop
'[x/y]' denotes [preference/metric]
'%<string>' in via output denotes VRF <string>
172.16.1.0/24, ubest/mbest: 1/0
*via 192.168.16.6, Eth4/44, [110/20], 00:16:54, ospf-1, nssa type-2 >>>> Type 7
route is installed in RIB.
Comme R1 ne dispose pas de la commande rfc1583compatibilité configurée dans le processus du routeur OSPF et que l'ID de routeur avancé de la LSA de type 5 est accessible dans la zone 0 (routeur de backbone), OSPF prend toujours le chemin de la route via la zone non de backbone. Dans ce cas, le tronçon suivant est choisi dans la zone 2 (conformément à la RFC 2328).
R1-NXOS(config)# router ospf 1
R1-NXOS(config-router)# rfc1583compatibility
R1-NXOS# sh ip route 172.16.1.0
IP Route Table for VRF "default"
'*' denotes best ucast next-hop
'**' denotes best mcast next-hop
'[x/y]' denotes [preference/metric]
'%<string>' in via output denotes VRF <string>
172.16.1.0/24, ubest/mbest: 1/0
*via 192.168.12.2, Eth4/43, [110/20], 00:00:04, ospf-1, type-2 >>>> Type 5
route is installed in RIB.
