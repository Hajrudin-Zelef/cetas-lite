---
id: ai-industry-kb-2026/10-training-fine-tuning/part-9
title: "10. Training & Fine-Tuning (part 9)"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Nvidia", "Samsung", "Unsloth", "Z.ai"]
dates: ["2026-02-09", "2026-02-10", "2026-03-23", "2026-04", "2026-05", "2026-05-06", "2026-07", "2026-07-22", "2026-08-30", "2026-09", "2026-09-02", "2026-09-09", "2026-09-15", "2026-09-17"]
keywords: ["fine-tuning", "training", "awq", "benchmark", "claude", "cost", "deepseek", "funding", "gguf", "glm", "gpu", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5708, 5755]
section: "10. Training & Fine-Tuning"
sha256: 5b2fe0c5987e5be4efe6414fe989573ed86635d7e537334c750c6e7767f621ac
---

# 10. Training & Fine-Tuning (part 9)

- [VENDOR] Unsloth × NVIDIA collaboration blog, 2026-05-06 (3 measured optimizations): https://unsloth.ai/blog/nvidia-collab
- [VENDOR] Unsloth benchmark table (unsloth-zoo README): https://github.com/unslothai/unsloth-zoo/blob/HEAD/README.md
- [VENDOR] Unsloth official changelog (latest entry 2026-09-17): https://unsloth.ai/docs/new/changelog
- [VENDOR] Unsloth releases v0.1.805-beta / v0.1.806-beta (MTP, ~early Sept 2026): https://github.com/unslothai/unsloth/releases/tag/v0.1.805-beta
- [VENDOR] https://github.com/unslothai/unsloth/releases/tag/v0.1.806-beta
- [VENDOR] Unsloth releases v0.1.807-beta / v0.1.808-beta (perf gains, ~early-mid Sept 2026): https://github.com/unslothai/unsloth/releases/tag/v0.1.807-beta
- [VENDOR] https://github.com/unslothai/unsloth/releases/tag/v0.1.808-beta
- [VENDOR] Unsloth release v0.1.481-beta ("DeepSeek-V4 + NVFP4 Exporting", July 2026): https://github.com/unslothai/unsloth/releases/tag/v0.1.481-beta
- [VENDOR] Unsloth versioning PR (2026.9.2, 2026-09-02): https://github.com/unslothai/unsloth/pull/10199
- [VENDOR] Studio release-body fixture v0.1.501-beta (MoE expert offload, GGUF export fixes, MCP): https://github.com/unslothai/unsloth/blob/HEAD/tests/studio/fixtures/release_bodies/v0.1.501-beta.md
- [VENDOR] Studio release-body fixture v0.1.471-beta (GLM-5.2, MTP auto-fit): https://github.com/unslothai/unsloth/blob/HEAD/tests/studio/fixtures/release_bodies/v0.1.471-beta.md
- [VENDOR] Unsloth Qwen3.8 fine-tuning guide (September 2026): https://unsloth.ai/docs/models/qwen3.8/train
- [VENDOR] Unsloth HF organization (1,374 models, GGUF/NVFP4 uploads): https://huggingface.co/unsloth
- [VENDOR] Unsloth Dynamic v3.0 docs: https://unsloth.ai/docs/basics/dynamic-3.0-ggufs
- [VENDOR] Unsloth Dynamic 2.0 docs (superseded lineage): https://unsloth.ai/docs/basics/unsloth-dynamic-2.0-ggufs
- Unsloth MoE Triton grouped-GEMM correctness fix (2026-09-09): https://github.com/danielhanchen/unsloth-zoo-staging/commit/57aa8fa32a3fceac11b7d1795d3b828f7342c6ba
- HF TRL Unsloth integration docs (save_pretrained_gguf API, 23-entry quant list, OOM guard; updated ~2026-09-17): https://github.com/huggingface/trl/blob/HEAD/docs/source/unsloth_integration.md
- Hugging Face Unsloth–TRL benchmark blog (59 runs, T4/A100): https://huggingface.co/blog/unsloth-trl
- torchtune paper, May 2026 (independent torchtune/Axolotl/Unsloth comparison): https://arxiv.org/pdf/2605.21442v1.pdf
- MarkTechPost framework comparison, 2026-07-22 (Unsloth vs Axolotl vs TRL vs Llama-Factory): https://www.marktechpost.com/2026/07/22/unsloth-vs-axolotl-vs-trl-vs-llama-factory-a-fine-tuning-framework-comparison-on-speed-vram-and-multi-gpu/
- Independent cross-entropy loss-head analysis (Chew Loong Nian, July 2026): https://medium.com/@chewloongnian/unsloth-vs-axolotl-vs-trl-87-of-your-fine-tuning-vram-goes-to-a-tensor-you-never-wrote-d21b8326d89d
- Independent practitioner comparison, 2026-02-09 (incl. loss): https://medium.com/@balci.pelin/unsloth-vs-standard-training-92d4c35b8ad8
- "Run AI Locally in 2026" (April 2026 adoption snapshot, 61K stars, RTX 4090 example): https://medium.com/@computeleap/run-ai-locally-in-2026-dgx-spark-unsloth-beyond-67c2b3627dba
- [COMMUNITY] Unsloth performance-reality notes (community skill doc, July 2026): https://github.com/ericrisco/rsc-harness/blob/HEAD/skills/unsloth/SKILL.md
- [COMMUNITY] Unsloth implementation notes (community skill doc): https://github.com/maragudk/fabrik/blob/HEAD/skills/unsloth/SKILL.md
- [COMMUNITY] Unsloth GGUF deployment skill doc: https://github.com/scientiacapital/unsloth-mcp-server/blob/HEAD/.claude/skills/model-deployment/SKILL.md
- [COMMUNITY] Third-party unsloth-cli 0.7.0/0.7.1 (2026-09-15): https://github.com/agentculture/unsloth-cli/blob/HEAD/CHANGELOG.md
- [COMMUNITY] unsloth-cli fine-tuning docs (NVFP4/AWQ serving-tested, Sept 2026): https://github.com/agentculture/unsloth-cli/blob/HEAD/docs/fine-tuning.md
- [COMMUNITY] unsloth-cli /finetune skill: https://github.com/agentculture/unsloth-cli/blob/HEAD/.claude/skills/finetune/SKILL.md
- [COMMUNITY] Community DGX Spark speedup project (May 2026) — UNVERIFIED figures: https://github.com/megastood/dgx_spark_unsloth_lossless_speedup
- [COMMUNITY] Unsloth performance analysis (RTX 4090 vs torchtune): https://github.com/akaszubski/realign/blob/HEAD/docs/research/UNSLOTH_ANALYSIS.md
- [COMMUNITY] Unsloth ~$500K seed figure (Redpoint Ventures scout, Samsung NEXT) — agentvc-index, 2026-03-23 (third-party analyst, not a funding announcement): https://github.com/lucy-cxy/agentvc-index/blob/main/cases/2026-03-23_unsloth.md
- [UNVERIFIED] GLM-5.2 Dynamic GGUF figures (secondary/LinkedIn): https://www.linkedin.com/posts/linasbeliunas_huge-unsloth-just-shrunk-the-strongest-open-activity-7474842446586179584-KoHK
- [UNVERIFIED] GLM-5.2 local self-hosting cost math (secondary): https://www.thundercompute.com/blog/glm-5-2-unsloth
- Axolotl README "Latest Updates" (2026 cadence): https://github.com/stars1233/axolotl/blob/HEAD/README.md
- DeepSpeed MoE tutorial: https://github.com/deepspeedai/deepspeed/blob/HEAD/docs/_tutorials/mixture-of-experts.md
- DeepSpeed AutoEP blog, 2026-02-10: https://deepspeed.ai/blog/2026/02/10/AutoEP/
- PyTorch release cadence (2.11–2.16): https://github.com/pytorch/pytorch/blob/main/RELEASE.md
- ESFT paper (DeepSeek AI + Northwestern, arXiv 2407.01906): https://export.arxiv.org/pdf/2407.01906v1.pdf
- ESFT official code: https://github.com/shahils01/ESFT
- DoRA paper: https://ar5iv.labs.arxiv.org/html/2402.09353
- LoRA-FA / DoRA comparison (ICLR 2026 submission): https://arxiv.org/abs/2511.04021
- Liger Kernel paper: https://arxiv.org/pdf/2410.10989
- Liger Kernel TRL integration docs: https://github.com/huggingface/trl/blob/HEAD/docs/source/liger_kernel_integration.md
- GitHub AI radar report (Unsloth star trajectory, 2026-08-30): https://github.com/juliayu907/github-ai-radar/blob/HEAD/reports/2026-08-30/github_ai_hot_repo_2026-08-30_en.md
- Soup layer-streaming fine-tuning on 4 GB laptop GPU: https://hackernoon.com/how-two-engineers-built-a-6-gb-tool-that-outperforms-32-gb-of-vram
- MoE Parallel Folding (fine-grained vs coarse-grained training efficiency, arXiv 2504.14960v2): https://arxiv.org/pdf/2504.14960v2.pdf

