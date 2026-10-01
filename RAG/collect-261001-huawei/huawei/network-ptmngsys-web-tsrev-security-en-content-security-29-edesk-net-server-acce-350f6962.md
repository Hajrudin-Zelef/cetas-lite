---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-security-en-content-security-29-edesk-net-server-acce-350f6962
title: "network-ptmngsys-web-tsrev-security-en-content-security-29-edesk-net-server-acce-350f6962"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-security-en-content-security-29-edesk-net-server-acce-350f6962.md
source_anchor: ""
source_lines: [1, 68]
sha256: 6551c0046253868fc35d21338fea1c4f2ddb0f9c6784853235f00af838126c03
---

# network-ptmngsys-web-tsrev-security-en-content-security-29-edesk-net-server-acce-350f6962

Run the display ip interface brief command to check whether IP addresses are correctly configured for interfaces.
<HUAWEI> display ip interface brief 
*down: administratively down                                                                                                         
!down: FIB overload down                                                                                                             
^down: standby                                                                                                                       
(l): loopback                                                                                                                        
(s): spoofing                                                                                                                        
(d): Dampening Suppressed                                                                                                            
(E): E-Trunk down                                                                                                                    
The number of interface that is UP in Physical is 3                                                                                 
The number of interface that is DOWN in Physical is 3                                                                              
The number of interface that is UP in Protocol is 4                                                                                  
The number of interface that is DOWN in Protocol is 4                                                                              
 
Interface                         IP Address/Mask      Physical   Protocol  
GigabitEthernet1/0/1              1.1.1.1              up         up 
GigabitEthernet1/0/2              10.2.0.1             up         up 
GigabitEthernet1/0/3              10.3.0.1             up         up 
GigabitEthernet1/0/4              unassigned           down       down 
GigabitEthernet1/0/5              unassigned           down       down  
GigabitEthernet1/0/6              unassigned           down       down  
GigabitEthernet1/0/7              unassigned           down       down
Check the IP Address column. If any configuration is incorrect, run the ip address ip-address mask command in the interface view to re-configure an IP address.
<HUAWEI> display zone 
local 
priority is 100 
interface of the zone is (0):     
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
If any interface is added to an inappropriate zone, run the add interface interface-type interface-number command in the desired security zone view to add the interface to the zone.
For services using multi-channel protocols such as FTP and H.323, check whether the ASPF function is configured for the protocols.
Run the display firewall detect command to check whether the ASPF function is enabled for multi-channel protocols globally or in interzones.
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
If the ASPF function is disabled, run the firewall detect command in the system view to enable it.
<HUAWEI> system 
[HUAWEI]firewall detect ftp 
[HUAWEI]firewall detect h323 
[HUAWEI]firewall detect sip
Currently, the USG6000 series supports the following multi-channel protocols: ACTIVEX BLOCKING, DNS, FTP, H.323, ICQ, ILS, JAVA BLOCKING, MMS, MSN, NetBIOS, PPTP, QQ, RTSP, RSH, SCCP, SIP, and SQL.NET.
