---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-vip-dnat-port-forwarding-not-working-ngfw-policy-ta-p-5f87d2e3
title: "t5-fortigate-technical-tip-vip-dnat-port-forwarding-not-working-ngfw-policy-ta-p-5f87d2e3"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-vip-dnat-port-forwarding-not-working-ngfw-policy-ta-p-5f87d2e3.md
source_anchor: ""
source_lines: [1, 4]
sha256: 11071c569df2aab7c23a3106eb31ef3962b2acdf936d521cee521b607115d57d
---

# t5-fortigate-technical-tip-vip-dnat-port-forwarding-not-working-ngfw-policy-ta-p-5f87d2e3

Technical Tip: VIP/DNAT port forwarding not working (NGFW Policy based central NAT)
| Description | This article describes the scenario for VIP port forwarding in an NGFW policy-based central NAT setup. | 
| Scope | FortiGate Central NAT. | 
| Solution | In the scenario of 2 DNATs are configured. One DNAT with port forwarding but the other DNAT without.   If the DNAT without port forwarding is on top, then it will not match the port forwarding VIP.  Debug flow:  id=20085 trace_id=31 func=ipv4_fast_cb line=53 msg="enter fast path" id=20085 trace_id=31 func=ip_session_run_all_tuple line=7140 msg="DNAT x.x.3.26:33389->y.y.3.23:33389"  The firewall will not do the port mapping from 33389 to 3389.  It is necessary to move the VIP object with port forwarding to the top from GUI.  Or from CLI:  # config firewall vip move test2 before test1 end   With the above changes, FortiGate will match test2 prior to test 1.  Debug flow:  id=20085 trace_id=36 func=ipv4_fast_cb line=53 msg="enter fast path" id=20085 trace_id=36 func=ip_session_run_all_tuple line=7140 msg="DNAT x.x.3.26:33389->y.y.3.23:3389"  Now the port is forwarding from 33389 to 3389. |
