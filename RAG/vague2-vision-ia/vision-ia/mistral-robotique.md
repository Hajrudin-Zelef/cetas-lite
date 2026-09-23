---
id: vague2-vision-ia/vision-ia/mistral-robotique
title: "Mistral débarque dans la robotique, et une seule caméra suffit à son robot"
domain: vision-ia
role: reference
task: article
actors: ["AWS", "Anthropic", "China", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Nvidia", "OpenAI", "United States", "xAI"]
dates: ["2026-07-09", "2026-09-23"]
keywords: ["mistral", "agent", "agentic", "apache", "benchmark", "benchmarks", "chatgpt", "consumer", "cost", "fable 5", "gpt-live", "gpus"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/mistral-robotique.md
source_anchor: ""
source_lines: [1, 52]
sha256: fbd84314c41b3c105bc74ef59ed56bb63cb8f003bb2d0f23e43b8e436641e6a0
---

# Mistral débarque dans la robotique, et une seule caméra suffit à son robot

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/mistral-robotique
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This July 9, 2026 edition of the Vision-IA newsletter is led by Mistral's official entry into robotics with Robostral Navigate, an 8-billion-parameter model capable of guiding a robot through an unknown environment using a single RGB camera—no lidar or expensive sensor stack—following natural-language instructions. The model was trained entirely in simulation then fine-tuned with reinforcement learning (CISPO method), and achieves 76.6% success on the R2R-CE benchmark, which measures following verbal instructions in a continuous environment. No availability date has been announced. The newsletter frames this as attacking the real barrier to consumer robotics: hardware cost. If a robot can navigate with a few-euro webcam instead of a multi-thousand-dollar lidar, domestic robotics becomes far more credible—and a European player is setting this milestone on terrain dominated by US and Chinese labs.

The edition also covers Yann LeCun calling LLMs a dead end: in a Bloomberg interview, the former Meta AI research head argues language is only a very approximate, reduced description of the world, that large models are pre-trained on ~30,000 billion tokens (~10^14 bytes, nearly all public internet text)—exactly the amount of information a 4-year-old absorbs through vision alone in four years. Without vision and multimodal learning, he says, there is no true general intelligence.

OpenAI revamped ChatGPT's voice mode with GPT-Live-1, designed for more natural conversation: it interrupts less, respects pauses and hesitations, and routes queries to the best text models (like GPT-5.5) for reasoning or web search. xAI launched Grok 4.5, which Musk calls "Opus class," priced at $2 per million input tokens, consuming 4.2x fewer tokens than Opus 4.8 for the same task, trained on tens of thousands of Nvidia GB300 GPUs; it trails Fable 5 and GPT-5.5 in coding but far cheaper, with Europe availability expected mid-July.

The "Research" section covers MiniMax preparing a 2,700-billion-parameter open-source LLM; self-improving AI no longer limited to frontier labs (per Wired); General Intuition training robot foundation models on millions of hours of video games; Amazon's "Moonraker" project to make Alexa truly agentic; Google Photos adding Video Remix; OpenAI pointing out flaws in SWE-Bench Pro; and OpenAI clarifying its doctrine on national security partnerships.

Additional briefs include Meta testing "Super Sensing" Ray-Ban glasses that continuously record (privacy concerns), a pickup-artist guru's alleged chatbot affair, Microsoft Research's Flint language for AI-agent-generated charts (compiles to Vega-Lite, Apache ECharts, Chart.js; 20+ chart types; MCP integration), Anthropic's "Advisor" pattern where Fable 5 plans and delegates to Sonnet 5 (92% of Fable 5's performance at 63% of the cost), Google's SynthID deepfake detector exposing a fake McConnell image, OpenAI training K-12 teachers, ex-DeepMind's Verity Harding warning on an AI arms race, SambaNova's $11B valuation, Meta's first big Canada data center, Kevin Weil joining Stoke Space, and Blue Origin raising funds at a $130B valuation.

## Key points

- Mistral enters robotics with Robostral Navigate, an 8B model navigating unknown environments from a single RGB camera.
- Trained in simulation, fine-tuned via reinforcement learning (CISPO); 76.6% success on the R2R-CE benchmark.
- No availability date announced; the bet is that cheap hardware (webcam vs lidar) unlocks consumer robotics.
- Yann LeCun argues LLMs are a dead end; language captures only a fraction of reality, vision is essential for general intelligence.
- OpenAI's GPT-Live-1 voice mode respects pauses and routes to text models like GPT-5.5 for reasoning.
- xAI's Grok 4.5 at $2/million input tokens is "Opus class" and uses 4.2x fewer tokens than Opus 4.8.
- MiniMax is preparing a 2,700B-parameter open-source LLM.
- Anthropic's "Advisor" pattern (Fable 5 planning + Sonnet 5 executing) reaches 92% performance at 63% cost.

## Technical data / figures

| Item | Figure |
|---|---|
| Robostral Navigate | 8B parameters, single RGB camera |
| R2R-CE benchmark success | 76.6% |
| LLM pretraining data (LeCun) | ~30,000 billion tokens (~10^14 bytes) |
| Grok 4.5 price | $2 / million input tokens |
| Grok 4.5 token efficiency | 4.2x fewer than Opus 4.8 |
| Grok 4.5 training hardware | tens of thousands of Nvidia GB300 GPUs |
| MiniMax model size | 2,700 billion parameters |
| Anthropic Advisor pattern | 92% performance at 63% cost |
| SambaNova valuation | $11B |
| Blue Origin valuation | $130B |

## Why this source matters for the RAG

It documents Mistral's strategic expansion from language models into embodied AI/robotics with a cost-focused approach, alongside a high-profile expert critique of the LLM paradigm (LeCun) and the emergence of price-per-token competition (Grok 4.5). It provides concrete quantitative benchmarks useful for tracking the robotics-AI convergence and the performance/cost trade-offs central to the industry in 2026.
