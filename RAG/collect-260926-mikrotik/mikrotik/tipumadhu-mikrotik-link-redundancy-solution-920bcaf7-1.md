---
id: collect-260926-mikrotik/mikrotik/tipumadhu-mikrotik-link-redundancy-solution-920bcaf7-1
title: "tipumadhu-mikrotik-link-redundancy-solution-920bcaf7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "benchmarks", "cost", "incident", "pricing"]
source: docs/RAG/lot-mikrotik/forum/misc/tipumadhu-mikrotik-link-redundancy-solution-920bcaf7.md
source_anchor: ""
source_lines: [1, 64]
sha256: ff822f498bc4ff1c5874656270cbd3302d76c459f004d886ae31f7bbdbfbb58e
---

# tipumadhu-mikrotik-link-redundancy-solution-920bcaf7

Téléchargé 27 fois
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
Publicité
 Ouvre dans une nouvelle fenêtre Ouvre un site Web externe Ouvre un site Web externe dans une nouvelle fenêtre 
      Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant :     Politique de confidentialité    
   Préférences en matière de conservation des données
Passer au contenu principal
Téléchargé parS M  Tipu
PPTX, PDF2 508 vues
This document provides information on various techniques for load balancing and redundancy, including: - Load balancing over multiple gateways using policy routing based on client IP address, firewall mangle rules, and default routes. - VRRP (Virtual Router Redundancy Protocol) for high availability using a virtual IP address, master and backup routers, and fast failover detection. - Mikrotik PCC (Per Connection Classifier) load balancing which divides traffic into equal streams using a hashing algorithm on packet header fields and marks connections for policy routing out specific gateways.
PDF
L2/L3 für Fortgeschrittene - Helle und dunkle Magie im Linux-Netzwerkstack
36 diapositives5.4K vues
PPTX
Innovative Applications of Nanotechnology in Fisheries and Aquaculture Value ...
parB. BHASKAR
21 diapositives19 vues
Potential Scope, Advantages, and Challenges of RAS and Bio-floc Technology in...
parB. BHASKAR
6 diapositives15 vues
BotGentz AI Playbook: Seamless AI Agents Integrated into Your Workflow
parZloadr
8 diapositives33 vues
OhhPro: Enhancing Community Engagement and Connectivity in Residential Societies
9 diapositives13 vues
Comprehensive Azure Migration Readiness Guide: Inventory, Security, Networkin...
paracaptacloud
10 diapositives56 vues
Comprehensive Guide to Vulnerability Remediation and Incident Response Workflow
8 diapositives29 vues
Centralized Application-Context Aware Firewall with AI/ML for Enhanced Cybers...
parndaindna55
7 diapositives40 vues
Your npm dependencies can see everything | How we lock them down with SES + L...
parAmbire
23 diapositives7 vues
HITCON 2026 Slide- Not Just Spies Anymore: DPRK's Espionage Actors Are Coming...
36 diapositives193 vues
How Much Does Data Annotation Cost? Pricing Models, Rates & Budget Benchmarks
parHabile  Data
12 diapositives44 vues
- 1. S M YeaserHossain Tipu
- 2. CLASS -12 Link Redundancysolution
- 3. Topics We CoverHere. • Load Balancing • Load Balancing with fail over. • VRRP (High Availability) • Mikrotik PCC Load Balancing. • Load Balance and Redundancy with OSPF. • Load Balance other mechanism. • Bandwidth merge of different link. We also covered some load balancing and redundancy technique in our routing lecture.
- 4. Load Balancing overMultiple Gateways The typical situation where you got one router and want to connect to two ISPs, Of course, you want to do load balancing! There are several ways how to do it. Depending on the particular situation, you may find one best suited for you.
- 5. Policy Routing basedon Client IP Address If you have a number of hosts, you may group them by IP addresses. Then, depending on the source IP address, send the traffic out through Gateway #1 or #2. This is not really the best approach, giving you perfect load balancing, but it's easy to implement, and gives you some control too. Let us assume we use for our workstations IP addresses from network 192.168.100.0/24. The IP addresses are assigned as follows: 192.168.100.1-127 are used for Group A workstations 192.168.100.128-253 are used for Group B workstations 192.168.100.254 is used for the router. All workstations have IP configuration with the IP address from the relevant group, they all have network mask 255.255.255.0, and 192.168.100.254 is the default gateway for them. We will talk about DNS servers later. Now, when we have workstations divided into groups, we can refer to them using subnet addressing: Group A is 192.168.100.0/25, i.e., addresses 192.168.100.0-127 Group B is 192.168.100.128/25, i.e., addresses 192.168.100.128-255
- 6. We need toadd two IP Firewall Mangle rules to mark the packets originated from Group A or Group B workstations. For Group A, specify • Chain prerouting and Src. Address 192.168.100.0/25 • Action mark routing and New Routing Mark GroupA.
- 7. • It isa good practice to add a comment as well. Your mangle rules might be interesting for someone else and for yourself as well after some time. • For Group B, specify • Chain prerouting and Src. Address 192.168.100.128/25 • Action mark routing and New Routing Mark GroupB All IP traffic coming from workstations is marked with the routing marks GroupA or GroupB. We can use these marks in the routing table.
- 8. Next, we shouldspecify two default routes (destination 0.0.0.0/0) with appropriate routing marks and gateways:
- 9. This thing isnot going to work, unless you do masquerading for your LAN! The simplest way to do it is by adding one NAT rule for Src. Address 192.168.100.0/24 and Action masquerade:
- 10. ECMP load balancingwith masquerade This example is improved (different) version of round-robin load balancing example. It adds persistent user sessions, i.e. a particular user would use the same source IP address for all outgoing connections. Consider the following network layout:
- 11. Quick Start / ipaddress add address=192.168.0.1/24 network=192.168.0.0 broadcast=192.168.0.255 interface=Local add address=10.111.0.2/24 network=10.111.0.0 broadcast=10.111.0.255 interface=wlan2 add address=10.112.0.2/24 network=10.112.0.0 broadcast=10.112.0.255 interface=wlan1 / ip route add dst-address=0.0.0.0/0 gateway=10.111.0.1,10.112.0.1 check-gateway=ping / ip firewall nat add chain=srcnat out-interface=wlan1 action=masquerade add chain=srcnat out-interface=wlan2 action=masquerade / ip firewall mangle add chain=input in-interface=wlan1 action=mark-connection new-connection-mark=wlan1_conn add chain=input in-interface=wlan2 action=mark-connection new-connection-mark=wlan2_conn add chain=output connection-mark=wlan1_conn action=mark-routing new-routing-mark=to_wla1 add chain=output connection-mark=wlan1_conn action=mark-routing new-routing-mark=to_wla2 / ip route add dst-address=0.0.0.0/0 gateway=10.111.0.1 routing-mark=to_wla1 add dst-address=0.0.0.0/0 gateway=10.112.0.1 routing-mark=to_wla2
- 12. Explanation First we givea code snippet and then explain what it actually does. IP Addresses / ip address add address=192.168.0.1/24 network=192.168.0.0 broadcast=192.168.0.255 interface=Local add address=10.111.0.2/24 network=10.111.0.0 broadcast=10.111.0.255 interface=wlan2 add address=10.112.0.2/24 network=10.112.0.0 broadcast=10.112.0.255 interface=wlan1 The router has two upstream (WAN) interfaces with the addresses of 10.111.0.2/24 and 10.112.0.2/24. The LAN interface has the name "Local" and IP address of 192.168.0.1/24. NAT / ip firewall nat add chain=srcnat out-interface=wlan1 action=masquerade add chain=srcnat out-interface=wlan2 action=masquerade As routing decision is already made we just need rules that will fix src-addresses for all outgoing packets. if this packet will leave via wlan1 it will be NATed to 10.112.0.2/24, if via wlan2 then NATed to 10.111.0.2/24 Routing / ip route add dst-address=0.0.0.0/0 gateway=10.111.0.1,10.112.0.1 check-gateway=ping This is typical ECMP (Equal Cost Multi-Path) gateway with check-gateway. ECMP is "persistent per-connection load balancing" or "per-src-dst-address combination load balancing". As soon as one of the gateway will not be reachable, check-gateway will remove it from gateway list. And you will have a "failover" effect.
