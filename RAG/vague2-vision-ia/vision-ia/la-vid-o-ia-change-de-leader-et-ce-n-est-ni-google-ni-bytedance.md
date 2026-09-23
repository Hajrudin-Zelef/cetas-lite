---
id: vague2-vision-ia/vision-ia/la-vid-o-ia-change-de-leader-et-ce-n-est-ni-google-ni-bytedance
title: "La vidéo IA change de leader, et ce n'est ni Google ni ByteDance"
domain: vision-ia
role: reference
task: article
actors: ["Anthropic", "ByteDance", "Google", "Meta", "Mistral", "Moonshot", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "apache", "benchmarks", "claude", "consumer", "cost", "gemini", "gpu", "kimi", "mistral", "muse"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/la-vid-o-ia-change-de-leader-et-ce-n-est-ni-google-ni-bytedance.md
source_anchor: ""
source_lines: [1, 44]
sha256: 4a4070ecd64f3779b600787fc1963f63b34cde1f78b6ae7027101967683e4a45
---

# La vidéo IA change de leader, et ce n'est ni Google ni ByteDance

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/la-vid-o-ia-change-de-leader-et-ce-n-est-ni-google-ni-bytedance
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter (Aug 6, 2026) leads with Black Forest Labs (BFL), the German studio behind the FLUX image generators, opening FLUX 3 Video to all users after restricted access since late July. The model produces Full HD clips up to 20 seconds with native audio and synchronized dialogue in over 14 languages. On BFL's own Elo rankings it claims 1st in text-to-video (1135 points) and 1st in image-to-video (1051), ahead of ByteDance's Seedance 2.0, Google's Gemini Omni Flash and Minimax H3. It generates from text, image, keyframes, or video continuation, with multiple scenes and camera angles in one clip; audio (dialogue, SFX, ambience) is produced alongside video; lip-sync works in 14+ languages and it can embed legible text into scenes. Pricing is per second: $0.06/s Draft HD, $0.17/s full-quality HD, $0.29/s Full HD (audio included) — a 10s Full HD clip costs ~$2.90. Access is via the BFL API and select partners; no consumer app yet. The newsletter cautions the Elo ranking is self-reported. Other major items: on Aug 5, Google DeepMind lost both Demis Hassabis (moving to Alphabet chief scientist, refocusing on Isomorphic Labs) and Jeff Dean (leaving after 27 years to found Discovery Loop, joined by Sanjay Ghemawat, Oriol Vinyals and Quoc Le); Koray Kavukcuoglu becomes SVP of Google DeepMind. Meta launched Muse Code, a terminal-based coding agent built on Muse Spark 1.2, scoring 82.9% on Terminal-Bench 2.1 (vs 86.7% for Claude Opus 5). Mistral released Shieldstral 1.0, an open 3B-parameter safety classifier (Apache 2.0) running on a single 16 GB GPU. Additional briefs cover Google Assistant's Sept 4 shutdown in favor of Gemini, the Inouye solar telescope, NASA's Roman telescope as asteroid hunter, NVIDIA's Alpamayo 2 Super, Kimi K3's 2,800B-parameter technical report, and more.

## Key points

- Black Forest Labs' FLUX 3 Video tops BFL's own Elo rankings (1135 T2V, 1051 I2V), ahead of Seedance 2.0, Gemini Omni Flash and Minimax H3.
- Full HD clips up to 20s with native synchronized audio and dialogue in 14+ languages; pricing $0.06–$0.29/s.
- Google DeepMind loses Hassabis (to Alphabet chief scientist) and Jeff Dean (founds Discovery Loop) on the same day.
- Meta enters the coding-agent race with Muse Code (Terminal-Bench 2.1: 82.9% vs Opus 5's 86.7%).
- Mistral's Shieldstral 1.0 is a 3B open safety classifier (Apache 2.0) runnable on one 16 GB GPU with natural-language rules.
- Google Assistant is discontinued Sept 4, 2026, replaced everywhere by Gemini.

## Technical data / figures

| Item | Value |
|---|---|
| FLUX 3 Video max length | 20s Full HD |
| FLUX 3 Elo text-to-video | 1135 (1st) |
| FLUX 3 Elo image-to-video | 1051 (1st) |
| FLUX 3 pricing | $0.06 / $0.17 / $0.29 per second |
| 10s Full HD clip cost | ~$2.90 |
| Languages (lip-sync) | 14+ |
| Muse Code Terminal-Bench 2.1 | 82.9% (Opus 5: 86.7%) |
| Muse Code API pricing | $1.25/M in, $0.15/M cache, $4.25/M out |
| Shieldstral params | 3B (Ministral-3B + Pixtral vision) |
| Shieldstral F1 | 84.9% text, 83.8% image+text, 91.3% unseen rules |
| Shieldstral training | ~54.1M harmful examples, 12 languages |
| Kimi K3 params | 2,800B (47-page report) |

## Why this source matters for the RAG

It captures a genuine leadership shift in generative video toward a European lab, with concrete pricing and capability benchmarks that matter for media-production use cases. It also records a pivotal day of leadership departures at Google DeepMind and the entry of Meta and Mistral into coding agents and safety tooling, useful for mapping competitive dynamics.
