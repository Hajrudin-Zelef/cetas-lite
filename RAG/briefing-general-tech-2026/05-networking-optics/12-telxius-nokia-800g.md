---
id: briefing-general-tech-2026/05-networking-optics/12-telxius-nokia-800g
title: "Telxius deploys Nokia 800G coherent pluggables"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Nokia", "OIF", "Telxius"]
dates: ["2026-09", "2026-09-15"]
keywords: ["dci", "dsp", "optics", "wavelength"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-12"
source_lines: [6414, 6486]
sha256: 4e6a4b7571330fd97a29018f9b3a1d19f46cdf9f267e925c8452fa708da9a70e
---

# Telxius deploys Nokia 800G coherent pluggables

<a id="g06-12"></a>
### 6.12 Telxius deploys Nokia 800G coherent pluggables

On September 15, 2026, **Telxius and Nokia issued a joint press release**
announcing the deployment of **Nokia ICE-X 800G ZR/ZR+ coherent pluggable
optics** on Telxius's terrestrial networks across **Europe, the United
States and Latin America**. The architecture is **IP-over-DWDM** — 800G
coherent pluggables directly in routers, collapsing the optical-transport
layer — managed with **Nokia's Transcend automation** platform.

The deployment follows a record-setting proof point: a demonstration on the
**BRUSA submarine cable** pushing **400 Gb/s per wavelength over more than
5,600 km** — the kind of long-haul validation that de-risks the terrestrial
rollout, since the same coherent DSP technology that survives a transatlantic
submarine span has ample margin on terrestrial links.

The deployment nuance is the honest one: the press release announces an
**ongoing deployment — no volume figures and no completion date were
published.** The reader should therefore treat this as a confirmed strategic
deployment in progress, not as a finished network. Its significance in this
chapter is as the carrier-side mirror of the data-center story: while the
hyperscalers were ramping 1.6T inside their AI fabrics (6.6), the long-haul
carriers were standardizing on 800G coherent pluggables to move that traffic
between continents and regions — the same IP-over-DWDM collapse, one
generation behind the data-center leading edge.

#### IP-over-DWDM, now at 800G per wavelength

The architecture keyword is **IP-over-DWDM**: 800G ZR/ZR+ coherent
pluggables sitting directly in router faceplates, no separate optical
transport layer. What was exotic a decade ago is now the default
procurement model for DCI — the coherent DSP has shrunk far enough to
live in a pluggable, and the economics of eliminating transponder shelves
are overwhelming. Nokia's **ICE-X** is the pluggable family; **Transcend**
is the automation layer that provisions and monitors the optical paths,
because a collapsed architecture still needs someone — or something — to
manage wavelengths, and at 800G per lambda the operational tooling matters
as much as the optics. The three-region footprint (Europe, US, Latin
America) is the point of the announcement as much as the technology:
Telxius is a wholesale carrier, and a wholesale carrier standardizing on
one coherent pluggable family across three continents is a volume signal
for the 800G coherent supply chain.

#### BRUSA: the 5,600 km proof point

The submarine-cable demonstration on **BRUSA** — 400 Gb/s per wavelength
over more than 5,600 km — functions as the technical de-risking for the
terrestrial deployment. Submarine spans are the hardest coherent
environment (thousands of kilometers, no mid-span access, accumulated
noise); a DSP and optics combination that closes a 5,600 km submarine link
at 400G per wavelength has enormous margin on a terrestrial span of a few
hundred kilometers. The record framing ("record demo") is vendor language,
but the underlying logic is sound engineering communication: prove it on
the hardest path, then deploy it on the easier ones. It also places the
Telxius–Nokia announcement in the coherent roadmap: 400G per wavelength
submarine-proven in 2026, 800G ZR/ZR+ deploying terrestrially the same
year, 1.6T coherent (1600ZR, section 6.7) standardizing for the next step.

#### A deployment in progress, not a completed network

The honest scope note: the September 15 release announces a deployment
that is **ongoing — no volume figures, no completion date, no
per-region rollout schedule were published.** The reader should therefore
file this under "confirmed strategic deployment in progress" and resist
turning it into a finished-network fact. That is not a criticism of the
announcement — carriers essentially never publish deployment volumes for
competitive reasons — but it bounds what the dossier can claim. What *is*
claimable, and significant: a major wholesale carrier chose the
IP-over-DWDM 800G-pluggable architecture as its terrestrial standard in
September 2026, which is the carrier-side ratification of the same
architectural collapse the data-center world completed a generation
earlier.

