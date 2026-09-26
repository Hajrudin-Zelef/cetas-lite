---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-1ardkqg-mikrotik-rosv7-dual-dhcp-wan-recursive-failover-w-263bdd0b
title: "r-mikrotik-comments-1ardkqg-mikrotik-rosv7-dual-dhcp-wan-recursive-failover-w-263bdd0b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["chatgpt"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-1ardkqg-mikrotik-rosv7-dual-dhcp-wan-recursive-failover-w-263bdd0b.md
source_anchor: ""
source_lines: [1, 16]
sha256: 1de4769108aa8e5249f373138fcf808393720b7b9f6e378cd71363e2e6072166
---

# r-mikrotik-comments-1ardkqg-mikrotik-rosv7-dual-dhcp-wan-recursive-failover-w-263bdd0b

MikroTik ROSv7 dual DHCP WAN recursive failover w/ PCC load-balancing question
Hi
I ran across this setup on github.com
Mikrotik DHCP WAN recursive failover w/pcc load-balancing
I have already set recursive failover up on my Mikrotik device, so I can make sense most of the code.
However, since I still know little about mangle, routing and policy routing, I am not sure what exactly a few lines of the code are about and do, from line 50 to 91 in particular.
Would anyone so kind to explain them to me? Thanks
Section des commentaires
Run it through chatGPT and ask for a breakdown of how it works. Seriously!
I need to give it a try soon or later. Meanwhile if someone can help me with it is more than welcome. Thanks
Here's my currently working solution (I posted this quite a long time ago):
https://www.reddit.com/r/mikrotik/comments/110nzso/ama_rb5009_initial_success_with_dual_isp_ziply/
Thanks. I might come in handy, even though I am focusing on how it works and what exactly does this kind of setup step by step. A step by step explanations of the code would be appreciated. Thanks
In particular, I didn't understand why we need "recursive routes for ECMP default gateways, dst-address are public DNS servers" and "recursive routes for default gateways, dst-address are public DNS servers" on the same setup."
It’s easy. If you ping your next jump to determine if your gateway is alive, you don’t have a real mechanism to check if your internet connection is working, so you have to ping a server outside your isp to check if it’s working
I know that. But if you take a look at the conf file you'll see that there is the gateway check is on both the "ECMP" and "Recursive Routes". This seems redundant to me. Thanks
