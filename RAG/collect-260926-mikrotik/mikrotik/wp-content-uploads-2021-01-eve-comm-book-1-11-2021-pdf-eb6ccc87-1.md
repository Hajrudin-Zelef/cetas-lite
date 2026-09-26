---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-1
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [1, 91]
sha256: 55682d46d1191efb1769fb9c03efbbfbeffeca650c1848f792c44ff372e41164
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community 
Cookbook 
 
 
Version 1.11  
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
© EVE-NG LTD 
The information contained in this document is the property of EVE-NG Limited 
The contents of the document must not be reproduced or disclosed wholly or in part or used 
for purposes other than that for which it is supplied without the prior written permission of 
EVE-NG Limited. 
Author:  
Uldis Dzerkals 
 
Editors:  
Michael Doe 
Christopher Lim

EVE-NG Community Cookbook 
Version 1.11 
Page 2 of 165 © EVE-NG LTD 
 
Table of Contents 
 
PREFACE .................................................................................................................................. 7 
1 INTRODUCTION ................................................................................................................ 8 
1.1 WHAT IS EVE-NG? ....................................................................................................... 8 
1.2 WHAT IS EVE-NG USED FOR? ....................................................................................... 8 
1.3 WHO IS EVE-NG FOR? ................................................................................................. 8 
2 SYSTEM REQUIREMENTS............................................................................................... 9 
2.1 HARDWARE REQUIREMENTS ........................................................................................... 9 
2.1.1 Minimal Laptop/PC Desktop system requirements .............................................. 9 
2.1.2 Recommended Laptop/PC Desktop system requirements ................................ 10 
2.1.3 Virtual Server system requirements ................................................................... 10 
2.1.4 Dedicated Server (bare) system requirements .................................................. 11 
2.1.5 Nodes per lab calculator .................................................................................... 11 
2.2 SUPPORTED VIRTUALIZATION PLATFORMS AND SOFTWARE ............................................. 11 
2.3 UNSUPPORTED HARDWARE AND SYSTEMS .................................................................... 12 
3 INSTALLATION ............................................................................................................... 13 
3.1 VMWARE WORKSTATION OR VM PLAYER ..................................................................... 13 
3.1.1 VMware workstation EVE VM installation using ISO image (preferred)  ............ 13 
3.1.1.1 EVE VM Setup and Settings ....................................................................................................... 13 
3.1.1.2 EVE-NG VM Installation steps ................................................................................................... 16 
3.1.2 VMware workstation OVF deployment .............................................................. 22 
3.1.2.1 Deployment and VM machine settings ..................................................................................... 22 
3.1.2.2 OVF VM update to the latest EVE version ................................................................................. 24 
3.1.2.3 OVF VM HDD Size expansion ..................................................................................................... 24 
3.2 VMWARE ESXI ........................................................................................................... 24 
3.2.1 VMware ESXi EVE installation using ISO image (preferred)............................. 24 
3.2.1.1 EVE-NG ESXi VM Setup and Settings ......................................................................................... 24 
3.2.1.2 EVE-NG ESXi VM Installation steps ............................................................................................ 26 
3.2.2 VMware ESXi OVF deployment ......................................................................... 32 
3.2.2.1 ESXi OVF VM Setup and Settings ............................................................................................... 32 
3.2.2.2 ESXi OVF VM update to the latest EVE version ......................................................................... 34 
3.2.2.3 ESXi OVF VM HDD Size expansion ............................................................................................. 34 
3.3 BARE HARDWARE SERVER EVE INSTALLATION .............................................................. 34 
3.3.1 Ubuntu Server Installation Phase 1 ................................................................... 34 
3.3.2 EVE Community Installation Phase 2 ................................................................ 42 
3.3.3 EVE Community Installation Phase 3 ................................................................ 42 
3.4 GOOGLE CLOUD PLATFORM ......................................................................................... 43 
3.4.1 Google account .................................................................................................. 43 
3.4.2 Goggle Cloud project ......................................................................................... 43 
3.4.3 Preparing Ubuntu boot disk template ................................................................ 45 
3.4.4 Creating VM ....................................................................................................... 46 
3.4.5 EVE-NG-Community installation ........................................................................ 48 
3.4.6 Access to Google Cloud EVE-COMM ............................................................... 50 
3.4.7 Optional: GCP Firewall rules for native console use ......................................... 50 
3.5 EVE MANAGEMENT IP ADDRESS SETUP ....................................................................... 53 
3.5.1 Management static IP address setup (preferred) .............................................. 53

