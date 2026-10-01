---
id: collect-261001-ia-llm/ia-llm/kimi-vendor-verifier
title: "kimi-vendor-verifier"
domain: ia-llm
role: reference
task: reference
actors: ["Moonshot", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["kimi", "agentic", "agents", "benchmark", "benchmarks", "cost", "gpu", "inference", "kv cache", "leaderboard", "multimodal", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/kimi-vendor-verifier.md
source_anchor: ""
source_lines: [1, 54]
sha256: a41bb2e086a12aacc61ba357bc276bc4c6fb4eb78658c4903006b55bca167790
---

# kimi-vendor-verifier

Alongside the release of the Kimi K2.6 model, we are open-sourcing the Kimi Vendor Verifier (KVV) project, designed to help users of open-source models verify the accuracy of their inference implementations.

Not as an afterthought, but because we learned the hard way that open-sourcing a model is only half the battle. The other half is ensuring it runs correctly everywhere else.

## Official Evaluation Results

| Benchmark | Metric | Temperature | TopP | MaxTokens | Kimi API | 
|---|---|---|---|---|---|
| OCRBench | acc | 1.0 | 0.95 | 16384 | 91.0 | 
| AIME2025 | avg@32 | 1.0 | 0.95 | 98304 | 98.4 | 
| MMMU Pro Vision | acc | 1.0 | 0.95 | 65536 | 78.8 | 

You can click here to access the Kimi API K2VV evaluation results for calculating the F1 score.

## Why We Built KVV

**From Isolated Incidents to Systemic Issues**

Since the release of K2 Thinking, we have received frequent feedback from the community regarding anomalies in benchmark scores. Our investigation confirmed that a significant portion of these cases stemmed from the misuse of Decoding parameters. To mitigate this immediately, we built our first line of defense at the API level: enforcing Temperature=1.0 and TopP=0.95 in Thinking mode, with mandatory validation that thinking content is correctly passed back.

However, more subtle anomalies soon triggered our alarm. In a specific evaluation on LiveBenchmark, we observed a stark contrast between third-party API and official API. After extensive testing of various infrastructure providers, we found this difference is widespread.

This exposed a deeper problem in the open-source model ecosystem: The more open the weights are, and the more diverse the deployment channels become, the less controllable the quality becomes.

If users cannot distinguish between "model capability defects" and "engineering implementation deviations," trust in the open-source ecosystem will inevitably collapse.

## Our Solution

**Six Critical Benchmarks** (selected to expose specific infra failures):

1. Pre-Verification: Validates that API parameter constraints (temperature, top_p, etc.) are correctly enforced. All tests must pass before proceeding to benchmark evaluation.
2. OCRBench: 5 minutes smoke test for multimodal pipelines.
3. MMMU Pro: Verify Vision input preprocessing by testing diverse visual inputs.
4. AIME2025: Long-output stress test. Catches KV cache bugs and quantization degradation that short benchmarks hide.
5. K2VV ToolCall: Measures trigger consistency (F1) and JSON Schema accuracy. Tool errors compound in agents; we catch them early.
6. SWE-Bench: Full agentic coding test. (Not open sourced due to dependency of sandbox)

**Upstream Fix**: We embed with vLLM/SGLang/KTransformers communities to fix root causes, not just detect symptoms.

**Pre-Release Validation**:
Rather than waiting for post-deployment complaints, we provide early access to test models. This lets infrastructure providers validate their stacks before users encounter issues.

**Continuous Benchmarking**:
We will maintain a public leaderboard of vendor results. This transparency encourages vendors to prioritize accuracy.

## Testing Cost Estimation

We completed full evaluation workflow validation on Two NVIDIA H20 8-GPU servers, with sequential execution taking approximately 15 hours. To improve evaluation efficiency, scripts have been optimized for long-running inference scenarios, including streaming inference, automatic retry, and checkpoint resumption mechanisms.

## An Open Invitation

**Weights are open. The knowledge to run them correctly must be too.**

We are expanding vendor coverage and seeking lighter agentic tests. **Contact Us: [email protected]**
