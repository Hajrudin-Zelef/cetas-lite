---
id: collect-240926-huggingface/huggingface/unsloth-kimi-k3-gguf-hugging-face-1
title: "Read our How to Run Kimi K3 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Moonshot", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "context window", "fable 5", "gguf", "glm"]
source: docs/RAG/clean_en/huggingface/unsloth-kimi-k3-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 99]
sha256: 4f076427d93e23dae6025bd9416b03fe38337bc7dbd743e0be7383747f0844d8
---

# Read our How to Run Kimi K3 Guide!

<!-- source: https://huggingface.co/unsloth/Kimi-K3-GGUF -->

# Read our How to Run Kimi K3 Guide!

    *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

- To run Kimi K3, use our llama.cpp PR fork or run in Unsloth Studio
- Full precision lossless = Q8 (UD-Q8_K_XL), which is 50GB bigger than Q4 (UD-Q4_K_XL).
- Has vision support!
- Kimi K3 has toggles for High and Max thinking in Unsloth Studio
- Read our Kimi K3 guide for analysis and instructions.

Kimi K3 is an open-weight, native multimodal agentic model and our most capable model to date. It is a 2.8T-parameter model built on Kimi Delta Attention (KDA) and Attention Residuals (AttnRes), with native vision capabilities and a 1-million-token context window. It is the world's first open 3T-class model, designed for frontier intelligence across long-horizon coding, knowledge work, and reasoning.

- **New Architecture** : Kimi K3 is built on Kimi Delta Attention (KDA) and Attention Residuals (AttnRes), and scales up MoE sparsity with a Stable LatentMoE framework that activates 16 out of 896 experts — yielding an approximate 2.5× improvement in overall scaling efficiency over Kimi K2.
- **Long-Horizon Coding** : Operating with minimal human oversight, Kimi K3 sustains long engineering sessions, navigates massive repositories, and orchestrates terminal tools — from GPU kernel optimization and compiler development to vision-in-the-loop game dev, CAD, and even chip design.
- **Agentic Knowledge Work** : Kimi K3 advances end-to-end knowledge work, producing deep research with interactive visualizations, widgets and dashboards, and motion design and video editing, powered by its native multimodal architecture.
- **Native Multimodality & Long Context** : Kimi K3 understands text, images, and video within the same model, and supports a 1-million-token context window.
- **Open Frontier Weights** : We release the full Kimi K3 model weights under the Kimi K3 License, making frontier intelligence openly available for research, deployment, and further innovation.

| **Architecture** | Mixture-of-Experts (MoE) | 
| **Total Parameters** | 2.8T | 
| **Activated Parameters** | 104B | 
| **Number of Layers** | 93 | 
| **Number of Dense Layers** | 1 | 
| **Attention-Layer Composition** | 69 KDA + 24 Gated MLA | 
| **Attention Hidden Dimension** | 7168 | 
| **Number of Attention Heads** | 96 | 
| **Latent MoE Dimension** | 3584 | 
| **MoE Hidden Dimension** (per Expert) | 3072 | 
| **Number of Experts** | 896 | 
| **Selected Experts per Token** | 16 | 
| **Number of Shared Experts** | 2 | 
| **Vocabulary Size** | 160K | 
| **Context Length** | 1048576 | 
| **Attention Mechanism** | KDA & Gated MLA | 
| **Activation Function** | SiTU-GLU | 
| **Vision Encoder** | MoonViT-V2 | 
| **Parameters of Vision Encoder** | 401M | 
| **Quantization** | MXFP4 weights / MXFP8 activations (quantization-aware training) | 
| **Modality** | Text, Image | 

