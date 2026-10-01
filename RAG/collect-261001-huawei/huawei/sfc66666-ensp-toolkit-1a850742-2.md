---
id: collect-261001-huawei/huawei/sfc66666-ensp-toolkit-1a850742-2
title: "Clone anywhere"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/sfc66666-ensp-toolkit-1a850742.md
source_anchor: ""
source_lines: [195, 266]
sha256: b3111ba368b95e06f7edbfa859848f055f54a5ac68949023c0e9847bcff52601
---

# Clone anywhere

             " green pass yellow pass red discard")
cfg.add_line("#")
cfg.add_line("traffic policy QOS_POLICY")
cfg.add_line(" classifier guest_night behavior limit_2m")
cfg.add_line("#")
cfg.set_interface("GigabitEthernet1/0/1", traffic_policy_in="QOS_POLICY")
# OSPF to cores
cfg.set_ospf(process_id=1, router_id="10.0.0.1",
    networks=[("10.0.0.0", "0.0.0.3")])
print(cfg.build())
See examples/campus_network.py for a
fully-worked 16-device campus network featuring:
- Dual-core VRRP hot standby (S5700 × 2)
- 3 access zones (机房 / 教室 / 宿舍) with DHCP
- Firewall with security zones, NAT, and time-based QoS
- OSPF internal routing with default-route advertisement
Run it:
python examples/campus_network.py
# Output → ~/Desktop/campus_network/
| Model | Description | Interfaces | 
|---|---|---|
| S3700 | Layer 3 access switch | 22 Ethernet + 2 GE | 
| S5700 | Layer 3 core switch | 24 GE | 
| AR2240 | Enterprise router | 3 GE | 
| USG6000V | Next-gen firewall | 8 GE | 
| Router | Generic router | 2 Eth + 4 GE + 4 Serial | 
| PC | Endpoint (static or DHCP) | 1 Ethernet | 
| Server | Server endpoint | 1 Ethernet | 
| Client | Client endpoint (HTTP) | 1 Ethernet | 
| Cloud | Host bridge | 1 Ethernet | 
| Method | Description | 
|---|---|
| add_vlan(id, name?) | Create a VLAN | 
| add_vlan_batch(ids) | Create multiple VLANs | 
| set_interface(name, **opts) | Configure an interface | 
set_interface options: ip, description, link_type, access_vlan,
trunk_allow, default_vlan, pvid, hybrid_tagged, hybrid_untagged,
undo_shutdown, stp_disable, trust_dscp, traffic_policy_in,
traffic_policy_out, nat_outbound, nat_server, eth_trunk,
qos_queue_profile, dhcp_select_global
| Method | Description | 
|---|---|
| add_static_route(dest, mask, next_hop) | IPv4 static route | 
| set_default_route(next_hop) | Default route (0.0.0.0/0) | 
| set_ospf(pid, router_id, networks, ...) | OSPF process with area 0 | 
| add_ospf_network(net, wc, area) | Add network to OSPF area | 
| Method | Description | 
|---|---|
| set_stp(mode, region_name, instances) | STP / MSTP | 
| add_eth_trunk(id, members, **opts) | Link aggregation | 
| add_vrrp(ifname, vrid, vip, prio, preempt, auth_key) | VRRP group | 
| Method | Description | 
|---|---|
| enable_dhcp_global() | Enable DHCP server globally | 
| add_dhcp_pool(name, net, mask, gw, ...) | Create DHCP pool | 
| add_acl(number, rules) | Numbered ACL | 
| add_acl_named(name, number, rules) | Named ACL | 
| add_nat_outbound(acl) | Source NAT (hide behind interface IP) | 
| add_nat_server(proto, gip, gport, iip, iport) | Destination NAT / port forward | 
| Method | Description | 
|---|---|
| add_line(raw) | Inject raw VRP line (firewall zones, time-range, etc.) | 
| add_traffic_policy(name, cb_pairs) | Traffic classifier + behavior + policy | 
| build() | Render complete VRP config string | 
output/
├── my_network.topo              # Topology XML → open in eNSP
├── {UUID}/                      # Per-device directory
│   └── vrpcfg.zip               # Contains vrpcfg.cfg → import in device console
└── ...
Workflow: Open my_network.topo in eNSP → start devices → right-click each
device → "Import Configuration" → select its vrpcfg.zip.
MIT
