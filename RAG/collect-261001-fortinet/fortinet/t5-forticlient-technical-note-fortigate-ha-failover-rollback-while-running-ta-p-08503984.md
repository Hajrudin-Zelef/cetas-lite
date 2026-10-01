---
id: collect-261001-fortinet/fortinet/t5-forticlient-technical-note-fortigate-ha-failover-rollback-while-running-ta-p-08503984
title: "t5-forticlient-technical-note-fortigate-ha-failover-rollback-while-running-ta-p--08503984"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-forticlient-technical-note-fortigate-ha-failover-rollback-while-running-ta-p--08503984.md
source_anchor: ""
source_lines: [1, 52]
sha256: a5633aeeb7bd2deab0a39abd879e068c38a98b466becdbe5f3bbe64d668dd02c
---

# t5-forticlient-technical-note-fortigate-ha-failover-rollback-while-running-ta-p--08503984

Technical Note: FortiGate HA failover - Rollback while running IPsec traffic from FortiClient
Description
FortiGate tuning proposals to support cluster failover and rollback while running traffic in IPsec tunnel from/to FortiClient.
 
Scope
FortiOS 5.2.10
 FortiClient 5.6.0
 Both IPsec setting using IKEv1
 
Solution
FortiGate HA commands
  config system ha
  
Solution #1
Modify the FortiGate to propose a single phase-2 Diffie-Hellman group. Use group 5 instead of default value proposing group 14 and group 5.
  fgt (phase2-interface) # config vpn ipsec phase2-interface 
  
Solution #2
Modify Phase-2 replay detection value to 'DISABLE' on both sides.
On the FortiGate:
  fgt (phase2-interface) # config vpn ipsec phase2-interface 
  
On FortiClient:
Edit the IPSec VPN connection
Click on "Advanced Setting" > "Phase-2" >
Remove "Enable Replay Detection"
 config system ha
set mode a-p
set hbdev <portname> 50 <portname> 50
set session-pickup enable
set session-pickup-connectionless enable
set ha-mgmt-status enable
set ha-mgmt-interface <port>"
set ha-mgmt-interface-gateway <ip addr>
set override disable
set priority 250
Solution #1
Modify the FortiGate to propose a single phase-2 Diffie-Hellman group. Use group 5 instead of default value proposing group 14 and group 5.
fgt (phase2-interface) # config vpn ipsec phase2-interface 
edit "client_tunnel"
set phase1name " client_tunnel "
set dhgrp 5
Solution #2
Modify Phase-2 replay detection value to 'DISABLE' on both sides.
On the FortiGate:
fgt (phase2-interface) # config vpn ipsec phase2-interface 
edit " client_tunnel "
set replay disable
On FortiClient:
Edit the IPSec VPN connection
Click on "Advanced Setting" > "Phase-2" >
Remove "Enable Replay Detection"
