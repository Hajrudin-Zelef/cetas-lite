---
id: collect-huggingface/huggingface/hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf
title: "Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "agentic", "apache", "attention", "benchmarks", "blackwell", "license", "llama", "llama.cpp", "open-weight", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/HauhauCS-Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF.md
source_anchor: ""
source_lines: [1, 52]
sha256: 488c077ec75a5c09f872063ad906ffdc16e23162e3a531d9bc35baf09ab9d135
---

# Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/HauhauCS/Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is HauhauCS's uncensored GGUF release of Qwen3.8-27B in the "Aggressive" variant (direct answers, no refusal behavior, minimal preamble), reporting 0/465 refusals. The model preserves Qwen3.8-27B's native text, reasoning, agentic, image, and video capabilities while applying the HauhauCS Aggressive uncensoring profile. It is a dense 27B causal language model with a vision encoder (Image-Text-to-Text): 64 layers, hidden size 5,120, FFN 17,408, padded vocab 248,320, 48 Gated DeltaNet layers + 16 gated-attention layers, 262,144 native context extensible to 1M, and native embedded MTP/NextN preserved. Base model: Qwen/Qwen3.8-27B; license Apache 2.0.

Every text GGUF preserves Qwen3.8's native NextN head; this release adds HauhauCS FastMTP, a custom acceleration sidecar (903 MB, 32K draft) qualified across the full quant lineup at maximum native context. Benchmarks (RTX PRO 6000 Blackwell, 204800 context): FastMTP delivers up to 3.02x document TG and 1.93x reasoning TG vs non-MTP, and up to +35.2% document / +21.1% reasoning TG vs standard embedded MTP (e.g., Q8_K_P: 138.18 document TG, 90.07 reasoning TG, 3.02x/1.93x vs MTP-off). Vision is provided through a separate BF16 projector (mmproj, 931 MB). K_P ("Perfect") quants are model-specific profiles bumping quality one or two levels at only 5-15% more size, standard GGUF compatible (may display as "?" in LM Studio).

File lineup: Q8_K_P 31.46 GB (9.21 BPW), Q6_K_P 25.92 GB, Q5_K_P 20.22 GB, Q4_K_P 17.92 GB, IQ4_XS 15.71 GB, Q3_K_P 13.44 GB, IQ3_M 12.79 GB, IQ3_XS 12.18 GB, Q2_K_P 10.68 GB, IQ2_M 10.32 GB, plus the FastMTP sidecar and BF16 vision projector. FastMTP requires the HauhauCS-FastMTP-llama.cpp.patch on a pinned llama.cpp commit (4df29be). Recommended settings: thinking mode temp 1.0 / top_p 0.95 / top_k 20 / reasoning_effort xhigh; non-thinking temp 0.7 / top_p 0.8 / presence_penalty 1.5 / enable_thinking=false. Thinking on by default; disable via chat_template_kwargs. Authenticity via signed release manifest (Ed25519). Hub: 2B hub-tagged params (27B model), 2.05M downloads/month.

## Key points

- Uncensored Aggressive variant of Qwen3.8-27B: 0/465 refusals, direct answers.
- Dense 27B causal LM + vision encoder; 48 Gated DeltaNet + 16 gated-attention layers.
- 262K native context (extendable to 1M); native embedded MTP/NextN preserved.
- HauhauCS FastMTP sidecar: up to 3.02x document / 1.93x reasoning TG vs MTP-off.
- K_P "Perfect" quant profiles boost quality 1-2 levels at 5-15% more size.
- Quants from IQ2_M (10.32 GB) to Q8_K_P (31.46 GB); separate BF16 vision projector.
- Apache 2.0 (inherited from Qwen3.8-27B).
- Requires llama.cpp patch for FastMTP; thinking mode default, disable via enable_thinking.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | HauhauCS (community) |
| Model name | Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF |
| Base model | Qwen/Qwen3.8-27B |
| Architecture | Dense causal LM + vision encoder (Gated DeltaNet + gated attention) |
| Total params | 27B |
| Layers | 64 |
| Hidden size | 5120 |
| Vocabulary | 248,320 (padded) |
| Context length | 262,144 native; up to 1M |
| FastMTP | up to 3.02x doc / 1.93x reason TG vs MTP-off |
| Quants | IQ2_M (10.32 GB) to Q8_K_P (31.46 GB) + K_P line |
| Vision projector | BF16, 931 MB |
| License | Apache 2.0 |
| Downloads/month | 2,052,516 |

## Why this source matters for the RAG

This card documents a highly popular community uncensored variant of a frontier open-weight model, including custom K_P quantization methodology, a FastMTP speculative-decoding sidecar with detailed speedups, and signed-release authenticity practices. It is valuable reference material for retrieval on uncensored model variants, GGUF quantization, and speculative decoding acceleration.
