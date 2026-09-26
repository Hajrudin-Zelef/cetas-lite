---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-6
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["cost", "distribution", "ethernet", "intel", "packaging", "training"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [334, 507]
sha256: ed507f6db50aedc8aef7e438e426de427d1bd98be3821d6947280bde217f6467
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 8 of 165 © EVE-NG LTD 
1 Introduction 
1.1 What is EVE-NG? 
To describe what Emulated Virtual Environment – Next Generation (EVE-NG) is without solely 
stating dry facts about features, we need to elaborate more on what EVE -NG can be used for 
and whom it would be useful for. 
In some trivial dry words, EVE -NG gives you tools to use around virtual devices and 
interconnect them with other virtual or physical devices. Many of its features greatly simplify the 
usabilities, re-usability, manageability, interconnectivity, distribution and therefore the ability to 
understand and share topologies, work, ideas, concepts or simply “labs”. This can simply mean 
it will reduce the cost and time to s et up what you need or it might enable you to do tasks you 
would not have thought could be done this simple. 
1.2 What is EVE-NG used for? 
This is the real question but there is no finite answer, the possibilities are almost limitless and 
depends on what you want to use it for. 
It can be used for studying all kinds of technologies. You can learn about general technologies 
or vendor specific topics. You can test new technologies like network automation, SDN, etc. 
It can be used to recreate corporate networks and test changes before putting them i nto 
production. You can create proof of concepts for clients . You can troubleshoot network issues 
by recreating them and e.g. use Wireshark to inspect packets. 
It is most definitely not just for networking, it can be used to test software in simulated networks, 
test out security vulnerabilities of any kind, system engineering like LDAP and AD servers and 
many more areas. 
You could set it up to automate sandboxing unknown files/software and use software to analyse 
short and long term behaviour for malicious intent much simpler than without EVE -NG. 
The list of what EVE -NG can be used for co uld go on indefinitely, possibilities are limited by 
knowledge and imagination only. Both of which can be improved with EVE-NG. 
To get a ve ry small idea of what can be done with EVE -NG, check out the tested/ supported 
images (many have not been tested, almost everything virtual should run on EVE-NG) and refer 
to section 12. 
EVE-NG helps you achieve what you want to and more. 
1.3 Who is EVE-NG for? 
EVE-NG is for everyone working in the Information Technology Sector, period. 
It is for very large enterprise companies, training facilities, service providers, consultants, 
people who want to train themselves; it is for everyone, it is for YOU! 
Use-cases that are more than worth it, almost priceless even, can be found everywhere.  
 
The EVE-NG community version is free for everyone; while the paid professional version adds 
a few things that make your life easier. Almost everything can still be done with the free version, 
just less conveniently and therefore more time-consuming. 
However, with the free version, the possibility to train yourself with  technologies, hone your 
skills and become an expert even with very no monetary possibilities. For some this is and has 
been life changing.

EVE-NG Community Cookbook 
Version 1.11 
Page 9 of 165 © EVE-NG LTD 
2 System requirements 
EVE-NG is available in the OVF or ISO file format. The Open Virtualization Format (OVF) is an 
open standard for packaging and distributing virtual appliances. It can be used to deploy a VM 
in hypervisors like VMware Workstation, Player and ESXi. Please note that installing EVE as a 
Virtual Machine (VM) will mean any nodes deployed within EVE will be nested. Nested 
virtualization causes degraded performance in deployed nodes . This should be fine for  lab 
purposes as long as the host meets or exceed s the resource requirements for the deployed 
nodes.  
EVE-NG can also be installed directly on physical hard ware, without a hypervisor, using the 
provided ISO image. This is referred to as a “bare metal” install and is the most recommended 
method of installing EVE-NG. 
2.1 Hardware requirements 
2.1.1 Minimal Laptop/PC Desktop system requirements 
Prerequisites: 
CPU: Intel CPU supporting Intel® VT-x /EPT virtualization 
Operating System: Windows 7, 8, 10 or Linux Desktop 
VMware Workstation 12.5 or later 
VMware Player 12.5 or later 
 
PC/Laptop HW requirements 
CPU Intel i5/i7 (4 Logical processors), Enabled Intel virtualization 
in BIOS 
RAM 8Gb 
HDD Space 40Gb 
Network LAN/WLAN 
EVE Virtual machine requirements 
CPU 4/1 (Number of processors/Number of cores per processor) 
Enabled Intel VT-x/EPT virtualization engine 
 
RAM 6Gb or more 
HDD 40Gb or more 
Network VMware NAT or Bridged network adapter 
Note: Minimal PC Desktop/Laptop will be able to run small Labs. The performance and quantity 
of nodes per lab depend on the types of nodes deployed in the lab. 
Example:  
IOL image-based nodes: up to 40-50 nodes per lab 
Dynamips image-based nodes: up to 20-25 nodes per lab 
vIOS image-based nodes: up to 8-10 nodes per lab 
CSRv1000 or XRv image-based nodes: up to 2-3 per lab

EVE-NG Community Cookbook 
Version 1.11 
Page 10 of 165 © EVE-NG LTD 
 
2.1.2 Recommended Laptop/PC Desktop system requirements 
Prerequisites: 
CPU: Intel CPU supporting Intel® VT-x /EPT virtualization  
Operation System: Windows 7, 8, 10 or Linux Desktop 
VMware Workstation 12.5 or later 
VW Ware Player 12.5 or later 
 
PC/Laptop HW requirements 
CPU Intel i7 (8 Logical processors), Enabled Intel virtualization in 
BIOS 
RAM 32Gb 
HDD Space 200Gb 
Network LAN/WLAN 
EVE Virtual machine requirements 
CPU 8/1 (Number of processors/Number of cores per processor) 
Enabled Intel VT-x/EPT virtualization engine 
 
RAM 24Gb or more 
HDD 200Gb or more 
Network VMware NAT or Bridged network adapter 
 
Note: PC Desktops/Laptops will be able to run small to medium Labs. Performance and quantity 
of nodes per lab depend on the type of nodes deployed in the lab. 
Example:  
IOL image-based nodes: up to 120 nodes per lab 
vIOS image-based nodes: up to 20-40 nodes per lab 
CSR image-based nodes: up to 10 per lab 
2.1.3 Virtual Server system requirements 
Prerequisites: 
CPU: Intel Xeon CPU supporting Intel® VT-x with Extended Page Tables (EPT) 
Operation System: ESXi 6.0 or later 
 
Server HW requirements 
CPU Recommended CPU 2x Intel E5-2650v3 (40 Logical processors) 
or better supporting Intel® VT -x with Extended Page Tables 
(EPT) 
Minimum CPU is any Intel Xeon CPU supporting Intel® VT -x 
with Extended Page Tables (EPT) 
RAM 128Gb 
HDD Space 2Tb 
Network LAN Ethernet

EVE-NG Community Cookbook 
Version 1.11 
Page 11 of 165 © EVE-NG LTD 
EVE Virtual machine requirements 
CPU 32/1 (Number of processors/Number of cores per processor) 
Enabled Intel VT-x/EPT virtualization engine 
 
RAM 64Gb or more 
HDD 800Gb or more 
Network vSwitch/VMnet 
Note: Performance and quantity of nodes per lab depends from the type of nodes used in the 
lab.  
Example:  
120 IOL image-based lab 
20 CSRv1000 image-based nodes per lab 
2.1.4 Dedicated Server (bare) system requirements 
Prerequisites: 
CPU: Intel Xeon CPU supporting Intel® VT-x with Extended Page Tables (EPT) 
Operation System: Ubuntu Server 16.04.4 LTS x64 
 
Server HW requirements 
CPU Recommended CPU Intel E5-2650v3 (40 Logical processors) or 
better supporting Intel® VT-x with Extended Page Tables (EPT) 
Minimum CPU is any Intel Xeon CPU supporting Intel® VT -x 
with Extended Page Tables (EPT) 
RAM 128Gb 
HDD Space 2Tb 
Network LAN Ethernet 
 
Note: Performance and quantity of nodes per lab depends from type of nodes used in the lab.  
2.1.5 Nodes per lab calculator 
It is recommended to use the “nodes per lab calculator” to achieve best performance and avoid 
overloading your EVE system. 
https://drive.google.com/file/d/1Rbu7KDNSNuWiv_AphWx0vCek8CKVB1WI/view  
2.2 Supported virtualization platforms and software 
• VMware Workstation 12.5 or later 
• VMware Player 12.5 or later 
• VMware ESXi 6.0 or later

