---
id: collect-261001-cisco/cisco/github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-40-2
title: "github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena"
domain: cisco
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "capex", "cost", "governance", "safeguards"]
source: docs/RAG/collect-261001-cisco/github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena.md
source_anchor: ""
source_lines: [159, 225]
sha256: 89fa264b0cbe184bc8455ab97a39c25359d590af816d0df315a70ed970de8918
---

# github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena

- **Firepower 4100 clusters** (Active/Active) at every internet edge
- **Palo Alto PA-series** as alternative vendor at one region (multi-vendor demonstration)
- **Cisco ISE cluster** (2 PAN + 4 PSN + 2 MnT) for AAA
- **TrustSec SGTs** campus → WAN → DC
- **MACsec (802.1AE)** on all inter-DC and inter-region links
- **Cisco Umbrella** DNS-layer security
- **Arbor APS** DDoS mitigation at internet edge

- **Infoblox Grid** (4-member HA) for IPAM
- **Splunk Enterprise** cluster for logs
- **ThousandEyes** for end-user experience monitoring
- **Streaming telemetry** via gNMI to Prometheus
- **Flexible NetFlow v9** export
- **SLIs + SLOs + error budgets** defined per service

- **NetBox** as single source of truth
- **Ansible AWX/Tower** for config management + compliance scanning
- **Terraform** for cloud infrastructure
- **Python** (Nornir + Netmiko + NAPALM) for operational scripts
- **Cisco DNA Center API** integration
- **NETCONF / RESTCONF / gNMI** for model-driven ops
- **CI/CD pipelines** (GitLab CI, GitHub Actions) with Batfish pre-deploy validation

**99.99% (four nines)** across core routing and data center fabric — ~52 minutes of downtime per year maximum. See HA Analysis for the math, audit, and evolution path to five nines.

Designed against the following frameworks (see Compliance Mapping):

- **PCI-DSS v4.0** — network segmentation, encryption in transit, access logging
- **HIPAA Security Rule** — technical safeguards, audit controls
- **SOC 2 Type II** — network-layer controls for all five trust service criteria
- **FedRAMP Moderate** — NIST 800-53 network controls mapped
- **GDPR** — data residency via regional VRF segmentation

**CCNA Production Network** — the associate-level foundation. Recommended reading order:

1. CCNA repo (small → medium → large enterprise)
2. This repo (Fortune 500 reference architecture)

Together they demonstrate the full trajectory from associate to architect.

- Master Architecture — everything in one scrollable reference
- Design Decisions — 50+ ADRs covering every major choice
- HA Analysis — availability math, SPOFs, 5-nines evolution
- Scale Limitations — where this breaks, how to extend
- Compliance Mapping — control matrix across 5 frameworks
- Cost Model — TCO, CapEx/OpEx, cloud vs. on-prem
- Migration Playbook — greenfield + brownfield paths

- 00 Executive Summary — business requirements, SLAs, service catalog
- 01 Foundation Concepts — BGP, advanced OSPF, MPLS, VRF, QoS, IPsec, DMVPN
- 02 Global WAN — multi-region BGP + SD-WAN + DMVPN
- 03 Data Center Fabric — VXLAN EVPN on Nexus 9000
- 04 SD-Access Campus — DNA Center + LISP + ISE + TrustSec
- 05 Cloud Integration — AWS / Azure / GCP
- 06 Security Architecture — Firepower + ISE + zero-trust
- 07 Services & Observability — Infoblox, Splunk, telemetry, SLOs
- 08 NetDevOps & Automation — NetBox, Ansible, Terraform, Python

- 09 Compliance & Governance — PCI/HIPAA/SOC2/FedRAMP/GDPR
- 10 Operations — NOC ops, DR/BC, runbooks
- 11 Troubleshooting — protocol-specific methodology

- 12 Cert Prep — ENCOR + ENARSI + CCIE blueprint mapping, 6-month study plan
- 13 Lab Setup — CML-Personal, Containerlab, EVE-NG

- Diagrams — Mermaid source for all topology views
- Screenshots — capture guide with ~100 checkpoints
