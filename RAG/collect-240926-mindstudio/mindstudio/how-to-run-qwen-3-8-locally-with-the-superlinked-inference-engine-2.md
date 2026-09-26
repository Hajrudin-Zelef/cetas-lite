---
id: collect-240926-mindstudio/mindstudio/how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine-2
title: "how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["inference", "qwen", "gpu", "gpus", "speculative decoding"]
source: docs/RAG/clean_en/mindstudio/how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine.md
source_anchor: ""
source_lines: [67, 87]
sha256: ae5a21f2a9b5a3ffa8604d078d3aa5ddbd00df5c7a084fd529ca94471edb7489
---

# how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine

Qwen 3.8 is a large language model in Alibaba’s Qwen family, with a 27 billion parameter version used in the SIE demonstration for both general question answering and larger code/content generation tasks.

### Does SIE support GPUs other than RTX Pro 6000?

Yes. SIE includes a safe default profile that runs on most GPUs without speculative decoding, plus tuned profiles for H100 and H200 data center cards in addition to the RTX Pro 6000 256K profile.

### Do I need to configure speculative decoding manually?

No. Speculative decoding is built into specific hardware profiles. Selecting a profile like the RTX Pro 6000 256K configuration turns it on automatically rather than requiring manual setup.

### How long does it take to install SIE?

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

The core installation involves two commands, setting up the environment and running the install, and takes about a minute on a fresh Ubuntu system.

### Can I use SIE without owning a GPU?

Yes. SIE Cloud offers the same engine and models fully hosted, with free inference grants available, so you can test models like Qwen 3.8 without local hardware.
