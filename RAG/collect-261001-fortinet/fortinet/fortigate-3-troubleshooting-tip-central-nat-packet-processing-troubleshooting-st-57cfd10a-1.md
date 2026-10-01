---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a-1
title: "fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a.md
source_anchor: ""
source_lines: [1, 147]
sha256: 3a7e76ff2114ecf5e896ef17e22392b691bcc15ac2d20718c33067ad6af51d86
---

# fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a

**Description**

This article describes how to troubleshoot Central NAT and traffic flow on FortiGate. It explains the packet processing order, routing considerations, firewall policy matching, and NAT behavior, and provides CLI and GUI troubleshooting methods.**Scope**

FortiGate, FortiOS.**Solution**

Central NAT on FortiGate is a powerful feature that enables flexible source and destination NAT policies. It is often used for complex deployments like VIPs, VPNs, and multi-VRF environments. However, troubleshooting Central NAT issues can be challenging due to its deep integration with routing, firewall policies, and UTM policies. In Central NAT mode, NATs and firewall policies are 2 independent modules.

This article provides a step-by-step troubleshooting guide focusing on:

- Packet flow and processing order.
- Routing considerations.
- Firewall policy and NAT interactions.
- Debugging commands and log analysis.


**Understanding packet flow in Central NAT:**

When a packet arrives at a FortiGate interface, it follows a specific sequence. This process ensures that the packet is handled correctly based on the configured rules and policies.

- **Ingress Interface:** Packet enters ingress interface and virtual domain (VDOM).
- **Session Lookup** : FortiGate checks for an existing session. If found, the packet is processed and forwarded based on the session's established rules, bypassing some subsequent checks.
- **DNAT Check:** Applies if there is a matching Central NAT DNAT or VIP policy; the destination IP rewritten
- **Routing Lookup:** Based on the post-DNAT destination IP, FortiGate determines the outgoing interface and VRF.
- **Firewall Policy Lookup:** Matches source and (post/Pre-NAT) destination IP, ports, zones, and interfaces.
- **SNAT Check:** If enabled, the source IP is translated after policy match.
- **Packet Forwarding:** The packet is sent via the determined egress interface.


Example traffic flow and steps in red:


```
trace_id=76 func=print_pkt_detail line=6005 msg="vd-root:0 received a packet(proto=6, 192.168.1.10:49928->192.168.2.10:80) tun_id=0.0.0.0 from port2. flag [S], seq 4102116648, ack 0, win 64240"
trace_id=76 func=init_ip_session_common line=6204 msg="allocate a new session-00014233"
trace_id=76 func=iprope_dnat_check line=5481 msg="in-[port2], out-[]"
trace_id=76 func=iprope_dnat_tree_check line=824 msg="len=0"
trace_id=76 func=iprope_dnat_check line=5506 msg="result: skb_flags-02000000, vid-0, ret-no-match, act-accept, flag-00000000"
trace_id=76 func=__vf_ip_route_input_rcu line=1989 msg="find a route: flag=00000000 gw-192.168.10.2 via port3"
trace_id=76 func=__iprope_fwd_check line=810 msg="in-[port2], out-[port3], skb_flags-02000000, vid-0, app_id: 0, url_cat_id: 0"
trace_id=76 func=__iprope_tree_check line=524 msg="gnum-100004, use int hash, slot=37, len=2"
trace_id=76 func=__iprope_check_one_policy line=2140 msg="checked gnum-100004 policy-1, ret-matched, act-accept"
trace_id=76 func=__iprope_user_identity_check line=1903 msg="ret-matched"
trace_id=76 func=__iprope_check line=2404 msg="gnum-4e20, check-ffffffffa002cb97"
trace_id=76 func=__iprope_check_one_policy line=2140 msg="checked gnum-4e20 policy-6, ret-no-match, act-accept"
trace_id=76 func=__iprope_check_one_policy line=2140 msg="checked gnum-4e20 policy-6, ret-no-match, act-accept"
trace_id=76 func=__iprope_check_one_policy line=2140 msg="checked gnum-4e20 policy-6, ret-no-match, act-accept"
trace_id=76 func=__iprope_check line=2421 msg="gnum-4e20 check result: ret-no-match, act-accept, flag-00000000, flag2-00000000"
trace_id=76 func=__iprope_check_one_policy line=2374 msg="policy-1 is matched, act-accept"
trace_id=76 func=__iprope_fwd_check line=847 msg="after iprope_captive_check(): is_captive-0, ret-matched, act-accept, idx-1"
trace_id=76 func=iprope_fwd_auth_check line=876 msg="after iprope_captive_check(): is_captive-0, ret-matched, act-accept, idx-1"
trace_id=76 func=iprope_reverse_dnat_check line=1353 msg="in-[port2], out-[port3], skb_flags-02000000, vid-0"
trace_id=76 func=iprope_reverse_dnat_tree_check line=916 msg="len=0"
trace_id=76 func=iprope_central_nat_check line=1376 msg="in-[port2], out-[port3], skb_flags-02000000, vid-0"
trace_id=76 func=__iprope_check_one_policy line=2140 msg="checked gnum-10000d policy-1, ret-matched, act-accept"
trace_id=76 func=get_new_addr line=1274 msg="find DNAT: IP-192.168.0.10, port-49928"
trace_id=76 func=__iprope_check_one_policy line=2374 msg="policy-1 is matched, act-accept"
trace_id=76 func=fw_forward_handler line=1002 msg="Allowed by Policy-1: SNAT"
trace_id=76 func=ip_session_confirm_final line=3179 msg="npu_state=0x100, hook=4"
trace_id=76 func=__ip_session_run_tuple line=3512 msg="SNAT 192.168.1.10->192.168.0.10:49928"
```
**Note:**

