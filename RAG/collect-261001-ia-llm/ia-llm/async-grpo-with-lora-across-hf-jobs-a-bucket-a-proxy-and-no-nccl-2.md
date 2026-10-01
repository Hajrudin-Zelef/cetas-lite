---
id: collect-261001-ia-llm/ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl-2
title: "every Job gets the same bucket at the same absolute path"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "vLLM"]
dates: []
keywords: ["agent", "attention", "compute", "decode", "deepseek", "gpu", "inference", "kv cache", "lora", "preemption", "prefill", "qwen"]
source: docs/RAG/collect-261001-ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl.md
source_anchor: ""
source_lines: [64, 131]
sha256: 40881115dce104023c96e81bd5cdc4b5223c71ce7c9a8f9d95e0ce05a17bd4da
---

# every Job gets the same bucket at the same absolute path

We chose `sail/Sanity-Test-R1D-1.5B`, the dataset from Defeating the Training-Inference Mismatch via FP16 (Qi et al., 2025). The reproduction code is in `sail-sg/Precision-RL`.

The authors generated 40 answers for each MATH problem with DeepSeek-R1-Distill-Qwen-1.5B. They kept problems with a success rate between 20% and 80%, yielding 1,460 questions. This dataset is really good for RL validation because the questions are neither already solved nor completely hopeless for that model, meaning the model can get a good early signal to train on and improve.

This is awesome as a robust end-to-end test: if one vLLM replica silently serves the base model under an adapter name, we want to see that in the curve within a few dozen steps. Also, this dataset is small enough to cycle through in less than two hours.

We also take the hyperparameters from the paper's LoRA scripts in `oat/scripts/lora`: `Qwen/Qwen2.5-Math-1.5B`, LoRA rank 1 with alpha 2, a learning rate of 4e-5, 8 samples per prompt, 128 completions per step, a maximum of 3,000 generated tokens and a 4,096-token context.

The trainer uses the same `vllm/vllm-openai:v0.27.1` image with TRL installed on top. We ran the PR branch at the time; the same code now ships in TRL v1.14. The training script is a normal `AsyncGRPOTrainer` script. The only Job-specific values are the output directory and the server URL.

```
from peft import LoraConfig
from trl.experimental.async_grpo import AsyncGRPOConfig, AsyncGRPOTrainer
config = AsyncGRPOConfig(
    output_dir="/lora/sanity-lora-r1",       # on the bucket: adapters, checkpoints and the final adapter all land here
    vllm_server_base_url="http://localhost:8000",   # the proxy, not a vLLM Job; TRL never sees the Jobs URLs
    max_staleness=4,
    weight_sync_steps=4,                     # publish an adapter every 4 optimizer steps
    save_strategy="steps", save_steps=50,    # checkpoints go to the same bucket -> resume after preemption
    ...
)
trainer = AsyncGRPOTrainer(
    model="Qwen/Qwen2.5-Math-1.5B",
    args=config,
    peft_config=LoraConfig(r=1, lora_alpha=2, target_modules="all-linear"),  # plain LoRA vLLM can serve as-is
    ...
)
```
During initialization, TRL calls `/server_info`. If it finds a `lora_config`, it uses adapter-only sync. Configurations vLLM cannot serve directly, such as DoRA, `modules_to_save`, or a rank above `--max-lora-rank`, fall back to merged-weight sync with a warning. The log should contain `Adapter-only vLLM sync enabled`.


Now onto the fun stuff. We need a proxy between the trainer and the vLLM Jobs for two reasons:

1. Exposed Job ports require an `Authorization: Bearer <HF token>` header on every request. The proxy is where that header gets added, so TRL does not need to know about it.
2. We want more than one GPU generating. On a single vLLM server, the usual way to get that is `--data-parallel-size > 1` , but TRL refuses adapter-only sync in that mode, for a good reason: a call to`/v1/load_lora_adapter` only reaches the DP rank that answers it, so the other ranks would keep serving the base model under the new policy name. On Jobs the question does not even arise, since each replica is its own machine. So the data parallelism has to live one level up, in something that fans the adapter load out to every replica.

