---
id: ai-industry-kb-2026/10-training-fine-tuning/main-actors
title: "Main actors"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02", "2026-02-09", "2026-02-10", "2026-03-18", "2026-03-23", "2026-04", "2026-05", "2026-05-06", "2026-05-13", "2026-06", "2026-07", "2026-07-08", "2026-07-22", "2026-08-12", "2026-08-19", "2026-08-30", "2026-09", "2026-09-02", "2026-09-09", "2026-09-15", "2026-09-17", "2026-09-22", "2026-10-28"]
keywords: ["amd", "apache", "attention", "awq", "benchmark", "benchmarks", "bitnet", "claude", "consumer", "cost", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5602, 5755]
section: "10. Training & Fine-Tuning"
sha256: feeb4650c25931ffb5c13bda186da38433a164395d86c8bd33314d5c2c1d8e7a
---

# Main actors

## Main actors

### Unsloth (the vendor)

- **Daniel Han and Michael Han** (brothers, Australia-based; Daniel has an NVIDIA background) — founders. Documented official fine-tuning partnerships with Google, OpenAI, Meta, and NVIDIA (partners, not investors).
- **Three products by September 2026**: Unsloth Desktop (Tauri app), Unsloth Studio (web UI), Unsloth Core (code library). Model hub ships per-quant downloads with thinking toggle and `reasoning_effort` controls.
- **Dual license**: Apache 2.0 for core, AGPL-3.0 for Studio. Commercial tiers Pro/Enterprise add multi-GPU/multi-node and full-parameter training — reflecting the boundary between Unsloth's core strength (1–2 GPU PEFT) and the multi-node world.
- **RL integration**: plugs into TRL's `GRPOTrainer`/`GRPOConfig` with a built-in vLLM engine (`fast_inference=True`) for rollouts; the 2026-09-17 changelog confirms GRPO with Qwen3.5 plus latest TRL/vLLM. 2026 RL items: 7x longer-context RL via new batching; FP8 & Vision (VLM) GRPO on consumer GPUs; DoRA support; embedding fine-tuning 1.8–3.3x faster. Docs cite ~80% less VRAM for GRPO vs standard setup (flagged "verify" even by community skill docs — treat as vendor claim).
- **Model catalog (Sept 2026)**: Qwen3.8, Qwen3.6 (MTP 1.4–2.2x faster inference, NVFP4), Kimi K3, Kimi K2.7 Code, MiniMax M3, MiniMax-H3, GLM-5.3-Flash, GLM-5.2 (Dynamic GGUF), DeepSeek-V4-Flash, Gemma 4, DiffusionGemma, Qwen3.8-Flash-Next, Muse Glimmer (Meta), Ornith, FLUX, Wan. Catalog scope caveat: verified current Unsloth materials confirm the named families plus NVFP4/GGUF availability in *parts* of the catalog — availability is family/model/hardware-specific, and "Nemotron" was not confirmed in searched official materials. Do not present the catalog as uniform.

### Framework competitors

