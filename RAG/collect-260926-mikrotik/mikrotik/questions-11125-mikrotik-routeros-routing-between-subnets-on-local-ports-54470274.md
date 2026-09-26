---
id: collect-260926-mikrotik/mikrotik/questions-11125-mikrotik-routeros-routing-between-subnets-on-local-ports-54470274
title: "questions-11125-mikrotik-routeros-routing-between-subnets-on-local-ports-54470274"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-11125-mikrotik-routeros-routing-between-subnets-on-local-ports-54470274.md
source_anchor: ""
source_lines: [1, 17]
sha256: ef2adb4bb0969d9d340e0528e192e89ba2503489c1b0715894d71643b672fe37
---

# questions-11125-mikrotik-routeros-routing-between-subnets-on-local-ports-54470274

The fault was with the DHCP server setup on the Mikrotik. Before:
[xxxx@MikroTik] /ip dhcp-server> network print 
 # ADDRESS            GATEWAY         DNS-SERVER      WINS-SERVER     DOMAIN                 
 0 192.168.100.0/32   192.168.100.1   8.8.8.8                                       
 1 192.168.102.0/32   192.168.102.1   8.8.8.8                                        
 2 192.168.105.0/32   192.168.105.1   8.8.8.8                                       
 3 192.168.200.0/32   192.168.200.1   8.8.8.8   
After:
[xxxx@MikroTik] /ip dhcp-server> network print
 # ADDRESS            GATEWAY         DNS-SERVER      WINS-SERVER     DOMAIN
 0 192.168.100.0/24   192.168.100.1   8.8.8.8                                           
 1 192.168.102.0/24   192.168.102.1   8.8.8.8                                       
 2 192.168.105.0/24   192.168.105.1   8.8.8.8                                            
 3 192.168.200.0/24   192.168.200.1   8.8.8.8
The difference is the subnet mask (/32 -> /24).
Even though the Mikrotik has an option to add the Netmask as /24 on the DHCP Network screen in Winbox, it does not automatically pull that mask value through to the address value and thus needs to be explicitly added to the address value as well.
(I know this export example is not exactly like my question, but I did not want to over-complicate the question)
