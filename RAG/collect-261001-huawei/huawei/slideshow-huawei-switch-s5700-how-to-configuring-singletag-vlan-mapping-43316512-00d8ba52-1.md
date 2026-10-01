---
id: collect-261001-huawei/huawei/slideshow-huawei-switch-s5700-how-to-configuring-singletag-vlan-mapping-43316512-00d8ba52-1
title: "slideshow-huawei-switch-s5700-how-to-configuring-singletag-vlan-mapping-43316512-00d8ba52"
domain: huawei
role: reference
task: reference
actors: ["Google", "Huawei"]
dates: []
keywords: ["agents", "benchmarks", "cost", "incident", "mcp", "pricing"]
source: docs/RAG/collect-261001-huawei/slideshow-huawei-switch-s5700-how-to-configuring-singletag-vlan-mapping-43316512-00d8ba52.md
source_anchor: ""
source_lines: [1, 82]
sha256: da521adbf5a8c93026a1ed98d81fb3c04b52c2b9f2524b5e9d554002c573ab8a
---

# slideshow-huawei-switch-s5700-how-to-configuring-singletag-vlan-mapping-43316512-00d8ba52

Télécharger en tant que PDF, PPTX
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
Publicité
 Ouvre dans une nouvelle fenêtre Ouvre un site Web externe Ouvre un site Web externe dans une nouvelle fenêtre 
      Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant :     Politique de confidentialité    
   Préférences en matière de conservation des données
