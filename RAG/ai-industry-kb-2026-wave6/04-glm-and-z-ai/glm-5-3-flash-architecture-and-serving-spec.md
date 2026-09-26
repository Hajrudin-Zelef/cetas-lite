---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/glm-5-3-flash-architecture-and-serving-spec
title: "GLM-5.3-Flash — architecture and serving spec"
domain: glm-and-z-ai
role: deep-dive
task: actor-profile
actors: ["DeepSeek", "Hugging Face", "Moonshot", "Nvidia", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-08-14", "2026-08-26", "2026-08-28"]
keywords: ["glm", "agent", "attention", "benchmark", "context window", "cyber", "cybersecurity", "deepseek", "distribution", "exploit", "fp8", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1765, 1791]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: 79f6949b0df75c3abfc8300afb28f75b0642d689c7fbbad46f266f5cc4b110a0
---

# GLM-5.3-Flash — architecture and serving spec

### GLM-5.3-Flash — architecture and serving spec
- 320B total parameters / **18B active per token** (A18B); 1,048,576-token context window (exactly 2^20) [SECONDARY] (MarkTechPost, 2026-08-26; startupfortune.com; RadixArk model card).
- **45 decoder layers: 3 dense MLP layers + 42 MoE layers; 288 routed experts per MoE layer; one shared expert** [SECONDARY] (RadixArk/GLM-5.3-Flash-NVFP4 model card, describing the upstream BF16 checkpoint).
- Hybrid attention stack: **34 Kimi Delta Attention (KDA) linear-attention layers + 11 KPool-indexed DeepSeek Sparse Attention (DSA) layers**, per NVIDIA NeMo AutoModel documentation; only the sparse layers carry a growing KV cache — the 34 KDA layers use a fixed-size recurrent state, which is what makes long context cheaper to serve [SECONDARY] (startupfortune.com, citing NeMo AutoModel docs).
- Architecture additionally includes **MLA, manifold-constrained hyper-connections, a vision encoder, and a native MTP/NextN draft layer** [SECONDARY] (RadixArk model card reading of the upstream checkpoint).
- Routing top-k count for Flash was not disclosed in the sources reviewed; do not quote one [DIRECTIONAL].
- **Pretraining: ~30T multimodal tokens** [SECONDARY] (model-guide research note, github.com/mattrobenolt pi-configs; startupfortune.com context — treat as [UNVERIFIED] pending the vendor model card, single-source chain).
- Deployment support listed on the model card: **vLLM, SGLang, Transformers, Unsloth, KTransformers** [SECONDARY] (startupfortune.com; MarkTechPost, 2026-08-26).
- Self-hosting floor: default **FP8 checkpoint ≈ 306 GiB of weights before KV cache**; current vLLM path supports **NVIDIA Hopper and newer only**; realistic minimum is an **8-GPU node (or a GB200 tray at TP4)**; everyone below that consumes Flash as an API [SECONDARY] (MarkTechPost, 2026-08-26 — single source, mark [UNVERIFIED] for procurement decisions).
- Community quantization: **RadixArk/GLM-5.3-Flash-NVFP4** released on Hugging Face **August 28, 2026**, pinned to source revision `a6c167b6` of `zai-org/GLM-5.3-Flash-BF16`; converted to mixed-precision **NVFP4 W4A4 using NVIDIA Model Optimizer** with plain abs-max scaling and 256 tensor-scale normalization; card declares **MIT license** and **global deployment geography** [COMMUNITY] (RadixArk model card). A matching RadixArk NVFP4 quant of the GLM-5.3 flagship exists at `RadixArk/GLM-5.3-NVFP4` [COMMUNITY] (Hugging Face listing).

### GLM-5.3 flagship — staged weights release and custom license
- At the **August 14, 2026** launch, GLM-5.3 was **not yet downloadable**: weights and license terms were both unpublished, so "open" described a roadmap, not a file [SECONDARY] (emergent.sh guide, 2026-08).
- Z.ai said it would delay the public weights release by **about two weeks** for additional security checks and strengthened safeguards; the most sensitive cybersecurity functions were initially restricted to a **"trusted access" program** for selected partners and verified users [SECONDARY] (kr-asia.com; deeplearning.ai/The Batch, 2026-08).
- As of **August 28, 2026** the flagship's weights remained unpublished — the two-week window closed with the MIT-licensed Flash shipping instead [SECONDARY] (ainvest.com, 2026-08; emergent.sh).
- When the flagship weights did release, they came under a **new custom license (not MIT)**: companies with **more than $10 billion in annual revenue over 12 months must pass a Z.ai security review** before using the model or derivative works commercially [SECONDARY] (Techmeme summary of Frederic Lardinois/The New Stack, 2026-08-28; techbooky.com citing the same reporting). GLM-5.2's MIT license does not carry this gate [SECONDARY] (deeplearning.ai).
- The revenue gate follows the Llama-style pattern of keeping the largest rivals from free commercial advantage while preserving ecosystem distribution [DIRECTIONAL] (techbooky.com analysis).

### GLM-5.3 flagship — serving spec and cyber capability framing
- Text in up to **1M tokens**; text out up to **128,000 tokens**; API throughput claimed at **90 tokens/second** [SECONDARY] (deeplearning.ai/The Batch, 2026-08).
- Architecture: MoE transformer, **753B total parameters / 40B active per token** — consistent with the 753B tensor-class figure in existing §4 [SECONDARY] (deeplearning.ai).
- Features: adjustable reasoning levels (**low / high / max**), tool calling, structured output, streaming, context caching [SECONDARY] (deeplearning.ai).
- Z.ai's framing: GLM-5.3 was **not trained as a dedicated cybersecurity model**; cyber capabilities are described as **"emergent"** from expanded reinforcement learning and training in longer, more varied software-development environments [VENDOR via secondary] (kr-asia.com, 2026-08).
- Z.ai acknowledged the model is currently **better at early vulnerability stages — code review, discovery, verification — than at deeper exploitation or complete offensive/defensive operations** [VENDOR via secondary] (kr-asia.com, 2026-08).
- Training recipe details reported: **single-rollout asynchronous optimization** (trains on attempts one at a time instead of full batches); long agent-attempt records are split into **compacted segments** so the model learns from long-running tasks, not only short ones [SECONDARY] (deeplearning.ai).
- On Artificial Analysis' index the model scored **60 points and effectively ties Kimi K3 as open-weights leader**; it took the **best score among all models on CyberGym** (exploit-detection benchmark) in Z.ai's tests [SECONDARY reporting vendor tests] (deeplearning.ai — vendor numbers, label [VENDOR] for the CyberGym claim).

