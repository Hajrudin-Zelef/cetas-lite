---
id: collect-261001-ia-llm/ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-1
title: "Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia"]
dates: []
keywords: ["agents", "multimodal", "nvidia", "omni", "agentic", "alignment", "attention", "benchmark", "benchmarks", "cost", "embedding", "fp8"]
source: docs/RAG/collect-261001-ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-audio-and.md
source_anchor: ""
source_lines: [1, 65]
sha256: 1dc369013a108064c4c66c76b8638c474b9bbaf9531e3f32a25d765e7901e233
---

# Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents

- **NVIDIA Nemotron 3 Nano Omni** is a new omni-modal understanding model built for**real-world document analysis, multiple image reasoning, automatic speech recognition, long audio-video understanding, agentic computer use, and general reasoning** .
- It extends the Nemotron multimodal line from a strong vision-language system to a broader **text + image + video + audio** model.
- Nemotron 3 Nano Omni delivers **best-in-class accuracy** on complex document intelligence leaderboards such as MMlongbench-Doc, OCRBenchV2, while also leading in video and audio leaderboards like WorldSense and DailyOmni. It achieves top accuracy on VoiceBench for audio understanding and ranks as the most cost‑efficient open video understanding model on MediaPerf.
- Under the hood, it combines the **Nemotron 3 hybrid Mamba-Transformer Mixture-of-Experts backbone** with a**C-RADIOv4-H** vision encoder and**Parakeet-TDT-0.6B-v2** audio encoder.
- The architecture is designed to preserve fine visual detail, add native audio understanding, and scale to **very long multimodal contexts** for dense images, documents, videos, and mixed-modality reasoning.
- The training recipe uses **staged multimodal alignment and context extension** , followed by**preference optimization and multimodal reinforcement learning** .
- Nemotron 3 Nano Omni delivers up to 9x higher throughput and 2.9x the single-stream reasoning speed on multimodal use-cases, compared to alternatives.
- Download the BF16, FP8 and NVFP4 checkpoints at HuggingFace.
- For more information about the model architecture, training recipe, data pipelines and benchmarks, read the full Nemotron 3 Nano Omni report.

**Benchmark highlights**

Building on Nemotron Nano V2 VL, Nemotron 3 Nano Omni delivers substantial visual gains and adds entirely new audio and video+audio capabilities - while also leading another open-weights omni model, Qwen3-Omni, in many domains.

| Task | Benchmark | Nemotron 3 Nano Omni | Nemotron Nano V2 VL | Qwen3-Omni 30B-A3B | 
|---|---|---|---|---|
| Document understanding | OCRBenchV2-En | **65.8** | 61.2 | - | 
|  | MMLongBench-Doc | **57.5** | 38.0 | 49.5 | 
|  | CharXiv reasoning | **63.6** | 41.3 | 61.1 | 
| GUI | ScreenSpot-Pro | 57.8 | 5.5 | **59.7** | 
|  | OSWorld | **47.4** | 11.0 | 29.0 | 
| Video understanding | Video-MME | **72.2** | 63.0 | 70.5 | 
| Video + Audio understanding | WorldSense | **55.4** | - | 54.0 | 
|  | DailyOmni | **74.1** | - | 73.6 | 
| Voice interaction | VoiceBench | **89.4** | - | 88.8 | 
| ASR | HF Open ASR (lower is better) | **5.95** | - | 6.55 | 

**Efficiency highlights**

Compared to other open omni models with the same interactivity, Nemotron 3 Nano Omni delivers 7.4x higher system efficiency for multi-document use cases and 9.2x higher system efficiency for video use cases

*Figure 1. Total system throughput for multi-document and video use cases sustained by each model at a fixed per‑user interactivity threshold (tokens/sec/user)* 

At a high level, Nemotron 3 Nano Omni is aimed at five classes of workloads:

