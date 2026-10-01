---
id: collect-261001-general-networking/general-networking/hugging-face-models-on-foundry-managed-compute-2
title: "hugging-face-models-on-foundry-managed-compute"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Hugging Face", "Microsoft", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["compute", "foundry", "accelerator", "agent", "agentic", "agents", "amd", "cost", "embedding", "embeddings", "gguf", "governance"]
source: docs/RAG/collect-261001-general-networking/hugging-face-models-on-foundry-managed-compute.md
source_anchor: ""
source_lines: [70, 139]
sha256: 8f1444a61f50a95ee9feb8b9c9f1485738cab83e48f37aeb99a818726130fa09
---

# hugging-face-models-on-foundry-managed-compute

Because weights are pre-staged in Azure storage and runtime images live in a Microsoft-managed registry, your deployments won't need outbound network access to Hugging Face Hub — you can deploy to production inside a private network.

Hugging Face models on Foundry are powered by a versatile collection of community-built, open-source inference runtimes — each selected and tuned for Foundry Managed Compute, and matched to the model architectures it serves best. Across all runtimes, the systematic curation process means new versions and patches land on Foundry quickly, and existing model deployments are upgraded automatically — without requiring you to redeploy.

- **vLLM** — the default high-throughput serving engine for open large language models, tuned for production GPU workloads. Because Hugging Face is a direct contributor to vLLM, any model in the Transformers library can run on vLLM out of the box — so when a new model lands on Hugging Face, it can be served on Foundry the same day, with no waiting on a custom integration.
- **SGLang** — a serving engine for language and multi-modal models, with strong support for structured outputs (JSON, regex, grammar-constrained generation) that agentic and tool-using workloads depend on. Hugging Face and the SGLang team have built a Transformers backend integration for SGLang, so any model in the Transformers library runs on SGLang out of the box — and reaches Foundry the same day it lands on Hugging Face.
- **Text Embeddings Inference (TEI)** — the runtime for embedding, reranker, and sequence-classification models. Accelerator-specific images ship with kernels compiled for each GPU and CPU family Foundry supports, keeping the embedding hot path lean for RAG and semantic-search workloads.
- **llama.cpp** — the CPU and small-GPU path for GGUF-quantized models. Useful for cost-optimized deployments, smaller models, and CPU-only regions, with the same OpenAI-compatible API as vLLM and SGLang.
- **TensorRT-LLM and NIM** — used on NVIDIA hardware where NVIDIA's optimized kernels and Triton-based serving deliver meaningfully better latency or throughput for specific model families.
- **hf-serve** — Hugging Face's own multi-model inference server, used for model architectures outside the LLM and embedding fast paths (vision, audio, segmentation, and other Transformers-native pipelines) so the Collection can cover every modality with a consistent serving layer.

The Hugging Face Collection in the Foundry Model Catalog is where you start, and deployment is five steps:

1. **Browse the catalog and pick a model** — the deploy wizard also surfaces the model id, deployment template id, and`acceleratorType` you'll need if you're scripting the deploy via SDK or REST.
2. **Choose a deployment template** — latency- vs throughput-optimized, accelerator family, context length, quantization.
3. **Configure instance count** — scale throughput by adding model instances.
4. **Deploy** — from the portal, CLI, SDK, or REST.
5. **Score** via the unified Foundry endpoint with the SDK you already use.

A deployment template is the unit of choice in step 2: a named, versioned asset that pins the runtime, the accelerator family and count, the context length, and the runtime-specific tuning needed to serve the model well — so picking a template is the only knob you turn for "how do I want this model to run."

`qwen3-32b`, for example, ships with four templates the deploy wizard exposes side by side:

| Template | Runtime | Accelerator | Context | 
|---|---|---|---|
| `qwen–qwen3-32b–40k-nvidia-a100` | vLLM | 1 × A100 80 GB | 40K | 
| `qwen–qwen3-32b–40k-nvidia-h100` | vLLM | 1 × H100 80 GB | 40K | 
| `qwen–qwen3-32b–128k-nvidia-2xa100` | vLLM | 2 × A100 80 GB | 128K | 
| `qwen–qwen3-32b–128k-nvidia-2xh100` | vLLM | 2 × H100 80 GB | 128K | 

Each template arrives pre-tuned for the model — runtime settings, tool-call and reasoning parsers, scoring path, health probes, request concurrency, and any model-specific context-extension settings are all set by Microsoft, with any trade-offs called out inline in the template description. When you script the deploy, you reference the template and Foundry handles the rest.

```
from azure.identity import DefaultAzureCredential
from azure.mgmt.cognitiveservices import CognitiveServicesManagementClient
client = CognitiveServicesManagementClient(DefaultAzureCredential(), SUBSCRIPTION_ID)
deployment = client.managed_compute_deployments.begin_create_or_update(
    resource_group_name=RESOURCE_GROUP,
    account_name=ACCOUNT_NAME,
    deployment_name="qwen3-32b",
    resource={
        "sku": {"name": "GlobalManagedCompute", "capacity": 1},
        "properties": {
            "model": "azureml://registries/azure-huggingface/models/qwen--qwen3-32b/versions/1",
            "deploymentTemplate": "azureml://registries/azure-huggingface/deploymenttemplates/qwen--qwen3-32b--40k-nvidia-h100/labels/latest",
            "acceleratorType": "H100_80GB",
        },
    },
).result()
```
The deployment is reachable through the unified Foundry endpoint with the OpenAI SDK — the `model` field takes the deployment name you just created:

```
from openai import OpenAI
api_key  = client.accounts.list_keys(RESOURCE_GROUP, ACCOUNT_NAME).key1
endpoint = f"https://{ACCOUNT_NAME}.services.ai.azure.com/openai/v1"
openai_client = OpenAI(base_url=endpoint, api_key=api_key)
completion = openai_client.chat.completions.create(
    model=deployment.name,
    messages=[{"role": "user", "content": "What is the capital of France?"}],
)
print(completion.choices[0].message)
```
A chat-completions model from the Collection slots into Foundry Agents as an admin-connected model and is callable through the Foundry Responses API with the same OpenAI SDK — same auth, same endpoint, same observability.

**Available now in preview:** the Hugging Face Collection in the Microsoft Foundry Model Catalog — thousands of models across every modality, refreshed weekly, deployable onto Foundry Managed Compute with NVIDIA A100, NVIDIA H100, or AMD MI300X accelerators in Global and Data Zone scopes, behind a unified Foundry endpoint with Playground support, first-class Azure Monitor metrics, per-deployment billing tags, and curated runtime upgrades and CVE patching applied automatically to your deployments.

**On the roadmap:** broader coverage of the Hugging Face ecosystem, additional accelerator families, and Bring Your Own Weights for fine-tuned and proprietary variants deployed through the same templates and governance as Collection models.

Hugging Face is where open models are published and discovered. Microsoft Foundry is where enterprises operationalize them — on curated, license-screened, security-screened weights hosted in Azure; on community-built and CVE-scanned runtimes; behind a single endpoint with enterprise identity, networking, observability, and agent integration on top. The breadth of the open-source ecosystem, with the operational layer Microsoft runs underneath. For a deep dive on Foundry Managed Compute — pricing, accelerator SKUs, data residency, enterprise readiness, observability, and the full Responses API + memory pattern — see the Managed Compute launch blog.
