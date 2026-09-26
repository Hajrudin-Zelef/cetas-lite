---
id: ai-industry-kb-2026/06-inference-engines/sources-and-urls-continued
title: "Sources and URLs (continued)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "DeepSeek", "Huawei", "Moonshot", "Nvidia", "SGLang", "vLLM", "xAI"]
dates: ["2025-10-29", "2026-01", "2026-05", "2026-05-04", "2026-05-12", "2026-08", "2026-08-26", "2026-09", "2026-09-04", "2026-09-20"]
keywords: ["agent", "agentic", "amd", "ascend", "attention", "benchmark", "benchmarks", "blackwell", "datacenter", "deepseek", "distribution", "funding"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2942, 2991]
section: "6. Inference Engines"
sha256: db6e058647b5cd2f07c6ecaed89ace7193a89430ffd9b5a700a7b61207653fa8
---

# Sources and URLs (continued)

1. **Engine selection is workload-shaped, not per-religion:** the same-kernel +29% SGLang win, vLLM's TTFT/100+-concurrency counter-evidence, and the operator serving DeepSeek V4 Flash on vLLM while keeping SGLang elsewhere all point to per-workload, per-model choice behind a gateway — not engine loyalty.
2. **The `lpm` opt-in gap is the cheapest win in the SGLang track:** cache-aware scheduling (`--schedule-policy lpm`/`dfs-weight`) is off by default and likely underused in the wild; `--enable-cache-report` turns hit rates into an observable metric before any tuning spend.
3. **Don't conflate throughput gains with cache-hit ratios:** the brief's "2.5x cache hit rate vs competition" was a misread of a 2.5x jump-forward throughput figure — the consolidation should cite 3.8x-vs-round-robin router figures and the paper's ~96%-of-optimal figure instead.
4. **Commercialization changes procurement calculus:** RadixArk ($100M, May 2026) narrows SGLang's enterprise-packaging gap vs Red Hat/IBM vLLM; as of September 2026 the comparison is a two-commercial-stack one, still favoring vLLM's shipped server — revisit after RadixArk's paid-hosting tiers materialize.
5. **"4x batch size per GPU" is a planning figure, not a benchmark:** until a controlled OOM/SLO-breaching batch test is published, budget MLA capacity with the paper's 93.3% KV reduction as the ceiling and treat ~4x as conservative-derivative guidance.
6. **Time-to-serve is the 2026 differentiation metric:** both vLLM v0.28.0 and SGLang v0.5.19 raced the same flagships (Qwen3.8, Kimi-K3) across NVIDIA/AMD/Ascend in the same month — watch release-note race dynamics, not just architecture.
7. **Measure on your hardware class:** the August 2026 TRT-LLM vs vLLM measurement and the SM121 NaN-bisecting community patches show workstation-Blackwell results don't transfer from datacenter-Hopper claims — the SM90+-first reality of 2026 MLA/sparse kernels demands per-silicon validation.
8. **The post-training feedback loop is SGLang's durable moat:** verl/AReaL/slime/Tunix run on SGLang, optimizations land there first, frontier open-weight models serve best there — this is where tomorrow's production models are born, and vLLM cannot match it through distribution alone.
9. **Date-stamp every version claim:** vLLM moved v0.27 → v0.28.0 (2026-08-26) with v0.29.0 already attempted; SGLang v0.5.18 → v0.5.19 (2026-09-04) — any "latest version" claim without a date rots within weeks.
10. **Structural decoding is commoditized; the edge is forward-pass elimination:** XGrammar is universal in 2026 — SGLang's differentiator is specifically jump-forward decoding, and XGrammar-2's Structural Tag is the protocol the agent stack (xAI, Databricks, DeepSeek) is converging on.
11. **January 2026 priced the duopoly's commercial layer:** $100M-at-$400M (SGLang/RadixArk) vs $150M-at-$800M (vLLM/Inferact) on the same day — future buyers should read engine choice against these two commercial stacks, not against "open vs proprietary."
12. **Prefix discipline is an operator skill:** with hit rates spanning 50–99% on prompt design, the 2026 practitioner's first tuning pass is prompt templating and `--enable-cache-report` instrumentation — engine upgrades come second.

## Sources and URLs (continued)

