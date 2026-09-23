---
id: vague2-vision-ia/vision-ia/une-ia-vient-de-traduire-des-tablettes-vieilles-de-plus-de-4000-ans
title: "Une IA vient de traduire des tablettes vieilles de plus de 4000 ans"
domain: vision-ia
role: reference
task: article
actors: ["Anthropic", "Google", "Moonshot", "OpenRouter", "Stripe", "Z.ai"]
dates: ["2026-07", "2026-09-23"]
keywords: ["alignment", "claude", "cost", "gemini", "glm", "governance", "kimi", "mythos 5", "open source", "open weights", "protein", "research"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/une-ia-vient-de-traduire-des-tablettes-vieilles-de-plus-de-4000-ans.md
source_anchor: ""
source_lines: [1, 42]
sha256: 56049aa3e8b4c61163408fe815b077be5a1ffb89d0f577a30160a3d08b53474e
---

# Une IA vient de traduire des tablettes vieilles de plus de 4000 ans

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/une-ia-vient-de-traduire-des-tablettes-vieilles-de-plus-de-4000-ans
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter (Aug 20, 2026) leads with TabletCraft, a translation model for cuneiform Akkadian developed by Zhaohui Wang of the USC Viterbi School of Engineering. Trained on 116,000 Akkadian-to-English examples and built on the ByT5 architecture (a specialized translation model, not a general LLM; parameter count unpublished), it scores a BLEU of 49.1 on the Akkademia test set (2,812 examples), versus 37.5 for the previous best system. Crucially, it also translates English back into Akkadian and renders the output in Unicode cuneiform signs and as a virtual clay-tablet image — a first. An integrated sign converter covers 14,240 mappings (95.3% of the repertoire). Released open source under Creative Commons BY 4.0 with code, pretrained model, CLI and web demo on GitHub (`pip install cuneiscribe`). Presented in July 2026 and accepted at the ACL 2026 C3NLP workshop. Roughly 500,000 cuneiform tablets sit untranslated in museum reserves; the bottleneck is that only a few hundred people can read the script. Other headlines: Z.ai launched GLM-5.3 (Aug 14), scoring 60 on the Artificial Analysis Intelligence Index (tied with Kimi K3, +7 over GLM-5.2), but withheld the open weights because the model is too good at finding code vulnerabilities; weights expected around Aug 28. Stripe told investors the "singularity" began Jan 1, 2026, as a reason to stay private, while acquiring OpenRouter (est. $7.5–8B). Axios revealed Anthropic holds two successors to Claude Mythos 5 (Model 1 and Model 2), with Model 2 too capable for public release pending safety evals; its 186-page alignment report covers ~2,900 test sessions. The research section covers the Vesuvius Challenge reading new Philodemus works, Generalist AI's robots learning from one video, Claude designing proteins with up to 35% success, Google turning Search/Gemini into a study room, and more.

## Key points

- TabletCraft (USC) translates Akkadian↔English with BLEU 49.1 (vs 37.5 prior best); bidirectional output in cuneiform Unicode is a first.
- Built on ByT5, trained on 116,000 examples, open-sourced (CC BY 4.0) with a one-line install (`pip install cuneiscribe`).
- Z.ai's GLM-5.3 tops open models (60 on Artificial Analysis index) but delays open weights due to strong vulnerability-detection ability.
- Stripe invokes the "singularity" to stay private; acquires OpenRouter (~$7.5–8B).
- Anthropic internally holds a more capable "Model 2" it refuses to release pending safety evaluations.
- Claude achieved up to 35% success in protein design (vs 10–15% industry average), orchestrating existing specialized tools.

## Technical data / figures

| Item | Value |
|---|---|
| TabletCraft BLEU (Akkadian→English) | 49.1 (vs 37.5 prior) |
| TabletCraft BLEU (English→Akkadian) | 48.5 |
| Training examples | 116,000 |
| Akkademia test set | 2,812 examples |
| Sign converter coverage | 14,240 mappings (95.3%) |
| GLM-5.3 Intelligence Index | 60 (tied Kimi K3; +7 vs GLM-5.2) |
| GLM-5.3 cost/task | $0.68 (Kimi K3 $0.84) |
| Anthropic alignment report | ~186 pages, ~2,900 test sessions |
| Untranslated tablets | ~500,000 |
| Claude protein design success | up to 35% |

## Why this source matters for the RAG

It documents a concrete, open-source breakthrough in applying specialized AI to a millennia-old scholarly bottleneck, illustrating the shift from generalist LLMs to purpose-built translation models. It also captures frontier-lab policy signals (withheld open weights, unreleased frontier models, Stripe's positioning) useful for tracking AI governance and economics narratives.
