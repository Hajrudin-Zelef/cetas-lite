---
id: collect-261001-huawei/huawei/slideshow-huawei-switch-configuration-commands-29238665-0332bc29-1
title: "slideshow-huawei-switch-configuration-commands-29238665-0332bc29"
domain: huawei
role: reference
task: reference
actors: ["Google", "Huawei"]
dates: []
keywords: ["agentic", "agents", "cost", "incident", "mcp"]
source: docs/RAG/collect-261001-huawei/slideshow-huawei-switch-configuration-commands-29238665-0332bc29.md
source_anchor: ""
source_lines: [1, 49]
sha256: 607bb99f2571dc0493b8cae820eef3857a91ccd394636b4b2f522ed2501c8833
---

# slideshow-huawei-switch-configuration-commands-29238665-0332bc29

Téléchargé 1 191 fois
Publicité
Publicité
Publicité
 Ouvre dans une nouvelle fenêtre Ouvre un site Web externe Ouvre un site Web externe dans une nouvelle fenêtre 
      Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant :     Politique de confidentialité    
   Préférences en matière de conservation des données
Passer au contenu principal
Téléchargé parHuanetwork
DOC, PDF60 182 vues
The document provides a comprehensive guide on Huawei router and switch configuration commands, detailing various commands for managing IP addresses, routing, VLANs, and user privileges. It covers both static and dynamic routing commands, access control list configurations, and NAT settings, enabling users to effectively set up and manage their networking equipment. Additional information and resources for Huawei networking products are also mentioned, including contact details for inquiries.
PDF
Huawei Switch S5700  How To - Configuring single-tag vlan mapping
parIPMAX s.r.l.
10 diapositives28K vues
DOC
Huanetwork x dsl solution - huawei adsl2+ and vdsl2 solution)
parHuanetwork
12 diapositives972 vues
What are the differences between huawei and cisco wlan products
parHuanetwork
12 diapositives1.7K vues
PPTX
Optimizing Write-Intensive Database Workloads Masterclass: Why Scaling Writes...
parScyllaDB
42 diapositives11 vues
When AI Agents Work Together: Exploring A2A, MCP, and Connected AI Frameworks...
24 diapositives10 vues
Comprehensive Guide to Vulnerability Remediation and Incident Response Workflow
8 diapositives29 vues
Comprehensive Guide to Building and Deploying Applications on Google Cloud Pl...
33 diapositives64 vues
Agile Gurugram & Delhi National Capital Region 2026 | AISDLC - Why, What and ...
parAgileNetwork
26 diapositives15 vues
Centralized Application-Context Aware Firewall with AI/ML for Enhanced Cybers...
parndaindna55
7 diapositives40 vues
Optimizing Write-Intensive Database Workloads Masterclass: Strategies for Opt...
parScyllaDB
29 diapositives10 vues
Salesforce Headless 360: Revolutionizing Agentic Enterprise with AI and Multi...
parDele Amefo
17 diapositives16 vues
Innovative Applications of Nanotechnology in Fisheries and Aquaculture Value ...
parB. BHASKAR
21 diapositives19 vues
- 1. Huawei switch configurationcommands Huawei router switch configuration commands: computer command PCAlogin:root; root users The password:linux password is Linux; #shutdown-hnow shutdown; #init0 shutdown; #logout; user logoff #login; user login #ifconfig displays the IP address; #ifconfigeth0netmask; set the IP address #ifconfigeht0netmaskdown; disable IP address #routeadd0.0.0.0gw; set the gateway #routedel0.0.0.0gw delete gateway; #routeadddefaultgw; set the gateway #routedeldefaultgw delete gateway; #route display gateway; #ping; send ECHO packets #telnet; remote login Huawei router switch configuration commands: switch command [Quidway]discur displays the current configuration; [Quidway]displaycurrent-configuration displays the current configuration; [Quidway]displayinterfaces display interface information; [Quidway]displayvlanall displays the routing information; [Quidway]displayversion; display version information [Quidway]superpassword; modify privileged user password [Quidway]sysname switch name; [Quidway]interfaceethernet0/1; into the interface view [Quidway]interfacevlanx; into the interface view [Quidway-Vlan-interfacex]ipaddress10.65.1.1255.255.0.0; VLAN IP address configuration [Quidway]iproute-static0.0.0.00.0.0.010.65.1.2; static routing = gateway [Quidway]rip; three layer exchange support [Quidway]local-userftp [Quidway]user-interfacevty04; enter virtual terminal [S3026-ui-vty0-4]authentication-modepassword; set the password mode [S3026-ui-vty0-4]setauthentication-modepasswordsimple222; set the password [S3026-ui-vty0-4]userprivilegelevel3; user level [Quidway]interfaceethernet0/1 entered the port mode; [Quidway]inte0/1 entered the port mode; [Quidway-Ethernet0/1]duplex{half|full|auto}; working state port configuration [Quidway-Ethernet0/1]speed{10|100|auto} configure the port work rate; [Quidway-Ethernet0/1]flow-control; configuration port flow control 1
- 2. [Quidway-Ethernet0/1]mdi{across|auto|normal}; configuration portflush twisting [Quidway-Ethernet0/1]portlink-type{trunk|access|hybrid}; set the port mode [Quidway-Ethernet0/1]portaccessvlan3; add the port to VLAN [Quidway-Ethernet0/2]porttrunkpermitvlan{ID|All}; trunk VLAN allowed [Quidway-Ethernet0/3]porttrunkpvidvlan3; set the trunk port of PVID [Quidway-Ethernet0/1]undoshutdown activation port; [Quidway-Ethernet0/1]shutdown; closed port [Quidway-Ethernet0/1]quit; return Create a VLAN [Quidway]vlan3; [Quidway-vlan3]portethernet0/1; add port in VLAN [Quidway-vlan3]porte0/1 shorthand method; [Quidway-vlan3]portethernet0/1toethernet0/4; add port in VLAN [Quidway-vlan3]porte0/1toe0/4 shorthand method; [Quidway]monitor-port; designated port mirroring [Quidway]portmirror; assigned port mirroring [Quidway]portmirrorint_listobserving-portint_typeint_num; specify a mirror and mirror [Quidway]descriptionstring specifies the VLAN description character; [Quidway]description; deletion of VLAN description character Check the VLAN settings to [Quidway]displayvlan[vlan_id]; [Quidway]stp{enable|disable}; set the spanning tree, is turned off by default [Quidway]stppriority4096; set spantree priority [Quidway]stproot{primary|secondary}; set as the root or root backup [Quidway-Ethernet0/1]stpcost200; set the switch port cost [Quidway]link-aggregatione0/1toe0/4ingress|both port polymerization; [Quidway]undolink-aggregatione0/1|all; start port channel number [SwitchA-vlanx]isolate-user-vlanenable; set master VLAN [SwitchA]isolate-user-vlansecondary; set the main VLAN including VLAN [Quidway-Ethernet0/2]porthybridpvidvlan; VLAN PVID Remove VLAN PVID [Quidway-Ethernet0/2]porthybridpvid; [Quidway-Ethernet0/2]porthybridvlanvlan_id_listuntagged; set no identification VLAN If the package of vlanid is consistent with PVId, then remove the default PVID=1 VLAN information. So the PVID is set for the vlanid, set VLAN untagged. Exchange Huawei router switch configuration commands: router command [Quidway]displayversion; display version information [Quidway]displaycurrent-configuration displays the current configuration; [Quidway]displayinterfaces display interface information; [Quidway]displayiproute displays the routing information; [Quidway]sysnameaabbcc; change the host name [Quidway]superpasswrod123456; set the password [Quidway]interfaceserial0 into the interface; [Quidway-serial0]ipaddress configure the port IP address; 2
