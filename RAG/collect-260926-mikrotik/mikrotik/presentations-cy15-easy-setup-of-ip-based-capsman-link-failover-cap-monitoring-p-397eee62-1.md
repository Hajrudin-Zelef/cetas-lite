---
id: collect-260926-mikrotik/mikrotik/presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62-1
title: "presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62"
domain: mikrotik
role: reference
task: reference
actors: ["Apple", "EU"]
dates: ["2015-06"]
keywords: ["license", "training", "voice"]
source: docs/RAG/lot-mikrotik/RouterOS/presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62.md
source_anchor: ""
source_lines: [1, 292]
sha256: ffa00eea489d0bdbacd05d8f89c8b3b460f55e2e1b32ea553ebb5d9a6797947f
---

# presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62

Easy Setup of IP Based CAPsMAN 
with link failover & CAPs monitor 
Georgios Argyrides 
MUM Middle East  (Cyprus-Larnaca) 
12nd June 2015

About Me 
 
My Name: 
Georgios Argyrides 
 You can call me “George” 
(its easier) 
 
 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 2

About Me 
• Born in Cyprus  
(Europe, Near Greece)  
– Can Speak English & Greek 
 
 
• Have been working in Industry since 2006 
 ITSP Consultant / Voice Engineer 
 Systems / Network Administrator 
 Internet Security Consultant 
 ISP / WISP Consultant 
 
 Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 3

About Me 
• 1st MikroTik Certified Consultant  
in Greece since 2011 
[MTCRE,MTCWE,MTCTCE,MTCINE]  
 
• MikroTik Certified Trainer  
in Greece since 2012 
 
• Cyberoam Certified Network & Security Professional (CCNSP) 
• BSc (Hon) Applied Computing , Sheffield Hallam University(UK) 
 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 4

About Me 
• Providing  MikroTik Training (On-Site)  
& Consultancy (On-site or remote) 
 in Greece & Cyprus(EU Region)  
 Worldwide as well 
 
• I have conducted the training events just  
before MUM Middle East 2015 
 Thanks to MyTelco Limited(Cyprus)   
for hosting  and helping me organize: 
– 6-8 June 2015 1st Official MTCNA Training in Larnaca,Cyprus  
– 9-11 June 2015 1st Official MTCWE Training in Larnaca,Cyprus 
 
• More information and contact details at the end of this Presentation  
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 5

This Presentation Objective 
• Introduce the IP Based CAPsMAN  even for  new users of 
MikroTik products 
– Through  an Easy SetUP 
 
• Small Routed Network  
– CAPs communicating with CAPsMAN by redundant connections 
(routing protocol will take care of this) 
 
• Monitor CAPs 
– Get notified when one goes down 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 6

CAPsMAN Features 
• Centralized management of RouterOS APs  
• Dual Band AP support 
• Provisioning of APs 
• MAC and IP Layer communication with APs 
• Certificate support for AP communication  
• Full and Local data forwarding mode 
• RADIUS MAC authentication 
• Custom configuration support 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 7

Definitions  
?CAP?? CAPs? CAPsMAN? AP? Router?  
CAPsMAN 
• Controlled Access Point 
system Manager 
CAP 
• Controlled Access Point 
 CAPsMAN = a MikroTik router  CAP = a MikroTik router 
 CAPs = many Mikrotik routers 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 8

Requirements 
CAPsMAN 
1. x86 or RouterBOARD 
based device 
2. RouterOS v6.11+ 
version (Use Latest!) 
3. Wireless-fp package  
installed and enabled 
 
CAPs 
1. X86 or RouterBOARD 
based device  
2. RouterOS v6.11+ 
version(Use Latest!) 
3. Atheros chipset 
(a/b/g/n/ac) wireless card  
4. Wireless-fp package 
installed and enabled  
5. At least Level4 RouterOS 
license 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 9

CAPsMAN v1 & v2(New) 
 CAPsMAN v.6.23+ introduces CAPsMAN  v2 
• Improvements 
• Some new features 
 
 CAPsMAN v1 is already stable and can be used for production 
 
 Warning: CAPsMAN/CAP v1 is not compatible with  v2! 
 Upgrade or downgrade everything in the network 
 
  Try CAPsMAN v2 initially on non-production environment 
 Help us make it better by reporting any possible  issues appeared in v2 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 10

CAPsMAN Simple Setup 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 11

CAPsMAN Simple Setup 
• Enable CAPsMAN service 
• Create Bridge interface 
• Add IP configuration to Bridge interface 
• Create CAPsMAN Configuration  
• Create Provisioning rule 
• Enable CAP mode on the APs 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 12

CAPsMAN Simple Setup 
• Enable the CAPsMAN service 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 13

CAPsMAN Simple Setup 
• Create Bridge Interface 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 14

CAPsMAN Simple Setup 
1. Add IP 
address 
 
2. Add DHCP 
Server 
 
3. Add NAT 
rule 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 15

CAPsMAN Simple Setup 
• Add new CAPsMAN Configuration 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 16

CAP to CAPsMAN IP Based Connection 
IP (UDP) Layer3 
 CAP communicates CAPsMAN 
using IP protocol 
 Can traverse NAT when 
required 
 Management connection 
between CAP and CAPsMAN is 
secured using DTLS 
 CAP client data traffic is not 
secured  
 If encryption is 
required IPSec  or 
encrypted tunnels can 
be used  
Specify IP on The CAP 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 17

CAPsMAN and CAP in one board 
• Does your CAPsMAN router has a wireless interface too? 
 Enable CAP & Connect it to it self (127.0.0.1) for central 
management 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 18

CAPsMAN Simple Setup 
• Add new Provisioning rule 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 19

CAPsMAN Simple Setup 
• Check the “Interface” status on:  
CAPsMAN CAP 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 20

CAPsMAN Registration table 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 21

Manual Provisioning 
• Changing Provisioning rules doesn't effect already 
configured CAPs, manual Provisioning required: 
 Remove CAP interface 
 Initiate Provision command on the CAP 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 22

CAP Identification 
• MAC/IP address 
• RouterBoard model 
• Serial Number of the Board 
• RouterOS version 
• System Identity 
• Main wireless MAC 
• State of the CAP 
• Radio count 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 23

CAPsMAN static CAP interface 
• Interface name or setting does not change after a reboot  
• Additional manual setting override 
• Copy dynamic interface to make static interface 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 24

CAPsMAN Virtual AP 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 25

CAPsMAN VirtualAP Configuration 
• Create new Bridge interface and IP 
configuration for the VirtualAPs  
 Or use the same bridge interface  
used for  Master AP 
• Create a new configuration for the VirtualAP 
• Specify the new configuration in Provisioning 
rule as Slave Configuration 
• Remove all CAP interfaces 
• Initiate Manual Provisioning on all the CAPs 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 26

CAPsMAN VirtualAP Setup 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 27

CAPsMAN VirtualAP Setup 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 28

CAPsMAN static VirtualAP 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 29

CAPsMAN Access List Features 
• MAC Authentication  
• Radius Query support  
• MAC Mask support  
• Signal Range 
• Time 
• Private Passphrase  
• VLAN ID assignment 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 30

CAPsMAN Access List 
• Allow Apple devices to connect 
• Let RADIUS server decide for the rest of devices 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 31

CAPsMAN Configuration override 
 
 
 
• Configuration 
overrides Channel 
setting 
• Interface overrides 
Channel and 
Configuration setting 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 32

CAPsMAN Auto Certificate 
• Enable Certificate and CA Certificate on CAPsMAN 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 33

CAPsMAN Auto Certificate 
• Enable “Request” Certificate on CAP 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 34

