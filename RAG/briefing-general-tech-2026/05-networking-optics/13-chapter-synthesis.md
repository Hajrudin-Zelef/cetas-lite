---
id: briefing-general-tech-2026/05-networking-optics/13-chapter-synthesis
title: "Chapter synthesis: 2026, the year the fabric became the product"
domain: networking-optics
role: deep-dive
task: analysis
actors: ["Broadcom", "China", "Credo", "FS.com", "Google", "MACOM", "Meta", "Nokia", "Nvidia", "OIF", "Telxius"]
dates: ["2025-12", "2025-12-22", "2026-08-24", "2026-08-25", "2026-09", "2026-09-02", "2026-09-07", "2026-09-09", "2026-09-11", "2026-09-12", "2026-09-15", "2026-09-16", "2026-09-21", "2026-09-22", "2026-09-23", "2027-05"]
keywords: ["asic", "cpo", "dci", "dsp", "ethernet", "hyperscaler", "lpo", "muxponder", "npo", "optics", "serdes", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: ""
source_lines: [6487, 6641]
sha256: 833bb47f508b0634dd7558c7fe330b13a9d7adf426685026a9247d56914ac631
---

# Chapter synthesis: 2026, the year the fabric became the product

### Chapter synthesis: 2026, the year the fabric became the product

Step back from the twelve sections and a single arc emerges. In 2026 the
networking layer of AI infrastructure stopped being background plumbing and
became a first-order product category — with its own launch cadence, its
own seller's market, and its own financial guidance measured in hundreds of
millions of dollars. The year opened with a vendor (FS.com, December 2025)
planting a 1.6T flag ahead of the market; it closed the third quarter with
1.6T shipping in volume, capacity reserved rather than available, analysts
raising forecasts, a standards body publishing the coherent 1.6T agreement,
and a chip company guiding to $600M+ of optical revenue. Between those two
points, every layer of the stack moved one step: modules to 1.6T, coherent
to 800G pluggable as the carrier default, lane rates toward 448G, and
architectures toward linear and co-packaged optics.

The through-line for the reader is practical: anyone planning network
capacity for 2027 now works from a verified baseline — 1.6T shipping, 800G
coherent standard, linear optics deploying, 3.2T components demoed — and a
set of open questions (volumes, shares, ECOC results, guidance
confirmation) with known dates attached. That is what a reference dossier
is for: not a prediction, but a dated, sourced floor to stand on while the
next wave of facts arrives.

Three structural shifts deserve to be named explicitly, because they will
still be unfolding in 2027:

1. **Power, not bandwidth, is the binding constraint.** The 25–26 W 1.6T
   module (6.2), the LPO bet on removing the DSP (6.9), and the CPO/NPO
   sampling at 3.2T/6.4T (6.6, 6.10) are all answers to the same problem:
   the bits keep doubling, but the watts per faceplate cannot. Whoever
   solves power-per-bit owns the next two generations.
2. **Interoperability moved from press release to process.** The chapter's
   worked example — FS's "verified on NVIDIA" self-claim (6.2) versus the
   Ethernet Alliance multi-vendor demo (6.8) and the OIF 1600ZR agreement
   (6.7) — is the year's methodological lesson. In a ramp year, vendor
   claims outrun verification; the institutions that do multi-vendor
   interworking are where trust is actually manufactured.
3. **The hyperscaler and the carrier converged on the same architecture,
   one generation apart.** IP-over-DWDM with coherent pluggables is the
   carrier standard at 800G (6.12) and the data-center reality at 1.6T
   (6.6); the DCI "Scale-Across" domain (6.1) is where the two worlds
   meet. The 2027 builds both sides are reserving capacity for will be the
   first designed end-to-end around this convergence.

#### What to watch after September 22, 2026

Everything in this subsection is announced or planned, not accomplished —
the dossier's cutoff is September 22, and these are the milestones the
verified facts point toward:

- **ECOC 2026 show reports** (expected after the September 23 close):
  the outcomes of the Ethernet Alliance booth #2173 interop demo, the OIF
  1600ZR booth #2126 demo, and MACOM's booth #1055 448G-per-lane demos.
  These will turn September's announced demos into reviewed results.
- **AmpCon-Campus 3.0**: FS described it as an "upcoming release" on
  September 21 — watch for the actual release notes, and for whether the
  ".0" arrives with the usual point-fix tail.
- **The 1.6T second wave**: Meta and Oracle integrating in 2027 (per
  LightCounting), which will stretch the ramp and test whether the
  "sold out, capacity being reserved" seller's market persists or breaks.
- **Credo's FY2027 quarters**: the >$600M optical guidance (fiscal year
  ending ~May 2027) will be tested quarter by quarter — guidance versus
  booked revenue is the metric to watch.