Passer au contenu principal
Téléchargé parIPMAX s.r.l.
PDF, PPTX27 993 vues
The document discusses configuring single-tag VLAN mapping on Huawei S5700 switches to allow communication between client devices in different VLANs. It involves creating VLANs 10, 20 and 100 on switches, adding ports to the VLANs, and configuring single-tag VLAN mapping on trunk ports between switches to map VLANs 10 and 20 to VLAN 100 to allow inter-VLAN communication. The configuration is verified by pinging from a client in VLAN 10 to a client in VLAN 20 to confirm connectivity across the VLANs.
PDF
HUAWEI Switch HOW-TO - Configuring link aggregation in static LACP mode
parIPMAX s.r.l.
8 diapositives36K vues
003 obf600105 gpon ma5608 t basic operation and maintenance v8r15 issue1.02 (...
79 diapositives3.8K vues
Cisco Switch Configuration Basics for Beginners | CCNA Certification
parNiya Kohli
15 diapositives26 vues
Huawei ARG3 Router How To - Troubleshooting OSPF: Netmask mismatch
parIPMAX s.r.l.
11 diapositives1.3K vues
Huawei ARG3 Router How To - Troubleshooting OSPF: Router ID Confusion
parIPMAX s.r.l.
9 diapositives1.5K vues
Huawei SAN Storage How To - Configuring the i-SCSI Communication Protocol
parIPMAX s.r.l.
17 diapositives3.8K vues
Huawei SAN Storage How To - ISM management application setup
parIPMAX s.r.l.
14 diapositives3.1K vues
Huawei SAN Storage How To - Assigning Management IP Address
parIPMAX s.r.l.
10 diapositives6.2K vues
Comprehensive Guide to Building and Deploying Applications on Google Cloud Pl...
33 diapositives64 vues
PPTX
Comprehensive Guide to Vulnerability Remediation and Incident Response Workflow
8 diapositives29 vues
Agile Gurugram & Delhi National Capital Region 2026 _ The AI Shift_ Skill Gap...
parAgileNetwork
11 diapositives27 vues
Comprehensive Introduction to Artificial Intelligence: Concepts, History, and...
14 diapositives7 vues
How Much Does Data Annotation Cost? Pricing Models, Rates & Budget Benchmarks
parHabile  Data
12 diapositives44 vues
When AI Agents Work Together: Exploring A2A, MCP, and Connected AI Frameworks...
24 diapositives10 vues
How to Make AI-Assisted Writing More Authentic: A Practical Guide to Original...
parAdeel Ali 
16 diapositives33 vues
Innovative Applications of Nanotechnology in Fisheries and Aquaculture Value ...
parB. BHASKAR
21 diapositives19 vues
Foundations and Future of AI & Autonomous Systems: From Basics to Machine Aut...
parAdeel Ali 
9 diapositives26 vues
Potential Scope, Advantages, and Challenges of RAS and Bio-floc Technology in...
parB. BHASKAR
6 diapositives15 vues
- 1. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping HUAWEI SWITCH S5700 HOW TO
- 2. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping VLAN Mapping is a function that maps the customer VLAN ID to the carrier VLAN ID by replacing VLAN tags of data frames. VLAN mapping implements VLAN aggregation (users in different vlan can communicate) and allows service data to be transmitted according to carriers' network plans, saving carrier vlan resources. The S5700 supports the following VLAN mapping features: - Single-tag VLAN mapping based on the interface and VLAN - Double-tag VLAN mapping based on the interface and VLAN. We will explain single-tag vlan mappinq through the simple network scenario shown in the picture. We will configure VLAN mapping on SwitchA and SwitchB, so Client PCs in VLAN 10 will be able to communicate with Client PCs in VLAN 20, through carrier’s VLAN 100. General infos and application environment
- 3. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping 1. Create vlans on the switches Configuring vlans on the switches <SwitchA> system-view [SwitchA] vlan batch 10 100 Enter the configuration view Create vlans 10 and 100 <SwitchB> system-view [SwitchB] vlan batch 20 100 <SwitchC> system-view [SwitchC] vlan batch 100 <SwitchD> system-view [SwitchD] vlan batch 100
- 4. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping 2. Add interfaces to vlans Adding interfaces to vlans (1) [SwitchA] interface gigabitethernet 0/0/1 [SwitchA-GigabitEthernet0/0/1] port link-type access [SwitchA-GigabitEthernet0/0/1] port default vlan 10 [SwitchA-GigabitEthernet0/0/1] quit [SwitchA] interface gigabitethernet 0/0/2 [SwitchA-GigabitEthernet0/0/2] port link-type access [SwitchA-GigabitEthernet0/0/2] port default vlan 10 [SwitchA-GigabitEthernet0/0/2] quit [SwitchA] interface gigabitethernet 0/0/3 [SwitchA-GigabitEthernet0/0/3] port link-type trunk [SwitchA-GigabitEthernet0/0/3] port trunk allow-pass vlan 10 [SwitchA-GigabitEthernet0/0/3] quit Enter the interface view Configure the port type Configure the default vlan
- 5. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping Adding interfaces to vlans (2) [SwitchB] interface gigabitethernet 0/0/1 [SwitchB-GigabitEthernet0/0/1] port link-type access [SwitchB-GigabitEthernet0/0/1] port default vlan 20 [SwitchB-GigabitEthernet0/0/1] quit [SwitchB] interface gigabitethernet 0/0/2 [SwitchB-GigabitEthernet0/0/2] port link-type access [SwitchB-GigabitEthernet0/0/2] port default vlan 20 [SwitchB-GigabitEthernet0/0/2] quit [SwitchB] interface gigabitethernet 0/0/3 [SwitchB-GigabitEthernet0/0/3] port link-type trunk [SwitchB-GigabitEthernet0/0/3] port trunk allow-pass vlan 20 [SwitchB-GigabitEthernet0/0/3] quit
- 6. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping Adding interfaces to vlans (3) [SwitchC] interface gigabitethernet 0/0/1 [SwitchC-GigabitEthernet0/0/1] port link-type trunk [SwitchC-GigabitEthernet0/0/1] port trunk allow-pass vlan 100 [SwitchC-GigabitEthernet0/0/1] quit [SwitchD] interface gigabitethernet 0/0/1 [SwitchD-GigabitEthernet0/0/1] port link-type trunk [SwitchD-GigabitEthernet0/0/1] port trunk allow-pass vlan 100 [SwitchD-GigabitEthernet0/0/1] quit
- 7. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping Configuring single-tag vlan mapping on the switches [SwitchA] interface gigabitethernet 0/0/3 [SwitchA-GigabitEthernet0/0/3] qinq vlan-traslation enable [SwitchA-GigabitEthernet0/0/3] port vlan-mapping vlan 100 map-vlan 10 [SwitchA-GigabitEthernet0/0/3] quit [SwitchB] interface gigabitethernet 0/0/3 [SwitchB-GigabitEthernet0/0/3] qinq vlan-traslation enable [SwitchB-GigabitEthernet0/0/3] port vlan-mapping vlan 100 map-vlan 20 [SwitchB-GigabitEthernet0/0/3] quit Enter the interface view Enable vlan mapping Define vlan mapping
- 8. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping Checking the configuration 3. Check the configuration (ping from Client1 to Client4) Client1> ping 192.168.100.5 Ping 192.168.100.5: 32 data bytes, Press Ctrl_C to break From 192.168.100.5: bytes=32 seq=1 ttl=128 time=78 ms From 192.168.100.5: bytes=32 seq=2 ttl=128 time=93 ms From 192.168.100.5: bytes=32 seq=3 ttl=128 time=109 ms From 192.168.100.5: bytes=32 seq=4 ttl=128 time=78 ms From 192.168.100.5: bytes=32 seq=5 ttl=128 time=94 ms --- 192.168.100.5 ping statistics --- 5 packet(s) transmitted 5 packet(s) received 0.00% packet loss round-trip min/avg/max = 78/90/109 ms Client1>
- 9. HUAWEI SWITCH S5700- HOW TO Configuring single-tag vlan mapping More needs? See hints on www.ipmax.it Or email us your questions to info_ipmax@ipmax.it
