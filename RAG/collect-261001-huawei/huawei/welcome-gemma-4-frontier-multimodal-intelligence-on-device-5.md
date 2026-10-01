---
id: collect-261001-huawei/huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device-5
title: "Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ..."
domain: huawei
role: reference
task: reference
actors: ["Apple", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmarks", "compute", "decode", "diffusion", "fp8", "gguf", "inference"]
source: docs/RAG/collect-261001-huawei/welcome-gemma-4-frontier-multimodal-intelligence-on-device.md
source_anchor: ""
source_lines: [346, 510]
sha256: 2352835145f99b4ebba666e563587a2003bd62f561969aa46af450678b311d8e
---

# Cette image montre un personnage anthropomorphe ressemblant à un lapin, vêtu d'un manteau bleu et d'un pantalon beige, se tenant sur un chemin de terre dans un paysage rural idyllique ...

```
from transformers import AutoModelForMultimodalLM, AutoProcessor
model = AutoModelForMultimodalLM.from_pretrained("google/gemma-4-E2B-it", device_map="auto")
processor = AutoProcessor.from_pretrained("google/gemma-4-E2B-it")
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "video",
                "image": "https://huggingface.co/datasets/merve/vlm_test_images/resolve/main/rockets.mp4",
            },
            {"type": "text", "text": "What is happening in this video?"},
        ],
    }
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    add_generation_prompt=True,
    return_dict=True,
    return_tensors="pt"
).to(model.device)
generated_ids = model.generate(**inputs, max_new_tokens=128)
generated_ids_trimmed = [
    out_ids[len(in_ids) :] for in_ids, out_ids in zip(inputs.input_ids, generated_ids)
]
output_text = processor.batch_decode(
    generated_ids_trimmed, skip_special_tokens=True, clean_up_tokenization_spaces=False
)
print(output_text)
```
Gemma 4 models come with image+text support in llama.cpp from the get-go! This unlocks using Gemma 4 with all of your favorite local apps: llama-cpp server, lmstudio, Jan as well as coding agents like Pi across many backends such as Metal and CUDA.

You can install llama-cpp as follows.

```
curl -LsSf https://llama.app/install.sh | sh
```
You can then start a server compatible with the OpenAI API Replace the quantization scheme at the end of the command with the precision of your choice.

```
llama serve -hf ggml-org/gemma-4-E2B-it-GGUF
```
Check out this link for more options on combining llama.cpp with different coding agents and local apps. Find all the GGUF checkpoints in this collection.

We worked on making sure the new models work locally with agents like **openclaw, hermes, pi, and open code**. All thanks to llama.cpp! Run the following to try Gemma 4 right away.

First, start your local server:

```
llama serve -hf ggml-org/gemma-4-26b-a4b-it-GGUF:Q4_K_M
```
For **hermes:**

```
hermes model
```
For **openclaw:**

```
openclaw onboard
```
For **pi** define a `~/.pi/agent/models.json`:

```
{
  "providers": {
    "llama-cpp": {
      "baseUrl": "http://localhost:8080/v1",	
      "api": "openai-completions",
      "apiKey": "none",
      "models": [
        {
          "id": "ggml-org-gemma-4-26b-4b-gguf"
        }
      ]
    }
  }
}
```
For **open code** define a `~/.config/opencode/opencode.json`:

```
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "llama.cpp": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "llama server (local)",
      "options": {
        "baseURL": "http://127.0.0.1:8080/v1"
      },
      "models": {
        "gemma-4-26b-4b-it": {
          "name": "Gemma 4 (local)",
          "limit": {
            "context": 128000,
            "output": 8192
          }
        }
      }
    }
  }
}
```
transformers.js enables running Gemma 4 right inside browser. You can check out the model card to see text-only, image & text, audio & text inference in detail here. We also shipped a demo for you to test the model here.

Full multimodal support of Gemma 4 is available using the open-source `mlx-vlm` library. Here's how to ask the model to describe an image:

```
pip install -U mlx-vlm
```
```
mlx_vlm.generate \
--model google/gemma-4-E4B-it \
--image https://huggingface.co/datasets/huggingface/documentation-images/resolve/0052a70beed5bf71b92610a43a52df6d286cd5f3/diffusers/rabbit.jpg \
--prompt "Describe this image in detail"
```
mlx-vlm supports TurboQuant, which delivers the same accuracy as the uncompressed baseline while using ~4x less active memory and running a lot faster end-to-end. This makes long-context inference practical on Apple Silicon without sacrificing quality. Use it like this:

```
mlx_vlm.generate \
--model "mlx-community/gemma-4-26b-a4b-it-4bit" \
--prompt "Your prompt here" \
--kv-bits 3.5 \
--kv-quant-scheme turboquant
```
For audio examples and more details, please check the MLX collection.

mistral.rs is a Rust-native inference engine with day-0 Gemma 4 support across all modalities (text, image, video, audio) and builtin tool-calling and agentic functionality. Install mistral.rs:

```
curl --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/EricLBuehler/mistral.rs/master/install.sh | sh # Linux/macOS
irm https://raw.githubusercontent.com/EricLBuehler/mistral.rs/master/install.ps1 | iex # Windows
```
You can then start an OpenAI-compatible HTTP server:

```
mistralrs serve mistralrs-community/gemma-4-E4B-it-UQFF --from-uqff 8
```
Or, use interactive mode:

```
mistralrs run -m google/gemma-4-E4B-it --isq 8 --image image.png -i "Describe this image in detail."
mistralrs run -m google/gemma-4-E4B-it --isq 8 --audio audio.mp3 -i "Transcribe this fully."
```
Find all models here. Please, follow the instructions in the model cards for installation and inference guidelines.

Google has released Multi-Token Prediction (MTP) drafters for the Gemma 4 family: small *assistant* models that accelerate inference via speculative decoding. The drafter proposes several future tokens at once, and the target model verifies them in a single forward pass. You get the same outputs as the target model, just faster — no quality loss, no changes to reasoning behaviour. Reported end-to-end speedups go up to ~3x depending on hardware, batch size, and workload.

Assistants are available for all four Gemma 4 sizes (E2B, E4B, 26B A4B, 31B). They share the KV cache with the target model to avoid recomputing context, and the smaller edge variants additionally use an embedder clustering trick to keep memory and compute low on-device.

Find the checkpoints in the Gemma 4 collection and the mlx-community collection.

Alongside the autoregressive Gemma 4 family, Google DeepMind is releasing **DiffusionGemma**, a multimodal model that generates text using **discrete diffusion** instead of token-by-token autoregression. It's built on the same 26B A4B Mixture-of-Experts foundation (25.2B total / 3.8B active parameters, 8 active experts out of 128 plus 1 shared), takes text and image inputs, generates text, and supports up to 256K context — all under the same Apache 2.0 license.

Where a standard causal LM emits one token at a time, DiffusionGemma denoises whole blocks of tokens in parallel. The architecture is encoder-decoder: an autoregressive encoder prefills the prompt and builds the KV cache, while a decoder applies **bidirectional attention** over a "canvas" of 256 tokens. During **multi-canvas sampling**, the model iteratively denoises a full canvas with a diffusion sampler; once a canvas is finalized it's encoded and appended to the KV cache, then the next canvas begins. This block-autoregressive approach increases generation speed.

The headline benefit is **throughput**: parallel denoising generates roughly 15–20 tokens per forward pass, reaching per-user generation speeds **exceeding 1100 tokens/second** at low batch sizes (H100, FP8). Inference compute is adaptive too — simpler prompts and structured tasks like code need fewer denoising steps, so tokens-per-second scales with task complexity. It keeps the broader Gemma 4 toolkit: thinking mode, function calling, long context, native system prompts, and image understanding (OCR, document parsing, object detection, pointing) at variable aspect ratios and resolutions.

Benchmarks show an expected trade-off between speed and evaluation metrics — DiffusionGemma trails the autoregressive 26B A4B on most tasks (e.g. MMLU Pro 77.6% vs 82.6%, AIME 2026 69.1% vs 88.3%, GPQA Diamond 73.2% vs 82.3%) in exchange for its large speed advantage, while edging ahead on a few (HLE no tools 11.0% vs 8.7%).

Getting started looks just like the rest of Gemma 4, via the dedicated diffusion class:

