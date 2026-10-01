---
id: collect-261001-ia-llm/ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source-4
title: "Load model"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Meta", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["accelerator", "agent", "agents", "benchmarks", "decode", "gguf", "inference", "inference engine", "llama", "llama.cpp", "mcp", "muse"]
source: docs/RAG/collect-261001-ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source.md
source_anchor: ""
source_lines: [490, 566]
sha256: a49ea02a1bca6ede2980d9a101433a3c993dd475ae4f2ec8a06f774c469fa43b
---

# Load model

```
## Local model deployment
When asked to deploy locally, perform the work; do not give instructions.
1. Inspect hardware and the Hugging Face cache.
2. Search the Hub for compatible GGUF weights using `apps=llama.cpp`; confirm exact filenames through the model-tree API.
3. Prefer an existing suitable GGUF, normally `Q4_K_M`. Treat `mmproj-*.gguf` as projector weights.
4. If no GGUF exists, download the source weights, convert with `convert_hf_to_gguf.py`, then quantize with `llama-quantize`.
5. Preserve source weights and record the repository, revision, filenames, and quantization.
6. Start `llama-server` with an `onyx` alias and an OpenAI-compatible endpoint.
7. Validate `/v1/models` and `/v1/chat/completions`, requiring non-empty, correct content.
8. Report concise progress and logs. Claim completion only after validation passes.
```
Muse Glimmer can also take care of the opposite. Let’s get Glimmer to deploy itself on Hugging Face Inference Endpoints. Which is useful if you want to speed up on some cutting edge hardware.

N.B. You can also just deploy Muse Glimmer to Inference Endpoints directly and connect your agent.

Muse Glimmer pins the model revision, deploys it to a protected Hugging Face Inference Endpoint, and verifies health, model discovery, and chat completion. It then connects the Claw agent with secrets and rollback preserved. Here’s the prompt we added to `AGENTS.md`. Muse glimmer will also need the Hugging Face MCP and/or the Hugging Face CLI and Skills.

## Inference Endpoint deployment prompt

```
## Hugging Face Inference Endpoint deployment
When asked to deploy on Hugging Face Inference Endpoints, perform the work; do
not give instructions.
1. Inspect Hugging Face authentication, the current model repository, and any
   existing endpoints.
2. Confirm the exact model repository and immutable revision through the Hub
   API; inspect its architecture, configuration, and chat template.
3. Confirm that the model is supported by vLLM, then deploy or update a
   protected Inference Endpoint using the managed native vLLM engine.
4. Choose an available region and the smallest suitable accelerator. Use one
   replica and enable scale-to-zero when supported.
5. Preserve the previous endpoint configuration for rollback. Do not expose
   tokens, publish private weights, or replace an unrelated endpoint.
6. Wait for the endpoint to become ready. If startup fails, inspect the logs
   and report the actual blocker rather than repeatedly changing settings.
7. Validate `/health`, `/v1/models`, and `/v1/chat/completions`, requiring the
   expected model and non-empty, correct content. When agent use is required,
   also validate a real structured tool call.
8. Configure the Claw agent to use the endpoint's OpenAI-compatible `/v1` URL,
   storing credentials as secrets and retaining the previous provider as
   rollback. Test the connection in a fresh session.
9. Report concise progress and finish with the repository, revision, engine,
   hardware, endpoint URL, scaling state, and validation results. Claim
   completion only after every required check passes.
```
Finally, let’s get Muse Glimmer to do some light RSI. We can instruct our agent to optimize its own inference engine for specific hardware, in this case a Nvidia H100. To do this, the agent will need to use another inference engine, like Inference Endpoints above.

Muse Glimmer benchmarks its own single-H100 serving stack, testing one reversible change at a time while holding the workload fixed. It keeps only correctness-passing gains and finishes with the fastest reproducible configuration. Here’s the prompt we added to `AGENTS.md`. Muse glimmer need the Hugging Face MCP and the Hugging Face CLI and Skills.

## Self-optimization prompt

```
You are Muse Glimmer acting as an autonomous inference-optimization engineer for your own serving stack.
Goal: maximize valid single-H100 aggregate completion throughput in tokens/second.
Protocol:
1. Establish a correctness-passing baseline.
2. Test one reversible optimization at a time.
3. Keep the prompt, concurrency, sampling, request count, warm-up, and decode length fixed.
4. Reject results that fail correctness or prefix checks.
5. Record every experiment chronologically with its configuration, raw throughput, correctness, and delta.
6. Keep improvements and revert regressions.
7. Stop after six consecutive regressions or when the experiment budget is exhausted.
8. Report the best valid configuration and exact reproduction command.
Create a minimal scientific animation of the results:
- white background;
- raw tokens/second—never normalize;
- one point revealed per experiment;
- connect every point chronologically;
- begin with the lowest valid result;
- stop at the best result;
- export as a GIF.
Never fabricate, interpolate, or count correctness-failing measurements.
```
Try Muse Glimmer as a Hugging Face research agent. The Gradio Space sends each model request to a private Hugging Face Inference Endpoint through its OpenAI-compatible API. It also connects to the official Hugging Face MCP server, giving the agent read-only tools to search and inspect Hub repositories, models, datasets, Spaces, documentation, and papers.

We are happy to welcome Muse Glimmer to the Hugging Face Hub. Try Muse Glimmer with your local coding setups today!
