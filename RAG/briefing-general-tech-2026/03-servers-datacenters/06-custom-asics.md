---
id: briefing-general-tech-2026/03-servers-datacenters/06-custom-asics
title: "Custom ASICs gain ground"
domain: servers-datacenters
role: deep-dive
task: infrastructure
actors: ["Amazon", "Anthropic", "Broadcom", "Google", "Meta", "Microsoft", "Nvidia", "TSMC", "UALink"]
dates: ["2026-09"]
keywords: ["asic", "accelerator", "capex", "custom silicon", "foundry", "gpu", "gpus", "hbm", "hyperscaler", "inference", "maia", "mtia"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04-6"
source_lines: [3631, 3762]
canonical_for: ["custom-asics"]
sha256: f04ceaf4d3c510fb0919266883f53161e55b9da95461cf134496d3c2bd953496
---

# Custom ASICs gain ground

<a id="g04-6"></a>
### 4.6 Custom ASICs gain ground

The most strategic response to the economics of 2026 — expensive merchant
accelerators, scarce memory, three-quarters of a trillion dollars in capex
chasing both — was to design the silicon yourself. Every hyperscaler with the
scale to justify the fixed cost pushed custom AI ASICs harder in 2026, and the
year brought concrete production milestones, not just roadmaps and renders.
Custom silicon moved, in 2026, from strategic intention to shipping reality.

**Meta — "Iris" (4th-generation MTIA).** Meta's fourth-generation Meta Training
and Inference Accelerator, developed with Broadcom and manufactured by TSMC,
entered **mass production in September 2026**. The reported target frames the
ambition in the industry's new unit of account: supporting **14 GW of deployed
capacity in 2027** — gigawatts, not chips or servers, because at hyperscaler
scale the fleet is planned the way utilities plan generation capacity. Iris is
Meta's bid to own its inference economics outright. The logic is brutal and
simple: at Meta's scale — billions of users, inference workloads running
continuously — the margin between a merchant GPU and a custom ASIC, multiplied
across gigawatts of deployment, is measured in billions of dollars per year.
Mass production starting in September 2026 means Iris-powered capacity comes
online through 2027, exactly when the memory crisis (4.4) is expected to still
be binding — and a custom chip can be co-designed around the memory you can
actually procure, rather than around a merchant GPU's fixed HBM configuration.

**Google — TPU v7 "Ironwood".** Google's seventh-generation TPU is the reference
case for the entire custom-silicon strategy — the program that proved, over a
decade of iteration, that a hyperscaler could design competitive AI accelerators
in-house and deploy them at a scale no other custom program has matched.
Ironwood continued that line in 2026: full-stack co-design with Google's
software (JAX, XLA, the TPU software stack), deployment across Google Cloud as
both internal capacity and customer-facing product, and the institutional
knowledge of seven generations informing the eighth. Where Meta's Iris is the
aggressive newcomer ramping into mass production, Ironwood is the incumbent
custom architecture — the existence proof the others are chasing.

**Amazon — Trainium3 ("Project Rainier").** Amazon's third-generation Trainium,
developed under the "Project Rainier" effort widely associated with Anthropic's
compute needs. The association matters: custom silicon is increasingly built not
just for the hyperscaler's own first-party workloads but as the physical
substrate of its largest AI partnerships. Trainium3's volumes are anchored, in
significant part, by a single customer's demand — which de-risks the enormous
fixed cost of a leading-edge ASIC program and illustrates the new structure of
the industry, in which model companies and cloud companies co-design the
hardware layer. (The Anthropic–Amazon relationship itself is covered in the AI
dossier; here it matters as the demand anchor that makes Trainium3's economics
work.)

**Microsoft — Maia 200 ("Braga").** Microsoft's second-generation Maia
accelerator, codenamed Braga, continuing the company's push — begun with the
first Maia generation — to bring a meaningful share of Azure's AI workloads onto
first-party silicon. Microsoft's position is distinctive among the four: Azure
sells GPU capacity to the broadest enterprise customer base in the industry, so
Maia's success is measured not only in internal cost savings but in whether
enterprise customers will run production AI workloads on non-Nvidia silicon
inside Azure. Braga is the generation where that question starts to get a
production-scale answer.

The strategic logic is uniform across all four programs, and it sharpened
considerably in 2026. At hyperscaler scale, the unit economics of a custom ASIC
outweigh the enormous fixed cost of design along three axes. *Cost*: lower per-
chip cost at volume, no merchant margin, pricing you control. *Efficiency*:
workload-specific architectures that do more useful work per watt and per memory
byte than general-purpose accelerators — a decisive advantage when memory is the
binding constraint (4.3) and power delivery is being rebuilt around 660 kW racks
(4.9). *Sovereignty*: freedom from a single supplier's pricing, allocation
decisions and roadmap — the strategic independence that the UALink coalition
(4.10) is pursuing at the interconnect layer, pursued here at the silicon layer.
The memory crisis did not create this logic, but it poured fuel on it: every
quarter of 3.5× memory pricing is a quarter in which the ASIC's efficiency
advantage pays for more of its development cost.

One important nuance, stated with emphasis because the number circulates freely
and is easy to misread as fact: the claim that **custom ASICs would exceed 50%
of AI accelerator shipments** is a **J.P. Morgan projection for 2027**, not a
measured fact about 2026. It is a reasonable projection — the production ramps
documented above (Iris in mass production from September 2026, Trainium3 and
Maia 200 in the pipeline, TPU v7 deploying) all point toward 2027 as the year
custom silicon reaches real volume — but it remains a projection, subject to the
usual risks of ramps, yields and software readiness. In 2026 itself, merchant
GPUs (Nvidia overwhelmingly) still dominated deployed accelerator value, as the
server-market figures in 4.2 imply: $87.4 billion in GPU-accelerated server
value in a single quarter did not flow to custom ASICs. The custom-ASIC story of
2026 is a story of **production milestones and 2027 ramps**, not of a crown
already taken. The dossier will be able to report the outcome — whether J.P.
Morgan's 50% materialized — only in a future edition.
**The software half of the custom-silicon bet.** A persistent question shadows
every custom ASIC program in this section, and 2026 did not resolve it: the
software moat. Merchant accelerators ship with a mature, decade-deep software
stack — compilers, libraries, profilers, and a global workforce trained on
it — that makes raw silicon performance actually usable. Custom ASICs must
replicate enough of that stack to run real workloads at competitive
efficiency, and the history of AI chips is littered with excellent silicon
stranded by inadequate software. Each of the four programs addresses this
differently: Google's TPU has the deepest stack (a decade of XLA/JAX
co-design); Meta's MTIA benefits from PyTorch lineage and inference-focused
scopes that are easier to compile well; Amazon and Microsoft lean on
partnership models (Anthropic's workloads; Azure's enterprise developers) to
guarantee software investment follows the silicon. None of this is a 2026
measurement — it is structural context, offered because any assessment of the
custom-ASIC ramp that counts only chips is counting only half the bet. The
J.P. Morgan 50% projection for 2027 (4.6) implicitly assumes the software
keeps pace; if it does not, the silicon milestones above will under-deliver
regardless of yields.

