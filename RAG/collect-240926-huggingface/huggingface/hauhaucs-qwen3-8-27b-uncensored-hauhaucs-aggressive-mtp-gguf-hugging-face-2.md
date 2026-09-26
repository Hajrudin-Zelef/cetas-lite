---
id: collect-240926-huggingface/huggingface/hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face-2
title: "hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "OpenAI", "Unsloth"]
dates: []
keywords: ["gguf", "qwen", "agents", "apache", "cost", "license", "llama", "llama.cpp", "memory", "reasoning"]
source: docs/RAG/clean_en/huggingface/hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face.md
source_anchor: ""
source_lines: [129, 227]
sha256: 9caa0e8767ca0e9495de751018276e278fde69785352c0dee9457294cbf6e591
---

# hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face

| Quant | PP tok/s | TG tok/s | 
|---|---|---|
| Q2_K_P | 1959.14 | 121.88 | 
| Q3_K_P | 1944.73 | 112.76 | 
| Q4_K_P | 1860.34 | 92.60 | 
| Q5_K_P | 1737.51 | 83.25 | 
| Q6_K_P | 1747.60 | 72.89 | 
| Q8_K_P | 1827.29 | 59.00 | 
| IQ2_M | 1884.83 | 121.25 | 
| IQ3_M | 1867.48 | 108.45 | 
| IQ3_XS | 1880.59 | 111.77 | 
| IQ4_XS | 1978.46 | 104.25 | 

With HauhauCS FastMTP enabled, the final Q3_K_P reached **138.37 document TG and 87.95 reasoning TG on the same Ada—23.5% and 3.9% faster than the pinned Unsloth Q3 control.**

From the official Qwen3.8-27B model card:

**Thinking mode (default):**

- `temperature=1.0`
- `top_p=0.95`
- `top_k=20`
- `min_p=0.0`
- `presence_penalty=0.0`
- `repetition_penalty=1.0`
- `reasoning_effort=xhigh` for the deepest reasoning

**Instruct / non-thinking mode:**

- `temperature=0.7`
- `top_p=0.80`
- `top_k=20`
- `min_p=0.0`
- `presence_penalty=1.5`
- `repetition_penalty=1.0`
- `enable_thinking=false`

Qwen3.8 supports `xhigh`, `medium`, and `low` reasoning effort. Thinking and preserved reasoning are enabled by default in the official model contract.

**Important:**

- Use `--jinja` for the embedded chat template.
- Use the BF16 projector for Vision.
- The model's native maximum is `262144` .
- Context length and KV precision have a large VRAM cost. Reduce context before reducing model quality if your workload does not need maximum native context.
- Keep default F16 K/V on the lower tiers unless memory pressure requires otherwise.

If your llama.cpp build does not recognize the reasoning or MTP flags, update it. Older builds may still load the GGUF but will not expose the full Qwen3.8 serving path.

Qwen3.8 uses thinking mode by default. Disable it when you want shorter, faster direct responses.

**Example llama-server default for all requests:**

```
--chat-template-kwargs '{"enable_thinking":false}'
```
**Example per request through the OpenAI-compatible API:**

```
{
  "model": "qwen3.8-27b-aggressive-q3",
  "messages": [{"role": "user", "content": "..."}],
  "chat_template_kwargs": {"enable_thinking": false}
}
```
**Example for multi-turn agents, preserve prior reasoning context with:**

```
{
  "chat_template_kwargs": {"preserve_thinking": true}
}
```
- **llama.cpp:** recommended; use a current Qwen3.8/MTP-capable build
- **LM Studio, Jan, KoboldCpp, and other GGUF frontends:** base compatibility depends on their bundled llama.cpp version
- **Embedded MTP:** optional and stock-compatible in current llama.cpp
- **HauhauCS FastMTP:** optional; requires the sidecar and HauhauCS-FastMTP-llama.cpp.patch
- **Vision:** requires the separate BF16 projector
- **K_P display:** may appear as`?` in UIs that do not recognize the suffix

Every GGUF is covered by the signed HauhauCS release manifest. Exact SHA-256 values identify byte-for-byte mirrors after renaming; canonical tensor fingerprints continue to identify HauhauCS tensors after metadata-only rewriting.

The FastMTP sidecar's exact file SHA-256 is `115e618e1f73cb50817ed5856f0551c6bf9c3d94df96f440eaca78dc63b8968b`; its canonical tensor fingerprint is `49e248e799f169b6ccc6a8127b9300a95f06cf3d96a8353266f5d457e81d1c87`. The public-key DER fingerprint is `f7be4a2335582ab7b2e393ca1c40ce70e483f1492c0f57b8c6e05d8a7223833c`.

Download HauhauCS-RELEASE-MANIFEST.json, its signature, FastMTP-PROVENANCE.json, its signature, and HauhauCS-FastMTP-Ed25519-PUBLIC.pem, then verify:

```
openssl pkeyutl -verify -rawin -pubin \
  -inkey HauhauCS-FastMTP-Ed25519-PUBLIC.pem \
  -in FastMTP-PROVENANCE.json \
  -sigfile FastMTP-PROVENANCE.json.sig
openssl pkeyutl -verify -rawin -pubin \
  -inkey HauhauCS-FastMTP-Ed25519-PUBLIC.pem \
  -in HauhauCS-RELEASE-MANIFEST.json \
  -sigfile HauhauCS-RELEASE-MANIFEST.json.sig
```
Qwen3.8-27B is released by Qwen under the Apache 2.0 license. This quantized Aggressive variant retains that license.

- Downloads last month
- 2,046,935
