---
id: collect-261001-ia-llm/ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl-3
title: "every Job gets the same bucket at the same absolute path"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "vLLM"]
dates: []
keywords: ["decode", "inference", "kv cache", "lora", "prefill", "qwen", "training", "vllm"]
source: docs/RAG/collect-261001-ia-llm/async-grpo-with-lora-across-hf-jobs-a-bucket-a-proxy-and-no-nccl.md
source_anchor: ""
source_lines: [132, 211]
sha256: 254dc04fb74e2d03a2f520a921d9a85b16f3f3519a6571ae8bb2e79cf81567a7
---

# every Job gets the same bucket at the same absolute path

- If one replica has specific blocks and it is not swamped, meaning it is at most 8 requests ahead of the least-loaded replica, the request goes there. We call this an **affinity** hit.
- If a replica has specific blocks but it is more than 8 requests ahead, we give up on the cache and send the request to the least-loaded replica. We call this a **spill** .
- If no replica has specific blocks, this is a new prompt. It goes to the least-loaded replica, round-robin on ties. We call this **unmatched** .

Here is the rule applied to four requests. Start from a state where replicas A and B both have 3 requests in flight, and only the template block `h1` is known on both replicas.

- **Request 1, problem 0, rollout 1.** Both replicas match one block, the template, and that block is common. So nothing specific matches anywhere. The request is unmatched, both replicas are equally loaded, and round-robin sends it to A. The router records`h2` to`h8` as owned by A. A now has 4 requests in flight.
- **Request 2, problem 0, rollout 2.** Same prompt. A matches all 8 blocks, B matches only the template. After removing the one common block, A has 7 specific blocks and B has none. A is only 1 request ahead of B, well within the limit of 8, so the request goes to A. This is an affinity hit: A already has the whole prompt in its KV cache.
- **Request 3, problem 1, rollout 1.** A new prompt. Both replicas match only the template block, so nothing specific matches. The request is unmatched and goes to the least-loaded replica, B, which has 3 in flight against A's 5. The router records`h2'` to`h6'` as owned by B.
- **Request 4, problem 0, rollout 9.** Suppose that by now A has 12 requests in flight while B is back to 3. A still has the 7 specific blocks, but it is now 9 requests ahead of B, more than the limit. The request spills to B. B prefills problem 0 once, and the router records`h2` to`h8` as owned by B as well. Every replica has now served those blocks, so problem 0 becomes common too, and its later rollouts are placed by load alone.

**6. Reuse the prefill.** Request 2 is the reason for doing all this. It reuses the prefill computed by Request 1: blocks 1 to 8 are already in A's KV cache, so A skips straight to decoding the completion. Had it gone to B, B would have prefilled all 135 tokens again while A's cache sat unused. Request 3 shows why the `common` rule is needed. Without it, the shared chat template would make every new prompt look like a cache hit. Request 4 keeps the load bounded. Saving one prefill is not worth letting a replica fall far behind.

```
def choose(self, upstreams, model, prompt):
    hashes = self.block_hashes(model, prompt)         # chained blake2b over 16-token blocks, seeded with `model`
    matched = self.matched_prefix(hashes)             # per replica: leading blocks it has served
    common = self.common_prefix_len(hashes)           # leading blocks that identify no prompt (see below)
    specific = [max(0, m - common) for m in matched]  # what actually distinguishes replicas
    least = min(u.inflight for u in upstreams)
    best = max(range(self.n), key=lambda i: (specific[i], -upstreams[i].inflight))
    if specific[best] > 0 and upstreams[best].inflight - least <= self.cfg.imbalance:
        pick = best                                   # affinity: the replica that has this prompt, and is not swamped
    else:
        candidates = [i for i in range(self.n) if upstreams[i].inflight == least]
        pick = candidates[self.rr % len(candidates)]  # spill or new prompt: least-loaded, round robin on ties
        self.rr += 1
    ...record `pick` as an owner of every block, and each block's successor...
    return upstreams[pick]
```
The `common` prefix is the annoying part. Every request starts with the same system prompt and chat template. A simple longest-prefix match would give the first replica a match for almost every new prompt. We detect the shared prefix through fan-out instead: a block with several different successors is common, while a block that always leads to the same successor belongs to a particular prompt. Only the blocks after that common prefix count as affinity.

