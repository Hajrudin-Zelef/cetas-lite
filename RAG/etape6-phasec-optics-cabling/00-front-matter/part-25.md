---
id: etape6-phasec-optics-cabling/00-front-matter/part-25
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 25)"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory", "parameters"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [926, 934]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 69f0ff6d52a24bfba899409e4e5020f35a1d53b3b7025419f19eaf74b3455dba
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 25)

#### 5.7.2 Fluke DSX certification [official — Fluke Networks DSX CableAnalyzer datasheet, media.fluke.com]
- **DSX2-8000**: certifies Cat5e through Cat8/Class I/II to 2000 MHz; **Cat6A/Class EA autotest ≈ 8 s; Cat8/Class I/II ≈ 16 s**; Level VI/2G (2 GHz) accuracy; measures wire map, length, propagation delay, delay skew, DC loop resistance, insertion loss, return loss, NEXT, ACR-N, ACR-F/ELFEXT, PS NEXT, PS ACR-N, **PS ANEXT and PS AACR-F (alien crosstalk)**; resistance unbalance; shield integrity check with distance-to-fault; TCL/ELTCTL/CDNEXT/CMRL; internal memory ≈ 12,000 Cat6A / 5,000 Cat8 results with plots; LinkWare Live cloud project management; optional TERA and GG45/ARJ45 adapters; 2 GHz channel and permanent-link adapters.
- **DSX2-5000**: to 1000 MHz (Cat6A/Class EA/FA); Cat6A autotest ≈ 10 s; built-in alien crosstalk capability [official].
- DSX series is Intertek (ETL) verified to IEC 61935-1 and ANSI/TIA-1152-A accuracy levels [official/vendor via itm.com listing].

#### 5.7.3 Alien crosstalk testing for Cat6A
- PS ANEXT / PS AACR-F are the alien-crosstalk parameters measured by the DSX family [official — Fluke datasheet]. Field AXT testing is the mechanism for certifying UTP Cat6A installations against the alien-crosstalk spec [secondary — general industry knowledge; specific "mandatory in all cases" guidance not found in sources — flagged as partial gap].
- Practical notes: Cat6A certification costs/time are dominated by AXT sampling on UTP; shielded systems simplify AXT compliance (vendor positioning) [unverified — inferred from vendor materials; no independent cost-per-link figure found].

