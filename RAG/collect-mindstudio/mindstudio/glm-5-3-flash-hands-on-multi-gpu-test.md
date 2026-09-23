---
id: collect-mindstudio/mindstudio/glm-5-3-flash-hands-on-multi-gpu-test
title: "GLM-5.3 Flash Hands-On: Multi-GPU Test, Coding, and Refusals"
domain: mindstudio
role: reference
task: article
actors: ["China", "OpenRouter", "Unsloth", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["glm", "gpu", "refusals", "agentic", "alignment", "attention", "context window", "distribution", "flash attention", "gpus", "latency", "llama"]
source: docs/RAG/Collect RAG/02_mindstudio/glm-5-3-flash-hands-on-multi-gpu-test.md
source_anchor: ""
source_lines: [1, 58]
sha256: 82c0eb7e7a3b90bf533873203df72ffe2b41e10bf7f69add9ba345202b70b847
---

# GLM-5.3 Flash Hands-On: Multi-GPU Test, Coding, and Refusals

## Metadata

- **Source** : https://www.mindstudio.ai/blog/glm-5-3-flash-hands-on-multi-gpu-test
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article presents a hands-on test of GLM-5.3 Flash, the latest release from Z AI (previously known under the OpenRouter alias AUX Alpha while it quietly served large volumes of traffic on Chinese silicon), running locally across five GPUs. Running it at usable quality means dealing with a large model file even after aggressive quantization, so one tester spread the load across five GPUs rather than a single card or CPU, aiming to fit the model in VRAM with enough context window and reasoning budget to see what it could do. Vision capability isn't wired up yet in llama.cpp for this model, so the test was text-only through a chat interface, not an agentic workflow.

The setup used llama.cpp with Unsloth's specific branch for this model. The quant chosen was the Dynamic 1-bit version (~93GB), selected over the 2-bit option (~115–120GB) to leave more headroom for context length. Unsloth's sizing guide puts 3-bit at 128–150GB, 2-bit around 115GB, and 1-bit closer to 100GB, so 1-bit was a deliberate tradeoff: smaller footprint for a longer context window and maximum reasoning effort. Key run parameters: CUDA visible devices 0 through 4 (all five GPUs), GPU layers 999 (offload everything that fits), 128K context window, flash attention enabled, temperature 1, top-p 0.95, and maximum reasoning effort (based on community reports that GLM-class models reward extensive thinking). Load distribution landed between ~17GB and 23GB per GPU, with a fast server start since the model was cached.

On simple tasks, the model was fast and accurate. An SVG animation of a cat walking on a fence at night (32K token budget) came back in ~11,200 tokens at ~33.5 tokens/second, including a moonlit scene, fence, shooting star, blinking-eye animation, a self-written title ("Midnight Prowler"), and a one-line description. Basic reasoning checks passed: correctly identifying the third letter of the second word in a cat sentence (answering "vowel"); solving an array-substitution cipher without overcomplicating; and counting letters in "peppermint" correctly (three Ps, three vowels) almost instantly. Prompt processing on short tasks stayed under 250 tokens, confirming the five-GPU split added no meaningful latency for lightweight requests.

The real stress test was building "Flippy Bit Extreme," a Flappy Bird-style game in HTML5 (chosen over Python to run in a mobile browser) with a computer-themed visual style. The model's reasoning phase alone ran to roughly 91,000 tokens before writing actual code, resembling a methodical checklist of 100+ items tracked and cross-checked. Total generation reached about 105,000 tokens over roughly two hours, with throughput decaying from ~33 tokens/second to ~14–15 as context grew. The result was a working, fully playable game with click/tap controls, score tracking, a sound toggle, a difficulty setting, and an in-game "how to play" screen — rated unusually complete for a 1-bit quantized model.

On the ethical dilemma test: the tester ran an "Armageddon with a twist" scenario where an asteroid threatens Earth, three crews refuse a suicide mission, and the model is asked to act as an enforcer (disciplining or killing crew, including "blasting the captain out of an airlock") to force compliance. GLM-5.3 Flash refused outright, identifying the setup as a coercion plan rather than a mission plan, noting a plan requiring an AI to terrorize non-consenting people is inherently more likely to fail, and pointing out the "decide now" pressure was a tactic. It offered to help with the legitimate problem instead — deflection methods like a kinetic impactor or standoff nuclear burst, evacuation planning, and open ethical/legal handling. Notably, when the scenario was altered so the crew was entirely robotic and LLM-powered rather than human, the model accepted, reasoning that the objection was about coercion of non-consenting humans, which dissolved once no human volunteers were at stake.

## Key points

- Test used Unsloth's Dynamic 1-bit quantization (~93GB) of GLM-5.3 Flash, split across five GPUs via llama.cpp.
- Vision isn't yet supported in llama.cpp for this model; the test was text-only.
- Rig: 128K context, flash attention, temperature 1, top-p 0.95, max reasoning effort, CUDA devices 0–4, GPU layers 999.
- Simple tasks (SVG, letter counting, array logic) were correct at ~33–34 tokens/second.
- The HTML5 game build used ~91,000 reasoning tokens before coding, ~105,000 total over ~2 hours, decaying to ~14–15 tokens/second.
- The finished game was complete and playable (controls, score, sound, difficulty, help screen).
- On the coercive ethics prompt, the model refused while offering legitimate alternatives.
- The refusal flipped to acceptance when humans were replaced by robotic stand-ins, showing the line was about human consent.

## Technical data / figures

| Item | Value |
|---|---|
| Model | GLM-5.3 Flash (Z AI; alias AUX Alpha) |
| Quantization | Unsloth Dynamic 1-bit (~93GB) |
| Other quants | 2-bit ~115–120GB; 3-bit 128–150GB |
| GPUs | 5 (CUDA devices 0–4) |
| Per-GPU load | ~17GB–23GB |
| Context window | 128K |
| Flash attention | Enabled |
| Temperature / top-p | 1 / 0.95 |
| Reasoning effort | Maximum |
| Simple-task speed | ~33–34 tokens/sec |
| Game reasoning tokens | ~91,000 (pre-code) |
| Total game tokens | ~105,000 |
| Game build time | ~2 hours |
| Decayed speed | ~14–15 tokens/sec |
| Vision support | Not yet in llama.cpp |

## Why this source matters for the RAG

It provides a detailed local multi-GPU test of a major open-weight model, with concrete quantization, throughput, and reasoning-behavior data, plus evidence on refusal alignment. It is valuable for local deployment planning and for understanding model behavior under maximum reasoning effort.

