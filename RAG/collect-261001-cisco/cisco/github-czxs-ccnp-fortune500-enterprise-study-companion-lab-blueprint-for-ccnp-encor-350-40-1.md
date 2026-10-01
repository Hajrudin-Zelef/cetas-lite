---
id: collect-261001-cisco/cisco/github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-40-1
title: "github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena"
domain: cisco
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["aws", "capex", "cost", "dci", "governance", "incident", "license"]
source: docs/RAG/collect-261001-cisco/github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena.md
source_anchor: ""
source_lines: [1, 158]
sha256: 1e3db6c5b6a34e7557d2bd6fc1a5984d3bdbde971f088cd78ebbdf3c7901b522
---

# github-czxs-ccnp-fortune500-enterprise-study-companion-lab-blueprint-for-ccnp-encor-350-401-ccnp-ena

Production-grade reference architecture for a **global Fortune 500 enterprise network** — 4 regions, 2 data centers, 30+ branch offices, VXLAN EVPN fabric, Cisco SD-Access campus, AWS + Azure + GCP cloud integration, zero-trust security, and full NetDevOps automation. Covers 100% of the Cisco CCNP ENCOR (350-401) and ENARSI (300-410) blueprints and ~80% of the CCIE Enterprise Infrastructure blueprint.

This is the **senior network architect** companion to my CCNA Production Network portfolio. That project is associate-level; this one is principal-level.

```
graph TB
    subgraph Cloud[Cloud Providers]
        AWS[AWS<br/>Direct Connect]
        AZ[Azure<br/>ExpressRoute]
        GCP[GCP<br/>Cloud Interconnect]
    end
    subgraph Internet[Dual ISP Per Region]
        ISPA1[ISP-A NAMER]
        ISPB1[ISP-B NAMER]
    end
    subgraph NAMER[NAMER Region]
        NEDGE1[R-EDGE-NAMER-1]
        NEDGE2[R-EDGE-NAMER-2]
        NHUB[Regional Hub]
    end
    subgraph EMEA[EMEA Region]
        EEDGE[Edge Pair]
        EHUB[Regional Hub]
    end
    subgraph APAC[APAC Region]
        APEDGE[Edge Pair]
        APHUB[Regional Hub]
    end
    subgraph LATAM[LATAM Region]
        LEDGE[Edge Pair]
        LHUB[Regional Hub]
    end
    subgraph DC1[Primary DC - VXLAN EVPN]
        SPINE1[Spines 4x]
        LEAF1[Leaves 8x]
    end
    subgraph DC2[DR DC - VXLAN EVPN]
        SPINE2[Spines 4x]
        LEAF2[Leaves 4x]
    end
    subgraph Campus[HQ Campus - SD-Access]
        DNA[DNA Center]
        FBORDER[Fabric Border]
        FEDGE[Fabric Edge]
        WLC[WLC 9800 HA]
    end
    subgraph Branches[30+ SD-WAN Branches]
        BR[Branch Spokes]
    end
    Cloud --> NEDGE1
    ISPA1 --> NEDGE1
    ISPB1 --> NEDGE2
    NEDGE1 --> NHUB
    NEDGE2 --> NHUB
    NHUB --> DC1
    NHUB --> Campus
    EHUB --> DC2
    NHUB -.SD-WAN Overlay.- EHUB
    NHUB -.SD-WAN Overlay.- APHUB
    NHUB -.SD-WAN Overlay.- LHUB
    NHUB -.SD-WAN Overlay.- Branches
    DC1 <-.DCI VXLAN Stretch.-> DC2
```
    - **Senior Network Engineer** roles ($130-180K USD) — demonstrates BGP, MPLS, VXLAN EVPN, SD-WAN, cloud integration
- **Principal Network Engineer** roles ($180-250K USD) — reference-architecture-level documentation + ADRs + compliance mapping
- **Network Architect** roles ($150-220K USD) — design decisions, trade-offs, cost modeling, migration playbooks
- **CCIE Enterprise** candidate portfolio — ~80% of the v1.1 blueprint represented

This is not a cookbook. It is an **opinionated reference architecture** with decisions, alternatives considered, scale limits, availability math, and compliance controls all documented at the depth a senior audit or design review would require.

```
git clone https://github.com/calvinanglo/CCNP-Fortune500-Enterprise.git
cd CCNP-Fortune500-Enterprise
```
**Reading paths by audience:**

| Audience | Start Here | 
|---|---|
| Hiring manager (30 seconds) | Executive Summary → Master Architecture | 
| Senior engineer / architect (2 hours) | Design Decisions → HA Analysis → Scale Limitations | 
| Auditor / compliance | Compliance Mapping → section 09 | 
| CFO / procurement | Cost Model | 
| Lab builder / studier | Lab Setup → Cert Prep | 
| Operations / NOC | Operations → Runbooks | 

Full coverage of the CCNP ENCOR + ENARSI blueprints:

