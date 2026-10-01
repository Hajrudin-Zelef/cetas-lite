---
id: collect-261001-automatisation-infra/automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead-3
title: "etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead.md
source_anchor: ""
source_lines: [200, 311]
sha256: 604a781092a7717f4596d7412205e50ed65d5204b115fb959705e36e3783f46a
---

# etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead

- IPsec protection: IKEv2 + PSK over the mGRE tunnels. Later, replace PSK with certificates (Alpine VM as CA if available).
- MikroTik cross-check: site-to-site IPsec between two CHRs — IKEv2, AES-256-GCM, SHA-256, DH group 19. No DMVPN — just the IPsec part. Compare the conceptual IKE phase flow with your Cisco notes.
Prove it:
- show ip nhrp on a spoke shows dynamic entries for other spokes (Phase 2/3).
- show crypto ikev2 sa detailed shows active IKEv2 SAs.
- Packet capture on hub's WAN interface: spoke-to-spoke traffic absent after Phase 3 shortcut established.
GNS3vault labs (gns3vault-archive/Security/) — the archive is strong here:
| Lab | What to do | 
|---|---|
| aaa-authentication | RADIUS + TACACS+ auth, local fallback. | 
| aaa-command-authorization | Command authorization — two privilege tiers. | 
| aaa-exec-authorization | EXEC authorization. | 
| standard-access-list | ACL refresher. | 
| extended-access-list | Extended ACL. | 
| reflexive-access-list | Reflexive ACL — stateful inspection precursor. | 
| unicast-reverse-path-forwarding-urpf | uRPF strict vs loose mode. | 
| control-place-policing | CoPP — MQC for the control plane. | 
| basic-zone-based-firewall | ZBF — zone pairs, policy-map inspect, self zone. | 
IPv6 FHS extension (hand-written): GNS3vault has no IPv6 FHS lab. Build on vIOS-L2 access segment:
Tasks:
- RA Guard: ipv6 nd raguard policy HOST /device-role host on access ports. Send a rogue RA from a host port — blocked. Confirm withshow ipv6 snooping events .
- DHCPv6 Guard: ipv6 dhcp guard policy SERVER /device-role server only on the uplink to the DHCPv6 server. Rogue DHCPv6 offer from a host port — dropped.
- IPv6 Snooping: ipv6 snooping policy SNOOP applied to the access VLAN. Verify binding table withshow ipv6 neighbors binding .
- IPv6 Source Guard: ipv6 source-guard policy SGP on host ports. Host with spoofed IPv6 source — dropped.
CoPP extension: after the GNS3vault CoPP lab, build a production-grade 4-class policy on a CSR1000v:
class MANAGEMENT  → SSH, SNMP, NTP, TACACS+ — police 64 kbps
class ROUTING     → OSPF, BGP, EIGRP, LDP hello — police 256 kbps
class UNDESIRABLE → ICMP redirects, IP options, fragments — police 8 kbps
class default     → everything else — police 1 Mbps
Verify: send a flood of malformed packets at the management interface — show policy-map control-plane shows drops incrementing in the right class.
GNS3vault labs (gns3vault-archive/Network Management/) — good coverage of the classical stack:
| Lab | What to do | 
|---|---|
| snmpv2-server + snmpv3-server | Both back-to-back. Compare config complexity. | 
| syslog-server-logging + system-message-logging | Remote + local logging, severity levels. | 
| ntp-network-time-protocol | NTP server + client, stratum, authentication. | 
| ip-service-level-agreement-sla | IP SLA probes. | 
| kron-task-scheduler | Scheduled config tasks — production-relevant. | 
Modern telemetry extensions (hand-written): GNS3vault is too old for these. Build on a CSR1000v + Docker VM:
- SNMPv3 → LibreNMS: point LibreNMS (running in Docker) at your CSR1000v. Add all routers. Verify interface graphs appear.
- NetFlow v9 → ntopng: flow exporter on the CSR, export to ntopng on port 2055. Add Flexible NetFlow with a custom record (src/dst IP, ports, TOS, BGP next-hop). View top-talkers per AS.
- IP SLA + track + failover: UDP-jitter probe between two routers. track 1 ip sla 1 reachability . Floating static conditional on track state. Pull the primary link — verify static route switches within the SLA reaction time.
- DHCP relay across VRFs: DHCPv4 server in VRF MGMT , clients in VRFCUST_A .ip helper-address vrf MGMT X.X.X.X global on the SVI. Confirm leases:show ip dhcp binding .
- gNMI subscription (stretch): install gnmic on the Docker VM. Subscribe to/interfaces/interface/state/counters on the CSR1000v. Pipe output to Prometheus. Build a Grafana panel showing real-time interface counters.
Prove it:
- LibreNMS shows traffic graphs from every device.
- ntopng shows top-talkers per AS from NetFlow.
- Pull primary link → IP SLA tracker flaps → show ip route shows standby static activated.
- Grafana panel shows live counter increments from gNMI.
No new EVE-NG topology. One structured activity:
Walk the Library C design modules with a notebook. For each module, draw the reference architecture on paper. Then, for every Cisco-proprietary concept, write its vendor-neutral / open-source equivalent:
| Cisco concept | Vendor-neutral equivalent | 
|---|---|
| SD-Access fabric edge | VTEP (EVPN-VXLAN) | 
| LISP control-plane node | BGP-EVPN route-reflector | 
| DNA Center | OpenConfig + Prometheus + NetBox | 
| vSmart (SD-WAN) | Open-source SD-WAN controller (flexiWAN, VyOS-based) | 
| DMVPN NHS | WireGuard hub, strongSwan IKEv2 server | 
| MP-BGP VPNv4 | Same — it's an IETF standard, all vendors implement it | 
This translation exercise is more valuable than any CLI lab this week. Senior engineers think in concepts, not vendor menus.
No video. The entire week is building a working network-automation Git repo.
Repo structure:
network-automation/
  inventory/
    hosts.yaml          # Nornir host inventory
    groups.yaml         # cisco_iosxe, mikrotik_routeros, juniper_junos
  playbooks/
    backup.yaml         # Ansible
    gather_state.py     # Nornir + NAPALM
  scripts/
    netconf_pull.py     # ncclient — pull OSPF neighbors
    restconf_pull.py    # requests — same data via REST
    gnmic_subscribe.sh  # gnmic → Prometheus
  telemetry/
    prometheus.yml
    grafana_dashboards/
  Makefile              # make backup, make report, make telemetry
  README.md
