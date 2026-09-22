---
id: etape4-trackd-unsloth-training/00-unsloth-training/2-axolotl
title: "2. AXOLOTL"
domain: step-4-track-d-unsloth-training-fine-tuning-tooling-2026
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "China", "Cohere", "DeepSeek", "Falcon", "Huawei", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-06-17", "2026-07-17", "2026-08", "2026-09"]
keywords: ["agentic", "ascend", "attention", "bitnet", "cohere", "deepseek", "distillation", "dpo", "fine-tuning", "flash attention", "funding", "glm"]
source: docs/RAG/etape4_trackD_unsloth_training.md
source_anchor: ""
source_lines: [66, 119]
section: "Step 4 — Track D: Unsloth + Training / Fine-Tuning Tooling (2026)"
sha256: 8940fa5533a27e5d544bae4883aba01b0859061c2670406a982f4f36ec33385a
---

# 2. AXOLOTL

## 2. AXOLOTL

### 2.1 v0.19.0 release — September 2026 (latest)
- Release: **axolotl-ai-cloud/axolotl v0.19.0**, ~12 days before Sep 22, 2026; **72 commits since v0.18.0 (July 17, 2026)** [official — https://github.com/axolotl-ai-cloud/axolotl/releases/tag/v0.19.0].
- Highlights:
  - **Declarative Model Support Profiles**: new-arch support via family templates (`VANILLA_CAUSAL_LM`, `IMAGE_TEXT_TO_TEXT`) instead of scattered loader edits; no config changes [official].
  - **Qwen3.8-Flash-Next**: fine-tune Qwen's 176.94B multimodal MoE on a single B300, as little as 120 GiB with `ple_cpu_offload: true`; QLoRA, vision QLoRA, NVFP4 MoE-LoRA configs ship [official].
  - **Ternary (BitNet b1.58) QAT**: `qat.weight_dtype: ternary` — weights restricted to {-1,0,1} natively, no extra dependency; Neutrino-1 two-stage recipe (ternary continued pretraining → instruction tuning) [official].
  - Seven new model families; two new optimizers; chat-template tokenization speedup; **PyTorch 2.13 support** [official].
- Community summary (README, Sep 2026): Unsloth 34k+ vs **Axolotl 17k+ GitHub stars**; positioning: Axolotl = production-grade distributed training [secondary].

### 2.2 2026 feature timeline (from README "Latest Updates")
- **2026/07**: NVFP4 (4-bit) MoE LoRA via **ScatterMoE (W4A16)** and **SonicMoE (W4A4)**, incl. lossless adapter merge back into NVFP4 checkpoint [official/secondary].
- **2026/06**: **Expert Parallelism (EP)** for distributed MoE via **DeepEP**; remote training through **Tinker-compatible APIs** (Thinking Machines' Tinker); **Context Parallelism for hybrid SSM models** (Nemotron-H, Falcon-H1, Bamba); **BitNet 1.58-bit fine-tuning**; multimodal assistant-only loss-masking fix [secondary].
- **2026/04**: model support for **Mistral Medium 3.5** and **Gemma 4**; **Async GRPO** (up to 58% faster steps); **Flash Attention 4**; **NeMo Gym** integration; **EBFT**; **uv-first** packaging; **SonicMoE fused LoRA** [secondary].
- **2026/03**: Mistral Small 4, Qwen3.5 (+MoE), GLM-4.7-Flash, GLM-4.6V, GLM-4.5-Air; **MoE expert quantization** (`quantize_moe_experts: true`, FSDP2-compatible) [secondary].
- **2026/02**: **ScatterMoE LoRA** (LoRA directly on MoE expert weights, custom Triton kernels); **SageAttention**; **GDPO** (Generalized DPO) [secondary].
- **2026/01**: **EAFT** (Entropy-Aware Focal Training); **Scalable Softmax** for long-context attention [secondary].
- **2025/12** (context): Kimi-Linear, Plano-Orchestrator, MiMo, InternVL 3.5, Olmo3, Trinity, Ministral3; **Distributed Muon Optimizer** for FSDP2 pretraining [secondary].
- Note: Axolotl positions against Unsloth as YAML-config-driven, production-grade, **FSDP2 multi-node** training [secondary].

### 2.3 Company status
- Maintained by **axolotl-ai-cloud** org (the company behind Axolotl AI Cloud). No funding announcement found in 2026 research — **[unverified/no data]**. Treat "Axolotl raised $X" as unconfirmed.

---

## 3. LLAMA-FACTORY (hiyouga/LLaMA-Factory)

- **v0.9.5 — June 17, 2026** (latest found): support for **Qwen 3.5/3.6**, Gemma 4, Phi-4-mini, **GLM-4.7-Flash**; Tencent Hunyuan/Youtu; **Huawei Ascend NPU** (Partial RoPE, Hybrid Attention); ROCm 7.2; VLMs (MiniCPM-V-4.6, Qwen3-VL, LiquidAI LFM 2.5-VL); **Transformers v5 compatibility**; **FSDP2**; Liger Kernels for Qwen 3.5/Qwen3-Next [secondary — AI China News].
- v0.9.3 (Mar 2026): Llama 4, DeepSeek-R1, Qwen3, InternVL3, Gemma 3 [secondary].
- v0.9.2 (early 2026): DeepSeek V3 (671B), Qwen2.5-VL, InternLM3; QLoRA on Ascend via CANN; APOLLO optimizer; Ray Trainer; vLLM batch inference; Ollama modelfile auto-generation [secondary].
- Role: "Swiss Army Knife" of the Chinese developer community; no-code **LlamaBoard** web UI; 30K+ GitHub stars (2025) [secondary].
- ⚠️ All LLaMA-Factory version details sourced from AI China News auto-generated posts; treat specific feature claims as [secondary].

---

## 4. TORCHTUNE (PyTorch/Meta)

- PyPI description (crawled Sep 2026; package page itself dated ~Dec 2025, v0.6.x era): hackable recipes for **SFT, knowledge distillation, DPO, PPO, GRPO, QAT**; native Llama/Gemma/Mistral/Phi/Qwen implementations; YAML configs [official — pypi.org/project/torchtune].
- Recipe coverage matrix (from docs): full/single/multi-node SFT; LoRA/QLoRA single-device & multi-device (not multi-node); DPO full (distributed) + LoRA; PPO full single-device; **GRPO marked 🚧 (in progress)**; QAT distributed [official].
- **No 2026 torchtune release notes found** in this research window — flag as **[gap]**: Meta's torchtune releases (0.7+) were not located; do not assume features. Known context: torchtune 0.6 (Dec 2025) added Qwen3, Gemma 3, Llama 4 Scout support and torchao QAT improvements — outside or at the edge of window [background knowledge, unverified in-window].

---

## 5. TRL (Hugging Face)

- **TRL v1.0** landed early 2026: unified post-training stack (SFT, reward modeling, DPO, GRPO) with one API [secondary — press]. Docs/community notes put the line at ~**v1.8** by September 2026 [secondary].
- **TRL v1.4** (announced by HF staff @qgallouedec, ~Sep 15, 2026): chunked NLL loss for SFT; first-class **agentic RL environments** (`trl.experimental.openreward`, `OpenRewardSpec("Eigent/SETA", num_tasks=64)` — one-string wiring of dataset/factory/reward; dynamic tool binding from JSON Schema); MFU helpers for dense + MoE; **GRPO support for Liger 0.8.0** (delta clipping + VESPO + KL bias correction); Tülu 3 length-normalized DPO loss; 4 new chat templates (Cohere, Cohere2, Gemma 3, Qwen3-2507); 5+ GB CUDA memory-leak fix in activation offloading [secondary — HF post].
- **As of August 2026**: `GRPOTrainer` `loss_type` defaults to **"dapo"**, not original GRPO (docstring: "Not recommended due to length bias"); GSPO via `importance_sampling_level="sequence"`; paper index covers DAPO, Dr. GRPO, GSPO, CISPO reproductions [independent — HF blog "How frontier models train on outcomes in 2026", Aug 2026].
- CLI: `trl sft|dpo|grpo|kto|rloo|reward|distillation`; `trl vllm-serve` **deprecated** in favor of `vllm serve` [official — trl docs].
- Trend note (HF blog, Aug 2026): frontier labs train **in environments** (sandboxed tool-use rollouts, outcome-based rewards combining LLM-judge rubrics + programmatic checks + safety gates); cited: LiquidAI LFM2.5-2.6B, Cursor Composer, Kimi K3, GLM-5, Nemotron 3 Ultra, MiniMax [independent].

---

