---
id: briefing-general-tech-2026/02-gpus-accelerators/19-competitive-landscape
title: "Competitive landscape: who stands where"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "Anthropic", "China", "CoreWeave", "Dell", "Fujitsu", "Huawei", "Intel", "MLCommons", "Meta", "Nvidia", "OCP", "OpenAI", "Qualcomm"]
dates: ["2026-03", "2026-04", "2026-06", "2026-06-01", "2026-07", "2026-09", "2026-09-22", "2026-11"]
keywords: ["18a", "accelerator", "agentic", "agi", "ai200", "ascend", "benchmark", "blackwell", "clearwater forest", "cost per token", "crescent island", "ethernet"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-19"
source_lines: [2643, 2881]
canonical_for: ["gpu-landscape"]
sha256: 82531f9acc782b439c73bcd290ce668a99be745782b52ba72a466eb9ca1f89b1
---

# Competitive landscape: who stands where

<a id="g03-19"></a>
### 3.19 Competitive landscape: who stands where

Stepping back from the preceding sections, the accelerator market of September 2026 sorts into a clear structure — with one dominant player, one credible challenger, one incumbent fighting back on two fronts, one sovereign contender playing a different game, and two specialists.

#### Nvidia: dominant but memory-constrained

Rubin is in production, shipping in volume in H2 2026, benchmarking first in MLPerf, and already deployed in a 504-GPU multi-rack cluster. The roadmap (NVL144, Rubin Ultra NVL576 in H2 2027, Feynman 2028) is the industry's metronome. The constraints are physical — 190–230 kW per rack, liquid cooling, datacenter power delivery — and perceptual: Huang's "AGI has arrived" post shows a CEO marketing the installed base as much as the new platform. Nvidia's lead is real and measured; it is also the lead everyone else's roadmap is now explicitly calibrated against. The company's structural advantage remains the full stack: silicon, NVLink fabric, and the software (Dynamo, TensorRT-LLM, NVFP4 tooling) that turned the MLPerf multiples of §3.3 from hardware specs into measured results.

#### Nvidia's moat, inventoried

Nvidia's lead in September 2026 rests on five layers, and challengers must breach all five:

1. **Silicon** — Rubin shipping; the only measured platform in the chapter.
2. **Fabric** — NVLink 6 scale-up; the domain is the product.
3. **Software** — Dynamo, TensorRT-LLM, NVFP4 tooling; the MLPerf multiples are software multiples.
4. **Cadence** — Rubin → Rubin Ultra → Feynman; buyers purchase the roadmap, not the rack.
5. **Channel** — hyperscalers plus neoclouds; the fastest deployment path in the industry.

Each layer reinforces the others: the cadence justifies the software investment, the software extracts the silicon's value, the channel delivers it first. A challenger that matches one layer (AMD's memory, Huawei's optics, Intel's value) still faces four. That is why the chapter's verdict is "dominant" rather than merely "ahead" — and why the settlement dates that matter most are the ones where challengers attack multiple layers at once (Helios: silicon + deals; SuperPoD: system + optics + software).

#### AMD: gaining via memory-heavy scale-up and deals

