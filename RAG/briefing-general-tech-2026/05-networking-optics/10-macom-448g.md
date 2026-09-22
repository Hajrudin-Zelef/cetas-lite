---
id: briefing-general-tech-2026/05-networking-optics/10-macom-448g
title: "MACOM: 448G/lane toward 3.2T"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Credo", "MACOM"]
dates: ["2026-09", "2026-09-15"]
keywords: ["cpo", "lpo", "npo", "optics"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-10"
source_lines: [6250, 6322]
canonical_for: ["448g-lane"]
sha256: ec9aab20aa98f7f6b2b9cc2199d58435dd00a0a6e33e9c7fcfb6589d6034f3cd
---

# MACOM: 448G/lane toward 3.2T

<a id="g06-10"></a>
### 6.10 MACOM: 448G/lane toward 3.2T

On September 15, 2026 (via GlobeNewswire), MACOM announced the **live demos
it would stage at ECOC 2026, booth #1055** — including what the release
called "3.2T Optical Solutions," featuring **448 Gbps-per-lane modulator
drivers, transimpedance amplifiers (TIAs) and photodetectors**. The arithmetic
is the roadmap: 8 lanes × 448 Gb/s = 3.2T (3.584 Tb/s raw). The release also
covered MACOM's 1.6T products — built on the current 224G-per-lane generation —
and a **1.6T active copper cable** offering.

Two nuances matter. First, the 448G-per-lane components are the *enabling
building blocks* of a 3.2T module, not a 3.2T module itself: drivers, TIAs and
photodetectors are what transceiver vendors will integrate in the next
generation. Second — and this is a scope warning for the reader — the
September 15 release **announces the demos**; with ECOC ongoing as of
September 22, no show report on how the demos went had been published yet.
What is verified is the announcement and its technical content, not the
demo outcomes.

Placed against the CIOE signal from three days earlier (3.2T/6.4T NPO/CPO in
sampling), MACOM's announcement confirms the pipeline: the 224G-per-lane
generation is the shipping product of 2026 (Credo's ZeroFlap, section 6.11,
is built on it), while the 448G-per-lane generation is the demoed,
sampled, next step. The active-copper-cable line is worth a footnote too: at
1.6T, copper's reach collapses, so an *active* copper cable — with signal
conditioning in the connector — is how copper keeps a foothold in the
shortest intra-rack links against DACs and linear optics.

#### 448G per lane: the arithmetic of the next doubling

The "3.2T" in MACOM's demo title is arithmetic, and it is worth showing
the work: **8 lanes × 448 Gb/s = 3.584 Tb/s raw**, marketed as 3.2T after
coding overhead — the same convention that made 8 × 224G into "1.6T."
The per-lane doubling (224G → 448G) is the fundamental unit of progress in
optical networking: every generation since 100G has been built by doubling
the lane rate and multiplying the lane count. What makes 448G notable is
that it is the first lane rate the industry is pursuing *while* the power
problem (section 6.9) is binding — which is why the demos pair the
448G-per-lane analog components with the LPO/CPO architectural discussion
happening around them. A 448G lane in a retimed pluggable would be a
thermal impossibility; in a linear or co-packaged architecture, it is a
product plan.

#### The component layer beneath the module vendors

MACOM's position in the stack explains the announcement's shape. The
company does not sell transceivers; it sells the analog components inside
them — modulator drivers (which drive the optical modulator at 448G lane
rates), transimpedance amplifiers or TIAs (which amplify the tiny
photocurrent on the receive side), and photodetectors. Announcing these as
ECOC demos is a message to the *transceiver vendors* in the audience: the
building blocks for your 3.2T modules exist and are demonstrable. This is
the merchant-silicon layer of the optical supply chain, and its health is
a leading indicator — component demos in September 2026 mean module
sampling in 2027 and volume some quarters after. The release's parallel
coverage of 1.6T products (224G/lane) serves the other half of the
audience: the buyers placing this year's 1.6T orders.

#### 1.6T active copper: copper's last stand in the rack

The 1.6T active copper cable in the MACOM announcement deserves its own
note because it marks a boundary. Passive DACs — the catalog items from
section 6.1 — work at 1.6T only over very short reaches; as lane rates
rise, copper's usable distance collapses. An *active* copper cable puts
signal conditioning (retiming or equalization) into the connector,
extending copper's reach enough to keep the shortest intra-rack links on
copper rather than fiber. It is a cost play — copper remains cheaper than
optics per link — but a bounded one: each generation shrinks the distance
over which it works, and linear optics is coming down-market to contest
exactly these links. The active copper cable is therefore best read as a
bridge product for the 1.6T generation, not as copper's long-term answer.

