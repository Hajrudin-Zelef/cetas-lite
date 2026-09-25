---
id: collect-250926-servers-hardware/servers-hardware/lora-hot-swapping-guide-unsloth-documentation
title: "lora-hot-swapping-guide-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["lora", "fp8", "gpu", "llama", "memory", "quantization", "training", "vllm"]
source: docs/RAG/clean4/lora-hot-swapping-guide-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 37]
sha256: 3e569343a461403f4c64717fd5fe68796fa63ded82757122d3675a2aaa60699b
---

# lora-hot-swapping-guide-unsloth-documentation

To enable LoRA serving for at most 4 LoRAs at 1 time (these are hot swapped / changed), first set the environment flag to allow hot swapping:

`export VLLM_ALLOW_RUNTIME_LORA_UPDATING=True`
Then, serve it with LoRA support:

```
export VLLM_ALLOW_RUNTIME_LORA_UPDATING=True
vllm serve unsloth/Llama-3.1-8B-Instruct \
    --quantization fp8 \
    --kv-cache-dtype fp8 \
    --gpu-memory-utilization 0.8 \
    --max-model-len 65536 \
    --enable-lora \
    --max-loras 4 \
    --max-lora-rank 64
```
To load a LoRA dynamically (set the lora name as well), do:

```
curl -X POST http://localhost:8000/v1/load_lora_adapter \
    -H "Content-Type: application/json" \
    -d '{
        "lora_name": "LORA_NAME",
        "lora_path": "/path/to/LORA"
    }'
```
To remove it from the pool:

For example when finetuning with Unsloth:

Then after training, we save the LoRAs:

We can then load the LoRA:

Last updated

Was this helpful?
