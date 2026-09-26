---
id: ai-industry-kb-2026/10-training-fine-tuning/sources-and-urls
title: "Sources and URLs"
domain: training-fine-tuning
role: deep-dive
task: training
actors: ["Google", "Meta", "Nvidia", "OpenAI", "Unsloth"]
dates: ["2026-07", "2026-07-22", "2026-09", "2026-10-28"]
keywords: ["benchmark", "consumer", "decode", "fine-tuning", "funding", "gpu", "gpus", "lora", "mcp", "memory", "moe", "nvfp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5699, 5707]
section: "10. Training & Fine-Tuning"
sha256: 7eb8959d94fd73aa57b5ae5db56a4fb97a99aae8dd1af7c8324ac2322e4b5c0f
---

# Sources and URLs

- **For RAG retrieval design:** architecture claims (expert counts, routing formulas) and systems claims (bandwidth, VRAM, speedups) must live as separate facts with separate evidence grades — the 2026 literature shows they move in opposite directions for fine-grained MoE. Speedup claims must carry their baseline (HF+FA2 QLoRA vs torchtune vs stock), hardware (B200 vs 80 GB A100 vs consumer), dataset, batch, and rank — otherwise they are not facts.
- **For practitioners:** attach hardware + dataset + batch + rank + baseline to every Unsloth figure before acting on it; treat 12x as a B200-class vendor headline and the NVIDIA-blog component gains as the plannable numbers. For QLoRA on one 80 GB card, Unsloth remains the lowest-memory choice at every size tested; for throughput-maximizing research on H100, torchtune with torch.compile is competitive or faster; for multi-GPU YAML workflows with expert parallelism, Axolotl is the 2026 pick. If the goal is task-specialized MoE adaptation rather than general instruction tuning, ESFT is the method of choice. Watch the PyTorch 2.14 clamp-gradient change when upgrading mid-experiment.
- **For the brief's claims, the consolidated verdicts:** (1) fine-grained experts improve routing *without hurting inter-GPU bandwidth* — [UNVERIFIED / CONTRADICTED AS STATED] (§9). (2) Unsloth reduces VRAM by 70–80% on 70B+ fine-tuning — [PARTIALLY VERIFIED, WORKLOAD-SPECIFIC]. (3) A 70B+ model can be fine-tuned on one A100/H100 — [VERIFIED FOR QLORA], not full-parameter. (4) Unsloth enables practical MoE/70B fine-tuning on modest GPUs including RTX 4090/5090 clusters — [PARTIALLY VERIFIED] for small-MoE/8B-class on consumer GPUs; [UNVERIFIED] for 70B on a single 24 GB card and for 4090/5090 clusters.
- **Watch items:** latent-space MoE routing spreading to other labs (§9); NVFP4 MoE LoRA (Axolotl, July 2026) maturing from export-and-serve toward actual low-precision training; expert parallelism becoming a standard fine-tuning primitive; Dynamic v3.0's decode-speed/tail-quality replication caveats (recorded per Wave 2.1, not re-verified here); PyTorch 2.15 scheduled 2026-10-28.
- **What the funding/correction records mean for retrieval:** Unsloth's World's Fair presence, seed figure, and investors must never be returned as established facts. The only defensible funding statement is "founders Daniel & Michael Han; Google/OpenAI/Meta/NVIDIA are documented partners; funding figures conflict (~$40K claimed [UNVERIFIED] vs ~$500K third-party analyst figure [UNVERIFIED])."
- **Trend note [DIRECTIONAL]:** the September 2026 Studio releases shift Unsloth's public story from training-kernel speedups toward platform features (MTP-by-default, Vulkan, MLX, MCP endpoints, multi-user Docker). No new training-kernel benchmark was published after July 22, 2026 — the core training claims are older than the platform cadence suggests.

## Sources and URLs

