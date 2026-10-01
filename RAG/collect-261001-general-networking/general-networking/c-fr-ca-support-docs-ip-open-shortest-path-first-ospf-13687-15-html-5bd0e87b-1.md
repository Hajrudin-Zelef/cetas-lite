---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b-1
title: "c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "diffusion"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b.md
source_anchor: ""
source_lines: [1, 79]
sha256: 0d3b51b7b6d55b3b12978706bb5009622fa402a8d08a6ba27499c25c43bb3d8c
---

# c-fr-ca-support-docs-ip-open-shortest-path-first-ospf-13687-15-html-5bd0e87b

Sur les liaisons point-à-point, telles que le High-Level Data Link Control (HDLC) et le protocole point-à-point (PPP), Open Shortest Path First (OSPF) fonctionne comme un type de réseau point-à-point. Ce type de réseau est activé par défaut. Ce document propose des exemples de configuration pour OSPF sur une liaison point-à-point. Les autres types de réseau que l’OSPF prend en charge sont les réseaux point-à-multipoint, de diffusion et de non-diffusion. Afin de vérifier le type de réseau d’une interface qui exécute OSPF, exécutez la commande show ip ospf interface.
Les exemples décrits dans ce document fonctionnent avec les encapsulations HDLC et PPP.
Les lecteurs de ce document doivent avoir une connaissance de base du protocole de routage OSPF. Pour plus d'informations sur le protocole de routage OSPF, reportez-vous à la documentation Open Shortest Path First.
Les informations de ce document s'appliquent à ces versions de logiciel et de matériel.
Routeurs Cisco 2500
Logiciel Cisco IOS® version 12.2(27) exécuté sur les routeurs
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, make sure that you understand the potential impact of any command.
For more information on document conventions, refer to the Cisco Technical Tips Conventions.
Cette section vous présente les informations que vous pouvez utiliser pour configurer les routeurs qui exécutent OSPF et qui sont connectés via une liaison série point à point avec l'utilisation de l'encapsulation HDLC avec des adresses IP configurées sur des interfaces série.
Remarque : Afin de trouver des informations supplémentaires sur les commandes utilisées dans ce document, utilisez l'Outil de recherche de commandes (clients enregistrés uniquement).
Cette section utilise ces configurations .
| Routeur 1 | 
|---|
| Router1 ! interface Loopback0   ip address 3.3.3.3 255.255.255.255 ! interface Serial0  ip address 1.1.1.2 255.255.255.0 ! router ospf 1 network 1.1.1.0 0.0.0.255 area 0  !--- Configures the Serial Interface S0 under OSPF area 0.  ! | 
| Routeur 2 | 
|---|
| Router2 ! interface Loopback0  ip address 2.2.2.2 255.255.255.255 ! interface Serial0  ip address 1.1.1.1 255.255.255.0  clockrate 2000000 ! router ospf 1 network 1.1.1.0 0.0.0.255 area 0  !--- Configures the Serial Interface S0 under OSPF area 0..   | 
Cette section présente les informations que vous pouvez utiliser pour configurer les routeurs qui exécutent OSPF comme protocole de routage et qui sont connectés via une liaison série point à point avec encapsulation PPP et avec des interfaces série non numérotées. Les interfaces série non numérotées sont des interfaces qui ne possèdent pas leur propre adresse IP. Ces interfaces empruntent l'adresse IP à une autre interface du routeur dont l'adresse IP est configurée. Pour plus d'informations sur les interfaces série non numérotées, référez-vous à Présentation et configuration de la commande ip unnumbered.
| Routeur 1 | 
|---|
| !  interface Loopback0  ip address 3.3.3.3 255.255.255.255 ! interface Serial0  ip unnumbered loopback 0  !--- Configures Serial 0 as an unnumbered interface.   encapsulation ppp  !--- Configures the Encapsulation on the interface as PPP.  ! router ospf 1   network 3.3.3.0 0.0.0.255 area 0  !--- Configures the Loopback Interface L0 under OSPF area 0.  ! | 
| Routeur 2 | 
|---|
| ! interface Loopback0  ip address 2.2.2.2 255.255.255.255 ! interface Serial0  ip unnumbered loopback 0  !--- Configures Serial 0 as an unnumbered interface.   encapsulation ppp  !--- Configures the Encapsulation on the interface as PPP.  ! router ospf 1   network 2.2.2.0 0.0.0.255 area 0  !--- Configures the Loopback Interface L0 under OSPF area 0..  ! | 
Cette section fournit des informations que vous pouvez utiliser pour confirmer que votre configuration fonctionne correctement.
Certaines commandes show sont prises en charge par l'Output Interpreter Tool (clients enregistrés uniquement), qui vous permet de voir une analyse de la sortie de la commande show.
Il s'agit de la sortie de commande show ip ospf neighbor émise sur Router1.
 
    Router1#show ip ospf neighbor
   Neighbor ID   Pri   State       Dead Time       Address        Interface
     2.2.2.2      1   FULL/ -       00:00:32       1.1.1.1         Serial 0 
    
   Ce résultat montre que Router1 a un voisin avec un ID de routeur de 2.2.2.2 et a formé une contiguïté complète avec ce voisin. L’adresse 1.1.1.1 est l’adresse IP de l’interface Serial0 sur Router2.
