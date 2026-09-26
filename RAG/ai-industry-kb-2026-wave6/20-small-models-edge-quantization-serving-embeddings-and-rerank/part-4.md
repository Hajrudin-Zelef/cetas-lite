---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-4
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 4)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Intel", "Microsoft", "Mistral", "Qualcomm"]
dates: []
keywords: ["amd", "benchmarks", "copilot", "inference", "intel", "license", "llama", "memory", "mistral", "multimodal", "qwen"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9720, 9726]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: 3f62b139657f5637ed7e4ab3e5310aaeb7a0450ad9595aaec7039a3136337174
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 4)

- Copilot+ PC hardware floor (Microsoft-defined): NPU capable of 40+ TOPS, high-performance processor (Qualcomm Snapdragon X Elite/Plus, Intel Core Ultra, or AMD Ryzen AI), 16GB RAM, 256GB SSD, Windows 11 24H2+, Pluton security processor with TPM 2.0 [VENDOR/SECONDARY]. Sources: https://news.microsoft.com/source/features/ai/how-the-npu-is-paving-the-way-toward-a-more-intelligent-windows/ and https://www.techtarget.com/enterprise-software/feature/Are-Copilot-PCs-worth-it-Evaluating-AI-powered-Windows-PCs
- Copilot+ exclusive experiences: Recall (photographic-memory device search), Cocreator (near-real-time on-device image generation), Live Captions translating 40+ languages to English in real time [VENDOR]. Sources: https://news.microsoft.com/source/features/ai/how-the-npu-is-paving-the-way-toward-a-more-intelligent-windows/ and https://www.nasdaq.com/articles/microsoft-msft-debuts-copilot-pcs-with-ai-focused-features
- Copilot+ efficiency claims: up to 22 hours video playback / 15 hours web browsing; up to 20x more powerful and 100x more efficient on AI workloads versus 12th-gen Intel Core i7 baseline [VENDOR]. Source: https://www.nasdaq.com/articles/microsoft-msft-debuts-copilot-pcs-with-ai-focused-features
- 2026 AI-PC reality check (community): chips at 11–34 TOPS are AI-branded but not AI-capable — Windows Recall requires 40 TOPS, on-device Phi-3 inference 40+, real-time AI video enhancement 45+ [COMMUNITY]. Source: https://medium.com/@svnkrmkr/ai-pc-npu-dashboard-check-your-copilot-tops-rating-2026-452e0a89660c
- Llama 3.2 small line reference: 11B multimodal at 128K under Llama 3 license appears in a 2026 small-model comparison table [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
- The same 2026 small-model table gives: Phi-4 14B MMLU 84.8% / HumanEval 82.6% / MATH 80.4% / 16K / MIT; Gemma 3 12B MMLU 83.2% / HumanEval 78.4% / MATH 75.1% / 128K; Mistral Small 3.1 24B MMLU 82.7%; Llama 3.2 11B MMLU 80.1%; Qwen 2.5 14B MMLU 82.9% — note these are 2025-era figures, do not treat as 2026 re-benchmarks [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf

