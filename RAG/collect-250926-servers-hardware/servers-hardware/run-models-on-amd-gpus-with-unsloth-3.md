---
id: collect-250926-servers-hardware/servers-hardware/run-models-on-amd-gpus-with-unsloth-3
title: "Train & run models on AMD GPUs with Unsloth"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Unsloth"]
dates: []
keywords: ["amd", "gpu", "gpus", "agent", "agents", "inference", "llama", "llama.cpp", "memory", "qlora", "safeguards", "training"]
source: docs/RAG/clean4/run models on AMD GPUs with Unsloth.md
source_anchor: ""
source_lines: [252, 282]
sha256: 5651d68ff885bbe64422c68d53b5bdb9519476a957b4add6a776ec848b3e191a
---

# Train & run models on AMD GPUs with Unsloth

### :question:How AMD support was built for Unsloth
Unsloth’s custom kernels were ported to AMD in close collaboration with the AMD team. This required changes across the installer, hardware detection, runtime routing, monitoring, and prebuilt asset selection.
* **ROCm kernels:** HIP/Triton ports tuned for RDNA and CDNA, with architecture-specific precision, performance, and stability fixes. ([#2520](https://github.com/unslothai/unsloth/pull/2520))
* **4-bit training:** bitsandbytes-based QLoRA support for Radeon and CDNA GPUs. ([#3748](https://github.com/unslothai/unsloth/pull/3748))
* **Automatic setup:** robust AMD GPU detection and installation of the correct ROCm PyTorch wheel, including repair and platform safeguards. ([#4770](https://github.com/unslothai/unsloth/pull/4770))
* **Windows and Strix Halo:** native Windows ROCm installation, runtime compatibility patches, unified-memory reporting, and WSL support. ([#5301](https://github.com/unslothai/unsloth/pull/5301), [#6227](https://github.com/unslothai/unsloth/pull/6227))
* **ROCm inference:** architecture-specific llama.cpp prebuilts, source-build fallback, AMD VRAM detection, and monitoring. ([#5172](https://github.com/unslothai/unsloth/pull/5172))
### :clock1: Future work
Thanks so much to the AMD team for collaborating with us on making AMD work well! Upcoming priorities include:
* Continued development and engagement with the AMD team.
* Expanded hardware-backend CI/CD using AMD Strix Halo and Radeon Lab hardware.
* Support for new AMD GPU cards and architectures as they launch.
* Clearer multi-GPU guidance: AMD distributed training is currently Linux-only. On Linux, ROCm uses RCCL, so distributed training works. On Windows, AMD ROCm does not yet provide a GPU-collective backend, with GLOO/CPU-only support as of [TheRock #5694](https://github.com/ROCm/TheRock/pull/5694), use Linux if you plan to do AMD multi-GPU training.
* Further speed optimizations with respect to the ROCm stack.
### :love\_letter: Thanks to AMD for collaborating!
Thanks so much to the AMD team and our amazing Unsloth community for collaborating with us on this release. We thank Strahinja Stamenkovic, Filip Jankovic, Iswarya Alex, Yue Yuan Bill He, Eda Zhou, Mahdi Ghodsi, Guruprasad and the entire AMD team for collaborating on making AMD support work great with Unsloth! Also huge thanks to community members: Nate, UBER6, AiwendilOfMirk..., Gene, Jimster480, VELICAN, Vintheboy, archmaker, BichonFriseMax, Calandracas, Chigoma333, Edd, electroglyph, jslowbell, mk 27B UD-, Satyam, snapcast3r, TotoMC for testing and debugging with us.
An easy way to use Unsloth on AMD is through [AMD Developer Cloud](https://www.amd.com/en/developer/resources/cloud-access/amd-developer-cloud.html), with one-click notebooks powered by MI300X GPUs with 192GB VRAM. AMD also offers free credits through the [AMD AI Developer Program](https://www.amd.com/en/developer/ai-dev-program.html). To run an Unsloth notebook, use any AMD notebook from [unslothai/notebooks](https://github.com/unslothai/notebooks) and swap the GitHub domain with the AMD Developer Cloud URL, since the notebook paths are the same.
Need help or want to share feedback? You can join our [Discord](https://discord.com/invite/unsloth), read the [AMD docs](https://unsloth.ai/docs/get-started/install/amd), open an issue or PR on [GitHub](https://github.com/unslothai/unsloth/issues), or post on our [Reddit r/unsloth](https://www.reddit.com/r/unsloth/).
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/amd.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
