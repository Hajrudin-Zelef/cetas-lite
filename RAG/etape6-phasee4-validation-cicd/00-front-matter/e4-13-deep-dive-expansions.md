---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-13-deep-dive-expansions
title: "E4.13 — Deep-dive expansions"
domain: front-matter
role: reference
task: reference
actors: ["AWS", "Huawei"]
dates: ["2025-10", "2026-05", "2026-07"]
keywords: ["agents", "aws", "embedding", "license", "licenses", "mcp", "research", "sandbox"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [465, 542]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: c8324b750b92669a9ec58f7d03b9b7edfd63095d508280bdafab4f995424e333
---

# E4.13 — Deep-dive expansions

## E4.13 — Deep-dive expansions

### E4.13.1 IP Fabric: vendor coverage breadth and version 8.0

- IP Fabric publishes a per-version feature matrix of supported vendors and discovery tasks: v7.10.0 and v8.0.0 matrices exist at matrix.ipfabric.io `[official]` — https://matrix.ipfabric.io/ and https://matrix.ipfabric.io/7.10.0.
- The v8.0.0 matrix (updated ~63 days before research, i.e., ~July 2026) lists ~60 vendor/platform targets including: AWS, Azure, GCP, Cisco (IOS, IOS-XE, NX-OS, IOS-XR, ASA, FTD, ACI, Meraki, Viptela, WLC, SG, ENCS), Arista EOS, Juniper JunOS + Mist, Palo Alto PAN-OS + Prisma, Fortinet FortiGate/FortiSwitch, Check Point Gaia, F5 BIG-IP, Dell (FTOS, OS10, PowerConnect), HPE (Comware, Aruba, ArubaCX, Aruba switches, Instant APs), Extreme (BOSS, Enterasys, VOSS, EXOS), Huawei VRP, Nokia TIMOS, MikroTik RouterOS, D-Link, FS (FSOS), ZPE Nodegrid, FRR/Quagga, Citrix ADC, Riverbed, Opengear, Ruckus VSZ, Stormshield, Forcepoint, Versa VOS, Silver Peak, VMware NSX-T/VeloCloud `[official]` (v8.0.0 matrix).
- Per-task granularity: the matrix scores each vendor on 17 basic-discovery tasks (ARP, device info, etc.) plus advanced tasks; new additions in v8.0 include ZPE Nodegrid and expanded D-Link coverage `[official]` (matrix comparison 7.10 → 8.0).
- Dedicated Vendor APIs complement CLI discovery and are essential for SD-WAN/cloud: AWS, Azure, GCP, Check Point Management, Cisco APIC, Cisco FMC, Cisco Meraki REST, F5OS, Forcepoint, Juniper Mist, Palo Alto Prisma, Ruckus VSZ, Silver Peak, VeloCloud, Versa, Viptela, VMware NSX-T `[official]` — https://docs.ipfabric.io/latest/IP_Fabric_Settings/Discovery_and_Snapshots/Discovery_Settings/Vendors_API/.
- MFA limitation: IP Fabric does not support MFA on vendor API logins — use non-MFA service accounts or application tokens `[official]` (docs portal).
- Technology coverage (datasheet 7.11): ARP/DHCP/IPv4/IPv6/NAT/SNMP/NTP/AAA/Syslog/CDP/LLDP/MPLS/LDP/RSVP/DMVPN/flow exports/SSL certificates; STP family; HSRP/VRRP/GLBP; PIM/IGMP snooping/RPs/multicast tables; wireless controllers/APs/clients; URL filtering in path lookups `[vendor-reported]` — https://ipfabric.kinsta.cloud/wp-content/uploads/2026/04/IPF-Datasheet-7.11_EDITED.pdf.
- Inventory depth: modules/part numbers, OS versions across the fleet, interface inventory (state, description, speed, duplex, media), host discovery rules, and EoX/EoL milestone reports for Arista, Aruba, Check Point, Cisco, Meraki, Extreme, F5, Fortinet, HP/H3C/3COM, Juniper, Palo Alto, Riverbed `[official]` — https://github.com/ipfabric/docs/blob/HEAD/docs/IP_Fabric_GUI/inventory.md.
- Version note: research earlier established 7.9 (Jan 2026); the v8.0.0 matrix (~July 2026) indicates a major 8.0 release exists — feature-level 8.0 release notes were not pulled; treat 8.0 specifics as `[unverified]` pending https://matrix.ipfabric.io/ and official release notes.

### E4.13.2 CML licensing and positioning detail

- CML tiers (2026, per community/vendor-adjacent sources): Free (5 nodes, no license purchase, Cisco account required — announced October 2025), Personal $199/year (20 Cisco nodes), Personal Plus $349/year (40 nodes), Enterprise (up to 300 nodes, clustering, multi-user) `[secondary]` — https://www.pinglabz.com/cisco-modeling-labs-cml-2-8-free/, https://blog.cloudmylab.com/lab-as-a-service-cml, https://learningnetwork.cisco.com/s/question/0D56e0000DmSovTCQS/cml-licensing-period.
- Licenses are 1-year subscriptions; on expiry CML reverts to the free tier `[secondary]` (Cisco Learning Network community).
- CML-EDU-BASE academic licensing exists for education `[secondary]` — https://www.hummingbirdnetworks.com/cisco-modeling-labs-base-license-1-user-academic-esd-cml-edu-base.
- Sizing: Cisco's stated minimum is 8 GB RAM and 4 physical cores; real labs need the sum of every node's CPU/RAM/disk `[secondary]` (cloudmylab).
- PeerSpot 2026 reviews are mixed-to-negative on price/complexity (one reviewer: expense 8–9/10, "leave it in the classroom") — single-reviewer anecdote, not a survey `[secondary]` — https://www.peerspot.com/products/cisco-modeling-labs-cml-reviews.
- DevNet Sandbox offers short CML sessions (~4 hours) for try-before-buy `[secondary]` (cloudmylab).

### E4.13.3 Batfish in CI: how the integration works (detail)

- Typical CI embedding: (1) configs pulled from Git/source of truth; (2) Batfish snapshot initialized via pybatfish in a CI job; (3) question suite runs (reachability, no unintended ACL widening, BGP session symmetry); (4) differential questions compare base vs proposed snapshot; (5) job fails on violations, blocking merge `[official]` (Batfish README: "validate configuration changes before deployment"; "including Batfish in automation workflows").
- Because Batfish needs no device access, CI runners need only the config files — no lab hardware, no credentials in the validation step `[official]` (README).
- Differential analysis is the CI killer feature: prove the planned change alters *only* intended behavior `[official]` (question catalog: "End-to-end reachability is identical across the current and a planned configuration").
- pybatfish + Jupyter: interactive development of question suites, then promotion to CI scripts `[official]` (allinone notebooks).
- Gap: no 2026-published reference architecture for Batfish + GitHub Actions specifically was found; the pattern is documented generically `[unverified]`.

### E4.13.4 SuzieQ data model (detail)

- Poller → normalizer → columnar store → CLI/REST/API analysis; tables are the unit of analysis (device, interface, bgp, ospf, lldp, arpnd, macs, routes, etc.) `[official]` — https://suzieq.readthedocs.io/en/latest/poller/ and https://github.com/gelbox/suzieq.
- Time-series retention enables "what changed and when" forensics — the basis of drift detection and pre/post-change testing `[official]` (NetBox Labs partnership blog).
- Schema normalization is the differentiator vs raw SNMP/CLI scraping: one query works across vendors `[official]` (README: "stores it in a vendor independent way").

### E4.13.5 Forward Enterprise platform (detail)

- Forward Enterprise: the digital-twin platform; Forward AI (natural-language interface) and Forward Predict (pre-change verification) are capabilities on top `[vendor-reported]` (IDC Spotlight May 2026).
- Forward MCP Server: exposes twin intelligence to custom agents; closed-loop workflows across ServiceNow, Slack, Infoblox `[vendor-reported]` (IDC Spotlight).
- Every Forward AI answer is cited to the exact path, policy, or config line as evidence `[vendor-reported]` (IDC Spotlight).
- Multi-vendor + multi-cloud: AWS, Azure, GCP, IBM `[vendor-reported]` (IDC Spotlight).

### E4.13.6 pyATS component map (detail)

- pyATS core: test framework, topology/testbed YAML, aetest (test sections: common_setup/testcase/cleanup), Easypy runner, CLI `[official]` (DevNet).
- Unicon: connection plugin library (SSH/Telnet/console) `[official]` (DevNet release notes).
- Genie: parser library (hundreds of `show` command parsers), Ops (operational models), libs (conf, clean, filetransferutils, health, sdk, telemetry, trafficgen), robot integration `[official]` (Cisco Live lab package list: genie 26.5 with libs.clean/conf/filetransferutils/health/ops/parser/robot/sdk/telemetry/trafficgen).
- Clean: device bring-up/recovery automation (image upgrade, password recovery in 24.8) `[official]` (DevNet 24.8 notes).
- Dyntopo: dynamic topology/testbed generation (LaaSv2 gateway IP fixes in 24.7) `[official]` (DevNet 24.7 notes).
- Version support: Python 3.8 dropped after Oct 2024; 3.11 in the 2026 lab `[official]` (DevNet).

### E4.13.7 End-to-end validation workflow (worked example, synthesis)

Scenario: add a new leaf pair to an EVPN-VXLAN fabric.

1. **Design in Git.** Topology, configs, and intent (VLANs, VNIs, BGP ASN plan) committed to a feature branch. Source of truth: NetBox/Nautobot `[analysis]`.
2. **Static validation (CI, no devices).** Batfish snapshot of proposed configs: questions — (a) new VTEPs reachable from all existing VTEPs; (b) BGP underlay sessions form; (c) no ACL change widens access; (d) differential vs current snapshot shows only intended deltas. Fail = block merge `[official]` (Batfish question catalog) + `[analysis]`.
3. **Virtual lab test (CI).** Containerlab/CML spins the exact topology; Robot/pyATS suite: BGP Established on all peerings, EVPN RT-2/RT-5 exchange, BUM handling, failover convergence timing, config idempotency rerun `[secondary]` (BRKATO-1009-2026 patterns) + `[analysis]`.
4. **Human review.** Plan artifact + Batfish diff + lab results attached to the PR; CODEOWNERS approval `[secondary]` (AUTOM-07) + `[analysis]`.
5. **Staged deploy.** Merge → apply to canary devices → EDA rulebook watches for anomalies (BGP flaps, interface errors) with throttle/dedupe → auto-rollback or page `[secondary]` (EDA patterns) + `[analysis]`.
6. **Post-change assurance.** SuzieQ 90-second polls confirm observed == intended; IP Fabric intent checks green; NetBox updated via discovery sync `[official]` (NetBox Labs × SuzieQ) + `[vendor-reported]` (IP Fabric) + `[analysis]`.
7. **Continuous.** Drift detection jobs, EoL reports, periodic re-validation `[analysis]`.

### E4.13.8 Glossary

- **Digital twin (network):** software model of the network built from configs/state, used for simulation and verification (Batfish offline model; IP Fabric/Forward live snapshots) `[analysis]`.
- **Intent verification:** checking observed state against declared intent (IP Fabric checks; Batfish questions) `[analysis]`.
- **Differential analysis:** comparing two config snapshots to prove only intended changes `[official]` (Batfish).
- **VTEP/NVE/VNI:** VXLAN tunnel endpoint / network virtualization edge / network identifier (see Phase D2) `[secondary]`.
- **Rulebook:** EDA YAML binding event sources → conditions → actions `[secondary]`.
- **Event Streams:** Red Hat EDA production event routing enhancement `[secondary]`.
- **Genie Learn:** pyATS operational-state snapshot feature `[secondary]`.
- **Netem:** Linux network emulation (delay/loss/jitter) exposed by Containerlab `[secondary]`.
- **Freemium (EVE-NG):** Pro ISO unlicensed, 7-node cap `[secondary]`.

---

