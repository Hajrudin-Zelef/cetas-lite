---
id: briefing-general-tech-2026/05-networking-optics/11-credo-zeroflap
title: "Credo ZeroFlap 1.6T"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Credo"]
dates: ["2025-12", "2026-06", "2026-09", "2026-09-02", "2026-09-15", "2027-05"]
keywords: ["dsp", "optics", "packaging", "serdes"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-11"
source_lines: [6323, 6413]
canonical_for: ["credo-zeroflap"]
sha256: 3ef2d2b4fc38ede10a509365790720c9763a8b1e0ff51bd451cf0897ea4a904d
---

# Credo ZeroFlap 1.6T

<a id="g06-11"></a>
### 6.11 Credo ZeroFlap 1.6T

On September 15, 2026 (via BusinessWire), Credo announced **ZeroFlap (ZF)**,
a 1.6T transceiver family built on **224G-per-lane** signaling in three
configurations — **2xDR4, 2xFR4 and DR8**. The family combines three Credo
technology pillars: a **224G optical DSP**, the **Kfir200 silicon-photonics
PIC** (photonic integrated circuit), and the **PILOT diagnostics platform**.
The DR8 variant is the 500 m–class intra-data-center part; the FR4 variants
extend reach to the 2 km class — the same two-reach structure as FS's 1.6T
line in section 6.2, which is no coincidence: DR8/FR4 at 224G/lane is the
industry's 2026 consensus 1.6T form.

The announcement cited **Cignal AI**: the 1.6T ramp was described as "over
11 million forecast to ship in 2026" — a third-party market-research figure
relayed in a vendor press release, which is how it should be read (Cignal's
forecast, cited by Credo, not an independently verified shipment count).

The financial frame around Credo's optical push is official guidance, and the
dossier treats it strictly as such. CEO Bill Brennan stated on the **Q1
FY2027 earnings call (September 2, 2026)** that Credo expected **over $600
million in optical revenue in FY2027** — guidance reaffirmed around September
16 (via TipRanks) — with **three optical families (ZeroFlap Optics, PICs,
DSPs) each expected to exceed $100 million** (via Zacks). Two nuances are
non-negotiable here: first, this is **guidance, not booked results** — a
forward-looking statement about a fiscal year still in progress; second,
Credo runs an **offset fiscal year, so FY2027 ends around May 2027**, meaning
the $600M+ figure spans roughly June 2026–May 2027. Read with those caveats,
the guidance is still the strongest vendor-side quantification of the 1.6T
ramp available in September 2026: a single connectivity-chip company guiding
to more than half a billion dollars of optical revenue on the back of the
224G/lane generation.

#### The 224G-per-lane consensus of 2026

ZeroFlap's three configurations — 2xDR4, 2xFR4, DR8 — are not a Credo
invention; they are the industry's converged answer to "what does a 1.6T
module look like." Eight lanes at 200G (DR8, 500 m class) or four lanes at
400G in FR4 packaging (2 km class), all built on 224-gigabit electrical
lanes: this is the same two-reach structure as FS's December 2025 modules
(section 6.2), because the underlying physics and the OSFP/faceplate
constraints leave little room for differentiation at the form-factor level.
The convergence matters for buyers — a DR8 from Credo and a DR8 from FS
are meant to be interchangeable — and it is what makes the
interoperability demos of section 6.8 meaningful rather than ceremonial.

#### Vertical integration: DSP + PIC + diagnostics

What Credo *does* claim as differentiation is vertical integration across
three layers: the **224G optical DSP** (Credo's historical strength —
high-speed SerDes and DSP silicon), the **Kfir200 silicon-photonics PIC**
(the photonic integrated circuit: lasers, modulators and detectors on one
chip), and the **PILOT diagnostics platform** (telemetry and fault
isolation for operating the modules at scale). The strategic logic is
margin capture: in a transceiver, the value concentrates in the DSP and the
photonics, while final module assembly commoditizes. Owning the DSP and
the PIC lets Credo sell to module makers *and* compete with them — the
three "optical families" of the guidance (ZeroFlap Optics, PICs, DSPs)
map exactly onto the layers of this integration. It is the same playbook
the company ran in the copper/AEC (active electrical cable) market, now
applied to optics.

#### Reading the $600M guidance: what Brennan actually said

The guidance deserves a careful parse because it is the strongest
vendor-side quantification of the 1.6T ramp — and the easiest to misread.
CEO Bill Brennan's statement on the Q1 FY2027 earnings call (September 2,
2026): **over $600 million in optical revenue in FY2027**, reaffirmed
around September 16 (via TipRanks), with each of the three optical
families expected above $100 million (via Zacks). The non-negotiable
nuances:

1. **Guidance is not results.** This is a forward-looking statement about
   a fiscal year still in progress — a commitment of expectation, not a
   report of achievement. It belongs in forecasts, not in shipment tables.
2. **The fiscal year is offset.** Credo's FY2027 ends around May 2027, so
   the $600M+ spans roughly June 2026–May 2027 — it straddles two calendar
   half-years and cannot be compared directly with calendar-2026 market
   figures like Cignal AI's "over 11 million" 1.6T units.
3. **The 3×$100M floor is structural.** Three families each above $100M
   means the guidance does not depend on a single product line hitting —
   it is diversified across finished transceivers, bare PICs and merchant
   DSPs, which is also a hedge against the module-assembly
   commoditization noted above.

Read with those caveats, the number still says something real: a
connectivity-chip company is willing to guide publicly to more than half
a billion dollars of optical revenue on the 224G/lane generation. In
September 2026, no other single vendor had put a number that large on the
1.6T ramp.

