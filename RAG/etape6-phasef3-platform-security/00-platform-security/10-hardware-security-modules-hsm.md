---
id: etape6-phasef3-platform-security/00-platform-security/10-hardware-security-modules-hsm
title: "10. Hardware Security Modules (HSM)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AWS", "Apple", "CISA", "EU"]
dates: ["2024-04", "2024-11", "2025-01", "2025-09", "2026-03", "2026-06", "2026-08"]
keywords: ["acquisition", "aws", "cost", "cyber", "cybersecurity", "luna", "pricing", "revenue"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [102, 124]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 5077451faf230fcade9367874bdd87f1bc153f3b955d30a7e4db624beb2397a4
---

# 10. Hardware Security Modules (HSM)

## 10. Hardware Security Modules (HSM)

- HSMs are dedicated tamper-resistant devices that generate, store, and protect cryptographic keys in hardware; keys never leave the module in unprotected form. Core use cases: PKI/CA hierarchies, TLS key protection, code signing, database encryption (TDE), blockchain/digital assets, 5G [secondary](https://www.saitechincorporated.com/saitech-inc-strengthens-port-of-oaklands-cybersecurity-with-entrust-nshield-hsm-solutions/).
- **Market concentration:** Thales, Entrust, Utimaco, IBM, and Futurex hold roughly four-fifths of global HSM revenue; the competitive contest has shifted to crypto-agility and certification [secondary](https://highways.today/2026/08/05/quantum-deadlines/).
- **Thales Luna (2026):**
  - Luna HSMs are FIPS 140-3 Level 3 validated and Common Criteria EAL4+ certified; core claims: on-prem/cloud/hybrid deployment, quantum-safe crypto agility, high-performance crypto, broad ecosystem integration [vendor-reported](https://www.gartner.com/reviews/product/thales-luna-hsm?marketSeoName=operational-technology-security&vendorSeoName=thales&productSeoName=thales-luna-network-hsm&industry=265).
  - **Luna 8:** launched ~August 2026, the first HSM on a new Thales platform, available as a network appliance; independently assessed against FIPS 140-3 Level 3 and EU Common Criteria [secondary](https://highways.today/2026/08/05/quantum-deadlines/).
  - **Thales TCT (Trusted Cyber Technologies) Luna T-Series:** first U.S.-manufactured HSMs to achieve FIPS 140-3 Level 3 (developed/produced/supported entirely in the U.S. for federal supply-chain requirements). Firmware v7.15.1 passed NIST testing for ML-DSA, ML-KEM, and LMS — the first HSM to include all NSA CNSA 2.0 post-quantum algorithms in a FIPS 140-3 validation; validated module "Luna T7" under certificate **5450**, built into Luna PCIe HSM, Luna Network HSM, Luna as a Service, and CipherTrust Manager [vendor-reported](https://www.executivebiz.com/articles/thales-tct-luna-t-series-hsm-fips-140-3-validation).
  - Thales TCT product brief (older): Luna Network HSM T-2000 (standard) / T-5000 (enterprise) performance tiers; up to 10x performance vs last generation [vendor-reported](https://www.thalestct.com/wp-content/uploads/2022/09/luna-network-hsm-pb-11-29-23.pdf).
  - DigiCert Trust Lifecycle Manager integrates with Luna for PKI root-of-trust, ACME/CMPv2/EST protocols, post-quantum readiness [vendor-reported](https://www.digicert.com/content/dam/digicert/pdfs/solution-brief/secure-trust-with-digicert-and-thales.pdf).
  - PeerSpot 2026 comparison (Feb 2026): Luna ranked 1st in General Purpose HSM (10.0/1 review); nShield 3rd [secondary](https://www.peerspot.com/products/comparisons/entrust-nshield-hsm_vs_thales-luna-hsm).
- **Entrust nShield (2026):**
  - FIPS 140-2 and 140-3 certified; Security World key-management architecture; CodeSafe secure execution environment for custom applications; network, PCIe, and USB form factors; PQC algorithm support [secondary](https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm).
  - Pricing signal: Techjockey lists nShield Connect **starting at $10,000** (callback for exact pricing — indicative, not a list price) [secondary](https://www.techjockey.com/detail/entrust-nshield-connect). PeerSpot notes nShield has higher initial cost than Luna but justifies it with features/long-term value [secondary](https://www.peerspot.com/products/comparisons/entrust-nshield-hsm_vs_thales-luna-hsm).
  - Corporate changes: CEO changed 31 March 2026 — Tony Ball succeeded Todd Wilkinson after 17 years [secondary](https://www.thetechbag.com/entrust). Entrust sold its entire public TLS certificate business to Sectigo (announced January 2025, completed 18 September 2025) after Chrome/Apple/Mozilla distrusted its roots from November 2024; remaining portfolio: nShield HSMs, private PKI/CLM, and identity verification (Onfido acquisition, April 2024) [secondary](https://www.thetechbag.com/entrust).
  - nShield Connect XC holds Bureau of Indian Standards (BIS) certification — a deciding factor for Indian government/BFSI procurement [secondary](https://www.thetechbag.com/entrust).
- **YubiHSM 2:** Yubico's entry-level HSM (USB form factor, ~$650 street price historically) for key storage, code signing, and small PKI; no 2026 refresh found in this pass — flagged as gap [unverified].
- **AWS CloudHSM:** FIPS 140-2 Level 3 validated, single-tenant dedicated HSM instances, pay-as-you-go pricing, integration with AWS KMS-adjacent services; used for customer-managed roots of trust in cloud [secondary](https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm). Exact 2026 hourly pricing not captured — flagged as gap [unverified].
- **Regulatory drivers 2026:** June 2026 U.S. executive order set federal PQC deadlines of end-2030 (key establishment) and end-2031 (digital signatures); France's ANSSI will stop certifying security products lacking quantum-safe encryption from 2027, with government/critical-infrastructure buyers expected to procure only quantum-safe products by 2030 [secondary](https://highways.today/2026/08/05/quantum-deadlines/). Harvest-now-decrypt-later risk is acute for OT/long-lived infrastructure [secondary].
- FIPS 140-3 levels: Level 1 (software), Level 2 (tamper-evidence), Level 3 (tamper resistance, identity-based auth — the HSM standard tier), Level 4 (active tamper response, environmental protection) [secondary].

## 11. Self-encrypting drives (SED): TCG Opal, FIPS drives, key management

