---
id: briefing-general-tech-2026/04-training-inference-quantization/15-what-quantization-changes
title: "What quantization changes for the industry"
domain: training-inference-quantization
role: deep-dive
task: quantization
actors: ["AMD", "Amazon", "Apple", "BitNet", "Cerebras", "Crusoe", "Groq", "Intel", "MLCommons", "Nvidia", "Positron", "PrismML", "Taalas"]
dates: ["2024-02-27", "2024-10-17", "2025-04", "2026-02-20", "2026-03", "2026-06", "2026-07", "2026-07-14", "2026-08", "2026-09", "2026-09-11", "2026-09-16", "2026-09-22"]
keywords: ["quantization", "agentic", "awq", "benchmark", "bitnet", "blackwell", "cost per token", "cowos", "decode", "disaggregated", "disaggregated serving", "fp4"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-15"
source_lines: [5454, 5587]
sha256: b20fb3dd55a69614913e196b98fe73fcc46a63f265833c378257e96b692c1cd6
---

# What quantization changes for the industry

<a id="g05-15"></a>
### 5.15 What quantization changes for the industry

Step back from the formats, the methods, the kernels and the flags, and the
picture of 2026 resolves into a single economic mechanism with architectural
consequences. Quantization — from FP8's safe tier through INT4's aggressive
standard, through the 3-bit KV cache, to the 1.58-bit frontier — is the
technology that let inference economics detach from training economics. And
that detachment is what made the great specialization of Section 5.1
physically possible: without the bit-level economics, the inference rack could
not have diverged from the training rack, because it could not have afforded
to.

The causal chain runs in one direction and is worth stating end to end.
Training still needs dense HBM racks at 190–230 kW because synchronous
training is bandwidth- and power-bound with no compression escape hatch — you
cannot quantize a gradient the way you quantize a weight, and the NVL72's
power envelope is the price of keeping 72 accelerators in lockstep. Inference,
memory-bound and decomposable, *can* be compressed, disaggregated and
relocated: prefill separated from decode (Dynamo GA March 2026, v1.4 August;
the Trainium-plus-Cerebras summer split), decode served from cheap LPDDR5X
pools (Positron's $875M September bet) or on-chip SRAM (Groq's LPX) or
packaging-free silicon (Taalas's HC1, February), capacity pushed toward the
telco edge (eight AI-RAN operators in field trials by September). None of that
hardware divergence pays off without the bit-level economics underneath it:
FP16→INT4 at 2–4× throughput per McKinsey, −62% per-token cost on the measured
AWS A10 configuration, model optimization driving the 85–95% cost-per-token
reduction in McKinsey's June 2026 accounting, inference at 65–90% of lifetime
operating cost making every one of those percentages a budget line rather than
a benchmark curiosity.

Three structural consequences follow, and they are the ones to carry into 2027
planning:

1. **Cost per token is the industry's unit of account.** MLPerf v6.1's scale-out framing, the 5.75M tokens/s Crusoe/MI355X record on GPT-OSS-120B, MangoBoost's disaggregated AMD results — the benchmark world now measures what operators actually buy: tokens per second per dollar, at cluster scale, on contestable hardware. Hardware vendors are priced on it, cloud providers bill on it, model providers compete on it, and agentic workloads — which multiply tokens per task by ten or a hundred — multiply its importance with every deployment. The training-era metrics (peak FLOPS, HBM capacity per rack) haven't disappeared; they've been demoted to inputs of the token equation.

2. **The format stack has stabilized into tiers, not a single winner.** FP8 is the safe default and the industry's baseline assumption; INT4-with-grouping is the aggressive standard, deployed after evals; NVFP4/MXFP4 are the rising Blackwell tier, with MXFP4 carrying the gpt-oss distribution endorsement and the Intel-measured caveat (2× worse KLD than GPTQ-INT4 on Qwen3.5-35B) proving that FP4's superiority over INT4 is hardware-dependent, not settled. AWQ is the dominant method with Marlin-class kernels underneath; TurboQuant opened the KV-cache second frontier behind an SGLang flag. Operators no longer bet on *a* format — they maintain a tiered policy with quality gates between tiers, and the gates (perplexity budgets, KLD checks, long-context evals) are as much the infrastructure as the formats.

3. **The edge became credible because the bits got small.** A 27B model at 3.9 GB (Bonsai, July 2026), a 3-bit KV cache (TurboQuant, ICLR 2026), 1.58-bit training proven at 2B scale (BitNet b1.58 2B4T, April 2025) — each step down in bits per weight is a step outward in deployment geography: from the datacenter core to colocation, to telco points of presence, to the phone in the user's hand. The AI-RAN field trials and the iPhone throughput claims are early points on the same curve, and the curve points at a 2027 in which "where can this model run" is answered in gigabytes rather than in GPU counts.

The reserve to keep is the one this chapter has flagged throughout, and the
synthesis restates it because syntheses are where caveats go to die. Vendor
claims are labeled (MangoBoost's disaggregation first, PrismML's 11 tokens/s);
single-configuration figures are bounded (the A10's −62%, TurboQuant's ×8 on
H100); single-table figures are flagged (NVFP4's "~95% retention"); unverified
hardware claims are not asserted (FP8 on MI300+); and frontier artifacts are
distinguished from production tiers (Bonsai's quality retention unestablished,
BitNet not a serving format, SpinQuant-class refinements not verified here).
What is *not* reserved — what survived every verification pass this dossier
ran — is the direction. By September 2026, the industry had stopped asking
*whether* to quantize and started asking *how far down the bit ladder* each
workload could go. The infrastructure — Dynamo's disaggregated fleets, the
LPDDR inference racks, the telco trials, the tiered format policies — was
being rebuilt around the answer. Training got denser; inference got cheaper;
and the bit was the wedge that split them.


**Working vocabulary.** A few terms recur throughout the chapter with precise
meanings worth fixing. *Prefill* is the compute-intensive ingestion of the
prompt and context in one pass; *decode* is the memory-bound, token-by-token
generation that follows — the split Dynamo schedules independently. *PTQ*
(post-training quantization) compresses an already-trained model without
retraining, as opposed to BitNet's *native* low-bit training. *Perplexity*
(PPL) is the standard language-modeling quality metric — how surprised the
model is by held-out text — and small degradations there are the industry's
unit for quantization quality cost. *KLD* (Kullback–Leibler divergence)
measures how far the quantized model's output distribution drifted from the
original's — the Intel cookbook's chosen yardstick. *Grouping* and
*microscaling* are both answers to the same problem: precision is local, so
scales are assigned per group or per block rather than per tensor.
*Disaggregated serving* is the deployment pattern — prefill and decode on
separately sized, priced and placed hardware — that turns the prefill/decode
distinction into fleet architecture.

**Open questions for 2027.** 
The chapter closes, as the dossier's method requires, not with predictions but
with the questions the verified facts leave open — the ones a reader planning
2027 infrastructure should be tracking.

First, does the training power-density curve hold? The NVL144's 600 kW is
announced/planned for 2027; whether facilities absorb it or force a plateau
determines whether training efficiency — the one workload quantization can't
rescue — returns to the agenda.

Second, does 4-bit become the new safe tier? INT4 spent 2026 accumulating
runbooks, kernels and eval-gate procedures. If that maturation completes, the
industry's baseline assumption shifts down another factor of two, and FP8
becomes what FP16 is now: the fallback.

Third, does KV-cache quantization standardize the way weight quantization did?
TurboQuant is one method behind one flag in one runtime. A standardized,
multi-runtime 3-bit cache — the cache equivalent of what AWQ did for weights —
would be the next structural cost step.

Fourth, does ternary training reach production? BitNet proved trainability;
Bonsai proved post-training ternary compression of an existing model. The gap
between "provable" and "deployed at scale" is where 2027's most interesting
infrastructure bets will sit — including, in principle, the ternary-simplified
silicon the arithmetic implies.

| Key figures at a glance (all verified) | Value |
|---|---|
| Rubin NVL72 power | ~190–230 kW (Blackwell GB200 NVL72: ~100–120 kW — not Rubin's) |
| NVL144 power | 600 kW, planned 2027 |
| MLPerf v6.1 multi-node submissions | 16 (record); 16/09/2026 |
| Crusoe/MI355X record | 512 GPUs, 5.75M tokens/s offline, GPT-OSS-120B |
| Dynamo GA / v1.4 | ~GTC March 2026 / August 2026 |
| Positron Series C | $875M, 11/09/2026, LPDDR5X inference-only |
| Taalas HC1 | 20/02/2026, no HBM, no CoWoS |
| AI-RAN trials | 8 operators, September 2026 (from GTC March 2026) |
| Inference share of LLM operating cost | 65–90% (McKinsey AI Economics Report 2025) |
| A10 FP16 → INT4 GPTQ | $8.12 → $3.10/M tokens, −62% (AWS, Nov 2025) |
| McKinsey model-optimization figure | 85–95% cost-per-token reduction (June 2026, verified quote) |
| FP16→INT4 throughput | 2–4× (McKinsey) |
| FP8 quality | ~99% retained (working figure); Hopper/Blackwell acceleration |
| INT4 memory | −75% vs FP32 |
| AWQ salient weights | ~1% protected; 0.5–1.5% perplexity cost; ~3× speedup vs FP16 (contextual) |
| BitNet paper / bitnet.cpp 1.0 / b1.58 2B4T | 27/02/2024 / 17/10/2024 / April 2025 |
| Bonsai | 14/07/2026; 27B → 3.9 GB (1.125 bits/wt), ternary 5.9 GB |
| TurboQuant | arXiv 2504.19874 (April 2025); ICLR 2026; ≥6× KV memory, ×8 attention (bounded configs); SGLang flag |


**A final methodological note.** Three figures in this chapter survived
explicit hallucination suspicion during verification — McKinsey's 85–95% cost-
per-token claim, PrismML Bonsai's existence, and TurboQuant's paper trail —
and the text records each clearance where the figure appears, rather than
silently asserting it. That discipline is the chapter's contract with the
reader: the numbers that look too dramatic to be true are the ones whose
verification is documented, and the numbers that couldn't be verified (FP8 on
MI300+, Bonsai's fine benchmarks, NVFP4's single-table retention) are flagged
rather than smoothed over. A reference dossier is only as trustworthy as its
handling of the facts it cannot confirm.

**How to use this chapter.** As a reference: each section stands alone — the tables collect the verified figures, the reserve flags mark what not to quote. As analysis: the through-line is the one stated in 5.15, that quantization is the wedge that split training from inference. And as a planning input for 2027: the tiered-policy decision tree at the end of 5.14 is the operational takeaway, the open questions above are the watch list, and the key-figures table is the numbers sheet. The facts are dated September 22, 2026; the direction they point is the reader's to judge.<a id="g06"></a>