- **Axolotl** — the YAML-driven multi-GPU fine-tuning framework. 2026 MoE cadence (from its README "Latest Updates"): **February** — ScatterMoE LoRA (LoRA directly on MoE expert weights via custom Triton kernels), SageAttention, GDPO; **March** — `quantize_moe_experts` (greatly reduces VRAM, FSDP2-compatible), new-model support (Mistral Small 4, Qwen3.5/MoE, GLM-4.7-Flash); **April** — Async GRPO (up to 58% faster steps), Flash Attention 4, SonicMoE fused LoRA, EBFT, uv-first packaging; **June** — Expert Parallelism via DeepEP for distributed MoE training, Tinker-compatible APIs for remote training, context parallelism for hybrid SSM models, BitNet 1.58-bit fine-tuning; **July** — NVFP4 (4-bit) MoE LoRA via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint. Axolotl owns the multi-GPU research lane alongside torchtune.
- **torchtune** (Meta/PyTorch ecosystem; May 2026 paper, arXiv 2605.21442v1) — positions itself between high-level automated trainers and kernel-specialized systems: modular YAML recipes, in-backward optimizer fusion (reducing gradient-buffer lifetime), composable DTensor parallelism (FSDP2 / tensor / sequence / loss / context parallelism plus a custom expert-parallel plan for MoE), linear cross-entropy, and an **asynchronous GRPO recipe** (Ray-coordinated queue + replay buffer, vLLM rollout workers, on-policy and bounded-off-policy modes — a system design, with head-to-head comparisons left to future work). Its stated conclusion: competitive or superior efficiency vs Axolotl and Unsloth while remaining transparent and hackable.
- **DeepSpeed** (Microsoft) — the max-scale path. February 2026 blog introduced **AutoEP** (automatic expert-parallelism configuration); the MoE tutorial supports expert parallelism combined with data parallelism, ZeRO, and model parallelism (docs refreshed ~September 2026); communication-optimization notes identify MoE all-reduce penalties of **3–12 seconds** in problematic large-scale layouts, with multi-rank bucketing as the fix.
- **Liger Kernel** — open-source Triton kernels integrated into TRL, FSDP2, and DeepSpeed; published end-to-end benchmarks on 4× A100 80 GB (Alpaca, BF16, seq 512): Llama 3 8B +42.8% throughput / −54.8% memory; Qwen2 +25.5% / −56.8%; Gemma +11.9% / −51.8%; Mistral +27% / −21%; Phi-3 +17% / −13%. Trainer support for SFT, DPO, GRPO, KTO, GKD. **Complementary** to Unsloth (open, integrated into the HF stack) rather than a drop-in substitute.

### PEFT method researchers (adapter-quality axis)

- **ESFT — Expert-Specialized Fine-Tuning** (DeepSeek AI + Northwestern, arXiv 2407.01906, EMNLP 2024). The MoE-native PEFT method: score expert affinity on task data, fine-tune only the most relevant experts per layer, freeze the rest. Reported: storage ↓ up to 90%, training time ↓ up to 30% vs full-parameter fine-tuning; matching or exceeding full-FT on math/code while retaining general-task performance better than LoRA. Key architectural finding for this topic: **finer-grained experts are more advantageous for ESFT**. Official code: github.com/shahils01/ESFT.
- **DoRA** (arXiv 2402.09353): decomposes weights into magnitude + direction, applies low-rank adaptation to directional components; mergeable at inference with no serving overhead. Reported: LLaMA-7B 78.1 vs LoRA 74.7; LLaMA-13B 81.5 vs 80.5 (average scores); partial-DoRA can beat LoRA with under half the trainable parameters in reported tests.
- **LoRA-FA** (ICLR 2026 submission; arXiv 2511.04021): at r=64, scored 5.7/28.1/57.0 (MT-Bench/HumanEval/GSM8K) with gains over LoRA of +0.5/+19.1/+14.4; DoRA scored 5.9/19.0/52.2 at the same rank.
- **RoRA** reportedly beats DoRA; **Dual LoRA** (ICLR 2026) reportedly outperforms DoRA with fewer parameters; **PiSSA** is best framed as an initialization/quality method, not a memory-system breakthrough.
- **PEFT-Arena** (May 2026, arXiv): stability–plasticity comparison across LoRA, AdaLoRA, DoRA, VeRA, PiSSA, MiLoRA, OFT, IA3, and full fine-tuning. OFT variants may offer the better retention/adaptation Pareto — do not rank methods by downstream accuracy alone.

### Community and third-party ecosystem

- **Chew Loong Nian** (July 2026) — independent practitioner whose first-principles loss-head analysis confirmed the cut-cross-entropy mechanism story.
- **Pelin Balci** (2026-02-09) — early independent comparison including loss parity.
- **agentculture/unsloth-cli** (0.7.0/0.7.1, 2026-09-15) — container-backed multi-format export, quantization-loss eval (`sloth eval`).
- **megastood DGX Spark project** (~May 2026) — [COMMUNITY, UNVERIFIED] speedups for Qwen3.5 on GB10 vs stock Unsloth; self-hosting breakeven analysis.
- **Community skill-documentation authors** (e.g. rsc-harness, fabrik, unsloth-mcp-server, local-model-autotuning) — the ecosystem's de-facto hedge culture ("do not quote a single number as gospel"; treat 80%-GRPO as vendor claim).
- **MarkTechPost** (2026-07-22) — the independent 4-framework comparison.

