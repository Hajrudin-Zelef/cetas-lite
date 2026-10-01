---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-fortigate-ha-failover-via-fortimanager-94190-ba6096cc
title: "fortigate-3-technical-tip-fortigate-ha-failover-via-fortimanager-94190-ba6096cc"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-fortigate-ha-failover-via-fortimanager-94190-ba6096cc.md
source_anchor: ""
source_lines: [1, 13]
sha256: 26a46e078a95421ebcfcea0c3c7a0bac223d20f8135a7b7a1e59e2fbae38dd0a
---

# fortigate-3-technical-tip-fortigate-ha-failover-via-fortimanager-94190-ba6096cc

Technical Tip: FortiGate HA failover via FortiManager
Description
This article describes how to perform FortiGate HA failover via the FortiManager 'promote' option.
Scope
FortiGate.
Solution
To select the 'promote' option from  GUI:
- Go to Device Manager, and select the 'HA device' on the left panel, which will open System: Dashboard.
- To make the secondary becomes the primary, select 'Promote', then select 'OK' to confirm failover. 
- It will take a while to change the 'HA device' sequence.
- Once completed, the HA role will be changed.
- To confirm this on the production HA FortiGate, go to System -> HA.
Related article:
