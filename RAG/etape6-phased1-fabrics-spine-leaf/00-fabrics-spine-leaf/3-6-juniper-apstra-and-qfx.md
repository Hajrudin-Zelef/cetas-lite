---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/3-6-juniper-apstra-and-qfx
title: "3.6 Juniper: Apstra and QFX"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Microsoft", "Nvidia"]
dates: []
keywords: ["acquisition", "cost", "dci", "nvidia"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [142, 157]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: cfd84cfb91702a92b98181f2a1387298ac93d1d464be8b54bac0c7b9d793cbde
---

# 3.6 Juniper: Apstra and QFX

HPE's acquisition of Juniper (completed mid-2025, ~$14B) [secondary] reshapes this portfolio: Juniper Apstra + QFX and Aruba CX/Fabric Composer now sit under one roof; integration roadmaps were still being published through 2026 [secondary]. **[unverified]** — long-term NOS consolidation outcomes (AOS-CX vs Junos vs Apstra-managed) remain officially uncommitted as of this writing.

### 3.6 Juniper: Apstra and QFX

Juniper's data-center story is intent-based: **Apstra** (acquired 2021) manages QFX/PTX fabrics from declared intent, with built-in validation, telemetry, and support for SONiC as well as Junos devices [official]. Reference blueprints cover 2-stage and 5-stage Clos (collapsed, 3-stage, 5-stage templates are Apstra's standard vocabulary) with EVPN-VXLAN overlays [official].

Hardware: QFX5120/5130/5220/5230/5240/5210 lines span leaf to spine; PTX for super-spine/DCI roles [official]. Juniper publishes "data center fabric" validated designs with eBGP underlays and documented oversubscription options per blueprint [official].

### 3.7 Whitebox / community SONiC designs

Community SONiC (Software for Open Networking in the Cloud, originally Microsoft) runs on whitebox/brite-box switches (Accton/EdgeCore, Celestica, etc.) and on vendor platforms (Dell, NVIDIA, Arista via AVD-adjacent tooling) [secondary]. Reference fabric designs are community-maintained: the classic Microsoft Azure SONiC deployment (tens of thousands of switches, multi-tier Clos) is the existence proof at hyperscale [secondary].

Design tooling: SONiC fabrics are typically provisioned with Ansible/Nornir + Jinja2 templates or vendor wrappers; Batfish/SuzieQ-style validation (see Phase E) is common in the community [secondary]. Trade-off vs vendor NOS: maximum openness and cost leverage against self-supported integration burden [independent].

---