- **SNAT:** Use the pre-NAT source IP in the policy. SNAT applies after policy check.
- **DNAT:** Use the post-NAT destination IP in the policy. DNAT applies before policy check


**When using the DNAT mapped IP (VIP) in the policy:**

- DNAT translates the destination IP before policy check.
- Policy lookup tries to match the post-DNAT (real) IP.

If the policy still uses the VIP, it will not match, and traffic will hit the implicit deny (policy-0).

**DNAT on outgoing traffic:**

In some rare cases, DNAT applied on outgoing (egress) interface. The logic remains the same:

- The destination IP is rewritten before routing and policy lookup.
- The firewall policy must use the translated destination (Post-NAT) IP for matching.

**Note:** 

Missing routes or policy for the translated IP cause policy drop issues.**DNAT with pre-NAT IP not on incoming interface:**

Use Central NAT where the VIP is associated with an interface different from the incoming interface.

For example:

- **Port 2 with IP:** 10.0.2.1/24.
- **Port 3 with IP:** 10.0.3.1/24.


A packet with destination IP 10.0.3.50 come in from the Port 2 and is configured in a VIP with 10.0.3.50 as extip (Pre-NAT IP) and with extintf 'any'.

That VIP 'belongs' to the port , so a firewall policy is required to allow the traffic between the port2 and the port3, additionally to the one allowing the traffic between the port3 and the interface matched by the routing lookup for the Post-NAT IP (mappedip).

If the extip (Pre-NAT IP) dosen't belong to any subnet configured on FortiGate ports it consider beloning to the ingress interface so the double firewall policy is not necessary.

A solution to avoid the double firewall for each new "inter-interface traffic flow" is enabling a firewall policy on top enabling traffic between port2 and port3.**Debug traffic flow and packet capture:**

The diagnose debug flow command is essential for tracing packet flow. Use the commands below to enable and use it for gathering traffic flow.

Reset debug filters and levels:

```
diagnose debug reset
diagnose debug flow filter clear
```

Set filter and parameters:

```
diagnose debug flow filter <parameter> <value>
diagnose debug flow show iprope enable
diagnose debug flow show function-name enable
diagnose debug console timestamp enable
diagnose debug flow trace start <number of packets>
```

View applied filters:

`diagnose debug flow filter`

Enable/disable debugging:

```
diagnose debug enable
diagnose debug disable
```

Sometimes, it is necessary to see the packet's contents before and after the NAT process to troubleshoot. Use the **diagnose sniffer packet** command with a specific filter to capture packets.

`diagnose sniffer packet any 'host < IP> and port <port>' 4`

**GUI tools for troubleshooting:**

Under **Policy & Objects -> Central NAT**, find the rule, right-click, and select 'Show matching logs'.



Under **Log & Report -> Forward Traffic**, find the traffic and check NAT translations and policies.


**Common debug scenarios and error messages:**


