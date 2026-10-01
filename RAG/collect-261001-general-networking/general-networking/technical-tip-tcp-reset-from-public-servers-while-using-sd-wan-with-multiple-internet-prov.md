---
id: collect-261001-general-networking/general-networking/technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-with-multiple-internet-prov
title: "technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-with-multiple-internet-providers-comm"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-with-multiple-internet-providers-comm.md
source_anchor: ""
source_lines: [1, 5]
sha256: f3cc9fbef7ce118a163f439345eea5b4f571b5fef5a9020bf9008a1c47633cb6
---

# technical-tip-tcp-reset-from-public-servers-while-using-sd-wan-with-multiple-internet-providers-comm

| Description | This article describes a possible scenario where the user is applying SD-WAN configuration with 3 ISP links. | 
| Scope | FortiGate. | 
| Solution | However, the user is seeing in logs multiple TCP resets from public servers on the internet while traffic is being allowed by the proper SD-WAN rule 3 which has the below settings :  **config system sdwan**    **config service** edit 3 set name "test" set addr-mode ipv4 set input-device-negate disable set mode load-balance set minimum-sla-meet-members 0 set hash-mode round-robin set role standalone set standalone-action disable set tos 0x00 set tos-mask 0x00 set protocol 0 set route-tag 0 set dst "all" set dst-negate disable set src "Name" set src-negate disable set internet-service disable set dscp-forward disable set dscp-reverse disable                 **config sla** edit "SLA" set id 1 next end set priority-members 3 4 5 set status enable set default disable set passive-measurement disable next end  As shown above, the SD-WAN rule has a round-robin hash-mode which may result in public servers receiving the request from different source IPs and eventually will lead to TCP reset.  Change the SD-WAN rule hash mode to be source-ip-based as shown below:  **config system sdwan**    **config service** edit 3 set hash-mode source-ip-based next end | 

Enter your E-mail address. We'll send you an e-mail with instructions to reset your password.
