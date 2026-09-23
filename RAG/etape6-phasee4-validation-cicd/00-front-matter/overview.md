---
id: etape6-phasee4-validation-cicd/00-front-matter/overview
title: "Phase E4 — Network Validation, Observability & CI/CD"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-01", "2026-09-22"]
keywords: ["funding", "research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [1, 62]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 64464bec8fde5c504b14d98b12ec330cef3f227b6bed1e36a0ec09bd69b594b7
---

# Phase E4 — Network Validation, Observability & CI/CD

> Scope: network configuration validation (Batfish), network observability (SuzieQ, IP Fabric, Forward Networks, Cisco Crosswork), lab/emulation environments (Containerlab, EVE-NG, GNS3, Cisco Modeling Labs), testing frameworks (pyATS, Robot Framework, pytest), GitOps CI/CD pipelines for networks, event-driven automation (StackStorm, event-driven Ansible).
> Date: 2026-09-22. Method: read-only web research (browser_search; select page fetches). Nothing was purchased, signed into, or executed; no credentials used.
> Provenance legend: `[official]` = vendor/project official documentation or announcement; `[vendor-reported]` = vendor marketing/press claims not independently verified; `[independent]` = independent press, analyst, or community verification; `[secondary]` = third-party blogs/tutorials/aggregators; `[unverified]` = claim found but not corroborated or marked by the source itself as provisional.
> Line-count and integrity verification are recorded at the end of the file.

---

## E4.1 — Batfish: open-source network configuration analysis

### E4.1.1 Project identity and 2026 status

- Batfish is an open-source network configuration analysis tool that "finds bugs and guarantees the correctness of (planned or current) network configurations" `[official]` — https://github.com/batfish/batfish (project README).
- Core value proposition: enable network engineers to validate configuration changes *before* deployment; the analysis requires only device configurations, not direct device access `[official]` (README: "Batfish does NOT require direct access to network devices").
- Optional enrichment inputs: BGP routes received from external peers, topology from LLDP/CDP `[official]` (project README).
- The project repository `batfish/batfish` is actively maintained: architecture documentation (`docs/architecture/README.md`) was updated ~5 days before the research date, and release notes show continuous multi-vendor parser work `[official]` — https://github.com/batfish/batfish/blob/HEAD/docs/architecture/README.md.
- Batfish was originally developed inside Intent (the company); the SIGCOMM'23 paper "Batfish: A Network Configuration Analysis Tool" evolution summary describes the current architecture `[secondary]` — https://www.halper.in/pubs/halperin_batfish_sigcomm23.pdf.

### E4.1.2 Architecture (2026)

- Pipeline stages: Parsing (ANTLR-generated parsers produce parse trees) → Extraction (parse trees → vendor-specific models) → Conversion (→ vendor-independent model) → Post-processing → Data Plane Generation (RIBs/FIBs) → Forwarding Analysis `[official]` — https://github.com/batfish/batfish/blob/HEAD/docs/architecture/README.md.
- Design decisions: vendor-independent model so analysis code works across vendors; parse-tree-based extraction separating syntax from semantics; symbolic analysis over packet sets (BDD-based engine) so "all possible packets" are reasoned about without enumeration `[official]` (architecture README).
- Per the SIGCOMM'23 evolution paper: data-plane generation improved by three orders of magnitude and data-plane verification by an order of magnitude versus the original version; analysis finishes in minutes even on networks with thousands of nodes `[secondary]` (paper summary, 11 real networks benchmarked).
- Build system migrated from Maven to Bazel; CI moved from Buildkite to GitHub Actions; developer docs live under `docs/` `[official]` (release notes).

### E4.1.3 Capabilities (question-based analysis)

- Question families include: correctness (end-to-end reachability, no black holes), reliability (reachability preserved under any single-link/single-device failure), security (sensitive services reachable only from specific subnets; paths traverse firewalls / have ≥2-way ECMP), and change analysis (reachability identical across current vs planned configs; ACL/firewall changes provably collateral-free; functional equivalence of two configs, potentially cross-vendor) `[official]` (project README question catalog).
- Recent development focus (release notes): validating *incremental* changes to configurations — e.g., Arista support for deleting BGP neighbors/peer groups/prefix-list seqs, NX-OS/IOS-XR removal of BGP `aggregate-address`, IOS interfaces defined in incremental changes after router OSPF `[official]` — https://github.com/batfish/batfish/releases.
- VXLAN/EVPN support is under active development: JunOS VXLAN/EVPN work credited to contributor @jeffkala in recent release notes `[official]` (releases page) — status of full EVPN coverage remains partial `[unverified]` (no complete EVPN support matrix found in research).

### E4.1.4 Vendor coverage (from recent release notes)

- Cisco: IOS, NX-OS, IOS-XR, ASA/Firepower-era parsers; NX-OS EIGRP→EIGRP redistribution support added recently `[official]`.
- Arista EOS: BGP incremental-change support, MLAG peer-address heartbeat VRF parsing fixes `[official]`.
- Juniper JunOS: static routes with qualified-next-hops, dotted BGP ASNs, firewall filter source-interface filtering, ongoing VXLAN/EVPN `[official]`.
- FRR/SONiC: BGP parsing and inheritance improvements `[official]`.
- Palo Alto PAN-OS: template variables support, NAT address reference tracking `[official]`.
- Community contributions are significant (e.g., @Katsuya414 NX-OS work, @network-dave IOS fixes, @pawelhaj JunOS) `[official]` (release notes).

### E4.1.5 Interfaces and deployment

- `pybatfish`: Python 3 SDK for interacting with the Batfish service `[official]` — https://www.github.com/batfish/pybatfish.
- `batfish/allinone` Docker image bundles Batfish with example Jupyter notebooks; recommended run: `docker run --name batfish -v batfish-data:/data -p 8888:8888 -p 9997:9997 -p 9996:9996 batfish/allinone` `[official]` (README).
- API: V2 API on port 9996 is now sufficient for all client features; port 9997 (V1) is legacy and will be disabled by default in a future release; `pybatfish >= 2022.9.7` required for V2-only access `[official]` (release notes).
- Jupyter notebooks shipped in the allinone image walk through capabilities ("Getting Started with Batfish", question notebooks) `[official]` (README).

### E4.1.6 Integrations and ecosystem

- Batfish is positioned as the pre-deployment validation step inside network automation/CI workflows `[official]` (README: "By including Batfish in automation workflows, network engineers can close this gap").
- Slack community at `batfish-org` for questions/feedback/feature requests `[official]` (README).
- NetBox Labs ecosystem partnerships (announced January 2026) include Forward Networks, IP Fabric, Slurp'it, and SuzieQ for discovery/reconciliation — Batfish itself is not listed as a NetBox Labs partner in that announcement `[official]` — https://netboxlabs.com/blog/netbox-labs-stardust-systems-network-observability-partnership/.

### E4.1.7 Gaps and unverified items

- No 2026 funding/company-status news found for Batfish (Intent); open-source maintenance appears community + core-team driven `[unverified]`.
- Exact latest release version number and date were not captured (releases page fetched as search-result text, not the version header) `[unverified]` — verify at https://github.com/batfish/batfish/releases before citing a version.
- EVPN/VXLAN analysis completeness: partial, in progress `[unverified]`.

---

