---
id: collect-261001-general-networking/general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431-3
title: "wp-content-uploads-2026-06-preview-v1-pdf-24e88431"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431.md
source_anchor: ""
source_lines: [97, 179]
sha256: 64f46fa2e3a958473b5918ba48e6173cd75750291b05c08b3da6ec13334e850b
---

# wp-content-uploads-2026-06-preview-v1-pdf-24e88431

3.1.4 IP Sec .................................................................................................................. Error! Bookmark not defined. 
3.1.5 GRE ................................................................................................................. Error! Bookmark not defined. 
3.2 Remote Access VPN ............................................................................................... Error! Bookmark not defined. 
3.2.1 IPsec VPN ........................................................................................................... Error! Bookmark not defined. 
3.2.2 SSL VPN ............................................................................................................. Error! Bookmark not defined. 
4 HA LAB ............................................................................................................................... Error! Bookmark not defined. 
4.1 Topology .................................................................................................................. Error! Bookmark not defined. 
4.2 IP addressing .......................................................................................................... Error! Bookmark not defined. 
4.3 LAB Setup ............................................................................................................... Error! Bookmark not defined. 
4.4 Verification ............................................................................................................... Error! Bookmark not defined. 
5 SD-WAN LAB ..................................................................................................................... Error! Bookmark not defined. 
5.1 Topology .................................................................................................................. Error! Bookmark not defined. 
5.2 IP Addressing .......................................................................................................... Error! Bookmark not defined. 
5.3 LAB Setup ............................................................................................................... Error! Bookmark not defined. 
5.3.1 DNS-Root ............................................................................................................ Error! Bookmark not defined. 
5.3.2 DNS-Google ........................................................................................................ Error! Bookmark not defined. 
5.3.3 DNS-CloudFlare .................................................................................................. Error! Bookmark not defined. 
5.3.4 ISP1 ..................................................................................................................... Error! Bookmark not defined. 
5.3.5 ISP2 ..................................................................................................................... Error! Bookmark not defined. 
5.3.6 Verification Commands on FortiGate: ................................................................. Error! Bookmark not defined. 
5.3.7 FortiGate BGP Configuration .............................................................................. Error! Bookmark not defined. 
5.4 SD WAN Zone ......................................................................................................... Error! Bookmark not defined. 
5.5 Performance SLA .................................................................................................... Error! Bookmark not defined. 
5.6 SD-WAN Rules ........................................................................................................ Error! Bookmark not defined. 
5.7 verification ............................................................................................................... Error! Bookmark not defined. 
6 VDOM .................................................................................................................................. Error! Bookmark not defined. 
6.1 Enable VDOM mode ............................................................................................... Error! Bookmark not defined. 
6.2 Verify VDOM Status ................................................................................................ Error! Bookmark not defined. 
6.3 Create New VDOMs ................................................................................................ Error! Bookmark not defined.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 1-4 
 
6.4 Assign Interfaces to VDOMs ................................................................................... Error! Bookmark not defined. 
6.5 Configure Interfaces Within a VDOM ...................................................................... Error! Bookmark not defined. 
6.6 Create Inter-VDOM Links (VDOM Links) ................................................................ Error! Bookmark not defined. 
6.7 Configure Firewall Policies in Specific VDOM ........................................................ Error! Bookmark not defined. 
6.8 View Current VDOM Context .................................................................................. Error! Bookmark not defined. 
6.9 Disable VDOM Mode (FortiOS 7.4) ......................................................................... Error! Bookmark not defined. 
7 Logging and Monitoring ................................................................................................... Error! Bookmark not defined. 
7.1 Traffic Logs .............................................................................................................. Error! Bookmark not defined. 
7.2 Log & Report ........................................................................................................... Error! Bookmark not defined. 
7.2.1 Log Settings ........................................................................................................ Error! Bookmark not defined. 
7.2.2 Forward Traffic .................................................................................................... Error! Bookmark not defined. 
7.2.3 Security Events ................................................................................................... Error! Bookmark not defined. 
7.2.4 System Events .................................................................................................... Error! Bookmark not defined. 
7.3 FortiView ................................................................................................................. Error! Bookmark not defined.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 1-5 
 
 
1.1 Files and Documents 
The table below shows the files and materials that come with this workbook 
File or Folder Name Comments 
FortiGateAdministrator7.4.pdf This current file that you are reading 
FortiGate-LAB1.unl Should be imported into eVe-NG 
FORTIGATE-HA  
FortiGate-LAB-SDwan  
FortiGate-LAB-IPsec-VPN  
FortiGate (directory) 
This folder contains a virtioa.qcow2 file. 
Should already be placed on your eVe-NG server at 
/opt/unetlab/addons/qemu/ 
 
1. ensure that these files have been imported into the respective directories on your eVe-NG server as indicated 
in the table above. 
 x86_64_crb_linux-adventerprisek9-ms.bin 
 asav-922-1/virtioa.qcow2 
2. Access the Lab in eVe-NG 
 Log in to your eVe-NG web interface. 
 Navigate to “File manager”. 
 Select the “Import” icon. 
 Choose the  unl then hit open. 
 Once the file is uploaded, locate the lab in the GUI and cli ck it to open the lab. 
 The lab topology and configurations will load automatically. 
1.2 Overview 
This document provides a comprehensive, hands-on guide to the most common FortiGate features, structured 
across four progressive laboratories: FortiGate LAB, VPN Lab, HA Lab, and SD-WAN Lab, covering concepts 
from basic setup to advanced enterprise features.

