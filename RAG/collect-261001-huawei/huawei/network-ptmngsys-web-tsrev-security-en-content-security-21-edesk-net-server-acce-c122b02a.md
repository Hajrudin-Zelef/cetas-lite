---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-security-en-content-security-21-edesk-net-server-acce-c122b02a
title: "network-ptmngsys-web-tsrev-security-en-content-security-21-edesk-net-server-acce-c122b02a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-security-en-content-security-21-edesk-net-server-acce-c122b02a.md
source_anchor: ""
source_lines: [1, 52]
sha256: 91415659f87429373a0f90d04f7afff57b539b9192321e93a0be63d2a64e3766
---

# network-ptmngsys-web-tsrev-security-en-content-security-21-edesk-net-server-acce-c122b02a

<HUAWEI> display ip interface brief 
*down: administratively down 
(s): spoofing 
Interface            IP  Address  Physical Protocol Description   
GigabitEthernet1/0/1  1.1.1.1    up      up 
GigabitEthernet1/0/2  10.2.0.1   up      up 
GigabitEthernet1/0/3  10.3.0.1   up      up 
GigabitEthernet1/0/4  unassigned  down  down 
GigabitEthernet1/0/5  unassigned  down  down   
GigabitEthernet1/0/6  unassigned  down  down 
GigabitEthernet1/0/7  unassigned  down  down
Check the information in the IP Address column. If any configuration is incorrect, run the ip address ip-address mask command in the interface view to re-configure the IP address.
<HUAWEI> display zone 
local 
priority is 100 
# 
trust 
 priority is 85 
 interface of the zone is (1): 
  GigabitEthernet1/0/3 
# 
untrust 
 priority is 5 
 interface of the zone is (1): 
  GigabitEthernet1/0/1 
# 
dmz 
 priority is 50 
 interface of the zone is (1): 
  GigabitEthernet1/0/2 
# 
For services using the multichannel protocol such as FTP and H.323, check whether the ASPF function is configured.
Run the display firewall detect command to check whether the ASPF function is enabled for the multichannel protocols globally or between zones.
<HUAWEI> display firewall detect   
global configruation 
--------------------------------  
detect ftp  
detect rtsp  
detect h323   
zone untrust 
--------------------------------  
detect h323   
interzone trust untrust 
--------------------------------  
detect ftp  
detect pptp  
detect ils
If not enabled, run the firewall detect command in the system view to enable the ASPF function.
<HUAWEI> system 
[HUAWEI] firewall detect ftp 
[HUAWEI] firewall detect h323 
[HUAWEI] firewall detect sip
