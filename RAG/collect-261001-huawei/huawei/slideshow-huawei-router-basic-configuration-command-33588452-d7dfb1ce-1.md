---
id: collect-261001-huawei/huawei/slideshow-huawei-router-basic-configuration-command-33588452-d7dfb1ce-1
title: "slideshow-huawei-router-basic-configuration-command-33588452-d7dfb1ce"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agents", "incident", "mcp"]
source: docs/RAG/collect-261001-huawei/slideshow-huawei-router-basic-configuration-command-33588452-d7dfb1ce.md
source_anchor: ""
source_lines: [1, 58]
sha256: c76f452ed724f8d28ad514bc23198e4c93d93c079cf9422fad5377352f8f46fb
---

# slideshow-huawei-router-basic-configuration-command-33588452-d7dfb1ce

Téléchargé 603 fois
Playback speed
1x Normal
Back
0.25x
0.5x
1x Normal
1.5x
2x
Skip
Ads by 
Publicité
Publicité
Publicité
 Ouvre dans une nouvelle fenêtre Ouvre un site Web externe Ouvre un site Web externe dans une nouvelle fenêtre 
      Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant :     Politique de confidentialité    
   Préférences en matière de conservation des données
Passer au contenu principal
Téléchargé parHuanetwork
DOC, PDF25 259 vues
This document provides instructions for basic configuration of Huawei routers, including commands to display device information, configure interfaces and IP addresses, set the host name and password, enable routing protocols like RIP and OSPF, and configure static and dynamic routing. It also provides examples of extended access list configuration and address translation. Additional links are included for VPN configuration, PPPOE client setup, and information on purchasing Huawei networking equipment from a distributor.
DOCX
Student Name _________________________________  Date _____________SE.docx
paremelyvalg9
25 diapositives32 vues
DOC
Huanetwork x dsl solution - huawei adsl2+ and vdsl2 solution)
parHuanetwork
12 diapositives972 vues
What are the differences between huawei and cisco wlan products
parHuanetwork
12 diapositives1.7K vues
PDF
OhhPro: Enhancing Community Engagement and Connectivity in Residential Societies
9 diapositives13 vues
PPTX
When AI Agents Work Together: Exploring A2A, MCP, and Connected AI Frameworks...
24 diapositives10 vues
Agile Gurugram & Delhi National Capital Region 2026 _ The AI Shift_ Skill Gap...
parAgileNetwork
11 diapositives27 vues
Comprehensive Guide to Vulnerability Remediation and Incident Response Workflow
8 diapositives29 vues
Centralized Application-Context Aware Firewall with AI/ML for Enhanced Cybers...
parndaindna55
7 diapositives40 vues
HITCON 2026 Slide- Not Just Spies Anymore: DPRK's Espionage Actors Are Coming...
36 diapositives193 vues
Comprehensive MEAN Stack Developer Roadmap for 2026: From Fundamentals to Job...
15 diapositives8 vues
Optimizing Write-Intensive Database Workloads Masterclass: Strategies for Opt...
parScyllaDB
29 diapositives10 vues
- 1. Huawei Router BasicConfiguration Command Huawei Router Basic Configuration Command [Liufei]display version; display version information [Liufei]display current-configuration; display current configuration [Liufei]display interfaces; display interface information; [Liufei]display IP route; show that the routing information; [Liufei]sysname lf123; change the host name [Liufei]super passwrod 123456; setup password; [Liufei]interface serial0; into the interface; [Liufei-serial0]ip address <ip> <mask|mask_len> IP address configured ports; [Liufei-serial0]undo shutdown; activation port; [Liufei]link-protocol HDLC; binds HDLC protocol; [Liufei]user-interface vty 04 [Liufei-ui-vty0-4]authentication-mode password [Liufei-ui-vty0-4]set authentication-mode password simple 222 [Liufei-ui-vty0-4]user privilege level 3 [Liufei-ui-vty0-4]quit [Liufei]debugging HDLC all serial0; to display all of the information; [Liufei]debugging HDLC event serial0; debug event information [Liufei]debugging HDLC packet serial0; show the package information 1
- 2. Static routing: [Liufei]ip route-static<ip><mask>{interface number|nexthop}[value][reject|blackhole] For example: [Liufei]ip route-static 129.1.0.0 16 10.0.0.2 [Liufei]ip route-static 129.1.0.0 255.255.0.0 10.0.0.2 [Liufei]ip route-static 129.1.0.0 16 Serial 2 [Liufei]ip route-static 0.0.0.0 0.0.0.0 10.0.0.2 Dynamic routing: [Liufei]rip; set dynamic routing [Liufei]rip work is allowed to work; [Liufei]rip input set the entrance permit; [Liufei]rip output; set the export permit [Liufei-rip]network 1.0.0.0; set the exchange routing network [Liufei-rip]network all; set the exchange with all network [Liufei-rip]peer ip-address; [Liufei-rip]summary routing aggregation; [Liufei]rip version 1; installed in version 1 [Liufei]rip version 2 multicast; with version 2, multicast [Liufei-Ethernet0]rip split-horizon; horizontal partitioning [Liufei]router ID A.B.C.D; configure router ID [Liufei]ospf enable start OSPF protocol; 2
- 3. [Liufei-ospf]import-route direct; introducingthe direct route [Liufei-Serial0]ospf enable area <area_id> OSPF region; Standard access list command format is as follows: ACL <acl-number> [match-order config|auto]; the default of the former sequence matching. Rule [normal|special]{permit|deny} [source source-addr source-wildcard|any] An example: [Liufei]acl 10 [Liufei-acl-10]rule normal permit source 10.0.0.0 0.0.0.255 [Liufei-acl-10]rule normal deny source any Extended access list configuration command Extended configuration of TCP/UDP protocol access list: Rule {normal|special}{permit|deny}{tcp|udp}source {<ip wild>|any}destination <ip wild>| any} [operate] Extended configuration of ICMP protocol access list: Rule {normal|special}{permit|deny}icmp source {<ip wild>|any]destination {<ip wild>|any] [icmp-code] [logging] Extended access control list the meaning of the operators Equal PortNumber; equal 3
- 4. Greater-than PortNumber; greaterthan Less-than is PortNumber; less than Not-equal PortNumber; unequal Range portnumber1 portnumber2; interval Extended access control list example [Liufei]acl 101 [Liufei-acl-101]rule deny souce any destination any [Liufei-acl-101]rule permit ICMP source any destination any icmp-type echo [Liufei-acl-101]rule permit ICMP source any destination any icmp-type echo-reply [Liufei]acl 102 [Liufei-acl-102]rule permit IP source 10.0.0.1 0.0.0.0 destination 202.0.0.1 0.0.0.0 [Liufei-acl-102]rule deny IP source any destination any [Liufei]acl 103 [Liufei-acl-103]rule permit TCP source any destination 10.0.0.1 0.0.0.0 destination-port equalFTP [Liufei-acl-103]rule permit TCP source any destination 10.0.0.2 0.0.0.0 destination-port equalwww [Liufei]firewall enable [Liufei]firewall default permit|deny [Liufei]int E0 [Liufei-Ethernet0]firewall packet-filter 101 inbound|outbound Address conversion configuration example [Liufei]firewall enable 4
- 5. [Liufei]firewall default permit [Liufei]acl101; the internal host can enter the E0 [Liufei-acl-101]rule deny IP source any destination any [Liufei-acl-101]rule permit IP source 129.38.1.1 0 destination any [Liufei-acl-101]rule permit IP source 129.38.1.2 0 destination any [Liufei-acl-101]rule permit IP source 129.38.1.3 0 destination any [Liufei-acl-101]rule permit IP source 129.38.1.4 0 destination any [Liufei-acl-101]quit [Liufei]int E0 [Liufei-Ethernet0]firewall packet-filter 101 inbound [Liufei]acl 102; the external host specific and more than 1024 port packet is allowed to enter the S0 [Liufei-acl-102]rule deny IP source any destination any [Liufei-acl-102]rule permit TCP source 202.39.2.3 0 destination 202.38.160.1 0 [Liufei-acl-102]rule permit TCP source any destination 202.38.160.1 0 destination-port great-than 1024 [Liufei-acl-102]quit [Liufei]int S0 [Liufei-Serial0]firewall packet-filter 102 inbound; 202.38.160.1 is the egress router IP. [Liufei-Serial0]nat outbound 101 interface; Easy IP, ACL 101 allows IP from this interfacetransform source address. 5
