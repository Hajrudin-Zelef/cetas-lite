---
id: collect-261001-cisco/cisco/github-nasser-souidi-enterprise-multi-site-network-cisco-lab-gns3-core-sw1-sw2-etherchanne
title: "github-nasser-souidi-enterprise-multi-site-network-cisco-lab-gns3-core-sw1-sw2-etherchannel-port-cha"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/github-nasser-souidi-enterprise-multi-site-network-cisco-lab-gns3-core-sw1-sw2-etherchannel-port-cha.md
source_anchor: ""
source_lines: [1, 38]
sha256: da0ad013991c91bc1e6503706f1bf8d2e26f6fc1f70ceebf501c9a31b426bd04
---

# github-nasser-souidi-enterprise-multi-site-network-cisco-lab-gns3-core-sw1-sw2-etherchannel-port-cha

End-to-end simulation of an enterprise network connecting a **Headquarters (HQ)** site and a **Remote Branch** office over a simulated ISP, built entirely in **GNS3** with Cisco IOS. The lab reproduces real-world enterprise WAN/LAN design: multi-VLAN segmentation, Layer 3 high availability, dynamic routing with redistribution, and a secure site-to-site VPN.

**HQ site**

- Core: SW1 & SW2 — EtherChannel (Port-channel) + HSRP + OSPF, 4 parallel links
- Access: ESW1–ESW4, one VLAN per switch (10/20/30/40) + management VLAN 99
- Edge: Router R1 — OSPF ↔ EIGRP redistribution, NAT/PAT, VPN endpoint

**Remote site**

- Distribution: SW3–SW6 in a routed chain (dual links between each pair, ECMP via OSPF)
- Access: ESW5 & ESW6, single user VLAN with HSRP active/standby (SW3/SW6)
- Edge: Router R2 — 4 routed uplinks (one per distribution switch), OSPF ↔ EIGRP redistribution, VPN endpoint

**WAN**

- Simulated ISP router (public transit)
- R1 ↔ R2: GRE tunnel encrypted with IPsec (site-to-site VPN)

*(Drop your topology screenshot here, e.g. `docs/topology.png`)*

- **Routing** : OSPF (per site, area 0), EIGRP (inter-site, over the GRE tunnel), mutual OSPF↔EIGRP redistribution
- **VPN** : GRE over IPsec (ISAKMP/IKEv1, AES, SHA, pre-shared key)
- **High availability** : HSRP (tuned priorities, preempt), EtherChannel (LACP active)
- **Switching** : VTP (server/client), 802.1Q trunking, Rapid-PVST+ with per-VLAN root bridge tuning
- **Services** : DHCP (per-VLAN scopes), NAT/PAT (overload) for Internet access
- **Management** : SSH v2 only, local authentication

- SSH-only remote access (Telnet disabled)
- Encrypted passwords (service password-encryption) + secret hashing
- Port-security (sticky MAC, violation restrict) on all access ports
- BPDU Guard + PortFast on edge ports
- Extended ACLs for inter-VLAN traffic segmentation
- All inter-site traffic encrypted via IPsec

Network design · L2/L3 routing & switching · Route redistribution · Site-to-site VPN (GRE/IPsec) · High availability (HSRP/EtherChannel) · VLAN segmentation & security · Structured troubleshooting

nasser souidi — built as a hands-on lab to practice enterprise-grade network design, security, and troubleshooting.