Pour plus d'informations sur la commande show ip ospf neighbor, référez-vous à Que révèle la commande show ip ospf neighbor ?
Ceci est le résultat de la commande show ip ospf interface serial 0 sur Router1. Ce résultat montre le type de réseau OSPF pour l'interface Serial0 en tant que point à point.
 
    Router1# show ip ospf interface serial 0 Serial0 is up, line protocol is up Internet Address 1.1.1.2/24, Area 0 Process ID 1, Router ID 3.3.3.3, Network Type POINT_TO_POINT, Cost: 64 Transmit Delay is 1 sec, State POINT_TO_POINT, Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5 Hello due in 00:00:05 Index 1/1, flood queue length 0 Next 0x0(0)/0x0(0) Last flood scan length is 1, maximum is 1 Last flood scan time is 0 msec, maximum is 0 msec Neighbor Count is 1, Adjacent neighbor count is 1 Adjacent with neighbor 2.2.2.2 Suppress hello for 0 neighbor(s)
De même, cet exemple montre comment afficher le résultat des commandes show ip ospf neighbor et show ip ospf interface sur Router2.
 
    Router2# show ip ospf neighbor
   Neighbor ID     Pri    State     Dead Time     Address      Interface
    3.3.3.3         1     FULL/ -    00:00:32      1.1.1.2      Serial0
Router2# show ip ospf interface serial 0
   Serial0 is up, line protocol is up 
   Internet Address 1.1.1.1/24, Area 0 
   Process ID 1, Router ID 2.2.2.2, Network Type POINT_TO_POINT, Cost: 64
   Transmit Delay is 1 sec, State POINT_TO_POINT,
   Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:05
   Index 1/1, flood queue length 0
   Next 0x0(0)/0x0(0)
   Last flood scan length is 1, maximum is 1
   Last flood scan time is 0 msec, maximum is 0 msec
   Neighbor Count is 1, Adjacent neighbor count is 1 
    Adjacent with neighbor 3.3.3.3
   Suppress hello for 0 neighbor(s) 
    
   La commande show ip ospf neighbor sur Router1 et Router2 confirme la contiguïté totale. Ce résultat montre le type de réseau OSPF pour l'interface Serial0 en tant que point à point. Vous pouvez également voir que l’adresse Internet est 0.0.0.0 pour l’interface série des deux routeurs. En effet, une adresse IP est empruntée à l'interface de bouclage (dans ce cas) en vertu de la commande ip unnumbered.
 
    Router1#show ip ospf neighbor
      Neighbor ID  Pri   State    Dead Time    Address      Interface
      2.2.2.2      1   FULL/ -    00:00:37    2.2.2.2       Serial0
Router1# show ip ospf interface serial 0
Serial0 is up, line protocol is up
  Internet Address 0.0.0.0/24, Area 0
  Process ID 1, Router ID 3.3.3.3, Network Type POINT_TO_POINT, Cost: 64
  Transmit Delay is 1 sec, State POINT_TO_POINT,
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:08
  Index 2/2, flood queue length 0
  Next 0x0(0)/0x0(0)
  Last flood scan length is 1, maximum is 1
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 1, Adjacent neighbor count is 1
    Adjacent with neighbor 2.2.2.2
  Suppress hello for 0 neighbor(s) 
    
   Cet exemple montre comment afficher le résultat des commandes show ip ospf neighbor et show ip ospf interface serial 0 sur Router2.
 