| Domain | Weight | Where | 
|---|---|---|
| Architecture | 15% | Sections 02, 03, 04 | 
| Virtualization | 10% | VRF design, VXLAN | 
| Infrastructure (L2/L3/IPv6) | 30% | Foundation concepts, Global WAN | 
| Network Assurance | 10% | Services & Observability | 
| Security | 20% | Security Architecture | 
| Automation | 15% | NetDevOps | 

| Domain | Weight | Where | 
|---|---|---|
| Layer 3 Technologies | 35% | BGP/OSPF/EIGRP deep dives, MPLS L3VPN | 
| VPN Technologies | 20% | DMVPN, Cloud Integration | 
| Infrastructure Security | 20% | Security Architecture | 
| Infrastructure Services | 25% | Services & Observability, DNS | 

```
CCNP-Fortune500-Enterprise/
├── README.md                                <- you are here
├── MASTER-ARCHITECTURE.md                   <- single-scroll consolidated reference
├── DESIGN-DECISIONS.md                      <- 50+ ADRs (Architecture Decision Records)
├── SCALE-LIMITATIONS.md                     <- where the design breaks + evolution path
├── HA-ANALYSIS.md                           <- 5-nines target + availability math + SPOFs
├── COMPLIANCE-MAPPING.md                    <- PCI-DSS / HIPAA / SOC2 / FedRAMP / GDPR
├── COST-MODEL.md                            <- CapEx + OpEx + TCO per component
├── MIGRATION-PLAYBOOK.md                    <- greenfield vs. brownfield adoption paths
├── LICENSE / .gitignore
│
├── 00-Executive-Summary/                    <- business context + SLA targets
├── 01-Foundation-Concepts/                  <- BGP, advanced OSPF, MPLS, VRF, QoS, IPsec, DMVPN (10 deep dives)
├── 02-Global-WAN/                           <- 4 regions, BGP, SD-WAN, DMVPN, 15 configs, 7 setup guides
├── 03-Data-Center-Fabric/                   <- VXLAN EVPN spine-leaf, 16 NX-OS configs, 9 setup guides
├── 04-SD-Access-Campus/                     <- DNA Center, LISP/VXLAN overlay, ISE integration, TrustSec
├── 05-Cloud-Integration/                    <- AWS DX, Azure ExpressRoute, GCP Interconnect, Terraform
├── 06-Security-Architecture/                <- Firepower clusters, ISE, TrustSec, MACsec, zero-trust
├── 07-Services-Observability/               <- Infoblox IPAM, streaming telemetry, SLOs, golden signals
├── 08-NetDevOps-Automation/                 <- NetBox, Ansible Tower, Terraform, Python, YANG, CI/CD
├── 09-Compliance-Governance/                <- PCI/HIPAA/SOC2/FedRAMP/GDPR control matrices
├── 10-Operations/                           <- NOC ops, DR/BC, 15 named-incident runbooks
├── 11-Troubleshooting/                      <- protocol-specific methodology (BGP/VXLAN/SD-WAN/IPsec)
├── 12-Cert-Prep/                            <- ENCOR/ENARSI/CCIE blueprint mapping + 6mo study plan
├── 13-Lab-Setup/                            <- CML-Personal, Containerlab, EVE-NG setup
├── diagrams/                                <- Mermaid source files for all topology views
└── screenshots/                             <- proof-of-build capture guide + placeholders
```
- **BGP multi-homed** internet at every regional edge with provider-independent address space
- **SD-WAN fabric** (Cisco Catalyst SD-WAN / Viptela architecture) overlaying MPLS + internet
- **DMVPN Phase 3** demonstrated as alternative/backup WAN pattern
- 30+ branch offices as SD-WAN spokes with local internet breakout
- Full IPv4 + IPv6 addressing plans

- **4 spine + 8 leaf** primary DC on Nexus 9000 (NX-OS)
- **Border leaves** for external routing handoff
- **Service leaves** for stateful service insertion (firewall, load balancer)
- **VPC** for multi-homed server attachment
- **Multi-tenant VRFs** — one per business unit
- DR DC in different region with **OTV / VXLAN stretch** DCI

- **Fabric roles:** Border (2), Control Plane (2), Edge (many)
- **Underlay:** IS-IS (documented trade-off vs. OSPF)
- **Overlay:** LISP control + VXLAN data plane
- **Wireless:** WLC 9800 HA SSO pair, fabric-enabled SSIDs
- **Client onboarding:** Cisco ISE with 802.1X / MAB / TrustSec SGTs end-to-end

- **AWS Direct Connect** 2× diverse paths to Transit Gateway
- **Azure ExpressRoute** with Microsoft + Private peering
- **GCP Cloud Interconnect** (dedicated)
- **Hybrid DNS** with split-horizon + conditional forwarders
- **SD-WAN Cloud OnRamp** for SaaS (O365, Salesforce) + IaaS
- **Terraform** modules for all cloud networking

