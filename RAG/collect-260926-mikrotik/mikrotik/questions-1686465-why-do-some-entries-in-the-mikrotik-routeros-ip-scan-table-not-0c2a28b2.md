---
id: collect-260926-mikrotik/mikrotik/questions-1686465-why-do-some-entries-in-the-mikrotik-routeros-ip-scan-table-not-0c2a28b2
title: "questions-1686465-why-do-some-entries-in-the-mikrotik-routeros-ip-scan-table-not-0c2a28b2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1686465-why-do-some-entries-in-the-mikrotik-routeros-ip-scan-table-not-0c2a28b2.md
source_anchor: ""
source_lines: [1, 40]
sha256: 7638bfa3b574dc6440e99b3533f28607ce13dcd7a8341641887013cbe4a1458a
---

# questions-1686465-why-do-some-entries-in-the-mikrotik-routeros-ip-scan-table-not-0c2a28b2

There is a command /tool ip-scan in the RouterOS, which basically does network scanning. It works quite efficiently, however, there is something I can't understand in its behaviour.
For some network the answer settled to something like this (MAC addresses are obfuscated):
> tool ip-scan address-range=192.168.26.0/24
Flags: D - dhcp 
  ADDRESS         MAC-ADDRESS        TIME DNS      SNMP      NETBIOS                                   
  192.168.26.2    FE:02:4F:2C:47:06                                                                    
  192.168.26.9    DC:13:75:57:7A:2A 186ms                                                              
  192.168.26.254                    254ms                                                              
  192.168.26.28   98:35:1C:62:D0:6B 189ms                                                              
  192.168.26.12   76:46:E1:C2:03:5B 278ms                                                              
  192.168.26.11   54:57:83:05:36:A0 419ms                                                              
  192.168.26.31   32:68:DA:E4:69:EB 202ms                                                              
  192.168.26.19   10:79:8B:99:9C:DF 293ms                                                              
  192.168.26.27   FE:8A:7A:B1:CF:FF 291ms                                                              
  192.168.26.32   DC:9B:EE:25:F2:C8 279ms                                                              
  192.168.26.33   BA:AC:9F:08:25:FB 205ms                                                              
  192.168.26.60   98:BD:46:B3:58:7C 308ms                                                              
  192.168.26.87   76:CE:15:20:8B:F9 271ms                                                              
  192.168.26.14   54:DF:4B:5C:BE:9B                                                                    
  192.168.26.91                     386ms                                                              
  192.168.26.29   10:F1:C6:2F:14:8F                                                                    
  192.168.26.8    FE:02:03:1D:47:96                                                                    
  192.168.26.20   DC:13:6D:43:7A:A3                                                                    
  192.168.26.100  BA:24:6B:31:AD:D7 328ms                                                              
  192.168.26.101  98:35:3A:B8:D0:7B 577ms                                                              
  192.168.26.99   76:46:53:E8:03:79 578ms                                                              
  192.168.26.16   54:57:F9:60:36:8B                                                                    
  192.168.26.124  32:68:34:F9:69:E8 454ms                                                              
  192.168.26.40   10:79:7D:D4:9C:A0                                                                    
  192.168.26.41   FE:8A:0F:E0:CF:BA                                                                    
  192.168.26.93   DC:9B:53:E1:F2:77                                                                    
  192.168.26.154                    514ms                                                              
  192.168.26.200  98:BD:5D:C5:58:6A 221ms                                                              
  192.168.26.15   76:CE:69:8D:8B:F4                                                                    
  192.168.26.1    54:DF:44:D5:BE:A8                                                                    
  192.168.26.7    32:E0:E8:27:E1:FA                                                                    
-- [Q quit|D dump|C-z continue]
This output was settled, so nothing was changed for some time, like a minute (I paused it to be able to copy).
The RouterOS is 192.168.24.254 in this case. It is directly in the network. There are no "too smart" devices, no proxy-arps, etc., just switches and wireless access points.
Question is: why are there entries without MAC address (except for .254)? How it can see that some IP address is taken without knowing its MAC address e.g. from the ARP reply or other packet it sent? Also, why some entries lack a RTT delay field, while otherwise being successfully detected?