We therefore run a small proxy at `127.0.0.1:8000` on the trainer Job and point TRL to it as if it were a single vLLM server. Besides adding the header, the proxy does two things functionally:

- It sends each completion request to one replica, chosen so that the eight rollouts of a prompt land where their prefix is already cached (details on this below).
- It broadcasts every *state-changing* request, such as**adapter loads, pause and resume** , to all replicas, so that a policy name means the same thing everywhere.

A quick reminder of why this matters. Generating a completion has two phases with very different workload profiles:

- The prefill processes the **whole prompt at once** and computes the attention keys and values for every prompt token.
- The decode phase then produces one token at a time, and each new token attends to the keys and values of all the tokens before it.

Those keys and values are the KV cache. Because attention is causal, the KV of a token depends only on the tokens before it, not on what comes after. Two requests that share a prefix therefore share the KV of that prefix, and a replica that already has it in cache can **skip that part of the prefill entirely**. The whole game now is to find that replica, so a request can benefit from landing on the replica that has already seen its prefix.

vLLM stores its prefix KV cache in blocks of 16 tokens. Because of GRPO, the rollout worker sends `G` requests with the same prompt (in our case `G=8`). If they all reach the same replica, the first request computes the prefill and the next seven reuse it. With round-robin routing, half would go to a replica that doesn't have the prefix cached, and those four requests would redo the prefill work and waste valuable GPU compute.

The job of our router is to track which replica has seen which **block hash**. One important detail is that the hashes are chained, so the hash of block 3 represents blocks 1, 2 and 3, not just block 3. This mirrors causal attention: the KV of block 3 is only valid if blocks 1 and 2 are the same too. We also seed the chain with the adapter name because the KV cache also depends on the adapter that generated it: a prefix cached for policy v3 is useless for policy v4!

The video walks through the entire decision process for choosing a replica. The steps below go through a real 135-token completion request example (from the Sanity dataset problems):

**1. Split the prompt into blocks.** The router receives token ids and cuts them into 16-token blocks, just like vLLM. It only hashes complete blocks, so the last 7 tokens are ignored here.

**2. Hash the prefix.** Each block is hashed with the previous hash, starting from the adapter seed. `h3` therefore identifies blocks 1, 2 and 3 in order. Two prompts with the same first `k` blocks get the same hashes up to `hk`. Once one block changes, every hash after it changes too. This is why we seed the hashes with the adapter name: the same prompt under `trl-policy-v4` starts from another seed and cannot match entries from `v3`. This is what we want because the old KV blocks were computed with different weights.

**3. Compare two prompts.** Problem 1 has 103 tokens. Both prompts start with the same 23-token chat template. Their first block is identical, but block 2 already contains the problem text. The hashes differ from there.

**4. Store the owners.** For every hash, the router remembers which replicas served it and which hashes came after it (we cap the successor set at two because we only need to know whether a block has one continuation or several). After a few prompts, the template block `h1` is owned by both replicas and already has several successors, `h2` to `h8` are owned by A only and each has a single successor, and `h2'` to `h6'` are owned by B only.

In practice, every prompt in a run starts with the same tokens. Here, it is the chat template and the system prompt, which amount to the first 23 tokens of all 1,460 problems. In an agent setting, it would be the tool descriptions, and in a multi-turn environment it would be the shared conversation history. These blocks are in every replica's cache within seconds, so matching on them tells us nothing about where a particular prompt lives.

A block is *common* if every replica has served it, or if it has more than one successor. Common blocks are ignored during routing because they do not identify a particular prompt. More on this below!

**5. Pick a replica.** The router counts how many leading blocks match on each replica and removes the common prefix. What is left is the number of blocks specific to this prompt. Then:

