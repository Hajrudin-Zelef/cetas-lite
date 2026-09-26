---
id: ai-industry-kb-2026/09-moe-architectures/part-24
title: "9. MoE Architectures (part 24)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Huawei", "Hugging Face", "Moonshot", "Nvidia", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-10", "2026-05", "2026-07-22", "2026-08-26", "2026-09-17"]
keywords: ["moe", "ascend", "attention", "benchmark", "blackwell", "claude", "fine-tuning", "glm", "gpu", "inference", "kimi", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5357, 5386]
section: "9. MoE Architectures"
sha256: 2ad32c44707159737855facbdfd10cbffd8c5a4ffdd023775c118137eeefa862
---

# 9. MoE Architectures (part 24)

- NVIDIA dense-vs-MoE technical blog: https://developer.nvidia.com/blog/dense-vs-moe-models-active-parameters-throughput-and-when-to-choose-each/ [VENDOR]
- MoE KV-cache economics research notes: https://github.com/eduardogbg/mnemo/blob/HEAD/docs/research/dag-kv-cache-economics.md
- MoE inference optimization (Spheron): https://www.spheron.network/blog/moe-inference-optimization-gpu-cloud/
- vLLM-Ascend expert parallelism load balancer: https://github.com/vllm-hust/vllm-ascend-hust/blob/HEAD/docs/source/user_guide/feature_guide/expert_parallelism_load_balancer.md
- vLLM Mixtral MoE optimization case: https://github.com/ineshreddy249/vllm-mixtral-moe-optimization/blob/HEAD/README.md
- ROCm vLLM inference optimization (EP topology guidance): https://github.com/yusrildestroy1/rocm/blob/HEAD/docs/how-to/rocm-for-ai/inference-optimization/vllm-optimization.rst
- vLLM distributed inference cookbook: https://github.com/wahajsayyed/ai-system-engineering/blob/HEAD/vLLM-cookbook/chapter-07-distributed-inference.md
- MoE tutorial (ESFT, aux-loss discussion): https://github.com/wanshuiyin/auto-claude-code-research-in-sleep/blob/HEAD/docs/tutorials/moe_tutorial_en.md
- MoE at Blackwell lecture notes: https://github.com/kingarthurv1/ai-hardware-engineer-roadmap/blob/HEAD/Phase%205%20-%20Advanced%20Topics%20and%20Specialization/7.%20ML%20Systems%20Engineering/AI%20Inference%20Engineer%202026/Part%203%20-%20MoE%20at%20Blackwell/Lecture-01.md [COMMUNITY]
- MoE vs dense serving profile (OLMoE): https://github.com/malcomzww/moe-vs-dense-serving-profile [COMMUNITY]
- Controlled MoE vs dense study: https://github.com/oliversundaram/moe-study [COMMUNITY]
- GLM-5.3-Flash release (MarkTechPost, 2026-08-26): https://www.marktechpost.com/2026/08/26/z-ai-releases-glm-5-3-flash-a-320b-a18b-natively-multimodal-moe-with-a-1m-token-context/
- GLM-5.3-Flash model page (Atomic): https://atomic.chat/models/glm-5-3-flash
- GLM-5.3-Flash analysis (LumaDock): https://lumadock.com/blog/glm-5-3-flash
- Kimi K3 breakdown (Medium, secondary): https://medium.com/@jamilxt/kimi-k3-is-now-open-weight-the-full-breakdown-of-the-2-8t-frontier-model-65e7033aa917 [COMMUNITY]
- Kimi K3 concept wiki: https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/concepts/kimi-k3.md [COMMUNITY]
- Kimi K3 technical report (primary, Moonshot): https://raw.githubusercontent.com/MoonshotAI/Kimi-K3/master/k3_tech_report.pdf
- Kimi K3 overview (ai-stack): https://ai-stack.ai/en/kimi-k3-moonshot-ai-open-source
- MoE Parallel Folding (arXiv 2504.14960v2): https://arxiv.org/pdf/2504.14960v2.pdf
- Expert-parallel dispatch two-phase analysis (arXiv 2503.04398): https://arxiv.org/pdf/2503.04398
- llm-d wide expert parallelism docs: https://github.com/llm-d/llm-d.github.io/blob/HEAD/versioned_docs/version-0.7/well-lit-paths/wide-expert-parallelism.md
- torchtune paper, May 2026 (independent framework comparison): https://arxiv.org/pdf/2605.21442v1.pdf
- MarkTechPost 4-framework comparison, 2026-07-22: https://www.marktechpost.com/2026/07/22/unsloth-vs-axolotl-vs-trl-vs-llama-factory-a-fine-tuning-framework-comparison-on-speed-vram-and-multi-gpu/
- Hugging Face Unsloth–TRL benchmark blog: https://huggingface.co/blog/unsloth-trl
- DeepSpeed communication optimization (MoE all-reduce penalties): https://github.com/deepspeedai/deepspeed/blob/HEAD/blogs/comm-opt/README.md
- DeepSpeed AutoEP blog, 2026-02-10: https://deepspeed.ai/blog/2026/02/10/AutoEP/
- Unsloth Studio v0.1.501-beta release-body fixture (MoE expert offload to system RAM): https://github.com/unslothai/unsloth/blob/HEAD/tests/studio/fixtures/release_bodies/v0.1.501-beta.md
- Unsloth Qwen3.8 fine-tuning guide (Sept 2026; VRAM tiers, Flash Linear Attention kernels): https://unsloth.ai/docs/models/qwen3.8/train [VENDOR]
- Unsloth official changelog (latest entry 2026-09-17): https://unsloth.ai/docs/new/changelog [VENDOR]

