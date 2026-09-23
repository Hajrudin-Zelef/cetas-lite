---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/9-distributed-training-systems-deepspeed-fsdp-megatron-lm-to
title: "9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: training
actors: ["Meta", "Microsoft", "Nvidia"]
dates: []
keywords: ["training", "benchmarks", "moe", "nvidia", "tpu"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [608, 661]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 630cc153bc92e0243930f39ce5bc9ca701709c1cfeac173e60180571115e14c9
---

# 9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan

## 9. Distributed training systems: DeepSpeed, FSDP, Megatron-LM, TorchTitan

### 9.1 DeepSpeed

- DeepSpeed (Microsoft) documentation shows **release 0.19.5 dated August
  6, 2026**, confirming active 2026 development [secondary].
- Core techniques: **ZeRO** stages (optimizer/gradient/parameter
  partitioning across data-parallel ranks), CPU/NVMe offload, activation
  checkpointing, pipeline parallelism, and MoE support [secondary].
- Historical vendor performance claims (e.g., "100B-parameter training
  10× faster" from the ZeRO-2 era) are era-labeled and must not be read
  as 2026 benchmarks — recorded in the gaps register (G-14) [secondary].
- DeepSpeed remains the pragmatic choice for teams training on PyTorch
  without adopting Meta's TorchTitan or NVIDIA's Megatron-LM [secondary].

### 9.2 FSDP (native PyTorch)

- FSDP2 (`fully_shard`) is stock PyTorch's sharded-data-parallel answer
  to ZeRO-3, integrated with `torch.compile` and DTensor (Section 1.2)
  [secondary].
- The 2026 trend is FSDP2 + `torch.compile` as the default large-model
  recipe inside the PyTorch-native ecosystem (notably TorchTitan)
  [secondary].

### 9.3 Megatron-LM

- Megatron-LM is NVIDIA's framework for tensor/pipeline/data-parallel
  training of large transformers, the reference implementation behind
  many frontier training runs [secondary].
- No verified 2026 version captured in this pass — gap G-15 [unverified].
- Its ecosystem includes Megatron-DeepSpeed bridges and the NeMo
  framework for curated training workflows [secondary].

### 9.4 TorchTitan

- TorchTitan is Meta's PyTorch-native platform for large-scale LLM
  training, built on FSDP2, `torch.compile`, and DTensor-based
  parallelism [secondary].
- No verified 2026 release state captured in this pass — gap G-16
  [unverified].
- Positioning: the "stay inside stock PyTorch" alternative to DeepSpeed
  and Megatron-LM for teams standardizing on the PyTorch 2.x compiler
  stack [secondary].

### 9.5 Distributed-training decision matrix

| System | Parallelism model | Base | Best fit |
|---|---|---|---|
| DeepSpeed 0.19.x | ZeRO 1/2/3, offload, pipeline, MoE | PyTorch | Pragmatic large-model training [secondary] |
| FSDP2 (stock) | Sharded data-parallel | PyTorch 2.x | PyTorch-native, compiles well [secondary] |
| Megatron-LM | Tensor+pipeline+data | PyTorch | Frontier-scale, NVIDIA-tuned [secondary] |
| TorchTitan | FSDP2 + compile + DTensor | PyTorch 2.x | Meta-style native stack [secondary] |
| JAX `pjit` | Named-axis sharding | XLA/TPU | TPU-first training (see §2.3) [secondary] |

