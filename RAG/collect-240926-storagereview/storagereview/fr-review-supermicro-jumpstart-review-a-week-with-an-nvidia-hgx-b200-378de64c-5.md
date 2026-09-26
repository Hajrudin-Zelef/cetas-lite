---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c-5
title: "fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["blackwell", "fp4", "fp8", "gpu", "llama", "nvfp4", "tensorrt", "tensorrt-llm", "throughput", "vllm"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c.md
source_anchor: ""
source_lines: [86, 90]
sha256: 437a86f273c8363f20fcd5214a25775016ea905ee6d8651a4fb60550b6941c3d
---

# fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c

The results obtained with the quantized model were unexpected and warrant further analysis. In several cases, the NVFP4 and FP8 quantized versions of the models did not achieve the expected performance gains compared to their native precision counterparts. For example, the Llama 3.1 8B FP4 model reached only 830.46 tok/s total throughput at BS=1, versus 1,727.62 tok/s for the standard precision variant, despite similar per-user throughput. For larger batch sizes, although the quantized models approached standard precision throughput (29,200 to 29,300 tok/s vs. 32,800 tok/s at BS=256), the overall results suggest that the current vLLM implementation may not be fully optimized for Blackwell.
We plan to conduct additional testing on vLLM and also compare it to TensorRT-LLM to determine the performance that end users can expect today.
Supermicro JumpStart revolutionizes the AI proof-of-concept game.
Supermicro's JumpStart program delivers on its promises: real, unrestricted access to production hardware, without the logistical constraints, lead times, or lab costs of a traditional PoC. Our JumpStart week on the X14 HGX B200 platform started quickly, ran smoothly, and allowed us to evaluate performance as if we were using hardware installed in our own racks.
For organizations making rapid decisions about their AI infrastructure, this type of platform access reduces weeks of planning to a few days. Whether validating GPU throughput, storage behavior, model performance, or simply confirming software stack compatibility, JumpStart provides answers with ease. This approach, which prioritizes confidence in hardware evaluation, is very promising, and we would like to see more vendors adopt it.
