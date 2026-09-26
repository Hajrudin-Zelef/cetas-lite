---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-13-8-glossary
title: "E4.13.8 Glossary"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [521, 542]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: cdf9efec514c6058bd7cd5de8796995dae4a8d6f13beebd02b2673c6a11e56bf
---

# E4.13.8 Glossary

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

