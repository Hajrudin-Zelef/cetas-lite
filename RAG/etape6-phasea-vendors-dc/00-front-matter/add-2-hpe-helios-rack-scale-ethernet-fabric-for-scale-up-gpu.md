---
id: etape6-phasea-vendors-dc/00-front-matter/add-2-hpe-helios-rack-scale-ethernet-fabric-for-scale-up-gpu
title: "ADD-2. HPE \"Helios\" rack-scale Ethernet fabric for scale-up GPU (distinct from AMD Helios) [secondary]"
domain: front-matter
role: reference
task: hardware
actors: ["AMD", "Broadcom"]
dates: ["2026-09-22"]
keywords: ["amd", "ethernet", "gpu", "helios", "rack-scale", "acquisition", "agentic", "compute", "gpus", "inference", "mi455x", "optics"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1791, 1822]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: af26fd62725b7e66189e7d711741dd9792c2caf46db20ee8247f0ce84cf36b25
---

# ADD-2. HPE "Helios" rack-scale Ethernet fabric for scale-up GPU (distinct from AMD Helios) [secondary]

- HPE (post-Juniper-acquisition) announced a new **QFX5140 network switch**, purpose-built for the next wave of AI infrastructure with a focus on **inferencing clusters at the edge** — HPE is seeing "explosive growth and demand" for AI inferencing, per Rami Rahim, President of HPE Networking [secondary — CRN, 2026 reporting; quote verbatim from article].
- Hardware profile reported: **1RU, 16 Tbps**, built on **Broadcom Trident 5** switching silicon; includes **AI load balancing and congestion control**; supports speeds **from 25 Gbps to 800 Gbps** per port [secondary — CRN].
- Part of a broader announcement bringing Juniper network integration as a **full-stack validated solution for HPE AI Factory** deployments: "This includes the switches themselves managed through HPE Networking Data Center Director, creating a fully integrated full-stack solution" (Rami Rahim) [secondary — CRN].
- Juniper QFX switches named as supported in the validated stack: **QFX 5230, 5240, 5250**, with self-driving network operations powered by the **HPE Mist AI Networking Data Center Assurance** platform [secondary — CRN].
- AI-data-center ops additions in the same announcement: "predictive analytics" capabilities for DC operations — Mist AI / Marvis AI engine proactively analyzing telemetry across AI infrastructure including **power, temperature, optics and system health** to "identify potential issues before they become outages"; HPE claims AI/ML prediction of system and optics failures "with a high confidence level well before they occur" [secondary — CRN; vendor-reported per Rahim quotes].
- Relation to this file's HPE/Aruba coverage: the QFX5140 sits alongside the Juniper QFX5230/5240/5250 line already referenced in earlier passes, but is the **first inference-edge-specific, Trident-5-based SKU** in the post-acquisition HPE Networking DC portfolio; managed under HPE Networking Data Center Director rather than AOS-CX fabric tooling.

### ADD-2. HPE "Helios" rack-scale Ethernet fabric for scale-up GPU (distinct from AMD Helios) [secondary]

- In a 2026 interview at HPE Discover (Las Vegas), **Rami Rahim** (EVP, President and GM of HPE Networking, formerly Juniper CEO) described the **"Helios" rack-scale Ethernet fabric** — an **open Ethernet solution for scale-up GPU connectivity** — as a **first-of-its-kind architecture enabled specifically by combining HPE and Juniper capabilities** [secondary — sixfivemedia interview with Daniel Newman; Rahim's characterization].
- Positioning given: networking is **10–15% of AI data center spend** and determines throughput efficiency of GPU infrastructure; running at 50% capacity due to blocking/congestion erodes returns on the entire stack — networking as "force multiplier on compute investment, not a commodity line item" [secondary — sixfivemedia].
- HPE–Juniper integration status per Rahim: **ahead of schedule and already producing new product surface area**; the Helios fabric is cited as unlocked by the close of the acquisition [secondary — sixfivemedia].
- Marvis AI automation claims in the same interview: Marvis has ~a decade of real-world telemetry and **currently resolves 70–80% of network issues through closed-loop automation**; Rahim projects agentic-AI advances bringing that to 100% "within a few years", with HPE intending to reach it before competitors [secondary — sixfivemedia].
- Security integration strategy: spans firewall, SD-WAN, NAC, cloud security, converging toward **universal zero trust network access**; enterprises moving from requiring approval before remediation to **authorizing Marvis to resolve issues without prior notification** [secondary — sixfivemedia].
- ⚠️ **Disambiguation:** this "Helios" (HPE Networking scale-up Ethernet fabric, HPE–Juniper) is a different product from **AMD's "Helios" rack-scale AI architecture** (72× MI455X GPUs, UALink-over-Ethernet, first-OEM-commit HPE Helios AI Racks 2026) already documented in this file's earlier round-4 pass. Treat as separate entities; note the potential confusion.
- Status as of 2026-09-22: public technical specifications of the HPE Networking Helios fabric (bandwidth, radix, switch silicon, topology, availability) were **not located** — mark [unverified] until an official HPE datasheet or press release is found. The CRN piece (ADD-1) and sixfivemedia interview (ADD-2) are the only sources surfaced; both are reputable but secondary.

### ADD-3. Addendum verification log (open items)

1. **QFX5140 datasheet details** — exact port configuration, optics, power, AOS/Junos OS version, pricing, and availability: not located in the CRN reporting; seek official HPE/Juniper datasheet [gap — needs official source].
2. **HPE Networking "Helios" fabric specs** — silicon, bandwidth per lane, scale-up protocol detail, GA date: no official source located as of 2026-09-22 [unverified — needs official source].
3. **Dell SN6600-LD (Spectrum-6) GA/pricing** — still not confirmed by any source located across all round-4 passes [unverified — carried forward].

### ADD-4. Addendum sources (verbatim URLs)

- https://www.crn.com/news/networking/2026/hpe-networking-president-rami-rahim-on-latest-self-driving-network-innovation-and-why-hpe-remains-years-ahead-of-competitors
- https://www.sixfivemedia.com/content/networking-becomes-the-bottleneck-rami-rahim-on-building-infrastructure-for-ai-at-scale

**Addendum collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; open items are listed in §ADD-3. Earlier sections (§§1–8 and all prior supplementary passes) were not modified.

---

