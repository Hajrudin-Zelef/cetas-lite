---
id: collect-240926-huggingface/huggingface/poolside-laguna-s-2-1-nvfp4-hugging-face-2
title: "Python headers: Triton JIT needs them and DGX OS ships without them."
domain: huggingface
role: reference
task: reference
actors: ["Poolside", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agentic", "decode", "fp4", "gpu", "guardrails", "license", "memory", "nvfp4", "prefill", "quantization", "reasoning", "sglang"]
source: docs/RAG/clean_en/huggingface/poolside-laguna-s-2-1-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [117, 185]
sha256: ae31323ab3887e5cb901fcc1a0302a3a4f7b3b49786fd770b09d23ec2cff8f1d
---

# Python headers: Triton JIT needs them and DGX OS ships without them.

```
export CUTE_DSL_ARCH=sm_121a          # arch string for FP4 kernel JIT
export PATH=/usr/local/cuda/bin:$PATH # nvcc for JIT
export MAX_JOBS=4                     # cap JIT fan-out; see warning below
source ~/venvs/vllm025/bin/activate
vllm serve poolside/Laguna-S-2.1-NVFP4 \
  --speculative-config '{"model":"poolside/Laguna-S-2.1-DFlash-NVFP4","num_speculative_tokens":7}' \
  --enable-auto-tool-choice \
  --tool-call-parser poolside_v1 \
  --reasoning-parser poolside_v1 \
  --max-num-seqs 32 \
  --max-model-len 262144 \
  --gpu-memory-utilization 0.85 \
  --host 0.0.0.0 --port 8000
```
- You do not need backend flags: auto-selection picks FlashInferCutlass, which
runs natively on `sm_121` . Do not set`--linear-backend flashinfer_b12x` on
0.25.1; the opt-in is broken there and it is slower anyway.
- Sampling comes from the checkpoint's `generation_config.json` (`top_k 20` eval-certified truncation); serve with those defaults rather than overriding
them. Do not add`min_p` : vLLM rejects`min_p` and`logit_bias` under
speculative decoding, so putting it in the defaults returns a 400 on every
sampled request.
- `--max-num-seqs 32` is required: DFlash crashes vLLM at the default of 256.
- The first start takes about 15 minutes (weight load from NVMe, JIT, and graph capture).

Never drop `MAX_JOBS=4` on a cold `~/.cache/flashinfer`. An uncapped nvcc
fan-out can exhaust the 128 GB unified memory and take the whole machine down.
A warm cache makes it a no-op, but the cache is cold again after any config or
shape change.


Prefill runs 600-800 tok/s, decode is around 15 tok/s on prose and 22-24 on code, and DFlash accepts 2.9-3.1 tokens per step. At the 256K setting you get 830-870K KV tokens. Two requests at once roughly double total throughput. Without speculation, decode sits at 13-14 tok/s on every engine we tried on the GB10, which is the memory-bandwidth ceiling for this model. Speculation is how you get past it.

*pool CLI against the box:*

```
POOLSIDE_STANDALONE_BASE_URL=http://<spark-ip>:8000/v1 \
POOLSIDE_STANDALONE_MODEL=poolside/Laguna-S-2.1-NVFP4 \
POOLSIDE_STANDALONE_CONTEXT_LENGTH=262144 \
POOLSIDE_API_KEY=dummy pool
```
Use the box's IPv4 address rather than the `.local` mDNS name: Go's resolver can
pick the link-local IPv6 and fail with "no route to host". On macOS the terminal
app needs Local Network permission, and the system curl is exempt from that
check, so if curl works but pool does not, it is the permission and not the
network.

The Laguna S 2.1 architecture is supported in SGLang via sgl-project/sglang#24204. Quantization is detected automatically from `quantization_config`, so no extra flags are required. See the SGLang cookbook entry and the main Laguna S 2.1 model card for a serving recipe.

The full Transformers recipe is on the main Laguna S 2.1 model card. Substitute `poolside/Laguna-S-2.1-NVFP4` for the model ID; quantization is detected automatically from `quantization_config`.

Laguna S 2.1 support ships in TensorRT-LLM `>=1.3.0rc16`; see the install recipe on the main Laguna S 2.1 model card. Substitute `poolside/Laguna-S-2.1-NVFP4` for the model ID; quantization is detected automatically from `quantization_config`, no extra flags required.

```
from tensorrt_llm import LLM
llm = LLM(model="poolside/Laguna-S-2.1-NVFP4", trust_remote_code=True)
```
Available on the Ollama library.

Laguna S 2.1-NVFP4 uses the same reasoning controls (interleaved thinking, preserved reasoning, and the `enable_thinking` flag) as the base model. See the Controlling reasoning section of the main Laguna S 2.1 model card.

This model is licensed under the OpenMDW-1.1 License.

Laguna S 2.1-NVFP4 is designed for software engineering and agentic coding use cases, and you are responsible for confirming that it is appropriate for your intended application. Laguna S 2.1-NVFP4 is subject to the OpenMDW-1.1 License, and should be used consistently with Poolside's Acceptable Use Policy. We advise against circumventing Laguna S 2.1-NVFP4 safety guardrails without implementing substantially equivalent mitigations appropriate for your use case.

Please report security vulnerabilities or safety concerns to security@poolside.ai.

- Downloads last month
- 435,542
