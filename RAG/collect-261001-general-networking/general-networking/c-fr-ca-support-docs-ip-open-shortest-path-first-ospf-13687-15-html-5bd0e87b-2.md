---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b-2
title: "c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b.md
source_anchor: ""
source_lines: [80, 168]
sha256: d4832f2ae07448c10edf3915f974bea70f8efa18062b95bb7991e4bb499af769
---

# c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b

    Router2#show ip ospf neighbor
      Neighbor ID   Pri   State   Dead Time   Address   Interface
      3.3.3.3        1   FULL/ -  00:00:30    3.3.3.3    Serial0
Router2# show ip ospf interface serial 0
Serial0 is up, line protocol is up
  Internet Address 0.0.0.0/24, Area 0
  Process ID 1, Router ID 2.2.2.2, Network Type POINT_TO_POINT, Cost: 64
  Transmit Delay is 1 sec, State POINT_TO_POINT,
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:08
  Index 2/2, flood queue length 0
  Next 0x0(0)/0x0(0)
  Last flood scan length is 1, maximum is 1
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 1, Adjacent neighbor count is 1
    Adjacent with neighbor 3.3.3.3
  Suppress hello for 0 neighbor(s) 
    
   Cet exemple montre le résultat de la commande show ip route sur Router1 avec encapsulation PPP et l'utilisation d'interfaces non numérotées.
 
    Router1#show ip route
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
     2.0.0.0/32 is subnetted, 1 subnets
C       2.2.2.2 is directly connected, Serial0
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Loopback0 
    
   Cet exemple montre comment afficher le résultat de la commande show ip route sur le routeur 2 avec encapsulation PPP et l'utilisation d'interfaces non numérotées.
 
    Router2#show ip route
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
     2.0.0.0/32 is subnetted, 1 subnets
C       2.2.2.2 is directly connected, Loopback0
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Serial0
 
    
   Cet exemple montre comment afficher le résultat de la commande show ip route sur Router1 avec encapsulation HDLC et l'utilisation d'interfaces non numérotées.
 
    Router1#show ip route
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/65] via 2.2.2.2, 00:00:08, Serial0
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Loopback0 
    
   Cet exemple montre comment afficher le résultat de la commande show ip route sur Router2 avec encapsulation HDLC et l'utilisation d'interfaces non numérotées.
 
    Router1#show ip route
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
        2.0.0.0/32 is subnetted, 1 subnets
C       2.2.2.2 is directly connected, Loopback0
     3.0.0.0/32 is subnetted, 1 subnets
O       3.3.3.3 [110/65] via 3.3.3.3, 00:01:28, Serial0 
    
   Remarque : Le résultat de la commande show ip route peut différer entre les encapsulations PPP et HDLC lorsque la configuration IP non numérotée est utilisée sur les interfaces série. Le protocole PPP installe une route hôte vers l’adresse IP qui est utilisée sur l’interface série à l’autre extrémité comme réseau directement connecté. Si le même préfixe est également appris via OSPF comme dans cette configuration, il s'affiche uniquement en tant que route connectée (comme le montre la sortie show ip route). En effet, les routes connectées ont une distance administrative inférieure à celle du protocole OSPF et sont plus privilégiées. Vous pouvez modifier ce comportement lorsque vous émettez la commande no peer neighbor-route sous les interfaces série qui empêche l'installation d'une route hôte et la traite comme une route OSPF.
Ce n’est pas le cas avec HDLC, car il n’installe pas de route hôte. HDLC installe une route OSPF pour l’adresse de l’autre extrémité lorsque l’adresse IP non numérotée est utilisée.
Pour plus d'informations sur le dépannage des problèmes OSPF, référez-vous à Dépannage OSPF.
| Révision | Date de publication | Commentaires | 
|---|---|---|
| 1.0 |                                                                                               06-Jul-2007                                                                                       | Première publication |
