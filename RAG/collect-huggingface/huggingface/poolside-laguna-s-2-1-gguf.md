---
id: collect-huggingface/huggingface/poolside-laguna-s-2-1-gguf
title: "Laguna S 2.1 GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["gguf", "agentic", "attention", "context window", "embeddings", "license", "llama", "llama.cpp", "moe", "open-weight", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/poolside-Laguna-S-2.1-GGUF.md
source_anchor: ""
source_lines: [1, 50]
sha256: 27eeaca41ce8b7c35a9d8d9fe9b05c812ccfb8734d27cc1a3d3a8058faa4381b
---

# Laguna S 2.1 GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/poolside/Laguna-S-2.1-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Laguna S 2.1 GGUF is poolside's official llama.cpp/GGUF conversion of the Laguna S 2.1 Mixture-of-Experts model (117.6B total / 8.5B active parameters), plus the DFlash speculative-decoding draft model. It is designed for agentic coding and long-horizon work on a local machine, and is published under the OpenMDW-1.1 license (commercial and non-commercial use permitted). The conversion includes four main files: a full-precision F16 GGUF at 235 GB; Q8_0 at 129 GB (routed experts in Q8_0, signal path — attention, shared experts, embeddings — kept in BF16); Q4_K_M at 68 GB (routed experts in Q4_K via imatrix, signal path in Q8_0); and a 2.2 GB DFlash BF16 drafter for speculative decoding, plus a 0.4 GB importance matrix file used for the K-quants.

The GGUFs ship configured for a 262,144-token (256K) context window, which poolside recommends for best output quality. The underlying weights are native 1M checkpoints — training included a long-context extension stage up to 1,048,576 tokens — and users can override the RoPE configuration at load time (`--ctx-size 1048576 --rope-scaling yarn --rope-scale 128 --yarn-orig-ctx 8192`) to access the full 1M context, with some expected quality degradation (recommended sampling `--temp 0.7 --top-p 0.95`). The release corrects the embedded `yarn_attn_factor` metadata (now 1.0; llama.cpp derives YaRN attention scaling internally).

Serving is done with poolside's llama.cpp fork on the `laguna` branch, which carries full Laguna support including DFlash speculative decoding (`--spec-type draft-dflash --spec-draft-n-max 15 --spec-draft-override-tensor '.*=CUDA0'`). Base Laguna support is also in upstream review (ggml-org/llama.cpp#25165). The architecture tag is `laguna`; the hub lists hardware compatibility: Q4_K_M needs ~96 GB, Q8_0 ~129 GB, F16 ~235 GB. It is a highly popular download with 657,230 downloads in the last month, and is the base for 94 downstream quantizations plus Ollama, LM Studio, Jan, Unsloth, and Pi integrations. The card points to the base model card (poolside/Laguna-S-2.1) for architecture details, license and full usage guidance, and to the release blog post.

## Key points

- GGUF conversions of Laguna S 2.1 (117.6B-A8.5B MoE) for llama.cpp, incl. DFlash drafter.
- Files: F16 235 GB, Q8_0 129 GB, Q4_K_M 68 GB, DFlash-BF16 2.2 GB, imatrix 0.4 GB.
- Default 256K context; native 1M checkpoints reachable via YaRN RoPE override.
- Served via poolside's llama.cpp `laguna` branch (supports DFlash speculative decoding).
- OpenMDW-1.1 license; commercial and non-commercial use allowed.
- 657K downloads/month; 94 downstream quantizations.
- Corrected `yarn_attn_factor` metadata (1.0).

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | poolside |
| Model name | Laguna S 2.1 GGUF |
| Architecture | laguna (MoE, SWA + global attention, 48 layers) |
| Total params | 117.6B |
| Active params | 8.5B |
| Hub model size | 118B params |
| Context length | 256K default; 1M (1,048,576) via RoPE override |
| License | OpenMDW-1.1 |
| Quantizations | F16 (235 GB), Q8_0 (129 GB), Q4_K_M (68 GB), DFlash-BF16 (2.2 GB) |
| Hardware | Q4_K_M ~96 GB, Q8_0 ~129 GB, F16 ~235 GB |
| Base model | poolside/Laguna-S-2.1 |
| Downloads/month | 657,230 |
| Serving | poolside llama.cpp `laguna` branch; Ollama, LM Studio, Jan, Pi, OpenClaw, Hermes |

## Why this source matters for the RAG

This card is the authoritative reference for running the open-weight Laguna S 2.1 locally, documenting exact GGUF file sizes, quantization strategy (signal path vs. routed experts), 256K vs. 1M context configuration, and DFlash speculative decoding with llama.cpp. It provides practical, citable deployment data for a strong open agentic coding model on commodity hardware.