## Timeline and context

### Full trajectory, February → September 2026

| Date | Event |
|---|---|
| 2026-02-09 | Pelin Balci publishes independent "Unsloth vs Standard Training" comparison (speed, memory, loss) |
| 2026-02-10 | DeepSpeed AutoEP blog (automatic expert-parallelism configuration) |
| 2026-02 | Unsloth February release: MoE training pipeline, "12x faster MoE training" headline [VENDOR] |
| 2026-03-18 | PyTorch 2.11 released |
| 2026-04 | Unsloth ~61K GitHub stars; "default tool for local fine-tuning"; Studio (data recipes from PDFs/CSVs/DOCXs); FP8 training; 500K-context training on consumer hardware; Llama 3.1 8B QLoRA < 2 h on RTX 4090 |
| 2026-05-06 | Unsloth×NVIDIA collaboration blog: 3 measured optimizations (packed metadata +43.3% fwd / double-buffered checkpointing +4.6–8.4% / bincount MoE routing +23% fwd +13% bwd) |
| 2026-05 | torchtune paper (arXiv 2605.21442v1): Unsloth = memory leader at every tested size; torchtune+torch.compile = throughput leader at 3 of 4 sizes |
| 2026-05 | PEFT-Arena: OFT variants may offer the better retention/adaptation Pareto |
| 2026-05-13 | PyTorch 2.12 released |
| ~2026-05 | Community DGX Spark project claims 7.67x LoRA / 8.35x full-FT for Qwen3.5 on GB10 vs stock Unsloth [COMMUNITY, UNVERIFIED] |
| 2026-07-08 | PyTorch 2.13 released |
| 2026-07 | Studio v0.1.481-beta: NVFP4/FP8/imatrix GGUF export; DeepSeek-V4-Flash; "GRPO 1.3x", "MoE training 3–5x" [VENDOR] |
| 2026-07 | Chew Loong Nian's loss-head analysis: 87.3% of sequence-scaling memory in Llama 3.1 8B LoRA is the cross-entropy head; 95.2% for gpt-oss-20b; 49.1 GB at 16K context |
| 2026-07 | Axolotl NVFP4 MoE LoRA (ScatterMoE W4A16, SonicMoE W4A4); Unsloth NVFP4 export — parallel tracks |
| 2026-07-22 | MarkTechPost 4-framework comparison: gpt-oss-20b in 12.8 GB; Qwen3-30B-A3B 16-bit LoRA 63 GB; B200 8K 47.43 vs 73.80 GB, 16K 55.13 vs OOM |
| 2026-08-12 → 2026-08-30 | Stars 70.4K → 75.2K (~400–570/day) |
| 2026-08 | "Soup"-style layer-streaming approach fine-tunes 8B models on a 4 GB laptop GPU |
| 2026-08-19 | **Dynamic v3.0 launched — supersedes Dynamic Quantization 2.0** (format details: §8) |
| 2026-09-02 | PyTorch 2.14 released (silent clamp/min/max boundary-subgradient change — training-reproducibility hazard); `unsloth/_version.py` bumped to 2026.9.2 |
| ~2026-09-early | Studio v0.1.805/806-beta: MTP-by-default for Qwen3.8-Flash-Next + GLM-5.3-Flash ("up to 2x faster generation" [VENDOR]); 170+ improvements; MLX MoE training |
| 2026-09-09 | MoE Triton grouped-GEMM correctness fix (down-LoRA row order; scatter via `gather_indices`) + dtype-aware backend selection |
| ~2026-09-mid | Studio v0.1.807/808-beta: AMD Vulkan by default; INT8/FP8 diffusion 1.2–1.7x (downward revision of the Sept-17 "2x"); gated-delta training +25% on Apple Silicon [VENDOR]; DoRA/DPO on MLX |
| 2026-09-15 | Community `unsloth-cli` 0.7.0/0.7.1 (container export, quantization-loss eval) [COMMUNITY] |
| 2026-09-17 | Official changelog "Docker + MultiUser + AMD Support" (latest entry as of 2026-09-22) |
| 2026-09-17 | TRL Unsloth-integration docs (GGUF one-call API, 23-entry quant list, OOM guard) verified current |
| September 2026 | Qwen3.8-27B fine-tuning guide: ~1.5x/50% per-model figures; QLoRA 24 GB / LoRA >36 GB / FFT 4x VRAM |
| 2026-09-22 | No releases newer than the 2026-09-17 changelog found; HF `unsloth/` org (1,374 models) shows uploads within days |

