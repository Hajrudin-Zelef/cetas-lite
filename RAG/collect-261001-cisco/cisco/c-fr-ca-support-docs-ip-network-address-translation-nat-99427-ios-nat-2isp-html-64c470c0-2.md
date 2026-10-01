---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0-2
title: "c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0.md
source_anchor: ""
source_lines: [26, 53]
sha256: 15d558933c41d99d6de65d9ae4f219283dc88ee6b2bb27aadaeabfb9b93cc237
---

# c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0

      Router# sh ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, 
       L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, 
       U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is 172.16.108.1 to network 0.0.0.0
C    192.168.108.0/24 is directly connected, Vlan1
     172.16.0.0/24 is subnetted, 2 subnets
C       172.16.108.0 is directly connected, 
        FastEthernet4
C       172.16.106.0 is directly connected, Vlan106
S*   0.0.0.0/0 [1/0] via 172.16.108.1
               [1/0] via 172.16.106.1
Router#
 
      Après avoir configuré le routeur Cisco IOS avec la fonction NAT, si les connexions ne fonctionnent pas, assurez-vous des éléments suivants :
NAT est appliqué convenablement sur les interfaces externes et internes.
La configuration NAT est complète et la liste reflète le trafic qui doit être soumis à NAT.
Plusieurs itinéraires vers Internet/WAN sont disponibles.
Si vous utilisez le suivi de route pour vous assurer que les connexions Internet sont disponibles, vérifiez l'état du suivi de route.
| Révision | Date de publication | Commentaires | 
|---|---|---|
| 1.0 |                                                                                               15-Nov-2007                                                                                       | Première publication |
