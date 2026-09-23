---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/3-the-hugging-face-ecosystem
title: "3. The Hugging Face ecosystem"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: platform
actors: ["Alibaba", "Hugging Face"]
dates: ["2025-12", "2025-12-18", "2026-01", "2026-01-26", "2026-05", "2026-09-18"]
keywords: ["diffusion", "dpo", "fine-tuning", "gpus", "lora", "qlora", "qwen", "reasoning", "training"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [187, 266]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 37bb429dc3ee392bf65843d2d5ac32929c07cb692620ee4a139318dfc24f829e
---

# 3. The Hugging Face ecosystem

## 3. The Hugging Face ecosystem

### 3.1 Hugging Face Hub — scale and the numbers conflict

Third-party sources disagree on Hub totals; all figures below are secondary
and conflict with each other, so none is presented as authoritative:

- Source A: 2.4M+ models, 730K+ datasets, ~1M Spaces (as of May 2026)
  [secondary].
- Source B: 1.8M models, 450K datasets, 720K Spaces [secondary].
- Source C: 2M+ models, 500K+ datasets, ~1M Spaces [secondary].
- A September 18, 2026 API-based study reported Qwen occupying 27 of the top
  60 downloaded open text/VLM repositories, with Qwen3-0.6B at 22.5M rolling
  30-day downloads; this is download-count methodology, not unique users
  [independent for the study, secondary for Hub totals].
- Recorded as conflict C-01: verify against a dated Hub API snapshot before
  quoting any total.

### 3.2 Transformers v5 — the 2026 major release

Transformers v5 is the first major version in roughly five years and is the
defining library event of this phase's window:

- Community migration guides date the v5 landing to late December 2025 /
  January 2026 (one guide says v5.0.0 landed January 26, 2026; another
  describes a December 18, 2025 release) — recorded as conflict C-02
  [secondary].
- **PyTorch is the sole backend**: all `TF*` and `Flax*` classes were
  removed; importing them raises `ImportError` [secondary].
- Minimum requirements per migration guides: Python 3.10+, PyTorch 2.4+,
  `huggingface_hub` 1.x [secondary].
- `huggingface_hub` 1.x switched the HTTP backend from `requests` to
  `httpx`, dropped `hf_transfer` in favor of `hf_xet`, and removed legacy
  cache env vars (`TRANSFORMERS_CACHE`, etc.) in favor of `HF_HOME`
  [secondary].
- The tokenizer redesign removes the Fast/Slow distinction and makes
  tokenizers composable, trainable-from-scratch templates [secondary].
- Hugging Face's own framing (via secondary coverage) positions v5 around
  interoperability, simpler model building, and better performance on modern
  GPUs [vendor-reported via secondary].
- Vendor-scale claims via secondary coverage: ~3M daily installs, 1.2B+
  cumulative installs, 400+ model architectures supported, 750K+
  community models at the v4→v5 transition — all [vendor-reported], not
  independently confirmed.
- Security note: CVE-2026-1839 (arbitrary code execution via
  `Trainer._load_rng_state()` calling `torch.load()` without
  `weights_only=True`) affected versions supporting `torch>=2.2` on PyTorch
  <2.6 and was fixed in v5.0.0rc3 [secondary].

### 3.3 Diffusers, Datasets, TRL, PEFT, Accelerate

- **Diffusers** is the Hugging Face library for diffusion models
  (image/video/audio generation: Stable Diffusion, Flux, etc.); no 2026
  version was captured in this pass — gap G-03 [unverified].
- **Datasets** is the Arrow-backed dataset library and Hub dataset loader
  underlying most open-data pipelines; the Hub dataset counts above
  (450K–730K, conflicting) include these [secondary].
- **TRL** (Transformer Reinforcement Learning) implements post-training
  methods (SFT, DPO, PPO/GRPO-style RL) on top of Transformers; central to
  the 2026 reasoning-model post-training wave [secondary].
- **PEFT** (Parameter-Efficient Fine-Tuning) implements LoRA, QLoRA
  (with bitsandbytes), AdaLoRA, and related adapters; LoRA remains the
  default fine-tuning interface for open models in 2026 [secondary].
- **Accelerate** abstracts distributed training and mixed precision across
  DDP/FSDP/DeepSpeed backends for the Transformers `Trainer` [secondary].
- **huggingface_hub** 1.x is the client library for the Hub (models,
  datasets, Spaces), now `httpx`-based per Section 3.2 [secondary].

### 3.4 Library quick-reference matrix

| Library | Role | 2026 status |
|---|---|---|
| `transformers` v5 | Model APIs, tokenizers, Trainer | PyTorch-only, tokenizer redesign [secondary] |
| `diffusers` | Diffusion models | Version not captured (G-03) [unverified] |
| `datasets` | Dataset loading/sharing | Arrow-backed, Hub-integrated [secondary] |
| `trl` | Post-training (SFT/DPO/RL) | Active, reasoning-era relevance [secondary] |
| `peft` | LoRA/QLoRA adapters | Default fine-tune interface [secondary] |
| `accelerate` | Distributed/mixed-precision glue | DDP/FSDP/DeepSpeed abstraction [secondary] |
| `huggingface_hub` 1.x | Hub client | `httpx`, `hf_xet`, `HF_HOME` [secondary] |

