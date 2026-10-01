---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a-2
title: "fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a.md
source_anchor: ""
source_lines: [148, 206]
sha256: 7f39d2be2780bdf91e8f898575076a5b014d273f92f34dc76e8a8ebc9aca881f
---

# fortigate-3-troubleshooting-tip-central-nat-packet-processing-troubleshooting-st-57cfd10a

Depending on the specific condition, such as an incorrect policy configuration or a routing issue, different error messages may be encountered. By analyzing these messages, it is possible to identify the root cause of the problem.


**SNAT issues:****No route to destination:**

- **New session:**


`func=fw_forward_handler line=839 msg="Denied by forward policy check (policy 0)"`
 

- **Existing session after route change:**

`func=fw_forward_dirty_handler line=443 msg="state=00000208, state2=00000000, npu_state=40000100"                                                                             func=fw_post_route_handler line=1158 msg="Session is in BLOCK state. Drop the packet."`

**No matching policy:**

```
func=fw_forward_handler line=839 msg="Denied by forward policy check (policy 0)"
```

**Central-NAT policy problem:**


```
func=__iprope_check_one_policy line=2374 msg="policy-0 is matched, act-accept"
func=fw_snat_check line=688 msg="NAT disabled by central SNAT policy!"
```

**DNAT issues:**

 

**No matching DNAT policy:**

`func=iprope_dnat_check line=5506 msg="result: skb_flags-02000000, vid-0, ret-no-match, act-accept, flag-00000000"`
 

**Host behind NAT unreachable:**


The debug flow is normal, but diagnose sniffer packet shows:

```
551.958499 port3 in 192.168.2.10.60761 -> 192.168.5.10.80: syn 641016312
554.826495 port3 out 192.168.10.1 -> 192.168.2.10: icmp: host 192.168.5.10 unreachable
```

In some cases, it shows ARP resolution failure:

`17.511812 port2 out arp who-has 192.168.1.10 tell 192.168.1.1`

**Key troubleshooting:**

- Always check if the policy uses the correct IP (post-NAT for DNAT, pre-NAT for SNAT).
- Routing issues cause packets to be dropped before NAT is applied.
- VRF configuration must match interface assignments for routes to work.
- policy-0 drops indicate policy mismatch or routing failure.
- Use both CLI debug and GUI logs for comprehensive visibility
