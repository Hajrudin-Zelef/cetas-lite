---
id: briefing-general-tech-2026/02-gpus-accelerators/03-mlperf-inference-v61
title: "MLPerf Inference v6.1: Rubin's first benchmark outing"
domain: gpus-accelerators
role: deep-dive
task: benchmark
actors: ["AMD", "CoreWeave", "MLCommons", "Nebius", "Nvidia", "OpenAI"]
dates: ["2026-03", "2026-09-16", "2026-09-22"]
keywords: ["benchmark", "inference", "mlperf", "blackwell", "cost per token", "decode", "disaggregated", "disaggregated serving", "gpu", "gpus", "helios", "kv cache"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-3"
source_lines: [1582, 1667]
canonical_for: ["mlperf-v61"]
sha256: f4a5fc936241b8f94642d58ee184e2cc86c1241eeedeac13f73accc6cfb35eec
---

# MLPerf Inference v6.1: Rubin's first benchmark outing

<a id="g03-3"></a>
### 3.3 MLPerf Inference v6.1: Rubin's first benchmark outing

MLPerf Inference v6.1, published by MLCommons on September 16, 2026 with results from 30 organizations, was the first independent benchmark round to include the Vera Rubin platform. Nvidia submitted Rubin results the same day — entries 6.1-0106 and 6.1-0074 in the Closed Division — and published a company blog the same day under an unambiguous headline: "Vera Rubin NVL72 Delivers Leading Performance in MLPerf Inference v6.1 Debut."

The framing Nvidia chose is worth quoting precisely: the company called the submission a **"preview submission"**, explicitly noting that Rubin was not yet shipping in volume at the time of publication. That honesty matters because MLPerf Closed Division results are supposed to compare like-for-like hardware — and a preview of a not-yet-shipping platform sits in a gray zone between a launch benchmark and a paper launch. The numbers are real measurements on real silicon; the availability is announced, not delivered.

#### The headline results

All comparisons below are versus the previous-generation GB300 NVL72:

- **Up to 3.7x token throughput on Qwen3-VL** — 1,307 queries/s versus 349 — measured with vLLM plus Nvidia Dynamo.
- **Up to 2.5x on DeepSeek-R1**, measured with TensorRT-LLM.

Both are large margins, and both are multimodal-or-MoE-flavored workloads — exactly the workload classes where Nvidia's software stack changes do the most work. The company attributed the gains to a software bundle rather than raw silicon: **NVFP4** applied across weights, attention, and KV cache; **disaggregated serving** splitting prefill and decode onto different resources; and **expert parallelism** for mixture-of-experts models. A caution on the first item: the verified sources describe NVFP4 as the quantization format in play, but no source in this verification round measured or claimed a specific KV-cache precision reduction — that detail should not be asserted as a measured result.

| MLPerf Inference v6.1 — Rubin NVL72 (Closed Division) | Result | Baseline (GB300 NVL72) | Multiple |
|---|---|---|---|
| Qwen3-VL token throughput (vLLM + Dynamo) | 1,307 queries/s | 349 queries/s | ~3.7x |
| DeepSeek-R1 throughput (TensorRT-LLM) | — | — | ~2.5x (vendor-reported) |
| Submitter entries | 6.1-0106, 6.1-0074 | — | — |
| Status | "Preview submission" (not yet shipping) | — | — |

#### Queries per second: what the unit measures

MLPerf Inference reports throughput in queries (or samples) per second — completed inference requests, not raw FLOPS. The 1,307 vs. 349 queries/s on Qwen3-VL therefore measures *system* throughput: the GPU, the memory system, vLLM, and Dynamo working together to serve multimodal requests. Queries per second is the buyer's unit — it answers "how many users can this serve?" — which is why Nvidia headlined it rather than a FLOPS figure. The 3.7x multiple inherits the unit's honesty and its limits: it is workload-specific (Qwen3-VL, a specific configuration), stack-specific (vLLM + Dynamo), and generation-specific (vs. GB300 NVL72). Change any of the three and the multiple changes — which is why the dossier reports all three alongside the number.

#### Reading the software stack behind the numbers

Each of the three software items Nvidia credited deserves a plain-language gloss, because the 3.7x/2.5x multiples are software multiples as much as hardware ones:

- **NVFP4 quantization** (weights, attention, KV cache): running the model at 4-bit precision shrinks memory footprint and multiplies effective bandwidth — the single largest lever in memory-bound inference.
- **Disaggregated serving** (prefill/decode split): separating the prompt-processing phase from the token-generation phase onto different resources lets each be scaled and scheduled independently, rather than forcing one configuration to serve both.
- **Expert parallelism** for MoE: distributing a mixture-of-experts model's experts across devices so that only the activated experts consume resources per token.

The point of listing them is not to endorse the multiples but to locate them: a buyer reading "3.7x" should understand that a substantial share of it comes from the software stack (vLLM + Dynamo, TensorRT-LLM) that Nvidia co-develops, not from the Rubin die alone. That is not a criticism — the stack is part of the product — but it is the correct attribution, and it explains why AMD's response has emphasized its own software (see §3.6) rather than conceding the benchmark outright.

#### The serving stacks, decoded

The MLPerf multiples were measured *through* specific software stacks, and the stacks are worth identifying:

- **vLLM** is the open-source inference serving engine that became the industry's default high-throughput server — itsPagedAttention memory management is the reason long-context serving is economical at all.
- **Nvidia Dynamo** is Nvidia's open-source inference-serving framework: the disaggregated-serving layer that splits prefill and decode across resources and orchestrates them. The Qwen3-VL result (1,307 queries/s) was measured as vLLM *plus* Dynamo — the combination, not either alone.
- **TensorRT-LLM** is Nvidia's inference optimizer and runtime — the DeepSeek-R1 result's stack.

The pattern: Nvidia's benchmark wins run on Nvidia's (or Nvidia-adjacent) software. That is not a methodological flaw — MLPerf's Closed Division permits it — but it means the multiples measure the *platform* (silicon + stack), and a buyer without the stack gets less than the multiple. It also explains AMD's strategic imperative: matching Rubin requires matching Dynamo and TensorRT-LLM as much as matching the MI455X die, which is why the software-readiness evidence (OpenAI's April pilots, §3.7) matters as much as the spec sheet.

#### The cost-per-token envelope: vendor claims, March vintage

The cost-per-token framing that accompanied the results came from further back: Nvidia's GTC March 2026 claims that Rubin would deliver **MoE training with one-quarter the GPUs of Blackwell**, and **10x inference throughput per watt at one-tenth the cost per token**. Those are **vendor claims from March, not independently verified measurements**, and they should be read as the marketing envelope around the September numbers, not as the numbers themselves. The September MLPerf results measure throughput; they do not measure cost per token, watts per token, or training-GPU counts. The GTC claims remain untested against third-party measurement as of September 22, 2026.

#### Context: Nebius, CoreWeave, and the multi-rack milestone

Two contextual notes complete the picture. First, Nvidia was not the only Rubin submitter: **Nebius also submitted results on a VR200 NVL72 system**, an early sign that the neocloud channel Nvidia leaned on for the H2 ramp is benchmarking, not just buying. A neocloud submitting to MLPerf is a commercial signal — it says "our Rubin capacity is performant and available," which is marketing to the next customer as much as it is engineering.

Second, the same day — September 16, 2026 — CoreWeave confirmed an official **multi-rack 504-GPU Rubin cluster**, the first public evidence of Rubin operating beyond a single rack. Seven racks of NVL72 running as one fabric is the scale at which the platform's economics actually get tested: multi-rack is where interconnect behavior, failure domains, and scheduling efficiency stop being spec-sheet items and start being operational realities. It arrived three months after the first single rack went into production — a fast escalation from one rack to seven, and a preview of the H2 2026 deployment pattern.

#### What to watch

- **General-availability MLPerf submissions**: the "preview" label comes off when Rubin ships in volume; expect resubmissions with shipping firmware.
- **AMD's answer**: Helios has no MLPerf presence as of September 22 (no shipped racks to test). The first Helios submission — whenever it lands — will be the first third-party-comparable data point in the Rubin-vs-Helios contest.
- **Cost-per-token verification**: the GTC 10x/1/10th claims await any independent measurement.

---

#### MLPerf in one paragraph

For readers outside the benchmarking world: MLPerf Inference is the industry-standard AI inference benchmark suite administered by MLCommons, a multi-vendor consortium. The **Closed Division** — the division Nvidia submitted in — requires submitters to run equivalent models under fixed rules, so results are comparable across hardware; it is the division the press treats as the "official" scoreboard. Results are published in versioned rounds (v6.1 here), with entries numbered (6.1-0106, 6.1-0074) and attributed to submitters. Thirty organizations submitted to v6.1, which makes it one of the larger rounds — and makes Rubin's debut in it the closest thing 2026 offered to a neutral measurement of the new platform.

#### The workloads, decoded

The two headline workloads were chosen by Nvidia, not by MLCommons, and the choice is informative:

| Workload | Type | Why Nvidia chose it |
|---|---|---|
| Qwen3-VL | Vision-language model | Multimodal: exercises the full stack (vision encoder + LLM), where Dynamo's scheduling helps most |
| DeepSeek-R1 | MoE reasoning model | Sparse expert activation: where expert parallelism and NVFP4 pay off most |

Both are the workload classes where software matters most — which is consistent with Nvidia crediting the software stack rather than raw silicon. A skeptic's reading: Nvidia benchmarked where its software advantage is largest. A fair reading: those workload classes (multimodal, MoE reasoning) are also where the industry's growth is, so the advantage is pointed at the relevant target. Both readings can be true, and the dossier records the workload selection alongside the multiples so readers can apply their own discount.

#### "Preview submission": the gray zone, examined

Nvidia's "preview submission" label deserves a final paragraph because it will recur as other vendors' platforms reach MLPerf. A preview submission is a real measurement on real silicon, run under Closed Division rules — it is not a simulation. What it is not is a *shipping* product's measurement: firmware, drivers, and system tuning may all change before volume delivery, and historically they usually improve (early silicon is rarely the fastest silicon). The practical consequence: treat preview multiples as a floor for the architecture's potential and a ceiling for launch-day claims — the truth at general availability typically lands between, usually closer to the preview than to the marketing envelope.

---

