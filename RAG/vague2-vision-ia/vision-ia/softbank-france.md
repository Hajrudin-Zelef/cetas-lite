---
id: vague2-vision-ia/vision-ia/softbank-france
title: "SoftBank promet 75 milliards d'euros à la France : révolution IA ou mirage ?"
domain: vision-ia
role: reference
task: article
actors: ["Anthropic", "Apple", "China", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI"]
dates: ["2026-06-01", "2026-09-23"]
keywords: ["agent", "agents", "benchmark", "chatgpt", "claude", "compute", "consumer", "datacenter", "gemini", "gpu", "gpus", "humanoid"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/softbank-france.md
source_anchor: ""
source_lines: [1, 49]
sha256: 98b0485bad85969841ab82a560354ceab001f9d4f0719c7246471dbaa339e1a6
---

# SoftBank promet 75 milliards d'euros à la France : révolution IA ou mirage ?

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/softbank-france
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This June 1, 2026 edition of the Vision-IA newsletter leads with SoftBank planning to build AI data centers in France with a total capacity of 5 GW—reportedly its largest AI infrastructure investment in Europe. The Japanese group mentions €75 billion, with €45 billion of operational facilities expected by 2031 across three sites concentrated in Normandy. If realized, France would become a major European AI compute hub with direct effects on digital sovereignty, local jobs, and attractiveness to labs. The newsletter cautions that SoftBank has a habit of announcing spectacular megaprojects that struggle to materialize, so it advises waiting for the first shovels in Normandy.

Nvidia entered the humanoid robotics race with Chinese startup Unitree, launching its first publicly accessible humanoid robotics platform based on Unitree's hardware—as Unitree prepares an IPO. Nvidia now provides the full platform, not just GPUs, potentially accelerating humanoid adoption, and its choice of a Chinese partner says a lot about who leads consumer robotics hardware. Nvidia also unveiled its first ARM-based chip for laptops (with Microsoft, Dell, HP, ASUS), attacking the x86 PC market with on-device AI as the selling point. Google made Nano Banana 2 (Gemini 3.1 Flash Image) and Nano Banana Pro (Gemini 3 Pro Image) generally available, now accepting video as a prompt input to analyze subjects, actions, and visual styles—all in Google Cloud's enterprise infrastructure.

The "Research" section covers Biohub (Chan Zuckerberg initiative, in Meta's orbit) publishing Evolutionary Scale Models to predict and design proteins; its flagship ESMFold2, trained on 2.8 billion protein sequences, achieves SOTA structure prediction with applications against cancer and immune diseases. Kog AI announced an inference optimization reaching 3,000 tokens/s per user on standard datacenter GPUs (10-30x the usual 100-300 tokens/s), exploiting a hidden inefficiency in GPU token generation; demonstrated on a 2B model, expected to scale to large MoE models. Datacurve launched DeepSWE, a "long-horizon" software engineering benchmark built on novel tasks (not public GitHub issues) to avoid training contamination: GPT-5.5 at 70%, GPT-5.4 at 56%, Claude Opus 4.7 at 54%—differences old tests masked. NVIDIA published Cosmos 3, an open-source omnidirectional model for robotics and physical AI, generating/understanding video to simulate autonomous agent behavior.

Additional briefs include Apple rebuilding Siri on Google Gemini (generative capabilities like ChatGPT, in Dynamic Island, third-party agents, on-device data); Blue Origin's New Glenn exploding during a ground engine test at Cape Canaveral, destroying the booster and the only operational launch pad, hitting NASA's Artemis program; Kaikaku.AI's Epicure (three models trained on 4.14M recipes and FlavorDB—the pure chemistry model unexpectedly beats recipe models at ranking tastes and nutritional values); Boston Dynamics' Atlas learning football from videos, including the "Ghost Rabona" (planned at the 2026 World Cup); Anthropic banning AI during its job interviews (up to 5 rounds, salaries up to $850,000, paid coaching at $4,600); Bonsai Image 4B (4B-parameter 1-bit quantized local image generation, 357 HN upvotes); Turkey's hair-transplant industry using machine learning; and building a full game with Codex in a single prompt using "goals."

## Key points

- SoftBank plans €75B of AI data centers in France (5 GW total), with €45B operational by 2031 across three Normandy sites.
- The newsletter cautions that SoftBank's megaprojects often struggle to materialize.
- Nvidia enters humanoid robotics with Unitree, launching its first public humanoid platform as Unitree preps an IPO.
- Nvidia unveils its first ARM PC chip with Microsoft, Dell, HP, and ASUS, targeting the x86 laptop market.
- Google's Nano Banana 2 and Pro now accept video as prompt input for image generation/editing.
- Biohub's ESMFold2 (2.8B protein sequences) achieves SOTA protein structure prediction.
- Kog AI claims 3,000 tokens/s per user on standard GPUs (10-30x normal).
- DeepSWE benchmark reveals real model gaps: GPT-5.5 70%, GPT-5.4 56%, Claude Opus 4.7 54%.

## Technical data / figures

| Item | Figure |
|---|---|
| SoftBank France investment | €75B total, €45B by 2031, 5 GW |
| France sites | 3 in Normandy |
| ESMFold2 training | 2.8 billion protein sequences |
| Kog AI inference | 3,000 tokens/s/user (10-30x) |
| DeepSWE scores | GPT-5.5 70%, GPT-5.4 56%, Opus 4.7 54% |
| Epicure training | 4.14M recipes + FlavorDB |
| Bonsai Image 4B | 4B params, 1-bit, 357 HN upvotes |
| Anthropic interview salary | up to $850,000; coaching $4,600 |
| Unitree H2 Plus | 1.80 m (in later edition) |

## Why this source matters for the RAG

It documents a huge, if uncertain, AI infrastructure bet in Europe (SoftBank/France) that bears on digital sovereignty and the compute race. It also captures hardware/architecture diversification (Nvidia ARM PC chip, Nvidia-Unitree humanoid platform) and rigorous benchmarking/efficiency advances (DeepSWE, Kog AI, ESMFold2) that matter for evaluating real progress beyond marketing claims.