- [VENDOR] https://www.businesswire.com/news/home/20260505077157/en/RadixArk-Launches-with-%24100-Million-in-Seed-Funding-Led-by-Accel-to-Grow-SGLang-and-Democratize-Frontier-AI-Infrastructure
- [VENDOR] https://github.com/sgl-project/sglang/releases/tag/v0.5.19
- [VENDOR] https://github.com/sgl-project/sglang-omni/releases/tag/v0.1.4
- [VENDOR] https://github.com/lm-sys/lm-sys.github.io/blob/HEAD/blog/2025-10-29-sglang-jax.md
- [VENDOR] https://github.com/sgl-project/sglang/blob/HEAD/docs/docs/hardware-platforms/tpu.mdx
- [VENDOR] https://github.com/sgl-project/sglang/pull/30514
- [VENDOR] https://github.com/vllm-project/vllm/releases/tag/v0.28.0
- [VENDOR] https://github.com/mlc-ai/blog/blob/HEAD/_posts/2026-05-04-xgrammar-2-fast-customizable-structured-generation.md
- [VENDOR] https://unsloth.ai/docs/basics/unsloth-dynamic-2.0-ggufs
- [VENDOR] https://huggingface.co/unsloth/Qwen3-0.6B-GGUF
- [VENDOR] https://openreview.net/pdf?id=VqkAKQibpq
- [DIRECTIONAL] https://www.morningstar.com/news/business-wire/20260505077157/radixark-launches-with-100-million-in-seed-funding-led-by-accel-to-grow-sglang-and-democratize-frontier-ai-infrastructure
- [DIRECTIONAL] https://pulse2.com/radixark-launches-with-100-million-in-seed-funding-to-democratize-frontier-ai-infrastructure/
- [DIRECTIONAL] https://techcrunch.com/2026/01/22/inference-startup-inferact-lands-150m-to-commercialize-vllm/
- [DIRECTIONAL] https://pulse2.com/inferact-launches-with-150-million-funding-at-800m-valuation-to-commercialize-vllm-as-inference-demand-surges/
- [DIRECTIONAL] https://theaiinsider.tech/2026/01/27/inferact-launches-with-150m-seed-round-to-commercialize-vllm-inference-engine/
- [DIRECTIONAL] https://www.bbw9n.io/blog/sglang_vs_vllm
- [DIRECTIONAL] https://privocto.com/blog/vllm-sglang
- [DIRECTIONAL] https://felloai.com/it/deepseek-v4/
- [DIRECTIONAL] https://apidog.com/blog/what-is-deepseek-v4/
- [COMMUNITY] https://explore.n1n.ai/blog/sglang-vs-vllm-architecture-radixattention-benchmarks-2026-09-20
- [COMMUNITY] https://github.com/xianlubird/sglang
- [COMMUNITY] https://github.com/rishikinger10/radixscope/blob/HEAD/Docs/sglang-integration.md
- [COMMUNITY] https://github.com/softwealth/eval-report-skills/blob/HEAD/skills/inference/xgrammar-structured-output.md
- [COMMUNITY] https://github.com/tabnas/parser/blob/HEAD/ts/doc/gbnf-feasibility.md
- [COMMUNITY] https://github.com/jedmund/homelab/blob/HEAD/roles/sglang/README.md
- [COMMUNITY] https://github.com/profsynapse/synaptic-tuner/blob/HEAD/docs/preparation/vllm-vs-sglang-inference-serving-research.md
- [COMMUNITY] https://github.com/shurankain/agentic-ai-course/blob/HEAD/17_Production_Inference/07_SGLang_and_Alternatives.md
- [COMMUNITY] https://github.com/iopsystems/llm-calc/blob/HEAD/docs/superpowers/specs/2026-05-12-dsa-design.md
- [COMMUNITY] https://github.com/yusanxy/llm_flops/blob/HEAD/operators/references/deepseek_v32_dsa_sparse_attention/README.md
- [COMMUNITY] https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc21-released-nvidia-gb300-achieves-moe-w-20260728
- [COMMUNITY] https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc23-released-amd-mi450-nvidia-rtx-5090-o-20260731
- [COMMUNITY] https://github.com/dcharlot-physicalai-bmi/ferric/blob/HEAD/docs/runtime-landscape-2026-08.md
- [COMMUNITY] https://data.safetycli.com/packages/pypi/tensorrt-llm/changelog?page=2

