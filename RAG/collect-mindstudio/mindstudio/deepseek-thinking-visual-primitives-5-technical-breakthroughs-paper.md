---
id: collect-mindstudio/mindstudio/deepseek-thinking-visual-primitives-5-technical-breakthroughs-paper
title: "DeepSeek's 'Thinking with Visual Primitives': 5 Technical Breakthroughs in the Paper That Briefly Disappeared"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI"]
dates: ["2024-03", "2026-09-23"]
keywords: ["deepseek", "attention", "benchmark", "claude", "consumer", "cost", "distillation", "gemini", "gpu", "kv cache", "moe", "multimodal"]
source: docs/RAG/Collect RAG/02_mindstudio/deepseek-thinking-visual-primitives-5-technical-breakthroughs-paper.md
source_anchor: ""
source_lines: [1, 51]
sha256: a0588e57a520f912a083bc1c04cef9e140c6ea19d1ae260b6e25e51210a6b946
---

# DeepSeek's 'Thinking with Visual Primitives': 5 Technical Breakthroughs in the Paper That Briefly Disappeared

## Metadata

- **Source** : https://www.mindstudio.ai/blog/deepseek-thinking-visual-primitives-5-technical-breakthroughs-paper
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article analyzes DeepSeek's "Thinking with Visual Primitives" paper — published, then pulled — extracting five technical breakthroughs. The paper describes how DeepSeek's vision model solves the "reference gap" (as opposed to the perception gap) in multimodal reasoning.

1) Visual primitives as inline tokens, not tool calls: the model emits `<ref>label</ref><box>x1,y1,x2,y2</box>` special vocabulary tokens directly in its chain of thought — not a function call or separate grounding module. When counting people in a crowded photo, the model emits a bounding box for each identified entity within its reasoning trace, so references don't drift. This addresses the "reference gap" (pointing precisely) vs the "perception gap" (seeing better) that 2024-era multimodal work targeted.

2) Training pipeline with five stages and three reward heads: the model is a 284B-parameter MoE with 13B active parameters (the DeepSeek V4 Flash backbone). Stage 1: multimodal pre-training on trillions of tokens. Stage 2: split SFT — two separate models, one for grounding (bounding boxes), one for pointing (coordinate points). Stage 3: reinforcement learning with GRPO on each specialist using three reward heads (format, quality, accuracy). Stage 4: unified RFT merging the two specialists. Stage 5: on-policy distillation into a single student model. The separation of concerns — train specialists then consolidate — mirrors MoE logic applied to training.

3) Specific and honestly scoped benchmark numbers: maze navigation — DeepSeek 67% vs Gemini Flash 3 49%, GPT-5.4 50%, Claude Sonnet 4.6 49% (~17-point gap over GPT-5.4). Path tracing similar; counting/spatial reasoning more mixed (Gemini Flash 3 ahead on raw count QA). The paper includes a footnote scoping the results to dimensions relevant to its research focus — not overall capabilities.

4) Two-year vision lineage telling one consistent story: DeepSeek VL (March 2024, 1.3B/7B, hybrid SigLIP+SAM encoder); Janus (Oct 2024, decoupled encoders for understanding vs generation); VL2 (Dec 2024, MoE + multi-head latent attention; 1B-activated version scored 80.9 OCR Bench, 88.9 DocVQA); Janus Pro 7B (Jan 2025, 80% GenEval, single consumer GPU); DeepSeek OCR (Oct 2025, 1,000 text tokens → rendered as image → 100 vision tokens reconstruct text at 97% accuracy, 10x compression; Karpathy: "the tokenizer must go, pixel may be better inputs to language models than text"). Visual primitives is the next step: compress text into pixels and make spatial coordinates first-class reasoning tokens.

5) Three admitted limitations: resolution-bound (fine-grain scenes fail; 756×756 image → 81 KV cache entries); visual primitives mode must be explicitly triggered (no auto-selection); point-based topological reasoning doesn't generalize well across all spatial tasks.

Compression architecture: DeepSeek Vision Transformer supports arbitrary resolution with 14×14 patches. A 756×756 image (~571,000 pixels) → 2,916 patch tokens → 3×3 spatial compression → 324 tokens → compressed sparse attention (from V4 paper) compresses KV cache by another 4× → 81 KV cache entries ≈ 7,000× total compression. For an 80×80 image: DeepSeek ~90 KV entries vs Claude Sonnet 4.6 ~870 vs Gemini Flash 3 ~1,000 (~10x more efficient, ~one-tenth the cost). Deployment: vision mode rolled out April 29 in limited form alongside fast/expert modes in the DeepSeek app/web.

## Key points

- Visual primitives are inline `<ref>`/`<box>` tokens in the chain of thought — the model points to objects as it reasons, solving the reference gap.
- Five-stage training pipeline with separate grounding/pointing specialists and three GRPO reward heads (format, quality, accuracy).
- Maze navigation: DeepSeek 67% vs GPT-5.4 50%, Gemini Flash 3 49%, Claude Sonnet 4.6 49%.
- ~7,000x compression from raw pixels to KV cache; ~90 vs ~870 vs ~1,000 KV entries per 80×80 image vs Sonnet 4.6 / Flash 3.
- Backbone: 284B-param MoE with 13B active parameters (DeepSeek V4 Flash).
- Three admitted limitations: resolution-bound, explicit triggering required, limited topological generalization.

## Technical data / figures

| Element | Value |
|---|---|
| Backbone | DeepSeek V4 Flash, 284B MoE / 13B active |
| Maze navigation | DeepSeek 67%, GPT-5.4 50%, Flash 3 49%, Sonnet 4.6 49% |
| 756×756 image → patch tokens | 571k px → 2,916 → 324 (3×3 compression) → 81 KV (4× sparse attention) |
| Compression ratio | ~7,000x pixels → KV cache |
| 80×80 KV entries | DeepSeek ~90, Sonnet 4.6 ~870, Flash 3 ~1,000 |
| DeepSeek OCR lineage | 1,000 text tokens → 100 vision tokens, 97% reconstruction |
| VL2 (1B active) | OCR Bench 80.9, DocVQA 88.9 |

## Why this source matters for the RAG

Explains the technical foundations (inline spatial tokens, five-stage training, compression architecture) behind DeepSeek's vision model — key context for multimodal RAG on image/diagram-heavy documents. Provides per-image KV cache cost figures essential for comparing vision-model economics in document pipelines.