- **LPO/CPO in 2027**: volume LPO rollouts through 2026 into 2027, and
  whether co-packaged optics moves from sampling (3.2T/6.4T) toward
  deployment — the power-per-bit contest's next round.
- **FS.com H2 2026 results**: the first half-year to include a full
  period of 1.6T volume shipments, and the test of whether the H1 margin
  expansion was a mix-shift trend or a one-half spike.

#### A note on sources: how this chapter labels its evidence

The reader will have noticed a recurring discipline: every claim carries
its provenance, and vendor claims are fenced off from independent ones.
The chapter's sources fall into three tiers. **Tier 1 — vendor press
releases** (BusinessWire, GlobeNewswire): FS.com's portfolio, 1.6T
modules, D7070 and Wi-Fi announcements; Credo's ZeroFlap; MACOM's ECOC
demos; the Telxius–Nokia joint release. These are confirmed as *issued*,
and their technical content is reported faithfully — but compatibility
and performance claims inside them are labeled as vendor claims.
**Tier 2 — standards and industry bodies**: the OIF 1600ZR agreement, the
Ethernet Alliance demo announcement, IEEE P802.3dj progress — multi-vendor
processes, inherently more trustworthy on interoperability.
**Tier 3 — independent analysis and press**: LightCounting via EETimes,
Cignal AI (as cited), Goldman Sachs, TipRanks, Zacks, Photonics Spectra.
Analyst forecasts and financial-press relays are attributed, never
laundered into facts. The chapter's one worked correction — "verified on
NVIDIA" downgraded to "FS claim (22/12/2025)" — is the method applied
throughout: in a ramp year, the press release is the beginning of
verification, not its end.

#### What this chapter does not know: acknowledged blind spots

The anti-fabrication rule cuts both ways: the dossier says what it
verified, and it names what it could not. **No verified 1.6T shipment
volumes** exist in this chapter's sources — LightCounting's "shipping
now," Cignal AI's "over 11 million forecast," and Goldman Sachs's raised
forecasts are directional, not a counted installed base. **No verified
market shares** for the 1.6T transceiver market could be established;
vendor press releases do not disclose them and no independent share table
surfaced in verification. **The FS "verified on NVIDIA" claim** remains a
self-claim — no Nvidia statement either confirming or denying it was
found, and the dossier does not fill that silence. **ECOC 2026 outcomes**
(the interop demos, the 448G-per-lane demos) were unverifiable at the
September 22 cutoff because the show was still running. **The exact 100G+
segment growth figure** for FS.com is unverifiable — no public breakdown
exists, and none is cited. And **the Telxius deployment's scale and
timeline** are undisclosed by design. These are not gaps to be papered
over with estimates; they are the questions the next verification wave —
post-ECOC reporting, Q4 earnings, H2 results — is positioned to answer.

#### Master chronology: the chapter's verified dates

