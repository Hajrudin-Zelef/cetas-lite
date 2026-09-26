---
id: collect-260926-mikrotik/mikrotik/questions-71078-static-routing-host-to-host-ping-does-not-work-mikrotik-2ee6f757
title: "questions-71078-static-routing-host-to-host-ping-does-not-work-mikrotik-2ee6f757"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-71078-static-routing-host-to-host-ping-does-not-work-mikrotik-2ee6f757.md
source_anchor: ""
source_lines: [1, 7]
sha256: ad74a18884517cce24e99ddcffa04728f145c162ce559d182d489d98e07c98fb
---

# questions-71078-static-routing-host-to-host-ping-does-not-work-mikrotik-2ee6f757

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
Static route configuration in both R1 & R2 router configuration is correct . According to your configuration should be able to establish connectivity among both laptop s without any issues.
check whether gateway is reachable from laptops , verify it by pinging gateway from laptop .
check ARP table in router whether Mac address of laptop is visible with ip address of laptop
Sh ip arp | i 192.168.88.252
Check connectivity from router whether 192.168.77.254 is reachable from router1 ,Verify it with pinging from router1.
If it's real time please verify firewall status on laptop end . If it's windows laptop please ensure windows firewall is off.