Day-by-day build order:
| Day | Task | 
|---|---|
| 1 | python3 -m venv .venv , install:nornir nornir-napalm nornir-utils nornir-netmiko ansible netmiko napalm ncclient requests pyang gnmic . Firstgit commit . | 
| 2 | make backup — Nornir+NAPALM pulls running configs from every device. Commits toconfigs/ with timestamp. | 
| 3 | make report — Nornir+NAPALM pulls interface error counters. Outputs a Markdown table. | 
| 4 | netconf_pull.py — pull/ietf-routing:routing/control-plane-protocols from CSR1000v via NETCONF. Pretty-print OSPF neighbor list. | 
| 5 | restconf_pull.py — same data via RESTCONF + JSON. Write a paragraph in your notes comparing the two API ergonomics. | 
| 6 | gnmic_subscribe.sh — subscribe to interface counters. Stream to Prometheus. Grafana dashboard. | 
| 7 | Ansible role: deploy a new VRF + OSPF process to a Cisco IOS-XE device and a MikroTik CHR from the same playbook (vendor-conditional tasks). | 
Prove it:
- git log shows ≥ 30 commits with meaningful messages.
- make backup && make report runs end-to-end without manual intervention.
- One playbook successfully touches both a Cisco and a MikroTik device.
The portfolio piece. Everything from Labs 1–15 in one topology, deployed by Ansible from your Lab 15 repo.
Topology:
              [TACACS+ / Prometheus / Grafana / LibreNMS / ntopng]
                                    |
                                [MGMT vSwitch]
                                    |
     +------- AS 65000 (your ISP/enterprise) -------+
     |                                               |
    PE1 ============ MPLS Core (LDP) ============= PE2
    / \                                             / \
CE1A  CE1B (VPN_RED)                         CE2A  CE2B (VPN_BLUE)
                                               |
                                          [DMVPN Phase 3 spoke
                                           over simulated Internet
                                           terminating in VPN_BLUE]
Mandatory checklist:
- OSPF in MPLS core, LDP, MP-BGP VPNv4 between PEs
- VPN_RED: OSPF PE-CE. VPN_BLUE: eBGP PE-CE.
-  Overlapping 10.0.0.0/24 between VRFs — isolation proven
- DMVPN Phase 3 from one branch into VPN_BLUE on PE2, IKEv2 protection
