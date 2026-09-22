---
id: briefing-general-tech-2026/05-networking-optics/02-16t-osfp-verified-nvidia
title: "1.6T OSFP and the 'verified on NVIDIA' claim"
domain: networking-optics
role: deep-dive
task: hardware
actors: ["Broadcom", "FS.com", "Nvidia", "OIF"]
dates: ["2025-12-22", "2026-09", "2026-09-07"]
keywords: ["asic", "dsp", "ethernet", "gpus", "optics"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-2"
source_lines: [5687, 5772]
sha256: 1206d971a9ca0e14c0b439580c7f61665c57e4b5c4249114dc6c6f07148d4c03
---

# 1.6T OSFP and the 'verified on NVIDIA' claim

<a id="g06-2"></a>
### 6.2 1.6T OSFP and the "verified on NVIDIA" claim

FS.com's December 22, 2025 press release (via BusinessWire) announced two
1.6T OSFP modules: the **OSFP-DR8-1.6T**, built on silicon photonics (SiPh)
with reach up to 500 m, and the **OSFP-2FR4-1.6T**, built on electro-absorption
modulated lasers (EML) with reach up to 2 km. Both use a Broadcom 3 nm DSP and
run at 25–26 W — a power envelope that matters, because 1.6T modules pushed
into racks must stay within the thermal budget of existing switch line cards.
The claimed application was explicit: links between Nvidia Quantum-X800
switches and Nvidia B300/GB300 GPUs running XDR InfiniBand.

The same release described the modules as "100% verified on NVIDIA
Quantum-X800 switches." This formulation needs a correction, because the
wording invites a misreading. What exists in the record is an **FS self-claim**,
dated December 22, 2025: no Nvidia co-announcement, no independent
certification, and no third-party test report were found during verification.
"Verified on" here means tested by FS against Nvidia switches — it is not
certification *by* Nvidia, and it should be read as a vendor claim, not as an
endorsement. Notably, FS's own September 7, 2026 release — the one that lays
out the Scale-Out / Scale-Across portfolio — makes no mention of any Nvidia
validation. The correct way to present this is therefore: **"FS claim
(22/12/2025)"**, with the caveat attached.

The episode illustrates a pattern that recurs throughout this chapter: in a
market where 1.6T is ramping faster than standards bodies and certification
programs can follow, vendor press releases outpace independently verifiable
interoperability claims. The real interoperability evidence for 1.6T Ethernet
in September 2026 comes instead from the Ethernet Alliance multi-vendor demo at
ECOC (section 6.8) and the OIF 1600ZR agreement (section 6.7) — processes where
multiple vendors demonstrate interworking rather than one vendor asserting
compatibility.

#### Anatomy of the two modules

| Module | Optical technology | Reach | DSP | Power |
|---|---|---|---|---|
| OSFP-DR8-1.6T | Silicon photonics (SiPh) | ≤500 m | Broadcom 3 nm | 25–26 W |
| OSFP-2FR4-1.6T | EML (electro-absorption modulated laser) | ≤2 km | Broadcom 3 nm | 25–26 W |

The pairing is the industry's standard two-reach 1.6T structure: a
silicon-photonics DR8 variant for the 500-meter class (intra-data-center,
the bulk of AI-fabric links) and an EML-based 2xFR4 variant for the 2 km
class (campus and metro-edge spans). Silicon photonics integrates the
optical functions onto a chip process, favoring volume economics; EML
pairs a laser with an electro-absorption modulator for longer-reach,
higher-extinction links. Both share the same Broadcom 3 nm DSP — the
digital engine that does the heavy lifting of equalization and forward
error correction at 1.6T line rates.

#### Why 25–26 W is a headline figure

A 1.6T module's power draw is not a footnote; it is a deployment
constraint. Switch line cards are designed around a per-port thermal and
power budget, and the move from 800G to 1.6T roughly doubled the heat each
faceplate port must dissipate. At 25–26 W, FS's modules sit in the envelope
that existing OSFP switch designs can absorb without a chassis redesign —
which is precisely why the figure leads the press release. The power story
is also the thread that connects this section to section 6.9: the DSP is
the hungriest component in the module, and the industry's answer — linear
pluggable optics, moving equalization to the host ASIC — exists largely to
push module power down from exactly this level.

#### Reading "verified on": a method note for the whole chapter

The "100% verified on NVIDIA Quantum-X800 switches" claim deserves to be
treated as this chapter's worked example of source hygiene, because the
same pattern — a true-sounding compatibility statement with no named
counterparty — recurs across vendor marketing in a ramp year. The checklist
applied during verification:

1. **Is there a co-announcement?** No Nvidia statement accompanied the FS
   release, and none was found in the verification window.
2. **Is there an independent certification?** No third-party test report or
   certification program result was found.
3. **Does the vendor repeat the claim?** FS's own September 7, 2026
   portfolio release — nine months later — mentions no Nvidia validation.

A claim that fails all three checks is a self-claim, full stop. That does
not make it false — FS may well have tested the modules against Quantum-X800
switches exactly as stated — but it makes it *unverifiable by the reader*,
and this dossier's rule is that unverifiable compatibility is labeled, not
repeated. The contrast case is the Ethernet Alliance multi-vendor demo at
ECOC (6.8) and the OIF 1600ZR process (6.7): interoperability demonstrated
by several vendors together, on a show floor, against each other's gear.

