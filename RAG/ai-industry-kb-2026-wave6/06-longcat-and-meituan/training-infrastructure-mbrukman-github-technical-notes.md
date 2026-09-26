---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/training-infrastructure-mbrukman-github-technical-notes
title: "Training infrastructure (mbrukman GitHub technical notes)"
domain: longcat-and-meituan
role: deep-dive
task: training
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "LongCat", "MiniMax", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2025-08-29", "2025-09-05", "2025-09-22", "2025-12-22", "2026-01-14", "2026-02-05", "2026-03", "2026-03-11", "2026-03-12", "2026-05-29", "2026-08"]
keywords: ["training", "agent", "agentic", "alignment", "attention", "benchmarks", "claude", "compute", "cost", "decode", "deepseek", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2668, 2732]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: 1727502469cc349574743d3097c1883a9d97bd4fd6c6c26de0fa484aa6e12420
---

# Training infrastructure (mbrukman GitHub technical notes)

### Training infrastructure (mbrukman GitHub technical notes)
- **Determinism & reliability**: enforced determinism for training reproducibility; numerical reliability via optimized foundational operators; **automated monitoring for seamless fault recovery** to secure stable production operations [SECONDARY] (github.com/mbrukman/longcat-2.0).
- **Training at scale**: 6D parallelism integrated with **super-node architectures**; multi-dimensional memory optimizations; **pioneering large-scale deployment of a customized Muon optimizer** [SECONDARY].
- **Long-context training**: optimized in-house operators; context scaled to 1M via an **all-gather-based CP (context-parallelism) scheme**; compute-communication overlap to minimize synchronization overhead [SECONDARY].
- Inference optimization on domestic superpod accelerators: **indexer pipelining and KV-cache parallelism** to mitigate KV-cache overhead; **explicit per-core control** for fully parallel dense+MoE execution; super kernels; L2 weight prefetching; high-speed interconnects for scale-up/out [SECONDARY].
- Serving: prefill-decode disaggregation with tailored schemes — **CPP and SP for prefill, KVP and large EP for decode** — plus **asynchronous load balancing** against stage-specific bottlenecks [SECONDARY].

