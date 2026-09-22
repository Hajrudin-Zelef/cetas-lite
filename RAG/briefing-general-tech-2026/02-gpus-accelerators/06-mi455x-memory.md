---
id: briefing-general-tech-2026/02-gpus-accelerators/06-mi455x-memory
title: "MI455X and the memory play"
domain: gpus-accelerators
role: deep-dive
task: memory-crisis
actors: ["AMD", "Anthropic", "Intel", "Micron", "Nvidia", "OpenAI", "Qualcomm"]
dates: ["2026-07", "2026-09-22"]
keywords: ["mi455x", "accelerator", "agentic", "ai200", "benchmark", "cost per token", "crescent island", "dram", "fp4", "fp8", "gpu", "gpus"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-6"
source_lines: [1797, 1869]
canonical_for: ["amd-mi455x"]
sha256: eb2110af6cb81cabe2f154d73f225ce3ff9fb062c330079681177098ab6fe692
---

# MI455X and the memory play

<a id="g03-6"></a>
### 3.6 MI455X and the memory play

If Rubin is a compute-density story, Helios is a memory story. The MI455X — the lead SKU of the MI450 family — is positioned by AMD not as a FLOPS leader but as the accelerator for the workloads where the industry's binding constraint actually sits: memory capacity and memory bandwidth. The verified Helios rack figures:

| Specification | Helios rack (72× MI455X) | Notes |
|---|---|---|
| GPUs per rack | 72 MI455X | Same rack unit as NVL72 |
| Unified memory | 31 TB HBM4 | Pooled across the rack |
| Compute | 2.9 exaFLOPS FP4 / 1.4 exaFLOPS FP8 | Per rack |
| Scale-out bandwidth | 43 TB/s | — |
| Scale-up bandwidth (aggregate) | 260 TB/s | Aggregate across the domain, as with NVL72 |

Two things stand out in comparison with Rubin. First, the scale-up bandwidth figure is the same 260 TB/s — the same aggregate-domain number, subject to the same per-link misreading warning as in §3.2. Second, the memory capacity is the differentiator: 31 TB of HBM4 pooled in one rack is the foundation of every AMD claim that follows.

#### HBM4 in one paragraph: why the memory generation matters

High-bandwidth memory stacks DRAM dies vertically on or beside the accelerator package, connecting them with thousands of microscopic interconnects — delivering an order of magnitude more bandwidth than discrete memory at the cost of capacity limits and manufacturing complexity. HBM4, the 2026 generation, increases both bandwidth and capacity over HBM3E. Its significance in this chapter is twofold. First, it is the *constrained* resource: HBM supply was allocated through 2026, which is why Micron sits in the trillion-dollar club (§3.8) and why Intel's Crescent Island (§3.11) and Qualcomm's AI200 (§3.17) deliberately avoid it. Second, it is AMD's chosen battlefield: "50% more HBM4 capacity and bandwidth" is a claim about out-provisioning Nvidia on the scarcest, most valuable component in the rack. In a memory-bound era, the HBM4 allocation is the moat — and AMD's pitch is that its moat is wider.

#### 43 TB/s scale-out: the between-racks fabric

Helios's 43 TB/s scale-out bandwidth is the lesser-quoted of its two fabric numbers, but it answers a different question than the 260 TB/s scale-up figure: how fast can *racks* talk to *racks*. Scale-out is what turns a 72-GPU system into a multi-rack cluster — the domain where the largest training runs and the largest serving deployments actually live. The order-of-magnitude gap between scale-up (260 TB/s) and scale-out (43 TB/s) is the industry's standard hierarchy: in-rack bandwidth is plentiful, between-rack bandwidth is the constraint, and workload placement (keeping communication-intensive work inside the rack) is the optimization. Any multi-rack Helios deployment — OpenAI's "massive scale," Anthropic's gigawatts — will be gated by this number long before it is gated by per-rack compute.

#### Su's two headline claims — and their status

| Claim (Lisa Su, Advancing AI, July 2026) | Status |
|---|---|
| "50% more HBM4 capacity and bandwidth" than Rubin | **AMD claim — vendor benchmark, not independently verified** |
| Up to 30% more tokens per dollar vs. Rubin | **AMD claim — vendor benchmark, not independently verified** |

Both are **AMD benchmarks and AMD's framing — vendor claims, not independently verified.** No third party had measured a Helios rack against an NVL72 as of September 22, 2026, because no third party had a Helios rack. The claims are nevertheless strategically legible: AMD is not contesting raw FP4 throughput (Rubin's 3.6 vs. Helios's 2.9 exaFLOPS per rack favors Nvidia). It is contesting the economics of serving large models — tokens per dollar — on the theory that inference at scale is memory-bound, and that the buyer who cares about cost per token will pick the fatter memory subsystem.

#### 31 TB, worked out: what pooled HBM4 holds

Thirty-one terabytes of HBM4 pooled across 72 GPUs works out to roughly 430 GB per GPU — and the pooling matters as much as the total. In a unified memory domain, the rack's memory behaves as one addressable pool: a single model instance can span GPUs without the per-device capacity walls that fragment multi-GPU serving. Concretely, 31 TB holds models in the hundreds-of-billions-of-parameters class at 4-bit precision with room left for large-batch KV caches — the configuration where Su's tokens-per-dollar claim is strongest. The comparison point is architectural, not just numerical: Nvidia's NVL72 is also a unified domain, so both vendors offer pooling; AMD's claim is that its pool is ~50% deeper. Depth of pool, not just its existence, is the contested metric — and it is the one no third party has measured.

#### 2.9 vs. 3.6: what a 19% FLOPS deficit means

The FP4 gap between Helios (2.9 EFLOPS) and Rubin (3.6 EFLOPS) is roughly 19% — close enough that software, memory, and pricing decide the contest, far enough that AMD cannot claim compute parity. Nineteen percent is the worst kind of deficit for a challenger: too small to concede the segment, too large to ignore. It is precisely why AMD's pitch leads with memory and tokens-per-dollar rather than FLOPS — the company is competing *around* the deficit, not through it. And it is why Su deferred compute leadership to MI500: closing 19% takes a generation, and AMD's roadmap says that generation is 2027. Until then, every Helios sale is a sale made on economics despite the FLOPS gap — which, if the memory-bound thesis holds, is exactly the right sale to make.

#### The theory behind the memory play

The theory deserves a paragraph, because it is the most consequential strategic bet any challenger made in 2026. Large-model inference in 2026 is dominated by two costs: storing the model weights (capacity) and streaming them to the compute units once per generated token (bandwidth). The KV cache — the stored attention state that grows with context length and batch size — adds a third capacity demand that scales with exactly the agentic, long-context workloads the industry is pivoting toward. A rack with 50% more HBM4 capacity can hold larger models, larger batches, or longer contexts without spilling to slower tiers; 50% more bandwidth turns directly into tokens per second. If inference economics are set by memory, then the vendor with the fattest memory subsystem wins tokens-per-dollar even while losing FLOPS — which is precisely AMD's claim.

The honest caveats: the claim assumes the workload is in fact memory-bound (true for many serving workloads, less true for training and for compute-heavy phases), and it assumes AMD's software stack extracts the hardware's potential (the open question behind every AMD generation). But as a positioning statement it is coherent, differentiated, and aimed at the market segment — inference — where the growth is.

#### "Ideal for memory-bound workloads": whose words?

Analysts spent the summer describing Helios as **"ideal for memory-bound workloads"** — a characterization to note carefully, because it is analyst framing, not an AMD quote. AMD's own language was the tokens-per-dollar claim. The distinction matters for the dossier's anti-fabrication discipline: the "memory-bound" label is the market's interpretation of AMD's positioning, and it should be attributed that way. The market heard "50% more HBM4" and translated it into "the memory-bound workload chip" — a reasonable translation, but a translation.

#### The MI500 deferral

One more piece of Su's positioning deserves quoting precisely because of what it concedes: she promised **scale-up compute leadership only with the MI500 generation in 2027**. That is an explicit deferral — an admission that on raw scale-up compute, Rubin leads through 2026 and AMD's answer is a next-generation part. It is also a roadmap commitment the market will hold her to: MI500 is now the named vehicle for AMD's compute-leadership claim, and its 2027 delivery is the date that claim lives or dies on. Note the symmetry with Nvidia's cadence logic (§3.1): AMD is asking buyers to commit to Helios in 2026 on the promise of MI500 in 2027 — the same "buy the tick, get the tock" pitch Nvidia makes with Rubin → Rubin Ultra.

---

#### Precision in one paragraph: why FP4 and FP8 are the rack's headline numbers

Both Rubin (3.6 EFLOPS NVFP4) and Helios (2.9 EFLOPS FP4 / 1.4 EFLOPS FP8) headline their compute in 4-bit and 8-bit floating point rather than the 16-bit precision of earlier generations. The reason is arithmetic, not marketing: inference at scale is dominated by moving weights, and halving the bits per weight halves the memory traffic per token. A rack's FP4 figure is therefore a rough proxy for its token-generation ceiling on quantized models — which is why both vendors lead with it, and why the FP4 gap (3.6 vs. 2.9) is the honest headline of the compute comparison. FP8, at double the bits, is the training-relevant figure and the precision where Helios's 1.4 EFLOPS sits. Neither vendor's 2026 headline used FP16; the industry has moved on.

#### Tokens per dollar: the claim's internal logic

Su's "up to 30% more tokens per dollar" claim has a structure worth exposing, because it shows what AMD is — and isn't — asserting:

1. **Premise:** inference serving is memory-bound (capacity for weights + KV cache, bandwidth for token streaming).
2. **Hardware fact (claimed):** Helios has ~50% more HBM4 capacity and bandwidth per rack than Rubin.
3. **Inference:** more memory headroom → larger batches / longer contexts per rack → more tokens per unit of hardware cost.
4. **Conclusion (claimed):** up to 30% more tokens per dollar.

The logic is valid *if* the premise holds and *if* the hardware facts are right. Step 2 is the vendor claim awaiting verification; step 1 holds for many serving workloads but not all (compute-heavy phases, small-batch latency-sensitive serving). The "up to" is doing the usual work of capping the claim at the favorable end of the distribution. A buyer evaluating the claim should ask for the workload mix behind the 30% — which AMD had not published as of September 22.

---