Helios concedes raw FP4 throughput (2.9 vs. 3.6 exaFLOPS per rack) and contests the economics: 31 TB of pooled HBM4, "50% more HBM4 capacity and bandwidth," up to 30% more tokens per dollar — all vendor claims, all pointed at the memory-bound inference market. The three lab deals (14 GW forward-committed: OpenAI's 6 GW, Meta's 6 GW custom, Anthropic's 2 GW) are the demand-side proof of the strategy; the September "on track" earnings language and the $1 trillion crossing are the market's verdict so far. The risk is execution: no customer deliveries confirmed as of September 22, and scale-up compute leadership is explicitly deferred to MI500 in 2027. AMD's 2026 is a bet that the market will wait — and at $1,005 billion, the market has placed it.

#### AMD's path to winning a segment first

AMD does not need to beat Nvidia everywhere to change the industry — it needs to win one segment decisively. The most plausible first win is inference value: if Helios delivers near its tokens-per-dollar claims on memory-bound serving workloads, cost-sensitive clouds and enterprises have a rational reason to second-source, and second-sourcing is the foot in the door that becomes a platform shift. The three lab deals are the mechanism: OpenAI, Meta, and Anthropic deploying Helios at gigawatt scale would create the operational proof — and the software maturity — that smaller buyers need. History suggests challengers win by owning the new workload (inference, agents) rather than the old one (training); AMD's entire 2026 positioning reads as an attempt to run that playbook. The MI500 compute-leadership promise is the second act; the inference-value win is the first.

#### Intel: fighting on CPU plus inference value

Clearwater Forest (288 E-cores, Intel 18A, June 2026) gives Intel a current datacenter CPU with a "control plane of agentic workloads" thesis and startling demand signals ("only supply 50% of customers"). Panther Lake/X9 gives it an 18A client story. Crescent Island is the wildcard: a 350 W air-cooled LPDDR5X inference GPU with no HBM, no display output, and a tokens-as-a-service pitch — the most aggressive value play in the inference market, sampling H2 2026, generally available 2027. Intel is not trying to out-Nvidia Nvidia; it is trying to own the CPU side of the agentic rack and the value side of inference. If the control-plane thesis (§3.10) is right, the CPU:GPU ratio shift toward 1:1 is a structural tailwind no accelerator vendor can capture — it accrues to Intel by default.

#### Huawei: system-level compensation for China

Approximately two years behind per chip (Blackwell-class, ~10x under Rubin per GPU on paper), Huawei answers with the most ambitious system in the chapter: a 15,488-chip, 220-cabinet SuperPoD with first-to-mass-production NPO optics, a unified 12-protocol fabric, an annual "Tao's Law" cadence, and PyTorch backend recognition. The strategy is China-first by stated choice and capacity constraint — Eric Xu's words — which means Huawei competes for the narrative of parity and for Chinese demand, not (yet) for global market share. Every claim is unverified by independent benchmarks; the Q4 2027 SuperPoD delivery is the date the thesis becomes testable. The bull case (Chan: "the effective gap is narrowing") and the bear case (The Register: two years behind per chip) are both on the record; silicon will settle it.

#### Qualcomm and Fujitsu: niche inference plays

Qualcomm's AI200 (768 GB LPDDR per card, 160 kW liquid-cooled racks, HUMAIN's 200 MW as anchor customer) bets on hyperscale inference capacity for sovereign-scale buyers; 2026 shipping is announced, not confirmed. Fujitsu's MONAKA (144-core 3D-stacked Arm CPU, made in Japan, sales starting November 2026) bets on sovereignty as a product category. Both are inference-first, both avoid the training fight entirely, and both will be judged on deliveries in the next two quarters — the shortest fuse in the chapter.

#### The two scoreboards

A final structural observation: 2026 produced not one accelerator ranking but two, and they measure different things.

**Scoreboard 1 — frontier training and large-scale inference** (Nvidia > AMD > Intel > Huawei): ranked on measured or claimed throughput, memory systems, and lab deals. This is the scoreboard the press covers, and the one §3.19's table summarizes.

