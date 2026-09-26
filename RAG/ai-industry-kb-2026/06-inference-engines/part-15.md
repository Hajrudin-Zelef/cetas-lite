---
id: ai-industry-kb-2026/06-inference-engines/part-15
title: "6. Inference Engines (part 15)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Groq", "Hugging Face", "Nvidia", "SGLang", "Stripe", "TensorRT-LLM", "vLLM"]
dates: ["2025-05", "2026-01-22", "2026-01-27", "2026-02-02", "2026-03", "2026-03-20", "2026-04-28", "2026-06-05", "2026-07-12", "2026-07-14", "2026-08-17", "2026-08-26", "2026-08-29", "2026-09-01", "2026-09-09", "2026-09-12"]
keywords: ["inference", "inference engine", "agentic", "aws", "capex", "claude", "cost", "decode", "deepseek", "diffusion", "disaggregated", "disaggregated serving"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2549, 2597]
section: "6. Inference Engines"
sha256: 0eda7d693daea46fc8ee9a731c1a94cefd9e78dcb0e632843fc46450bf1fcb4e
---

# 6. Inference Engines (part 15)

- vLLM v0.28.0 release notes (tagged 2026-08-26; 584 commits, 270 contributors): https://github.com/vllm-project/vllm/releases/tag/v0.28.0
- vLLM v0.26.0 release notes (DeepSeek-V4 push, KV offloading maturation): https://github.com/vllm-project/vllm/releases/tag/v0.26.0
- vLLM v0.25.0 release notes (Model Runner V2 default, PagedAttention removed, Streaming Parser Engine): https://github.com/vllm-project/vllm/releases/tag/v0.25.0
- vLLM releases page: https://github.com/vllm-project/vllm/releases/
- vLLM v0.20.0 release notes (CUDA 13.0, PyTorch 2.11, Transformers v5, DeepSeek V4, TurboQuant): https://releaseport.com/r/vllm-project-vllm/v0-20-0
- vLLM V0 deprecation RFC #18571 (v0.9 last V0 / v0.10 removal start / v0.11 first without V0): https://github.com/vllm-project/vllm/issues/18571
- vLLM V1 guide, "We have fully deprecated V0": https://docs.vllm.ai/en/latest/usage/v1_guide/
- Forensic code review confirming V0 removal in v0.16.0 (March 2026): https://github.com/kender242/portfolio_dev/blob/HEAD/docs/vLLM_016_REVIEW.md
- "vLLM 0.25 Deletes PagedAttention" analysis (kenashe, 2026-07-12): https://github.com/kenashe/kenashe/blob/HEAD/src/content/blog/2026-07-12-vllm-0-25-deletes-pagedattention-and-the-transformers-backend-catches-up.mdx
- vLLM + TileRT decode pairing (vLLM blog, 2026-07-14): https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2026-07-14-vllm-tilert-pd.md
- Community production log, v0.28.0 upgrade 2026-09-01 and v0.29.0 attempt 2026-09-09: https://github.com/jsboige/vllm/blob/HEAD/CLAUDE.md
- vLLM GGUF migration (PR #39612) and plugin status: https://github.com/nathanmaine/nvidia-dgx-spark-with-vllm/blob/HEAD/docs/ALTERNATIVES.md
- GGUF on vLLM + ROCm findings (2026-09-12/13, vLLM 0.27.0): https://github.com/badger-labs-dev/llamacpp-rocm-forge/blob/HEAD/docs/gguf-on-vllm-rocm-findings.md
- vLLM quantization deep-dive (plugin kernels, loading syntax, tested models): https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/vllm-quant-deep-dive.md
- unsloth/Qwen3-0.6B-GGUF on Hugging Face (repo exists, quants incl. Q4_K_M): https://huggingface.co/unsloth/Qwen3-0.6B-GGUF
- Qwen3-0.6B-GGUF via vLLM (xinference model spec): https://github.com/xorbitsai/inference/blob/HEAD/doc/source/models/builtin/llm/qwen3.rst
- vLLM-Omni paper (arXiv:2602.02204, 2026-02-02, 91.4% JCT claim): https://arxiv.org/pdf/2602.02204
- vLLM-Omni deployment guide (disaggregated encoder/prefill/decode, NIXL): https://www.spheron.network/blog/deploy-vllm-omni-disaggregated-multimodal-serving-gpu-cloud/
- AWS Deep Learning Containers vLLM-Omni docs (routes, bundled stack): https://github.com/aws/deep-learning-containers/blob/HEAD/docs/vllm-omni/index.md
- vLLM-Omni releases (v0.28.0, 2026-08; v0.12.0rc1, 2026-01): https://github.com/huyuwei/vllm-omni and https://github.com/liunix61/vllm-omni/blob/HEAD/README.md
- vLLM-Omni diffusion cookbook (v0.14–v0.20 version arc, Wan2.2 perf progression): https://github.com/hsliuustc0106/vllm-omni-cookbook/blob/HEAD/diffusion/wan2.2/index.md
- Inferact launch, $150M seed to commercialize vLLM (TechCrunch, 2026-01-22): https://techcrunch.com/2026/01/22/inference-startup-inferact-lands-150m-to-commercialize-vllm/
- Inferact launch coverage (Pulse2, 2026-01-22): https://pulse2.com/inferact-launches-with-150-million-funding-at-800m-valuation-to-commercialize-vllm-as-inference-demand-surges/
- Inferact launch coverage (The AI Insider, 2026-01-27): https://theaiinsider.tech/2026/01/27/inferact-launches-with-150m-seed-round-to-commercialize-vllm-inference-engine/
- Red Hat AI Inference Server (hardened vLLM + Neural Magic, announced May 2025): https://siliconangle.com/2025/05/20/red-hat-expands-ai-offerings-inference-server-validated-models/
- Red Hat AI platform / inference server coverage (The New Stack): https://Thenewstack.Io/red-hats-ai-platform-now-has-an-ai-inference-server/
- Futurum Group: vLLM as "de facto open-source LLM inference engine" at PyTorch Conference 2026 (2026-08-29): https://futurumgroup.com/insights/vllm-becomes-production-infrastructure-at-pytorch-conference-2026/
- PyTorch Conference 2026 session video (vLLM positioning): https://www.youtube.com/watch?v=3YUhsEuyZHY
- PyTorch Conference 2026 "Pluggable PyTorch LLM" session PDF: https://hosted-files.sched.co/pytorchconferenceeu2026/40/Pluggable%20PyTorch%20LLM%20-%20pytorch%2026.pdf
- NVIDIA NIM 2.0 vLLM-backend analysis (third-party deep-dive, reported 2026-08): https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/nvidia-deep-dive.md
- Gartner AI-optimized IaaS: inference $23.3B > training $19B, $42B total +96% (CXO Today, Aug 2026): https://cxotoday.com/cloud/gartner-global-ai-optimized-iaas-spending-to-surge-96-to-42b-in-2026/
- Gartner AI-optimized infrastructure spending, Oct 2025 $37.5B forecast (ChannelDive): https://www.channeldive.com/news/ai-optimized-infrastructure-spending-nvidia-gartner/802992/
- Training vs inference spend analysis, $757.7B hyperscaler capex (AInvest, Aug 2026): https://www.ainvest.com/news/training-built-nvidia-empire-inference-money-2608/
- "Training is dead, inference is the real AI war" (blog, 2026-03-20): https://medium.com/@Reiki32/training-is-dead-inference-is-the-real-ai-war-now-e8b679f8c703
- NVIDIA Rubin + Groq 3 inference processor analysis (AutoCoreAI, Sep 2026): https://www.autocoreai.net/blog/nvidia-rubin-why-ai-keeps-getting-cheaper-small-business-2026
- Stripe 73% cost reduction via vLLM (secondary, reported 2026): https://www.bbw9n.io/blog/sglang_vs_vllm
- ISG open-weight AI adoption (65% piloting, ~20% local LLM) via TradersUnion (Sep 2026): https://tradersunion.com/news/companies/show/3420099-ai-local-llm-trend/
- Disaggregation practitioner view, 4–10x cost-per-token (Shurankain agentic-ai-course, 2026): https://github.com/shurankain/agentic-ai-course/blob/HEAD/17_Production_Inference/07_SGLang_and_Alternatives.md
- NVIDIA Dynamo launch analysis, 7x throughput-per-GPU vendor claims (LinkedIn, 2026): https://www.linkedin.com/posts/paoloperrone_nvidia-just-open-sourced-the-inference-orchestration-activity-7463238312615645184-k3UV
- AWS EKS Dynamo blueprints for vLLM/SGLang/TensorRT-LLM disaggregated serving: https://github.com/awslabs/ai-on-eks/blob/HEAD/blueprints/inference/nvidia-dynamo/README.md
- Dynamo TRT-LLM backend guide, disaggregated gpt-oss-120b on B200: https://github.com/stars1233/dynamo/blob/HEAD/docs/backends/trtllm/trtllm-gpt-oss.md
- NIXL vs NCCL vs UCCL comparison (llm-systems-wiki, GPU-Communication): https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/GPU-Communication/15-nccl-vs-nixl-vs-uccl.md
- NIXL + DeepEP + connectors evidence table (llm-systems-wiki references): https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/GPU-Communication/21-references-and-research.md
- Distributed AI inference best practices 2026 (KV connector table, LMCache/NIXL, llm-d cache-aware routing): https://github.com/terrytangyuan/terrytangyuan.github.com/blob/HEAD/_posts/2026-06-05-distributed-ai-inference-best-practices-gotchas.md
- llm-d glossary v0.8 (NIXL, prefix caching, prefill/decode definitions): https://github.com/llm-d/llm-d.github.io/blob/HEAD/versioned_docs/version-0.8/api-reference/glossary.md
- AWS Deep Learning Containers vLLM changelog v1.1.0 (LMCache bidirectional NIXL probe, 2026-04-28): https://github.com/huggingface/deep-learning-containers/blob/HEAD/docs/vllm/changelog/index.md
- vLLM on Neuron (vllm-neuron plugin Beta + NxD Inference; Inf2/Trn1/Trn2): https://awsdocs-neuron.readthedocs-hosted.com/en/latest/libraries/nxd-inference/vllm/index.html
- AWS Neuron SDK, vLLM V1 API compatibility on Trainium/Inferentia: https://github.com/aws-neuron/aws-neuron-sdk/blob/HEAD/about-neuron/what-is-neuron.rst
- AWS Neuron SDK release notes, Neuron 2.32.0 (2026-08-17): https://github.com/aws-neuron/aws-neuron-sdk/blob/HEAD/index.rst
