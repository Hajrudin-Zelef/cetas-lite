---
id: briefing-general-tech-2026/05-networking-optics/01-fscom-portfolio
title: "FS.com: the AI networking portfolio"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Broadcom", "FS.com", "Nvidia"]
dates: ["2025-12-22", "2026-09", "2026-09-07"]
keywords: ["dci", "gpu", "optics"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-1"
source_lines: [5605, 5686]
sha256: ff43c74d11a1ffc490661f9257b9cc171eb9b5c7158b8e2041ae1ad1389ecf01
---

# FS.com: the AI networking portfolio

<a id="g06-1"></a>
### 6.1 FS.com: the AI networking portfolio

FS.com is a network-hardware vendor that built its reputation on third-party
optical transceivers and now positions itself as a one-stop supplier for AI
data centers. In a press release dated September 7, 2026 (via BusinessWire),
the company laid out an optical portfolio organized around two domains with
FS's own terminology: **"Scale-Out"** — 400G, 800G and 1.6T transceivers for
GPU-cluster fabrics inside the data center — and **"Scale-Across"** — 400G and
800G coherent transceivers for data-center interconnect (DCI), with reach up
to 500 km. The "Scale-Out / Scale-Across" vocabulary is FS.com's, not an
industry standard, and the September 7 release is the confirmed source for
this framing.

The 1.6T leg of the portfolio had been introduced earlier: FS launched its
1.6T OSFP transceiver line on December 22, 2025, planting its flag well ahead
of the 2026 volume ramp. The catalog around the optics fills in the rest of
the fabric stack: 400G RoCE network interface cards — catalog listings include
Nvidia's ConnectX-7 and ConnectX-8 as well as the Broadcom BCM57608 — and a
direct-attach copper (DAC) cable range for short-reach intra-rack links, where
copper remains cheaper than optics below a few meters.

On the commercial side, FS reported strong sales growth in high-speed
(100G-and-above) optics tied to AI-cluster demand. One figure is deliberately
absent from this dossier: no public breakdown of the 100G+ segment exists, so
the exact growth rate is unverifiable and is not cited here. What is confirmed
is the direction — AI clusters pulling the high-rate portfolio — corroborated
by the company's overall first-half results, which are detailed in section
6.5.

| Portfolio segment (FS terminology) | Rates | Reach / role |
|---|---|---|
| Scale-Out (GPU cluster fabric) | 400G / 800G / 1.6T | Intra-data-center, ≤2 km |
| Scale-Across (DCI) | 400G / 800G coherent | Up to 500 km |
| NICs | 400G RoCE (ConnectX-7/8, BCM57608) | Server attach |
| DACs | Direct-attach copper | Short-reach, intra-rack |

#### From compatible optics to AI-networking vendor

FS.com's trajectory is the optical industry's 2026 story in miniature. A
vendor that built its reputation on third-party-compatible transceivers —
the aftermarket alternative to OEM-branded optics — used the AI-cluster
build-out to climb the stack: from selling cheaper equivalents of other
people's modules to marketing a named, structured portfolio for GPU fabrics.
The September 7, 2026 release is the codification of that climb. It is a
marketing document, and it should be read as one — but it is also a
revealing map of where the money is: the release organizes everything around
the two traffic domains that AI created, inside the cluster and between
data centers, because those are the domains where 2026 demand outran supply.

#### What "Scale-Out" and "Scale-Across" describe

Strip away the branding and the two terms map onto the two fundamental
traffic patterns of the AI era. "Scale-Out" is the east-west fabric *inside*
a GPU cluster: thousands of accelerators exchanging gradients and activations
at full bisection bandwidth, where the link budget is dominated by density,
power per port and cost — hence 400G/800G/1.6T short-reach optics (≤2 km).
"Scale-Across" is the traffic *between* data-center sites: coherent 400G/800G
wavelengths carrying aggregated fabric traffic across metro and regional
distances up to 500 km, where the link budget is dominated by reach and
spectral efficiency. The 500 km figure matters because it covers the
regional DCI use case — the second data-center hall across the metro area,
the disaster-recovery site in the next region — without requiring
regeneration. The NICs and DACs in the catalog are the connective tissue at
the bottom of the stack: 400G RoCE cards to attach servers, copper DACs for
the few meters inside a rack where even the cheapest optics lose on price.

#### The number this dossier refuses to cite

FS reported strong sales growth in its high-speed (100G-and-above) optics,
tied explicitly to AI-cluster demand — and this dossier deliberately cites
no growth figure for that segment. The reason is methodological: no public
breakdown of the 100G+ segment exists, so any number would be either a
vendor talking point repeated without a source or an analyst estimate
presented as fact. The anti-fabrication rule of this dossier treats both as
unacceptable. What can be said, and is verified: the direction of travel is
confirmed by the portfolio emphasis (three 1.6T-centered releases in nine
months), by the September 2026 industry signals on 1.6T volume (section
6.6), and by the company's own half-year P&L (section 6.5), where profit
growth far outpaced revenue growth — the fingerprint of a mix shift toward
premium high-rate parts.

