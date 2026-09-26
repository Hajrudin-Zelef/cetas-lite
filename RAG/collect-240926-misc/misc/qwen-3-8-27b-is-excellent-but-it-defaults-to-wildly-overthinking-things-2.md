---
id: collect-240926-misc/misc/qwen-3-8-27b-is-excellent-but-it-defaults-to-wildly-overthinking-things-2
title: "qwen-3-8-27b-is-excellent-but-it-defaults-to-wildly-overthinking-things"
domain: simonwillison
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI"]
dates: ["2026-09"]
keywords: ["qwen", "agent", "agents", "astra", "benchmark", "chatgpt", "claude", "consumer", "datacenter", "gguf", "gpt-5.6", "gpt-6"]
source: docs/RAG/clean_en/misc/qwen-3-8-27b-is-excellent-but-it-defaults-to-wildly-overthinking-things.md
source_anchor: ""
source_lines: [115, 196]
sha256: c31ffcfe92c1f66cebd968bd331e15e0e5ef25b864cac5b601197f0f41df02c5
---

# qwen-3-8-27b-is-excellent-but-it-defaults-to-wildly-overthinking-things

So without reasoning it didn’t quite one-shot a working tool. I’m sure it could get there with some follow-up prompts, but this is a good example of how reasoning can make a difference.

#### Yes, it can drive coding agents

One of the biggest questions around local models is whether or not they have enough horsepower to successfully run a coding agent loop. Coding agents require long context, strong code generation support and reliable tool-calling. On paper Qwen 3.8 27B has all three of these, so is it up to the task?

My initial experiments with Pi have been very promising. I chose Pi because it has a shorter system prompt than most other options, making it a better fit for trying out smaller models.

I configured Pi to use Qwen 3.8 27B running in LM Studio on the Spark (shared via `tailscale serve`) by adding this to `~/.pi/agent/models.json`:

```
{
  "providers": {
    "spark": {
      "baseUrl": "https://spark-18b3.tail68a31.ts.net/v1",
      "api": "openai-responses",
      "apiKey": "dummy",
      "models": [
        {
          "id": "qwen3.8-27b",
          "reasoning": true
        }
      ]
    }
  }
}
```
Then ran `pi --provider spark --model qwen3.8-27b` in my `~/dev/datasette` folder and prompted:

`how does auth work?`


After a sequence of reasoning and tool calls that accessed a bunch of different files it produced this reply, which is very solid.

Just one problem: I wanted to share that transcript. So I pointed Pi and Qwen 3.8 27B at the JSONL transcript file in `~/.pi/agent/sessions/--Users-simon-Dropbox-dev-datasette--` and prompted:

`Write Python code to convert this jsonl to markdown`


And it built and tested this pi_jsonl_to_md.py, which did exactly what I needed. Here’s that session transcript, published using the tool that it created.

#### The quest for speed

So far this is all looking *very* promising. We have a 17GB model that runs on high-end consumer hardware and can write code, drive tools, annotate images and generally do everything that I need from an LLM for getting real work done.

There’s one very significant catch: it feels slow—especially when it starts over-thinking, but even without that it’s not particularly sprightly.

I’ve been getting around 15-30 tokens a second from LM Studio. That’s not terrible, but it’s slow enough that it’s going to be hard to win me away from hosted API models, which can return results a whole lot faster. Artificial Analysis track token speed and show OpenAI 5.6 Sol at 74 tokens/second and 5.6 Luna at an impressive 184/second.

The good news is that the community have been exploring ways to speed things up since the model was first released two days ago.

One of the most promising optimizations is baked into the model itself. Qwen supports Multi-Token Prediction, an architecture trick where a cheaper mechanism guesses several tokens ahead and the main model can then quickly verify if the guesses were correct. This can have quite a dramatic effect on inference performance.

Based on this tweet from `llama.cpp` creator Georgi Gerganov I tried running the model with MTP like this on the Spark:

```
llama serve \
 -hf  ggml-org/Qwen3.8-27B-GGUF:Q4_K_M \
 -hfd ggml-org/Qwen3.8-27B-GGUF:Q4_0 \
 --spec-default \
 --spec-type draft-mtp \
 --reasoning-preserve
```
And sure enough, this gave me a significant boost. I had GPT-5.6 in Codex run a comparative benchmark on the Spark and the `--spec-type draft-mtp` server outperformed the LM Studio default GGUF by around 72%.

I expect we’ll see a whole lot more innovation around serving this model faster over the next few weeks. The MLX community likely have some tricks brewing as well.

#### Some observations

The fact that a 17GB file can do all of this stuff on my home machines is a *miracle*. Once again, I’m delighted and amazed at how much progress local models have made this year. A year ago this would have been competitive with the best and most expensive of the proprietary models—today it can run on a capable laptop.

The only thing holding this back from being a daily driver is performance. It feels pretty slow on both the M5 Mac and the DGX Spark. That’s the catch with these dense (non-Mixture-of-Experts) models—they require a whole lot of memory bandwidth to perform well, and neither of the machines I have access to are top performers in that regard.

The most important thing about Qwen 3.8 27B is **what it demonstrates**. We can have an open weights general purpose model with a long context, effective tool calling, strong vision ability, and competent code generation, and we can fit the whole thing in just a 17GB file.

The models at this size continue to get better at an impressive rate. We don’t need to spend half a million dollars on datacenter-class hardware just to run a competent model.

## More recent articles

- Claude Opus 5.5, GPT-6 Sol, GPT-6 Luna, and a new price war - 22nd September 2026
- Jev introduces a new shape of LLM - System One, aka Decision Models - 21st September 2026
- Generating running routes with GPT-6 Astra and ChatGPT Work - 12th September 2026
