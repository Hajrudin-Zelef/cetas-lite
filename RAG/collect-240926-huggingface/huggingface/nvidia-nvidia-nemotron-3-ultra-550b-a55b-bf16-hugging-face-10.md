---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-10
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Moonshot", "Nvidia", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "benchmark", "benchmarks", "blackwell", "deepseek", "distribution", "glm", "guardrails", "kimi", "leaderboard", "llama"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [769, 869]
sha256: 1db5d71ad9ae261f6a90e492baa5c47c5628388e3b122c9962225dcec4f0e874
---

# Set the IP for the head node in RAY_HEAD_IP

| Synthetic Science Diversity MCQ from GPT-OSS and Kimi-K2 | Text | 532,942 | [doubtnut]; [Pile-FreeLaw]; [Llama Nemotron Dataset]; [askfilo]; [EssentialAI/essential-web-v1.0]; [Vedantu]; [auxiliary_train]; [cdquestions.com]; [AMC 8 Problems and Solutions, AMC 10 Problems and Solution, and AIME Problems and Solutions]; [AAPT]; [ICHO-IPH0 Dataset]; [LIMO dataset (Less is More for Reasoning)] | [GPT-OSS]; [Kimi-K2] | 
| Synthetic Science Diversity OpenQ from GPT-OSS and Kimi-K2 | Text | 131,045 | [doubtnut]; [Pile-FreeLaw]; [Llama Nemotron Dataset]; [askfilo]; [EssentialAI/essential-web-v1.0]; [Vedantu]; [auxiliary_train]; [cdquestions.com]; [AMC 8 Problems and Solutions, AMC 10 Problems and Solution, and AIME Problems and Solutions]; [AAPT]; [ICHO-IPH0 Dataset]; [LIMO dataset (Less is More for Reasoning)] | [GPT-OSS]; [Kimi-K2] | 
| Synthetic Science Reasoning No-Tool from GPT-OSS and Kimi-K2 | Text | 2,085,600 | [doubtnut]; [Pile-FreeLaw]; [Llama Nemotron Dataset]; [askfilo]; [EssentialAI/essential-web-v1.0]; [Vedantu]; [auxiliary_train]; [cdquestions.com]; [AMC 8 Problems and Solutions, AMC 10 Problems and Solution, and AIME Problems and Solutions]; [AAPT]; [ICHO-IPH0 Dataset]; [LIMO dataset (Less is More for Reasoning)] | [GPT-OSS]; [Kimi-K2] | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | 62,333 | [Long-context SFT data: lc_nothink 256k] | [Qwen/Qwen3-235B-A22B-Thinking-2507 and deepseek-ai/DeepSeek-R1] | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | 49,698 | [Long-context SFT data: MRCR 200k] | [Qwen/Qwen3-235B-A22B-Thinking-2507 and deepseek-ai/DeepSeek-R1] | 
| Synthetic Text-To-SQL | Text | 96,564 | [Undisclosed - no seed data listed] | [gpt-oss-120b] | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | 397,538 | [Long-context SFT data: RULER 256k] | [Qwen/Qwen3-235B-A22B-Thinking-2507 and deepseek-ai/DeepSeek-R1] | 
| Synthetic SWE Unverified | Text | 27,960 | [NVAgenticCLIPrompts-v1]; [NVAgenticSkills-v1] | [gpt-oss-120b]; [Qwen/Qwen3-Coder-480B-A35B-Instruct]; [GLM-4.7-Flash] | 
| Synthetic SWE Unverified | Text | 24,632 | [NVAgenticCLIPrompts-v1]; [NVAgenticSkills-v1] | [gpt-oss-120b]; [Qwen/Qwen3-Coder-480B-A35B-Instruct]; [GLM-4.7-Flash] | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | 49,902 | [Long-context SFT data] | [Qwen/Qwen3-235B-A22B-Thinking-2507 and deepseek-ai/DeepSeek-R1] | 
| Synthetic Tool Call Schema for RL | Text | 469,983 | [UltraTool]; [ToolEyes]; [AutoTools]; [API-Bank]; [Nemotron-Personas-USA]; [Salesforce xLAM function-calling]; [Glaive function-calling-v2]; [Agent-Ark/Toucan-1.5M] | [DeepSeek-V3.2]; [GLM-4.6]; [gpt-oss-120b]; [Kimi-K2-Instruct] | 
| Synthetic Tool Call Schema for RL | Text | 707,967 | [UltraTool]; [ToolEyes]; [AutoTools]; [API-Bank]; [Nemotron-Personas-USA]; [Salesforce xLAM function-calling]; [Glaive function-calling-v2]; [Agent-Ark/Toucan-1.5M] | [DeepSeek-V3.2]; [GLM-4.6]; [gpt-oss-120b]; [Kimi-K2-Instruct] | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | 52,630 | [AALCR seed blend: SEC Filings]; [CC]; [Wikipedia]; [FinePDFs]; [ArXiv]; [Pile-NIH ExPorter]; [BioRxiv]; [PMC Article]; [USPTO Backgrounds]; [peS20]; [Global Regulations]; [CORE]; [Gutenberg (PG-19)]; [DOAB CC-BY]; [NDLTD]; [Amps]; [StackExchange]; [MathPile]; [Numinas] | [Qwen3-30B-A3B] | 
| Synthetic Safety from gemma-3-4b-it, Nemotron-Nano-9B-v2, and gpt-oss-120b | Text | 44,091 | [Safety SFT Data] | [google/gemma-3-4b-it]; [Nemotron-Nano-9B-v2]; [gpt-oss-120b] | 