### Distributed training in 2026: FSDP2 as the default

- **FSDP2 is the 2026 default distributed substrate** for multi-GPU fine-tuning (torchtune's DTensor stack; Axolotl's FSDP2-compatible `quantize_moe_experts`; Liger's FSDP2 integration).
- **Axolotl (June 2026)** added **Expert Parallelism (EP) for distributed MoE training via DeepEP**, documented under its N-dimensional parallelism docs — moving DeepEP from the pretraining world into the fine-tuning YAML-config world.
- **torchtune (May 2026 paper)** builds its parallelism stack on PyTorch DTensor: FSDP2, tensor and sequence parallel, **a custom expert-parallel plan for MoE models**, **loss parallel** (sharding output features over the vocabulary dimension across the TP mesh so full logits are never materialized), context parallelism via Ring Attention. The same recipes scale from a single H100 to multi-node FSDP2 clusters without rewriting the training loop.
- Together these confirm that **expert parallelism is becoming a configurable primitive in fine-tuning frameworks**, not just a pretraining-systems concern.

### The framework-selection rule (falls out of the 2026 evidence)

- **Single/dual-GPU PEFT and local workflows → Unsloth** (lowest memory, simplest).
- **YAML-driven multi-GPU with FSDP2/EP and broad model coverage → Axolotl**.
- **Hackable native-PyTorch research with custom parallelism → torchtune**.
- **Trillion-parameter or heavily customized scale → DeepSpeed/Megatron-style**.

### Benchmark methodology: reading Unsloth numbers correctly

1. **The baseline dominates the headline.** Unsloth's figures are vs Hugging Face + FlashAttention 2 QLoRA. Against that baseline: 1.87–2.74x speedups and 12–74% VRAM savings (HF's own blog, 59 runs). Against torchtune with torch.compile (May 2026 paper), Unsloth wins on memory but loses on throughput at most sizes. Always record the baseline.
2. **Report the config, not just the model.** The >75% VRAM figure is Llama 3.3 70B, Alpaca, batch 2, grad-accum 4, rank 32, QLoRA on all linear layers, 80 GB GPU. Change any of these and the figure moves.
3. **Separate the three VRAM figures.** ~70% = general marketing claim (still current); >75% = Llama 3.3 70B QLoRA vendor benchmark; ~80% = GRPO-specific (Qwen3 4B row). The 80%-as-universal reading is unverified.
4. **Separate MoE-specific from general figures.** The 12x MoE headline is hardware-specific (B200 in Unsloth's own post) and vendor-reported; the component-level measurements (NVIDIA blog: +43.3% forward packed metadata, +8.4%/+6.7%/+4.6% checkpoint reload, +23%/+13% bincount routing) are the auditable part.
5. **Vendor vs independent.** Vendor: unsloth-zoo tables, Unsloth docs/blog. Independent: HF's Unsloth–TRL blog, torchtune paper, MarkTechPost comparison, Balci's loss-aware comparison, Chew's loss-head analysis. **No independent reproduction of the 12x MoE multiplier or the 12.8 GB gpt-oss-20b figure was found.**
6. **Loss parity is claimed, and partly checked.** Unsloth claims 0% accuracy degradation (exact optimization, no approximations). The February 2026 Balci comparison is the independent practitioner check that included loss; the HF blog predates the MoE work but also reported 0% degradation for standard QLoRA.

### Fine-grained MoE and bandwidth — cross-ref to §9

The MoE architecture question (whether fine-grained experts hurt inter-GPU bandwidth) is resolved in §9: fine granularity improves routing specialization and combinatorial flexibility, but MoE Parallel Folding (arXiv 2504.14960v2) shows fine-grained MoE has *lower* training efficiency than coarse-grained MoE across tested parallelism strategies — bandwidth is preserved only by co-designed mitigations (DeepEP-class dispatch, latent-space routing, Quantile Balancing). The claim that fine granularity improves routing "without hurting inter-GPU bandwidth" is **not verified and contradicted as stated**; this section does not duplicate that analysis.

## Implications

- **For RAG retrieval design:** architecture claims (expert counts, routing formulas) and systems claims (bandwidth, VRAM, speedups) must live as separate facts with separate evidence grades — the 2026 literature shows they move in opposite directions for fine-grained MoE. Speedup claims must carry their baseline (HF+FA2 QLoRA vs torchtune vs stock), hardware (B200 vs 80 GB A100 vs consumer), dataset, batch, and rank — otherwise they are not facts.
- **For practitioners:** attach hardware + dataset + batch + rank + baseline to every Unsloth figure before acting on it; treat 12x as a B200-class vendor headline and the NVIDIA-blog component gains as the plannable numbers. For QLoRA on one 80 GB card, Unsloth remains the lowest-memory choice at every size tested; for throughput-maximizing research on H100, torchtune with torch.compile is competitive or faster; for multi-GPU YAML workflows with expert parallelism, Axolotl is the 2026 pick. If the goal is task-specialized MoE adaptation rather than general instruction tuning, ESFT is the method of choice. Watch the PyTorch 2.14 clamp-gradient change when upgrading mid-experiment.
- **For the brief's claims, the consolidated verdicts:** (1) fine-grained experts improve routing *without hurting inter-GPU bandwidth* — [UNVERIFIED / CONTRADICTED AS STATED] (§9). (2) Unsloth reduces VRAM by 70–80% on 70B+ fine-tuning — [PARTIALLY VERIFIED, WORKLOAD-SPECIFIC]. (3) A 70B+ model can be fine-tuned on one A100/H100 — [VERIFIED FOR QLORA], not full-parameter. (4) Unsloth enables practical MoE/70B fine-tuning on modest GPUs including RTX 4090/5090 clusters — [PARTIALLY VERIFIED] for small-MoE/8B-class on consumer GPUs; [UNVERIFIED] for 70B on a single 24 GB card and for 4090/5090 clusters.
- **Watch items:** latent-space MoE routing spreading to other labs (§9); NVFP4 MoE LoRA (Axolotl, July 2026) maturing from export-and-serve toward actual low-precision training; expert parallelism becoming a standard fine-tuning primitive; Dynamic v3.0's decode-speed/tail-quality replication caveats (recorded per Wave 2.1, not re-verified here); PyTorch 2.15 scheduled 2026-10-28.
- **What the funding/correction records mean for retrieval:** Unsloth's World's Fair presence, seed figure, and investors must never be returned as established facts. The only defensible funding statement is "founders Daniel & Michael Han; Google/OpenAI/Meta/NVIDIA are documented partners; funding figures conflict (~$40K claimed [UNVERIFIED] vs ~$500K third-party analyst figure [UNVERIFIED])."
- **Trend note [DIRECTIONAL]:** the September 2026 Studio releases shift Unsloth's public story from training-kernel speedups toward platform features (MTP-by-default, Vulkan, MLX, MCP endpoints, multi-user Docker). No new training-kernel benchmark was published after July 22, 2026 — the core training claims are older than the platform cadence suggests.

## Sources and URLs

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