| Date | Event | Section |
|---|---|---|
| 22/12/2025 | FS.com launches 1.6T OSFP line (OSFP-DR8-1.6T SiPh ≤500 m, OSFP-2FR4-1.6T EML ≤2 km, Broadcom 3 nm DSP, 25–26 W); "verified on NVIDIA Quantum-X800" vendor self-claim | 6.2 |
| 24/08/2026 | FS.com D7070 800G muxponder confirmed (3.2 Tb/s aggregated) | 6.3 |
| ~25/08/2026 | FS.com H1 2026 results: revenue +25.9% to RMB 1.76B, net profit +65.1% to RMB 451.1M (TipRanks) | 6.5 |
| 2/09/2026 | Credo CEO Bill Brennan, Q1 FY2027 earnings call: >$600M optical FY2027 guidance | 6.11 |
| 7/09/2026 | FS.com AI-networking portfolio release: "Scale-Out" 400G/800G/1.6T + "Scale-Across" 400G/800G coherent to 500 km (BusinessWire) | 6.1 |
| 7/09/2026 | Goldman Sachs raises 1.6T shipment forecasts | 6.6 |
| 9/09/2026 | OIF publishes 1600ZR Implementation Agreement: 1.6T coherent single-wavelength, DWDM single-span to 120 km, 2× 800ZR | 6.7 |
| 9/09/2026 | Ethernet Alliance announces ECOC multi-vendor interop demo, booth #2173 (GlobeNewswire) | 6.8 |
| 11/09/2026 | Photonics Spectra covers the OIF 1600ZR agreement | 6.7 |
| 12/09/2026 | CIOE China: 1.6T ramping to volume ("sold out, capacity being reserved"); 3.2T/6.4T NPO/CPO in sampling | 6.6 |
| 15/09/2026 | Credo announces ZeroFlap 1.6T family (224G/lane, 2xDR4/2xFR4/DR8; 224G DSP + Kfir200 SiPh PIC + PILOT) (BusinessWire) | 6.11 |
| 15/09/2026 | Telxius–Nokia joint release: ICE-X 800G ZR/ZR+ deployment, terrestrial, EU/US/LatAm, IP-over-DWDM + Transcend | 6.12 |
| 15/09/2026 | MACOM announces ECOC booth #1055 demos: 3.2T solutions (448G/lane drivers, TIAs, photodetectors), 1.6T products, 1.6T active copper cable (GlobeNewswire) | 6.10 |
| ~16/09/2026 | Credo reaffirms >$600M optical guidance (TipRanks); LightCounting via EETimes: "1.6T is shipping now" (Nvidia, Google integrating; Meta, Oracle 2027); LPO: 4–5 players incl. Oracle, Google 1.6T linear-receive deploying | 6.11, 6.6, 6.9 |
| 21/09/2026 | FS.com Wi-Fi 7 campus announcement: AP-N755, 24.436 Gb/s; AmpCon-Campus 3.0 "upcoming release" (BusinessWire) | 6.4 |
| 21–23/09/2026 | ECOC 2026, Malaga — ongoing as of 22/09; no show reports published yet | 6.8 |

#### Glossary of the chapter's technical terms

- **1.6T / 800G / 400G**: per-port Ethernet rates in gigabits/terabits per second; 1.6T = 1,600 Gb/s, the generation that moved to volume in 2026.
- **OSFP**: the pluggable module form factor used for 800G/1.6T optics (octal small form-factor pluggable).
- **DR8 / FR4 / 2xDR4 / 2xFR4**: 1.6T module configurations — DR8: 8 short-reach lanes (500 m class); FR4: 4 longer-reach lanes (2 km class); the "2x" variants double the lane count.
- **SiPh (silicon photonics)**: integrating optical functions onto silicon-chip processes, favoring volume economics.
- **EML**: electro-absorption modulated laser, the laser technology behind longer-reach 1.6T modules.
- **DSP**: digital signal processor — the in-module chip doing equalization and error correction; the hungriest component in a retimed transceiver.
- **224G/lane**: the per-lane electrical signaling rate of the 2026 1.6T generation (8 × 200G ≈ 1.6T).
- **448G/lane**: the next per-lane rate, demoed/sampled in 2026 as the building block of 3.2T.
- **LPO (linear pluggable optics)**: DSP removed from the module, equalization on the host ASIC; lower power/latency, pluggable form factor kept.
- **CPO / NPO**: co-packaged / near-packaged optics — moving optics onto or next to the switch package; in sampling for 3.2T/6.4T in 2026.
- **Coherent (ZR/ZR+)**: coherent optical transmission in pluggable form; 800ZR/ZR+ the 2026 carrier DCI standard, 1600ZR the OIF's September 2026 1.6T agreement.
- **IP-over-DWDM**: routers connecting directly over DWDM fiber via coherent pluggables, collapsing the optical-transport layer.
- **DCI**: data-center interconnect — links between data-center sites (the "Scale-Across" domain).
- **Muxponder**: device multiplexing multiple client signals onto coherent line wavelengths (e.g., 4 × 800G = 3.2 Tb/s).
- **RoCE**: RDMA over Converged Ethernet — the NIC technology for low-latency server attach in AI fabrics.
- **DAC**: direct-attach copper cable — cheapest short-reach intra-rack link; active copper adds in-connector signal conditioning.
- **SerDes**: serializer/deserializer — the high-speed electrical lanes in a switch ASIC; "link training" is their equalization negotiation at link-up.
- **LLR / CBFC**: link-level retry / credit-based flow control — mechanisms for lossless Ethernet fabrics in AI clusters.
- **TIA**: transimpedance amplifier — amplifies the photodetector's tiny current on the optical receive path.
- **PIC**: photonic integrated circuit — lasers, modulators, detectors integrated on one chip (e.g., Credo's Kfir200).
- **P802.3dj**: the IEEE project standardizing 1.6 Tb/s Ethernet — progressing in 2026, shipments preceding it.
