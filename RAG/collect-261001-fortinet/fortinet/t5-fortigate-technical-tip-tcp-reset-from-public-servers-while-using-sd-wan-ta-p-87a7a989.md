---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-ta-p-87a7a989
title: "t5-fortigate-technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-ta-p-87a7a989"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-ta-p-87a7a989.md
source_anchor: ""
source_lines: [1, 5]
sha256: 42246c65d996be681f2fbbdf25b470ad3c2c09245f2afebb59bdacf930458272
---

# t5-fortigate-technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-ta-p-87a7a989

| Description | This article describes a possible scenario where the user is applying SD-WAN configuration with 3 ISP links. | 
| Scope | FortiGate. | 
| Solution | However, the user is seeing in logs multiple TCP resets from public servers on the internet while traffic is being allowed by the proper SD-WAN rule 3 which has the below settings :  **config system sdwan**    **config service** edit 3 set name "test" set addr-mode ipv4 set input-device-negate disable set mode load-balance set minimum-sla-meet-members 0 set hash-mode round-robin set role standalone set standalone-action disable set tos 0x00 set tos-mask 0x00 set protocol 0 set route-tag 0 set dst "all" set dst-negate disable set src "Name" set src-negate disable set internet-service disable set dscp-forward disable set dscp-reverse disable                 **config sla** edit "SLA" set id 1 next end set priority-members 3 4 5 set status enable set default disable set passive-measurement disable next end  As shown above, the SD-WAN rule has a round-robin hash-mode which may result in public servers receiving the request from different source IPs and eventually will lead to TCP reset.  Change the SD-WAN rule hash mode to be source-ip-based as shown below:  **config system sdwan**    **config service** edit 3 set hash-mode source-ip-based next end | 

The Fortinet Security Fabric brings together the concepts of convergence and consolidation to provide comprehensive cybersecurity protection for all users, devices, and applications and across all network edges.
