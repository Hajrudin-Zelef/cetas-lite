---
id: briefing-general-tech-2026/05-networking-optics/09-linear-pluggable-optics
title: "Linear Pluggable Optics (LPO)"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Google"]
dates: ["2026-09"]
keywords: ["lpo", "optics", "asic", "cpo", "dsp", "ethernet", "hyperscaler", "npo"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-9"
source_lines: [6167, 6249]
canonical_for: ["lpo-cpo-npo"]
sha256: 86040dabf8fd2a9f585d78b8c42c2e8593da3e035332c32a66c66db05f4c8a95
---

# Linear Pluggable Optics (LPO)

<a id="g06-9"></a>
### 6.9 Linear Pluggable Optics (LPO)

Linear Pluggable Optics is the architecture bet that 2026 made concrete.
The principle is simple to state: **remove the DSP from the optical module
and move equalization onto the host ASIC** (the switch or NIC chip). The DSP
is the hungriest component in a conventional retimed transceiver; pushing its
function to the host trades module complexity for lower module power and
lower latency, at the price of tighter co-design between the module and the
switch ASIC it plugs into.

By September 2026 the technology had crossed from PowerPoint into purchase
orders. EETimes/LightCounting reporting around September 16 placed **4–5
players already using LPO, including Oracle**, and reported **Google deploying
1.6T linear-receive** optics, with **volume rollouts in 2026** and
acceleration of both LPO and CPO (co-packaged optics) expected through
2026–2027. Google's involvement is the tell: when a hyperscaler with
in-house switch-ASIC design commits to linear optics, it validates the
co-design model that LPO depends on — Google controls both ends of the link,
which is exactly what linear optics needs to be robust.

LPO and CPO are related but distinct answers to the same power problem. LPO
keeps the pluggable form factor (operational familiarity, field
replaceability) and linearizes the module; CPO moves the optics onto the
switch package itself (maximum power and density savings, maximum
operational disruption). The 2026 market verdict, per the CIOE and
LightCounting signals: LPO is shipping now at 1.6T, CPO and near-packaged
optics are in sampling for the 3.2T/6.4T generation. Both trajectories are
worth tracking into 2027, because the power-per-bit curve — not the
bits-per-second curve — is now the binding constraint on AI-fabric scaling.

#### The power-per-bit problem, stated plainly

Every generation of Ethernet doubles the bits and roughly doubles the
module power — until the power breaks the system around it. A 1.6T retimed
module at 25–26 W (section 6.2) is at the edge of what an OSFP faceplate
and its switch cooling can handle; the next doubling to 3.2T in a retimed
pluggable would not fit. The industry therefore has to remove power from
the module, and the DSP — the digital signal processor doing equalization
and error correction — is the single hungriest component to remove. LPO's
bargain is exactly that: take the DSP out of the module, do the
equalization in the host switch ASIC (which has the thermal budget and the
process node to do it efficiently), and accept a "linear" optical path
through the module. The payoff is lower module power and lower latency
(no DSP in the data path means no DSP delay); the price is that the module
and the host ASIC must be designed — or at least qualified — together.

#### Three answers to the same problem: LPO, CPO, NPO

The 2026 market did not converge on one answer; it is pursuing three, at
different stages:

| Approach | What changes | Status in 9/2026 |
|---|---|---|
| **LPO** (linear pluggable optics) | DSP removed from module, equalization on host ASIC; pluggable form factor kept | Shipping at 1.6T (Google linear-receive deploying; 4–5 players incl. Oracle) |
| **CPO** (co-packaged optics) | Optics moved onto the switch package itself | In sampling for 3.2T/6.4T generation |
| **NPO** (near-packaged optics) | Optics placed adjacent to the switch ASIC, not quite co-packaged | In sampling for 3.2T/6.4T generation |

LPO keeps operational familiarity — field-replaceable modules, existing
faceplates — and is therefore the least disruptive path; CPO maximizes the
power and density savings and maximizes the operational disruption (a
failed optic means a failed switch package). The CIOE September signal
places CPO/NPO at the 3.2T/6.4T sampling stage while LPO ships at 1.6T —
which reads as the industry's sequencing decision: linearize first, then
co-package.

#### Google's linear-receive bet and the co-design bargain

The most significant LPO data point of September 2026 is Google deploying
**1.6T linear-receive** optics. "Linear-receive" specifies which half of
the link went linear first — the receive path, where equalization is the
harder problem — and Google's involvement validates the co-design model
LPO requires. A hyperscaler that designs its own switch ASICs controls both
ends of the link: it can qualify the host equalization against known
linear modules, which is precisely the tight coupling that makes LPO
robust. Merchant-silicon buyers cannot do this as easily, which is why LPO
adoption in 2026 is led by the hyperscalers (Google deploying, Oracle
among the 4–5 players already using it) while the broader market watches.
The EETimes/LightCounting expectation — volume rollouts through 2026,
LPO/CPO acceleration in 2026–2027 — makes the next twelve months the
proving ground for whether linear optics stays a hyperscaler specialty or
becomes the industry default.

