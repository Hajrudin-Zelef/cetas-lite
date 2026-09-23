---
id: collect-tutoriels/tutoriels/qwen-38-27b
title: "Qwen 3.8 27B is excellent, but it defaults to wildly overthinking things"
domain: tutoriels
role: reference
task: review
actors: ["Alibaba", "Nvidia", "OpenAI", "OpenRouter"]
dates: ["2026-08", "2026-09-23"]
keywords: ["qwen", "agent", "agents", "apache", "benchmark", "benchmarks", "consumer", "datacenter", "gguf", "gpt-5.6", "license", "llama"]
source: docs/RAG/Collect RAG/07_tutoriels/qwen-38-27b.md
source_anchor: ""
source_lines: [1, 90]
sha256: bd6fbf332746ca2565aacf44731e9283fac0f05bc9124d77bca6b5a601bd27a0
---

# Qwen 3.8 27B is excellent, but it defaults to wildly overthinking things

## Metadata

- **Source** : https://simonwillison.net/2026/Aug/16/qwen-38-27b/
- **Site** : Simon Willison
- **Type** : Review
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Simon Willison's blog post (16 August 2026) reviews Qwen 3.8 27B, the Apache 2 licensed 27B-parameter vision-capable LLM released by Alibaba's Qwen research lab on Friday 14 August 2026. He had been looking forward to it because 27B is an excellent size for running on a reasonably specced laptop, and its predecessor Qwen 3.6 27B was impressive. Qwen's self-reported benchmarks show a boost from both Qwen 3.6 27B and the closed-weight Qwen 3.7-Plus. He notes it will be interesting to hear what independent benchmarks say.

He ran the model on two machines: a 128GB M5 Max MacBook Pro and an NVIDIA DGX Spark, using LM Studio with the 17GB Q4_K_M quantized build, and also tried `llama-server` directly on the Spark. His central finding is that the default reasoning effort of "extra high" (xhigh) results in spectacular over-thinking. Qwen's documentation describes the model as defaulting to xhigh for `reasoning_effort`, with medium and low alternatives. He calls this "a hilarious default" that is absolutely not a good way to run the model, especially on consumer hardware. He hit LM Studio's default 8,192-token context limit because Qwen used it all up thinking about mundane problems; loading the full 262,144 maximum context fixed it.

His first pelican-riding-a-bicycle SVG took 21 minutes to generate, using 22,276 reasoning tokens to produce 3,223 output tokens. He calls it by far the best pelican SVG he has generated with a local model — correct bicycle frame shape, legs on each side (very rare), clear pouch, wings touching the handlebars, motion lines behind, tasteful background — but asks "was that worth waiting 21 minutes for? Absolutely not." With reasoning off, the same prompt produced 3,715 tokens in 137s (just over two minutes). He also ran the prompt through the much larger Qwen 3.8 2.4T-A95B on OpenRouter, getting an animated SVG. A simpler prompt ("draw an svg of a circle") with xhigh still produced an elaborate animated "circle study" entirely not what was asked.

He tested bounding boxes: asking for JSON bounding boxes around pelicans in a photo on a 0-1000 scale produced an excellent match, which he rendered using a custom HTML tool that Qwen 3.8 27B built for him from a single prompt (massively over-engineered because he forgot to dial down thinking). Notably, the model added an unrequested "demo scene" feature and drew its own pelicans because he used the label "pelicans" in the example JSON. Without reasoning the tool "nearly works" but shows boxes in the wrong place, illustrating where reasoning matters.

He also found the model can drive coding agents: he configured Pi to use Qwen 3.8 27B in LM Studio on the Spark (shared via tailscale serve), asked "how does auth work?" in his datasette folder, and got a very solid reply. He then had it write Python code to convert a JSONL transcript to markdown, which it built and tested (`pi_jsonl_to_md.py`).

The catch is speed: 15-30 tokens/s from LM Studio, versus OpenAI 5.6 Sol at 74 tokens/s and 5.6 Luna at 184 tokens/s per Artificial Analysis. He explored optimizations, including Multi-Token Prediction (MTP). Following a tweet from llama.cpp creator Georgi Gerganov, he ran the model with `--spec-default --spec-type draft-mtp --reasoning-preserve`, which outperformed the LM Studio default GGUF by around 72% in a benchmark run via GPT-5.6 in Codex. He concludes that a 17GB file doing all this on home machines is a miracle, and that the most important thing Qwen 3.8 27B demonstrates is that we can have an open-weights general purpose model with long context, effective tool calling, strong vision, and competent code generation in just a 17GB file — no half-million-dollar datacenter needed.

## Key points

- Qwen 3.8 27B: Apache 2, 27B vision-capable LLM from Alibaba's Qwen lab, released 14 August 2026.
- Defaults to `reasoning_effort=xhigh`, causing spectacular over-thinking; recommendation is to run on low or no reasoning at first.
- 21-minute pelican SVG used 22,276 reasoning tokens for 3,223 output tokens; reasoning off took 137s.
- LM Studio default 8,192-token context is insufficient; full 262,144 context fixed the problem.
- Excellent at vision bounding boxes (0-1000 scale JSON) and built a full bbox-labeling HTML tool from one prompt.
- Can drive coding agents (Pi harness) for repo questions and code generation.
- Speed: 15-30 tokens/s from LM Studio, slow for a daily driver; MTP spec decoding gave ~72% boost.
- Run on 128GB M5 Max MacBook Pro and NVIDIA DGX Spark with the 17GB Q4_K_M quant.

## Technical data / figures

| Item | Value |
| --- | --- |
| Model | Qwen 3.8 27B |
| License | Apache 2 |
| Parameters | 27B (vision-capable) |
| Quant used | Q4_K_M, ~17 GB |
| Default reasoning_effort | xhigh |
| Max context | 262,144 tokens |
| Pelican SVG (xhigh) | 21 min, 22,276 reasoning + 3,223 output tokens |
| Pelican SVG (no reasoning) | 137s, 3,715 tokens |
| Local speed | 15-30 tokens/s |
| OpenAI 5.6 Sol | 74 tokens/s (Artificial Analysis) |
| OpenAI 5.6 Luna | 184 tokens/s (Artificial Analysis) |
| MTP speedup | ~72% over LM Studio default GGUF |

Hardware/setup: 128GB M5 Max MacBook Pro; NVIDIA DGX Spark; LM Studio; llama-server.

llama.cpp MTP command:
```
llama serve \
 -hf ggml-org/Qwen3.8-27B-GGUF:Q4_K_M \
 -hfd ggml-org/Qwen3.8-27B-GGUF:Q4_0 \
 --spec-default \
 --spec-type draft-mtp \
 --reasoning-preserve
```

Pi config snippet (`~/.pi/agent/models.json`):
```json
{
  "providers": {
    "spark": {
      "baseUrl": "https://spark-18b3.tail68a31.ts.net/v1",
      "api": "openai-responses",
      "apiKey": "dummy",
      "models": [{ "id": "qwen3.8-27b", "reasoning": true }]
    }
  }
}
```

Bounding box example output:
```json
[
  {"bbox_2d": [195, 290, 370, 780], "label": "pelicans"},
  {"bbox_2d": [445, 320, 675, 850], "label": "pelicans"}
]
```

## Why this source matters for the RAG

It is a hands-on, independent practitioner review with reproducible commands, token counts, timings, and hardware details for running a major open-weight model locally. It documents the critical `reasoning_effort` default pitfall and a concrete MTP speedup, making it valuable for RAG queries about local LLM deployment, reasoning controls, and real-world model performance.