| Benchmark | <sup>Kimi K3 <sup>(max)</sup></sup> | <sup>Claude Fable 5 <sup>(max, w/ fallback)</sup></sup> | <sup>GPT-5.6 Sol <sup>(max)</sup></sup> | <sup>Claude Opus 4.8 <sup>(max)</sup></sup> | <sup>GPT-5.5 <sup>(xhigh)</sup></sup> | <sup>GLM-5.2 <sup>(max)</sup></sup> | 
|---|---|---|---|---|---|---|
| **Reasoning & Knowledge** |  |  |  |  |  |  | 
| GPQA Diamond | 93.5 | 92.6 | 94.1 | 91.0 | 93.5 | 91.2 | 
| CritPt | 23.4 | 28.6 | 32.3 | 20.9 | 27.1 | 20.9 | 
| AA-LCR | 74.7 | 70.0 | 73.7 | 67.7 | 74.3 | 71.3 | 
| HLE-Full | 43.5 / 56.0 | 53.3 / 63.0 | 44.5 / 58.0 | 49.8 / 57.9 | 41.4 / 52.2 | — | 
| **Coding** |  |  |  |  |  |  | 
| DeepSWE | 67.5 | 70.0 | 73.0 | 59.0 | 67.0 | 46.2 | 
| ProgramBench | 77.8 | 76.8 | 77.6 | 71.9 | 70.8 | 63.7 | 
| Terminal-Bench 2.1 | 88.3 | 88.0 | 88.8 | 84.6 | 83.4 | 82.7 | 
| FrontierSWE | 81.2 | 86.6 | 71.3 | 66.7 | 64.9 | 67.3 | 
| SWE-Marathon | 42.0 | 35.0 | 39.0 | 40.0 | 14.0 | 13.0 | 
| PostTrainBench | 36.6 | 41.4 | 34.6 | 34.1 | 28.4 | 34.3 | 
| MLS-Bench-Lite | 48.3 | 49.9 | 46.2 | 42.8 | 35.5 | 40.4 | 
| SciCode | 58.7 | 60.2 | 56.1 | 53.5 | 56.1 | 50.5 | 
| Kimi Code Bench 2.0 | 72.9 | 76.9 | 64.8 | 71.7 | 69.0 | 64.2 | 
| **Agentic** |  |  |  |  |  |  | 
| BrowseComp | 91.2 | 88.0 | 90.4 | 84.3 | 84.4 | — | 
| DeepSearchQA (F1) | 95.0 | 94.2 | — | 93.1 | — | — | 
| ResearchRubrics | 76.2 | — | 73.8 | 73.5 | 64.0 | 71.1 | 
| GDPval-AA v2 (Elo) | 1686 | 1747 | 1736 | 1593 | 1491 | 1510 | 
| Toolathlon-Verified | 76.5 | 77.9 | 74.9 | 76.2 | 73.5 | 59.9 | 
| MCPMark-Verified | 94.5 | 87.4 | 92.9 | 76.4 | 92.9 | — | 
| MCP-Atlas | 84.2 | 84.7 | 83.6 | 83.6 | 82.8 | 82.6 | 
| AutomationBench | 30.8 | 29.1 | 29.7 | 27.2 | 22.7 | 12.9 | 
| JobBench | 54.3 | 57.4 | 45.4 | 48.4 | 38.3 | 43.4 | 
| AA-Briefcase (Elo) | 1548 | 1583 | 1495 | 1354 | 1158 | 1260 | 
| Agents' Last Exam | 28.3 | 25.7 <sup>†</sup> | 29.6 | 27.0 | 26.6 | 20.4 | 
| APEX-Agents | 41.0 | 43.3 | 39.9 | 39.4 | 38.5 | 35.6 | 
| OfficeQA Pro | 63.3 | 69.9 | 63.2 | 63.9 | 60.9 | 41.4 | 
| SpreadsheetBench 2 | 34.8 | 34.7 | 32.4 | 31.6 | 29.1 | 28.1 | 
| OSWorld-Verified | 84.8 | 85.0 | 83.0 | 83.4 | 79.0 | — | 
| OSWorld 2.0 | 58.3 | 66.1 | 62.6 | 55.7 | 49.5 | — | 
| SaaS-Bench | 60.1 | — | 61.4 | 56.1 | 43.8 | — | 
| τ³-Banking | 33.4 | 26.8 | 33.0 | 27.6 | 31.3 | 26.8 | 
| Harvey Lab-AA | 94.6 | 93.6 | 87.2 | 91.1 | 86.3 | 91.0 | 
| CorpFin v2 | 71.6 | 71.8 | 64.4 | 66.7 | 68.4 | 66.1 | 
| Finance Agent v2 | 54.4 | 56.3 | 53.8 | 53.9 | 51.8 | 49.7 | 
| Legal Research Bench | 44.2 | 49.5 | 48.1 | 43.8 | 40.4 | 31.3 | 
| **Vision** |  |  |  |  |  |  | 
| WorldVQA ForceAnswer | 51.0 | 56.7 | 41.8 | 39.1 | 38.5 | — | 
| OmniDocBench | 91.1 | 89.8 | 85.8 | 87.9 | 89.4 | — | 
| PerceptionBench | 58.5 | 57.2 | 59.7 | 47.2 | 55.8 | — | 
| Video-MME (w. sub) | 90.0 | — | 89.5 | 86.0 | 89.3 | — | 
| MMVU | 82.1 | — | 81.2 | 79.2 | 81.7 | — | 
| BabyVision w/ python | 85.7 | 90.5 | 88.9 | 81.2 | 83.6 | — | 
| MMMU-Pro | 81.6 / 83.4 | 81.2 / 86.5 | 83.0 / 84.6 | 78.9 / 82.7 | 81.2 / 83.2 | — | 
| CharXiv (RQ) | 84.8 / 91.3 | 88.9 / 93.5 | 84.6 / 89.1 | 80.5 / 89.9 | 84.1 / 89.0 | — | 
| MathVision | 94.3 / 97.8 | 94.8 / 98.6 | 95.8 / 97.8 | 86.7 / 97.1 | 92.2 / 96.8 | — | 
| ZeroBench (pass@5) | 23.0 / 41.0 | 23.0 / 46.0 | 17.0 / 35.0 | 17.0 / 34.0 | 22.0 / 41.0 | — | 

## **Footnotes**

All Kimi K3 results are obtained with reasoning effort set to 'max' and temperature = 1.0. For single-step tasks, such as GPQA Diamond, HLE-Full, and vision benchmarks without tools, we set top-p = 0.95; for agentic tasks, we set top-p = 1.0. For HLE-Full, MMMU-Pro, CharXiv (RQ), MathVision, and ZeroBench, each cell reports the scores without and with tool augmentation (general tools for HLE-Full, Python for the vision benchmarks), in that order.