**Why 2026 favored the custom strategy anyway.** Against the software risk,
three 2026-specific forces pushed toward custom silicon — stated here as
analysis, not as vendor claims. First, *allocation*: merchant GPU supply was
rationed (4.4's logic applied to accelerators as well as memory), and a
custom chip on a TSMC allocation you control beats a merchant chip on an
allocation you queue for. Second, *memory co-design*: with memory the binding
constraint, an ASIC architected around obtainable memory configurations
extracts more useful work per scarce byte than a fixed-configuration merchant
part. Third, *the efficiency mandate*: 660 kW racks (4.9) and power-limited
datacenter campuses make performance-per-watt a first-order procurement
criterion, and workload-specific silicon wins that metric by construction.
2026 was, in short, the year the environment rewarded exactly the properties
custom ASICs are designed for — which is why four programs hit production
milestones simultaneously rather than coincidentally.

**The foundry behind the custom wave.** One dependency underlies all four
custom-ASIC programs in this section, and it deserves explicit statement:
leading-edge foundry capacity. Meta's Iris is manufactured by TSMC (verified);
the industry's custom silicon overwhelmingly tapes out on the same leading
edge. This creates a second-order concentration worth noting as structural
context: the industry's escape from single-supplier dependence on merchant
accelerators runs through single-supplier dependence on advanced foundry
capacity — and that capacity, like memory fab capacity, is built on
multi-year lead times. The custom-ASIC strategy diversifies *who designs the
chip*; it does not diversify *who can physically build it*. In a chapter
about physical limits, that is not a footnote.