For our post-training recipe, we focused on the following languages in addition to English: French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, and Chinese.

Those languages were represented in the form of multilingual reasoning and translation tasks.

The following table depicts our sample distribution.

| Language | Size | 
|---|---|
| English | 8.6M | 
| Italian | 138k | 
| German | 138k | 
| Spanish | 138k | 
| French | 138k | 
| Japanese | 138k | 
| Chinese | 138k | 
| Hindi | 138k | 
| Korean | 138k | 
| Brazilian Portuguese | 138k | 

**Data Collection Method by dataset** 

- Hybrid: Automated, Human, Synthetic

**Labeling Method by dataset** 

- Hybrid: Automated, Human, Synthetic

**Properties:** This corpus comprises a mix of high-quality standard benchmarks and test suites for modern agentic AI as outlined in the benchmark section of the model card.

**Data Collection Method by dataset** 

- Hybrid: Automated, Human, Synthetic

**Labeling Method by dataset** 

- Hybrid: Automated, Human, Synthetic

**Properties:** This corpus comprises a mix of high-quality standard benchmarks and test suites for modern agentic AI as outlined in the benchmark section of the model card.

- **Acceleration Engine:** PyTorch
- **Test Hardware:**
  - NVIDIA Hopper
    - H100
    - H200
  - NVIDIA Grace Blackwell
    - GB200
    - GB300
  - NVIDIA Blackwell
    - B200
    - B300
- NVIDIA Hopper

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case. For more details: Safety and Explainability Subcards.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, and Privacy Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

```
@misc{nvidia_nemotron_3_ultra_2026,
  title  = {Nemotron 3 Ultra: Open, Efficient Mixture-of-Experts Hybrid Mamba-Transformer Model for Agentic Reasoning},
  author = {{NVIDIA}},
  year   = {2026},
  url    = {https://research.nvidia.com/labs/nemotron/files/NVIDIA-Nemotron-3-Ultra-Technical-Report.pdf},
  note   = {White Paper}
}
```
- Downloads last month
- 171,251

## Spaces using nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 4

## Collections including nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16

## Article mentioning nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  87<sup>*</sup>
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results    leaderboard  71.9
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  86.8
- harborframework/terminal-bench-2.1 · Terminalbench 2 1 View evaluation results    leaderboard  56.4<sup>*</sup>
- SWE-bench/SWE-bench_Multilingual · Swe Bench Multilingual Resolved View evaluation results    leaderboard  67.7
- cais/hle · Hle
- No tools View evaluation results26.7<sup>*</sup>
- With tools View evaluation results37.4<sup>*</sup>