### Pricing and positioning (cryptobriefing)
- **SWE-bench Pro 59.5** edges **GPT-5.5's 58.6** and surpasses Claude Opus variants on the same test [SECONDARY] (cryptobriefing.com/meituan-longcat-2-undercuts-gpt-claude-pricing/).
- API pricing **$0.75/M input / $2.95/M output** — "almost confrontational" vs Western frontier rates [SECONDARY].
- **1.6T total, 33–56B dynamic activation (~48B average)**; native 1M context [SECONDARY].
- **50,000-chip domestic cluster** — "the largest reported training run using non-NVIDIA hardware" [SECONDARY] (cryptobriefing.com/china-meituan-frontier-ai-domestic-chips/).
- Second cryptobriefing piece: pretrained on **over 30T tokens** (vs 35T+ in other coverage — record the variance); **Terminal-Bench 2.1: 70.8** (cryptobriefing's "Terminal Bench 2" normalized to the vendor-confirmed 2.1), "competitive territory with leading models globally" [SECONDARY].

### Community field review (tamago-labs, August 2026)
- **"Secretly topped OpenRouter for two months"** — the model's stealth popularity claim [COMMUNITY] (medium.com/tamago-labs).
- Pricing model: **no monthly plan, no subscription — token packs like mobile games with massive discounts; a starter pack for ~$2** [COMMUNITY].
- The reviewer's team tried **Claude, GLM, and MiniMax** before landing on LongCat-2.0 via OpenCode — direct cross-lab comparison shopping [COMMUNITY].
- Recap figures: 1.6T MoE, ~48B active, LSA (evolution of DeepSeek's sparse attention), **35+T tokens on 50K+ Chinese ASICs, no Nvidia**; integrates **Claude Code, OpenClaw, Hermes**; benchmarks **59.5 SWE-bench Pro / 70.8 Terminal-Bench 2.1 / 77.3 SWE-bench Multilingual**; three expert groups via multi-teacher distillation [COMMUNITY].

### Dynamic activation vs fixed MoE (ayinedjimi-consultants.fr, French analysis)
- The French analysis traces the design to **DenseNet/highway-network lineage (2015–2017)** adapted to trillion-parameter MoE scale [SECONDARY].
- MoE comparison table: **LongCat-2.0 1600B total / 33–56B dynamic (avg 48B)** vs **DeepSeek-V3 671B / 37B fixed (256 experts, top-8)** vs **Mixtral 8×22B 141B / 39B fixed (top-2/8)** vs **Qwen3-235B 235B / 22B fixed (128 experts, top-8)** — LongCat is the only one with dynamic activation [SECONDARY].
- Practical consequence: per-token "cost" is no longer directly measurable from token counters — **you need the average query complexity**; the model is economical on simple tasks while keeping deep reasoning for complex ones [SECONDARY analysis].


### New verified facts — expansion (continued — LongCat platform changelog history (deprecations resolved))

### The May 29, 2026 six-model sunset — exact list recovered
The LongCat API Platform Change Log (official) gives the exact six models retired **effective May 29, 2026** [VENDOR] (longcat.chat/platform/docs/ChangeLog.html):
1. **LongCat-Flash-Chat**
2. **LongCat-Flash-Thinking**
3. **LongCat-Flash-Thinking-2601**
4. **LongCat-Flash-Omni-2603**
5. **LongCat-Flash-Lite**
6. **LongCat-Flash-Chat-2602-Exp**
- Reason given: after the **LongCat-2.0-Preview** launch, demand for service resources kept growing; retiring legacy models consolidates resources for 2.0-Preview testing and iteration [VENDOR].
- Migration path: users were invited to apply for **LongCat-2.0-Preview**; a limited number of beta slots released daily at **09:00:00 and 21:00:00 (UTC+8)**, first-come first-served, availability expanding gradually [VENDOR].
- Note: the sunset list includes models released as recently as March 2026 (Omni-2603) — a **~2.5-month API lifecycle** for some checkpoints [DIRECTIONAL].

### LongCat-Flash-Lite (2026-02-05)
- **68.5B total parameters, ~3B activated per inference** — highly efficient MoE, specifically optimized for inference efficiency and specialized use cases [VENDOR] (ChangeLog).
- Efficiency mechanism: **N-gram embedding table** mitigating I/O bottlenecks within MoE layers, plus specialized cache and kernel optimizations — reduces latency, boosts efficiency [VENDOR].
- Agentic and coding performance described as competitive relative to model scale [VENDOR].
- Open-sourced on Hugging Face: `meituan-longcat/LongCat-Flash-Lite` [VENDOR].

### LongCat-Flash-Omni-2603 (2026-03-11)
- Upgraded version of LongCat-Flash-Omni: **end-to-end Omni interaction model** with more human-like responsiveness and stronger full-modality perception [VENDOR] (ChangeLog).
- Improvements over predecessor: natural dialogue flow via **deep semantic alignment and personalized style adaptation**; improved accuracy across **vision, speech, and text** multimodal tasks; enhanced complex problem-solving, emotional understanding, casual entertainment [VENDOR].
- **Native voice Function Calls**: direct parsing of audio commands with real-time interaction at near-zero latency [VENDOR].

### LongCat-Flash-Thinking auto-upgrade (2026-03-12)
- On **2026-03-12 at 20:00:00 (UTC+8)**, all `LongCat-Flash-Thinking` API requests were **automatically routed to LongCat-Flash-Thinking-2601** — no code changes required [VENDOR] (ChangeLog).
- ChangeLog's 2026-01-14 release entry for 2601 adds: **560B total MoE**; large-scale multi-environment RL; **environment dependency maps with over 60 types of tools**; multi-environment expansion and large-scale exploration training for OOD generalization [VENDOR].
- **Advanced Deep Thinking Mode**: parallel reasoning to expand thinking breadth + a **summarization mechanism with recursive feedback** to expand thinking depth [VENDOR].
- **Extreme robustness in noisy environments** via systematic curriculum training on noise/uncertainty in real-world environments — cross-confirms the noise-injection strategy from the earlier block [VENDOR].
- GitHub: `meituan-longcat/LongCat-Flash-Thinking-2601` [VENDOR].

### LongCat-Flash-Chat evolution
- **2025-08-29**: LongCat-Flash-Chat released and open-sourced; free on LongCat Chat [VENDOR].
- **2025-09-05**: LongCat API Platform launched with Flash-Chat calls [VENDOR].
- **2025-09-22**: LongCat-Flash-Thinking released and open-sourced [VENDOR].
- **2025-12-22**: Flash-Chat upgraded in place (same model name/API): **context doubled from 128K to 256K**; significantly enhanced programming (code generation, debugging, explanation); **high-quality support for nine languages** (Spanish, French, Arabic, Portuguese, Russian, Indonesian among them); more robust agent capabilities [VENDOR].
- This in-place upgrade pattern (same name, new weights) contrasts with the dated-checkpoint naming of the 2601/2602/2603 releases [DIRECTIONAL].