**Scoreboard 2 — sovereign and value inference** (Fujitsu, Qualcomm, Intel's Crescent Island): ranked on supply-chain provenance, power efficiency, capacity-per-dollar, and deployability in existing datacenters. This scoreboard barely registers in FLOPS comparisons, but it is where procurement mandates and national budgets are spent.

The two scoreboards have different buyers, different metrics, and different settlement dates — and a vendor can lose the first while winning the second. Huawei straddles both (frontier claims, sovereign market); Qualcomm and Fujitsu play only the second; Nvidia and AMD play only the first. Reading 2026 as a single race misses half the industry.

#### The buyer's guide: who should buy what

For readers using this chapter as a procurement reference — with all the evidentiary caveats above — the 2026 lineup sorts by buyer need:

| Buyer need | Best fit (Sept 2026) | Caveat |
|---|---|---|
| Frontier training, shipping now | Nvidia Rubin NVL72 | Power/cooling envelope; premium pricing |
| Inference at scale, cost-optimized | AMD Helios (if claims hold) | Unshipped; verify tokens/$ on your workload |
| Agentic/long-context inference, value | Intel Crescent Island (2027) | Unsampled; the value thesis is untested |
| CPU-heavy agentic orchestration | Intel Xeon 6+ Clearwater Forest | The control-plane thesis is Intel's, not proven |
| China-domestic frontier scale | Huawei Atlas 960 / Ascend 960 | China-first availability; no independent benchmarks |
| Sovereign-scale inference capacity | Qualcomm AI200 | Announced; HUMAIN deployment is the proof |
| Sovereign-provenance inference | Fujitsu MONAKA | Sales Nov 2026; verify the 2x claim |

The table is interpretive — a synthesis of the chapter's verified facts, not itself a verified fact. Its honest summary: as of September 2026, exactly one row ("frontier training, shipping now") could be bought without qualification. Every other row is a bet on a settlement date. That is the state of the industry the chapter documents.

#### The landscape at a glance

| Vendor | Platform | Rack unit | Memory story | Status (22/09/2026) | Decisive date |
|---|---|---|---|---|---|
| Nvidia | Vera Rubin NVL72 | 72 GPUs + 36 CPUs, 3.6 EFLOPS FP4 | NVLink 6, 260 TB/s aggregate | In production; volume H2 2026 | H2 2026 ramp |
| AMD | Helios (MI455X) | 72 GPUs, 2.9 EFLOPS FP4 | 31 TB HBM4 pooled; tokens/$ claim | Announced; first deliveries Sept (unconfirmed) | H1 2027 (Anthropic GW) |
| Intel | Xeon 6+ / Crescent Island | CPU + inference GPU | LPDDR5X value play (Crescent Island) | Xeon shipping; Crescent sampling H2 2026 | 2027 (Crescent GA) |
| Huawei | Ascend 960 / Atlas 960 SuperPoD | 15,488-chip system, 60 EFLOPS FP4 | NPO optics, unified fabric | Announced; China-first | Q4 2027 (SuperPoD) |
| Qualcomm | AI200 | 160 kW rack, liquid-cooled | 768 GB LPDDR/card | Announced; 2026 shipping unconfirmed | HUMAIN deployment |
| Fujitsu | MONAKA | 1U/2U inference servers | 144-core 3D-stacked CPU | Sales start Nov 2026 (announced) | Nov 2026 – Apr 2027 |

#### Master chronology: every dated event in the chapter

| Date | Event | Section |
|---|---|---|
| Oct 6, 2025 | OpenAI–AMD deal: 6 GW / 5 years, 160M warrants | §3.7 |
| Oct 27, 2025 | Qualcomm AI200 announced | §3.17 |
| OCP Global Summit 2025 | Intel Crescent Island announced | §3.11 |
| Jan 5–6, 2026 | CES: Core Ultra X9 388H launches; Huang declares "Vera Rubin is in full production" | §3.9, §3.1 |
| Jan 27, 2026 | First Panther Lake laptops on sale | §3.9 |
| Feb 24, 2026 | Meta–AMD deal: 6 GW custom MI450 GPU, 160M warrants | §3.7 |
| March 2026 | MWC: Xeon 6+ Clearwater Forest announced | §3.10 |
| Mar 16, 2026 | GTC: full Vera Rubin reveal; cost-per-token claims | §3.1 |
| April 2026 | Core Ultra X9 378H listed on Intel ARK; GPT-class pilots on Helios begin (~April) | §3.9, §3.7 |
| June 1, 2026 | First NVL72 rack in production (CoreWeave/Dell, Livingston); Xeon 6+ launched at Computex | §3.1, §3.10 |
| Jul 22, 2026 | Anthropic–AMD deal: 2 GW MI450, first GW H1 2027 (Advancing AI day 1) | §3.7 |
| Jul 23, 2026 | Su keynote: Helios, MI455X, EPYC Venice; Katti on stage | §3.5, §3.7 |
| Late July 2026 | AMD guidance: first Helios deliveries September | §3.5 |
| Week of Aug 24, 2026 | Crescent Island technical details at Hot Chips 2026 | §3.11 |
| Sept 3, 2026 | GPT-6 Astra launches (Reuters) | §3.4 |
| Sept 6, 2026 | Huang: "AGI has arrived" (X; 300K→~100K); AMD Q2 earnings: Helios "on track" | §3.4, §3.5 |
| ~Sept 7, 2026 | PyTorch backend recognition via TorchNPU (Linux Foundation) | §3.16 |
| Sept 14, 2026 | Fujitsu MONAKA press release (Kawasaki) | §3.18 |
| Sept 16, 2026 | MLPerf Inference v6.1 (Rubin "preview submission"); CoreWeave 504-GPU cluster | §3.3 |
| Sept 17, 2026 | Huawei Connect keynote (David Wang); Eric Xu to Reuters; SCMP on NPO; The Register on per-chip gap; Nikkei Asia on chipsets; Morgan Stanley via Bloomberg | §3.12–§3.16 |
| Sept 17–19, 2026 | Huawei Connect 2026, Shanghai | §3.12 |
| Sept 21, 2026 | AMD crosses $1T (+9.6% to $613.50, ~$1,005B) | §3.8 |
| Sept 22, 2026 | Dossier cutoff | — |

#### The claims ledger: every vendor claim and its status

| # | Claim | Maker, date | Status (22/09/2026) |
|---|---|---|---|
| 1 | Rubin: MoE training with 1/4 the GPUs vs Blackwell | Nvidia, GTC Mar 2026 | Vendor claim — unverified |
| 2 | Rubin: 10x inference throughput/watt at 1/10 cost per token | Nvidia, GTC Mar 2026 | Vendor claim — unverified |
| 3 | Rubin NVL72: up to 3.7x token throughput on Qwen3-VL | Nvidia, MLPerf Sept 2026 | Measured (preview submission) |
| 4 | Rubin NVL72: up to 2.5x on DeepSeek-R1 | Nvidia, MLPerf Sept 2026 | Measured (preview submission) |
| 5 | "~100,000 Grace Blackwell systems" behind Astra-era scale | Huang, X Sept 2026 | CEO claim, revised down from 300K |
| 6 | Helios: "50% more HBM4 capacity and bandwidth" vs Rubin | AMD (Su), Jul 2026 | Vendor claim — unverified |
| 7 | Helios: up to 30% more tokens/dollar vs Rubin | AMD (Su), Jul 2026 | Vendor claim — unverified |
| 8 | "Ideal for memory-bound workloads" (Helios) | Analysts, summer 2026 | Analyst framing, not AMD quote |
| 9 | Scale-up compute leadership with MI500 (2027) | AMD (Su), Jul 2026 | Roadmap promise |
| 10 | Helios "at massive scale" end of 2026, accelerating 2027 | OpenAI (Katti), Jul 2026 | Customer statement — pending |
| 11 | CPU:GPU ratio moving from 1:8 toward ~1:1 | Intel (Tan), Jun 2026 | CEO thesis — unverified |
| 12 | "Only supply 50% of customers" (CPU demand) | Intel (Tan), Sept 2026 | CEO statement |
| 13 | Crescent Island for "tokens-as-a-service" inference | Intel, 2026 | Positioning — sampling H2 2026 |
| 14 | Ascend 960DT: 2 PFLOPS FP8 / 4 PFLOPS FP4, 288 GB HBM | Huawei, Sept 2026 | Vendor claim, unshipped silicon |
| 15 | Atlas 960 SuperPoD: 2.3x training / 2.5x inference vs 950 SuperPoD | Huawei, Sept 2026 | Vendor claim — no independent benchmarks |
| 16 | NPO: 5,500 Hi-ONE modules replace ~48,000 800G; >550 kW saved | Huawei, Sept 2026 | Vendor claim (960E config only) |
| 17 | 99.8% availability (<18h unplanned downtime/year) | Huawei, Sept 2026 | Vendor claim |
| 18 | UnifiedBus: 12+ protocols, latency 7 µs → 2 µs | Huawei, Sept 2026 | Vendor claim |
| 19 | "The effective gap is narrowing" (system-level thesis) | Morgan Stanley (Chan), Sept 2026 | Analyst assessment |
| 20 | 960DT ≈ 1/2 FP8, ~1/3 FP4 of B300 (~2 years behind) | The Register, Sept 2026 | Press analysis of claimed specs |
| 21 | AI200: 768 GB LPDDR/card; 2026 shipping | Qualcomm, 2025–2026 | Announced — no deliveries confirmed |
| 22 | MONAKA: "2x throughput vs other CPUs" | Fujitsu, Sept 2026 | Internal estimate — unverified |
| 23 | $60–100B implied valuation (Meta–AMD deal) | Sources (press), 2026 | Unofficial, per sources |

The ledger's pattern is the chapter's pattern: Nvidia's claims 3–4 are measured (the only measured vendor claims in the chapter); everything else is a keynote claim, a roadmap promise, an analyst assessment, or a CEO statement — each labeled, each awaiting its settlement date.

#### The 2026 scoreboard: measured vs. claimed

| Vendor | Measured results (third-party rules) | Claimed figures | Ratio |
|---|---|---|---|
| Nvidia | MLPerf v6.1: 3.7x (Qwen3-VL), 2.5x (DeepSeek-R1) | GTC cost/token; roadmap to Feynman | The only vendor with both |
| AMD | None (no shipped racks to test) | Helios specs; tokens/$; 14 GW deals | All claimed/committed |
| Intel | None (Crescent Island unsampled) | Xeon 6+ shipping; Crescent specs | Half shipped, half claimed |
| Huawei | None | 960DT specs; SuperPoD; NPO; 2.3x/2.5x | All claimed |
| Qualcomm | None | AI200 specs; HUMAIN 200 MW | All claimed |
| Fujitsu | None | MONAKA specs; 2x estimate | Announced; sales Nov 2026 |

One vendor with measurements, five with announcements. That is the 2026 scoreboard in its entirety — and the reason the dossier's evidentiary discipline matters more than any single figure in it.

#### What would change the ranking

The §3.19 ranking (Nvidia > AMD > Intel > Huawei > Qualcomm/Fujitsu) is a September 2026 snapshot. Five developments would force a rewrite, in rough order of likelihood:

1. **Helios ships and benchmarks.** If first deliveries land in Q4 2026 with third-party MLPerf numbers near AMD's claims, the "credible challenger" becomes the "measured challenger" — the single most ranking-relevant event available.
2. **The 14 GW converts.** OpenAI's end-of-2026 ramp and Anthropic's H1 2027 gigawatt are the demand-side settlement dates; conversion confirms the deals as revenue, slippage reprices them as options.
3. **Rubin stumbles.** A production or deployment problem in the H2 2026 volume ramp — power-delivery issues at 190–230 kW, yield problems, schedule slips — would narrow the maturity gap that underwrites Nvidia's lead.
4. **Huawei's SuperPoD delivers early or benchmarks independently.** Any third-party measurement of Atlas hardware — even the 950 SuperPoD in Q4 2026 — would move Huawei's claims from the "unverifiable" column for the first time.
5. **A sovereign buyer deploys at scale.** HUMAIN's 200 MW on AI200, or MONAKA's November sales converting to visible deployments, would validate the niche-inference thesis and create a second scoreboard beside the frontier-training one.

#### Open questions the chapter leaves

- **The tokens-per-dollar contest:** AMD's and Intel's shared metric has no neutral referee. Who builds the industry-standard cost-per-token benchmark — MLCommons, a cloud vendor, or nobody?
- **The memory wall:** HBM4 allocations, LPDDR5X capacity plays, NPO optics, near-memory (AI250's HBC) — four different answers to the same constraint. Which scales?
- **The fabric wars:** NVLink 6, AMD's scale-up, UnifiedBus, Ethernet scale-out — the interconnect is now a strategic layer. Does it consolidate on one standard or fragment by vendor?
- **The CPU's revenge:** if Tan's control-plane thesis is right, the 2027 datacenter looks very different from the 2025 one. Is the 1:1 ratio real, and who captures the CPU-side value?
- **China's clock speed:** "Tao's Law" promises annual generations. The 960DT's Q1 2027 date is the first test of whether Huawei's cadence is real or rhetorical.

#### The settlement calendar: every decisive date in one place

The chapter's claims resolve on specific dates. All of them, in chronological order:

| Date | What gets tested | Section |
|---|---|---|
| H2 2026 | Rubin volume shipments (8 partners) | §3.1 |
| H2 2026 | Crescent Island customer sampling | §3.11 |
| September 2026 | Helios first deliveries (announced) | §3.5 |
| Q4 2026 | Ascend 950DT delivery; Atlas 950 SuperPoD (baseline for 2.3x/2.5x) | §3.12, §3.14 |
| November 2026 | MONAKA sales start (Japan, Europe) | §3.18 |
| End of 2026 | OpenAI Helios deployment "at massive scale" begins (Katti) | §3.7 |
| 2026 (unspecified) | AI200 shipping; HUMAIN 200 MW progress | §3.17 |
| Q1 2027 | Ascend 960DT (pulled in 3 quarters — the pull-in claim) | §3.12 |
| Q1 2027 | MONAKA ramp into April; US/APAC globalization (Q4 FY2026) | §3.18 |
| Early 2027 | Qualcomm AI250 (HBC near-memory) | §3.17 |
| H1 2027 | First Anthropic gigawatt online (MI450) | §3.7 |
| Q3 2027 | Ascend 960PR (inference SKU; specs still undisclosed) | §3.13 |
| H2 2027 | Rubin Ultra NVL576; AMD MI500 (compute-leadership claim) | §3.1, §3.6 |
| 2027 | Crescent Island general availability | §3.11 |
| Q4 2027 | Full Atlas 960 SuperPoD delivery (15,488 chips — the system thesis, testable) | §3.14 |
| 2028 | Ascend 970; Feynman (Nvidia); AI300 (Qualcomm roadmap) | §3.12, §3.1, §3.17 |
| 2029 | Ascend 980 | §3.12 |

#### The power story: the meta-constraint

Every power figure in this chapter, in one place:

| System | Power | Cooling | Status |
|---|---|---|---|
| Rubin NVL72 | ~190–230 kW / rack | Liquid (required) | Shipping |
| Blackwell GB200 NVL72 | ~100–120 kW / rack | — | Previous gen |
| Qualcomm AI200 | 160 kW / rack | Liquid | Announced |
| Crescent Island | 350 W / card | Air | Sampling H2 2026 |
| Atlas 960E optics saving | >550 kW (optical subsystem) | — | Claimed |

The pattern: flagship rack power roughly doubled from Blackwell (~100–120 kW) to Rubin (~190–230 kW) in one generation — and Huawei's optical power *savings* alone (>550 kW) exceed an entire Rubin rack's draw. Power is the meta-constraint because it gates everything else: a datacenter's megawatt capacity determines how many racks it can host, liquid-cooling retrofits determine how fast old buildings can take new racks, and national power grids determine where new datacenters get built at all. The vendors who designed for the constraint (neocloud greenfield sites, Crescent Island's air-cooled 350 W, Huawei's optical savings) gain deployment speed; those who didn't inherit a retrofit problem. In 2026, power delivery — not silicon — was the most common reason a "shipped" platform wasn't yet "deployed."

#### The precision story: FP4 everywhere

One underappreciated uniformity: every vendor in this chapter converged on 4-bit precision as the headline format. Nvidia (NVFP4), AMD (FP4), Intel (FP4/MXFP4), Huawei (FP4), Qualcomm and Fujitsu (positioned for low-bit inference). The convergence reflects a shared technical judgment — that 4-bit inference preserves model quality well enough for serving while roughly doubling effective throughput over 8-bit — and a shared economic one: in a memory-bound era, bits per weight are dollars per token. The format *names* differ (NVFP4 vs. MXFP4 vs. generic FP4), and portability between them is an open question the OCP Microscaling effort (MXFP4) exists to solve. But the direction is unanimous: the industry's 2026 was the year 4-bit went from quantization trick to headline spec.

#### By the numbers: the chapter on one page

| Vendor | Platform | Pinnacle figure (claimed) | Memory figure | Power figure | Status |
|---|---|---|---|---|---|
| Nvidia | Rubin NVL72 | 3.6 EFLOPS FP4 / rack | 260 TB/s aggregate scale-up | ~190–230 kW / rack | Shipping |
| AMD | Helios | 2.9 EFLOPS FP4 / rack | 31 TB HBM4 / rack | n/d | Announced |
| Intel CPU | Xeon 6+ | 288 E-cores / socket | n/a (CPU) | n/d | Shipping |
| Intel GPU | Crescent Island | 32 Xe3P cores | 160–480 GB LPDDR5X | 350 W / card | Sampling H2 2026 |
| Huawei chip | Ascend 960DT | 4 PFLOPS FP4 / chip | 288 GB HBM / chip | n/d | Q1 2027 |
| Huawei system | Atlas 960 SuperPoD | 60 EFLOPS FP4 | 4,460 TB | n/d (optics save >550 kW) | Q4 2027 |
| Qualcomm | AI200 | n/d (bandwidth undisclosed) | 768 GB LPDDR / card | 160 kW / rack | Announced |
| Fujitsu | MONAKA | 144 cores, 3.8 GHz | 3D-stacked SRAM | n/d | Sales Nov 2026 |

("n/d" = not disclosed in verified sources — the empty cells are data, not omissions.)

#### How to use this chapter in 2027

Three habits will keep this chapter useful as the settlement dates arrive. First, **check the announced against the delivered**: every "announced/planned" in the text is a promise with a date, and the calendar above is the checklist. Second, **weight claims by evidence distance**: measured (MLPerf) outranks claimed (keynotes), shipped outranks announced, and independent outranks vendor — the gradient in the methodology note below. Third, **watch the memory numbers**: across every vendor, the 2026–2027 contest is being fought on HBM4 capacity, LPDDR5X capacity-per-dollar, NPO optics, and near-memory architectures — not on FLOPS. The vendor that moves its memory story from claimed to measured first wins the next round of the argument.

#### How this chapter connects to the rest of the dossier

Hardware does not move alone, and the chapter's events rhyme with the dossier's other chapters. The September pile-up — Astra's launch (models), the "AGI" post (narrative), MLPerf (hardware measurement), Huawei Connect (sovereign hardware), AMD's trillion (capital) — is a single story told in five chapters: models create demand, hardware narratives amplify it, capital prices it. The Anthropic–AMD deal (§3.7) belongs equally to a chapter on lab strategy; the export-control backdrop (§3.16) to a chapter on geopolitics; the $1 trillion crossing (§3.8) to a chapter on markets. Readers using this dossier for Q&A should treat the chapters as facets: the hardware facts here supply the *substrate* that the model and capital chapters assume. When a model chapter asks "what made this launch possible," the answer is often a row in §3.19's tables.

#### The through-line

The through-line of 2026 is that the industry stopped selling chips and started selling systems on annual cadences — and that the scarcest resource in the system is no longer FLOPS but memory, power, and the datacenter to house them. Every vendor in this chapter has a plan for 2027. None of the challengers' plans have survived contact with customers yet. That is what 2027 is for.

A final methodological note for readers using this chapter as a reference: the evidentiary gradient runs downhill from Nvidia (shipping, benchmarked) through AMD (announced, committed, unshipped) and Intel (half-shipped) to Huawei, Qualcomm, and Fujitsu (announced, largely unshipped). Claims should be weighted accordingly — not dismissed, but discounted by distance from delivery. The dates in the settlement calendar above are the dates to check. Check them when they arrive.

---
