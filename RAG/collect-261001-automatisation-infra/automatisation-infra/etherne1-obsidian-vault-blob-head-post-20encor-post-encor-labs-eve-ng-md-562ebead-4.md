---
id: collect-261001-automatisation-infra/automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead-4
title: "etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead.md
source_anchor: ""
source_lines: [312, 341]
sha256: daca65f956c8ef4da61975783e253abfe0d77fe3b382f743d2bb698bf7bb906e
---

# etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead

- AAA via TACACS+ on every router, two privilege tiers
- CoPP 4-class policy on every router
- IPv6 dual-stack everywhere: BGP IPv6 AF, OSPFv3, IPv6 ACLs, IPv6 FHS on access segment
- NetFlow v9 → ntopng; SNMPv3 → LibreNMS; gNMI → Prometheus
- IP SLA monitoring critical paths with track-based floating statics
- Ansible playbook deploys the whole topology from the Git repo
- One-page network diagram (draw.io / Excalidraw / mermaid) committed to the repo
- README: explains how to rebuild from scratch in under 30 minutes
Self-test:
- Run the CCNP ENARSI 300-410 practice exam (Pearson Test Prep). Target ≥ 85% on the 200-question pool.
- Hand the repo to someone else. They git clone +make all and rebuild your topology without asking you anything.
| Post-ENCOR topic | GNS3vault coverage | Hand-written supplement | 
|---|---|---|
| VRF-Lite | Full ( vrf-lite ,vrf-routing ) | Route-leaking extension, MikroTik cross-check | 
| BFD | None | Full hand-written | 
| Route manipulation / PBR | Partial ( policy-based-routing , BGP filtering) | Multi-AS extension, MikroTik cross-check | 
| Redistribution | Indirect (OSPF + EIGRP baselines) | Full redistribution + loop-prevention hand-written | 
| EIGRP classic | Full (23 labs in archive) | Named-mode extension hand-written | 
| OSPFv3 AFs | None | Full hand-written | 
| OSPF areas / network types | Full (12+ labs) | Modern network-type extension (no FR) | 
| IS-IS | None (containerlab topology only) | Full hand-written on FRR | 
| BGP deep | Excellent (50+ labs) | IPv6 AF extension hand-written | 
| BGP troubleshooting | Partial (3 failure-mode labs) | Gauntlet hand-written | 
| MPLS L3VPN | Good ( mpls-ldp ,basic-mpls-vpn , PE-CE OSPF) | eBGP PE-CE + 6VPE extension, Juniper cross-check | 
| DMVPN | None | Full hand-written (GRE/IPsec warm-up from archive) | 
| Infrastructure security | Excellent (AAA, ACL, CoPP, ZBF) | IPv6 FHS extension, production CoPP extension | 
| SNMP / NTP / syslog / IP SLA | Full | — | 
| Modern telemetry (gNMI / Grafana) | None | Full hand-written | 
| Automation repo | None | Full hand-written | 
| Capstone | None | Full hand-written |
