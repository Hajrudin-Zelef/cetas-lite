---
id: collect-261001-general-networking/general-networking/welcome-inkling-by-thinking-machines-4
title: "model_id = \"thinkingmachines/Inkling-NVFP4\""
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Z.ai", "vLLM"]
dates: []
keywords: ["nvfp4", "agentic", "claude", "deepseek", "fable 5", "gemini", "glm", "inference", "kimi", "kv cache", "mcp", "reasoning"]
source: docs/RAG/collect-261001-general-networking/welcome-inkling-by-thinking-machines.md
source_anchor: ""
source_lines: [368, 435]
sha256: be672bce2925eaf9322fe977217ecd914f924abd58cfb70375dbf8f10cdf76ff
---

# model_id = "thinkingmachines/Inkling-NVFP4"

To deploy Inkling on a cluster, we provide SLURM scripts serving with transformers API, as well as how to query the endpoint with different modalities. You can adapt these scripts to vLLM or SGlang by updating the commands. These scripts live here.

You can deploy an Inkling-Small NVFP4 checkpoint using Inference Endpoints. We provide a pre-tested configuration to deploy it on 8 RTX PRO 6000s with an aggregated VRAM of·768 GB, giving large space for KV cache, costs $ 22 per hourly uptime (scales to zero when unused). To deploy, run hf endpoints catalog deploy --repo thinkingmachines/Inkling-Small-NVFP4 using the Hugging Face CLI or alternatively, head to https://endpoints.huggingface.co/new/thinkingmachines/Inkling-Small-NVFP4. With this setup, you can get 140 TPS (single user inference, prepared for multi-user support with continuous batching)

Once the endpoint is up, you can query as follows.

```
curl "YOUR_ENDPOINT_HERE" \
-X POST \
-H "Authorization: Bearer $HF_TOKEN" \
-H "Content-Type: application/json" \
-d '{
    "model": "thinkingmachines/Inkling-Small-NVFP4",
    "messages": [
        {
            "role": "user",
            "content": [
                {
                    "type": "image_url",
                    "image_url": {
                        "url": "https://endpoints.hf.co/media-examples/img1.png"
                    }
                },
                {
                    "type": "text",
                    "text": "Describe this image in one sentence."
                }
            ]
        }
    ],
    "stream": true,
    "max_tokens": 100
}'
```
|  |  | Inkling Small | Inkling | Nemotron 3 Ultra | Kimi K2.5 | Kimi K2.6 | GLM 5.2 | DeepSeek V4 Pro | Gemini 3.1 Pro (high) | Claude Fable 5 (max) | GPT 5.6 Sol (xhigh) | 
|---|---|---|---|---|---|---|---|---|---|---|---|
| **Reasoning** |  |  |  |  |  |  |  |  |  |  |  | 
|  | HLE (text only) | 31.6% | 29.7% | 26.6% | 29.4% | 35.9% | 40.1% | 35.9% | 44.7% | 53.3% | 47.2% | 
|  | HLE (with tools) | 47.8% | 46.0% | 37.4% | 50.2% | 54.0% | 54.7% | 48.2% | 51.4% | 64.5% | 55.0% | 
|  | AIME 2026 | 95.5% | 97.1% | 94.2% | 95.8% | 96.4% | 99.2% | 96.7% | 98.3% | – | 99.9% | 
|  | GPQA Diamond | 89.5% | 87.2% | 86.7% 0 | 87.9% | 91.1% | 89.5% | 88.8% | 94.1% | 92.6% | 94.1% | 
| **Agentic (coding)** |  |  |  |  |  | 0 |  |  |  |  |  | 
|  | SWEBench Verified | 80.2% | 77.6% | 70.7% | 76.8% | 80.2% | – | 80.6% | 80.6% | 95.0% | – | 
|  | SWEBench Pro (Public) | 55.9% | 54.3% | 46.4% | 50.7% | 58.6% | 62.1% | 55.4% | 54.2% | 80.0% | 64.6% | 
|  | Terminal Bench 2.1 (Best Harness) | 64.69 | 63.8 | 56.4 | 51.3 | 71.3 | 82.7 | 64 | 73.8 | 84.6 | 89.5 | 
|  | GDPVal-AA v2 | 1269 | 1233 | 1164 | 1009 | 1190 | 1514 | 1307 | 962 | 1760 | 1748 | 
| **Agentic (general)** |  |  |  |  |  |  |  |  |  |  |  | 
|  | MCP Atlas | 79.2% | 74.1% | 44.7% | 64.0% | 68.1% | 77.8% | 73.2% | 78.2% | 83.3% | 81.8% | 
|  | Tau 3 Banking | 15.5% | 23.7% | 13.8% | 13.2% | 20.6% | 26.8% | 25.8% | 16.5% | 26.8% | 33.0% | 
| **Factuality** |  |  | 0 |  |  |  |  |  |  |  |  | 
|  | BrowseComp (w/ Ctx) | 77.4% | 77.1% | – | 74.9% | 83.2% | – | 83.4% | 85.9% | 88.0% | 89.4% | 
|  | SimpleQA Verified | 20.6% | 43.9% | 32.4% | 36.9% | 38.7% | 38.1% | 57.0% | 77.3% | 68.3% | 71.6% | 
|  | AA Omniscience | -9 | 1.0% | -1.0% | -8.0% | 6.0% | 4.0% | -10.0% | 33.0% | 40.0% | 22.0% | 
| **Chat** |  |  |  |  | 0 |  |  |  |  |  |  | 
|  | IFBench | 82.2% | 79.8% | 81.4% | 70.2% | 76.0% | 73.3% | 76.5% | 77.1% | 63.5% | 72.7% | 
|  | Global-MMLU-Lite | 86.7% | 88.7% | 85.6% | 84.0% | 88.4% | 89.2% | 89.3% | 92.7% | 93.3% | 91.8% | 
| **Vision** |  |  |  |  |  |  |  | 0 | 0 |  | 0 | 
|  | MMMU Pro (Standard 10) | 74.0% | 73.3% | – | 75.0% | 79.0% | – | – | 82.0% | 84.2% | 83.0% | 
|  | Charxiv RQ | 77.4% | 78.1% | – | 77.5% | 80.4% | – | – | 80.2% | 86.5% | 84.7% | 
|  | Charxiv RQ (with python) | 82.3% | 82.0% | – | 78.7% | 86.7% | – | – | 89.9% | 89.4% | 87.8% | 
| **Audio** |  |  |  |  |  | 0 |  |  |  |  |  | 
|  | Audio MC | 54.9% | 56.6% | – | – | – | – | – | 66.8% | – | – | 
|  | MMAU | 77.0% | 77.2% | – | – | – | – | – | 82.5% | – | – | 
|  | VoiceBench | 90.1% | 91.4% | – | – | – | – | – | 94.3% 0 | – | – | 
| **Safety** |  |  |  |  |  |  | 0 |  |  |  |  | 
|  | FORTRESS (Adversarial) | 71.6% | 78.0% | 77.6% | 54.1% | 65.6% | 71.3% | 36.0% | 65.2% | 96.0% | 82.4% stationary | 
|  | FORTRESS (Benign) | 96.9% | 95.9% | 90.5% | 98.3% | 97.2% | 90.0% | 98.5% | 98.0% | 55.1% | 98.1% | 
|  | StrongREJECT | 98.4% | 98.6% | 98.7% | 99.5% | 99.8% | 98.5% | 98.6% | 98.0% | 98.7% | 98.5% |
