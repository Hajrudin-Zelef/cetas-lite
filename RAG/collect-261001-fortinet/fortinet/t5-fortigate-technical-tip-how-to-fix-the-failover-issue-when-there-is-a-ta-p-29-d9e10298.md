---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-how-to-fix-the-failover-issue-when-there-is-a-ta-p-29-d9e10298
title: "t5-fortigate-technical-tip-how-to-fix-the-failover-issue-when-there-is-a-ta-p-29-d9e10298"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-how-to-fix-the-failover-issue-when-there-is-a-ta-p-29-d9e10298.md
source_anchor: ""
source_lines: [1, 35]
sha256: bd2d8a110c544d79876fcceec35e56694d1a96d7b4743d7ddc2cf5c4c403cf58
---

# t5-fortigate-technical-tip-how-to-fix-the-failover-issue-when-there-is-a-ta-p-29-d9e10298

Technical Tip: How to fix the failover issue when there is a FortiGate-VM HA cluster with multiple IPSec tunnels
Description
This article describes how to fix the failover issue when there is a FortiGate-VM HA cluster with multiple IPSec tunnels.
There is an Active-Passive HA cluster. There is more than one Public address assigned to the External interface of the FGT-VM. Each IP is used in a different IPSec tunnel, so it is necessary to specify a 'local-gw' in each phase1-interface configuration.
Since the local-gw is the same on both FGT-A and FGT-B but the WAN interface private IPs are different, if the fail-over happens, the new primary will not be able to bring up the IPsec tunnel.
This may happen on the HA clusters that do not have a floating private IP moving between the HA members.
Scope
FortiGate-VM, OCI-FortiGate, FortiGate v7.0.x, v7.2.x, v7.4.x.
Solution
To make both HA members able to bring up the IPsec tunnel it is necessary to, first, prevent the phase1-interface configuration sync between the HA members and, then, change the local-gw on the FGT-B according to its WAN interface IPs.
Process:
- Prevent the phase1-interface configuration sync between the HA members:
config system vdom-exception
    edit 0
        set object vpn.ipsec.phase1-interface
    next
end
-  Change the local-gw on the FGT-B according to its WAN interface IPs: config vpn ipsec phase1-interface 
 edit VPN-1
 set local-gw <primary private IP>
 next
 edit VPN-2
 set local-gw <secondary private IP 1>
 next
 edit VPN-3
 set local-gw <secondary private IP 2>
 next
end
After following these steps, trigger the failover and test the IPSec tunnels.
Related documents:
Phase 1 configuration | FortiGate / FortiOS 7.4.2 | Fortinet Document Library
VDOM exceptions | FortiGate / FortiOS 7.4.2 | Fortinet Document Library
IPsec VPN in an HA environment | FortiGate / FortiOS 6.2.15 | Fortinet Document Library
Related article:
Technical Tip: IPSec VPN in HA Environment - Fortinet Community
