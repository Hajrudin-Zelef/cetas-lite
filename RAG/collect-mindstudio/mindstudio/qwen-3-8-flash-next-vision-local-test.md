---
id: collect-mindstudio/mindstudio/qwen-3-8-flash-next-vision-local-test
title: "Qwen 3.8 Flash Next Vision at Q4: Does Quantization Cost Accuracy?"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["cost", "quantization", "qwen", "agents", "benchmark", "consumer", "context window", "gpu", "gpus", "inference", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/02_mindstudio/qwen-3-8-flash-next-vision-local-test.md
source_anchor: ""
source_lines: [1, 57]
sha256: 0a8f969c2054cc49280f4690fce258287f3635d1430330737c059fc8a53c1cdb
---

# Qwen 3.8 Flash Next Vision at Q4: Does Quantization Cost Accuracy?

## Metadata

- **Source** : https://www.mindstudio.ai/blog/qwen-3-8-flash-next-vision-local-test
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article presents a hands-on quad-RTX-3090 test of Qwen 3.8 Flash Next's vision support at Q4 quantization, compared against a full-precision Qwen 3.8 27B model. Qwen 3.8 Flash Next is a large multimodal (image-text-to-text) model from Alibaba's Qwen team, distributed on Hugging Face across 131 safetensors shards. It recently gained working vision support in llama.cpp, merged from an Unsloth branch, along with meaningful inference speed gains. The article notes it's rare to get both a capability update and a real throughput bump in the same merge. The catch: most people run the model quantized, not at full precision, and quantization has historically hit vision accuracy harder than text accuracy.

Test setup: the run was entirely local on a rig with four RTX 3090 GPUs (a fifth card, an RTX 4090, present but unused), managed through a Proxmox host and driven by llama.cpp built fresh from source (the vision merge and performance improvements weren't in a stable release at testing time). Key configuration: mmproj set to BF16 even though the main model was quantized; mmap loading enabled; context window set to the full 262,144 tokens; parallel set to 1, batch size 2048, ubatch 512; reasoning effort set to high with a generous min/max image token budget. The BF16 mmproj detail matters: even in a quantized setup, keeping the vision projector at higher precision limits some accuracy loss, though it didn't eliminate the gap versus full precision.

Token generation speed roughly jumped from 43.99 to 62.42 tokens/second in one creator's benchmark after the merge and follow-up updates, independent of the vision feature. The comparison point was Qwen 3.8's 27B variant at full precision in an earlier test, which produced "amazing results" — correctly identifying a photo of a camel as consistent with Texas Hill Country terrain (shrub type, single-hump anatomy), described as nailing the call immediately. The Q4 Flash Next model, tested on the same camel photo, got the broad strokes right (dromedary, one hump, shrubs, likely Texas) but hedged toward South Texas generally rather than confidently landing on Hill Country. The creator scored it good but not as sharp as full precision.

That pattern repeated. General scene understanding and object classification were largely accurate: correctly identifying a false robber fly by name; identifying an assassin bug (hedging on some family-level detail); naming all six ingredients in a stir-fry photo (bell peppers, mushrooms, avocado, onions, cheese, green beans); accurately describing a griddle of burger patties (melted cheese, peppers, pan drippings) though slightly fumbling which toppings were on which patty. Fine-precision tasks struggled: misreading a Kill A Watt meter's LCD display (landing close but wrong on the digit); failing to identify a landscape as Lake Austin, initially proposing the visually similar Lake Travis before backing off; and declining to commit to a specific GPU model given a close-up of a bare GPU PCB with two 8-pin power connectors, where the creator expected the full-precision 27B would likely have identified it directly. Location/time-period reasoning without EXIF data was a high point: correctly narrowing an old Austin skyline photo to around 2008–2012 based on which buildings were present.

The article concludes Q4 is worth it for general-purpose image description, object recognition, and casual visual Q&A, especially paired with the throughput gains (50+ tokens/second on consumer GPUs is a practical win). For reading fine text, distinguishing near-identical scenes, or identifying specific hardware by exact model, the quantization cost shows up clearly. The recommended next test is running the full safetensors version of Qwen 3.8 Flash Next in split mode across the same rig, unquantized, to isolate whether the accuracy gap comes from quantization specifically or other model differences.

## Key points

- Qwen 3.8 Flash Next vision support landed in llama.cpp via a merged Unsloth branch, requiring a fresh source build.
- Speed roughly jumped from 43.99 to 62.42 tokens/second after the merge, independent of vision.
- Tested Q4 against an earlier full-precision Qwen 3.8 27B run; Q4 was noticeably less precise on detail-heavy tasks.
- General scene description, object ID, and ingredient listing held up well at Q4 (false robber fly, assassin bug, six stir-fry ingredients).
- Fine tasks failed or landed close-but-wrong: misread a Kill A Watt LCD; failed to identify Lake Austin; hedged on a GPU PCB model.
- Location/time-period reasoning without EXIF was strong (narrowed an Austin skyline to ~2008–2012).
- Test rig: quad RTX 3090, llama.cpp from source, 262,144-token context, BF16 mmproj, reasoning high.
- Next step: test full safetensors in split mode to isolate quantization's effect.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Qwen 3.8 Flash Next (multimodal) |
| HF shards | 131 safetensors |
| Vision merge source | Unsloth branch → llama.cpp |
| Speed before/after | 43.99 → 62.42 tokens/sec |
| Quantization | Q4 |
| Comparison model | Qwen 3.8 27B (full precision) |
| GPUs | 4x RTX 3090 (4090 unused) |
| Host | Proxmox |
| Context window | 262,144 tokens |
| mmproj | BF16 |
| parallel / batch / ubatch | 1 / 2048 / 512 |
| Reasoning effort | High |
| Strengths | Scene description, object ID, ingredients, location reasoning |
| Weaknesses | LCD reading, exact GPU ID, near-identical scenes |

## Why this source matters for the RAG

It provides concrete evidence on how Q4 quantization affects vision accuracy versus full precision, with specific task-level successes and failures. This is directly useful for anyone deciding quantization levels for local multimodal agents and for understanding the precision/throughput tradeoff.