The proxy also needs to broadcast the adapter loads to every replica. We treat the operation as **all-or-nothing**. Each replica has its own bucket mount, so they do not necessarily see a new adapter at exactly the same time. A `No adapter found for <path>` error usually means that one bucket mount has not caught up yet, and we retry only that replica. For any other error, we unload the adapter from the replicas that accepted it so that a policy name never exists on only part of the replicas.

```
async def load_one(u):
    while True:
        status, _, out = await send(u, "POST", "/v1/load_lora_adapter", headers, body)
        if status == 200 or "No adapter found" not in out.decode() or time.monotonic() > deadline:
            return u, status, out
        await asyncio.sleep(cfg.lora_retry_s)          # this replica's mount has not seen the directory yet
results = await asyncio.gather(*(load_one(u) for u in ups))
if any(st != 200 for _, st, _ in results):
    await asyncio.gather(*(send(u, "POST", "/v1/unload_lora_adapter", headers, unload) for u, st, _ in results if st == 200))
    return web.Response(status=504 if timed_out else st, text="rolled back on the others")
```
We also broadcast `/pause`, `/resume` and `/v1/unload_lora_adapter` in the same way. `/health` returns 200 only if every replica is healthy. `/server_info` and `/v1/models` only need one answer. From TRL's point of view, the proxy is a single `data_parallel_size=1` server, so it selects adapter-only sync.

We initially wondered whether a Python asyncio proxy would become a bottleneck. It does not (at least at this scale). There are at most 128 non-streaming JSON requests in flight, and routing only computes a few hashes. One thread handles this easily. A more refined router that needs to handle more traffic would probably need to be written in a faster language (We see you 🦀).

The numbers below come from the trainer's logged metrics on trackio. The run uses `Qwen/Qwen2.5-Math-1.5B`, LoRA `r=1` on `all-linear`, 128 completions per step and 8 rollouts per prompt. It runs for 500 steps and saves a checkpoint every 50 steps. The trainer uses an `h200x2` Job and each of the two vLLM replicas uses one `h200` Job. Running all three costs ~$20 per hour.

| per sync, trainer's clock, 126 syncs | before | now (p50) | 
|---|---|---|
| whole sync | 30.8 s | **8.5 s** (min 6.6, max 9.2) | 
| of which: pause both replicas | 0.3 s | 0.3 s | 
| adapter all-gather and save to the bucket | 0.6 s | 1.1 s | 
| both replicas accept the adapter | ~29 s | ~7 s | 

All 252 adapter loads succeeded 🎉: 126 syncs times 2 replicas. Six succeeded on the second attempt and 246 on the third.

At the end of the run, after 64,728 rollouts, the proxy's counters read:

```
routed [31928, 32800]  affinity 54712  spilled 820  unmatched 9196
```
With 8 rollouts per prompt, at least one of the eight requests must be cold. The theoretical minimum is therefore 12.5 %. The router gets 14.2 % unmatched requests, 84.5 % affinity hits, and 1.3 % spills. There is not much left to gain here unless we start looking at each replica's load using deeper inference-side metrics based on real measured load.

The first configuration has a pretty obvious problem: the trainer is the bottleneck, not generation. Over the 500 steps, we have:

| per optimizer step, p50 |  | 
|---|---|
| step | 22.9 s | 
| forward + backward | 21.9 s | 
| waiting for rollouts | 0.02 s | 
| rollout queue occupancy | 476 of 512 | 
| trainer MFU | 3.9 % | 

The rollout queue stays full, and the worker is mostly blocked by backpressure. The second replica is basically useless in this configuration. We'll see later in the post how we went through runs that moved the bottleneck between training and generation to make the run 3.9× faster.