This is not only about OCR. The model is positioned for long, messy, high-value documents where understanding depends on layout, tables, figures, formulas, section structure, and cross-page references. Think contracts, technical papers, reports, manuals, multi-page forms, or compliance packets. The model can handle 100+ page documents.

Nemotron 3 Nano Omni includes strong speech understanding capabilities that enable high-quality transcription across diverse audio conditions. It handles long-form audio with varying speakers, accents, and background noise. These capabilities can be integrated into broader workflows, allowing spoken content to be transcribed, analyzed, and combined with other modalities for tasks like summarization, question answering, and cross-modal reasoning.

Many enterprise and developer workflows depend on mixed audio and visual evidence: screen recordings with narration, training videos, meetings with slides, tutorials, product demos, customer support captures, and long-form video archives. Nemotron 3 Nano Omni is built to reason over those inputs jointly.

The Nemotron 3 Nano Omni model is specifically trained for agentic computer use, enabling it to assist with tasks in graphical user interface (GUI) environments. Its capabilities include interpreting screenshots, monitoring the state of the user interface, grounding its reasoning in on-screen visuals, and helping with action selection or workflow automation.

The model is designed for more than perception. It excels at reasoning-intensive tasks that require synthesizing information across long context windows, multiple modalities, and structured or semi-structured evidence. It can carry out multi-step reasoning, perform calculations, and connect signals from text, images, tables, and other inputs to arrive at coherent, well-supported answers.

Nemotron 3 Nano Omni uses a unified **encoder-projector-decoder** design. The language backbone is Nemotron 3 Nano 30B-A3B, paired with the C-RADIOv4-H vision encoder and the Parakeet-TDT-0.6B-v2 audio encoder. The modality-specific encoders connect into the LLM backbone through lightweight projectors. 

Figure 2. Model architecture of NVIDIA Nemotron 3 Nano Omni 30B-A3B

The model backbone interleaves three key components: **23 Mamba selective state-space layers** for efficient long-context processing;  **23 MoE layers** with **128 experts, top-6 routing**, and a **shared expert** for conditional capacity; and **6 grouped-query attention layers** to preserve strong global interaction and expressivity.

Nemotron 3 Nano Omni combines state-space models, attention, and MoE in a unified design that maintains strong reasoning performance while remaining practical for long, multimodal contexts.

On the vision side, the Nemotron 3 Nano Omni replaces the tiling strategy used in the v2 model with **dynamic resolution processing at native aspect ratio**. Each image can be represented using a variable number of 16 x 16 patches, with **a minimum of 1,024 to a maximum of 13,312 visual patches per image**. For square images, this is equivalent to 512 x 512 and 1840 x 1840, respectively.

That flexibility is critical for handling high-resolution, complex visual inputs such as OCR-heavy documents, financial tables, slides, research figures, screenshots, and GUI layouts—especially when both fine details and overall structure need to be understood together.

For video, Nemotron 3 Nano Omni uses a dedicated **Conv3D tubelet embedding** path. Instead of embedding each frame independently, every pair of consecutive frames is fused into a single "tubelet" before the ViT, halving the number of vision tokens the language model has to attend to. This allows us to either double the number of frames with the same token budget, or halve the number of tokens with the same number of frames

EVS is an important feature, enabled during inference time, that drops redundant video tokens after the vision encoder. This reduces latency and improves throughput while maintaining accuracy. The first frame of the video is kept entirely, then for each subsequent frame, EVS keeps the “dynamic” tokens where the video is changing and drops the “static” ones where nothing has changed from the previous frame. We combine this with Conv3D to enable superior compression: Conv3D fuses tokens from pairs of frames into one, and then EVS prunes redundant static information.

The audio side is powered by **Parakeet-TDT-0.6B-v2**, connected to the backbone through its own 2-layer MLP projector. Audio is sampled at **16 kHz**, and the model is trained with inputs up to **1,200 seconds (20 minutes)**, while the LLM max context length supports 5+ hours. 

