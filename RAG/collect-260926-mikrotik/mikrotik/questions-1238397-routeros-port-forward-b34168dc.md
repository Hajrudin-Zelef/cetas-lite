---
id: collect-260926-mikrotik/mikrotik/questions-1238397-routeros-port-forward-b34168dc
title: "questions-1238397-routeros-port-forward-b34168dc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1238397-routeros-port-forward-b34168dc.md
source_anchor: ""
source_lines: [1, 15]
sha256: 71e71535107f0602382e133d3a2d45fc28d6a5078f8cab470b356dcf33de2ec2
---

# questions-1238397-routeros-port-forward-b34168dc

I want to port-forward HTTP traffic from 1.1.66.166:80 to LAN-ip 10.13.37.10:80 on my router running RouterOS 6.40
Configuration:
/ip address print
Flags: X - disabled, I - invalid, D - dynamic 
 #   ADDRESS            NETWORK         INTERFACE                                                                                                                                                                                                          
 0   10.13.37.62/26     10.13.37.0      ether2                                                                                                                                                                                                             
 1   1.1.66.166/27      1.1.66.160      ether1 
/ip firewall nat print
Flags: X - disabled, I - invalid, D - dynamic 
 0    chain=dstnat action=dst-nat to-addresses=10.13.37.10 protocol=tcp dst-address=1.1.66.166 dst-port=80 log=yes log-prefix="" 
 1    chain=srcnat action=masquerade out-interface=ether1 log=no
But no traffic is reaching the internal host. Am I missing something here?
I have verified that I can reach 10.13.37.10:80 from the router.
I dó get logging that the rule has been triggered:
firewall,info dstnat: in:ether1 out:(none), src-mac xx:xx:d0:xx:3c:00, proto TCP (SYN), xx.xx.174.107:22880->1.1.66.166:80, len 60
