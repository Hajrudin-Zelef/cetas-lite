---
id: collect-240926-huggingface/huggingface/juspay-xor-hugging-face
title: "juspay-xor-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang"]
dates: []
keywords: ["apache", "blackwell", "distribution", "gpu", "gpus", "inference", "kv cache", "latency", "license", "memory", "nvidia", "packaging"]
source: docs/RAG/clean_en/huggingface/juspay-xor-hugging-face.md
source_anchor: ""
source_lines: [1, 90]
sha256: c7c250c3f962b04bab52c00fbd7072e148780c74c3e61fec53c0a301f929fad8
---

# juspay-xor-hugging-face

<!-- source: https://huggingface.co/juspay/jev-one -->

Xor (`xor`) is a post-trained version of Qwen/Qwen3.6-35B-A3B for typed decision tasks. It is served through a TypeSafe-compatible `/v1/systemone` API

| Field | Value | 
|---|---|
| Base checkpoint | `Qwen/Qwen3.6-35B-A3B` | 
| Architecture | Mixture-of-experts causal language model | 
| Total parameters | 35 billion | 
| Activated parameters | Approximately 3 billion per token | 
| Released precision | BF16 | 
| Base license | Apache License 2.0 | 
| Packaging | Fully merged weights; no adapter loading or merging is required | 

The server accepts a state and a map of typed questions:

- `noul` : binary probability
- `choice` : categorical decision and full probability distribution
- `score` : expected ordinal score and full probability distribution

Requests may also include an `images` array containing up to eight image URLs or data URLs. Typed questions can classify information from the supplied images and text

The serving layer performs deterministic single-token candidate readout, forward and reverse option-order evaluation, probability calibration, and schema conversion. The serving layer is part of the released inference configuration and must be used for reproducible results

The following commands download the validated release, verify and extract the serving bundle, and start Xor on two GPUs:

```
export XOR_REVISION=xor-v1
export INSTALL_ROOT="$PWD/xor-eval"
export MODEL_DIR="$INSTALL_ROOT/model"
hf download juspay/xor   --revision "$XOR_REVISION"   --local-dir "$MODEL_DIR"
(
  cd "$MODEL_DIR/serving"
  sha256sum -c xor-serving.tar.gz.sha256
)
tar -tzf "$MODEL_DIR/serving/xor-serving.tar.gz"
mkdir -p "$INSTALL_ROOT/runtime"
tar -xzf "$MODEL_DIR/serving/xor-serving.tar.gz"   -C "$INSTALL_ROOT/runtime"   --strip-components=1
cd "$INSTALL_ROOT/runtime"
export CUDA_VISIBLE_DEVICES=0,1
export TP_SIZE=2
export API_PORT=49001
./run.sh
```
The setup requires Linux x86-64, the Hugging Face CLI, Docker Engine with Docker Compose v2, the NVIDIA Container Toolkit, and approximately 120 GB of free disk space

When the smoke test succeeds, the API is available at `http://127.0.0.1:49001/v1/systemone`

Test it with the included request:

```
curl -sS -X POST http://127.0.0.1:49001/v1/systemone   -H 'Content-Type: application/json'   --data @examples/request.json
```
Test image input with the included data-URL request:

```
curl -sS -X POST http://127.0.0.1:49001/v1/systemone \
  -H 'Content-Type: application/json' \
  --data @examples/image-request.json
```
Xor was evaluated locally on the public tiers of JEVBench using harness commit `fd51755eb0c0b546ca206d764faf3302feca913e`, the existing `typesafe` adapter, and one request at a time

The run used 2 x NVIDIA RTX PRO 6000 Blackwell Server Edition GPUs with 96 GB each, tensor parallelism 2, and the pinned serving image listed below

| Tier | Attempted | Valid | Correct | Accuracy | Macro accuracy | Brier mean | ECE | p50 | p95 | 
|---|---|---|---|---|---|---|---|---|---|
| Easy | 48 | 48 | 48 | 1.0000 | 1.0000 | 0.0018 | 0.0257 | 0.0766 s | 0.0888 s | 
| Original | 72 | 72 | 70 | 0.9722 | 0.9722 | 0.0897 | 0.1285 | 0.0773 s | 0.0812 s | 
| Hard public | 111 | 111 | 86 | 0.7748 | 0.8033 | 0.3460 | 0.0579 | 0.1361 s | 0.3887 s | 

Operational success, coverage, schema validity, and strict schema validity were 1.0000 for all three public tiers

These are self-run public-tier results, not an official JEVBench rank. An independent full-suite evaluation, including held-out items, has been requested in JEVBench issue #20

| Setting | Value | 
|---|---|
| SGLang image | `lmsysorg/sglang@sha256:6bcaa47db52f78ce0d67863b8b2431221b79bc23204a80cad757fa819d00e921` | 
| Tensor parallelism | 2 | 
| Validated GPUs | 2 x NVIDIA RTX PRO 6000 Blackwell Server Edition, 96 GB each | 
| Maximum prefill tokens | 250,000 | 
| Static memory fraction | 0.85 | 

For directly comparable latency measurements, use the same GPU model, GPU count, tensor-parallel configuration, and pinned serving image. Latency from other hardware should be identified as hardware-specific

The model files occupy approximately 66 GB. The validated SGLang configuration reserves substantial GPU memory for its KV cache. Smaller prefill budgets may work on other two-GPU configurations but are not covered by the validated setup

The service does not require an API key when bound to localhost. Remote deployments must add authentication, TLS, rate limits, and request-size limits at the ingress layer

- Downloads last month
- 1,174
