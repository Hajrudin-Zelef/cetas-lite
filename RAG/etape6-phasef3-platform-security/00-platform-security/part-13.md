---
id: etape6-phasef3-platform-security/00-platform-security/part-13
title: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 13)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AWS"]
dates: []
keywords: ["aws", "cost", "cybersecurity", "licenses", "luna", "pricing", "training"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [344, 355]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: a9e96895445d3abf3f749c1fabff743d279b2d4696d13eb8fc1f539ac7e4d393
---

# Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 13)

- **Form factors:** network appliance (Luna Network HSM, nShield Connect), PCIe card (Luna PCIe, nShield Solo), USB (YubiHSM 2), cloud dedicated (AWS CloudHSM), cloud multi-tenant service (Luna as a Service, Azure Dedicated HSM) [secondary].
- **Tenancy models:** single-tenant dedicated (CloudHSM, Dedicated HSM) vs multi-tenant with partitions (Luna: up to 100 partitions per appliance, each with own SO/users/keys) [secondary].
- **APIs:** PKCS#11 (universal), JCE (Java), CNG/KSP (Windows), OpenSSL ENGINE/provider, REST (CloudHSM, Luna SA), vendor SDKs [secondary].
- **M-of-N quorum:** key-management operations (backup, partition init) require M of N smart-card/PED tokens — split-knowledge control [secondary].
- **Key backup:** encrypted cloning domains (Luna), smart-card sets (nShield Security World), CloudHSM backup to S3 [secondary].
- **Use-case mapping:** public CA roots → network HSM (FIPS 140-3 L3, CC EAL4+); database TDE → PCIe or network; code signing → network + timestamping; blockchain/digital assets → FIPS L3 with custom CodeSafe apps; dev/test → YubiHSM 2 or cloud [secondary].
- **Certification as procurement gate:** FIPS 140-3 Level 3 is the standard HSM bar; Common Criteria EAL4+ adds assurance; U.S. federal may require U.S.-manufactured (Thales TCT line) [vendor-reported](https://www.thalestct.com/wp-content/uploads/2022/09/luna-network-hsm-pb-11-29-23.pdf).
- **PQC posture (2026):** Thales TCT Luna T-Series v7.15.1 (cert 5450) — ML-DSA/ML-KEM/LMS validated; Luna 8 — new PQC-oriented platform; nShield — PQC support advertised; verify per-firmware-version validation status before procurement [vendor-reported][secondary].
- **Price visibility (2026, indicative):** nShield Connect aggregator listing "starting at $10,000" (callback pricing); Port of Oakland nShield delivery reported >$19K per unit (secondary snapshot, not OEM list); exact list prices unpublished — G7 [secondary](https://www.techjockey.com/detail/entrust-nshield-connect) [secondary](https://www.saitechincorporated.com/saitech-inc-strengthens-port-of-oaklands-cybersecurity-with-entrust-nshield-hsm-solutions/).
- **Hidden costs:** per-partition licensing, client licenses, support contracts, PED/smart-card kits, HA clustering (second appliance), training [secondary].
- **Cloud vs on-prem trade-off:** CloudHSM = no hardware to manage, single-tenant, but cloud-region jurisdiction and hourly cost; on-prem = capital cost + operational burden, full physical control [secondary](https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm).

