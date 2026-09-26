---
id: collect-260926-mikrotik/mikrotik/questions-554129-mikrotik-router-rdp-brute-force-blocking-with-whitelist-ccf78954
title: "questions-554129-mikrotik-router-rdp-brute-force-blocking-with-whitelist-ccf78954"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-554129-mikrotik-router-rdp-brute-force-blocking-with-whitelist-ccf78954.md
source_anchor: ""
source_lines: [1, 22]
sha256: 106da0e9c312e148e3090a9ac4f6bbf35335df61e717854b78355cbe1c96ca38
---

# questions-554129-mikrotik-router-rdp-brute-force-blocking-with-whitelist-ccf78954

Masters,
We have a miktorik router, and we would like to block RDP brute force, but we need whitelist also to allow our collegaues to connect without blocking them. (they sometimes hitting wrong password..)
So as an earlier post and @Regan suggested, we have block RDP attacks with this rules:
add chain=forward protocol=tcp dst-port=3389 src-address-list=rdp_blacklist action=drop \
comment="drop rdp brute forcers" disabled=no
add chain=forward protocol=tcp dst-port=3389 connection-state=new \
src-address-list=rdp_stage3 action=add-src-to-address-list address-list=rdp_blacklist \
address-list-timeout=10d comment="" disabled=no
add chain=forward protocol=tcp dst-port=3389 connection-state=new \
src-address-list=rdp_stage2 action=add-src-to-address-list address-list=rdp_stage3 \
address-list-timeout=1m comment="" disabled=no
add chain=forward protocol=tcp dst-port=3389 connection-state=new src-address-list=rdp_stage1 \
action=add-src-to-address-list address-list=rdp_stage2 address-list-timeout=1m comment="" disabled=no
add chain=forward protocol=tcp dst-port=3389 connection-state=new action=add-src-to-address-list \
address-list=rdp_stage1 address-list-timeout=1m comment="" disabled=no
But in this way, this rules often take our collegaue to RDP Blacklist (if hit the password, or login different computer with same IP..).
So i need to modify this code, or add new rule, if the RDP_White_List contain an IP than always allow connection, and not put to RDP_Black_list
So i need to create an rdp_whitelist...
add chain=forward dst-port=3389 src-address-list=rdp_whitelist action=accept
It's ok. But what is the correct rule order? Because this accept rule is not wokking...
Thank you for suggestion...
This is my rule set now:
